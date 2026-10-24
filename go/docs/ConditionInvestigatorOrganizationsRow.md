# ConditionInvestigatorOrganizationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionId** | Pointer to **NullableString** | The condition&#39;s MeSH descriptor id, the id in the conditions page address (D001943). The entry key of the conditions object type. | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s MeSH name. | [optional] 
**OrganizationName** | Pointer to **NullableString** | The organization&#39;s name in the registry that resolved it (ROR, else SEC EDGAR), never the investigator&#39;s affiliation text as filed. | [optional] 
**OrganizationOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the investigators&#39; affiliation resolved to. | [optional] 
**OrganizationCompanyId** | Pointer to **NullableString** | Company level: the company the organization belongs to (null when unresolved). | [optional] 
**OrganizationCompanyName** | Pointer to **NullableString** | Display name of the organization&#39;s company. | [optional] 
**OrganizationGroupId** | Pointer to **NullableString** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**OrganizationGroupName** | Pointer to **NullableString** | Display name of the organization&#39;s group. | [optional] 
**NInvestigators** | Pointer to **NullableInt64** | Distinct investigators at the organization named on a trial of the condition, in any role. Always at least 2. | [optional] 
**NPrincipalInvestigators** | Pointer to **NullableInt64** | Of n_investigators, those who hold the principal investigator role on one of the condition&#39;s trials. | [optional] 
**NInvestigatorsTotal** | Pointer to **NullableInt64** | Distinct investigators across the condition&#39;s trials, at any organization or none. A constant of the condition repeated on every row: never sum it. | [optional] 
**NAtResolvedOrganizations** | Pointer to **NullableInt64** | Of n_investigators_total, those whose affiliation resolved to an organization at a high-confidence tier, listed or not. A constant of the condition. | [optional] 
**NAtUnresolvedAffiliations** | Pointer to **NullableInt64** | Of n_investigators_total, those with no resolved organization, including those with no affiliation on file. A constant of the condition. | [optional] 
**NPrincipalInvestigatorsTotal** | Pointer to **NullableInt64** | Of n_investigators_total, those who hold the principal investigator role on one of the condition&#39;s trials. A constant of the condition. | [optional] 

## Methods

### NewConditionInvestigatorOrganizationsRow

`func NewConditionInvestigatorOrganizationsRow() *ConditionInvestigatorOrganizationsRow`

NewConditionInvestigatorOrganizationsRow instantiates a new ConditionInvestigatorOrganizationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionInvestigatorOrganizationsRowWithDefaults

`func NewConditionInvestigatorOrganizationsRowWithDefaults() *ConditionInvestigatorOrganizationsRow`

NewConditionInvestigatorOrganizationsRowWithDefaults instantiates a new ConditionInvestigatorOrganizationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionId

`func (o *ConditionInvestigatorOrganizationsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionInvestigatorOrganizationsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionInvestigatorOrganizationsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionInvestigatorOrganizationsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ConditionInvestigatorOrganizationsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ConditionInvestigatorOrganizationsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ConditionInvestigatorOrganizationsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ConditionInvestigatorOrganizationsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ConditionInvestigatorOrganizationsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ConditionInvestigatorOrganizationsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ConditionInvestigatorOrganizationsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ConditionInvestigatorOrganizationsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetOrganizationName

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ConditionInvestigatorOrganizationsRow) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ConditionInvestigatorOrganizationsRow) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationOrgId

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationOrgId() string`

GetOrganizationOrgId returns the OrganizationOrgId field if non-nil, zero value otherwise.

### GetOrganizationOrgIdOk

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationOrgIdOk() (*string, bool)`

GetOrganizationOrgIdOk returns a tuple with the OrganizationOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationOrgId

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationOrgId(v string)`

SetOrganizationOrgId sets OrganizationOrgId field to given value.

### HasOrganizationOrgId

`func (o *ConditionInvestigatorOrganizationsRow) HasOrganizationOrgId() bool`

HasOrganizationOrgId returns a boolean if a field has been set.

### SetOrganizationOrgIdNil

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationOrgIdNil(b bool)`

 SetOrganizationOrgIdNil sets the value for OrganizationOrgId to be an explicit nil

### UnsetOrganizationOrgId
`func (o *ConditionInvestigatorOrganizationsRow) UnsetOrganizationOrgId()`

UnsetOrganizationOrgId ensures that no value is present for OrganizationOrgId, not even an explicit nil
### GetOrganizationCompanyId

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationCompanyId() string`

