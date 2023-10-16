package operationstatus

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationOperationStatus struct {
	EndTime         *string              `json:"endTime,omitempty"`
	Error           *StorageSyncApiError `json:"error,omitempty"`
	Id              *string              `json:"id,omitempty"`
	Name            *string              `json:"name,omitempty"`
	PercentComplete *int64               `json:"percentComplete,omitempty"`
	StartTime       *string              `json:"startTime,omitempty"`
	Status          *string              `json:"status,omitempty"`
}

func (o *LocationOperationStatus) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *LocationOperationStatus) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *LocationOperationStatus) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *LocationOperationStatus) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
