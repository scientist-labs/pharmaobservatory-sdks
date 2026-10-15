# ApiV1InvestorProfileGet200ResponseResolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verdict** | **string** |  | 
**MatchTier** | **string** |  | 
**Message** | **interface{}** |  | 
**Candidates** | **[]interface{}** |  | 

## Methods

### NewApiV1InvestorProfileGet200ResponseResolution

`func NewApiV1InvestorProfileGet200ResponseResolution(verdict string, matchTier string, message interface{}, candidates []interface{}, ) *ApiV1InvestorProfileGet200ResponseResolution`

NewApiV1InvestorProfileGet200ResponseResolution instantiates a new ApiV1InvestorProfileGet200ResponseResolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1InvestorProfileGet200ResponseResolutionWithDefaults

`func NewApiV1InvestorProfileGet200ResponseResolutionWithDefaults() *ApiV1InvestorProfileGet200ResponseResolution`

NewApiV1InvestorProfileGet200ResponseResolutionWithDefaults instantiates a new ApiV1InvestorProfileGet200ResponseResolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerdict

`func (o *ApiV1InvestorProfileGet200ResponseResolution) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ApiV1InvestorProfileGet200ResponseResolution) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ApiV1InvestorProfileGet200ResponseResolution) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetMatchTier

`func (o *ApiV1InvestorProfileGet200ResponseResolution) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *ApiV1InvestorProfileGet200ResponseResolution) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *ApiV1InvestorProfileGet200ResponseResolution) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.


### GetMessage

`func (o *ApiV1InvestorProfileGet200ResponseResolution) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1InvestorProfileGet200ResponseResolution) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1InvestorProfileGet200ResponseResolution) SetMessage(v interface{})`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *ApiV1InvestorProfileGet200ResponseResolution) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ApiV1InvestorProfileGet200ResponseResolution) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCandidates

`func (o *ApiV1InvestorProfileGet200ResponseResolution) GetCandidates() []interface{}`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *ApiV1InvestorProfileGet200ResponseResolution) GetCandidatesOk() (*[]interface{}, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *ApiV1InvestorProfileGet200ResponseResolution) SetCandidates(v []interface{})`

SetCandidates sets Candidates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


