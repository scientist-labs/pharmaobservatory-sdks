# BiomarkerEvidenceByYearRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The gene&#39;s route key, the entry key for the biomarkers object type; the gene symbol. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**EvidenceType** | Pointer to **NullableString** | The evidence arm. trial_mention counts trials that mention the gene (the year is the trial start year); publication_mention counts papers cited by registered trials that mention it (the year is the publication year). | [optional] 
**EvidenceYear** | Pointer to **NullableInt32** | The year of the mentions. Null for the undated class. | [optional] 
**YearClass** | Pointer to **NullableString** | How the year is treated. plotted: inside the window the web chart draws. pre_window: before 1990. future_dated: after the current year. corpus_incomplete: after the latest year the arm&#39;s corpus covers. undated: the source has no year. | [optional] 
**NMentions** | Pointer to **NullableInt32** | Mentions of the gene in that arm, year and class. A document that mentions the gene more than once counts once per mention. | [optional] 
**NDocuments** | Pointer to **NullableInt32** | Distinct documents (trials or papers) that mention the gene in that arm, year and class. | [optional] 
**YearDenominator** | Pointer to **NullableInt32** | Distinct documents in that arm and year that mention any gene, the denominator of the gene&#39;s share of the year. | [optional] 

## Methods

### NewBiomarkerEvidenceByYearRow

`func NewBiomarkerEvidenceByYearRow() *BiomarkerEvidenceByYearRow`

NewBiomarkerEvidenceByYearRow instantiates a new BiomarkerEvidenceByYearRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerEvidenceByYearRowWithDefaults

`func NewBiomarkerEvidenceByYearRowWithDefaults() *BiomarkerEvidenceByYearRow`

NewBiomarkerEvidenceByYearRowWithDefaults instantiates a new BiomarkerEvidenceByYearRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerEvidenceByYearRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerEvidenceByYearRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerEvidenceByYearRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerEvidenceByYearRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerEvidenceByYearRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerEvidenceByYearRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerEvidenceByYearRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerEvidenceByYearRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerEvidenceByYearRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerEvidenceByYearRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerEvidenceByYearRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerEvidenceByYearRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetEvidenceType

`func (o *BiomarkerEvidenceByYearRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *BiomarkerEvidenceByYearRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *BiomarkerEvidenceByYearRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *BiomarkerEvidenceByYearRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *BiomarkerEvidenceByYearRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *BiomarkerEvidenceByYearRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetEvidenceYear

`func (o *BiomarkerEvidenceByYearRow) GetEvidenceYear() int32`

GetEvidenceYear returns the EvidenceYear field if non-nil, zero value otherwise.

### GetEvidenceYearOk

`func (o *BiomarkerEvidenceByYearRow) GetEvidenceYearOk() (*int32, bool)`

GetEvidenceYearOk returns a tuple with the EvidenceYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceYear

`func (o *BiomarkerEvidenceByYearRow) SetEvidenceYear(v int32)`

SetEvidenceYear sets EvidenceYear field to given value.

### HasEvidenceYear

`func (o *BiomarkerEvidenceByYearRow) HasEvidenceYear() bool`

HasEvidenceYear returns a boolean if a field has been set.

### SetEvidenceYearNil

`func (o *BiomarkerEvidenceByYearRow) SetEvidenceYearNil(b bool)`

 SetEvidenceYearNil sets the value for EvidenceYear to be an explicit nil

### UnsetEvidenceYear
`func (o *BiomarkerEvidenceByYearRow) UnsetEvidenceYear()`

UnsetEvidenceYear ensures that no value is present for EvidenceYear, not even an explicit nil
### GetYearClass

`func (o *BiomarkerEvidenceByYearRow) GetYearClass() string`

GetYearClass returns the YearClass field if non-nil, zero value otherwise.

### GetYearClassOk

`func (o *BiomarkerEvidenceByYearRow) GetYearClassOk() (*string, bool)`

GetYearClassOk returns a tuple with the YearClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearClass

`func (o *BiomarkerEvidenceByYearRow) SetYearClass(v string)`

SetYearClass sets YearClass field to given value.

### HasYearClass

`func (o *BiomarkerEvidenceByYearRow) HasYearClass() bool`

HasYearClass returns a boolean if a field has been set.

### SetYearClassNil

`func (o *BiomarkerEvidenceByYearRow) SetYearClassNil(b bool)`

 SetYearClassNil sets the value for YearClass to be an explicit nil

### UnsetYearClass
`func (o *BiomarkerEvidenceByYearRow) UnsetYearClass()`

UnsetYearClass ensures that no value is present for YearClass, not even an explicit nil
### GetNMentions

`func (o *BiomarkerEvidenceByYearRow) GetNMentions() int32`

GetNMentions returns the NMentions field if non-nil, zero value otherwise.

### GetNMentionsOk

`func (o *BiomarkerEvidenceByYearRow) GetNMentionsOk() (*int32, bool)`

GetNMentionsOk returns a tuple with the NMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNMentions

`func (o *BiomarkerEvidenceByYearRow) SetNMentions(v int32)`

SetNMentions sets NMentions field to given value.

### HasNMentions

`func (o *BiomarkerEvidenceByYearRow) HasNMentions() bool`

HasNMentions returns a boolean if a field has been set.

### SetNMentionsNil

`func (o *BiomarkerEvidenceByYearRow) SetNMentionsNil(b bool)`

 SetNMentionsNil sets the value for NMentions to be an explicit nil

### UnsetNMentions
`func (o *BiomarkerEvidenceByYearRow) UnsetNMentions()`

UnsetNMentions ensures that no value is present for NMentions, not even an explicit nil
### GetNDocuments

`func (o *BiomarkerEvidenceByYearRow) GetNDocuments() int32`

GetNDocuments returns the NDocuments field if non-nil, zero value otherwise.

### GetNDocumentsOk

`func (o *BiomarkerEvidenceByYearRow) GetNDocumentsOk() (*int32, bool)`

GetNDocumentsOk returns a tuple with the NDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDocuments

`func (o *BiomarkerEvidenceByYearRow) SetNDocuments(v int32)`

SetNDocuments sets NDocuments field to given value.

### HasNDocuments

`func (o *BiomarkerEvidenceByYearRow) HasNDocuments() bool`

HasNDocuments returns a boolean if a field has been set.

### SetNDocumentsNil

`func (o *BiomarkerEvidenceByYearRow) SetNDocumentsNil(b bool)`

 SetNDocumentsNil sets the value for NDocuments to be an explicit nil

### UnsetNDocuments
`func (o *BiomarkerEvidenceByYearRow) UnsetNDocuments()`

UnsetNDocuments ensures that no value is present for NDocuments, not even an explicit nil
### GetYearDenominator

`func (o *BiomarkerEvidenceByYearRow) GetYearDenominator() int32`

GetYearDenominator returns the YearDenominator field if non-nil, zero value otherwise.

### GetYearDenominatorOk

`func (o *BiomarkerEvidenceByYearRow) GetYearDenominatorOk() (*int32, bool)`

GetYearDenominatorOk returns a tuple with the YearDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearDenominator

`func (o *BiomarkerEvidenceByYearRow) SetYearDenominator(v int32)`

SetYearDenominator sets YearDenominator field to given value.

### HasYearDenominator

`func (o *BiomarkerEvidenceByYearRow) HasYearDenominator() bool`

HasYearDenominator returns a boolean if a field has been set.

### SetYearDenominatorNil

`func (o *BiomarkerEvidenceByYearRow) SetYearDenominatorNil(b bool)`

 SetYearDenominatorNil sets the value for YearDenominator to be an explicit nil

### UnsetYearDenominator
`func (o *BiomarkerEvidenceByYearRow) UnsetYearDenominator()`

UnsetYearDenominator ensures that no value is present for YearDenominator, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


