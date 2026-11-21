# RwdUsNetworkProjectsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **NullableString** | The N3C project id (a Data Use Request uid), the grain. | [optional] 
**NetworkKey** | Pointer to **NullableString** | The network key, the entry key of the rwd_us_networks object type. | [optional] 
**Title** | Pointer to **NullableString** | The project title as N3C publishes it. | [optional] 
**InstitutionName** | Pointer to **NullableString** | The accessing institution as published. | [optional] 
**InstitutionNameKey** | Pointer to **NullableString** | The normalized name key; joins to rwd_us_network_memberships.institution_name_key. | [optional] 
**PiName** | Pointer to **NullableString** | Personal data: the principal investigator&#39;s name as N3C&#39;s public dashboard publishes it. Landed whole, tagged, never a gate (#5143). | [optional] 
**WorkspaceStatus** | Pointer to **NullableString** | The project&#39;s workspace status as N3C publishes it. | [optional] 
**ProjectType** | Pointer to **NullableString** | The project type as N3C publishes it. | [optional] 
**IsTaskTeam** | Pointer to **NullableBool** | Whether N3C marks the project as a task team. | [optional] 
**ContainsPersonalData** | Pointer to **NullableBool** | True wherever pi_name is populated. A tag, never a gate (#5143). | [optional] 
**PersonalDataFields** | Pointer to **NullableString** | The personal-data columns of the row, as a JSON list: [\&quot;pi_name\&quot;]. | [optional] 
**AsOf** | Pointer to **NullableString** | The retrieval date of the roster snapshot. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The API endpoint the roster came from. | [optional] 
**RetrievedAt** | Pointer to **NullableTime** | When the snapshot was retrieved, in UTC. | [optional] 
**OrgId** | Pointer to **NullableString** | The organization spine id of the accessing institution, or null. Carried, never minted. | [optional] 
**OrgName** | Pointer to **NullableString** | The spine organization&#39;s name. | [optional] 
**InstitutionRorId** | Pointer to **NullableString** | The bare ROR id N3C published for the institution where it is a network member; null otherwise. | [optional] 
**OrgMatchRung** | Pointer to **NullableString** | How the institution reached the organization: ror_exact, xref, match_keys, ambiguous or unresolved. | [optional] 
**OrgMatchTier** | Pointer to **NullableString** | The tier of the organization match. | [optional] 
**OrgMatchConfidence** | Pointer to **NullableFloat32** | The organization match confidence. | [optional] 
**InstitutionIsNetworkMember** | Pointer to **NullableBool** | Whether the accessing institution itself holds a published agreement. | [optional] 
**PiPersonId** | Pointer to **NullableString** | A resolved person on the spine, by name plus organization; see pi_match_method. The entry key of the investigators object type; null when unresolved. | [optional] 
**PiMatchMethod** | Pointer to **NullableString** | How the PI was matched: name_org_spine (name and organization agreed with one spine person), name_org_reporter (a unique NIH RePORTER PI at the same organization), ambiguous (more than one candidate, never picked) or unresolved. | [optional] 
**PiCorroborated** | Pointer to **NullableBool** | False on every row: no shared trial or paper corroborates the name-plus-organization match. | [optional] 
**PiReporterPiId** | Pointer to **NullableString** | NIH RePORTER&#39;s persistent PI_ID, an external key published as a labelled attribute, never as a resolved person; set where a unique RePORTER PI shares the name and organization. | [optional] 
**CandidatePersonCount** | Pointer to **NullableInt32** | On an ambiguous row, the number of distinct spine persons the evidence could not choose between; otherwise 0. | [optional] 
**AlsoMatchedByReporter** | Pointer to **NullableBool** | True on a name_org_spine row whose person the RePORTER arm also reached. | [optional] 
**ProjectTotal** | Pointer to **NullableInt32** | Projects in the roster, the denominator of the resolution rate. | [optional] 
**ProjectsWithInstitutionOnSpine** | Pointer to **NullableInt32** | Projects whose institution resolved to a spine organization. | [optional] 
**ProjectsWithResolvedPi** | Pointer to **NullableInt32** | Projects whose PI resolved to a spine person; over project_total it is the measured resolution rate. | [optional] 
**ProjectsWithReporterPiId** | Pointer to **NullableInt32** | Projects carrying a RePORTER PI_ID. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit line every surface carrying this data prints. | [optional] 

## Methods

### NewRwdUsNetworkProjectsRow

`func NewRwdUsNetworkProjectsRow() *RwdUsNetworkProjectsRow`

NewRwdUsNetworkProjectsRow instantiates a new RwdUsNetworkProjectsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdUsNetworkProjectsRowWithDefaults

`func NewRwdUsNetworkProjectsRowWithDefaults() *RwdUsNetworkProjectsRow`

NewRwdUsNetworkProjectsRowWithDefaults instantiates a new RwdUsNetworkProjectsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *RwdUsNetworkProjectsRow) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RwdUsNetworkProjectsRow) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RwdUsNetworkProjectsRow) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RwdUsNetworkProjectsRow) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *RwdUsNetworkProjectsRow) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *RwdUsNetworkProjectsRow) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetNetworkKey

