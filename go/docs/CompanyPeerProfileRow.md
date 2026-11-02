# CompanyPeerProfileRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**PeerState** | Pointer to **NullableString** | measured when the company has at least one peer; no_live_owned_program when it has no live owned program; phase_not_reported when its live owned programs carry no phase; no_candidates when stage and modality are known and no other company shares them. | [optional] 
**NLiveOwnedPrograms** | Pointer to **NullableInt32** | Distinct live owned (drug, condition) programs of the company&#39;s subtree. | [optional] 
**NLiveOwnedDrugs** | Pointer to **NullableInt32** | Distinct live owned drugs of the company&#39;s subtree. | [optional] 
**LeadStage** | Pointer to **NullableString** | Furthest phase of the live owned programs: approved, phase3, phase2 (Phase 2/3 included), phase1 (Phase 1/2 included) or early_phase1. Null when no program carries a phase. | [optional] 
**LeadMaxPhase** | Pointer to **NullableInt32** | The phase code behind lead_stage: 6 Phase 4 or approved, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1. | [optional] 
**LeadDrugId** | Pointer to **NullableString** | The canonical drug id of the lead drug, the highest-phase live owned drug (ties: most trials, then drug id). | [optional] 
**LeadDrugName** | Pointer to **NullableString** | Name of the lead drug. | [optional] 
**NBiologicDrugs** | Pointer to **NullableInt32** | Live owned drugs whose GSRS substance type is protein or nucleic acid. | [optional] 
**BiologicShare** | Pointer to **NullableFloat32** | n_biologic_drugs over n_live_owned_drugs, 0 to 1. Null when the company has no live owned drug. | [optional] 
**ModalityBand** | Pointer to **NullableString** | biologic when biologic_share is at least 0.5, else non_biologic. | [optional] 
**PrimaryArea** | Pointer to **NullableString** | MeSH top-level category with the most live owned program trials. | [optional] 
**PrimaryAreaTrials** | Pointer to **NullableInt32** | Live owned program trials in primary_area. | [optional] 
**NumLeadTrials** | Pointer to **NullableInt32** | Distinct lead-sponsored trials of the company&#39;s subtree. | [optional] 
**SizeBand** | Pointer to **NullableString** | Band of num_lead_trials: 1-4, 5-19, 20-99 or 100+. | [optional] 
**NCellCandidates** | Pointer to **NullableInt32** | Other measured companies outside the company&#39;s group that share stage, modality, area and size (match level 1), before the top 10 are kept. | [optional] 
**NPeers** | Pointer to **NullableInt32** | Peer rows kept for the company, at most 10. | [optional] 
**CoverageMeasured** | Pointer to **NullableInt32** | Companies in this dataset with peer_state measured. The same value on every row. | [optional] 
**CoverageTotal** | Pointer to **NullableInt32** | Companies in this dataset. The same value on every row. | [optional] 

## Methods

### NewCompanyPeerProfileRow

`func NewCompanyPeerProfileRow() *CompanyPeerProfileRow`

NewCompanyPeerProfileRow instantiates a new CompanyPeerProfileRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyPeerProfileRowWithDefaults

`func NewCompanyPeerProfileRowWithDefaults() *CompanyPeerProfileRow`

NewCompanyPeerProfileRowWithDefaults instantiates a new CompanyPeerProfileRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyPeerProfileRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyPeerProfileRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyPeerProfileRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyPeerProfileRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyPeerProfileRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyPeerProfileRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyPeerProfileRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyPeerProfileRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyPeerProfileRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyPeerProfileRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyPeerProfileRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyPeerProfileRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyPeerProfileRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyPeerProfileRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyPeerProfileRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyPeerProfileRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyPeerProfileRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyPeerProfileRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyPeerProfileRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyPeerProfileRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyPeerProfileRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyPeerProfileRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyPeerProfileRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyPeerProfileRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetPeerState

`func (o *CompanyPeerProfileRow) GetPeerState() string`

GetPeerState returns the PeerState field if non-nil, zero value otherwise.

### GetPeerStateOk

`func (o *CompanyPeerProfileRow) GetPeerStateOk() (*string, bool)`

