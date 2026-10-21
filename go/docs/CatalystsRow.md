# CatalystsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **NullableString** | The event key: event_type, a colon and source_record_id. Unique. | [optional] 
**EventType** | Pointer to **NullableString** | What the event is: readout (a trial&#39;s primary completion), patent_cliff (loss of exclusivity), pdufa (an FDA decision date or advisory committee) or press_release (a Benzinga press-release event). | [optional] 
**EventCategory** | Pointer to **NullableString** | The kind of event within its type. readout: primary_completion. patent_cliff: the cliff basis (curated_loe, orange_book_patent, pte_expiry, bpcia_exclusivity). pdufa: the calendar&#39;s own category (pdufa, advisory_committee and others). press_release: the Benzinga event token, lowercased (clinical_trial_result, pdufa_assignment, data_presentation and others). Open vocabulary on the last two. | [optional] 
**Source** | Pointer to **NullableString** | The lane the row came from: ctgov (ClinicalTrials.gov), drug_exclusivity (the exclusivity registers), benzinga, both or sec_extraction (the PDUFA calendar, the same values as pdufa_dates) or benzinga_press_release. | [optional] 
**SourceRecordId** | Pointer to **NullableString** | The key of the row in its source, to join back to trials.trial_id, drug_patent_cliff.drug_id, pdufa_dates.pdufa_date_id or the press-release event id. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | none for an open row (ClinicalTrials.gov, the exclusivity registers, the SEC lane of the PDUFA calendar); licensable for a Benzinga row (the calendar or a press release). The api view gates rows on it. | [optional] 
**EventDate** | Pointer to **NullableString** | The date the calendar sorts on: a readout&#39;s primary completion date, the cliff date, the decision date, or the end of the press-release event&#39;s date window. | [optional] 
**EventYear** | Pointer to **NullableInt32** | The calendar year of event_date; the x axis of a by-year count. | [optional] 
**DateStart** | Pointer to **NullableString** | Start of the date interval the source states (the first day of the month for a readout the registry filed by month). | [optional] 
**DateEnd** | Pointer to **NullableString** | End of the date interval the source states. | [optional] 
**DatePrecision** | Pointer to **NullableString** | Precision of the date (day, month, quarter, half, year, fuzzy, relative). A readout is month where the registry filed only a month, in which case event_date is the last day of that month. | [optional] 
**DatePhrase** | Pointer to **NullableString** | The date as the source words it (an ISO date for a readout and a cliff; the press release&#39;s own phrase, such as H2 2026). | [optional] 
**DateStatus** | Pointer to **NullableString** | For a readout, whether the registry&#39;s completion date is estimated or actual (an anticipated date counts as estimated); null on the other legs. | [optional] 
**TrialId** | Pointer to **NullableString** | The ClinicalTrials.gov registry id (readout rows). | [optional] 
**TrialTitle** | Pointer to **NullableString** | The trial&#39;s brief title (readout rows). | [optional] 
**Phase** | Pointer to **NullableString** | The registry&#39;s phase code (readout rows). | [optional] 
**PhaseLabel** | Pointer to **NullableString** | The trial phase as a display label (readout rows). | [optional] 
**OverallStatus** | Pointer to **NullableString** | The registry&#39;s overall status code (readout rows). | [optional] 
**RelatedTrialIds** | Pointer to **[]string** | The registry ids a press-release event names; null on the other legs. | [optional] 
**Outcome** | Pointer to **NullableString** | The outcome a press-release event reports (a decided event stays in the forward book); null on the other legs. | [optional] 
**IndicationText** | Pointer to **NullableString** | The indication as the source words it (PDUFA and press-release rows). | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition key (PDUFA rows). | [optional] 
**ConditionName** | Pointer to **NullableString** | The condition&#39;s name (PDUFA rows). | [optional] 
**SourceType** | Pointer to **NullableString** | Kind of source document (SEC filing for the SEC lane, Press release for a press-release event); null on the open legs. | [optional] 
**SourceUrl** | Pointer to **NullableString** | Link to the source document (PDUFA and press-release rows). | [optional] 
**DrugName** | Pointer to **NullableString** | The drug name as filed or its display name; null on a readout (its drugs are the list drug_ids). | [optional] 
**DrugId** | Pointer to **NullableString** | The drug key of a patent-cliff, PDUFA or press-release row; null on a readout and where the drug did not resolve. | [optional] 
**MoleculeId** | Pointer to **NullableString** | The molecule key (PDUFA and press-release rows). | [optional] 
**DrugIds** | Pointer to **[]string** | Every drug key of the event: a readout carries each drug of its trial, sorted; the other legs carry their one drug, or an empty array. The drug predicate of a watchlist is &#39;&lt;drug&gt;&#39; &#x3D; ANY(drug_ids). | [optional] 
**SponsorRole** | Pointer to **NullableString** | What the organization is on the row: lead_sponsor (readout), applicant (patent cliff, the FDA applicant of record), sponsor (PDUFA) or issuer (press release, the company the ticker names). | [optional] 
**SponsorName** | Pointer to **NullableString** | The organization&#39;s name as filed; null on a press-release row (the source names a ticker, not an organization). | [optional] 
**SponsorOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record (a readout&#39;s lead sponsor, a cliff&#39;s applicant, a PDUFA row&#39;s filing organization; for a press release the identity hub&#39;s resolved entity for the ticker). Null when unresolved. | [optional] 
**SponsorCompanyId** | Pointer to **NullableString** | Company level: the company the organization belongs to (null when unresolved). | [optional] 
**SponsorCompanyName** | Pointer to **NullableString** | Display name of the organization&#39;s company. | [optional] 
**SponsorGroupId** | Pointer to **NullableString** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SponsorGroupName** | Pointer to **NullableString** | Display name of the organization&#39;s group. | [optional] 
**Ticker** | Pointer to **NullableString** | Exchange ticker (PDUFA and press-release rows). | [optional] 

## Methods

### NewCatalystsRow

`func NewCatalystsRow() *CatalystsRow`

NewCatalystsRow instantiates a new CatalystsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsRowWithDefaults

`func NewCatalystsRowWithDefaults() *CatalystsRow`

NewCatalystsRowWithDefaults instantiates a new CatalystsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *CatalystsRow) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *CatalystsRow) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *CatalystsRow) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *CatalystsRow) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventIdNil

