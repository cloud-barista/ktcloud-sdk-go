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
	"encoding/json"
)

type QueryAsyncJobResultResponse struct {
	Queryasyncjobresultresponse struct {
		AccountId     string  `json:"accountid"`
		UserId        string  `json:"userid"`
		Cmd           string  `json:"cmd"`
		JobStatus     float64 `json:"jobstatus"`
		JobProcStatus float64 `json:"jobprocstatus"`
		JobResultCode float64 `json:"jobresultcode"`
		JobResultType string  `json:"jobresulttype"`
		State 		  string  `json:"state"`

		JobResult 	json.RawMessage `json:"jobresult"`

		JobInstanceType string `json:"jobinstancetype"`
		JobInstanceId string `json:"jobinstanceid"`
		Created       string  `json:"created"`
		JobId         string  `json:"jobid"`

	} `json:"queryasyncjobresultresponse"`
}

// Query KT Cloud for the state of a scheduled job
func (c KtCloudClient) QueryAsyncJobResult(jobId string) (QueryAsyncJobResultResponse, error) {
	var resp QueryAsyncJobResultResponse

	params := url.Values{}
	params.Set("jobid", jobId)

	response, err := NewRequest(c, "queryAsyncJobResult", params)
	if err != nil {
		return resp, err
	}

	resp = response.(QueryAsyncJobResultResponse)

	return resp, nil
}
