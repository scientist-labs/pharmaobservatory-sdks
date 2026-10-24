# CompanyCompetitorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**CompetitorCompanyId** | Pointer to **NullableString** | Company level: the competing company. | [optional] 
**CompetitorCompanyName** | Pointer to **NullableString** | Display name of the competitor&#39;s company. | [optional] 
**CompetitorGroupId** | Pointer to **NullableString** | Group level: the competitor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**CompetitorGroupName** | Pointer to **NullableString** | Display name of the competitor&#39;s group. | [optional] 
**CompetitorRank** | Pointer to **NullableInt32** | The competitor&#39;s rank for the company by competitor_score (1 is the highest). | [optional] 
**CompetitorScore** | Pointer to **NullableFloat32** | 100 times the weighted mean of the four similarities below, rounded to four places. | [optional] 
**TargetSim** | Pointer to **NullableFloat32** | Jaccard similarity, 0 to 1, of the two companies&#39; molecular target sets. | [optional] 
**MechanismSim** | Pointer to **NullableFloat32** | Jaccard similarity, 0 to 1, of their mechanism sets (a target and its action). | [optional] 
**IndicationSim** | Pointer to **NullableFloat32** | Jaccard similarity, 0 to 1, of their top-25 disease area sets. | [optional] 
**PhaseSim** | Pointer to **NullableFloat32** | The fraction, 0 to 1, of the shared targets that both companies have taken to phase 2 or later. | [optional] 
**NSharedTargets** | Pointer to **NullableInt32** | Molecular targets both companies pursue. | [optional] 
**NSharedMechanisms** | Pointer to **NullableInt32** | Mechanisms (target and action) both companies pursue. | [optional] 
**NSharedIndications** | Pointer to **NullableInt32** | Disease areas both companies are in, within each one&#39;s top 25. | [optional] 
**NPhaseOverlapTargets** | Pointer to **NullableInt32** | Shared targets that both companies have taken to phase 2 or later. | [optional] 
**NCompanyTargets** | Pointer to **NullableInt32** | Targets the company pursues (the Jaccard denominator). | [optional] 
**NCompetitorTargets** | Pointer to **NullableInt32** | Targets the competitor pursues. | [optional] 
**NCompanyIndications** | Pointer to **NullableInt32** | Disease areas in the company&#39;s top 25. | [optional] 
**NCompetitorIndications** | Pointer to **NullableInt32** | Disease areas in the competitor&#39;s top 25. | [optional] 
**TopSharedTargets** | Pointer to **[]string** | Up to 5 shared target gene symbols, by combined phase then combined trials. An empty list when none. | [optional] 
**TopSharedIndications** | Pointer to **[]string** | Up to 5 shared disease area names, by combined trial count. An empty list when none. | [optional] 
**TopSharedMechanisms** | Pointer to **[]string** | Up to 3 shared mechanisms written as a gene symbol and an action, by combined phase then combined trials. An empty list when none. | [optional] 

## Methods

### NewCompanyCompetitorsRow

`func NewCompanyCompetitorsRow() *CompanyCompetitorsRow`

NewCompanyCompetitorsRow instantiates a new CompanyCompetitorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyCompetitorsRowWithDefaults

`func NewCompanyCompetitorsRowWithDefaults() *CompanyCompetitorsRow`

NewCompanyCompetitorsRowWithDefaults instantiates a new CompanyCompetitorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CompanyCompetitorsRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyCompetitorsRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyCompetitorsRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyCompetitorsRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyCompetitorsRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyCompetitorsRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyCompetitorsRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyCompetitorsRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyCompetitorsRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyCompetitorsRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyCompetitorsRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyCompetitorsRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyCompetitorsRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyCompetitorsRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyCompetitorsRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyCompetitorsRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyCompetitorsRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyCompetitorsRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyCompetitorsRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyCompetitorsRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyCompetitorsRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyCompetitorsRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyCompetitorsRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyCompetitorsRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetCompetitorCompanyId

`func (o *CompanyCompetitorsRow) GetCompetitorCompanyId() string`

GetCompetitorCompanyId returns the CompetitorCompanyId field if non-nil, zero value otherwise.

### GetCompetitorCompanyIdOk

