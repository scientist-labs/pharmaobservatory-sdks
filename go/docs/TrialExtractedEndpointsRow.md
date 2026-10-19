# TrialExtractedEndpointsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**ResultReported** | Pointer to **NullableBool** | Whether the extraction found a formal primary-endpoint result. | [optional] 
**Signal** | Pointer to **NullableString** | The extracted reading, verbatim: met, missed, mixed, reported_indeterminate or not_reported. | [optional] 
**OutcomeCategory** | Pointer to **NullableString** | The signal grouped as met, missed, mixed or unclear. | [optional] 
**HasCleanOutcome** | Pointer to **NullableBool** | Whether outcome_category is met, missed or mixed. | [optional] 
**PrimaryEndpoints** | Pointer to **[]string** | The primary endpoint descriptions the extraction found. | [optional] 
**PrimaryEndpointCount** | Pointer to **NullableInt32** | Number of primary endpoints extracted. | [optional] 
**Confidence** | Pointer to **NullableFloat32** | The extraction&#39;s confidence, from 0 to 1. | [optional] 
**EvidenceQuote** | Pointer to **NullableString** | The text quoted as evidence for the reading. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The registry page the reading was taken from. | [optional] 
**Model** | Pointer to **NullableString** | The model that produced the reading. | [optional] 
**ExtractedAt** | Pointer to **NullableTime** | When the extraction ran. | [optional] 

## Methods

### NewTrialExtractedEndpointsRow

`func NewTrialExtractedEndpointsRow() *TrialExtractedEndpointsRow`

NewTrialExtractedEndpointsRow instantiates a new TrialExtractedEndpointsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialExtractedEndpointsRowWithDefaults

`func NewTrialExtractedEndpointsRowWithDefaults() *TrialExtractedEndpointsRow`

NewTrialExtractedEndpointsRowWithDefaults instantiates a new TrialExtractedEndpointsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialExtractedEndpointsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialExtractedEndpointsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialExtractedEndpointsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialExtractedEndpointsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialExtractedEndpointsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialExtractedEndpointsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetResultReported

`func (o *TrialExtractedEndpointsRow) GetResultReported() bool`

GetResultReported returns the ResultReported field if non-nil, zero value otherwise.

### GetResultReportedOk

`func (o *TrialExtractedEndpointsRow) GetResultReportedOk() (*bool, bool)`

GetResultReportedOk returns a tuple with the ResultReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReported

`func (o *TrialExtractedEndpointsRow) SetResultReported(v bool)`

SetResultReported sets ResultReported field to given value.

### HasResultReported

`func (o *TrialExtractedEndpointsRow) HasResultReported() bool`

HasResultReported returns a boolean if a field has been set.

### SetResultReportedNil

`func (o *TrialExtractedEndpointsRow) SetResultReportedNil(b bool)`

 SetResultReportedNil sets the value for ResultReported to be an explicit nil

### UnsetResultReported
`func (o *TrialExtractedEndpointsRow) UnsetResultReported()`

UnsetResultReported ensures that no value is present for ResultReported, not even an explicit nil
### GetSignal