`func (o *RwdUsNetworkProjectsRow) GetNetworkKey() string`

GetNetworkKey returns the NetworkKey field if non-nil, zero value otherwise.

### GetNetworkKeyOk

`func (o *RwdUsNetworkProjectsRow) GetNetworkKeyOk() (*string, bool)`

GetNetworkKeyOk returns a tuple with the NetworkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkKey

`func (o *RwdUsNetworkProjectsRow) SetNetworkKey(v string)`

SetNetworkKey sets NetworkKey field to given value.

### HasNetworkKey

`func (o *RwdUsNetworkProjectsRow) HasNetworkKey() bool`

HasNetworkKey returns a boolean if a field has been set.

### SetNetworkKeyNil

`func (o *RwdUsNetworkProjectsRow) SetNetworkKeyNil(b bool)`

 SetNetworkKeyNil sets the value for NetworkKey to be an explicit nil

### UnsetNetworkKey
`func (o *RwdUsNetworkProjectsRow) UnsetNetworkKey()`

UnsetNetworkKey ensures that no value is present for NetworkKey, not even an explicit nil
### GetTitle

`func (o *RwdUsNetworkProjectsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RwdUsNetworkProjectsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RwdUsNetworkProjectsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RwdUsNetworkProjectsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *RwdUsNetworkProjectsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *RwdUsNetworkProjectsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetInstitutionName

`func (o *RwdUsNetworkProjectsRow) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *RwdUsNetworkProjectsRow) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *RwdUsNetworkProjectsRow) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.

### HasInstitutionName

`func (o *RwdUsNetworkProjectsRow) HasInstitutionName() bool`

HasInstitutionName returns a boolean if a field has been set.

### SetInstitutionNameNil

`func (o *RwdUsNetworkProjectsRow) SetInstitutionNameNil(b bool)`

 SetInstitutionNameNil sets the value for InstitutionName to be an explicit nil

### UnsetInstitutionName
`func (o *RwdUsNetworkProjectsRow) UnsetInstitutionName()`

UnsetInstitutionName ensures that no value is present for InstitutionName, not even an explicit nil
### GetInstitutionNameKey

`func (o *RwdUsNetworkProjectsRow) GetInstitutionNameKey() string`

GetInstitutionNameKey returns the InstitutionNameKey field if non-nil, zero value otherwise.

### GetInstitutionNameKeyOk

`func (o *RwdUsNetworkProjectsRow) GetInstitutionNameKeyOk() (*string, bool)`

GetInstitutionNameKeyOk returns a tuple with the InstitutionNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionNameKey

`func (o *RwdUsNetworkProjectsRow) SetInstitutionNameKey(v string)`

SetInstitutionNameKey sets InstitutionNameKey field to given value.

### HasInstitutionNameKey

`func (o *RwdUsNetworkProjectsRow) HasInstitutionNameKey() bool`

HasInstitutionNameKey returns a boolean if a field has been set.

### SetInstitutionNameKeyNil

`func (o *RwdUsNetworkProjectsRow) SetInstitutionNameKeyNil(b bool)`

 SetInstitutionNameKeyNil sets the value for InstitutionNameKey to be an explicit nil

### UnsetInstitutionNameKey
`func (o *RwdUsNetworkProjectsRow) UnsetInstitutionNameKey()`

UnsetInstitutionNameKey ensures that no value is present for InstitutionNameKey, not even an explicit nil
### GetPiName

`func (o *RwdUsNetworkProjectsRow) GetPiName() string`

GetPiName returns the PiName field if non-nil, zero value otherwise.

### GetPiNameOk

`func (o *RwdUsNetworkProjectsRow) GetPiNameOk() (*string, bool)`

GetPiNameOk returns a tuple with the PiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiName

`func (o *RwdUsNetworkProjectsRow) SetPiName(v string)`

SetPiName sets PiName field to given value.

### HasPiName

`func (o *RwdUsNetworkProjectsRow) HasPiName() bool`

HasPiName returns a boolean if a field has been set.

### SetPiNameNil

`func (o *RwdUsNetworkProjectsRow) SetPiNameNil(b bool)`

 SetPiNameNil sets the value for PiName to be an explicit nil

### UnsetPiName
`func (o *RwdUsNetworkProjectsRow) UnsetPiName()`

