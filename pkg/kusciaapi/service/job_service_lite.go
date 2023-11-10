// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"
	"fmt"

	"github.com/secretflow/kuscia/pkg/kusciaapi/errorcode"
	"github.com/secretflow/kuscia/pkg/kusciaapi/proxy"
	utils2 "github.com/secretflow/kuscia/pkg/web/utils"
	"github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapi"
)

type jobServiceLite struct {
	Initiator       string
	kusciaAPIClient proxy.KusciaAPIClient
}

func (h *jobServiceLite) CreateJob(ctx context.Context, request *kusciaapi.CreateJobRequest) *kusciaapi.CreateJobResponse {
	// do validate
	if err := validateCreateJobRequest(request, h.Initiator); err != nil {
		return &kusciaapi.CreateJobResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestValidate, err.Error()),
		}
	}
	// request the master api
	resp, err := h.kusciaAPIClient.CreateJob(ctx, request)
	if err != nil {
		return &kusciaapi.CreateJobResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestMasterFailed, err.Error()),
		}
	}
	return resp
}

func (h *jobServiceLite) QueryJob(ctx context.Context, request *kusciaapi.QueryJobRequest) *kusciaapi.QueryJobResponse {
	// do validate
	jobID := request.JobId
	if jobID == "" {
		return &kusciaapi.QueryJobResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestValidate, "job id can not be empty"),
		}
	}
	// request the master api
	resp, err := h.kusciaAPIClient.QueryJob(ctx, request)
	if err != nil {
		return &kusciaapi.QueryJobResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestMasterFailed, err.Error()),
		}
	}
	return resp
}

func (h *jobServiceLite) DeleteJob(ctx context.Context, request *kusciaapi.DeleteJobRequest) *kusciaapi.DeleteJobResponse {
	// do validate
	jobID := request.JobId
	if jobID == "" {
		return &kusciaapi.DeleteJobResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestValidate, "job id can not be empty"),
		}
	}
	// request the master api
	resp, err := h.kusciaAPIClient.DeleteJob(ctx, request)
	if err != nil {
		return &kusciaapi.DeleteJobResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestMasterFailed, err.Error()),
		}
	}
	return resp
}

func (h *jobServiceLite) StopJob(ctx context.Context, request *kusciaapi.StopJobRequest) *kusciaapi.StopJobResponse {
	// do validate
	jobID := request.JobId
	if jobID == "" {
		return &kusciaapi.StopJobResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestValidate, "job id can not be empty"),
		}
	}
	// request the master api
	resp, err := h.kusciaAPIClient.StopJob(ctx, request)
	if err != nil {
		return &kusciaapi.StopJobResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestMasterFailed, err.Error()),
		}
	}
	return resp
}

func (h *jobServiceLite) BatchQueryJobStatus(ctx context.Context, request *kusciaapi.BatchQueryJobStatusRequest) *kusciaapi.BatchQueryJobStatusResponse {
	// do validate
	if err := validateBatchQueryJobStatusRequest(request); err != nil {
		return &kusciaapi.BatchQueryJobStatusResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestValidate, err.Error()),
		}
	}
	// request the master api
	resp, err := h.kusciaAPIClient.BatchQueryJob(ctx, request)
	if err != nil {
		return &kusciaapi.BatchQueryJobStatusResponse{
			Status: utils2.BuildErrorResponseStatus(errorcode.ErrRequestMasterFailed, err.Error()),
		}
	}
	return resp
}

func (h *jobServiceLite) WatchJob(ctx context.Context, request *kusciaapi.WatchJobRequest, eventCh chan<- *kusciaapi.WatchJobEventResponse) error {
	// do validate
	if request.TimeoutSeconds < 0 {
		return fmt.Errorf("timeout seconds must be greater than or equal to 0")
	}
	// request the master api
	return h.kusciaAPIClient.WatchJob(ctx, request, eventCh)
}