package home

import (
	"encoding/json"
	"net/http"
)

type AppInfo struct {
	Version string
	Name    string
	Author  string
}

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	appInfo := AppInfo{"0.0.1", "freed", "Dennis Schoepf"}
	json, err := json.Marshal(appInfo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(json)
}
