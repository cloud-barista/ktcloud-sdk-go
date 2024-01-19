// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This code is based on the following material: https://github.com/mindjiver/gopherstack (MIT License)
//
// KT Cloud SDK go
//
// by ETRI, 2024.01.

package ktcloudsdk

import (
	"net/url"
)
// KT Cloud (G1/G2 Platform) Load-Balancer API : https://cloud.kt.com/docs/open-api-guide/g/network/load-balancer

type CreateNLBReqInfo struct {
	Name             string `json:"name"`				// Required
	ZoneId           string `json:"zoneid"`				// Required. Zone ID that has the 'ServiceIP'
	NLBOption 		 string `json:"loadbalanceroption"`	// Required. roundrobin / leastconnection / leastresponse / sourceiphash / srcipsrcporthash
	ServiceIP        string `json:"serviceip"`			// Required. KT Cloud Virtual IP. $$$ In case of an empty value(""), it is newly created.
	ServicePort      string `json:"serviceport"`		// Required
	ServiceType      string `json:"servicetype"`		// Required. NLB ServiceType : https / http / sslbridge / tcp / ftp
	HealthCheckType  string `json:"healthchecktype"`	// Required. HealthCheckType : http / https / tcp
	HealthCheckURL   string `json:"healthcheckurl"`		// Required. URL when the HealthCheckType is 'http' or 'https'.
	CipherGroupName  string `json:"ciphergroupname"`	// Required when ServiceType is 'https'. Set CipherGroup Name
	SSLv3        	 string `json:"sslv3"`				// Required when ServiceType is 'https'. Use SSLv3? : 'DISABLED' / 'ENABLED'
	TLSv1        	 string `json:"tlsv1"`				// Required when ServiceType is 'https'. Use TLSv1? : 'DISABLED' / 'ENABLED'
	TLSv11         	 string `json:"tlsv11"`				// Required when ServiceType is 'https'. Use TLSv11? : 'DISABLED' / 'ENABLED'
	TLSv12        	 string `json:"tlsv12"`				// Required when ServiceType is 'https'. Use TLSv12? : 'DISABLED' / 'ENABLED'
	NetworkId        string `json:"networkid"` 			// Tier Network ID. Required in case of 'Enterprise Security'
}

type ListNLBsReqInfo struct {
	Name             string `json:"name"`				// NLB Name
	ZoneId           string `json:"zoneid"`
	ServiceIP        string `json:"serviceip"`
	NLBId 		 	 string `json:"loadbalancerid"`
	MemId            string `json:"memid"` 				// User account ID
}

type AddNLBVMReqInfo struct {
	NLBId 		 	 string `json:"loadbalancerid"`		// Required
	VMId 			 string `json:"virtualmachineid"`	// Required
	IpAddress   	 string	`json:"ipaddress"`			// Required. 'Public IP' of VM to be added to NLB
	PublicPort   	 string	`json:"publicport"`			// Required. Port of VM to be added
}

// # Create a Load-Balancer 
func (c KtCloudClient) CreateNLB(req CreateNLBReqInfo) (CreateNLBResponse, error) {
	var resp CreateNLBResponse	
	params := url.Values{}

	params.Add("name", req.Name)
	params.Add("zoneid", req.ZoneId)
	params.Add("loadbalanceroption", req.NLBOption)
	params.Add("serviceip", req.ServiceIP)
	params.Add("serviceport", req.ServicePort)
	params.Add("servicetype", req.ServiceType)
	params.Add("healthchecktype", req.HealthCheckType)

	if req.HealthCheckURL != "" {
		params.Add("healthcheckurl", req.HealthCheckURL)
	}
	if req.CipherGroupName != "" {
		params.Add("ciphergroupname", req.CipherGroupName)
	}		
	if req.SSLv3 != "" {
		params.Add("sslv3", req.SSLv3)
	}
	if req.TLSv1 != "" {
		params.Add("tlsv1", req.TLSv1)
	}
	if req.TLSv11 != "" {
		params.Add("tlsv11", req.TLSv11)
	}
	if req.TLSv12 != "" {
		params.Add("tlsv12", req.TLSv12)
	}
	if req.NetworkId != "" {
		params.Add("networkid", req.NetworkId)
	}

	response, err := NewRequest(c, "createLoadBalancer", params) // Request Command according to KT Cloud API doc.
	if err != nil {
		return resp, err
	}
	resp = response.(CreateNLBResponse)
	return resp, nil
}

// # List Load-Balancers 
func (c KtCloudClient) ListNLBs(req ListNLBsReqInfo) (ListNLBsResponse, error) {
	var resp ListNLBsResponse	
	params := url.Values{}

	if req.Name != "" {
		params.Add("name", req.Name)
	}
	if req.ZoneId != "" {
		params.Add("zoneid", req.ZoneId)
	}		
	if req.ServiceIP != "" {
		params.Add("serviceip", req.ServiceIP)
	}
	if req.NLBId != "" {
		params.Add("loadbalancerid", req.NLBId)
	}
	if req.MemId != "" {
		params.Add("memid", req.MemId)
	}

	response, err := NewRequest(c, "listLoadBalancers", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListNLBsResponse)
	return resp, nil
}

// # Delete a Load-Balancer
func (c KtCloudClient) DeleteNLB(nlbId string) (DeleteNLBResponse, error) {
	var resp DeleteNLBResponse
	params := url.Values{}

	params.Set("loadbalancerid", nlbId)

	response, err := NewRequest(c, "deleteLoadBalancer", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DeleteNLBResponse)
	return resp, err
}

