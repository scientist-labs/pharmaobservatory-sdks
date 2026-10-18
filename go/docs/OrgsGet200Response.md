# OrgsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **string** |  | 
**Id** | **string** |  | 
**Identity** | [**OrgsGet200ResponseIdentity**](OrgsGet200ResponseIdentity.md) |  | 
**Sections** | [**[]OrgsGet200ResponseSectionsInner**](OrgsGet200ResponseSectionsInner.md) |  | 
**Attribution** | [**[]BiomarkersList200ResponseAttributionInner**](BiomarkersList200ResponseAttributionInner.md) |  | 
**WebUrl** | **string** |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 
**ResolutionConfidence** | [**BiomarkersGet200ResponseResolutionConfidence**](BiomarkersGet200ResponseResolutionConfidence.md) |  | 

## Methods

### NewOrgsGet200Response

`func NewOrgsGet200Response(page string, id string, identity OrgsGet200ResponseIdentity, sections []OrgsGet200ResponseSectionsInner, attribution []BiomarkersList200ResponseAttributionInner, webUrl string, buildState BiomarkersList200ResponseBuildState, resolutionConfidence BiomarkersGet200ResponseResolutionConfidence, ) *OrgsGet200Response`

NewOrgsGet200Response instantiates a new OrgsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsGet200ResponseWithDefaults

`func NewOrgsGet200ResponseWithDefaults() *OrgsGet200Response`

NewOrgsGet200ResponseWithDefaults instantiates a new OrgsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrgsGet200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgsGet200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgsGet200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *OrgsGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgsGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgsGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetIdentity

`func (o *OrgsGet200Response) GetIdentity() OrgsGet200ResponseIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *OrgsGet200Response) GetIdentityOk() (*OrgsGet200ResponseIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *OrgsGet200Response) SetIdentity(v OrgsGet200ResponseIdentity)`

SetIdentity sets Identity field to given value.


### GetSections

`func (o *OrgsGet200Response) GetSections() []OrgsGet200ResponseSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *OrgsGet200Response) GetSectionsOk() (*[]OrgsGet200ResponseSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *OrgsGet200Response) SetSections(v []OrgsGet200ResponseSectionsInner)`

SetSections sets Sections field to given value.


### GetAttribution

`func (o *OrgsGet200Response) GetAttribution() []BiomarkersList200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsGet200Response) GetAttributionOk() (*[]BiomarkersList200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsGet200Response) SetAttribution(v []BiomarkersList200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetWebUrl

`func (o *OrgsGet200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OrgsGet200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OrgsGet200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetBuildState

`func (o *OrgsGet200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *OrgsGet200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *OrgsGet200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetResolutionConfidence

`func (o *OrgsGet200Response) GetResolutionConfidence() BiomarkersGet200ResponseResolutionConfidence`

GetResolutionConfidence returns the ResolutionConfidence field if non-nil, zero value otherwise.

### GetResolutionConfidenceOk

`func (o *OrgsGet200Response) GetResolutionConfidenceOk() (*BiomarkersGet200ResponseResolutionConfidence, bool)`

GetResolutionConfidenceOk returns a tuple with the ResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionConfidence

`func (o *OrgsGet200Response) SetResolutionConfidence(v BiomarkersGet200ResponseResolutionConfidence)`

SetResolutionConfidence sets ResolutionConfidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


