# DrugTargetsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**EnsemblId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id: the entry key of the targets object type. | [optional] 
**Symbol** | Pointer to **NullableString** | The gene symbol: the route key of the biomarker page for the same gene. | [optional] 
**TargetName** | Pointer to **NullableString** | The target&#39;s name. | [optional] 
**Uniprot** | Pointer to **NullableString** | The UniProt accession of the target protein. | [optional] 
**HasPdb** | Pointer to **NullableBool** | True when the Protein Data Bank holds a structure for the protein. | [optional] 
**PrimaryActionType** | Pointer to **NullableString** | The mechanism of action (for example INHIBITOR, AGONIST) from the best-ranked evidence that states one. | [optional] 
**EvidenceSource** | Pointer to **NullableString** | The best-ranked evidence source for the link. | [optional] 
**EvidenceRank** | Pointer to **NullableInt32** | The rank of that source; lower is stronger. | [optional] 
**EvidenceSources** | Pointer to **[]string** | Every evidence source for the link. | [optional] 

## Methods

### NewDrugTargetsRow

`func NewDrugTargetsRow() *DrugTargetsRow`

NewDrugTargetsRow instantiates a new DrugTargetsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugTargetsRowWithDefaults

`func NewDrugTargetsRowWithDefaults() *DrugTargetsRow`

NewDrugTargetsRowWithDefaults instantiates a new DrugTargetsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugTargetsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugTargetsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugTargetsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugTargetsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugTargetsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugTargetsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetEnsemblId

`func (o *DrugTargetsRow) GetEnsemblId() string`

GetEnsemblId returns the EnsemblId field if non-nil, zero value otherwise.

### GetEnsemblIdOk

`func (o *DrugTargetsRow) GetEnsemblIdOk() (*string, bool)`

GetEnsemblIdOk returns a tuple with the EnsemblId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblId

`func (o *DrugTargetsRow) SetEnsemblId(v string)`

SetEnsemblId sets EnsemblId field to given value.

### HasEnsemblId

`func (o *DrugTargetsRow) HasEnsemblId() bool`

HasEnsemblId returns a boolean if a field has been set.

### SetEnsemblIdNil

`func (o *DrugTargetsRow) SetEnsemblIdNil(b bool)`

 SetEnsemblIdNil sets the value for EnsemblId to be an explicit nil

### UnsetEnsemblId
`func (o *DrugTargetsRow) UnsetEnsemblId()`

UnsetEnsemblId ensures that no value is present for EnsemblId, not even an explicit nil
### GetSymbol

`func (o *DrugTargetsRow) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *DrugTargetsRow) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *DrugTargetsRow) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *DrugTargetsRow) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### SetSymbolNil

`func (o *DrugTargetsRow) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *DrugTargetsRow) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
### GetTargetName

