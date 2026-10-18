# BiomarkersCgi200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Biomarker** | **string** |  | 
**VariantCount** | **int32** |  | 
**Drug** | **string** |  | 
**CanonicalDrugId** | **NullableString** |  | 
**Response** | **string** |  | 
**Disease** | **string** |  | 
**Source** | **string** |  | 
**SourceLabel** | **string** |  | 
**Score** | **float32** |  | 

## Methods

### NewBiomarkersCgi200ResponseRowsInner

`func NewBiomarkersCgi200ResponseRowsInner(biomarker string, variantCount int32, drug string, canonicalDrugId NullableString, response string, disease string, source string, sourceLabel string, score float32, ) *BiomarkersCgi200ResponseRowsInner`

NewBiomarkersCgi200ResponseRowsInner instantiates a new BiomarkersCgi200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkersCgi200ResponseRowsInnerWithDefaults

`func NewBiomarkersCgi200ResponseRowsInnerWithDefaults() *BiomarkersCgi200ResponseRowsInner`

NewBiomarkersCgi200ResponseRowsInnerWithDefaults instantiates a new BiomarkersCgi200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiomarker

`func (o *BiomarkersCgi200ResponseRowsInner) GetBiomarker() string`

GetBiomarker returns the Biomarker field if non-nil, zero value otherwise.

### GetBiomarkerOk

`func (o *BiomarkersCgi200ResponseRowsInner) GetBiomarkerOk() (*string, bool)`

GetBiomarkerOk returns a tuple with the Biomarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarker

`func (o *BiomarkersCgi200ResponseRowsInner) SetBiomarker(v string)`

SetBiomarker sets Biomarker field to given value.


### GetVariantCount

`func (o *BiomarkersCgi200ResponseRowsInner) GetVariantCount() int32`

GetVariantCount returns the VariantCount field if non-nil, zero value otherwise.

### GetVariantCountOk

`func (o *BiomarkersCgi200ResponseRowsInner) GetVariantCountOk() (*int32, bool)`

GetVariantCountOk returns a tuple with the VariantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantCount

`func (o *BiomarkersCgi200ResponseRowsInner) SetVariantCount(v int32)`

SetVariantCount sets VariantCount field to given value.


### GetDrug

`func (o *BiomarkersCgi200ResponseRowsInner) GetDrug() string`

GetDrug returns the Drug field if non-nil, zero value otherwise.

### GetDrugOk

`func (o *BiomarkersCgi200ResponseRowsInner) GetDrugOk() (*string, bool)`

GetDrugOk returns a tuple with the Drug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrug

`func (o *BiomarkersCgi200ResponseRowsInner) SetDrug(v string)`

SetDrug sets Drug field to given value.


### GetCanonicalDrugId

`func (o *BiomarkersCgi200ResponseRowsInner) GetCanonicalDrugId() string`

GetCanonicalDrugId returns the CanonicalDrugId field if non-nil, zero value otherwise.

### GetCanonicalDrugIdOk

`func (o *BiomarkersCgi200ResponseRowsInner) GetCanonicalDrugIdOk() (*string, bool)`

GetCanonicalDrugIdOk returns a tuple with the CanonicalDrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalDrugId

`func (o *BiomarkersCgi200ResponseRowsInner) SetCanonicalDrugId(v string)`

SetCanonicalDrugId sets CanonicalDrugId field to given value.


### SetCanonicalDrugIdNil

`func (o *BiomarkersCgi200ResponseRowsInner) SetCanonicalDrugIdNil(b bool)`

 SetCanonicalDrugIdNil sets the value for CanonicalDrugId to be an explicit nil

### UnsetCanonicalDrugId
`func (o *BiomarkersCgi200ResponseRowsInner) UnsetCanonicalDrugId()`

UnsetCanonicalDrugId ensures that no value is present for CanonicalDrugId, not even an explicit nil
### GetResponse

`func (o *BiomarkersCgi200ResponseRowsInner) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *BiomarkersCgi200ResponseRowsInner) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *BiomarkersCgi200ResponseRowsInner) SetResponse(v string)`

SetResponse sets Response field to given value.


### GetDisease

`func (o *BiomarkersCgi200ResponseRowsInner) GetDisease() string`

GetDisease returns the Disease field if non-nil, zero value otherwise.

### GetDiseaseOk

`func (o *BiomarkersCgi200ResponseRowsInner) GetDiseaseOk() (*string, bool)`

GetDiseaseOk returns a tuple with the Disease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisease

`func (o *BiomarkersCgi200ResponseRowsInner) SetDisease(v string)`

SetDisease sets Disease field to given value.


### GetSource

`func (o *BiomarkersCgi200ResponseRowsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BiomarkersCgi200ResponseRowsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BiomarkersCgi200ResponseRowsInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceLabel

`func (o *BiomarkersCgi200ResponseRowsInner) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *BiomarkersCgi200ResponseRowsInner) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *BiomarkersCgi200ResponseRowsInner) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.


### GetScore

`func (o *BiomarkersCgi200ResponseRowsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BiomarkersCgi200ResponseRowsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BiomarkersCgi200ResponseRowsInner) SetScore(v float32)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


