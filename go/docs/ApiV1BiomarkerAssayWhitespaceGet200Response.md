# ApiV1BiomarkerAssayWhitespaceGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**Gene** | Pointer to **string** |  | [optional] 
**GeneRecognized** | Pointer to **bool** |  | [optional] 
**InPopulation** | Pointer to **bool** |  | [optional] 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Filters** | Pointer to [**ApiV1BiomarkerAssayWhitespaceGet200ResponseFilters**](ApiV1BiomarkerAssayWhitespaceGet200ResponseFilters.md) |  | [optional] 
**BucketCounts** | Pointer to [**ApiV1BiomarkerAssayWhitespaceGet200ResponseBucketCounts**](ApiV1BiomarkerAssayWhitespaceGet200ResponseBucketCounts.md) |  | [optional] 
**NGenesMatching** | Pointer to **int32** |  | [optional] 
**NGenesReturned** | Pointer to **int32** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**Genes** | Pointer to [**[]ApiV1BiomarkerAssayWhitespaceGet200ResponseGenesInner**](ApiV1BiomarkerAssayWhitespaceGet200ResponseGenesInner.md) |  | [optional] 
**ProfileLink** | Pointer to [**ApiV1BiomarkerAssayWhitespaceGet200ResponseProfileLink**](ApiV1BiomarkerAssayWhitespaceGet200ResponseProfileLink.md) |  | [optional] 
**GeneSymbol** | Pointer to **string** |  | [optional] 
**BiomarkerDisplayName** | Pointer to **string** |  | [optional] 
**WhitespaceBucket** | Pointer to **string** |  | [optional] 
**WhitespaceBucketLabel** | Pointer to **string** |  | [optional] 
**WhitespaceBucketRuleId** | Pointer to **string** |  | [optional] 
**GeneEvidenceTier** | Pointer to **string** |  | [optional] 
**GeneEvidenceTierLabel** | Pointer to **string** |  | [optional] 
**GeneSupplyStatus** | Pointer to **string** |  | [optional] 
**GeneSupplyStatusLabel** | Pointer to **string** |  | [optional] 
**PipelineProgramCount** | Pointer to **interface{}** |  | [optional] 
**PipelineMaxPhase** | Pointer to **interface{}** |  | [optional] 
**ValidatedEvidence** | Pointer to [**[]ApiV1BiomarkerAssayWhitespaceGet200ResponseValidatedEvidenceInner**](ApiV1BiomarkerAssayWhitespaceGet200ResponseValidatedEvidenceInner.md) |  | [optional] 
**LiteratureEvidence** | Pointer to [**[]ApiV1BiomarkerAssayWhitespaceGet200ResponseLiteratureEvidenceInner**](ApiV1BiomarkerAssayWhitespaceGet200ResponseLiteratureEvidenceInner.md) |  | [optional] 
**ConfidenceCaveat** | Pointer to **string** |  | [optional] 
**SupplierDeclarations** | Pointer to [**[]ApiV1BiomarkerAssayWhitespaceGet200ResponseSupplierDeclarationsInner**](ApiV1BiomarkerAssayWhitespaceGet200ResponseSupplierDeclarationsInner.md) |  | [optional] 
**ValidatedEvidenceBuilt** | Pointer to **bool** |  | [optional] 
**LiteratureEvidenceBuilt** | Pointer to **bool** |  | [optional] 
**SupplierDeclarationsBuilt** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiV1BiomarkerAssayWhitespaceGet200Response

`func NewApiV1BiomarkerAssayWhitespaceGet200Response(mode string, buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1BiomarkerAssayWhitespaceGet200Response`

NewApiV1BiomarkerAssayWhitespaceGet200Response instantiates a new ApiV1BiomarkerAssayWhitespaceGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1BiomarkerAssayWhitespaceGet200ResponseWithDefaults

`func NewApiV1BiomarkerAssayWhitespaceGet200ResponseWithDefaults() *ApiV1BiomarkerAssayWhitespaceGet200Response`

NewApiV1BiomarkerAssayWhitespaceGet200ResponseWithDefaults instantiates a new ApiV1BiomarkerAssayWhitespaceGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetMode(v string)`

SetMode sets Mode field to given value.


### GetGene

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGene() string`

GetGene returns the Gene field if non-nil, zero value otherwise.

### GetGeneOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneOk() (*string, bool)`

GetGeneOk returns a tuple with the Gene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGene

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetGene(v string)`

SetGene sets Gene field to given value.

### HasGene

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasGene() bool`

