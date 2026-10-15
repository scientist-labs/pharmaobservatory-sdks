# ApiV1EnrollmentBenchmarkGet200ResponseBenchmark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CohortLevel** | **string** |  | 
**CohortLabel** | **string** |  | 
**ConditionName** | **string** |  | 
**Phase** | **string** |  | 
**Country** | **string** |  | 
**CountryBacked** | **bool** |  | 
**AccrualRatePerSiteMonth** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkAccrualRatePerSiteMonth**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkAccrualRatePerSiteMonth.md) |  | 
**TypicalDurationMonths** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkTypicalDurationMonths**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkTypicalDurationMonths.md) |  | 
**Enrollment** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment.md) |  | 
**ActiveSites** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment.md) |  | 
**CohortSize** | **int32** |  | 
**CohortSizeTotalIndication** | **int32** |  | 
**Confidence** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkConfidence**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkConfidence.md) |  | 

## Methods

### NewApiV1EnrollmentBenchmarkGet200ResponseBenchmark

`func NewApiV1EnrollmentBenchmarkGet200ResponseBenchmark(cohortLevel string, cohortLabel string, conditionName string, phase string, country string, countryBacked bool, accrualRatePerSiteMonth ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkAccrualRatePerSiteMonth, typicalDurationMonths ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkTypicalDurationMonths, enrollment ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment, activeSites ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment, cohortSize int32, cohortSizeTotalIndication int32, confidence ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkConfidence, ) *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark`

NewApiV1EnrollmentBenchmarkGet200ResponseBenchmark instantiates a new ApiV1EnrollmentBenchmarkGet200ResponseBenchmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1EnrollmentBenchmarkGet200ResponseBenchmarkWithDefaults

`func NewApiV1EnrollmentBenchmarkGet200ResponseBenchmarkWithDefaults() *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark`

NewApiV1EnrollmentBenchmarkGet200ResponseBenchmarkWithDefaults instantiates a new ApiV1EnrollmentBenchmarkGet200ResponseBenchmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCohortLevel

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCohortLevel() string`

GetCohortLevel returns the CohortLevel field if non-nil, zero value otherwise.

### GetCohortLevelOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCohortLevelOk() (*string, bool)`

GetCohortLevelOk returns a tuple with the CohortLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortLevel

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetCohortLevel(v string)`

SetCohortLevel sets CohortLevel field to given value.


### GetCohortLabel

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCohortLabel() string`

GetCohortLabel returns the CohortLabel field if non-nil, zero value otherwise.

### GetCohortLabelOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCohortLabelOk() (*string, bool)`

GetCohortLabelOk returns a tuple with the CohortLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortLabel

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetCohortLabel(v string)`

SetCohortLabel sets CohortLabel field to given value.


### GetConditionName

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.


### GetPhase

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetCountry

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCountryBacked

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCountryBacked() bool`

GetCountryBacked returns the CountryBacked field if non-nil, zero value otherwise.

### GetCountryBackedOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCountryBackedOk() (*bool, bool)`

GetCountryBackedOk returns a tuple with the CountryBacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryBacked

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetCountryBacked(v bool)`

SetCountryBacked sets CountryBacked field to given value.


### GetAccrualRatePerSiteMonth

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetAccrualRatePerSiteMonth() ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkAccrualRatePerSiteMonth`

GetAccrualRatePerSiteMonth returns the AccrualRatePerSiteMonth field if non-nil, zero value otherwise.

### GetAccrualRatePerSiteMonthOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetAccrualRatePerSiteMonthOk() (*ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkAccrualRatePerSiteMonth, bool)`

GetAccrualRatePerSiteMonthOk returns a tuple with the AccrualRatePerSiteMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualRatePerSiteMonth

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetAccrualRatePerSiteMonth(v ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkAccrualRatePerSiteMonth)`

SetAccrualRatePerSiteMonth sets AccrualRatePerSiteMonth field to given value.


### GetTypicalDurationMonths

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetTypicalDurationMonths() ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkTypicalDurationMonths`

GetTypicalDurationMonths returns the TypicalDurationMonths field if non-nil, zero value otherwise.

### GetTypicalDurationMonthsOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetTypicalDurationMonthsOk() (*ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkTypicalDurationMonths, bool)`

GetTypicalDurationMonthsOk returns a tuple with the TypicalDurationMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypicalDurationMonths

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetTypicalDurationMonths(v ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkTypicalDurationMonths)`

SetTypicalDurationMonths sets TypicalDurationMonths field to given value.


### GetEnrollment

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetEnrollment() ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetEnrollmentOk() (*ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetEnrollment(v ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment)`

SetEnrollment sets Enrollment field to given value.


### GetActiveSites

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetActiveSites() ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment`

GetActiveSites returns the ActiveSites field if non-nil, zero value otherwise.

### GetActiveSitesOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetActiveSitesOk() (*ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment, bool)`

GetActiveSitesOk returns a tuple with the ActiveSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSites

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetActiveSites(v ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment)`

SetActiveSites sets ActiveSites field to given value.


### GetCohortSize

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCohortSize() int32`

GetCohortSize returns the CohortSize field if non-nil, zero value otherwise.

### GetCohortSizeOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCohortSizeOk() (*int32, bool)`

GetCohortSizeOk returns a tuple with the CohortSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortSize

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetCohortSize(v int32)`

SetCohortSize sets CohortSize field to given value.


### GetCohortSizeTotalIndication

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCohortSizeTotalIndication() int32`

GetCohortSizeTotalIndication returns the CohortSizeTotalIndication field if non-nil, zero value otherwise.

### GetCohortSizeTotalIndicationOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetCohortSizeTotalIndicationOk() (*int32, bool)`

GetCohortSizeTotalIndicationOk returns a tuple with the CohortSizeTotalIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortSizeTotalIndication

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetCohortSizeTotalIndication(v int32)`

SetCohortSizeTotalIndication sets CohortSizeTotalIndication field to given value.


### GetConfidence

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetConfidence() ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) GetConfidenceOk() (*ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ApiV1EnrollmentBenchmarkGet200ResponseBenchmark) SetConfidence(v ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkConfidence)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


