# TrialDemographicsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**HasSexData** | Pointer to **NullableBool** | Whether a sex baseline breakdown was posted. | [optional] 
**HasRaceData** | Pointer to **NullableBool** | Whether a race baseline breakdown was posted. | [optional] 
**HasEthnicityData** | Pointer to **NullableBool** | Whether an ethnicity baseline breakdown was posted. | [optional] 
**HasAgeData** | Pointer to **NullableBool** | Whether a categorical age baseline breakdown was posted. | [optional] 
**PctFemale** | Pointer to **NullableFloat32** | Percent of baseline participants recorded as female (null when sex was not posted). | [optional] 
**RaceBreakdown** | Pointer to **NullableString** | A JSON object of race category to percent of baseline participants, as posted. | [optional] 
**NAnalyzed** | Pointer to **NullableInt32** | Baseline participants analyzed. | [optional] 

## Methods

### NewTrialDemographicsRow

`func NewTrialDemographicsRow() *TrialDemographicsRow`

NewTrialDemographicsRow instantiates a new TrialDemographicsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialDemographicsRowWithDefaults

`func NewTrialDemographicsRowWithDefaults() *TrialDemographicsRow`

NewTrialDemographicsRowWithDefaults instantiates a new TrialDemographicsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialDemographicsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialDemographicsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialDemographicsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialDemographicsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialDemographicsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialDemographicsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetHasSexData

`func (o *TrialDemographicsRow) GetHasSexData() bool`

GetHasSexData returns the HasSexData field if non-nil, zero value otherwise.

### GetHasSexDataOk

`func (o *TrialDemographicsRow) GetHasSexDataOk() (*bool, bool)`

GetHasSexDataOk returns a tuple with the HasSexData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSexData

`func (o *TrialDemographicsRow) SetHasSexData(v bool)`

SetHasSexData sets HasSexData field to given value.

### HasHasSexData

`func (o *TrialDemographicsRow) HasHasSexData() bool`

HasHasSexData returns a boolean if a field has been set.

### SetHasSexDataNil

`func (o *TrialDemographicsRow) SetHasSexDataNil(b bool)`

 SetHasSexDataNil sets the value for HasSexData to be an explicit nil

### UnsetHasSexData
`func (o *TrialDemographicsRow) UnsetHasSexData()`

UnsetHasSexData ensures that no value is present for HasSexData, not even an explicit nil
### GetHasRaceData

`func (o *TrialDemographicsRow) GetHasRaceData() bool`

GetHasRaceData returns the HasRaceData field if non-nil, zero value otherwise.

### GetHasRaceDataOk

`func (o *TrialDemographicsRow) GetHasRaceDataOk() (*bool, bool)`

GetHasRaceDataOk returns a tuple with the HasRaceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRaceData

`func (o *TrialDemographicsRow) SetHasRaceData(v bool)`

SetHasRaceData sets HasRaceData field to given value.

### HasHasRaceData

`func (o *TrialDemographicsRow) HasHasRaceData() bool`

HasHasRaceData returns a boolean if a field has been set.

### SetHasRaceDataNil

`func (o *TrialDemographicsRow) SetHasRaceDataNil(b bool)`

 SetHasRaceDataNil sets the value for HasRaceData to be an explicit nil

### UnsetHasRaceData
`func (o *TrialDemographicsRow) UnsetHasRaceData()`

UnsetHasRaceData ensures that no value is present for HasRaceData, not even an explicit nil
### GetHasEthnicityData

`func (o *TrialDemographicsRow) GetHasEthnicityData() bool`

GetHasEthnicityData returns the HasEthnicityData field if non-nil, zero value otherwise.

### GetHasEthnicityDataOk

`func (o *TrialDemographicsRow) GetHasEthnicityDataOk() (*bool, bool)`

GetHasEthnicityDataOk returns a tuple with the HasEthnicityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEthnicityData

`func (o *TrialDemographicsRow) SetHasEthnicityData(v bool)`

SetHasEthnicityData sets HasEthnicityData field to given value.

### HasHasEthnicityData

`func (o *TrialDemographicsRow) HasHasEthnicityData() bool`

HasHasEthnicityData returns a boolean if a field has been set.

### SetHasEthnicityDataNil

`func (o *TrialDemographicsRow) SetHasEthnicityDataNil(b bool)`

 SetHasEthnicityDataNil sets the value for HasEthnicityData to be an explicit nil

