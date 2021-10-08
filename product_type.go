// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista
//
// KT Cloud SDK go
//
// by ETRI, 2021.07.

package ktcloudsdk

import (
	"net/url"
)

func (c KtCloudClient) ListAvailableProductTypes(zoneId string) (ListAvailableProductTypesResponse, error) {
	var resp ListAvailableProductTypesResponse
	params := url.Values{}

	if zoneId != "" {
		params.Set("zoneid", zoneId)
	}

	response, err := NewRequest(c, "listAvailableProductTypes", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListAvailableProductTypesResponse)
	return resp, err
}

type ProductTypes struct {
	DiskOfferingDesc  		string            `json:"diskofferingdesc"`
	DiskOfferingId        	string            `json:"diskofferingid"`
	Product			      	string            `json:"product"`
	ProductId       	    string            `json:"productid"`
	ProductState            string            `json:"productstate"`
	ServiceOfferingDesc 	string            `json:"serviceofferingdesc"`
	ServiceOfferingId     	string            `json:"serviceofferingid"`
	TemplateDesc		   	string            `json:"templatedesc"`
	TemplateId       	    string            `json:"templateid"`
	ZoneDesc            	string            `json:"zonedesc"`
	ZoneId            		string            `json:"zoneid"`
}

type ListAvailableProductTypesResponse struct {
	Listavailableproducttypesresponse struct {
		Count          int          `json:"count"`
		ProductTypes []ProductTypes `json:"producttypes"`
	} `json:"listavailableproducttypesresponse"`
}
