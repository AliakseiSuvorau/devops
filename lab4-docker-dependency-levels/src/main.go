package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Status string

const (
	StatusOK        Status = "ok"
	GreetingMessage        = "Welcome to the custom app"
)

type StatusReport struct {
	Status Status `json:"status"`
}

type LogMsg struct {
	Message string `json:"message"`
}

const (
	dirname  = "logs"
	filename = "app.log"
	filepath = dirname + "/" + filename
)

func PrepareLogFile() {
	if mkdirErr := os.Mkdir(dirname, 0755); mkdirErr != nil && !os.IsExist(mkdirErr) {
		log.Fatalf("Error while creating subdirectories for log file: %v", mkdirErr)
	}

	//file, fileCreateErr := os.Create(filepath)
	//if fileCreateErr != nil {
	//	log.Fatalf("Error while creating a log file: %v", fileCreateErr)
	//}
	//
	//if err := file.Close(); err != nil {
	//	log.Fatalf("Error while closing the log file: %v", err)
	//}
}

func RunServer() {
	http.HandleFunc("/", Greet)
	http.HandleFunc("/status", GetStatus)
	http.HandleFunc("/log", WriteLog)
	http.HandleFunc("/logs", GetLogs)

	port := "6029" // Default port
	if len(os.Args) > 2 && os.Args[1] == "-port" {
		port = os.Args[2]
	}

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	if _, writeErr := w.Write([]byte(GreetingMessage)); writeErr != nil {
		log.Printf("Error while writing response: %v", writeErr)
	}
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, jsonErr := json.Marshal(StatusReport{Status: StatusOK})
	if jsonErr != nil {
		log.Printf("Error while getting status: %v", jsonErr)
	}

	if _, writeErr := w.Write(jsonData); writeErr != nil {
		log.Printf("Error while writing response: %v", writeErr)
	}
}

func WriteLog(w http.ResponseWriter, r *http.Request) {
	reader := r.Body
	defer func() {
		if err := reader.Close(); err != nil {
			log.Fatalf("Error while closing the log file: %v", err)
		}
	}()

	body, readErr := io.ReadAll(reader)
	if readErr != nil && readErr != io.EOF {
		log.Printf("Error while reading request body: %v", readErr)
	}

	var requestBody LogMsg
	if jsonUnmarshalErr := json.Unmarshal(body, &requestBody); jsonUnmarshalErr != nil {
		log.Printf("Error while unmarshalling request body: %v", jsonUnmarshalErr)
	}

	logFile, logFileOpenErr := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if logFileOpenErr != nil {
		log.Fatalf("Error while opening log file: %v", logFileOpenErr)
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Fatalf("Error while closing the log file: %v", err)
		}
	}()

	if _, logWriteErr := logFile.Write([]byte(requestBody.Message + "\n")); logWriteErr != nil {
		log.Printf("Error while writing to log file: %v", logWriteErr)
	}
}

func GetLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	logFile, logFileOpenErr := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
	if logFileOpenErr != nil {
		log.Printf("Error while opening log file: %v", logFileOpenErr)
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Fatalf("Error while closing the log file: %v", err)
		}
	}()

	stat, getLogFileStatsErr := logFile.Stat()
	if getLogFileStatsErr != nil {
		log.Printf("Error while getting log file: %v", getLogFileStatsErr)
	}

	data := make([]byte, stat.Size())
	_, getLogFileStatsErr = logFile.Read(data)
	if getLogFileStatsErr != nil && getLogFileStatsErr != io.EOF {
		log.Printf("Error while reading log file: %v", getLogFileStatsErr)
	}

	if _, writeRequestErr := w.Write(data); writeRequestErr != nil {
		log.Printf("Error while writing to request from log file: %v", writeRequestErr)
	}
}

func main() {
	PrepareLogFile()
	RunServer()
}
