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
**CompetitorClass** | Pointer to **NullableString** | The competitor&#39;s class: always industry since the not-classed-as-industry group left Competitors; the column is kept so a query written against it keeps working. competitor_rank is 1-based within the class. | [optional] 
**NClassCandidates** | Pointer to **NullableInt32** | How many companies of this competitor_class had at least one asset collision with the company before the top 25 were kept. The same value on every row of one company and class. | [optional] 
**CompetitorRank** | Pointer to **NullableInt32** | The competitor&#39;s rank within its competitor_class for the company by competitor_score (1 is the highest). | [optional] 
**CompetitorScore** | Pointer to **NullableFloat32** | The weighted share of the company&#39;s live target families that the competitor&#39;s live assets collide with, 0 to 100, rounded to four places. Each colliding family counts its best collision weight (rarity of the family across companies, times factors for same action, a shared indication, the least advanced stage, recency and a head-to-head trial), capped at the family&#39;s rarity weight; the sum is divided by the rarity weights of ALL the company&#39;s families. | [optional] 
**NCollisions** | Pointer to **NullableInt32** | Asset pairs (a drug of the company, a drug of the competitor) that act on a common target. | [optional] 
**NCollidingFamilies** | Pointer to **NullableInt32** | Target families of the company&#39;s live assets that the competitor collides with. | [optional] 
**NSubjectFamilies** | Pointer to **NullableInt32** | Target families of the company&#39;s own live, owned assets; the denominator of competitor_score is their rarity weights. | [optional] 
**NSameAction** | Pointer to **NullableInt32** | Asset pairs whose drugs act on the shared target with the same action type. | [optional] 
**NSharedIndicationPairs** | Pointer to **NullableInt32** | Asset pairs whose programs share at least one disease area. Every listed pair does, so it equals n_collisions. | [optional] 
**NHeadToHeadTrials** | Pointer to **NullableInt32** | Distinct trials led by either company in which one drug of a pair is experimental and the other an active comparator. | [optional] 
**BestStage** | Pointer to **NullableInt32** | The most advanced pair&#39;s stage, where a pair&#39;s stage is the less advanced of its two programs&#39; furthest phase: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1, -1 phase not applicable. A tie-break key after the score. | [optional] 
**LastYear** | Pointer to **NullableInt32** | The latest trial start year over the company pair&#39;s programs (the maximum over its asset pairs of the later of the two programs&#39; last start year). A tie-break key after the score; null when no program has a start year. | [optional] 
**TopCollisions** | Pointer to **[]string** | Up to 3 strongest asset pairs as text: \&quot;drug A / drug B: target family, stage, N shared indications\&quot;. The stage is the less advanced of the two programs. An ownership_basis of inferred_lead on the underlying pair (see company_asset_collisions) is an inference. | [optional] 
**NCosponsoredTrials** | Pointer to **NullableInt32** | Distinct trials on which the company and the competitor are both named as sponsors, counted over their subtrees; 0 when none. A consortium trial (more than 25 sponsor companies) is left out. | [optional] 

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
### GetCompetitorClass

`func (o *CompanyCompetitorsRow) GetCompetitorClass() string`

GetCompetitorClass returns the CompetitorClass field if non-nil, zero value otherwise.

### GetCompetitorClassOk

`func (o *CompanyCompetitorsRow) GetCompetitorClassOk() (*string, bool)`

GetCompetitorClassOk returns a tuple with the CompetitorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorClass

`func (o *CompanyCompetitorsRow) SetCompetitorClass(v string)`

SetCompetitorClass sets CompetitorClass field to given value.

### HasCompetitorClass

`func (o *CompanyCompetitorsRow) HasCompetitorClass() bool`

HasCompetitorClass returns a boolean if a field has been set.

### SetCompetitorClassNil

`func (o *CompanyCompetitorsRow) SetCompetitorClassNil(b bool)`

 SetCompetitorClassNil sets the value for CompetitorClass to be an explicit nil

### UnsetCompetitorClass
`func (o *CompanyCompetitorsRow) UnsetCompetitorClass()`

UnsetCompetitorClass ensures that no value is present for CompetitorClass, not even an explicit nil
### GetNClassCandidates

`func (o *CompanyCompetitorsRow) GetNClassCandidates() int32`

