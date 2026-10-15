# ApiV1OrgFinancingAndOwnershipGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgName** | Pointer to **string** |  | [optional] 
**Resolution** | Pointer to [**ApiV1InvestorProfileGet200ResponseResolution**](ApiV1InvestorProfileGet200ResponseResolution.md) |  | [optional] 
**OrgId** | **NullableString** |  | 
**RequestedOrgId** | **NullableString** |  | 
**RolledUpToParent** | **bool** |  | 
**Financing** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing**](ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing.md) |  | 
**Investors** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseInvestors**](ApiV1OrgFinancingAndOwnershipGet200ResponseInvestors.md) |  | 
**OwnershipPositions** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPositions**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPositions.md) |  | 
**Semantics** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseSemantics**](ApiV1OrgFinancingAndOwnershipGet200ResponseSemantics.md) |  | 
**BuildState** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  | 
**Sections** | [**ApiV1OrgFinancingAndOwnershipGet200ResponseSections**](ApiV1OrgFinancingAndOwnershipGet200ResponseSections.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OwnershipPercentages** | Pointer to [**ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages**](ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages.md) |  | [optional] 
**NameWithheldNote** | Pointer to **NullableString** |  | [optional] 
**UnmatchedProductRevenue** | Pointer to [**ApiV1OrgFinancingAndOwnershipGet200ResponseUnmatchedProductRevenue**](ApiV1OrgFinancingAndOwnershipGet200ResponseUnmatchedProductRevenue.md) |  | [optional] 

## Methods

### NewApiV1OrgFinancingAndOwnershipGet200Response

`func NewApiV1OrgFinancingAndOwnershipGet200Response(orgId NullableString, requestedOrgId NullableString, rolledUpToParent bool, financing ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing, investors ApiV1OrgFinancingAndOwnershipGet200ResponseInvestors, ownershipPositions ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPositions, semantics ApiV1OrgFinancingAndOwnershipGet200ResponseSemantics, buildState ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, sections ApiV1OrgFinancingAndOwnershipGet200ResponseSections, ) *ApiV1OrgFinancingAndOwnershipGet200Response`

NewApiV1OrgFinancingAndOwnershipGet200Response instantiates a new ApiV1OrgFinancingAndOwnershipGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1OrgFinancingAndOwnershipGet200ResponseWithDefaults

`func NewApiV1OrgFinancingAndOwnershipGet200ResponseWithDefaults() *ApiV1OrgFinancingAndOwnershipGet200Response`

NewApiV1OrgFinancingAndOwnershipGet200ResponseWithDefaults instantiates a new ApiV1OrgFinancingAndOwnershipGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgName

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetResolution

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetResolution() ApiV1InvestorProfileGet200ResponseResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetResolutionOk() (*ApiV1InvestorProfileGet200ResponseResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetResolution(v ApiV1InvestorProfileGet200ResponseResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetOrgId

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### SetOrgIdNil

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetRequestedOrgId

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetRequestedOrgId() string`

GetRequestedOrgId returns the RequestedOrgId field if non-nil, zero value otherwise.

### GetRequestedOrgIdOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetRequestedOrgIdOk() (*string, bool)`

GetRequestedOrgIdOk returns a tuple with the RequestedOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedOrgId

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetRequestedOrgId(v string)`

SetRequestedOrgId sets RequestedOrgId field to given value.


### SetRequestedOrgIdNil

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetRequestedOrgIdNil(b bool)`

 SetRequestedOrgIdNil sets the value for RequestedOrgId to be an explicit nil

### UnsetRequestedOrgId
`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) UnsetRequestedOrgId()`

UnsetRequestedOrgId ensures that no value is present for RequestedOrgId, not even an explicit nil
### GetRolledUpToParent

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetRolledUpToParent() bool`

GetRolledUpToParent returns the RolledUpToParent field if non-nil, zero value otherwise.

### GetRolledUpToParentOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetRolledUpToParentOk() (*bool, bool)`

GetRolledUpToParentOk returns a tuple with the RolledUpToParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledUpToParent

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetRolledUpToParent(v bool)`

SetRolledUpToParent sets RolledUpToParent field to given value.


### GetFinancing

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetFinancing() ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing`

GetFinancing returns the Financing field if non-nil, zero value otherwise.

### GetFinancingOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetFinancingOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing, bool)`

GetFinancingOk returns a tuple with the Financing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancing

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetFinancing(v ApiV1OrgFinancingAndOwnershipGet200ResponseFinancing)`

SetFinancing sets Financing field to given value.


### GetInvestors

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetInvestors() ApiV1OrgFinancingAndOwnershipGet200ResponseInvestors`

GetInvestors returns the Investors field if non-nil, zero value otherwise.

### GetInvestorsOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetInvestorsOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseInvestors, bool)`

GetInvestorsOk returns a tuple with the Investors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestors

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetInvestors(v ApiV1OrgFinancingAndOwnershipGet200ResponseInvestors)`

SetInvestors sets Investors field to given value.


### GetOwnershipPositions

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetOwnershipPositions() ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPositions`

GetOwnershipPositions returns the OwnershipPositions field if non-nil, zero value otherwise.

### GetOwnershipPositionsOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetOwnershipPositionsOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPositions, bool)`

GetOwnershipPositionsOk returns a tuple with the OwnershipPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipPositions

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetOwnershipPositions(v ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPositions)`

SetOwnershipPositions sets OwnershipPositions field to given value.


### GetSemantics

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetSemantics() ApiV1OrgFinancingAndOwnershipGet200ResponseSemantics`

GetSemantics returns the Semantics field if non-nil, zero value otherwise.

### GetSemanticsOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetSemanticsOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseSemantics, bool)`

GetSemanticsOk returns a tuple with the Semantics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantics

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetSemantics(v ApiV1OrgFinancingAndOwnershipGet200ResponseSemantics)`

SetSemantics sets Semantics field to given value.


### GetBuildState

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetBuildState() ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetBuildStateOk() (*ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetBuildState(v ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetSections

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetSections() ApiV1OrgFinancingAndOwnershipGet200ResponseSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetSectionsOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetSections(v ApiV1OrgFinancingAndOwnershipGet200ResponseSections)`

SetSections sets Sections field to given value.


### GetError

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOwnershipPercentages

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetOwnershipPercentages() ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages`

GetOwnershipPercentages returns the OwnershipPercentages field if non-nil, zero value otherwise.

### GetOwnershipPercentagesOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetOwnershipPercentagesOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages, bool)`

GetOwnershipPercentagesOk returns a tuple with the OwnershipPercentages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipPercentages

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetOwnershipPercentages(v ApiV1OrgFinancingAndOwnershipGet200ResponseOwnershipPercentages)`

SetOwnershipPercentages sets OwnershipPercentages field to given value.

### HasOwnershipPercentages

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) HasOwnershipPercentages() bool`

HasOwnershipPercentages returns a boolean if a field has been set.

### GetNameWithheldNote

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetNameWithheldNote() string`

GetNameWithheldNote returns the NameWithheldNote field if non-nil, zero value otherwise.

### GetNameWithheldNoteOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetNameWithheldNoteOk() (*string, bool)`

GetNameWithheldNoteOk returns a tuple with the NameWithheldNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameWithheldNote

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetNameWithheldNote(v string)`

SetNameWithheldNote sets NameWithheldNote field to given value.

### HasNameWithheldNote

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) HasNameWithheldNote() bool`

HasNameWithheldNote returns a boolean if a field has been set.

### SetNameWithheldNoteNil

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetNameWithheldNoteNil(b bool)`

 SetNameWithheldNoteNil sets the value for NameWithheldNote to be an explicit nil

### UnsetNameWithheldNote
`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) UnsetNameWithheldNote()`

UnsetNameWithheldNote ensures that no value is present for NameWithheldNote, not even an explicit nil
### GetUnmatchedProductRevenue

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetUnmatchedProductRevenue() ApiV1OrgFinancingAndOwnershipGet200ResponseUnmatchedProductRevenue`

GetUnmatchedProductRevenue returns the UnmatchedProductRevenue field if non-nil, zero value otherwise.

### GetUnmatchedProductRevenueOk

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) GetUnmatchedProductRevenueOk() (*ApiV1OrgFinancingAndOwnershipGet200ResponseUnmatchedProductRevenue, bool)`

GetUnmatchedProductRevenueOk returns a tuple with the UnmatchedProductRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedProductRevenue

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) SetUnmatchedProductRevenue(v ApiV1OrgFinancingAndOwnershipGet200ResponseUnmatchedProductRevenue)`

SetUnmatchedProductRevenue sets UnmatchedProductRevenue field to given value.

### HasUnmatchedProductRevenue

`func (o *ApiV1OrgFinancingAndOwnershipGet200Response) HasUnmatchedProductRevenue() bool`

HasUnmatchedProductRevenue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


