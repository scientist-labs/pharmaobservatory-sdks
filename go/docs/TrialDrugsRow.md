# TrialDrugsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**DrugKey** | Pointer to **NullableString** | The molecule key, or the drug name as registered prefixed unresolved: for a drug that reached no molecule. | [optional] 
**MoleculeId** | Pointer to **NullableString** | The molecule key the drugs resolved to (null for a drug with no molecule). | [optional] 
**MoleculeName** | Pointer to **NullableString** | The molecule&#39;s display name (null for a drug with no molecule). | [optional] 
**RegisteredAs** | Pointer to **[]string** | The drug names as registered on the trial that led to this row. | [optional] 
**MatchedFromNamePart** | Pointer to **NullableBool** | Whether every match to the molecule came from one part of a multi-part drug name. | [optional] 
**BrandCount** | Pointer to **NullableInt32** | Number of brand names recorded for the molecule. | [optional] 

## Methods

### NewTrialDrugsRow

`func NewTrialDrugsRow() *TrialDrugsRow`

NewTrialDrugsRow instantiates a new TrialDrugsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialDrugsRowWithDefaults

`func NewTrialDrugsRowWithDefaults() *TrialDrugsRow`

NewTrialDrugsRowWithDefaults instantiates a new TrialDrugsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialDrugsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialDrugsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialDrugsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialDrugsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialDrugsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialDrugsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetDrugKey

`func (o *TrialDrugsRow) GetDrugKey() string`

GetDrugKey returns the DrugKey field if non-nil, zero value otherwise.

### GetDrugKeyOk

`func (o *TrialDrugsRow) GetDrugKeyOk() (*string, bool)`

GetDrugKeyOk returns a tuple with the DrugKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugKey

`func (o *TrialDrugsRow) SetDrugKey(v string)`

SetDrugKey sets DrugKey field to given value.

### HasDrugKey

`func (o *TrialDrugsRow) HasDrugKey() bool`

HasDrugKey returns a boolean if a field has been set.

### SetDrugKeyNil

`func (o *TrialDrugsRow) SetDrugKeyNil(b bool)`

 SetDrugKeyNil sets the value for DrugKey to be an explicit nil

### UnsetDrugKey
`func (o *TrialDrugsRow) UnsetDrugKey()`

UnsetDrugKey ensures that no value is present for DrugKey, not even an explicit nil
### GetMoleculeId

`func (o *TrialDrugsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *TrialDrugsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *TrialDrugsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *TrialDrugsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *TrialDrugsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *TrialDrugsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetMoleculeName

`func (o *TrialDrugsRow) GetMoleculeName() string`

GetMoleculeName returns the MoleculeName field if non-nil, zero value otherwise.

### GetMoleculeNameOk

`func (o *TrialDrugsRow) GetMoleculeNameOk() (*string, bool)`

GetMoleculeNameOk returns a tuple with the MoleculeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeName

`func (o *TrialDrugsRow) SetMoleculeName(v string)`

SetMoleculeName sets MoleculeName field to given value.

### HasMoleculeName

`func (o *TrialDrugsRow) HasMoleculeName() bool`

HasMoleculeName returns a boolean if a field has been set.

### SetMoleculeNameNil

`func (o *TrialDrugsRow) SetMoleculeNameNil(b bool)`

 SetMoleculeNameNil sets the value for MoleculeName to be an explicit nil

### UnsetMoleculeName
`func (o *TrialDrugsRow) UnsetMoleculeName()`

UnsetMoleculeName ensures that no value is present for MoleculeName, not even an explicit nil
### GetRegisteredAs

`func (o *TrialDrugsRow) GetRegisteredAs() []string`

GetRegisteredAs returns the RegisteredAs field if non-nil, zero value otherwise.

### GetRegisteredAsOk

`func (o *TrialDrugsRow) GetRegisteredAsOk() (*[]string, bool)`

GetRegisteredAsOk returns a tuple with the RegisteredAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAs

`func (o *TrialDrugsRow) SetRegisteredAs(v []string)`

SetRegisteredAs sets RegisteredAs field to given value.

### HasRegisteredAs

`func (o *TrialDrugsRow) HasRegisteredAs() bool`

HasRegisteredAs returns a boolean if a field has been set.

### SetRegisteredAsNil

`func (o *TrialDrugsRow) SetRegisteredAsNil(b bool)`

 SetRegisteredAsNil sets the value for RegisteredAs to be an explicit nil

### UnsetRegisteredAs
`func (o *TrialDrugsRow) UnsetRegisteredAs()`

UnsetRegisteredAs ensures that no value is present for RegisteredAs, not even an explicit nil
### GetMatchedFromNamePart

`func (o *TrialDrugsRow) GetMatchedFromNamePart() bool`

GetMatchedFromNamePart returns the MatchedFromNamePart field if non-nil, zero value otherwise.

### GetMatchedFromNamePartOk

`func (o *TrialDrugsRow) GetMatchedFromNamePartOk() (*bool, bool)`

GetMatchedFromNamePartOk returns a tuple with the MatchedFromNamePart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedFromNamePart

`func (o *TrialDrugsRow) SetMatchedFromNamePart(v bool)`

SetMatchedFromNamePart sets MatchedFromNamePart field to given value.

### HasMatchedFromNamePart

`func (o *TrialDrugsRow) HasMatchedFromNamePart() bool`

HasMatchedFromNamePart returns a boolean if a field has been set.

### SetMatchedFromNamePartNil

`func (o *TrialDrugsRow) SetMatchedFromNamePartNil(b bool)`

 SetMatchedFromNamePartNil sets the value for MatchedFromNamePart to be an explicit nil

### UnsetMatchedFromNamePart
`func (o *TrialDrugsRow) UnsetMatchedFromNamePart()`

UnsetMatchedFromNamePart ensures that no value is present for MatchedFromNamePart, not even an explicit nil
### GetBrandCount

`func (o *TrialDrugsRow) GetBrandCount() int32`

GetBrandCount returns the BrandCount field if non-nil, zero value otherwise.

### GetBrandCountOk

`func (o *TrialDrugsRow) GetBrandCountOk() (*int32, bool)`

GetBrandCountOk returns a tuple with the BrandCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandCount

`func (o *TrialDrugsRow) SetBrandCount(v int32)`

SetBrandCount sets BrandCount field to given value.

### HasBrandCount

`func (o *TrialDrugsRow) HasBrandCount() bool`

HasBrandCount returns a boolean if a field has been set.

### SetBrandCountNil

`func (o *TrialDrugsRow) SetBrandCountNil(b bool)`

 SetBrandCountNil sets the value for BrandCount to be an explicit nil

### UnsetBrandCount
`func (o *TrialDrugsRow) UnsetBrandCount()`

UnsetBrandCount ensures that no value is present for BrandCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


