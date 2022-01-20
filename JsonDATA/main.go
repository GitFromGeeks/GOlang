package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Website  string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	println("Json in Golang Class")
	// EncodedJson()
	DecodeJson()
}

func EncodedJson() {
	courses := []course{
		{"Reactjs", 299, "learncodeonline.in", "123@123", []string{"facebook"}},
		{"Flutter", 799, "udemy", "123", nil},
		{"GoLang", 199, "learncodeonline", "123", []string{"google"}},
	}
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "GoLang",
        "Price": 199,
        "Website": "learncodeonline",
        "tags": ["google"]
	}
	`)

	var lcocourse course
	checkvlaid := json.Valid(jsonDataFromWeb)

	if checkvlaid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		print(" JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myonlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlineData)
	fmt.Printf("%#v\n", myonlineData)

	for k, v := range myonlineData {
		fmt.Printf("key is %v and value is %v and Type is %T\n", k, v, v)
	}

}
