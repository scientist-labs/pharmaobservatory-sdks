# BiomarkersSampleScreening200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AminoAcidChange** | **string** |  | 
**GdcProjectId** | **string** |  | 
**SourceDiseaseTypes** | **[]string** |  | 
**SourcePrimarySites** | **[]string** |  | 
**AffectedCaseCount** | **int32** |  | 
**SsmProfiledCaseCount** | **int32** |  | 
**ObservedFrequency** | **float32** |  | 
**FrequencyCi95Low** | **float32** |  | 
**FrequencyCi95High** | **float32** |  | 
**GdcDataRelease** | **string** |  | 
**ReferenceScreens** | Pointer to [**BiomarkersSampleScreening200ResponseRowsInnerReferenceScreens**](BiomarkersSampleScreening200ResponseRowsInnerReferenceScreens.md) |  | [optional] 

## Methods

### NewBiomarkersSampleScreening200ResponseRowsInner

`func NewBiomarkersSampleScreening200ResponseRowsInner(aminoAcidChange string, gdcProjectId string, sourceDiseaseTypes []string, sourcePrimarySites []string, affectedCaseCount int32, ssmProfiledCaseCount int32, observedFrequency float32, frequencyCi95Low float32, frequencyCi95High float32, gdcDataRelease string, ) *BiomarkersSampleScreening200ResponseRowsInner`

NewBiomarkersSampleScreening200ResponseRowsInner instantiates a new BiomarkersSampleScreening200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkersSampleScreening200ResponseRowsInnerWithDefaults

`func NewBiomarkersSampleScreening200ResponseRowsInnerWithDefaults() *BiomarkersSampleScreening200ResponseRowsInner`

NewBiomarkersSampleScreening200ResponseRowsInnerWithDefaults instantiates a new BiomarkersSampleScreening200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAminoAcidChange

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetAminoAcidChange() string`

GetAminoAcidChange returns the AminoAcidChange field if non-nil, zero value otherwise.

### GetAminoAcidChangeOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetAminoAcidChangeOk() (*string, bool)`

GetAminoAcidChangeOk returns a tuple with the AminoAcidChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidChange

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetAminoAcidChange(v string)`

SetAminoAcidChange sets AminoAcidChange field to given value.


### GetGdcProjectId

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetGdcProjectId() string`

GetGdcProjectId returns the GdcProjectId field if non-nil, zero value otherwise.

### GetGdcProjectIdOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetGdcProjectIdOk() (*string, bool)`

GetGdcProjectIdOk returns a tuple with the GdcProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdcProjectId

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetGdcProjectId(v string)`

SetGdcProjectId sets GdcProjectId field to given value.


### GetSourceDiseaseTypes

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetSourceDiseaseTypes() []string`

GetSourceDiseaseTypes returns the SourceDiseaseTypes field if non-nil, zero value otherwise.

### GetSourceDiseaseTypesOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetSourceDiseaseTypesOk() (*[]string, bool)`

GetSourceDiseaseTypesOk returns a tuple with the SourceDiseaseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDiseaseTypes

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetSourceDiseaseTypes(v []string)`

SetSourceDiseaseTypes sets SourceDiseaseTypes field to given value.


### GetSourcePrimarySites

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetSourcePrimarySites() []string`

GetSourcePrimarySites returns the SourcePrimarySites field if non-nil, zero value otherwise.

### GetSourcePrimarySitesOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetSourcePrimarySitesOk() (*[]string, bool)`

GetSourcePrimarySitesOk returns a tuple with the SourcePrimarySites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePrimarySites

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetSourcePrimarySites(v []string)`

SetSourcePrimarySites sets SourcePrimarySites field to given value.


### GetAffectedCaseCount

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetAffectedCaseCount() int32`

GetAffectedCaseCount returns the AffectedCaseCount field if non-nil, zero value otherwise.

### GetAffectedCaseCountOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetAffectedCaseCountOk() (*int32, bool)`

GetAffectedCaseCountOk returns a tuple with the AffectedCaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedCaseCount

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetAffectedCaseCount(v int32)`

SetAffectedCaseCount sets AffectedCaseCount field to given value.


### GetSsmProfiledCaseCount

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetSsmProfiledCaseCount() int32`

GetSsmProfiledCaseCount returns the SsmProfiledCaseCount field if non-nil, zero value otherwise.

### GetSsmProfiledCaseCountOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetSsmProfiledCaseCountOk() (*int32, bool)`

GetSsmProfiledCaseCountOk returns a tuple with the SsmProfiledCaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsmProfiledCaseCount

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetSsmProfiledCaseCount(v int32)`

SetSsmProfiledCaseCount sets SsmProfiledCaseCount field to given value.


### GetObservedFrequency

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetObservedFrequency() float32`

GetObservedFrequency returns the ObservedFrequency field if non-nil, zero value otherwise.

### GetObservedFrequencyOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetObservedFrequencyOk() (*float32, bool)`

GetObservedFrequencyOk returns a tuple with the ObservedFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedFrequency

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetObservedFrequency(v float32)`

SetObservedFrequency sets ObservedFrequency field to given value.


### GetFrequencyCi95Low

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetFrequencyCi95Low() float32`

GetFrequencyCi95Low returns the FrequencyCi95Low field if non-nil, zero value otherwise.

### GetFrequencyCi95LowOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetFrequencyCi95LowOk() (*float32, bool)`

GetFrequencyCi95LowOk returns a tuple with the FrequencyCi95Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyCi95Low

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetFrequencyCi95Low(v float32)`

SetFrequencyCi95Low sets FrequencyCi95Low field to given value.


### GetFrequencyCi95High

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetFrequencyCi95High() float32`

GetFrequencyCi95High returns the FrequencyCi95High field if non-nil, zero value otherwise.

### GetFrequencyCi95HighOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetFrequencyCi95HighOk() (*float32, bool)`

GetFrequencyCi95HighOk returns a tuple with the FrequencyCi95High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyCi95High

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetFrequencyCi95High(v float32)`

SetFrequencyCi95High sets FrequencyCi95High field to given value.


### GetGdcDataRelease

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetGdcDataRelease() string`

GetGdcDataRelease returns the GdcDataRelease field if non-nil, zero value otherwise.

### GetGdcDataReleaseOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetGdcDataReleaseOk() (*string, bool)`

GetGdcDataReleaseOk returns a tuple with the GdcDataRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdcDataRelease

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetGdcDataRelease(v string)`

SetGdcDataRelease sets GdcDataRelease field to given value.


### GetReferenceScreens

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetReferenceScreens() BiomarkersSampleScreening200ResponseRowsInnerReferenceScreens`

GetReferenceScreens returns the ReferenceScreens field if non-nil, zero value otherwise.

### GetReferenceScreensOk

`func (o *BiomarkersSampleScreening200ResponseRowsInner) GetReferenceScreensOk() (*BiomarkersSampleScreening200ResponseRowsInnerReferenceScreens, bool)`

GetReferenceScreensOk returns a tuple with the ReferenceScreens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceScreens

`func (o *BiomarkersSampleScreening200ResponseRowsInner) SetReferenceScreens(v BiomarkersSampleScreening200ResponseRowsInnerReferenceScreens)`

SetReferenceScreens sets ReferenceScreens field to given value.

### HasReferenceScreens

`func (o *BiomarkersSampleScreening200ResponseRowsInner) HasReferenceScreens() bool`

HasReferenceScreens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


