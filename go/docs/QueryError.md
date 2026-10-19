# QueryError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | 
**Message** | **string** |  | 
**Hint** | Pointer to **NullableString** |  | [optional] 
**Docs** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryError

`func NewQueryError(errorCode string, message string, ) *QueryError`

NewQueryError instantiates a new QueryError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryErrorWithDefaults

`func NewQueryErrorWithDefaults() *QueryError`

NewQueryErrorWithDefaults instantiates a new QueryError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *QueryError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *QueryError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *QueryError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *QueryError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QueryError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QueryError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetHint

`func (o *QueryError) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *QueryError) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *QueryError) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *QueryError) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *QueryError) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *QueryError) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetDocs

`func (o *QueryError) GetDocs() string`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *QueryError) GetDocsOk() (*string, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *QueryError) SetDocs(v string)`

SetDocs sets Docs field to given value.

### HasDocs

`func (o *QueryError) HasDocs() bool`

HasDocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


