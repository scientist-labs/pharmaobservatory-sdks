# ApiV1ProgramCatalystsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrugIds** | **[]string** |  | 
**CompanyIds** | **[]string** |  | 
**Calendar** | [**[]ApiV1ProgramCatalystsGet200ResponseCalendarInner**](ApiV1ProgramCatalystsGet200ResponseCalendarInner.md) |  | 
**Pdufa** | [**NullableApiV1ProgramCatalystsGet200ResponsePdufa**](ApiV1ProgramCatalystsGet200ResponsePdufa.md) |  | 
**BuildState** | [**ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState.md) |  | 
**Sections** | [**ApiV1ProgramCatalystsGet200ResponseSections**](ApiV1ProgramCatalystsGet200ResponseSections.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1ProgramCatalystsGet200Response

`func NewApiV1ProgramCatalystsGet200Response(drugIds []string, companyIds []string, calendar []ApiV1ProgramCatalystsGet200ResponseCalendarInner, pdufa NullableApiV1ProgramCatalystsGet200ResponsePdufa, buildState ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, sections ApiV1ProgramCatalystsGet200ResponseSections, ) *ApiV1ProgramCatalystsGet200Response`

NewApiV1ProgramCatalystsGet200Response instantiates a new ApiV1ProgramCatalystsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ProgramCatalystsGet200ResponseWithDefaults

`func NewApiV1ProgramCatalystsGet200ResponseWithDefaults() *ApiV1ProgramCatalystsGet200Response`

NewApiV1ProgramCatalystsGet200ResponseWithDefaults instantiates a new ApiV1ProgramCatalystsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrugIds

`func (o *ApiV1ProgramCatalystsGet200Response) GetDrugIds() []string`

GetDrugIds returns the DrugIds field if non-nil, zero value otherwise.

### GetDrugIdsOk

`func (o *ApiV1ProgramCatalystsGet200Response) GetDrugIdsOk() (*[]string, bool)`

GetDrugIdsOk returns a tuple with the DrugIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugIds

`func (o *ApiV1ProgramCatalystsGet200Response) SetDrugIds(v []string)`

SetDrugIds sets DrugIds field to given value.


### GetCompanyIds

`func (o *ApiV1ProgramCatalystsGet200Response) GetCompanyIds() []string`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *ApiV1ProgramCatalystsGet200Response) GetCompanyIdsOk() (*[]string, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *ApiV1ProgramCatalystsGet200Response) SetCompanyIds(v []string)`

SetCompanyIds sets CompanyIds field to given value.


### GetCalendar

`func (o *ApiV1ProgramCatalystsGet200Response) GetCalendar() []ApiV1ProgramCatalystsGet200ResponseCalendarInner`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *ApiV1ProgramCatalystsGet200Response) GetCalendarOk() (*[]ApiV1ProgramCatalystsGet200ResponseCalendarInner, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *ApiV1ProgramCatalystsGet200Response) SetCalendar(v []ApiV1ProgramCatalystsGet200ResponseCalendarInner)`

SetCalendar sets Calendar field to given value.


### GetPdufa

`func (o *ApiV1ProgramCatalystsGet200Response) GetPdufa() ApiV1ProgramCatalystsGet200ResponsePdufa`

GetPdufa returns the Pdufa field if non-nil, zero value otherwise.

### GetPdufaOk

`func (o *ApiV1ProgramCatalystsGet200Response) GetPdufaOk() (*ApiV1ProgramCatalystsGet200ResponsePdufa, bool)`

GetPdufaOk returns a tuple with the Pdufa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdufa

`func (o *ApiV1ProgramCatalystsGet200Response) SetPdufa(v ApiV1ProgramCatalystsGet200ResponsePdufa)`

SetPdufa sets Pdufa field to given value.


### SetPdufaNil

`func (o *ApiV1ProgramCatalystsGet200Response) SetPdufaNil(b bool)`

 SetPdufaNil sets the value for Pdufa to be an explicit nil

### UnsetPdufa
`func (o *ApiV1ProgramCatalystsGet200Response) UnsetPdufa()`

UnsetPdufa ensures that no value is present for Pdufa, not even an explicit nil
### GetBuildState

`func (o *ApiV1ProgramCatalystsGet200Response) GetBuildState() ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1ProgramCatalystsGet200Response) GetBuildStateOk() (*ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1ProgramCatalystsGet200Response) SetBuildState(v ApiV1ConditionsConditionIdEvidenceFreshnessGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetSections

`func (o *ApiV1ProgramCatalystsGet200Response) GetSections() ApiV1ProgramCatalystsGet200ResponseSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ApiV1ProgramCatalystsGet200Response) GetSectionsOk() (*ApiV1ProgramCatalystsGet200ResponseSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ApiV1ProgramCatalystsGet200Response) SetSections(v ApiV1ProgramCatalystsGet200ResponseSections)`

SetSections sets Sections field to given value.


### GetError

`func (o *ApiV1ProgramCatalystsGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1ProgramCatalystsGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1ProgramCatalystsGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1ProgramCatalystsGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1ProgramCatalystsGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1ProgramCatalystsGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1ProgramCatalystsGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1ProgramCatalystsGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


