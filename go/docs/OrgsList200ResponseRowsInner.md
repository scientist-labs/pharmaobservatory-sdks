# OrgsList200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CanonicalOrgId** | **NullableString** |  | 
**Name** | **string** |  | 
**Kind** | **string** |  | 
**OrgType** | **string** |  | 
**OrgKinds** | **[]string** |  | 
**Country** | **NullableString** |  | 
**Ticker** | **NullableString** |  | 
**LeadTrials** | **int32** |  | 
**RecentLeadTrials** | **int32** |  | 
**LeadEnrollment** | **int32** |  | 
**CollaboratorTrials** | **interface{}** |  | 
**GroupSize** | **int32** |  | 
**WebUrl** | **NullableString** |  | 

## Methods

### NewOrgsList200ResponseRowsInner

`func NewOrgsList200ResponseRowsInner(id string, canonicalOrgId NullableString, name string, kind string, orgType string, orgKinds []string, country NullableString, ticker NullableString, leadTrials int32, recentLeadTrials int32, leadEnrollment int32, collaboratorTrials interface{}, groupSize int32, webUrl NullableString, ) *OrgsList200ResponseRowsInner`

NewOrgsList200ResponseRowsInner instantiates a new OrgsList200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsList200ResponseRowsInnerWithDefaults

`func NewOrgsList200ResponseRowsInnerWithDefaults() *OrgsList200ResponseRowsInner`

NewOrgsList200ResponseRowsInnerWithDefaults instantiates a new OrgsList200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrgsList200ResponseRowsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgsList200ResponseRowsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgsList200ResponseRowsInner) SetId(v string)`

SetId sets Id field to given value.


### GetCanonicalOrgId

`func (o *OrgsList200ResponseRowsInner) GetCanonicalOrgId() string`

GetCanonicalOrgId returns the CanonicalOrgId field if non-nil, zero value otherwise.

### GetCanonicalOrgIdOk

`func (o *OrgsList200ResponseRowsInner) GetCanonicalOrgIdOk() (*string, bool)`

GetCanonicalOrgIdOk returns a tuple with the CanonicalOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalOrgId

`func (o *OrgsList200ResponseRowsInner) SetCanonicalOrgId(v string)`

SetCanonicalOrgId sets CanonicalOrgId field to given value.


### SetCanonicalOrgIdNil

`func (o *OrgsList200ResponseRowsInner) SetCanonicalOrgIdNil(b bool)`

 SetCanonicalOrgIdNil sets the value for CanonicalOrgId to be an explicit nil

### UnsetCanonicalOrgId
`func (o *OrgsList200ResponseRowsInner) UnsetCanonicalOrgId()`

UnsetCanonicalOrgId ensures that no value is present for CanonicalOrgId, not even an explicit nil
### GetName

`func (o *OrgsList200ResponseRowsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgsList200ResponseRowsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgsList200ResponseRowsInner) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *OrgsList200ResponseRowsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgsList200ResponseRowsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgsList200ResponseRowsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetOrgType

`func (o *OrgsList200ResponseRowsInner) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *OrgsList200ResponseRowsInner) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *OrgsList200ResponseRowsInner) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.


### GetOrgKinds

`func (o *OrgsList200ResponseRowsInner) GetOrgKinds() []string`

GetOrgKinds returns the OrgKinds field if non-nil, zero value otherwise.

### GetOrgKindsOk

`func (o *OrgsList200ResponseRowsInner) GetOrgKindsOk() (*[]string, bool)`

GetOrgKindsOk returns a tuple with the OrgKinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgKinds

`func (o *OrgsList200ResponseRowsInner) SetOrgKinds(v []string)`

SetOrgKinds sets OrgKinds field to given value.


### GetCountry

`func (o *OrgsList200ResponseRowsInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrgsList200ResponseRowsInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrgsList200ResponseRowsInner) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *OrgsList200ResponseRowsInner) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *OrgsList200ResponseRowsInner) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetTicker

