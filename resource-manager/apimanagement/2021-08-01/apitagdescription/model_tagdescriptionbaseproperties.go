package apitagdescription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagDescriptionBaseProperties struct {
	Description             *string `json:"description,omitempty"`
	ExternalDocsDescription *string `json:"externalDocsDescription,omitempty"`
	ExternalDocsUrl         *string `json:"externalDocsUrl,omitempty"`
}