UnsetPiName ensures that no value is present for PiName, not even an explicit nil
### GetWorkspaceStatus

`func (o *RwdUsNetworkProjectsRow) GetWorkspaceStatus() string`

GetWorkspaceStatus returns the WorkspaceStatus field if non-nil, zero value otherwise.

### GetWorkspaceStatusOk

`func (o *RwdUsNetworkProjectsRow) GetWorkspaceStatusOk() (*string, bool)`

GetWorkspaceStatusOk returns a tuple with the WorkspaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceStatus

`func (o *RwdUsNetworkProjectsRow) SetWorkspaceStatus(v string)`

SetWorkspaceStatus sets WorkspaceStatus field to given value.

### HasWorkspaceStatus

`func (o *RwdUsNetworkProjectsRow) HasWorkspaceStatus() bool`

HasWorkspaceStatus returns a boolean if a field has been set.

### SetWorkspaceStatusNil

`func (o *RwdUsNetworkProjectsRow) SetWorkspaceStatusNil(b bool)`

 SetWorkspaceStatusNil sets the value for WorkspaceStatus to be an explicit nil

### UnsetWorkspaceStatus
`func (o *RwdUsNetworkProjectsRow) UnsetWorkspaceStatus()`

UnsetWorkspaceStatus ensures that no value is present for WorkspaceStatus, not even an explicit nil
### GetProjectType

`func (o *RwdUsNetworkProjectsRow) GetProjectType() string`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *RwdUsNetworkProjectsRow) GetProjectTypeOk() (*string, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *RwdUsNetworkProjectsRow) SetProjectType(v string)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *RwdUsNetworkProjectsRow) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### SetProjectTypeNil

`func (o *RwdUsNetworkProjectsRow) SetProjectTypeNil(b bool)`

 SetProjectTypeNil sets the value for ProjectType to be an explicit nil

### UnsetProjectType
`func (o *RwdUsNetworkProjectsRow) UnsetProjectType()`

UnsetProjectType ensures that no value is present for ProjectType, not even an explicit nil
### GetIsTaskTeam

`func (o *RwdUsNetworkProjectsRow) GetIsTaskTeam() bool`

GetIsTaskTeam returns the IsTaskTeam field if non-nil, zero value otherwise.

### GetIsTaskTeamOk

`func (o *RwdUsNetworkProjectsRow) GetIsTaskTeamOk() (*bool, bool)`

GetIsTaskTeamOk returns a tuple with the IsTaskTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaskTeam

`func (o *RwdUsNetworkProjectsRow) SetIsTaskTeam(v bool)`

SetIsTaskTeam sets IsTaskTeam field to given value.

### HasIsTaskTeam

`func (o *RwdUsNetworkProjectsRow) HasIsTaskTeam() bool`

HasIsTaskTeam returns a boolean if a field has been set.

### SetIsTaskTeamNil

`func (o *RwdUsNetworkProjectsRow) SetIsTaskTeamNil(b bool)`

 SetIsTaskTeamNil sets the value for IsTaskTeam to be an explicit nil

### UnsetIsTaskTeam
`func (o *RwdUsNetworkProjectsRow) UnsetIsTaskTeam()`

UnsetIsTaskTeam ensures that no value is present for IsTaskTeam, not even an explicit nil
### GetContainsPersonalData

`func (o *RwdUsNetworkProjectsRow) GetContainsPersonalData() bool`

GetContainsPersonalData returns the ContainsPersonalData field if non-nil, zero value otherwise.

### GetContainsPersonalDataOk

`func (o *RwdUsNetworkProjectsRow) GetContainsPersonalDataOk() (*bool, bool)`

GetContainsPersonalDataOk returns a tuple with the ContainsPersonalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsPersonalData

`func (o *RwdUsNetworkProjectsRow) SetContainsPersonalData(v bool)`

SetContainsPersonalData sets ContainsPersonalData field to given value.

### HasContainsPersonalData

`func (o *RwdUsNetworkProjectsRow) HasContainsPersonalData() bool`

HasContainsPersonalData returns a boolean if a field has been set.

### SetContainsPersonalDataNil

`func (o *RwdUsNetworkProjectsRow) SetContainsPersonalDataNil(b bool)`

 SetContainsPersonalDataNil sets the value for ContainsPersonalData to be an explicit nil

### UnsetContainsPersonalData
`func (o *RwdUsNetworkProjectsRow) UnsetContainsPersonalData()`

UnsetContainsPersonalData ensures that no value is present for ContainsPersonalData, not even an explicit nil
### GetPersonalDataFields

`func (o *RwdUsNetworkProjectsRow) GetPersonalDataFields() string`