`func (o *TrialExtractedEndpointsRow) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *TrialExtractedEndpointsRow) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *TrialExtractedEndpointsRow) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *TrialExtractedEndpointsRow) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### SetSignalNil

`func (o *TrialExtractedEndpointsRow) SetSignalNil(b bool)`

 SetSignalNil sets the value for Signal to be an explicit nil

### UnsetSignal
`func (o *TrialExtractedEndpointsRow) UnsetSignal()`

UnsetSignal ensures that no value is present for Signal, not even an explicit nil
### GetOutcomeCategory

`func (o *TrialExtractedEndpointsRow) GetOutcomeCategory() string`

GetOutcomeCategory returns the OutcomeCategory field if non-nil, zero value otherwise.

### GetOutcomeCategoryOk

`func (o *TrialExtractedEndpointsRow) GetOutcomeCategoryOk() (*string, bool)`

GetOutcomeCategoryOk returns a tuple with the OutcomeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeCategory

`func (o *TrialExtractedEndpointsRow) SetOutcomeCategory(v string)`

SetOutcomeCategory sets OutcomeCategory field to given value.

### HasOutcomeCategory

`func (o *TrialExtractedEndpointsRow) HasOutcomeCategory() bool`

HasOutcomeCategory returns a boolean if a field has been set.

### SetOutcomeCategoryNil

`func (o *TrialExtractedEndpointsRow) SetOutcomeCategoryNil(b bool)`

 SetOutcomeCategoryNil sets the value for OutcomeCategory to be an explicit nil

### UnsetOutcomeCategory
`func (o *TrialExtractedEndpointsRow) UnsetOutcomeCategory()`

UnsetOutcomeCategory ensures that no value is present for OutcomeCategory, not even an explicit nil
### GetHasCleanOutcome

`func (o *TrialExtractedEndpointsRow) GetHasCleanOutcome() bool`

GetHasCleanOutcome returns the HasCleanOutcome field if non-nil, zero value otherwise.

### GetHasCleanOutcomeOk

`func (o *TrialExtractedEndpointsRow) GetHasCleanOutcomeOk() (*bool, bool)`

GetHasCleanOutcomeOk returns a tuple with the HasCleanOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCleanOutcome

`func (o *TrialExtractedEndpointsRow) SetHasCleanOutcome(v bool)`

SetHasCleanOutcome sets HasCleanOutcome field to given value.

### HasHasCleanOutcome

`func (o *TrialExtractedEndpointsRow) HasHasCleanOutcome() bool`

HasHasCleanOutcome returns a boolean if a field has been set.

### SetHasCleanOutcomeNil

`func (o *TrialExtractedEndpointsRow) SetHasCleanOutcomeNil(b bool)`

 SetHasCleanOutcomeNil sets the value for HasCleanOutcome to be an explicit nil

### UnsetHasCleanOutcome
`func (o *TrialExtractedEndpointsRow) UnsetHasCleanOutcome()`

UnsetHasCleanOutcome ensures that no value is present for HasCleanOutcome, not even an explicit nil
### GetPrimaryEndpoints

`func (o *TrialExtractedEndpointsRow) GetPrimaryEndpoints() []string`

GetPrimaryEndpoints returns the PrimaryEndpoints field if non-nil, zero value otherwise.

### GetPrimaryEndpointsOk

`func (o *TrialExtractedEndpointsRow) GetPrimaryEndpointsOk() (*[]string, bool)`

GetPrimaryEndpointsOk returns a tuple with the PrimaryEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEndpoints

`func (o *TrialExtractedEndpointsRow) SetPrimaryEndpoints(v []string)`

SetPrimaryEndpoints sets PrimaryEndpoints field to given value.

### HasPrimaryEndpoints

`func (o *TrialExtractedEndpointsRow) HasPrimaryEndpoints() bool`

HasPrimaryEndpoints returns a boolean if a field has been set.

### SetPrimaryEndpointsNil

`func (o *TrialExtractedEndpointsRow) SetPrimaryEndpointsNil(b bool)`

 SetPrimaryEndpointsNil sets the value for PrimaryEndpoints to be an explicit nil

### UnsetPrimaryEndpoints
`func (o *TrialExtractedEndpointsRow) UnsetPrimaryEndpoints()`

UnsetPrimaryEndpoints ensures that no value is present for PrimaryEndpoints, not even an explicit nil
### GetPrimaryEndpointCount

`func (o *TrialExtractedEndpointsRow) GetPrimaryEndpointCount() int32`

GetPrimaryEndpointCount returns the PrimaryEndpointCount field if non-nil, zero value otherwise.

### GetPrimaryEndpointCountOk

`func (o *TrialExtractedEndpointsRow) GetPrimaryEndpointCountOk() (*int32, bool)`

GetPrimaryEndpointCountOk returns a tuple with the PrimaryEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEndpointCount

`func (o *TrialExtractedEndpointsRow) SetPrimaryEndpointCount(v int32)`

SetPrimaryEndpointCount sets PrimaryEndpointCount field to given value.

### HasPrimaryEndpointCount

`func (o *TrialExtractedEndpointsRow) HasPrimaryEndpointCount() bool`

HasPrimaryEndpointCount returns a boolean if a field has been set.

### SetPrimaryEndpointCountNil

`func (o *TrialExtractedEndpointsRow) SetPrimaryEndpointCountNil(b bool)`

 SetPrimaryEndpointCountNil sets the value for PrimaryEndpointCount to be an explicit nil

### UnsetPrimaryEndpointCount
`func (o *TrialExtractedEndpointsRow) UnsetPrimaryEndpointCount()`

UnsetPrimaryEndpointCount ensures that no value is present for PrimaryEndpointCount, not even an explicit nil
### GetConfidence

`func (o *TrialExtractedEndpointsRow) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *TrialExtractedEndpointsRow) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *TrialExtractedEndpointsRow) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *TrialExtractedEndpointsRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *TrialExtractedEndpointsRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *TrialExtractedEndpointsRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetEvidenceQuote

`func (o *TrialExtractedEndpointsRow) GetEvidenceQuote() string`

GetEvidenceQuote returns the EvidenceQuote field if non-nil, zero value otherwise.

### GetEvidenceQuoteOk

`func (o *TrialExtractedEndpointsRow) GetEvidenceQuoteOk() (*string, bool)`

GetEvidenceQuoteOk returns a tuple with the EvidenceQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceQuote

`func (o *TrialExtractedEndpointsRow) SetEvidenceQuote(v string)`

SetEvidenceQuote sets EvidenceQuote field to given value.

### HasEvidenceQuote

`func (o *TrialExtractedEndpointsRow) HasEvidenceQuote() bool`

HasEvidenceQuote returns a boolean if a field has been set.

### SetEvidenceQuoteNil

`func (o *TrialExtractedEndpointsRow) SetEvidenceQuoteNil(b bool)`

 SetEvidenceQuoteNil sets the value for EvidenceQuote to be an explicit nil

### UnsetEvidenceQuote
`func (o *TrialExtractedEndpointsRow) UnsetEvidenceQuote()`

UnsetEvidenceQuote ensures that no value is present for EvidenceQuote, not even an explicit nil
### GetSourceUrl

`func (o *TrialExtractedEndpointsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *TrialExtractedEndpointsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *TrialExtractedEndpointsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *TrialExtractedEndpointsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *TrialExtractedEndpointsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *TrialExtractedEndpointsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetModel

`func (o *TrialExtractedEndpointsRow) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TrialExtractedEndpointsRow) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TrialExtractedEndpointsRow) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TrialExtractedEndpointsRow) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *TrialExtractedEndpointsRow) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *TrialExtractedEndpointsRow) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetExtractedAt

`func (o *TrialExtractedEndpointsRow) GetExtractedAt() time.Time`

GetExtractedAt returns the ExtractedAt field if non-nil, zero value otherwise.

### GetExtractedAtOk

`func (o *TrialExtractedEndpointsRow) GetExtractedAtOk() (*time.Time, bool)`

GetExtractedAtOk returns a tuple with the ExtractedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedAt

`func (o *TrialExtractedEndpointsRow) SetExtractedAt(v time.Time)`

SetExtractedAt sets ExtractedAt field to given value.

### HasExtractedAt

`func (o *TrialExtractedEndpointsRow) HasExtractedAt() bool`

HasExtractedAt returns a boolean if a field has been set.

### SetExtractedAtNil

`func (o *TrialExtractedEndpointsRow) SetExtractedAtNil(b bool)`

 SetExtractedAtNil sets the value for ExtractedAt to be an explicit nil

### UnsetExtractedAt
`func (o *TrialExtractedEndpointsRow) UnsetExtractedAt()`

UnsetExtractedAt ensures that no value is present for ExtractedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


