# ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**RegistryValue** | **bool** |  | 
**Marker** | **string** |  | 
**EstimateForm** | **string** |  | 
**State** | **string** |  | 
**Method** | **string** |  | 
**ModelName** | **string** |  | 
**ModelVersion** | **string** |  | 
**Cutoff** | **string** |  | 
**Reason** | [**NullableApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerReason**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerReason.md) |  | 
**Statement** | **NullableString** |  | 
**Figure** | [**NullableApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerFigure**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerFigure.md) |  | 
**Baseline** | [**NullableApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerBaseline**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerBaseline.md) |  | 
**RegisteredPrimaryCompletion** | [**NullableApiV1TrialsHistoryPredictionsInnerRegisteredPrimaryCompletion**](ApiV1TrialsHistoryPredictionsInnerRegisteredPrimaryCompletion.md) |  | 
**MeasuredError** | [**ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredError**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredError.md) |  | 
**InputsAsRegistered** | [**NullableApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerInputsAsRegistered**](ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerInputsAsRegistered.md) |  | 
**CutoffStatement** | **NullableString** |  | 
**HistoryRunsPastCutoff** | **NullableBool** |  | 
**NewestHistoryAsOf** | **NullableString** |  | 

## Methods

### NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner

`func NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner(kind string, registryValue bool, marker string, estimateForm string, state string, method string, modelName string, modelVersion string, cutoff string, reason NullableApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerReason, statement NullableString, figure NullableApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerFigure, baseline NullableApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerBaseline, registeredPrimaryCompletion NullableApiV1TrialsHistoryPredictionsInnerRegisteredPrimaryCompletion, measuredError ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredError, inputsAsRegistered NullableApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerInputsAsRegistered, cutoffStatement NullableString, historyRunsPastCutoff NullableBool, newestHistoryAsOf NullableString, ) *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner`

NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner instantiates a new ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerWithDefaults

`func NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerWithDefaults() *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner`

NewApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerWithDefaults instantiates a new ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRegistryValue

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetRegistryValue() bool`

GetRegistryValue returns the RegistryValue field if non-nil, zero value otherwise.

### GetRegistryValueOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetRegistryValueOk() (*bool, bool)`

GetRegistryValueOk returns a tuple with the RegistryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryValue

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetRegistryValue(v bool)`

SetRegistryValue sets RegistryValue field to given value.


### GetMarker

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetMarker(v string)`

SetMarker sets Marker field to given value.


### GetEstimateForm

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetEstimateForm() string`

GetEstimateForm returns the EstimateForm field if non-nil, zero value otherwise.

### GetEstimateFormOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetEstimateFormOk() (*string, bool)`

GetEstimateFormOk returns a tuple with the EstimateForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateForm

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetEstimateForm(v string)`

SetEstimateForm sets EstimateForm field to given value.


### GetState

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetState(v string)`

SetState sets State field to given value.


### GetMethod

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetModelName

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelVersion

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.


### GetCutoff

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetCutoff() string`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetCutoffOk() (*string, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoff

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetCutoff(v string)`

SetCutoff sets Cutoff field to given value.


### GetReason

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetReason() ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetReasonOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetReason(v ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerReason)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetStatement

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetStatement(v string)`

SetStatement sets Statement field to given value.


### SetStatementNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetStatementNil(b bool)`

 SetStatementNil sets the value for Statement to be an explicit nil

### UnsetStatement
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) UnsetStatement()`

UnsetStatement ensures that no value is present for Statement, not even an explicit nil
### GetFigure

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetFigure() ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerFigure`

GetFigure returns the Figure field if non-nil, zero value otherwise.

### GetFigureOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetFigureOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerFigure, bool)`

GetFigureOk returns a tuple with the Figure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigure

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetFigure(v ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerFigure)`

SetFigure sets Figure field to given value.


### SetFigureNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetFigureNil(b bool)`

 SetFigureNil sets the value for Figure to be an explicit nil

### UnsetFigure
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) UnsetFigure()`

UnsetFigure ensures that no value is present for Figure, not even an explicit nil
### GetBaseline

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetBaseline() ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerBaseline`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetBaselineOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerBaseline, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetBaseline(v ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerBaseline)`

SetBaseline sets Baseline field to given value.


### SetBaselineNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetBaselineNil(b bool)`

 SetBaselineNil sets the value for Baseline to be an explicit nil

### UnsetBaseline
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) UnsetBaseline()`

UnsetBaseline ensures that no value is present for Baseline, not even an explicit nil
### GetRegisteredPrimaryCompletion

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetRegisteredPrimaryCompletion() ApiV1TrialsHistoryPredictionsInnerRegisteredPrimaryCompletion`

GetRegisteredPrimaryCompletion returns the RegisteredPrimaryCompletion field if non-nil, zero value otherwise.

### GetRegisteredPrimaryCompletionOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetRegisteredPrimaryCompletionOk() (*ApiV1TrialsHistoryPredictionsInnerRegisteredPrimaryCompletion, bool)`

