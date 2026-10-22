# RareDiseasesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**DiseaseKey** | Pointer to **NullableString** | The disease key (orpha: plus the ORPHA code). An attribute, not an id: the entry key is orpha_code. | [optional] 
**PreferredName** | Pointer to **NullableString** | The disease&#39;s preferred English name. | [optional] 
**Aliases** | Pointer to **[]string** | English synonyms Orphanet lists for the disease. | [optional] 
**ConceptType** | Pointer to **NullableString** | Whether the concept is a group of disorders, a disorder or a subtype. | [optional] 
**Status** | Pointer to **NullableString** | The concept&#39;s lifecycle state in Orphanet (active or deprecated). | [optional] 
**OrdoEntityType** | Pointer to **NullableString** | The entity type the Orphanet Rare Disease Ontology gives the concept. | [optional] 
**OrphanetDisorderType** | Pointer to **NullableString** | The disorder type in Orphanet&#39;s classification of rare diseases. | [optional] 
**OrphanetDisorderGroup** | Pointer to **NullableString** | The disorder group in Orphanet&#39;s classification of rare diseases. | [optional] 
**GeneEvidenceState** | Pointer to **NullableString** | known_causal when at least one gene assertion is a disease-causing mutation assessed by Orphanet, weaker_only when only other gene assertions exist, none when there is no gene assertion. | [optional] 
**NGenes** | Pointer to **NullableInt64** | Distinct genes asserted for the disease. | [optional] 
**NKnownCausalGenes** | Pointer to **NullableInt64** | Distinct genes whose assertion is a disease-causing mutation assessed by Orphanet. | [optional] 
**PrevalenceBand** | Pointer to **NullableString** | The Orphanet point-prevalence band chosen for the disease, one of six literals; null when no documented band exists. | [optional] 
**RarityStatusGb** | Pointer to **NullableString** | The disease&#39;s rarity status under the UK definition, from the latest assessment (rare, not_rare, unknown or conflicting); null when not assessed. | [optional] 
**RarityStatusEu** | Pointer to **NullableString** | The disease&#39;s rarity status under the European Union definition, from the latest assessment (rare, not_rare, unknown or conflicting); null when not assessed. | [optional] 
**TrialCoverageState** | Pointer to **NullableString** | no_bridge when the disease maps to no MeSH condition, bridged_no_trials when it maps to a condition with no trial, trials when at least one trial is linked. | [optional] 
**NTrialsExactDirect** | Pointer to **NullableInt64** | Trials linked through an exact mapping whose own condition list names the disease. Count within this column only. | [optional] 
**NTrialsExactBroader** | Pointer to **NullableInt64** | Trials linked through an exact mapping to a broader condition than the trial studies. Count within this column only. | [optional] 
**NTrialsUnspecified** | Pointer to **NullableInt64** | Trials linked through a MONDO cross-reference that states no relation. Count within this column only. | [optional] 
**NExactMeshConditions** | Pointer to **NullableInt64** | MeSH conditions the disease maps to with an exact relation. | [optional] 
**MergesIntoConditionId** | Pointer to **NullableString** | The single MeSH condition the web merges this disease into, set only when exactly one exact MeSH match exists. | [optional] 
**Icd10CodingState** | Pointer to **NullableString** | How Orphanet aligns the disease to ICD-10: exact, directional, unstated or unaligned. | [optional] 
**Icd11CodingState** | Pointer to **NullableString** | How Orphanet aligns the disease to ICD-11: exact, directional, unstated or unaligned. | [optional] 
**SourceRelease** | Pointer to **NullableString** | The Orphadata Science release the row was read from. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the Orphadata Science release. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseasesRow

`func NewRareDiseasesRow() *RareDiseasesRow`

NewRareDiseasesRow instantiates a new RareDiseasesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseasesRowWithDefaults

`func NewRareDiseasesRowWithDefaults() *RareDiseasesRow`

NewRareDiseasesRowWithDefaults instantiates a new RareDiseasesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrphaCode

