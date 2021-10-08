// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista
//
// KT Cloud SDK go
//
// by ETRI, 2021.08.

package ktcloudsdk

import (
	"net/url"
)

type AssociatePublicIpReqInfo struct {
	ZoneId 				string
	UsagePlanType		string
	Account 			string
	DomainId 			string
	NetworkId			string
}

type ListPublicIpReqInfo struct {
	ID 					string   // PublicIP ID
	VLanId				string
	IpAddress			string
	Account 			string
	DomainId 			string
	IsRecursive 		bool
	Keyword 			string
	Page				string
	PageSize			string
	ZoneId 				string
	IsSourceNat			bool
	IsStaticNat			bool
	AssociatedNetworkId	string
	PhysicalNetworkId	string
	ForVirtualNetwork	string
	AllocatedOnly		string
	ListAll				bool
}


func (c KtCloudClient) AssociateIpAddress(ipReqInfo AssociatePublicIpReqInfo) (AssociateIpAddressResponse, error) {
	var resp AssociateIpAddressResponse
	params := url.Values{}

	params.Set("zoneid", ipReqInfo.ZoneId)

	if ipReqInfo.UsagePlanType != "" {
		params.Set("usageplantype", ipReqInfo.UsagePlanType)
	}
	
	if ipReqInfo.Account != "" {
		params.Set("account", ipReqInfo.Account)
	}

	if ipReqInfo.DomainId != "" {
		params.Set("domainid", ipReqInfo.DomainId)
	}

	if ipReqInfo.NetworkId != "" {
		params.Set("networkid", ipReqInfo.NetworkId)
	}

	response, err := NewRequest(c, "associateIpAddress", params)
	if err != nil {
		return resp, err
	}
	resp = response.(AssociateIpAddressResponse)
	return resp, err
}


func (c KtCloudClient) ListPublicIpAddresses(ipListReqInfo ListPublicIpReqInfo) (ListPublicIpAddressesResponse, error) {
	var resp ListPublicIpAddressesResponse
	params := url.Values{}

	if ipListReqInfo.ID != "" {
		params.Set("id", ipListReqInfo.ID)
	}

	if ipListReqInfo.VLanId != "" {
		params.Set("vlanid", ipListReqInfo.VLanId)
	}

	if ipListReqInfo.IpAddress != "" {
		params.Set("ipaddress", ipListReqInfo.IpAddress)
	}

	if ipListReqInfo.Account != "" {
		params.Set("account", ipListReqInfo.Account)
	}

	if ipListReqInfo.DomainId != "" {
		params.Set("domainid", ipListReqInfo.DomainId)
	}

	if ipListReqInfo.IsRecursive {
		params.Set("isrecursive", "true")
	}
	
	if ipListReqInfo.Keyword != "" {
		params.Set("keyword", ipListReqInfo.Keyword)
	}

	if ipListReqInfo.Page != "" {
		params.Set("page", ipListReqInfo.Page)
	}

	if ipListReqInfo.PageSize != "" {
		params.Set("pagesize", ipListReqInfo.PageSize)
	}

	if ipListReqInfo.ZoneId != "" {
		params.Set("zoneid", ipListReqInfo.ZoneId)
	}

	if ipListReqInfo.IsSourceNat {
		params.Set("issourcenat", "true")
	}

	if ipListReqInfo.IsStaticNat {
		params.Set("isstaticnat", "true")
	}

	if ipListReqInfo.AssociatedNetworkId != "" {
		params.Set("associatednetworkid", ipListReqInfo.AssociatedNetworkId)
	}

	if ipListReqInfo.PhysicalNetworkId != "" {
		params.Set("physicalnetworkid", ipListReqInfo.PhysicalNetworkId)
	}

	if ipListReqInfo.ForVirtualNetwork != "" {
		params.Set("forvirtualnetwork", ipListReqInfo.ForVirtualNetwork)
	}

	if ipListReqInfo.AllocatedOnly != "" {
		params.Set("allocatedonly", ipListReqInfo.AllocatedOnly)
	}

	if ipListReqInfo.ListAll {
		params.Set("listall", "true")
	}

	response, err := NewRequest(c, "listPublicIpAddresses", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListPublicIpAddressesResponse)
	return resp, err
}

func (c KtCloudClient) DisassociateIpAddress(publicIpId string) (DisassociateIpAddressResponse, error) {
	var resp DisassociateIpAddressResponse
	params := url.Values{}

	params.Set("id", publicIpId)

	response, err := NewRequest(c, "disassociateIpAddress", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DisassociateIpAddressResponse)
	return resp, err
}

type PublicIpAddress struct {
	ID                    string            `json:"id"`
	IpAddress			  string 			`json:"ipaddress"`
	Allocated		      string 			`json:"allocated"`
	ZoneId           	  string        	`json:"zoneid"`
	ZoneName              string        	`json:"zonename"`
	IsSourcenat 		  bool              `json:"issourcenat"`
	Account           	  string        	`json:"account"`
	DomainId         	  string        	`json:"domainid"`
	Domain                string        	`json:"domain"`
	ForVirtualNetwork	  string        	`json:"forvirtualnetwork"`
	IsStaticNat 		  bool              `json:"isstaticnat"`
	IsSystem 		 	  bool              `json:"issystem"`
	AssociatedNetworkId   string            `json:"associatednetworkid"`
	AssociatedNetworkName   string          `json:"associatednetworkname"`
	NetworktId            string            `json:"networkid"`
	State                 string            `json:"state"`
	PhysicalNetworkId     string            `json:"physicalnetworkid"`
	Tags                  []interface{} 	`json:"tags"`
	IsPortable 		 	  bool              `json:"isportable"`
	UsagePlanType         string            `json:"usageplantype"`
	Desc                  string            `json:"desc"`
}

type AssociateIpAddressResponse struct {
	Associateipaddressresponse struct {
		ID    string `json:"id"`  // PublicIP ID
		JobId string `json:"jobid"`
	} `json:"associateipaddressresponse"`
}

type ListPublicIpAddressesResponse struct {
	Listpublicipaddressesresponse struct {
		Count          	  int          			`json:"count"`
		PublicIpAddress   []PublicIpAddress 	`json:"publicipaddress"`
	} `json:"listpublicipaddressesresponse"`
}

type DisassociateIpAddressResponse struct {
	Disassociateipaddressresponse struct {
		JobId string `json:"jobid"`
	} `json:"disassociateipaddressresponse"`
}