`func (o *CompanyCompetitorsRow) GetCompetitorCompanyIdOk() (*string, bool)`

GetCompetitorCompanyIdOk returns a tuple with the CompetitorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorCompanyId

`func (o *CompanyCompetitorsRow) SetCompetitorCompanyId(v string)`

SetCompetitorCompanyId sets CompetitorCompanyId field to given value.

### HasCompetitorCompanyId

`func (o *CompanyCompetitorsRow) HasCompetitorCompanyId() bool`

HasCompetitorCompanyId returns a boolean if a field has been set.

### SetCompetitorCompanyIdNil

`func (o *CompanyCompetitorsRow) SetCompetitorCompanyIdNil(b bool)`

 SetCompetitorCompanyIdNil sets the value for CompetitorCompanyId to be an explicit nil

### UnsetCompetitorCompanyId
`func (o *CompanyCompetitorsRow) UnsetCompetitorCompanyId()`

UnsetCompetitorCompanyId ensures that no value is present for CompetitorCompanyId, not even an explicit nil
### GetCompetitorCompanyName

`func (o *CompanyCompetitorsRow) GetCompetitorCompanyName() string`

GetCompetitorCompanyName returns the CompetitorCompanyName field if non-nil, zero value otherwise.

### GetCompetitorCompanyNameOk

`func (o *CompanyCompetitorsRow) GetCompetitorCompanyNameOk() (*string, bool)`

GetCompetitorCompanyNameOk returns a tuple with the CompetitorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorCompanyName

`func (o *CompanyCompetitorsRow) SetCompetitorCompanyName(v string)`

SetCompetitorCompanyName sets CompetitorCompanyName field to given value.

### HasCompetitorCompanyName

`func (o *CompanyCompetitorsRow) HasCompetitorCompanyName() bool`

HasCompetitorCompanyName returns a boolean if a field has been set.

### SetCompetitorCompanyNameNil

`func (o *CompanyCompetitorsRow) SetCompetitorCompanyNameNil(b bool)`

 SetCompetitorCompanyNameNil sets the value for CompetitorCompanyName to be an explicit nil

### UnsetCompetitorCompanyName
`func (o *CompanyCompetitorsRow) UnsetCompetitorCompanyName()`

UnsetCompetitorCompanyName ensures that no value is present for CompetitorCompanyName, not even an explicit nil
### GetCompetitorGroupId

`func (o *CompanyCompetitorsRow) GetCompetitorGroupId() string`

GetCompetitorGroupId returns the CompetitorGroupId field if non-nil, zero value otherwise.

### GetCompetitorGroupIdOk

`func (o *CompanyCompetitorsRow) GetCompetitorGroupIdOk() (*string, bool)`

GetCompetitorGroupIdOk returns a tuple with the CompetitorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorGroupId

`func (o *CompanyCompetitorsRow) SetCompetitorGroupId(v string)`

SetCompetitorGroupId sets CompetitorGroupId field to given value.

### HasCompetitorGroupId

`func (o *CompanyCompetitorsRow) HasCompetitorGroupId() bool`

HasCompetitorGroupId returns a boolean if a field has been set.

### SetCompetitorGroupIdNil

`func (o *CompanyCompetitorsRow) SetCompetitorGroupIdNil(b bool)`

 SetCompetitorGroupIdNil sets the value for CompetitorGroupId to be an explicit nil

### UnsetCompetitorGroupId
`func (o *CompanyCompetitorsRow) UnsetCompetitorGroupId()`

UnsetCompetitorGroupId ensures that no value is present for CompetitorGroupId, not even an explicit nil
### GetCompetitorGroupName

`func (o *CompanyCompetitorsRow) GetCompetitorGroupName() string`

GetCompetitorGroupName returns the CompetitorGroupName field if non-nil, zero value otherwise.

### GetCompetitorGroupNameOk

`func (o *CompanyCompetitorsRow) GetCompetitorGroupNameOk() (*string, bool)`

GetCompetitorGroupNameOk returns a tuple with the CompetitorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorGroupName

`func (o *CompanyCompetitorsRow) SetCompetitorGroupName(v string)`

SetCompetitorGroupName sets CompetitorGroupName field to given value.

