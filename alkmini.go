package main

import (
	"log"
	"net"
	"os/exec"
	"time"
)

func getStatus() string {
	timeout := 3 * time.Second
	conn, err := net.DialTimeout("tcp", "alkmini:22", timeout)
	if err != nil {
		log.Println("host unreachable, error: ", err.Error())
		return "Host unreachable"
	}
	defer conn.Close()
	return "Host up and running"
}

func startServer() error {
	cmd := exec.Command("alkmini-wake")
	err := cmd.Run()

	if err != nil {
		log.Println("Error starting server: ", err.Error())
		return err
	}
	return nil
}