GetPeerStateOk returns a tuple with the PeerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerState

`func (o *CompanyPeerProfileRow) SetPeerState(v string)`

SetPeerState sets PeerState field to given value.

### HasPeerState

`func (o *CompanyPeerProfileRow) HasPeerState() bool`

HasPeerState returns a boolean if a field has been set.

### SetPeerStateNil

`func (o *CompanyPeerProfileRow) SetPeerStateNil(b bool)`

 SetPeerStateNil sets the value for PeerState to be an explicit nil

### UnsetPeerState
`func (o *CompanyPeerProfileRow) UnsetPeerState()`

UnsetPeerState ensures that no value is present for PeerState, not even an explicit nil
### GetNLiveOwnedPrograms

`func (o *CompanyPeerProfileRow) GetNLiveOwnedPrograms() int32`

GetNLiveOwnedPrograms returns the NLiveOwnedPrograms field if non-nil, zero value otherwise.

### GetNLiveOwnedProgramsOk

`func (o *CompanyPeerProfileRow) GetNLiveOwnedProgramsOk() (*int32, bool)`

GetNLiveOwnedProgramsOk returns a tuple with the NLiveOwnedPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLiveOwnedPrograms

`func (o *CompanyPeerProfileRow) SetNLiveOwnedPrograms(v int32)`

SetNLiveOwnedPrograms sets NLiveOwnedPrograms field to given value.

### HasNLiveOwnedPrograms

`func (o *CompanyPeerProfileRow) HasNLiveOwnedPrograms() bool`

HasNLiveOwnedPrograms returns a boolean if a field has been set.

### SetNLiveOwnedProgramsNil

`func (o *CompanyPeerProfileRow) SetNLiveOwnedProgramsNil(b bool)`

 SetNLiveOwnedProgramsNil sets the value for NLiveOwnedPrograms to be an explicit nil

### UnsetNLiveOwnedPrograms
`func (o *CompanyPeerProfileRow) UnsetNLiveOwnedPrograms()`

UnsetNLiveOwnedPrograms ensures that no value is present for NLiveOwnedPrograms, not even an explicit nil
### GetNLiveOwnedDrugs

`func (o *CompanyPeerProfileRow) GetNLiveOwnedDrugs() int32`

GetNLiveOwnedDrugs returns the NLiveOwnedDrugs field if non-nil, zero value otherwise.

### GetNLiveOwnedDrugsOk

`func (o *CompanyPeerProfileRow) GetNLiveOwnedDrugsOk() (*int32, bool)`

GetNLiveOwnedDrugsOk returns a tuple with the NLiveOwnedDrugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNLiveOwnedDrugs

`func (o *CompanyPeerProfileRow) SetNLiveOwnedDrugs(v int32)`

SetNLiveOwnedDrugs sets NLiveOwnedDrugs field to given value.

### HasNLiveOwnedDrugs

`func (o *CompanyPeerProfileRow) HasNLiveOwnedDrugs() bool`

HasNLiveOwnedDrugs returns a boolean if a field has been set.

### SetNLiveOwnedDrugsNil

`func (o *CompanyPeerProfileRow) SetNLiveOwnedDrugsNil(b bool)`

 SetNLiveOwnedDrugsNil sets the value for NLiveOwnedDrugs to be an explicit nil

### UnsetNLiveOwnedDrugs
`func (o *CompanyPeerProfileRow) UnsetNLiveOwnedDrugs()`

UnsetNLiveOwnedDrugs ensures that no value is present for NLiveOwnedDrugs, not even an explicit nil
### GetLeadStage

`func (o *CompanyPeerProfileRow) GetLeadStage() string`

GetLeadStage returns the LeadStage field if non-nil, zero value otherwise.

### GetLeadStageOk

`func (o *CompanyPeerProfileRow) GetLeadStageOk() (*string, bool)`

GetLeadStageOk returns a tuple with the LeadStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadStage

`func (o *CompanyPeerProfileRow) SetLeadStage(v string)`

SetLeadStage sets LeadStage field to given value.

### HasLeadStage

`func (o *CompanyPeerProfileRow) HasLeadStage() bool`