### HasCompetitorGroupName

`func (o *CompanyCompetitorsRow) HasCompetitorGroupName() bool`

HasCompetitorGroupName returns a boolean if a field has been set.

### SetCompetitorGroupNameNil

`func (o *CompanyCompetitorsRow) SetCompetitorGroupNameNil(b bool)`

 SetCompetitorGroupNameNil sets the value for CompetitorGroupName to be an explicit nil

### UnsetCompetitorGroupName
`func (o *CompanyCompetitorsRow) UnsetCompetitorGroupName()`

UnsetCompetitorGroupName ensures that no value is present for CompetitorGroupName, not even an explicit nil
### GetCompetitorRank

`func (o *CompanyCompetitorsRow) GetCompetitorRank() int32`

GetCompetitorRank returns the CompetitorRank field if non-nil, zero value otherwise.

### GetCompetitorRankOk

`func (o *CompanyCompetitorsRow) GetCompetitorRankOk() (*int32, bool)`

GetCompetitorRankOk returns a tuple with the CompetitorRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorRank

`func (o *CompanyCompetitorsRow) SetCompetitorRank(v int32)`

SetCompetitorRank sets CompetitorRank field to given value.

### HasCompetitorRank

`func (o *CompanyCompetitorsRow) HasCompetitorRank() bool`

HasCompetitorRank returns a boolean if a field has been set.

### SetCompetitorRankNil

`func (o *CompanyCompetitorsRow) SetCompetitorRankNil(b bool)`

 SetCompetitorRankNil sets the value for CompetitorRank to be an explicit nil

### UnsetCompetitorRank
`func (o *CompanyCompetitorsRow) UnsetCompetitorRank()`

UnsetCompetitorRank ensures that no value is present for CompetitorRank, not even an explicit nil
### GetCompetitorScore

`func (o *CompanyCompetitorsRow) GetCompetitorScore() float32`

GetCompetitorScore returns the CompetitorScore field if non-nil, zero value otherwise.

### GetCompetitorScoreOk

`func (o *CompanyCompetitorsRow) GetCompetitorScoreOk() (*float32, bool)`

GetCompetitorScoreOk returns a tuple with the CompetitorScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorScore

`func (o *CompanyCompetitorsRow) SetCompetitorScore(v float32)`

SetCompetitorScore sets CompetitorScore field to given value.

### HasCompetitorScore

`func (o *CompanyCompetitorsRow) HasCompetitorScore() bool`

HasCompetitorScore returns a boolean if a field has been set.

### SetCompetitorScoreNil

`func (o *CompanyCompetitorsRow) SetCompetitorScoreNil(b bool)`

 SetCompetitorScoreNil sets the value for CompetitorScore to be an explicit nil

### UnsetCompetitorScore
`func (o *CompanyCompetitorsRow) UnsetCompetitorScore()`

UnsetCompetitorScore ensures that no value is present for CompetitorScore, not even an explicit nil
### GetTargetSim

`func (o *CompanyCompetitorsRow) GetTargetSim() float32`

GetTargetSim returns the TargetSim field if non-nil, zero value otherwise.

### GetTargetSimOk

`func (o *CompanyCompetitorsRow) GetTargetSimOk() (*float32, bool)`

GetTargetSimOk returns a tuple with the TargetSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSim

`func (o *CompanyCompetitorsRow) SetTargetSim(v float32)`

SetTargetSim sets TargetSim field to given value.

### HasTargetSim

`func (o *CompanyCompetitorsRow) HasTargetSim() bool`

HasTargetSim returns a boolean if a field has been set.

### SetTargetSimNil

`func (o *CompanyCompetitorsRow) SetTargetSimNil(b bool)`

 SetTargetSimNil sets the value for TargetSim to be an explicit nil

### UnsetTargetSim
`func (o *CompanyCompetitorsRow) UnsetTargetSim()`

UnsetTargetSim ensures that no value is present for TargetSim, not even an explicit nil
### GetMechanismSim

`func (o *CompanyCompetitorsRow) GetMechanismSim() float32`

GetMechanismSim returns the MechanismSim field if non-nil, zero value otherwise.

### GetMechanismSimOk

