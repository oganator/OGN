package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (e *Entity) AzureSim(tempentitydata *EntityData) {
	e.Metrics = Metrics{}
	tempentitydata.SampleForEntity(e)
	e.MC = true

	buf := new(bytes.Buffer)
	// json.NewEncoder(buf).Encode(e)
	fmt.Println(json.NewEncoder(buf).Encode(e.Metrics))
	// req, _ := http.NewRequest("POST", "http://localhost:8081/test", buf)

	// req, _ := http.NewRequest("GET", "https://oganica.azurewebsites.net", nil)

	// client := &http.Client{}
	// res, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer res.Body.Close()

	// fmt.Println("response Status:", res.Status)

	// // 	Print the body to the stdout
	// io.Copy(os.Stdout, res.Body)

	// e.UpdateEntity(true, tempentitydata)
}
