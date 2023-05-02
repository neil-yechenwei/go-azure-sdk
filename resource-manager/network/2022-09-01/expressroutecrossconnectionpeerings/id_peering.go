package expressroutecrossconnectionpeerings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PeeringId{}

// PeeringId is a struct representing the Resource ID for a Peering
type PeeringId struct {
	SubscriptionId                  string
	ResourceGroupName               string
	ExpressRouteCrossConnectionName string
	PeeringName                     string
}

// NewPeeringID returns a new PeeringId struct
func NewPeeringID(subscriptionId string, resourceGroupName string, expressRouteCrossConnectionName string, peeringName string) PeeringId {
	return PeeringId{
		SubscriptionId:                  subscriptionId,
		ResourceGroupName:               resourceGroupName,
		ExpressRouteCrossConnectionName: expressRouteCrossConnectionName,
		PeeringName:                     peeringName,
	}
}

// ParsePeeringID parses 'input' into a PeeringId
func ParsePeeringID(input string) (*PeeringId, error) {
	parser := resourceids.NewParserFromResourceIdType(PeeringId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PeeringId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ExpressRouteCrossConnectionName, ok = parsed.Parsed["expressRouteCrossConnectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'expressRouteCrossConnectionName' was not found in the resource id %q", input)
	}

	if id.PeeringName, ok = parsed.Parsed["peeringName"]; !ok {
		return nil, fmt.Errorf("the segment 'peeringName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePeeringIDInsensitively parses 'input' case-insensitively into a PeeringId
// note: this method should only be used for API response data and not user input
func ParsePeeringIDInsensitively(input string) (*PeeringId, error) {
	parser := resourceids.NewParserFromResourceIdType(PeeringId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PeeringId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ExpressRouteCrossConnectionName, ok = parsed.Parsed["expressRouteCrossConnectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'expressRouteCrossConnectionName' was not found in the resource id %q", input)
	}

	if id.PeeringName, ok = parsed.Parsed["peeringName"]; !ok {
		return nil, fmt.Errorf("the segment 'peeringName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePeeringID checks that 'input' can be parsed as a Peering ID
func ValidatePeeringID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePeeringID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Peering ID
func (id PeeringId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/expressRouteCrossConnections/%s/peerings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ExpressRouteCrossConnectionName, id.PeeringName)
}

// Segments returns a slice of Resource ID Segments which comprise this Peering ID
func (id PeeringId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticExpressRouteCrossConnections", "expressRouteCrossConnections", "expressRouteCrossConnections"),
		resourceids.UserSpecifiedSegment("expressRouteCrossConnectionName", "expressRouteCrossConnectionValue"),
		resourceids.StaticSegment("staticPeerings", "peerings", "peerings"),
		resourceids.UserSpecifiedSegment("peeringName", "peeringValue"),
	}
}

// String returns a human-readable description of this Peering ID
func (id PeeringId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Express Route Cross Connection Name: %q", id.ExpressRouteCrossConnectionName),
		fmt.Sprintf("Peering Name: %q", id.PeeringName),
	}
	return fmt.Sprintf("Peering (%s)", strings.Join(components, "\n"))
}