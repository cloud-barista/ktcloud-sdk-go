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
	"strconv"
	"strings"
)

type CreateTags struct {
	ResourceIds  []string `json:"resourceids"`
	ResourceType string   `json:"resourcetype`
	Tags         []TagArg `json:"tags`
}

type ListTags struct {
	Account      string `json:"account"`
	DomainId     string `json:"domainid"`
	Key          string `json:"key"`
	Value        string `json:"value`
	ResourceIds  string `json:"resourceids"`
	ResourceType string `json:"resourcetype`
}

type DeleteTags struct {
	ResourceIds  []string `json:"resourceids"`
	ResourceType string   `json:"resourcetype`
	Tags         []TagArg `json:"tags`
}

type TagArg struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Add tags to specified resources
func (c KtCloudClient) CreateTags(options *CreateTags) (CreateTagsResponse, error) {
	var resp CreateTagsResponse
	params := url.Values{}

	params.Set("resourceids", strings.Join(options.ResourceIds, ","))
	params.Set("resourcetype", options.ResourceType)
	for j, tag := range options.Tags {
		params.Set("tags["+strconv.Itoa(j+1)+"].key", tag.Key)
		params.Set("tags["+strconv.Itoa(j+1)+"].value", tag.Value)
	}

	response, err := NewRequest(c, "createTags", params)
	if err != nil {
		return resp, err
	}

	resp = response.(CreateTagsResponse)
	return resp, err
}

// Returns all items with a particular tag
func (c KtCloudClient) ListTags(options *ListTags) (ListTagsResponse, error) {
	var resp ListTagsResponse
	params := url.Values{}

	if options.Account != "" {
		params.Set("account", options.Account)
	}

	if options.DomainId != "" {
		params.Set("domainid", options.DomainId)
	}

	if options.Key != "" {
		params.Set("key", options.Key)
	}

	if options.Value != "" {
		params.Set("value", options.Value)
	}

	if options.ResourceIds != "" {
		params.Set("resourceid", options.ResourceIds)
	}

	if options.ResourceType != "" {
		params.Set("resourcetype", options.ResourceType)
	}

	response, err := NewRequest(c, "listTags", params)
	if err != nil {
		return resp, err
	}

	resp = response.(ListTagsResponse)
	return resp, err
}

// Remove tags from specified resources
func (c KtCloudClient) DeleteTags(options *DeleteTags) (DeleteTagsResponse, error) {
	var resp DeleteTagsResponse
	params := url.Values{}

	params.Set("resourceids", strings.Join(options.ResourceIds, ","))
	params.Set("resourcetype", options.ResourceType)
	
	for j, tag := range options.Tags {
		params.Set("tags["+strconv.Itoa(j+1)+"].key", tag.Key)
		params.Set("tags["+strconv.Itoa(j+1)+"].value", tag.Value)
	}

	response, err := NewRequest(c, "deleteTags", params)
	if err != nil {
		return resp, err
	}

	resp = response.(DeleteTagsResponse)
	return resp, err
}

type Tag struct {
	Key          string `json:"key"`
	Value        string `json:"value`
	ResourceType string `json:"resourcetype`
	ResourceId   string `json:"resourceid"`
	Account      string `json:"account"`
	DomainId     string `json:"domainid"`
	Domain       string `json:"domain"`
}

type CreateTagsResponse struct {
	Createtagsresponse struct {
		JobId string `json:"jobid"`
	} `json:"createtagsresponse"`
}

type ListTagsResponse struct {
	Listtagsresponse struct {
		Count int `json:"count"`
		Tag   []Tag   `json:"tag"`
	} `json:"listtagsresponse"`
}

type DeleteTagsResponse struct {
	Deletetagsresponse struct {
		JobId string `json:"jobid"`
	} `json:"deletetagsresponse"`
}
