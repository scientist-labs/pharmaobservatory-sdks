# TargetEvidenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEvidenceId** | Pointer to **NullableString** | The row id: a hash of the target, drug, indication, trial and sponsor organization. | [optional] 
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] 
**DrugId** | Pointer to **NullableString** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**ApprovedAnywhere** | Pointer to **NullableBool** | Whether the drug is approved by any authority we hold. | [optional] 
**PrimaryActionType** | Pointer to **NullableString** | The action type of the drug&#39;s best-evidence mechanism on the target; null where the source gives none. | [optional] 
**EvidenceSource** | Pointer to **NullableString** | The best evidence tier that links the drug to the target: curated_moa (an Open Targets curated mechanism edge), chembl_moa (a ChEMBL mechanism) or intervention_seed (a curated investigational asset). | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition as a MeSH descriptor id (D followed by digits), the id the conditions datasets use. | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s name. | [optional] 
**TrialId** | Pointer to **NullableString** | The ClinicalTrials.gov NCT id, the id the trials datasets use. | [optional] 
**MaxPhase** | Pointer to **NullableInt32** | The furthest phase the program (drug and indication) reached, 0 to 4 (0 is phase not reported). | [optional] 
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the trial&#39;s lead sponsor organization as filed; null when it does not resolve. | [optional] 
**SponsorName** | Pointer to **NullableString** | The lead sponsor&#39;s name as filed. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the lead sponsor organization belongs to; null when none. | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 

## Methods

### NewTargetEvidenceRow

`func NewTargetEvidenceRow() *TargetEvidenceRow`

NewTargetEvidenceRow instantiates a new TargetEvidenceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetEvidenceRowWithDefaults

`func NewTargetEvidenceRowWithDefaults() *TargetEvidenceRow`

NewTargetEvidenceRowWithDefaults instantiates a new TargetEvidenceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEvidenceId

`func (o *TargetEvidenceRow) GetTargetEvidenceId() string`

GetTargetEvidenceId returns the TargetEvidenceId field if non-nil, zero value otherwise.

### GetTargetEvidenceIdOk

`func (o *TargetEvidenceRow) GetTargetEvidenceIdOk() (*string, bool)`

GetTargetEvidenceIdOk returns a tuple with the TargetEvidenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEvidenceId

`func (o *TargetEvidenceRow) SetTargetEvidenceId(v string)`

SetTargetEvidenceId sets TargetEvidenceId field to given value.

### HasTargetEvidenceId

`func (o *TargetEvidenceRow) HasTargetEvidenceId() bool`

HasTargetEvidenceId returns a boolean if a field has been set.

### SetTargetEvidenceIdNil

`func (o *TargetEvidenceRow) SetTargetEvidenceIdNil(b bool)`

 SetTargetEvidenceIdNil sets the value for TargetEvidenceId to be an explicit nil

### UnsetTargetEvidenceId
`func (o *TargetEvidenceRow) UnsetTargetEvidenceId()`

UnsetTargetEvidenceId ensures that no value is present for TargetEvidenceId, not even an explicit nil
### GetTargetId