`func (o *RareDiseasesRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseasesRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseasesRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseasesRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseasesRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseasesRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetDiseaseKey

`func (o *RareDiseasesRow) GetDiseaseKey() string`

GetDiseaseKey returns the DiseaseKey field if non-nil, zero value otherwise.

### GetDiseaseKeyOk

`func (o *RareDiseasesRow) GetDiseaseKeyOk() (*string, bool)`

GetDiseaseKeyOk returns a tuple with the DiseaseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiseaseKey

`func (o *RareDiseasesRow) SetDiseaseKey(v string)`

SetDiseaseKey sets DiseaseKey field to given value.

### HasDiseaseKey

`func (o *RareDiseasesRow) HasDiseaseKey() bool`

HasDiseaseKey returns a boolean if a field has been set.

### SetDiseaseKeyNil

`func (o *RareDiseasesRow) SetDiseaseKeyNil(b bool)`

 SetDiseaseKeyNil sets the value for DiseaseKey to be an explicit nil

### UnsetDiseaseKey
`func (o *RareDiseasesRow) UnsetDiseaseKey()`

UnsetDiseaseKey ensures that no value is present for DiseaseKey, not even an explicit nil
### GetPreferredName

`func (o *RareDiseasesRow) GetPreferredName() string`

GetPreferredName returns the PreferredName field if non-nil, zero value otherwise.

### GetPreferredNameOk

`func (o *RareDiseasesRow) GetPreferredNameOk() (*string, bool)`

GetPreferredNameOk returns a tuple with the PreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredName

`func (o *RareDiseasesRow) SetPreferredName(v string)`

SetPreferredName sets PreferredName field to given value.

### HasPreferredName

`func (o *RareDiseasesRow) HasPreferredName() bool`

HasPreferredName returns a boolean if a field has been set.

### SetPreferredNameNil

`func (o *RareDiseasesRow) SetPreferredNameNil(b bool)`

 SetPreferredNameNil sets the value for PreferredName to be an explicit nil

### UnsetPreferredName
`func (o *RareDiseasesRow) UnsetPreferredName()`

UnsetPreferredName ensures that no value is present for PreferredName, not even an explicit nil
### GetAliases

`func (o *RareDiseasesRow) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *RareDiseasesRow) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *RareDiseasesRow) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *RareDiseasesRow) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *RareDiseasesRow) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *RareDiseasesRow) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil
### GetConceptType

`func (o *RareDiseasesRow) GetConceptType() string`

GetConceptType returns the ConceptType field if non-nil, zero value otherwise.

### GetConceptTypeOk

`func (o *RareDiseasesRow) GetConceptTypeOk() (*string, bool)`

GetConceptTypeOk returns a tuple with the ConceptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConceptType

`func (o *RareDiseasesRow) SetConceptType(v string)`

SetConceptType sets ConceptType field to given value.

### HasConceptType

`func (o *RareDiseasesRow) HasConceptType() bool`

HasConceptType returns a boolean if a field has been set.

### SetConceptTypeNil

`func (o *RareDiseasesRow) SetConceptTypeNil(b bool)`

 SetConceptTypeNil sets the value for ConceptType to be an explicit nil

### UnsetConceptType
`func (o *RareDiseasesRow) UnsetConceptType()`

UnsetConceptType ensures that no value is present for ConceptType, not even an explicit nil
### GetStatus

