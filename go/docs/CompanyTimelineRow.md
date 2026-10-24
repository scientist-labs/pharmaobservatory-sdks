# CompanyTimelineRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimelineEventId** | Pointer to **NullableString** | md5 of the lane, event type, source record and company: a stable surrogate id for the row. | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: the company the row is about; the entry key for the companies object type. | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**ReportingOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the lane reported the event on (the lowest id when several records of the company report it). | [optional] 
**SourceKey** | Pointer to **NullableString** | The lane that supplied the event. | [optional] 
**SourceEntityKey** | Pointer to **NullableString** | The lane&#39;s own key for the thing that happened (trial id, funding event id, patent id, enforcement event id, statement id). | [optional] 
**EventType** | Pointer to **NullableString** | The fine event vocabulary; funding types are the funding round types verbatim. | [optional] 
**EventFamily** | Pointer to **NullableString** | The family of the event type, assigned by name from the lane registry: clinical, capital, intellectual_property, regulatory, compliance or other. | [optional] 
**EventDate** | Pointer to **NullableString** | The first day consistent with the evidence, never the raw value when the evidence is coarser than a day. | [optional] 
**DatePrecision** | Pointer to **NullableString** | How precise the evidence is: day, month, quarter, half or year. | [optional] 
**PrecisionWindowEnd** | Pointer to **NullableString** | The last day consistent with the evidence; equals event_date at day precision. | [optional] 
**DatePrecisionBasis** | Pointer to **NullableString** | Where the precision claim comes from. Provenance for an audit, not a vocabulary to branch on. | [optional] 
**IsPeriod** | Pointer to **NullableBool** | True when the company did something across an interval (a trial). Independent of date_precision. | [optional] 
**SpanEndDate** | Pointer to **NullableString** | The end of a real period; null for a non-period row and for an open-ended period. | [optional] 
**SpanEndIsProjected** | Pointer to **NullableBool** | True when the end (or the event itself, for a decision still ahead) is planned rather than realized. | [optional] 
**DisplayLabel** | Pointer to **NullableString** | Presentation label for the event, not a stable contract. | [optional] 
**DetailText** | Pointer to **NullableString** | Presentation detail, including any amount as formatted text. | [optional] 
**LinkKind** | Pointer to **NullableString** | What link_ref points at: trial, funding_event, patent, external or none. | [optional] 
**LinkRef** | Pointer to **NullableString** | The natural id of the linked thing, or an absolute URL when link_kind is external; null exactly when link_kind is none. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The tier vocabulary is none or licensable. Every row here is none: every lane is an open source, and a data test fails if a restricted lane is ever admitted. | [optional] 
**SourceLicense** | Pointer to **NullableString** | The licence of the lane&#39;s source. | [optional] 
**CommercialUseAllowed** | Pointer to **NullableBool** | Whether the lane&#39;s licence permits commercial use. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit line the lane&#39;s source requires; show it wherever the event is shown. Null when none is owed. | [optional] 
**SourceLabel** | Pointer to **NullableString** | The lane&#39;s display name. | [optional] 

## Methods

### NewCompanyTimelineRow

`func NewCompanyTimelineRow() *CompanyTimelineRow`

NewCompanyTimelineRow instantiates a new CompanyTimelineRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyTimelineRowWithDefaults

`func NewCompanyTimelineRowWithDefaults() *CompanyTimelineRow`

NewCompanyTimelineRowWithDefaults instantiates a new CompanyTimelineRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimelineEventId

`func (o *CompanyTimelineRow) GetTimelineEventId() string`

GetTimelineEventId returns the TimelineEventId field if non-nil, zero value otherwise.

### GetTimelineEventIdOk

`func (o *CompanyTimelineRow) GetTimelineEventIdOk() (*string, bool)`

GetTimelineEventIdOk returns a tuple with the TimelineEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineEventId

`func (o *CompanyTimelineRow) SetTimelineEventId(v string)`

SetTimelineEventId sets TimelineEventId field to given value.

