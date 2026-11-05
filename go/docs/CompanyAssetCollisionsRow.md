# CompanyAssetCollisionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**MoleculeId** | Pointer to **NullableString** | Molecule level: the company&#39;s asset, the resolved drug (a drug with two canonical ids is one asset). | [optional] 
**MoleculeName** | Pointer to **NullableString** | The asset&#39;s display name, from the company&#39;s most-trialled record of it. | [optional] 
**CompetitorCompanyId** | Pointer to **NullableString** | Company level: the company whose asset collides. | [optional] 
**CompetitorCompanyName** | Pointer to **NullableString** | Display name of the competitor&#39;s company. | [optional] 
**CompetitorGroupId** | Pointer to **NullableString** | Group level: the competitor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**CompetitorGroupName** | Pointer to **NullableString** | Display name of the competitor&#39;s group. | [optional] 
**CompetitorMoleculeId** | Pointer to **NullableString** | Molecule level: the competitor&#39;s asset, the resolved drug. | [optional] 
**CompetitorMoleculeName** | Pointer to **NullableString** | The competitor asset&#39;s display name. | [optional] 
**TargetFamily** | Pointer to **NullableString** | The company drug&#39;s target family of the shared gene (the Open Targets target name, so a taxane&#39;s tubulin isotypes are one family; the ChEMBL mechanism text for a drug Open Targets lacks; the gene symbol for a curated investigational asset). | [optional] 
**SharedSymbols** | Pointer to **[]string** | Gene symbols of the shared genes within the family. | [optional] 
**ActionType** | Pointer to **NullableString** | The company drug&#39;s action on the family (Open Targets or ChEMBL vocabulary); null when no source states one. | [optional] 
**CompetitorActionType** | Pointer to **NullableString** | The competitor drug&#39;s action on the shared genes; null when no source states one. | [optional] 
**SameAction** | Pointer to **NullableBool** | True when the two drugs have the same action type on a shared gene of the family. | [optional] 
**NSharedIndications** | Pointer to **NullableInt32** | Disease areas both programs are in, among live programs. | [optional] 
**TopSharedIndications** | Pointer to **[]string** | Up to 3 shared disease area names, by the pair&#39;s combined trial count. An empty list when none. | [optional] 
**Stage** | Pointer to **NullableInt32** | The less advanced of the two programs&#39; furthest phase: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1, -1 phase not applicable. | [optional] 
**LastYear** | Pointer to **NullableInt32** | The later of the two programs&#39; last trial start year. | [optional] 
**NHeadToHeadTrials** | Pointer to **NullableInt32** | Trials led by either company in which one drug is experimental and the other an active comparator. | [optional] 
**OwnershipBasis** | Pointer to **NullableString** | Why the company owns the drug. experimental_arm and regulatory_holder rest on evidence; inferred_lead is an inference (no arm data says otherwise). | [optional] 
**CompetitorOwnershipBasis** | Pointer to **NullableString** | Why the competitor owns its drug, same vocabulary as ownership_basis. | [optional] 
**FamilyIdf** | Pointer to **NullableFloat32** | The rarity weight of the family: the natural log of (companies holding any live, owned target family) over (companies holding this one). | [optional] 
**CollisionWeight** | Pointer to **NullableFloat32** | The weight of the pair, family_idf times a factor for same action, a shared indication, the stage, recency and a head-to-head trial. | [optional] 

## Methods

### NewCompanyAssetCollisionsRow

`func NewCompanyAssetCollisionsRow() *CompanyAssetCollisionsRow`

NewCompanyAssetCollisionsRow instantiates a new CompanyAssetCollisionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyAssetCollisionsRowWithDefaults

`func NewCompanyAssetCollisionsRowWithDefaults() *CompanyAssetCollisionsRow`

NewCompanyAssetCollisionsRowWithDefaults instantiates a new CompanyAssetCollisionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyAssetCollisionsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyAssetCollisionsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyAssetCollisionsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyAssetCollisionsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyAssetCollisionsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyAssetCollisionsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyAssetCollisionsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyAssetCollisionsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyAssetCollisionsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyAssetCollisionsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyAssetCollisionsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyAssetCollisionsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyAssetCollisionsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyAssetCollisionsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyAssetCollisionsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyAssetCollisionsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyAssetCollisionsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyAssetCollisionsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyAssetCollisionsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyAssetCollisionsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyAssetCollisionsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyAssetCollisionsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyAssetCollisionsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyAssetCollisionsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetMoleculeId

`func (o *CompanyAssetCollisionsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *CompanyAssetCollisionsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *CompanyAssetCollisionsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *CompanyAssetCollisionsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *CompanyAssetCollisionsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *CompanyAssetCollisionsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetMoleculeName

`func (o *CompanyAssetCollisionsRow) GetMoleculeName() string`

GetMoleculeName returns the MoleculeName field if non-nil, zero value otherwise.

### GetMoleculeNameOk

`func (o *CompanyAssetCollisionsRow) GetMoleculeNameOk() (*string, bool)`

GetMoleculeNameOk returns a tuple with the MoleculeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeName

`func (o *CompanyAssetCollisionsRow) SetMoleculeName(v string)`

SetMoleculeName sets MoleculeName field to given value.

### HasMoleculeName

`func (o *CompanyAssetCollisionsRow) HasMoleculeName() bool`

HasMoleculeName returns a boolean if a field has been set.

### SetMoleculeNameNil

`func (o *CompanyAssetCollisionsRow) SetMoleculeNameNil(b bool)`

 SetMoleculeNameNil sets the value for MoleculeName to be an explicit nil

### UnsetMoleculeName
`func (o *CompanyAssetCollisionsRow) UnsetMoleculeName()`

UnsetMoleculeName ensures that no value is present for MoleculeName, not even an explicit nil
### GetCompetitorCompanyId

`func (o *CompanyAssetCollisionsRow) GetCompetitorCompanyId() string`

GetCompetitorCompanyId returns the CompetitorCompanyId field if non-nil, zero value otherwise.

### GetCompetitorCompanyIdOk

`func (o *CompanyAssetCollisionsRow) GetCompetitorCompanyIdOk() (*string, bool)`

GetCompetitorCompanyIdOk returns a tuple with the CompetitorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorCompanyId

`func (o *CompanyAssetCollisionsRow) SetCompetitorCompanyId(v string)`

SetCompetitorCompanyId sets CompetitorCompanyId field to given value.

### HasCompetitorCompanyId

`func (o *CompanyAssetCollisionsRow) HasCompetitorCompanyId() bool`

HasCompetitorCompanyId returns a boolean if a field has been set.

### SetCompetitorCompanyIdNil

`func (o *CompanyAssetCollisionsRow) SetCompetitorCompanyIdNil(b bool)`

 SetCompetitorCompanyIdNil sets the value for CompetitorCompanyId to be an explicit nil

### UnsetCompetitorCompanyId
`func (o *CompanyAssetCollisionsRow) UnsetCompetitorCompanyId()`

UnsetCompetitorCompanyId ensures that no value is present for CompetitorCompanyId, not even an explicit nil
### GetCompetitorCompanyName

`func (o *CompanyAssetCollisionsRow) GetCompetitorCompanyName() string`

GetCompetitorCompanyName returns the CompetitorCompanyName field if non-nil, zero value otherwise.

### GetCompetitorCompanyNameOk

`func (o *CompanyAssetCollisionsRow) GetCompetitorCompanyNameOk() (*string, bool)`

GetCompetitorCompanyNameOk returns a tuple with the CompetitorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorCompanyName

`func (o *CompanyAssetCollisionsRow) SetCompetitorCompanyName(v string)`

SetCompetitorCompanyName sets CompetitorCompanyName field to given value.

### HasCompetitorCompanyName

`func (o *CompanyAssetCollisionsRow) HasCompetitorCompanyName() bool`

HasCompetitorCompanyName returns a boolean if a field has been set.

### SetCompetitorCompanyNameNil

`func (o *CompanyAssetCollisionsRow) SetCompetitorCompanyNameNil(b bool)`

 SetCompetitorCompanyNameNil sets the value for CompetitorCompanyName to be an explicit nil

### UnsetCompetitorCompanyName
`func (o *CompanyAssetCollisionsRow) UnsetCompetitorCompanyName()`

UnsetCompetitorCompanyName ensures that no value is present for CompetitorCompanyName, not even an explicit nil
### GetCompetitorGroupId

`func (o *CompanyAssetCollisionsRow) GetCompetitorGroupId() string`

GetCompetitorGroupId returns the CompetitorGroupId field if non-nil, zero value otherwise.

### GetCompetitorGroupIdOk

`func (o *CompanyAssetCollisionsRow) GetCompetitorGroupIdOk() (*string, bool)`

GetCompetitorGroupIdOk returns a tuple with the CompetitorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorGroupId

`func (o *CompanyAssetCollisionsRow) SetCompetitorGroupId(v string)`

SetCompetitorGroupId sets CompetitorGroupId field to given value.

### HasCompetitorGroupId

`func (o *CompanyAssetCollisionsRow) HasCompetitorGroupId() bool`

HasCompetitorGroupId returns a boolean if a field has been set.

### SetCompetitorGroupIdNil

`func (o *CompanyAssetCollisionsRow) SetCompetitorGroupIdNil(b bool)`

 SetCompetitorGroupIdNil sets the value for CompetitorGroupId to be an explicit nil

### UnsetCompetitorGroupId
`func (o *CompanyAssetCollisionsRow) UnsetCompetitorGroupId()`

UnsetCompetitorGroupId ensures that no value is present for CompetitorGroupId, not even an explicit nil
### GetCompetitorGroupName

`func (o *CompanyAssetCollisionsRow) GetCompetitorGroupName() string`

GetCompetitorGroupName returns the CompetitorGroupName field if non-nil, zero value otherwise.

### GetCompetitorGroupNameOk

`func (o *CompanyAssetCollisionsRow) GetCompetitorGroupNameOk() (*string, bool)`

GetCompetitorGroupNameOk returns a tuple with the CompetitorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorGroupName

`func (o *CompanyAssetCollisionsRow) SetCompetitorGroupName(v string)`

SetCompetitorGroupName sets CompetitorGroupName field to given value.

### HasCompetitorGroupName

`func (o *CompanyAssetCollisionsRow) HasCompetitorGroupName() bool`

HasCompetitorGroupName returns a boolean if a field has been set.

### SetCompetitorGroupNameNil

`func (o *CompanyAssetCollisionsRow) SetCompetitorGroupNameNil(b bool)`

 SetCompetitorGroupNameNil sets the value for CompetitorGroupName to be an explicit nil

### UnsetCompetitorGroupName
`func (o *CompanyAssetCollisionsRow) UnsetCompetitorGroupName()`

UnsetCompetitorGroupName ensures that no value is present for CompetitorGroupName, not even an explicit nil
### GetCompetitorMoleculeId

`func (o *CompanyAssetCollisionsRow) GetCompetitorMoleculeId() string`

GetCompetitorMoleculeId returns the CompetitorMoleculeId field if non-nil, zero value otherwise.

### GetCompetitorMoleculeIdOk

`func (o *CompanyAssetCollisionsRow) GetCompetitorMoleculeIdOk() (*string, bool)`

GetCompetitorMoleculeIdOk returns a tuple with the CompetitorMoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorMoleculeId

`func (o *CompanyAssetCollisionsRow) SetCompetitorMoleculeId(v string)`

SetCompetitorMoleculeId sets CompetitorMoleculeId field to given value.

### HasCompetitorMoleculeId

`func (o *CompanyAssetCollisionsRow) HasCompetitorMoleculeId() bool`

HasCompetitorMoleculeId returns a boolean if a field has been set.

### SetCompetitorMoleculeIdNil

`func (o *CompanyAssetCollisionsRow) SetCompetitorMoleculeIdNil(b bool)`

 SetCompetitorMoleculeIdNil sets the value for CompetitorMoleculeId to be an explicit nil

### UnsetCompetitorMoleculeId
`func (o *CompanyAssetCollisionsRow) UnsetCompetitorMoleculeId()`

UnsetCompetitorMoleculeId ensures that no value is present for CompetitorMoleculeId, not even an explicit nil
### GetCompetitorMoleculeName

`func (o *CompanyAssetCollisionsRow) GetCompetitorMoleculeName() string`

GetCompetitorMoleculeName returns the CompetitorMoleculeName field if non-nil, zero value otherwise.

### GetCompetitorMoleculeNameOk

`func (o *CompanyAssetCollisionsRow) GetCompetitorMoleculeNameOk() (*string, bool)`

GetCompetitorMoleculeNameOk returns a tuple with the CompetitorMoleculeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorMoleculeName

`func (o *CompanyAssetCollisionsRow) SetCompetitorMoleculeName(v string)`

SetCompetitorMoleculeName sets CompetitorMoleculeName field to given value.

### HasCompetitorMoleculeName

`func (o *CompanyAssetCollisionsRow) HasCompetitorMoleculeName() bool`

HasCompetitorMoleculeName returns a boolean if a field has been set.

### SetCompetitorMoleculeNameNil

`func (o *CompanyAssetCollisionsRow) SetCompetitorMoleculeNameNil(b bool)`

 SetCompetitorMoleculeNameNil sets the value for CompetitorMoleculeName to be an explicit nil

### UnsetCompetitorMoleculeName
`func (o *CompanyAssetCollisionsRow) UnsetCompetitorMoleculeName()`

UnsetCompetitorMoleculeName ensures that no value is present for CompetitorMoleculeName, not even an explicit nil
### GetTargetFamily

`func (o *CompanyAssetCollisionsRow) GetTargetFamily() string`

GetTargetFamily returns the TargetFamily field if non-nil, zero value otherwise.

### GetTargetFamilyOk

`func (o *CompanyAssetCollisionsRow) GetTargetFamilyOk() (*string, bool)`

GetTargetFamilyOk returns a tuple with the TargetFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFamily

`func (o *CompanyAssetCollisionsRow) SetTargetFamily(v string)`

SetTargetFamily sets TargetFamily field to given value.

### HasTargetFamily

`func (o *CompanyAssetCollisionsRow) HasTargetFamily() bool`

HasTargetFamily returns a boolean if a field has been set.

### SetTargetFamilyNil

`func (o *CompanyAssetCollisionsRow) SetTargetFamilyNil(b bool)`

 SetTargetFamilyNil sets the value for TargetFamily to be an explicit nil

### UnsetTargetFamily
`func (o *CompanyAssetCollisionsRow) UnsetTargetFamily()`

UnsetTargetFamily ensures that no value is present for TargetFamily, not even an explicit nil
### GetSharedSymbols

`func (o *CompanyAssetCollisionsRow) GetSharedSymbols() []string`

GetSharedSymbols returns the SharedSymbols field if non-nil, zero value otherwise.

### GetSharedSymbolsOk

`func (o *CompanyAssetCollisionsRow) GetSharedSymbolsOk() (*[]string, bool)`

GetSharedSymbolsOk returns a tuple with the SharedSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSymbols

`func (o *CompanyAssetCollisionsRow) SetSharedSymbols(v []string)`

SetSharedSymbols sets SharedSymbols field to given value.

### HasSharedSymbols

`func (o *CompanyAssetCollisionsRow) HasSharedSymbols() bool`

HasSharedSymbols returns a boolean if a field has been set.

### SetSharedSymbolsNil

`func (o *CompanyAssetCollisionsRow) SetSharedSymbolsNil(b bool)`

 SetSharedSymbolsNil sets the value for SharedSymbols to be an explicit nil

### UnsetSharedSymbols
`func (o *CompanyAssetCollisionsRow) UnsetSharedSymbols()`

UnsetSharedSymbols ensures that no value is present for SharedSymbols, not even an explicit nil
### GetActionType

`func (o *CompanyAssetCollisionsRow) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *CompanyAssetCollisionsRow) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *CompanyAssetCollisionsRow) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *CompanyAssetCollisionsRow) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionTypeNil

`func (o *CompanyAssetCollisionsRow) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *CompanyAssetCollisionsRow) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetCompetitorActionType

`func (o *CompanyAssetCollisionsRow) GetCompetitorActionType() string`

GetCompetitorActionType returns the CompetitorActionType field if non-nil, zero value otherwise.

### GetCompetitorActionTypeOk

`func (o *CompanyAssetCollisionsRow) GetCompetitorActionTypeOk() (*string, bool)`

GetCompetitorActionTypeOk returns a tuple with the CompetitorActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorActionType

`func (o *CompanyAssetCollisionsRow) SetCompetitorActionType(v string)`

SetCompetitorActionType sets CompetitorActionType field to given value.

### HasCompetitorActionType

`func (o *CompanyAssetCollisionsRow) HasCompetitorActionType() bool`

HasCompetitorActionType returns a boolean if a field has been set.

### SetCompetitorActionTypeNil

`func (o *CompanyAssetCollisionsRow) SetCompetitorActionTypeNil(b bool)`

 SetCompetitorActionTypeNil sets the value for CompetitorActionType to be an explicit nil

### UnsetCompetitorActionType
`func (o *CompanyAssetCollisionsRow) UnsetCompetitorActionType()`

UnsetCompetitorActionType ensures that no value is present for CompetitorActionType, not even an explicit nil
### GetSameAction

`func (o *CompanyAssetCollisionsRow) GetSameAction() bool`

GetSameAction returns the SameAction field if non-nil, zero value otherwise.

### GetSameActionOk

`func (o *CompanyAssetCollisionsRow) GetSameActionOk() (*bool, bool)`

GetSameActionOk returns a tuple with the SameAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameAction

`func (o *CompanyAssetCollisionsRow) SetSameAction(v bool)`

SetSameAction sets SameAction field to given value.

### HasSameAction

`func (o *CompanyAssetCollisionsRow) HasSameAction() bool`

HasSameAction returns a boolean if a field has been set.

### SetSameActionNil

`func (o *CompanyAssetCollisionsRow) SetSameActionNil(b bool)`

 SetSameActionNil sets the value for SameAction to be an explicit nil

### UnsetSameAction
`func (o *CompanyAssetCollisionsRow) UnsetSameAction()`

UnsetSameAction ensures that no value is present for SameAction, not even an explicit nil
### GetNSharedIndications

`func (o *CompanyAssetCollisionsRow) GetNSharedIndications() int32`

GetNSharedIndications returns the NSharedIndications field if non-nil, zero value otherwise.

### GetNSharedIndicationsOk

`func (o *CompanyAssetCollisionsRow) GetNSharedIndicationsOk() (*int32, bool)`

GetNSharedIndicationsOk returns a tuple with the NSharedIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedIndications

`func (o *CompanyAssetCollisionsRow) SetNSharedIndications(v int32)`

SetNSharedIndications sets NSharedIndications field to given value.

### HasNSharedIndications

`func (o *CompanyAssetCollisionsRow) HasNSharedIndications() bool`

HasNSharedIndications returns a boolean if a field has been set.

### SetNSharedIndicationsNil

`func (o *CompanyAssetCollisionsRow) SetNSharedIndicationsNil(b bool)`

 SetNSharedIndicationsNil sets the value for NSharedIndications to be an explicit nil

### UnsetNSharedIndications
`func (o *CompanyAssetCollisionsRow) UnsetNSharedIndications()`

UnsetNSharedIndications ensures that no value is present for NSharedIndications, not even an explicit nil
### GetTopSharedIndications

`func (o *CompanyAssetCollisionsRow) GetTopSharedIndications() []string`

GetTopSharedIndications returns the TopSharedIndications field if non-nil, zero value otherwise.

### GetTopSharedIndicationsOk

`func (o *CompanyAssetCollisionsRow) GetTopSharedIndicationsOk() (*[]string, bool)`

GetTopSharedIndicationsOk returns a tuple with the TopSharedIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSharedIndications

`func (o *CompanyAssetCollisionsRow) SetTopSharedIndications(v []string)`

SetTopSharedIndications sets TopSharedIndications field to given value.

### HasTopSharedIndications

`func (o *CompanyAssetCollisionsRow) HasTopSharedIndications() bool`

HasTopSharedIndications returns a boolean if a field has been set.

### SetTopSharedIndicationsNil

`func (o *CompanyAssetCollisionsRow) SetTopSharedIndicationsNil(b bool)`

 SetTopSharedIndicationsNil sets the value for TopSharedIndications to be an explicit nil

### UnsetTopSharedIndications
`func (o *CompanyAssetCollisionsRow) UnsetTopSharedIndications()`

UnsetTopSharedIndications ensures that no value is present for TopSharedIndications, not even an explicit nil
### GetStage

`func (o *CompanyAssetCollisionsRow) GetStage() int32`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CompanyAssetCollisionsRow) GetStageOk() (*int32, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CompanyAssetCollisionsRow) SetStage(v int32)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CompanyAssetCollisionsRow) HasStage() bool`

