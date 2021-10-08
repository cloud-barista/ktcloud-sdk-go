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
	"net/url"
)

// Create a SSH key pair
func (c KtCloudClient) CreateSSHKeyPair(name string) (CreateSshKeyPairResponse, error) {
	var resp CreateSshKeyPairResponse
	params := url.Values{}
	params.Set("name", name)

	response, err := NewRequest(c, "createSSHKeyPair", params)
	if err != nil {
		return resp, err
	}
	resp = response.(CreateSshKeyPairResponse)
	return resp, nil
}

// List SSH keypairs
func (c KtCloudClient) ListSSHKeyPairs(name string) (ListSshKeyPairsResponse, error) {
	var resp ListSshKeyPairsResponse
	params := url.Values{}

	if name != "" {
		params.Set("name", name)
	}
	
	response, err := NewRequest(c, "listSSHKeyPairs", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListSshKeyPairsResponse)
	return resp, nil
}
//~~~ keypairs res... 에서 중간 s자 주의

// Deletes an SSH key pair
func (c KtCloudClient) DeleteSSHKeyPair(name string) (DeleteSshKeyPairResponse, error) {
	var resp DeleteSshKeyPairResponse
	params := url.Values{}
	params.Set("name", name)
	response, err := NewRequest(c, "deleteSSHKeyPair", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DeleteSshKeyPairResponse)
	return resp, err
}

type KeyPair struct {
	PrivateKey     string 		`json:"privatekey"`
	// CreateSSHKeyPair() 할때만 response로 받음.

	Name           string 		`json:"name"`
	Account        string 		`json:"account"`
	DomainId       string 		`json:"domainid"`
	Domain         string 		`json:"domain"`
	Fingerprint    string 		`json:"fingerprint"`
}

type CreateSshKeyPairResponse struct {
	Createsshkeypairresponse struct {
		KeyPair    KeyPair  	`json:"keypair"`
	} `json:"createsshkeypairresponse"`
}

type ListSshKeyPairsResponse struct {
	Listsshkeypairsresponse struct {
		Count      	   int      	`json:"count"`
		KeyPair 	   []KeyPair 	`json:"sshkeypair"`		
	} `json:"listsshkeypairsresponse"`
}
//~~~ keypair가 아닌 sshkeypair로 받음에 주의
//~~~ keypairs res... 에서 중간 s자 주의

type DeleteSshKeyPairResponse struct {
	Deletesshkeypairresponse struct {
		Success 		string 		`json:"success"`
	} `json:"deletesshkeypairresponse"`
}
