# OrgsCompetitors200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**OrgName** | **string** |  | 
**CompetitorOrgId** | **string** |  | 
**CompetitorName** | **string** |  | 
**CompetitorRank** | **int32** |  | 
**CompetitorScore** | **float32** |  | 
**TargetSim** | **float32** |  | 
**MechanismSim** | **float32** |  | 
**IndicationSim** | **float32** |  | 
**PhaseSim** | **float32** |  | 
**NSharedTargets** | **int32** |  | 
**NSharedMechanisms** | **int32** |  | 
**NSharedIndications** | **int32** |  | 
**NPhaseOverlapTargets** | **int32** |  | 
**NOrgTargets** | **int32** |  | 
**NCompetitorTargets** | **int32** |  | 
**NOrgIndications** | **int32** |  | 
**NCompetitorIndications** | **int32** |  | 
**TopSharedTargets** | **[]string** |  | 
**TopSharedIndications** | **[]string** |  | 
**TopSharedMechanisms** | **[]string** |  | 

## Methods

### NewOrgsCompetitors200ResponseRowsInner

`func NewOrgsCompetitors200ResponseRowsInner(orgId string, orgName string, competitorOrgId string, competitorName string, competitorRank int32, competitorScore float32, targetSim float32, mechanismSim float32, indicationSim float32, phaseSim float32, nSharedTargets int32, nSharedMechanisms int32, nSharedIndications int32, nPhaseOverlapTargets int32, nOrgTargets int32, nCompetitorTargets int32, nOrgIndications int32, nCompetitorIndications int32, topSharedTargets []string, topSharedIndications []string, topSharedMechanisms []string, ) *OrgsCompetitors200ResponseRowsInner`

NewOrgsCompetitors200ResponseRowsInner instantiates a new OrgsCompetitors200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsCompetitors200ResponseRowsInnerWithDefaults

`func NewOrgsCompetitors200ResponseRowsInnerWithDefaults() *OrgsCompetitors200ResponseRowsInner`

NewOrgsCompetitors200ResponseRowsInnerWithDefaults instantiates a new OrgsCompetitors200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrgsCompetitors200ResponseRowsInner) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgsCompetitors200ResponseRowsInner) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetOrgName

`func (o *OrgsCompetitors200ResponseRowsInner) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *OrgsCompetitors200ResponseRowsInner) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetCompetitorOrgId

`func (o *OrgsCompetitors200ResponseRowsInner) GetCompetitorOrgId() string`

GetCompetitorOrgId returns the CompetitorOrgId field if non-nil, zero value otherwise.

### GetCompetitorOrgIdOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetCompetitorOrgIdOk() (*string, bool)`

GetCompetitorOrgIdOk returns a tuple with the CompetitorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorOrgId

`func (o *OrgsCompetitors200ResponseRowsInner) SetCompetitorOrgId(v string)`

SetCompetitorOrgId sets CompetitorOrgId field to given value.


### GetCompetitorName

`func (o *OrgsCompetitors200ResponseRowsInner) GetCompetitorName() string`

GetCompetitorName returns the CompetitorName field if non-nil, zero value otherwise.

### GetCompetitorNameOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetCompetitorNameOk() (*string, bool)`

GetCompetitorNameOk returns a tuple with the CompetitorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorName

`func (o *OrgsCompetitors200ResponseRowsInner) SetCompetitorName(v string)`

SetCompetitorName sets CompetitorName field to given value.


### GetCompetitorRank

`func (o *OrgsCompetitors200ResponseRowsInner) GetCompetitorRank() int32`

GetCompetitorRank returns the CompetitorRank field if non-nil, zero value otherwise.

### GetCompetitorRankOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetCompetitorRankOk() (*int32, bool)`

GetCompetitorRankOk returns a tuple with the CompetitorRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorRank

`func (o *OrgsCompetitors200ResponseRowsInner) SetCompetitorRank(v int32)`

SetCompetitorRank sets CompetitorRank field to given value.


### GetCompetitorScore

`func (o *OrgsCompetitors200ResponseRowsInner) GetCompetitorScore() float32`

GetCompetitorScore returns the CompetitorScore field if non-nil, zero value otherwise.

### GetCompetitorScoreOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetCompetitorScoreOk() (*float32, bool)`

GetCompetitorScoreOk returns a tuple with the CompetitorScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitorScore

`func (o *OrgsCompetitors200ResponseRowsInner) SetCompetitorScore(v float32)`

SetCompetitorScore sets CompetitorScore field to given value.


### GetTargetSim

`func (o *OrgsCompetitors200ResponseRowsInner) GetTargetSim() float32`

GetTargetSim returns the TargetSim field if non-nil, zero value otherwise.

### GetTargetSimOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetTargetSimOk() (*float32, bool)`

GetTargetSimOk returns a tuple with the TargetSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSim

`func (o *OrgsCompetitors200ResponseRowsInner) SetTargetSim(v float32)`

SetTargetSim sets TargetSim field to given value.


### GetMechanismSim

`func (o *OrgsCompetitors200ResponseRowsInner) GetMechanismSim() float32`

GetMechanismSim returns the MechanismSim field if non-nil, zero value otherwise.

### GetMechanismSimOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetMechanismSimOk() (*float32, bool)`

GetMechanismSimOk returns a tuple with the MechanismSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanismSim

`func (o *OrgsCompetitors200ResponseRowsInner) SetMechanismSim(v float32)`

SetMechanismSim sets MechanismSim field to given value.


### GetIndicationSim

`func (o *OrgsCompetitors200ResponseRowsInner) GetIndicationSim() float32`

GetIndicationSim returns the IndicationSim field if non-nil, zero value otherwise.

### GetIndicationSimOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetIndicationSimOk() (*float32, bool)`

GetIndicationSimOk returns a tuple with the IndicationSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicationSim

`func (o *OrgsCompetitors200ResponseRowsInner) SetIndicationSim(v float32)`

SetIndicationSim sets IndicationSim field to given value.


### GetPhaseSim

`func (o *OrgsCompetitors200ResponseRowsInner) GetPhaseSim() float32`

GetPhaseSim returns the PhaseSim field if non-nil, zero value otherwise.

### GetPhaseSimOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetPhaseSimOk() (*float32, bool)`

GetPhaseSimOk returns a tuple with the PhaseSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseSim

`func (o *OrgsCompetitors200ResponseRowsInner) SetPhaseSim(v float32)`

SetPhaseSim sets PhaseSim field to given value.


### GetNSharedTargets

`func (o *OrgsCompetitors200ResponseRowsInner) GetNSharedTargets() int32`

GetNSharedTargets returns the NSharedTargets field if non-nil, zero value otherwise.

### GetNSharedTargetsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetNSharedTargetsOk() (*int32, bool)`

GetNSharedTargetsOk returns a tuple with the NSharedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedTargets

`func (o *OrgsCompetitors200ResponseRowsInner) SetNSharedTargets(v int32)`

SetNSharedTargets sets NSharedTargets field to given value.


### GetNSharedMechanisms

`func (o *OrgsCompetitors200ResponseRowsInner) GetNSharedMechanisms() int32`

GetNSharedMechanisms returns the NSharedMechanisms field if non-nil, zero value otherwise.

### GetNSharedMechanismsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetNSharedMechanismsOk() (*int32, bool)`

GetNSharedMechanismsOk returns a tuple with the NSharedMechanisms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedMechanisms

`func (o *OrgsCompetitors200ResponseRowsInner) SetNSharedMechanisms(v int32)`

SetNSharedMechanisms sets NSharedMechanisms field to given value.


### GetNSharedIndications

`func (o *OrgsCompetitors200ResponseRowsInner) GetNSharedIndications() int32`

GetNSharedIndications returns the NSharedIndications field if non-nil, zero value otherwise.

### GetNSharedIndicationsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetNSharedIndicationsOk() (*int32, bool)`

GetNSharedIndicationsOk returns a tuple with the NSharedIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSharedIndications

`func (o *OrgsCompetitors200ResponseRowsInner) SetNSharedIndications(v int32)`

SetNSharedIndications sets NSharedIndications field to given value.


### GetNPhaseOverlapTargets

`func (o *OrgsCompetitors200ResponseRowsInner) GetNPhaseOverlapTargets() int32`

GetNPhaseOverlapTargets returns the NPhaseOverlapTargets field if non-nil, zero value otherwise.

### GetNPhaseOverlapTargetsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetNPhaseOverlapTargetsOk() (*int32, bool)`

GetNPhaseOverlapTargetsOk returns a tuple with the NPhaseOverlapTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPhaseOverlapTargets

`func (o *OrgsCompetitors200ResponseRowsInner) SetNPhaseOverlapTargets(v int32)`

SetNPhaseOverlapTargets sets NPhaseOverlapTargets field to given value.


### GetNOrgTargets

`func (o *OrgsCompetitors200ResponseRowsInner) GetNOrgTargets() int32`

GetNOrgTargets returns the NOrgTargets field if non-nil, zero value otherwise.

### GetNOrgTargetsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetNOrgTargetsOk() (*int32, bool)`

GetNOrgTargetsOk returns a tuple with the NOrgTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOrgTargets

`func (o *OrgsCompetitors200ResponseRowsInner) SetNOrgTargets(v int32)`

SetNOrgTargets sets NOrgTargets field to given value.


### GetNCompetitorTargets

`func (o *OrgsCompetitors200ResponseRowsInner) GetNCompetitorTargets() int32`

GetNCompetitorTargets returns the NCompetitorTargets field if non-nil, zero value otherwise.

### GetNCompetitorTargetsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetNCompetitorTargetsOk() (*int32, bool)`

GetNCompetitorTargetsOk returns a tuple with the NCompetitorTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompetitorTargets

`func (o *OrgsCompetitors200ResponseRowsInner) SetNCompetitorTargets(v int32)`

SetNCompetitorTargets sets NCompetitorTargets field to given value.


### GetNOrgIndications

`func (o *OrgsCompetitors200ResponseRowsInner) GetNOrgIndications() int32`

GetNOrgIndications returns the NOrgIndications field if non-nil, zero value otherwise.

### GetNOrgIndicationsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetNOrgIndicationsOk() (*int32, bool)`

GetNOrgIndicationsOk returns a tuple with the NOrgIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOrgIndications

`func (o *OrgsCompetitors200ResponseRowsInner) SetNOrgIndications(v int32)`

SetNOrgIndications sets NOrgIndications field to given value.


### GetNCompetitorIndications

`func (o *OrgsCompetitors200ResponseRowsInner) GetNCompetitorIndications() int32`

GetNCompetitorIndications returns the NCompetitorIndications field if non-nil, zero value otherwise.

### GetNCompetitorIndicationsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetNCompetitorIndicationsOk() (*int32, bool)`

GetNCompetitorIndicationsOk returns a tuple with the NCompetitorIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCompetitorIndications

`func (o *OrgsCompetitors200ResponseRowsInner) SetNCompetitorIndications(v int32)`

SetNCompetitorIndications sets NCompetitorIndications field to given value.


### GetTopSharedTargets

`func (o *OrgsCompetitors200ResponseRowsInner) GetTopSharedTargets() []string`

GetTopSharedTargets returns the TopSharedTargets field if non-nil, zero value otherwise.

### GetTopSharedTargetsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetTopSharedTargetsOk() (*[]string, bool)`

GetTopSharedTargetsOk returns a tuple with the TopSharedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSharedTargets

`func (o *OrgsCompetitors200ResponseRowsInner) SetTopSharedTargets(v []string)`

SetTopSharedTargets sets TopSharedTargets field to given value.


### GetTopSharedIndications

`func (o *OrgsCompetitors200ResponseRowsInner) GetTopSharedIndications() []string`

GetTopSharedIndications returns the TopSharedIndications field if non-nil, zero value otherwise.

### GetTopSharedIndicationsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetTopSharedIndicationsOk() (*[]string, bool)`

GetTopSharedIndicationsOk returns a tuple with the TopSharedIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSharedIndications

`func (o *OrgsCompetitors200ResponseRowsInner) SetTopSharedIndications(v []string)`

SetTopSharedIndications sets TopSharedIndications field to given value.


### GetTopSharedMechanisms

`func (o *OrgsCompetitors200ResponseRowsInner) GetTopSharedMechanisms() []string`

GetTopSharedMechanisms returns the TopSharedMechanisms field if non-nil, zero value otherwise.

### GetTopSharedMechanismsOk

`func (o *OrgsCompetitors200ResponseRowsInner) GetTopSharedMechanismsOk() (*[]string, bool)`

GetTopSharedMechanismsOk returns a tuple with the TopSharedMechanisms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSharedMechanisms

`func (o *OrgsCompetitors200ResponseRowsInner) SetTopSharedMechanisms(v []string)`

SetTopSharedMechanisms sets TopSharedMechanisms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


