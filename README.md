# JSON GENERATOR - jgen

---

A dynamic test data generator based on JSON templates with placeholder support.

## Features

- Generate complex nested data structures
- Dynamic placeholders for various data types
- Type conversion (`:to_int`, `:to_float`, `:to_bool`)
- Enum support
- Flexible date/time formatting

## Installation

wip...

## Usage

```bash
jgen input.json [-o output.json]
```

## Supported Placeholders

| Placeholder | Parameters                 | Example                                                 |
| ----------- | -------------------------- | ------------------------------------------------------- |
| `{index}`   | Padding length (optional)  | `{index:3}` → "001"                                     |
| `{uuid}`    | -                          | Generates UUID v4                                       |
| `{date}`    | Format (default: %Y-%m-%d) | `{date:%d/%m/%Y}` → "15/05/2023"                        |
| `{rdate}`   | Format, start, end         | `{rdate:%Y-%m,2020-01-01 15:20:00,2023-12-31 23:59:59}` |
| `{email}`   | -                          | `{email}` → "Fb3oq@example.com"                         |
| `{phone}`   | -                          | `{phone}` → "+7 (999) 999-99-99"                        |
| `{name}`    | -                          | `{name}` → "MR John Doe"                                |
| `{fname}`   | -                          | `{fname}` → "John"                                      |
| `{lname}`   | -                          | `{lname}` → "Doe"                                       |
| `{address}` | -                          | `{address}` → "123 Main St"                             |
| `{int}`     | min, max                   | `{int:1,100}` → "42"                                    |
| `{float}`   | min, max                   | `{float:0,1}` → "0.75"                                  |
| `{ichar}`   | Number length              | `{ichar:3}` → "456"                                     |
| `{oneof}`   | Enum name                  | `{oneof:statues}` → "quisquam3"                         |
| `{bool}`    | -                          | `{bool}` → "true"                                       |
| `{char}`    | Length                     | `{char:8}` → "6FFmxtWH"                                 |

### Example Input (input.json)

```json
{
  "statuslist": [
    { "count": 5 },
    {
      "status": "{word}{index}!{enum:statues}"
    }
  ],
  "products": [
    {
      "count": 2
    },
    {
      "id": "{uuid}",
      "index": "{index}",
      "name": "{word}-{word}-{int:100-999}",
      "sku": "SKU-{index:4}",
      "status": "{oneof:statues}",
      "dates": {
        "random": "{rdate}",
        "random with format": "{rdate:%Y-%m-%d %H:%M:%S}",
        "random with format and range": "{rdate:%Y-%m-%d %H:%M:%S,2025-01-01 00:00:00,2025-01-31 23:59:59}",
        "now with format": "{date:%Y-%m-%d %H:%M:%S}",
        "now": "{date}"
      },
      "names": {
        "full_name": "{name}",
        "first_name": "{fname}",
        "last_name": "{lname}"
      }
    }
  ]
}
```

### Example Output (output.json)

```json
{
  "products": [
    {
      "dates": {
        "now": "2025-01-24 08:45:51",
        "now with format": "2025-01-24 08:45:51",
        "random": "2025-05-05 23:19:42",
        "random with format": "2025-05-19 03:53:11",
        "random with format and range": "2025-01-11 10:10:05"
      },
      "id": "268692d0-6d30-4e89-a31f-fea04544364b",
      "index": "1",
      "name": "fugiat-unde-417",
      "names": {
        "first_name": "Cade",
        "full_name": "Dr. Allison Corkery",
        "last_name": "Keebler"
      },
      "sku": "SKU-0001",
      "status": "quisquam3"
    },
    {
      "dates": {
        "now": "2025-01-24 08:45:51",
        "now with format": "2025-01-24 08:45:51",
        "random": "2025-10-03 11:55:31",
        "random with format": "2025-02-25 05:27:18",
        "random with format and range": "2025-01-11 05:34:51"
      },
      "id": "d5995863-5613-4008-aab0-288b44e31c70",
      "index": "2",
      "name": "nisi-expedita-438",
      "names": {
        "first_name": "Dario",
        "full_name": "Mrs. Kasey Ritchie",
        "last_name": "Jakubowski"
      },
      "sku": "SKU-0002",
      "status": "nemo5"
    }
  ],
  "statuslist": [
    {
      "status": "odio1"
    },
    {
      "status": "id2"
    },
    {
      "status": "quisquam3"
    },
    {
      "status": "beatae4"
    },
    {
      "status": "nemo5"
    }
  ]
}
```
