# BiomarkerEvidenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiomarkerEvidenceId** | Pointer to **NullableString** | The evidence assertion id. | [optional] 
**RouteKey** | Pointer to **NullableString** | The biomarker route key of the gene this evidence is about (the gene symbol); the entry key for the biomarkers object type. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol the evidence is about. | [optional] 
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id. | [optional] 
**UniprotAccession** | Pointer to **NullableString** | The UniProt accession. | [optional] 
**Biomarker** | Pointer to **NullableString** | The biomarker as the source names it (the gene or the variant). | [optional] 
**EvidenceType** | Pointer to **NullableString** | cancer_actionable_cgi, ot_association or trial_mention. | [optional] 
**Source** | Pointer to **NullableString** | The source the assertion comes from. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key. | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition as the source names it. | [optional] 
**TherapeuticArea** | Pointer to **NullableString** | The therapeutic area. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug as filed. | [optional] 
**DrugId** | Pointer to **NullableString** | The drug key. | [optional] 
**MoleculeId** | Pointer to **NullableString** | The molecule key. Null today: the evidence mart carries no molecule column. | [optional] 
**Response** | Pointer to **NullableString** | The reported response (for example Responsive, Resistant). | [optional] 
**Score** | Pointer to **NullableFloat64** | The association score, for an Open Targets row. | [optional] 
**VariantCount** | Pointer to **NullableInt32** | Variants the assertion covers. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial registry id, for a trial-mention row. | [optional] 
**SampleType** | Pointer to **NullableString** | The sample type, for a trial-mention row. | [optional] 

## Methods

### NewBiomarkerEvidenceRow

`func NewBiomarkerEvidenceRow() *BiomarkerEvidenceRow`

NewBiomarkerEvidenceRow instantiates a new BiomarkerEvidenceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerEvidenceRowWithDefaults

`func NewBiomarkerEvidenceRowWithDefaults() *BiomarkerEvidenceRow`

NewBiomarkerEvidenceRowWithDefaults instantiates a new BiomarkerEvidenceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiomarkerEvidenceId

`func (o *BiomarkerEvidenceRow) GetBiomarkerEvidenceId() string`

GetBiomarkerEvidenceId returns the BiomarkerEvidenceId field if non-nil, zero value otherwise.

### GetBiomarkerEvidenceIdOk

`func (o *BiomarkerEvidenceRow) GetBiomarkerEvidenceIdOk() (*string, bool)`

GetBiomarkerEvidenceIdOk returns a tuple with the BiomarkerEvidenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerEvidenceId

`func (o *BiomarkerEvidenceRow) SetBiomarkerEvidenceId(v string)`

SetBiomarkerEvidenceId sets BiomarkerEvidenceId field to given value.

### HasBiomarkerEvidenceId

`func (o *BiomarkerEvidenceRow) HasBiomarkerEvidenceId() bool`

HasBiomarkerEvidenceId returns a boolean if a field has been set.

### SetBiomarkerEvidenceIdNil

`func (o *BiomarkerEvidenceRow) SetBiomarkerEvidenceIdNil(b bool)`

 SetBiomarkerEvidenceIdNil sets the value for BiomarkerEvidenceId to be an explicit nil

### UnsetBiomarkerEvidenceId
`func (o *BiomarkerEvidenceRow) UnsetBiomarkerEvidenceId()`

UnsetBiomarkerEvidenceId ensures that no value is present for BiomarkerEvidenceId, not even an explicit nil
### GetRouteKey