`func (o *CatalystsRow) SetEventIdNil(b bool)`

 SetEventIdNil sets the value for EventId to be an explicit nil

### UnsetEventId
`func (o *CatalystsRow) UnsetEventId()`

UnsetEventId ensures that no value is present for EventId, not even an explicit nil
### GetEventType

`func (o *CatalystsRow) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CatalystsRow) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CatalystsRow) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CatalystsRow) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *CatalystsRow) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *CatalystsRow) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetEventCategory

`func (o *CatalystsRow) GetEventCategory() string`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *CatalystsRow) GetEventCategoryOk() (*string, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *CatalystsRow) SetEventCategory(v string)`

SetEventCategory sets EventCategory field to given value.

### HasEventCategory

`func (o *CatalystsRow) HasEventCategory() bool`

HasEventCategory returns a boolean if a field has been set.

### SetEventCategoryNil

`func (o *CatalystsRow) SetEventCategoryNil(b bool)`

 SetEventCategoryNil sets the value for EventCategory to be an explicit nil

### UnsetEventCategory
`func (o *CatalystsRow) UnsetEventCategory()`

UnsetEventCategory ensures that no value is present for EventCategory, not even an explicit nil
### GetSource

`func (o *CatalystsRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CatalystsRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CatalystsRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CatalystsRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CatalystsRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CatalystsRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceRecordId

`func (o *CatalystsRow) GetSourceRecordId() string`

GetSourceRecordId returns the SourceRecordId field if non-nil, zero value otherwise.

### GetSourceRecordIdOk

`func (o *CatalystsRow) GetSourceRecordIdOk() (*string, bool)`

GetSourceRecordIdOk returns a tuple with the SourceRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRecordId

`func (o *CatalystsRow) SetSourceRecordId(v string)`

SetSourceRecordId sets SourceRecordId field to given value.

### HasSourceRecordId

`func (o *CatalystsRow) HasSourceRecordId() bool`

HasSourceRecordId returns a boolean if a field has been set.

### SetSourceRecordIdNil

`func (o *CatalystsRow) SetSourceRecordIdNil(b bool)`

 SetSourceRecordIdNil sets the value for SourceRecordId to be an explicit nil

### UnsetSourceRecordId
`func (o *CatalystsRow) UnsetSourceRecordId()`

UnsetSourceRecordId ensures that no value is present for SourceRecordId, not even an explicit nil
### GetRestrictionTier

`func (o *CatalystsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *CatalystsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *CatalystsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *CatalystsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *CatalystsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *CatalystsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetEventDate

`func (o *CatalystsRow) GetEventDate() string`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *CatalystsRow) GetEventDateOk() (*string, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *CatalystsRow) SetEventDate(v string)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *CatalystsRow) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### SetEventDateNil

