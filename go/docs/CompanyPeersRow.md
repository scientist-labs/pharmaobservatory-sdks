# CompanyPeersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**PeerCompanyId** | Pointer to **NullableString** | Company level: the peer company. | [optional] 
**PeerCompanyName** | Pointer to **NullableString** | Display name of the peer company. | [optional] 
**PeerGroupId** | Pointer to **NullableString** | Group level: the peer&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**PeerGroupName** | Pointer to **NullableString** | Display name of the peer&#39;s group. | [optional] 
**PeerRank** | Pointer to **NullableInt32** | The peer&#39;s rank for the company, 1 is the closest match: match_level first, then the closeness of lead-trial counts, then of biologic share. Dense 1 to N. | [optional] 
**MatchLevel** | Pointer to **NullableInt32** | What the peer matched the company on. 1 stage, modality, area and size; 2 stage, modality and area; 3 stage, modality and size; 4 stage and modality. Stage and modality always match. | [optional] 
**MatchedArea** | Pointer to **NullableBool** | True when the peer shares the company&#39;s primary therapeutic area. | [optional] 
**MatchedSize** | Pointer to **NullableBool** | True when the peer shares the company&#39;s size band. | [optional] 
**LeadStage** | Pointer to **NullableString** | The peer&#39;s lead stage, the furthest phase of its live owned programs. Equal to the company&#39;s lead stage on every row. | [optional] 
**LeadDrugName** | Pointer to **NullableString** | The peer&#39;s lead drug, the highest-phase live owned drug (ties: most trials). | [optional] 
**BiologicShare** | Pointer to **NullableFloat32** | Share, 0 to 1, of the peer&#39;s live owned drugs that are biologic (GSRS protein or nucleic acid substance type). | [optional] 
**ModalityBand** | Pointer to **NullableString** | biologic when biologic_share is at least 0.5, else non_biologic. Equal to the company&#39;s band on every row. | [optional] 
**PrimaryArea** | Pointer to **NullableString** | The peer&#39;s primary therapeutic area: the MeSH top-level category with the most live owned program trials. | [optional] 
**NumLeadTrials** | Pointer to **NullableInt32** | Distinct lead-sponsored trials of the peer&#39;s subtree. | [optional] 
**SizeBand** | Pointer to **NullableString** | Band of num_lead_trials: 1-4, 5-19, 20-99 or 100+. | [optional] 

## Methods

### NewCompanyPeersRow

`func NewCompanyPeersRow() *CompanyPeersRow`

NewCompanyPeersRow instantiates a new CompanyPeersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyPeersRowWithDefaults

`func NewCompanyPeersRowWithDefaults() *CompanyPeersRow`

NewCompanyPeersRowWithDefaults instantiates a new CompanyPeersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyPeersRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyPeersRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyPeersRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyPeersRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyPeersRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyPeersRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyPeersRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyPeersRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyPeersRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyPeersRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyPeersRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyPeersRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyPeersRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyPeersRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyPeersRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyPeersRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyPeersRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyPeersRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyPeersRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyPeersRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyPeersRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyPeersRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyPeersRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyPeersRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetPeerCompanyId

`func (o *CompanyPeersRow) GetPeerCompanyId() string`

GetPeerCompanyId returns the PeerCompanyId field if non-nil, zero value otherwise.

### GetPeerCompanyIdOk

`func (o *CompanyPeersRow) GetPeerCompanyIdOk() (*string, bool)`

GetPeerCompanyIdOk returns a tuple with the PeerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCompanyId

`func (o *CompanyPeersRow) SetPeerCompanyId(v string)`

SetPeerCompanyId sets PeerCompanyId field to given value.

### HasPeerCompanyId

`func (o *CompanyPeersRow) HasPeerCompanyId() bool`

HasPeerCompanyId returns a boolean if a field has been set.

### SetPeerCompanyIdNil

`func (o *CompanyPeersRow) SetPeerCompanyIdNil(b bool)`

 SetPeerCompanyIdNil sets the value for PeerCompanyId to be an explicit nil

### UnsetPeerCompanyId
`func (o *CompanyPeersRow) UnsetPeerCompanyId()`

UnsetPeerCompanyId ensures that no value is present for PeerCompanyId, not even an explicit nil
### GetPeerCompanyName

`func (o *CompanyPeersRow) GetPeerCompanyName() string`

GetPeerCompanyName returns the PeerCompanyName field if non-nil, zero value otherwise.

### GetPeerCompanyNameOk

`func (o *CompanyPeersRow) GetPeerCompanyNameOk() (*string, bool)`

GetPeerCompanyNameOk returns a tuple with the PeerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCompanyName

`func (o *CompanyPeersRow) SetPeerCompanyName(v string)`

