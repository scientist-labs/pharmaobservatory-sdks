# BiomarkersList200ResponseQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Q** | **NullableString** |  | 
**Filters** | [**BiomarkersList200ResponseQueryFilters**](BiomarkersList200ResponseQueryFilters.md) |  | 
**Sort** | **string** |  | 
**Dir** | **string** |  | 
**EvidenceFloor** | **bool** |  | 
**ValidatedFloor** | **bool** |  | 

## Methods

### NewBiomarkersList200ResponseQuery

`func NewBiomarkersList200ResponseQuery(q NullableString, filters BiomarkersList200ResponseQueryFilters, sort string, dir string, evidenceFloor bool, validatedFloor bool, ) *BiomarkersList200ResponseQuery`

NewBiomarkersList200ResponseQuery instantiates a new BiomarkersList200ResponseQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkersList200ResponseQueryWithDefaults

`func NewBiomarkersList200ResponseQueryWithDefaults() *BiomarkersList200ResponseQuery`

NewBiomarkersList200ResponseQueryWithDefaults instantiates a new BiomarkersList200ResponseQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQ

`func (o *BiomarkersList200ResponseQuery) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *BiomarkersList200ResponseQuery) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *BiomarkersList200ResponseQuery) SetQ(v string)`

SetQ sets Q field to given value.


### SetQNil

`func (o *BiomarkersList200ResponseQuery) SetQNil(b bool)`

 SetQNil sets the value for Q to be an explicit nil

### UnsetQ
`func (o *BiomarkersList200ResponseQuery) UnsetQ()`

UnsetQ ensures that no value is present for Q, not even an explicit nil
### GetFilters

`func (o *BiomarkersList200ResponseQuery) GetFilters() BiomarkersList200ResponseQueryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BiomarkersList200ResponseQuery) GetFiltersOk() (*BiomarkersList200ResponseQueryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BiomarkersList200ResponseQuery) SetFilters(v BiomarkersList200ResponseQueryFilters)`

SetFilters sets Filters field to given value.


### GetSort

`func (o *BiomarkersList200ResponseQuery) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *BiomarkersList200ResponseQuery) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *BiomarkersList200ResponseQuery) SetSort(v string)`

SetSort sets Sort field to given value.


### GetDir

`func (o *BiomarkersList200ResponseQuery) GetDir() string`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *BiomarkersList200ResponseQuery) GetDirOk() (*string, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *BiomarkersList200ResponseQuery) SetDir(v string)`

SetDir sets Dir field to given value.


### GetEvidenceFloor

`func (o *BiomarkersList200ResponseQuery) GetEvidenceFloor() bool`

GetEvidenceFloor returns the EvidenceFloor field if non-nil, zero value otherwise.

### GetEvidenceFloorOk

`func (o *BiomarkersList200ResponseQuery) GetEvidenceFloorOk() (*bool, bool)`

GetEvidenceFloorOk returns a tuple with the EvidenceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceFloor

`func (o *BiomarkersList200ResponseQuery) SetEvidenceFloor(v bool)`

SetEvidenceFloor sets EvidenceFloor field to given value.


### GetValidatedFloor

`func (o *BiomarkersList200ResponseQuery) GetValidatedFloor() bool`

GetValidatedFloor returns the ValidatedFloor field if non-nil, zero value otherwise.

### GetValidatedFloorOk

`func (o *BiomarkersList200ResponseQuery) GetValidatedFloorOk() (*bool, bool)`

GetValidatedFloorOk returns a tuple with the ValidatedFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedFloor

`func (o *BiomarkersList200ResponseQuery) SetValidatedFloor(v bool)`

SetValidatedFloor sets ValidatedFloor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


