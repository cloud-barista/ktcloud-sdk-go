// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This code is based on the following material: https://github.com/mindjiver/gopherstack (MIT License)
//
// KT Cloud SDK go
//
// by ETRI, 2021.07.

package ktcloudsdk

import (
	"encoding/base64"
	"net/url"
	"strings"
)

type DeployVMReqInfo struct {
	ZoneId 				string
	ServiceOfferingId 	string  // VMSpec ID  
	TemplateId 			string  // VMImage ID
	DiskOfferingId 		string
	ProductCode 		string
	VMHostName 			string
	DisplayName 		string
	UsagePlanType 		string
	RunSysPrep 			bool
	Account 			string
	DomainId 			string
	Group 				string
	Hypervisor 			string
	KeyPair 			string
	IPtoNetworkList 	[]string
	Keyboard 			string
	NetworkIds 			[]string
	ProjectId 			string
	UserData 			string
}

type ListVMReqInfo struct {
	ZoneId 				string
	VMId 				string
}

// Deploys a Virtual Machine and returns it's id
func (c KtCloudClient) DeployVirtualMachine(vmReqInfo DeployVMReqInfo) (DeployVirtualMachineResponse, error) {
	var resp DeployVirtualMachineResponse

	params := url.Values{}

	params.Set("zoneid", vmReqInfo.ZoneId)
	params.Set("serviceofferingid", vmReqInfo.ServiceOfferingId)
	params.Set("templateid", vmReqInfo.TemplateId)

	if vmReqInfo.DiskOfferingId != "" {
		params.Set("diskofferingid", vmReqInfo.DiskOfferingId)
	}

	if vmReqInfo.ProductCode != "" {
		params.Set("productcode", vmReqInfo.ProductCode)
	}

	if vmReqInfo.VMHostName != "" {
		params.Set("name", vmReqInfo.VMHostName)
	}

	if vmReqInfo.DisplayName != "" {
		params.Set("displayname", vmReqInfo.DisplayName)
	}

	if vmReqInfo.UsagePlanType != "" {
		params.Set("usageplantype", vmReqInfo.UsagePlanType)
	}

	if vmReqInfo.RunSysPrep {
		params.Set("runsysprep", "true")
	}

	if vmReqInfo.Account != "" {
		params.Set("account", vmReqInfo.Account)
	}

	if vmReqInfo.DomainId != "" {
		params.Set("domainid", vmReqInfo.DomainId)
	}

	if vmReqInfo.Group != "" {
		params.Set("group", vmReqInfo.Group)
	}

	if vmReqInfo.Hypervisor != "" {
		params.Set("hypervisor", vmReqInfo.Hypervisor)
	}

	if vmReqInfo.KeyPair != "" {
		params.Set("keypair", vmReqInfo.KeyPair)
	}

	if len(vmReqInfo.IPtoNetworkList) > 0 {
		params.Set("iptonetworklist", strings.Join(vmReqInfo.IPtoNetworkList, ","))
	}

	if vmReqInfo.Keyboard != "" {
		params.Set("keyboard", vmReqInfo.Keyboard)
	}

	if len(vmReqInfo.NetworkIds) > 0 {
		params.Set("networkids", strings.Join(vmReqInfo.NetworkIds, ","))
	}

	if vmReqInfo.UserData != "" {
		params.Set("userdata", base64.StdEncoding.EncodeToString([]byte(vmReqInfo.UserData)))
	}

	response, err := NewRequest(c, "deployVirtualMachine", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DeployVirtualMachineResponse)
	return resp, nil
}

// Start a Virtual Machine
func (c KtCloudClient) StartVirtualMachine(vmId string) (StartVirtualMachineResponse, error) {
	var resp StartVirtualMachineResponse
	params := url.Values{}
	params.Set("id", vmId)
	response, err := NewRequest(c, "startVirtualMachine", params)
	if err != nil {
		return resp, err
	}
	resp = response.(StartVirtualMachineResponse)
	return resp, err
}

// Stops a Virtual Machine
func (c KtCloudClient) StopVirtualMachine(vmId string) (StopVirtualMachineResponse, error) {
	var resp StopVirtualMachineResponse
	params := url.Values{}
	params.Set("id", vmId)
	response, err := NewRequest(c, "stopVirtualMachine", params)
	if err != nil {
		return resp, err
	}
	resp = response.(StopVirtualMachineResponse)
	return resp, err
}

// Reboot a Virtual Machine
func (c KtCloudClient) RebootVirtualMachine(vmId string) (RebootVirtualMachineResponse, error) {
	var resp RebootVirtualMachineResponse
	params := url.Values{}
	params.Set("id", vmId)
	response, err := NewRequest(c, "rebootVirtualMachine", params)
	if err != nil {
		return resp, err
	}
	resp = response.(RebootVirtualMachineResponse)
	return resp, err
}

// Destroys a Virtual Machine
func (c KtCloudClient) DestroyVirtualMachine(vmId string) (DestroyVirtualMachineResponse, error) {
	var resp DestroyVirtualMachineResponse
	params := url.Values{}
	params.Set("id", vmId)

	response, err := NewRequest(c, "destroyVirtualMachine", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DestroyVirtualMachineResponse)
	return resp, nil
}