### HasTimelineEventId

`func (o *CompanyTimelineRow) HasTimelineEventId() bool`

HasTimelineEventId returns a boolean if a field has been set.

### SetTimelineEventIdNil

`func (o *CompanyTimelineRow) SetTimelineEventIdNil(b bool)`

 SetTimelineEventIdNil sets the value for TimelineEventId to be an explicit nil

### UnsetTimelineEventId
`func (o *CompanyTimelineRow) UnsetTimelineEventId()`

UnsetTimelineEventId ensures that no value is present for TimelineEventId, not even an explicit nil
### GetCompanyId

`func (o *CompanyTimelineRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyTimelineRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyTimelineRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CompanyTimelineRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CompanyTimelineRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CompanyTimelineRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *CompanyTimelineRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyTimelineRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyTimelineRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyTimelineRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyTimelineRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyTimelineRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *CompanyTimelineRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CompanyTimelineRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CompanyTimelineRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CompanyTimelineRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CompanyTimelineRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CompanyTimelineRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *CompanyTimelineRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CompanyTimelineRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CompanyTimelineRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CompanyTimelineRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CompanyTimelineRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CompanyTimelineRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetReportingOrgId

`func (o *CompanyTimelineRow) GetReportingOrgId() string`

GetReportingOrgId returns the ReportingOrgId field if non-nil, zero value otherwise.

### GetReportingOrgIdOk

`func (o *CompanyTimelineRow) GetReportingOrgIdOk() (*string, bool)`

GetReportingOrgIdOk returns a tuple with the ReportingOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOrgId

`func (o *CompanyTimelineRow) SetReportingOrgId(v string)`

SetReportingOrgId sets ReportingOrgId field to given value.

### HasReportingOrgId

`func (o *CompanyTimelineRow) HasReportingOrgId() bool`

HasReportingOrgId returns a boolean if a field has been set.

### SetReportingOrgIdNil

`func (o *CompanyTimelineRow) SetReportingOrgIdNil(b bool)`

 SetReportingOrgIdNil sets the value for ReportingOrgId to be an explicit nil

### UnsetReportingOrgId
`func (o *CompanyTimelineRow) UnsetReportingOrgId()`

UnsetReportingOrgId ensures that no value is present for ReportingOrgId, not even an explicit nil
### GetSourceKey

`func (o *CompanyTimelineRow) GetSourceKey() string`

GetSourceKey returns the SourceKey field if non-nil, zero value otherwise.

### GetSourceKeyOk

`func (o *CompanyTimelineRow) GetSourceKeyOk() (*string, bool)`

GetSourceKeyOk returns a tuple with the SourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKey

`func (o *CompanyTimelineRow) SetSourceKey(v string)`

SetSourceKey sets SourceKey field to given value.

### HasSourceKey

`func (o *CompanyTimelineRow) HasSourceKey() bool`

HasSourceKey returns a boolean if a field has been set.

### SetSourceKeyNil

`func (o *CompanyTimelineRow) SetSourceKeyNil(b bool)`

 SetSourceKeyNil sets the value for SourceKey to be an explicit nil

### UnsetSourceKey
`func (o *CompanyTimelineRow) UnsetSourceKey()`

UnsetSourceKey ensures that no value is present for SourceKey, not even an explicit nil
### GetSourceEntityKey

`func (o *CompanyTimelineRow) GetSourceEntityKey() string`

GetSourceEntityKey returns the SourceEntityKey field if non-nil, zero value otherwise.

### GetSourceEntityKeyOk

`func (o *CompanyTimelineRow) GetSourceEntityKeyOk() (*string, bool)`

GetSourceEntityKeyOk returns a tuple with the SourceEntityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityKey

`func (o *CompanyTimelineRow) SetSourceEntityKey(v string)`

SetSourceEntityKey sets SourceEntityKey field to given value.

### HasSourceEntityKey

`func (o *CompanyTimelineRow) HasSourceEntityKey() bool`

HasSourceEntityKey returns a boolean if a field has been set.

### SetSourceEntityKeyNil

`func (o *CompanyTimelineRow) SetSourceEntityKeyNil(b bool)`

 SetSourceEntityKeyNil sets the value for SourceEntityKey to be an explicit nil

### UnsetSourceEntityKey
`func (o *CompanyTimelineRow) UnsetSourceEntityKey()`

UnsetSourceEntityKey ensures that no value is present for SourceEntityKey, not even an explicit nil
### GetEventType

`func (o *CompanyTimelineRow) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CompanyTimelineRow) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CompanyTimelineRow) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CompanyTimelineRow) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *CompanyTimelineRow) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *CompanyTimelineRow) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetEventFamily

