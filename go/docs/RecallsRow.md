# RecallsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecallNumber** | Pointer to **NullableString** | The FDA recall number. | [optional] 
**Domain** | Pointer to **NullableString** | The regulatory domain of the recalled product (drug or device here). | [optional] 
**Status** | Pointer to **NullableString** | The recall status (Ongoing, Completed or Terminated). | [optional] 
**RecallClass** | Pointer to **NullableString** | The severity class: Class I, II or III. | [optional] 
**ProductType** | Pointer to **NullableString** | The product type the FDA reports (Drugs, Devices and so on). | [optional] 
**FirmName** | Pointer to **NullableString** | The firm&#39;s name as filed with the regulator (the recalling firm). | [optional] 
**FirmOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] 
**FirmCompanyId** | Pointer to **NullableString** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] 
**FirmCompanyName** | Pointer to **NullableString** | Display name of the firm&#39;s company. | [optional] 
**FirmGroupId** | Pointer to **NullableString** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FirmGroupName** | Pointer to **NullableString** | Display name of the firm&#39;s group. | [optional] 
**IsOrgResolved** | Pointer to **NullableBool** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] 
**MatchTier** | Pointer to **NullableString** | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous (several candidates, none chosen) or unresolved. Warning letters, untitled letters and debarments attach by firm name alone. | [optional] 
**OrgResolutionMethod** | Pointer to **NullableString** | The route that attached the firm: fei_bridge (the FDA establishment identifier, FEI, bridge), fei_firm_mint or fei_firm_fold (the operating firm&#39;s registry record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional] 
**ProductDescription** | Pointer to **NullableString** | The recalled product as described in the report. | [optional] 
**ReasonForRecall** | Pointer to **NullableString** | The reason the FDA gives for the recall. | [optional] 
**RecallInitiationDate** | Pointer to **NullableString** | The date the firm began the recall. | [optional] 
**ReportDate** | Pointer to **NullableString** | The date the FDA reported the recall. | [optional] 
**TerminationDate** | Pointer to **NullableString** | The date the recall was terminated, where it was. | [optional] 

## Methods

### NewRecallsRow

`func NewRecallsRow() *RecallsRow`

NewRecallsRow instantiates a new RecallsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecallsRowWithDefaults

`func NewRecallsRowWithDefaults() *RecallsRow`

NewRecallsRowWithDefaults instantiates a new RecallsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecallNumber

`func (o *RecallsRow) GetRecallNumber() string`

GetRecallNumber returns the RecallNumber field if non-nil, zero value otherwise.

### GetRecallNumberOk

`func (o *RecallsRow) GetRecallNumberOk() (*string, bool)`

GetRecallNumberOk returns a tuple with the RecallNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallNumber

`func (o *RecallsRow) SetRecallNumber(v string)`

SetRecallNumber sets RecallNumber field to given value.

### HasRecallNumber

`func (o *RecallsRow) HasRecallNumber() bool`

HasRecallNumber returns a boolean if a field has been set.

### SetRecallNumberNil

`func (o *RecallsRow) SetRecallNumberNil(b bool)`

 SetRecallNumberNil sets the value for RecallNumber to be an explicit nil

### UnsetRecallNumber
`func (o *RecallsRow) UnsetRecallNumber()`

UnsetRecallNumber ensures that no value is present for RecallNumber, not even an explicit nil
### GetDomain

`func (o *RecallsRow) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RecallsRow) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RecallsRow) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RecallsRow) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *RecallsRow) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *RecallsRow) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetStatus

`func (o *RecallsRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecallsRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecallsRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecallsRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RecallsRow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RecallsRow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRecallClass

`func (o *RecallsRow) GetRecallClass() string`

GetRecallClass returns the RecallClass field if non-nil, zero value otherwise.

### GetRecallClassOk

`func (o *RecallsRow) GetRecallClassOk() (*string, bool)`

GetRecallClassOk returns a tuple with the RecallClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallClass

`func (o *RecallsRow) SetRecallClass(v string)`

SetRecallClass sets RecallClass field to given value.

### HasRecallClass

`func (o *RecallsRow) HasRecallClass() bool`

HasRecallClass returns a boolean if a field has been set.

### SetRecallClassNil

`func (o *RecallsRow) SetRecallClassNil(b bool)`

 SetRecallClassNil sets the value for RecallClass to be an explicit nil

### UnsetRecallClass
`func (o *RecallsRow) UnsetRecallClass()`

UnsetRecallClass ensures that no value is present for RecallClass, not even an explicit nil
### GetProductType

`func (o *RecallsRow) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *RecallsRow) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *RecallsRow) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *RecallsRow) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### SetProductTypeNil

`func (o *RecallsRow) SetProductTypeNil(b bool)`

 SetProductTypeNil sets the value for ProductType to be an explicit nil

### UnsetProductType
`func (o *RecallsRow) UnsetProductType()`

UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
### GetFirmName

`func (o *RecallsRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *RecallsRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *RecallsRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *RecallsRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *RecallsRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *RecallsRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *RecallsRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *RecallsRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *RecallsRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *RecallsRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *RecallsRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *RecallsRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *RecallsRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *RecallsRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *RecallsRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *RecallsRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *RecallsRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *RecallsRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *RecallsRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *RecallsRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *RecallsRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *RecallsRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *RecallsRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *RecallsRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *RecallsRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *RecallsRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *RecallsRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *RecallsRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *RecallsRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *RecallsRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *RecallsRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *RecallsRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *RecallsRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *RecallsRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *RecallsRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *RecallsRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetIsOrgResolved

`func (o *RecallsRow) GetIsOrgResolved() bool`

GetIsOrgResolved returns the IsOrgResolved field if non-nil, zero value otherwise.

### GetIsOrgResolvedOk

`func (o *RecallsRow) GetIsOrgResolvedOk() (*bool, bool)`

GetIsOrgResolvedOk returns a tuple with the IsOrgResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgResolved

`func (o *RecallsRow) SetIsOrgResolved(v bool)`

SetIsOrgResolved sets IsOrgResolved field to given value.

### HasIsOrgResolved

`func (o *RecallsRow) HasIsOrgResolved() bool`

HasIsOrgResolved returns a boolean if a field has been set.

### SetIsOrgResolvedNil

`func (o *RecallsRow) SetIsOrgResolvedNil(b bool)`

 SetIsOrgResolvedNil sets the value for IsOrgResolved to be an explicit nil

### UnsetIsOrgResolved
`func (o *RecallsRow) UnsetIsOrgResolved()`

UnsetIsOrgResolved ensures that no value is present for IsOrgResolved, not even an explicit nil
### GetMatchTier

`func (o *RecallsRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *RecallsRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *RecallsRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *RecallsRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *RecallsRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *RecallsRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetOrgResolutionMethod

`func (o *RecallsRow) GetOrgResolutionMethod() string`

GetOrgResolutionMethod returns the OrgResolutionMethod field if non-nil, zero value otherwise.

### GetOrgResolutionMethodOk

`func (o *RecallsRow) GetOrgResolutionMethodOk() (*string, bool)`

GetOrgResolutionMethodOk returns a tuple with the OrgResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgResolutionMethod

`func (o *RecallsRow) SetOrgResolutionMethod(v string)`

SetOrgResolutionMethod sets OrgResolutionMethod field to given value.

### HasOrgResolutionMethod

`func (o *RecallsRow) HasOrgResolutionMethod() bool`

HasOrgResolutionMethod returns a boolean if a field has been set.

### SetOrgResolutionMethodNil

`func (o *RecallsRow) SetOrgResolutionMethodNil(b bool)`

 SetOrgResolutionMethodNil sets the value for OrgResolutionMethod to be an explicit nil

### UnsetOrgResolutionMethod
`func (o *RecallsRow) UnsetOrgResolutionMethod()`

UnsetOrgResolutionMethod ensures that no value is present for OrgResolutionMethod, not even an explicit nil
### GetProductDescription

`func (o *RecallsRow) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *RecallsRow) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *RecallsRow) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *RecallsRow) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### SetProductDescriptionNil

