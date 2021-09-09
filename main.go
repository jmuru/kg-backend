package main

import (
	mux "github.com/gorilla/mux"
	service "github.com/kat-generator/KGB/service"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("PORT", "8080");
	port := os.Getenv("PORT")



	if port == "" {
		log.Fatal("$PORT must be set")
	}

	svc := service.NewService()

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", svc.HelloWorld)


	myRouter.HandleFunc("/accessory/{placement}", svc.GetAccessory).Methods("GET")
	myRouter.HandleFunc("/face", svc.GetFace).Methods("GET")
	myRouter.HandleFunc("/background", svc.GetBackground).Methods("GET")
	myRouter.HandleFunc("/generate/kat", svc.GetKat).Methods("GET")
	myRouter.HandleFunc("/palette/{type}", svc.GetPalette).Methods("GET")


	myRouter.HandleFunc("/accessory/create", svc.CreateAccessory).Methods("POST")
	myRouter.HandleFunc("/palette/create", svc.CreatePalette).Methods("POST")

	log.Printf("starting service on %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}
