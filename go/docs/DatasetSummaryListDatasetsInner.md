# DatasetSummaryListDatasetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Grain** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **NullableString** |  | [optional] 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**ColumnCount** | Pointer to **int32** |  | [optional] 
**SourceCredit** | Pointer to **string** |  | [optional] 

## Methods

### NewDatasetSummaryListDatasetsInner

`func NewDatasetSummaryListDatasetsInner() *DatasetSummaryListDatasetsInner`

NewDatasetSummaryListDatasetsInner instantiates a new DatasetSummaryListDatasetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetSummaryListDatasetsInnerWithDefaults

`func NewDatasetSummaryListDatasetsInnerWithDefaults() *DatasetSummaryListDatasetsInner`

NewDatasetSummaryListDatasetsInnerWithDefaults instantiates a new DatasetSummaryListDatasetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatasetSummaryListDatasetsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetSummaryListDatasetsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetSummaryListDatasetsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatasetSummaryListDatasetsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *DatasetSummaryListDatasetsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DatasetSummaryListDatasetsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DatasetSummaryListDatasetsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DatasetSummaryListDatasetsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *DatasetSummaryListDatasetsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasetSummaryListDatasetsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasetSummaryListDatasetsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatasetSummaryListDatasetsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGrain

`func (o *DatasetSummaryListDatasetsInner) GetGrain() string`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *DatasetSummaryListDatasetsInner) GetGrainOk() (*string, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *DatasetSummaryListDatasetsInner) SetGrain(v string)`

SetGrain sets Grain field to given value.

### HasGrain

`func (o *DatasetSummaryListDatasetsInner) HasGrain() bool`

HasGrain returns a boolean if a field has been set.

### GetTier

`func (o *DatasetSummaryListDatasetsInner) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *DatasetSummaryListDatasetsInner) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *DatasetSummaryListDatasetsInner) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *DatasetSummaryListDatasetsInner) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *DatasetSummaryListDatasetsInner) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *DatasetSummaryListDatasetsInner) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetObjectTypes

`func (o *DatasetSummaryListDatasetsInner) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *DatasetSummaryListDatasetsInner) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *DatasetSummaryListDatasetsInner) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *DatasetSummaryListDatasetsInner) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetColumnCount

`func (o *DatasetSummaryListDatasetsInner) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *DatasetSummaryListDatasetsInner) GetColumnCountOk() (*int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnCount

`func (o *DatasetSummaryListDatasetsInner) SetColumnCount(v int32)`

SetColumnCount sets ColumnCount field to given value.

### HasColumnCount

`func (o *DatasetSummaryListDatasetsInner) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### GetSourceCredit

`func (o *DatasetSummaryListDatasetsInner) GetSourceCredit() string`

GetSourceCredit returns the SourceCredit field if non-nil, zero value otherwise.

### GetSourceCreditOk

`func (o *DatasetSummaryListDatasetsInner) GetSourceCreditOk() (*string, bool)`

GetSourceCreditOk returns a tuple with the SourceCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCredit

`func (o *DatasetSummaryListDatasetsInner) SetSourceCredit(v string)`

SetSourceCredit sets SourceCredit field to given value.

### HasSourceCredit

`func (o *DatasetSummaryListDatasetsInner) HasSourceCredit() bool`

HasSourceCredit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