### UnsetHasEthnicityData
`func (o *TrialDemographicsRow) UnsetHasEthnicityData()`

UnsetHasEthnicityData ensures that no value is present for HasEthnicityData, not even an explicit nil
### GetHasAgeData

`func (o *TrialDemographicsRow) GetHasAgeData() bool`

GetHasAgeData returns the HasAgeData field if non-nil, zero value otherwise.

### GetHasAgeDataOk

`func (o *TrialDemographicsRow) GetHasAgeDataOk() (*bool, bool)`

GetHasAgeDataOk returns a tuple with the HasAgeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAgeData

`func (o *TrialDemographicsRow) SetHasAgeData(v bool)`

SetHasAgeData sets HasAgeData field to given value.

### HasHasAgeData

`func (o *TrialDemographicsRow) HasHasAgeData() bool`

HasHasAgeData returns a boolean if a field has been set.

### SetHasAgeDataNil

`func (o *TrialDemographicsRow) SetHasAgeDataNil(b bool)`

 SetHasAgeDataNil sets the value for HasAgeData to be an explicit nil

### UnsetHasAgeData
`func (o *TrialDemographicsRow) UnsetHasAgeData()`

UnsetHasAgeData ensures that no value is present for HasAgeData, not even an explicit nil
### GetPctFemale

`func (o *TrialDemographicsRow) GetPctFemale() float32`

GetPctFemale returns the PctFemale field if non-nil, zero value otherwise.

### GetPctFemaleOk

`func (o *TrialDemographicsRow) GetPctFemaleOk() (*float32, bool)`

GetPctFemaleOk returns a tuple with the PctFemale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctFemale

`func (o *TrialDemographicsRow) SetPctFemale(v float32)`

SetPctFemale sets PctFemale field to given value.

### HasPctFemale

`func (o *TrialDemographicsRow) HasPctFemale() bool`

HasPctFemale returns a boolean if a field has been set.

### SetPctFemaleNil

`func (o *TrialDemographicsRow) SetPctFemaleNil(b bool)`

 SetPctFemaleNil sets the value for PctFemale to be an explicit nil

### UnsetPctFemale
`func (o *TrialDemographicsRow) UnsetPctFemale()`

UnsetPctFemale ensures that no value is present for PctFemale, not even an explicit nil
### GetRaceBreakdown

`func (o *TrialDemographicsRow) GetRaceBreakdown() string`

GetRaceBreakdown returns the RaceBreakdown field if non-nil, zero value otherwise.

### GetRaceBreakdownOk

`func (o *TrialDemographicsRow) GetRaceBreakdownOk() (*string, bool)`

GetRaceBreakdownOk returns a tuple with the RaceBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaceBreakdown

`func (o *TrialDemographicsRow) SetRaceBreakdown(v string)`

SetRaceBreakdown sets RaceBreakdown field to given value.

### HasRaceBreakdown

`func (o *TrialDemographicsRow) HasRaceBreakdown() bool`

HasRaceBreakdown returns a boolean if a field has been set.

### SetRaceBreakdownNil

`func (o *TrialDemographicsRow) SetRaceBreakdownNil(b bool)`

 SetRaceBreakdownNil sets the value for RaceBreakdown to be an explicit nil

### UnsetRaceBreakdown
`func (o *TrialDemographicsRow) UnsetRaceBreakdown()`

UnsetRaceBreakdown ensures that no value is present for RaceBreakdown, not even an explicit nil
### GetNAnalyzed

`func (o *TrialDemographicsRow) GetNAnalyzed() int32`

GetNAnalyzed returns the NAnalyzed field if non-nil, zero value otherwise.

### GetNAnalyzedOk

`func (o *TrialDemographicsRow) GetNAnalyzedOk() (*int32, bool)`

GetNAnalyzedOk returns a tuple with the NAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAnalyzed

`func (o *TrialDemographicsRow) SetNAnalyzed(v int32)`

SetNAnalyzed sets NAnalyzed field to given value.

### HasNAnalyzed

`func (o *TrialDemographicsRow) HasNAnalyzed() bool`

HasNAnalyzed returns a boolean if a field has been set.

### SetNAnalyzedNil

`func (o *TrialDemographicsRow) SetNAnalyzedNil(b bool)`

 SetNAnalyzedNil sets the value for NAnalyzed to be an explicit nil

### UnsetNAnalyzed
`func (o *TrialDemographicsRow) UnsetNAnalyzed()`

UnsetNAnalyzed ensures that no value is present for NAnalyzed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


