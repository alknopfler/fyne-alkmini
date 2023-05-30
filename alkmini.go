package main

import (
	"errors"
	"log"
	"net"
	"os"
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
	if _, err := os.Stat("/tmp/sshuttle.pid"); err == nil {
		return STATUS_TUNNELED
	}
	return STATUS_UP
}

func startServer() error {

	state := getStatus()

	if state == STATUS_DOWN {
		cmd := exec.Command("/Users/alknopfler/bin/alkmini-wake")
		_, err := cmd.Output()
		if err != nil {
			log.Println("Error executing command startServer: ", err.Error())
			return err
		}
	}
	return nil
}

func stopServer() error {
	state := getStatus()
	if state != STATUS_DOWN {
		if state == STATUS_TUNNELED {
			err := removeTunnel()
			if err != nil {
				log.Println("Error executing command stopTunnel: ", err.Error())
				return err
			}
		}
		cmd := exec.Command("ssh", "alkmini", "sudo systemctl suspend")
		err := cmd.Run()
		if err != nil {
			log.Println("Error executing command stopServer: ", err.Error())
			return err
		}
	}
	return nil
}

func createTunnel() error {
	state := getStatus()
	if state == STATUS_DOWN {
		log.Println("Cannot create tunnel, server is down")
		return errors.New("Cannot create tunnel, server is down")
	}
	if state == STATUS_TUNNELED {
		log.Println("Server is up and tunneled already")
		return nil
	}
	cmd := exec.Command("sshuttle", "-D", "--pidfile", "/tmp/sshuttle.pid", "-r", "alkmini", "192.168.122.0/24")
	err := cmd.Run()
	if err != nil {
		log.Println("Error executing command createTunnel: ", err.Error())
		return err
	}
	return nil
}

func removeTunnel() error {
	cmd := exec.Command("pkill", "-F", "/tmp/sshuttle.pid")
	err := cmd.Run()
	if err != nil {
		log.Println("Error executing command removeTunnel: ", err.Error())
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
