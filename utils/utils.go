package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
}

var Config Configuration

// Convenience function for printing to stdout
func P(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatalln("Cannot open config file: ", err)
	}

	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// version
func Version() string {
	return "0.1"
}


// Message function
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond response writer to incomming requests
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
