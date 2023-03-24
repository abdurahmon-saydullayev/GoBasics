package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JSON struct {
	ID       int    `json:"id"`
	Code     string `json:"Code"`
	Ccy      string `json:"Ccy"`
	CcyNmRU  string `json:"CcyNm_RU"`
	CcyNmUZ  string `json:"CcyNm_UZ"`
	CcyNmUZC string `json:"CcyNm_UZC"`
	CcyNmEN  string `json:"CcyNm_EN"`
	Nominal  string `json:"Nominal"`
	Rate     string `json:"Rate"`
	Diff     string `json:"Diff"`
	Date     string `json:"Date"`
}

func main() {
    data:=ReadFileFromJSON()
	// fmt.Println(data)
	var str string
	fmt.Println("Kirit:")
	fmt.Scan(&str)
	for _,v:= range data{
		if v.Ccy==str{
			fmt.Println(v)
		}
	}
}

func ReadFileFromJSON()[]JSON{
	var datas [] JSON
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return datas
	}

	err = json.Unmarshal(data, &datas)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return datas
	}
	return datas
}

// func WriteFileToJSON(datas []JSON) {
// 	data, err := json.Marshal(datas)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON:", err)
// 		return
// 	}

// 	err = ioutil.WriteFile("data.json", data, 0644)
// 	if err != nil {
// 		fmt.Println("Error writing JSON to file:", err)
// 		return
// 	}
// }