HasStage returns a boolean if a field has been set.

### SetStageNil

`func (o *CompanyAssetCollisionsRow) SetStageNil(b bool)`

 SetStageNil sets the value for Stage to be an explicit nil

### UnsetStage
`func (o *CompanyAssetCollisionsRow) UnsetStage()`

UnsetStage ensures that no value is present for Stage, not even an explicit nil
### GetLastYear

`func (o *CompanyAssetCollisionsRow) GetLastYear() int32`

GetLastYear returns the LastYear field if non-nil, zero value otherwise.

### GetLastYearOk

`func (o *CompanyAssetCollisionsRow) GetLastYearOk() (*int32, bool)`

GetLastYearOk returns a tuple with the LastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYear

`func (o *CompanyAssetCollisionsRow) SetLastYear(v int32)`

SetLastYear sets LastYear field to given value.

### HasLastYear

`func (o *CompanyAssetCollisionsRow) HasLastYear() bool`

HasLastYear returns a boolean if a field has been set.

### SetLastYearNil

`func (o *CompanyAssetCollisionsRow) SetLastYearNil(b bool)`

 SetLastYearNil sets the value for LastYear to be an explicit nil

### UnsetLastYear
`func (o *CompanyAssetCollisionsRow) UnsetLastYear()`

UnsetLastYear ensures that no value is present for LastYear, not even an explicit nil
### GetNHeadToHeadTrials

