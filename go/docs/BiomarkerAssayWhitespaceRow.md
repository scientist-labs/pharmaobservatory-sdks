# BiomarkerAssayWhitespaceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiomarkerId** | Pointer to **NullableString** | The gene&#39;s biomarker key. | [optional] 
**RouteKey** | Pointer to **NullableString** | The gene&#39;s route key, the entry key for the biomarkers object type; the gene symbol. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol. | [optional] 
**BiomarkerDisplayName** | Pointer to **NullableString** | The gene&#39;s display name. | [optional] 
**BiomarkerIdentityStatus** | Pointer to **NullableString** | Whether the gene&#39;s identity is confirmed. | [optional] 
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id; null when the gene does not resolve one, in which case the pipeline columns are null. | [optional] 
**WhitespaceBucket** | Pointer to **NullableString** | The classification, one of nine values crossing gene_evidence_tier with gene_supply_status. Labels: no_evidence_no_supply &#39;No evidence, explicit no-assay&#39;; no_evidence_supply_ambiguous &#39;No evidence, ambiguous supply&#39;; supply_no_evidence &#39;Supply declared, no evidence&#39;; literature_evidence_no_supply &#39;Literature evidence, no confirmed supply&#39;; literature_evidence_supply_ambiguous &#39;Literature evidence, ambiguous supply&#39;; literature_evidence_served &#39;Literature evidence, supply declared&#39;; validated_evidence_no_supply &#39;Validated-ladder evidence, no confirmed supply&#39;; validated_evidence_supply_ambiguous &#39;Validated-ladder evidence, ambiguous supply&#39;; served &#39;Validated-ladder evidence, supply declared&#39;. | [optional] 
**WhitespaceBucketRuleId** | Pointer to **NullableString** | The rule that assigned the bucket, one rule per bucket. | [optional] 
**GeneEvidenceTier** | Pointer to **NullableString** | Evidence axis, a precedence and not a score. validated: at least one row of the validation ladder (biomarker_validation) for the gene. literature_only: no validated row but at least one extracted PubMed assay fact. none: neither. | [optional] 
**GeneSupplyStatus** | Pointer to **NullableString** | Supply axis, a precedence. declared_positive: a supplier declares an assay (validated or not). declared_ambiguous: a response conflicts or gives detail with no option chosen. declared_no_assay: a supplier answered that it has no assay. not_on_survey: no answer at all. | [optional] 
**ValidatedEvidenceRowCount** | Pointer to **NullableInt32** | Rows of the validation ladder for the gene (biomarker_validation). A count, not a score, and not combined with the literature count. | [optional] 
**LiteratureEvidenceRowCount** | Pointer to **NullableInt32** | Extracted PubMed assay facts resolved to the gene (biomarker_assay_evidence). A count, not a score. | [optional] 
**EvidenceTypesPresent** | Pointer to **[]string** | The sorted set of evidence types behind the gene&#39;s evidence, including literature_assay_validation when literature evidence exists. A set, never a rank. | [optional] 
**EvidenceTierRuleIdsPresent** | Pointer to **[]string** | The sorted set of validation ladder rule ids behind the validated evidence. For display; do not derive a rank from it. | [optional] 
**EvidenceValidationIds** | Pointer to **[]string** | biomarker_validation_id values of the validated evidence. Empty when gene_evidence_tier is not validated. | [optional] 
**LiteratureEvidenceIds** | Pointer to **[]string** | biomarker_assay_evidence_id values of the extracted PubMed facts. Populated whichever evidence tier governs the bucket. | [optional] 
**LiteratureEvidenceDocumentIds** | Pointer to **[]string** | The source document ids behind literature_evidence_ids. | [optional] 
**SupplierDeclarationIds** | Pointer to **[]string** | Ids of the supplier survey declarations behind gene_supply_status (the winning precedence tier). Empty when not_on_survey. The declarations are not a published dataset. | [optional] 
**SupplierPositiveResponseCount** | Pointer to **NullableInt32** | Supplier declarations that an assay exists, validated or unvalidated, for the gene. | [optional] 
**SupplierConflictingResponseCount** | Pointer to **NullableInt32** | Declarations whose own cells disagree. | [optional] 
**SupplierDetailOnlyResponseCount** | Pointer to **NullableInt32** | Declarations that give free text and no selected option. | [optional] 
**SupplierNoAssayResponseCount** | Pointer to **NullableInt32** | Declarations that the supplier has no assay for the gene. | [optional] 
**PipelineProgramCount** | Pointer to **NullableInt32** | Drug programs against the gene&#39;s target across every phase band. Null means the gene has no target row, never zero activity. Heat only; not part of the bucket. | [optional] 
**PipelineActiveProgramCount** | Pointer to **NullableInt32** | Active drug programs against the gene&#39;s target. Same nullability as pipeline_program_count. | [optional] 
**PipelineMaxPhase** | Pointer to **NullableInt32** | The highest phase band (0 to 4) any program against the target reached; the app maps it to a phase label. Null as pipeline_program_count. | [optional] 

