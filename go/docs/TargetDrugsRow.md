# TargetDrugsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] 
**DrugId** | Pointer to **NullableString** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] 
**MoleculeId** | Pointer to **NullableString** | The molecule the drug resolves to, the id condition_drugs also calls molecule_id; equal to drug_id for a drug that resolves to no other. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**PrimaryActionType** | Pointer to **NullableString** | The action type of the best-evidence mechanism (INHIBITOR, AGONIST and so on); null where the source gives none. | [optional] 
**EvidenceSource** | Pointer to **NullableString** | The best evidence tier that links the drug to the target: curated_moa (an Open Targets curated mechanism edge), chembl_moa (a ChEMBL mechanism) or intervention_seed (a curated investigational asset). | [optional] 
**EvidenceRank** | Pointer to **NullableInt32** | The rank of evidence_source, 1 for curated_moa, 2 for chembl_moa, 3 for intervention_seed. | [optional] 
**EvidenceSources** | Pointer to **[]string** | Every evidence tier that links the drug to the target; more than one is a corroboration signal. | [optional] 
**ApprovedAnywhere** | Pointer to **NullableBool** | Whether the drug is approved by any authority we hold. | [optional] 
**DrugNTrials** | Pointer to **NullableFloat32** | Distinct trials of the drug, on any target. | [optional] 

## Methods

### NewTargetDrugsRow

`func NewTargetDrugsRow() *TargetDrugsRow`

NewTargetDrugsRow instantiates a new TargetDrugsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetDrugsRowWithDefaults

`func NewTargetDrugsRowWithDefaults() *TargetDrugsRow`

NewTargetDrugsRowWithDefaults instantiates a new TargetDrugsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetDrugsRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetDrugsRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetDrugsRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetDrugsRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetDrugsRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetDrugsRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *TargetDrugsRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *TargetDrugsRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *TargetDrugsRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *TargetDrugsRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *TargetDrugsRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *TargetDrugsRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetDrugId

