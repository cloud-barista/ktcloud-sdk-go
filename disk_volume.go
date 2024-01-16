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
func (c KtCloudClient) CreateVolume(volumeCreateReqInfo CreateVolumeReqInfo) (CreateVolumeResponse, error) {
	var resp CreateVolumeResponse
	params := url.Values{}
	params.Set("name", volumeCreateReqInfo.Name)
	params.Set("diskofferingid", volumeCreateReqInfo.DiskOfferingId)
	params.Set("zoneid", volumeCreateReqInfo.ZoneId)

	if volumeCreateReqInfo.UsagePlanType != "" {
		params.Set("usageplantype", volumeCreateReqInfo.UsagePlanType)
	}

	if volumeCreateReqInfo.Account != "" {
		params.Set("account", volumeCreateReqInfo.Account)
	}

	if volumeCreateReqInfo.DomainId != "" {
		params.Set("domainid", volumeCreateReqInfo.DomainId)
	}

	if volumeCreateReqInfo.Size != "" {
		params.Set("size", volumeCreateReqInfo.Size)
	}

	if volumeCreateReqInfo.SnapshotId != "" {
		params.Set("snapshotid", volumeCreateReqInfo.SnapshotId)
	}

	if volumeCreateReqInfo.VMId != "" {
		params.Set("virtualmachineid", volumeCreateReqInfo.VMId)
	}

	if volumeCreateReqInfo.ProductCode != "" {
		params.Set("productcode", volumeCreateReqInfo.ProductCode)
	}

	if volumeCreateReqInfo.IOPS != "" {
		params.Set("iops", volumeCreateReqInfo.IOPS)
	}

	response, err := NewRequest(c, "createVolume", params)
	if err != nil {
		return resp, err
	}
	resp = response.(CreateVolumeResponse)
	return resp, nil
}

// # List Disk Volumes
func (c KtCloudClient) ListVolumes(volumeListReqInfo ListVolumeReqInfo) (ListVolumesResponse, error) {
	var resp ListVolumesResponse
	params := url.Values{}

	if volumeListReqInfo.Account != "" {
		params.Set("account", volumeListReqInfo.Account)
	}

	if volumeListReqInfo.DomainId != "" {
		params.Set("domainid", volumeListReqInfo.DomainId)
	}

	if volumeListReqInfo.IsRecursive {
		params.Set("isrecursive", "true")
	}

	if volumeListReqInfo.ID != "" {
		params.Set("id", volumeListReqInfo.ID)
	}

	if volumeListReqInfo.Keyword != "" {
		params.Set("keyword", volumeListReqInfo.Keyword)
	}

	if volumeListReqInfo.Name != "" {
		params.Set("name", volumeListReqInfo.Name)
	}

	if volumeListReqInfo.Page != "" {
		params.Set("page", volumeListReqInfo.Page)
	}

	if volumeListReqInfo.PageSize != "" {
		params.Set("pagesize", volumeListReqInfo.PageSize)
	}

	if volumeListReqInfo.Type != "" {
		params.Set("type", volumeListReqInfo.Type)
	}

	if volumeListReqInfo.VMId != "" {
		params.Set("virtualmachineid", volumeListReqInfo.VMId)
	}

	if volumeListReqInfo.ZoneId != "" {
		params.Set("zoneid", volumeListReqInfo.ZoneId)
	}

	if volumeListReqInfo.Install {
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
func (c KtCloudClient) ResizeVolume(volumeResizeReqInfo ResizeVolumeReqInfo) (ResizeVolumeResponse, error) {
	var resp ResizeVolumeResponse
	params := url.Values{}

	params.Set("id", volumeResizeReqInfo.ID) // Volume ID
	params.Set("vmid", volumeResizeReqInfo.VMId)
	params.Set("size", volumeResizeReqInfo.Size)
	params.Set("isLinux", volumeResizeReqInfo.IsLinux)

	response, err := NewRequest(c, "resizeVolume", params)
	if err != nil {
		return resp, err
	}
	resp = response.(ResizeVolumeResponse)
	return resp, err
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
	return resp, err
}

// # Attach a Disk Volume to VM
func (c KtCloudClient) AttachVolume(volumeAttachReqInfo AttachVolumeReqInfo) (AttachVolumeResponse, error) {
	var resp AttachVolumeResponse
	params := url.Values{}
	
	params.Set("id", volumeAttachReqInfo.ID) // Volume ID
	params.Set("virtualmachineid", volumeAttachReqInfo.VMId) // Not 'vmid' but 'virtualmachineid'
	
	if volumeAttachReqInfo.DeviceId != "" {
		params.Set("deviceid", volumeAttachReqInfo.DeviceId)
	}

	response, err := NewRequest(c, "attachVolume", params)
	if err != nil {
		return resp, err
	}
	resp = response.(AttachVolumeResponse)
	return resp, err
}

// # Detach a Disk Volume from VM
func (c KtCloudClient) DetachVolume(volumeDetachReqInfo DetachVolumeReqInfo) (DetachVolumeResponse, error) {
	var resp DetachVolumeResponse
	params := url.Values{}

	params.Set("id", volumeDetachReqInfo.ID) // Volume ID

	if volumeDetachReqInfo.VMId != "" {
		params.Set("virtualmachineid", volumeDetachReqInfo.VMId) // Not 'vmid' but 'virtualmachineid'
	}

	if volumeDetachReqInfo.DeviceId != "" {
		params.Set("deviceid", volumeDetachReqInfo.DeviceId)
	}

	response, err := NewRequest(c, "detachVolume", params)
	if err != nil {
		return resp, err
	}
	resp = response.(DetachVolumeResponse)
	return resp, err
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
		Success string `json:"success"` // Successful deletion ('true' / 'false')
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
