# InstitutionalHoldingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoldingId** | Pointer to **NullableString** | The holding line id. | [optional] 
**IssuerName** | Pointer to **NullableString** | The security issuer as filed. | [optional] 
**IssuerOrgId** | Pointer to **NullableString** | Legal-entity level: the issuer&#39;s organization record (null when unresolved). | [optional] 
**IssuerCompanyId** | Pointer to **NullableString** | Company level: the issuer&#39;s company (null when unresolved). | [optional] 
**IssuerCompanyName** | Pointer to **NullableString** | Display name of the issuer&#39;s company. | [optional] 
**IssuerGroupId** | Pointer to **NullableString** | Group level: the issuer&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**IssuerGroupName** | Pointer to **NullableString** | Display name of the issuer&#39;s group. | [optional] 
**SecurityRef** | Pointer to **NullableString** | The security&#39;s CUSIP (Committee on Uniform Securities Identification Procedures) identifier. | [optional] 
**HolderName** | Pointer to **NullableString** | The institutional manager&#39;s name. | [optional] 
**HolderIdentityKey** | Pointer to **NullableString** | The institutional manager&#39;s stable identity (an SEC registration number), the key that groups all of one manager&#39;s holdings. | [optional] 
**HolderKind** | Pointer to **NullableString** | Manager kind. | [optional] 
**HolderOrgId** | Pointer to **NullableString** | Legal-entity level: the manager&#39;s organization record (null when unresolved). | [optional] 
**HolderCompanyId** | Pointer to **NullableString** | Company level: the manager&#39;s company (null when unresolved). | [optional] 
**HolderCompanyName** | Pointer to **NullableString** | Display name of the manager&#39;s company. | [optional] 
**HolderGroupId** | Pointer to **NullableString** | Group level: the manager&#39;s current ultimate parent. | [optional] 
**HolderGroupName** | Pointer to **NullableString** | Display name of the manager&#39;s group. | [optional] 
**ReportedValueAmount** | Pointer to **NullableFloat32** | Reported value of the position. | [optional] 
**ReportedValueCurrency** | Pointer to **NullableString** | Currency of the reported value. | [optional] 
**Shares** | Pointer to **NullableFloat32** | Shares standing. | [optional] 
**PercentBasis** | Pointer to **NullableString** | The basis the percent is measured on. | [optional] 
**ReportedPercent** | Pointer to **NullableFloat32** | Reported share of the issuer. | [optional] 
**EdgeState** | Pointer to **NullableString** | State of the position (standing, ceased, no longer reported). | [optional] 
**FirstObservedDate** | Pointer to **NullableString** | First filing date the position was observed. | [optional] 
**LastFilingDate** | Pointer to **NullableString** | Latest filing date the position was reported. | [optional] 
**ContainsPersonalData** | Pointer to **NullableBool** | Whether the holder is a natural person. | [optional] 

## Methods

### NewInstitutionalHoldingsRow

`func NewInstitutionalHoldingsRow() *InstitutionalHoldingsRow`

NewInstitutionalHoldingsRow instantiates a new InstitutionalHoldingsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionalHoldingsRowWithDefaults

`func NewInstitutionalHoldingsRowWithDefaults() *InstitutionalHoldingsRow`

NewInstitutionalHoldingsRowWithDefaults instantiates a new InstitutionalHoldingsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoldingId

`func (o *InstitutionalHoldingsRow) GetHoldingId() string`

GetHoldingId returns the HoldingId field if non-nil, zero value otherwise.

### GetHoldingIdOk

`func (o *InstitutionalHoldingsRow) GetHoldingIdOk() (*string, bool)`

GetHoldingIdOk returns a tuple with the HoldingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingId

`func (o *InstitutionalHoldingsRow) SetHoldingId(v string)`

SetHoldingId sets HoldingId field to given value.

### HasHoldingId

`func (o *InstitutionalHoldingsRow) HasHoldingId() bool`

HasHoldingId returns a boolean if a field has been set.

### SetHoldingIdNil

`func (o *InstitutionalHoldingsRow) SetHoldingIdNil(b bool)`

 SetHoldingIdNil sets the value for HoldingId to be an explicit nil

### UnsetHoldingId
`func (o *InstitutionalHoldingsRow) UnsetHoldingId()`

UnsetHoldingId ensures that no value is present for HoldingId, not even an explicit nil
### GetIssuerName

