package registries

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Action string

const (
	ActionAllow Action = "Allow"
)

func PossibleValuesForAction() []string {
	return []string{
		string(ActionAllow),
	}
}

func parseAction(input string) (*Action, error) {
	vals := map[string]Action{
		"allow": ActionAllow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Action(input)
	return &out, nil
}

type ActionsRequired string

const (
	ActionsRequiredNone     ActionsRequired = "None"
	ActionsRequiredRecreate ActionsRequired = "Recreate"
)

func PossibleValuesForActionsRequired() []string {
	return []string{
		string(ActionsRequiredNone),
		string(ActionsRequiredRecreate),
	}
}

func parseActionsRequired(input string) (*ActionsRequired, error) {
	vals := map[string]ActionsRequired{
		"none":     ActionsRequiredNone,
		"recreate": ActionsRequiredRecreate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActionsRequired(input)
	return &out, nil
}

type ConnectionStatus string

const (
	ConnectionStatusApproved     ConnectionStatus = "Approved"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
	ConnectionStatusPending      ConnectionStatus = "Pending"
	ConnectionStatusRejected     ConnectionStatus = "Rejected"
)

func PossibleValuesForConnectionStatus() []string {
	return []string{
		string(ConnectionStatusApproved),
		string(ConnectionStatusDisconnected),
		string(ConnectionStatusPending),
		string(ConnectionStatusRejected),
	}
}

func parseConnectionStatus(input string) (*ConnectionStatus, error) {
	vals := map[string]ConnectionStatus{
		"approved":     ConnectionStatusApproved,
		"disconnected": ConnectionStatusDisconnected,
		"pending":      ConnectionStatusPending,
		"rejected":     ConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionStatus(input)
	return &out, nil
}

type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

func PossibleValuesForDefaultAction() []string {
	return []string{
		string(DefaultActionAllow),
		string(DefaultActionDeny),
	}
}

func parseDefaultAction(input string) (*DefaultAction, error) {
	vals := map[string]DefaultAction{
		"allow": DefaultActionAllow,
		"deny":  DefaultActionDeny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultAction(input)
	return &out, nil
}

type EncryptionStatus string

const (
	EncryptionStatusDisabled EncryptionStatus = "disabled"
	EncryptionStatusEnabled  EncryptionStatus = "enabled"
)

func PossibleValuesForEncryptionStatus() []string {
	return []string{
		string(EncryptionStatusDisabled),
		string(EncryptionStatusEnabled),
	}
}

func parseEncryptionStatus(input string) (*EncryptionStatus, error) {
	vals := map[string]EncryptionStatus{
		"disabled": EncryptionStatusDisabled,
		"enabled":  EncryptionStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionStatus(input)
	return &out, nil
}

type ExportPolicyStatus string

const (
	ExportPolicyStatusDisabled ExportPolicyStatus = "disabled"
	ExportPolicyStatusEnabled  ExportPolicyStatus = "enabled"
)

func PossibleValuesForExportPolicyStatus() []string {
	return []string{
		string(ExportPolicyStatusDisabled),
		string(ExportPolicyStatusEnabled),
	}
}

func parseExportPolicyStatus(input string) (*ExportPolicyStatus, error) {
	vals := map[string]ExportPolicyStatus{
		"disabled": ExportPolicyStatusDisabled,
		"enabled":  ExportPolicyStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExportPolicyStatus(input)
	return &out, nil
}

type ImportMode string

const (
	ImportModeForce   ImportMode = "Force"
	ImportModeNoForce ImportMode = "NoForce"
)

func PossibleValuesForImportMode() []string {
	return []string{
		string(ImportModeForce),
		string(ImportModeNoForce),
	}
}

func parseImportMode(input string) (*ImportMode, error) {
	vals := map[string]ImportMode{
		"force":   ImportModeForce,
		"noforce": ImportModeNoForce,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportMode(input)
	return &out, nil
}

type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = "AzureServices"
	NetworkRuleBypassOptionsNone          NetworkRuleBypassOptions = "None"
)

func PossibleValuesForNetworkRuleBypassOptions() []string {
	return []string{
		string(NetworkRuleBypassOptionsAzureServices),
		string(NetworkRuleBypassOptionsNone),
	}
}

func parseNetworkRuleBypassOptions(input string) (*NetworkRuleBypassOptions, error) {
	vals := map[string]NetworkRuleBypassOptions{
		"azureservices": NetworkRuleBypassOptionsAzureServices,
		"none":          NetworkRuleBypassOptionsNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkRuleBypassOptions(input)
	return &out, nil
}

type PasswordName string

const (
	PasswordNamePassword    PasswordName = "password"
	PasswordNamePasswordTwo PasswordName = "password2"
)

func PossibleValuesForPasswordName() []string {
	return []string{
		string(PasswordNamePassword),
		string(PasswordNamePasswordTwo),
	}
}

func parsePasswordName(input string) (*PasswordName, error) {
	vals := map[string]PasswordName{
		"password":  PasswordNamePassword,
		"password2": PasswordNamePasswordTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PasswordName(input)
	return &out, nil
}

type PolicyStatus string

const (
	PolicyStatusDisabled PolicyStatus = "disabled"
	PolicyStatusEnabled  PolicyStatus = "enabled"
)

func PossibleValuesForPolicyStatus() []string {
	return []string{
		string(PolicyStatusDisabled),
		string(PolicyStatusEnabled),
	}
}

func parsePolicyStatus(input string) (*PolicyStatus, error) {
	vals := map[string]PolicyStatus{
		"disabled": PolicyStatusDisabled,
		"enabled":  PolicyStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyStatus(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":  ProvisioningStateCanceled,
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

func parsePublicNetworkAccess(input string) (*PublicNetworkAccess, error) {
	vals := map[string]PublicNetworkAccess{
		"disabled": PublicNetworkAccessDisabled,
		"enabled":  PublicNetworkAccessEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicNetworkAccess(input)
	return &out, nil
}

type RegistryUsageUnit string

const (
	RegistryUsageUnitBytes RegistryUsageUnit = "Bytes"
	RegistryUsageUnitCount RegistryUsageUnit = "Count"
)

func PossibleValuesForRegistryUsageUnit() []string {
	return []string{
		string(RegistryUsageUnitBytes),
		string(RegistryUsageUnitCount),
	}
}

func parseRegistryUsageUnit(input string) (*RegistryUsageUnit, error) {
	vals := map[string]RegistryUsageUnit{
		"bytes": RegistryUsageUnitBytes,
		"count": RegistryUsageUnitCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryUsageUnit(input)
	return &out, nil
}

type SkuName string

const (
	SkuNameBasic    SkuName = "Basic"
	SkuNameClassic  SkuName = "Classic"
	SkuNamePremium  SkuName = "Premium"
	SkuNameStandard SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameBasic),
		string(SkuNameClassic),
		string(SkuNamePremium),
		string(SkuNameStandard),
	}
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"basic":    SkuNameBasic,
		"classic":  SkuNameClassic,
		"premium":  SkuNamePremium,
		"standard": SkuNameStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierClassic  SkuTier = "Classic"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierClassic),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"basic":    SkuTierBasic,
		"classic":  SkuTierClassic,
		"premium":  SkuTierPremium,
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}

type TokenPasswordName string

const (
	TokenPasswordNamePasswordOne TokenPasswordName = "password1"
	TokenPasswordNamePasswordTwo TokenPasswordName = "password2"
)

func PossibleValuesForTokenPasswordName() []string {
	return []string{
		string(TokenPasswordNamePasswordOne),
		string(TokenPasswordNamePasswordTwo),
	}
}

func parseTokenPasswordName(input string) (*TokenPasswordName, error) {
	vals := map[string]TokenPasswordName{
		"password1": TokenPasswordNamePasswordOne,
		"password2": TokenPasswordNamePasswordTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TokenPasswordName(input)
	return &out, nil
}

type TrustPolicyType string

const (
	TrustPolicyTypeNotary TrustPolicyType = "Notary"
)

func PossibleValuesForTrustPolicyType() []string {
	return []string{
		string(TrustPolicyTypeNotary),
	}
}

func parseTrustPolicyType(input string) (*TrustPolicyType, error) {
	vals := map[string]TrustPolicyType{
		"notary": TrustPolicyTypeNotary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrustPolicyType(input)
	return &out, nil
}

type ZoneRedundancy string

const (
	ZoneRedundancyDisabled ZoneRedundancy = "Disabled"
	ZoneRedundancyEnabled  ZoneRedundancy = "Enabled"
)

func PossibleValuesForZoneRedundancy() []string {
	return []string{
		string(ZoneRedundancyDisabled),
		string(ZoneRedundancyEnabled),
	}
}

func parseZoneRedundancy(input string) (*ZoneRedundancy, error) {
	vals := map[string]ZoneRedundancy{
		"disabled": ZoneRedundancyDisabled,
		"enabled":  ZoneRedundancyEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZoneRedundancy(input)
	return &out, nil
}
