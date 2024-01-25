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
// KT Cloud (G1/G2 Platform) Disk Volume API : https://cloud.kt.com/docs/open-api-guide/g/computing/disk-volume

type CreateVolumeReqInfo struct {
	Name             string // Required. Volume Name.
	DiskOfferingId   string // Required. DiskOfferingId
	ZoneId           string // Required
	UsagePlanType    string // default : hourly
	Account          string
	DomainId         string
	Size             string // Custom size of Volume
	SnapshotId       string
	VMId 			 string // It is only valid if the 'SnapshotId' field is present, and allows the volume to be automatically connected to the VM after creation.
	ProductCode      string // ### Create volume using product abbreviations (ex. STG 100G, SSD 300G, etc.)
	// ### If the 'ProductCode' field is used, the 'DiskOfferingId' field value is ignored.
	IOPS string
}

type ListVolumeReqInfo struct {
	Account          string
	DomainId         string
	IsRecursive      bool   // Default : false
	ID               string // Volume ID
	Keyword          string
	Name             string // Volume Name
	Page             string
	PageSize         string
	Type             string
	VMId 			 string
	ZoneId           string
	Install          bool // Default : false
}

type ResizeVolumeReqInfo struct {
	ID               string // Required. Volume ID
	VMId 			 string // Required
	Size             string // Required. Only 50(Linux series only), 80, and 100 are available.
	IsLinux          string // Required. 'Y' for Linux series, 'N' for Windows series
}

type AttachVolumeReqInfo struct {
	ID               string // Required. Volume ID
	VMId 			 string // Required
	DeviceId         string // ID of VM Device where Volume is connected. When unused, a Device ID that can be used sequentially is selected automatically.
}

type DetachVolumeReqInfo struct {
	ID               string // Required. Volume ID
	VMId 			 string
	DeviceId         string
}

// # Create a Disk Volume
// DiskOfferingId : https://cloud.kt.com/docs/open-api-guide/g/computing/disk-volume
func (c KtCloudClient) CreateVolume(req CreateVolumeReqInfo) (CreateVolumeResponse, error) {
	var resp CreateVolumeResponse
	params := url.Values{}
	params.Set("name", req.Name)
	params.Set("diskofferingid", req.DiskOfferingId)
	params.Set("zoneid", req.ZoneId)

	if req.UsagePlanType != "" {
		params.Set("usageplantype", req.UsagePlanType)
	}
	if req.Account != "" {
		params.Set("account", req.Account)
	}
	if req.DomainId != "" {
		params.Set("domainid", req.DomainId)
	}
	if req.Size != "" {
		params.Set("size", req.Size)
	}
	if req.SnapshotId != "" {
		params.Set("snapshotid", req.SnapshotId)
	}
	if req.VMId != "" {
		params.Set("virtualmachineid", req.VMId)
	}
	if req.ProductCode != "" {
		params.Set("productcode", req.ProductCode)
	}
	if req.IOPS != "" {
		params.Set("iops", req.IOPS)
	}

	response, err := NewRequest(c, "createVolume", params)
	if err != nil {
		return resp, err
	}
	resp = response.(CreateVolumeResponse)
	return resp, nil
}

// # List Disk Volumes
func (c KtCloudClient) ListVolumes(req ListVolumeReqInfo) (ListVolumesResponse, error) {
	var resp ListVolumesResponse
	params := url.Values{}

	if req.Account != "" {
		params.Set("account", req.Account)
	}
	if req.DomainId != "" {
		params.Set("domainid", req.DomainId)
	}
	if req.IsRecursive {
		params.Set("isrecursive", "true")
	}
	if req.ID != "" {
		params.Set("id", req.ID)
	}
	if req.Keyword != "" {
		params.Set("keyword", req.Keyword)
	}
	if req.Name != "" {
		params.Set("name", req.Name)
	}
	if req.Page != "" {
		params.Set("page", req.Page)
	}
	if req.PageSize != "" {
		params.Set("pagesize", req.PageSize)
	}
	if req.Type != "" {
		params.Set("type", req.Type)
	}
	if req.VMId != "" {
		params.Set("virtualmachineid", req.VMId)
	}
	if req.ZoneId != "" {
		params.Set("zoneid", req.ZoneId)
	}
	if req.Install {
		params.Set("install", "true")
	}

	response, err := NewRequest(c, "listVolumes", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ListVolumesResponse)
	return resp, nil
}

// # Resize(Change) Disk Volume Size (This is only for Bootalbe Disk of a VM.)
func (c KtCloudClient) ResizeVolume(req ResizeVolumeReqInfo) (ResizeVolumeResponse, error) {
	var resp ResizeVolumeResponse
	params := url.Values{}

	params.Set("id", req.ID) // Volume ID
	params.Set("vmid", req.VMId)
	params.Set("size", req.Size)
	params.Set("isLinux", req.IsLinux)

	response, err := NewRequest(c, "resizeVolume", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ResizeVolumeResponse)
	return resp, nil
}

