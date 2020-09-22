package signal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := Person{
		Name:       "Vinicios",
		Age:        28,
		Occupation: "Software Engineer",
	}
	json, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(json); err != nil {
		e := fmt.Sprintf("ISE: %s", err.Error())
		http.Error(w, e, http.StatusInternalServerError)
	}
}
