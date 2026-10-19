# BiomarkerBqpAssertionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAssertionId** | Pointer to **NullableString** | The assertion id. | [optional] 
**RouteKey** | Pointer to **NullableString** | The biomarker&#39;s route key, the entry key for the biomarkers object type. | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The biomarker key. | [optional] 
**DdtNumber** | Pointer to **NullableString** | The FDA drug development tool number of the project. | [optional] 
**DdtProjectId** | Pointer to **NullableString** | The FDA project id. | [optional] 
**ProjectName** | Pointer to **NullableString** | The project&#39;s name. | [optional] 
**BqpSignal** | Pointer to **NullableString** | The source status of the project. Only qualified means FDA-qualified; every other value is a project status. | [optional] 
**IsFdaQualified** | Pointer to **NullableBool** | True only when the source status is exactly qualified. | [optional] 
**SubmissionStatus** | Pointer to **NullableString** | The submission status as the source words it. | [optional] 
**SubmissionStage** | Pointer to **NullableString** | The submission stage as the source words it. | [optional] 
**SubmitterName** | Pointer to **NullableString** | The submitter as the source names it. | [optional] 
**SubmitterOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the submitter resolved to (null when unresolved). | [optional] 
**SubmitterCompanyId** | Pointer to **NullableString** | Company level: the company the submitter belongs to (null when unresolved). | [optional] 
**SubmitterCompanyName** | Pointer to **NullableString** | Display name of the submitter&#39;s company. | [optional] 
**SubmitterGroupId** | Pointer to **NullableString** | Group level: the submitter&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SubmitterGroupName** | Pointer to **NullableString** | Display name of the submitter&#39;s group. | [optional] 
**SubmitterMatchTier** | Pointer to **NullableString** | How the submitter name matched an organization record. | [optional] 
**SubmitterRole** | Pointer to **NullableString** | The submitter&#39;s role on the project. | [optional] 
**BiomarkerRole** | Pointer to **NullableString** | The biomarker&#39;s role in the project (for example a diagnostic or a surrogate endpoint). | [optional] 
**BiomarkerRoleMethod** | Pointer to **NullableString** | How the biomarker role was assigned. | [optional] 
**IsSurrogateEndpointClaim** | Pointer to **NullableBool** | True when the project claims a surrogate endpoint. | [optional] 
**SourceUrl** | Pointer to **NullableString** | Link to the FDA project page. | [optional] 
**SnapshotDate** | Pointer to **NullableString** | The date the source list was captured. | [optional] 

## Methods

### NewBiomarkerBqpAssertionsRow

`func NewBiomarkerBqpAssertionsRow() *BiomarkerBqpAssertionsRow`

NewBiomarkerBqpAssertionsRow instantiates a new BiomarkerBqpAssertionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerBqpAssertionsRowWithDefaults

`func NewBiomarkerBqpAssertionsRowWithDefaults() *BiomarkerBqpAssertionsRow`

NewBiomarkerBqpAssertionsRowWithDefaults instantiates a new BiomarkerBqpAssertionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAssertionId

`func (o *BiomarkerBqpAssertionsRow) GetSourceAssertionId() string`

GetSourceAssertionId returns the SourceAssertionId field if non-nil, zero value otherwise.

### GetSourceAssertionIdOk

`func (o *BiomarkerBqpAssertionsRow) GetSourceAssertionIdOk() (*string, bool)`

GetSourceAssertionIdOk returns a tuple with the SourceAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAssertionId

`func (o *BiomarkerBqpAssertionsRow) SetSourceAssertionId(v string)`

SetSourceAssertionId sets SourceAssertionId field to given value.

### HasSourceAssertionId

`func (o *BiomarkerBqpAssertionsRow) HasSourceAssertionId() bool`

HasSourceAssertionId returns a boolean if a field has been set.

### SetSourceAssertionIdNil

`func (o *BiomarkerBqpAssertionsRow) SetSourceAssertionIdNil(b bool)`

 SetSourceAssertionIdNil sets the value for SourceAssertionId to be an explicit nil

### UnsetSourceAssertionId
`func (o *BiomarkerBqpAssertionsRow) UnsetSourceAssertionId()`