`func (o *InstitutionalHoldingsRow) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *InstitutionalHoldingsRow) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *InstitutionalHoldingsRow) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *InstitutionalHoldingsRow) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### SetIssuerNameNil

`func (o *InstitutionalHoldingsRow) SetIssuerNameNil(b bool)`

 SetIssuerNameNil sets the value for IssuerName to be an explicit nil

### UnsetIssuerName
`func (o *InstitutionalHoldingsRow) UnsetIssuerName()`

UnsetIssuerName ensures that no value is present for IssuerName, not even an explicit nil
### GetIssuerOrgId

`func (o *InstitutionalHoldingsRow) GetIssuerOrgId() string`

GetIssuerOrgId returns the IssuerOrgId field if non-nil, zero value otherwise.

### GetIssuerOrgIdOk

`func (o *InstitutionalHoldingsRow) GetIssuerOrgIdOk() (*string, bool)`

GetIssuerOrgIdOk returns a tuple with the IssuerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerOrgId

`func (o *InstitutionalHoldingsRow) SetIssuerOrgId(v string)`

SetIssuerOrgId sets IssuerOrgId field to given value.

### HasIssuerOrgId

`func (o *InstitutionalHoldingsRow) HasIssuerOrgId() bool`

HasIssuerOrgId returns a boolean if a field has been set.

### SetIssuerOrgIdNil

`func (o *InstitutionalHoldingsRow) SetIssuerOrgIdNil(b bool)`

 SetIssuerOrgIdNil sets the value for IssuerOrgId to be an explicit nil

### UnsetIssuerOrgId
`func (o *InstitutionalHoldingsRow) UnsetIssuerOrgId()`

UnsetIssuerOrgId ensures that no value is present for IssuerOrgId, not even an explicit nil
### GetIssuerCompanyId

`func (o *InstitutionalHoldingsRow) GetIssuerCompanyId() string`

GetIssuerCompanyId returns the IssuerCompanyId field if non-nil, zero value otherwise.

### GetIssuerCompanyIdOk

`func (o *InstitutionalHoldingsRow) GetIssuerCompanyIdOk() (*string, bool)`

GetIssuerCompanyIdOk returns a tuple with the IssuerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyId

`func (o *InstitutionalHoldingsRow) SetIssuerCompanyId(v string)`

SetIssuerCompanyId sets IssuerCompanyId field to given value.

### HasIssuerCompanyId

`func (o *InstitutionalHoldingsRow) HasIssuerCompanyId() bool`

HasIssuerCompanyId returns a boolean if a field has been set.

### SetIssuerCompanyIdNil

`func (o *InstitutionalHoldingsRow) SetIssuerCompanyIdNil(b bool)`

 SetIssuerCompanyIdNil sets the value for IssuerCompanyId to be an explicit nil

### UnsetIssuerCompanyId
`func (o *InstitutionalHoldingsRow) UnsetIssuerCompanyId()`

UnsetIssuerCompanyId ensures that no value is present for IssuerCompanyId, not even an explicit nil
### GetIssuerCompanyName

`func (o *InstitutionalHoldingsRow) GetIssuerCompanyName() string`

GetIssuerCompanyName returns the IssuerCompanyName field if non-nil, zero value otherwise.

### GetIssuerCompanyNameOk

`func (o *InstitutionalHoldingsRow) GetIssuerCompanyNameOk() (*string, bool)`

GetIssuerCompanyNameOk returns a tuple with the IssuerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCompanyName

`func (o *InstitutionalHoldingsRow) SetIssuerCompanyName(v string)`

SetIssuerCompanyName sets IssuerCompanyName field to given value.

### HasIssuerCompanyName

`func (o *InstitutionalHoldingsRow) HasIssuerCompanyName() bool`

HasIssuerCompanyName returns a boolean if a field has been set.

### SetIssuerCompanyNameNil

`func (o *InstitutionalHoldingsRow) SetIssuerCompanyNameNil(b bool)`

 SetIssuerCompanyNameNil sets the value for IssuerCompanyName to be an explicit nil

### UnsetIssuerCompanyName
`func (o *InstitutionalHoldingsRow) UnsetIssuerCompanyName()`

UnsetIssuerCompanyName ensures that no value is present for IssuerCompanyName, not even an explicit nil
### GetIssuerGroupId

`func (o *InstitutionalHoldingsRow) GetIssuerGroupId() string`

GetIssuerGroupId returns the IssuerGroupId field if non-nil, zero value otherwise.

### GetIssuerGroupIdOk

`func (o *InstitutionalHoldingsRow) GetIssuerGroupIdOk() (*string, bool)`

GetIssuerGroupIdOk returns a tuple with the IssuerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupId

`func (o *InstitutionalHoldingsRow) SetIssuerGroupId(v string)`

SetIssuerGroupId sets IssuerGroupId field to given value.

### HasIssuerGroupId

`func (o *InstitutionalHoldingsRow) HasIssuerGroupId() bool`

HasIssuerGroupId returns a boolean if a field has been set.

### SetIssuerGroupIdNil

`func (o *InstitutionalHoldingsRow) SetIssuerGroupIdNil(b bool)`

 SetIssuerGroupIdNil sets the value for IssuerGroupId to be an explicit nil

### UnsetIssuerGroupId
`func (o *InstitutionalHoldingsRow) UnsetIssuerGroupId()`

UnsetIssuerGroupId ensures that no value is present for IssuerGroupId, not even an explicit nil
### GetIssuerGroupName

`func (o *InstitutionalHoldingsRow) GetIssuerGroupName() string`

GetIssuerGroupName returns the IssuerGroupName field if non-nil, zero value otherwise.

### GetIssuerGroupNameOk

`func (o *InstitutionalHoldingsRow) GetIssuerGroupNameOk() (*string, bool)`

GetIssuerGroupNameOk returns a tuple with the IssuerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerGroupName

`func (o *InstitutionalHoldingsRow) SetIssuerGroupName(v string)`

SetIssuerGroupName sets IssuerGroupName field to given value.

### HasIssuerGroupName

`func (o *InstitutionalHoldingsRow) HasIssuerGroupName() bool`

HasIssuerGroupName returns a boolean if a field has been set.

### SetIssuerGroupNameNil

`func (o *InstitutionalHoldingsRow) SetIssuerGroupNameNil(b bool)`

 SetIssuerGroupNameNil sets the value for IssuerGroupName to be an explicit nil

### UnsetIssuerGroupName
`func (o *InstitutionalHoldingsRow) UnsetIssuerGroupName()`

UnsetIssuerGroupName ensures that no value is present for IssuerGroupName, not even an explicit nil
### GetSecurityRef

`func (o *InstitutionalHoldingsRow) GetSecurityRef() string`

GetSecurityRef returns the SecurityRef field if non-nil, zero value otherwise.

### GetSecurityRefOk

`func (o *InstitutionalHoldingsRow) GetSecurityRefOk() (*string, bool)`

GetSecurityRefOk returns a tuple with the SecurityRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRef

`func (o *InstitutionalHoldingsRow) SetSecurityRef(v string)`

SetSecurityRef sets SecurityRef field to given value.

### HasSecurityRef

`func (o *InstitutionalHoldingsRow) HasSecurityRef() bool`

HasSecurityRef returns a boolean if a field has been set.

### SetSecurityRefNil

`func (o *InstitutionalHoldingsRow) SetSecurityRefNil(b bool)`

 SetSecurityRefNil sets the value for SecurityRef to be an explicit nil

### UnsetSecurityRef
`func (o *InstitutionalHoldingsRow) UnsetSecurityRef()`

UnsetSecurityRef ensures that no value is present for SecurityRef, not even an explicit nil
### GetHolderName

`func (o *InstitutionalHoldingsRow) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *InstitutionalHoldingsRow) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *InstitutionalHoldingsRow) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *InstitutionalHoldingsRow) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### SetHolderNameNil