HasGene returns a boolean if a field has been set.

### GetGeneRecognized

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneRecognized() bool`

GetGeneRecognized returns the GeneRecognized field if non-nil, zero value otherwise.

### GetGeneRecognizedOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneRecognizedOk() (*bool, bool)`

GetGeneRecognizedOk returns a tuple with the GeneRecognized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneRecognized

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetGeneRecognized(v bool)`

SetGeneRecognized sets GeneRecognized field to given value.

### HasGeneRecognized

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasGeneRecognized() bool`

HasGeneRecognized returns a boolean if a field has been set.

### GetInPopulation

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetInPopulation() bool`

GetInPopulation returns the InPopulation field if non-nil, zero value otherwise.

### GetInPopulationOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetInPopulationOk() (*bool, bool)`

GetInPopulationOk returns a tuple with the InPopulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInPopulation

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetInPopulation(v bool)`

SetInPopulation sets InPopulation field to given value.

### HasInPopulation

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasInPopulation() bool`

HasInPopulation returns a boolean if a field has been set.

### GetBuildState

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetFilters

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetFilters() ApiV1BiomarkerAssayWhitespaceGet200ResponseFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetFiltersOk() (*ApiV1BiomarkerAssayWhitespaceGet200ResponseFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetFilters(v ApiV1BiomarkerAssayWhitespaceGet200ResponseFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetBucketCounts

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetBucketCounts() ApiV1BiomarkerAssayWhitespaceGet200ResponseBucketCounts`

GetBucketCounts returns the BucketCounts field if non-nil, zero value otherwise.

### GetBucketCountsOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetBucketCountsOk() (*ApiV1BiomarkerAssayWhitespaceGet200ResponseBucketCounts, bool)`

GetBucketCountsOk returns a tuple with the BucketCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketCounts

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetBucketCounts(v ApiV1BiomarkerAssayWhitespaceGet200ResponseBucketCounts)`

SetBucketCounts sets BucketCounts field to given value.

### HasBucketCounts

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasBucketCounts() bool`

HasBucketCounts returns a boolean if a field has been set.

### GetNGenesMatching

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetNGenesMatching() int32`

GetNGenesMatching returns the NGenesMatching field if non-nil, zero value otherwise.

### GetNGenesMatchingOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetNGenesMatchingOk() (*int32, bool)`

GetNGenesMatchingOk returns a tuple with the NGenesMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNGenesMatching

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetNGenesMatching(v int32)`

SetNGenesMatching sets NGenesMatching field to given value.

### HasNGenesMatching

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasNGenesMatching() bool`

HasNGenesMatching returns a boolean if a field has been set.

### GetNGenesReturned

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetNGenesReturned() int32`

GetNGenesReturned returns the NGenesReturned field if non-nil, zero value otherwise.

### GetNGenesReturnedOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetNGenesReturnedOk() (*int32, bool)`

GetNGenesReturnedOk returns a tuple with the NGenesReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNGenesReturned

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetNGenesReturned(v int32)`

SetNGenesReturned sets NGenesReturned field to given value.

### HasNGenesReturned

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasNGenesReturned() bool`

HasNGenesReturned returns a boolean if a field has been set.

### GetTruncated

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetGenes

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGenes() []ApiV1BiomarkerAssayWhitespaceGet200ResponseGenesInner`

GetGenes returns the Genes field if non-nil, zero value otherwise.

### GetGenesOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGenesOk() (*[]ApiV1BiomarkerAssayWhitespaceGet200ResponseGenesInner, bool)`

GetGenesOk returns a tuple with the Genes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenes

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetGenes(v []ApiV1BiomarkerAssayWhitespaceGet200ResponseGenesInner)`

SetGenes sets Genes field to given value.

### HasGenes

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasGenes() bool`

HasGenes returns a boolean if a field has been set.

### GetProfileLink

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetProfileLink() ApiV1BiomarkerAssayWhitespaceGet200ResponseProfileLink`

GetProfileLink returns the ProfileLink field if non-nil, zero value otherwise.

### GetProfileLinkOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetProfileLinkOk() (*ApiV1BiomarkerAssayWhitespaceGet200ResponseProfileLink, bool)`

GetProfileLinkOk returns a tuple with the ProfileLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileLink

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetProfileLink(v ApiV1BiomarkerAssayWhitespaceGet200ResponseProfileLink)`

SetProfileLink sets ProfileLink field to given value.

### HasProfileLink

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasProfileLink() bool`

HasProfileLink returns a boolean if a field has been set.

### GetGeneSymbol

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### GetBiomarkerDisplayName

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetBiomarkerDisplayName() string`

GetBiomarkerDisplayName returns the BiomarkerDisplayName field if non-nil, zero value otherwise.

### GetBiomarkerDisplayNameOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetBiomarkerDisplayNameOk() (*string, bool)`

GetBiomarkerDisplayNameOk returns a tuple with the BiomarkerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerDisplayName

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetBiomarkerDisplayName(v string)`

SetBiomarkerDisplayName sets BiomarkerDisplayName field to given value.

### HasBiomarkerDisplayName

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasBiomarkerDisplayName() bool`

HasBiomarkerDisplayName returns a boolean if a field has been set.

### GetWhitespaceBucket

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetWhitespaceBucket() string`

GetWhitespaceBucket returns the WhitespaceBucket field if non-nil, zero value otherwise.

### GetWhitespaceBucketOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetWhitespaceBucketOk() (*string, bool)`

GetWhitespaceBucketOk returns a tuple with the WhitespaceBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitespaceBucket

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetWhitespaceBucket(v string)`

SetWhitespaceBucket sets WhitespaceBucket field to given value.

### HasWhitespaceBucket

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasWhitespaceBucket() bool`

HasWhitespaceBucket returns a boolean if a field has been set.

### GetWhitespaceBucketLabel

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetWhitespaceBucketLabel() string`

GetWhitespaceBucketLabel returns the WhitespaceBucketLabel field if non-nil, zero value otherwise.

### GetWhitespaceBucketLabelOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetWhitespaceBucketLabelOk() (*string, bool)`

GetWhitespaceBucketLabelOk returns a tuple with the WhitespaceBucketLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitespaceBucketLabel

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetWhitespaceBucketLabel(v string)`

SetWhitespaceBucketLabel sets WhitespaceBucketLabel field to given value.

### HasWhitespaceBucketLabel

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasWhitespaceBucketLabel() bool`

HasWhitespaceBucketLabel returns a boolean if a field has been set.

### GetWhitespaceBucketRuleId

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetWhitespaceBucketRuleId() string`

GetWhitespaceBucketRuleId returns the WhitespaceBucketRuleId field if non-nil, zero value otherwise.

### GetWhitespaceBucketRuleIdOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetWhitespaceBucketRuleIdOk() (*string, bool)`

GetWhitespaceBucketRuleIdOk returns a tuple with the WhitespaceBucketRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitespaceBucketRuleId

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetWhitespaceBucketRuleId(v string)`

SetWhitespaceBucketRuleId sets WhitespaceBucketRuleId field to given value.

### HasWhitespaceBucketRuleId

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasWhitespaceBucketRuleId() bool`

HasWhitespaceBucketRuleId returns a boolean if a field has been set.

### GetGeneEvidenceTier

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneEvidenceTier() string`

GetGeneEvidenceTier returns the GeneEvidenceTier field if non-nil, zero value otherwise.

### GetGeneEvidenceTierOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneEvidenceTierOk() (*string, bool)`

GetGeneEvidenceTierOk returns a tuple with the GeneEvidenceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneEvidenceTier

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetGeneEvidenceTier(v string)`

SetGeneEvidenceTier sets GeneEvidenceTier field to given value.

### HasGeneEvidenceTier

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasGeneEvidenceTier() bool`

HasGeneEvidenceTier returns a boolean if a field has been set.

### GetGeneEvidenceTierLabel

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneEvidenceTierLabel() string`

GetGeneEvidenceTierLabel returns the GeneEvidenceTierLabel field if non-nil, zero value otherwise.

### GetGeneEvidenceTierLabelOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneEvidenceTierLabelOk() (*string, bool)`

GetGeneEvidenceTierLabelOk returns a tuple with the GeneEvidenceTierLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneEvidenceTierLabel

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetGeneEvidenceTierLabel(v string)`

SetGeneEvidenceTierLabel sets GeneEvidenceTierLabel field to given value.

### HasGeneEvidenceTierLabel

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasGeneEvidenceTierLabel() bool`

HasGeneEvidenceTierLabel returns a boolean if a field has been set.

### GetGeneSupplyStatus

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneSupplyStatus() string`

GetGeneSupplyStatus returns the GeneSupplyStatus field if non-nil, zero value otherwise.

### GetGeneSupplyStatusOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneSupplyStatusOk() (*string, bool)`

GetGeneSupplyStatusOk returns a tuple with the GeneSupplyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSupplyStatus

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetGeneSupplyStatus(v string)`

SetGeneSupplyStatus sets GeneSupplyStatus field to given value.

### HasGeneSupplyStatus

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasGeneSupplyStatus() bool`

HasGeneSupplyStatus returns a boolean if a field has been set.

### GetGeneSupplyStatusLabel

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneSupplyStatusLabel() string`

GetGeneSupplyStatusLabel returns the GeneSupplyStatusLabel field if non-nil, zero value otherwise.

### GetGeneSupplyStatusLabelOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetGeneSupplyStatusLabelOk() (*string, bool)`

GetGeneSupplyStatusLabelOk returns a tuple with the GeneSupplyStatusLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSupplyStatusLabel

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetGeneSupplyStatusLabel(v string)`

SetGeneSupplyStatusLabel sets GeneSupplyStatusLabel field to given value.

### HasGeneSupplyStatusLabel

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasGeneSupplyStatusLabel() bool`

HasGeneSupplyStatusLabel returns a boolean if a field has been set.

### GetPipelineProgramCount

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetPipelineProgramCount() interface{}`

GetPipelineProgramCount returns the PipelineProgramCount field if non-nil, zero value otherwise.

### GetPipelineProgramCountOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetPipelineProgramCountOk() (*interface{}, bool)`

GetPipelineProgramCountOk returns a tuple with the PipelineProgramCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineProgramCount

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetPipelineProgramCount(v interface{})`

SetPipelineProgramCount sets PipelineProgramCount field to given value.

### HasPipelineProgramCount

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasPipelineProgramCount() bool`

HasPipelineProgramCount returns a boolean if a field has been set.

### SetPipelineProgramCountNil

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetPipelineProgramCountNil(b bool)`

 SetPipelineProgramCountNil sets the value for PipelineProgramCount to be an explicit nil

### UnsetPipelineProgramCount
`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) UnsetPipelineProgramCount()`

UnsetPipelineProgramCount ensures that no value is present for PipelineProgramCount, not even an explicit nil
### GetPipelineMaxPhase

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetPipelineMaxPhase() interface{}`

GetPipelineMaxPhase returns the PipelineMaxPhase field if non-nil, zero value otherwise.

### GetPipelineMaxPhaseOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetPipelineMaxPhaseOk() (*interface{}, bool)`

GetPipelineMaxPhaseOk returns a tuple with the PipelineMaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineMaxPhase

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetPipelineMaxPhase(v interface{})`

SetPipelineMaxPhase sets PipelineMaxPhase field to given value.

### HasPipelineMaxPhase

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasPipelineMaxPhase() bool`

HasPipelineMaxPhase returns a boolean if a field has been set.

### SetPipelineMaxPhaseNil

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetPipelineMaxPhaseNil(b bool)`

 SetPipelineMaxPhaseNil sets the value for PipelineMaxPhase to be an explicit nil

### UnsetPipelineMaxPhase
`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) UnsetPipelineMaxPhase()`

UnsetPipelineMaxPhase ensures that no value is present for PipelineMaxPhase, not even an explicit nil
### GetValidatedEvidence

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetValidatedEvidence() []ApiV1BiomarkerAssayWhitespaceGet200ResponseValidatedEvidenceInner`