`func (o *CompanyTimelineRow) GetEventFamily() string`

GetEventFamily returns the EventFamily field if non-nil, zero value otherwise.

### GetEventFamilyOk

`func (o *CompanyTimelineRow) GetEventFamilyOk() (*string, bool)`

GetEventFamilyOk returns a tuple with the EventFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFamily

`func (o *CompanyTimelineRow) SetEventFamily(v string)`

SetEventFamily sets EventFamily field to given value.

### HasEventFamily

`func (o *CompanyTimelineRow) HasEventFamily() bool`

HasEventFamily returns a boolean if a field has been set.

### SetEventFamilyNil

`func (o *CompanyTimelineRow) SetEventFamilyNil(b bool)`

 SetEventFamilyNil sets the value for EventFamily to be an explicit nil

### UnsetEventFamily
`func (o *CompanyTimelineRow) UnsetEventFamily()`

UnsetEventFamily ensures that no value is present for EventFamily, not even an explicit nil
### GetEventDate

`func (o *CompanyTimelineRow) GetEventDate() string`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *CompanyTimelineRow) GetEventDateOk() (*string, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *CompanyTimelineRow) SetEventDate(v string)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *CompanyTimelineRow) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### SetEventDateNil

`func (o *CompanyTimelineRow) SetEventDateNil(b bool)`

 SetEventDateNil sets the value for EventDate to be an explicit nil

### UnsetEventDate
`func (o *CompanyTimelineRow) UnsetEventDate()`

UnsetEventDate ensures that no value is present for EventDate, not even an explicit nil
### GetDatePrecision

`func (o *CompanyTimelineRow) GetDatePrecision() string`

GetDatePrecision returns the DatePrecision field if non-nil, zero value otherwise.

### GetDatePrecisionOk

`func (o *CompanyTimelineRow) GetDatePrecisionOk() (*string, bool)`

GetDatePrecisionOk returns a tuple with the DatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePrecision

`func (o *CompanyTimelineRow) SetDatePrecision(v string)`

SetDatePrecision sets DatePrecision field to given value.

### HasDatePrecision

`func (o *CompanyTimelineRow) HasDatePrecision() bool`

HasDatePrecision returns a boolean if a field has been set.

### SetDatePrecisionNil

`func (o *CompanyTimelineRow) SetDatePrecisionNil(b bool)`

 SetDatePrecisionNil sets the value for DatePrecision to be an explicit nil

### UnsetDatePrecision
`func (o *CompanyTimelineRow) UnsetDatePrecision()`

UnsetDatePrecision ensures that no value is present for DatePrecision, not even an explicit nil
### GetPrecisionWindowEnd

`func (o *CompanyTimelineRow) GetPrecisionWindowEnd() string`

GetPrecisionWindowEnd returns the PrecisionWindowEnd field if non-nil, zero value otherwise.

### GetPrecisionWindowEndOk

`func (o *CompanyTimelineRow) GetPrecisionWindowEndOk() (*string, bool)`

GetPrecisionWindowEndOk returns a tuple with the PrecisionWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionWindowEnd

`func (o *CompanyTimelineRow) SetPrecisionWindowEnd(v string)`

SetPrecisionWindowEnd sets PrecisionWindowEnd field to given value.

