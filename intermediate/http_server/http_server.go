package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	store = map[string]Lake{}
	lock  = &sync.Mutex{}
)

type (
	Lake struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Area int32  `json:"area"`
	}
	Action struct {
		Type    string `json:"type"`
		Payload string `json:"payload"`
	}
)

func PostHandler(w http.ResponseWriter, req *http.Request) {
	var action Action
	err := json.NewDecoder(req.Body).Decode(&action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var lake Lake
	if err := json.Unmarshal([]byte(action.Payload), &lake); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lock.Lock()
	store[lake.Id] = lake
	lock.Unlock()

	w.WriteHeader(http.StatusOK)
}

func DeleteHandler(w http.ResponseWriter, req *http.Request) {
	var action Action
	err := json.NewDecoder(req.Body).Decode(&action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lock.Lock()
	defer lock.Unlock()

	lake, ok := store[action.Payload]
	if !ok {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	delete(store, lake.Id)
}

func GetHandler(w http.ResponseWriter, req *http.Request) {
	var action Action
	err := json.NewDecoder(req.Body).Decode(&action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lake, ok := store[action.Payload]
	if !ok {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n%v", lake.Name, lake.Area)
}
