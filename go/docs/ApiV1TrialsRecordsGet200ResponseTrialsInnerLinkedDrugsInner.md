# ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | **NullableString** |  | 
**MoleculeName** | **NullableString** |  | 
**RegisteredAs** | **[]string** |  | 
**MatchedFromNamePart** | **bool** |  | 
**BrandCount** | **int32** |  | 
**Brands** | [**[]ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerBrandsInner**](ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerBrandsInner.md) |  | 

## Methods

### NewApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner

`func NewApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner(moleculeId NullableString, moleculeName NullableString, registeredAs []string, matchedFromNamePart bool, brandCount int32, brands []ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerBrandsInner, ) *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner`

NewApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner instantiates a new ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerWithDefaults

`func NewApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerWithDefaults() *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner`

NewApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerWithDefaults instantiates a new ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.


### SetMoleculeIdNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetMoleculeName

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetMoleculeName() string`

GetMoleculeName returns the MoleculeName field if non-nil, zero value otherwise.

### GetMoleculeNameOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetMoleculeNameOk() (*string, bool)`

GetMoleculeNameOk returns a tuple with the MoleculeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeName

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) SetMoleculeName(v string)`

SetMoleculeName sets MoleculeName field to given value.


### SetMoleculeNameNil

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) SetMoleculeNameNil(b bool)`

 SetMoleculeNameNil sets the value for MoleculeName to be an explicit nil

### UnsetMoleculeName
`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) UnsetMoleculeName()`

UnsetMoleculeName ensures that no value is present for MoleculeName, not even an explicit nil
### GetRegisteredAs

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetRegisteredAs() []string`

GetRegisteredAs returns the RegisteredAs field if non-nil, zero value otherwise.

### GetRegisteredAsOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetRegisteredAsOk() (*[]string, bool)`

GetRegisteredAsOk returns a tuple with the RegisteredAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAs

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) SetRegisteredAs(v []string)`

SetRegisteredAs sets RegisteredAs field to given value.


### GetMatchedFromNamePart

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetMatchedFromNamePart() bool`

GetMatchedFromNamePart returns the MatchedFromNamePart field if non-nil, zero value otherwise.

### GetMatchedFromNamePartOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetMatchedFromNamePartOk() (*bool, bool)`

GetMatchedFromNamePartOk returns a tuple with the MatchedFromNamePart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedFromNamePart

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) SetMatchedFromNamePart(v bool)`

SetMatchedFromNamePart sets MatchedFromNamePart field to given value.


### GetBrandCount

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetBrandCount() int32`

GetBrandCount returns the BrandCount field if non-nil, zero value otherwise.

### GetBrandCountOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetBrandCountOk() (*int32, bool)`

GetBrandCountOk returns a tuple with the BrandCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandCount

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) SetBrandCount(v int32)`

SetBrandCount sets BrandCount field to given value.


### GetBrands

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetBrands() []ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerBrandsInner`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) GetBrandsOk() (*[]ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerBrandsInner, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInner) SetBrands(v []ApiV1TrialsRecordsGet200ResponseTrialsInnerLinkedDrugsInnerBrandsInner)`

SetBrands sets Brands field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