`func (o *CatalystsRow) SetEventDateNil(b bool)`

 SetEventDateNil sets the value for EventDate to be an explicit nil

### UnsetEventDate
`func (o *CatalystsRow) UnsetEventDate()`

UnsetEventDate ensures that no value is present for EventDate, not even an explicit nil
### GetEventYear

`func (o *CatalystsRow) GetEventYear() int32`

GetEventYear returns the EventYear field if non-nil, zero value otherwise.

### GetEventYearOk

`func (o *CatalystsRow) GetEventYearOk() (*int32, bool)`

GetEventYearOk returns a tuple with the EventYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventYear

`func (o *CatalystsRow) SetEventYear(v int32)`

SetEventYear sets EventYear field to given value.

### HasEventYear

`func (o *CatalystsRow) HasEventYear() bool`

HasEventYear returns a boolean if a field has been set.

### SetEventYearNil

`func (o *CatalystsRow) SetEventYearNil(b bool)`

 SetEventYearNil sets the value for EventYear to be an explicit nil

### UnsetEventYear
`func (o *CatalystsRow) UnsetEventYear()`

UnsetEventYear ensures that no value is present for EventYear, not even an explicit nil
### GetDateStart

`func (o *CatalystsRow) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *CatalystsRow) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *CatalystsRow) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *CatalystsRow) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### SetDateStartNil

`func (o *CatalystsRow) SetDateStartNil(b bool)`

 SetDateStartNil sets the value for DateStart to be an explicit nil

### UnsetDateStart
`func (o *CatalystsRow) UnsetDateStart()`

UnsetDateStart ensures that no value is present for DateStart, not even an explicit nil
### GetDateEnd

`func (o *CatalystsRow) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *CatalystsRow) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *CatalystsRow) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *CatalystsRow) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### SetDateEndNil

`func (o *CatalystsRow) SetDateEndNil(b bool)`

 SetDateEndNil sets the value for DateEnd to be an explicit nil

### UnsetDateEnd
`func (o *CatalystsRow) UnsetDateEnd()`

UnsetDateEnd ensures that no value is present for DateEnd, not even an explicit nil
### GetDatePrecision

`func (o *CatalystsRow) GetDatePrecision() string`

GetDatePrecision returns the DatePrecision field if non-nil, zero value otherwise.

### GetDatePrecisionOk

`func (o *CatalystsRow) GetDatePrecisionOk() (*string, bool)`

GetDatePrecisionOk returns a tuple with the DatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePrecision

`func (o *CatalystsRow) SetDatePrecision(v string)`

SetDatePrecision sets DatePrecision field to given value.

### HasDatePrecision

`func (o *CatalystsRow) HasDatePrecision() bool`

HasDatePrecision returns a boolean if a field has been set.

### SetDatePrecisionNil

`func (o *CatalystsRow) SetDatePrecisionNil(b bool)`

 SetDatePrecisionNil sets the value for DatePrecision to be an explicit nil

### UnsetDatePrecision
`func (o *CatalystsRow) UnsetDatePrecision()`

UnsetDatePrecision ensures that no value is present for DatePrecision, not even an explicit nil
### GetDatePhrase

