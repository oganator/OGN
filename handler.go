package main

import (
	_ "OGN/routers"
	"log"
	"os"

	ogn "OGN/controllers"
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tempentity := ogn.Entity{}
	// slicebyte, err := r.URL.MarshalBinary()
	err := json.NewDecoder(r.Body).Decode(&tempentity)
	if err != nil {
		fmt.Println("Azure Function Error: NewDecoder/Decode", err)
	}
	tempentity.MCSetup.Sims = tempentity.MCSetup.Sims / 100
	tempentity.MonteCarlo("Azure")
	fmt.Println("Azure IRR: ", tempentity.MCResultSlice.IRR)
	response, err := json.Marshal(tempentity.MCResultSlice)
	if err != nil {
		fmt.Println("Azure Function Error: Marshal:", err)
	}
	fmt.Fprint(w, response)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/OGNTrigger", helloHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
