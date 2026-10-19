# CompanyTargetsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company whose lead-sponsored trials pursue the target. | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**TargetId** | Pointer to **NullableString** | The target&#39;s gene identifier (Ensembl). | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s gene symbol. | [optional] 
**TargetName** | Pointer to **NullableString** | The target&#39;s name. | [optional] 
**MaxPhase** | Pointer to **NullableInt32** | The furthest phase any qualifying program of the company reached against the target. | [optional] 
**NDrugs** | Pointer to **NullableInt32** | Distinct drugs of the company&#39;s qualifying programs against the target. | [optional] 
**NConditions** | Pointer to **NullableInt32** | Distinct conditions of the company&#39;s qualifying programs against the target. | [optional] 
**NTrials** | Pointer to **NullableInt32** | Distinct trials the company leads in those programs. | [optional] 

## Methods

### NewCompanyTargetsRow

`func NewCompanyTargetsRow() *CompanyTargetsRow`

NewCompanyTargetsRow instantiates a new CompanyTargetsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyTargetsRowWithDefaults

`func NewCompanyTargetsRowWithDefaults() *CompanyTargetsRow`

NewCompanyTargetsRowWithDefaults instantiates a new CompanyTargetsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponsorCompanyId

`func (o *CompanyTargetsRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *CompanyTargetsRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *CompanyTargetsRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *CompanyTargetsRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *CompanyTargetsRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *CompanyTargetsRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *CompanyTargetsRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *CompanyTargetsRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *CompanyTargetsRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *CompanyTargetsRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *CompanyTargetsRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *CompanyTargetsRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *CompanyTargetsRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *CompanyTargetsRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *CompanyTargetsRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *CompanyTargetsRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *CompanyTargetsRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *CompanyTargetsRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *CompanyTargetsRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *CompanyTargetsRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *CompanyTargetsRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *CompanyTargetsRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *CompanyTargetsRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *CompanyTargetsRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetTargetId

`func (o *CompanyTargetsRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *CompanyTargetsRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *CompanyTargetsRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *CompanyTargetsRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *CompanyTargetsRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *CompanyTargetsRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *CompanyTargetsRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *CompanyTargetsRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *CompanyTargetsRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *CompanyTargetsRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *CompanyTargetsRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *CompanyTargetsRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetTargetName

`func (o *CompanyTargetsRow) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *CompanyTargetsRow) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *CompanyTargetsRow) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *CompanyTargetsRow) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *CompanyTargetsRow) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *CompanyTargetsRow) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetMaxPhase

`func (o *CompanyTargetsRow) GetMaxPhase() int32`

GetMaxPhase returns the MaxPhase field if non-nil, zero value otherwise.

### GetMaxPhaseOk

`func (o *CompanyTargetsRow) GetMaxPhaseOk() (*int32, bool)`

GetMaxPhaseOk returns a tuple with the MaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhase

`func (o *CompanyTargetsRow) SetMaxPhase(v int32)`

SetMaxPhase sets MaxPhase field to given value.

### HasMaxPhase

`func (o *CompanyTargetsRow) HasMaxPhase() bool`

HasMaxPhase returns a boolean if a field has been set.

### SetMaxPhaseNil

`func (o *CompanyTargetsRow) SetMaxPhaseNil(b bool)`

 SetMaxPhaseNil sets the value for MaxPhase to be an explicit nil

### UnsetMaxPhase
`func (o *CompanyTargetsRow) UnsetMaxPhase()`

UnsetMaxPhase ensures that no value is present for MaxPhase, not even an explicit nil
### GetNDrugs

`func (o *CompanyTargetsRow) GetNDrugs() int32`

GetNDrugs returns the NDrugs field if non-nil, zero value otherwise.

### GetNDrugsOk

`func (o *CompanyTargetsRow) GetNDrugsOk() (*int32, bool)`

GetNDrugsOk returns a tuple with the NDrugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDrugs

`func (o *CompanyTargetsRow) SetNDrugs(v int32)`

SetNDrugs sets NDrugs field to given value.

### HasNDrugs

`func (o *CompanyTargetsRow) HasNDrugs() bool`

HasNDrugs returns a boolean if a field has been set.

### SetNDrugsNil

`func (o *CompanyTargetsRow) SetNDrugsNil(b bool)`

 SetNDrugsNil sets the value for NDrugs to be an explicit nil

### UnsetNDrugs
`func (o *CompanyTargetsRow) UnsetNDrugs()`

UnsetNDrugs ensures that no value is present for NDrugs, not even an explicit nil
### GetNConditions

`func (o *CompanyTargetsRow) GetNConditions() int32`

GetNConditions returns the NConditions field if non-nil, zero value otherwise.

### GetNConditionsOk

`func (o *CompanyTargetsRow) GetNConditionsOk() (*int32, bool)`

GetNConditionsOk returns a tuple with the NConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNConditions

`func (o *CompanyTargetsRow) SetNConditions(v int32)`

SetNConditions sets NConditions field to given value.

### HasNConditions

`func (o *CompanyTargetsRow) HasNConditions() bool`

HasNConditions returns a boolean if a field has been set.

### SetNConditionsNil

`func (o *CompanyTargetsRow) SetNConditionsNil(b bool)`

 SetNConditionsNil sets the value for NConditions to be an explicit nil

### UnsetNConditions
`func (o *CompanyTargetsRow) UnsetNConditions()`

UnsetNConditions ensures that no value is present for NConditions, not even an explicit nil
### GetNTrials

`func (o *CompanyTargetsRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *CompanyTargetsRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *CompanyTargetsRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *CompanyTargetsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *CompanyTargetsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *CompanyTargetsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


