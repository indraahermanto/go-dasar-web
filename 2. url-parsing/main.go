package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlStr = "http://kalipare.com:80/hello?name=john wick&age=27"
	var u, e = url.Parse(urlStr)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlStr)
	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	// var query = u.Query() // tipe data map
	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)
}
