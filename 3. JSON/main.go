package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name":"indra","Age":27}`
	var jsonData = []byte(jsonString)

	var data User
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(jsonString)
	fmt.Println(data)

	fmt.Println("user: ", data.Fullname)
	fmt.Println("age: ", data.Age)

	// decode json to map
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println(data1)

	// decode array json ke array object
	var jsonStr = `[
		{"Name":"jonathan","Age":32},
		{"Name":"dulah","Age":52}
	]`

	var data2 []User
	err = json.Unmarshal([]byte(jsonStr), &data2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data2)

	// encode object ke json string
	var object = []User{{"jonathan", 52}, {"dulah", 30}}
	var jsonData2, err2 = json.Marshal(object) // jsonData2 is byte
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	var jsonStr2 = string(jsonData2) // convert byte to string
	fmt.Println(jsonStr2)
}
