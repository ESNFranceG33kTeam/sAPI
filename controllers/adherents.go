package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/models"
	"github.com/gorilla/mux"
)

func AdherentsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(models.AllAdherents())
	if err != nil {
		logger.LogError("adherent", "problem with indexation.", err)
	} else {
		logger.LogInfo("adherent", "request GET : "+r.RequestURI)
	}
}

func AdherentsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.LogError("adherent", "problem with create.", err)
	}

	var adh models.Adherent

	err = json.Unmarshal(body, &adh)
	if err != nil {
		logger.LogError("adherent", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", adh.Dateofbirth)

	if err != nil {
		logger.LogInfo("adherent", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	models.NewAdherent(&adh)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		logger.LogError("adherent", "problem with encoder.", err)
	} else {
		logger.LogInfo("adherent", "request POST : "+r.RequestURI)
	}
}

func AdherentsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.LogError("adherent", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	adh := models.FindAdherentById(id)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		logger.LogError("adherent", "problem with encoder.", err)
	} else {
		logger.LogInfo("adherent", "request GET : "+r.RequestURI)
	}
}

func AdherentsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.LogError("adherent", "unable to get id.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.LogError("adherent", "problem with update.", err)
	}

	adh := models.FindAdherentById(id)

	err = json.Unmarshal(body, &adh)
	if err != nil {
		logger.LogError("adherent", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", adh.Dateofbirth)

	if err != nil {
		logger.LogInfo("adherent", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	models.UpdateAdherent(adh)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		logger.LogError("adherent", "problem with encoder.", err)
	} else {
		logger.LogInfo("adherent", "request PUT : "+r.RequestURI)
	}
}

func AdherentsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.LogError("adherent", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = models.DeleteAdherentById(id)
	if err != nil {
		logger.LogError("adherent", "unable to delete adherent.", err)
	} else {
		logger.LogInfo("adherent", "request DELETE : "+r.RequestURI)
	}
}
