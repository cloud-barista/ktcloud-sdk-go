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

	apiurl := os.Getenv("KTCLOUD_API_URL")
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
	if len(secret) == 0 {
		fmt.Println("Needed environment variable KTCLOUD_SECRET_KEY not found, exiting")
		os.Exit(1)
	}

	cs := ktcloudsdk.KtCloudClient{}.New(apiurl, apikey, secretkey)

	vmId := "19d2acfb-e281-4a13-8d62-e04ab501271d"
	zoneId := "XXXXXX"

	vmListReqInfo := ktsdk.ListVMReqInfo{
		ZoneId: 	zoneId,
		VMId:       vmId,
	}

	response, err := cs.ListVirtualMachines(vmListReqInfo)
	if err != nil {
		fmt.Errorf("Error listing virtual machine: %s", err)
		os.Exit(1)
	}
	
	if len(response.Listvirtualmachinesresponse.Virtualmachine) > 0 {
		ip := response.Listvirtualmachinesresponse.Virtualmachine[0].Nic[0].Ipaddress
		state := response.Listvirtualmachinesresponse.Virtualmachine[0].State
		fmt.Printf("%s has IP : %s and state : %s\n", vmid, ip, state)
	} else {
		fmt.Printf("No VM with UUID: %s found\n", vmid)
	}

}
```
