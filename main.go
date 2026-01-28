package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"time"
)

func flagLib() {
	serverPort := flag.Int("port", 8080, "The port to run the server on")
	env := flag.String("env", "dev", "The environment (dev/prod)")

	flag.Parse()

	fmt.Printf("Starting server on port %d in %s mode\n", *serverPort, *env)

}

func timeLib() {
	//This library provides functionality for measuring and displaying time
	//Allows us to find current time , format dates , sleep or pause

	start := time.Now()
	fmt.Println("The current time is ", start)
}

func osLib() {
	//This library helps us to look upto user account information used to know who is running the current program

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Username:", currentUser.Username)
	fmt.Println("Home Directory:", currentUser.HomeDir)
	fmt.Println("User ID (UID):", currentUser.Uid)

}

func main() {

}