GetOrganizationCompanyId returns the OrganizationCompanyId field if non-nil, zero value otherwise.

### GetOrganizationCompanyIdOk

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationCompanyIdOk() (*string, bool)`

GetOrganizationCompanyIdOk returns a tuple with the OrganizationCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyId

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationCompanyId(v string)`

SetOrganizationCompanyId sets OrganizationCompanyId field to given value.

### HasOrganizationCompanyId

`func (o *ConditionInvestigatorOrganizationsRow) HasOrganizationCompanyId() bool`

HasOrganizationCompanyId returns a boolean if a field has been set.

### SetOrganizationCompanyIdNil

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationCompanyIdNil(b bool)`

 SetOrganizationCompanyIdNil sets the value for OrganizationCompanyId to be an explicit nil

### UnsetOrganizationCompanyId
`func (o *ConditionInvestigatorOrganizationsRow) UnsetOrganizationCompanyId()`

UnsetOrganizationCompanyId ensures that no value is present for OrganizationCompanyId, not even an explicit nil
### GetOrganizationCompanyName

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationCompanyName() string`

GetOrganizationCompanyName returns the OrganizationCompanyName field if non-nil, zero value otherwise.

### GetOrganizationCompanyNameOk

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationCompanyNameOk() (*string, bool)`

GetOrganizationCompanyNameOk returns a tuple with the OrganizationCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyName

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationCompanyName(v string)`

SetOrganizationCompanyName sets OrganizationCompanyName field to given value.

### HasOrganizationCompanyName

`func (o *ConditionInvestigatorOrganizationsRow) HasOrganizationCompanyName() bool`

HasOrganizationCompanyName returns a boolean if a field has been set.

### SetOrganizationCompanyNameNil

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationCompanyNameNil(b bool)`

 SetOrganizationCompanyNameNil sets the value for OrganizationCompanyName to be an explicit nil

### UnsetOrganizationCompanyName
`func (o *ConditionInvestigatorOrganizationsRow) UnsetOrganizationCompanyName()`

UnsetOrganizationCompanyName ensures that no value is present for OrganizationCompanyName, not even an explicit nil
### GetOrganizationGroupId

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationGroupId() string`

GetOrganizationGroupId returns the OrganizationGroupId field if non-nil, zero value otherwise.

### GetOrganizationGroupIdOk

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationGroupIdOk() (*string, bool)`

GetOrganizationGroupIdOk returns a tuple with the OrganizationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupId

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationGroupId(v string)`

SetOrganizationGroupId sets OrganizationGroupId field to given value.

### HasOrganizationGroupId

`func (o *ConditionInvestigatorOrganizationsRow) HasOrganizationGroupId() bool`

HasOrganizationGroupId returns a boolean if a field has been set.

### SetOrganizationGroupIdNil

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationGroupIdNil(b bool)`

 SetOrganizationGroupIdNil sets the value for OrganizationGroupId to be an explicit nil

### UnsetOrganizationGroupId
`func (o *ConditionInvestigatorOrganizationsRow) UnsetOrganizationGroupId()`

UnsetOrganizationGroupId ensures that no value is present for OrganizationGroupId, not even an explicit nil
### GetOrganizationGroupName

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationGroupName() string`

GetOrganizationGroupName returns the OrganizationGroupName field if non-nil, zero value otherwise.

### GetOrganizationGroupNameOk

`func (o *ConditionInvestigatorOrganizationsRow) GetOrganizationGroupNameOk() (*string, bool)`

