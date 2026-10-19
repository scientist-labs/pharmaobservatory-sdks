# PdufaDatesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PdufaDateId** | Pointer to **NullableString** | The calendar entry id. | [optional] 
**Source** | Pointer to **NullableString** | Which lane the row came from: benzinga, sec_extraction or both. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | none for a row from SEC filings; licensable for a row from the Benzinga calendar (the filings are those made to the US Securities and Exchange Commission). | [optional] 
**DecisionDate** | Pointer to **NullableString** | The decision date of the US Food and Drug Administration (FDA); a PDUFA date is the goal date under the Prescription Drug User Fee Act. | [optional] 
**DateStart** | Pointer to **NullableString** | Start of the date interval the source states. | [optional] 
**DateEnd** | Pointer to **NullableString** | End of the date interval the source states. | [optional] 
**DatePrecision** | Pointer to **NullableString** | Precision of the date (day, month, quarter). | [optional] 
**DatePhrase** | Pointer to **NullableString** | The date as the source words it. | [optional] 
**EventType** | Pointer to **NullableString** | Event type. | [optional] 
**EventCategory** | Pointer to **NullableString** | Event category (pdufa, adcomm and so on). | [optional] 
**DrugName** | Pointer to **NullableString** | The drug name as filed. | [optional] 
**DrugId** | Pointer to **NullableString** | The drug key. | [optional] 
**MoleculeId** | Pointer to **NullableString** | The molecule key. | [optional] 
**SponsorName** | Pointer to **NullableString** | The sponsor name as filed. | [optional] 
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the sponsor resolved to (null when unresolved). | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the sponsor belongs to (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the sponsor&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the sponsor&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the sponsor&#39;s group. | [optional] 
**Ticker** | Pointer to **NullableString** | Exchange ticker. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key. | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s name. | [optional] 
**IndicationText** | Pointer to **NullableString** | The indication as the source words it. | [optional] 
**SourceType** | Pointer to **NullableString** | Kind of source document. | [optional] 
**SourceLink** | Pointer to **NullableString** | Link to the source document. | [optional] 

## Methods

### NewPdufaDatesRow

`func NewPdufaDatesRow() *PdufaDatesRow`

NewPdufaDatesRow instantiates a new PdufaDatesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPdufaDatesRowWithDefaults

`func NewPdufaDatesRowWithDefaults() *PdufaDatesRow`

NewPdufaDatesRowWithDefaults instantiates a new PdufaDatesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPdufaDateId

`func (o *PdufaDatesRow) GetPdufaDateId() string`

GetPdufaDateId returns the PdufaDateId field if non-nil, zero value otherwise.

### GetPdufaDateIdOk

`func (o *PdufaDatesRow) GetPdufaDateIdOk() (*string, bool)`

GetPdufaDateIdOk returns a tuple with the PdufaDateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdufaDateId

`func (o *PdufaDatesRow) SetPdufaDateId(v string)`

SetPdufaDateId sets PdufaDateId field to given value.

### HasPdufaDateId

`func (o *PdufaDatesRow) HasPdufaDateId() bool`

HasPdufaDateId returns a boolean if a field has been set.

### SetPdufaDateIdNil

`func (o *PdufaDatesRow) SetPdufaDateIdNil(b bool)`

 SetPdufaDateIdNil sets the value for PdufaDateId to be an explicit nil

### UnsetPdufaDateId
`func (o *PdufaDatesRow) UnsetPdufaDateId()`

UnsetPdufaDateId ensures that no value is present for PdufaDateId, not even an explicit nil
### GetSource

`func (o *PdufaDatesRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PdufaDatesRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PdufaDatesRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PdufaDatesRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *PdufaDatesRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *PdufaDatesRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetRestrictionTier

`func (o *PdufaDatesRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *PdufaDatesRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *PdufaDatesRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *PdufaDatesRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *PdufaDatesRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *PdufaDatesRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetDecisionDate

`func (o *PdufaDatesRow) GetDecisionDate() string`

GetDecisionDate returns the DecisionDate field if non-nil, zero value otherwise.

### GetDecisionDateOk

`func (o *PdufaDatesRow) GetDecisionDateOk() (*string, bool)`

GetDecisionDateOk returns a tuple with the DecisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDate

`func (o *PdufaDatesRow) SetDecisionDate(v string)`

SetDecisionDate sets DecisionDate field to given value.

### HasDecisionDate

`func (o *PdufaDatesRow) HasDecisionDate() bool`

HasDecisionDate returns a boolean if a field has been set.

### SetDecisionDateNil

`func (o *PdufaDatesRow) SetDecisionDateNil(b bool)`

 SetDecisionDateNil sets the value for DecisionDate to be an explicit nil

### UnsetDecisionDate
`func (o *PdufaDatesRow) UnsetDecisionDate()`

UnsetDecisionDate ensures that no value is present for DecisionDate, not even an explicit nil
### GetDateStart

`func (o *PdufaDatesRow) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *PdufaDatesRow) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *PdufaDatesRow) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *PdufaDatesRow) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### SetDateStartNil

`func (o *PdufaDatesRow) SetDateStartNil(b bool)`

 SetDateStartNil sets the value for DateStart to be an explicit nil

### UnsetDateStart
`func (o *PdufaDatesRow) UnsetDateStart()`

UnsetDateStart ensures that no value is present for DateStart, not even an explicit nil
### GetDateEnd

`func (o *PdufaDatesRow) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *PdufaDatesRow) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *PdufaDatesRow) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *PdufaDatesRow) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### SetDateEndNil

