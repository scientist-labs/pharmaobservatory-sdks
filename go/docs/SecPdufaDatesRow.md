# SecPdufaDatesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**PdufaObservationId** | Pointer to **NullableString** | The extracted disclosure id (filing document id and disclosure index). | [optional] 
**EventCategory** | Pointer to **NullableString** | pdufa, advisory_committee, priority_review or other. | [optional] 
**Status** | Pointer to **NullableString** | The disclosure&#39;s status: pending, approved, crl (complete response letter), or withdrawn. | [optional] 
**IsDecided** | Pointer to **NullableBool** | True when the status is approved, crl or withdrawn. | [optional] 
**IsForward** | Pointer to **NullableBool** | True when the disclosure is a pending, forward-looking date. | [optional] 
**ApplicationType** | Pointer to **NullableString** | The application type the filing names (NDA, BLA and so on). | [optional] 
**Designation** | Pointer to **NullableString** | The review designation the filing names (priority review, fast track and so on). | [optional] 
**Indication** | Pointer to **NullableString** | The indication the filing names. | [optional] 
**DatePhrase** | Pointer to **NullableString** | The date as the filing words it. | [optional] 
**DateStart** | Pointer to **NullableString** | Start of the date interval the phrase means. | [optional] 
**DateEnd** | Pointer to **NullableString** | End of the date interval the phrase means. | [optional] 
**DatePrecision** | Pointer to **NullableString** | How exact the phrase is (day, month, quarter and so on). | [optional] 
**IsInPeriodNow** | Pointer to **NullableBool** | True when today falls in the interval and the disclosure is not decided. | [optional] 
**DrugName** | Pointer to **NullableString** | The drug name as filed. | [optional] 
**SponsorName** | Pointer to **NullableString** | The filer&#39;s name. | [optional] 
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the SEC registrant&#39;s organization record (null when it is not in the spine). | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the filer&#39;s company (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the filer&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the filer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the filer&#39;s group. | [optional] 
**Ticker** | Pointer to **NullableString** | The stock ticker the filing names. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The filing&#39;s address on the SEC site. | [optional] 
**Confidence** | Pointer to **NullableFloat64** | The extraction&#39;s own confidence, from 0 to 1. | [optional] 

## Methods

### NewSecPdufaDatesRow

`func NewSecPdufaDatesRow() *SecPdufaDatesRow`

NewSecPdufaDatesRow instantiates a new SecPdufaDatesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecPdufaDatesRowWithDefaults

`func NewSecPdufaDatesRowWithDefaults() *SecPdufaDatesRow`

NewSecPdufaDatesRowWithDefaults instantiates a new SecPdufaDatesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *SecPdufaDatesRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *SecPdufaDatesRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *SecPdufaDatesRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *SecPdufaDatesRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *SecPdufaDatesRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *SecPdufaDatesRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetPdufaObservationId

`func (o *SecPdufaDatesRow) GetPdufaObservationId() string`

GetPdufaObservationId returns the PdufaObservationId field if non-nil, zero value otherwise.

### GetPdufaObservationIdOk

`func (o *SecPdufaDatesRow) GetPdufaObservationIdOk() (*string, bool)`

GetPdufaObservationIdOk returns a tuple with the PdufaObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdufaObservationId

`func (o *SecPdufaDatesRow) SetPdufaObservationId(v string)`

SetPdufaObservationId sets PdufaObservationId field to given value.

### HasPdufaObservationId

`func (o *SecPdufaDatesRow) HasPdufaObservationId() bool`

HasPdufaObservationId returns a boolean if a field has been set.

### SetPdufaObservationIdNil

`func (o *SecPdufaDatesRow) SetPdufaObservationIdNil(b bool)`

 SetPdufaObservationIdNil sets the value for PdufaObservationId to be an explicit nil

### UnsetPdufaObservationId
`func (o *SecPdufaDatesRow) UnsetPdufaObservationId()`

