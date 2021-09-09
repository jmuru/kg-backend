package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kat-generator/KGB/client"
	"io/ioutil"
	"net/http"
)

type Service struct {
	kc *client.KatClient
}

func NewService() *Service {
	//https://immense-shelf-23449.herokuapp.com/
	c, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	return &Service{
		kc: c,
	}
}

func (s *Service) HelloWorld(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello world")
}

func (s *Service) GetAccessory(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	placement, ok := vars["placement"] // top, mid, bottom
	if !ok {
		fmt.Println("missing query parameter")
	}
	result, err := s.kc.GetAccessoryData(placement)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	fmt.Fprint(w, string(data))
}

func (s *Service) GetFace(w http.ResponseWriter, req *http.Request) {
	result, err := s.kc.GetFaceData()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	fmt.Fprint(w, "result: %+v", string(data))
}

func (s *Service) GetKat(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := s.kc.GetKat()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	//json.NewEncoder(w).Encode(data)
	fmt.Fprint(w, string(data))
}

func (s *Service) GetBackground(w http.ResponseWriter, req *http.Request) {
	result, err := s.kc.GetBackgroundData()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	fmt.Fprint(w, "result: %+v", string(data))
}

func (s *Service) GetPalette(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pt, ok := vars["type"]
	if !ok {
		fmt.Println("missing query parameter")
	}
	result, err := s.kc.GetPaletteData(pt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	fmt.Fprint(w, string(data))
}

func (s *Service) CreateAccessory(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var a client.AccessoryDataRequest
	err = json.Unmarshal(b, &a)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Printf(fmt.Sprintf("accessory: %+v \n", a.Data.Accessory))
	fmt.Printf(fmt.Sprintf("subtype: %+v \n", a.Data.SubType))
	fmt.Printf(fmt.Sprintf("placement: %+v \n", a.Data.Placement))

	//err := json.NewDecoder(req.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dataObj := &client.AccessoryData{
		Placement: a.Data.Placement,
		Accessory: a.Data.Accessory,
		SubType:   a.Data.SubType,
	}
	fmt.Printf(fmt.Sprintf("accessory to save: %+v \n", dataObj))
	if err := s.kc.CreateAccessoryData(dataObj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Service) CreatePalette(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	var p client.PaletteData
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dataObj := &client.PaletteData{
		Palette: p.Palette,
		Type:    p.Type,
	}
	if err := s.kc.CreatePaletteData(dataObj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}