`func (o *RareDiseasesRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RareDiseasesRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RareDiseasesRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RareDiseasesRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RareDiseasesRow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RareDiseasesRow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetOrdoEntityType

`func (o *RareDiseasesRow) GetOrdoEntityType() string`

GetOrdoEntityType returns the OrdoEntityType field if non-nil, zero value otherwise.

### GetOrdoEntityTypeOk

`func (o *RareDiseasesRow) GetOrdoEntityTypeOk() (*string, bool)`

GetOrdoEntityTypeOk returns a tuple with the OrdoEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdoEntityType

`func (o *RareDiseasesRow) SetOrdoEntityType(v string)`

SetOrdoEntityType sets OrdoEntityType field to given value.

### HasOrdoEntityType

`func (o *RareDiseasesRow) HasOrdoEntityType() bool`

HasOrdoEntityType returns a boolean if a field has been set.

### SetOrdoEntityTypeNil

`func (o *RareDiseasesRow) SetOrdoEntityTypeNil(b bool)`

 SetOrdoEntityTypeNil sets the value for OrdoEntityType to be an explicit nil

### UnsetOrdoEntityType
`func (o *RareDiseasesRow) UnsetOrdoEntityType()`

UnsetOrdoEntityType ensures that no value is present for OrdoEntityType, not even an explicit nil
### GetOrphanetDisorderType

`func (o *RareDiseasesRow) GetOrphanetDisorderType() string`

GetOrphanetDisorderType returns the OrphanetDisorderType field if non-nil, zero value otherwise.

### GetOrphanetDisorderTypeOk

`func (o *RareDiseasesRow) GetOrphanetDisorderTypeOk() (*string, bool)`

GetOrphanetDisorderTypeOk returns a tuple with the OrphanetDisorderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanetDisorderType

`func (o *RareDiseasesRow) SetOrphanetDisorderType(v string)`

SetOrphanetDisorderType sets OrphanetDisorderType field to given value.

### HasOrphanetDisorderType

`func (o *RareDiseasesRow) HasOrphanetDisorderType() bool`

HasOrphanetDisorderType returns a boolean if a field has been set.

### SetOrphanetDisorderTypeNil

`func (o *RareDiseasesRow) SetOrphanetDisorderTypeNil(b bool)`

 SetOrphanetDisorderTypeNil sets the value for OrphanetDisorderType to be an explicit nil

### UnsetOrphanetDisorderType
`func (o *RareDiseasesRow) UnsetOrphanetDisorderType()`

UnsetOrphanetDisorderType ensures that no value is present for OrphanetDisorderType, not even an explicit nil
### GetOrphanetDisorderGroup

`func (o *RareDiseasesRow) GetOrphanetDisorderGroup() string`

GetOrphanetDisorderGroup returns the OrphanetDisorderGroup field if non-nil, zero value otherwise.

### GetOrphanetDisorderGroupOk

`func (o *RareDiseasesRow) GetOrphanetDisorderGroupOk() (*string, bool)`

GetOrphanetDisorderGroupOk returns a tuple with the OrphanetDisorderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanetDisorderGroup

`func (o *RareDiseasesRow) SetOrphanetDisorderGroup(v string)`

SetOrphanetDisorderGroup sets OrphanetDisorderGroup field to given value.

### HasOrphanetDisorderGroup

`func (o *RareDiseasesRow) HasOrphanetDisorderGroup() bool`

HasOrphanetDisorderGroup returns a boolean if a field has been set.

### SetOrphanetDisorderGroupNil

`func (o *RareDiseasesRow) SetOrphanetDisorderGroupNil(b bool)`

 SetOrphanetDisorderGroupNil sets the value for OrphanetDisorderGroup to be an explicit nil

### UnsetOrphanetDisorderGroup
`func (o *RareDiseasesRow) UnsetOrphanetDisorderGroup()`

UnsetOrphanetDisorderGroup ensures that no value is present for OrphanetDisorderGroup, not even an explicit nil
### GetGeneEvidenceState

`func (o *RareDiseasesRow) GetGeneEvidenceState() string`

GetGeneEvidenceState returns the GeneEvidenceState field if non-nil, zero value otherwise.

### GetGeneEvidenceStateOk

`func (o *RareDiseasesRow) GetGeneEvidenceStateOk() (*string, bool)`

GetGeneEvidenceStateOk returns a tuple with the GeneEvidenceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneEvidenceState

`func (o *RareDiseasesRow) SetGeneEvidenceState(v string)`

SetGeneEvidenceState sets GeneEvidenceState field to given value.

### HasGeneEvidenceState

`func (o *RareDiseasesRow) HasGeneEvidenceState() bool`

HasGeneEvidenceState returns a boolean if a field has been set.

### SetGeneEvidenceStateNil

`func (o *RareDiseasesRow) SetGeneEvidenceStateNil(b bool)`

 SetGeneEvidenceStateNil sets the value for GeneEvidenceState to be an explicit nil

### UnsetGeneEvidenceState
`func (o *RareDiseasesRow) UnsetGeneEvidenceState()`

UnsetGeneEvidenceState ensures that no value is present for GeneEvidenceState, not even an explicit nil
### GetNGenes

`func (o *RareDiseasesRow) GetNGenes() int64`

GetNGenes returns the NGenes field if non-nil, zero value otherwise.

### GetNGenesOk

`func (o *RareDiseasesRow) GetNGenesOk() (*int64, bool)`

GetNGenesOk returns a tuple with the NGenes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNGenes

`func (o *RareDiseasesRow) SetNGenes(v int64)`

SetNGenes sets NGenes field to given value.

### HasNGenes

`func (o *RareDiseasesRow) HasNGenes() bool`

HasNGenes returns a boolean if a field has been set.

### SetNGenesNil

`func (o *RareDiseasesRow) SetNGenesNil(b bool)`

 SetNGenesNil sets the value for NGenes to be an explicit nil

### UnsetNGenes
`func (o *RareDiseasesRow) UnsetNGenes()`

UnsetNGenes ensures that no value is present for NGenes, not even an explicit nil
### GetNKnownCausalGenes

`func (o *RareDiseasesRow) GetNKnownCausalGenes() int64`

GetNKnownCausalGenes returns the NKnownCausalGenes field if non-nil, zero value otherwise.

### GetNKnownCausalGenesOk

`func (o *RareDiseasesRow) GetNKnownCausalGenesOk() (*int64, bool)`

GetNKnownCausalGenesOk returns a tuple with the NKnownCausalGenes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNKnownCausalGenes

`func (o *RareDiseasesRow) SetNKnownCausalGenes(v int64)`

SetNKnownCausalGenes sets NKnownCausalGenes field to given value.

### HasNKnownCausalGenes

`func (o *RareDiseasesRow) HasNKnownCausalGenes() bool`

HasNKnownCausalGenes returns a boolean if a field has been set.

### SetNKnownCausalGenesNil

`func (o *RareDiseasesRow) SetNKnownCausalGenesNil(b bool)`

 SetNKnownCausalGenesNil sets the value for NKnownCausalGenes to be an explicit nil

### UnsetNKnownCausalGenes
`func (o *RareDiseasesRow) UnsetNKnownCausalGenes()`

UnsetNKnownCausalGenes ensures that no value is present for NKnownCausalGenes, not even an explicit nil
### GetPrevalenceBand

`func (o *RareDiseasesRow) GetPrevalenceBand() string`

GetPrevalenceBand returns the PrevalenceBand field if non-nil, zero value otherwise.

### GetPrevalenceBandOk

`func (o *RareDiseasesRow) GetPrevalenceBandOk() (*string, bool)`

GetPrevalenceBandOk returns a tuple with the PrevalenceBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceBand

`func (o *RareDiseasesRow) SetPrevalenceBand(v string)`

SetPrevalenceBand sets PrevalenceBand field to given value.

### HasPrevalenceBand

`func (o *RareDiseasesRow) HasPrevalenceBand() bool`

HasPrevalenceBand returns a boolean if a field has been set.

### SetPrevalenceBandNil

`func (o *RareDiseasesRow) SetPrevalenceBandNil(b bool)`

 SetPrevalenceBandNil sets the value for PrevalenceBand to be an explicit nil

### UnsetPrevalenceBand
`func (o *RareDiseasesRow) UnsetPrevalenceBand()`

UnsetPrevalenceBand ensures that no value is present for PrevalenceBand, not even an explicit nil
### GetRarityStatusGb

`func (o *RareDiseasesRow) GetRarityStatusGb() string`

GetRarityStatusGb returns the RarityStatusGb field if non-nil, zero value otherwise.

### GetRarityStatusGbOk

`func (o *RareDiseasesRow) GetRarityStatusGbOk() (*string, bool)`

GetRarityStatusGbOk returns a tuple with the RarityStatusGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarityStatusGb

`func (o *RareDiseasesRow) SetRarityStatusGb(v string)`

SetRarityStatusGb sets RarityStatusGb field to given value.

### HasRarityStatusGb

`func (o *RareDiseasesRow) HasRarityStatusGb() bool`

HasRarityStatusGb returns a boolean if a field has been set.

### SetRarityStatusGbNil

`func (o *RareDiseasesRow) SetRarityStatusGbNil(b bool)`

 SetRarityStatusGbNil sets the value for RarityStatusGb to be an explicit nil

### UnsetRarityStatusGb
`func (o *RareDiseasesRow) UnsetRarityStatusGb()`

UnsetRarityStatusGb ensures that no value is present for RarityStatusGb, not even an explicit nil
### GetRarityStatusEu

`func (o *RareDiseasesRow) GetRarityStatusEu() string`

GetRarityStatusEu returns the RarityStatusEu field if non-nil, zero value otherwise.

### GetRarityStatusEuOk

`func (o *RareDiseasesRow) GetRarityStatusEuOk() (*string, bool)`

GetRarityStatusEuOk returns a tuple with the RarityStatusEu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarityStatusEu

`func (o *RareDiseasesRow) SetRarityStatusEu(v string)`

SetRarityStatusEu sets RarityStatusEu field to given value.

### HasRarityStatusEu

`func (o *RareDiseasesRow) HasRarityStatusEu() bool`

HasRarityStatusEu returns a boolean if a field has been set.

### SetRarityStatusEuNil

`func (o *RareDiseasesRow) SetRarityStatusEuNil(b bool)`

 SetRarityStatusEuNil sets the value for RarityStatusEu to be an explicit nil

### UnsetRarityStatusEu
`func (o *RareDiseasesRow) UnsetRarityStatusEu()`

UnsetRarityStatusEu ensures that no value is present for RarityStatusEu, not even an explicit nil
### GetTrialCoverageState

`func (o *RareDiseasesRow) GetTrialCoverageState() string`

GetTrialCoverageState returns the TrialCoverageState field if non-nil, zero value otherwise.

### GetTrialCoverageStateOk

`func (o *RareDiseasesRow) GetTrialCoverageStateOk() (*string, bool)`

GetTrialCoverageStateOk returns a tuple with the TrialCoverageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialCoverageState

`func (o *RareDiseasesRow) SetTrialCoverageState(v string)`

SetTrialCoverageState sets TrialCoverageState field to given value.

### HasTrialCoverageState

`func (o *RareDiseasesRow) HasTrialCoverageState() bool`

HasTrialCoverageState returns a boolean if a field has been set.

### SetTrialCoverageStateNil

`func (o *RareDiseasesRow) SetTrialCoverageStateNil(b bool)`

 SetTrialCoverageStateNil sets the value for TrialCoverageState to be an explicit nil

### UnsetTrialCoverageState
`func (o *RareDiseasesRow) UnsetTrialCoverageState()`

UnsetTrialCoverageState ensures that no value is present for TrialCoverageState, not even an explicit nil
### GetNTrialsExactDirect

`func (o *RareDiseasesRow) GetNTrialsExactDirect() int64`

GetNTrialsExactDirect returns the NTrialsExactDirect field if non-nil, zero value otherwise.

### GetNTrialsExactDirectOk

`func (o *RareDiseasesRow) GetNTrialsExactDirectOk() (*int64, bool)`

GetNTrialsExactDirectOk returns a tuple with the NTrialsExactDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsExactDirect

`func (o *RareDiseasesRow) SetNTrialsExactDirect(v int64)`

SetNTrialsExactDirect sets NTrialsExactDirect field to given value.

### HasNTrialsExactDirect

`func (o *RareDiseasesRow) HasNTrialsExactDirect() bool`

HasNTrialsExactDirect returns a boolean if a field has been set.

### SetNTrialsExactDirectNil

`func (o *RareDiseasesRow) SetNTrialsExactDirectNil(b bool)`

 SetNTrialsExactDirectNil sets the value for NTrialsExactDirect to be an explicit nil

### UnsetNTrialsExactDirect
`func (o *RareDiseasesRow) UnsetNTrialsExactDirect()`

UnsetNTrialsExactDirect ensures that no value is present for NTrialsExactDirect, not even an explicit nil
### GetNTrialsExactBroader

`func (o *RareDiseasesRow) GetNTrialsExactBroader() int64`

GetNTrialsExactBroader returns the NTrialsExactBroader field if non-nil, zero value otherwise.

### GetNTrialsExactBroaderOk

`func (o *RareDiseasesRow) GetNTrialsExactBroaderOk() (*int64, bool)`

GetNTrialsExactBroaderOk returns a tuple with the NTrialsExactBroader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsExactBroader

`func (o *RareDiseasesRow) SetNTrialsExactBroader(v int64)`

SetNTrialsExactBroader sets NTrialsExactBroader field to given value.

### HasNTrialsExactBroader

`func (o *RareDiseasesRow) HasNTrialsExactBroader() bool`

HasNTrialsExactBroader returns a boolean if a field has been set.

### SetNTrialsExactBroaderNil

`func (o *RareDiseasesRow) SetNTrialsExactBroaderNil(b bool)`

 SetNTrialsExactBroaderNil sets the value for NTrialsExactBroader to be an explicit nil

### UnsetNTrialsExactBroader
`func (o *RareDiseasesRow) UnsetNTrialsExactBroader()`

UnsetNTrialsExactBroader ensures that no value is present for NTrialsExactBroader, not even an explicit nil
### GetNTrialsUnspecified

`func (o *RareDiseasesRow) GetNTrialsUnspecified() int64`

GetNTrialsUnspecified returns the NTrialsUnspecified field if non-nil, zero value otherwise.

### GetNTrialsUnspecifiedOk

`func (o *RareDiseasesRow) GetNTrialsUnspecifiedOk() (*int64, bool)`

GetNTrialsUnspecifiedOk returns a tuple with the NTrialsUnspecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsUnspecified

`func (o *RareDiseasesRow) SetNTrialsUnspecified(v int64)`

SetNTrialsUnspecified sets NTrialsUnspecified field to given value.

### HasNTrialsUnspecified

`func (o *RareDiseasesRow) HasNTrialsUnspecified() bool`

HasNTrialsUnspecified returns a boolean if a field has been set.

### SetNTrialsUnspecifiedNil

`func (o *RareDiseasesRow) SetNTrialsUnspecifiedNil(b bool)`

 SetNTrialsUnspecifiedNil sets the value for NTrialsUnspecified to be an explicit nil

### UnsetNTrialsUnspecified
`func (o *RareDiseasesRow) UnsetNTrialsUnspecified()`

UnsetNTrialsUnspecified ensures that no value is present for NTrialsUnspecified, not even an explicit nil
### GetNExactMeshConditions

`func (o *RareDiseasesRow) GetNExactMeshConditions() int64`

GetNExactMeshConditions returns the NExactMeshConditions field if non-nil, zero value otherwise.

### GetNExactMeshConditionsOk

`func (o *RareDiseasesRow) GetNExactMeshConditionsOk() (*int64, bool)`

GetNExactMeshConditionsOk returns a tuple with the NExactMeshConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNExactMeshConditions

`func (o *RareDiseasesRow) SetNExactMeshConditions(v int64)`

SetNExactMeshConditions sets NExactMeshConditions field to given value.

### HasNExactMeshConditions

`func (o *RareDiseasesRow) HasNExactMeshConditions() bool`

HasNExactMeshConditions returns a boolean if a field has been set.

### SetNExactMeshConditionsNil

`func (o *RareDiseasesRow) SetNExactMeshConditionsNil(b bool)`

 SetNExactMeshConditionsNil sets the value for NExactMeshConditions to be an explicit nil

### UnsetNExactMeshConditions
`func (o *RareDiseasesRow) UnsetNExactMeshConditions()`

UnsetNExactMeshConditions ensures that no value is present for NExactMeshConditions, not even an explicit nil
### GetMergesIntoConditionId

`func (o *RareDiseasesRow) GetMergesIntoConditionId() string`

GetMergesIntoConditionId returns the MergesIntoConditionId field if non-nil, zero value otherwise.

### GetMergesIntoConditionIdOk

`func (o *RareDiseasesRow) GetMergesIntoConditionIdOk() (*string, bool)`

GetMergesIntoConditionIdOk returns a tuple with the MergesIntoConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergesIntoConditionId

`func (o *RareDiseasesRow) SetMergesIntoConditionId(v string)`

SetMergesIntoConditionId sets MergesIntoConditionId field to given value.

### HasMergesIntoConditionId

`func (o *RareDiseasesRow) HasMergesIntoConditionId() bool`

HasMergesIntoConditionId returns a boolean if a field has been set.

### SetMergesIntoConditionIdNil

`func (o *RareDiseasesRow) SetMergesIntoConditionIdNil(b bool)`

 SetMergesIntoConditionIdNil sets the value for MergesIntoConditionId to be an explicit nil

### UnsetMergesIntoConditionId
`func (o *RareDiseasesRow) UnsetMergesIntoConditionId()`

UnsetMergesIntoConditionId ensures that no value is present for MergesIntoConditionId, not even an explicit nil
### GetIcd10CodingState

`func (o *RareDiseasesRow) GetIcd10CodingState() string`

GetIcd10CodingState returns the Icd10CodingState field if non-nil, zero value otherwise.

### GetIcd10CodingStateOk

`func (o *RareDiseasesRow) GetIcd10CodingStateOk() (*string, bool)`

GetIcd10CodingStateOk returns a tuple with the Icd10CodingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcd10CodingState

`func (o *RareDiseasesRow) SetIcd10CodingState(v string)`

SetIcd10CodingState sets Icd10CodingState field to given value.

### HasIcd10CodingState

`func (o *RareDiseasesRow) HasIcd10CodingState() bool`

HasIcd10CodingState returns a boolean if a field has been set.

### SetIcd10CodingStateNil

`func (o *RareDiseasesRow) SetIcd10CodingStateNil(b bool)`

 SetIcd10CodingStateNil sets the value for Icd10CodingState to be an explicit nil

### UnsetIcd10CodingState
`func (o *RareDiseasesRow) UnsetIcd10CodingState()`

UnsetIcd10CodingState ensures that no value is present for Icd10CodingState, not even an explicit nil
### GetIcd11CodingState

`func (o *RareDiseasesRow) GetIcd11CodingState() string`

GetIcd11CodingState returns the Icd11CodingState field if non-nil, zero value otherwise.

### GetIcd11CodingStateOk

`func (o *RareDiseasesRow) GetIcd11CodingStateOk() (*string, bool)`

GetIcd11CodingStateOk returns a tuple with the Icd11CodingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcd11CodingState

`func (o *RareDiseasesRow) SetIcd11CodingState(v string)`

SetIcd11CodingState sets Icd11CodingState field to given value.

### HasIcd11CodingState

`func (o *RareDiseasesRow) HasIcd11CodingState() bool`

HasIcd11CodingState returns a boolean if a field has been set.

### SetIcd11CodingStateNil

`func (o *RareDiseasesRow) SetIcd11CodingStateNil(b bool)`

 SetIcd11CodingStateNil sets the value for Icd11CodingState to be an explicit nil

### UnsetIcd11CodingState
`func (o *RareDiseasesRow) UnsetIcd11CodingState()`

UnsetIcd11CodingState ensures that no value is present for Icd11CodingState, not even an explicit nil
### GetSourceRelease

`func (o *RareDiseasesRow) GetSourceRelease() string`

GetSourceRelease returns the SourceRelease field if non-nil, zero value otherwise.

### GetSourceReleaseOk

`func (o *RareDiseasesRow) GetSourceReleaseOk() (*string, bool)`

GetSourceReleaseOk returns a tuple with the SourceRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRelease

`func (o *RareDiseasesRow) SetSourceRelease(v string)`

SetSourceRelease sets SourceRelease field to given value.

### HasSourceRelease

`func (o *RareDiseasesRow) HasSourceRelease() bool`

HasSourceRelease returns a boolean if a field has been set.

### SetSourceReleaseNil

`func (o *RareDiseasesRow) SetSourceReleaseNil(b bool)`

 SetSourceReleaseNil sets the value for SourceRelease to be an explicit nil

### UnsetSourceRelease
`func (o *RareDiseasesRow) UnsetSourceRelease()`

UnsetSourceRelease ensures that no value is present for SourceRelease, not even an explicit nil
### GetLicenseName

`func (o *RareDiseasesRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseasesRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseasesRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseasesRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseasesRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseasesRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseasesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseasesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseasesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseasesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseasesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseasesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


