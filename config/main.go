package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func main() {
	envs, err := godotenv.Read(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	envMap := map[string]string{}
	for key, value := range envs {
		envMap[key] = value
	}

	f, err := os.Create("config.json")
	_, err = f.Write(injectEnvs(config, envMap))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(config))
}

func injectEnvs(config []byte, env map[string]string) []byte {
	var data map[string]interface{}
	err := json.Unmarshal(config, &data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	re := regexp.MustCompile(`{{\$ENV\.(.*?)}}`)

	var processValue func(interface{}) interface{}
	processValue = func(value interface{}) interface{} {
		switch v := value.(type) {
		case map[string]interface{}:
			for k, innerValue := range v {
				v[k] = processValue(innerValue)
			}
			return v
		case []interface{}:
			for i, item := range v {
				v[i] = processValue(item)
			}
			return v
		case string:
			return re.ReplaceAllStringFunc(v, func(match string) string {
				key := re.FindStringSubmatch(match)[1]
				return env[key]
			})
		default:
			return value
		}
	}

	data = processValue(data).(map[string]interface{})

	// Converte o objeto json em bytes
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return bytes
}
