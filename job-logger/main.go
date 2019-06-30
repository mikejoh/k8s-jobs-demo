package main

import (
	"github.com/gorilla/mux"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func JobHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	q := r.URL.Query()

	pod := q.Get("podname")
	if pod == "" {
		pod = "N/A"
	}

	switch p {
	case "/parallell":
		log.Infof("Pod %s runnning a parallell job.", pod)
	case "/sequential":
		log.Infof("Pod %s runnning a sequential job.", pod)
	case "/single":
		log.Infof("Pod %s runnning a single job.", pod)
	default:
		log.Infof("Path %s not recognized\n", p)
	}
}

func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	r := mux.NewRouter()
	r.HandleFunc("/parallell", JobHandler).Methods("GET")
	r.HandleFunc("/sequential", JobHandler).Methods("GET")
	r.HandleFunc("/single", JobHandler).Methods("GET")

	log.Infof("Starting job logger..")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}