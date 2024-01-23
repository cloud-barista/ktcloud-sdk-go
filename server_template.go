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
// KT Cloud (G1/G2 Platform) Image-Snapshot API : https://cloud.kt.com/docs/open-api-guide/g/computing/image-snapshot

type CreateTemplateReqInfo struct {
	Name             string `json:"name"`				// Required
	DisplayText      string `json:"displaytext"`		// Required
	OsTypeId         string `json:"ostypeid"`			// Required
	VolumeId         string `json:"volumeid"`			// Required
	SnapshotId       string `json:"snapshotid"`
	VMId 			 string `json:"virtualmachineid"`
	Bits             string	`json:"bits"`
	IsFeatured       bool   `json:"isfeatured"`
	IsPublic         bool   `json:"ispublic"`
	RequiresHVM      bool   `json:"requireshvm"`
}

type ListTemplateReqInfo struct {
	TemplateFilter	 string // 'self' : image created by the user. 'selfexecutable' : created by the user and currently available.
	Account          string // Accounts with generated Image belong
	DomainId         string
	IsRecursive      bool   // Used with "DomainId" field, in case of 'true', all account inquiry included in the domain (Default: false)
	Keyword          string
	Name             string // Volume Name
	Page             string
	PageSize         string
	Install          bool
		// false : Only inquiry of resources belonging to the account (Default : false)
		// true: Inquiry of all resources lists that the account can list
}

// # Create a Image Template (Server Image) of a VM
func (c KtCloudClient) CreateTemplate(req *CreateTemplateReqInfo) (CreateTemplateResponse, error) {
	var resp CreateTemplateResponse
	params := url.Values{}

	params.Set("name", req.Name)
	params.Set("displaytext", req.DisplayText)
	params.Set("ostypeid", req.OsTypeId)
	params.Set("volumeid", req.VolumeId)

	if req.SnapshotId != "" {
		params.Set("snapshotid", req.SnapshotId)
	}
	if req.SnapshotId != "" {
		params.Set("virtualmachineid", req.VMId)
	}
	if req.SnapshotId != "" {
		params.Set("bits", req.Bits)
	}
	if req.IsFeatured {
		params.Set("isfeatured", "true")
	}
	if req.IsPublic {
		params.Set("ispublic", "true")
	}
	if req.RequiresHVM {
		params.Set("requireshvm", "true")
	}

	response, err := NewRequest(c, "createTemplate", params) // Request Command according to KT Cloud API doc.
	if err != nil {
		return resp, err
	}
	resp = response.(CreateTemplateResponse)
	return resp, nil
}
// (Note) The 'queryasyncjobresultresponse' processing method is the same as in 'DeployVirtualMachine()'.

// # List Available Image Templates
func (c KtCloudClient) ListTemplates(req *ListTemplateReqInfo) (ListTemplatesResponse, error) {
	var resp ListTemplatesResponse
	params := url.Values{}

	params.Set("templatefilter", req.TemplateFilter)

	if req.Account != "" {
		params.Set("account", req.Account)
	}
	if req.DomainId != "" {
		params.Set("domainid", req.DomainId)
	}
	if req.IsRecursive {
		params.Set("isrecursive", "true")
	}
	if req.Keyword != "" {
		params.Set("keyword", req.Keyword)
	}
	if req.Page != "" {
		params.Set("page", req.Page)
	}
	if req.PageSize != "" {
		params.Set("pagesize", req.PageSize)
	}
	if req.Install {
		params.Set("install", "true")
	}

	response, err := NewRequest(c, "listTemplates", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListTemplatesResponse)
	return resp, nil
}

// # Delete a Image Template
func (c KtCloudClient) DeleteTemplate(id string, zoneId string) (DeleteTemplateResponse, error) {
	var resp DeleteTemplateResponse
	params := url.Values{}
	
	params.Set("id", id)
	params.Set("zoneid", zoneId)

	response, err := NewRequest(c, "deleteTemplate", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DeleteTemplateResponse)
	return resp, nil
}

type Template struct {
	ID                    string            `json:"id"`
	Name                  string            `json:"name"`
	DisplayText           string            `json:"displaytext"`
	IsPublic              bool              `json:"ispublic"`
	Created               string            `json:"created"`
	IsReady               bool              `json:"isready"`
	PasswordEnabled       bool              `json:"passwordenabled"`
	Format                string            `json:"format"`
	IsFeatured            bool              `json:"isfeatured"`
	CrossZones            bool              `json:"crossZones"`
	OSTypeId              string            `json:"ostypeid"`
	OSTypeName            string            `json:"ostypename"`
	Account               string            `json:"account"`
	ZoneId                string            `json:"zoneid"`
	ZoneName              string            `json:"zonename"`
	Status                string            `json:"status"`
	Size                  int	            `json:"size"`
	PhysicalSize          int   	        `json:"physicalsize"`
	TemplateType          string            `json:"templatetype"`
	Hypervisor            string            `json:"hypervisor"`
	Domain                string            `json:"domain"`
	DomainId              string            `json:"domainid"`
	IsExtractable         bool              `json:"isextractable"`
	SourceTemplateId      string            `json:"sourcetemplateid"`
	Details               map[string]string `json:"details"`
	Bits                  int               `json:"bits"`
	SshKeyEnabled         bool              `json:"sshkeyenabled"`
	IsDynamicallyScalable bool              `json:"isdynamicallyscalable"`
	CacheInPrimary        bool              `json:"cacheinprimary"`
	Tags                  []interface{}     `json:"tags"`
}

type CreateTemplateResponse struct {
	Createtemplateresponse struct {
		ID    string `json:"id"`
		JobId string `json:"jobid"`
	} `json:"createtemplateresponse"`
}

type ListTemplatesResponse struct {
	Listtemplatesresponse struct {
		Count    int    	`json:"count"`
		Template []Template `json:"template"`
	} `json:"listtemplatesresponse"`
}

type DeleteTemplateResponse struct {
	Deletetemplateresponse struct {
		JobId string `json:"jobid"`
	} `json:"deletetemplateresponse"`
}