HasLeadStage returns a boolean if a field has been set.

### SetLeadStageNil

`func (o *CompanyPeerProfileRow) SetLeadStageNil(b bool)`

 SetLeadStageNil sets the value for LeadStage to be an explicit nil

### UnsetLeadStage
`func (o *CompanyPeerProfileRow) UnsetLeadStage()`

UnsetLeadStage ensures that no value is present for LeadStage, not even an explicit nil
### GetLeadMaxPhase

`func (o *CompanyPeerProfileRow) GetLeadMaxPhase() int32`

GetLeadMaxPhase returns the LeadMaxPhase field if non-nil, zero value otherwise.

### GetLeadMaxPhaseOk

`func (o *CompanyPeerProfileRow) GetLeadMaxPhaseOk() (*int32, bool)`

GetLeadMaxPhaseOk returns a tuple with the LeadMaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadMaxPhase

`func (o *CompanyPeerProfileRow) SetLeadMaxPhase(v int32)`

SetLeadMaxPhase sets LeadMaxPhase field to given value.

### HasLeadMaxPhase

`func (o *CompanyPeerProfileRow) HasLeadMaxPhase() bool`

HasLeadMaxPhase returns a boolean if a field has been set.

### SetLeadMaxPhaseNil

`func (o *CompanyPeerProfileRow) SetLeadMaxPhaseNil(b bool)`

 SetLeadMaxPhaseNil sets the value for LeadMaxPhase to be an explicit nil

### UnsetLeadMaxPhase
`func (o *CompanyPeerProfileRow) UnsetLeadMaxPhase()`

UnsetLeadMaxPhase ensures that no value is present for LeadMaxPhase, not even an explicit nil
### GetLeadDrugId

`func (o *CompanyPeerProfileRow) GetLeadDrugId() string`

GetLeadDrugId returns the LeadDrugId field if non-nil, zero value otherwise.

### GetLeadDrugIdOk

`func (o *CompanyPeerProfileRow) GetLeadDrugIdOk() (*string, bool)`

GetLeadDrugIdOk returns a tuple with the LeadDrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadDrugId

`func (o *CompanyPeerProfileRow) SetLeadDrugId(v string)`

SetLeadDrugId sets LeadDrugId field to given value.

### HasLeadDrugId

`func (o *CompanyPeerProfileRow) HasLeadDrugId() bool`

HasLeadDrugId returns a boolean if a field has been set.

### SetLeadDrugIdNil

`func (o *CompanyPeerProfileRow) SetLeadDrugIdNil(b bool)`

 SetLeadDrugIdNil sets the value for LeadDrugId to be an explicit nil

### UnsetLeadDrugId
`func (o *CompanyPeerProfileRow) UnsetLeadDrugId()`

UnsetLeadDrugId ensures that no value is present for LeadDrugId, not even an explicit nil
### GetLeadDrugName

`func (o *CompanyPeerProfileRow) GetLeadDrugName() string`

GetLeadDrugName returns the LeadDrugName field if non-nil, zero value otherwise.

### GetLeadDrugNameOk

`func (o *CompanyPeerProfileRow) GetLeadDrugNameOk() (*string, bool)`

GetLeadDrugNameOk returns a tuple with the LeadDrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadDrugName

`func (o *CompanyPeerProfileRow) SetLeadDrugName(v string)`

SetLeadDrugName sets LeadDrugName field to given value.

### HasLeadDrugName

`func (o *CompanyPeerProfileRow) HasLeadDrugName() bool`

HasLeadDrugName returns a boolean if a field has been set.

### SetLeadDrugNameNil

`func (o *CompanyPeerProfileRow) SetLeadDrugNameNil(b bool)`

 SetLeadDrugNameNil sets the value for LeadDrugName to be an explicit nil

### UnsetLeadDrugName
`func (o *CompanyPeerProfileRow) UnsetLeadDrugName()`

UnsetLeadDrugName ensures that no value is present for LeadDrugName, not even an explicit nil
### GetNBiologicDrugs

`func (o *CompanyPeerProfileRow) GetNBiologicDrugs() int32`

