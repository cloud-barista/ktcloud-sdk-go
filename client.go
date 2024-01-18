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
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type KtCloudClient struct {
	// The http client for communicating
	client *http.Client

	// The base URL of the API
	BaseURL string

	// Credentials
	APIKey    string
	SecretKey string
}

// Creates a new client for communicating with KT Cloud
func (ktcloud KtCloudClient) New(apiUrl string, apiKey string, secretKey string, insecureSkipVerify bool) *KtCloudClient {
	c := &KtCloudClient{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
				Proxy:           http.ProxyFromEnvironment,
			},
		},
		BaseURL:   apiUrl,
		APIKey:    apiKey,
		SecretKey: secretKey,
	}
	return c
}

func NewRequest(c KtCloudClient, request string, params url.Values) (interface{}, error) {
	client := c.client

	params.Set("apikey", c.APIKey)
	params.Set("command", request)
	params.Set("response", "json")

	// Generate signature for API call
	// * Serialize parameters and sort them by key, done by Encode()
	// * Use byte sequence for '+' character as KT Cloud requires this
	// * For the signature only, un-encode [ and ].
	// * Convert the entire argument string to lowercase
	// * Calculate HMAC SHA1 of argument string with KT Cloud secret key
	// * URL encode the string and convert to base64
	s := params.Encode()
	s2 := strings.Replace(s, "+", "%20", -1)
	s3 := strings.ToLower(strings.Replace(strings.Replace(s2, "%5B", "[", -1), "%5D", "]", -1))
	mac := hmac.New(sha1.New, []byte(c.SecretKey))
	mac.Write([]byte(s3))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	signature = url.QueryEscape(signature)

	// Create the final URL before we issue the request
	// For some reason KT Cloud refuses to accept '+' as a space character so we byte escape it instead.
	url := c.BaseURL + "?" + s2 + "&signature=" + signature
	// log.Printf("\n\n### Request URI : %s\n\n", url)	// For Testing

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	// log.Printf("\n\n### Response from KT Cloud: %d, %s\n\n", resp.StatusCode, body)	// For Testing
	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Received HTTP client/server error from KT Cloud: %d - %s", resp.StatusCode, body))
		return nil, err
	}

	switch request {

	default:
		log.Printf("Unknown request %s", request)

	// SSH Key
	case "createSSHKeyPair": // Request Command according to KT Cloud API doc.
		var decodedResponse CreateSshKeyPairResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listSSHKeyPairs":
		var decodedResponse ListSshKeyPairsResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil
		// Caution!!) When list ~, ~ KeyPair's'

	case "deleteSSHKeyPair":
		var decodedResponse DeleteSshKeyPairResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// Virtual Machine (Server)
	case "deployVirtualMachine":
		var decodedResponse DeployVirtualMachineResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "destroyVirtualMachine":
		var decodedResponse DestroyVirtualMachineResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "startVirtualMachine":
		var decodedResponse StartVirtualMachineResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "stopVirtualMachine":
		var decodedResponse StopVirtualMachineResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "rebootVirtualMachine":
		var decodedResponse RebootVirtualMachineResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listVirtualMachines":
		var decodedResponse ListVirtualMachinesResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// Product Type
	case "listAvailableProductTypes":
		var decodedResponse ListAvailableProductTypesResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// AsyncJob
	case "queryAsyncJobResult":
		var decodedResponse QueryAsyncJobResultResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// Tag
	case "createTags":
		var decodedResponse CreateTagsResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listTags":
		var decodedResponse ListTagsResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "deleteTags":
		var decodedResponse DeleteTagsResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// Zone
	case "listZones":
		var decodedResponse ListZonesResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// Firewallrule
	case "createFirewallRule":
		var decodedResponse CreateFirewallRuleResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listFirewallRules":
		var decodedResponse ListFirewallRulesResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "deleteFirewallRule":
		var decodedResponse DeleteFirewallRuleResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// Public IP
	case "associateIpAddress":
		var decodedResponse AssociateIpAddressResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listPublicIpAddresses":
		var decodedResponse ListPublicIpAddressesResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "disassociateIpAddress":
		var decodedResponse DisassociateIpAddressResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// PortForwarding Rule
	case "createPortForwardingRule":
		var decodedResponse CreatePortForwardingRuleResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listPortForwardingRules":
		var decodedResponse ListPortForwardingRulesResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "deletePortForwardingRule":
		var decodedResponse DeletePortForwardingRuleResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// Disk Volume
	case "createVolume":
		var decodedResponse CreateVolumeResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listVolumes":
		var decodedResponse ListVolumesResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil
	
	case "resizeVolume":
		var decodedResponse ResizeVolumeResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "deleteVolume":
		var decodedResponse DeleteVolumeResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "attachVolume":
		var decodedResponse AttachVolumeResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "detachVolume":
		var decodedResponse DetachVolumeResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	// Load Balancer
	case "createLoadBalancer": // Request Command according to KT Cloud API doc.
		var decodedResponse CreateNLBResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listLoadBalancers":
		var decodedResponse ListNLBsResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "deleteLoadBalancer":
		var decodedResponse DeleteNLBResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "addLoadBalancerWebServer":
		var decodedResponse AddNLBVMResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	case "listLoadBalancerWebServers":
		var decodedResponse ListNLBVMsResponse
		json.Unmarshal(body, &decodedResponse)
		return decodedResponse, nil

	}

	// only reached with unknown request
	return "", nil
}