GetPersonalDataFields returns the PersonalDataFields field if non-nil, zero value otherwise.

### GetPersonalDataFieldsOk

`func (o *RwdUsNetworkProjectsRow) GetPersonalDataFieldsOk() (*string, bool)`

GetPersonalDataFieldsOk returns a tuple with the PersonalDataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDataFields

`func (o *RwdUsNetworkProjectsRow) SetPersonalDataFields(v string)`

SetPersonalDataFields sets PersonalDataFields field to given value.

### HasPersonalDataFields

`func (o *RwdUsNetworkProjectsRow) HasPersonalDataFields() bool`

HasPersonalDataFields returns a boolean if a field has been set.

### SetPersonalDataFieldsNil

`func (o *RwdUsNetworkProjectsRow) SetPersonalDataFieldsNil(b bool)`

 SetPersonalDataFieldsNil sets the value for PersonalDataFields to be an explicit nil

### UnsetPersonalDataFields
`func (o *RwdUsNetworkProjectsRow) UnsetPersonalDataFields()`

UnsetPersonalDataFields ensures that no value is present for PersonalDataFields, not even an explicit nil
### GetAsOf

`func (o *RwdUsNetworkProjectsRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *RwdUsNetworkProjectsRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *RwdUsNetworkProjectsRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *RwdUsNetworkProjectsRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *RwdUsNetworkProjectsRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *RwdUsNetworkProjectsRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetSourceUrl

`func (o *RwdUsNetworkProjectsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RwdUsNetworkProjectsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RwdUsNetworkProjectsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RwdUsNetworkProjectsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RwdUsNetworkProjectsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RwdUsNetworkProjectsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetRetrievedAt

`func (o *RwdUsNetworkProjectsRow) GetRetrievedAt() time.Time`

GetRetrievedAt returns the RetrievedAt field if non-nil, zero value otherwise.

### GetRetrievedAtOk

`func (o *RwdUsNetworkProjectsRow) GetRetrievedAtOk() (*time.Time, bool)`

GetRetrievedAtOk returns a tuple with the RetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedAt

`func (o *RwdUsNetworkProjectsRow) SetRetrievedAt(v time.Time)`

SetRetrievedAt sets RetrievedAt field to given value.

### HasRetrievedAt

`func (o *RwdUsNetworkProjectsRow) HasRetrievedAt() bool`

HasRetrievedAt returns a boolean if a field has been set.

### SetRetrievedAtNil

`func (o *RwdUsNetworkProjectsRow) SetRetrievedAtNil(b bool)`

 SetRetrievedAtNil sets the value for RetrievedAt to be an explicit nil

### UnsetRetrievedAt
`func (o *RwdUsNetworkProjectsRow) UnsetRetrievedAt()`

UnsetRetrievedAt ensures that no value is present for RetrievedAt, not even an explicit nil
### GetOrgId

`func (o *RwdUsNetworkProjectsRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RwdUsNetworkProjectsRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RwdUsNetworkProjectsRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *RwdUsNetworkProjectsRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *RwdUsNetworkProjectsRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *RwdUsNetworkProjectsRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOrgName

`func (o *RwdUsNetworkProjectsRow) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *RwdUsNetworkProjectsRow) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *RwdUsNetworkProjectsRow) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *RwdUsNetworkProjectsRow) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *RwdUsNetworkProjectsRow) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *RwdUsNetworkProjectsRow) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetInstitutionRorId

`func (o *RwdUsNetworkProjectsRow) GetInstitutionRorId() string`

GetInstitutionRorId returns the InstitutionRorId field if non-nil, zero value otherwise.

### GetInstitutionRorIdOk

`func (o *RwdUsNetworkProjectsRow) GetInstitutionRorIdOk() (*string, bool)`

GetInstitutionRorIdOk returns a tuple with the InstitutionRorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionRorId

`func (o *RwdUsNetworkProjectsRow) SetInstitutionRorId(v string)`

SetInstitutionRorId sets InstitutionRorId field to given value.

### HasInstitutionRorId

`func (o *RwdUsNetworkProjectsRow) HasInstitutionRorId() bool`

HasInstitutionRorId returns a boolean if a field has been set.

### SetInstitutionRorIdNil

`func (o *RwdUsNetworkProjectsRow) SetInstitutionRorIdNil(b bool)`

 SetInstitutionRorIdNil sets the value for InstitutionRorId to be an explicit nil

### UnsetInstitutionRorId
`func (o *RwdUsNetworkProjectsRow) UnsetInstitutionRorId()`

UnsetInstitutionRorId ensures that no value is present for InstitutionRorId, not even an explicit nil
### GetOrgMatchRung

`func (o *RwdUsNetworkProjectsRow) GetOrgMatchRung() string`

