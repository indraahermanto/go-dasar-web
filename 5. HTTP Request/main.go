package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseUrl = "http://localhost:8002"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// func fetchUser(ID string) (student, error) {
// 	var err error
// 	var client = &http.Client{}
// 	var data student

// 	var param = url.Values{}
// 	param.Set("id", ID)
// 	var payload = bytes.NewBufferString(param.Encode())
// 	fmt.Println(payload)

// 	request, err := http.NewRequest("POST", baseUrl+"/user", payload)
// 	if err != nil {
// 		return data, err
// 	}
// 	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	response, err := client.Do(request)
// 	if err != nil {
// 		return data, err
// 	}
// 	defer response.Body.Close()

// 	err = json.NewDecoder(response.Body).Decode(&data)
// 	if err != nil {
// 		return data, err
// 	}

// 	return data, nil
// }

func fetchUser(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseUrl+"/user", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	// /users
	// var users, err = fetchUsers()
	// if err != nil {
	// 	fmt.Println("error", err.Error())
	// 	return
	// }

	// for _, each := range users {
	// 	fmt.Println(each)
	// }

	// user/id
	var user1, err1 = fetchUser("E0001")
	if err1 != nil {
		fmt.Println("error", err1.Error())
		return
	}

	fmt.Println(user1)
	// error EOF
}