UnsetPdufaObservationId ensures that no value is present for PdufaObservationId, not even an explicit nil
### GetEventCategory

`func (o *SecPdufaDatesRow) GetEventCategory() string`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *SecPdufaDatesRow) GetEventCategoryOk() (*string, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *SecPdufaDatesRow) SetEventCategory(v string)`

SetEventCategory sets EventCategory field to given value.

### HasEventCategory

`func (o *SecPdufaDatesRow) HasEventCategory() bool`

HasEventCategory returns a boolean if a field has been set.

### SetEventCategoryNil

`func (o *SecPdufaDatesRow) SetEventCategoryNil(b bool)`

 SetEventCategoryNil sets the value for EventCategory to be an explicit nil

### UnsetEventCategory
`func (o *SecPdufaDatesRow) UnsetEventCategory()`

UnsetEventCategory ensures that no value is present for EventCategory, not even an explicit nil
### GetStatus

`func (o *SecPdufaDatesRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecPdufaDatesRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecPdufaDatesRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecPdufaDatesRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SecPdufaDatesRow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SecPdufaDatesRow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsDecided

`func (o *SecPdufaDatesRow) GetIsDecided() bool`

GetIsDecided returns the IsDecided field if non-nil, zero value otherwise.

### GetIsDecidedOk

`func (o *SecPdufaDatesRow) GetIsDecidedOk() (*bool, bool)`

GetIsDecidedOk returns a tuple with the IsDecided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDecided

`func (o *SecPdufaDatesRow) SetIsDecided(v bool)`

SetIsDecided sets IsDecided field to given value.

### HasIsDecided

`func (o *SecPdufaDatesRow) HasIsDecided() bool`

HasIsDecided returns a boolean if a field has been set.

### SetIsDecidedNil

`func (o *SecPdufaDatesRow) SetIsDecidedNil(b bool)`

 SetIsDecidedNil sets the value for IsDecided to be an explicit nil

### UnsetIsDecided
`func (o *SecPdufaDatesRow) UnsetIsDecided()`

UnsetIsDecided ensures that no value is present for IsDecided, not even an explicit nil
### GetIsForward

`func (o *SecPdufaDatesRow) GetIsForward() bool`

GetIsForward returns the IsForward field if non-nil, zero value otherwise.

### GetIsForwardOk

`func (o *SecPdufaDatesRow) GetIsForwardOk() (*bool, bool)`

GetIsForwardOk returns a tuple with the IsForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForward

`func (o *SecPdufaDatesRow) SetIsForward(v bool)`

SetIsForward sets IsForward field to given value.

### HasIsForward

`func (o *SecPdufaDatesRow) HasIsForward() bool`

HasIsForward returns a boolean if a field has been set.

### SetIsForwardNil

`func (o *SecPdufaDatesRow) SetIsForwardNil(b bool)`

 SetIsForwardNil sets the value for IsForward to be an explicit nil

### UnsetIsForward
`func (o *SecPdufaDatesRow) UnsetIsForward()`

UnsetIsForward ensures that no value is present for IsForward, not even an explicit nil
### GetApplicationType

`func (o *SecPdufaDatesRow) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *SecPdufaDatesRow) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *SecPdufaDatesRow) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *SecPdufaDatesRow) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationTypeNil

`func (o *SecPdufaDatesRow) SetApplicationTypeNil(b bool)`

 SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil

### UnsetApplicationType
`func (o *SecPdufaDatesRow) UnsetApplicationType()`

UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
### GetDesignation

`func (o *SecPdufaDatesRow) GetDesignation() string`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *SecPdufaDatesRow) GetDesignationOk() (*string, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *SecPdufaDatesRow) SetDesignation(v string)`

SetDesignation sets Designation field to given value.

### HasDesignation

`func (o *SecPdufaDatesRow) HasDesignation() bool`

HasDesignation returns a boolean if a field has been set.

### SetDesignationNil

`func (o *SecPdufaDatesRow) SetDesignationNil(b bool)`

 SetDesignationNil sets the value for Designation to be an explicit nil

### UnsetDesignation
`func (o *SecPdufaDatesRow) UnsetDesignation()`

UnsetDesignation ensures that no value is present for Designation, not even an explicit nil
### GetIndication

`func (o *SecPdufaDatesRow) GetIndication() string`

GetIndication returns the Indication field if non-nil, zero value otherwise.

### GetIndicationOk

`func (o *SecPdufaDatesRow) GetIndicationOk() (*string, bool)`

GetIndicationOk returns a tuple with the Indication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndication

`func (o *SecPdufaDatesRow) SetIndication(v string)`

SetIndication sets Indication field to given value.

### HasIndication

`func (o *SecPdufaDatesRow) HasIndication() bool`

HasIndication returns a boolean if a field has been set.

### SetIndicationNil

`func (o *SecPdufaDatesRow) SetIndicationNil(b bool)`

 SetIndicationNil sets the value for Indication to be an explicit nil

### UnsetIndication
`func (o *SecPdufaDatesRow) UnsetIndication()`

UnsetIndication ensures that no value is present for Indication, not even an explicit nil
### GetDatePhrase

`func (o *SecPdufaDatesRow) GetDatePhrase() string`

GetDatePhrase returns the DatePhrase field if non-nil, zero value otherwise.

### GetDatePhraseOk

`func (o *SecPdufaDatesRow) GetDatePhraseOk() (*string, bool)`

GetDatePhraseOk returns a tuple with the DatePhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePhrase

`func (o *SecPdufaDatesRow) SetDatePhrase(v string)`

SetDatePhrase sets DatePhrase field to given value.

### HasDatePhrase

`func (o *SecPdufaDatesRow) HasDatePhrase() bool`

HasDatePhrase returns a boolean if a field has been set.

### SetDatePhraseNil

`func (o *SecPdufaDatesRow) SetDatePhraseNil(b bool)`

 SetDatePhraseNil sets the value for DatePhrase to be an explicit nil

### UnsetDatePhrase
`func (o *SecPdufaDatesRow) UnsetDatePhrase()`

UnsetDatePhrase ensures that no value is present for DatePhrase, not even an explicit nil
### GetDateStart

`func (o *SecPdufaDatesRow) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *SecPdufaDatesRow) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *SecPdufaDatesRow) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *SecPdufaDatesRow) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### SetDateStartNil

`func (o *SecPdufaDatesRow) SetDateStartNil(b bool)`

 SetDateStartNil sets the value for DateStart to be an explicit nil

### UnsetDateStart
`func (o *SecPdufaDatesRow) UnsetDateStart()`

UnsetDateStart ensures that no value is present for DateStart, not even an explicit nil
### GetDateEnd

`func (o *SecPdufaDatesRow) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *SecPdufaDatesRow) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *SecPdufaDatesRow) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *SecPdufaDatesRow) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### SetDateEndNil