`func (o *CompanyAssetCollisionsRow) GetNHeadToHeadTrials() int32`

GetNHeadToHeadTrials returns the NHeadToHeadTrials field if non-nil, zero value otherwise.

### GetNHeadToHeadTrialsOk

`func (o *CompanyAssetCollisionsRow) GetNHeadToHeadTrialsOk() (*int32, bool)`

GetNHeadToHeadTrialsOk returns a tuple with the NHeadToHeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNHeadToHeadTrials

`func (o *CompanyAssetCollisionsRow) SetNHeadToHeadTrials(v int32)`

SetNHeadToHeadTrials sets NHeadToHeadTrials field to given value.

### HasNHeadToHeadTrials

`func (o *CompanyAssetCollisionsRow) HasNHeadToHeadTrials() bool`

HasNHeadToHeadTrials returns a boolean if a field has been set.

### SetNHeadToHeadTrialsNil

`func (o *CompanyAssetCollisionsRow) SetNHeadToHeadTrialsNil(b bool)`

 SetNHeadToHeadTrialsNil sets the value for NHeadToHeadTrials to be an explicit nil

### UnsetNHeadToHeadTrials
`func (o *CompanyAssetCollisionsRow) UnsetNHeadToHeadTrials()`

UnsetNHeadToHeadTrials ensures that no value is present for NHeadToHeadTrials, not even an explicit nil
### GetOwnershipBasis

