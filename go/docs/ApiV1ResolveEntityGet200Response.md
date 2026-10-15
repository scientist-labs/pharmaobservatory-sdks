# ApiV1ResolveEntityGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spine** | **string** |  | 
**Mode** | **string** |  | 
**Query** | **NullableString** |  | 
**SourceType** | **NullableString** |  | 
**RawKey** | **NullableString** |  | 
**Candidates** | [**[]ApiV1ResolveEntityGet200ResponseCandidatesInner**](ApiV1ResolveEntityGet200ResponseCandidatesInner.md) |  | 
**PersistenceStable** | **NullableBool** |  | 
**Message** | Pointer to **string** |  | [optional] 
**ResolutionConfidence** | [**ApiV1ResolveEntityGet200ResponseResolutionConfidence**](ApiV1ResolveEntityGet200ResponseResolutionConfidence.md) |  | 
**BuildState** | Pointer to [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | [optional] 

## Methods

### NewApiV1ResolveEntityGet200Response

`func NewApiV1ResolveEntityGet200Response(spine string, mode string, query NullableString, sourceType NullableString, rawKey NullableString, candidates []ApiV1ResolveEntityGet200ResponseCandidatesInner, persistenceStable NullableBool, resolutionConfidence ApiV1ResolveEntityGet200ResponseResolutionConfidence, ) *ApiV1ResolveEntityGet200Response`

NewApiV1ResolveEntityGet200Response instantiates a new ApiV1ResolveEntityGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ResolveEntityGet200ResponseWithDefaults

`func NewApiV1ResolveEntityGet200ResponseWithDefaults() *ApiV1ResolveEntityGet200Response`

NewApiV1ResolveEntityGet200ResponseWithDefaults instantiates a new ApiV1ResolveEntityGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpine

`func (o *ApiV1ResolveEntityGet200Response) GetSpine() string`

GetSpine returns the Spine field if non-nil, zero value otherwise.

### GetSpineOk

`func (o *ApiV1ResolveEntityGet200Response) GetSpineOk() (*string, bool)`

GetSpineOk returns a tuple with the Spine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpine

`func (o *ApiV1ResolveEntityGet200Response) SetSpine(v string)`

SetSpine sets Spine field to given value.


### GetMode

`func (o *ApiV1ResolveEntityGet200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApiV1ResolveEntityGet200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApiV1ResolveEntityGet200Response) SetMode(v string)`

SetMode sets Mode field to given value.


### GetQuery

`func (o *ApiV1ResolveEntityGet200Response) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ApiV1ResolveEntityGet200Response) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ApiV1ResolveEntityGet200Response) SetQuery(v string)`

SetQuery sets Query field to given value.


### SetQueryNil

`func (o *ApiV1ResolveEntityGet200Response) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *ApiV1ResolveEntityGet200Response) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetSourceType

`func (o *ApiV1ResolveEntityGet200Response) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ApiV1ResolveEntityGet200Response) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ApiV1ResolveEntityGet200Response) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### SetSourceTypeNil

`func (o *ApiV1ResolveEntityGet200Response) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *ApiV1ResolveEntityGet200Response) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetRawKey

`func (o *ApiV1ResolveEntityGet200Response) GetRawKey() string`

GetRawKey returns the RawKey field if non-nil, zero value otherwise.

### GetRawKeyOk

`func (o *ApiV1ResolveEntityGet200Response) GetRawKeyOk() (*string, bool)`

GetRawKeyOk returns a tuple with the RawKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawKey

`func (o *ApiV1ResolveEntityGet200Response) SetRawKey(v string)`

SetRawKey sets RawKey field to given value.


### SetRawKeyNil

`func (o *ApiV1ResolveEntityGet200Response) SetRawKeyNil(b bool)`

 SetRawKeyNil sets the value for RawKey to be an explicit nil

### UnsetRawKey
`func (o *ApiV1ResolveEntityGet200Response) UnsetRawKey()`

UnsetRawKey ensures that no value is present for RawKey, not even an explicit nil
### GetCandidates

`func (o *ApiV1ResolveEntityGet200Response) GetCandidates() []ApiV1ResolveEntityGet200ResponseCandidatesInner`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *ApiV1ResolveEntityGet200Response) GetCandidatesOk() (*[]ApiV1ResolveEntityGet200ResponseCandidatesInner, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *ApiV1ResolveEntityGet200Response) SetCandidates(v []ApiV1ResolveEntityGet200ResponseCandidatesInner)`

SetCandidates sets Candidates field to given value.


### GetPersistenceStable

`func (o *ApiV1ResolveEntityGet200Response) GetPersistenceStable() bool`

GetPersistenceStable returns the PersistenceStable field if non-nil, zero value otherwise.

### GetPersistenceStableOk

`func (o *ApiV1ResolveEntityGet200Response) GetPersistenceStableOk() (*bool, bool)`

GetPersistenceStableOk returns a tuple with the PersistenceStable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceStable

`func (o *ApiV1ResolveEntityGet200Response) SetPersistenceStable(v bool)`

SetPersistenceStable sets PersistenceStable field to given value.


### SetPersistenceStableNil

`func (o *ApiV1ResolveEntityGet200Response) SetPersistenceStableNil(b bool)`

 SetPersistenceStableNil sets the value for PersistenceStable to be an explicit nil

### UnsetPersistenceStable
`func (o *ApiV1ResolveEntityGet200Response) UnsetPersistenceStable()`

UnsetPersistenceStable ensures that no value is present for PersistenceStable, not even an explicit nil
### GetMessage

`func (o *ApiV1ResolveEntityGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1ResolveEntityGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1ResolveEntityGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1ResolveEntityGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResolutionConfidence

`func (o *ApiV1ResolveEntityGet200Response) GetResolutionConfidence() ApiV1ResolveEntityGet200ResponseResolutionConfidence`

GetResolutionConfidence returns the ResolutionConfidence field if non-nil, zero value otherwise.

### GetResolutionConfidenceOk

`func (o *ApiV1ResolveEntityGet200Response) GetResolutionConfidenceOk() (*ApiV1ResolveEntityGet200ResponseResolutionConfidence, bool)`

GetResolutionConfidenceOk returns a tuple with the ResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionConfidence

`func (o *ApiV1ResolveEntityGet200Response) SetResolutionConfidence(v ApiV1ResolveEntityGet200ResponseResolutionConfidence)`

SetResolutionConfidence sets ResolutionConfidence field to given value.


### GetBuildState

`func (o *ApiV1ResolveEntityGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1ResolveEntityGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1ResolveEntityGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.

### HasBuildState

`func (o *ApiV1ResolveEntityGet200Response) HasBuildState() bool`

HasBuildState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