func (c KtCloudClient) ListVirtualMachines(vmListReqInfo ListVMReqInfo) (ListVirtualMachinesResponse, error) {
	var resp ListVirtualMachinesResponse
	params := url.Values{}

	if vmListReqInfo.ZoneId != "" {
		params.Set("zoneid", vmListReqInfo.ZoneId)
	}

	if vmListReqInfo.VMId != "" {
		params.Set("id", vmListReqInfo.VMId)
	}

	response, err := NewRequest(c, "listVirtualMachines", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListVirtualMachinesResponse)
	return resp, err
}

// KT Cloud > Computing > Server Management > 'Server : 부가 정보 변경'
func (c KtCloudClient) UpdateVirtualMachine(vmId string, displayname string, haenable string) (UpdateVirtualMachineResponse, error) {
	var resp UpdateVirtualMachineResponse

	params := url.Values{}
	params.Set("id", vmId)
	params.Set("displayname", displayname)
	params.Set("haenable", haenable)

	response, err := NewRequest(c, "updateVirtualMachine", params)
	if err != nil {
		return resp, err
	}
	resp = response.(UpdateVirtualMachineResponse)
	return resp, err
}

type DeployVirtualMachineResponse struct {
	Deployvirtualmachineresponse struct {
		ID    string `json:"id"`	// Virtualmachine ID
		JobId string `json:"jobid"`
		RootId string `json:"rootid"`	// Created Root Volume ID
	} `json:"deployvirtualmachineresponse"`
}

type DestroyVirtualMachineResponse struct {
	Destroyvirtualmachineresponse struct {
		JobId string `json:"jobid"`
	} `json:"destroyvirtualmachineresponse"`
}

type StartVirtualMachineResponse struct {
	Startvirtualmachineresponse struct {
		JobId string `json:"jobid"`
	} `json:"startvirtualmachineresponse"`
}

type StopVirtualMachineResponse struct {
	Stopvirtualmachineresponse struct {
		JobId string `json:"jobid"`
	} `json:"stopvirtualmachineresponse"`
}

type RebootVirtualMachineResponse struct {
	Rebootvirtualmachineresponse struct {
		JobId string `json:"jobid"`
	} `json:"rebootvirtualmachineresponse"`
}

type Nic struct {
	ID          	string	 `json:"id"`
	NetworkId   	string	 `json:"networkid"`
	NetworkName 	string	 `json:"networkname"`
	Netmask     	string	 `json:"netmask"`
	Gateway     	string	 `json:"gateway"`
	IpAddress   	string	 `json:"ipaddress"`
	IsolationUri   	string	 `json:"isolationuri"`
	BroadcastUri   	string	 `json:"broadcasturi"`
	TrafficType 	string	 `json:"traffictype"`
	Type        	string	 `json:"type"`
	IsDefault   	bool  	 `json:"isdefault"`
	MacAddress  	string	 `json:"macaddress"`
	SecondaryIp   	string	 `json:"secondaryip"`
}

type Virtualmachine struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name"`
	DisplayName         string        `json:"displayname"`
	Account             string        `json:"account"`
	UserId         	    string        `json:"userid"`
	UserName            string        `json:"username"`
	DomainId            string        `json:"domainid"`
	Domain              string        `json:"domain"`
	Created             string        `json:"created"`
	State               string        `json:"state"`
	Haenable            bool          `json:"haenable"`
	ZoneId              string        `json:"zoneid"`
	ZoneName            string        `json:"zonename"`
	TemplateId          string        `json:"templateid"`  // VMImage ID
	TemplateName        string        `json:"templatename"`
	TemplateDisplayText string        `json:"templatedisplaytext"`
	PasswordEnabled     bool          `json:"passwordenabled"`
	ServiceOfferingId   string        `json:"serviceofferingid"`  // VMSpec ID  
	ServiceOfferingName string        `json:"serviceofferingname"`
	CpuNumber           float64       `json:"cpunumber"`
	CpuSpeed            float64       `json:"cpuspeed"`
	Memory              float64       `json:"memory"`
	GuestOsId           string        `json:"guestosid"`
	RootDeviceId        float64       `json:"rootdeviceid"`
	RootDeviceType      string        `json:"rootdevicetype"`
	SecurityGroup       []interface{} `json:"securitygroup"`
	Password      		string        `json:"password"`	// KT Cloud API로는 지원하지 않는다고함.(Blank)
	Nic                 []Nic         `json:"nic"`
	Hypervisor          string        `json:"hypervisor"`
	KeyPair             string        `json:"keypair"`	// ### Manual에는 parameter가 없으나 response 값 존재
	AffinityGroup       string        `json:"affinitygroup"`
	IsDynamicallyScalable string      `json:"isdynamicallyscalable"` // VM의 cpu 와 memory 에 대한 Scale up/down을 지원하기 위한 tools 포함 여부
	OsTypeId            string        `json:"ostypeid"`
	Tags                []interface{} `json:"tags"`
}

type ListVirtualMachinesResponse struct {
	Listvirtualmachinesresponse struct {
		Count          float64          `json:"count"`
		Virtualmachine []Virtualmachine `json:"virtualmachine"`
	} `json:"listvirtualmachinesresponse"`
}

type UpdateVirtualMachineResponse struct {
	Updatevirtualmachineresponse struct {
		Virtualmachine []Virtualmachine `json:"virtualmachine"`
	} `json:"updatevirtualmachineresponse"`
}