SetPeerCompanyName sets PeerCompanyName field to given value.

### HasPeerCompanyName

`func (o *CompanyPeersRow) HasPeerCompanyName() bool`

HasPeerCompanyName returns a boolean if a field has been set.

### SetPeerCompanyNameNil

`func (o *CompanyPeersRow) SetPeerCompanyNameNil(b bool)`

 SetPeerCompanyNameNil sets the value for PeerCompanyName to be an explicit nil

### UnsetPeerCompanyName
`func (o *CompanyPeersRow) UnsetPeerCompanyName()`

UnsetPeerCompanyName ensures that no value is present for PeerCompanyName, not even an explicit nil
### GetPeerGroupId

`func (o *CompanyPeersRow) GetPeerGroupId() string`

GetPeerGroupId returns the PeerGroupId field if non-nil, zero value otherwise.

### GetPeerGroupIdOk

`func (o *CompanyPeersRow) GetPeerGroupIdOk() (*string, bool)`

GetPeerGroupIdOk returns a tuple with the PeerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGroupId

`func (o *CompanyPeersRow) SetPeerGroupId(v string)`

SetPeerGroupId sets PeerGroupId field to given value.

### HasPeerGroupId

`func (o *CompanyPeersRow) HasPeerGroupId() bool`

HasPeerGroupId returns a boolean if a field has been set.

### SetPeerGroupIdNil

`func (o *CompanyPeersRow) SetPeerGroupIdNil(b bool)`

 SetPeerGroupIdNil sets the value for PeerGroupId to be an explicit nil

### UnsetPeerGroupId
`func (o *CompanyPeersRow) UnsetPeerGroupId()`

UnsetPeerGroupId ensures that no value is present for PeerGroupId, not even an explicit nil
### GetPeerGroupName

`func (o *CompanyPeersRow) GetPeerGroupName() string`

GetPeerGroupName returns the PeerGroupName field if non-nil, zero value otherwise.

### GetPeerGroupNameOk

`func (o *CompanyPeersRow) GetPeerGroupNameOk() (*string, bool)`

GetPeerGroupNameOk returns a tuple with the PeerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGroupName

`func (o *CompanyPeersRow) SetPeerGroupName(v string)`

SetPeerGroupName sets PeerGroupName field to given value.

### HasPeerGroupName

`func (o *CompanyPeersRow) HasPeerGroupName() bool`

HasPeerGroupName returns a boolean if a field has been set.

### SetPeerGroupNameNil

`func (o *CompanyPeersRow) SetPeerGroupNameNil(b bool)`

 SetPeerGroupNameNil sets the value for PeerGroupName to be an explicit nil

### UnsetPeerGroupName
`func (o *CompanyPeersRow) UnsetPeerGroupName()`

UnsetPeerGroupName ensures that no value is present for PeerGroupName, not even an explicit nil
### GetPeerRank

`func (o *CompanyPeersRow) GetPeerRank() int32`

GetPeerRank returns the PeerRank field if non-nil, zero value otherwise.

### GetPeerRankOk

`func (o *CompanyPeersRow) GetPeerRankOk() (*int32, bool)`

GetPeerRankOk returns a tuple with the PeerRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerRank

`func (o *CompanyPeersRow) SetPeerRank(v int32)`

SetPeerRank sets PeerRank field to given value.

### HasPeerRank

`func (o *CompanyPeersRow) HasPeerRank() bool`

HasPeerRank returns a boolean if a field has been set.

### SetPeerRankNil

`func (o *CompanyPeersRow) SetPeerRankNil(b bool)`

 SetPeerRankNil sets the value for PeerRank to be an explicit nil

### UnsetPeerRank
`func (o *CompanyPeersRow) UnsetPeerRank()`

UnsetPeerRank ensures that no value is present for PeerRank, not even an explicit nil
### GetMatchLevel

`func (o *CompanyPeersRow) GetMatchLevel() int32`

GetMatchLevel returns the MatchLevel field if non-nil, zero value otherwise.

### GetMatchLevelOk

`func (o *CompanyPeersRow) GetMatchLevelOk() (*int32, bool)`

GetMatchLevelOk returns a tuple with the MatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchLevel

`func (o *CompanyPeersRow) SetMatchLevel(v int32)`

SetMatchLevel sets MatchLevel field to given value.

### HasMatchLevel

`func (o *CompanyPeersRow) HasMatchLevel() bool`

HasMatchLevel returns a boolean if a field has been set.

### SetMatchLevelNil

`func (o *CompanyPeersRow) SetMatchLevelNil(b bool)`

 SetMatchLevelNil sets the value for MatchLevel to be an explicit nil

### UnsetMatchLevel
`func (o *CompanyPeersRow) UnsetMatchLevel()`

