# DrugRecallsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**RecallNumber** | Pointer to **NullableString** | The FDA recall number. | [optional] 
**Classification** | Pointer to **NullableString** | The FDA recall class (Class I, Class II or Class III). | [optional] 
**Status** | Pointer to **NullableString** | The recall status (for example Ongoing, Completed, Terminated). | [optional] 
**RecallingFirm** | Pointer to **NullableString** | The firm that recalled the product, as filed. | [optional] 
**ProductDescription** | Pointer to **NullableString** | The recalled product as described in the report. | [optional] 
**ReasonForRecall** | Pointer to **NullableString** | The reason given for the recall. | [optional] 
**RecallInitiationDate** | Pointer to **NullableString** | The date the recall began. | [optional] 
**ReportDate** | Pointer to **NullableString** | The date the FDA posted the report. | [optional] 

## Methods

### NewDrugRecallsRow

`func NewDrugRecallsRow() *DrugRecallsRow`

NewDrugRecallsRow instantiates a new DrugRecallsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugRecallsRowWithDefaults

`func NewDrugRecallsRowWithDefaults() *DrugRecallsRow`

NewDrugRecallsRowWithDefaults instantiates a new DrugRecallsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugRecallsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugRecallsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugRecallsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugRecallsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugRecallsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugRecallsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetRecallNumber

`func (o *DrugRecallsRow) GetRecallNumber() string`

GetRecallNumber returns the RecallNumber field if non-nil, zero value otherwise.

### GetRecallNumberOk

`func (o *DrugRecallsRow) GetRecallNumberOk() (*string, bool)`

GetRecallNumberOk returns a tuple with the RecallNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallNumber

`func (o *DrugRecallsRow) SetRecallNumber(v string)`

SetRecallNumber sets RecallNumber field to given value.

### HasRecallNumber

`func (o *DrugRecallsRow) HasRecallNumber() bool`

HasRecallNumber returns a boolean if a field has been set.

### SetRecallNumberNil

`func (o *DrugRecallsRow) SetRecallNumberNil(b bool)`

 SetRecallNumberNil sets the value for RecallNumber to be an explicit nil

### UnsetRecallNumber
`func (o *DrugRecallsRow) UnsetRecallNumber()`

UnsetRecallNumber ensures that no value is present for RecallNumber, not even an explicit nil
### GetClassification

`func (o *DrugRecallsRow) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *DrugRecallsRow) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *DrugRecallsRow) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *DrugRecallsRow) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *DrugRecallsRow) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *DrugRecallsRow) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetStatus

`func (o *DrugRecallsRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DrugRecallsRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DrugRecallsRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DrugRecallsRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DrugRecallsRow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DrugRecallsRow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRecallingFirm

`func (o *DrugRecallsRow) GetRecallingFirm() string`

GetRecallingFirm returns the RecallingFirm field if non-nil, zero value otherwise.

### GetRecallingFirmOk

`func (o *DrugRecallsRow) GetRecallingFirmOk() (*string, bool)`

GetRecallingFirmOk returns a tuple with the RecallingFirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallingFirm

`func (o *DrugRecallsRow) SetRecallingFirm(v string)`

SetRecallingFirm sets RecallingFirm field to given value.

### HasRecallingFirm

`func (o *DrugRecallsRow) HasRecallingFirm() bool`

HasRecallingFirm returns a boolean if a field has been set.

### SetRecallingFirmNil

`func (o *DrugRecallsRow) SetRecallingFirmNil(b bool)`

 SetRecallingFirmNil sets the value for RecallingFirm to be an explicit nil

### UnsetRecallingFirm
`func (o *DrugRecallsRow) UnsetRecallingFirm()`

UnsetRecallingFirm ensures that no value is present for RecallingFirm, not even an explicit nil
### GetProductDescription

`func (o *DrugRecallsRow) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *DrugRecallsRow) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *DrugRecallsRow) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *DrugRecallsRow) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### SetProductDescriptionNil

`func (o *DrugRecallsRow) SetProductDescriptionNil(b bool)`

 SetProductDescriptionNil sets the value for ProductDescription to be an explicit nil

### UnsetProductDescription
`func (o *DrugRecallsRow) UnsetProductDescription()`

UnsetProductDescription ensures that no value is present for ProductDescription, not even an explicit nil
### GetReasonForRecall

`func (o *DrugRecallsRow) GetReasonForRecall() string`

GetReasonForRecall returns the ReasonForRecall field if non-nil, zero value otherwise.

### GetReasonForRecallOk

`func (o *DrugRecallsRow) GetReasonForRecallOk() (*string, bool)`

GetReasonForRecallOk returns a tuple with the ReasonForRecall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForRecall

`func (o *DrugRecallsRow) SetReasonForRecall(v string)`

SetReasonForRecall sets ReasonForRecall field to given value.

### HasReasonForRecall

`func (o *DrugRecallsRow) HasReasonForRecall() bool`

HasReasonForRecall returns a boolean if a field has been set.

### SetReasonForRecallNil

`func (o *DrugRecallsRow) SetReasonForRecallNil(b bool)`

 SetReasonForRecallNil sets the value for ReasonForRecall to be an explicit nil

### UnsetReasonForRecall
`func (o *DrugRecallsRow) UnsetReasonForRecall()`

UnsetReasonForRecall ensures that no value is present for ReasonForRecall, not even an explicit nil
### GetRecallInitiationDate

`func (o *DrugRecallsRow) GetRecallInitiationDate() string`

GetRecallInitiationDate returns the RecallInitiationDate field if non-nil, zero value otherwise.

### GetRecallInitiationDateOk

`func (o *DrugRecallsRow) GetRecallInitiationDateOk() (*string, bool)`

GetRecallInitiationDateOk returns a tuple with the RecallInitiationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallInitiationDate

`func (o *DrugRecallsRow) SetRecallInitiationDate(v string)`

SetRecallInitiationDate sets RecallInitiationDate field to given value.

### HasRecallInitiationDate

`func (o *DrugRecallsRow) HasRecallInitiationDate() bool`

HasRecallInitiationDate returns a boolean if a field has been set.

### SetRecallInitiationDateNil

`func (o *DrugRecallsRow) SetRecallInitiationDateNil(b bool)`

 SetRecallInitiationDateNil sets the value for RecallInitiationDate to be an explicit nil

### UnsetRecallInitiationDate
`func (o *DrugRecallsRow) UnsetRecallInitiationDate()`

UnsetRecallInitiationDate ensures that no value is present for RecallInitiationDate, not even an explicit nil
### GetReportDate

`func (o *DrugRecallsRow) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *DrugRecallsRow) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *DrugRecallsRow) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *DrugRecallsRow) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### SetReportDateNil

`func (o *DrugRecallsRow) SetReportDateNil(b bool)`

 SetReportDateNil sets the value for ReportDate to be an explicit nil

### UnsetReportDate
`func (o *DrugRecallsRow) UnsetReportDate()`

UnsetReportDate ensures that no value is present for ReportDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