GetValidatedEvidence returns the ValidatedEvidence field if non-nil, zero value otherwise.

### GetValidatedEvidenceOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetValidatedEvidenceOk() (*[]ApiV1BiomarkerAssayWhitespaceGet200ResponseValidatedEvidenceInner, bool)`

GetValidatedEvidenceOk returns a tuple with the ValidatedEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedEvidence

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetValidatedEvidence(v []ApiV1BiomarkerAssayWhitespaceGet200ResponseValidatedEvidenceInner)`

SetValidatedEvidence sets ValidatedEvidence field to given value.

### HasValidatedEvidence

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasValidatedEvidence() bool`

HasValidatedEvidence returns a boolean if a field has been set.

### GetLiteratureEvidence

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetLiteratureEvidence() []ApiV1BiomarkerAssayWhitespaceGet200ResponseLiteratureEvidenceInner`

GetLiteratureEvidence returns the LiteratureEvidence field if non-nil, zero value otherwise.

### GetLiteratureEvidenceOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetLiteratureEvidenceOk() (*[]ApiV1BiomarkerAssayWhitespaceGet200ResponseLiteratureEvidenceInner, bool)`

GetLiteratureEvidenceOk returns a tuple with the LiteratureEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiteratureEvidence

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetLiteratureEvidence(v []ApiV1BiomarkerAssayWhitespaceGet200ResponseLiteratureEvidenceInner)`

SetLiteratureEvidence sets LiteratureEvidence field to given value.

### HasLiteratureEvidence

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasLiteratureEvidence() bool`

