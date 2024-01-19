package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type MyData struct {
	BuildName string `json:"buildName"`
	// Include other fields if needed
}

func main() {

	url := "https://fibakod.fibabanka.com.tr/FibaCollection/Individual%20Banking%20and%20Technology/_apis/release/definitions?api-version=6.0&name=ipm-fx-market-frontend-CoE"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic Om41eXd1MjNldm82ZTR2dm5rYnlkanpkc2phcGxhbXVpbWpra3NyN2NlcjUzdTJ3azJzYmE=")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reflect.TypeOf(body))
	// convert this json string(body)
	//var datas map[string]interface{}
	jsonData := []byte(body)
	var data MyData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(jsonData)
	// Access the buildName field
	fmt.Println("buildName:", data.BuildName)
	// convert this json string(body) to map[string]interface{}
	//var datas map[string]interface{}
	//err = json.Unmarshal(body, &datas)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(reflect.TypeOf(datas))
	////fmt.Println(datas["value"][0]["id"])
	//var value = datas["value"]
	//fmt.Println(reflect.TypeOf(value))
	//for _, v := range value {
	//	fmt.Println(v)
	//}
}