### HasPrecisionWindowEnd

`func (o *CompanyTimelineRow) HasPrecisionWindowEnd() bool`

HasPrecisionWindowEnd returns a boolean if a field has been set.

### SetPrecisionWindowEndNil

`func (o *CompanyTimelineRow) SetPrecisionWindowEndNil(b bool)`

 SetPrecisionWindowEndNil sets the value for PrecisionWindowEnd to be an explicit nil

### UnsetPrecisionWindowEnd
`func (o *CompanyTimelineRow) UnsetPrecisionWindowEnd()`

UnsetPrecisionWindowEnd ensures that no value is present for PrecisionWindowEnd, not even an explicit nil
### GetDatePrecisionBasis

`func (o *CompanyTimelineRow) GetDatePrecisionBasis() string`

GetDatePrecisionBasis returns the DatePrecisionBasis field if non-nil, zero value otherwise.

### GetDatePrecisionBasisOk

`func (o *CompanyTimelineRow) GetDatePrecisionBasisOk() (*string, bool)`

GetDatePrecisionBasisOk returns a tuple with the DatePrecisionBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePrecisionBasis

`func (o *CompanyTimelineRow) SetDatePrecisionBasis(v string)`

SetDatePrecisionBasis sets DatePrecisionBasis field to given value.

### HasDatePrecisionBasis

`func (o *CompanyTimelineRow) HasDatePrecisionBasis() bool`

HasDatePrecisionBasis returns a boolean if a field has been set.

### SetDatePrecisionBasisNil

`func (o *CompanyTimelineRow) SetDatePrecisionBasisNil(b bool)`

 SetDatePrecisionBasisNil sets the value for DatePrecisionBasis to be an explicit nil

### UnsetDatePrecisionBasis
`func (o *CompanyTimelineRow) UnsetDatePrecisionBasis()`

UnsetDatePrecisionBasis ensures that no value is present for DatePrecisionBasis, not even an explicit nil
### GetIsPeriod

`func (o *CompanyTimelineRow) GetIsPeriod() bool`

GetIsPeriod returns the IsPeriod field if non-nil, zero value otherwise.

### GetIsPeriodOk

`func (o *CompanyTimelineRow) GetIsPeriodOk() (*bool, bool)`

GetIsPeriodOk returns a tuple with the IsPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPeriod

`func (o *CompanyTimelineRow) SetIsPeriod(v bool)`

SetIsPeriod sets IsPeriod field to given value.

### HasIsPeriod

`func (o *CompanyTimelineRow) HasIsPeriod() bool`

HasIsPeriod returns a boolean if a field has been set.

### SetIsPeriodNil

`func (o *CompanyTimelineRow) SetIsPeriodNil(b bool)`

 SetIsPeriodNil sets the value for IsPeriod to be an explicit nil

### UnsetIsPeriod
`func (o *CompanyTimelineRow) UnsetIsPeriod()`

UnsetIsPeriod ensures that no value is present for IsPeriod, not even an explicit nil
### GetSpanEndDate

`func (o *CompanyTimelineRow) GetSpanEndDate() string`

GetSpanEndDate returns the SpanEndDate field if non-nil, zero value otherwise.

### GetSpanEndDateOk

`func (o *CompanyTimelineRow) GetSpanEndDateOk() (*string, bool)`

GetSpanEndDateOk returns a tuple with the SpanEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanEndDate

`func (o *CompanyTimelineRow) SetSpanEndDate(v string)`

SetSpanEndDate sets SpanEndDate field to given value.

### HasSpanEndDate

`func (o *CompanyTimelineRow) HasSpanEndDate() bool`

HasSpanEndDate returns a boolean if a field has been set.

### SetSpanEndDateNil

`func (o *CompanyTimelineRow) SetSpanEndDateNil(b bool)`

 SetSpanEndDateNil sets the value for SpanEndDate to be an explicit nil

### UnsetSpanEndDate
`func (o *CompanyTimelineRow) UnsetSpanEndDate()`