`func (o *CatalystsRow) GetDatePhrase() string`

GetDatePhrase returns the DatePhrase field if non-nil, zero value otherwise.

### GetDatePhraseOk

`func (o *CatalystsRow) GetDatePhraseOk() (*string, bool)`

GetDatePhraseOk returns a tuple with the DatePhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePhrase

`func (o *CatalystsRow) SetDatePhrase(v string)`

SetDatePhrase sets DatePhrase field to given value.

### HasDatePhrase

`func (o *CatalystsRow) HasDatePhrase() bool`

HasDatePhrase returns a boolean if a field has been set.

### SetDatePhraseNil

`func (o *CatalystsRow) SetDatePhraseNil(b bool)`

 SetDatePhraseNil sets the value for DatePhrase to be an explicit nil

### UnsetDatePhrase
`func (o *CatalystsRow) UnsetDatePhrase()`

UnsetDatePhrase ensures that no value is present for DatePhrase, not even an explicit nil
### GetDateStatus

`func (o *CatalystsRow) GetDateStatus() string`

GetDateStatus returns the DateStatus field if non-nil, zero value otherwise.

### GetDateStatusOk

`func (o *CatalystsRow) GetDateStatusOk() (*string, bool)`

GetDateStatusOk returns a tuple with the DateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStatus

`func (o *CatalystsRow) SetDateStatus(v string)`

SetDateStatus sets DateStatus field to given value.

### HasDateStatus

`func (o *CatalystsRow) HasDateStatus() bool`

HasDateStatus returns a boolean if a field has been set.

### SetDateStatusNil

`func (o *CatalystsRow) SetDateStatusNil(b bool)`

 SetDateStatusNil sets the value for DateStatus to be an explicit nil

### UnsetDateStatus
`func (o *CatalystsRow) UnsetDateStatus()`

UnsetDateStatus ensures that no value is present for DateStatus, not even an explicit nil
### GetTrialId

`func (o *CatalystsRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *CatalystsRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *CatalystsRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *CatalystsRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *CatalystsRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *CatalystsRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetTrialTitle

`func (o *CatalystsRow) GetTrialTitle() string`

GetTrialTitle returns the TrialTitle field if non-nil, zero value otherwise.

### GetTrialTitleOk

`func (o *CatalystsRow) GetTrialTitleOk() (*string, bool)`

GetTrialTitleOk returns a tuple with the TrialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialTitle

`func (o *CatalystsRow) SetTrialTitle(v string)`

SetTrialTitle sets TrialTitle field to given value.

### HasTrialTitle

`func (o *CatalystsRow) HasTrialTitle() bool`

HasTrialTitle returns a boolean if a field has been set.

### SetTrialTitleNil

`func (o *CatalystsRow) SetTrialTitleNil(b bool)`

 SetTrialTitleNil sets the value for TrialTitle to be an explicit nil

### UnsetTrialTitle
`func (o *CatalystsRow) UnsetTrialTitle()`

UnsetTrialTitle ensures that no value is present for TrialTitle, not even an explicit nil
### GetPhase

`func (o *CatalystsRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CatalystsRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CatalystsRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CatalystsRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *CatalystsRow) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *CatalystsRow) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseLabel

`func (o *CatalystsRow) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *CatalystsRow) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *CatalystsRow) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.

### HasPhaseLabel

`func (o *CatalystsRow) HasPhaseLabel() bool`

HasPhaseLabel returns a boolean if a field has been set.

### SetPhaseLabelNil

`func (o *CatalystsRow) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *CatalystsRow) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetOverallStatus

