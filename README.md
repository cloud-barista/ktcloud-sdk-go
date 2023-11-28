KT Cloud SDK Go
===========

Go SDK for KT Cloud G1/G2 Platform REST API.
This Go SDK is being used for the KT Cloud Classic connection driver of [Cloud-Barista](https://github.com/cloud-barista).

Example usage of the SDK
-------------

Showing the Guest OS and State of a Product Type(VM Image) :

```go
package main

import (
	"fmt"
	"strings"
	"os"
	"errors"
	ktsdk "github.com/cloud-barista/ktcloud-sdk-go"
)

func main() {
	imgIdToGetInfo := "87838094-af4f-449f-a2f4-f5b4b581eb29" // An Image ID on 'KOR-Seoul M' zone.
	zoneId := "95e2f517-d64a-4866-8585-5177c256f7c7" // KT Cloud 'KOR-Seoul M' zone ID

	guestOS, imgStatus, err := GetVMImageInfo(imgIdToGetInfo, zoneId)
	if err != nil {
		fmt.Errorf("Failed to Find the Image Info : [%v]", err)
		os.Exit(1)
	}
	fmt.Printf("# Guest OS : [%s], Image Status : [%s] of the Image ID.\n", guestOS, imgStatus)
}

func GetVMImageInfo(imageId string, zoneId string) (string, string, error) {
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

	// If KT Cloud Zone is 'KOR-Seoul M2' => use API v2, else use API v1.
	var apiUrl string
	if zoneId == "d7d0177e-6cda-404a-a46f-a5b356d2874e" { // 'KOR-Seoul M2' zone
	apiUrl = "https://api.ucloudbiz.olleh.com/server/v2/client/api" // API v2
	} else {
	apiUrl = "https://api.ucloudbiz.olleh.com/server/v1/client/api" // API v1
	}

	// Always validate any SSL certificates in the chain
	insecureskipverify := false
	cs := ktsdk.KtCloudClient{}.New(apiUrl, apiKey, secretKey, insecureskipverify)

	result, err := cs.ListAvailableProductTypes(zoneId)
	if err != nil {
		return "", "", fmt.Errorf("Failed to Find the List of Product Types : [%v]", err)
	}

	if len(result.Listavailableproducttypesresponse.ProductTypes) < 1 {
		return "", "", errors.New("Failed to Get Product Type List!!")
	} else {
		var guestOS string
		var imgStatus string
		for _, productType := range result.Listavailableproducttypesresponse.ProductTypes {
			if strings.EqualFold(productType.TemplateId, imageId) {	
				guestOS 	= productType.TemplateDesc
				imgStatus 	= productType.ProductState
				break
			}
		}
		if !strings.EqualFold(guestOS, "") {
			return guestOS, imgStatus, nil		
		} else {
			return "", "", errors.New("Failed to Find the Product Types in the Zone!!")
		}
	}
}
```
## Original source code of ktcloud-sdk-go
The original source code, [gopherstack](https://github.com/mindjiver/gopherstack) is a CloudStack Go SDK.

< Original code Licensed under the MIT >