GetNBiologicDrugs returns the NBiologicDrugs field if non-nil, zero value otherwise.

### GetNBiologicDrugsOk

`func (o *CompanyPeerProfileRow) GetNBiologicDrugsOk() (*int32, bool)`

GetNBiologicDrugsOk returns a tuple with the NBiologicDrugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNBiologicDrugs

`func (o *CompanyPeerProfileRow) SetNBiologicDrugs(v int32)`

SetNBiologicDrugs sets NBiologicDrugs field to given value.

### HasNBiologicDrugs

`func (o *CompanyPeerProfileRow) HasNBiologicDrugs() bool`

HasNBiologicDrugs returns a boolean if a field has been set.

### SetNBiologicDrugsNil

`func (o *CompanyPeerProfileRow) SetNBiologicDrugsNil(b bool)`

 SetNBiologicDrugsNil sets the value for NBiologicDrugs to be an explicit nil

### UnsetNBiologicDrugs
`func (o *CompanyPeerProfileRow) UnsetNBiologicDrugs()`

UnsetNBiologicDrugs ensures that no value is present for NBiologicDrugs, not even an explicit nil
### GetBiologicShare

`func (o *CompanyPeerProfileRow) GetBiologicShare() float32`

GetBiologicShare returns the BiologicShare field if non-nil, zero value otherwise.

### GetBiologicShareOk

`func (o *CompanyPeerProfileRow) GetBiologicShareOk() (*float32, bool)`

GetBiologicShareOk returns a tuple with the BiologicShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiologicShare

`func (o *CompanyPeerProfileRow) SetBiologicShare(v float32)`

SetBiologicShare sets BiologicShare field to given value.

### HasBiologicShare

`func (o *CompanyPeerProfileRow) HasBiologicShare() bool`

HasBiologicShare returns a boolean if a field has been set.

### SetBiologicShareNil

`func (o *CompanyPeerProfileRow) SetBiologicShareNil(b bool)`

 SetBiologicShareNil sets the value for BiologicShare to be an explicit nil

### UnsetBiologicShare
`func (o *CompanyPeerProfileRow) UnsetBiologicShare()`

UnsetBiologicShare ensures that no value is present for BiologicShare, not even an explicit nil
### GetModalityBand

`func (o *CompanyPeerProfileRow) GetModalityBand() string`

GetModalityBand returns the ModalityBand field if non-nil, zero value otherwise.

### GetModalityBandOk

`func (o *CompanyPeerProfileRow) GetModalityBandOk() (*string, bool)`

GetModalityBandOk returns a tuple with the ModalityBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalityBand

`func (o *CompanyPeerProfileRow) SetModalityBand(v string)`

SetModalityBand sets ModalityBand field to given value.

### HasModalityBand

`func (o *CompanyPeerProfileRow) HasModalityBand() bool`

HasModalityBand returns a boolean if a field has been set.

### SetModalityBandNil

`func (o *CompanyPeerProfileRow) SetModalityBandNil(b bool)`

 SetModalityBandNil sets the value for ModalityBand to be an explicit nil

### UnsetModalityBand
`func (o *CompanyPeerProfileRow) UnsetModalityBand()`

UnsetModalityBand ensures that no value is present for ModalityBand, not even an explicit nil
### GetPrimaryArea

`func (o *CompanyPeerProfileRow) GetPrimaryArea() string`

GetPrimaryArea returns the PrimaryArea field if non-nil, zero value otherwise.

### GetPrimaryAreaOk

`func (o *CompanyPeerProfileRow) GetPrimaryAreaOk() (*string, bool)`

GetPrimaryAreaOk returns a tuple with the PrimaryArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryArea

`func (o *CompanyPeerProfileRow) SetPrimaryArea(v string)`

SetPrimaryArea sets PrimaryArea field to given value.

### HasPrimaryArea

`func (o *CompanyPeerProfileRow) HasPrimaryArea() bool`

HasPrimaryArea returns a boolean if a field has been set.

### SetPrimaryAreaNil

`func (o *CompanyPeerProfileRow) SetPrimaryAreaNil(b bool)`

 SetPrimaryAreaNil sets the value for PrimaryArea to be an explicit nil

