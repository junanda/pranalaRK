package jawaban

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Items struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var DataItems []Items

func handleItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var dataResult, err = json.Marshal(DataItems)
		if err != nil {
			log.Println("Failed parse data: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(dataResult)
		return
	}

	if r.Method == "POST" {
		var dataReq Items

		dataBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("failed read data body:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(dataBody, &dataReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		DataItems = append(DataItems, dataReq)
		result, err := json.Marshal(DataItems)
		if err != nil {
			log.Println("Error Marshal data: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func Run() {
	http.HandleFunc("/items", handleItems)
	log.Println("server running....")

	http.ListenAndServe(":8080", nil)
}
