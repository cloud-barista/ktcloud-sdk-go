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

type CreatePortForwardingRuleReqInfo struct {
	IpAddressId         	    string
	PrivatePort      			string
	Protocol         			string   //TCP or UDP
	PublicPort       			string	 //Port of the public IP Address
	VirtualmachineId    		string
	OpenFirewall  				bool
	//CidrList      				string  // # Parameter cidrList is deprecated!!
	PrivateEndPort   			string
	PublicEndPort    			string
}

type ListPortForwardingRulesReqInfo struct {
	ID	 				string
	IpAddressId 		string
	Account 			string
	DomainId 			string
	IsRecursive 		bool
	Keyword 			string
	Page				string
	PageSize			string
	ListAll				bool
}

// Creates a PortForwardingRule
func (c KtCloudClient) CreatePortForwardingRule(portForwardingRuleCreateReqInfo CreatePortForwardingRuleReqInfo) (CreatePortForwardingRuleResponse, error) {
	var resp CreatePortForwardingRuleResponse
	params := url.Values{}

	params.Set("ipaddressid", portForwardingRuleCreateReqInfo.IpAddressId)
	params.Set("privateport", portForwardingRuleCreateReqInfo.PrivatePort)
	params.Set("protocol", portForwardingRuleCreateReqInfo.Protocol)
	params.Set("publicport", portForwardingRuleCreateReqInfo.PublicPort)
	params.Set("virtualmachineid", portForwardingRuleCreateReqInfo.VirtualmachineId)

	if portForwardingRuleCreateReqInfo.OpenFirewall {
		params.Set("openfirewall", "true")
	}

	// if portForwardingRuleCreateReqInfo.CidrList != "" {
	// 	params.Set("cidrlist", portForwardingRuleCreateReqInfo.CidrList)
	// }
	
	if portForwardingRuleCreateReqInfo.PrivateEndPort != "" {
		params.Set("privateendport", portForwardingRuleCreateReqInfo.PrivateEndPort)
	}

	if portForwardingRuleCreateReqInfo.PublicEndPort != "" {
		params.Set("publicendport", portForwardingRuleCreateReqInfo.PublicEndPort)
	}
	
	response, err := NewRequest(c, "createPortForwardingRule", params)
	if err != nil {
		return resp, err
	}

	resp = response.(CreatePortForwardingRuleResponse)
	return resp, err
}


// Returns all available templates
func (c KtCloudClient) ListPortForwardingRules(portForwardingRulesListReqInfo ListPortForwardingRulesReqInfo) (ListPortForwardingRulesResponse, error) {
	var resp ListPortForwardingRulesResponse
	params := url.Values{}

	if portForwardingRulesListReqInfo.ID != "" {
		params.Set("id", portForwardingRulesListReqInfo.ID)
	}

	if portForwardingRulesListReqInfo.IpAddressId != "" {
		params.Set("ipaddressid", portForwardingRulesListReqInfo.IpAddressId)
	}

	if portForwardingRulesListReqInfo.Account != "" {
		params.Set("account", portForwardingRulesListReqInfo.Account)
	}

	if portForwardingRulesListReqInfo.DomainId != "" {
		params.Set("domainid", portForwardingRulesListReqInfo.DomainId)
	}

	if portForwardingRulesListReqInfo.IsRecursive {
		params.Set("isrecursive", "true")
	}
	
	if portForwardingRulesListReqInfo.Keyword != "" {
		params.Set("keyword", portForwardingRulesListReqInfo.Keyword)
	}

	if portForwardingRulesListReqInfo.Page != "" {
		params.Set("page", portForwardingRulesListReqInfo.Page)
	}

	if portForwardingRulesListReqInfo.PageSize != "" {
		params.Set("pagesize", portForwardingRulesListReqInfo.PageSize)
	}

	if portForwardingRulesListReqInfo.ListAll {
		params.Set("listall", "true")
	}
	
	response, err := NewRequest(c, "listPortForwardingRules", params)
	if err != nil {
		return resp, err
	}

	resp = response.(ListPortForwardingRulesResponse)
	return resp, err
}

// Deletes a PortForwarding Rule by its ID.
func (c KtCloudClient) DeletePortForwardingRule(ruleId string) (DeletePortForwardingRuleResponse, error) {
	var resp DeletePortForwardingRuleResponse
	params := url.Values{}
	params.Set("id", ruleId)  // PortForwardingRule ID
	
	response, err := NewRequest(c, "deletePortForwardingRule", params)
	if err != nil {
		return resp, err
	}

	resp = response.(DeletePortForwardingRuleResponse)
	return resp, err
}

type PortForwardingRule struct {
	ID               			string  		`json:"id"`
	PrivatePort      			string  		`json:"privateport"`
	PrivateEndPort   			string  		`json:"privateendport"`
	Protocol         			string  		`json:"protocol"`
	PublicPort       			string  		`json:"publicport"`
	PublicEndPort    			string  		`json:"publicendport"`
	VirtualmachineId    		string  		`json:"virtualmachineid"`
	VirtualmachineName  		string  		`json:"virtualmachinename"`
	VirtualmachineDisplayName	string  		`json:"virtualmachinedisplayname"`
	IpAddressId         	    string  		`json:"ipaddressid"`
	IpAddress         	  		string  		`json:"ipaddress"`
	State            			string  		`json:"state"`
	CidrList      				string  		`json:"cidrlist"`
	Tags  			 	        []interface{}  	`json:"tags"`
	VmGuestIp       			string  		`json:"vmguestip"`
	NetworkId 					string  		`json:"networkid"`
	ForDisplay				    bool    		`json:"fordisplay"`
}


type CreatePortForwardingRuleResponse struct {
	Createportforwardingruleresponse struct {
		ID    string `json:"id"`  // PortForwardingRule ID
		JobId string `json:"jobid"`
	} `json:"createportforwardingruleresponse"`
}

type ListPortForwardingRulesResponse struct {
	Listportforwardingrulesresponse struct {
		PortForwardingRule []PortForwardingRule `json:"portforwardingrule"`
	} `json:"listportforwardingrulesresponse"`
}

type DeletePortForwardingRuleResponse struct {
	Deleteportforwardingruleresponse struct {
		JobId string `json:"jobid"`
	} `json:"deleteportforwardingruleresponse"`
}