### UnsetPrimaryArea
`func (o *CompanyPeerProfileRow) UnsetPrimaryArea()`

UnsetPrimaryArea ensures that no value is present for PrimaryArea, not even an explicit nil
### GetPrimaryAreaTrials

`func (o *CompanyPeerProfileRow) GetPrimaryAreaTrials() int32`

GetPrimaryAreaTrials returns the PrimaryAreaTrials field if non-nil, zero value otherwise.

### GetPrimaryAreaTrialsOk

`func (o *CompanyPeerProfileRow) GetPrimaryAreaTrialsOk() (*int32, bool)`

GetPrimaryAreaTrialsOk returns a tuple with the PrimaryAreaTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAreaTrials

`func (o *CompanyPeerProfileRow) SetPrimaryAreaTrials(v int32)`

SetPrimaryAreaTrials sets PrimaryAreaTrials field to given value.

### HasPrimaryAreaTrials

`func (o *CompanyPeerProfileRow) HasPrimaryAreaTrials() bool`

HasPrimaryAreaTrials returns a boolean if a field has been set.

### SetPrimaryAreaTrialsNil

`func (o *CompanyPeerProfileRow) SetPrimaryAreaTrialsNil(b bool)`

 SetPrimaryAreaTrialsNil sets the value for PrimaryAreaTrials to be an explicit nil

### UnsetPrimaryAreaTrials
`func (o *CompanyPeerProfileRow) UnsetPrimaryAreaTrials()`

UnsetPrimaryAreaTrials ensures that no value is present for PrimaryAreaTrials, not even an explicit nil
### GetNumLeadTrials

`func (o *CompanyPeerProfileRow) GetNumLeadTrials() int32`

GetNumLeadTrials returns the NumLeadTrials field if non-nil, zero value otherwise.

### GetNumLeadTrialsOk

`func (o *CompanyPeerProfileRow) GetNumLeadTrialsOk() (*int32, bool)`

GetNumLeadTrialsOk returns a tuple with the NumLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLeadTrials

`func (o *CompanyPeerProfileRow) SetNumLeadTrials(v int32)`

SetNumLeadTrials sets NumLeadTrials field to given value.

### HasNumLeadTrials

`func (o *CompanyPeerProfileRow) HasNumLeadTrials() bool`

HasNumLeadTrials returns a boolean if a field has been set.

### SetNumLeadTrialsNil

`func (o *CompanyPeerProfileRow) SetNumLeadTrialsNil(b bool)`

 SetNumLeadTrialsNil sets the value for NumLeadTrials to be an explicit nil

### UnsetNumLeadTrials
`func (o *CompanyPeerProfileRow) UnsetNumLeadTrials()`

UnsetNumLeadTrials ensures that no value is present for NumLeadTrials, not even an explicit nil
### GetSizeBand

`func (o *CompanyPeerProfileRow) GetSizeBand() string`

GetSizeBand returns the SizeBand field if non-nil, zero value otherwise.

### GetSizeBandOk

`func (o *CompanyPeerProfileRow) GetSizeBandOk() (*string, bool)`

GetSizeBandOk returns a tuple with the SizeBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBand

`func (o *CompanyPeerProfileRow) SetSizeBand(v string)`

SetSizeBand sets SizeBand field to given value.

### HasSizeBand

`func (o *CompanyPeerProfileRow) HasSizeBand() bool`

HasSizeBand returns a boolean if a field has been set.

### SetSizeBandNil

`func (o *CompanyPeerProfileRow) SetSizeBandNil(b bool)`

 SetSizeBandNil sets the value for SizeBand to be an explicit nil

### UnsetSizeBand
`func (o *CompanyPeerProfileRow) UnsetSizeBand()`

UnsetSizeBand ensures that no value is present for SizeBand, not even an explicit nil
### GetNCellCandidates

`func (o *CompanyPeerProfileRow) GetNCellCandidates() int32`

GetNCellCandidates returns the NCellCandidates field if non-nil, zero value otherwise.

### GetNCellCandidatesOk

`func (o *CompanyPeerProfileRow) GetNCellCandidatesOk() (*int32, bool)`

