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

type CreateFirewallRuleReqInfo struct {
	IpAddressId string
	Protocol    string
	CidrList    string
	// StartPort   int
	// EndPort     int
	StartPort   string
	EndPort     string
	IcmpCode    string
	IcmpType    string
	Type        string // Firewall 설정 타입으로 user 또는 system 입력(Default : user)
}

type ListFirewallRulesReqInfo struct {
	ID          string
	IpAddressId string
	Account     string
	DomainId    string
	IsRecursive bool
	Keyword     string
	Page        string
	PageSize    string
	ListAll     bool
}

func (c KtCloudClient) CreateFirewallRule(filewallRuleCreateReqInfo CreateFirewallRuleReqInfo) (CreateFirewallRuleResponse, error) {
	var resp CreateFirewallRuleResponse
	params := url.Values{}

	params.Set("ipaddressid", filewallRuleCreateReqInfo.IpAddressId)
	params.Set("protocol", filewallRuleCreateReqInfo.Protocol)

	if filewallRuleCreateReqInfo.CidrList != "" {
		params.Set("cidrlist", filewallRuleCreateReqInfo.CidrList)
	}

	params.Set("startport", filewallRuleCreateReqInfo.StartPort)

	if filewallRuleCreateReqInfo.EndPort != "" {
		params.Set("endport", filewallRuleCreateReqInfo.EndPort)
	}

	if filewallRuleCreateReqInfo.IcmpCode != "" {
		params.Set("icmpcode", filewallRuleCreateReqInfo.IcmpCode)
	}

	if filewallRuleCreateReqInfo.IcmpType != "" {
		params.Set("icmptype", filewallRuleCreateReqInfo.IcmpType)
	}

	if filewallRuleCreateReqInfo.Type != "" {
		params.Set("type", filewallRuleCreateReqInfo.Type)
	}

	response, err := NewRequest(c, "createFirewallRule", params)
	if err != nil {
		return resp, err
	}
	resp = response.(CreateFirewallRuleResponse)
	return resp, err
}

func (c KtCloudClient) ListFirewallRules(filewallRuleListReqInfo ListFirewallRulesReqInfo) (ListFirewallRulesResponse, error) {
	var resp ListFirewallRulesResponse
	params := url.Values{}

	if filewallRuleListReqInfo.ID != "" {
		params.Set("id", filewallRuleListReqInfo.ID)
	}

	if filewallRuleListReqInfo.IpAddressId != "" {
		params.Set("ipaddressid", filewallRuleListReqInfo.IpAddressId)
	}

	if filewallRuleListReqInfo.Account != "" {
		params.Set("account", filewallRuleListReqInfo.Account)
	}

	if filewallRuleListReqInfo.DomainId != "" {
		params.Set("domainid", filewallRuleListReqInfo.DomainId)
	}

	if filewallRuleListReqInfo.IsRecursive {
		params.Set("isrecursive", "true")
	}

	if filewallRuleListReqInfo.Keyword != "" {
		params.Set("keyword", filewallRuleListReqInfo.Keyword)
	}

	if filewallRuleListReqInfo.Page != "" {
		params.Set("page", filewallRuleListReqInfo.Page)
	}

	if filewallRuleListReqInfo.PageSize != "" {
		params.Set("pagesize", filewallRuleListReqInfo.PageSize)
	}

	if filewallRuleListReqInfo.ListAll {
		params.Set("listall", "true")
	}

	response, err := NewRequest(c, "listFirewallRules", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListFirewallRulesResponse)
	return resp, err
}

func (c KtCloudClient) DeleteFirewallRule(ruleId string) (DeleteFirewallRuleResponse, error) {
	var resp DeleteFirewallRuleResponse
	params := url.Values{}

	params.Set("id", ruleId) // FirewallRule ID

	response, err := NewRequest(c, "deleteFirewallRule", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DeleteFirewallRuleResponse)
	return resp, err
}

type FirewallRule struct {
	ID          string        `json:"id"`
	Protocol    string        `json:"protocol"`
	StartPort   int           `json:"startport"` // Caution!! (Parameter type)
	EndPort     int           `json:"endport"`   // Caution!! (Parameter type)
	IpAddressId string        `json:"ipaddressid"`
	NetworkId   string        `json:"networkid"`
	IpAddress   string        `json:"ipaddress"`
	State       string        `json:"state"`
	CidrList    string        `json:"cidrlist"`
	Tags        []interface{} `json:"tags"`
	ForDisplay  bool          `json:"fordisplay"`
}

type CreateFirewallRuleResponse struct {
	Createfirewallruleresponse struct {
		ID    string `json:"id"` // FirewallRule ID
		JobId string `json:"jobid"`
	} `json:"createfirewallruleresponse"`
}

type ListFirewallRulesResponse struct {
	Listfirewallrulesresponse struct {
		Count        int            `json:"count"`
		FirewallRule []FirewallRule `json:"firewallrule"`
	} `json:"listfirewallrulesresponse"`
}

type DeleteFirewallRuleResponse struct {
	Deletefirewallruleresponse struct {
		JobId string `json:"jobid"`
	} `json:"deletefirewallruleresponse"`
}
