package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(uuid.New())

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/Dataset", createDataset).Methods("POST")
	r.HandleFunc("/api/v1/Datasets", getFiles).Methods("GET")
	r.HandleFunc("/api/v1/files/{id}", getFile).Methods("GET")
	r.HandleFunc("/api/v1/files/{id}", updateFile).Methods("PATCH")
	r.HandleFunc("/api/v1/files/{id}", deleteFile).Methods("DELETE")

	// db := connect()
	// defer db.Close()

	// file := File{MCID: uuid.New()}
	// fmt.Println(file)

	// Starting Server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Create a Dataset
func createDataset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := connect()
	defer db.Close()

	
	//TODO Upload File to Estuary and Receive CID
	dataID := uuid.New().String()

	dataVersion := &DataVersion{
		DataID: dataID,
		Version: 0,
		//TODO add CID
		Date: time.Now(),
		//TODO add UploadedBy
	}
	

	dataSet := &Dataset{
		DataID: uuid.New().String(),
		CurrentVersionNumber: 0,
		CurrentVersion: *dataVersion,
		//TODO add owners
	}
	

	// Insert into Database
	_, err := db.Model(dataSet).Insert()

	
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

// Get Files
func getFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db := connect()
	defer db.Close()

	// Creating Files Slice
	var files []File
	if err := db.Model(&files).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Returning Files
	json.NewEncoder(w).Encode(files)
}

// Get File
func getFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db := connect()
	defer db.Close()

	// Get MCID
	params := mux.Vars(r)
	fileMCID := params["mcid"]

	// Creating file Instance
	file := &File{MCID: fileMCID}
	if err := db.Model(file).WherePK().Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Returning file
	json.NewEncoder(w).Encode(file)
}

// Update file
func updateFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db := connect()
	defer db.Close()

	// Get MCID
	params := mux.Vars(r)
	fileMCID := params["mcid"]

	// Creating file Instance
	file := &File{MCID: fileMCID}

	_ = json.NewDecoder(r.Body).Decode(&file)

	
	_, err := db.Model(file).WherePK().Set("MCID = ?, CID = ?, Name = ?, Collection = ?", file.MCID, file.CID, file.Name, file.Collection).Update()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Returning File
	json.NewEncoder(w).Encode(file)
}

// Delete File
func deleteFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db := connect()
	defer db.Close()

	// Get ID
	// Get MCID
	params := mux.Vars(r)
	fileMCID := params["mcid"]

	// Creating File Instance Alternative Way
	// file := &File{MCID: fileMCID}
	// result, err := db.Model(file).WherePK().Delete()

	// Creating File Instance
	file := &File{}
	result, err := db.Model(file).Where("id = ?", fileMCID).Delete()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Returning result
	json.NewEncoder(w).Encode(result)
}