# DatasetListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Grain** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **NullableString** |  | [optional] 
**Entitled** | Pointer to **bool** |  | [optional] 
**RowCount** | Pointer to **NullableInt32** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewDatasetListing

`func NewDatasetListing() *DatasetListing`

NewDatasetListing instantiates a new DatasetListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetListingWithDefaults

`func NewDatasetListingWithDefaults() *DatasetListing`

NewDatasetListingWithDefaults instantiates a new DatasetListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatasetListing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetListing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetListing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatasetListing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *DatasetListing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DatasetListing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DatasetListing) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DatasetListing) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *DatasetListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasetListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasetListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatasetListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGrain

`func (o *DatasetListing) GetGrain() string`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *DatasetListing) GetGrainOk() (*string, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *DatasetListing) SetGrain(v string)`

SetGrain sets Grain field to given value.

### HasGrain

`func (o *DatasetListing) HasGrain() bool`

HasGrain returns a boolean if a field has been set.

### GetTier

`func (o *DatasetListing) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *DatasetListing) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *DatasetListing) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *DatasetListing) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *DatasetListing) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *DatasetListing) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetEntitled

`func (o *DatasetListing) GetEntitled() bool`

GetEntitled returns the Entitled field if non-nil, zero value otherwise.

### GetEntitledOk

`func (o *DatasetListing) GetEntitledOk() (*bool, bool)`

GetEntitledOk returns a tuple with the Entitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitled

`func (o *DatasetListing) SetEntitled(v bool)`

SetEntitled sets Entitled field to given value.

### HasEntitled

`func (o *DatasetListing) HasEntitled() bool`

HasEntitled returns a boolean if a field has been set.

### GetRowCount

`func (o *DatasetListing) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DatasetListing) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DatasetListing) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DatasetListing) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCountNil

`func (o *DatasetListing) SetRowCountNil(b bool)`

 SetRowCountNil sets the value for RowCount to be an explicit nil

### UnsetRowCount
`func (o *DatasetListing) UnsetRowCount()`

UnsetRowCount ensures that no value is present for RowCount, not even an explicit nil
### GetHref

`func (o *DatasetListing) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DatasetListing) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DatasetListing) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DatasetListing) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