`func (o *CompanyAssetCollisionsRow) GetOwnershipBasis() string`

GetOwnershipBasis returns the OwnershipBasis field if non-nil, zero value otherwise.

### GetOwnershipBasisOk

`func (o *CompanyAssetCollisionsRow) GetOwnershipBasisOk() (*string, bool)`

GetOwnershipBasisOk returns a tuple with the OwnershipBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipBasis

`func (o *CompanyAssetCollisionsRow) SetOwnershipBasis(v string)`

SetOwnershipBasis sets OwnershipBasis field to given value.

### HasOwnershipBasis

`func (o *CompanyAssetCollisionsRow) HasOwnershipBasis() bool`

HasOwnershipBasis returns a boolean if a field has been set.

### SetOwnershipBasisNil

`func (o *CompanyAssetCollisionsRow) SetOwnershipBasisNil(b bool)`

 SetOwnershipBasisNil sets the value for OwnershipBasis to be an explicit nil

### UnsetOwnershipBasis
`func (o *CompanyAssetCollisionsRow) UnsetOwnershipBasis()`

UnsetOwnershipBasis ensures that no value is present for OwnershipBasis, not even an explicit nil
### GetCompetitorOwnershipBasis

`func (o *CompanyAssetCollisionsRow) GetCompetitorOwnershipBasis() string`

