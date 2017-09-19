package main

import (
	"fmt"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

func main() {
	sess := session.New()
	doCreateVMTest(sess)
}

func doCreateVMTest(sess *session.Session) {
	service := services.GetVirtualGuestService(sess)

	// Create a Virtual_Guest instance as a template
	vGuestTemplate := datatypes.Virtual_Guest{
		Hostname:                     sl.String("goapi"),
		Domain:                       sl.String("iptest.com"),
		MaxMemory:                    sl.Int(6144),
		StartCpus:                    sl.Int(2),
		Datacenter:                   &datatypes.Location{Name: sl.String("seo01")},
		OperatingSystemReferenceCode: sl.String("WIN_2012-STD-R2_64"),
		LocalDiskFlag:                sl.Bool(false),
		HourlyBillingFlag:            sl.Bool(true),
		PrivateNetworkOnlyFlag:       sl.Bool(false),
	}

	vGuest, err := service.Mask("id;domain").CreateObject(&vGuestTemplate)
	if err != nil {
		fmt.Printf("s\n", err)
		return
	} else {
		fmt.Printf("\nNew Virtual Guest created with ID %d\n", *vGuest.Id)
		fmt.Printf("Domain: %s\n", *vGuest.Domain)
	}
}
