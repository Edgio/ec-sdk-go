// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package task

import "github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

// NewTaskPostCreated creates a TaskPostCreated with default headers values
func NewTaskPostCreated() *TaskPostCreated {
	return &TaskPostCreated{}
}

/*
	TaskPostCreated describes a response with status code 201, with default header values.

Success
*/
type TaskPostCreated struct {
	models.TaskItem
}

// NewTaskPostBadRequest creates a TaskPostBadRequest with default headers values
func NewTaskPostBadRequest() *TaskPostBadRequest {
	return &TaskPostBadRequest{}
}

/*
	TaskPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TaskPostBadRequest struct {
	models.HyperionErrorReponse
}
