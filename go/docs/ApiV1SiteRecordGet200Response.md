# ApiV1SiteRecordGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] 
**Requested** | Pointer to **string** |  | [optional] 
**MatchedCount** | Pointer to **int32** |  | [optional] 
**ReturnedCount** | Pointer to **int32** |  | [optional] 
**Sites** | Pointer to [**[]ApiV1SiteRecordGet200ResponseSitesInner**](ApiV1SiteRecordGet200ResponseSitesInner.md) |  | [optional] 
**BuildState** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  | 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**DetailSections** | Pointer to [**ApiV1SiteRecordGet200ResponseDetailSections**](ApiV1SiteRecordGet200ResponseDetailSections.md) |  | [optional] 
**SiteKeyRedirect** | Pointer to [**ApiV1SiteRecordGet200ResponseSiteKeyRedirect**](ApiV1SiteRecordGet200ResponseSiteKeyRedirect.md) |  | [optional] 

## Methods

### NewApiV1SiteRecordGet200Response

`func NewApiV1SiteRecordGet200Response(buildState ApiV1BiomarkerAssayScreenGet200ResponseBuildState, ) *ApiV1SiteRecordGet200Response`

NewApiV1SiteRecordGet200Response instantiates a new ApiV1SiteRecordGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1SiteRecordGet200ResponseWithDefaults

`func NewApiV1SiteRecordGet200ResponseWithDefaults() *ApiV1SiteRecordGet200Response`

NewApiV1SiteRecordGet200ResponseWithDefaults instantiates a new ApiV1SiteRecordGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ApiV1SiteRecordGet200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApiV1SiteRecordGet200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApiV1SiteRecordGet200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ApiV1SiteRecordGet200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRequested

`func (o *ApiV1SiteRecordGet200Response) GetRequested() string`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *ApiV1SiteRecordGet200Response) GetRequestedOk() (*string, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *ApiV1SiteRecordGet200Response) SetRequested(v string)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *ApiV1SiteRecordGet200Response) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetMatchedCount

`func (o *ApiV1SiteRecordGet200Response) GetMatchedCount() int32`

GetMatchedCount returns the MatchedCount field if non-nil, zero value otherwise.

### GetMatchedCountOk

`func (o *ApiV1SiteRecordGet200Response) GetMatchedCountOk() (*int32, bool)`

GetMatchedCountOk returns a tuple with the MatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCount

`func (o *ApiV1SiteRecordGet200Response) SetMatchedCount(v int32)`

SetMatchedCount sets MatchedCount field to given value.

### HasMatchedCount

`func (o *ApiV1SiteRecordGet200Response) HasMatchedCount() bool`

HasMatchedCount returns a boolean if a field has been set.

### GetReturnedCount

`func (o *ApiV1SiteRecordGet200Response) GetReturnedCount() int32`

GetReturnedCount returns the ReturnedCount field if non-nil, zero value otherwise.

### GetReturnedCountOk

`func (o *ApiV1SiteRecordGet200Response) GetReturnedCountOk() (*int32, bool)`

GetReturnedCountOk returns a tuple with the ReturnedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCount

`func (o *ApiV1SiteRecordGet200Response) SetReturnedCount(v int32)`

SetReturnedCount sets ReturnedCount field to given value.

### HasReturnedCount

`func (o *ApiV1SiteRecordGet200Response) HasReturnedCount() bool`

HasReturnedCount returns a boolean if a field has been set.

### GetSites

`func (o *ApiV1SiteRecordGet200Response) GetSites() []ApiV1SiteRecordGet200ResponseSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ApiV1SiteRecordGet200Response) GetSitesOk() (*[]ApiV1SiteRecordGet200ResponseSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ApiV1SiteRecordGet200Response) SetSites(v []ApiV1SiteRecordGet200ResponseSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ApiV1SiteRecordGet200Response) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetBuildState

`func (o *ApiV1SiteRecordGet200Response) GetBuildState() ApiV1BiomarkerAssayScreenGet200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *ApiV1SiteRecordGet200Response) GetBuildStateOk() (*ApiV1BiomarkerAssayScreenGet200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *ApiV1SiteRecordGet200Response) SetBuildState(v ApiV1BiomarkerAssayScreenGet200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetError

`func (o *ApiV1SiteRecordGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1SiteRecordGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1SiteRecordGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1SiteRecordGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ApiV1SiteRecordGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV1SiteRecordGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV1SiteRecordGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV1SiteRecordGet200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetailSections

`func (o *ApiV1SiteRecordGet200Response) GetDetailSections() ApiV1SiteRecordGet200ResponseDetailSections`

GetDetailSections returns the DetailSections field if non-nil, zero value otherwise.

### GetDetailSectionsOk

`func (o *ApiV1SiteRecordGet200Response) GetDetailSectionsOk() (*ApiV1SiteRecordGet200ResponseDetailSections, bool)`

GetDetailSectionsOk returns a tuple with the DetailSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailSections

`func (o *ApiV1SiteRecordGet200Response) SetDetailSections(v ApiV1SiteRecordGet200ResponseDetailSections)`

SetDetailSections sets DetailSections field to given value.

### HasDetailSections

`func (o *ApiV1SiteRecordGet200Response) HasDetailSections() bool`

HasDetailSections returns a boolean if a field has been set.

### GetSiteKeyRedirect

`func (o *ApiV1SiteRecordGet200Response) GetSiteKeyRedirect() ApiV1SiteRecordGet200ResponseSiteKeyRedirect`

GetSiteKeyRedirect returns the SiteKeyRedirect field if non-nil, zero value otherwise.

### GetSiteKeyRedirectOk

`func (o *ApiV1SiteRecordGet200Response) GetSiteKeyRedirectOk() (*ApiV1SiteRecordGet200ResponseSiteKeyRedirect, bool)`

GetSiteKeyRedirectOk returns a tuple with the SiteKeyRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKeyRedirect

`func (o *ApiV1SiteRecordGet200Response) SetSiteKeyRedirect(v ApiV1SiteRecordGet200ResponseSiteKeyRedirect)`

SetSiteKeyRedirect sets SiteKeyRedirect field to given value.

### HasSiteKeyRedirect

`func (o *ApiV1SiteRecordGet200Response) HasSiteKeyRedirect() bool`

HasSiteKeyRedirect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


