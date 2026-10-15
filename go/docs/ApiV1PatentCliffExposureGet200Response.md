# ApiV1PatentCliffExposureGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrugIds** | **[]string** |  | 
**CompanyIds** | **[]string** |  | 
**Exposures** | [**[]ApiV1PatentCliffExposureGet200ResponseExposuresInner**](ApiV1PatentCliffExposureGet200ResponseExposuresInner.md) |  | 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1PatentCliffExposureGet200Response

`func NewApiV1PatentCliffExposureGet200Response(drugIds []string, companyIds []string, exposures []ApiV1PatentCliffExposureGet200ResponseExposuresInner, buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1PatentCliffExposureGet200Response`

NewApiV1PatentCliffExposureGet200Response instantiates a new ApiV1PatentCliffExposureGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1PatentCliffExposureGet200ResponseWithDefaults

`func NewApiV1PatentCliffExposureGet200ResponseWithDefaults() *ApiV1PatentCliffExposureGet200Response`

NewApiV1PatentCliffExposureGet200ResponseWithDefaults instantiates a new ApiV1PatentCliffExposureGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrugIds

`func (o *ApiV1PatentCliffExposureGet200Response) GetDrugIds() []string`

GetDrugIds returns the DrugIds field if non-nil, zero value otherwise.

### GetDrugIdsOk

`func (o *ApiV1PatentCliffExposureGet200Response) GetDrugIdsOk() (*[]string, bool)`

GetDrugIdsOk returns a tuple with the DrugIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugIds

`func (o *ApiV1PatentCliffExposureGet200Response) SetDrugIds(v []string)`

SetDrugIds sets DrugIds field to given value.


### GetCompanyIds

`func (o *ApiV1PatentCliffExposureGet200Response) GetCompanyIds() []string`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *ApiV1PatentCliffExposureGet200Response) GetCompanyIdsOk() (*[]string, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *ApiV1PatentCliffExposureGet200Response) SetCompanyIds(v []string)`

SetCompanyIds sets CompanyIds field to given value.


### GetExposures

`func (o *ApiV1PatentCliffExposureGet200Response) GetExposures() []ApiV1PatentCliffExposureGet200ResponseExposuresInner`

GetExposures returns the Exposures field if non-nil, zero value otherwise.

### GetExposuresOk

`func (o *ApiV1PatentCliffExposureGet200Response) GetExposuresOk() (*[]ApiV1PatentCliffExposureGet200ResponseExposuresInner, bool)`

GetExposuresOk returns a tuple with the Exposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposures

`func (o *ApiV1PatentCliffExposureGet200Response) SetExposures(v []ApiV1PatentCliffExposureGet200ResponseExposuresInner)`

SetExposures sets Exposures field to given value.


### GetBuildState

`func (o *ApiV1PatentCliffExposureGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1PatentCliffExposureGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1PatentCliffExposureGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetError

`func (o *ApiV1PatentCliffExposureGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1PatentCliffExposureGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1PatentCliffExposureGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1PatentCliffExposureGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1PatentCliffExposureGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1PatentCliffExposureGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1PatentCliffExposureGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1PatentCliffExposureGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