`func (o *RecallsRow) SetProductDescriptionNil(b bool)`

 SetProductDescriptionNil sets the value for ProductDescription to be an explicit nil

### UnsetProductDescription
`func (o *RecallsRow) UnsetProductDescription()`

UnsetProductDescription ensures that no value is present for ProductDescription, not even an explicit nil
### GetReasonForRecall

`func (o *RecallsRow) GetReasonForRecall() string`

GetReasonForRecall returns the ReasonForRecall field if non-nil, zero value otherwise.

### GetReasonForRecallOk

`func (o *RecallsRow) GetReasonForRecallOk() (*string, bool)`

GetReasonForRecallOk returns a tuple with the ReasonForRecall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForRecall

`func (o *RecallsRow) SetReasonForRecall(v string)`

SetReasonForRecall sets ReasonForRecall field to given value.

### HasReasonForRecall

`func (o *RecallsRow) HasReasonForRecall() bool`

HasReasonForRecall returns a boolean if a field has been set.

### SetReasonForRecallNil

`func (o *RecallsRow) SetReasonForRecallNil(b bool)`

 SetReasonForRecallNil sets the value for ReasonForRecall to be an explicit nil

### UnsetReasonForRecall
`func (o *RecallsRow) UnsetReasonForRecall()`

UnsetReasonForRecall ensures that no value is present for ReasonForRecall, not even an explicit nil
### GetRecallInitiationDate

`func (o *RecallsRow) GetRecallInitiationDate() string`

GetRecallInitiationDate returns the RecallInitiationDate field if non-nil, zero value otherwise.

### GetRecallInitiationDateOk

`func (o *RecallsRow) GetRecallInitiationDateOk() (*string, bool)`

GetRecallInitiationDateOk returns a tuple with the RecallInitiationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallInitiationDate

`func (o *RecallsRow) SetRecallInitiationDate(v string)`

SetRecallInitiationDate sets RecallInitiationDate field to given value.

### HasRecallInitiationDate

`func (o *RecallsRow) HasRecallInitiationDate() bool`

HasRecallInitiationDate returns a boolean if a field has been set.

### SetRecallInitiationDateNil

`func (o *RecallsRow) SetRecallInitiationDateNil(b bool)`

 SetRecallInitiationDateNil sets the value for RecallInitiationDate to be an explicit nil

### UnsetRecallInitiationDate
`func (o *RecallsRow) UnsetRecallInitiationDate()`

UnsetRecallInitiationDate ensures that no value is present for RecallInitiationDate, not even an explicit nil
### GetReportDate

`func (o *RecallsRow) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *RecallsRow) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *RecallsRow) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *RecallsRow) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### SetReportDateNil

`func (o *RecallsRow) SetReportDateNil(b bool)`

 SetReportDateNil sets the value for ReportDate to be an explicit nil

### UnsetReportDate
`func (o *RecallsRow) UnsetReportDate()`

UnsetReportDate ensures that no value is present for ReportDate, not even an explicit nil
### GetTerminationDate

`func (o *RecallsRow) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *RecallsRow) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *RecallsRow) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *RecallsRow) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### SetTerminationDateNil

`func (o *RecallsRow) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *RecallsRow) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


