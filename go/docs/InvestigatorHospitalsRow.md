# InvestigatorHospitalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**Ccn** | Pointer to **NullableString** | CMS Certification Number, the six-character Medicare provider id the hospital page address carries. | [optional] 
**HospitalName** | Pointer to **NullableString** | The hospital&#39;s facility name in the CMS directory. | [optional] 
**City** | Pointer to **NullableString** | The hospital&#39;s city in the CMS directory. | [optional] 
**State** | Pointer to **NullableString** | The hospital&#39;s state in the CMS directory. | [optional] 
**FacilityTypes** | Pointer to **[]string** | The CMS facility types of the affiliation. | [optional] 

## Methods

### NewInvestigatorHospitalsRow

`func NewInvestigatorHospitalsRow() *InvestigatorHospitalsRow`

NewInvestigatorHospitalsRow instantiates a new InvestigatorHospitalsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigatorHospitalsRowWithDefaults

`func NewInvestigatorHospitalsRowWithDefaults() *InvestigatorHospitalsRow`

NewInvestigatorHospitalsRowWithDefaults instantiates a new InvestigatorHospitalsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestigatorPersonId

`func (o *InvestigatorHospitalsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *InvestigatorHospitalsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *InvestigatorHospitalsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *InvestigatorHospitalsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *InvestigatorHospitalsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *InvestigatorHospitalsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetCcn

`func (o *InvestigatorHospitalsRow) GetCcn() string`

GetCcn returns the Ccn field if non-nil, zero value otherwise.

### GetCcnOk

`func (o *InvestigatorHospitalsRow) GetCcnOk() (*string, bool)`

GetCcnOk returns a tuple with the Ccn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcn

`func (o *InvestigatorHospitalsRow) SetCcn(v string)`

SetCcn sets Ccn field to given value.

### HasCcn

`func (o *InvestigatorHospitalsRow) HasCcn() bool`

HasCcn returns a boolean if a field has been set.

### SetCcnNil

`func (o *InvestigatorHospitalsRow) SetCcnNil(b bool)`

 SetCcnNil sets the value for Ccn to be an explicit nil

### UnsetCcn
`func (o *InvestigatorHospitalsRow) UnsetCcn()`

UnsetCcn ensures that no value is present for Ccn, not even an explicit nil
### GetHospitalName

`func (o *InvestigatorHospitalsRow) GetHospitalName() string`

GetHospitalName returns the HospitalName field if non-nil, zero value otherwise.

### GetHospitalNameOk

`func (o *InvestigatorHospitalsRow) GetHospitalNameOk() (*string, bool)`

GetHospitalNameOk returns a tuple with the HospitalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHospitalName

`func (o *InvestigatorHospitalsRow) SetHospitalName(v string)`

SetHospitalName sets HospitalName field to given value.

### HasHospitalName

`func (o *InvestigatorHospitalsRow) HasHospitalName() bool`

HasHospitalName returns a boolean if a field has been set.

### SetHospitalNameNil

`func (o *InvestigatorHospitalsRow) SetHospitalNameNil(b bool)`

 SetHospitalNameNil sets the value for HospitalName to be an explicit nil

### UnsetHospitalName
`func (o *InvestigatorHospitalsRow) UnsetHospitalName()`

UnsetHospitalName ensures that no value is present for HospitalName, not even an explicit nil
### GetCity

`func (o *InvestigatorHospitalsRow) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *InvestigatorHospitalsRow) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *InvestigatorHospitalsRow) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *InvestigatorHospitalsRow) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *InvestigatorHospitalsRow) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *InvestigatorHospitalsRow) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *InvestigatorHospitalsRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InvestigatorHospitalsRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InvestigatorHospitalsRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InvestigatorHospitalsRow) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *InvestigatorHospitalsRow) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *InvestigatorHospitalsRow) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetFacilityTypes

`func (o *InvestigatorHospitalsRow) GetFacilityTypes() []string`

GetFacilityTypes returns the FacilityTypes field if non-nil, zero value otherwise.

### GetFacilityTypesOk

`func (o *InvestigatorHospitalsRow) GetFacilityTypesOk() (*[]string, bool)`

GetFacilityTypesOk returns a tuple with the FacilityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityTypes

`func (o *InvestigatorHospitalsRow) SetFacilityTypes(v []string)`

SetFacilityTypes sets FacilityTypes field to given value.

### HasFacilityTypes

`func (o *InvestigatorHospitalsRow) HasFacilityTypes() bool`

HasFacilityTypes returns a boolean if a field has been set.

### SetFacilityTypesNil

`func (o *InvestigatorHospitalsRow) SetFacilityTypesNil(b bool)`

 SetFacilityTypesNil sets the value for FacilityTypes to be an explicit nil

### UnsetFacilityTypes
`func (o *InvestigatorHospitalsRow) UnsetFacilityTypes()`

UnsetFacilityTypes ensures that no value is present for FacilityTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