`func (o *BiomarkerEvidenceRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerEvidenceRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerEvidenceRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerEvidenceRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerEvidenceRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerEvidenceRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerEvidenceRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerEvidenceRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerEvidenceRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerEvidenceRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerEvidenceRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerEvidenceRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetTargetId

`func (o *BiomarkerEvidenceRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *BiomarkerEvidenceRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *BiomarkerEvidenceRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *BiomarkerEvidenceRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *BiomarkerEvidenceRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *BiomarkerEvidenceRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetUniprotAccession

`func (o *BiomarkerEvidenceRow) GetUniprotAccession() string`

GetUniprotAccession returns the UniprotAccession field if non-nil, zero value otherwise.

### GetUniprotAccessionOk

`func (o *BiomarkerEvidenceRow) GetUniprotAccessionOk() (*string, bool)`

GetUniprotAccessionOk returns a tuple with the UniprotAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotAccession

`func (o *BiomarkerEvidenceRow) SetUniprotAccession(v string)`

SetUniprotAccession sets UniprotAccession field to given value.

### HasUniprotAccession

`func (o *BiomarkerEvidenceRow) HasUniprotAccession() bool`

HasUniprotAccession returns a boolean if a field has been set.

### SetUniprotAccessionNil

`func (o *BiomarkerEvidenceRow) SetUniprotAccessionNil(b bool)`

 SetUniprotAccessionNil sets the value for UniprotAccession to be an explicit nil

### UnsetUniprotAccession
`func (o *BiomarkerEvidenceRow) UnsetUniprotAccession()`

UnsetUniprotAccession ensures that no value is present for UniprotAccession, not even an explicit nil
### GetBiomarker

`func (o *BiomarkerEvidenceRow) GetBiomarker() string`

GetBiomarker returns the Biomarker field if non-nil, zero value otherwise.

### GetBiomarkerOk

`func (o *BiomarkerEvidenceRow) GetBiomarkerOk() (*string, bool)`

GetBiomarkerOk returns a tuple with the Biomarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarker

`func (o *BiomarkerEvidenceRow) SetBiomarker(v string)`

SetBiomarker sets Biomarker field to given value.

### HasBiomarker

`func (o *BiomarkerEvidenceRow) HasBiomarker() bool`

HasBiomarker returns a boolean if a field has been set.

### SetBiomarkerNil

`func (o *BiomarkerEvidenceRow) SetBiomarkerNil(b bool)`

 SetBiomarkerNil sets the value for Biomarker to be an explicit nil

### UnsetBiomarker
`func (o *BiomarkerEvidenceRow) UnsetBiomarker()`

UnsetBiomarker ensures that no value is present for Biomarker, not even an explicit nil
### GetEvidenceType

`func (o *BiomarkerEvidenceRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *BiomarkerEvidenceRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *BiomarkerEvidenceRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *BiomarkerEvidenceRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *BiomarkerEvidenceRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *BiomarkerEvidenceRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetSource

`func (o *BiomarkerEvidenceRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BiomarkerEvidenceRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BiomarkerEvidenceRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BiomarkerEvidenceRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *BiomarkerEvidenceRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *BiomarkerEvidenceRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetConditionId

`func (o *BiomarkerEvidenceRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *BiomarkerEvidenceRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *BiomarkerEvidenceRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *BiomarkerEvidenceRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *BiomarkerEvidenceRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *BiomarkerEvidenceRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *BiomarkerEvidenceRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *BiomarkerEvidenceRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *BiomarkerEvidenceRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *BiomarkerEvidenceRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *BiomarkerEvidenceRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *BiomarkerEvidenceRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetTherapeuticArea

`func (o *BiomarkerEvidenceRow) GetTherapeuticArea() string`

GetTherapeuticArea returns the TherapeuticArea field if non-nil, zero value otherwise.

### GetTherapeuticAreaOk

`func (o *BiomarkerEvidenceRow) GetTherapeuticAreaOk() (*string, bool)`

GetTherapeuticAreaOk returns a tuple with the TherapeuticArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTherapeuticArea

`func (o *BiomarkerEvidenceRow) SetTherapeuticArea(v string)`

SetTherapeuticArea sets TherapeuticArea field to given value.

### HasTherapeuticArea

`func (o *BiomarkerEvidenceRow) HasTherapeuticArea() bool`

HasTherapeuticArea returns a boolean if a field has been set.

### SetTherapeuticAreaNil

`func (o *BiomarkerEvidenceRow) SetTherapeuticAreaNil(b bool)`

 SetTherapeuticAreaNil sets the value for TherapeuticArea to be an explicit nil

### UnsetTherapeuticArea
`func (o *BiomarkerEvidenceRow) UnsetTherapeuticArea()`

UnsetTherapeuticArea ensures that no value is present for TherapeuticArea, not even an explicit nil
### GetDrugName

`func (o *BiomarkerEvidenceRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *BiomarkerEvidenceRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *BiomarkerEvidenceRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *BiomarkerEvidenceRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *BiomarkerEvidenceRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *BiomarkerEvidenceRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetDrugId

`func (o *BiomarkerEvidenceRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *BiomarkerEvidenceRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *BiomarkerEvidenceRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *BiomarkerEvidenceRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *BiomarkerEvidenceRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *BiomarkerEvidenceRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetMoleculeId

`func (o *BiomarkerEvidenceRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *BiomarkerEvidenceRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *BiomarkerEvidenceRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *BiomarkerEvidenceRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *BiomarkerEvidenceRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *BiomarkerEvidenceRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetResponse

`func (o *BiomarkerEvidenceRow) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *BiomarkerEvidenceRow) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *BiomarkerEvidenceRow) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *BiomarkerEvidenceRow) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *BiomarkerEvidenceRow) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *BiomarkerEvidenceRow) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetScore

`func (o *BiomarkerEvidenceRow) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BiomarkerEvidenceRow) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BiomarkerEvidenceRow) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *BiomarkerEvidenceRow) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *BiomarkerEvidenceRow) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *BiomarkerEvidenceRow) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetVariantCount

`func (o *BiomarkerEvidenceRow) GetVariantCount() int32`

GetVariantCount returns the VariantCount field if non-nil, zero value otherwise.

### GetVariantCountOk

`func (o *BiomarkerEvidenceRow) GetVariantCountOk() (*int32, bool)`

GetVariantCountOk returns a tuple with the VariantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantCount

`func (o *BiomarkerEvidenceRow) SetVariantCount(v int32)`

SetVariantCount sets VariantCount field to given value.

### HasVariantCount

`func (o *BiomarkerEvidenceRow) HasVariantCount() bool`

HasVariantCount returns a boolean if a field has been set.

### SetVariantCountNil

`func (o *BiomarkerEvidenceRow) SetVariantCountNil(b bool)`

 SetVariantCountNil sets the value for VariantCount to be an explicit nil

### UnsetVariantCount
`func (o *BiomarkerEvidenceRow) UnsetVariantCount()`

UnsetVariantCount ensures that no value is present for VariantCount, not even an explicit nil
### GetTrialId

`func (o *BiomarkerEvidenceRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *BiomarkerEvidenceRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *BiomarkerEvidenceRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *BiomarkerEvidenceRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *BiomarkerEvidenceRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *BiomarkerEvidenceRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetSampleType

`func (o *BiomarkerEvidenceRow) GetSampleType() string`

GetSampleType returns the SampleType field if non-nil, zero value otherwise.

### GetSampleTypeOk

`func (o *BiomarkerEvidenceRow) GetSampleTypeOk() (*string, bool)`

GetSampleTypeOk returns a tuple with the SampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleType

`func (o *BiomarkerEvidenceRow) SetSampleType(v string)`

SetSampleType sets SampleType field to given value.

### HasSampleType

`func (o *BiomarkerEvidenceRow) HasSampleType() bool`

HasSampleType returns a boolean if a field has been set.

### SetSampleTypeNil

`func (o *BiomarkerEvidenceRow) SetSampleTypeNil(b bool)`

 SetSampleTypeNil sets the value for SampleType to be an explicit nil

### UnsetSampleType
`func (o *BiomarkerEvidenceRow) UnsetSampleType()`

UnsetSampleType ensures that no value is present for SampleType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


