package main

import (
	"errors"
	"log"
	"net"
	"os/exec"
	"time"
)

func getStatus() string {
	timeout := 2 * time.Second
	_, err := net.DialTimeout("tcp", "alkmini:22", timeout)
	if err != nil {
		log.Println("host unreachable, error: ", err.Error())
		return STATUS_DOWN
	}
	return STATUS_UP
}

func startServer() error {
	cmd := exec.Command("/Users/alknopfler/bin/alkmini-wake")

	_, err := cmd.Output()

	if err != nil {
		log.Println("Error executing command startServer: ", err.Error())
		return err
	}
	return nil
}

func stopServer() error {
	cmd := exec.Command("ssh", "alkmini", "sudo systemctl suspend")
	err := cmd.Run()
	if err != nil {
		log.Println("Error executing command stopServer: ", err.Error())
		return err
	}
	return nil
}

func createTunnel() error {
	if getStatus() != STATUS_UP {
		log.Println("Cannot create tunnel, server is down")
		return errors.New("Cannot create tunnel, server is down")
	}
	cmd := exec.Command("sshuttle", "-D", "-r", "alkmini", "192.168.122.0/24")
	err := cmd.Run()
	if err != nil {
		log.Println("Error executing command createTunnel: ", err.Error())
		return err
	}
	return nil
}

func waitUntilUp() {
	for getStatus() != STATUS_UP {
		time.Sleep(5 * time.Second)
	}
}

func waitUntilDown() {
	for getStatus() != STATUS_DOWN {
		time.Sleep(5 * time.Second)
	}
}