`func (o *SecPdufaDatesRow) SetDateEndNil(b bool)`

 SetDateEndNil sets the value for DateEnd to be an explicit nil

### UnsetDateEnd
`func (o *SecPdufaDatesRow) UnsetDateEnd()`

UnsetDateEnd ensures that no value is present for DateEnd, not even an explicit nil
### GetDatePrecision

`func (o *SecPdufaDatesRow) GetDatePrecision() string`

GetDatePrecision returns the DatePrecision field if non-nil, zero value otherwise.

### GetDatePrecisionOk

`func (o *SecPdufaDatesRow) GetDatePrecisionOk() (*string, bool)`

GetDatePrecisionOk returns a tuple with the DatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePrecision

`func (o *SecPdufaDatesRow) SetDatePrecision(v string)`

SetDatePrecision sets DatePrecision field to given value.

### HasDatePrecision

`func (o *SecPdufaDatesRow) HasDatePrecision() bool`

HasDatePrecision returns a boolean if a field has been set.

### SetDatePrecisionNil

`func (o *SecPdufaDatesRow) SetDatePrecisionNil(b bool)`

 SetDatePrecisionNil sets the value for DatePrecision to be an explicit nil

### UnsetDatePrecision
`func (o *SecPdufaDatesRow) UnsetDatePrecision()`