`func (o *OrgsList200ResponseRowsInner) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *OrgsList200ResponseRowsInner) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *OrgsList200ResponseRowsInner) SetTicker(v string)`

SetTicker sets Ticker field to given value.


### SetTickerNil

`func (o *OrgsList200ResponseRowsInner) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *OrgsList200ResponseRowsInner) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil
### GetLeadTrials

`func (o *OrgsList200ResponseRowsInner) GetLeadTrials() int32`

GetLeadTrials returns the LeadTrials field if non-nil, zero value otherwise.

### GetLeadTrialsOk

`func (o *OrgsList200ResponseRowsInner) GetLeadTrialsOk() (*int32, bool)`

GetLeadTrialsOk returns a tuple with the LeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadTrials

`func (o *OrgsList200ResponseRowsInner) SetLeadTrials(v int32)`

SetLeadTrials sets LeadTrials field to given value.


### GetRecentLeadTrials

`func (o *OrgsList200ResponseRowsInner) GetRecentLeadTrials() int32`

GetRecentLeadTrials returns the RecentLeadTrials field if non-nil, zero value otherwise.

### GetRecentLeadTrialsOk

`func (o *OrgsList200ResponseRowsInner) GetRecentLeadTrialsOk() (*int32, bool)`

GetRecentLeadTrialsOk returns a tuple with the RecentLeadTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentLeadTrials

`func (o *OrgsList200ResponseRowsInner) SetRecentLeadTrials(v int32)`

SetRecentLeadTrials sets RecentLeadTrials field to given value.


### GetLeadEnrollment

`func (o *OrgsList200ResponseRowsInner) GetLeadEnrollment() int32`

GetLeadEnrollment returns the LeadEnrollment field if non-nil, zero value otherwise.

### GetLeadEnrollmentOk

`func (o *OrgsList200ResponseRowsInner) GetLeadEnrollmentOk() (*int32, bool)`

GetLeadEnrollmentOk returns a tuple with the LeadEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadEnrollment

`func (o *OrgsList200ResponseRowsInner) SetLeadEnrollment(v int32)`

SetLeadEnrollment sets LeadEnrollment field to given value.


### GetCollaboratorTrials

`func (o *OrgsList200ResponseRowsInner) GetCollaboratorTrials() interface{}`

GetCollaboratorTrials returns the CollaboratorTrials field if non-nil, zero value otherwise.

### GetCollaboratorTrialsOk

`func (o *OrgsList200ResponseRowsInner) GetCollaboratorTrialsOk() (*interface{}, bool)`

GetCollaboratorTrialsOk returns a tuple with the CollaboratorTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorTrials

`func (o *OrgsList200ResponseRowsInner) SetCollaboratorTrials(v interface{})`

SetCollaboratorTrials sets CollaboratorTrials field to given value.


### SetCollaboratorTrialsNil

`func (o *OrgsList200ResponseRowsInner) SetCollaboratorTrialsNil(b bool)`

 SetCollaboratorTrialsNil sets the value for CollaboratorTrials to be an explicit nil

### UnsetCollaboratorTrials
`func (o *OrgsList200ResponseRowsInner) UnsetCollaboratorTrials()`

UnsetCollaboratorTrials ensures that no value is present for CollaboratorTrials, not even an explicit nil
### GetGroupSize

`func (o *OrgsList200ResponseRowsInner) GetGroupSize() int32`

GetGroupSize returns the GroupSize field if non-nil, zero value otherwise.

### GetGroupSizeOk

`func (o *OrgsList200ResponseRowsInner) GetGroupSizeOk() (*int32, bool)`

GetGroupSizeOk returns a tuple with the GroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSize

`func (o *OrgsList200ResponseRowsInner) SetGroupSize(v int32)`

SetGroupSize sets GroupSize field to given value.


### GetWebUrl

`func (o *OrgsList200ResponseRowsInner) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OrgsList200ResponseRowsInner) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OrgsList200ResponseRowsInner) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### SetWebUrlNil

`func (o *OrgsList200ResponseRowsInner) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *OrgsList200ResponseRowsInner) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


