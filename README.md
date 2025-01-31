# JSON GENERATOR - jgen

A dynamic test data generator based on JSON templates with placeholder support.

## Features

- Generate complex nested data structures
- Dynamic placeholders for various data types
- Type conversion (`:to_int`, `:to_float`, `:to_bool`)
- Enum support
- Flexible date/time formatting

## Installation

```bash
curl -fsSL https://raw.githubusercontent.com/savage-demon/jgen/master/install.sh | bash
```

## Usage

```bash
jgen input.json [-o output.json]
```

## Supported Placeholders

| Placeholder | Parameters                | Example                                            |
| ----------- | ------------------------- | -------------------------------------------------- |
| `{index}`   | Padding length (optional) | `{index:3}` → "001"                                |
| `{uuid}`    | -                         | "91ca6d4e-5376-4a09-8576-967381e234fe"             |
| `{word}`    | -                         | "minima"                                           |
| `{fname}`   | -                         | "John"                                             |
| `{lname}`   | -                         | "Doe"                                              |
| `{name}`    | -                         | "Mr. John Doe"                                     |
| `{address}` | -                         | "122 Coral Drive"                                  |
| `{phone}`   | -                         | "104-259-3681"                                     |
| `{email}`   | -                         | "hZsWIas@jLgKVxW.top"                              |
| `{char}`    | Length (optional)         | `{char:15}` → "JCUpmMeTa1IK7P4"                    |
| `{achar}`   | Length (optional)         | `{achar:15}` → "rpilyErHLNwdPtm"                   |
| `{ichar}`   | Length (optional)         | `{ichar:15}` → "594171701512527"                   |
| `{int}`     | Range (optional)          | `{int:0,100}` → "47"                               |
| `{float}`   | Range (optional)          | `{float:0,10}` → "3.34"                            |
| `{bool}`    | -                         | `{bool}` → "true"                                  |
| `{date}`    | Format                    | `{date:%d.%m.%Y %H:%M:%S}` → "31.01.2025 19:54:12" |
| `{rdate}`   | -                         | `{rdate}` → "2025-09-01 05:19:49"                  |

### Example Input (input.json)

```json
[
  {
    "Description": "This is the basic usage of the generator",
    "index": "{index}",
    "index lpad": "{index:5}",
    "current date": "{date}",
    "current date with mask": "{date:%d.%m.%Y %H:%M:%S}",
    "random date": "{rdate}",
    "random date with mask": "{rdate:%d.%m.%Y %H:%M:%S}",
    "random date with mask from to": "{rdate:%d.%m.%Y %H:%M:%S,2025-01-01 00:00:00,2025-01-31 23:59:59}",
    "uuid": "{uuid}",
    "word": "{word}",
    "first name": "{fname}",
    "last name": "{lname}",
    "full name": "{name}",
    "address": "{address}",
    "phone": "{phone}",
    "email": "{email}",
    "string of letters": "{achar:15}",
    "string of numbers": "{ichar:15}",
    "string of letters and numbers": "{char:15}",
    "int": "{int:0,100}",
    "int for real": "{int:0,100}:to_int",
    "float": "{float:0,10}",
    "float for real": "{float:0,10}:to_float",
    "bool": "{bool}",
    "bool for real": "{bool}:to_bool"
  },

  {
    "Generator with a '!count' flag": [
      {
        "!count": 5,
        "description": "next array element will be generated 5 times",
        "I will never appear in the output": "because of '!count' key"
      },
      {
        "number": "{index}",
        "I am being duplicated": "5 times"
      }
    ]
  },

  {
    "Description": "In case you want to generate some enumerations that shouldn't end up in the output file.",
    "Then use the exclude flag": [
      {
        "!exclude": 2
      },
      {
        "I will never appear in the output": "because of previous !exclude object"
      },
      {
        "I will never appear in the output": "because of previous !exclude object"
      },
      {
        "I'm 3rd after exclude 2, so I will appear": "in the output"
      }
    ]
  },

  { "!exclude": 1 },
  {
    "Description": "Enumerations should be described before they are used",
    "action": {
      "1": "talk!{enum:act}",
      "3": "play!{enum:act}",
      "5": "help!{enum:act}",
      "6": "laugh!{enum:act}"
    },
    "names": [
      { "!count": 10 },
      {
        "c": "{fname}!{enum:name}"
      }
    ]
  },
  {
    "Script": [
      { "!count": 5 },
      {
        "action_number": "{index}",
        "first_actor": "{oneof:name}",
        "second_actor": "{oneof:name}",
        "action": "{oneof:act}"
      }
    ],

    "fee for all actors": [
      { "!count": 10 },
      {
        "action": "{each:name} gets a fee of ${int:50,100} for '{oneof:act}ing'"
      }
    ]
  }
]
```

