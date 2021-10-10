package main

import ("fmt"
		"log"
		"os"
		"github.com/jainpiyush26/network-monitor/process_settings")

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("Welcome to the network monitor code!")
	settings.Run()
}
