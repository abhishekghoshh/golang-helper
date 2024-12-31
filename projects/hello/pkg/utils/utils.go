package utils

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/goombaio/namegenerator"
)

func RandomUserName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	return name
}

func RandomServerName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	return name
}

func KubernetesPodName() string {
	data, err := os.ReadFile("/etc/hostname")
	if err != nil {
		log.Printf("Error reading /etc/hostname: %v \n", err)
		return ""
	}
	// Convert the byte slice to a string and trim any whitespace
	hostname := string(data)
	hostname = hostname[:len(hostname)-1]
	hostname = strings.TrimSpace(hostname)
	log.Printf("Hostname: %s\n", hostname)
	return hostname
}