`func (o *CatalystsRow) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *CatalystsRow) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *CatalystsRow) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.

### HasOverallStatus

`func (o *CatalystsRow) HasOverallStatus() bool`

HasOverallStatus returns a boolean if a field has been set.

### SetOverallStatusNil

`func (o *CatalystsRow) SetOverallStatusNil(b bool)`

 SetOverallStatusNil sets the value for OverallStatus to be an explicit nil

### UnsetOverallStatus
`func (o *CatalystsRow) UnsetOverallStatus()`

UnsetOverallStatus ensures that no value is present for OverallStatus, not even an explicit nil
### GetRelatedTrialIds

`func (o *CatalystsRow) GetRelatedTrialIds() []string`

GetRelatedTrialIds returns the RelatedTrialIds field if non-nil, zero value otherwise.

### GetRelatedTrialIdsOk

`func (o *CatalystsRow) GetRelatedTrialIdsOk() (*[]string, bool)`

GetRelatedTrialIdsOk returns a tuple with the RelatedTrialIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTrialIds

`func (o *CatalystsRow) SetRelatedTrialIds(v []string)`

SetRelatedTrialIds sets RelatedTrialIds field to given value.

### HasRelatedTrialIds

`func (o *CatalystsRow) HasRelatedTrialIds() bool`

HasRelatedTrialIds returns a boolean if a field has been set.

### SetRelatedTrialIdsNil

`func (o *CatalystsRow) SetRelatedTrialIdsNil(b bool)`

 SetRelatedTrialIdsNil sets the value for RelatedTrialIds to be an explicit nil

### UnsetRelatedTrialIds
`func (o *CatalystsRow) UnsetRelatedTrialIds()`

UnsetRelatedTrialIds ensures that no value is present for RelatedTrialIds, not even an explicit nil
### GetOutcome

`func (o *CatalystsRow) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *CatalystsRow) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *CatalystsRow) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *CatalystsRow) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *CatalystsRow) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *CatalystsRow) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetIndicationText

`func (o *CatalystsRow) GetIndicationText() string`

GetIndicationText returns the IndicationText field if non-nil, zero value otherwise.

### GetIndicationTextOk

`func (o *CatalystsRow) GetIndicationTextOk() (*string, bool)`

GetIndicationTextOk returns a tuple with the IndicationText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicationText

`func (o *CatalystsRow) SetIndicationText(v string)`

SetIndicationText sets IndicationText field to given value.

### HasIndicationText

`func (o *CatalystsRow) HasIndicationText() bool`

HasIndicationText returns a boolean if a field has been set.

### SetIndicationTextNil

`func (o *CatalystsRow) SetIndicationTextNil(b bool)`

 SetIndicationTextNil sets the value for IndicationText to be an explicit nil

### UnsetIndicationText
`func (o *CatalystsRow) UnsetIndicationText()`

UnsetIndicationText ensures that no value is present for IndicationText, not even an explicit nil
### GetConditionId

`func (o *CatalystsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *CatalystsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *CatalystsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *CatalystsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *CatalystsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *CatalystsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetConditionName

`func (o *CatalystsRow) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *CatalystsRow) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *CatalystsRow) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *CatalystsRow) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### SetConditionNameNil

`func (o *CatalystsRow) SetConditionNameNil(b bool)`

 SetConditionNameNil sets the value for ConditionName to be an explicit nil

### UnsetConditionName
`func (o *CatalystsRow) UnsetConditionName()`

UnsetConditionName ensures that no value is present for ConditionName, not even an explicit nil
### GetSourceType