`func (o *InstitutionalHoldingsRow) SetHolderNameNil(b bool)`

 SetHolderNameNil sets the value for HolderName to be an explicit nil

### UnsetHolderName
`func (o *InstitutionalHoldingsRow) UnsetHolderName()`

UnsetHolderName ensures that no value is present for HolderName, not even an explicit nil
### GetHolderIdentityKey

`func (o *InstitutionalHoldingsRow) GetHolderIdentityKey() string`

GetHolderIdentityKey returns the HolderIdentityKey field if non-nil, zero value otherwise.

### GetHolderIdentityKeyOk

`func (o *InstitutionalHoldingsRow) GetHolderIdentityKeyOk() (*string, bool)`

GetHolderIdentityKeyOk returns a tuple with the HolderIdentityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderIdentityKey

`func (o *InstitutionalHoldingsRow) SetHolderIdentityKey(v string)`

SetHolderIdentityKey sets HolderIdentityKey field to given value.

### HasHolderIdentityKey

`func (o *InstitutionalHoldingsRow) HasHolderIdentityKey() bool`

HasHolderIdentityKey returns a boolean if a field has been set.

### SetHolderIdentityKeyNil

`func (o *InstitutionalHoldingsRow) SetHolderIdentityKeyNil(b bool)`

 SetHolderIdentityKeyNil sets the value for HolderIdentityKey to be an explicit nil

