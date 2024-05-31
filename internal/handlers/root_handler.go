package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nafnaufal/roadmap-forum/internal/db"
)

type Response struct {
	Count int `json:"count"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	session, err := db.ConnectToNeo4j()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.CloseSession(session)

	result, err := db.RunCypherQuery(session, "MATCH (n) RETURN count(n) AS count", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response Response
	for result.Next() {
		record := result.Record()
		if count, found := record.Get("count"); found {
			response.Count = int(count.(int64))
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
