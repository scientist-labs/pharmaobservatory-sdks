# CompanyMarketCollisionsRow

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
**ConditionUi** | Pointer to **NullableString** | The shared MeSH indication (a MeSH descriptor id): both assets are owned, live and Phase 3 or 4 in it. | [optional] 
**ConditionName** | Pointer to **NullableString** | The MeSH name of the shared indication. | [optional] 
**TargetSymbols** | Pointer to **[]string** | Gene symbols of the company drug&#39;s recorded targets. Never empty, and none is a target of the competitor drug. | [optional] 
**CompetitorTargetSymbols** | Pointer to **[]string** | Gene symbols of the competitor drug&#39;s recorded targets. | [optional] 
**Stage** | Pointer to **NullableInt32** | The less advanced of the two programs&#39; furthest phase: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1, -1 phase not applicable. | [optional] 
**LastYear** | Pointer to **NullableInt32** | The later of the two programs&#39; last trial start year. | [optional] 
**NHeadToHeadTrials** | Pointer to **NullableInt32** | Trials led by either company in which one drug is experimental and the other an active comparator. | [optional] 
**OwnershipBasis** | Pointer to **NullableString** | Why the company owns the drug. experimental_arm and regulatory_holder rest on evidence; inferred_lead is an inference (no arm data says otherwise). | [optional] 
**CompetitorOwnershipBasis** | Pointer to **NullableString** | Why the competitor owns its drug, same vocabulary as ownership_basis. | [optional] 
**IndicationIdf** | Pointer to **NullableFloat32** | The rarity weight of the indication: the natural log of (companies holding a live, owned Phase 3+ asset with a recorded target) over (companies holding one in this indication). | [optional] 
**MarketWeight** | Pointer to **NullableFloat32** | The weight of the pair, 0.5 times indication_idf times factors for recency and a head-to-head trial. There is no stage factor, both programs are Phase 3 or 4. | [optional] 

## Methods

### NewCompanyMarketCollisionsRow

`func NewCompanyMarketCollisionsRow() *CompanyMarketCollisionsRow`

NewCompanyMarketCollisionsRow instantiates a new CompanyMarketCollisionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyMarketCollisionsRowWithDefaults

`func NewCompanyMarketCollisionsRowWithDefaults() *CompanyMarketCollisionsRow`

NewCompanyMarketCollisionsRowWithDefaults instantiates a new CompanyMarketCollisionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyMarketCollisionsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyMarketCollisionsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyMarketCollisionsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyMarketCollisionsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyMarketCollisionsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyMarketCollisionsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyMarketCollisionsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyMarketCollisionsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyMarketCollisionsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyMarketCollisionsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyMarketCollisionsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyMarketCollisionsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyMarketCollisionsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyMarketCollisionsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyMarketCollisionsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyMarketCollisionsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyMarketCollisionsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyMarketCollisionsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyMarketCollisionsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyMarketCollisionsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyMarketCollisionsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyMarketCollisionsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyMarketCollisionsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyMarketCollisionsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetMoleculeId

`func (o *CompanyMarketCollisionsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *CompanyMarketCollisionsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *CompanyMarketCollisionsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *CompanyMarketCollisionsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *CompanyMarketCollisionsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *CompanyMarketCollisionsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetMoleculeName

`func (o *CompanyMarketCollisionsRow) GetMoleculeName() string`

GetMoleculeName returns the MoleculeName field if non-nil, zero value otherwise.

### GetMoleculeNameOk

`func (o *CompanyMarketCollisionsRow) GetMoleculeNameOk() (*string, bool)`

GetMoleculeNameOk returns a tuple with the MoleculeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeName

`func (o *CompanyMarketCollisionsRow) SetMoleculeName(v string)`

SetMoleculeName sets MoleculeName field to given value.

### HasMoleculeName

`func (o *CompanyMarketCollisionsRow) HasMoleculeName() bool`

HasMoleculeName returns a boolean if a field has been set.

### SetMoleculeNameNil

`func (o *CompanyMarketCollisionsRow) SetMoleculeNameNil(b bool)`

 SetMoleculeNameNil sets the value for MoleculeName to be an explicit nil

### UnsetMoleculeName
`func (o *CompanyMarketCollisionsRow) UnsetMoleculeName()`

UnsetMoleculeName ensures that no value is present for MoleculeName, not even an explicit nil
### GetCompetitorCompanyId

`func (o *CompanyMarketCollisionsRow) GetCompetitorCompanyId() string`

GetCompetitorCompanyId returns the CompetitorCompanyId field if non-nil, zero value otherwise.

### GetCompetitorCompanyIdOk

`func (o *CompanyMarketCollisionsRow) GetCompetitorCompanyIdOk() (*string, bool)`

GetCompetitorCompanyIdOk returns a tuple with the CompetitorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorCompanyId

`func (o *CompanyMarketCollisionsRow) SetCompetitorCompanyId(v string)`

SetCompetitorCompanyId sets CompetitorCompanyId field to given value.

### HasCompetitorCompanyId

`func (o *CompanyMarketCollisionsRow) HasCompetitorCompanyId() bool`

HasCompetitorCompanyId returns a boolean if a field has been set.

### SetCompetitorCompanyIdNil

`func (o *CompanyMarketCollisionsRow) SetCompetitorCompanyIdNil(b bool)`

 SetCompetitorCompanyIdNil sets the value for CompetitorCompanyId to be an explicit nil

### UnsetCompetitorCompanyId
`func (o *CompanyMarketCollisionsRow) UnsetCompetitorCompanyId()`

UnsetCompetitorCompanyId ensures that no value is present for CompetitorCompanyId, not even an explicit nil
### GetCompetitorCompanyName

`func (o *CompanyMarketCollisionsRow) GetCompetitorCompanyName() string`

GetCompetitorCompanyName returns the CompetitorCompanyName field if non-nil, zero value otherwise.

### GetCompetitorCompanyNameOk

`func (o *CompanyMarketCollisionsRow) GetCompetitorCompanyNameOk() (*string, bool)`

GetCompetitorCompanyNameOk returns a tuple with the CompetitorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorCompanyName

`func (o *CompanyMarketCollisionsRow) SetCompetitorCompanyName(v string)`

SetCompetitorCompanyName sets CompetitorCompanyName field to given value.

### HasCompetitorCompanyName

`func (o *CompanyMarketCollisionsRow) HasCompetitorCompanyName() bool`

HasCompetitorCompanyName returns a boolean if a field has been set.

### SetCompetitorCompanyNameNil

`func (o *CompanyMarketCollisionsRow) SetCompetitorCompanyNameNil(b bool)`

 SetCompetitorCompanyNameNil sets the value for CompetitorCompanyName to be an explicit nil

### UnsetCompetitorCompanyName
`func (o *CompanyMarketCollisionsRow) UnsetCompetitorCompanyName()`

UnsetCompetitorCompanyName ensures that no value is present for CompetitorCompanyName, not even an explicit nil
### GetCompetitorGroupId

`func (o *CompanyMarketCollisionsRow) GetCompetitorGroupId() string`

GetCompetitorGroupId returns the CompetitorGroupId field if non-nil, zero value otherwise.

### GetCompetitorGroupIdOk

`func (o *CompanyMarketCollisionsRow) GetCompetitorGroupIdOk() (*string, bool)`

GetCompetitorGroupIdOk returns a tuple with the CompetitorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorGroupId

`func (o *CompanyMarketCollisionsRow) SetCompetitorGroupId(v string)`

SetCompetitorGroupId sets CompetitorGroupId field to given value.

### HasCompetitorGroupId

`func (o *CompanyMarketCollisionsRow) HasCompetitorGroupId() bool`

HasCompetitorGroupId returns a boolean if a field has been set.

### SetCompetitorGroupIdNil

`func (o *CompanyMarketCollisionsRow) SetCompetitorGroupIdNil(b bool)`

 SetCompetitorGroupIdNil sets the value for CompetitorGroupId to be an explicit nil

### UnsetCompetitorGroupId
`func (o *CompanyMarketCollisionsRow) UnsetCompetitorGroupId()`

UnsetCompetitorGroupId ensures that no value is present for CompetitorGroupId, not even an explicit nil
### GetCompetitorGroupName

`func (o *CompanyMarketCollisionsRow) GetCompetitorGroupName() string`

GetCompetitorGroupName returns the CompetitorGroupName field if non-nil, zero value otherwise.

### GetCompetitorGroupNameOk

`func (o *CompanyMarketCollisionsRow) GetCompetitorGroupNameOk() (*string, bool)`

GetCompetitorGroupNameOk returns a tuple with the CompetitorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorGroupName

`func (o *CompanyMarketCollisionsRow) SetCompetitorGroupName(v string)`

SetCompetitorGroupName sets CompetitorGroupName field to given value.

### HasCompetitorGroupName

`func (o *CompanyMarketCollisionsRow) HasCompetitorGroupName() bool`

HasCompetitorGroupName returns a boolean if a field has been set.

### SetCompetitorGroupNameNil

`func (o *CompanyMarketCollisionsRow) SetCompetitorGroupNameNil(b bool)`

 SetCompetitorGroupNameNil sets the value for CompetitorGroupName to be an explicit nil

### UnsetCompetitorGroupName
`func (o *CompanyMarketCollisionsRow) UnsetCompetitorGroupName()`

UnsetCompetitorGroupName ensures that no value is present for CompetitorGroupName, not even an explicit nil
### GetCompetitorMoleculeId

`func (o *CompanyMarketCollisionsRow) GetCompetitorMoleculeId() string`

GetCompetitorMoleculeId returns the CompetitorMoleculeId field if non-nil, zero value otherwise.

### GetCompetitorMoleculeIdOk

`func (o *CompanyMarketCollisionsRow) GetCompetitorMoleculeIdOk() (*string, bool)`

GetCompetitorMoleculeIdOk returns a tuple with the CompetitorMoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorMoleculeId

`func (o *CompanyMarketCollisionsRow) SetCompetitorMoleculeId(v string)`

SetCompetitorMoleculeId sets CompetitorMoleculeId field to given value.

### HasCompetitorMoleculeId

`func (o *CompanyMarketCollisionsRow) HasCompetitorMoleculeId() bool`

HasCompetitorMoleculeId returns a boolean if a field has been set.

### SetCompetitorMoleculeIdNil

`func (o *CompanyMarketCollisionsRow) SetCompetitorMoleculeIdNil(b bool)`

 SetCompetitorMoleculeIdNil sets the value for CompetitorMoleculeId to be an explicit nil

### UnsetCompetitorMoleculeId
`func (o *CompanyMarketCollisionsRow) UnsetCompetitorMoleculeId()`

UnsetCompetitorMoleculeId ensures that no value is present for CompetitorMoleculeId, not even an explicit nil
### GetCompetitorMoleculeName

`func (o *CompanyMarketCollisionsRow) GetCompetitorMoleculeName() string`

GetCompetitorMoleculeName returns the CompetitorMoleculeName field if non-nil, zero value otherwise.

### GetCompetitorMoleculeNameOk

`func (o *CompanyMarketCollisionsRow) GetCompetitorMoleculeNameOk() (*string, bool)`

GetCompetitorMoleculeNameOk returns a tuple with the CompetitorMoleculeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorMoleculeName

`func (o *CompanyMarketCollisionsRow) SetCompetitorMoleculeName(v string)`

SetCompetitorMoleculeName sets CompetitorMoleculeName field to given value.

### HasCompetitorMoleculeName

`func (o *CompanyMarketCollisionsRow) HasCompetitorMoleculeName() bool`

HasCompetitorMoleculeName returns a boolean if a field has been set.

### SetCompetitorMoleculeNameNil

`func (o *CompanyMarketCollisionsRow) SetCompetitorMoleculeNameNil(b bool)`

 SetCompetitorMoleculeNameNil sets the value for CompetitorMoleculeName to be an explicit nil

### UnsetCompetitorMoleculeName
`func (o *CompanyMarketCollisionsRow) UnsetCompetitorMoleculeName()`

UnsetCompetitorMoleculeName ensures that no value is present for CompetitorMoleculeName, not even an explicit nil
### GetConditionUi

`func (o *CompanyMarketCollisionsRow) GetConditionUi() string`

GetConditionUi returns the ConditionUi field if non-nil, zero value otherwise.

### GetConditionUiOk

`func (o *CompanyMarketCollisionsRow) GetConditionUiOk() (*string, bool)`

GetConditionUiOk returns a tuple with the ConditionUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionUi

`func (o *CompanyMarketCollisionsRow) SetConditionUi(v string)`

SetConditionUi sets ConditionUi field to given value.

### HasConditionUi

`func (o *CompanyMarketCollisionsRow) HasConditionUi() bool`

HasConditionUi returns a boolean if a field has been set.

### SetConditionUiNil

`func (o *CompanyMarketCollisionsRow) SetConditionUiNil(b bool)`

 SetConditionUiNil sets the value for ConditionUi to be an explicit nil

### UnsetConditionUi
`func (o *CompanyMarketCollisionsRow) UnsetConditionUi()`

UnsetConditionUi ensures that no value is present for ConditionUi, not even an explicit nil
### GetConditionName

`func (o *CompanyMarketCollisionsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *CompanyMarketCollisionsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *CompanyMarketCollisionsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *CompanyMarketCollisionsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *CompanyMarketCollisionsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *CompanyMarketCollisionsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetTargetSymbols

`func (o *CompanyMarketCollisionsRow) GetTargetSymbols() []string`

GetTargetSymbols returns the TargetSymbols field if non-nil, zero value otherwise.

### GetTargetSymbolsOk

`func (o *CompanyMarketCollisionsRow) GetTargetSymbolsOk() (*[]string, bool)`

GetTargetSymbolsOk returns a tuple with the TargetSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbols

`func (o *CompanyMarketCollisionsRow) SetTargetSymbols(v []string)`

SetTargetSymbols sets TargetSymbols field to given value.

### HasTargetSymbols

`func (o *CompanyMarketCollisionsRow) HasTargetSymbols() bool`

HasTargetSymbols returns a boolean if a field has been set.

### SetTargetSymbolsNil

`func (o *CompanyMarketCollisionsRow) SetTargetSymbolsNil(b bool)`

 SetTargetSymbolsNil sets the value for TargetSymbols to be an explicit nil

### UnsetTargetSymbols
`func (o *CompanyMarketCollisionsRow) UnsetTargetSymbols()`

UnsetTargetSymbols ensures that no value is present for TargetSymbols, not even an explicit nil
### GetCompetitorTargetSymbols

`func (o *CompanyMarketCollisionsRow) GetCompetitorTargetSymbols() []string`

GetCompetitorTargetSymbols returns the CompetitorTargetSymbols field if non-nil, zero value otherwise.

### GetCompetitorTargetSymbolsOk

`func (o *CompanyMarketCollisionsRow) GetCompetitorTargetSymbolsOk() (*[]string, bool)`

GetCompetitorTargetSymbolsOk returns a tuple with the CompetitorTargetSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorTargetSymbols

`func (o *CompanyMarketCollisionsRow) SetCompetitorTargetSymbols(v []string)`

SetCompetitorTargetSymbols sets CompetitorTargetSymbols field to given value.

### HasCompetitorTargetSymbols

`func (o *CompanyMarketCollisionsRow) HasCompetitorTargetSymbols() bool`

HasCompetitorTargetSymbols returns a boolean if a field has been set.

### SetCompetitorTargetSymbolsNil

`func (o *CompanyMarketCollisionsRow) SetCompetitorTargetSymbolsNil(b bool)`

 SetCompetitorTargetSymbolsNil sets the value for CompetitorTargetSymbols to be an explicit nil

### UnsetCompetitorTargetSymbols
`func (o *CompanyMarketCollisionsRow) UnsetCompetitorTargetSymbols()`

UnsetCompetitorTargetSymbols ensures that no value is present for CompetitorTargetSymbols, not even an explicit nil
### GetStage

`func (o *CompanyMarketCollisionsRow) GetStage() int32`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CompanyMarketCollisionsRow) GetStageOk() (*int32, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CompanyMarketCollisionsRow) SetStage(v int32)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CompanyMarketCollisionsRow) HasStage() bool`

HasStage returns a boolean if a field has been set.

### SetStageNil

`func (o *CompanyMarketCollisionsRow) SetStageNil(b bool)`

 SetStageNil sets the value for Stage to be an explicit nil

### UnsetStage
`func (o *CompanyMarketCollisionsRow) UnsetStage()`

UnsetStage ensures that no value is present for Stage, not even an explicit nil
### GetLastYear

`func (o *CompanyMarketCollisionsRow) GetLastYear() int32`

GetLastYear returns the LastYear field if non-nil, zero value otherwise.

### GetLastYearOk

`func (o *CompanyMarketCollisionsRow) GetLastYearOk() (*int32, bool)`

GetLastYearOk returns a tuple with the LastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYear

`func (o *CompanyMarketCollisionsRow) SetLastYear(v int32)`

SetLastYear sets LastYear field to given value.

### HasLastYear

`func (o *CompanyMarketCollisionsRow) HasLastYear() bool`

HasLastYear returns a boolean if a field has been set.

### SetLastYearNil

`func (o *CompanyMarketCollisionsRow) SetLastYearNil(b bool)`

 SetLastYearNil sets the value for LastYear to be an explicit nil

### UnsetLastYear
`func (o *CompanyMarketCollisionsRow) UnsetLastYear()`

UnsetLastYear ensures that no value is present for LastYear, not even an explicit nil
### GetNHeadToHeadTrials

`func (o *CompanyMarketCollisionsRow) GetNHeadToHeadTrials() int32`

GetNHeadToHeadTrials returns the NHeadToHeadTrials field if non-nil, zero value otherwise.

### GetNHeadToHeadTrialsOk

`func (o *CompanyMarketCollisionsRow) GetNHeadToHeadTrialsOk() (*int32, bool)`

GetNHeadToHeadTrialsOk returns a tuple with the NHeadToHeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNHeadToHeadTrials

`func (o *CompanyMarketCollisionsRow) SetNHeadToHeadTrials(v int32)`

SetNHeadToHeadTrials sets NHeadToHeadTrials field to given value.

### HasNHeadToHeadTrials

`func (o *CompanyMarketCollisionsRow) HasNHeadToHeadTrials() bool`

HasNHeadToHeadTrials returns a boolean if a field has been set.

### SetNHeadToHeadTrialsNil

`func (o *CompanyMarketCollisionsRow) SetNHeadToHeadTrialsNil(b bool)`

 SetNHeadToHeadTrialsNil sets the value for NHeadToHeadTrials to be an explicit nil

### UnsetNHeadToHeadTrials
`func (o *CompanyMarketCollisionsRow) UnsetNHeadToHeadTrials()`

UnsetNHeadToHeadTrials ensures that no value is present for NHeadToHeadTrials, not even an explicit nil
### GetOwnershipBasis

`func (o *CompanyMarketCollisionsRow) GetOwnershipBasis() string`

GetOwnershipBasis returns the OwnershipBasis field if non-nil, zero value otherwise.

### GetOwnershipBasisOk

`func (o *CompanyMarketCollisionsRow) GetOwnershipBasisOk() (*string, bool)`

GetOwnershipBasisOk returns a tuple with the OwnershipBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipBasis

`func (o *CompanyMarketCollisionsRow) SetOwnershipBasis(v string)`

SetOwnershipBasis sets OwnershipBasis field to given value.

### HasOwnershipBasis

`func (o *CompanyMarketCollisionsRow) HasOwnershipBasis() bool`

HasOwnershipBasis returns a boolean if a field has been set.

### SetOwnershipBasisNil

`func (o *CompanyMarketCollisionsRow) SetOwnershipBasisNil(b bool)`

 SetOwnershipBasisNil sets the value for OwnershipBasis to be an explicit nil

### UnsetOwnershipBasis
`func (o *CompanyMarketCollisionsRow) UnsetOwnershipBasis()`

UnsetOwnershipBasis ensures that no value is present for OwnershipBasis, not even an explicit nil
### GetCompetitorOwnershipBasis

`func (o *CompanyMarketCollisionsRow) GetCompetitorOwnershipBasis() string`

GetCompetitorOwnershipBasis returns the CompetitorOwnershipBasis field if non-nil, zero value otherwise.

### GetCompetitorOwnershipBasisOk

`func (o *CompanyMarketCollisionsRow) GetCompetitorOwnershipBasisOk() (*string, bool)`

GetCompetitorOwnershipBasisOk returns a tuple with the CompetitorOwnershipBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorOwnershipBasis

`func (o *CompanyMarketCollisionsRow) SetCompetitorOwnershipBasis(v string)`

SetCompetitorOwnershipBasis sets CompetitorOwnershipBasis field to given value.

### HasCompetitorOwnershipBasis

`func (o *CompanyMarketCollisionsRow) HasCompetitorOwnershipBasis() bool`

HasCompetitorOwnershipBasis returns a boolean if a field has been set.

### SetCompetitorOwnershipBasisNil

`func (o *CompanyMarketCollisionsRow) SetCompetitorOwnershipBasisNil(b bool)`

 SetCompetitorOwnershipBasisNil sets the value for CompetitorOwnershipBasis to be an explicit nil

### UnsetCompetitorOwnershipBasis
`func (o *CompanyMarketCollisionsRow) UnsetCompetitorOwnershipBasis()`

UnsetCompetitorOwnershipBasis ensures that no value is present for CompetitorOwnershipBasis, not even an explicit nil
### GetIndicationIdf

`func (o *CompanyMarketCollisionsRow) GetIndicationIdf() float32`

GetIndicationIdf returns the IndicationIdf field if non-nil, zero value otherwise.

### GetIndicationIdfOk

`func (o *CompanyMarketCollisionsRow) GetIndicationIdfOk() (*float32, bool)`

GetIndicationIdfOk returns a tuple with the IndicationIdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicationIdf

`func (o *CompanyMarketCollisionsRow) SetIndicationIdf(v float32)`

SetIndicationIdf sets IndicationIdf field to given value.

### HasIndicationIdf

`func (o *CompanyMarketCollisionsRow) HasIndicationIdf() bool`

HasIndicationIdf returns a boolean if a field has been set.

### SetIndicationIdfNil

`func (o *CompanyMarketCollisionsRow) SetIndicationIdfNil(b bool)`

 SetIndicationIdfNil sets the value for IndicationIdf to be an explicit nil

### UnsetIndicationIdf
`func (o *CompanyMarketCollisionsRow) UnsetIndicationIdf()`

UnsetIndicationIdf ensures that no value is present for IndicationIdf, not even an explicit nil
### GetMarketWeight

`func (o *CompanyMarketCollisionsRow) GetMarketWeight() float32`

GetMarketWeight returns the MarketWeight field if non-nil, zero value otherwise.

### GetMarketWeightOk

`func (o *CompanyMarketCollisionsRow) GetMarketWeightOk() (*float32, bool)`

GetMarketWeightOk returns a tuple with the MarketWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketWeight

`func (o *CompanyMarketCollisionsRow) SetMarketWeight(v float32)`

SetMarketWeight sets MarketWeight field to given value.

### HasMarketWeight

`func (o *CompanyMarketCollisionsRow) HasMarketWeight() bool`

HasMarketWeight returns a boolean if a field has been set.

### SetMarketWeightNil

`func (o *CompanyMarketCollisionsRow) SetMarketWeightNil(b bool)`

 SetMarketWeightNil sets the value for MarketWeight to be an explicit nil

### UnsetMarketWeight
`func (o *CompanyMarketCollisionsRow) UnsetMarketWeight()`

UnsetMarketWeight ensures that no value is present for MarketWeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


