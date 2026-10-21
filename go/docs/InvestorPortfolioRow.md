# InvestorPortfolioRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestorId** | Pointer to **NullableString** | The investor key. | [optional] 
**IssuerCik** | Pointer to **NullableString** | The issuer&#39;s SEC Central Index Key (CIK), ten digits. | [optional] 
**IssuerName** | Pointer to **NullableString** | The issuer&#39;s name: its organization record&#39;s name, else the SEC company name. | [optional] 
**IssuerOrgId** | Pointer to **NullableString** | Legal-entity level: the issuer&#39;s organization record (null when unresolved). | [optional] 
**IssuerCompanyId** | Pointer to **NullableString** | Company level: the issuer&#39;s company (null when unresolved). | [optional] 
**IssuerCompanyName** | Pointer to **NullableString** | Display name of the issuer&#39;s company. | [optional] 
**IssuerGroupId** | Pointer to **NullableString** | Group level: the issuer&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**IssuerGroupName** | Pointer to **NullableString** | Display name of the issuer&#39;s group. | [optional] 
**RelationshipTypes** | Pointer to **[]string** | How the investor is named against the issuer: holder_5pct (a 5 percent holder) or selling_stockholder. | [optional] 
**EdgeCount** | Pointer to **NullableInt64** | Filing mentions behind the row, across every filing and relationship type. | [optional] 
**FirstObservedDate** | Pointer to **NullableString** | The earliest filing date the investor was named against the issuer. | [optional] 
**LastObservedDate** | Pointer to **NullableString** | The latest filing date the investor was named against the issuer. | [optional] 
**IsLifeSciencesIssuer** | Pointer to **NullableBool** | True when the issuer&#39;s SIC (Standard Industrial Classification) code is a life-sciences code, false when it is another code, null when the issuer has no SIC evidence. | [optional] 

## Methods

### NewInvestorPortfolioRow

`func NewInvestorPortfolioRow() *InvestorPortfolioRow`

NewInvestorPortfolioRow instantiates a new InvestorPortfolioRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestorPortfolioRowWithDefaults

`func NewInvestorPortfolioRowWithDefaults() *InvestorPortfolioRow`

NewInvestorPortfolioRowWithDefaults instantiates a new InvestorPortfolioRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestorId

`func (o *InvestorPortfolioRow) GetInvestorId() string`

GetInvestorId returns the InvestorId field if non-nil, zero value otherwise.

### GetInvestorIdOk

`func (o *InvestorPortfolioRow) GetInvestorIdOk() (*string, bool)`

GetInvestorIdOk returns a tuple with the InvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorId

`func (o *InvestorPortfolioRow) SetInvestorId(v string)`

SetInvestorId sets InvestorId field to given value.

### HasInvestorId

`func (o *InvestorPortfolioRow) HasInvestorId() bool`

HasInvestorId returns a boolean if a field has been set.

### SetInvestorIdNil

`func (o *InvestorPortfolioRow) SetInvestorIdNil(b bool)`

 SetInvestorIdNil sets the value for InvestorId to be an explicit nil

### UnsetInvestorId
`func (o *InvestorPortfolioRow) UnsetInvestorId()`

UnsetInvestorId ensures that no value is present for InvestorId, not even an explicit nil
### GetIssuerCik

`func (o *InvestorPortfolioRow) GetIssuerCik() string`

GetIssuerCik returns the IssuerCik field if non-nil, zero value otherwise.

### GetIssuerCikOk

`func (o *InvestorPortfolioRow) GetIssuerCikOk() (*string, bool)`

GetIssuerCikOk returns a tuple with the IssuerCik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCik

`func (o *InvestorPortfolioRow) SetIssuerCik(v string)`

SetIssuerCik sets IssuerCik field to given value.

### HasIssuerCik

`func (o *InvestorPortfolioRow) HasIssuerCik() bool`

HasIssuerCik returns a boolean if a field has been set.

### SetIssuerCikNil

`func (o *InvestorPortfolioRow) SetIssuerCikNil(b bool)`

 SetIssuerCikNil sets the value for IssuerCik to be an explicit nil

### UnsetIssuerCik
`func (o *InvestorPortfolioRow) UnsetIssuerCik()`