UnsetSpanEndDate ensures that no value is present for SpanEndDate, not even an explicit nil
### GetSpanEndIsProjected

`func (o *CompanyTimelineRow) GetSpanEndIsProjected() bool`

GetSpanEndIsProjected returns the SpanEndIsProjected field if non-nil, zero value otherwise.

### GetSpanEndIsProjectedOk

`func (o *CompanyTimelineRow) GetSpanEndIsProjectedOk() (*bool, bool)`

GetSpanEndIsProjectedOk returns a tuple with the SpanEndIsProjected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanEndIsProjected

`func (o *CompanyTimelineRow) SetSpanEndIsProjected(v bool)`

SetSpanEndIsProjected sets SpanEndIsProjected field to given value.

### HasSpanEndIsProjected

`func (o *CompanyTimelineRow) HasSpanEndIsProjected() bool`

HasSpanEndIsProjected returns a boolean if a field has been set.

### SetSpanEndIsProjectedNil

`func (o *CompanyTimelineRow) SetSpanEndIsProjectedNil(b bool)`

 SetSpanEndIsProjectedNil sets the value for SpanEndIsProjected to be an explicit nil

### UnsetSpanEndIsProjected
`func (o *CompanyTimelineRow) UnsetSpanEndIsProjected()`

UnsetSpanEndIsProjected ensures that no value is present for SpanEndIsProjected, not even an explicit nil
### GetDisplayLabel