GetNClassCandidates returns the NClassCandidates field if non-nil, zero value otherwise.

### GetNClassCandidatesOk

`func (o *CompanyCompetitorsRow) GetNClassCandidatesOk() (*int32, bool)`

GetNClassCandidatesOk returns a tuple with the NClassCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNClassCandidates

`func (o *CompanyCompetitorsRow) SetNClassCandidates(v int32)`

SetNClassCandidates sets NClassCandidates field to given value.

### HasNClassCandidates

`func (o *CompanyCompetitorsRow) HasNClassCandidates() bool`

HasNClassCandidates returns a boolean if a field has been set.

### SetNClassCandidatesNil

`func (o *CompanyCompetitorsRow) SetNClassCandidatesNil(b bool)`

 SetNClassCandidatesNil sets the value for NClassCandidates to be an explicit nil

### UnsetNClassCandidates
`func (o *CompanyCompetitorsRow) UnsetNClassCandidates()`

UnsetNClassCandidates ensures that no value is present for NClassCandidates, not even an explicit nil
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
### GetNCollisions

`func (o *CompanyCompetitorsRow) GetNCollisions() int32`

GetNCollisions returns the NCollisions field if non-nil, zero value otherwise.

### GetNCollisionsOk

`func (o *CompanyCompetitorsRow) GetNCollisionsOk() (*int32, bool)`

GetNCollisionsOk returns a tuple with the NCollisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCollisions

`func (o *CompanyCompetitorsRow) SetNCollisions(v int32)`

SetNCollisions sets NCollisions field to given value.

### HasNCollisions

`func (o *CompanyCompetitorsRow) HasNCollisions() bool`

HasNCollisions returns a boolean if a field has been set.

### SetNCollisionsNil

`func (o *CompanyCompetitorsRow) SetNCollisionsNil(b bool)`

 SetNCollisionsNil sets the value for NCollisions to be an explicit nil

### UnsetNCollisions
`func (o *CompanyCompetitorsRow) UnsetNCollisions()`

UnsetNCollisions ensures that no value is present for NCollisions, not even an explicit nil
### GetNCollidingFamilies

`func (o *CompanyCompetitorsRow) GetNCollidingFamilies() int32`

GetNCollidingFamilies returns the NCollidingFamilies field if non-nil, zero value otherwise.

### GetNCollidingFamiliesOk

`func (o *CompanyCompetitorsRow) GetNCollidingFamiliesOk() (*int32, bool)`

GetNCollidingFamiliesOk returns a tuple with the NCollidingFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCollidingFamilies

`func (o *CompanyCompetitorsRow) SetNCollidingFamilies(v int32)`

SetNCollidingFamilies sets NCollidingFamilies field to given value.

### HasNCollidingFamilies

`func (o *CompanyCompetitorsRow) HasNCollidingFamilies() bool`

HasNCollidingFamilies returns a boolean if a field has been set.

### SetNCollidingFamiliesNil

`func (o *CompanyCompetitorsRow) SetNCollidingFamiliesNil(b bool)`

 SetNCollidingFamiliesNil sets the value for NCollidingFamilies to be an explicit nil

### UnsetNCollidingFamilies
`func (o *CompanyCompetitorsRow) UnsetNCollidingFamilies()`

UnsetNCollidingFamilies ensures that no value is present for NCollidingFamilies, not even an explicit nil
### GetNSubjectFamilies

`func (o *CompanyCompetitorsRow) GetNSubjectFamilies() int32`

GetNSubjectFamilies returns the NSubjectFamilies field if non-nil, zero value otherwise.

### GetNSubjectFamiliesOk

`func (o *CompanyCompetitorsRow) GetNSubjectFamiliesOk() (*int32, bool)`

GetNSubjectFamiliesOk returns a tuple with the NSubjectFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSubjectFamilies

`func (o *CompanyCompetitorsRow) SetNSubjectFamilies(v int32)`

SetNSubjectFamilies sets NSubjectFamilies field to given value.

### HasNSubjectFamilies

`func (o *CompanyCompetitorsRow) HasNSubjectFamilies() bool`

HasNSubjectFamilies returns a boolean if a field has been set.

### SetNSubjectFamiliesNil