UnsetSourceAssertionId ensures that no value is present for SourceAssertionId, not even an explicit nil
### GetRouteKey

`func (o *BiomarkerBqpAssertionsRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerBqpAssertionsRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerBqpAssertionsRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerBqpAssertionsRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerBqpAssertionsRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerBqpAssertionsRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerId

`func (o *BiomarkerBqpAssertionsRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkerBqpAssertionsRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkerBqpAssertionsRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkerBqpAssertionsRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkerBqpAssertionsRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkerBqpAssertionsRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetDdtNumber

`func (o *BiomarkerBqpAssertionsRow) GetDdtNumber() string`

GetDdtNumber returns the DdtNumber field if non-nil, zero value otherwise.

### GetDdtNumberOk

`func (o *BiomarkerBqpAssertionsRow) GetDdtNumberOk() (*string, bool)`

GetDdtNumberOk returns a tuple with the DdtNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdtNumber

`func (o *BiomarkerBqpAssertionsRow) SetDdtNumber(v string)`

SetDdtNumber sets DdtNumber field to given value.

### HasDdtNumber

`func (o *BiomarkerBqpAssertionsRow) HasDdtNumber() bool`

HasDdtNumber returns a boolean if a field has been set.

### SetDdtNumberNil

`func (o *BiomarkerBqpAssertionsRow) SetDdtNumberNil(b bool)`

 SetDdtNumberNil sets the value for DdtNumber to be an explicit nil

### UnsetDdtNumber
`func (o *BiomarkerBqpAssertionsRow) UnsetDdtNumber()`

UnsetDdtNumber ensures that no value is present for DdtNumber, not even an explicit nil
### GetDdtProjectId

`func (o *BiomarkerBqpAssertionsRow) GetDdtProjectId() string`

GetDdtProjectId returns the DdtProjectId field if non-nil, zero value otherwise.

### GetDdtProjectIdOk

`func (o *BiomarkerBqpAssertionsRow) GetDdtProjectIdOk() (*string, bool)`

GetDdtProjectIdOk returns a tuple with the DdtProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdtProjectId

`func (o *BiomarkerBqpAssertionsRow) SetDdtProjectId(v string)`

SetDdtProjectId sets DdtProjectId field to given value.

### HasDdtProjectId

`func (o *BiomarkerBqpAssertionsRow) HasDdtProjectId() bool`

HasDdtProjectId returns a boolean if a field has been set.

### SetDdtProjectIdNil

`func (o *BiomarkerBqpAssertionsRow) SetDdtProjectIdNil(b bool)`

 SetDdtProjectIdNil sets the value for DdtProjectId to be an explicit nil

### UnsetDdtProjectId
`func (o *BiomarkerBqpAssertionsRow) UnsetDdtProjectId()`

UnsetDdtProjectId ensures that no value is present for DdtProjectId, not even an explicit nil
### GetProjectName

`func (o *BiomarkerBqpAssertionsRow) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BiomarkerBqpAssertionsRow) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BiomarkerBqpAssertionsRow) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BiomarkerBqpAssertionsRow) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BiomarkerBqpAssertionsRow) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BiomarkerBqpAssertionsRow) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetBqpSignal

`func (o *BiomarkerBqpAssertionsRow) GetBqpSignal() string`

GetBqpSignal returns the BqpSignal field if non-nil, zero value otherwise.

### GetBqpSignalOk

`func (o *BiomarkerBqpAssertionsRow) GetBqpSignalOk() (*string, bool)`

GetBqpSignalOk returns a tuple with the BqpSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqpSignal

`func (o *BiomarkerBqpAssertionsRow) SetBqpSignal(v string)`

SetBqpSignal sets BqpSignal field to given value.

### HasBqpSignal

`func (o *BiomarkerBqpAssertionsRow) HasBqpSignal() bool`

HasBqpSignal returns a boolean if a field has been set.

### SetBqpSignalNil

`func (o *BiomarkerBqpAssertionsRow) SetBqpSignalNil(b bool)`

 SetBqpSignalNil sets the value for BqpSignal to be an explicit nil

### UnsetBqpSignal
`func (o *BiomarkerBqpAssertionsRow) UnsetBqpSignal()`

UnsetBqpSignal ensures that no value is present for BqpSignal, not even an explicit nil
### GetIsFdaQualified

`func (o *BiomarkerBqpAssertionsRow) GetIsFdaQualified() bool`

GetIsFdaQualified returns the IsFdaQualified field if non-nil, zero value otherwise.

### GetIsFdaQualifiedOk

`func (o *BiomarkerBqpAssertionsRow) GetIsFdaQualifiedOk() (*bool, bool)`

GetIsFdaQualifiedOk returns a tuple with the IsFdaQualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFdaQualified

`func (o *BiomarkerBqpAssertionsRow) SetIsFdaQualified(v bool)`

SetIsFdaQualified sets IsFdaQualified field to given value.

### HasIsFdaQualified

`func (o *BiomarkerBqpAssertionsRow) HasIsFdaQualified() bool`

HasIsFdaQualified returns a boolean if a field has been set.

### SetIsFdaQualifiedNil

`func (o *BiomarkerBqpAssertionsRow) SetIsFdaQualifiedNil(b bool)`

 SetIsFdaQualifiedNil sets the value for IsFdaQualified to be an explicit nil

### UnsetIsFdaQualified
`func (o *BiomarkerBqpAssertionsRow) UnsetIsFdaQualified()`

UnsetIsFdaQualified ensures that no value is present for IsFdaQualified, not even an explicit nil
### GetSubmissionStatus

`func (o *BiomarkerBqpAssertionsRow) GetSubmissionStatus() string`

GetSubmissionStatus returns the SubmissionStatus field if non-nil, zero value otherwise.

### GetSubmissionStatusOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmissionStatusOk() (*string, bool)`

GetSubmissionStatusOk returns a tuple with the SubmissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionStatus

`func (o *BiomarkerBqpAssertionsRow) SetSubmissionStatus(v string)`

SetSubmissionStatus sets SubmissionStatus field to given value.

### HasSubmissionStatus

`func (o *BiomarkerBqpAssertionsRow) HasSubmissionStatus() bool`

HasSubmissionStatus returns a boolean if a field has been set.

### SetSubmissionStatusNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmissionStatusNil(b bool)`

 SetSubmissionStatusNil sets the value for SubmissionStatus to be an explicit nil

### UnsetSubmissionStatus
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmissionStatus()`

UnsetSubmissionStatus ensures that no value is present for SubmissionStatus, not even an explicit nil
### GetSubmissionStage

`func (o *BiomarkerBqpAssertionsRow) GetSubmissionStage() string`

GetSubmissionStage returns the SubmissionStage field if non-nil, zero value otherwise.

### GetSubmissionStageOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmissionStageOk() (*string, bool)`

GetSubmissionStageOk returns a tuple with the SubmissionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionStage

`func (o *BiomarkerBqpAssertionsRow) SetSubmissionStage(v string)`

SetSubmissionStage sets SubmissionStage field to given value.

### HasSubmissionStage

`func (o *BiomarkerBqpAssertionsRow) HasSubmissionStage() bool`

HasSubmissionStage returns a boolean if a field has been set.

### SetSubmissionStageNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmissionStageNil(b bool)`

 SetSubmissionStageNil sets the value for SubmissionStage to be an explicit nil

### UnsetSubmissionStage
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmissionStage()`

UnsetSubmissionStage ensures that no value is present for SubmissionStage, not even an explicit nil
### GetSubmitterName

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterName() string`

GetSubmitterName returns the SubmitterName field if non-nil, zero value otherwise.

### GetSubmitterNameOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterNameOk() (*string, bool)`

GetSubmitterNameOk returns a tuple with the SubmitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterName

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterName(v string)`

SetSubmitterName sets SubmitterName field to given value.

### HasSubmitterName

`func (o *BiomarkerBqpAssertionsRow) HasSubmitterName() bool`

HasSubmitterName returns a boolean if a field has been set.

### SetSubmitterNameNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterNameNil(b bool)`

 SetSubmitterNameNil sets the value for SubmitterName to be an explicit nil

### UnsetSubmitterName
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmitterName()`

UnsetSubmitterName ensures that no value is present for SubmitterName, not even an explicit nil
### GetSubmitterOrgId

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterOrgId() string`

GetSubmitterOrgId returns the SubmitterOrgId field if non-nil, zero value otherwise.

### GetSubmitterOrgIdOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterOrgIdOk() (*string, bool)`

GetSubmitterOrgIdOk returns a tuple with the SubmitterOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterOrgId

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterOrgId(v string)`

SetSubmitterOrgId sets SubmitterOrgId field to given value.

### HasSubmitterOrgId

`func (o *BiomarkerBqpAssertionsRow) HasSubmitterOrgId() bool`

HasSubmitterOrgId returns a boolean if a field has been set.

### SetSubmitterOrgIdNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterOrgIdNil(b bool)`

 SetSubmitterOrgIdNil sets the value for SubmitterOrgId to be an explicit nil

### UnsetSubmitterOrgId
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmitterOrgId()`

UnsetSubmitterOrgId ensures that no value is present for SubmitterOrgId, not even an explicit nil
### GetSubmitterCompanyId

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterCompanyId() string`

GetSubmitterCompanyId returns the SubmitterCompanyId field if non-nil, zero value otherwise.

### GetSubmitterCompanyIdOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterCompanyIdOk() (*string, bool)`

GetSubmitterCompanyIdOk returns a tuple with the SubmitterCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterCompanyId

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterCompanyId(v string)`

SetSubmitterCompanyId sets SubmitterCompanyId field to given value.

### HasSubmitterCompanyId

`func (o *BiomarkerBqpAssertionsRow) HasSubmitterCompanyId() bool`

HasSubmitterCompanyId returns a boolean if a field has been set.

### SetSubmitterCompanyIdNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterCompanyIdNil(b bool)`

 SetSubmitterCompanyIdNil sets the value for SubmitterCompanyId to be an explicit nil

### UnsetSubmitterCompanyId
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmitterCompanyId()`

UnsetSubmitterCompanyId ensures that no value is present for SubmitterCompanyId, not even an explicit nil
### GetSubmitterCompanyName

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterCompanyName() string`

GetSubmitterCompanyName returns the SubmitterCompanyName field if non-nil, zero value otherwise.

### GetSubmitterCompanyNameOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterCompanyNameOk() (*string, bool)`

GetSubmitterCompanyNameOk returns a tuple with the SubmitterCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterCompanyName

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterCompanyName(v string)`

SetSubmitterCompanyName sets SubmitterCompanyName field to given value.

### HasSubmitterCompanyName

`func (o *BiomarkerBqpAssertionsRow) HasSubmitterCompanyName() bool`

HasSubmitterCompanyName returns a boolean if a field has been set.

### SetSubmitterCompanyNameNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterCompanyNameNil(b bool)`

 SetSubmitterCompanyNameNil sets the value for SubmitterCompanyName to be an explicit nil

### UnsetSubmitterCompanyName
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmitterCompanyName()`

UnsetSubmitterCompanyName ensures that no value is present for SubmitterCompanyName, not even an explicit nil
### GetSubmitterGroupId

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterGroupId() string`

GetSubmitterGroupId returns the SubmitterGroupId field if non-nil, zero value otherwise.

### GetSubmitterGroupIdOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterGroupIdOk() (*string, bool)`

GetSubmitterGroupIdOk returns a tuple with the SubmitterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterGroupId

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterGroupId(v string)`

SetSubmitterGroupId sets SubmitterGroupId field to given value.

### HasSubmitterGroupId

`func (o *BiomarkerBqpAssertionsRow) HasSubmitterGroupId() bool`

HasSubmitterGroupId returns a boolean if a field has been set.

### SetSubmitterGroupIdNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterGroupIdNil(b bool)`

 SetSubmitterGroupIdNil sets the value for SubmitterGroupId to be an explicit nil

### UnsetSubmitterGroupId
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmitterGroupId()`

UnsetSubmitterGroupId ensures that no value is present for SubmitterGroupId, not even an explicit nil
### GetSubmitterGroupName

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterGroupName() string`

GetSubmitterGroupName returns the SubmitterGroupName field if non-nil, zero value otherwise.

### GetSubmitterGroupNameOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterGroupNameOk() (*string, bool)`

GetSubmitterGroupNameOk returns a tuple with the SubmitterGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterGroupName

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterGroupName(v string)`

SetSubmitterGroupName sets SubmitterGroupName field to given value.

### HasSubmitterGroupName

`func (o *BiomarkerBqpAssertionsRow) HasSubmitterGroupName() bool`

HasSubmitterGroupName returns a boolean if a field has been set.

### SetSubmitterGroupNameNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterGroupNameNil(b bool)`

 SetSubmitterGroupNameNil sets the value for SubmitterGroupName to be an explicit nil

### UnsetSubmitterGroupName
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmitterGroupName()`

UnsetSubmitterGroupName ensures that no value is present for SubmitterGroupName, not even an explicit nil
### GetSubmitterMatchTier

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterMatchTier() string`

GetSubmitterMatchTier returns the SubmitterMatchTier field if non-nil, zero value otherwise.

### GetSubmitterMatchTierOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterMatchTierOk() (*string, bool)`

GetSubmitterMatchTierOk returns a tuple with the SubmitterMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterMatchTier

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterMatchTier(v string)`

SetSubmitterMatchTier sets SubmitterMatchTier field to given value.

### HasSubmitterMatchTier

`func (o *BiomarkerBqpAssertionsRow) HasSubmitterMatchTier() bool`

HasSubmitterMatchTier returns a boolean if a field has been set.

### SetSubmitterMatchTierNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterMatchTierNil(b bool)`

 SetSubmitterMatchTierNil sets the value for SubmitterMatchTier to be an explicit nil

### UnsetSubmitterMatchTier
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmitterMatchTier()`

UnsetSubmitterMatchTier ensures that no value is present for SubmitterMatchTier, not even an explicit nil
### GetSubmitterRole

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterRole() string`

GetSubmitterRole returns the SubmitterRole field if non-nil, zero value otherwise.

### GetSubmitterRoleOk

`func (o *BiomarkerBqpAssertionsRow) GetSubmitterRoleOk() (*string, bool)`

GetSubmitterRoleOk returns a tuple with the SubmitterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterRole

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterRole(v string)`

SetSubmitterRole sets SubmitterRole field to given value.

### HasSubmitterRole

`func (o *BiomarkerBqpAssertionsRow) HasSubmitterRole() bool`

HasSubmitterRole returns a boolean if a field has been set.

### SetSubmitterRoleNil

`func (o *BiomarkerBqpAssertionsRow) SetSubmitterRoleNil(b bool)`

 SetSubmitterRoleNil sets the value for SubmitterRole to be an explicit nil

### UnsetSubmitterRole
`func (o *BiomarkerBqpAssertionsRow) UnsetSubmitterRole()`

UnsetSubmitterRole ensures that no value is present for SubmitterRole, not even an explicit nil
### GetBiomarkerRole

`func (o *BiomarkerBqpAssertionsRow) GetBiomarkerRole() string`

GetBiomarkerRole returns the BiomarkerRole field if non-nil, zero value otherwise.

### GetBiomarkerRoleOk

`func (o *BiomarkerBqpAssertionsRow) GetBiomarkerRoleOk() (*string, bool)`

GetBiomarkerRoleOk returns a tuple with the BiomarkerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerRole

`func (o *BiomarkerBqpAssertionsRow) SetBiomarkerRole(v string)`

SetBiomarkerRole sets BiomarkerRole field to given value.

### HasBiomarkerRole

`func (o *BiomarkerBqpAssertionsRow) HasBiomarkerRole() bool`

HasBiomarkerRole returns a boolean if a field has been set.

### SetBiomarkerRoleNil

`func (o *BiomarkerBqpAssertionsRow) SetBiomarkerRoleNil(b bool)`

 SetBiomarkerRoleNil sets the value for BiomarkerRole to be an explicit nil

### UnsetBiomarkerRole
`func (o *BiomarkerBqpAssertionsRow) UnsetBiomarkerRole()`

UnsetBiomarkerRole ensures that no value is present for BiomarkerRole, not even an explicit nil
### GetBiomarkerRoleMethod

`func (o *BiomarkerBqpAssertionsRow) GetBiomarkerRoleMethod() string`

GetBiomarkerRoleMethod returns the BiomarkerRoleMethod field if non-nil, zero value otherwise.

### GetBiomarkerRoleMethodOk

`func (o *BiomarkerBqpAssertionsRow) GetBiomarkerRoleMethodOk() (*string, bool)`

GetBiomarkerRoleMethodOk returns a tuple with the BiomarkerRoleMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerRoleMethod

`func (o *BiomarkerBqpAssertionsRow) SetBiomarkerRoleMethod(v string)`

SetBiomarkerRoleMethod sets BiomarkerRoleMethod field to given value.

### HasBiomarkerRoleMethod

`func (o *BiomarkerBqpAssertionsRow) HasBiomarkerRoleMethod() bool`

HasBiomarkerRoleMethod returns a boolean if a field has been set.

### SetBiomarkerRoleMethodNil

`func (o *BiomarkerBqpAssertionsRow) SetBiomarkerRoleMethodNil(b bool)`

 SetBiomarkerRoleMethodNil sets the value for BiomarkerRoleMethod to be an explicit nil

### UnsetBiomarkerRoleMethod
`func (o *BiomarkerBqpAssertionsRow) UnsetBiomarkerRoleMethod()`

UnsetBiomarkerRoleMethod ensures that no value is present for BiomarkerRoleMethod, not even an explicit nil
### GetIsSurrogateEndpointClaim

`func (o *BiomarkerBqpAssertionsRow) GetIsSurrogateEndpointClaim() bool`

GetIsSurrogateEndpointClaim returns the IsSurrogateEndpointClaim field if non-nil, zero value otherwise.

### GetIsSurrogateEndpointClaimOk

`func (o *BiomarkerBqpAssertionsRow) GetIsSurrogateEndpointClaimOk() (*bool, bool)`

GetIsSurrogateEndpointClaimOk returns a tuple with the IsSurrogateEndpointClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSurrogateEndpointClaim

`func (o *BiomarkerBqpAssertionsRow) SetIsSurrogateEndpointClaim(v bool)`

SetIsSurrogateEndpointClaim sets IsSurrogateEndpointClaim field to given value.

### HasIsSurrogateEndpointClaim

`func (o *BiomarkerBqpAssertionsRow) HasIsSurrogateEndpointClaim() bool`

HasIsSurrogateEndpointClaim returns a boolean if a field has been set.

### SetIsSurrogateEndpointClaimNil

`func (o *BiomarkerBqpAssertionsRow) SetIsSurrogateEndpointClaimNil(b bool)`

 SetIsSurrogateEndpointClaimNil sets the value for IsSurrogateEndpointClaim to be an explicit nil

### UnsetIsSurrogateEndpointClaim
`func (o *BiomarkerBqpAssertionsRow) UnsetIsSurrogateEndpointClaim()`

UnsetIsSurrogateEndpointClaim ensures that no value is present for IsSurrogateEndpointClaim, not even an explicit nil
### GetSourceUrl

`func (o *BiomarkerBqpAssertionsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *BiomarkerBqpAssertionsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *BiomarkerBqpAssertionsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *BiomarkerBqpAssertionsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *BiomarkerBqpAssertionsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *BiomarkerBqpAssertionsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSnapshotDate

`func (o *BiomarkerBqpAssertionsRow) GetSnapshotDate() string`

GetSnapshotDate returns the SnapshotDate field if non-nil, zero value otherwise.

### GetSnapshotDateOk

`func (o *BiomarkerBqpAssertionsRow) GetSnapshotDateOk() (*string, bool)`

GetSnapshotDateOk returns a tuple with the SnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDate

`func (o *BiomarkerBqpAssertionsRow) SetSnapshotDate(v string)`

SetSnapshotDate sets SnapshotDate field to given value.

### HasSnapshotDate

`func (o *BiomarkerBqpAssertionsRow) HasSnapshotDate() bool`

HasSnapshotDate returns a boolean if a field has been set.

### SetSnapshotDateNil

`func (o *BiomarkerBqpAssertionsRow) SetSnapshotDateNil(b bool)`

 SetSnapshotDateNil sets the value for SnapshotDate to be an explicit nil

### UnsetSnapshotDate
`func (o *BiomarkerBqpAssertionsRow) UnsetSnapshotDate()`

UnsetSnapshotDate ensures that no value is present for SnapshotDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


