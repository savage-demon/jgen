package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/google/uuid"
	"github.com/itchyny/timefmt-go"
	"github.com/labstack/gommon/log"
)

type Template interface{}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var enums map[string][]string

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generateData(template Template, parentIndex int) interface{} {
	switch v := template.(type) {
	case []interface{}:
		var result []interface{}
		lenv := len(v)
		for i := 0; i < lenv; i++ {
			item := v[i]
			if itemMap, ok := item.(map[string]interface{}); ok {
				if count, ok := itemMap["count"].(float64); ok {
					var childTemplate Template
					if i+1 < len(v) {
						childTemplate = v[i+1]
					} else {
						childTemplate = map[string]interface{}{}
					}
					for j := 0; j < int(count); j++ {
						result = append(result, generateData(childTemplate, j+1))
					}

					i++
				} else {
					result = append(result, generateData(item, parentIndex))
				}
			}
		}
		return result

	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			result[key] = generateData(value, parentIndex)
		}
		return result

	case string:
		resultString := replaceInString(v, parentIndex)
		resultString = strings.ReplaceAll(resultString, "\n", "")
		if strings.HasSuffix(resultString, ":to_int") {
			intValue, _ := strconv.Atoi(resultString[:len(resultString)-len(":to_int")])
			return intValue
		}
		if strings.HasSuffix(resultString, ":to_float") {
			floatValue, _ := strconv.ParseFloat(resultString[:len(resultString)-len(":to_float")], 64)
			return floatValue
		}
		if strings.HasSuffix(resultString, ":to_bool") {
			boolValue, _ := strconv.ParseBool(resultString[:len(resultString)-len(":to_bool")])
			return boolValue
		}
		return resultString

	default:
		return template
	}
}

func randomDate(start, end time.Time) time.Time {
	delta := end.Sub(start).Seconds()
	randomSeconds := rand.Int63n(int64(delta))
	return start.Add(time.Duration(randomSeconds) * time.Second)
}

func parsePlaceholder(placeholder string, index int) string {
	parts := strings.SplitN(placeholder, ":", 2)
	phtype := parts[0]
	params := []string{}
	if len(parts) > 1 {
		params = strings.Split(parts[1], ",")
	}

	switch phtype {
	case "index":
		lpad := 1
		if len(params) > 0 {
			lpad, _ = strconv.Atoi(params[0])
		}
		return fmt.Sprintf("%0*d", lpad, index)

	case "date":
		format := "%Y-%m-%d %H:%M:%S"
		if len(params) > 0 {
			format = strings.Join(params, ",")
		}
		return timefmt.Format(time.Now(), format)

	case "rdate":
		format := "%Y-%m-%d %H:%M:%S"

		var start, end time.Time
		var err error

		if len(params) > 0 {
			format = params[0]
		}

		if len(params) > 1 {
			start, err = time.Parse(time.DateTime, params[1])
			if err != nil {
				log.Errorf("Error parsing start date: %v", err)
			}
		} else {
			start = time.Now()
			end = time.Now().Add(365 * 24 * time.Hour)
		}

		if len(params) > 2 {
			end, err = time.Parse(time.DateTime, params[2])
			if err != nil {
				log.Errorf("Error parsing end date: %v", err)
			}
		}

		return timefmt.Format(randomDate(start, end), format)

	case "uuid":
		return uuid.New().String()

	case "email":
		return faker.Email()

	case "phone":
		return faker.Phonenumber()

	case "name":
		return faker.Name()

	case "fname":
		return faker.FirstName()

	case "lname":
		return faker.LastName()

	case "word":
		return faker.Word(func(oo *options.Options) {
			oo.RandomMaxSliceSize = 1
			oo.RandomMinSliceSize = 1
		})

	case "int":
		minVal := 0
		maxVal := 1000
		if len(params) > 0 {
			minVal, _ = strconv.Atoi(params[0])
		}
		if len(params) > 1 {
			maxVal, _ = strconv.Atoi(params[1])
		}
		return strconv.Itoa(rand.Intn(maxVal-minVal) + minVal)

	case "float":
		minVal := 0.0
		maxVal := 1.0
		if len(params) > 0 {
			minVal, _ = strconv.ParseFloat(params[0], 64)
		}
		if len(params) > 1 {
			maxVal, _ = strconv.ParseFloat(params[1], 64)
		}
		return fmt.Sprintf("%.2f", minVal+rand.Float64()*(maxVal-minVal))

	case "address":
		return faker.GetRealAddress().Address

	case "char":
		length := 1
		if len(params) > 0 {
			length, _ = strconv.Atoi(params[0])
		}
		return RandStringBytes(length)

	case "ichar":
		length := 1
		if len(params) > 0 {
			length, _ = strconv.Atoi(params[0])
		}
		min := intPow(10, length)
		max := intPow(10, length+1)
		return strconv.Itoa(rand.Intn(max-min) + min)

	case "bool":
		return strconv.FormatBool(rand.Intn(2) == 1)

	case "oneof":
		en, ok := enums[params[0]]
		if !ok {
			fmt.Printf("Enum %s not found\n", params[0])
			return placeholder
		}

		if len(en) == 0 {
			fmt.Printf("Enum %s not found\n", params[0])
			return placeholder
		}

		idx := rand.Intn(len(en))
		return en[idx]

	case "each":
		en, ok := enums[params[0]]
		if !ok {
			fmt.Printf("Enum %s not found\n", params[0])
			return placeholder
		}

		if len(en) == 0 {
			fmt.Printf("Enum %s not found\n", params[0])
			return placeholder
		}

		if len(en) < index {
			fmt.Printf("Not enough values in enum '%s' for each element\n", params[0])
			return placeholder
		}

		return en[index-1]

	default:
		return placeholder
	}
}

func replaceInString(s string, index int) string {
	var result strings.Builder

	split := strings.Split(s, "!")
	s = split[0]

	i := 0
	for i < len(s) {
		if s[i] == '{' && strings.Contains(s[i:], "}") {
			end := strings.Index(s[i:], "}") + i
			placeholder := s[i+1 : end]
			result.WriteString(parsePlaceholder(placeholder, index))
			i = end + 1
		} else {
			result.WriteByte(s[i])
			i++
		}
	}

	resultstr := result.String()

	if len(split) > 1 {
		param := split[1]

		paramt := strings.TrimPrefix(param, "{")
		paramt = strings.TrimSuffix(paramt, "}")

		params := strings.Split(paramt, ":")
		if len(params) != 2 {
			return result.String()
		}

		left := params[0]
		right := params[1]

		switch left {
		case "enum":
			enums[right] = append(enums[right], resultstr)
		}
	}

	return resultstr
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jgen <input_schema_file> [-o output_file]")
		return
	}

	inputFile := os.Args[1]
	outputFile := "output.json"

	if len(os.Args) > 3 && os.Args[2] == "-o" {
		outputFile = os.Args[3]
	}

	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	var schema interface{}
	if err = json.Unmarshal(file, &schema); err != nil {
		jsonerr := &json.SyntaxError{}
		if errors.As(err, &jsonerr) {
			fmt.Printf("Error unmarshalling JSON: %s at %d\n", jsonerr.Error(), jsonerr.Offset)
			return
		}
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	enums = make(map[string][]string)

	result := generateData(schema, 1)

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	if err := os.WriteFile(outputFile, output, 0o644); err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	fmt.Println("Data generated and saved to", outputFile)
}
