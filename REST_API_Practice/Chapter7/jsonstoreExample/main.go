package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	models "github.com/SaravananPitchaimuthu/REST_API_Practice/Chapter7/jsonstoreExample/helper"
	"github.com/gorilla/mux"
)

type DBClient struct {
	db *gorm.DB
}

func (driver *DBClient) PostPackage(w http.ResponseWriter,
	r *http.Request) {
	var Package = models.Package{}
	postBody, _ := ioutil.ReadAll(r.Body)
	Package.Data = string(postBody)
	driver.db.Save(&Package)
	responseMap := map[string]interface{}{"id": Package.ID}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

type PackageResponse struct {
	Package models.Package `json:"Package"`
}

func (driver *DBClient) GetPackage(w http.ResponseWriter,
	r *http.Request) {
	var Package = models.Package{}
	vars := mux.Vars(r)
	driver.db.First(&Package, vars["id"])
	var PackageData interface{}
	json.Unmarshal([]byte(Package.Data), &PackageData)
	var response = PackageResponse{Package: Package}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(response)
	w.Write(respJSON)
}

func (driver *DBClient) GetPackagesbyWeight(w http.ResponseWriter,
	r *http.Request) {
	var packages []models.Package
	weight := r.FormValue("weight")
	// Handle response details
	var query = "select * from \"Package\" where data->>'weight'=?"
	driver.db.Raw(query, weight).Scan(&packages)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(packages)
	w.Write(respJSON)
}

func main() {

	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{db: db}
	defer db.Close()
	r := mux.NewRouter()
	r.HandleFunc("/v1/package/{id:[a-zA-Z0-9]*}",
		dbclient.GetPackage).Methods("GET")
	r.HandleFunc("/v1/package",
		dbclient.PostPackage).Methods("POST")
	r.HandleFunc("/v1/package",
		dbclient.GetPackagesbyWeight).Methods("GET")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