`func (o *DrugTargetsRow) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DrugTargetsRow) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DrugTargetsRow) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DrugTargetsRow) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *DrugTargetsRow) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *DrugTargetsRow) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetUniprot

`func (o *DrugTargetsRow) GetUniprot() string`

GetUniprot returns the Uniprot field if non-nil, zero value otherwise.

### GetUniprotOk

`func (o *DrugTargetsRow) GetUniprotOk() (*string, bool)`

GetUniprotOk returns a tuple with the Uniprot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprot

`func (o *DrugTargetsRow) SetUniprot(v string)`

SetUniprot sets Uniprot field to given value.

### HasUniprot

`func (o *DrugTargetsRow) HasUniprot() bool`

HasUniprot returns a boolean if a field has been set.

### SetUniprotNil

`func (o *DrugTargetsRow) SetUniprotNil(b bool)`

 SetUniprotNil sets the value for Uniprot to be an explicit nil

### UnsetUniprot
`func (o *DrugTargetsRow) UnsetUniprot()`

UnsetUniprot ensures that no value is present for Uniprot, not even an explicit nil
### GetHasPdb

`func (o *DrugTargetsRow) GetHasPdb() bool`

GetHasPdb returns the HasPdb field if non-nil, zero value otherwise.

### GetHasPdbOk

`func (o *DrugTargetsRow) GetHasPdbOk() (*bool, bool)`

GetHasPdbOk returns a tuple with the HasPdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPdb

`func (o *DrugTargetsRow) SetHasPdb(v bool)`

SetHasPdb sets HasPdb field to given value.

### HasHasPdb

`func (o *DrugTargetsRow) HasHasPdb() bool`

HasHasPdb returns a boolean if a field has been set.

### SetHasPdbNil

`func (o *DrugTargetsRow) SetHasPdbNil(b bool)`

 SetHasPdbNil sets the value for HasPdb to be an explicit nil

### UnsetHasPdb
`func (o *DrugTargetsRow) UnsetHasPdb()`

UnsetHasPdb ensures that no value is present for HasPdb, not even an explicit nil
### GetPrimaryActionType

`func (o *DrugTargetsRow) GetPrimaryActionType() string`

GetPrimaryActionType returns the PrimaryActionType field if non-nil, zero value otherwise.

### GetPrimaryActionTypeOk

`func (o *DrugTargetsRow) GetPrimaryActionTypeOk() (*string, bool)`

GetPrimaryActionTypeOk returns a tuple with the PrimaryActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryActionType

`func (o *DrugTargetsRow) SetPrimaryActionType(v string)`

SetPrimaryActionType sets PrimaryActionType field to given value.

### HasPrimaryActionType

`func (o *DrugTargetsRow) HasPrimaryActionType() bool`

HasPrimaryActionType returns a boolean if a field has been set.

### SetPrimaryActionTypeNil

`func (o *DrugTargetsRow) SetPrimaryActionTypeNil(b bool)`

 SetPrimaryActionTypeNil sets the value for PrimaryActionType to be an explicit nil

### UnsetPrimaryActionType
`func (o *DrugTargetsRow) UnsetPrimaryActionType()`

UnsetPrimaryActionType ensures that no value is present for PrimaryActionType, not even an explicit nil
### GetEvidenceSource

`func (o *DrugTargetsRow) GetEvidenceSource() string`

GetEvidenceSource returns the EvidenceSource field if non-nil, zero value otherwise.

### GetEvidenceSourceOk

`func (o *DrugTargetsRow) GetEvidenceSourceOk() (*string, bool)`

GetEvidenceSourceOk returns a tuple with the EvidenceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSource

`func (o *DrugTargetsRow) SetEvidenceSource(v string)`

SetEvidenceSource sets EvidenceSource field to given value.

### HasEvidenceSource

`func (o *DrugTargetsRow) HasEvidenceSource() bool`

HasEvidenceSource returns a boolean if a field has been set.

### SetEvidenceSourceNil

`func (o *DrugTargetsRow) SetEvidenceSourceNil(b bool)`

 SetEvidenceSourceNil sets the value for EvidenceSource to be an explicit nil

### UnsetEvidenceSource
`func (o *DrugTargetsRow) UnsetEvidenceSource()`

UnsetEvidenceSource ensures that no value is present for EvidenceSource, not even an explicit nil
### GetEvidenceRank

`func (o *DrugTargetsRow) GetEvidenceRank() int32`

GetEvidenceRank returns the EvidenceRank field if non-nil, zero value otherwise.

### GetEvidenceRankOk

`func (o *DrugTargetsRow) GetEvidenceRankOk() (*int32, bool)`

GetEvidenceRankOk returns a tuple with the EvidenceRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRank

`func (o *DrugTargetsRow) SetEvidenceRank(v int32)`

SetEvidenceRank sets EvidenceRank field to given value.

### HasEvidenceRank

`func (o *DrugTargetsRow) HasEvidenceRank() bool`

HasEvidenceRank returns a boolean if a field has been set.

### SetEvidenceRankNil

`func (o *DrugTargetsRow) SetEvidenceRankNil(b bool)`

 SetEvidenceRankNil sets the value for EvidenceRank to be an explicit nil

### UnsetEvidenceRank
`func (o *DrugTargetsRow) UnsetEvidenceRank()`

UnsetEvidenceRank ensures that no value is present for EvidenceRank, not even an explicit nil
### GetEvidenceSources

`func (o *DrugTargetsRow) GetEvidenceSources() []string`

GetEvidenceSources returns the EvidenceSources field if non-nil, zero value otherwise.

### GetEvidenceSourcesOk

`func (o *DrugTargetsRow) GetEvidenceSourcesOk() (*[]string, bool)`

GetEvidenceSourcesOk returns a tuple with the EvidenceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSources

`func (o *DrugTargetsRow) SetEvidenceSources(v []string)`

SetEvidenceSources sets EvidenceSources field to given value.

### HasEvidenceSources

`func (o *DrugTargetsRow) HasEvidenceSources() bool`

HasEvidenceSources returns a boolean if a field has been set.

### SetEvidenceSourcesNil

`func (o *DrugTargetsRow) SetEvidenceSourcesNil(b bool)`

 SetEvidenceSourcesNil sets the value for EvidenceSources to be an explicit nil

### UnsetEvidenceSources
`func (o *DrugTargetsRow) UnsetEvidenceSources()`

UnsetEvidenceSources ensures that no value is present for EvidenceSources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