`func (o *CompanyCompetitorsRow) SetNSubjectFamiliesNil(b bool)`

 SetNSubjectFamiliesNil sets the value for NSubjectFamilies to be an explicit nil

### UnsetNSubjectFamilies
`func (o *CompanyCompetitorsRow) UnsetNSubjectFamilies()`

UnsetNSubjectFamilies ensures that no value is present for NSubjectFamilies, not even an explicit nil
### GetNSameAction

`func (o *CompanyCompetitorsRow) GetNSameAction() int32`

GetNSameAction returns the NSameAction field if non-nil, zero value otherwise.

### GetNSameActionOk

`func (o *CompanyCompetitorsRow) GetNSameActionOk() (*int32, bool)`

GetNSameActionOk returns a tuple with the NSameAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSameAction

`func (o *CompanyCompetitorsRow) SetNSameAction(v int32)`

SetNSameAction sets NSameAction field to given value.

### HasNSameAction

`func (o *CompanyCompetitorsRow) HasNSameAction() bool`

HasNSameAction returns a boolean if a field has been set.

### SetNSameActionNil

`func (o *CompanyCompetitorsRow) SetNSameActionNil(b bool)`

 SetNSameActionNil sets the value for NSameAction to be an explicit nil

### UnsetNSameAction
`func (o *CompanyCompetitorsRow) UnsetNSameAction()`

UnsetNSameAction ensures that no value is present for NSameAction, not even an explicit nil
### GetNSharedIndicationPairs

`func (o *CompanyCompetitorsRow) GetNSharedIndicationPairs() int32`

GetNSharedIndicationPairs returns the NSharedIndicationPairs field if non-nil, zero value otherwise.

### GetNSharedIndicationPairsOk

`func (o *CompanyCompetitorsRow) GetNSharedIndicationPairsOk() (*int32, bool)`

GetNSharedIndicationPairsOk returns a tuple with the NSharedIndicationPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedIndicationPairs

`func (o *CompanyCompetitorsRow) SetNSharedIndicationPairs(v int32)`

SetNSharedIndicationPairs sets NSharedIndicationPairs field to given value.

### HasNSharedIndicationPairs

`func (o *CompanyCompetitorsRow) HasNSharedIndicationPairs() bool`

HasNSharedIndicationPairs returns a boolean if a field has been set.

### SetNSharedIndicationPairsNil

`func (o *CompanyCompetitorsRow) SetNSharedIndicationPairsNil(b bool)`

 SetNSharedIndicationPairsNil sets the value for NSharedIndicationPairs to be an explicit nil

### UnsetNSharedIndicationPairs
`func (o *CompanyCompetitorsRow) UnsetNSharedIndicationPairs()`

UnsetNSharedIndicationPairs ensures that no value is present for NSharedIndicationPairs, not even an explicit nil
### GetNHeadToHeadTrials

`func (o *CompanyCompetitorsRow) GetNHeadToHeadTrials() int32`

GetNHeadToHeadTrials returns the NHeadToHeadTrials field if non-nil, zero value otherwise.

### GetNHeadToHeadTrialsOk

`func (o *CompanyCompetitorsRow) GetNHeadToHeadTrialsOk() (*int32, bool)`

GetNHeadToHeadTrialsOk returns a tuple with the NHeadToHeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNHeadToHeadTrials

`func (o *CompanyCompetitorsRow) SetNHeadToHeadTrials(v int32)`

SetNHeadToHeadTrials sets NHeadToHeadTrials field to given value.

### HasNHeadToHeadTrials

`func (o *CompanyCompetitorsRow) HasNHeadToHeadTrials() bool`

HasNHeadToHeadTrials returns a boolean if a field has been set.

### SetNHeadToHeadTrialsNil

`func (o *CompanyCompetitorsRow) SetNHeadToHeadTrialsNil(b bool)`

 SetNHeadToHeadTrialsNil sets the value for NHeadToHeadTrials to be an explicit nil

### UnsetNHeadToHeadTrials
`func (o *CompanyCompetitorsRow) UnsetNHeadToHeadTrials()`

UnsetNHeadToHeadTrials ensures that no value is present for NHeadToHeadTrials, not even an explicit nil
### GetBestStage

`func (o *CompanyCompetitorsRow) GetBestStage() int32`

