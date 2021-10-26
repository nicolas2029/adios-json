package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func adios(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Adios api"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/api/v1/adios", adios)
	fmt.Println("Server iniciado en el puerto 3001")
	http.ListenAndServe(":3001", nil)
}
