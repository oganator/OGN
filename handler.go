package main

import (
	_ "OGN/routers"
	"sync"

	ogn "OGN/controllers"
	"encoding/json"
	"fmt"
	"net/http"
)

func azureHandler(w http.ResponseWriter, r *http.Request) {
	tempentity := ogn.EntityModel{}
	err := json.NewDecoder(r.Body).Decode(&tempentity)
	if err != nil {
		fmt.Println("Azure Function Error: NewDecoder/Decode - ", err)
	}
	tempentity.MCSetup.Sims = tempentity.MCSetup.Sims / 100
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(t *ogn.EntityModel) {
		defer wg.Done()
		temp := make(map[int]*ogn.Unit)
		for i, v := range t.ChildUnitsMC {
			v.Parent = t
			temp[i] = &v
			// temp[i].Mutex = &sync.Mutex{}
		}
		t.ChildUnits = temp
		// tempentity.UpdateEntity(false, &tempentity.EntityData, "Azure")
		// fmt.Println("Azure Handler: ", tempentity.GrowthInput)
		// ogn.StructPrint("azureHandler - pre Monte Carlo: ", tempentity)
		// Entity Setup
		t.Mutex = &sync.Mutex{}

		t.MonteCarlo("Azure")
	}(&tempentity)
	response, err := json.Marshal(tempentity.MCResultSlice)
	if err != nil {
		fmt.Println("Azure Function Error: Marshal:", err)
		response, _ = json.Marshal(ogn.EntityModel{})
	}
	fmt.Println("Handler IRR: ", tempentity.Metrics.IRR.NetLeveredAfterTax)
	fmt.Fprint(w, response)
	fmt.Println("___________________________________________________________")
}

// func main() {
// 	listenAddr := ":8080"
// 	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
// 		listenAddr = ":" + val
// 	}
// 	http.HandleFunc("/api/OGNTrigger", azureHandler)
// 	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
// 	log.Fatal(http.ListenAndServe(listenAddr, nil))
// }
