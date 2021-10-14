// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista
//
// KT Cloud SDK go
//
// by ETRI, 2021.08.

package main

import (
	"os"
	"fmt"
	"github.com/davecgh/go-spew/spew"

	ktsdk "github.com/cloud-barista/ktcloud-sdk-go"
)

func main() {

	//apiurl := os.Getenv("KTCLOUD_API_URL")

	//When Zoneid, Zonename
    // Id: (string) (len=36) "eceb5d65-6571-4696-875f-5a17949f3317",
	// Name: (string) (len=13) "KOR-Central A"

	// Id: (string) (len=36) "9845bd17-d438-4bde-816d-1b12f37d5080",
	// Name: (string) (len=13) "KOR-Central B"

	// Id: (string) (len=36) "dfd6f03d-dae5-458e-a2ea-cb6a55d0d994",
	// Name: (string) (len=6) "KOR-HA"

	// Id: (string) (len=36) "95e2f517-d64a-4866-8585-5177c256f7c7",
	// Name: (string) (len=11) "KOR-Seoul M"

	// Id: (string) (len=36) "b7eb18c8-876d-4dc6-9215-3bd455bb05be",
	// Name: (string) (len=7) "US-West"
	// apiurl := "https://api.ucloudbiz.olleh.com/server/v1/client/api"	


	//When Zoneid, Zonename
    // Id: (string) (len=36) "d7d0177e-6cda-404a-a46f-a5b356d2874e",
    // Name: (string) (len=12) "KOR-Seoul M2"	
	apiurl := "https://api.ucloudbiz.olleh.com/server/v2/client/api"

	if len(apiurl) == 0 {
		fmt.Println("Needed environment variable KTCLOUD_API_URL not found, exiting")
		os.Exit(1)
	}

	apikey := os.Getenv("KTCLOUD_API_KEY")
	if len(apikey) == 0 {
		fmt.Println("Needed environment variable KTCLOUD_API_KEY not found, exiting")
		os.Exit(1)
	}

	secretkey := os.Getenv("KTCLOUD_SECRET_KEY")
	if len(secretkey) == 0 {
		fmt.Println("Needed environment variable KTCLOUD_SECRET_KEY not found, exiting")
		os.Exit(1)
	}

	// Always validate any SSL certificates in the chain
	insecureskipverify := false
	cs := ktsdk.KtCloudClient{}.New(apiurl, apikey, secretkey, insecureskipverify)

	//zoneid := "eceb5d65-6571-4696-875f-5a17949f3317"
	response, err := cs.ListZones(true, "", "", "")
	if err != nil {
		fmt.Errorf("Failed to Get the List of Zones: %s", err)
		os.Exit(1)
	}
	
	spew.Dump(response)
}