UnsetMatchLevel ensures that no value is present for MatchLevel, not even an explicit nil
### GetMatchedArea

`func (o *CompanyPeersRow) GetMatchedArea() bool`

GetMatchedArea returns the MatchedArea field if non-nil, zero value otherwise.

### GetMatchedAreaOk

`func (o *CompanyPeersRow) GetMatchedAreaOk() (*bool, bool)`

GetMatchedAreaOk returns a tuple with the MatchedArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedArea

`func (o *CompanyPeersRow) SetMatchedArea(v bool)`

SetMatchedArea sets MatchedArea field to given value.

### HasMatchedArea

`func (o *CompanyPeersRow) HasMatchedArea() bool`

HasMatchedArea returns a boolean if a field has been set.

### SetMatchedAreaNil

`func (o *CompanyPeersRow) SetMatchedAreaNil(b bool)`

 SetMatchedAreaNil sets the value for MatchedArea to be an explicit nil

### UnsetMatchedArea
`func (o *CompanyPeersRow) UnsetMatchedArea()`

UnsetMatchedArea ensures that no value is present for MatchedArea, not even an explicit nil
### GetMatchedSize

`func (o *CompanyPeersRow) GetMatchedSize() bool`

GetMatchedSize returns the MatchedSize field if non-nil, zero value otherwise.

### GetMatchedSizeOk

`func (o *CompanyPeersRow) GetMatchedSizeOk() (*bool, bool)`

GetMatchedSizeOk returns a tuple with the MatchedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedSize

`func (o *CompanyPeersRow) SetMatchedSize(v bool)`

SetMatchedSize sets MatchedSize field to given value.

### HasMatchedSize

`func (o *CompanyPeersRow) HasMatchedSize() bool`

HasMatchedSize returns a boolean if a field has been set.

### SetMatchedSizeNil

`func (o *CompanyPeersRow) SetMatchedSizeNil(b bool)`

 SetMatchedSizeNil sets the value for MatchedSize to be an explicit nil

### UnsetMatchedSize
`func (o *CompanyPeersRow) UnsetMatchedSize()`

UnsetMatchedSize ensures that no value is present for MatchedSize, not even an explicit nil
### GetLeadStage

`func (o *CompanyPeersRow) GetLeadStage() string`

GetLeadStage returns the LeadStage field if non-nil, zero value otherwise.

### GetLeadStageOk

`func (o *CompanyPeersRow) GetLeadStageOk() (*string, bool)`

GetLeadStageOk returns a tuple with the LeadStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadStage

`func (o *CompanyPeersRow) SetLeadStage(v string)`

SetLeadStage sets LeadStage field to given value.

### HasLeadStage

`func (o *CompanyPeersRow) HasLeadStage() bool`

HasLeadStage returns a boolean if a field has been set.

### SetLeadStageNil

`func (o *CompanyPeersRow) SetLeadStageNil(b bool)`

 SetLeadStageNil sets the value for LeadStage to be an explicit nil

### UnsetLeadStage
`func (o *CompanyPeersRow) UnsetLeadStage()`

UnsetLeadStage ensures that no value is present for LeadStage, not even an explicit nil
### GetLeadDrugName

`func (o *CompanyPeersRow) GetLeadDrugName() string`

GetLeadDrugName returns the LeadDrugName field if non-nil, zero value otherwise.

### GetLeadDrugNameOk

`func (o *CompanyPeersRow) GetLeadDrugNameOk() (*string, bool)`

GetLeadDrugNameOk returns a tuple with the LeadDrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadDrugName

`func (o *CompanyPeersRow) SetLeadDrugName(v string)`

SetLeadDrugName sets LeadDrugName field to given value.

### HasLeadDrugName

`func (o *CompanyPeersRow) HasLeadDrugName() bool`

HasLeadDrugName returns a boolean if a field has been set.

### SetLeadDrugNameNil

`func (o *CompanyPeersRow) SetLeadDrugNameNil(b bool)`

 SetLeadDrugNameNil sets the value for LeadDrugName to be an explicit nil

### UnsetLeadDrugName
`func (o *CompanyPeersRow) UnsetLeadDrugName()`

UnsetLeadDrugName ensures that no value is present for LeadDrugName, not even an explicit nil
### GetBiologicShare

`func (o *CompanyPeersRow) GetBiologicShare() float32`

GetBiologicShare returns the BiologicShare field if non-nil, zero value otherwise.

### GetBiologicShareOk

`func (o *CompanyPeersRow) GetBiologicShareOk() (*float32, bool)`

GetBiologicShareOk returns a tuple with the BiologicShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiologicShare

`func (o *CompanyPeersRow) SetBiologicShare(v float32)`

SetBiologicShare sets BiologicShare field to given value.