// # Ad a VM to Load-Balancer
func (c KtCloudClient) AddNLBVM(req AddNLBVMReqInfo) (AddNLBVMResponse, error) {
	var resp AddNLBVMResponse
	params := url.Values{}
	
	params.Set("loadbalancerid", req.NLBId)
	params.Set("virtualmachineid", req.VMId) // Not 'vmid' but 'virtualmachineid'
	params.Set("ipaddress", req.IpAddress)
	params.Set("publicport", req.PublicPort)

	response, err := NewRequest(c, "addLoadBalancerWebServer", params)
	if err != nil {
		return resp, err
	}
	resp = response.(AddNLBVMResponse)
	return resp, err
}

// # List Load-Balancers 
func (c KtCloudClient) ListNLBVMs(nlbId string) (ListNLBVMsResponse, error) {
	var resp ListNLBVMsResponse	
	params := url.Values{}

	params.Set("loadbalancerid", nlbId)

	response, err := NewRequest(c, "listLoadBalancerWebServers", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListNLBVMsResponse)
	return resp, nil
}

// # Delete a Load-Balancer
func (c KtCloudClient) RemoveNLBVM(serviceId string) (RemoveNLBVMResponse, error) {
	var resp RemoveNLBVMResponse
	params := url.Values{}

	params.Set("serviceid", serviceId)

	response, err := NewRequest(c, "removeLoadBalancerWebServer", params)
	if err != nil {
		return resp, err
	}
	resp = response.(RemoveNLBVMResponse)
	return resp, err
}

type NLB struct {
    CertificateName     string `json:"certificatename"`
    CipherGroupName     string `json:"cipherGroupName"`
    ClientIpYn          string `json:"clientIpYn"`
    EstablishedConn     string `json:"establishedconn"`
    HealthCheckType     string `json:"healthchecktype"` // Health CheckType : http / https / tcp
    HealthCheckURL      string `json:"healthcheckurl"`
    NLBId      			int    `json:"loadbalancerid"`
    NLBOption  			string `json:"loadbalanceroption"`
    Name                string `json:"name"`
    NetworkId           string `json:"networkid"`
    RequestsRate        int    `json:"requestsrate"`
    ServiceIP           string `json:"serviceip"`
    ServicePort         string `json:"serviceport"`
    ServiceType         string `json:"servicetype"`	// NLB Service Type : https / http / sslbridge / tcp / ftp
    Sslv2               string `json:"sslv2"`
    Sslv3               string `json:"sslv3"`
    State               string `json:"state"`
    Tag                 string `json:"tag"`
    Tlsv1               string `json:"tlsv1"`
    Tlsv11              string `json:"tlsv11"`
    Tlsv12              string `json:"tlsv12"`
    ZoneId              string `json:"zoneid"`
    ZoneName            string `json:"zonename"`
}

type CreateNLBResponse struct {
	Createnlbresponse struct {
		NLBId    			string `json:"loadbalancerid"`
		ZoneId            	string `json:"zoneid"`
		ZoneName          	string `json:"zonename"`
		ServiceIP         	string `json:"serviceip"`
		ServicePort       	string `json:"serviceport"`
		ServiceType       	string `json:"servicetype"`
		Name              	string `json:"name"`
		NLBOption 			string `json:"loadbalanceroption"`
		HealthCheckType   	string `json:"healthchecktype"`
		HealthCheckURL    	string `json:"healthcheckurl"`
		ErrorCode    		string `json:"errorcode"`
		ErrorText    		string `json:"errortext"`		
	} `json:"createLoadBalancerresponse"`
}

type ListNLBsResponse struct {   // Note) Plural
	Listnlbsresponse struct {	 // Note) Plural
		Count       		int 	`json:"count"`
		NLB 				[]NLB 	`json:"loadbalancer"`
	} `json:"listloadbalancersresponse"`
}

type DeleteNLBResponse struct {
	Deletenlbresponse struct {
		Success 			bool	`json:"success"` // 'bool' type of value (Not like DeleteVolumeResponse)
		Displaytext 		string 	`json:"displaytext"`
	} `json:"deleteloadbalancerresponse"`
}

type AddNLBVMResponse struct {
	Addnlbvmresponse struct {
		ServiceId         	int    `json:"serviceid"`
		NLBId    			string `json:"loadbalancerid"`
		VMId 			 	string `json:"virtualmachineid"`
		IpAddress   	 	string `json:"ipaddress"`	// Public IP of VM added
		PublicPort   	 	string `json:"publicport"`	// Port of VM added
	} `json:"addLoadBalancerWebServerresponse"`
}

type NLBVM struct {
	NLBId        		int 	`json:"loadbalancerid"`
	ServiceId           int 	`json:"serviceid"`
	VMId      			string `json:"virtualmachineid"`
	IPAddress           string `json:"ipaddress"`
	PublicPort          string `json:"publicport"`
	CurVMConnections    int    `json:"cursrvrconnections"`
	State               string `json:"state"`
	ThroughputRate      int    `json:"throughputrate"` // Throughput (Mbps)
	AvgSvrttfb          int    `json:"avgsvrttfb"`	   // TTFB (VM response time, unit: msec)
	RequestsRate        int    `json:"requestsrate"`
}

type ListNLBVMsResponse struct { // Note) Plural
	Listnlbvmsresponse struct {  // Note) Plural
		NLBVM 				[]NLBVM `json:"loadbalancerwebserver"`
		Count       		int 	`json:"count"`
	} `json:"listLoadBalancerWebServersresponse"`
}

type RemoveNLBVMResponse struct {
	Removenlbvmresponse struct {
		Success 			bool	`json:"success"` // 'bool' type of value (Not like DeleteVolumeResponse)
		Displaytext 		string 	`json:"displaytext"`
	} `json:"removeLoadbalancerWebServerresponse"`
}