HasLiteratureEvidence returns a boolean if a field has been set.

### GetConfidenceCaveat

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetConfidenceCaveat() string`

GetConfidenceCaveat returns the ConfidenceCaveat field if non-nil, zero value otherwise.

### GetConfidenceCaveatOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetConfidenceCaveatOk() (*string, bool)`

GetConfidenceCaveatOk returns a tuple with the ConfidenceCaveat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceCaveat

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetConfidenceCaveat(v string)`

SetConfidenceCaveat sets ConfidenceCaveat field to given value.

### HasConfidenceCaveat

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasConfidenceCaveat() bool`

HasConfidenceCaveat returns a boolean if a field has been set.

### GetSupplierDeclarations

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetSupplierDeclarations() []ApiV1BiomarkerAssayWhitespaceGet200ResponseSupplierDeclarationsInner`

GetSupplierDeclarations returns the SupplierDeclarations field if non-nil, zero value otherwise.

### GetSupplierDeclarationsOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetSupplierDeclarationsOk() (*[]ApiV1BiomarkerAssayWhitespaceGet200ResponseSupplierDeclarationsInner, bool)`

GetSupplierDeclarationsOk returns a tuple with the SupplierDeclarations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierDeclarations

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetSupplierDeclarations(v []ApiV1BiomarkerAssayWhitespaceGet200ResponseSupplierDeclarationsInner)`

SetSupplierDeclarations sets SupplierDeclarations field to given value.

### HasSupplierDeclarations

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasSupplierDeclarations() bool`