### HasBiologicShare

`func (o *CompanyPeersRow) HasBiologicShare() bool`

HasBiologicShare returns a boolean if a field has been set.

### SetBiologicShareNil

`func (o *CompanyPeersRow) SetBiologicShareNil(b bool)`

 SetBiologicShareNil sets the value for BiologicShare to be an explicit nil

### UnsetBiologicShare
`func (o *CompanyPeersRow) UnsetBiologicShare()`

UnsetBiologicShare ensures that no value is present for BiologicShare, not even an explicit nil
### GetModalityBand

`func (o *CompanyPeersRow) GetModalityBand() string`

GetModalityBand returns the ModalityBand field if non-nil, zero value otherwise.

### GetModalityBandOk

`func (o *CompanyPeersRow) GetModalityBandOk() (*string, bool)`

GetModalityBandOk returns a tuple with the ModalityBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalityBand

`func (o *CompanyPeersRow) SetModalityBand(v string)`

SetModalityBand sets ModalityBand field to given value.

### HasModalityBand

`func (o *CompanyPeersRow) HasModalityBand() bool`

HasModalityBand returns a boolean if a field has been set.

### SetModalityBandNil

`func (o *CompanyPeersRow) SetModalityBandNil(b bool)`

 SetModalityBandNil sets the value for ModalityBand to be an explicit nil

### UnsetModalityBand
`func (o *CompanyPeersRow) UnsetModalityBand()`

UnsetModalityBand ensures that no value is present for ModalityBand, not even an explicit nil
### GetPrimaryArea

`func (o *CompanyPeersRow) GetPrimaryArea() string`

GetPrimaryArea returns the PrimaryArea field if non-nil, zero value otherwise.

### GetPrimaryAreaOk

`func (o *CompanyPeersRow) GetPrimaryAreaOk() (*string, bool)`

GetPrimaryAreaOk returns a tuple with the PrimaryArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryArea

`func (o *CompanyPeersRow) SetPrimaryArea(v string)`

SetPrimaryArea sets PrimaryArea field to given value.

### HasPrimaryArea

`func (o *CompanyPeersRow) HasPrimaryArea() bool`

HasPrimaryArea returns a boolean if a field has been set.

### SetPrimaryAreaNil

`func (o *CompanyPeersRow) SetPrimaryAreaNil(b bool)`

 SetPrimaryAreaNil sets the value for PrimaryArea to be an explicit nil

### UnsetPrimaryArea
`func (o *CompanyPeersRow) UnsetPrimaryArea()`

UnsetPrimaryArea ensures that no value is present for PrimaryArea, not even an explicit nil
### GetNumLeadTrials

`func (o *CompanyPeersRow) GetNumLeadTrials() int32`

GetNumLeadTrials returns the NumLeadTrials field if non-nil, zero value otherwise.

### GetNumLeadTrialsOk

`func (o *CompanyPeersRow) GetNumLeadTrialsOk() (*int32, bool)`

GetNumLeadTrialsOk returns a tuple with the NumLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLeadTrials

`func (o *CompanyPeersRow) SetNumLeadTrials(v int32)`

SetNumLeadTrials sets NumLeadTrials field to given value.

### HasNumLeadTrials

`func (o *CompanyPeersRow) HasNumLeadTrials() bool`

HasNumLeadTrials returns a boolean if a field has been set.

### SetNumLeadTrialsNil

`func (o *CompanyPeersRow) SetNumLeadTrialsNil(b bool)`

 SetNumLeadTrialsNil sets the value for NumLeadTrials to be an explicit nil

### UnsetNumLeadTrials
`func (o *CompanyPeersRow) UnsetNumLeadTrials()`

UnsetNumLeadTrials ensures that no value is present for NumLeadTrials, not even an explicit nil
### GetSizeBand

`func (o *CompanyPeersRow) GetSizeBand() string`

GetSizeBand returns the SizeBand field if non-nil, zero value otherwise.

### GetSizeBandOk

`func (o *CompanyPeersRow) GetSizeBandOk() (*string, bool)`

GetSizeBandOk returns a tuple with the SizeBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBand

`func (o *CompanyPeersRow) SetSizeBand(v string)`

SetSizeBand sets SizeBand field to given value.

### HasSizeBand

`func (o *CompanyPeersRow) HasSizeBand() bool`

HasSizeBand returns a boolean if a field has been set.

### SetSizeBandNil

`func (o *CompanyPeersRow) SetSizeBandNil(b bool)`

 SetSizeBandNil sets the value for SizeBand to be an explicit nil

### UnsetSizeBand
`func (o *CompanyPeersRow) UnsetSizeBand()`

UnsetSizeBand ensures that no value is present for SizeBand, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