GetNCellCandidatesOk returns a tuple with the NCellCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCellCandidates

`func (o *CompanyPeerProfileRow) SetNCellCandidates(v int32)`

SetNCellCandidates sets NCellCandidates field to given value.

### HasNCellCandidates

`func (o *CompanyPeerProfileRow) HasNCellCandidates() bool`

HasNCellCandidates returns a boolean if a field has been set.

### SetNCellCandidatesNil

`func (o *CompanyPeerProfileRow) SetNCellCandidatesNil(b bool)`

 SetNCellCandidatesNil sets the value for NCellCandidates to be an explicit nil

### UnsetNCellCandidates
`func (o *CompanyPeerProfileRow) UnsetNCellCandidates()`

UnsetNCellCandidates ensures that no value is present for NCellCandidates, not even an explicit nil
### GetNPeers

`func (o *CompanyPeerProfileRow) GetNPeers() int32`

GetNPeers returns the NPeers field if non-nil, zero value otherwise.

### GetNPeersOk

`func (o *CompanyPeerProfileRow) GetNPeersOk() (*int32, bool)`

GetNPeersOk returns a tuple with the NPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPeers

`func (o *CompanyPeerProfileRow) SetNPeers(v int32)`

SetNPeers sets NPeers field to given value.

### HasNPeers

`func (o *CompanyPeerProfileRow) HasNPeers() bool`

HasNPeers returns a boolean if a field has been set.

### SetNPeersNil

`func (o *CompanyPeerProfileRow) SetNPeersNil(b bool)`

 SetNPeersNil sets the value for NPeers to be an explicit nil

### UnsetNPeers
`func (o *CompanyPeerProfileRow) UnsetNPeers()`

UnsetNPeers ensures that no value is present for NPeers, not even an explicit nil
### GetCoverageMeasured

`func (o *CompanyPeerProfileRow) GetCoverageMeasured() int32`

GetCoverageMeasured returns the CoverageMeasured field if non-nil, zero value otherwise.

### GetCoverageMeasuredOk

`func (o *CompanyPeerProfileRow) GetCoverageMeasuredOk() (*int32, bool)`

GetCoverageMeasuredOk returns a tuple with the CoverageMeasured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageMeasured

`func (o *CompanyPeerProfileRow) SetCoverageMeasured(v int32)`

SetCoverageMeasured sets CoverageMeasured field to given value.

### HasCoverageMeasured

`func (o *CompanyPeerProfileRow) HasCoverageMeasured() bool`

HasCoverageMeasured returns a boolean if a field has been set.

### SetCoverageMeasuredNil

`func (o *CompanyPeerProfileRow) SetCoverageMeasuredNil(b bool)`

 SetCoverageMeasuredNil sets the value for CoverageMeasured to be an explicit nil

### UnsetCoverageMeasured
`func (o *CompanyPeerProfileRow) UnsetCoverageMeasured()`

UnsetCoverageMeasured ensures that no value is present for CoverageMeasured, not even an explicit nil
### GetCoverageTotal

`func (o *CompanyPeerProfileRow) GetCoverageTotal() int32`

GetCoverageTotal returns the CoverageTotal field if non-nil, zero value otherwise.

### GetCoverageTotalOk

`func (o *CompanyPeerProfileRow) GetCoverageTotalOk() (*int32, bool)`

GetCoverageTotalOk returns a tuple with the CoverageTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageTotal

`func (o *CompanyPeerProfileRow) SetCoverageTotal(v int32)`

SetCoverageTotal sets CoverageTotal field to given value.

### HasCoverageTotal

`func (o *CompanyPeerProfileRow) HasCoverageTotal() bool`

HasCoverageTotal returns a boolean if a field has been set.

### SetCoverageTotalNil

`func (o *CompanyPeerProfileRow) SetCoverageTotalNil(b bool)`

 SetCoverageTotalNil sets the value for CoverageTotal to be an explicit nil

### UnsetCoverageTotal
`func (o *CompanyPeerProfileRow) UnsetCoverageTotal()`

UnsetCoverageTotal ensures that no value is present for CoverageTotal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


