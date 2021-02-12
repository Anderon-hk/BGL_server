package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Obj struct {
	Time string
	Key   string
	Value string
}

type PObj struct {
	Key   string
	Value string
}

func responsehd(w http.ResponseWriter, r *http.Request){
	list := make([]Obj, 0)

	for key, val := range(valMap){
		o := Obj{timeMap[key].UTC().Format("2006-01-02T15:04:05Z07:00"), key, val}
		list = append(list, o)
	}

	response, _ := json.Marshal(list)
	fmt.Println(string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func posthd(w http.ResponseWriter, r *http.Request){
	var po PObj
	json.NewDecoder(r.Body).Decode(&po)

	valMap[po.Key] = po.Value
	timeMap[po.Key] = time.Now()
}

var valMap = make(map[string]string)
var timeMap = make(map[string]time.Time)

func main(){

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/list", func (w http.ResponseWriter, r *http.Request){
		responsehd(w, r)
	})

	r.Post("/add", func (w http.ResponseWriter, r *http.Request){
		posthd(w, r)
	})

	http.ListenAndServe(":80", r)
}