// # Delete a Disk Volume
func (c KtCloudClient) DeleteVolume(id string) (DeleteVolumeResponse, error) {
	var resp DeleteVolumeResponse
	params := url.Values{}
	params.Set("id", id)
	response, err := NewRequest(c, "deleteVolume", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DeleteVolumeResponse)
	return resp, nil
}

// # Attach a Disk Volume to VM
func (c KtCloudClient) AttachVolume(req AttachVolumeReqInfo) (AttachVolumeResponse, error) {
	var resp AttachVolumeResponse
	params := url.Values{}
	
	params.Set("id", req.ID) // Volume ID
	params.Set("virtualmachineid", req.VMId) // Not 'vmid' but 'virtualmachineid'
	
	if req.DeviceId != "" {
		params.Set("deviceid", req.DeviceId)
	}

	response, err := NewRequest(c, "attachVolume", params)
	if err != nil {
		return resp, err
	}
	resp = response.(AttachVolumeResponse)
	return resp, nil
}

// # Detach a Disk Volume from VM
func (c KtCloudClient) DetachVolume(req DetachVolumeReqInfo) (DetachVolumeResponse, error) {
	var resp DetachVolumeResponse
	params := url.Values{}

	params.Set("id", req.ID) // Volume ID

	if req.VMId != "" {
		params.Set("virtualmachineid", req.VMId) // Not 'vmid' but 'virtualmachineid'
	}

	if req.DeviceId != "" {
		params.Set("deviceid", req.DeviceId)
	}

	response, err := NewRequest(c, "detachVolume", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DetachVolumeResponse)
	return resp, nil
}

type Volume struct {
	ID                         	string   `json:"id"`
	Name                       	string   `json:"name"`
	ZoneId                     	string   `json:"zoneid"`
	ZoneName                   	string   `json:"zonename"`
	Type                       	string   `json:"type"`
	DeviceId                   	int      `json:"deviceid"`
	VMId           			   	string   `json:"virtualmachineid"`
	VMName                     	string   `json:"vmname"`
	VMDisplayName              	string   `json:"vmdisplayname"`
	VMState                    	string   `json:"vmstate"`
	TemplateId                 	string   `json:"templateid"`   			// Volume with the OS installed
	TemplateName               	string   `json:"templatename"`  		// Volume with the OS installed
	TemplateDisplayText        	string   `json:"templatedisplaytext"`	// Volume with the OS installed
	ProvisioningType           	string   `json:"provisioningtype"`
	Size                       	int64    `json:"size"`
	MinIOPS                    	int64    `json:"miniops"`
	MaxIOPS                    	int64    `json:"maxiops"`
	Created                    	string   `json:"created"`
	State                      	string   `json:"state"`
	Account                    	string   `json:"account"`
	DomainId                   	string   `json:"domainid"`
	Domain                     	string   `json:"domain"`
	StorageType                	string   `json:"storagetype"`
	DiskOfferingId				string   `json:"diskofferingid"`
	DiskOfferingName        	string   `json:"diskofferingname"`
	DiskOfferingDisplayText 	string   `json:"diskofferingdisplaytext"`
	ServiceofferingId			string   `json:"serviceofferingid"`				// In case, Type value : ROOT
	ServiceofferingName        	string   `json:"serviceofferingname"`			// In case, Type value : ROOT
	ServiceofferingDisplayText 	string   `json:"serviceofferingdisplaytext"`	// In case, Type value : ROOT	
	AttachedTime               	string   `json:"attached"`
	Destroyed                  	bool     `json:"destroyed"`
	IsExtractable              	bool     `json:"isextractable"`
	QuiesceVM                  	bool     `json:"quiescevm"`
	Tags                       	[]string `json:"tags"`
	UsagePlanType              	string   `json:"usageplantype"`
	VolumeType                 	string   `json:"volumetype"`
}

type CreateVolumeResponse struct {
	Createvolumeresponse struct {
		Volume Volume `json:"volume"`
		JobId  string `json:"jobid"`
		ID     string `json:"id"`
	} `json:"createvolumeresponse"`
}

type ListVolumesResponse struct {
	Listvolumesresponse struct {
		Volume []Volume `json:"volume"`
		Count  int      `json:"count"`
	} `json:"listvolumesresponse"`
}

type ResizeVolumeResponse struct {
	Resizevolumeresponse struct {
		JobId string `json:"jobid"`
	} `json:"resizevolumeresponse"`
}

type DeleteVolumeResponse struct {
	Deletevolumeresponse struct {
		Success string `json:"success"` // 'string' type of value!! 'true' or 'false'
	} `json:"deletevolumeresponse"`
}

type AttachVolumeResponse struct {
	Attachvolumeresponse struct {
		JobId string `json:"jobid"`
	} `json:"attachvolumeresponse"`
}

type DetachVolumeResponse struct {
	Detachvolumeresponse struct {
		JobId string `json:"jobid"`
	} `json:"detachvolumeresponse"`
}
