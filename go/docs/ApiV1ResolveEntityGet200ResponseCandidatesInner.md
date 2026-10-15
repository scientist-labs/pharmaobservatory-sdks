# ApiV1ResolveEntityGet200ResponseCandidatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DisplayName** | **NullableString** |  | 
**Subtitle** | **NullableString** |  | 
**Restriction** | [**NullableApiV1ResolveEntityGet200ResponseCandidatesInnerRestriction**](ApiV1ResolveEntityGet200ResponseCandidatesInnerRestriction.md) |  | 
**NameWithheldNote** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiV1ResolveEntityGet200ResponseCandidatesInner

`func NewApiV1ResolveEntityGet200ResponseCandidatesInner(id string, displayName NullableString, subtitle NullableString, restriction NullableApiV1ResolveEntityGet200ResponseCandidatesInnerRestriction, ) *ApiV1ResolveEntityGet200ResponseCandidatesInner`

NewApiV1ResolveEntityGet200ResponseCandidatesInner instantiates a new ApiV1ResolveEntityGet200ResponseCandidatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ResolveEntityGet200ResponseCandidatesInnerWithDefaults

`func NewApiV1ResolveEntityGet200ResponseCandidatesInnerWithDefaults() *ApiV1ResolveEntityGet200ResponseCandidatesInner`

NewApiV1ResolveEntityGet200ResponseCandidatesInnerWithDefaults instantiates a new ApiV1ResolveEntityGet200ResponseCandidatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSubtitle

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.


### SetSubtitleNil

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
### GetRestriction

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetRestriction() ApiV1ResolveEntityGet200ResponseCandidatesInnerRestriction`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetRestrictionOk() (*ApiV1ResolveEntityGet200ResponseCandidatesInnerRestriction, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) SetRestriction(v ApiV1ResolveEntityGet200ResponseCandidatesInnerRestriction)`

SetRestriction sets Restriction field to given value.


### SetRestrictionNil

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) SetRestrictionNil(b bool)`

 SetRestrictionNil sets the value for Restriction to be an explicit nil

### UnsetRestriction
`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) UnsetRestriction()`

UnsetRestriction ensures that no value is present for Restriction, not even an explicit nil
### GetNameWithheldNote

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetNameWithheldNote() string`

GetNameWithheldNote returns the NameWithheldNote field if non-nil, zero value otherwise.

### GetNameWithheldNoteOk

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) GetNameWithheldNoteOk() (*string, bool)`

GetNameWithheldNoteOk returns a tuple with the NameWithheldNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameWithheldNote

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) SetNameWithheldNote(v string)`

SetNameWithheldNote sets NameWithheldNote field to given value.

### HasNameWithheldNote

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) HasNameWithheldNote() bool`

HasNameWithheldNote returns a boolean if a field has been set.

### SetNameWithheldNoteNil

`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) SetNameWithheldNoteNil(b bool)`

 SetNameWithheldNoteNil sets the value for NameWithheldNote to be an explicit nil

### UnsetNameWithheldNote
`func (o *ApiV1ResolveEntityGet200ResponseCandidatesInner) UnsetNameWithheldNote()`

UnsetNameWithheldNote ensures that no value is present for NameWithheldNote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