GetOrgMatchRung returns the OrgMatchRung field if non-nil, zero value otherwise.

### GetOrgMatchRungOk

`func (o *RwdUsNetworkProjectsRow) GetOrgMatchRungOk() (*string, bool)`

GetOrgMatchRungOk returns a tuple with the OrgMatchRung field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMatchRung

`func (o *RwdUsNetworkProjectsRow) SetOrgMatchRung(v string)`

SetOrgMatchRung sets OrgMatchRung field to given value.

### HasOrgMatchRung

`func (o *RwdUsNetworkProjectsRow) HasOrgMatchRung() bool`

HasOrgMatchRung returns a boolean if a field has been set.

### SetOrgMatchRungNil

`func (o *RwdUsNetworkProjectsRow) SetOrgMatchRungNil(b bool)`

 SetOrgMatchRungNil sets the value for OrgMatchRung to be an explicit nil

### UnsetOrgMatchRung
`func (o *RwdUsNetworkProjectsRow) UnsetOrgMatchRung()`

UnsetOrgMatchRung ensures that no value is present for OrgMatchRung, not even an explicit nil
### GetOrgMatchTier

`func (o *RwdUsNetworkProjectsRow) GetOrgMatchTier() string`

GetOrgMatchTier returns the OrgMatchTier field if non-nil, zero value otherwise.

### GetOrgMatchTierOk

`func (o *RwdUsNetworkProjectsRow) GetOrgMatchTierOk() (*string, bool)`

GetOrgMatchTierOk returns a tuple with the OrgMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMatchTier

`func (o *RwdUsNetworkProjectsRow) SetOrgMatchTier(v string)`

SetOrgMatchTier sets OrgMatchTier field to given value.

### HasOrgMatchTier

`func (o *RwdUsNetworkProjectsRow) HasOrgMatchTier() bool`

HasOrgMatchTier returns a boolean if a field has been set.

### SetOrgMatchTierNil

`func (o *RwdUsNetworkProjectsRow) SetOrgMatchTierNil(b bool)`

 SetOrgMatchTierNil sets the value for OrgMatchTier to be an explicit nil

### UnsetOrgMatchTier
`func (o *RwdUsNetworkProjectsRow) UnsetOrgMatchTier()`

UnsetOrgMatchTier ensures that no value is present for OrgMatchTier, not even an explicit nil
### GetOrgMatchConfidence

`func (o *RwdUsNetworkProjectsRow) GetOrgMatchConfidence() float32`

GetOrgMatchConfidence returns the OrgMatchConfidence field if non-nil, zero value otherwise.

### GetOrgMatchConfidenceOk

`func (o *RwdUsNetworkProjectsRow) GetOrgMatchConfidenceOk() (*float32, bool)`

GetOrgMatchConfidenceOk returns a tuple with the OrgMatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMatchConfidence

`func (o *RwdUsNetworkProjectsRow) SetOrgMatchConfidence(v float32)`

SetOrgMatchConfidence sets OrgMatchConfidence field to given value.

### HasOrgMatchConfidence

`func (o *RwdUsNetworkProjectsRow) HasOrgMatchConfidence() bool`

HasOrgMatchConfidence returns a boolean if a field has been set.

### SetOrgMatchConfidenceNil

`func (o *RwdUsNetworkProjectsRow) SetOrgMatchConfidenceNil(b bool)`

 SetOrgMatchConfidenceNil sets the value for OrgMatchConfidence to be an explicit nil

### UnsetOrgMatchConfidence
`func (o *RwdUsNetworkProjectsRow) UnsetOrgMatchConfidence()`

UnsetOrgMatchConfidence ensures that no value is present for OrgMatchConfidence, not even an explicit nil
### GetInstitutionIsNetworkMember

`func (o *RwdUsNetworkProjectsRow) GetInstitutionIsNetworkMember() bool`

GetInstitutionIsNetworkMember returns the InstitutionIsNetworkMember field if non-nil, zero value otherwise.

### GetInstitutionIsNetworkMemberOk

`func (o *RwdUsNetworkProjectsRow) GetInstitutionIsNetworkMemberOk() (*bool, bool)`

GetInstitutionIsNetworkMemberOk returns a tuple with the InstitutionIsNetworkMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionIsNetworkMember

`func (o *RwdUsNetworkProjectsRow) SetInstitutionIsNetworkMember(v bool)`

SetInstitutionIsNetworkMember sets InstitutionIsNetworkMember field to given value.

### HasInstitutionIsNetworkMember

`func (o *RwdUsNetworkProjectsRow) HasInstitutionIsNetworkMember() bool`

HasInstitutionIsNetworkMember returns a boolean if a field has been set.

### SetInstitutionIsNetworkMemberNil

