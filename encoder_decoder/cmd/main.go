package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	data := `
	{
	"data" : {
		"a" : 1,
		"b" : 2
	}
	}
	`
	var d interface{}
	var d0 map[string]interface{}
	var d1 map[string]map[string]interface{}
	var d2 map[string]map[string]float64
	_ = json.Unmarshal([]byte(data), &d)
	_ = json.Unmarshal([]byte(data), &d0)
	_ = json.Unmarshal([]byte(data), &d1)
	_ = json.Unmarshal([]byte(data), &d2)
	fmt.Println(d.(map[string]interface{})["data"].(map[string]interface{})["a"].(float64))
	fmt.Println(d0["data"].(map[string]interface{})["a"].(float64))
	fmt.Println(d1["data"]["a"].(float64))
	fmt.Println(d2["data"]["a"])

	// at one moment of time we can only deconstruct one interface

	// var x map[string]interface{}
	// err := json.Unmarshal([]byte(data), &x)
	// if err != nil {
	// }

	// fmt.Println(x)
	// var z map[string]interface{}

	// z, _ = x["data"].(map[string]interface{})
	// fmt.Printf("%T\n", z["a"])
	// switch val := z["a"].(type) {
	// case float64:
	// 	zz := int(val)
	// 	fmt.Printf("%T  and val is %v\n", zz, val)
	// 	break
	// default:
	// 	fmt.Println("NOOO")
	// }

	// fmt.Printf("%T\n", x["data"])

}

// https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/

// https://bitfieldconsulting.com/golang/map-string-interface
