package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"net/http"
)

func main() {
	podName := os.Getenv("POD_NAME")
	jobType := os.Getenv("JOB_TYPE")
	logEndpoint := os.Getenv("LOG_ENDPOINT")

	resp, err := http.Get(logEndpoint + "/" + jobType + "?podname=" + podName)
	if err != nil {
		log.Errorf("Error sending job log request to %s", logEndpoint)
	}

	defer resp.Body.Close()
}