GetRegisteredPrimaryCompletionOk returns a tuple with the RegisteredPrimaryCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredPrimaryCompletion

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetRegisteredPrimaryCompletion(v ApiV1TrialsHistoryPredictionsInnerRegisteredPrimaryCompletion)`

SetRegisteredPrimaryCompletion sets RegisteredPrimaryCompletion field to given value.


### SetRegisteredPrimaryCompletionNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetRegisteredPrimaryCompletionNil(b bool)`

 SetRegisteredPrimaryCompletionNil sets the value for RegisteredPrimaryCompletion to be an explicit nil

### UnsetRegisteredPrimaryCompletion
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) UnsetRegisteredPrimaryCompletion()`

UnsetRegisteredPrimaryCompletion ensures that no value is present for RegisteredPrimaryCompletion, not even an explicit nil
### GetMeasuredError

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetMeasuredError() ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredError`

GetMeasuredError returns the MeasuredError field if non-nil, zero value otherwise.

### GetMeasuredErrorOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetMeasuredErrorOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredError, bool)`

GetMeasuredErrorOk returns a tuple with the MeasuredError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredError

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetMeasuredError(v ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerMeasuredError)`

SetMeasuredError sets MeasuredError field to given value.


### GetInputsAsRegistered

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetInputsAsRegistered() ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerInputsAsRegistered`

GetInputsAsRegistered returns the InputsAsRegistered field if non-nil, zero value otherwise.

### GetInputsAsRegisteredOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetInputsAsRegisteredOk() (*ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerInputsAsRegistered, bool)`

GetInputsAsRegisteredOk returns a tuple with the InputsAsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsAsRegistered

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetInputsAsRegistered(v ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInnerInputsAsRegistered)`

SetInputsAsRegistered sets InputsAsRegistered field to given value.


### SetInputsAsRegisteredNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetInputsAsRegisteredNil(b bool)`

 SetInputsAsRegisteredNil sets the value for InputsAsRegistered to be an explicit nil

### UnsetInputsAsRegistered
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) UnsetInputsAsRegistered()`

UnsetInputsAsRegistered ensures that no value is present for InputsAsRegistered, not even an explicit nil
### GetCutoffStatement

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetCutoffStatement() string`

GetCutoffStatement returns the CutoffStatement field if non-nil, zero value otherwise.

### GetCutoffStatementOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetCutoffStatementOk() (*string, bool)`

GetCutoffStatementOk returns a tuple with the CutoffStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoffStatement

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetCutoffStatement(v string)`

SetCutoffStatement sets CutoffStatement field to given value.


### SetCutoffStatementNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetCutoffStatementNil(b bool)`

 SetCutoffStatementNil sets the value for CutoffStatement to be an explicit nil

### UnsetCutoffStatement
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) UnsetCutoffStatement()`

UnsetCutoffStatement ensures that no value is present for CutoffStatement, not even an explicit nil
### GetHistoryRunsPastCutoff

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetHistoryRunsPastCutoff() bool`

GetHistoryRunsPastCutoff returns the HistoryRunsPastCutoff field if non-nil, zero value otherwise.

### GetHistoryRunsPastCutoffOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetHistoryRunsPastCutoffOk() (*bool, bool)`

GetHistoryRunsPastCutoffOk returns a tuple with the HistoryRunsPastCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryRunsPastCutoff

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetHistoryRunsPastCutoff(v bool)`

SetHistoryRunsPastCutoff sets HistoryRunsPastCutoff field to given value.


### SetHistoryRunsPastCutoffNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetHistoryRunsPastCutoffNil(b bool)`

 SetHistoryRunsPastCutoffNil sets the value for HistoryRunsPastCutoff to be an explicit nil

### UnsetHistoryRunsPastCutoff
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) UnsetHistoryRunsPastCutoff()`

UnsetHistoryRunsPastCutoff ensures that no value is present for HistoryRunsPastCutoff, not even an explicit nil
### GetNewestHistoryAsOf

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetNewestHistoryAsOf() string`

GetNewestHistoryAsOf returns the NewestHistoryAsOf field if non-nil, zero value otherwise.

### GetNewestHistoryAsOfOk

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) GetNewestHistoryAsOfOk() (*string, bool)`

GetNewestHistoryAsOfOk returns a tuple with the NewestHistoryAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewestHistoryAsOf

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetNewestHistoryAsOf(v string)`

SetNewestHistoryAsOf sets NewestHistoryAsOf field to given value.


### SetNewestHistoryAsOfNil

`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) SetNewestHistoryAsOfNil(b bool)`

 SetNewestHistoryAsOfNil sets the value for NewestHistoryAsOf to be an explicit nil

### UnsetNewestHistoryAsOf
`func (o *ApiV1TrialsHistoryAndPredictionGet200ResponsePredictionsInner) UnsetNewestHistoryAsOf()`

UnsetNewestHistoryAsOf ensures that no value is present for NewestHistoryAsOf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