### UnsetHolderIdentityKey
`func (o *InstitutionalHoldingsRow) UnsetHolderIdentityKey()`

UnsetHolderIdentityKey ensures that no value is present for HolderIdentityKey, not even an explicit nil
### GetHolderKind

`func (o *InstitutionalHoldingsRow) GetHolderKind() string`

GetHolderKind returns the HolderKind field if non-nil, zero value otherwise.

### GetHolderKindOk

`func (o *InstitutionalHoldingsRow) GetHolderKindOk() (*string, bool)`

GetHolderKindOk returns a tuple with the HolderKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderKind

`func (o *InstitutionalHoldingsRow) SetHolderKind(v string)`

SetHolderKind sets HolderKind field to given value.

### HasHolderKind

`func (o *InstitutionalHoldingsRow) HasHolderKind() bool`

HasHolderKind returns a boolean if a field has been set.

### SetHolderKindNil

`func (o *InstitutionalHoldingsRow) SetHolderKindNil(b bool)`

 SetHolderKindNil sets the value for HolderKind to be an explicit nil

### UnsetHolderKind
`func (o *InstitutionalHoldingsRow) UnsetHolderKind()`

UnsetHolderKind ensures that no value is present for HolderKind, not even an explicit nil
### GetHolderOrgId

`func (o *InstitutionalHoldingsRow) GetHolderOrgId() string`

GetHolderOrgId returns the HolderOrgId field if non-nil, zero value otherwise.

### GetHolderOrgIdOk

`func (o *InstitutionalHoldingsRow) GetHolderOrgIdOk() (*string, bool)`

GetHolderOrgIdOk returns a tuple with the HolderOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderOrgId

`func (o *InstitutionalHoldingsRow) SetHolderOrgId(v string)`

SetHolderOrgId sets HolderOrgId field to given value.

### HasHolderOrgId

`func (o *InstitutionalHoldingsRow) HasHolderOrgId() bool`

HasHolderOrgId returns a boolean if a field has been set.

### SetHolderOrgIdNil

`func (o *InstitutionalHoldingsRow) SetHolderOrgIdNil(b bool)`

 SetHolderOrgIdNil sets the value for HolderOrgId to be an explicit nil

### UnsetHolderOrgId
`func (o *InstitutionalHoldingsRow) UnsetHolderOrgId()`

UnsetHolderOrgId ensures that no value is present for HolderOrgId, not even an explicit nil
### GetHolderCompanyId

`func (o *InstitutionalHoldingsRow) GetHolderCompanyId() string`

GetHolderCompanyId returns the HolderCompanyId field if non-nil, zero value otherwise.

### GetHolderCompanyIdOk

`func (o *InstitutionalHoldingsRow) GetHolderCompanyIdOk() (*string, bool)`

GetHolderCompanyIdOk returns a tuple with the HolderCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderCompanyId

`func (o *InstitutionalHoldingsRow) SetHolderCompanyId(v string)`

SetHolderCompanyId sets HolderCompanyId field to given value.

### HasHolderCompanyId

`func (o *InstitutionalHoldingsRow) HasHolderCompanyId() bool`

HasHolderCompanyId returns a boolean if a field has been set.

### SetHolderCompanyIdNil

`func (o *InstitutionalHoldingsRow) SetHolderCompanyIdNil(b bool)`

 SetHolderCompanyIdNil sets the value for HolderCompanyId to be an explicit nil

### UnsetHolderCompanyId
`func (o *InstitutionalHoldingsRow) UnsetHolderCompanyId()`

UnsetHolderCompanyId ensures that no value is present for HolderCompanyId, not even an explicit nil
### GetHolderCompanyName

`func (o *InstitutionalHoldingsRow) GetHolderCompanyName() string`

GetHolderCompanyName returns the HolderCompanyName field if non-nil, zero value otherwise.

### GetHolderCompanyNameOk