GetOrganizationGroupNameOk returns a tuple with the OrganizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupName

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationGroupName(v string)`

SetOrganizationGroupName sets OrganizationGroupName field to given value.

### HasOrganizationGroupName

`func (o *ConditionInvestigatorOrganizationsRow) HasOrganizationGroupName() bool`

HasOrganizationGroupName returns a boolean if a field has been set.

### SetOrganizationGroupNameNil

`func (o *ConditionInvestigatorOrganizationsRow) SetOrganizationGroupNameNil(b bool)`

 SetOrganizationGroupNameNil sets the value for OrganizationGroupName to be an explicit nil

### UnsetOrganizationGroupName
`func (o *ConditionInvestigatorOrganizationsRow) UnsetOrganizationGroupName()`

UnsetOrganizationGroupName ensures that no value is present for OrganizationGroupName, not even an explicit nil
### GetNInvestigators

`func (o *ConditionInvestigatorOrganizationsRow) GetNInvestigators() int64`

GetNInvestigators returns the NInvestigators field if non-nil, zero value otherwise.

### GetNInvestigatorsOk

`func (o *ConditionInvestigatorOrganizationsRow) GetNInvestigatorsOk() (*int64, bool)`

GetNInvestigatorsOk returns a tuple with the NInvestigators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNInvestigators

`func (o *ConditionInvestigatorOrganizationsRow) SetNInvestigators(v int64)`

SetNInvestigators sets NInvestigators field to given value.

### HasNInvestigators

`func (o *ConditionInvestigatorOrganizationsRow) HasNInvestigators() bool`

HasNInvestigators returns a boolean if a field has been set.

### SetNInvestigatorsNil

`func (o *ConditionInvestigatorOrganizationsRow) SetNInvestigatorsNil(b bool)`

 SetNInvestigatorsNil sets the value for NInvestigators to be an explicit nil

### UnsetNInvestigators
`func (o *ConditionInvestigatorOrganizationsRow) UnsetNInvestigators()`

UnsetNInvestigators ensures that no value is present for NInvestigators, not even an explicit nil
### GetNPrincipalInvestigators

`func (o *ConditionInvestigatorOrganizationsRow) GetNPrincipalInvestigators() int64`

GetNPrincipalInvestigators returns the NPrincipalInvestigators field if non-nil, zero value otherwise.

### GetNPrincipalInvestigatorsOk

`func (o *ConditionInvestigatorOrganizationsRow) GetNPrincipalInvestigatorsOk() (*int64, bool)`

GetNPrincipalInvestigatorsOk returns a tuple with the NPrincipalInvestigators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPrincipalInvestigators

`func (o *ConditionInvestigatorOrganizationsRow) SetNPrincipalInvestigators(v int64)`

SetNPrincipalInvestigators sets NPrincipalInvestigators field to given value.

### HasNPrincipalInvestigators

`func (o *ConditionInvestigatorOrganizationsRow) HasNPrincipalInvestigators() bool`

HasNPrincipalInvestigators returns a boolean if a field has been set.

### SetNPrincipalInvestigatorsNil

`func (o *ConditionInvestigatorOrganizationsRow) SetNPrincipalInvestigatorsNil(b bool)`

 SetNPrincipalInvestigatorsNil sets the value for NPrincipalInvestigators to be an explicit nil

### UnsetNPrincipalInvestigators
`func (o *ConditionInvestigatorOrganizationsRow) UnsetNPrincipalInvestigators()`

UnsetNPrincipalInvestigators ensures that no value is present for NPrincipalInvestigators, not even an explicit nil
### GetNInvestigatorsTotal

`func (o *ConditionInvestigatorOrganizationsRow) GetNInvestigatorsTotal() int64`

GetNInvestigatorsTotal returns the NInvestigatorsTotal field if non-nil, zero value otherwise.

### GetNInvestigatorsTotalOk

`func (o *ConditionInvestigatorOrganizationsRow) GetNInvestigatorsTotalOk() (*int64, bool)`

GetNInvestigatorsTotalOk returns a tuple with the NInvestigatorsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNInvestigatorsTotal

`func (o *ConditionInvestigatorOrganizationsRow) SetNInvestigatorsTotal(v int64)`

SetNInvestigatorsTotal sets NInvestigatorsTotal field to given value.

### HasNInvestigatorsTotal

`func (o *ConditionInvestigatorOrganizationsRow) HasNInvestigatorsTotal() bool`

HasNInvestigatorsTotal returns a boolean if a field has been set.

### SetNInvestigatorsTotalNil

`func (o *ConditionInvestigatorOrganizationsRow) SetNInvestigatorsTotalNil(b bool)`

 SetNInvestigatorsTotalNil sets the value for NInvestigatorsTotal to be an explicit nil

### UnsetNInvestigatorsTotal
`func (o *ConditionInvestigatorOrganizationsRow) UnsetNInvestigatorsTotal()`

UnsetNInvestigatorsTotal ensures that no value is present for NInvestigatorsTotal, not even an explicit nil
### GetNAtResolvedOrganizations

`func (o *ConditionInvestigatorOrganizationsRow) GetNAtResolvedOrganizations() int64`

GetNAtResolvedOrganizations returns the NAtResolvedOrganizations field if non-nil, zero value otherwise.

### GetNAtResolvedOrganizationsOk

`func (o *ConditionInvestigatorOrganizationsRow) GetNAtResolvedOrganizationsOk() (*int64, bool)`

GetNAtResolvedOrganizationsOk returns a tuple with the NAtResolvedOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAtResolvedOrganizations

`func (o *ConditionInvestigatorOrganizationsRow) SetNAtResolvedOrganizations(v int64)`

SetNAtResolvedOrganizations sets NAtResolvedOrganizations field to given value.

### HasNAtResolvedOrganizations

`func (o *ConditionInvestigatorOrganizationsRow) HasNAtResolvedOrganizations() bool`

HasNAtResolvedOrganizations returns a boolean if a field has been set.

### SetNAtResolvedOrganizationsNil

`func (o *ConditionInvestigatorOrganizationsRow) SetNAtResolvedOrganizationsNil(b bool)`

 SetNAtResolvedOrganizationsNil sets the value for NAtResolvedOrganizations to be an explicit nil

### UnsetNAtResolvedOrganizations
`func (o *ConditionInvestigatorOrganizationsRow) UnsetNAtResolvedOrganizations()`

UnsetNAtResolvedOrganizations ensures that no value is present for NAtResolvedOrganizations, not even an explicit nil
### GetNAtUnresolvedAffiliations

`func (o *ConditionInvestigatorOrganizationsRow) GetNAtUnresolvedAffiliations() int64`

GetNAtUnresolvedAffiliations returns the NAtUnresolvedAffiliations field if non-nil, zero value otherwise.

### GetNAtUnresolvedAffiliationsOk

`func (o *ConditionInvestigatorOrganizationsRow) GetNAtUnresolvedAffiliationsOk() (*int64, bool)`

GetNAtUnresolvedAffiliationsOk returns a tuple with the NAtUnresolvedAffiliations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAtUnresolvedAffiliations

`func (o *ConditionInvestigatorOrganizationsRow) SetNAtUnresolvedAffiliations(v int64)`

SetNAtUnresolvedAffiliations sets NAtUnresolvedAffiliations field to given value.

### HasNAtUnresolvedAffiliations

`func (o *ConditionInvestigatorOrganizationsRow) HasNAtUnresolvedAffiliations() bool`

HasNAtUnresolvedAffiliations returns a boolean if a field has been set.

### SetNAtUnresolvedAffiliationsNil

`func (o *ConditionInvestigatorOrganizationsRow) SetNAtUnresolvedAffiliationsNil(b bool)`

 SetNAtUnresolvedAffiliationsNil sets the value for NAtUnresolvedAffiliations to be an explicit nil

### UnsetNAtUnresolvedAffiliations
`func (o *ConditionInvestigatorOrganizationsRow) UnsetNAtUnresolvedAffiliations()`

UnsetNAtUnresolvedAffiliations ensures that no value is present for NAtUnresolvedAffiliations, not even an explicit nil
### GetNPrincipalInvestigatorsTotal

`func (o *ConditionInvestigatorOrganizationsRow) GetNPrincipalInvestigatorsTotal() int64`

GetNPrincipalInvestigatorsTotal returns the NPrincipalInvestigatorsTotal field if non-nil, zero value otherwise.

### GetNPrincipalInvestigatorsTotalOk

`func (o *ConditionInvestigatorOrganizationsRow) GetNPrincipalInvestigatorsTotalOk() (*int64, bool)`

GetNPrincipalInvestigatorsTotalOk returns a tuple with the NPrincipalInvestigatorsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPrincipalInvestigatorsTotal

`func (o *ConditionInvestigatorOrganizationsRow) SetNPrincipalInvestigatorsTotal(v int64)`

SetNPrincipalInvestigatorsTotal sets NPrincipalInvestigatorsTotal field to given value.

### HasNPrincipalInvestigatorsTotal

`func (o *ConditionInvestigatorOrganizationsRow) HasNPrincipalInvestigatorsTotal() bool`

HasNPrincipalInvestigatorsTotal returns a boolean if a field has been set.

### SetNPrincipalInvestigatorsTotalNil

`func (o *ConditionInvestigatorOrganizationsRow) SetNPrincipalInvestigatorsTotalNil(b bool)`

 SetNPrincipalInvestigatorsTotalNil sets the value for NPrincipalInvestigatorsTotal to be an explicit nil

### UnsetNPrincipalInvestigatorsTotal
`func (o *ConditionInvestigatorOrganizationsRow) UnsetNPrincipalInvestigatorsTotal()`

UnsetNPrincipalInvestigatorsTotal ensures that no value is present for NPrincipalInvestigatorsTotal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