`func (o *CatalystsRow) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CatalystsRow) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CatalystsRow) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *CatalystsRow) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *CatalystsRow) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *CatalystsRow) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetSourceUrl

`func (o *CatalystsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CatalystsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CatalystsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *CatalystsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *CatalystsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *CatalystsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetDrugName

`func (o *CatalystsRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *CatalystsRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *CatalystsRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *CatalystsRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *CatalystsRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *CatalystsRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetDrugId

`func (o *CatalystsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *CatalystsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *CatalystsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *CatalystsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *CatalystsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *CatalystsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetMoleculeId

`func (o *CatalystsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *CatalystsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *CatalystsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *CatalystsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *CatalystsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *CatalystsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetDrugIds

`func (o *CatalystsRow) GetDrugIds() []string`

GetDrugIds returns the DrugIds field if non-nil, zero value otherwise.

### GetDrugIdsOk

`func (o *CatalystsRow) GetDrugIdsOk() (*[]string, bool)`

GetDrugIdsOk returns a tuple with the DrugIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugIds

`func (o *CatalystsRow) SetDrugIds(v []string)`

SetDrugIds sets DrugIds field to given value.

### HasDrugIds

`func (o *CatalystsRow) HasDrugIds() bool`

HasDrugIds returns a boolean if a field has been set.

### SetDrugIdsNil

`func (o *CatalystsRow) SetDrugIdsNil(b bool)`

 SetDrugIdsNil sets the value for DrugIds to be an explicit nil

### UnsetDrugIds
`func (o *CatalystsRow) UnsetDrugIds()`

UnsetDrugIds ensures that no value is present for DrugIds, not even an explicit nil
### GetSponsorRole

`func (o *CatalystsRow) GetSponsorRole() string`

GetSponsorRole returns the SponsorRole field if non-nil, zero value otherwise.

### GetSponsorRoleOk

`func (o *CatalystsRow) GetSponsorRoleOk() (*string, bool)`

GetSponsorRoleOk returns a tuple with the SponsorRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorRole

`func (o *CatalystsRow) SetSponsorRole(v string)`

SetSponsorRole sets SponsorRole field to given value.

### HasSponsorRole

`func (o *CatalystsRow) HasSponsorRole() bool`

HasSponsorRole returns a boolean if a field has been set.

### SetSponsorRoleNil

`func (o *CatalystsRow) SetSponsorRoleNil(b bool)`

 SetSponsorRoleNil sets the value for SponsorRole to be an explicit nil

### UnsetSponsorRole
`func (o *CatalystsRow) UnsetSponsorRole()`

UnsetSponsorRole ensures that no value is present for SponsorRole, not even an explicit nil
### GetSponsorName

`func (o *CatalystsRow) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *CatalystsRow) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *CatalystsRow) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *CatalystsRow) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### SetSponsorNameNil

`func (o *CatalystsRow) SetSponsorNameNil(b bool)`

 SetSponsorNameNil sets the value for SponsorName to be an explicit nil

### UnsetSponsorName
`func (o *CatalystsRow) UnsetSponsorName()`

UnsetSponsorName ensures that no value is present for SponsorName, not even an explicit nil
### GetSponsorOrgId

`func (o *CatalystsRow) GetSponsorOrgId() string`

GetSponsorOrgId returns the SponsorOrgId field if non-nil, zero value otherwise.

### GetSponsorOrgIdOk

`func (o *CatalystsRow) GetSponsorOrgIdOk() (*string, bool)`

GetSponsorOrgIdOk returns a tuple with the SponsorOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorOrgId

`func (o *CatalystsRow) SetSponsorOrgId(v string)`

SetSponsorOrgId sets SponsorOrgId field to given value.

### HasSponsorOrgId

`func (o *CatalystsRow) HasSponsorOrgId() bool`

HasSponsorOrgId returns a boolean if a field has been set.

### SetSponsorOrgIdNil

`func (o *CatalystsRow) SetSponsorOrgIdNil(b bool)`

 SetSponsorOrgIdNil sets the value for SponsorOrgId to be an explicit nil

### UnsetSponsorOrgId
`func (o *CatalystsRow) UnsetSponsorOrgId()`

UnsetSponsorOrgId ensures that no value is present for SponsorOrgId, not even an explicit nil
### GetSponsorCompanyId

`func (o *CatalystsRow) GetSponsorCompanyId() string`

GetSponsorCompanyId returns the SponsorCompanyId field if non-nil, zero value otherwise.

### GetSponsorCompanyIdOk

`func (o *CatalystsRow) GetSponsorCompanyIdOk() (*string, bool)`

GetSponsorCompanyIdOk returns a tuple with the SponsorCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyId

`func (o *CatalystsRow) SetSponsorCompanyId(v string)`

SetSponsorCompanyId sets SponsorCompanyId field to given value.

### HasSponsorCompanyId

`func (o *CatalystsRow) HasSponsorCompanyId() bool`

HasSponsorCompanyId returns a boolean if a field has been set.

### SetSponsorCompanyIdNil

`func (o *CatalystsRow) SetSponsorCompanyIdNil(b bool)`

 SetSponsorCompanyIdNil sets the value for SponsorCompanyId to be an explicit nil

