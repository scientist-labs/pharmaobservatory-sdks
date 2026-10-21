# TrialKolsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key. | [optional] 
**InvestigatorPersonName** | Pointer to **NullableString** | The investigator&#39;s name as filed. | [optional] 
**InvestigatorAffiliation** | Pointer to **NullableString** | The affiliation as filed. | [optional] 
**Roles** | Pointer to **[]string** | The roles the person holds on this trial as registered. | [optional] 
**IsPrincipalInvestigator** | Pointer to **NullableBool** | True when the person is named principal investigator on this trial. | [optional] 
**AuthoredPubs** | Pointer to **NullableFloat64** | Publications the person authored. | [optional] 
**PiTrials** | Pointer to **NullableInt32** | Trials on which the person is principal investigator, across the registry. | [optional] 

## Methods

### NewTrialKolsRow

`func NewTrialKolsRow() *TrialKolsRow`

NewTrialKolsRow instantiates a new TrialKolsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialKolsRowWithDefaults

`func NewTrialKolsRowWithDefaults() *TrialKolsRow`

NewTrialKolsRowWithDefaults instantiates a new TrialKolsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialKolsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialKolsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialKolsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialKolsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialKolsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialKolsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetInvestigatorPersonId

`func (o *TrialKolsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *TrialKolsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *TrialKolsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *TrialKolsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *TrialKolsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *TrialKolsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetInvestigatorPersonName

`func (o *TrialKolsRow) GetInvestigatorPersonName() string`

GetInvestigatorPersonName returns the InvestigatorPersonName field if non-nil, zero value otherwise.

### GetInvestigatorPersonNameOk

`func (o *TrialKolsRow) GetInvestigatorPersonNameOk() (*string, bool)`

GetInvestigatorPersonNameOk returns a tuple with the InvestigatorPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonName

`func (o *TrialKolsRow) SetInvestigatorPersonName(v string)`

SetInvestigatorPersonName sets InvestigatorPersonName field to given value.

### HasInvestigatorPersonName

`func (o *TrialKolsRow) HasInvestigatorPersonName() bool`

HasInvestigatorPersonName returns a boolean if a field has been set.

### SetInvestigatorPersonNameNil

`func (o *TrialKolsRow) SetInvestigatorPersonNameNil(b bool)`

 SetInvestigatorPersonNameNil sets the value for InvestigatorPersonName to be an explicit nil

### UnsetInvestigatorPersonName
`func (o *TrialKolsRow) UnsetInvestigatorPersonName()`

UnsetInvestigatorPersonName ensures that no value is present for InvestigatorPersonName, not even an explicit nil
### GetInvestigatorAffiliation

`func (o *TrialKolsRow) GetInvestigatorAffiliation() string`

GetInvestigatorAffiliation returns the InvestigatorAffiliation field if non-nil, zero value otherwise.

### GetInvestigatorAffiliationOk

`func (o *TrialKolsRow) GetInvestigatorAffiliationOk() (*string, bool)`

GetInvestigatorAffiliationOk returns a tuple with the InvestigatorAffiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorAffiliation

`func (o *TrialKolsRow) SetInvestigatorAffiliation(v string)`

SetInvestigatorAffiliation sets InvestigatorAffiliation field to given value.

### HasInvestigatorAffiliation

`func (o *TrialKolsRow) HasInvestigatorAffiliation() bool`

HasInvestigatorAffiliation returns a boolean if a field has been set.

### SetInvestigatorAffiliationNil

`func (o *TrialKolsRow) SetInvestigatorAffiliationNil(b bool)`

 SetInvestigatorAffiliationNil sets the value for InvestigatorAffiliation to be an explicit nil

### UnsetInvestigatorAffiliation
`func (o *TrialKolsRow) UnsetInvestigatorAffiliation()`

UnsetInvestigatorAffiliation ensures that no value is present for InvestigatorAffiliation, not even an explicit nil
### GetRoles

`func (o *TrialKolsRow) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *TrialKolsRow) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *TrialKolsRow) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *TrialKolsRow) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *TrialKolsRow) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *TrialKolsRow) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetIsPrincipalInvestigator

`func (o *TrialKolsRow) GetIsPrincipalInvestigator() bool`

GetIsPrincipalInvestigator returns the IsPrincipalInvestigator field if non-nil, zero value otherwise.

### GetIsPrincipalInvestigatorOk

`func (o *TrialKolsRow) GetIsPrincipalInvestigatorOk() (*bool, bool)`

GetIsPrincipalInvestigatorOk returns a tuple with the IsPrincipalInvestigator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrincipalInvestigator

`func (o *TrialKolsRow) SetIsPrincipalInvestigator(v bool)`

SetIsPrincipalInvestigator sets IsPrincipalInvestigator field to given value.

### HasIsPrincipalInvestigator

`func (o *TrialKolsRow) HasIsPrincipalInvestigator() bool`

HasIsPrincipalInvestigator returns a boolean if a field has been set.

### SetIsPrincipalInvestigatorNil

`func (o *TrialKolsRow) SetIsPrincipalInvestigatorNil(b bool)`

 SetIsPrincipalInvestigatorNil sets the value for IsPrincipalInvestigator to be an explicit nil

### UnsetIsPrincipalInvestigator
`func (o *TrialKolsRow) UnsetIsPrincipalInvestigator()`

UnsetIsPrincipalInvestigator ensures that no value is present for IsPrincipalInvestigator, not even an explicit nil
### GetAuthoredPubs

`func (o *TrialKolsRow) GetAuthoredPubs() float64`

GetAuthoredPubs returns the AuthoredPubs field if non-nil, zero value otherwise.

### GetAuthoredPubsOk

`func (o *TrialKolsRow) GetAuthoredPubsOk() (*float64, bool)`

GetAuthoredPubsOk returns a tuple with the AuthoredPubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredPubs

`func (o *TrialKolsRow) SetAuthoredPubs(v float64)`

SetAuthoredPubs sets AuthoredPubs field to given value.

### HasAuthoredPubs

`func (o *TrialKolsRow) HasAuthoredPubs() bool`

HasAuthoredPubs returns a boolean if a field has been set.

### SetAuthoredPubsNil

`func (o *TrialKolsRow) SetAuthoredPubsNil(b bool)`

 SetAuthoredPubsNil sets the value for AuthoredPubs to be an explicit nil

### UnsetAuthoredPubs
`func (o *TrialKolsRow) UnsetAuthoredPubs()`

UnsetAuthoredPubs ensures that no value is present for AuthoredPubs, not even an explicit nil
### GetPiTrials

`func (o *TrialKolsRow) GetPiTrials() int32`

GetPiTrials returns the PiTrials field if non-nil, zero value otherwise.

### GetPiTrialsOk

`func (o *TrialKolsRow) GetPiTrialsOk() (*int32, bool)`

GetPiTrialsOk returns a tuple with the PiTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiTrials

`func (o *TrialKolsRow) SetPiTrials(v int32)`

SetPiTrials sets PiTrials field to given value.

### HasPiTrials

`func (o *TrialKolsRow) HasPiTrials() bool`

HasPiTrials returns a boolean if a field has been set.

### SetPiTrialsNil

`func (o *TrialKolsRow) SetPiTrialsNil(b bool)`

 SetPiTrialsNil sets the value for PiTrials to be an explicit nil

### UnsetPiTrials
`func (o *TrialKolsRow) UnsetPiTrials()`

UnsetPiTrials ensures that no value is present for PiTrials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