GetCompetitorOwnershipBasis returns the CompetitorOwnershipBasis field if non-nil, zero value otherwise.

### GetCompetitorOwnershipBasisOk

`func (o *CompanyAssetCollisionsRow) GetCompetitorOwnershipBasisOk() (*string, bool)`

GetCompetitorOwnershipBasisOk returns a tuple with the CompetitorOwnershipBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorOwnershipBasis

`func (o *CompanyAssetCollisionsRow) SetCompetitorOwnershipBasis(v string)`

SetCompetitorOwnershipBasis sets CompetitorOwnershipBasis field to given value.

### HasCompetitorOwnershipBasis

`func (o *CompanyAssetCollisionsRow) HasCompetitorOwnershipBasis() bool`

HasCompetitorOwnershipBasis returns a boolean if a field has been set.

### SetCompetitorOwnershipBasisNil

`func (o *CompanyAssetCollisionsRow) SetCompetitorOwnershipBasisNil(b bool)`

 SetCompetitorOwnershipBasisNil sets the value for CompetitorOwnershipBasis to be an explicit nil

### UnsetCompetitorOwnershipBasis
`func (o *CompanyAssetCollisionsRow) UnsetCompetitorOwnershipBasis()`

UnsetCompetitorOwnershipBasis ensures that no value is present for CompetitorOwnershipBasis, not even an explicit nil
### GetFamilyIdf

`func (o *CompanyAssetCollisionsRow) GetFamilyIdf() float32`

GetFamilyIdf returns the FamilyIdf field if non-nil, zero value otherwise.

### GetFamilyIdfOk

`func (o *CompanyAssetCollisionsRow) GetFamilyIdfOk() (*float32, bool)`

GetFamilyIdfOk returns a tuple with the FamilyIdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyIdf

`func (o *CompanyAssetCollisionsRow) SetFamilyIdf(v float32)`

SetFamilyIdf sets FamilyIdf field to given value.

### HasFamilyIdf

`func (o *CompanyAssetCollisionsRow) HasFamilyIdf() bool`

HasFamilyIdf returns a boolean if a field has been set.

### SetFamilyIdfNil

`func (o *CompanyAssetCollisionsRow) SetFamilyIdfNil(b bool)`

 SetFamilyIdfNil sets the value for FamilyIdf to be an explicit nil

### UnsetFamilyIdf
`func (o *CompanyAssetCollisionsRow) UnsetFamilyIdf()`

UnsetFamilyIdf ensures that no value is present for FamilyIdf, not even an explicit nil
### GetCollisionWeight

`func (o *CompanyAssetCollisionsRow) GetCollisionWeight() float32`

GetCollisionWeight returns the CollisionWeight field if non-nil, zero value otherwise.

### GetCollisionWeightOk

`func (o *CompanyAssetCollisionsRow) GetCollisionWeightOk() (*float32, bool)`

GetCollisionWeightOk returns a tuple with the CollisionWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollisionWeight

`func (o *CompanyAssetCollisionsRow) SetCollisionWeight(v float32)`

SetCollisionWeight sets CollisionWeight field to given value.

### HasCollisionWeight

`func (o *CompanyAssetCollisionsRow) HasCollisionWeight() bool`

HasCollisionWeight returns a boolean if a field has been set.

### SetCollisionWeightNil

`func (o *CompanyAssetCollisionsRow) SetCollisionWeightNil(b bool)`

 SetCollisionWeightNil sets the value for CollisionWeight to be an explicit nil

### UnsetCollisionWeight
`func (o *CompanyAssetCollisionsRow) UnsetCollisionWeight()`

UnsetCollisionWeight ensures that no value is present for CollisionWeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


