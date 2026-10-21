# TargetPublicationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id. | [optional] 
**Citation** | Pointer to **NullableString** | The citation string registrants gave. | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title from PubMed; null when PubMed has no record. | [optional] 
**PubYear** | Pointer to **NullableInt32** | The publication year. | [optional] 
**NLinkedTrials** | Pointer to **NullableInt32** | Distinct trials in the registry that cite the paper, for any target. | [optional] 
**NTargetTrials** | Pointer to **NullableInt32** | Distinct trials of this target that cite the paper. | [optional] 

## Methods

### NewTargetPublicationsRow

`func NewTargetPublicationsRow() *TargetPublicationsRow`

NewTargetPublicationsRow instantiates a new TargetPublicationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetPublicationsRowWithDefaults

`func NewTargetPublicationsRowWithDefaults() *TargetPublicationsRow`

NewTargetPublicationsRowWithDefaults instantiates a new TargetPublicationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetPublicationsRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetPublicationsRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetPublicationsRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetPublicationsRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetPublicationsRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetPublicationsRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *TargetPublicationsRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *TargetPublicationsRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *TargetPublicationsRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *TargetPublicationsRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *TargetPublicationsRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *TargetPublicationsRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetPmid

`func (o *TargetPublicationsRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *TargetPublicationsRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *TargetPublicationsRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *TargetPublicationsRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *TargetPublicationsRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *TargetPublicationsRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetCitation

`func (o *TargetPublicationsRow) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *TargetPublicationsRow) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *TargetPublicationsRow) SetCitation(v string)`

SetCitation sets Citation field to given value.

### HasCitation

`func (o *TargetPublicationsRow) HasCitation() bool`

HasCitation returns a boolean if a field has been set.

### SetCitationNil

`func (o *TargetPublicationsRow) SetCitationNil(b bool)`

 SetCitationNil sets the value for Citation to be an explicit nil

### UnsetCitation
`func (o *TargetPublicationsRow) UnsetCitation()`

UnsetCitation ensures that no value is present for Citation, not even an explicit nil
### GetTitle

`func (o *TargetPublicationsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TargetPublicationsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TargetPublicationsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TargetPublicationsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TargetPublicationsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TargetPublicationsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPubYear

`func (o *TargetPublicationsRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *TargetPublicationsRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *TargetPublicationsRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *TargetPublicationsRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *TargetPublicationsRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *TargetPublicationsRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetNLinkedTrials

`func (o *TargetPublicationsRow) GetNLinkedTrials() int32`

GetNLinkedTrials returns the NLinkedTrials field if non-nil, zero value otherwise.

### GetNLinkedTrialsOk

`func (o *TargetPublicationsRow) GetNLinkedTrialsOk() (*int32, bool)`

GetNLinkedTrialsOk returns a tuple with the NLinkedTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLinkedTrials

`func (o *TargetPublicationsRow) SetNLinkedTrials(v int32)`

SetNLinkedTrials sets NLinkedTrials field to given value.

### HasNLinkedTrials

`func (o *TargetPublicationsRow) HasNLinkedTrials() bool`

HasNLinkedTrials returns a boolean if a field has been set.

### SetNLinkedTrialsNil

`func (o *TargetPublicationsRow) SetNLinkedTrialsNil(b bool)`

 SetNLinkedTrialsNil sets the value for NLinkedTrials to be an explicit nil

### UnsetNLinkedTrials
`func (o *TargetPublicationsRow) UnsetNLinkedTrials()`

UnsetNLinkedTrials ensures that no value is present for NLinkedTrials, not even an explicit nil
### GetNTargetTrials

`func (o *TargetPublicationsRow) GetNTargetTrials() int32`

GetNTargetTrials returns the NTargetTrials field if non-nil, zero value otherwise.

### GetNTargetTrialsOk

`func (o *TargetPublicationsRow) GetNTargetTrialsOk() (*int32, bool)`

GetNTargetTrialsOk returns a tuple with the NTargetTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTargetTrials

`func (o *TargetPublicationsRow) SetNTargetTrials(v int32)`

SetNTargetTrials sets NTargetTrials field to given value.

### HasNTargetTrials

`func (o *TargetPublicationsRow) HasNTargetTrials() bool`

HasNTargetTrials returns a boolean if a field has been set.

### SetNTargetTrialsNil

`func (o *TargetPublicationsRow) SetNTargetTrialsNil(b bool)`

 SetNTargetTrialsNil sets the value for NTargetTrials to be an explicit nil

### UnsetNTargetTrials
`func (o *TargetPublicationsRow) UnsetNTargetTrials()`

UnsetNTargetTrials ensures that no value is present for NTargetTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


