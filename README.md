KT Cloud SDK go
===========

KT Cloud API library written in Go. Tested towards KT Cloud.

Example usage
-------------

Showing the IP and state of a virtual machine:

```go
package main

import (
	"fmt"
	"os"
	ktcloudsdk "github.com/cloud-barista/ktcloud-sdk-go"
)

func main() {
	
	apiKey := os.Getenv("KTCLOUD_API_KEY")
	if len(apiKey) == 0 {
		fmt.Println("Failed to Find KTCLOUD_API_KEY, exiting")
		os.Exit(1)
	}
	secretKey := os.Getenv("KTCLOUD_SECRET_KEY")
	if len(secretKey) == 0 {
		fmt.Println("Failed to Find KTCLOUD_SECRET_KEY, exiting")
		os.Exit(1)
	}

	// When Zone is "KOR-Seoul M2" => use API v2, else use API v1
	if zoneID == "d7d0177e-6cda-404a-a46f-a5b356d2874e" {
	apiUrl := "https://api.ucloudbiz.olleh.com/server/v2/client/api"
	} else {
	apiUrl := "https://api.ucloudbiz.olleh.com/server/v1/client/api"
	}

	cs := ktcloudsdk.KtCloudClient{}.New(apiUrl, apiKey, secretKey)

	zoneId := "XXXXXXXXXXXXXXXXXXXX"
	vmId := "XXXXXXXXXXXXXXXXXXXX"

	vmListReqInfo := ktcloudsdk.ListVMReqInfo{
		ZoneId: 	zoneId,
		VMId:       vmId,
	}

	response, err := cs.ListVirtualMachines(vmListReqInfo)
	if err != nil {
		fmt.Errorf("Failed to Find the List of Virtual Machine: %s", err)
		os.Exit(1)
	}
	
	if len(response.Listvirtualmachinesresponse.Virtualmachine) > 0 {
		ip := response.Listvirtualmachinesresponse.Virtualmachine[0].Nic[0].Ipaddress
		state := response.Listvirtualmachinesresponse.Virtualmachine[0].State
		fmt.Printf("%s has IP : %s and state : %s\n", vmid, ip, state)
	} else {
		fmt.Printf("Failed to Find the VM with the ID: %s found\n", vmid)
	}

}
```
