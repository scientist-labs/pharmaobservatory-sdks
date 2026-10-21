# DrugPatentMentionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**PatentNumber** | Pointer to **NullableString** | The patent publication number. | [optional] 
**Country** | Pointer to **NullableString** | The patent office country code. | [optional] 
**PublicationDate** | Pointer to **NullableString** | The publication date. | [optional] 
**Title** | Pointer to **NullableString** | The patent title as filed. | [optional] 
**Assignee** | Pointer to **NullableString** | The assignee as filed. Not an ownership statement. | [optional] 

## Methods

### NewDrugPatentMentionsRow

`func NewDrugPatentMentionsRow() *DrugPatentMentionsRow`

NewDrugPatentMentionsRow instantiates a new DrugPatentMentionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugPatentMentionsRowWithDefaults

`func NewDrugPatentMentionsRowWithDefaults() *DrugPatentMentionsRow`

NewDrugPatentMentionsRowWithDefaults instantiates a new DrugPatentMentionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugPatentMentionsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugPatentMentionsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugPatentMentionsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugPatentMentionsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugPatentMentionsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugPatentMentionsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetPatentNumber

`func (o *DrugPatentMentionsRow) GetPatentNumber() string`

GetPatentNumber returns the PatentNumber field if non-nil, zero value otherwise.

### GetPatentNumberOk

`func (o *DrugPatentMentionsRow) GetPatentNumberOk() (*string, bool)`

GetPatentNumberOk returns a tuple with the PatentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentNumber

`func (o *DrugPatentMentionsRow) SetPatentNumber(v string)`

SetPatentNumber sets PatentNumber field to given value.

### HasPatentNumber

`func (o *DrugPatentMentionsRow) HasPatentNumber() bool`

HasPatentNumber returns a boolean if a field has been set.

### SetPatentNumberNil

`func (o *DrugPatentMentionsRow) SetPatentNumberNil(b bool)`

 SetPatentNumberNil sets the value for PatentNumber to be an explicit nil

### UnsetPatentNumber
`func (o *DrugPatentMentionsRow) UnsetPatentNumber()`

UnsetPatentNumber ensures that no value is present for PatentNumber, not even an explicit nil
### GetCountry

`func (o *DrugPatentMentionsRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DrugPatentMentionsRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DrugPatentMentionsRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DrugPatentMentionsRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *DrugPatentMentionsRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *DrugPatentMentionsRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPublicationDate

`func (o *DrugPatentMentionsRow) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *DrugPatentMentionsRow) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *DrugPatentMentionsRow) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *DrugPatentMentionsRow) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### SetPublicationDateNil

`func (o *DrugPatentMentionsRow) SetPublicationDateNil(b bool)`

 SetPublicationDateNil sets the value for PublicationDate to be an explicit nil

### UnsetPublicationDate
`func (o *DrugPatentMentionsRow) UnsetPublicationDate()`

UnsetPublicationDate ensures that no value is present for PublicationDate, not even an explicit nil
### GetTitle

`func (o *DrugPatentMentionsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DrugPatentMentionsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DrugPatentMentionsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DrugPatentMentionsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DrugPatentMentionsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DrugPatentMentionsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAssignee

`func (o *DrugPatentMentionsRow) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *DrugPatentMentionsRow) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *DrugPatentMentionsRow) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *DrugPatentMentionsRow) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *DrugPatentMentionsRow) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *DrugPatentMentionsRow) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