`func (o *CompanyTimelineRow) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *CompanyTimelineRow) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *CompanyTimelineRow) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *CompanyTimelineRow) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### SetDisplayLabelNil

`func (o *CompanyTimelineRow) SetDisplayLabelNil(b bool)`

 SetDisplayLabelNil sets the value for DisplayLabel to be an explicit nil

### UnsetDisplayLabel
`func (o *CompanyTimelineRow) UnsetDisplayLabel()`

UnsetDisplayLabel ensures that no value is present for DisplayLabel, not even an explicit nil
### GetDetailText

`func (o *CompanyTimelineRow) GetDetailText() string`

GetDetailText returns the DetailText field if non-nil, zero value otherwise.

### GetDetailTextOk

`func (o *CompanyTimelineRow) GetDetailTextOk() (*string, bool)`

GetDetailTextOk returns a tuple with the DetailText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailText

`func (o *CompanyTimelineRow) SetDetailText(v string)`

SetDetailText sets DetailText field to given value.

### HasDetailText

`func (o *CompanyTimelineRow) HasDetailText() bool`

HasDetailText returns a boolean if a field has been set.

### SetDetailTextNil

`func (o *CompanyTimelineRow) SetDetailTextNil(b bool)`

 SetDetailTextNil sets the value for DetailText to be an explicit nil

### UnsetDetailText
`func (o *CompanyTimelineRow) UnsetDetailText()`

UnsetDetailText ensures that no value is present for DetailText, not even an explicit nil
### GetLinkKind

`func (o *CompanyTimelineRow) GetLinkKind() string`

GetLinkKind returns the LinkKind field if non-nil, zero value otherwise.

### GetLinkKindOk

`func (o *CompanyTimelineRow) GetLinkKindOk() (*string, bool)`

GetLinkKindOk returns a tuple with the LinkKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkKind

`func (o *CompanyTimelineRow) SetLinkKind(v string)`

SetLinkKind sets LinkKind field to given value.

### HasLinkKind

`func (o *CompanyTimelineRow) HasLinkKind() bool`

HasLinkKind returns a boolean if a field has been set.

### SetLinkKindNil

`func (o *CompanyTimelineRow) SetLinkKindNil(b bool)`

 SetLinkKindNil sets the value for LinkKind to be an explicit nil

### UnsetLinkKind
`func (o *CompanyTimelineRow) UnsetLinkKind()`

UnsetLinkKind ensures that no value is present for LinkKind, not even an explicit nil
### GetLinkRef

`func (o *CompanyTimelineRow) GetLinkRef() string`

GetLinkRef returns the LinkRef field if non-nil, zero value otherwise.

### GetLinkRefOk

`func (o *CompanyTimelineRow) GetLinkRefOk() (*string, bool)`

GetLinkRefOk returns a tuple with the LinkRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRef

`func (o *CompanyTimelineRow) SetLinkRef(v string)`

SetLinkRef sets LinkRef field to given value.

### HasLinkRef

`func (o *CompanyTimelineRow) HasLinkRef() bool`

HasLinkRef returns a boolean if a field has been set.

### SetLinkRefNil

`func (o *CompanyTimelineRow) SetLinkRefNil(b bool)`

 SetLinkRefNil sets the value for LinkRef to be an explicit nil

### UnsetLinkRef
`func (o *CompanyTimelineRow) UnsetLinkRef()`

UnsetLinkRef ensures that no value is present for LinkRef, not even an explicit nil
### GetRestrictionTier

`func (o *CompanyTimelineRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *CompanyTimelineRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *CompanyTimelineRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *CompanyTimelineRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *CompanyTimelineRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *CompanyTimelineRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetSourceLicense

`func (o *CompanyTimelineRow) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *CompanyTimelineRow) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *CompanyTimelineRow) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.

### HasSourceLicense

`func (o *CompanyTimelineRow) HasSourceLicense() bool`

HasSourceLicense returns a boolean if a field has been set.

### SetSourceLicenseNil

`func (o *CompanyTimelineRow) SetSourceLicenseNil(b bool)`

 SetSourceLicenseNil sets the value for SourceLicense to be an explicit nil

### UnsetSourceLicense
`func (o *CompanyTimelineRow) UnsetSourceLicense()`

UnsetSourceLicense ensures that no value is present for SourceLicense, not even an explicit nil
### GetCommercialUseAllowed

`func (o *CompanyTimelineRow) GetCommercialUseAllowed() bool`

GetCommercialUseAllowed returns the CommercialUseAllowed field if non-nil, zero value otherwise.

### GetCommercialUseAllowedOk

`func (o *CompanyTimelineRow) GetCommercialUseAllowedOk() (*bool, bool)`

GetCommercialUseAllowedOk returns a tuple with the CommercialUseAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialUseAllowed

`func (o *CompanyTimelineRow) SetCommercialUseAllowed(v bool)`

SetCommercialUseAllowed sets CommercialUseAllowed field to given value.

### HasCommercialUseAllowed

`func (o *CompanyTimelineRow) HasCommercialUseAllowed() bool`

HasCommercialUseAllowed returns a boolean if a field has been set.

### SetCommercialUseAllowedNil

`func (o *CompanyTimelineRow) SetCommercialUseAllowedNil(b bool)`

 SetCommercialUseAllowedNil sets the value for CommercialUseAllowed to be an explicit nil

### UnsetCommercialUseAllowed
`func (o *CompanyTimelineRow) UnsetCommercialUseAllowed()`

UnsetCommercialUseAllowed ensures that no value is present for CommercialUseAllowed, not even an explicit nil
### GetAttribution

`func (o *CompanyTimelineRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *CompanyTimelineRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *CompanyTimelineRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *CompanyTimelineRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *CompanyTimelineRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *CompanyTimelineRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetSourceLabel

`func (o *CompanyTimelineRow) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *CompanyTimelineRow) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *CompanyTimelineRow) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.

### HasSourceLabel

`func (o *CompanyTimelineRow) HasSourceLabel() bool`

HasSourceLabel returns a boolean if a field has been set.

### SetSourceLabelNil

`func (o *CompanyTimelineRow) SetSourceLabelNil(b bool)`

 SetSourceLabelNil sets the value for SourceLabel to be an explicit nil

### UnsetSourceLabel
`func (o *CompanyTimelineRow) UnsetSourceLabel()`

UnsetSourceLabel ensures that no value is present for SourceLabel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


