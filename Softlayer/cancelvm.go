package main

import (
	"fmt"
	"log"

	"github.com/softlayer/softlayer-go/services"
    "github.com/softlayer/softlayer-go/session"
)

var guestId = 38686969

func main() {
	// Create a session
	sess := session.New()

	// Get the virtual Guest Service
	service := services.GetVirtualGuestService(sess)

	// Set the object ID and delete the guest
	success, err := service.Id(guestId).DeleteObject()

	if err !=nil {
		log.Fatal(err)
	} else if success == false {
		log.Fatal("Error deleting virtual guest")
	} else {
		fmt.Println("Virtual Guest deleted successfully")
	}
}