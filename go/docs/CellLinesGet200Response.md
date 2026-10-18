# CellLinesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **string** |  | 
**Id** | **string** |  | 
**Identity** | [**CellLinesGet200ResponseIdentity**](CellLinesGet200ResponseIdentity.md) |  | 
**Sections** | [**[]BiomarkersGet200ResponseSectionsInner**](BiomarkersGet200ResponseSectionsInner.md) |  | 
**Attribution** | [**[]CellLinesList200ResponseAttributionInner**](CellLinesList200ResponseAttributionInner.md) |  | 
**WebUrl** | **string** |  | 
**BuildState** | [**BiomarkersList200ResponseBuildState**](BiomarkersList200ResponseBuildState.md) |  | 
**ResolutionConfidence** | [**BiomarkersGet200ResponseResolutionConfidence**](BiomarkersGet200ResponseResolutionConfidence.md) |  | 

## Methods

### NewCellLinesGet200Response

`func NewCellLinesGet200Response(page string, id string, identity CellLinesGet200ResponseIdentity, sections []BiomarkersGet200ResponseSectionsInner, attribution []CellLinesList200ResponseAttributionInner, webUrl string, buildState BiomarkersList200ResponseBuildState, resolutionConfidence BiomarkersGet200ResponseResolutionConfidence, ) *CellLinesGet200Response`

NewCellLinesGet200Response instantiates a new CellLinesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellLinesGet200ResponseWithDefaults

`func NewCellLinesGet200ResponseWithDefaults() *CellLinesGet200Response`

NewCellLinesGet200ResponseWithDefaults instantiates a new CellLinesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CellLinesGet200Response) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CellLinesGet200Response) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CellLinesGet200Response) SetPage(v string)`

SetPage sets Page field to given value.


### GetId

`func (o *CellLinesGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CellLinesGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CellLinesGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetIdentity

`func (o *CellLinesGet200Response) GetIdentity() CellLinesGet200ResponseIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CellLinesGet200Response) GetIdentityOk() (*CellLinesGet200ResponseIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CellLinesGet200Response) SetIdentity(v CellLinesGet200ResponseIdentity)`

SetIdentity sets Identity field to given value.


### GetSections

`func (o *CellLinesGet200Response) GetSections() []BiomarkersGet200ResponseSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *CellLinesGet200Response) GetSectionsOk() (*[]BiomarkersGet200ResponseSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *CellLinesGet200Response) SetSections(v []BiomarkersGet200ResponseSectionsInner)`

SetSections sets Sections field to given value.


### GetAttribution

`func (o *CellLinesGet200Response) GetAttribution() []CellLinesList200ResponseAttributionInner`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *CellLinesGet200Response) GetAttributionOk() (*[]CellLinesList200ResponseAttributionInner, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *CellLinesGet200Response) SetAttribution(v []CellLinesList200ResponseAttributionInner)`

SetAttribution sets Attribution field to given value.


### GetWebUrl

`func (o *CellLinesGet200Response) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *CellLinesGet200Response) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *CellLinesGet200Response) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetBuildState

`func (o *CellLinesGet200Response) GetBuildState() BiomarkersList200ResponseBuildState`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *CellLinesGet200Response) GetBuildStateOk() (*BiomarkersList200ResponseBuildState, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *CellLinesGet200Response) SetBuildState(v BiomarkersList200ResponseBuildState)`

SetBuildState sets BuildState field to given value.


### GetResolutionConfidence

`func (o *CellLinesGet200Response) GetResolutionConfidence() BiomarkersGet200ResponseResolutionConfidence`

GetResolutionConfidence returns the ResolutionConfidence field if non-nil, zero value otherwise.

### GetResolutionConfidenceOk

`func (o *CellLinesGet200Response) GetResolutionConfidenceOk() (*BiomarkersGet200ResponseResolutionConfidence, bool)`

GetResolutionConfidenceOk returns a tuple with the ResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionConfidence

`func (o *CellLinesGet200Response) SetResolutionConfidence(v BiomarkersGet200ResponseResolutionConfidence)`

SetResolutionConfidence sets ResolutionConfidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


