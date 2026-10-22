# BiomarkerTraversalStatusRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The biomarker&#39;s route key, the entry key for the biomarkers object type. | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The biomarker&#39;s own key. | [optional] 
**BiomarkerKind** | Pointer to **NullableString** | What kind of biomarker the row describes. | [optional] 
**TraversalKey** | Pointer to **NullableString** | The traversal. | [optional] 
**DisplayOrder** | Pointer to **NullableInt32** | The order the web page lists the traversals in, one to nine. | [optional] 
**TraversalName** | Pointer to **NullableString** | The traversal&#39;s display name. | [optional] 
**ApplicabilityState** | Pointer to **NullableString** | applicable: measured for this kind of biomarker. inapplicable: does not apply to this kind. not_measured: applies in principle but the corpus is not mined for this kind yet. | [optional] 
**ContextMode** | Pointer to **NullableString** | How the traversal reads for this biomarker. direct: its own rows. parent_gene: the variant&#39;s gene rows. parent_and_exact: the gene&#39;s rows plus the exact variant&#39;s. child_variants: the gene&#39;s variant children. source_linked: shown when a source assertion resolves to it. bridge_deferred: depends on a bridge that has not landed. none: no traversal. | [optional] 
**ReasonCode** | Pointer to **NullableString** | The code behind the applicability state and context mode; state_detail is its sentence. | [optional] 
**StateDetail** | Pointer to **NullableString** | The sentence the web page prints for the reason code. | [optional] 

## Methods

### NewBiomarkerTraversalStatusRow

`func NewBiomarkerTraversalStatusRow() *BiomarkerTraversalStatusRow`

NewBiomarkerTraversalStatusRow instantiates a new BiomarkerTraversalStatusRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerTraversalStatusRowWithDefaults

`func NewBiomarkerTraversalStatusRowWithDefaults() *BiomarkerTraversalStatusRow`

NewBiomarkerTraversalStatusRowWithDefaults instantiates a new BiomarkerTraversalStatusRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerTraversalStatusRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerTraversalStatusRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerTraversalStatusRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerTraversalStatusRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerTraversalStatusRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerTraversalStatusRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerId

`func (o *BiomarkerTraversalStatusRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkerTraversalStatusRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkerTraversalStatusRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkerTraversalStatusRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkerTraversalStatusRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkerTraversalStatusRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetBiomarkerKind

`func (o *BiomarkerTraversalStatusRow) GetBiomarkerKind() string`

GetBiomarkerKind returns the BiomarkerKind field if non-nil, zero value otherwise.

### GetBiomarkerKindOk

`func (o *BiomarkerTraversalStatusRow) GetBiomarkerKindOk() (*string, bool)`

GetBiomarkerKindOk returns a tuple with the BiomarkerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerKind

`func (o *BiomarkerTraversalStatusRow) SetBiomarkerKind(v string)`

SetBiomarkerKind sets BiomarkerKind field to given value.

### HasBiomarkerKind

`func (o *BiomarkerTraversalStatusRow) HasBiomarkerKind() bool`

HasBiomarkerKind returns a boolean if a field has been set.

### SetBiomarkerKindNil

`func (o *BiomarkerTraversalStatusRow) SetBiomarkerKindNil(b bool)`

 SetBiomarkerKindNil sets the value for BiomarkerKind to be an explicit nil

### UnsetBiomarkerKind
`func (o *BiomarkerTraversalStatusRow) UnsetBiomarkerKind()`

UnsetBiomarkerKind ensures that no value is present for BiomarkerKind, not even an explicit nil
### GetTraversalKey

`func (o *BiomarkerTraversalStatusRow) GetTraversalKey() string`

GetTraversalKey returns the TraversalKey field if non-nil, zero value otherwise.

### GetTraversalKeyOk

`func (o *BiomarkerTraversalStatusRow) GetTraversalKeyOk() (*string, bool)`

GetTraversalKeyOk returns a tuple with the TraversalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraversalKey

`func (o *BiomarkerTraversalStatusRow) SetTraversalKey(v string)`

SetTraversalKey sets TraversalKey field to given value.

### HasTraversalKey

`func (o *BiomarkerTraversalStatusRow) HasTraversalKey() bool`

HasTraversalKey returns a boolean if a field has been set.

### SetTraversalKeyNil

`func (o *BiomarkerTraversalStatusRow) SetTraversalKeyNil(b bool)`

 SetTraversalKeyNil sets the value for TraversalKey to be an explicit nil

### UnsetTraversalKey
`func (o *BiomarkerTraversalStatusRow) UnsetTraversalKey()`

UnsetTraversalKey ensures that no value is present for TraversalKey, not even an explicit nil
### GetDisplayOrder

`func (o *BiomarkerTraversalStatusRow) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *BiomarkerTraversalStatusRow) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *BiomarkerTraversalStatusRow) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *BiomarkerTraversalStatusRow) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *BiomarkerTraversalStatusRow) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *BiomarkerTraversalStatusRow) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetTraversalName

