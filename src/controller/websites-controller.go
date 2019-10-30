package controller

import (
	"encoding/json"
	"net/http"
	"time"

	. "github.com/dhirajsharma072/website-health-analyzer/src/dao"
	. "github.com/dhirajsharma072/website-health-analyzer/src/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = WebsiteDAO{}

// GET list of websites
func AllWebsitesEndPoint(w http.ResponseWriter, r *http.Request) {
	websites, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, websites)
}

// POST a new website
func CreateWebsiteEndPoint(w http.ResponseWriter, r *http.Request) {
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
func UpdateWebsiteEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	var website Website
	website.ID = vars["id"]
	if err := json.NewDecoder(r.Body).Decode(&website); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(website); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func PatchWebsiteEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	var wp WebsitePatch
	var id = vars["id"]
	wp.UpdatedAt = time.Now()

	if err := json.NewDecoder(r.Body).Decode(&wp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Patch(bson.M{"uuid": id}, bson.M{"$set": &wp}); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing website
func DeleteWebsiteEndPoint(w http.ResponseWriter, r *http.Request) {
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