`func (o *PdufaDatesRow) SetDateEndNil(b bool)`

 SetDateEndNil sets the value for DateEnd to be an explicit nil

### UnsetDateEnd
`func (o *PdufaDatesRow) UnsetDateEnd()`

UnsetDateEnd ensures that no value is present for DateEnd, not even an explicit nil
### GetDatePrecision

`func (o *PdufaDatesRow) GetDatePrecision() string`

GetDatePrecision returns the DatePrecision field if non-nil, zero value otherwise.

### GetDatePrecisionOk

`func (o *PdufaDatesRow) GetDatePrecisionOk() (*string, bool)`

GetDatePrecisionOk returns a tuple with the DatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePrecision

`func (o *PdufaDatesRow) SetDatePrecision(v string)`

SetDatePrecision sets DatePrecision field to given value.

### HasDatePrecision

`func (o *PdufaDatesRow) HasDatePrecision() bool`

HasDatePrecision returns a boolean if a field has been set.

### SetDatePrecisionNil

`func (o *PdufaDatesRow) SetDatePrecisionNil(b bool)`

 SetDatePrecisionNil sets the value for DatePrecision to be an explicit nil

### UnsetDatePrecision
`func (o *PdufaDatesRow) UnsetDatePrecision()`

UnsetDatePrecision ensures that no value is present for DatePrecision, not even an explicit nil
### GetDatePhrase

`func (o *PdufaDatesRow) GetDatePhrase() string`

GetDatePhrase returns the DatePhrase field if non-nil, zero value otherwise.

### GetDatePhraseOk

`func (o *PdufaDatesRow) GetDatePhraseOk() (*string, bool)`

GetDatePhraseOk returns a tuple with the DatePhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePhrase

`func (o *PdufaDatesRow) SetDatePhrase(v string)`

SetDatePhrase sets DatePhrase field to given value.

### HasDatePhrase

`func (o *PdufaDatesRow) HasDatePhrase() bool`

HasDatePhrase returns a boolean if a field has been set.

### SetDatePhraseNil

`func (o *PdufaDatesRow) SetDatePhraseNil(b bool)`

 SetDatePhraseNil sets the value for DatePhrase to be an explicit nil

### UnsetDatePhrase
`func (o *PdufaDatesRow) UnsetDatePhrase()`

UnsetDatePhrase ensures that no value is present for DatePhrase, not even an explicit nil
### GetEventType

`func (o *PdufaDatesRow) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PdufaDatesRow) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PdufaDatesRow) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *PdufaDatesRow) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *PdufaDatesRow) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *PdufaDatesRow) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetEventCategory

`func (o *PdufaDatesRow) GetEventCategory() string`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *PdufaDatesRow) GetEventCategoryOk() (*string, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *PdufaDatesRow) SetEventCategory(v string)`

SetEventCategory sets EventCategory field to given value.

### HasEventCategory

`func (o *PdufaDatesRow) HasEventCategory() bool`

HasEventCategory returns a boolean if a field has been set.

### SetEventCategoryNil

`func (o *PdufaDatesRow) SetEventCategoryNil(b bool)`

 SetEventCategoryNil sets the value for EventCategory to be an explicit nil

