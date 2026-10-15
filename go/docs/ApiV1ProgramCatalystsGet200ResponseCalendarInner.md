# ApiV1ProgramCatalystsGet200ResponseCalendarInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** |  | 
**Id** | **string** |  | 
**Date** | **string** |  | 
**NctId** | Pointer to **NullableString** |  | [optional] 
**Title** | **string** |  | 
**Phase** | Pointer to **NullableString** |  | [optional] 
**MatchedDrugIds** | Pointer to **[]interface{}** |  | [optional] 
**MatchedCompanyIds** | Pointer to **[]string** |  | [optional] 
**CanonicalDrugId** | Pointer to **NullableString** |  | [optional] 
**CanonicalOrgId** | Pointer to **NullableString** |  | [optional] 
**DrugName** | Pointer to **NullableString** |  | [optional] 
**CliffYear** | Pointer to **NullableInt32** |  | [optional] 
**CliffBasis** | Pointer to **NullableString** |  | [optional] 
**OrgName** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiV1ProgramCatalystsGet200ResponseCalendarInner

`func NewApiV1ProgramCatalystsGet200ResponseCalendarInner(eventType string, id string, date string, title string, ) *ApiV1ProgramCatalystsGet200ResponseCalendarInner`

NewApiV1ProgramCatalystsGet200ResponseCalendarInner instantiates a new ApiV1ProgramCatalystsGet200ResponseCalendarInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ProgramCatalystsGet200ResponseCalendarInnerWithDefaults

`func NewApiV1ProgramCatalystsGet200ResponseCalendarInnerWithDefaults() *ApiV1ProgramCatalystsGet200ResponseCalendarInner`

NewApiV1ProgramCatalystsGet200ResponseCalendarInnerWithDefaults instantiates a new ApiV1ProgramCatalystsGet200ResponseCalendarInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetId(v string)`

SetId sets Id field to given value.


### GetDate

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetNctId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetNctId() string`

GetNctId returns the NctId field if non-nil, zero value otherwise.

### GetNctIdOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetNctIdOk() (*string, bool)`

GetNctIdOk returns a tuple with the NctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNctId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetNctId(v string)`

SetNctId sets NctId field to given value.

### HasNctId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasNctId() bool`

HasNctId returns a boolean if a field has been set.

### SetNctIdNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetNctIdNil(b bool)`

 SetNctIdNil sets the value for NctId to be an explicit nil

### UnsetNctId
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetNctId()`

UnsetNctId ensures that no value is present for NctId, not even an explicit nil
### GetTitle

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPhase

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetMatchedDrugIds

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetMatchedDrugIds() []interface{}`

GetMatchedDrugIds returns the MatchedDrugIds field if non-nil, zero value otherwise.

### GetMatchedDrugIdsOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetMatchedDrugIdsOk() (*[]interface{}, bool)`

GetMatchedDrugIdsOk returns a tuple with the MatchedDrugIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedDrugIds

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetMatchedDrugIds(v []interface{})`

SetMatchedDrugIds sets MatchedDrugIds field to given value.

### HasMatchedDrugIds

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasMatchedDrugIds() bool`

HasMatchedDrugIds returns a boolean if a field has been set.

### SetMatchedDrugIdsNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetMatchedDrugIdsNil(b bool)`

 SetMatchedDrugIdsNil sets the value for MatchedDrugIds to be an explicit nil

### UnsetMatchedDrugIds
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetMatchedDrugIds()`

UnsetMatchedDrugIds ensures that no value is present for MatchedDrugIds, not even an explicit nil
### GetMatchedCompanyIds

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetMatchedCompanyIds() []string`

GetMatchedCompanyIds returns the MatchedCompanyIds field if non-nil, zero value otherwise.

### GetMatchedCompanyIdsOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetMatchedCompanyIdsOk() (*[]string, bool)`

GetMatchedCompanyIdsOk returns a tuple with the MatchedCompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCompanyIds

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetMatchedCompanyIds(v []string)`

SetMatchedCompanyIds sets MatchedCompanyIds field to given value.

### HasMatchedCompanyIds

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasMatchedCompanyIds() bool`

