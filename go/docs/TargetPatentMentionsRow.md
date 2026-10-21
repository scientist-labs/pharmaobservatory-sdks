# TargetPatentMentionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] 
**DrugId** | Pointer to **NullableString** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**PatentNumber** | Pointer to **NullableString** | The patent number. | [optional] 
**Country** | Pointer to **NullableString** | The patent office&#39;s country code. | [optional] 
**PublicationDate** | Pointer to **NullableString** | The patent&#39;s publication date. | [optional] 
**Title** | Pointer to **NullableString** | The patent&#39;s title. | [optional] 
**Assignee** | Pointer to **NullableString** | The patent&#39;s assignee as filed; null where none. | [optional] 

## Methods

### NewTargetPatentMentionsRow

`func NewTargetPatentMentionsRow() *TargetPatentMentionsRow`

NewTargetPatentMentionsRow instantiates a new TargetPatentMentionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetPatentMentionsRowWithDefaults

`func NewTargetPatentMentionsRowWithDefaults() *TargetPatentMentionsRow`

NewTargetPatentMentionsRowWithDefaults instantiates a new TargetPatentMentionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetPatentMentionsRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetPatentMentionsRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetPatentMentionsRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetPatentMentionsRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetPatentMentionsRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetPatentMentionsRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *TargetPatentMentionsRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *TargetPatentMentionsRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *TargetPatentMentionsRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *TargetPatentMentionsRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *TargetPatentMentionsRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *TargetPatentMentionsRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetDrugId

`func (o *TargetPatentMentionsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *TargetPatentMentionsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *TargetPatentMentionsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *TargetPatentMentionsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *TargetPatentMentionsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *TargetPatentMentionsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *TargetPatentMentionsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *TargetPatentMentionsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *TargetPatentMentionsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *TargetPatentMentionsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *TargetPatentMentionsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *TargetPatentMentionsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetPatentNumber

`func (o *TargetPatentMentionsRow) GetPatentNumber() string`

GetPatentNumber returns the PatentNumber field if non-nil, zero value otherwise.

### GetPatentNumberOk

`func (o *TargetPatentMentionsRow) GetPatentNumberOk() (*string, bool)`

GetPatentNumberOk returns a tuple with the PatentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentNumber

`func (o *TargetPatentMentionsRow) SetPatentNumber(v string)`

SetPatentNumber sets PatentNumber field to given value.

### HasPatentNumber

`func (o *TargetPatentMentionsRow) HasPatentNumber() bool`

HasPatentNumber returns a boolean if a field has been set.

### SetPatentNumberNil

`func (o *TargetPatentMentionsRow) SetPatentNumberNil(b bool)`

 SetPatentNumberNil sets the value for PatentNumber to be an explicit nil

### UnsetPatentNumber
`func (o *TargetPatentMentionsRow) UnsetPatentNumber()`

UnsetPatentNumber ensures that no value is present for PatentNumber, not even an explicit nil
### GetCountry

`func (o *TargetPatentMentionsRow) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TargetPatentMentionsRow) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TargetPatentMentionsRow) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TargetPatentMentionsRow) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *TargetPatentMentionsRow) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *TargetPatentMentionsRow) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPublicationDate

`func (o *TargetPatentMentionsRow) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *TargetPatentMentionsRow) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *TargetPatentMentionsRow) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *TargetPatentMentionsRow) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### SetPublicationDateNil

`func (o *TargetPatentMentionsRow) SetPublicationDateNil(b bool)`

 SetPublicationDateNil sets the value for PublicationDate to be an explicit nil

### UnsetPublicationDate
`func (o *TargetPatentMentionsRow) UnsetPublicationDate()`

UnsetPublicationDate ensures that no value is present for PublicationDate, not even an explicit nil
### GetTitle

`func (o *TargetPatentMentionsRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TargetPatentMentionsRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TargetPatentMentionsRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TargetPatentMentionsRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TargetPatentMentionsRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TargetPatentMentionsRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAssignee

`func (o *TargetPatentMentionsRow) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TargetPatentMentionsRow) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TargetPatentMentionsRow) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *TargetPatentMentionsRow) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *TargetPatentMentionsRow) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *TargetPatentMentionsRow) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


