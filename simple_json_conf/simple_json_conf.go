package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func readConfFile() {
	data, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		fmt.Println("Unable to reasd conf file.", err)
		return
	}

	type Config struct {
		Timeout int
		Delete  bool
	}

	var configs Config
	err = json.Unmarshal(data, &configs)
	if err != nil {
		fmt.Println("JSON decode error.", err)
		return
	}

	fmt.Println(configs.Delete)
	fmt.Println(configs.Timeout)
}

func main() {
	carsColors := map[string]int{
		"Black": 10,
		"Blue":  5,
		"Red":   1,
	}

	bytes, _ := json.Marshal(carsColors)
	fmt.Println(string(bytes))

	readConfFile()

}