`func (o *InstitutionalHoldingsRow) GetHolderCompanyNameOk() (*string, bool)`

GetHolderCompanyNameOk returns a tuple with the HolderCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderCompanyName

`func (o *InstitutionalHoldingsRow) SetHolderCompanyName(v string)`

SetHolderCompanyName sets HolderCompanyName field to given value.

### HasHolderCompanyName

`func (o *InstitutionalHoldingsRow) HasHolderCompanyName() bool`

HasHolderCompanyName returns a boolean if a field has been set.

### SetHolderCompanyNameNil

`func (o *InstitutionalHoldingsRow) SetHolderCompanyNameNil(b bool)`

 SetHolderCompanyNameNil sets the value for HolderCompanyName to be an explicit nil

### UnsetHolderCompanyName
`func (o *InstitutionalHoldingsRow) UnsetHolderCompanyName()`

UnsetHolderCompanyName ensures that no value is present for HolderCompanyName, not even an explicit nil
### GetHolderGroupId

`func (o *InstitutionalHoldingsRow) GetHolderGroupId() string`

GetHolderGroupId returns the HolderGroupId field if non-nil, zero value otherwise.

### GetHolderGroupIdOk

`func (o *InstitutionalHoldingsRow) GetHolderGroupIdOk() (*string, bool)`

GetHolderGroupIdOk returns a tuple with the HolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderGroupId

`func (o *InstitutionalHoldingsRow) SetHolderGroupId(v string)`

SetHolderGroupId sets HolderGroupId field to given value.

### HasHolderGroupId

`func (o *InstitutionalHoldingsRow) HasHolderGroupId() bool`

HasHolderGroupId returns a boolean if a field has been set.

### SetHolderGroupIdNil

`func (o *InstitutionalHoldingsRow) SetHolderGroupIdNil(b bool)`

 SetHolderGroupIdNil sets the value for HolderGroupId to be an explicit nil

### UnsetHolderGroupId
`func (o *InstitutionalHoldingsRow) UnsetHolderGroupId()`

UnsetHolderGroupId ensures that no value is present for HolderGroupId, not even an explicit nil
### GetHolderGroupName

`func (o *InstitutionalHoldingsRow) GetHolderGroupName() string`

GetHolderGroupName returns the HolderGroupName field if non-nil, zero value otherwise.

### GetHolderGroupNameOk

`func (o *InstitutionalHoldingsRow) GetHolderGroupNameOk() (*string, bool)`

GetHolderGroupNameOk returns a tuple with the HolderGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderGroupName

`func (o *InstitutionalHoldingsRow) SetHolderGroupName(v string)`

SetHolderGroupName sets HolderGroupName field to given value.

### HasHolderGroupName

`func (o *InstitutionalHoldingsRow) HasHolderGroupName() bool`

HasHolderGroupName returns a boolean if a field has been set.

### SetHolderGroupNameNil

`func (o *InstitutionalHoldingsRow) SetHolderGroupNameNil(b bool)`

 SetHolderGroupNameNil sets the value for HolderGroupName to be an explicit nil

### UnsetHolderGroupName
`func (o *InstitutionalHoldingsRow) UnsetHolderGroupName()`

UnsetHolderGroupName ensures that no value is present for HolderGroupName, not even an explicit nil
### GetReportedValueAmount

`func (o *InstitutionalHoldingsRow) GetReportedValueAmount() float32`

GetReportedValueAmount returns the ReportedValueAmount field if non-nil, zero value otherwise.

### GetReportedValueAmountOk

`func (o *InstitutionalHoldingsRow) GetReportedValueAmountOk() (*float32, bool)`

GetReportedValueAmountOk returns a tuple with the ReportedValueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedValueAmount

`func (o *InstitutionalHoldingsRow) SetReportedValueAmount(v float32)`

SetReportedValueAmount sets ReportedValueAmount field to given value.

### HasReportedValueAmount

`func (o *InstitutionalHoldingsRow) HasReportedValueAmount() bool`

HasReportedValueAmount returns a boolean if a field has been set.

### SetReportedValueAmountNil

`func (o *InstitutionalHoldingsRow) SetReportedValueAmountNil(b bool)`

 SetReportedValueAmountNil sets the value for ReportedValueAmount to be an explicit nil

### UnsetReportedValueAmount
`func (o *InstitutionalHoldingsRow) UnsetReportedValueAmount()`