UnsetIssuerCik ensures that no value is present for IssuerCik, not even an explicit nil
### GetIssuerName

`func (o *InvestorPortfolioRow) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *InvestorPortfolioRow) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *InvestorPortfolioRow) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *InvestorPortfolioRow) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### SetIssuerNameNil

`func (o *InvestorPortfolioRow) SetIssuerNameNil(b bool)`

 SetIssuerNameNil sets the value for IssuerName to be an explicit nil

### UnsetIssuerName
`func (o *InvestorPortfolioRow) UnsetIssuerName()`

UnsetIssuerName ensures that no value is present for IssuerName, not even an explicit nil
### GetIssuerOrgId

`func (o *InvestorPortfolioRow) GetIssuerOrgId() string`

GetIssuerOrgId returns the IssuerOrgId field if non-nil, zero value otherwise.

### GetIssuerOrgIdOk

`func (o *InvestorPortfolioRow) GetIssuerOrgIdOk() (*string, bool)`

GetIssuerOrgIdOk returns a tuple with the IssuerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerOrgId

`func (o *InvestorPortfolioRow) SetIssuerOrgId(v string)`

SetIssuerOrgId sets IssuerOrgId field to given value.

### HasIssuerOrgId

`func (o *InvestorPortfolioRow) HasIssuerOrgId() bool`

HasIssuerOrgId returns a boolean if a field has been set.

### SetIssuerOrgIdNil

`func (o *InvestorPortfolioRow) SetIssuerOrgIdNil(b bool)`

 SetIssuerOrgIdNil sets the value for IssuerOrgId to be an explicit nil

### UnsetIssuerOrgId
`func (o *InvestorPortfolioRow) UnsetIssuerOrgId()`

UnsetIssuerOrgId ensures that no value is present for IssuerOrgId, not even an explicit nil
### GetIssuerCompanyId

`func (o *InvestorPortfolioRow) GetIssuerCompanyId() string`

GetIssuerCompanyId returns the IssuerCompanyId field if non-nil, zero value otherwise.

### GetIssuerCompanyIdOk

`func (o *InvestorPortfolioRow) GetIssuerCompanyIdOk() (*string, bool)`

GetIssuerCompanyIdOk returns a tuple with the IssuerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyId

`func (o *InvestorPortfolioRow) SetIssuerCompanyId(v string)`

SetIssuerCompanyId sets IssuerCompanyId field to given value.

### HasIssuerCompanyId

`func (o *InvestorPortfolioRow) HasIssuerCompanyId() bool`

HasIssuerCompanyId returns a boolean if a field has been set.

### SetIssuerCompanyIdNil

`func (o *InvestorPortfolioRow) SetIssuerCompanyIdNil(b bool)`

 SetIssuerCompanyIdNil sets the value for IssuerCompanyId to be an explicit nil

### UnsetIssuerCompanyId
`func (o *InvestorPortfolioRow) UnsetIssuerCompanyId()`

UnsetIssuerCompanyId ensures that no value is present for IssuerCompanyId, not even an explicit nil
### GetIssuerCompanyName

`func (o *InvestorPortfolioRow) GetIssuerCompanyName() string`

GetIssuerCompanyName returns the IssuerCompanyName field if non-nil, zero value otherwise.

### GetIssuerCompanyNameOk

`func (o *InvestorPortfolioRow) GetIssuerCompanyNameOk() (*string, bool)`

GetIssuerCompanyNameOk returns a tuple with the IssuerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyName

`func (o *InvestorPortfolioRow) SetIssuerCompanyName(v string)`

SetIssuerCompanyName sets IssuerCompanyName field to given value.

### HasIssuerCompanyName

`func (o *InvestorPortfolioRow) HasIssuerCompanyName() bool`

HasIssuerCompanyName returns a boolean if a field has been set.

### SetIssuerCompanyNameNil

`func (o *InvestorPortfolioRow) SetIssuerCompanyNameNil(b bool)`

 SetIssuerCompanyNameNil sets the value for IssuerCompanyName to be an explicit nil

### UnsetIssuerCompanyName
`func (o *InvestorPortfolioRow) UnsetIssuerCompanyName()`

UnsetIssuerCompanyName ensures that no value is present for IssuerCompanyName, not even an explicit nil
### GetIssuerGroupId

`func (o *InvestorPortfolioRow) GetIssuerGroupId() string`

GetIssuerGroupId returns the IssuerGroupId field if non-nil, zero value otherwise.

### GetIssuerGroupIdOk

`func (o *InvestorPortfolioRow) GetIssuerGroupIdOk() (*string, bool)`

GetIssuerGroupIdOk returns a tuple with the IssuerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupId

`func (o *InvestorPortfolioRow) SetIssuerGroupId(v string)`

SetIssuerGroupId sets IssuerGroupId field to given value.

### HasIssuerGroupId

`func (o *InvestorPortfolioRow) HasIssuerGroupId() bool`

HasIssuerGroupId returns a boolean if a field has been set.

### SetIssuerGroupIdNil

`func (o *InvestorPortfolioRow) SetIssuerGroupIdNil(b bool)`

 SetIssuerGroupIdNil sets the value for IssuerGroupId to be an explicit nil

### UnsetIssuerGroupId
`func (o *InvestorPortfolioRow) UnsetIssuerGroupId()`

UnsetIssuerGroupId ensures that no value is present for IssuerGroupId, not even an explicit nil
### GetIssuerGroupName

`func (o *InvestorPortfolioRow) GetIssuerGroupName() string`

GetIssuerGroupName returns the IssuerGroupName field if non-nil, zero value otherwise.

### GetIssuerGroupNameOk

`func (o *InvestorPortfolioRow) GetIssuerGroupNameOk() (*string, bool)`

GetIssuerGroupNameOk returns a tuple with the IssuerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupName

`func (o *InvestorPortfolioRow) SetIssuerGroupName(v string)`

SetIssuerGroupName sets IssuerGroupName field to given value.

### HasIssuerGroupName

`func (o *InvestorPortfolioRow) HasIssuerGroupName() bool`

HasIssuerGroupName returns a boolean if a field has been set.

### SetIssuerGroupNameNil

`func (o *InvestorPortfolioRow) SetIssuerGroupNameNil(b bool)`

 SetIssuerGroupNameNil sets the value for IssuerGroupName to be an explicit nil

### UnsetIssuerGroupName
`func (o *InvestorPortfolioRow) UnsetIssuerGroupName()`

UnsetIssuerGroupName ensures that no value is present for IssuerGroupName, not even an explicit nil
### GetRelationshipTypes

`func (o *InvestorPortfolioRow) GetRelationshipTypes() []string`

GetRelationshipTypes returns the RelationshipTypes field if non-nil, zero value otherwise.

### GetRelationshipTypesOk

`func (o *InvestorPortfolioRow) GetRelationshipTypesOk() (*[]string, bool)`

GetRelationshipTypesOk returns a tuple with the RelationshipTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipTypes

`func (o *InvestorPortfolioRow) SetRelationshipTypes(v []string)`

SetRelationshipTypes sets RelationshipTypes field to given value.

### HasRelationshipTypes

`func (o *InvestorPortfolioRow) HasRelationshipTypes() bool`

HasRelationshipTypes returns a boolean if a field has been set.

### SetRelationshipTypesNil

`func (o *InvestorPortfolioRow) SetRelationshipTypesNil(b bool)`

 SetRelationshipTypesNil sets the value for RelationshipTypes to be an explicit nil

### UnsetRelationshipTypes
`func (o *InvestorPortfolioRow) UnsetRelationshipTypes()`

UnsetRelationshipTypes ensures that no value is present for RelationshipTypes, not even an explicit nil
### GetEdgeCount

`func (o *InvestorPortfolioRow) GetEdgeCount() int64`

GetEdgeCount returns the EdgeCount field if non-nil, zero value otherwise.

### GetEdgeCountOk

`func (o *InvestorPortfolioRow) GetEdgeCountOk() (*int64, bool)`

GetEdgeCountOk returns a tuple with the EdgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCount

`func (o *InvestorPortfolioRow) SetEdgeCount(v int64)`

SetEdgeCount sets EdgeCount field to given value.

### HasEdgeCount

`func (o *InvestorPortfolioRow) HasEdgeCount() bool`

HasEdgeCount returns a boolean if a field has been set.

### SetEdgeCountNil

`func (o *InvestorPortfolioRow) SetEdgeCountNil(b bool)`

 SetEdgeCountNil sets the value for EdgeCount to be an explicit nil

### UnsetEdgeCount
`func (o *InvestorPortfolioRow) UnsetEdgeCount()`

UnsetEdgeCount ensures that no value is present for EdgeCount, not even an explicit nil
### GetFirstObservedDate

`func (o *InvestorPortfolioRow) GetFirstObservedDate() string`

GetFirstObservedDate returns the FirstObservedDate field if non-nil, zero value otherwise.

### GetFirstObservedDateOk

`func (o *InvestorPortfolioRow) GetFirstObservedDateOk() (*string, bool)`

GetFirstObservedDateOk returns a tuple with the FirstObservedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstObservedDate

`func (o *InvestorPortfolioRow) SetFirstObservedDate(v string)`

SetFirstObservedDate sets FirstObservedDate field to given value.

### HasFirstObservedDate

`func (o *InvestorPortfolioRow) HasFirstObservedDate() bool`

HasFirstObservedDate returns a boolean if a field has been set.

### SetFirstObservedDateNil

`func (o *InvestorPortfolioRow) SetFirstObservedDateNil(b bool)`

 SetFirstObservedDateNil sets the value for FirstObservedDate to be an explicit nil

### UnsetFirstObservedDate
`func (o *InvestorPortfolioRow) UnsetFirstObservedDate()`

UnsetFirstObservedDate ensures that no value is present for FirstObservedDate, not even an explicit nil
### GetLastObservedDate

`func (o *InvestorPortfolioRow) GetLastObservedDate() string`

GetLastObservedDate returns the LastObservedDate field if non-nil, zero value otherwise.

### GetLastObservedDateOk

`func (o *InvestorPortfolioRow) GetLastObservedDateOk() (*string, bool)`

GetLastObservedDateOk returns a tuple with the LastObservedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservedDate

`func (o *InvestorPortfolioRow) SetLastObservedDate(v string)`

SetLastObservedDate sets LastObservedDate field to given value.

### HasLastObservedDate

`func (o *InvestorPortfolioRow) HasLastObservedDate() bool`

HasLastObservedDate returns a boolean if a field has been set.

### SetLastObservedDateNil

`func (o *InvestorPortfolioRow) SetLastObservedDateNil(b bool)`

 SetLastObservedDateNil sets the value for LastObservedDate to be an explicit nil

### UnsetLastObservedDate
`func (o *InvestorPortfolioRow) UnsetLastObservedDate()`

UnsetLastObservedDate ensures that no value is present for LastObservedDate, not even an explicit nil
### GetIsLifeSciencesIssuer

`func (o *InvestorPortfolioRow) GetIsLifeSciencesIssuer() bool`

GetIsLifeSciencesIssuer returns the IsLifeSciencesIssuer field if non-nil, zero value otherwise.

### GetIsLifeSciencesIssuerOk

`func (o *InvestorPortfolioRow) GetIsLifeSciencesIssuerOk() (*bool, bool)`

GetIsLifeSciencesIssuerOk returns a tuple with the IsLifeSciencesIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLifeSciencesIssuer

`func (o *InvestorPortfolioRow) SetIsLifeSciencesIssuer(v bool)`

SetIsLifeSciencesIssuer sets IsLifeSciencesIssuer field to given value.

### HasIsLifeSciencesIssuer

`func (o *InvestorPortfolioRow) HasIsLifeSciencesIssuer() bool`

HasIsLifeSciencesIssuer returns a boolean if a field has been set.

### SetIsLifeSciencesIssuerNil

`func (o *InvestorPortfolioRow) SetIsLifeSciencesIssuerNil(b bool)`

 SetIsLifeSciencesIssuerNil sets the value for IsLifeSciencesIssuer to be an explicit nil

### UnsetIsLifeSciencesIssuer
`func (o *InvestorPortfolioRow) UnsetIsLifeSciencesIssuer()`

UnsetIsLifeSciencesIssuer ensures that no value is present for IsLifeSciencesIssuer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


