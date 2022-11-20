package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// 参考资料： https://zetcode.com/golang/getpostrequest/
func main() {
	sendGet()
	sendPost()
}

func sendGet() {
	resp, err := http.Get("http://localhost:8088/health")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}

func sendPost() {
	values := map[string]string{"name": "John Doe", "occupation": "gardener"}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8088/post", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))

	// var res map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&res)

	// log.Println(res["json"])
}