`func (o *RwdUsNetworkProjectsRow) SetInstitutionIsNetworkMemberNil(b bool)`

 SetInstitutionIsNetworkMemberNil sets the value for InstitutionIsNetworkMember to be an explicit nil

### UnsetInstitutionIsNetworkMember
`func (o *RwdUsNetworkProjectsRow) UnsetInstitutionIsNetworkMember()`

UnsetInstitutionIsNetworkMember ensures that no value is present for InstitutionIsNetworkMember, not even an explicit nil
### GetPiPersonId

`func (o *RwdUsNetworkProjectsRow) GetPiPersonId() string`

GetPiPersonId returns the PiPersonId field if non-nil, zero value otherwise.

### GetPiPersonIdOk

`func (o *RwdUsNetworkProjectsRow) GetPiPersonIdOk() (*string, bool)`

GetPiPersonIdOk returns a tuple with the PiPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiPersonId

`func (o *RwdUsNetworkProjectsRow) SetPiPersonId(v string)`

SetPiPersonId sets PiPersonId field to given value.

### HasPiPersonId

`func (o *RwdUsNetworkProjectsRow) HasPiPersonId() bool`

HasPiPersonId returns a boolean if a field has been set.

### SetPiPersonIdNil

`func (o *RwdUsNetworkProjectsRow) SetPiPersonIdNil(b bool)`

 SetPiPersonIdNil sets the value for PiPersonId to be an explicit nil

### UnsetPiPersonId
`func (o *RwdUsNetworkProjectsRow) UnsetPiPersonId()`

UnsetPiPersonId ensures that no value is present for PiPersonId, not even an explicit nil
### GetPiMatchMethod

`func (o *RwdUsNetworkProjectsRow) GetPiMatchMethod() string`

GetPiMatchMethod returns the PiMatchMethod field if non-nil, zero value otherwise.

### GetPiMatchMethodOk

`func (o *RwdUsNetworkProjectsRow) GetPiMatchMethodOk() (*string, bool)`

GetPiMatchMethodOk returns a tuple with the PiMatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiMatchMethod

`func (o *RwdUsNetworkProjectsRow) SetPiMatchMethod(v string)`

SetPiMatchMethod sets PiMatchMethod field to given value.

### HasPiMatchMethod

`func (o *RwdUsNetworkProjectsRow) HasPiMatchMethod() bool`

HasPiMatchMethod returns a boolean if a field has been set.

### SetPiMatchMethodNil

`func (o *RwdUsNetworkProjectsRow) SetPiMatchMethodNil(b bool)`

 SetPiMatchMethodNil sets the value for PiMatchMethod to be an explicit nil

### UnsetPiMatchMethod
`func (o *RwdUsNetworkProjectsRow) UnsetPiMatchMethod()`

UnsetPiMatchMethod ensures that no value is present for PiMatchMethod, not even an explicit nil
### GetPiCorroborated

`func (o *RwdUsNetworkProjectsRow) GetPiCorroborated() bool`

GetPiCorroborated returns the PiCorroborated field if non-nil, zero value otherwise.

### GetPiCorroboratedOk

`func (o *RwdUsNetworkProjectsRow) GetPiCorroboratedOk() (*bool, bool)`

GetPiCorroboratedOk returns a tuple with the PiCorroborated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiCorroborated

`func (o *RwdUsNetworkProjectsRow) SetPiCorroborated(v bool)`

SetPiCorroborated sets PiCorroborated field to given value.

### HasPiCorroborated

`func (o *RwdUsNetworkProjectsRow) HasPiCorroborated() bool`

HasPiCorroborated returns a boolean if a field has been set.

### SetPiCorroboratedNil

`func (o *RwdUsNetworkProjectsRow) SetPiCorroboratedNil(b bool)`

 SetPiCorroboratedNil sets the value for PiCorroborated to be an explicit nil

### UnsetPiCorroborated
`func (o *RwdUsNetworkProjectsRow) UnsetPiCorroborated()`

UnsetPiCorroborated ensures that no value is present for PiCorroborated, not even an explicit nil
### GetPiReporterPiId

`func (o *RwdUsNetworkProjectsRow) GetPiReporterPiId() string`

GetPiReporterPiId returns the PiReporterPiId field if non-nil, zero value otherwise.

### GetPiReporterPiIdOk

`func (o *RwdUsNetworkProjectsRow) GetPiReporterPiIdOk() (*string, bool)`

GetPiReporterPiIdOk returns a tuple with the PiReporterPiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiReporterPiId

`func (o *RwdUsNetworkProjectsRow) SetPiReporterPiId(v string)`

SetPiReporterPiId sets PiReporterPiId field to given value.

### HasPiReporterPiId

