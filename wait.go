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

// Blocks until the the asynchronous job has executed or has timed out.
// time.Duration unit => 1 nanosecond.  timeOut * 1,000,000,000 => 1 second
func (c KtCloudClient) WaitForAsyncJob(jobId string, timeOut time.Duration) error {
	done := make(chan struct{})
	defer close(done)

	result := make(chan error, 1)
	go func() {
		attempts := 0
		for {
			attempts += 1

			log.Printf("Checking the async job status... (attempt: %d)", attempts)
			response, err := c.QueryAsyncJobResult(jobId)
			if err != nil {
				result <- err
				return
			}

			// Check status of the job
			// 0 - Pending / In progress, Continue job
			// 1 - Succeeded
			// 2 - Failed
			// 3 - Cancelled  //Not supports on KT Cloud
			status := response.Queryasyncjobresultresponse.JobStatus
			log.Printf("The job status : %d", status)
			switch status {
			case 1:
				result <- nil
				return
			case 2:
				err := fmt.Errorf("WaitForAsyncJob() failed")
				result <- err
				return
			// case 3:
			// 	err := fmt.Errorf("WaitForAsyncJob was cancelled")
			// 	result <- err
			// 	return
			}

			// Wait 3 seconds between requests
			time.Sleep(3 * time.Second)

			// Verify whether we shouldn't exit or ...
			select {
			case <-done:
				// Finished, so just exit the goroutine
				return
			default:
				// Keep going
			}
		}
	}()

	log.Printf("Waiting for up to %f seconds for async job : %s", timeOut.Seconds(), jobId)
	select {
	case err := <-result:
		return err
	case <-time.After(timeOut):
		err := fmt.Errorf("Timeout while waiting to for the async job to finish")
		return err
	}
}

// WaitForVirtualMachineState simply blocks until the virtual machine is in the specified state.
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

			log.Printf("Checking the virtual machine state... (attempt: %d)", attempts)
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
			// Check what the real state will be.
			log.Printf("Current state: %s", currentState)
			log.Printf("Wanted state:  %s", wantedState)
			if currentState == wantedState {
				result <- nil
				return
			}

			// Wait 3 seconds in between
			time.Sleep(3 * time.Second)

			// Verify whether we shouldn't exit or ...
			select {
			case <-done:
				// Finished, so just exit the goroutine
				return
			default:
				// Keep going
			}
		}
	}()

	log.Printf("Waiting for up to %f seconds for Virtual Machine state to converge", timeOut.Seconds())
	select {
	case err := <-result:
		return err
	case <-time.After(timeOut):
		err := fmt.Errorf("Timeout while waiting to for the Virtual Machine to converge")
		return err
	}
}