`func (o *CompanyCompetitorsRow) GetMechanismSimOk() (*float32, bool)`

GetMechanismSimOk returns a tuple with the MechanismSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanismSim

`func (o *CompanyCompetitorsRow) SetMechanismSim(v float32)`

SetMechanismSim sets MechanismSim field to given value.

### HasMechanismSim

`func (o *CompanyCompetitorsRow) HasMechanismSim() bool`

HasMechanismSim returns a boolean if a field has been set.

### SetMechanismSimNil

`func (o *CompanyCompetitorsRow) SetMechanismSimNil(b bool)`

 SetMechanismSimNil sets the value for MechanismSim to be an explicit nil

### UnsetMechanismSim
`func (o *CompanyCompetitorsRow) UnsetMechanismSim()`

UnsetMechanismSim ensures that no value is present for MechanismSim, not even an explicit nil
### GetIndicationSim

`func (o *CompanyCompetitorsRow) GetIndicationSim() float32`

GetIndicationSim returns the IndicationSim field if non-nil, zero value otherwise.

### GetIndicationSimOk

`func (o *CompanyCompetitorsRow) GetIndicationSimOk() (*float32, bool)`

GetIndicationSimOk returns a tuple with the IndicationSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicationSim

`func (o *CompanyCompetitorsRow) SetIndicationSim(v float32)`

SetIndicationSim sets IndicationSim field to given value.

### HasIndicationSim

`func (o *CompanyCompetitorsRow) HasIndicationSim() bool`

HasIndicationSim returns a boolean if a field has been set.

### SetIndicationSimNil

`func (o *CompanyCompetitorsRow) SetIndicationSimNil(b bool)`

 SetIndicationSimNil sets the value for IndicationSim to be an explicit nil

### UnsetIndicationSim
`func (o *CompanyCompetitorsRow) UnsetIndicationSim()`

UnsetIndicationSim ensures that no value is present for IndicationSim, not even an explicit nil
### GetPhaseSim

`func (o *CompanyCompetitorsRow) GetPhaseSim() float32`

GetPhaseSim returns the PhaseSim field if non-nil, zero value otherwise.

### GetPhaseSimOk

`func (o *CompanyCompetitorsRow) GetPhaseSimOk() (*float32, bool)`

GetPhaseSimOk returns a tuple with the PhaseSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseSim

`func (o *CompanyCompetitorsRow) SetPhaseSim(v float32)`

SetPhaseSim sets PhaseSim field to given value.

### HasPhaseSim

`func (o *CompanyCompetitorsRow) HasPhaseSim() bool`

HasPhaseSim returns a boolean if a field has been set.

### SetPhaseSimNil

`func (o *CompanyCompetitorsRow) SetPhaseSimNil(b bool)`

 SetPhaseSimNil sets the value for PhaseSim to be an explicit nil

### UnsetPhaseSim
`func (o *CompanyCompetitorsRow) UnsetPhaseSim()`

UnsetPhaseSim ensures that no value is present for PhaseSim, not even an explicit nil
### GetNSharedTargets

`func (o *CompanyCompetitorsRow) GetNSharedTargets() int32`

GetNSharedTargets returns the NSharedTargets field if non-nil, zero value otherwise.

### GetNSharedTargetsOk

`func (o *CompanyCompetitorsRow) GetNSharedTargetsOk() (*int32, bool)`

GetNSharedTargetsOk returns a tuple with the NSharedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedTargets

`func (o *CompanyCompetitorsRow) SetNSharedTargets(v int32)`

SetNSharedTargets sets NSharedTargets field to given value.

### HasNSharedTargets

`func (o *CompanyCompetitorsRow) HasNSharedTargets() bool`

HasNSharedTargets returns a boolean if a field has been set.

### SetNSharedTargetsNil

`func (o *CompanyCompetitorsRow) SetNSharedTargetsNil(b bool)`

 SetNSharedTargetsNil sets the value for NSharedTargets to be an explicit nil

### UnsetNSharedTargets
`func (o *CompanyCompetitorsRow) UnsetNSharedTargets()`

UnsetNSharedTargets ensures that no value is present for NSharedTargets, not even an explicit nil
### GetNSharedMechanisms