GetBestStage returns the BestStage field if non-nil, zero value otherwise.

### GetBestStageOk

`func (o *CompanyCompetitorsRow) GetBestStageOk() (*int32, bool)`

GetBestStageOk returns a tuple with the BestStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestStage

`func (o *CompanyCompetitorsRow) SetBestStage(v int32)`

SetBestStage sets BestStage field to given value.

### HasBestStage

`func (o *CompanyCompetitorsRow) HasBestStage() bool`

HasBestStage returns a boolean if a field has been set.

### SetBestStageNil

`func (o *CompanyCompetitorsRow) SetBestStageNil(b bool)`

 SetBestStageNil sets the value for BestStage to be an explicit nil

### UnsetBestStage
`func (o *CompanyCompetitorsRow) UnsetBestStage()`

UnsetBestStage ensures that no value is present for BestStage, not even an explicit nil
### GetLastYear

`func (o *CompanyCompetitorsRow) GetLastYear() int32`

GetLastYear returns the LastYear field if non-nil, zero value otherwise.

### GetLastYearOk

`func (o *CompanyCompetitorsRow) GetLastYearOk() (*int32, bool)`

GetLastYearOk returns a tuple with the LastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYear

`func (o *CompanyCompetitorsRow) SetLastYear(v int32)`

SetLastYear sets LastYear field to given value.

### HasLastYear

`func (o *CompanyCompetitorsRow) HasLastYear() bool`

HasLastYear returns a boolean if a field has been set.

### SetLastYearNil

`func (o *CompanyCompetitorsRow) SetLastYearNil(b bool)`

 SetLastYearNil sets the value for LastYear to be an explicit nil

### UnsetLastYear
`func (o *CompanyCompetitorsRow) UnsetLastYear()`

UnsetLastYear ensures that no value is present for LastYear, not even an explicit nil
### GetTopCollisions

`func (o *CompanyCompetitorsRow) GetTopCollisions() []string`

GetTopCollisions returns the TopCollisions field if non-nil, zero value otherwise.

### GetTopCollisionsOk

`func (o *CompanyCompetitorsRow) GetTopCollisionsOk() (*[]string, bool)`

GetTopCollisionsOk returns a tuple with the TopCollisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopCollisions

`func (o *CompanyCompetitorsRow) SetTopCollisions(v []string)`

SetTopCollisions sets TopCollisions field to given value.

### HasTopCollisions

`func (o *CompanyCompetitorsRow) HasTopCollisions() bool`

HasTopCollisions returns a boolean if a field has been set.

### SetTopCollisionsNil

`func (o *CompanyCompetitorsRow) SetTopCollisionsNil(b bool)`

 SetTopCollisionsNil sets the value for TopCollisions to be an explicit nil

### UnsetTopCollisions
`func (o *CompanyCompetitorsRow) UnsetTopCollisions()`

UnsetTopCollisions ensures that no value is present for TopCollisions, not even an explicit nil
### GetNCosponsoredTrials

`func (o *CompanyCompetitorsRow) GetNCosponsoredTrials() int32`

GetNCosponsoredTrials returns the NCosponsoredTrials field if non-nil, zero value otherwise.

### GetNCosponsoredTrialsOk

`func (o *CompanyCompetitorsRow) GetNCosponsoredTrialsOk() (*int32, bool)`

GetNCosponsoredTrialsOk returns a tuple with the NCosponsoredTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCosponsoredTrials

`func (o *CompanyCompetitorsRow) SetNCosponsoredTrials(v int32)`

SetNCosponsoredTrials sets NCosponsoredTrials field to given value.

### HasNCosponsoredTrials

`func (o *CompanyCompetitorsRow) HasNCosponsoredTrials() bool`

HasNCosponsoredTrials returns a boolean if a field has been set.

### SetNCosponsoredTrialsNil

`func (o *CompanyCompetitorsRow) SetNCosponsoredTrialsNil(b bool)`

 SetNCosponsoredTrialsNil sets the value for NCosponsoredTrials to be an explicit nil

### UnsetNCosponsoredTrials
`func (o *CompanyCompetitorsRow) UnsetNCosponsoredTrials()`

UnsetNCosponsoredTrials ensures that no value is present for NCosponsoredTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