`func (o *TargetEvidenceRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetEvidenceRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetEvidenceRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetEvidenceRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetEvidenceRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetEvidenceRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *TargetEvidenceRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *TargetEvidenceRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *TargetEvidenceRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *TargetEvidenceRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *TargetEvidenceRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *TargetEvidenceRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetDrugId

`func (o *TargetEvidenceRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *TargetEvidenceRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *TargetEvidenceRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *TargetEvidenceRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *TargetEvidenceRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *TargetEvidenceRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *TargetEvidenceRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *TargetEvidenceRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *TargetEvidenceRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *TargetEvidenceRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *TargetEvidenceRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *TargetEvidenceRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetApprovedAnywhere

`func (o *TargetEvidenceRow) GetApprovedAnywhere() bool`

GetApprovedAnywhere returns the ApprovedAnywhere field if non-nil, zero value otherwise.

### GetApprovedAnywhereOk

`func (o *TargetEvidenceRow) GetApprovedAnywhereOk() (*bool, bool)`

GetApprovedAnywhereOk returns a tuple with the ApprovedAnywhere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAnywhere

`func (o *TargetEvidenceRow) SetApprovedAnywhere(v bool)`

SetApprovedAnywhere sets ApprovedAnywhere field to given value.

### HasApprovedAnywhere

`func (o *TargetEvidenceRow) HasApprovedAnywhere() bool`

HasApprovedAnywhere returns a boolean if a field has been set.

### SetApprovedAnywhereNil

`func (o *TargetEvidenceRow) SetApprovedAnywhereNil(b bool)`

 SetApprovedAnywhereNil sets the value for ApprovedAnywhere to be an explicit nil

### UnsetApprovedAnywhere
`func (o *TargetEvidenceRow) UnsetApprovedAnywhere()`

UnsetApprovedAnywhere ensures that no value is present for ApprovedAnywhere, not even an explicit nil
### GetPrimaryActionType

`func (o *TargetEvidenceRow) GetPrimaryActionType() string`

GetPrimaryActionType returns the PrimaryActionType field if non-nil, zero value otherwise.

### GetPrimaryActionTypeOk

`func (o *TargetEvidenceRow) GetPrimaryActionTypeOk() (*string, bool)`

GetPrimaryActionTypeOk returns a tuple with the PrimaryActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryActionType

`func (o *TargetEvidenceRow) SetPrimaryActionType(v string)`

SetPrimaryActionType sets PrimaryActionType field to given value.

### HasPrimaryActionType

`func (o *TargetEvidenceRow) HasPrimaryActionType() bool`

HasPrimaryActionType returns a boolean if a field has been set.

### SetPrimaryActionTypeNil

`func (o *TargetEvidenceRow) SetPrimaryActionTypeNil(b bool)`

 SetPrimaryActionTypeNil sets the value for PrimaryActionType to be an explicit nil

### UnsetPrimaryActionType
`func (o *TargetEvidenceRow) UnsetPrimaryActionType()`

UnsetPrimaryActionType ensures that no value is present for PrimaryActionType, not even an explicit nil
### GetEvidenceSource

`func (o *TargetEvidenceRow) GetEvidenceSource() string`

GetEvidenceSource returns the EvidenceSource field if non-nil, zero value otherwise.

### GetEvidenceSourceOk

`func (o *TargetEvidenceRow) GetEvidenceSourceOk() (*string, bool)`

GetEvidenceSourceOk returns a tuple with the EvidenceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSource

`func (o *TargetEvidenceRow) SetEvidenceSource(v string)`

SetEvidenceSource sets EvidenceSource field to given value.

### HasEvidenceSource

`func (o *TargetEvidenceRow) HasEvidenceSource() bool`

HasEvidenceSource returns a boolean if a field has been set.

### SetEvidenceSourceNil

`func (o *TargetEvidenceRow) SetEvidenceSourceNil(b bool)`

 SetEvidenceSourceNil sets the value for EvidenceSource to be an explicit nil

### UnsetEvidenceSource
`func (o *TargetEvidenceRow) UnsetEvidenceSource()`

UnsetEvidenceSource ensures that no value is present for EvidenceSource, not even an explicit nil
### GetConditionId

`func (o *TargetEvidenceRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *TargetEvidenceRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *TargetEvidenceRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *TargetEvidenceRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *TargetEvidenceRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *TargetEvidenceRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *TargetEvidenceRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *TargetEvidenceRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *TargetEvidenceRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *TargetEvidenceRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *TargetEvidenceRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *TargetEvidenceRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetTrialId

`func (o *TargetEvidenceRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TargetEvidenceRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TargetEvidenceRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TargetEvidenceRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TargetEvidenceRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TargetEvidenceRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetMaxPhase

`func (o *TargetEvidenceRow) GetMaxPhase() int32`

GetMaxPhase returns the MaxPhase field if non-nil, zero value otherwise.

### GetMaxPhaseOk

`func (o *TargetEvidenceRow) GetMaxPhaseOk() (*int32, bool)`

GetMaxPhaseOk returns a tuple with the MaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhase

`func (o *TargetEvidenceRow) SetMaxPhase(v int32)`

SetMaxPhase sets MaxPhase field to given value.

### HasMaxPhase

`func (o *TargetEvidenceRow) HasMaxPhase() bool`

HasMaxPhase returns a boolean if a field has been set.

### SetMaxPhaseNil

`func (o *TargetEvidenceRow) SetMaxPhaseNil(b bool)`

 SetMaxPhaseNil sets the value for MaxPhase to be an explicit nil

### UnsetMaxPhase
`func (o *TargetEvidenceRow) UnsetMaxPhase()`

UnsetMaxPhase ensures that no value is present for MaxPhase, not even an explicit nil
### GetSponsorOrgId

`func (o *TargetEvidenceRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *TargetEvidenceRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *TargetEvidenceRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *TargetEvidenceRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *TargetEvidenceRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *TargetEvidenceRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorName

`func (o *TargetEvidenceRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *TargetEvidenceRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *TargetEvidenceRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *TargetEvidenceRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *TargetEvidenceRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *TargetEvidenceRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorCompanyId

`func (o *TargetEvidenceRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *TargetEvidenceRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *TargetEvidenceRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *TargetEvidenceRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *TargetEvidenceRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *TargetEvidenceRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *TargetEvidenceRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *TargetEvidenceRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *TargetEvidenceRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *TargetEvidenceRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *TargetEvidenceRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *TargetEvidenceRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *TargetEvidenceRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *TargetEvidenceRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *TargetEvidenceRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *TargetEvidenceRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *TargetEvidenceRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *TargetEvidenceRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *TargetEvidenceRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *TargetEvidenceRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *TargetEvidenceRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *TargetEvidenceRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *TargetEvidenceRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *TargetEvidenceRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


