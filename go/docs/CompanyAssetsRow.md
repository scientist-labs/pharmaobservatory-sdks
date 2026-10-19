# CompanyAssetsRow

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
**MaxPhase** | Pointer to **NullableInt32** | Highest phase the sponsor reached on the drug: 0 early phase 1 through 6 for phase 4 (marketed); -1 when unstated. | [optional] 
**NTrials** | Pointer to **NullableInt64** | Distinct trials this sponsor leads on the drug. Counted once per trial within the sponsor organization, so it is not additive across a company&#39;s organizations. | [optional] 
**NActive** | Pointer to **NullableInt64** | Of those trials, the number recruiting, not yet recruiting, active or enrolling by invitation. | [optional] 
**LastYear** | Pointer to **NullableInt32** | Latest start year among those trials. | [optional] 
**Withdrawn** | Pointer to **NullableBool** | True when the drug has been withdrawn from the market. | [optional] 

## Methods

### NewCompanyAssetsRow

`func NewCompanyAssetsRow() *CompanyAssetsRow`

NewCompanyAssetsRow instantiates a new CompanyAssetsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyAssetsRowWithDefaults

`func NewCompanyAssetsRowWithDefaults() *CompanyAssetsRow`

NewCompanyAssetsRowWithDefaults instantiates a new CompanyAssetsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponsorOrgId

`func (o *CompanyAssetsRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *CompanyAssetsRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *CompanyAssetsRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *CompanyAssetsRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *CompanyAssetsRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *CompanyAssetsRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorName

`func (o *CompanyAssetsRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *CompanyAssetsRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *CompanyAssetsRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *CompanyAssetsRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *CompanyAssetsRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *CompanyAssetsRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorCompanyId

`func (o *CompanyAssetsRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *CompanyAssetsRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *CompanyAssetsRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *CompanyAssetsRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *CompanyAssetsRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *CompanyAssetsRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *CompanyAssetsRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *CompanyAssetsRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *CompanyAssetsRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *CompanyAssetsRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *CompanyAssetsRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *CompanyAssetsRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *CompanyAssetsRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *CompanyAssetsRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *CompanyAssetsRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *CompanyAssetsRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *CompanyAssetsRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *CompanyAssetsRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *CompanyAssetsRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *CompanyAssetsRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *CompanyAssetsRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *CompanyAssetsRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *CompanyAssetsRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *CompanyAssetsRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetDrugId

`func (o *CompanyAssetsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *CompanyAssetsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *CompanyAssetsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *CompanyAssetsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *CompanyAssetsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *CompanyAssetsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *CompanyAssetsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *CompanyAssetsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *CompanyAssetsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *CompanyAssetsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *CompanyAssetsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *CompanyAssetsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetMaxPhase

`func (o *CompanyAssetsRow) GetMaxPhase() int32`

GetMaxPhase returns the MaxPhase field if non-nil, zero value otherwise.

### GetMaxPhaseOk

`func (o *CompanyAssetsRow) GetMaxPhaseOk() (*int32, bool)`

GetMaxPhaseOk returns a tuple with the MaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhase

`func (o *CompanyAssetsRow) SetMaxPhase(v int32)`

SetMaxPhase sets MaxPhase field to given value.

### HasMaxPhase

`func (o *CompanyAssetsRow) HasMaxPhase() bool`

HasMaxPhase returns a boolean if a field has been set.

### SetMaxPhaseNil

`func (o *CompanyAssetsRow) SetMaxPhaseNil(b bool)`

 SetMaxPhaseNil sets the value for MaxPhase to be an explicit nil

### UnsetMaxPhase
`func (o *CompanyAssetsRow) UnsetMaxPhase()`

UnsetMaxPhase ensures that no value is present for MaxPhase, not even an explicit nil
### GetNTrials

`func (o *CompanyAssetsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *CompanyAssetsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *CompanyAssetsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *CompanyAssetsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *CompanyAssetsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *CompanyAssetsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNActive

`func (o *CompanyAssetsRow) GetNActive() int64`

GetNActive returns the NActive field if non-nil, zero value otherwise.

### GetNActiveOk

`func (o *CompanyAssetsRow) GetNActiveOk() (*int64, bool)`

GetNActiveOk returns a tuple with the NActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActive

`func (o *CompanyAssetsRow) SetNActive(v int64)`

SetNActive sets NActive field to given value.

### HasNActive

`func (o *CompanyAssetsRow) HasNActive() bool`

HasNActive returns a boolean if a field has been set.

### SetNActiveNil

`func (o *CompanyAssetsRow) SetNActiveNil(b bool)`

 SetNActiveNil sets the value for NActive to be an explicit nil

### UnsetNActive
`func (o *CompanyAssetsRow) UnsetNActive()`

UnsetNActive ensures that no value is present for NActive, not even an explicit nil
### GetLastYear

`func (o *CompanyAssetsRow) GetLastYear() int32`

GetLastYear returns the LastYear field if non-nil, zero value otherwise.

### GetLastYearOk

`func (o *CompanyAssetsRow) GetLastYearOk() (*int32, bool)`

GetLastYearOk returns a tuple with the LastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYear

`func (o *CompanyAssetsRow) SetLastYear(v int32)`

SetLastYear sets LastYear field to given value.

### HasLastYear

`func (o *CompanyAssetsRow) HasLastYear() bool`

HasLastYear returns a boolean if a field has been set.

### SetLastYearNil

`func (o *CompanyAssetsRow) SetLastYearNil(b bool)`

 SetLastYearNil sets the value for LastYear to be an explicit nil

### UnsetLastYear
`func (o *CompanyAssetsRow) UnsetLastYear()`

UnsetLastYear ensures that no value is present for LastYear, not even an explicit nil
### GetWithdrawn

`func (o *CompanyAssetsRow) GetWithdrawn() bool`

GetWithdrawn returns the Withdrawn field if non-nil, zero value otherwise.

### GetWithdrawnOk

`func (o *CompanyAssetsRow) GetWithdrawnOk() (*bool, bool)`

GetWithdrawnOk returns a tuple with the Withdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawn

`func (o *CompanyAssetsRow) SetWithdrawn(v bool)`

SetWithdrawn sets Withdrawn field to given value.

### HasWithdrawn

`func (o *CompanyAssetsRow) HasWithdrawn() bool`

HasWithdrawn returns a boolean if a field has been set.

### SetWithdrawnNil

`func (o *CompanyAssetsRow) SetWithdrawnNil(b bool)`

 SetWithdrawnNil sets the value for Withdrawn to be an explicit nil

### UnsetWithdrawn
`func (o *CompanyAssetsRow) UnsetWithdrawn()`

UnsetWithdrawn ensures that no value is present for Withdrawn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


