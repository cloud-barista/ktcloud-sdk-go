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

func (c KtCloudClient) ListZones(isAvailable bool, domainId string, zoneId string, keyword string) (ListZonesResponse, error) {
	var resp ListZonesResponse
	params := url.Values{}

	if isAvailable {
		params.Set("available", "true")
	}

	if domainId != "" {
		params.Set("domainid", domainId)
	}

	if zoneId != "" {
		params.Set("id", zoneId)
	}

	if keyword != "" {
		params.Set("keyword", keyword)
	}

	response, err := NewRequest(c, "listZones", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListZonesResponse)
	return resp, err
}

type Zone struct {
	ID	                  string            `json:"id"`
	NetworkType           string            `json:"networktype"`
	SecurityGroupsEnabled bool              `json:"securitygroupsenabled"`
	AllocationState       string            `json:"allocationstate"`
	DhcpProvider          string            `json:"dhcpprovider"`
	LocalStorageEnabled   bool              `json:"localstorageenabled"`
	Tags                  []interface{} 	`json:"tags"`
	Name                  string            `json:"name"`
}

type ListZonesResponse struct {
	Listzonesresponse struct {
		Count          int          `json:"count"`
		Zone 		   []Zone 		`json:"zone"`
	} `json:"listZonesResponse"`
}