UnsetDatePrecision ensures that no value is present for DatePrecision, not even an explicit nil
### GetIsInPeriodNow

`func (o *SecPdufaDatesRow) GetIsInPeriodNow() bool`

GetIsInPeriodNow returns the IsInPeriodNow field if non-nil, zero value otherwise.

### GetIsInPeriodNowOk

`func (o *SecPdufaDatesRow) GetIsInPeriodNowOk() (*bool, bool)`

GetIsInPeriodNowOk returns a tuple with the IsInPeriodNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInPeriodNow

`func (o *SecPdufaDatesRow) SetIsInPeriodNow(v bool)`

SetIsInPeriodNow sets IsInPeriodNow field to given value.

### HasIsInPeriodNow

`func (o *SecPdufaDatesRow) HasIsInPeriodNow() bool`

HasIsInPeriodNow returns a boolean if a field has been set.

### SetIsInPeriodNowNil

`func (o *SecPdufaDatesRow) SetIsInPeriodNowNil(b bool)`

 SetIsInPeriodNowNil sets the value for IsInPeriodNow to be an explicit nil

### UnsetIsInPeriodNow
`func (o *SecPdufaDatesRow) UnsetIsInPeriodNow()`

UnsetIsInPeriodNow ensures that no value is present for IsInPeriodNow, not even an explicit nil
### GetDrugName

`func (o *SecPdufaDatesRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *SecPdufaDatesRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *SecPdufaDatesRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *SecPdufaDatesRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *SecPdufaDatesRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *SecPdufaDatesRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetSponsorName

`func (o *SecPdufaDatesRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *SecPdufaDatesRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *SecPdufaDatesRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *SecPdufaDatesRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *SecPdufaDatesRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *SecPdufaDatesRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorOrgId

`func (o *SecPdufaDatesRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *SecPdufaDatesRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *SecPdufaDatesRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *SecPdufaDatesRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *SecPdufaDatesRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *SecPdufaDatesRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorCompanyId

`func (o *SecPdufaDatesRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *SecPdufaDatesRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *SecPdufaDatesRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *SecPdufaDatesRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *SecPdufaDatesRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *SecPdufaDatesRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *SecPdufaDatesRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *SecPdufaDatesRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *SecPdufaDatesRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *SecPdufaDatesRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *SecPdufaDatesRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *SecPdufaDatesRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *SecPdufaDatesRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *SecPdufaDatesRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *SecPdufaDatesRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *SecPdufaDatesRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *SecPdufaDatesRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *SecPdufaDatesRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *SecPdufaDatesRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *SecPdufaDatesRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *SecPdufaDatesRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *SecPdufaDatesRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *SecPdufaDatesRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *SecPdufaDatesRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetTicker

`func (o *SecPdufaDatesRow) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *SecPdufaDatesRow) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *SecPdufaDatesRow) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *SecPdufaDatesRow) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### SetTickerNil

`func (o *SecPdufaDatesRow) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *SecPdufaDatesRow) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil
### GetSourceUrl

`func (o *SecPdufaDatesRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *SecPdufaDatesRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *SecPdufaDatesRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *SecPdufaDatesRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *SecPdufaDatesRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *SecPdufaDatesRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetConfidence

`func (o *SecPdufaDatesRow) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SecPdufaDatesRow) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SecPdufaDatesRow) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SecPdufaDatesRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *SecPdufaDatesRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *SecPdufaDatesRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


