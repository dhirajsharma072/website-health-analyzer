package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	. "github.com/dhirajsharma072/website-health-analyzer/src/dao"
	. "github.com/dhirajsharma072/website-health-analyzer/src/models"
	. "github.com/dhirajsharma072/website-health-analyzer/src/validators"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var dao = WebsiteDAO{}

// GET list of websites
func AllWebsites(w http.ResponseWriter, r *http.Request) {
	websites, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, websites)
}

// POST a new website
func CreateWebsite(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var website Website
	website.ID = uuid.New().String()
	website.CreatedAt = time.Now()
	website.UpdatedAt = time.Now()
	if err := json.NewDecoder(r.Body).Decode(&website); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Insert(website); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, website)
}

// PUT update an existing website
func UpdateWebsite(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	var website Website
	website.ID = vars["id"]
	if err := json.NewDecoder(r.Body).Decode(&website); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if IsValidRequestURL(website.URL) == false {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"result": "Invalid URL"})
		return
	}
	if err := dao.Update(website); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func PatchWebsite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] 

	var sp WebsitePatch
	sp.UpdatedAt = time.Now()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error AddSite", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddSite", err)
	}
	if err := json.Unmarshal(body, &sp); err != nil { // unmarshall body contents as a type Site
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddSite unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := dao.PatchSite(bson.M{"uuid": id}, bson.M{"$set": &sp}) // adds the site to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return
}

// DELETE an existing website
func DeleteWebsite(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var website Website
	vars := mux.Vars(r)
	website.ID = vars["id"]

	if err := dao.Delete(website); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
