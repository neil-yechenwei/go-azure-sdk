package simpolicies

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimPolicy struct {
	Id         *string                   `json:"id,omitempty"`
	Location   string                    `json:"location"`
	Name       *string                   `json:"name,omitempty"`
	Properties SimPolicyPropertiesFormat `json:"properties"`
	SystemData *systemdata.SystemData    `json:"systemData,omitempty"`
	Tags       *map[string]string        `json:"tags,omitempty"`
	Type       *string                   `json:"type,omitempty"`
}