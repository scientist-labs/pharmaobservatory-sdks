# DrugInvestigatorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The drug&#39;s resolved id (the same value as molecule_id on the drugs dataset). The entry key for the drugs object type. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name as the drugs dataset names it. | [optional] 
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**InvestigatorPersonName** | Pointer to **NullableString** | The investigator&#39;s name as the registry filed it, resolved across the person&#39;s merged name fragments. | [optional] 
**InvestigatorAffiliation** | Pointer to **NullableString** | The investigator&#39;s affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional] 
**NTrials** | Pointer to **NullableInt64** | Distinct trials the person is named on that are in scope of the row (the target&#39;s trials, or the drug&#39;s). | [optional] 
**IsPrincipalInvestigator** | Pointer to **NullableBool** | True when the person holds the principal investigator role on any registered trial (the person-level flag the investigators dataset and the people page publish), not only the trials of this row. | [optional] 
**Roles** | Pointer to **[]string** | The distinct roles the person holds on the trials of this row, as the registry filed them (PRINCIPAL_INVESTIGATOR, STUDY_DIRECTOR, STUDY_CHAIR, SUB_INVESTIGATOR). | [optional] 

## Methods

### NewDrugInvestigatorsRow

`func NewDrugInvestigatorsRow() *DrugInvestigatorsRow`

NewDrugInvestigatorsRow instantiates a new DrugInvestigatorsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugInvestigatorsRowWithDefaults

`func NewDrugInvestigatorsRowWithDefaults() *DrugInvestigatorsRow`

NewDrugInvestigatorsRowWithDefaults instantiates a new DrugInvestigatorsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugInvestigatorsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugInvestigatorsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugInvestigatorsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugInvestigatorsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugInvestigatorsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugInvestigatorsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetDrugName

`func (o *DrugInvestigatorsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *DrugInvestigatorsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *DrugInvestigatorsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *DrugInvestigatorsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *DrugInvestigatorsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *DrugInvestigatorsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetInvestigatorPersonId

`func (o *DrugInvestigatorsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *DrugInvestigatorsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *DrugInvestigatorsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *DrugInvestigatorsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *DrugInvestigatorsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *DrugInvestigatorsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetInvestigatorPersonName

`func (o *DrugInvestigatorsRow) GetInvestigatorPersonName() string`

GetInvestigatorPersonName returns the InvestigatorPersonName field if non-nil, zero value otherwise.

### GetInvestigatorPersonNameOk

`func (o *DrugInvestigatorsRow) GetInvestigatorPersonNameOk() (*string, bool)`

GetInvestigatorPersonNameOk returns a tuple with the InvestigatorPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonName

`func (o *DrugInvestigatorsRow) SetInvestigatorPersonName(v string)`

SetInvestigatorPersonName sets InvestigatorPersonName field to given value.

### HasInvestigatorPersonName

`func (o *DrugInvestigatorsRow) HasInvestigatorPersonName() bool`

HasInvestigatorPersonName returns a boolean if a field has been set.

### SetInvestigatorPersonNameNil

`func (o *DrugInvestigatorsRow) SetInvestigatorPersonNameNil(b bool)`

 SetInvestigatorPersonNameNil sets the value for InvestigatorPersonName to be an explicit nil

### UnsetInvestigatorPersonName
`func (o *DrugInvestigatorsRow) UnsetInvestigatorPersonName()`

UnsetInvestigatorPersonName ensures that no value is present for InvestigatorPersonName, not even an explicit nil
### GetInvestigatorAffiliation

`func (o *DrugInvestigatorsRow) GetInvestigatorAffiliation() string`

GetInvestigatorAffiliation returns the InvestigatorAffiliation field if non-nil, zero value otherwise.

### GetInvestigatorAffiliationOk

`func (o *DrugInvestigatorsRow) GetInvestigatorAffiliationOk() (*string, bool)`

GetInvestigatorAffiliationOk returns a tuple with the InvestigatorAffiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorAffiliation

`func (o *DrugInvestigatorsRow) SetInvestigatorAffiliation(v string)`

SetInvestigatorAffiliation sets InvestigatorAffiliation field to given value.

### HasInvestigatorAffiliation

`func (o *DrugInvestigatorsRow) HasInvestigatorAffiliation() bool`

HasInvestigatorAffiliation returns a boolean if a field has been set.

### SetInvestigatorAffiliationNil

`func (o *DrugInvestigatorsRow) SetInvestigatorAffiliationNil(b bool)`

 SetInvestigatorAffiliationNil sets the value for InvestigatorAffiliation to be an explicit nil

### UnsetInvestigatorAffiliation
`func (o *DrugInvestigatorsRow) UnsetInvestigatorAffiliation()`

UnsetInvestigatorAffiliation ensures that no value is present for InvestigatorAffiliation, not even an explicit nil
### GetNTrials

`func (o *DrugInvestigatorsRow) GetNTrials() int64`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *DrugInvestigatorsRow) GetNTrialsOk() (*int64, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *DrugInvestigatorsRow) SetNTrials(v int64)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *DrugInvestigatorsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *DrugInvestigatorsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *DrugInvestigatorsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetIsPrincipalInvestigator

`func (o *DrugInvestigatorsRow) GetIsPrincipalInvestigator() bool`

GetIsPrincipalInvestigator returns the IsPrincipalInvestigator field if non-nil, zero value otherwise.

### GetIsPrincipalInvestigatorOk

`func (o *DrugInvestigatorsRow) GetIsPrincipalInvestigatorOk() (*bool, bool)`

GetIsPrincipalInvestigatorOk returns a tuple with the IsPrincipalInvestigator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrincipalInvestigator

`func (o *DrugInvestigatorsRow) SetIsPrincipalInvestigator(v bool)`

SetIsPrincipalInvestigator sets IsPrincipalInvestigator field to given value.

### HasIsPrincipalInvestigator

`func (o *DrugInvestigatorsRow) HasIsPrincipalInvestigator() bool`

HasIsPrincipalInvestigator returns a boolean if a field has been set.

### SetIsPrincipalInvestigatorNil

`func (o *DrugInvestigatorsRow) SetIsPrincipalInvestigatorNil(b bool)`

 SetIsPrincipalInvestigatorNil sets the value for IsPrincipalInvestigator to be an explicit nil

### UnsetIsPrincipalInvestigator
`func (o *DrugInvestigatorsRow) UnsetIsPrincipalInvestigator()`

UnsetIsPrincipalInvestigator ensures that no value is present for IsPrincipalInvestigator, not even an explicit nil
### GetRoles

`func (o *DrugInvestigatorsRow) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DrugInvestigatorsRow) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DrugInvestigatorsRow) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *DrugInvestigatorsRow) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *DrugInvestigatorsRow) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *DrugInvestigatorsRow) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


