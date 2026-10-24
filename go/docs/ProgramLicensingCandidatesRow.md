# ProgramLicensingCandidatesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the lead sponsor&#39;s organization record. | [optional] 
**SponsorName** | Pointer to **NullableString** | The sponsor&#39;s organization name. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the sponsor&#39;s company (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**DrugId** | Pointer to **NullableString** | The drug key. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s name. | [optional] 
**ConditionAncestorIds** | Pointer to **[]string** | Every MeSH descriptor id at or above the condition in the tree, the condition itself included. Filter a subtree with &#x3D; ANY. | [optional] 
**MaxPhase** | Pointer to **NullableInt32** | Highest phase reached: 0 early phase 1 through 6 for phase 4 (marketed); -1 when unstated. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Lead-sponsored trials in the program. | [optional] 
**NActive** | Pointer to **NullableInt64** | Of those, trials in an active status. | [optional] 
**LastYear** | Pointer to **NullableInt32** | Latest trial start year. | [optional] 
**LastUpdateYear** | Pointer to **NullableInt32** | Year the program&#39;s most recently updated trial record was posted. | [optional] 
**Withdrawn** | Pointer to **NullableBool** | True when the drug has been withdrawn. | [optional] 
**CauseName** | Pointer to **NullableString** | The Global Burden of Disease cause matched to the condition. | [optional] 
**DalysMillions** | Pointer to **NullableFloat64** | Global disability-adjusted life years for the matched cause, in millions. | [optional] 
**DeathsMillions** | Pointer to **NullableFloat64** | Global deaths for the matched cause, in millions. | [optional] 
**Druggability** | Pointer to **NullableInt32** | Druggability class of the matched cause: 1 highly druggable, 2 druggable, 3 challenging, 4 very challenging. | [optional] 

## Methods

### NewProgramLicensingCandidatesRow

`func NewProgramLicensingCandidatesRow() *ProgramLicensingCandidatesRow`

NewProgramLicensingCandidatesRow instantiates a new ProgramLicensingCandidatesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramLicensingCandidatesRowWithDefaults

`func NewProgramLicensingCandidatesRowWithDefaults() *ProgramLicensingCandidatesRow`

NewProgramLicensingCandidatesRowWithDefaults instantiates a new ProgramLicensingCandidatesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponsorOrgId

`func (o *ProgramLicensingCandidatesRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *ProgramLicensingCandidatesRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *ProgramLicensingCandidatesRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *ProgramLicensingCandidatesRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *ProgramLicensingCandidatesRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *ProgramLicensingCandidatesRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorName

`func (o *ProgramLicensingCandidatesRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *ProgramLicensingCandidatesRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *ProgramLicensingCandidatesRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *ProgramLicensingCandidatesRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *ProgramLicensingCandidatesRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *ProgramLicensingCandidatesRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorCompanyId

`func (o *ProgramLicensingCandidatesRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *ProgramLicensingCandidatesRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *ProgramLicensingCandidatesRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *ProgramLicensingCandidatesRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *ProgramLicensingCandidatesRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *ProgramLicensingCandidatesRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *ProgramLicensingCandidatesRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *ProgramLicensingCandidatesRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *ProgramLicensingCandidatesRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *ProgramLicensingCandidatesRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *ProgramLicensingCandidatesRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *ProgramLicensingCandidatesRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *ProgramLicensingCandidatesRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *ProgramLicensingCandidatesRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *ProgramLicensingCandidatesRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *ProgramLicensingCandidatesRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *ProgramLicensingCandidatesRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *ProgramLicensingCandidatesRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *ProgramLicensingCandidatesRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *ProgramLicensingCandidatesRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *ProgramLicensingCandidatesRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *ProgramLicensingCandidatesRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *ProgramLicensingCandidatesRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *ProgramLicensingCandidatesRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetDrugId

`func (o *ProgramLicensingCandidatesRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *ProgramLicensingCandidatesRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *ProgramLicensingCandidatesRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *ProgramLicensingCandidatesRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *ProgramLicensingCandidatesRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *ProgramLicensingCandidatesRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *ProgramLicensingCandidatesRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *ProgramLicensingCandidatesRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *ProgramLicensingCandidatesRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *ProgramLicensingCandidatesRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *ProgramLicensingCandidatesRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *ProgramLicensingCandidatesRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetConditionId

`func (o *ProgramLicensingCandidatesRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ProgramLicensingCandidatesRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ProgramLicensingCandidatesRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ProgramLicensingCandidatesRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *ProgramLicensingCandidatesRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *ProgramLicensingCandidatesRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *ProgramLicensingCandidatesRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ProgramLicensingCandidatesRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ProgramLicensingCandidatesRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *ProgramLicensingCandidatesRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *ProgramLicensingCandidatesRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *ProgramLicensingCandidatesRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetConditionAncestorIds

`func (o *ProgramLicensingCandidatesRow) GetConditionAncestorIds() []string`

GetConditionAncestorIds returns the ConditionAncestorIds field if non-nil, zero value otherwise.

### GetConditionAncestorIdsOk

`func (o *ProgramLicensingCandidatesRow) GetConditionAncestorIdsOk() (*[]string, bool)`

GetConditionAncestorIdsOk returns a tuple with the ConditionAncestorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionAncestorIds

`func (o *ProgramLicensingCandidatesRow) SetConditionAncestorIds(v []string)`

SetConditionAncestorIds sets ConditionAncestorIds field to given value.

### HasConditionAncestorIds

`func (o *ProgramLicensingCandidatesRow) HasConditionAncestorIds() bool`

HasConditionAncestorIds returns a boolean if a field has been set.

### SetConditionAncestorIdsNil

`func (o *ProgramLicensingCandidatesRow) SetConditionAncestorIdsNil(b bool)`

 SetConditionAncestorIdsNil sets the value for ConditionAncestorIds to be an explicit nil

### UnsetConditionAncestorIds
`func (o *ProgramLicensingCandidatesRow) UnsetConditionAncestorIds()`

UnsetConditionAncestorIds ensures that no value is present for ConditionAncestorIds, not even an explicit nil
### GetMaxPhase

`func (o *ProgramLicensingCandidatesRow) GetMaxPhase() int32`

GetMaxPhase returns the MaxPhase field if non-nil, zero value otherwise.

### GetMaxPhaseOk

`func (o *ProgramLicensingCandidatesRow) GetMaxPhaseOk() (*int32, bool)`

GetMaxPhaseOk returns a tuple with the MaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhase

`func (o *ProgramLicensingCandidatesRow) SetMaxPhase(v int32)`

SetMaxPhase sets MaxPhase field to given value.

### HasMaxPhase

`func (o *ProgramLicensingCandidatesRow) HasMaxPhase() bool`

HasMaxPhase returns a boolean if a field has been set.

### SetMaxPhaseNil

`func (o *ProgramLicensingCandidatesRow) SetMaxPhaseNil(b bool)`

 SetMaxPhaseNil sets the value for MaxPhase to be an explicit nil

### UnsetMaxPhase
`func (o *ProgramLicensingCandidatesRow) UnsetMaxPhase()`

UnsetMaxPhase ensures that no value is present for MaxPhase, not even an explicit nil
### GetNTrials

`func (o *ProgramLicensingCandidatesRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *ProgramLicensingCandidatesRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *ProgramLicensingCandidatesRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *ProgramLicensingCandidatesRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *ProgramLicensingCandidatesRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *ProgramLicensingCandidatesRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNActive

`func (o *ProgramLicensingCandidatesRow) GetNActive() int64`

GetNActive returns the NActive field if non-nil, zero value otherwise.

### GetNActiveOk

`func (o *ProgramLicensingCandidatesRow) GetNActiveOk() (*int64, bool)`

GetNActiveOk returns a tuple with the NActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActive

`func (o *ProgramLicensingCandidatesRow) SetNActive(v int64)`

SetNActive sets NActive field to given value.

### HasNActive

`func (o *ProgramLicensingCandidatesRow) HasNActive() bool`

HasNActive returns a boolean if a field has been set.

### SetNActiveNil

`func (o *ProgramLicensingCandidatesRow) SetNActiveNil(b bool)`

 SetNActiveNil sets the value for NActive to be an explicit nil

### UnsetNActive
`func (o *ProgramLicensingCandidatesRow) UnsetNActive()`

UnsetNActive ensures that no value is present for NActive, not even an explicit nil
### GetLastYear

`func (o *ProgramLicensingCandidatesRow) GetLastYear() int32`

GetLastYear returns the LastYear field if non-nil, zero value otherwise.

### GetLastYearOk

`func (o *ProgramLicensingCandidatesRow) GetLastYearOk() (*int32, bool)`

GetLastYearOk returns a tuple with the LastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYear

`func (o *ProgramLicensingCandidatesRow) SetLastYear(v int32)`

SetLastYear sets LastYear field to given value.

### HasLastYear

`func (o *ProgramLicensingCandidatesRow) HasLastYear() bool`

HasLastYear returns a boolean if a field has been set.

### SetLastYearNil

`func (o *ProgramLicensingCandidatesRow) SetLastYearNil(b bool)`

 SetLastYearNil sets the value for LastYear to be an explicit nil

### UnsetLastYear
`func (o *ProgramLicensingCandidatesRow) UnsetLastYear()`

UnsetLastYear ensures that no value is present for LastYear, not even an explicit nil
### GetLastUpdateYear

`func (o *ProgramLicensingCandidatesRow) GetLastUpdateYear() int32`

GetLastUpdateYear returns the LastUpdateYear field if non-nil, zero value otherwise.

### GetLastUpdateYearOk

`func (o *ProgramLicensingCandidatesRow) GetLastUpdateYearOk() (*int32, bool)`

GetLastUpdateYearOk returns a tuple with the LastUpdateYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateYear

`func (o *ProgramLicensingCandidatesRow) SetLastUpdateYear(v int32)`

SetLastUpdateYear sets LastUpdateYear field to given value.

### HasLastUpdateYear

`func (o *ProgramLicensingCandidatesRow) HasLastUpdateYear() bool`

HasLastUpdateYear returns a boolean if a field has been set.

### SetLastUpdateYearNil

`func (o *ProgramLicensingCandidatesRow) SetLastUpdateYearNil(b bool)`

 SetLastUpdateYearNil sets the value for LastUpdateYear to be an explicit nil

### UnsetLastUpdateYear
`func (o *ProgramLicensingCandidatesRow) UnsetLastUpdateYear()`

UnsetLastUpdateYear ensures that no value is present for LastUpdateYear, not even an explicit nil
### GetWithdrawn

`func (o *ProgramLicensingCandidatesRow) GetWithdrawn() bool`

GetWithdrawn returns the Withdrawn field if non-nil, zero value otherwise.

### GetWithdrawnOk

`func (o *ProgramLicensingCandidatesRow) GetWithdrawnOk() (*bool, bool)`

GetWithdrawnOk returns a tuple with the Withdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawn

`func (o *ProgramLicensingCandidatesRow) SetWithdrawn(v bool)`

SetWithdrawn sets Withdrawn field to given value.

### HasWithdrawn

`func (o *ProgramLicensingCandidatesRow) HasWithdrawn() bool`

HasWithdrawn returns a boolean if a field has been set.

### SetWithdrawnNil

`func (o *ProgramLicensingCandidatesRow) SetWithdrawnNil(b bool)`

 SetWithdrawnNil sets the value for Withdrawn to be an explicit nil

### UnsetWithdrawn
`func (o *ProgramLicensingCandidatesRow) UnsetWithdrawn()`

UnsetWithdrawn ensures that no value is present for Withdrawn, not even an explicit nil
### GetCauseName

`func (o *ProgramLicensingCandidatesRow) GetCauseName() string`

GetCauseName returns the CauseName field if non-nil, zero value otherwise.

### GetCauseNameOk

`func (o *ProgramLicensingCandidatesRow) GetCauseNameOk() (*string, bool)`

GetCauseNameOk returns a tuple with the CauseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseName

`func (o *ProgramLicensingCandidatesRow) SetCauseName(v string)`

SetCauseName sets CauseName field to given value.

### HasCauseName

`func (o *ProgramLicensingCandidatesRow) HasCauseName() bool`

HasCauseName returns a boolean if a field has been set.

### SetCauseNameNil

`func (o *ProgramLicensingCandidatesRow) SetCauseNameNil(b bool)`

 SetCauseNameNil sets the value for CauseName to be an explicit nil

### UnsetCauseName
`func (o *ProgramLicensingCandidatesRow) UnsetCauseName()`

UnsetCauseName ensures that no value is present for CauseName, not even an explicit nil
### GetDalysMillions

`func (o *ProgramLicensingCandidatesRow) GetDalysMillions() float64`

GetDalysMillions returns the DalysMillions field if non-nil, zero value otherwise.

### GetDalysMillionsOk

`func (o *ProgramLicensingCandidatesRow) GetDalysMillionsOk() (*float64, bool)`

GetDalysMillionsOk returns a tuple with the DalysMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDalysMillions

`func (o *ProgramLicensingCandidatesRow) SetDalysMillions(v float64)`

SetDalysMillions sets DalysMillions field to given value.

### HasDalysMillions

`func (o *ProgramLicensingCandidatesRow) HasDalysMillions() bool`

HasDalysMillions returns a boolean if a field has been set.

### SetDalysMillionsNil

`func (o *ProgramLicensingCandidatesRow) SetDalysMillionsNil(b bool)`

 SetDalysMillionsNil sets the value for DalysMillions to be an explicit nil

### UnsetDalysMillions
`func (o *ProgramLicensingCandidatesRow) UnsetDalysMillions()`

UnsetDalysMillions ensures that no value is present for DalysMillions, not even an explicit nil
### GetDeathsMillions

`func (o *ProgramLicensingCandidatesRow) GetDeathsMillions() float64`

GetDeathsMillions returns the DeathsMillions field if non-nil, zero value otherwise.

### GetDeathsMillionsOk

`func (o *ProgramLicensingCandidatesRow) GetDeathsMillionsOk() (*float64, bool)`

GetDeathsMillionsOk returns a tuple with the DeathsMillions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathsMillions

`func (o *ProgramLicensingCandidatesRow) SetDeathsMillions(v float64)`

SetDeathsMillions sets DeathsMillions field to given value.

### HasDeathsMillions

`func (o *ProgramLicensingCandidatesRow) HasDeathsMillions() bool`

HasDeathsMillions returns a boolean if a field has been set.

### SetDeathsMillionsNil

`func (o *ProgramLicensingCandidatesRow) SetDeathsMillionsNil(b bool)`

 SetDeathsMillionsNil sets the value for DeathsMillions to be an explicit nil

### UnsetDeathsMillions
`func (o *ProgramLicensingCandidatesRow) UnsetDeathsMillions()`

UnsetDeathsMillions ensures that no value is present for DeathsMillions, not even an explicit nil
### GetDruggability

`func (o *ProgramLicensingCandidatesRow) GetDruggability() int32`

GetDruggability returns the Druggability field if non-nil, zero value otherwise.

### GetDruggabilityOk

`func (o *ProgramLicensingCandidatesRow) GetDruggabilityOk() (*int32, bool)`

GetDruggabilityOk returns a tuple with the Druggability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDruggability

`func (o *ProgramLicensingCandidatesRow) SetDruggability(v int32)`

SetDruggability sets Druggability field to given value.

### HasDruggability

`func (o *ProgramLicensingCandidatesRow) HasDruggability() bool`

HasDruggability returns a boolean if a field has been set.

### SetDruggabilityNil

`func (o *ProgramLicensingCandidatesRow) SetDruggabilityNil(b bool)`

 SetDruggabilityNil sets the value for Druggability to be an explicit nil

### UnsetDruggability
`func (o *ProgramLicensingCandidatesRow) UnsetDruggability()`

UnsetDruggability ensures that no value is present for Druggability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