### Example Output (output.json)

```json
[
  {
    "Description": "This is the basic usage of the generator",
    "address": "163 Highwood Drive",
    "bool": "false",
    "bool for real": true,
    "current date": "2025-01-31 20:03:53",
    "current date with mask": "31.01.2025 20:03:53",
    "email": "SwMmtok@TXggJCo.biz",
    "first name": "Ahmed",
    "float": "5.59",
    "float for real": 9.43,
    "full name": "Lady Verona Hettinger",
    "index": "1",
    "index lpad": "00001",
    "int": "3",
    "int for real": 37,
    "last name": "Kilback",
    "phone": "110-826-4375",
    "random date": "2025-10-21 20:47:29",
    "random date with mask": "21.01.2026 03:43:15",
    "random date with mask from to": "13.01.2025 05:15:37",
    "string of letters": "oOKFlLKtgYhCNBw",
    "string of letters and numbers": "49WFe0Z3gNaHq1a",
    "string of numbers": "729795246287984",
    "uuid": "72589fef-66c1-4978-957d-7f711abd5c14",
    "word": "et"
  },
  {
    "Generator with a '!count' flag": [
      {
        "I am being duplicated": "5 times",
        "number": "1"
      },
      {
        "I am being duplicated": "5 times",
        "number": "2"
      },
      {
        "I am being duplicated": "5 times",
        "number": "3"
      },
      {
        "I am being duplicated": "5 times",
        "number": "4"
      },
      {
        "I am being duplicated": "5 times",
        "number": "5"
      }
    ]
  },
  {
    "Description": "In case you want to generate some enumerations that shouldn't end up in the output file.",
    "Then use the exclude flag": [
      {
        "I'm 3rd after exclude 2, so I will appear": "in the output"
      }
    ]
  },
  {
    "Script": [
      {
        "action": "help",
        "action_number": "1",
        "first_actor": "Kristopher",
        "second_actor": "Melba"
      },
      {
        "action": "play",
        "action_number": "2",
        "first_actor": "Celine",
        "second_actor": "Charles"
      },
      {
        "action": "laugh",
        "action_number": "3",
        "first_actor": "Celine",
        "second_actor": "Charles"
      },
      {
        "action": "laugh",
        "action_number": "4",
        "first_actor": "Ludwig",
        "second_actor": "Melba"
      },
      {
        "action": "play",
        "action_number": "5",
        "first_actor": "Kristopher",
        "second_actor": "Melba"
      }
    ],
    "fee for all actors": [
      {
        "action": "Britney gets a fee of $99 for 'talking'"
      },
      {
        "action": "Melba gets a fee of $58 for 'helping'"
      },
      {
        "action": "Richard gets a fee of $55 for 'playing'"
      },
      {
        "action": "Kristopher gets a fee of $66 for 'helping'"
      },
      {
        "action": "Ludwig gets a fee of $54 for 'playing'"
      },
      {
        "action": "Celine gets a fee of $86 for 'laughing'"
      },
      {
        "action": "Khalil gets a fee of $55 for 'playing'"
      },
      {
        "action": "Taylor gets a fee of $77 for 'helping'"
      },
      {
        "action": "Charles gets a fee of $78 for 'laughing'"
      },
      {
        "action": "Libbie gets a fee of $80 for 'helping'"
      }
    ]
  }
]
```
