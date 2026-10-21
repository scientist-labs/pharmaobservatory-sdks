# DrugLabelSafetyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**HasBoxedWarning** | Pointer to **NullableBool** | True when the label carries a boxed warning. | [optional] 
**BoxedWarningText** | Pointer to **NullableString** | The boxed warning text, null when there is none. | [optional] 
**Warnings** | Pointer to **NullableString** | The warnings and precautions text. | [optional] 
**LabelEffectiveDate** | Pointer to **NullableString** | The effective date of the label. | [optional] 
**LabelCount** | Pointer to **NullableInt32** | Labels matched to the drug. | [optional] 
**SetId** | Pointer to **NullableString** | The label set id. The DailyMed page is https://dailymed.nlm.nih.gov/dailymed/drugInfo.cfm?setid&#x3D; followed by this id. | [optional] 
**MatchMethod** | Pointer to **NullableString** | How the label was matched to the drug. | [optional] 

## Methods

### NewDrugLabelSafetyRow

`func NewDrugLabelSafetyRow() *DrugLabelSafetyRow`

NewDrugLabelSafetyRow instantiates a new DrugLabelSafetyRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugLabelSafetyRowWithDefaults

`func NewDrugLabelSafetyRowWithDefaults() *DrugLabelSafetyRow`

NewDrugLabelSafetyRowWithDefaults instantiates a new DrugLabelSafetyRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugLabelSafetyRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugLabelSafetyRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugLabelSafetyRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugLabelSafetyRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugLabelSafetyRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugLabelSafetyRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetHasBoxedWarning

`func (o *DrugLabelSafetyRow) GetHasBoxedWarning() bool`

GetHasBoxedWarning returns the HasBoxedWarning field if non-nil, zero value otherwise.

### GetHasBoxedWarningOk

`func (o *DrugLabelSafetyRow) GetHasBoxedWarningOk() (*bool, bool)`

GetHasBoxedWarningOk returns a tuple with the HasBoxedWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBoxedWarning

`func (o *DrugLabelSafetyRow) SetHasBoxedWarning(v bool)`

SetHasBoxedWarning sets HasBoxedWarning field to given value.

### HasHasBoxedWarning

`func (o *DrugLabelSafetyRow) HasHasBoxedWarning() bool`

HasHasBoxedWarning returns a boolean if a field has been set.

### SetHasBoxedWarningNil

`func (o *DrugLabelSafetyRow) SetHasBoxedWarningNil(b bool)`

 SetHasBoxedWarningNil sets the value for HasBoxedWarning to be an explicit nil

### UnsetHasBoxedWarning
`func (o *DrugLabelSafetyRow) UnsetHasBoxedWarning()`

UnsetHasBoxedWarning ensures that no value is present for HasBoxedWarning, not even an explicit nil
### GetBoxedWarningText

`func (o *DrugLabelSafetyRow) GetBoxedWarningText() string`

GetBoxedWarningText returns the BoxedWarningText field if non-nil, zero value otherwise.

### GetBoxedWarningTextOk

`func (o *DrugLabelSafetyRow) GetBoxedWarningTextOk() (*string, bool)`

GetBoxedWarningTextOk returns a tuple with the BoxedWarningText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxedWarningText

`func (o *DrugLabelSafetyRow) SetBoxedWarningText(v string)`

SetBoxedWarningText sets BoxedWarningText field to given value.

### HasBoxedWarningText

`func (o *DrugLabelSafetyRow) HasBoxedWarningText() bool`

HasBoxedWarningText returns a boolean if a field has been set.

### SetBoxedWarningTextNil

`func (o *DrugLabelSafetyRow) SetBoxedWarningTextNil(b bool)`

 SetBoxedWarningTextNil sets the value for BoxedWarningText to be an explicit nil

### UnsetBoxedWarningText
`func (o *DrugLabelSafetyRow) UnsetBoxedWarningText()`

UnsetBoxedWarningText ensures that no value is present for BoxedWarningText, not even an explicit nil
### GetWarnings

`func (o *DrugLabelSafetyRow) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DrugLabelSafetyRow) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DrugLabelSafetyRow) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DrugLabelSafetyRow) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *DrugLabelSafetyRow) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *DrugLabelSafetyRow) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetLabelEffectiveDate

`func (o *DrugLabelSafetyRow) GetLabelEffectiveDate() string`

GetLabelEffectiveDate returns the LabelEffectiveDate field if non-nil, zero value otherwise.

### GetLabelEffectiveDateOk

`func (o *DrugLabelSafetyRow) GetLabelEffectiveDateOk() (*string, bool)`

GetLabelEffectiveDateOk returns a tuple with the LabelEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelEffectiveDate

`func (o *DrugLabelSafetyRow) SetLabelEffectiveDate(v string)`

SetLabelEffectiveDate sets LabelEffectiveDate field to given value.

### HasLabelEffectiveDate

`func (o *DrugLabelSafetyRow) HasLabelEffectiveDate() bool`

HasLabelEffectiveDate returns a boolean if a field has been set.

### SetLabelEffectiveDateNil

`func (o *DrugLabelSafetyRow) SetLabelEffectiveDateNil(b bool)`

 SetLabelEffectiveDateNil sets the value for LabelEffectiveDate to be an explicit nil

### UnsetLabelEffectiveDate
`func (o *DrugLabelSafetyRow) UnsetLabelEffectiveDate()`

UnsetLabelEffectiveDate ensures that no value is present for LabelEffectiveDate, not even an explicit nil
### GetLabelCount

`func (o *DrugLabelSafetyRow) GetLabelCount() int32`

GetLabelCount returns the LabelCount field if non-nil, zero value otherwise.

### GetLabelCountOk

`func (o *DrugLabelSafetyRow) GetLabelCountOk() (*int32, bool)`

GetLabelCountOk returns a tuple with the LabelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCount

`func (o *DrugLabelSafetyRow) SetLabelCount(v int32)`

SetLabelCount sets LabelCount field to given value.

### HasLabelCount

`func (o *DrugLabelSafetyRow) HasLabelCount() bool`

HasLabelCount returns a boolean if a field has been set.

### SetLabelCountNil

`func (o *DrugLabelSafetyRow) SetLabelCountNil(b bool)`

 SetLabelCountNil sets the value for LabelCount to be an explicit nil

### UnsetLabelCount
`func (o *DrugLabelSafetyRow) UnsetLabelCount()`

UnsetLabelCount ensures that no value is present for LabelCount, not even an explicit nil
### GetSetId

`func (o *DrugLabelSafetyRow) GetSetId() string`

GetSetId returns the SetId field if non-nil, zero value otherwise.

### GetSetIdOk

`func (o *DrugLabelSafetyRow) GetSetIdOk() (*string, bool)`

GetSetIdOk returns a tuple with the SetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetId

`func (o *DrugLabelSafetyRow) SetSetId(v string)`

SetSetId sets SetId field to given value.

### HasSetId

`func (o *DrugLabelSafetyRow) HasSetId() bool`

HasSetId returns a boolean if a field has been set.

### SetSetIdNil

`func (o *DrugLabelSafetyRow) SetSetIdNil(b bool)`

 SetSetIdNil sets the value for SetId to be an explicit nil

### UnsetSetId
`func (o *DrugLabelSafetyRow) UnsetSetId()`

UnsetSetId ensures that no value is present for SetId, not even an explicit nil
### GetMatchMethod

`func (o *DrugLabelSafetyRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *DrugLabelSafetyRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *DrugLabelSafetyRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *DrugLabelSafetyRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *DrugLabelSafetyRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *DrugLabelSafetyRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