### UnsetSponsorCompanyId
`func (o *CatalystsRow) UnsetSponsorCompanyId()`

UnsetSponsorCompanyId ensures that no value is present for SponsorCompanyId, not even an explicit nil
### GetSponsorCompanyName

`func (o *CatalystsRow) GetSponsorCompanyName() string`

GetSponsorCompanyName returns the SponsorCompanyName field if non-nil, zero value otherwise.

### GetSponsorCompanyNameOk

`func (o *CatalystsRow) GetSponsorCompanyNameOk() (*string, bool)`

GetSponsorCompanyNameOk returns a tuple with the SponsorCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCompanyName

`func (o *CatalystsRow) SetSponsorCompanyName(v string)`

SetSponsorCompanyName sets SponsorCompanyName field to given value.

### HasSponsorCompanyName

`func (o *CatalystsRow) HasSponsorCompanyName() bool`

HasSponsorCompanyName returns a boolean if a field has been set.

### SetSponsorCompanyNameNil

`func (o *CatalystsRow) SetSponsorCompanyNameNil(b bool)`

 SetSponsorCompanyNameNil sets the value for SponsorCompanyName to be an explicit nil

### UnsetSponsorCompanyName
`func (o *CatalystsRow) UnsetSponsorCompanyName()`

UnsetSponsorCompanyName ensures that no value is present for SponsorCompanyName, not even an explicit nil
### GetSponsorGroupId

`func (o *CatalystsRow) GetSponsorGroupId() string`

GetSponsorGroupId returns the SponsorGroupId field if non-nil, zero value otherwise.

### GetSponsorGroupIdOk

`func (o *CatalystsRow) GetSponsorGroupIdOk() (*string, bool)`

GetSponsorGroupIdOk returns a tuple with the SponsorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupId

`func (o *CatalystsRow) SetSponsorGroupId(v string)`

SetSponsorGroupId sets SponsorGroupId field to given value.

### HasSponsorGroupId

`func (o *CatalystsRow) HasSponsorGroupId() bool`

HasSponsorGroupId returns a boolean if a field has been set.

### SetSponsorGroupIdNil

`func (o *CatalystsRow) SetSponsorGroupIdNil(b bool)`

 SetSponsorGroupIdNil sets the value for SponsorGroupId to be an explicit nil

### UnsetSponsorGroupId
`func (o *CatalystsRow) UnsetSponsorGroupId()`

UnsetSponsorGroupId ensures that no value is present for SponsorGroupId, not even an explicit nil
### GetSponsorGroupName

`func (o *CatalystsRow) GetSponsorGroupName() string`

GetSponsorGroupName returns the SponsorGroupName field if non-nil, zero value otherwise.

### GetSponsorGroupNameOk

`func (o *CatalystsRow) GetSponsorGroupNameOk() (*string, bool)`

GetSponsorGroupNameOk returns a tuple with the SponsorGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorGroupName

`func (o *CatalystsRow) SetSponsorGroupName(v string)`

SetSponsorGroupName sets SponsorGroupName field to given value.

### HasSponsorGroupName

`func (o *CatalystsRow) HasSponsorGroupName() bool`

HasSponsorGroupName returns a boolean if a field has been set.

### SetSponsorGroupNameNil

`func (o *CatalystsRow) SetSponsorGroupNameNil(b bool)`

 SetSponsorGroupNameNil sets the value for SponsorGroupName to be an explicit nil

### UnsetSponsorGroupName
`func (o *CatalystsRow) UnsetSponsorGroupName()`

UnsetSponsorGroupName ensures that no value is present for SponsorGroupName, not even an explicit nil
### GetTicker

`func (o *CatalystsRow) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *CatalystsRow) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *CatalystsRow) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *CatalystsRow) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### SetTickerNil

`func (o *CatalystsRow) SetTickerNil(b bool)`

 SetTickerNil sets the value for Ticker to be an explicit nil

### UnsetTicker
`func (o *CatalystsRow) UnsetTicker()`

UnsetTicker ensures that no value is present for Ticker, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


