# SbtiCrosscheckRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **NullableString** | Legal-entity level: the organization the cross-check is for. | [optional] 
**OrgName** | Pointer to **NullableString** | Display name of the organization record. | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: the company the organization belongs to (null when unresolved). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the organization&#39;s company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the organization&#39;s group. | [optional] 
**ReportedCompanyName** | Pointer to **NullableString** | The company name as the SBTi list or the disclosure gives it. | [optional] 
**ReconStatus** | Pointer to **NullableString** | The cross-check result: agree, extracted_only, sbti_only, target_year_conflict or self_claim_conflict. | [optional] 
**HasNetZeroTarget** | Pointer to **NullableBool** | Whether a net-zero target is stated (SBTi&#39;s value where it lists the organization, else the disclosure&#39;s). | [optional] 
**NetZeroYear** | Pointer to **NullableInt32** | The net-zero target year. | [optional] 
**HasNearTermTarget** | Pointer to **NullableBool** | Whether a near-term target is stated. | [optional] 
**NearTermTargetYear** | Pointer to **NullableInt32** | The near-term target year. | [optional] 
**SbtiValidated** | Pointer to **NullableBool** | Whether SBTi lists the targets as validated; null when SBTi does not list the organization. | [optional] 
**ClaimedSbtiStatus** | Pointer to **NullableString** | The SBTi status the organization&#39;s own filing claims. | [optional] 
**InSbti** | Pointer to **NullableBool** | Whether the SBTi list carries the organization. | [optional] 
**InExtracted** | Pointer to **NullableBool** | Whether an extraction of the organization&#39;s own filing exists. | [optional] 

## Methods

### NewSbtiCrosscheckRow

`func NewSbtiCrosscheckRow() *SbtiCrosscheckRow`

NewSbtiCrosscheckRow instantiates a new SbtiCrosscheckRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSbtiCrosscheckRowWithDefaults

`func NewSbtiCrosscheckRowWithDefaults() *SbtiCrosscheckRow`

NewSbtiCrosscheckRowWithDefaults instantiates a new SbtiCrosscheckRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *SbtiCrosscheckRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SbtiCrosscheckRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SbtiCrosscheckRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SbtiCrosscheckRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *SbtiCrosscheckRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *SbtiCrosscheckRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *SbtiCrosscheckRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *SbtiCrosscheckRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *SbtiCrosscheckRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *SbtiCrosscheckRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *SbtiCrosscheckRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *SbtiCrosscheckRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetCompanyId

`func (o *SbtiCrosscheckRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *SbtiCrosscheckRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *SbtiCrosscheckRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *SbtiCrosscheckRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *SbtiCrosscheckRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *SbtiCrosscheckRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *SbtiCrosscheckRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *SbtiCrosscheckRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *SbtiCrosscheckRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *SbtiCrosscheckRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *SbtiCrosscheckRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *SbtiCrosscheckRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *SbtiCrosscheckRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SbtiCrosscheckRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SbtiCrosscheckRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SbtiCrosscheckRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *SbtiCrosscheckRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *SbtiCrosscheckRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *SbtiCrosscheckRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *SbtiCrosscheckRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *SbtiCrosscheckRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *SbtiCrosscheckRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *SbtiCrosscheckRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *SbtiCrosscheckRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetReportedCompanyName

`func (o *SbtiCrosscheckRow) GetReportedCompanyName() string`

GetReportedCompanyName returns the ReportedCompanyName field if non-nil, zero value otherwise.

### GetReportedCompanyNameOk

`func (o *SbtiCrosscheckRow) GetReportedCompanyNameOk() (*string, bool)`

GetReportedCompanyNameOk returns a tuple with the ReportedCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedCompanyName

`func (o *SbtiCrosscheckRow) SetReportedCompanyName(v string)`

SetReportedCompanyName sets ReportedCompanyName field to given value.

### HasReportedCompanyName

`func (o *SbtiCrosscheckRow) HasReportedCompanyName() bool`

HasReportedCompanyName returns a boolean if a field has been set.

### SetReportedCompanyNameNil

`func (o *SbtiCrosscheckRow) SetReportedCompanyNameNil(b bool)`

 SetReportedCompanyNameNil sets the value for ReportedCompanyName to be an explicit nil

### UnsetReportedCompanyName
`func (o *SbtiCrosscheckRow) UnsetReportedCompanyName()`

UnsetReportedCompanyName ensures that no value is present for ReportedCompanyName, not even an explicit nil
### GetReconStatus

`func (o *SbtiCrosscheckRow) GetReconStatus() string`

GetReconStatus returns the ReconStatus field if non-nil, zero value otherwise.

### GetReconStatusOk

`func (o *SbtiCrosscheckRow) GetReconStatusOk() (*string, bool)`

GetReconStatusOk returns a tuple with the ReconStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconStatus

`func (o *SbtiCrosscheckRow) SetReconStatus(v string)`

SetReconStatus sets ReconStatus field to given value.

### HasReconStatus

`func (o *SbtiCrosscheckRow) HasReconStatus() bool`

HasReconStatus returns a boolean if a field has been set.

### SetReconStatusNil

`func (o *SbtiCrosscheckRow) SetReconStatusNil(b bool)`

 SetReconStatusNil sets the value for ReconStatus to be an explicit nil

### UnsetReconStatus
`func (o *SbtiCrosscheckRow) UnsetReconStatus()`

UnsetReconStatus ensures that no value is present for ReconStatus, not even an explicit nil
### GetHasNetZeroTarget

`func (o *SbtiCrosscheckRow) GetHasNetZeroTarget() bool`

GetHasNetZeroTarget returns the HasNetZeroTarget field if non-nil, zero value otherwise.

### GetHasNetZeroTargetOk

`func (o *SbtiCrosscheckRow) GetHasNetZeroTargetOk() (*bool, bool)`

GetHasNetZeroTargetOk returns a tuple with the HasNetZeroTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetZeroTarget

`func (o *SbtiCrosscheckRow) SetHasNetZeroTarget(v bool)`

SetHasNetZeroTarget sets HasNetZeroTarget field to given value.

### HasHasNetZeroTarget

`func (o *SbtiCrosscheckRow) HasHasNetZeroTarget() bool`

HasHasNetZeroTarget returns a boolean if a field has been set.

### SetHasNetZeroTargetNil

`func (o *SbtiCrosscheckRow) SetHasNetZeroTargetNil(b bool)`

 SetHasNetZeroTargetNil sets the value for HasNetZeroTarget to be an explicit nil

### UnsetHasNetZeroTarget
`func (o *SbtiCrosscheckRow) UnsetHasNetZeroTarget()`

UnsetHasNetZeroTarget ensures that no value is present for HasNetZeroTarget, not even an explicit nil
### GetNetZeroYear

`func (o *SbtiCrosscheckRow) GetNetZeroYear() int32`

GetNetZeroYear returns the NetZeroYear field if non-nil, zero value otherwise.

### GetNetZeroYearOk

`func (o *SbtiCrosscheckRow) GetNetZeroYearOk() (*int32, bool)`

GetNetZeroYearOk returns a tuple with the NetZeroYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetZeroYear

`func (o *SbtiCrosscheckRow) SetNetZeroYear(v int32)`

SetNetZeroYear sets NetZeroYear field to given value.

### HasNetZeroYear

`func (o *SbtiCrosscheckRow) HasNetZeroYear() bool`

HasNetZeroYear returns a boolean if a field has been set.

### SetNetZeroYearNil

`func (o *SbtiCrosscheckRow) SetNetZeroYearNil(b bool)`

 SetNetZeroYearNil sets the value for NetZeroYear to be an explicit nil

### UnsetNetZeroYear
`func (o *SbtiCrosscheckRow) UnsetNetZeroYear()`

UnsetNetZeroYear ensures that no value is present for NetZeroYear, not even an explicit nil
### GetHasNearTermTarget

`func (o *SbtiCrosscheckRow) GetHasNearTermTarget() bool`

GetHasNearTermTarget returns the HasNearTermTarget field if non-nil, zero value otherwise.

### GetHasNearTermTargetOk

`func (o *SbtiCrosscheckRow) GetHasNearTermTargetOk() (*bool, bool)`

GetHasNearTermTargetOk returns a tuple with the HasNearTermTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNearTermTarget

`func (o *SbtiCrosscheckRow) SetHasNearTermTarget(v bool)`

SetHasNearTermTarget sets HasNearTermTarget field to given value.

### HasHasNearTermTarget

`func (o *SbtiCrosscheckRow) HasHasNearTermTarget() bool`

HasHasNearTermTarget returns a boolean if a field has been set.

### SetHasNearTermTargetNil

`func (o *SbtiCrosscheckRow) SetHasNearTermTargetNil(b bool)`

 SetHasNearTermTargetNil sets the value for HasNearTermTarget to be an explicit nil

### UnsetHasNearTermTarget
`func (o *SbtiCrosscheckRow) UnsetHasNearTermTarget()`

UnsetHasNearTermTarget ensures that no value is present for HasNearTermTarget, not even an explicit nil
### GetNearTermTargetYear

`func (o *SbtiCrosscheckRow) GetNearTermTargetYear() int32`

GetNearTermTargetYear returns the NearTermTargetYear field if non-nil, zero value otherwise.

### GetNearTermTargetYearOk

`func (o *SbtiCrosscheckRow) GetNearTermTargetYearOk() (*int32, bool)`

GetNearTermTargetYearOk returns a tuple with the NearTermTargetYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearTermTargetYear

`func (o *SbtiCrosscheckRow) SetNearTermTargetYear(v int32)`

SetNearTermTargetYear sets NearTermTargetYear field to given value.

### HasNearTermTargetYear

`func (o *SbtiCrosscheckRow) HasNearTermTargetYear() bool`

HasNearTermTargetYear returns a boolean if a field has been set.

### SetNearTermTargetYearNil

`func (o *SbtiCrosscheckRow) SetNearTermTargetYearNil(b bool)`

 SetNearTermTargetYearNil sets the value for NearTermTargetYear to be an explicit nil

### UnsetNearTermTargetYear
`func (o *SbtiCrosscheckRow) UnsetNearTermTargetYear()`

UnsetNearTermTargetYear ensures that no value is present for NearTermTargetYear, not even an explicit nil
### GetSbtiValidated

`func (o *SbtiCrosscheckRow) GetSbtiValidated() bool`

GetSbtiValidated returns the SbtiValidated field if non-nil, zero value otherwise.

### GetSbtiValidatedOk

`func (o *SbtiCrosscheckRow) GetSbtiValidatedOk() (*bool, bool)`

GetSbtiValidatedOk returns a tuple with the SbtiValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtiValidated

`func (o *SbtiCrosscheckRow) SetSbtiValidated(v bool)`

SetSbtiValidated sets SbtiValidated field to given value.

### HasSbtiValidated

`func (o *SbtiCrosscheckRow) HasSbtiValidated() bool`

HasSbtiValidated returns a boolean if a field has been set.

### SetSbtiValidatedNil

`func (o *SbtiCrosscheckRow) SetSbtiValidatedNil(b bool)`

 SetSbtiValidatedNil sets the value for SbtiValidated to be an explicit nil

### UnsetSbtiValidated
`func (o *SbtiCrosscheckRow) UnsetSbtiValidated()`

UnsetSbtiValidated ensures that no value is present for SbtiValidated, not even an explicit nil
### GetClaimedSbtiStatus

`func (o *SbtiCrosscheckRow) GetClaimedSbtiStatus() string`

GetClaimedSbtiStatus returns the ClaimedSbtiStatus field if non-nil, zero value otherwise.

### GetClaimedSbtiStatusOk

`func (o *SbtiCrosscheckRow) GetClaimedSbtiStatusOk() (*string, bool)`

GetClaimedSbtiStatusOk returns a tuple with the ClaimedSbtiStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedSbtiStatus

`func (o *SbtiCrosscheckRow) SetClaimedSbtiStatus(v string)`

SetClaimedSbtiStatus sets ClaimedSbtiStatus field to given value.

### HasClaimedSbtiStatus

`func (o *SbtiCrosscheckRow) HasClaimedSbtiStatus() bool`

HasClaimedSbtiStatus returns a boolean if a field has been set.

### SetClaimedSbtiStatusNil

`func (o *SbtiCrosscheckRow) SetClaimedSbtiStatusNil(b bool)`

 SetClaimedSbtiStatusNil sets the value for ClaimedSbtiStatus to be an explicit nil

### UnsetClaimedSbtiStatus
`func (o *SbtiCrosscheckRow) UnsetClaimedSbtiStatus()`

UnsetClaimedSbtiStatus ensures that no value is present for ClaimedSbtiStatus, not even an explicit nil
### GetInSbti

`func (o *SbtiCrosscheckRow) GetInSbti() bool`

GetInSbti returns the InSbti field if non-nil, zero value otherwise.

### GetInSbtiOk

`func (o *SbtiCrosscheckRow) GetInSbtiOk() (*bool, bool)`

GetInSbtiOk returns a tuple with the InSbti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSbti

`func (o *SbtiCrosscheckRow) SetInSbti(v bool)`

SetInSbti sets InSbti field to given value.

### HasInSbti

`func (o *SbtiCrosscheckRow) HasInSbti() bool`

HasInSbti returns a boolean if a field has been set.

### SetInSbtiNil

`func (o *SbtiCrosscheckRow) SetInSbtiNil(b bool)`

 SetInSbtiNil sets the value for InSbti to be an explicit nil

### UnsetInSbti
`func (o *SbtiCrosscheckRow) UnsetInSbti()`

UnsetInSbti ensures that no value is present for InSbti, not even an explicit nil
### GetInExtracted

`func (o *SbtiCrosscheckRow) GetInExtracted() bool`

GetInExtracted returns the InExtracted field if non-nil, zero value otherwise.

### GetInExtractedOk

`func (o *SbtiCrosscheckRow) GetInExtractedOk() (*bool, bool)`

GetInExtractedOk returns a tuple with the InExtracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInExtracted

`func (o *SbtiCrosscheckRow) SetInExtracted(v bool)`

SetInExtracted sets InExtracted field to given value.

### HasInExtracted

`func (o *SbtiCrosscheckRow) HasInExtracted() bool`

HasInExtracted returns a boolean if a field has been set.

### SetInExtractedNil

`func (o *SbtiCrosscheckRow) SetInExtractedNil(b bool)`

 SetInExtractedNil sets the value for InExtracted to be an explicit nil

### UnsetInExtracted
`func (o *SbtiCrosscheckRow) UnsetInExtracted()`

UnsetInExtracted ensures that no value is present for InExtracted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