## Methods

### NewBiomarkerAssayWhitespaceRow

`func NewBiomarkerAssayWhitespaceRow() *BiomarkerAssayWhitespaceRow`

NewBiomarkerAssayWhitespaceRow instantiates a new BiomarkerAssayWhitespaceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerAssayWhitespaceRowWithDefaults

`func NewBiomarkerAssayWhitespaceRowWithDefaults() *BiomarkerAssayWhitespaceRow`

NewBiomarkerAssayWhitespaceRowWithDefaults instantiates a new BiomarkerAssayWhitespaceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiomarkerId

`func (o *BiomarkerAssayWhitespaceRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkerAssayWhitespaceRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkerAssayWhitespaceRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkerAssayWhitespaceRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkerAssayWhitespaceRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkerAssayWhitespaceRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetRouteKey

`func (o *BiomarkerAssayWhitespaceRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerAssayWhitespaceRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerAssayWhitespaceRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerAssayWhitespaceRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerAssayWhitespaceRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerAssayWhitespaceRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerAssayWhitespaceRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerAssayWhitespaceRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerAssayWhitespaceRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerAssayWhitespaceRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerAssayWhitespaceRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerAssayWhitespaceRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetBiomarkerDisplayName

`func (o *BiomarkerAssayWhitespaceRow) GetBiomarkerDisplayName() string`

GetBiomarkerDisplayName returns the BiomarkerDisplayName field if non-nil, zero value otherwise.

### GetBiomarkerDisplayNameOk

`func (o *BiomarkerAssayWhitespaceRow) GetBiomarkerDisplayNameOk() (*string, bool)`

GetBiomarkerDisplayNameOk returns a tuple with the BiomarkerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerDisplayName

`func (o *BiomarkerAssayWhitespaceRow) SetBiomarkerDisplayName(v string)`

SetBiomarkerDisplayName sets BiomarkerDisplayName field to given value.

### HasBiomarkerDisplayName

`func (o *BiomarkerAssayWhitespaceRow) HasBiomarkerDisplayName() bool`

HasBiomarkerDisplayName returns a boolean if a field has been set.

### SetBiomarkerDisplayNameNil

`func (o *BiomarkerAssayWhitespaceRow) SetBiomarkerDisplayNameNil(b bool)`

 SetBiomarkerDisplayNameNil sets the value for BiomarkerDisplayName to be an explicit nil

### UnsetBiomarkerDisplayName
`func (o *BiomarkerAssayWhitespaceRow) UnsetBiomarkerDisplayName()`

UnsetBiomarkerDisplayName ensures that no value is present for BiomarkerDisplayName, not even an explicit nil
### GetBiomarkerIdentityStatus

`func (o *BiomarkerAssayWhitespaceRow) GetBiomarkerIdentityStatus() string`

GetBiomarkerIdentityStatus returns the BiomarkerIdentityStatus field if non-nil, zero value otherwise.

### GetBiomarkerIdentityStatusOk

`func (o *BiomarkerAssayWhitespaceRow) GetBiomarkerIdentityStatusOk() (*string, bool)`

GetBiomarkerIdentityStatusOk returns a tuple with the BiomarkerIdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerIdentityStatus

`func (o *BiomarkerAssayWhitespaceRow) SetBiomarkerIdentityStatus(v string)`

SetBiomarkerIdentityStatus sets BiomarkerIdentityStatus field to given value.

### HasBiomarkerIdentityStatus

`func (o *BiomarkerAssayWhitespaceRow) HasBiomarkerIdentityStatus() bool`

HasBiomarkerIdentityStatus returns a boolean if a field has been set.

### SetBiomarkerIdentityStatusNil

`func (o *BiomarkerAssayWhitespaceRow) SetBiomarkerIdentityStatusNil(b bool)`

 SetBiomarkerIdentityStatusNil sets the value for BiomarkerIdentityStatus to be an explicit nil

### UnsetBiomarkerIdentityStatus
`func (o *BiomarkerAssayWhitespaceRow) UnsetBiomarkerIdentityStatus()`

UnsetBiomarkerIdentityStatus ensures that no value is present for BiomarkerIdentityStatus, not even an explicit nil
### GetTargetId

`func (o *BiomarkerAssayWhitespaceRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *BiomarkerAssayWhitespaceRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *BiomarkerAssayWhitespaceRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *BiomarkerAssayWhitespaceRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *BiomarkerAssayWhitespaceRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *BiomarkerAssayWhitespaceRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetWhitespaceBucket

`func (o *BiomarkerAssayWhitespaceRow) GetWhitespaceBucket() string`

GetWhitespaceBucket returns the WhitespaceBucket field if non-nil, zero value otherwise.

### GetWhitespaceBucketOk

`func (o *BiomarkerAssayWhitespaceRow) GetWhitespaceBucketOk() (*string, bool)`

GetWhitespaceBucketOk returns a tuple with the WhitespaceBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitespaceBucket

`func (o *BiomarkerAssayWhitespaceRow) SetWhitespaceBucket(v string)`

SetWhitespaceBucket sets WhitespaceBucket field to given value.

### HasWhitespaceBucket

`func (o *BiomarkerAssayWhitespaceRow) HasWhitespaceBucket() bool`

HasWhitespaceBucket returns a boolean if a field has been set.

### SetWhitespaceBucketNil

`func (o *BiomarkerAssayWhitespaceRow) SetWhitespaceBucketNil(b bool)`

 SetWhitespaceBucketNil sets the value for WhitespaceBucket to be an explicit nil

### UnsetWhitespaceBucket
`func (o *BiomarkerAssayWhitespaceRow) UnsetWhitespaceBucket()`

UnsetWhitespaceBucket ensures that no value is present for WhitespaceBucket, not even an explicit nil
### GetWhitespaceBucketRuleId

`func (o *BiomarkerAssayWhitespaceRow) GetWhitespaceBucketRuleId() string`

GetWhitespaceBucketRuleId returns the WhitespaceBucketRuleId field if non-nil, zero value otherwise.

### GetWhitespaceBucketRuleIdOk

`func (o *BiomarkerAssayWhitespaceRow) GetWhitespaceBucketRuleIdOk() (*string, bool)`

GetWhitespaceBucketRuleIdOk returns a tuple with the WhitespaceBucketRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitespaceBucketRuleId

`func (o *BiomarkerAssayWhitespaceRow) SetWhitespaceBucketRuleId(v string)`

SetWhitespaceBucketRuleId sets WhitespaceBucketRuleId field to given value.

### HasWhitespaceBucketRuleId

`func (o *BiomarkerAssayWhitespaceRow) HasWhitespaceBucketRuleId() bool`

HasWhitespaceBucketRuleId returns a boolean if a field has been set.

### SetWhitespaceBucketRuleIdNil

`func (o *BiomarkerAssayWhitespaceRow) SetWhitespaceBucketRuleIdNil(b bool)`

 SetWhitespaceBucketRuleIdNil sets the value for WhitespaceBucketRuleId to be an explicit nil

### UnsetWhitespaceBucketRuleId
`func (o *BiomarkerAssayWhitespaceRow) UnsetWhitespaceBucketRuleId()`

UnsetWhitespaceBucketRuleId ensures that no value is present for WhitespaceBucketRuleId, not even an explicit nil
### GetGeneEvidenceTier

`func (o *BiomarkerAssayWhitespaceRow) GetGeneEvidenceTier() string`

GetGeneEvidenceTier returns the GeneEvidenceTier field if non-nil, zero value otherwise.

### GetGeneEvidenceTierOk

`func (o *BiomarkerAssayWhitespaceRow) GetGeneEvidenceTierOk() (*string, bool)`

GetGeneEvidenceTierOk returns a tuple with the GeneEvidenceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneEvidenceTier

`func (o *BiomarkerAssayWhitespaceRow) SetGeneEvidenceTier(v string)`

SetGeneEvidenceTier sets GeneEvidenceTier field to given value.

### HasGeneEvidenceTier

`func (o *BiomarkerAssayWhitespaceRow) HasGeneEvidenceTier() bool`

HasGeneEvidenceTier returns a boolean if a field has been set.

### SetGeneEvidenceTierNil

`func (o *BiomarkerAssayWhitespaceRow) SetGeneEvidenceTierNil(b bool)`

 SetGeneEvidenceTierNil sets the value for GeneEvidenceTier to be an explicit nil

### UnsetGeneEvidenceTier
`func (o *BiomarkerAssayWhitespaceRow) UnsetGeneEvidenceTier()`

UnsetGeneEvidenceTier ensures that no value is present for GeneEvidenceTier, not even an explicit nil
### GetGeneSupplyStatus

`func (o *BiomarkerAssayWhitespaceRow) GetGeneSupplyStatus() string`

GetGeneSupplyStatus returns the GeneSupplyStatus field if non-nil, zero value otherwise.

### GetGeneSupplyStatusOk

`func (o *BiomarkerAssayWhitespaceRow) GetGeneSupplyStatusOk() (*string, bool)`

GetGeneSupplyStatusOk returns a tuple with the GeneSupplyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSupplyStatus

`func (o *BiomarkerAssayWhitespaceRow) SetGeneSupplyStatus(v string)`

SetGeneSupplyStatus sets GeneSupplyStatus field to given value.

### HasGeneSupplyStatus

`func (o *BiomarkerAssayWhitespaceRow) HasGeneSupplyStatus() bool`

HasGeneSupplyStatus returns a boolean if a field has been set.

### SetGeneSupplyStatusNil

`func (o *BiomarkerAssayWhitespaceRow) SetGeneSupplyStatusNil(b bool)`

 SetGeneSupplyStatusNil sets the value for GeneSupplyStatus to be an explicit nil

### UnsetGeneSupplyStatus
`func (o *BiomarkerAssayWhitespaceRow) UnsetGeneSupplyStatus()`

UnsetGeneSupplyStatus ensures that no value is present for GeneSupplyStatus, not even an explicit nil
### GetValidatedEvidenceRowCount

`func (o *BiomarkerAssayWhitespaceRow) GetValidatedEvidenceRowCount() int32`

GetValidatedEvidenceRowCount returns the ValidatedEvidenceRowCount field if non-nil, zero value otherwise.

### GetValidatedEvidenceRowCountOk

`func (o *BiomarkerAssayWhitespaceRow) GetValidatedEvidenceRowCountOk() (*int32, bool)`

GetValidatedEvidenceRowCountOk returns a tuple with the ValidatedEvidenceRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedEvidenceRowCount

`func (o *BiomarkerAssayWhitespaceRow) SetValidatedEvidenceRowCount(v int32)`

SetValidatedEvidenceRowCount sets ValidatedEvidenceRowCount field to given value.

### HasValidatedEvidenceRowCount

`func (o *BiomarkerAssayWhitespaceRow) HasValidatedEvidenceRowCount() bool`

HasValidatedEvidenceRowCount returns a boolean if a field has been set.

### SetValidatedEvidenceRowCountNil

`func (o *BiomarkerAssayWhitespaceRow) SetValidatedEvidenceRowCountNil(b bool)`

 SetValidatedEvidenceRowCountNil sets the value for ValidatedEvidenceRowCount to be an explicit nil

### UnsetValidatedEvidenceRowCount
`func (o *BiomarkerAssayWhitespaceRow) UnsetValidatedEvidenceRowCount()`

UnsetValidatedEvidenceRowCount ensures that no value is present for ValidatedEvidenceRowCount, not even an explicit nil
### GetLiteratureEvidenceRowCount

`func (o *BiomarkerAssayWhitespaceRow) GetLiteratureEvidenceRowCount() int32`

GetLiteratureEvidenceRowCount returns the LiteratureEvidenceRowCount field if non-nil, zero value otherwise.

### GetLiteratureEvidenceRowCountOk

`func (o *BiomarkerAssayWhitespaceRow) GetLiteratureEvidenceRowCountOk() (*int32, bool)`

GetLiteratureEvidenceRowCountOk returns a tuple with the LiteratureEvidenceRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiteratureEvidenceRowCount

`func (o *BiomarkerAssayWhitespaceRow) SetLiteratureEvidenceRowCount(v int32)`

SetLiteratureEvidenceRowCount sets LiteratureEvidenceRowCount field to given value.

### HasLiteratureEvidenceRowCount

`func (o *BiomarkerAssayWhitespaceRow) HasLiteratureEvidenceRowCount() bool`

HasLiteratureEvidenceRowCount returns a boolean if a field has been set.

### SetLiteratureEvidenceRowCountNil

`func (o *BiomarkerAssayWhitespaceRow) SetLiteratureEvidenceRowCountNil(b bool)`

 SetLiteratureEvidenceRowCountNil sets the value for LiteratureEvidenceRowCount to be an explicit nil

### UnsetLiteratureEvidenceRowCount
`func (o *BiomarkerAssayWhitespaceRow) UnsetLiteratureEvidenceRowCount()`

UnsetLiteratureEvidenceRowCount ensures that no value is present for LiteratureEvidenceRowCount, not even an explicit nil
### GetEvidenceTypesPresent

`func (o *BiomarkerAssayWhitespaceRow) GetEvidenceTypesPresent() []string`

GetEvidenceTypesPresent returns the EvidenceTypesPresent field if non-nil, zero value otherwise.

### GetEvidenceTypesPresentOk

`func (o *BiomarkerAssayWhitespaceRow) GetEvidenceTypesPresentOk() (*[]string, bool)`

GetEvidenceTypesPresentOk returns a tuple with the EvidenceTypesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceTypesPresent

`func (o *BiomarkerAssayWhitespaceRow) SetEvidenceTypesPresent(v []string)`

SetEvidenceTypesPresent sets EvidenceTypesPresent field to given value.

### HasEvidenceTypesPresent

`func (o *BiomarkerAssayWhitespaceRow) HasEvidenceTypesPresent() bool`

HasEvidenceTypesPresent returns a boolean if a field has been set.

### SetEvidenceTypesPresentNil

`func (o *BiomarkerAssayWhitespaceRow) SetEvidenceTypesPresentNil(b bool)`

 SetEvidenceTypesPresentNil sets the value for EvidenceTypesPresent to be an explicit nil

### UnsetEvidenceTypesPresent
`func (o *BiomarkerAssayWhitespaceRow) UnsetEvidenceTypesPresent()`

UnsetEvidenceTypesPresent ensures that no value is present for EvidenceTypesPresent, not even an explicit nil
### GetEvidenceTierRuleIdsPresent

`func (o *BiomarkerAssayWhitespaceRow) GetEvidenceTierRuleIdsPresent() []string`

GetEvidenceTierRuleIdsPresent returns the EvidenceTierRuleIdsPresent field if non-nil, zero value otherwise.

### GetEvidenceTierRuleIdsPresentOk

`func (o *BiomarkerAssayWhitespaceRow) GetEvidenceTierRuleIdsPresentOk() (*[]string, bool)`

GetEvidenceTierRuleIdsPresentOk returns a tuple with the EvidenceTierRuleIdsPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceTierRuleIdsPresent

`func (o *BiomarkerAssayWhitespaceRow) SetEvidenceTierRuleIdsPresent(v []string)`

SetEvidenceTierRuleIdsPresent sets EvidenceTierRuleIdsPresent field to given value.

### HasEvidenceTierRuleIdsPresent

`func (o *BiomarkerAssayWhitespaceRow) HasEvidenceTierRuleIdsPresent() bool`

HasEvidenceTierRuleIdsPresent returns a boolean if a field has been set.

### SetEvidenceTierRuleIdsPresentNil

`func (o *BiomarkerAssayWhitespaceRow) SetEvidenceTierRuleIdsPresentNil(b bool)`

 SetEvidenceTierRuleIdsPresentNil sets the value for EvidenceTierRuleIdsPresent to be an explicit nil

### UnsetEvidenceTierRuleIdsPresent
`func (o *BiomarkerAssayWhitespaceRow) UnsetEvidenceTierRuleIdsPresent()`

UnsetEvidenceTierRuleIdsPresent ensures that no value is present for EvidenceTierRuleIdsPresent, not even an explicit nil
### GetEvidenceValidationIds

`func (o *BiomarkerAssayWhitespaceRow) GetEvidenceValidationIds() []string`

GetEvidenceValidationIds returns the EvidenceValidationIds field if non-nil, zero value otherwise.

### GetEvidenceValidationIdsOk

`func (o *BiomarkerAssayWhitespaceRow) GetEvidenceValidationIdsOk() (*[]string, bool)`

GetEvidenceValidationIdsOk returns a tuple with the EvidenceValidationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceValidationIds

`func (o *BiomarkerAssayWhitespaceRow) SetEvidenceValidationIds(v []string)`

SetEvidenceValidationIds sets EvidenceValidationIds field to given value.

### HasEvidenceValidationIds

`func (o *BiomarkerAssayWhitespaceRow) HasEvidenceValidationIds() bool`

HasEvidenceValidationIds returns a boolean if a field has been set.

### SetEvidenceValidationIdsNil

`func (o *BiomarkerAssayWhitespaceRow) SetEvidenceValidationIdsNil(b bool)`

 SetEvidenceValidationIdsNil sets the value for EvidenceValidationIds to be an explicit nil

### UnsetEvidenceValidationIds
`func (o *BiomarkerAssayWhitespaceRow) UnsetEvidenceValidationIds()`

UnsetEvidenceValidationIds ensures that no value is present for EvidenceValidationIds, not even an explicit nil
### GetLiteratureEvidenceIds

`func (o *BiomarkerAssayWhitespaceRow) GetLiteratureEvidenceIds() []string`

GetLiteratureEvidenceIds returns the LiteratureEvidenceIds field if non-nil, zero value otherwise.

### GetLiteratureEvidenceIdsOk

`func (o *BiomarkerAssayWhitespaceRow) GetLiteratureEvidenceIdsOk() (*[]string, bool)`

GetLiteratureEvidenceIdsOk returns a tuple with the LiteratureEvidenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiteratureEvidenceIds

`func (o *BiomarkerAssayWhitespaceRow) SetLiteratureEvidenceIds(v []string)`

SetLiteratureEvidenceIds sets LiteratureEvidenceIds field to given value.

### HasLiteratureEvidenceIds

`func (o *BiomarkerAssayWhitespaceRow) HasLiteratureEvidenceIds() bool`

HasLiteratureEvidenceIds returns a boolean if a field has been set.

### SetLiteratureEvidenceIdsNil

`func (o *BiomarkerAssayWhitespaceRow) SetLiteratureEvidenceIdsNil(b bool)`

 SetLiteratureEvidenceIdsNil sets the value for LiteratureEvidenceIds to be an explicit nil

### UnsetLiteratureEvidenceIds
`func (o *BiomarkerAssayWhitespaceRow) UnsetLiteratureEvidenceIds()`

UnsetLiteratureEvidenceIds ensures that no value is present for LiteratureEvidenceIds, not even an explicit nil
### GetLiteratureEvidenceDocumentIds

`func (o *BiomarkerAssayWhitespaceRow) GetLiteratureEvidenceDocumentIds() []string`

GetLiteratureEvidenceDocumentIds returns the LiteratureEvidenceDocumentIds field if non-nil, zero value otherwise.

### GetLiteratureEvidenceDocumentIdsOk

`func (o *BiomarkerAssayWhitespaceRow) GetLiteratureEvidenceDocumentIdsOk() (*[]string, bool)`

GetLiteratureEvidenceDocumentIdsOk returns a tuple with the LiteratureEvidenceDocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiteratureEvidenceDocumentIds

`func (o *BiomarkerAssayWhitespaceRow) SetLiteratureEvidenceDocumentIds(v []string)`

SetLiteratureEvidenceDocumentIds sets LiteratureEvidenceDocumentIds field to given value.

### HasLiteratureEvidenceDocumentIds

`func (o *BiomarkerAssayWhitespaceRow) HasLiteratureEvidenceDocumentIds() bool`

HasLiteratureEvidenceDocumentIds returns a boolean if a field has been set.

### SetLiteratureEvidenceDocumentIdsNil

`func (o *BiomarkerAssayWhitespaceRow) SetLiteratureEvidenceDocumentIdsNil(b bool)`

 SetLiteratureEvidenceDocumentIdsNil sets the value for LiteratureEvidenceDocumentIds to be an explicit nil

### UnsetLiteratureEvidenceDocumentIds
`func (o *BiomarkerAssayWhitespaceRow) UnsetLiteratureEvidenceDocumentIds()`

UnsetLiteratureEvidenceDocumentIds ensures that no value is present for LiteratureEvidenceDocumentIds, not even an explicit nil
### GetSupplierDeclarationIds

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierDeclarationIds() []string`

GetSupplierDeclarationIds returns the SupplierDeclarationIds field if non-nil, zero value otherwise.

### GetSupplierDeclarationIdsOk

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierDeclarationIdsOk() (*[]string, bool)`

GetSupplierDeclarationIdsOk returns a tuple with the SupplierDeclarationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierDeclarationIds

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierDeclarationIds(v []string)`

SetSupplierDeclarationIds sets SupplierDeclarationIds field to given value.

### HasSupplierDeclarationIds

`func (o *BiomarkerAssayWhitespaceRow) HasSupplierDeclarationIds() bool`

HasSupplierDeclarationIds returns a boolean if a field has been set.

### SetSupplierDeclarationIdsNil

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierDeclarationIdsNil(b bool)`

 SetSupplierDeclarationIdsNil sets the value for SupplierDeclarationIds to be an explicit nil

### UnsetSupplierDeclarationIds
`func (o *BiomarkerAssayWhitespaceRow) UnsetSupplierDeclarationIds()`

UnsetSupplierDeclarationIds ensures that no value is present for SupplierDeclarationIds, not even an explicit nil
### GetSupplierPositiveResponseCount

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierPositiveResponseCount() int32`

GetSupplierPositiveResponseCount returns the SupplierPositiveResponseCount field if non-nil, zero value otherwise.

### GetSupplierPositiveResponseCountOk

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierPositiveResponseCountOk() (*int32, bool)`

GetSupplierPositiveResponseCountOk returns a tuple with the SupplierPositiveResponseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierPositiveResponseCount

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierPositiveResponseCount(v int32)`

SetSupplierPositiveResponseCount sets SupplierPositiveResponseCount field to given value.

### HasSupplierPositiveResponseCount

`func (o *BiomarkerAssayWhitespaceRow) HasSupplierPositiveResponseCount() bool`

HasSupplierPositiveResponseCount returns a boolean if a field has been set.

### SetSupplierPositiveResponseCountNil

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierPositiveResponseCountNil(b bool)`

 SetSupplierPositiveResponseCountNil sets the value for SupplierPositiveResponseCount to be an explicit nil

### UnsetSupplierPositiveResponseCount
`func (o *BiomarkerAssayWhitespaceRow) UnsetSupplierPositiveResponseCount()`

UnsetSupplierPositiveResponseCount ensures that no value is present for SupplierPositiveResponseCount, not even an explicit nil
### GetSupplierConflictingResponseCount

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierConflictingResponseCount() int32`

GetSupplierConflictingResponseCount returns the SupplierConflictingResponseCount field if non-nil, zero value otherwise.

### GetSupplierConflictingResponseCountOk

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierConflictingResponseCountOk() (*int32, bool)`

GetSupplierConflictingResponseCountOk returns a tuple with the SupplierConflictingResponseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierConflictingResponseCount

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierConflictingResponseCount(v int32)`

SetSupplierConflictingResponseCount sets SupplierConflictingResponseCount field to given value.

### HasSupplierConflictingResponseCount

`func (o *BiomarkerAssayWhitespaceRow) HasSupplierConflictingResponseCount() bool`

HasSupplierConflictingResponseCount returns a boolean if a field has been set.

### SetSupplierConflictingResponseCountNil

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierConflictingResponseCountNil(b bool)`

 SetSupplierConflictingResponseCountNil sets the value for SupplierConflictingResponseCount to be an explicit nil

### UnsetSupplierConflictingResponseCount
`func (o *BiomarkerAssayWhitespaceRow) UnsetSupplierConflictingResponseCount()`

UnsetSupplierConflictingResponseCount ensures that no value is present for SupplierConflictingResponseCount, not even an explicit nil
### GetSupplierDetailOnlyResponseCount

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierDetailOnlyResponseCount() int32`

GetSupplierDetailOnlyResponseCount returns the SupplierDetailOnlyResponseCount field if non-nil, zero value otherwise.

### GetSupplierDetailOnlyResponseCountOk

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierDetailOnlyResponseCountOk() (*int32, bool)`

GetSupplierDetailOnlyResponseCountOk returns a tuple with the SupplierDetailOnlyResponseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierDetailOnlyResponseCount

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierDetailOnlyResponseCount(v int32)`

SetSupplierDetailOnlyResponseCount sets SupplierDetailOnlyResponseCount field to given value.

### HasSupplierDetailOnlyResponseCount

`func (o *BiomarkerAssayWhitespaceRow) HasSupplierDetailOnlyResponseCount() bool`

HasSupplierDetailOnlyResponseCount returns a boolean if a field has been set.

### SetSupplierDetailOnlyResponseCountNil

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierDetailOnlyResponseCountNil(b bool)`

 SetSupplierDetailOnlyResponseCountNil sets the value for SupplierDetailOnlyResponseCount to be an explicit nil

### UnsetSupplierDetailOnlyResponseCount
`func (o *BiomarkerAssayWhitespaceRow) UnsetSupplierDetailOnlyResponseCount()`

UnsetSupplierDetailOnlyResponseCount ensures that no value is present for SupplierDetailOnlyResponseCount, not even an explicit nil
### GetSupplierNoAssayResponseCount

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierNoAssayResponseCount() int32`

GetSupplierNoAssayResponseCount returns the SupplierNoAssayResponseCount field if non-nil, zero value otherwise.

### GetSupplierNoAssayResponseCountOk

`func (o *BiomarkerAssayWhitespaceRow) GetSupplierNoAssayResponseCountOk() (*int32, bool)`

GetSupplierNoAssayResponseCountOk returns a tuple with the SupplierNoAssayResponseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierNoAssayResponseCount

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierNoAssayResponseCount(v int32)`

SetSupplierNoAssayResponseCount sets SupplierNoAssayResponseCount field to given value.

### HasSupplierNoAssayResponseCount

`func (o *BiomarkerAssayWhitespaceRow) HasSupplierNoAssayResponseCount() bool`

HasSupplierNoAssayResponseCount returns a boolean if a field has been set.

### SetSupplierNoAssayResponseCountNil

`func (o *BiomarkerAssayWhitespaceRow) SetSupplierNoAssayResponseCountNil(b bool)`

 SetSupplierNoAssayResponseCountNil sets the value for SupplierNoAssayResponseCount to be an explicit nil

### UnsetSupplierNoAssayResponseCount
`func (o *BiomarkerAssayWhitespaceRow) UnsetSupplierNoAssayResponseCount()`

UnsetSupplierNoAssayResponseCount ensures that no value is present for SupplierNoAssayResponseCount, not even an explicit nil
### GetPipelineProgramCount

`func (o *BiomarkerAssayWhitespaceRow) GetPipelineProgramCount() int32`

GetPipelineProgramCount returns the PipelineProgramCount field if non-nil, zero value otherwise.

### GetPipelineProgramCountOk

`func (o *BiomarkerAssayWhitespaceRow) GetPipelineProgramCountOk() (*int32, bool)`

GetPipelineProgramCountOk returns a tuple with the PipelineProgramCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineProgramCount

`func (o *BiomarkerAssayWhitespaceRow) SetPipelineProgramCount(v int32)`

SetPipelineProgramCount sets PipelineProgramCount field to given value.

### HasPipelineProgramCount

`func (o *BiomarkerAssayWhitespaceRow) HasPipelineProgramCount() bool`

HasPipelineProgramCount returns a boolean if a field has been set.

### SetPipelineProgramCountNil

`func (o *BiomarkerAssayWhitespaceRow) SetPipelineProgramCountNil(b bool)`

 SetPipelineProgramCountNil sets the value for PipelineProgramCount to be an explicit nil

### UnsetPipelineProgramCount
`func (o *BiomarkerAssayWhitespaceRow) UnsetPipelineProgramCount()`

UnsetPipelineProgramCount ensures that no value is present for PipelineProgramCount, not even an explicit nil
### GetPipelineActiveProgramCount

`func (o *BiomarkerAssayWhitespaceRow) GetPipelineActiveProgramCount() int32`

GetPipelineActiveProgramCount returns the PipelineActiveProgramCount field if non-nil, zero value otherwise.

### GetPipelineActiveProgramCountOk

`func (o *BiomarkerAssayWhitespaceRow) GetPipelineActiveProgramCountOk() (*int32, bool)`

GetPipelineActiveProgramCountOk returns a tuple with the PipelineActiveProgramCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineActiveProgramCount

`func (o *BiomarkerAssayWhitespaceRow) SetPipelineActiveProgramCount(v int32)`

SetPipelineActiveProgramCount sets PipelineActiveProgramCount field to given value.

### HasPipelineActiveProgramCount

`func (o *BiomarkerAssayWhitespaceRow) HasPipelineActiveProgramCount() bool`

HasPipelineActiveProgramCount returns a boolean if a field has been set.

### SetPipelineActiveProgramCountNil

`func (o *BiomarkerAssayWhitespaceRow) SetPipelineActiveProgramCountNil(b bool)`

 SetPipelineActiveProgramCountNil sets the value for PipelineActiveProgramCount to be an explicit nil

### UnsetPipelineActiveProgramCount
`func (o *BiomarkerAssayWhitespaceRow) UnsetPipelineActiveProgramCount()`

UnsetPipelineActiveProgramCount ensures that no value is present for PipelineActiveProgramCount, not even an explicit nil
### GetPipelineMaxPhase

`func (o *BiomarkerAssayWhitespaceRow) GetPipelineMaxPhase() int32`

GetPipelineMaxPhase returns the PipelineMaxPhase field if non-nil, zero value otherwise.

### GetPipelineMaxPhaseOk

`func (o *BiomarkerAssayWhitespaceRow) GetPipelineMaxPhaseOk() (*int32, bool)`

GetPipelineMaxPhaseOk returns a tuple with the PipelineMaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineMaxPhase

`func (o *BiomarkerAssayWhitespaceRow) SetPipelineMaxPhase(v int32)`

SetPipelineMaxPhase sets PipelineMaxPhase field to given value.

### HasPipelineMaxPhase

`func (o *BiomarkerAssayWhitespaceRow) HasPipelineMaxPhase() bool`

HasPipelineMaxPhase returns a boolean if a field has been set.

### SetPipelineMaxPhaseNil

`func (o *BiomarkerAssayWhitespaceRow) SetPipelineMaxPhaseNil(b bool)`

 SetPipelineMaxPhaseNil sets the value for PipelineMaxPhase to be an explicit nil

### UnsetPipelineMaxPhase
`func (o *BiomarkerAssayWhitespaceRow) UnsetPipelineMaxPhase()`

UnsetPipelineMaxPhase ensures that no value is present for PipelineMaxPhase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