`func (o *BiomarkerTraversalStatusRow) GetTraversalName() string`

GetTraversalName returns the TraversalName field if non-nil, zero value otherwise.

### GetTraversalNameOk

`func (o *BiomarkerTraversalStatusRow) GetTraversalNameOk() (*string, bool)`

GetTraversalNameOk returns a tuple with the TraversalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraversalName

`func (o *BiomarkerTraversalStatusRow) SetTraversalName(v string)`

SetTraversalName sets TraversalName field to given value.

### HasTraversalName

`func (o *BiomarkerTraversalStatusRow) HasTraversalName() bool`

HasTraversalName returns a boolean if a field has been set.

### SetTraversalNameNil

`func (o *BiomarkerTraversalStatusRow) SetTraversalNameNil(b bool)`

 SetTraversalNameNil sets the value for TraversalName to be an explicit nil

### UnsetTraversalName
`func (o *BiomarkerTraversalStatusRow) UnsetTraversalName()`

UnsetTraversalName ensures that no value is present for TraversalName, not even an explicit nil
### GetApplicabilityState

`func (o *BiomarkerTraversalStatusRow) GetApplicabilityState() string`

GetApplicabilityState returns the ApplicabilityState field if non-nil, zero value otherwise.

### GetApplicabilityStateOk

`func (o *BiomarkerTraversalStatusRow) GetApplicabilityStateOk() (*string, bool)`

GetApplicabilityStateOk returns a tuple with the ApplicabilityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicabilityState

`func (o *BiomarkerTraversalStatusRow) SetApplicabilityState(v string)`

SetApplicabilityState sets ApplicabilityState field to given value.

### HasApplicabilityState

`func (o *BiomarkerTraversalStatusRow) HasApplicabilityState() bool`

HasApplicabilityState returns a boolean if a field has been set.

### SetApplicabilityStateNil

`func (o *BiomarkerTraversalStatusRow) SetApplicabilityStateNil(b bool)`

 SetApplicabilityStateNil sets the value for ApplicabilityState to be an explicit nil

### UnsetApplicabilityState
`func (o *BiomarkerTraversalStatusRow) UnsetApplicabilityState()`

UnsetApplicabilityState ensures that no value is present for ApplicabilityState, not even an explicit nil
### GetContextMode

`func (o *BiomarkerTraversalStatusRow) GetContextMode() string`

GetContextMode returns the ContextMode field if non-nil, zero value otherwise.

### GetContextModeOk

`func (o *BiomarkerTraversalStatusRow) GetContextModeOk() (*string, bool)`

GetContextModeOk returns a tuple with the ContextMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextMode

`func (o *BiomarkerTraversalStatusRow) SetContextMode(v string)`

SetContextMode sets ContextMode field to given value.

### HasContextMode

`func (o *BiomarkerTraversalStatusRow) HasContextMode() bool`

HasContextMode returns a boolean if a field has been set.

### SetContextModeNil

`func (o *BiomarkerTraversalStatusRow) SetContextModeNil(b bool)`

 SetContextModeNil sets the value for ContextMode to be an explicit nil

### UnsetContextMode
`func (o *BiomarkerTraversalStatusRow) UnsetContextMode()`

UnsetContextMode ensures that no value is present for ContextMode, not even an explicit nil
### GetReasonCode

`func (o *BiomarkerTraversalStatusRow) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *BiomarkerTraversalStatusRow) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *BiomarkerTraversalStatusRow) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *BiomarkerTraversalStatusRow) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *BiomarkerTraversalStatusRow) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *BiomarkerTraversalStatusRow) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetStateDetail

`func (o *BiomarkerTraversalStatusRow) GetStateDetail() string`

GetStateDetail returns the StateDetail field if non-nil, zero value otherwise.

### GetStateDetailOk

`func (o *BiomarkerTraversalStatusRow) GetStateDetailOk() (*string, bool)`

GetStateDetailOk returns a tuple with the StateDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDetail

`func (o *BiomarkerTraversalStatusRow) SetStateDetail(v string)`

SetStateDetail sets StateDetail field to given value.

### HasStateDetail

`func (o *BiomarkerTraversalStatusRow) HasStateDetail() bool`

HasStateDetail returns a boolean if a field has been set.

### SetStateDetailNil

`func (o *BiomarkerTraversalStatusRow) SetStateDetailNil(b bool)`

 SetStateDetailNil sets the value for StateDetail to be an explicit nil

### UnsetStateDetail
`func (o *BiomarkerTraversalStatusRow) UnsetStateDetail()`

UnsetStateDetail ensures that no value is present for StateDetail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