`func (o *CompanyCompetitorsRow) GetNSharedMechanisms() int32`

GetNSharedMechanisms returns the NSharedMechanisms field if non-nil, zero value otherwise.

### GetNSharedMechanismsOk

`func (o *CompanyCompetitorsRow) GetNSharedMechanismsOk() (*int32, bool)`

GetNSharedMechanismsOk returns a tuple with the NSharedMechanisms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedMechanisms

`func (o *CompanyCompetitorsRow) SetNSharedMechanisms(v int32)`

SetNSharedMechanisms sets NSharedMechanisms field to given value.

### HasNSharedMechanisms

`func (o *CompanyCompetitorsRow) HasNSharedMechanisms() bool`

HasNSharedMechanisms returns a boolean if a field has been set.

### SetNSharedMechanismsNil

`func (o *CompanyCompetitorsRow) SetNSharedMechanismsNil(b bool)`

 SetNSharedMechanismsNil sets the value for NSharedMechanisms to be an explicit nil

### UnsetNSharedMechanisms
`func (o *CompanyCompetitorsRow) UnsetNSharedMechanisms()`

UnsetNSharedMechanisms ensures that no value is present for NSharedMechanisms, not even an explicit nil
### GetNSharedIndications

`func (o *CompanyCompetitorsRow) GetNSharedIndications() int32`

GetNSharedIndications returns the NSharedIndications field if non-nil, zero value otherwise.

### GetNSharedIndicationsOk

`func (o *CompanyCompetitorsRow) GetNSharedIndicationsOk() (*int32, bool)`

GetNSharedIndicationsOk returns a tuple with the NSharedIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedIndications

`func (o *CompanyCompetitorsRow) SetNSharedIndications(v int32)`

SetNSharedIndications sets NSharedIndications field to given value.

### HasNSharedIndications

`func (o *CompanyCompetitorsRow) HasNSharedIndications() bool`

HasNSharedIndications returns a boolean if a field has been set.

### SetNSharedIndicationsNil

`func (o *CompanyCompetitorsRow) SetNSharedIndicationsNil(b bool)`

 SetNSharedIndicationsNil sets the value for NSharedIndications to be an explicit nil

### UnsetNSharedIndications
`func (o *CompanyCompetitorsRow) UnsetNSharedIndications()`

UnsetNSharedIndications ensures that no value is present for NSharedIndications, not even an explicit nil
### GetNPhaseOverlapTargets

`func (o *CompanyCompetitorsRow) GetNPhaseOverlapTargets() int32`

GetNPhaseOverlapTargets returns the NPhaseOverlapTargets field if non-nil, zero value otherwise.

### GetNPhaseOverlapTargetsOk

`func (o *CompanyCompetitorsRow) GetNPhaseOverlapTargetsOk() (*int32, bool)`

GetNPhaseOverlapTargetsOk returns a tuple with the NPhaseOverlapTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPhaseOverlapTargets

`func (o *CompanyCompetitorsRow) SetNPhaseOverlapTargets(v int32)`

SetNPhaseOverlapTargets sets NPhaseOverlapTargets field to given value.

### HasNPhaseOverlapTargets

`func (o *CompanyCompetitorsRow) HasNPhaseOverlapTargets() bool`

HasNPhaseOverlapTargets returns a boolean if a field has been set.

### SetNPhaseOverlapTargetsNil

`func (o *CompanyCompetitorsRow) SetNPhaseOverlapTargetsNil(b bool)`

 SetNPhaseOverlapTargetsNil sets the value for NPhaseOverlapTargets to be an explicit nil

### UnsetNPhaseOverlapTargets
`func (o *CompanyCompetitorsRow) UnsetNPhaseOverlapTargets()`

UnsetNPhaseOverlapTargets ensures that no value is present for NPhaseOverlapTargets, not even an explicit nil
### GetNCompanyTargets

`func (o *CompanyCompetitorsRow) GetNCompanyTargets() int32`

GetNCompanyTargets returns the NCompanyTargets field if non-nil, zero value otherwise.

### GetNCompanyTargetsOk

`func (o *CompanyCompetitorsRow) GetNCompanyTargetsOk() (*int32, bool)`