HasSupplierDeclarations returns a boolean if a field has been set.

### GetValidatedEvidenceBuilt

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetValidatedEvidenceBuilt() bool`

GetValidatedEvidenceBuilt returns the ValidatedEvidenceBuilt field if non-nil, zero value otherwise.

### GetValidatedEvidenceBuiltOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetValidatedEvidenceBuiltOk() (*bool, bool)`

GetValidatedEvidenceBuiltOk returns a tuple with the ValidatedEvidenceBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedEvidenceBuilt

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetValidatedEvidenceBuilt(v bool)`

SetValidatedEvidenceBuilt sets ValidatedEvidenceBuilt field to given value.

### HasValidatedEvidenceBuilt

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasValidatedEvidenceBuilt() bool`

HasValidatedEvidenceBuilt returns a boolean if a field has been set.

### GetLiteratureEvidenceBuilt

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetLiteratureEvidenceBuilt() bool`

GetLiteratureEvidenceBuilt returns the LiteratureEvidenceBuilt field if non-nil, zero value otherwise.

### GetLiteratureEvidenceBuiltOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetLiteratureEvidenceBuiltOk() (*bool, bool)`

GetLiteratureEvidenceBuiltOk returns a tuple with the LiteratureEvidenceBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiteratureEvidenceBuilt

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetLiteratureEvidenceBuilt(v bool)`

SetLiteratureEvidenceBuilt sets LiteratureEvidenceBuilt field to given value.

### HasLiteratureEvidenceBuilt

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasLiteratureEvidenceBuilt() bool`

HasLiteratureEvidenceBuilt returns a boolean if a field has been set.

### GetSupplierDeclarationsBuilt

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetSupplierDeclarationsBuilt() bool`

GetSupplierDeclarationsBuilt returns the SupplierDeclarationsBuilt field if non-nil, zero value otherwise.

### GetSupplierDeclarationsBuiltOk

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) GetSupplierDeclarationsBuiltOk() (*bool, bool)`

GetSupplierDeclarationsBuiltOk returns a tuple with the SupplierDeclarationsBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierDeclarationsBuilt

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) SetSupplierDeclarationsBuilt(v bool)`

SetSupplierDeclarationsBuilt sets SupplierDeclarationsBuilt field to given value.

### HasSupplierDeclarationsBuilt

`func (o *ApiV1BiomarkerAssayWhitespaceGet200Response) HasSupplierDeclarationsBuilt() bool`

HasSupplierDeclarationsBuilt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