`func (o *RwdUsNetworkProjectsRow) HasPiReporterPiId() bool`

HasPiReporterPiId returns a boolean if a field has been set.

### SetPiReporterPiIdNil

`func (o *RwdUsNetworkProjectsRow) SetPiReporterPiIdNil(b bool)`

 SetPiReporterPiIdNil sets the value for PiReporterPiId to be an explicit nil

### UnsetPiReporterPiId
`func (o *RwdUsNetworkProjectsRow) UnsetPiReporterPiId()`

UnsetPiReporterPiId ensures that no value is present for PiReporterPiId, not even an explicit nil
### GetCandidatePersonCount

`func (o *RwdUsNetworkProjectsRow) GetCandidatePersonCount() int32`

GetCandidatePersonCount returns the CandidatePersonCount field if non-nil, zero value otherwise.

### GetCandidatePersonCountOk

`func (o *RwdUsNetworkProjectsRow) GetCandidatePersonCountOk() (*int32, bool)`

GetCandidatePersonCountOk returns a tuple with the CandidatePersonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidatePersonCount

`func (o *RwdUsNetworkProjectsRow) SetCandidatePersonCount(v int32)`

SetCandidatePersonCount sets CandidatePersonCount field to given value.

### HasCandidatePersonCount

`func (o *RwdUsNetworkProjectsRow) HasCandidatePersonCount() bool`

HasCandidatePersonCount returns a boolean if a field has been set.

### SetCandidatePersonCountNil

`func (o *RwdUsNetworkProjectsRow) SetCandidatePersonCountNil(b bool)`

 SetCandidatePersonCountNil sets the value for CandidatePersonCount to be an explicit nil

### UnsetCandidatePersonCount
`func (o *RwdUsNetworkProjectsRow) UnsetCandidatePersonCount()`

UnsetCandidatePersonCount ensures that no value is present for CandidatePersonCount, not even an explicit nil
### GetAlsoMatchedByReporter

`func (o *RwdUsNetworkProjectsRow) GetAlsoMatchedByReporter() bool`

GetAlsoMatchedByReporter returns the AlsoMatchedByReporter field if non-nil, zero value otherwise.

### GetAlsoMatchedByReporterOk

`func (o *RwdUsNetworkProjectsRow) GetAlsoMatchedByReporterOk() (*bool, bool)`

GetAlsoMatchedByReporterOk returns a tuple with the AlsoMatchedByReporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlsoMatchedByReporter

`func (o *RwdUsNetworkProjectsRow) SetAlsoMatchedByReporter(v bool)`

SetAlsoMatchedByReporter sets AlsoMatchedByReporter field to given value.

### HasAlsoMatchedByReporter

`func (o *RwdUsNetworkProjectsRow) HasAlsoMatchedByReporter() bool`

HasAlsoMatchedByReporter returns a boolean if a field has been set.

### SetAlsoMatchedByReporterNil

`func (o *RwdUsNetworkProjectsRow) SetAlsoMatchedByReporterNil(b bool)`

 SetAlsoMatchedByReporterNil sets the value for AlsoMatchedByReporter to be an explicit nil

### UnsetAlsoMatchedByReporter
`func (o *RwdUsNetworkProjectsRow) UnsetAlsoMatchedByReporter()`

UnsetAlsoMatchedByReporter ensures that no value is present for AlsoMatchedByReporter, not even an explicit nil
### GetProjectTotal

`func (o *RwdUsNetworkProjectsRow) GetProjectTotal() int32`

GetProjectTotal returns the ProjectTotal field if non-nil, zero value otherwise.

### GetProjectTotalOk

`func (o *RwdUsNetworkProjectsRow) GetProjectTotalOk() (*int32, bool)`

GetProjectTotalOk returns a tuple with the ProjectTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTotal

`func (o *RwdUsNetworkProjectsRow) SetProjectTotal(v int32)`

SetProjectTotal sets ProjectTotal field to given value.

### HasProjectTotal

`func (o *RwdUsNetworkProjectsRow) HasProjectTotal() bool`

HasProjectTotal returns a boolean if a field has been set.

### SetProjectTotalNil

`func (o *RwdUsNetworkProjectsRow) SetProjectTotalNil(b bool)`

 SetProjectTotalNil sets the value for ProjectTotal to be an explicit nil

### UnsetProjectTotal
`func (o *RwdUsNetworkProjectsRow) UnsetProjectTotal()`

UnsetProjectTotal ensures that no value is present for ProjectTotal, not even an explicit nil
### GetProjectsWithInstitutionOnSpine

`func (o *RwdUsNetworkProjectsRow) GetProjectsWithInstitutionOnSpine() int32`

