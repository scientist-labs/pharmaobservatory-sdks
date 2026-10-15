# ApiV1InvestorProfileGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**InvestorId** | **NullableString** |  | 
**Profile** | [**NullableApiV1InvestorProfileGet200ResponseProfile**](ApiV1InvestorProfileGet200ResponseProfile.md) |  | 
**Funds** | [**ApiV1InvestorProfileGet200ResponseFunds**](ApiV1InvestorProfileGet200ResponseFunds.md) |  | 
**Ownership** | [**ApiV1InvestorProfileGet200ResponseOwnership**](ApiV1InvestorProfileGet200ResponseOwnership.md) |  | 
**Semantics** | [**ApiV1InvestorProfileGet200ResponseSemantics**](ApiV1InvestorProfileGet200ResponseSemantics.md) |  | 
**BuildState** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  | 
**Sections** | [**ApiV1InvestorProfileGet200ResponseSections**](ApiV1InvestorProfileGet200ResponseSections.md) |  | 
**InvestorName** | Pointer to **string** |  | [optional] 
**Resolution** | Pointer to [**ApiV1InvestorProfileGet200ResponseResolution**](ApiV1InvestorProfileGet200ResponseResolution.md) |  | [optional] 

## Methods

### NewApiV1InvestorProfileGet200Response

`func NewApiV1InvestorProfileGet200Response(investorId NullableString, profile NullableApiV1InvestorProfileGet200ResponseProfile, funds ApiV1InvestorProfileGet200ResponseFunds, ownership ApiV1InvestorProfileGet200ResponseOwnership, semantics ApiV1InvestorProfileGet200ResponseSemantics, buildState ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, sections ApiV1InvestorProfileGet200ResponseSections, ) *ApiV1InvestorProfileGet200Response`

NewApiV1InvestorProfileGet200Response instantiates a new ApiV1InvestorProfileGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1InvestorProfileGet200ResponseWithDefaults

`func NewApiV1InvestorProfileGet200ResponseWithDefaults() *ApiV1InvestorProfileGet200Response`

NewApiV1InvestorProfileGet200ResponseWithDefaults instantiates a new ApiV1InvestorProfileGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApiV1InvestorProfileGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1InvestorProfileGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1InvestorProfileGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1InvestorProfileGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1InvestorProfileGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1InvestorProfileGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1InvestorProfileGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1InvestorProfileGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetInvestorId

`func (o *ApiV1InvestorProfileGet200Response) GetInvestorId() string`

GetInvestorId returns the InvestorId field if non-nil, zero value otherwise.

### GetInvestorIdOk

`func (o *ApiV1InvestorProfileGet200Response) GetInvestorIdOk() (*string, bool)`

GetInvestorIdOk returns a tuple with the InvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorId

`func (o *ApiV1InvestorProfileGet200Response) SetInvestorId(v string)`

SetInvestorId sets InvestorId field to given value.


### SetInvestorIdNil

`func (o *ApiV1InvestorProfileGet200Response) SetInvestorIdNil(b bool)`

 SetInvestorIdNil sets the value for InvestorId to be an explicit nil

### UnsetInvestorId
`func (o *ApiV1InvestorProfileGet200Response) UnsetInvestorId()`

UnsetInvestorId ensures that no value is present for InvestorId, not even an explicit nil
### GetProfile

`func (o *ApiV1InvestorProfileGet200Response) GetProfile() ApiV1InvestorProfileGet200ResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ApiV1InvestorProfileGet200Response) GetProfileOk() (*ApiV1InvestorProfileGet200ResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ApiV1InvestorProfileGet200Response) SetProfile(v ApiV1InvestorProfileGet200ResponseProfile)`

SetProfile sets Profile field to given value.


### SetProfileNil

`func (o *ApiV1InvestorProfileGet200Response) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *ApiV1InvestorProfileGet200Response) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetFunds

`func (o *ApiV1InvestorProfileGet200Response) GetFunds() ApiV1InvestorProfileGet200ResponseFunds`

GetFunds returns the Funds field if non-nil, zero value otherwise.

### GetFundsOk

`func (o *ApiV1InvestorProfileGet200Response) GetFundsOk() (*ApiV1InvestorProfileGet200ResponseFunds, bool)`

GetFundsOk returns a tuple with the Funds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunds

`func (o *ApiV1InvestorProfileGet200Response) SetFunds(v ApiV1InvestorProfileGet200ResponseFunds)`

SetFunds sets Funds field to given value.


### GetOwnership

`func (o *ApiV1InvestorProfileGet200Response) GetOwnership() ApiV1InvestorProfileGet200ResponseOwnership`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *ApiV1InvestorProfileGet200Response) GetOwnershipOk() (*ApiV1InvestorProfileGet200ResponseOwnership, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *ApiV1InvestorProfileGet200Response) SetOwnership(v ApiV1InvestorProfileGet200ResponseOwnership)`

SetOwnership sets Ownership field to given value.


### GetSemantics

`func (o *ApiV1InvestorProfileGet200Response) GetSemantics() ApiV1InvestorProfileGet200ResponseSemantics`

GetSemantics returns the Semantics field if non-nil, zero value otherwise.

### GetSemanticsOk

`func (o *ApiV1InvestorProfileGet200Response) GetSemanticsOk() (*ApiV1InvestorProfileGet200ResponseSemantics, bool)`

GetSemanticsOk returns a tuple with the Semantics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantics

`func (o *ApiV1InvestorProfileGet200Response) SetSemantics(v ApiV1InvestorProfileGet200ResponseSemantics)`

SetSemantics sets Semantics field to given value.


### GetBuildState

`func (o *ApiV1InvestorProfileGet200Response) GetBuildState() ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1InvestorProfileGet200Response) GetBuildStateOk() (*ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1InvestorProfileGet200Response) SetBuildState(v ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetSections

`func (o *ApiV1InvestorProfileGet200Response) GetSections() ApiV1InvestorProfileGet200ResponseSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ApiV1InvestorProfileGet200Response) GetSectionsOk() (*ApiV1InvestorProfileGet200ResponseSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ApiV1InvestorProfileGet200Response) SetSections(v ApiV1InvestorProfileGet200ResponseSections)`

SetSections sets Sections field to given value.


### GetInvestorName

`func (o *ApiV1InvestorProfileGet200Response) GetInvestorName() string`

GetInvestorName returns the InvestorName field if non-nil, zero value otherwise.

### GetInvestorNameOk

`func (o *ApiV1InvestorProfileGet200Response) GetInvestorNameOk() (*string, bool)`

GetInvestorNameOk returns a tuple with the InvestorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorName

`func (o *ApiV1InvestorProfileGet200Response) SetInvestorName(v string)`

SetInvestorName sets InvestorName field to given value.

### HasInvestorName

`func (o *ApiV1InvestorProfileGet200Response) HasInvestorName() bool`

HasInvestorName returns a boolean if a field has been set.

### GetResolution

`func (o *ApiV1InvestorProfileGet200Response) GetResolution() ApiV1InvestorProfileGet200ResponseResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ApiV1InvestorProfileGet200Response) GetResolutionOk() (*ApiV1InvestorProfileGet200ResponseResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ApiV1InvestorProfileGet200Response) SetResolution(v ApiV1InvestorProfileGet200ResponseResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ApiV1InvestorProfileGet200Response) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


