# SponsorPatentMentionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the lead sponsor organization record. | [optional] 
**SponsorName** | Pointer to **NullableString** | Display name of the sponsor organization record. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the sponsor belongs to (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**DrugId** | Pointer to **NullableString** | The drug key. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**PatentNumber** | Pointer to **NullableString** | The patent&#39;s publication number. | [optional] 
**Country** | Pointer to **NullableString** | The patent office&#39;s country code. | [optional] 
**PublicationDate** | Pointer to **NullableString** | The patent&#39;s publication date. | [optional] 
**Title** | Pointer to **NullableString** | The patent title. | [optional] 
**Assignee** | Pointer to **NullableString** | The patent&#39;s applicant as filed, free text. | [optional] 

## Methods

### NewSponsorPatentMentionsRow

`func NewSponsorPatentMentionsRow() *SponsorPatentMentionsRow`

NewSponsorPatentMentionsRow instantiates a new SponsorPatentMentionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSponsorPatentMentionsRowWithDefaults

`func NewSponsorPatentMentionsRowWithDefaults() *SponsorPatentMentionsRow`

NewSponsorPatentMentionsRowWithDefaults instantiates a new SponsorPatentMentionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponsorOrgId

`func (o *SponsorPatentMentionsRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *SponsorPatentMentionsRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *SponsorPatentMentionsRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *SponsorPatentMentionsRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *SponsorPatentMentionsRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *SponsorPatentMentionsRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorName

`func (o *SponsorPatentMentionsRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *SponsorPatentMentionsRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *SponsorPatentMentionsRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *SponsorPatentMentionsRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *SponsorPatentMentionsRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *SponsorPatentMentionsRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorCompanyId

`func (o *SponsorPatentMentionsRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *SponsorPatentMentionsRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *SponsorPatentMentionsRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *SponsorPatentMentionsRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *SponsorPatentMentionsRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *SponsorPatentMentionsRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *SponsorPatentMentionsRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *SponsorPatentMentionsRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *SponsorPatentMentionsRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *SponsorPatentMentionsRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *SponsorPatentMentionsRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *SponsorPatentMentionsRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *SponsorPatentMentionsRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *SponsorPatentMentionsRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *SponsorPatentMentionsRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *SponsorPatentMentionsRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *SponsorPatentMentionsRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *SponsorPatentMentionsRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *SponsorPatentMentionsRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *SponsorPatentMentionsRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *SponsorPatentMentionsRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *SponsorPatentMentionsRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *SponsorPatentMentionsRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *SponsorPatentMentionsRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetDrugId

`func (o *SponsorPatentMentionsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *SponsorPatentMentionsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *SponsorPatentMentionsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *SponsorPatentMentionsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *SponsorPatentMentionsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *SponsorPatentMentionsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *SponsorPatentMentionsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *SponsorPatentMentionsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *SponsorPatentMentionsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *SponsorPatentMentionsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *SponsorPatentMentionsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *SponsorPatentMentionsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetPatentNumber

`func (o *SponsorPatentMentionsRow) GetPatentNumber() string`

GetPatentNumber returns the PatentNumber field if non-nil, zero value otherwise.

### GetPatentNumberOk

`func (o *SponsorPatentMentionsRow) GetPatentNumberOk() (*string, bool)`

GetPatentNumberOk returns a tuple with the PatentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentNumber

`func (o *SponsorPatentMentionsRow) SetPatentNumber(v string)`

SetPatentNumber sets PatentNumber field to given value.

### HasPatentNumber

`func (o *SponsorPatentMentionsRow) HasPatentNumber() bool`

HasPatentNumber returns a boolean if a field has been set.

### SetPatentNumberNil

`func (o *SponsorPatentMentionsRow) SetPatentNumberNil(b bool)`

 SetPatentNumberNil sets the value for PatentNumber to be an explicit nil

### UnsetPatentNumber
`func (o *SponsorPatentMentionsRow) UnsetPatentNumber()`

UnsetPatentNumber ensures that no value is present for PatentNumber, not even an explicit nil
### GetCountry

`func (o *SponsorPatentMentionsRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SponsorPatentMentionsRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SponsorPatentMentionsRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SponsorPatentMentionsRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *SponsorPatentMentionsRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *SponsorPatentMentionsRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPublicationDate

`func (o *SponsorPatentMentionsRow) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *SponsorPatentMentionsRow) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *SponsorPatentMentionsRow) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *SponsorPatentMentionsRow) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### SetPublicationDateNil

`func (o *SponsorPatentMentionsRow) SetPublicationDateNil(b bool)`

 SetPublicationDateNil sets the value for PublicationDate to be an explicit nil

### UnsetPublicationDate
`func (o *SponsorPatentMentionsRow) UnsetPublicationDate()`

UnsetPublicationDate ensures that no value is present for PublicationDate, not even an explicit nil
### GetTitle

`func (o *SponsorPatentMentionsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SponsorPatentMentionsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SponsorPatentMentionsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SponsorPatentMentionsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SponsorPatentMentionsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SponsorPatentMentionsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAssignee

`func (o *SponsorPatentMentionsRow) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *SponsorPatentMentionsRow) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *SponsorPatentMentionsRow) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *SponsorPatentMentionsRow) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *SponsorPatentMentionsRow) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *SponsorPatentMentionsRow) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