GetNCompanyTargetsOk returns a tuple with the NCompanyTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompanyTargets

`func (o *CompanyCompetitorsRow) SetNCompanyTargets(v int32)`

SetNCompanyTargets sets NCompanyTargets field to given value.

### HasNCompanyTargets

`func (o *CompanyCompetitorsRow) HasNCompanyTargets() bool`

HasNCompanyTargets returns a boolean if a field has been set.

### SetNCompanyTargetsNil

`func (o *CompanyCompetitorsRow) SetNCompanyTargetsNil(b bool)`

 SetNCompanyTargetsNil sets the value for NCompanyTargets to be an explicit nil

### UnsetNCompanyTargets
`func (o *CompanyCompetitorsRow) UnsetNCompanyTargets()`

UnsetNCompanyTargets ensures that no value is present for NCompanyTargets, not even an explicit nil
### GetNCompetitorTargets

`func (o *CompanyCompetitorsRow) GetNCompetitorTargets() int32`

GetNCompetitorTargets returns the NCompetitorTargets field if non-nil, zero value otherwise.

### GetNCompetitorTargetsOk

`func (o *CompanyCompetitorsRow) GetNCompetitorTargetsOk() (*int32, bool)`

GetNCompetitorTargetsOk returns a tuple with the NCompetitorTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompetitorTargets

`func (o *CompanyCompetitorsRow) SetNCompetitorTargets(v int32)`

SetNCompetitorTargets sets NCompetitorTargets field to given value.

### HasNCompetitorTargets

`func (o *CompanyCompetitorsRow) HasNCompetitorTargets() bool`

HasNCompetitorTargets returns a boolean if a field has been set.

### SetNCompetitorTargetsNil

`func (o *CompanyCompetitorsRow) SetNCompetitorTargetsNil(b bool)`

 SetNCompetitorTargetsNil sets the value for NCompetitorTargets to be an explicit nil

### UnsetNCompetitorTargets
`func (o *CompanyCompetitorsRow) UnsetNCompetitorTargets()`

UnsetNCompetitorTargets ensures that no value is present for NCompetitorTargets, not even an explicit nil
### GetNCompanyIndications

`func (o *CompanyCompetitorsRow) GetNCompanyIndications() int32`

GetNCompanyIndications returns the NCompanyIndications field if non-nil, zero value otherwise.

### GetNCompanyIndicationsOk

`func (o *CompanyCompetitorsRow) GetNCompanyIndicationsOk() (*int32, bool)`

GetNCompanyIndicationsOk returns a tuple with the NCompanyIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompanyIndications

`func (o *CompanyCompetitorsRow) SetNCompanyIndications(v int32)`

SetNCompanyIndications sets NCompanyIndications field to given value.

### HasNCompanyIndications

`func (o *CompanyCompetitorsRow) HasNCompanyIndications() bool`

HasNCompanyIndications returns a boolean if a field has been set.

### SetNCompanyIndicationsNil

`func (o *CompanyCompetitorsRow) SetNCompanyIndicationsNil(b bool)`

 SetNCompanyIndicationsNil sets the value for NCompanyIndications to be an explicit nil

### UnsetNCompanyIndications
`func (o *CompanyCompetitorsRow) UnsetNCompanyIndications()`

UnsetNCompanyIndications ensures that no value is present for NCompanyIndications, not even an explicit nil
### GetNCompetitorIndications

`func (o *CompanyCompetitorsRow) GetNCompetitorIndications() int32`

GetNCompetitorIndications returns the NCompetitorIndications field if non-nil, zero value otherwise.

### GetNCompetitorIndicationsOk

`func (o *CompanyCompetitorsRow) GetNCompetitorIndicationsOk() (*int32, bool)`

GetNCompetitorIndicationsOk returns a tuple with the NCompetitorIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompetitorIndications

`func (o *CompanyCompetitorsRow) SetNCompetitorIndications(v int32)`

SetNCompetitorIndications sets NCompetitorIndications field to given value.

### HasNCompetitorIndications

`func (o *CompanyCompetitorsRow) HasNCompetitorIndications() bool`

HasNCompetitorIndications returns a boolean if a field has been set.