HasMatchedCompanyIds returns a boolean if a field has been set.

### SetMatchedCompanyIdsNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetMatchedCompanyIdsNil(b bool)`

 SetMatchedCompanyIdsNil sets the value for MatchedCompanyIds to be an explicit nil

### UnsetMatchedCompanyIds
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetMatchedCompanyIds()`

UnsetMatchedCompanyIds ensures that no value is present for MatchedCompanyIds, not even an explicit nil
### GetCanonicalDrugId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetCanonicalDrugId() string`

GetCanonicalDrugId returns the CanonicalDrugId field if non-nil, zero value otherwise.

### GetCanonicalDrugIdOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetCanonicalDrugIdOk() (*string, bool)`

GetCanonicalDrugIdOk returns a tuple with the CanonicalDrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalDrugId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetCanonicalDrugId(v string)`

SetCanonicalDrugId sets CanonicalDrugId field to given value.

### HasCanonicalDrugId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasCanonicalDrugId() bool`

HasCanonicalDrugId returns a boolean if a field has been set.

### SetCanonicalDrugIdNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetCanonicalDrugIdNil(b bool)`

 SetCanonicalDrugIdNil sets the value for CanonicalDrugId to be an explicit nil

### UnsetCanonicalDrugId
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetCanonicalDrugId()`

UnsetCanonicalDrugId ensures that no value is present for CanonicalDrugId, not even an explicit nil
### GetCanonicalOrgId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetCanonicalOrgId() string`

GetCanonicalOrgId returns the CanonicalOrgId field if non-nil, zero value otherwise.

### GetCanonicalOrgIdOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetCanonicalOrgIdOk() (*string, bool)`

GetCanonicalOrgIdOk returns a tuple with the CanonicalOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalOrgId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetCanonicalOrgId(v string)`

SetCanonicalOrgId sets CanonicalOrgId field to given value.

### HasCanonicalOrgId

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasCanonicalOrgId() bool`

HasCanonicalOrgId returns a boolean if a field has been set.

### SetCanonicalOrgIdNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetCanonicalOrgIdNil(b bool)`

 SetCanonicalOrgIdNil sets the value for CanonicalOrgId to be an explicit nil

### UnsetCanonicalOrgId
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetCanonicalOrgId()`

UnsetCanonicalOrgId ensures that no value is present for CanonicalOrgId, not even an explicit nil
### GetDrugName

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetCliffYear

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetCliffYear() int32`

GetCliffYear returns the CliffYear field if non-nil, zero value otherwise.

### GetCliffYearOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetCliffYearOk() (*int32, bool)`

GetCliffYearOk returns a tuple with the CliffYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffYear

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetCliffYear(v int32)`

SetCliffYear sets CliffYear field to given value.

### HasCliffYear

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasCliffYear() bool`

HasCliffYear returns a boolean if a field has been set.

### SetCliffYearNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetCliffYearNil(b bool)`

 SetCliffYearNil sets the value for CliffYear to be an explicit nil

### UnsetCliffYear
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetCliffYear()`

UnsetCliffYear ensures that no value is present for CliffYear, not even an explicit nil
### GetCliffBasis

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetCliffBasis() string`

GetCliffBasis returns the CliffBasis field if non-nil, zero value otherwise.

### GetCliffBasisOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetCliffBasisOk() (*string, bool)`

GetCliffBasisOk returns a tuple with the CliffBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliffBasis

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetCliffBasis(v string)`

SetCliffBasis sets CliffBasis field to given value.

### HasCliffBasis

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasCliffBasis() bool`

HasCliffBasis returns a boolean if a field has been set.

### SetCliffBasisNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetCliffBasisNil(b bool)`

 SetCliffBasisNil sets the value for CliffBasis to be an explicit nil

### UnsetCliffBasis
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetCliffBasis()`

UnsetCliffBasis ensures that no value is present for CliffBasis, not even an explicit nil
### GetOrgName

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgNameNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetOrgNameNil(b bool)`

 SetOrgNameNil sets the value for OrgName to be an explicit nil

### UnsetOrgName
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetOrgName()`

UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil
### GetSource

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ApiV1ProgramCatalystsGet200ResponseCalendarInner) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