UnsetReportedValueAmount ensures that no value is present for ReportedValueAmount, not even an explicit nil
### GetReportedValueCurrency

`func (o *InstitutionalHoldingsRow) GetReportedValueCurrency() string`

GetReportedValueCurrency returns the ReportedValueCurrency field if non-nil, zero value otherwise.

### GetReportedValueCurrencyOk

`func (o *InstitutionalHoldingsRow) GetReportedValueCurrencyOk() (*string, bool)`

GetReportedValueCurrencyOk returns a tuple with the ReportedValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedValueCurrency

`func (o *InstitutionalHoldingsRow) SetReportedValueCurrency(v string)`

SetReportedValueCurrency sets ReportedValueCurrency field to given value.

### HasReportedValueCurrency

`func (o *InstitutionalHoldingsRow) HasReportedValueCurrency() bool`

HasReportedValueCurrency returns a boolean if a field has been set.

### SetReportedValueCurrencyNil

`func (o *InstitutionalHoldingsRow) SetReportedValueCurrencyNil(b bool)`

 SetReportedValueCurrencyNil sets the value for ReportedValueCurrency to be an explicit nil

### UnsetReportedValueCurrency
`func (o *InstitutionalHoldingsRow) UnsetReportedValueCurrency()`

UnsetReportedValueCurrency ensures that no value is present for ReportedValueCurrency, not even an explicit nil
### GetShares

`func (o *InstitutionalHoldingsRow) GetShares() float32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *InstitutionalHoldingsRow) GetSharesOk() (*float32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *InstitutionalHoldingsRow) SetShares(v float32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *InstitutionalHoldingsRow) HasShares() bool`

HasShares returns a boolean if a field has been set.

### SetSharesNil

`func (o *InstitutionalHoldingsRow) SetSharesNil(b bool)`

 SetSharesNil sets the value for Shares to be an explicit nil

### UnsetShares
`func (o *InstitutionalHoldingsRow) UnsetShares()`

UnsetShares ensures that no value is present for Shares, not even an explicit nil
### GetPercentBasis

`func (o *InstitutionalHoldingsRow) GetPercentBasis() string`

GetPercentBasis returns the PercentBasis field if non-nil, zero value otherwise.

### GetPercentBasisOk

`func (o *InstitutionalHoldingsRow) GetPercentBasisOk() (*string, bool)`

GetPercentBasisOk returns a tuple with the PercentBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentBasis

`func (o *InstitutionalHoldingsRow) SetPercentBasis(v string)`

SetPercentBasis sets PercentBasis field to given value.

### HasPercentBasis

`func (o *InstitutionalHoldingsRow) HasPercentBasis() bool`

HasPercentBasis returns a boolean if a field has been set.

### SetPercentBasisNil

`func (o *InstitutionalHoldingsRow) SetPercentBasisNil(b bool)`

 SetPercentBasisNil sets the value for PercentBasis to be an explicit nil

### UnsetPercentBasis
`func (o *InstitutionalHoldingsRow) UnsetPercentBasis()`

UnsetPercentBasis ensures that no value is present for PercentBasis, not even an explicit nil
### GetReportedPercent

`func (o *InstitutionalHoldingsRow) GetReportedPercent() float32`

GetReportedPercent returns the ReportedPercent field if non-nil, zero value otherwise.

### GetReportedPercentOk

`func (o *InstitutionalHoldingsRow) GetReportedPercentOk() (*float32, bool)`

GetReportedPercentOk returns a tuple with the ReportedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedPercent

`func (o *InstitutionalHoldingsRow) SetReportedPercent(v float32)`

SetReportedPercent sets ReportedPercent field to given value.

### HasReportedPercent

`func (o *InstitutionalHoldingsRow) HasReportedPercent() bool`

HasReportedPercent returns a boolean if a field has been set.

### SetReportedPercentNil

`func (o *InstitutionalHoldingsRow) SetReportedPercentNil(b bool)`

 SetReportedPercentNil sets the value for ReportedPercent to be an explicit nil

### UnsetReportedPercent
`func (o *InstitutionalHoldingsRow) UnsetReportedPercent()`

UnsetReportedPercent ensures that no value is present for ReportedPercent, not even an explicit nil
### GetEdgeState

`func (o *InstitutionalHoldingsRow) GetEdgeState() string`

GetEdgeState returns the EdgeState field if non-nil, zero value otherwise.

### GetEdgeStateOk