### UnsetEventCategory
`func (o *PdufaDatesRow) UnsetEventCategory()`

UnsetEventCategory ensures that no value is present for EventCategory, not even an explicit nil
### GetDrugName

`func (o *PdufaDatesRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *PdufaDatesRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *PdufaDatesRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *PdufaDatesRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *PdufaDatesRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *PdufaDatesRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetDrugId

`func (o *PdufaDatesRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *PdufaDatesRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *PdufaDatesRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *PdufaDatesRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *PdufaDatesRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *PdufaDatesRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetMoleculeId

`func (o *PdufaDatesRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *PdufaDatesRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *PdufaDatesRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *PdufaDatesRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *PdufaDatesRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *PdufaDatesRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetSponsorName

`func (o *PdufaDatesRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *PdufaDatesRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *PdufaDatesRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *PdufaDatesRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *PdufaDatesRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *PdufaDatesRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorOrgId

`func (o *PdufaDatesRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *PdufaDatesRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *PdufaDatesRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *PdufaDatesRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *PdufaDatesRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *PdufaDatesRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorCompanyId

`func (o *PdufaDatesRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *PdufaDatesRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *PdufaDatesRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *PdufaDatesRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *PdufaDatesRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *PdufaDatesRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *PdufaDatesRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *PdufaDatesRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *PdufaDatesRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *PdufaDatesRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *PdufaDatesRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *PdufaDatesRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *PdufaDatesRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *PdufaDatesRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *PdufaDatesRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *PdufaDatesRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *PdufaDatesRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *PdufaDatesRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *PdufaDatesRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *PdufaDatesRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *PdufaDatesRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *PdufaDatesRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *PdufaDatesRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *PdufaDatesRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetTicker

`func (o *PdufaDatesRow) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *PdufaDatesRow) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *PdufaDatesRow) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *PdufaDatesRow) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### SetTickerNil

`func (o *PdufaDatesRow) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *PdufaDatesRow) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil
### GetConditionId

`func (o *PdufaDatesRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *PdufaDatesRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *PdufaDatesRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *PdufaDatesRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *PdufaDatesRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *PdufaDatesRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *PdufaDatesRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *PdufaDatesRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *PdufaDatesRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *PdufaDatesRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *PdufaDatesRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *PdufaDatesRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetIndicationText

`func (o *PdufaDatesRow) GetIndicationText() string`

GetIndicationText returns the IndicationText field if non-nil, zero value otherwise.

### GetIndicationTextOk

`func (o *PdufaDatesRow) GetIndicationTextOk() (*string, bool)`

GetIndicationTextOk returns a tuple with the IndicationText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicationText

`func (o *PdufaDatesRow) SetIndicationText(v string)`

SetIndicationText sets IndicationText field to given value.

### HasIndicationText

`func (o *PdufaDatesRow) HasIndicationText() bool`

HasIndicationText returns a boolean if a field has been set.

### SetIndicationTextNil

`func (o *PdufaDatesRow) SetIndicationTextNil(b bool)`

 SetIndicationTextNil sets the value for IndicationText to be an explicit nil

### UnsetIndicationText
`func (o *PdufaDatesRow) UnsetIndicationText()`

UnsetIndicationText ensures that no value is present for IndicationText, not even an explicit nil
### GetSourceType

`func (o *PdufaDatesRow) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PdufaDatesRow) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PdufaDatesRow) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *PdufaDatesRow) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *PdufaDatesRow) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *PdufaDatesRow) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetSourceLink

`func (o *PdufaDatesRow) GetSourceLink() string`

GetSourceLink returns the SourceLink field if non-nil, zero value otherwise.

### GetSourceLinkOk

`func (o *PdufaDatesRow) GetSourceLinkOk() (*string, bool)`

GetSourceLinkOk returns a tuple with the SourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLink

`func (o *PdufaDatesRow) SetSourceLink(v string)`

SetSourceLink sets SourceLink field to given value.

### HasSourceLink

`func (o *PdufaDatesRow) HasSourceLink() bool`

HasSourceLink returns a boolean if a field has been set.

### SetSourceLinkNil

`func (o *PdufaDatesRow) SetSourceLinkNil(b bool)`

 SetSourceLinkNil sets the value for SourceLink to be an explicit nil

### UnsetSourceLink
`func (o *PdufaDatesRow) UnsetSourceLink()`

UnsetSourceLink ensures that no value is present for SourceLink, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


