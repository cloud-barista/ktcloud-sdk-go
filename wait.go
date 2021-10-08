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
	"fmt"
	"log"
	"time"
)

// waitForAsyncJob simply blocks until the the asynchronous job has
// executed or has timed out.
func (c KtCloudClient) WaitForAsyncJob(jobId string, timeOut time.Duration) error {
	done := make(chan struct{})
	defer close(done)

	result := make(chan error, 1)
	go func() {
		attempts := 0
		for {
			attempts += 1

			log.Printf("Checking async job status... (attempt: %d)", attempts)
			response, err := c.QueryAsyncJobResult(jobId)
			if err != nil {
				result <- err
				return
			}

			// Check status of the job we issued.
			// 0 - pending / in progress, we continue
			// 1 - succedded
			// 2 - failed
			// 3 - cancelled
			status := response.Queryasyncjobresultresponse.JobStatus
			switch status {
			case 1:
				result <- nil
				return
			case 2:
				err := fmt.Errorf("WaitForAsyncJob failed")
				result <- err
				return
			case 3:
				err := fmt.Errorf("WaitForAsyncJob was cancelled")
				result <- err
				return
			}

			// Wait 3 seconds between requests
			time.Sleep(3 * time.Second)

			// Verify we shouldn't exit
			select {
			case <-done:
				// We finished, so just exit the goroutine
				return
			default:
				// Keep going
			}
		}
	}()

	log.Printf("Waiting for up to %d seconds for async job %s", timeOut, jobId)
	select {
	case err := <-result:
		return err
	case <-time.After(timeOut):
		err := fmt.Errorf("Timeout while waiting to for async job to finish")
		return err
	}
}

// WaitForVirtualMachineState simply blocks until the virtual machine
// is in the specified state.
func (c KtCloudClient) WaitForVirtualMachineState(zoneId string, vmId string, wantedState string, timeOut time.Duration) error {

	vmListReqInfo := ListVMReqInfo{
		ZoneId: 	zoneId,
		VMId: 		vmId,
	}

	done := make(chan struct{})
	defer close(done)

	result := make(chan error, 1)
	go func() {
		attempts := 0
		for {
			attempts += 1

			log.Printf("Checking virtual machine state... (attempt: %d)", attempts)
			response, err := c.ListVirtualMachines(vmListReqInfo)
			if err != nil {
				result <- err
				return
			}

			count := response.Listvirtualmachinesresponse.Count
			if count != 1 {
				result <- err
				return
			}

			currentState := response.Listvirtualmachinesresponse.Virtualmachine[0].State
			// check what the real state will be.
			log.Printf("current state: %s", currentState)
			log.Printf("wanted state:  %s", wantedState)
			if currentState == wantedState {
				result <- nil
				return
			}

			// Wait 3 seconds in between
			time.Sleep(3 * time.Second)

			// Verify we shouldn't exit
			select {
			case <-done:
				// We finished, so just exit the goroutine
				return
			default:
				// Keep going
			}
		}
	}()

	log.Printf("Waiting for up to %d seconds for Virtual Machine state to converge", timeOut)
	select {
	case err := <-result:
		return err
	case <-time.After(timeOut):
		err := fmt.Errorf("Timeout while waiting to for Virtual Machine to converge")
		return err
	}
}