`func (o *InstitutionalHoldingsRow) GetEdgeStateOk() (*string, bool)`

GetEdgeStateOk returns a tuple with the EdgeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeState

`func (o *InstitutionalHoldingsRow) SetEdgeState(v string)`

SetEdgeState sets EdgeState field to given value.

### HasEdgeState

`func (o *InstitutionalHoldingsRow) HasEdgeState() bool`

HasEdgeState returns a boolean if a field has been set.

### SetEdgeStateNil

`func (o *InstitutionalHoldingsRow) SetEdgeStateNil(b bool)`

 SetEdgeStateNil sets the value for EdgeState to be an explicit nil

### UnsetEdgeState
`func (o *InstitutionalHoldingsRow) UnsetEdgeState()`

UnsetEdgeState ensures that no value is present for EdgeState, not even an explicit nil
### GetFirstObservedDate

`func (o *InstitutionalHoldingsRow) GetFirstObservedDate() string`

GetFirstObservedDate returns the FirstObservedDate field if non-nil, zero value otherwise.

### GetFirstObservedDateOk

`func (o *InstitutionalHoldingsRow) GetFirstObservedDateOk() (*string, bool)`

GetFirstObservedDateOk returns a tuple with the FirstObservedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstObservedDate

`func (o *InstitutionalHoldingsRow) SetFirstObservedDate(v string)`

SetFirstObservedDate sets FirstObservedDate field to given value.

### HasFirstObservedDate

`func (o *InstitutionalHoldingsRow) HasFirstObservedDate() bool`

HasFirstObservedDate returns a boolean if a field has been set.

### SetFirstObservedDateNil

`func (o *InstitutionalHoldingsRow) SetFirstObservedDateNil(b bool)`

 SetFirstObservedDateNil sets the value for FirstObservedDate to be an explicit nil

### UnsetFirstObservedDate
`func (o *InstitutionalHoldingsRow) UnsetFirstObservedDate()`

UnsetFirstObservedDate ensures that no value is present for FirstObservedDate, not even an explicit nil
### GetLastFilingDate

`func (o *InstitutionalHoldingsRow) GetLastFilingDate() string`

GetLastFilingDate returns the LastFilingDate field if non-nil, zero value otherwise.

### GetLastFilingDateOk

`func (o *InstitutionalHoldingsRow) GetLastFilingDateOk() (*string, bool)`

GetLastFilingDateOk returns a tuple with the LastFilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFilingDate

`func (o *InstitutionalHoldingsRow) SetLastFilingDate(v string)`

SetLastFilingDate sets LastFilingDate field to given value.

### HasLastFilingDate

`func (o *InstitutionalHoldingsRow) HasLastFilingDate() bool`

HasLastFilingDate returns a boolean if a field has been set.

### SetLastFilingDateNil

`func (o *InstitutionalHoldingsRow) SetLastFilingDateNil(b bool)`

 SetLastFilingDateNil sets the value for LastFilingDate to be an explicit nil

### UnsetLastFilingDate
`func (o *InstitutionalHoldingsRow) UnsetLastFilingDate()`

UnsetLastFilingDate ensures that no value is present for LastFilingDate, not even an explicit nil
### GetContainsPersonalData

`func (o *InstitutionalHoldingsRow) GetContainsPersonalData() bool`

GetContainsPersonalData returns the ContainsPersonalData field if non-nil, zero value otherwise.

### GetContainsPersonalDataOk

`func (o *InstitutionalHoldingsRow) GetContainsPersonalDataOk() (*bool, bool)`

GetContainsPersonalDataOk returns a tuple with the ContainsPersonalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsPersonalData

`func (o *InstitutionalHoldingsRow) SetContainsPersonalData(v bool)`

SetContainsPersonalData sets ContainsPersonalData field to given value.

### HasContainsPersonalData

`func (o *InstitutionalHoldingsRow) HasContainsPersonalData() bool`

HasContainsPersonalData returns a boolean if a field has been set.

### SetContainsPersonalDataNil

`func (o *InstitutionalHoldingsRow) SetContainsPersonalDataNil(b bool)`

 SetContainsPersonalDataNil sets the value for ContainsPersonalData to be an explicit nil

### UnsetContainsPersonalData
`func (o *InstitutionalHoldingsRow) UnsetContainsPersonalData()`

UnsetContainsPersonalData ensures that no value is present for ContainsPersonalData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


