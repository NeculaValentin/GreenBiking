package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, getErr := http.Get("https://carburanti.mise.gov.it/ospzApi/registry/servicearea/31451")
	if getErr != nil {
		log.Fatalln(getErr)
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatalln(readErr)
	}

	var s StazioneServizio
	marshalErr := json.Unmarshal(body, &s)
	if marshalErr != nil {
		log.Fatalln(marshalErr)
	}

	for _, element := range s.Fuels {
		if element.FuelID == DIESEL {
			log.Println("Diesel price: ", element.Price)
		}
	}
}
