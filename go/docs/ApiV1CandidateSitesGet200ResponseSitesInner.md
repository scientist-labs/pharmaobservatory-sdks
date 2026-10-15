# ApiV1CandidateSitesGet200ResponseSitesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rank** | **int32** |  | 
**SiteKey** | **string** |  | 
**FacilityName** | **string** |  | 
**City** | **string** |  | 
**State** | **string** |  | 
**Country** | **string** |  | 
**NTrialsInIndication** | **int32** |  | 
**EstimatedPatientsInIndication** | **NullableString** |  | 
**NActiveInIndication** | **int32** |  | 
**NActiveSamePhase** | **interface{}** |  | 
**CompositeScore** | **float32** |  | 
**CapacityZ** | **float32** |  | 
**ActivityZ** | **float32** |  | 
**CompetitionZ** | **float32** |  | 
**Saturated** | **bool** |  | 

## Methods

### NewApiV1CandidateSitesGet200ResponseSitesInner

`func NewApiV1CandidateSitesGet200ResponseSitesInner(rank int32, siteKey string, facilityName string, city string, state string, country string, nTrialsInIndication int32, estimatedPatientsInIndication NullableString, nActiveInIndication int32, nActiveSamePhase interface{}, compositeScore float32, capacityZ float32, activityZ float32, competitionZ float32, saturated bool, ) *ApiV1CandidateSitesGet200ResponseSitesInner`

NewApiV1CandidateSitesGet200ResponseSitesInner instantiates a new ApiV1CandidateSitesGet200ResponseSitesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1CandidateSitesGet200ResponseSitesInnerWithDefaults

`func NewApiV1CandidateSitesGet200ResponseSitesInnerWithDefaults() *ApiV1CandidateSitesGet200ResponseSitesInner`

NewApiV1CandidateSitesGet200ResponseSitesInnerWithDefaults instantiates a new ApiV1CandidateSitesGet200ResponseSitesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRank

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetSiteKey

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetSiteKey() string`

GetSiteKey returns the SiteKey field if non-nil, zero value otherwise.

### GetSiteKeyOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetSiteKeyOk() (*string, bool)`

GetSiteKeyOk returns a tuple with the SiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKey

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetSiteKey(v string)`

SetSiteKey sets SiteKey field to given value.


### GetFacilityName

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.


### GetCity

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetState(v string)`

SetState sets State field to given value.


### GetCountry

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetNTrialsInIndication

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetNTrialsInIndication() int32`

GetNTrialsInIndication returns the NTrialsInIndication field if non-nil, zero value otherwise.

### GetNTrialsInIndicationOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetNTrialsInIndicationOk() (*int32, bool)`

GetNTrialsInIndicationOk returns a tuple with the NTrialsInIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialsInIndication

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetNTrialsInIndication(v int32)`

SetNTrialsInIndication sets NTrialsInIndication field to given value.


### GetEstimatedPatientsInIndication

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetEstimatedPatientsInIndication() string`

GetEstimatedPatientsInIndication returns the EstimatedPatientsInIndication field if non-nil, zero value otherwise.

### GetEstimatedPatientsInIndicationOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetEstimatedPatientsInIndicationOk() (*string, bool)`

GetEstimatedPatientsInIndicationOk returns a tuple with the EstimatedPatientsInIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedPatientsInIndication

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetEstimatedPatientsInIndication(v string)`

SetEstimatedPatientsInIndication sets EstimatedPatientsInIndication field to given value.


### SetEstimatedPatientsInIndicationNil

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetEstimatedPatientsInIndicationNil(b bool)`

 SetEstimatedPatientsInIndicationNil sets the value for EstimatedPatientsInIndication to be an explicit nil

### UnsetEstimatedPatientsInIndication
`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) UnsetEstimatedPatientsInIndication()`

UnsetEstimatedPatientsInIndication ensures that no value is present for EstimatedPatientsInIndication, not even an explicit nil
### GetNActiveInIndication

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetNActiveInIndication() int32`

GetNActiveInIndication returns the NActiveInIndication field if non-nil, zero value otherwise.

### GetNActiveInIndicationOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetNActiveInIndicationOk() (*int32, bool)`

GetNActiveInIndicationOk returns a tuple with the NActiveInIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActiveInIndication

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetNActiveInIndication(v int32)`

SetNActiveInIndication sets NActiveInIndication field to given value.


### GetNActiveSamePhase

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetNActiveSamePhase() interface{}`

GetNActiveSamePhase returns the NActiveSamePhase field if non-nil, zero value otherwise.

### GetNActiveSamePhaseOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetNActiveSamePhaseOk() (*interface{}, bool)`

GetNActiveSamePhaseOk returns a tuple with the NActiveSamePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNActiveSamePhase

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetNActiveSamePhase(v interface{})`

SetNActiveSamePhase sets NActiveSamePhase field to given value.


### SetNActiveSamePhaseNil

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetNActiveSamePhaseNil(b bool)`

 SetNActiveSamePhaseNil sets the value for NActiveSamePhase to be an explicit nil

### UnsetNActiveSamePhase
`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) UnsetNActiveSamePhase()`

UnsetNActiveSamePhase ensures that no value is present for NActiveSamePhase, not even an explicit nil
### GetCompositeScore

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCompositeScore() float32`

GetCompositeScore returns the CompositeScore field if non-nil, zero value otherwise.

### GetCompositeScoreOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCompositeScoreOk() (*float32, bool)`

GetCompositeScoreOk returns a tuple with the CompositeScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeScore

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetCompositeScore(v float32)`

SetCompositeScore sets CompositeScore field to given value.


### GetCapacityZ

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCapacityZ() float32`

GetCapacityZ returns the CapacityZ field if non-nil, zero value otherwise.

### GetCapacityZOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCapacityZOk() (*float32, bool)`

GetCapacityZOk returns a tuple with the CapacityZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityZ

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetCapacityZ(v float32)`

SetCapacityZ sets CapacityZ field to given value.


### GetActivityZ

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetActivityZ() float32`

GetActivityZ returns the ActivityZ field if non-nil, zero value otherwise.

### GetActivityZOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetActivityZOk() (*float32, bool)`

GetActivityZOk returns a tuple with the ActivityZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityZ

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetActivityZ(v float32)`

SetActivityZ sets ActivityZ field to given value.


### GetCompetitionZ

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCompetitionZ() float32`

GetCompetitionZ returns the CompetitionZ field if non-nil, zero value otherwise.

### GetCompetitionZOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetCompetitionZOk() (*float32, bool)`

GetCompetitionZOk returns a tuple with the CompetitionZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionZ

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetCompetitionZ(v float32)`

SetCompetitionZ sets CompetitionZ field to given value.


### GetSaturated

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetSaturated() bool`

GetSaturated returns the Saturated field if non-nil, zero value otherwise.

### GetSaturatedOk

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) GetSaturatedOk() (*bool, bool)`

GetSaturatedOk returns a tuple with the Saturated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturated

`func (o *ApiV1CandidateSitesGet200ResponseSitesInner) SetSaturated(v bool)`

SetSaturated sets Saturated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