`func (o *TargetDrugsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *TargetDrugsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *TargetDrugsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *TargetDrugsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *TargetDrugsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *TargetDrugsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetMoleculeId

`func (o *TargetDrugsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *TargetDrugsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *TargetDrugsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *TargetDrugsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *TargetDrugsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *TargetDrugsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetDrugName

`func (o *TargetDrugsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *TargetDrugsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *TargetDrugsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *TargetDrugsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *TargetDrugsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *TargetDrugsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetPrimaryActionType

`func (o *TargetDrugsRow) GetPrimaryActionType() string`

GetPrimaryActionType returns the PrimaryActionType field if non-nil, zero value otherwise.

### GetPrimaryActionTypeOk

`func (o *TargetDrugsRow) GetPrimaryActionTypeOk() (*string, bool)`

GetPrimaryActionTypeOk returns a tuple with the PrimaryActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryActionType

`func (o *TargetDrugsRow) SetPrimaryActionType(v string)`

SetPrimaryActionType sets PrimaryActionType field to given value.

### HasPrimaryActionType

`func (o *TargetDrugsRow) HasPrimaryActionType() bool`

HasPrimaryActionType returns a boolean if a field has been set.

### SetPrimaryActionTypeNil

`func (o *TargetDrugsRow) SetPrimaryActionTypeNil(b bool)`

 SetPrimaryActionTypeNil sets the value for PrimaryActionType to be an explicit nil

### UnsetPrimaryActionType
`func (o *TargetDrugsRow) UnsetPrimaryActionType()`

UnsetPrimaryActionType ensures that no value is present for PrimaryActionType, not even an explicit nil
### GetEvidenceSource

`func (o *TargetDrugsRow) GetEvidenceSource() string`

GetEvidenceSource returns the EvidenceSource field if non-nil, zero value otherwise.

### GetEvidenceSourceOk

`func (o *TargetDrugsRow) GetEvidenceSourceOk() (*string, bool)`

GetEvidenceSourceOk returns a tuple with the EvidenceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSource

`func (o *TargetDrugsRow) SetEvidenceSource(v string)`

SetEvidenceSource sets EvidenceSource field to given value.

### HasEvidenceSource

`func (o *TargetDrugsRow) HasEvidenceSource() bool`

HasEvidenceSource returns a boolean if a field has been set.

### SetEvidenceSourceNil

`func (o *TargetDrugsRow) SetEvidenceSourceNil(b bool)`

 SetEvidenceSourceNil sets the value for EvidenceSource to be an explicit nil

### UnsetEvidenceSource
`func (o *TargetDrugsRow) UnsetEvidenceSource()`

UnsetEvidenceSource ensures that no value is present for EvidenceSource, not even an explicit nil
### GetEvidenceRank

`func (o *TargetDrugsRow) GetEvidenceRank() int32`

GetEvidenceRank returns the EvidenceRank field if non-nil, zero value otherwise.

### GetEvidenceRankOk

`func (o *TargetDrugsRow) GetEvidenceRankOk() (*int32, bool)`

GetEvidenceRankOk returns a tuple with the EvidenceRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRank

`func (o *TargetDrugsRow) SetEvidenceRank(v int32)`

SetEvidenceRank sets EvidenceRank field to given value.

### HasEvidenceRank

`func (o *TargetDrugsRow) HasEvidenceRank() bool`

HasEvidenceRank returns a boolean if a field has been set.

### SetEvidenceRankNil

`func (o *TargetDrugsRow) SetEvidenceRankNil(b bool)`

 SetEvidenceRankNil sets the value for EvidenceRank to be an explicit nil

### UnsetEvidenceRank
`func (o *TargetDrugsRow) UnsetEvidenceRank()`

UnsetEvidenceRank ensures that no value is present for EvidenceRank, not even an explicit nil
### GetEvidenceSources

`func (o *TargetDrugsRow) GetEvidenceSources() []string`

GetEvidenceSources returns the EvidenceSources field if non-nil, zero value otherwise.

### GetEvidenceSourcesOk

`func (o *TargetDrugsRow) GetEvidenceSourcesOk() (*[]string, bool)`

GetEvidenceSourcesOk returns a tuple with the EvidenceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSources

`func (o *TargetDrugsRow) SetEvidenceSources(v []string)`

SetEvidenceSources sets EvidenceSources field to given value.

### HasEvidenceSources

`func (o *TargetDrugsRow) HasEvidenceSources() bool`

HasEvidenceSources returns a boolean if a field has been set.

### SetEvidenceSourcesNil

`func (o *TargetDrugsRow) SetEvidenceSourcesNil(b bool)`

 SetEvidenceSourcesNil sets the value for EvidenceSources to be an explicit nil

### UnsetEvidenceSources
`func (o *TargetDrugsRow) UnsetEvidenceSources()`

UnsetEvidenceSources ensures that no value is present for EvidenceSources, not even an explicit nil
### GetApprovedAnywhere

`func (o *TargetDrugsRow) GetApprovedAnywhere() bool`

GetApprovedAnywhere returns the ApprovedAnywhere field if non-nil, zero value otherwise.

### GetApprovedAnywhereOk

`func (o *TargetDrugsRow) GetApprovedAnywhereOk() (*bool, bool)`

GetApprovedAnywhereOk returns a tuple with the ApprovedAnywhere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAnywhere

`func (o *TargetDrugsRow) SetApprovedAnywhere(v bool)`

SetApprovedAnywhere sets ApprovedAnywhere field to given value.

### HasApprovedAnywhere

`func (o *TargetDrugsRow) HasApprovedAnywhere() bool`

HasApprovedAnywhere returns a boolean if a field has been set.

### SetApprovedAnywhereNil

`func (o *TargetDrugsRow) SetApprovedAnywhereNil(b bool)`

 SetApprovedAnywhereNil sets the value for ApprovedAnywhere to be an explicit nil

### UnsetApprovedAnywhere
`func (o *TargetDrugsRow) UnsetApprovedAnywhere()`

UnsetApprovedAnywhere ensures that no value is present for ApprovedAnywhere, not even an explicit nil
### GetDrugNTrials

`func (o *TargetDrugsRow) GetDrugNTrials() float32`

GetDrugNTrials returns the DrugNTrials field if non-nil, zero value otherwise.

### GetDrugNTrialsOk

`func (o *TargetDrugsRow) GetDrugNTrialsOk() (*float32, bool)`

GetDrugNTrialsOk returns a tuple with the DrugNTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugNTrials

`func (o *TargetDrugsRow) SetDrugNTrials(v float32)`

SetDrugNTrials sets DrugNTrials field to given value.

### HasDrugNTrials

`func (o *TargetDrugsRow) HasDrugNTrials() bool`

HasDrugNTrials returns a boolean if a field has been set.

### SetDrugNTrialsNil

`func (o *TargetDrugsRow) SetDrugNTrialsNil(b bool)`

 SetDrugNTrialsNil sets the value for DrugNTrials to be an explicit nil

### UnsetDrugNTrials
`func (o *TargetDrugsRow) UnsetDrugNTrials()`

UnsetDrugNTrials ensures that no value is present for DrugNTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