### SetNCompetitorIndicationsNil

`func (o *CompanyCompetitorsRow) SetNCompetitorIndicationsNil(b bool)`

 SetNCompetitorIndicationsNil sets the value for NCompetitorIndications to be an explicit nil

### UnsetNCompetitorIndications
`func (o *CompanyCompetitorsRow) UnsetNCompetitorIndications()`

UnsetNCompetitorIndications ensures that no value is present for NCompetitorIndications, not even an explicit nil
### GetTopSharedTargets

`func (o *CompanyCompetitorsRow) GetTopSharedTargets() []string`

GetTopSharedTargets returns the TopSharedTargets field if non-nil, zero value otherwise.

### GetTopSharedTargetsOk

`func (o *CompanyCompetitorsRow) GetTopSharedTargetsOk() (*[]string, bool)`

GetTopSharedTargetsOk returns a tuple with the TopSharedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSharedTargets

`func (o *CompanyCompetitorsRow) SetTopSharedTargets(v []string)`

SetTopSharedTargets sets TopSharedTargets field to given value.

### HasTopSharedTargets

`func (o *CompanyCompetitorsRow) HasTopSharedTargets() bool`

HasTopSharedTargets returns a boolean if a field has been set.

### SetTopSharedTargetsNil

`func (o *CompanyCompetitorsRow) SetTopSharedTargetsNil(b bool)`

 SetTopSharedTargetsNil sets the value for TopSharedTargets to be an explicit nil

### UnsetTopSharedTargets
`func (o *CompanyCompetitorsRow) UnsetTopSharedTargets()`

UnsetTopSharedTargets ensures that no value is present for TopSharedTargets, not even an explicit nil
### GetTopSharedIndications

`func (o *CompanyCompetitorsRow) GetTopSharedIndications() []string`

GetTopSharedIndications returns the TopSharedIndications field if non-nil, zero value otherwise.

### GetTopSharedIndicationsOk

`func (o *CompanyCompetitorsRow) GetTopSharedIndicationsOk() (*[]string, bool)`

GetTopSharedIndicationsOk returns a tuple with the TopSharedIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSharedIndications

`func (o *CompanyCompetitorsRow) SetTopSharedIndications(v []string)`

SetTopSharedIndications sets TopSharedIndications field to given value.

### HasTopSharedIndications

`func (o *CompanyCompetitorsRow) HasTopSharedIndications() bool`

HasTopSharedIndications returns a boolean if a field has been set.

### SetTopSharedIndicationsNil

`func (o *CompanyCompetitorsRow) SetTopSharedIndicationsNil(b bool)`

 SetTopSharedIndicationsNil sets the value for TopSharedIndications to be an explicit nil

### UnsetTopSharedIndications
`func (o *CompanyCompetitorsRow) UnsetTopSharedIndications()`

UnsetTopSharedIndications ensures that no value is present for TopSharedIndications, not even an explicit nil
### GetTopSharedMechanisms

`func (o *CompanyCompetitorsRow) GetTopSharedMechanisms() []string`

GetTopSharedMechanisms returns the TopSharedMechanisms field if non-nil, zero value otherwise.

### GetTopSharedMechanismsOk

`func (o *CompanyCompetitorsRow) GetTopSharedMechanismsOk() (*[]string, bool)`

GetTopSharedMechanismsOk returns a tuple with the TopSharedMechanisms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSharedMechanisms

`func (o *CompanyCompetitorsRow) SetTopSharedMechanisms(v []string)`

SetTopSharedMechanisms sets TopSharedMechanisms field to given value.

### HasTopSharedMechanisms

`func (o *CompanyCompetitorsRow) HasTopSharedMechanisms() bool`

HasTopSharedMechanisms returns a boolean if a field has been set.

### SetTopSharedMechanismsNil

`func (o *CompanyCompetitorsRow) SetTopSharedMechanismsNil(b bool)`

 SetTopSharedMechanismsNil sets the value for TopSharedMechanisms to be an explicit nil

### UnsetTopSharedMechanisms
`func (o *CompanyCompetitorsRow) UnsetTopSharedMechanisms()`

UnsetTopSharedMechanisms ensures that no value is present for TopSharedMechanisms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