GetProjectsWithInstitutionOnSpine returns the ProjectsWithInstitutionOnSpine field if non-nil, zero value otherwise.

### GetProjectsWithInstitutionOnSpineOk

`func (o *RwdUsNetworkProjectsRow) GetProjectsWithInstitutionOnSpineOk() (*int32, bool)`

GetProjectsWithInstitutionOnSpineOk returns a tuple with the ProjectsWithInstitutionOnSpine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsWithInstitutionOnSpine

`func (o *RwdUsNetworkProjectsRow) SetProjectsWithInstitutionOnSpine(v int32)`

SetProjectsWithInstitutionOnSpine sets ProjectsWithInstitutionOnSpine field to given value.

### HasProjectsWithInstitutionOnSpine

`func (o *RwdUsNetworkProjectsRow) HasProjectsWithInstitutionOnSpine() bool`

HasProjectsWithInstitutionOnSpine returns a boolean if a field has been set.

### SetProjectsWithInstitutionOnSpineNil

`func (o *RwdUsNetworkProjectsRow) SetProjectsWithInstitutionOnSpineNil(b bool)`

 SetProjectsWithInstitutionOnSpineNil sets the value for ProjectsWithInstitutionOnSpine to be an explicit nil

### UnsetProjectsWithInstitutionOnSpine
`func (o *RwdUsNetworkProjectsRow) UnsetProjectsWithInstitutionOnSpine()`

UnsetProjectsWithInstitutionOnSpine ensures that no value is present for ProjectsWithInstitutionOnSpine, not even an explicit nil
### GetProjectsWithResolvedPi

`func (o *RwdUsNetworkProjectsRow) GetProjectsWithResolvedPi() int32`

GetProjectsWithResolvedPi returns the ProjectsWithResolvedPi field if non-nil, zero value otherwise.

### GetProjectsWithResolvedPiOk

`func (o *RwdUsNetworkProjectsRow) GetProjectsWithResolvedPiOk() (*int32, bool)`

GetProjectsWithResolvedPiOk returns a tuple with the ProjectsWithResolvedPi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsWithResolvedPi

`func (o *RwdUsNetworkProjectsRow) SetProjectsWithResolvedPi(v int32)`

SetProjectsWithResolvedPi sets ProjectsWithResolvedPi field to given value.

### HasProjectsWithResolvedPi

`func (o *RwdUsNetworkProjectsRow) HasProjectsWithResolvedPi() bool`

HasProjectsWithResolvedPi returns a boolean if a field has been set.

### SetProjectsWithResolvedPiNil

`func (o *RwdUsNetworkProjectsRow) SetProjectsWithResolvedPiNil(b bool)`

 SetProjectsWithResolvedPiNil sets the value for ProjectsWithResolvedPi to be an explicit nil

### UnsetProjectsWithResolvedPi
`func (o *RwdUsNetworkProjectsRow) UnsetProjectsWithResolvedPi()`

UnsetProjectsWithResolvedPi ensures that no value is present for ProjectsWithResolvedPi, not even an explicit nil
### GetProjectsWithReporterPiId

`func (o *RwdUsNetworkProjectsRow) GetProjectsWithReporterPiId() int32`

GetProjectsWithReporterPiId returns the ProjectsWithReporterPiId field if non-nil, zero value otherwise.

### GetProjectsWithReporterPiIdOk

`func (o *RwdUsNetworkProjectsRow) GetProjectsWithReporterPiIdOk() (*int32, bool)`

GetProjectsWithReporterPiIdOk returns a tuple with the ProjectsWithReporterPiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsWithReporterPiId

`func (o *RwdUsNetworkProjectsRow) SetProjectsWithReporterPiId(v int32)`

SetProjectsWithReporterPiId sets ProjectsWithReporterPiId field to given value.

### HasProjectsWithReporterPiId

`func (o *RwdUsNetworkProjectsRow) HasProjectsWithReporterPiId() bool`

HasProjectsWithReporterPiId returns a boolean if a field has been set.

### SetProjectsWithReporterPiIdNil

`func (o *RwdUsNetworkProjectsRow) SetProjectsWithReporterPiIdNil(b bool)`

 SetProjectsWithReporterPiIdNil sets the value for ProjectsWithReporterPiId to be an explicit nil

### UnsetProjectsWithReporterPiId
`func (o *RwdUsNetworkProjectsRow) UnsetProjectsWithReporterPiId()`

UnsetProjectsWithReporterPiId ensures that no value is present for ProjectsWithReporterPiId, not even an explicit nil
### GetAttribution

`func (o *RwdUsNetworkProjectsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdUsNetworkProjectsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdUsNetworkProjectsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdUsNetworkProjectsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdUsNetworkProjectsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdUsNetworkProjectsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


