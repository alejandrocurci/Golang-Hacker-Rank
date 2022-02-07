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
	var lake Lake
	err := json.NewDecoder(req.Body).Decode(&lake)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lock.Lock()
	store[lake.Id] = lake
	lock.Unlock()

	w.WriteHeader(http.StatusOK)
}

func DeleteHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query()["id"]

	lock.Lock()
	defer lock.Unlock()

	lake, ok := store[id[0]]
	if !ok {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	delete(store, lake.Id)
}

func GetHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query()["id"]
	lake, ok := store[id[0]]
	if !ok {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	b, err := json.Marshal(&lake)
	if err != nil {
		http.Error(w, "error marshalling response", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(b))
}
