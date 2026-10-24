# EnforcementEventsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **NullableString** | The event id: a hash of the lane and the lane&#39;s own key, unique across lanes. | [optional] 
**Source** | Pointer to **NullableString** | The lane the event came from: inspection_483 (an FDA inspection), import_refusal, warning_letter, untitled_letter, debarment, recall or aphis_inspection (a USDA APHIS Animal Welfare Act inspection report). | [optional] 
**EventDate** | Pointer to **NullableString** | The event date: the inspection date, refusal date, letter issue date (posted date when there is none), debarment effective date or recall initiation date. | [optional] 
**Domain** | Pointer to **NullableString** | The regulatory domain of the event, derived from the lane&#39;s own classifier: drug, biologic, device, animal_research (the APHIS lane), veterinary, food, cosmetic, tobacco or other. | [optional] 
**IsBiomedical** | Pointer to **NullableBool** | True when domain is drug, biologic or device. The APHIS lane is false here on purpose; the page&#39;s default view adds domain animal_research to this set. | [optional] 
**Headline** | Pointer to **NullableString** | A one-line description: the inspection product type, the refused product, the letter subject, the debarment type or the recalled product. | [optional] 
**Detail** | Pointer to **NullableString** | The lane&#39;s supporting text: the cited regulations of an inspection, the refusal charges, the issuing office of a warning letter, the debarment term and citation, or the recall reason. | [optional] 
**Classification** | Pointer to **NullableString** | The inspection outcome (NAI no action indicated, VAI voluntary action indicated, OAI official action indicated); null for every other lane. | [optional] 
**RecallClass** | Pointer to **NullableString** | The recall severity class (Class I, II or III); null for every other lane. | [optional] 
**Had483** | Pointer to **NullableBool** | True when the inspection ended with a Form 483 of observations; null for every other lane. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The regulator&#39;s page for the event, where the lane has one. | [optional] 
**SourceRef** | Pointer to **NullableString** | The lane&#39;s own reference, where it has one: the letter id, debarment id or recall number. | [optional] 
**DeviceClass** | Pointer to **NullableString** | The FDA device class of an import-refusal product, from the device classification dimension; null for every other lane. | [optional] 
**DeviceSpecialty** | Pointer to **NullableString** | The medical specialty of an import-refusal device product; null for every other lane. | [optional] 
**IsLowSignalDevice** | Pointer to **NullableBool** | True for a device event the biomedical view hides by default (a low-signal commodity class). | [optional] 
**AwaOutcome** | Pointer to **NullableString** | The APHIS report&#39;s verdict, the strongest citation kind on it: direct, critical, non_critical, no_citations or no_citations_no_animals; null for every FDA lane. | [optional] 
**AwaInspectionType** | Pointer to **NullableString** | The inspection type as printed on the APHIS report (for example ROUTINE INSPECTION); null for every FDA lane. | [optional] 
**AwaNCitations** | Pointer to **NullableInt32** | Citations on the APHIS report; null for every FDA lane. Read with awa_species_scope. | [optional] 
**AwaNDirect** | Pointer to **NullableInt32** | Of the report&#39;s citations, those cited as direct; null for every FDA lane. | [optional] 
**AwaNCritical** | Pointer to **NullableInt32** | Of the report&#39;s citations, those cited as critical; null for every FDA lane. | [optional] 
**AwaNNonCritical** | Pointer to **NullableInt32** | Of the report&#39;s citations, those cited as non-critical; null for every FDA lane. | [optional] 
**AwaNRepeatCitations** | Pointer to **NullableInt32** | Of the report&#39;s citations, those it marks as repeats; null for every FDA lane. | [optional] 
**AwaSpeciesSummary** | Pointer to **NullableString** | The species the APHIS report printed with their animal counts, as \&quot;Common name count; ...\&quot; in the report&#39;s order; null on an FDA lane or a report that printed none. | [optional] 
**AwaSpeciesScope** | Pointer to **NullableString** | The label AWA-covered species on every APHIS row; null for every FDA lane. The awa_* counts describe the species the Animal Welfare Act covers; read them with this label. | [optional] 
**FirmName** | Pointer to **NullableString** | The firm&#39;s name as filed with the regulator (the party named on the record, an individual for an individual debarment). | [optional] 
**FirmOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] 
**FirmCompanyId** | Pointer to **NullableString** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] 
**FirmCompanyName** | Pointer to **NullableString** | Display name of the firm&#39;s company. | [optional] 
**FirmGroupId** | Pointer to **NullableString** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**FirmGroupName** | Pointer to **NullableString** | Display name of the firm&#39;s group. | [optional] 
**MatchTier** | Pointer to **NullableString** | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous or unresolved. | [optional] 
**OrgResolutionMethod** | Pointer to **NullableString** | The route that attached the firm: fei_bridge (the FDA establishment identifier bridge), fei_firm_mint or fei_firm_fold (the operating firm&#39;s registry record), aphis_bridge, aphis_firm_mint or aphis_firm_fold (the APHIS registrant certificate bridge and its operating-firm record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional] 
**IsOrgResolved** | Pointer to **NullableBool** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] 

## Methods

### NewEnforcementEventsRow

`func NewEnforcementEventsRow() *EnforcementEventsRow`

NewEnforcementEventsRow instantiates a new EnforcementEventsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnforcementEventsRowWithDefaults

`func NewEnforcementEventsRowWithDefaults() *EnforcementEventsRow`

NewEnforcementEventsRowWithDefaults instantiates a new EnforcementEventsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *EnforcementEventsRow) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EnforcementEventsRow) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EnforcementEventsRow) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *EnforcementEventsRow) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventIdNil

`func (o *EnforcementEventsRow) SetEventIdNil(b bool)`

 SetEventIdNil sets the value for EventId to be an explicit nil

### UnsetEventId
`func (o *EnforcementEventsRow) UnsetEventId()`

UnsetEventId ensures that no value is present for EventId, not even an explicit nil
### GetSource

`func (o *EnforcementEventsRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EnforcementEventsRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EnforcementEventsRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EnforcementEventsRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *EnforcementEventsRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *EnforcementEventsRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetEventDate

`func (o *EnforcementEventsRow) GetEventDate() string`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *EnforcementEventsRow) GetEventDateOk() (*string, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *EnforcementEventsRow) SetEventDate(v string)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *EnforcementEventsRow) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### SetEventDateNil

`func (o *EnforcementEventsRow) SetEventDateNil(b bool)`

 SetEventDateNil sets the value for EventDate to be an explicit nil

### UnsetEventDate
`func (o *EnforcementEventsRow) UnsetEventDate()`

UnsetEventDate ensures that no value is present for EventDate, not even an explicit nil
### GetDomain

`func (o *EnforcementEventsRow) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EnforcementEventsRow) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EnforcementEventsRow) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EnforcementEventsRow) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *EnforcementEventsRow) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *EnforcementEventsRow) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetIsBiomedical

`func (o *EnforcementEventsRow) GetIsBiomedical() bool`

GetIsBiomedical returns the IsBiomedical field if non-nil, zero value otherwise.

### GetIsBiomedicalOk

`func (o *EnforcementEventsRow) GetIsBiomedicalOk() (*bool, bool)`

GetIsBiomedicalOk returns a tuple with the IsBiomedical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBiomedical

`func (o *EnforcementEventsRow) SetIsBiomedical(v bool)`

SetIsBiomedical sets IsBiomedical field to given value.

### HasIsBiomedical

`func (o *EnforcementEventsRow) HasIsBiomedical() bool`

HasIsBiomedical returns a boolean if a field has been set.

### SetIsBiomedicalNil

`func (o *EnforcementEventsRow) SetIsBiomedicalNil(b bool)`

 SetIsBiomedicalNil sets the value for IsBiomedical to be an explicit nil

### UnsetIsBiomedical
`func (o *EnforcementEventsRow) UnsetIsBiomedical()`

UnsetIsBiomedical ensures that no value is present for IsBiomedical, not even an explicit nil
### GetHeadline

`func (o *EnforcementEventsRow) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *EnforcementEventsRow) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *EnforcementEventsRow) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *EnforcementEventsRow) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### SetHeadlineNil

`func (o *EnforcementEventsRow) SetHeadlineNil(b bool)`

 SetHeadlineNil sets the value for Headline to be an explicit nil

### UnsetHeadline
`func (o *EnforcementEventsRow) UnsetHeadline()`

UnsetHeadline ensures that no value is present for Headline, not even an explicit nil
### GetDetail

`func (o *EnforcementEventsRow) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *EnforcementEventsRow) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *EnforcementEventsRow) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *EnforcementEventsRow) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *EnforcementEventsRow) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *EnforcementEventsRow) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetClassification

`func (o *EnforcementEventsRow) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *EnforcementEventsRow) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *EnforcementEventsRow) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *EnforcementEventsRow) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *EnforcementEventsRow) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *EnforcementEventsRow) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetRecallClass

`func (o *EnforcementEventsRow) GetRecallClass() string`

GetRecallClass returns the RecallClass field if non-nil, zero value otherwise.

### GetRecallClassOk

`func (o *EnforcementEventsRow) GetRecallClassOk() (*string, bool)`

GetRecallClassOk returns a tuple with the RecallClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecallClass

`func (o *EnforcementEventsRow) SetRecallClass(v string)`

SetRecallClass sets RecallClass field to given value.

### HasRecallClass

`func (o *EnforcementEventsRow) HasRecallClass() bool`

HasRecallClass returns a boolean if a field has been set.

### SetRecallClassNil

`func (o *EnforcementEventsRow) SetRecallClassNil(b bool)`

 SetRecallClassNil sets the value for RecallClass to be an explicit nil

### UnsetRecallClass
`func (o *EnforcementEventsRow) UnsetRecallClass()`

UnsetRecallClass ensures that no value is present for RecallClass, not even an explicit nil
### GetHad483

`func (o *EnforcementEventsRow) GetHad483() bool`

GetHad483 returns the Had483 field if non-nil, zero value otherwise.

### GetHad483Ok

`func (o *EnforcementEventsRow) GetHad483Ok() (*bool, bool)`

GetHad483Ok returns a tuple with the Had483 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHad483

`func (o *EnforcementEventsRow) SetHad483(v bool)`

SetHad483 sets Had483 field to given value.

### HasHad483

`func (o *EnforcementEventsRow) HasHad483() bool`

HasHad483 returns a boolean if a field has been set.

### SetHad483Nil

`func (o *EnforcementEventsRow) SetHad483Nil(b bool)`

 SetHad483Nil sets the value for Had483 to be an explicit nil

### UnsetHad483
`func (o *EnforcementEventsRow) UnsetHad483()`

UnsetHad483 ensures that no value is present for Had483, not even an explicit nil
### GetSourceUrl

`func (o *EnforcementEventsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *EnforcementEventsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *EnforcementEventsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *EnforcementEventsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *EnforcementEventsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *EnforcementEventsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceRef

`func (o *EnforcementEventsRow) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *EnforcementEventsRow) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *EnforcementEventsRow) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *EnforcementEventsRow) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### SetSourceRefNil

`func (o *EnforcementEventsRow) SetSourceRefNil(b bool)`

 SetSourceRefNil sets the value for SourceRef to be an explicit nil

### UnsetSourceRef
`func (o *EnforcementEventsRow) UnsetSourceRef()`

UnsetSourceRef ensures that no value is present for SourceRef, not even an explicit nil
### GetDeviceClass

`func (o *EnforcementEventsRow) GetDeviceClass() string`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *EnforcementEventsRow) GetDeviceClassOk() (*string, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *EnforcementEventsRow) SetDeviceClass(v string)`

SetDeviceClass sets DeviceClass field to given value.

### HasDeviceClass

`func (o *EnforcementEventsRow) HasDeviceClass() bool`

HasDeviceClass returns a boolean if a field has been set.

### SetDeviceClassNil

`func (o *EnforcementEventsRow) SetDeviceClassNil(b bool)`

 SetDeviceClassNil sets the value for DeviceClass to be an explicit nil

### UnsetDeviceClass
`func (o *EnforcementEventsRow) UnsetDeviceClass()`

UnsetDeviceClass ensures that no value is present for DeviceClass, not even an explicit nil
### GetDeviceSpecialty

`func (o *EnforcementEventsRow) GetDeviceSpecialty() string`

GetDeviceSpecialty returns the DeviceSpecialty field if non-nil, zero value otherwise.

### GetDeviceSpecialtyOk

`func (o *EnforcementEventsRow) GetDeviceSpecialtyOk() (*string, bool)`

GetDeviceSpecialtyOk returns a tuple with the DeviceSpecialty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSpecialty

`func (o *EnforcementEventsRow) SetDeviceSpecialty(v string)`

SetDeviceSpecialty sets DeviceSpecialty field to given value.

### HasDeviceSpecialty

`func (o *EnforcementEventsRow) HasDeviceSpecialty() bool`

HasDeviceSpecialty returns a boolean if a field has been set.

### SetDeviceSpecialtyNil

`func (o *EnforcementEventsRow) SetDeviceSpecialtyNil(b bool)`

 SetDeviceSpecialtyNil sets the value for DeviceSpecialty to be an explicit nil

### UnsetDeviceSpecialty
`func (o *EnforcementEventsRow) UnsetDeviceSpecialty()`

UnsetDeviceSpecialty ensures that no value is present for DeviceSpecialty, not even an explicit nil
### GetIsLowSignalDevice

`func (o *EnforcementEventsRow) GetIsLowSignalDevice() bool`

GetIsLowSignalDevice returns the IsLowSignalDevice field if non-nil, zero value otherwise.

### GetIsLowSignalDeviceOk

`func (o *EnforcementEventsRow) GetIsLowSignalDeviceOk() (*bool, bool)`

GetIsLowSignalDeviceOk returns a tuple with the IsLowSignalDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLowSignalDevice

`func (o *EnforcementEventsRow) SetIsLowSignalDevice(v bool)`

SetIsLowSignalDevice sets IsLowSignalDevice field to given value.

### HasIsLowSignalDevice

`func (o *EnforcementEventsRow) HasIsLowSignalDevice() bool`

HasIsLowSignalDevice returns a boolean if a field has been set.

### SetIsLowSignalDeviceNil

`func (o *EnforcementEventsRow) SetIsLowSignalDeviceNil(b bool)`

 SetIsLowSignalDeviceNil sets the value for IsLowSignalDevice to be an explicit nil

### UnsetIsLowSignalDevice
`func (o *EnforcementEventsRow) UnsetIsLowSignalDevice()`

UnsetIsLowSignalDevice ensures that no value is present for IsLowSignalDevice, not even an explicit nil
### GetAwaOutcome

`func (o *EnforcementEventsRow) GetAwaOutcome() string`

GetAwaOutcome returns the AwaOutcome field if non-nil, zero value otherwise.

### GetAwaOutcomeOk

`func (o *EnforcementEventsRow) GetAwaOutcomeOk() (*string, bool)`

GetAwaOutcomeOk returns a tuple with the AwaOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaOutcome

`func (o *EnforcementEventsRow) SetAwaOutcome(v string)`

SetAwaOutcome sets AwaOutcome field to given value.

### HasAwaOutcome

`func (o *EnforcementEventsRow) HasAwaOutcome() bool`

HasAwaOutcome returns a boolean if a field has been set.

### SetAwaOutcomeNil

`func (o *EnforcementEventsRow) SetAwaOutcomeNil(b bool)`

 SetAwaOutcomeNil sets the value for AwaOutcome to be an explicit nil

### UnsetAwaOutcome
`func (o *EnforcementEventsRow) UnsetAwaOutcome()`

UnsetAwaOutcome ensures that no value is present for AwaOutcome, not even an explicit nil
### GetAwaInspectionType

`func (o *EnforcementEventsRow) GetAwaInspectionType() string`

GetAwaInspectionType returns the AwaInspectionType field if non-nil, zero value otherwise.

### GetAwaInspectionTypeOk

`func (o *EnforcementEventsRow) GetAwaInspectionTypeOk() (*string, bool)`

GetAwaInspectionTypeOk returns a tuple with the AwaInspectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaInspectionType

`func (o *EnforcementEventsRow) SetAwaInspectionType(v string)`

SetAwaInspectionType sets AwaInspectionType field to given value.

### HasAwaInspectionType

`func (o *EnforcementEventsRow) HasAwaInspectionType() bool`

HasAwaInspectionType returns a boolean if a field has been set.

### SetAwaInspectionTypeNil

`func (o *EnforcementEventsRow) SetAwaInspectionTypeNil(b bool)`

 SetAwaInspectionTypeNil sets the value for AwaInspectionType to be an explicit nil

### UnsetAwaInspectionType
`func (o *EnforcementEventsRow) UnsetAwaInspectionType()`

UnsetAwaInspectionType ensures that no value is present for AwaInspectionType, not even an explicit nil
### GetAwaNCitations

`func (o *EnforcementEventsRow) GetAwaNCitations() int32`

GetAwaNCitations returns the AwaNCitations field if non-nil, zero value otherwise.

### GetAwaNCitationsOk

`func (o *EnforcementEventsRow) GetAwaNCitationsOk() (*int32, bool)`

GetAwaNCitationsOk returns a tuple with the AwaNCitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaNCitations

`func (o *EnforcementEventsRow) SetAwaNCitations(v int32)`

SetAwaNCitations sets AwaNCitations field to given value.

### HasAwaNCitations

`func (o *EnforcementEventsRow) HasAwaNCitations() bool`

HasAwaNCitations returns a boolean if a field has been set.

### SetAwaNCitationsNil

`func (o *EnforcementEventsRow) SetAwaNCitationsNil(b bool)`

 SetAwaNCitationsNil sets the value for AwaNCitations to be an explicit nil

### UnsetAwaNCitations
`func (o *EnforcementEventsRow) UnsetAwaNCitations()`

UnsetAwaNCitations ensures that no value is present for AwaNCitations, not even an explicit nil
### GetAwaNDirect

`func (o *EnforcementEventsRow) GetAwaNDirect() int32`

GetAwaNDirect returns the AwaNDirect field if non-nil, zero value otherwise.

### GetAwaNDirectOk

`func (o *EnforcementEventsRow) GetAwaNDirectOk() (*int32, bool)`

GetAwaNDirectOk returns a tuple with the AwaNDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaNDirect

`func (o *EnforcementEventsRow) SetAwaNDirect(v int32)`

SetAwaNDirect sets AwaNDirect field to given value.

### HasAwaNDirect

`func (o *EnforcementEventsRow) HasAwaNDirect() bool`

HasAwaNDirect returns a boolean if a field has been set.

### SetAwaNDirectNil

`func (o *EnforcementEventsRow) SetAwaNDirectNil(b bool)`

 SetAwaNDirectNil sets the value for AwaNDirect to be an explicit nil

### UnsetAwaNDirect
`func (o *EnforcementEventsRow) UnsetAwaNDirect()`

UnsetAwaNDirect ensures that no value is present for AwaNDirect, not even an explicit nil
### GetAwaNCritical

`func (o *EnforcementEventsRow) GetAwaNCritical() int32`

GetAwaNCritical returns the AwaNCritical field if non-nil, zero value otherwise.

### GetAwaNCriticalOk

`func (o *EnforcementEventsRow) GetAwaNCriticalOk() (*int32, bool)`

GetAwaNCriticalOk returns a tuple with the AwaNCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaNCritical

`func (o *EnforcementEventsRow) SetAwaNCritical(v int32)`

SetAwaNCritical sets AwaNCritical field to given value.

### HasAwaNCritical

`func (o *EnforcementEventsRow) HasAwaNCritical() bool`

HasAwaNCritical returns a boolean if a field has been set.

### SetAwaNCriticalNil

`func (o *EnforcementEventsRow) SetAwaNCriticalNil(b bool)`

 SetAwaNCriticalNil sets the value for AwaNCritical to be an explicit nil

### UnsetAwaNCritical
`func (o *EnforcementEventsRow) UnsetAwaNCritical()`

UnsetAwaNCritical ensures that no value is present for AwaNCritical, not even an explicit nil
### GetAwaNNonCritical

`func (o *EnforcementEventsRow) GetAwaNNonCritical() int32`

GetAwaNNonCritical returns the AwaNNonCritical field if non-nil, zero value otherwise.

### GetAwaNNonCriticalOk

`func (o *EnforcementEventsRow) GetAwaNNonCriticalOk() (*int32, bool)`

GetAwaNNonCriticalOk returns a tuple with the AwaNNonCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaNNonCritical

`func (o *EnforcementEventsRow) SetAwaNNonCritical(v int32)`

SetAwaNNonCritical sets AwaNNonCritical field to given value.

### HasAwaNNonCritical

`func (o *EnforcementEventsRow) HasAwaNNonCritical() bool`

HasAwaNNonCritical returns a boolean if a field has been set.

### SetAwaNNonCriticalNil

`func (o *EnforcementEventsRow) SetAwaNNonCriticalNil(b bool)`

 SetAwaNNonCriticalNil sets the value for AwaNNonCritical to be an explicit nil

### UnsetAwaNNonCritical
`func (o *EnforcementEventsRow) UnsetAwaNNonCritical()`

UnsetAwaNNonCritical ensures that no value is present for AwaNNonCritical, not even an explicit nil
### GetAwaNRepeatCitations

`func (o *EnforcementEventsRow) GetAwaNRepeatCitations() int32`

GetAwaNRepeatCitations returns the AwaNRepeatCitations field if non-nil, zero value otherwise.

### GetAwaNRepeatCitationsOk

`func (o *EnforcementEventsRow) GetAwaNRepeatCitationsOk() (*int32, bool)`

GetAwaNRepeatCitationsOk returns a tuple with the AwaNRepeatCitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaNRepeatCitations

`func (o *EnforcementEventsRow) SetAwaNRepeatCitations(v int32)`

SetAwaNRepeatCitations sets AwaNRepeatCitations field to given value.

### HasAwaNRepeatCitations

`func (o *EnforcementEventsRow) HasAwaNRepeatCitations() bool`

HasAwaNRepeatCitations returns a boolean if a field has been set.

### SetAwaNRepeatCitationsNil

`func (o *EnforcementEventsRow) SetAwaNRepeatCitationsNil(b bool)`

 SetAwaNRepeatCitationsNil sets the value for AwaNRepeatCitations to be an explicit nil

### UnsetAwaNRepeatCitations
`func (o *EnforcementEventsRow) UnsetAwaNRepeatCitations()`

UnsetAwaNRepeatCitations ensures that no value is present for AwaNRepeatCitations, not even an explicit nil
### GetAwaSpeciesSummary

`func (o *EnforcementEventsRow) GetAwaSpeciesSummary() string`

GetAwaSpeciesSummary returns the AwaSpeciesSummary field if non-nil, zero value otherwise.

### GetAwaSpeciesSummaryOk

`func (o *EnforcementEventsRow) GetAwaSpeciesSummaryOk() (*string, bool)`

GetAwaSpeciesSummaryOk returns a tuple with the AwaSpeciesSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaSpeciesSummary

`func (o *EnforcementEventsRow) SetAwaSpeciesSummary(v string)`

SetAwaSpeciesSummary sets AwaSpeciesSummary field to given value.

### HasAwaSpeciesSummary

`func (o *EnforcementEventsRow) HasAwaSpeciesSummary() bool`

HasAwaSpeciesSummary returns a boolean if a field has been set.

### SetAwaSpeciesSummaryNil

`func (o *EnforcementEventsRow) SetAwaSpeciesSummaryNil(b bool)`

 SetAwaSpeciesSummaryNil sets the value for AwaSpeciesSummary to be an explicit nil

### UnsetAwaSpeciesSummary
`func (o *EnforcementEventsRow) UnsetAwaSpeciesSummary()`

UnsetAwaSpeciesSummary ensures that no value is present for AwaSpeciesSummary, not even an explicit nil
### GetAwaSpeciesScope

`func (o *EnforcementEventsRow) GetAwaSpeciesScope() string`

GetAwaSpeciesScope returns the AwaSpeciesScope field if non-nil, zero value otherwise.

### GetAwaSpeciesScopeOk

`func (o *EnforcementEventsRow) GetAwaSpeciesScopeOk() (*string, bool)`

GetAwaSpeciesScopeOk returns a tuple with the AwaSpeciesScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaSpeciesScope

`func (o *EnforcementEventsRow) SetAwaSpeciesScope(v string)`

SetAwaSpeciesScope sets AwaSpeciesScope field to given value.

### HasAwaSpeciesScope

`func (o *EnforcementEventsRow) HasAwaSpeciesScope() bool`

HasAwaSpeciesScope returns a boolean if a field has been set.

### SetAwaSpeciesScopeNil

`func (o *EnforcementEventsRow) SetAwaSpeciesScopeNil(b bool)`

 SetAwaSpeciesScopeNil sets the value for AwaSpeciesScope to be an explicit nil

### UnsetAwaSpeciesScope
`func (o *EnforcementEventsRow) UnsetAwaSpeciesScope()`

UnsetAwaSpeciesScope ensures that no value is present for AwaSpeciesScope, not even an explicit nil
### GetFirmName

`func (o *EnforcementEventsRow) GetFirmName() string`

GetFirmName returns the FirmName field if non-nil, zero value otherwise.

### GetFirmNameOk

`func (o *EnforcementEventsRow) GetFirmNameOk() (*string, bool)`

GetFirmNameOk returns a tuple with the FirmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmName

`func (o *EnforcementEventsRow) SetFirmName(v string)`

SetFirmName sets FirmName field to given value.

### HasFirmName

`func (o *EnforcementEventsRow) HasFirmName() bool`

HasFirmName returns a boolean if a field has been set.

### SetFirmNameNil

`func (o *EnforcementEventsRow) SetFirmNameNil(b bool)`

 SetFirmNameNil sets the value for FirmName to be an explicit nil

### UnsetFirmName
`func (o *EnforcementEventsRow) UnsetFirmName()`

UnsetFirmName ensures that no value is present for FirmName, not even an explicit nil
### GetFirmOrgId

`func (o *EnforcementEventsRow) GetFirmOrgId() string`

GetFirmOrgId returns the FirmOrgId field if non-nil, zero value otherwise.

### GetFirmOrgIdOk

`func (o *EnforcementEventsRow) GetFirmOrgIdOk() (*string, bool)`

GetFirmOrgIdOk returns a tuple with the FirmOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmOrgId

`func (o *EnforcementEventsRow) SetFirmOrgId(v string)`

SetFirmOrgId sets FirmOrgId field to given value.

### HasFirmOrgId

`func (o *EnforcementEventsRow) HasFirmOrgId() bool`

HasFirmOrgId returns a boolean if a field has been set.

### SetFirmOrgIdNil

`func (o *EnforcementEventsRow) SetFirmOrgIdNil(b bool)`

 SetFirmOrgIdNil sets the value for FirmOrgId to be an explicit nil

### UnsetFirmOrgId
`func (o *EnforcementEventsRow) UnsetFirmOrgId()`

UnsetFirmOrgId ensures that no value is present for FirmOrgId, not even an explicit nil
### GetFirmCompanyId

`func (o *EnforcementEventsRow) GetFirmCompanyId() string`

GetFirmCompanyId returns the FirmCompanyId field if non-nil, zero value otherwise.

### GetFirmCompanyIdOk

`func (o *EnforcementEventsRow) GetFirmCompanyIdOk() (*string, bool)`

GetFirmCompanyIdOk returns a tuple with the FirmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyId

`func (o *EnforcementEventsRow) SetFirmCompanyId(v string)`

SetFirmCompanyId sets FirmCompanyId field to given value.

### HasFirmCompanyId

`func (o *EnforcementEventsRow) HasFirmCompanyId() bool`

HasFirmCompanyId returns a boolean if a field has been set.

### SetFirmCompanyIdNil

`func (o *EnforcementEventsRow) SetFirmCompanyIdNil(b bool)`

 SetFirmCompanyIdNil sets the value for FirmCompanyId to be an explicit nil

### UnsetFirmCompanyId
`func (o *EnforcementEventsRow) UnsetFirmCompanyId()`

UnsetFirmCompanyId ensures that no value is present for FirmCompanyId, not even an explicit nil
### GetFirmCompanyName

`func (o *EnforcementEventsRow) GetFirmCompanyName() string`

GetFirmCompanyName returns the FirmCompanyName field if non-nil, zero value otherwise.

### GetFirmCompanyNameOk

`func (o *EnforcementEventsRow) GetFirmCompanyNameOk() (*string, bool)`

GetFirmCompanyNameOk returns a tuple with the FirmCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmCompanyName

`func (o *EnforcementEventsRow) SetFirmCompanyName(v string)`

SetFirmCompanyName sets FirmCompanyName field to given value.

### HasFirmCompanyName

`func (o *EnforcementEventsRow) HasFirmCompanyName() bool`

HasFirmCompanyName returns a boolean if a field has been set.

### SetFirmCompanyNameNil

`func (o *EnforcementEventsRow) SetFirmCompanyNameNil(b bool)`

 SetFirmCompanyNameNil sets the value for FirmCompanyName to be an explicit nil

### UnsetFirmCompanyName
`func (o *EnforcementEventsRow) UnsetFirmCompanyName()`

UnsetFirmCompanyName ensures that no value is present for FirmCompanyName, not even an explicit nil
### GetFirmGroupId

`func (o *EnforcementEventsRow) GetFirmGroupId() string`

GetFirmGroupId returns the FirmGroupId field if non-nil, zero value otherwise.

### GetFirmGroupIdOk

`func (o *EnforcementEventsRow) GetFirmGroupIdOk() (*string, bool)`

GetFirmGroupIdOk returns a tuple with the FirmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupId

`func (o *EnforcementEventsRow) SetFirmGroupId(v string)`

SetFirmGroupId sets FirmGroupId field to given value.

### HasFirmGroupId

`func (o *EnforcementEventsRow) HasFirmGroupId() bool`

HasFirmGroupId returns a boolean if a field has been set.

### SetFirmGroupIdNil

`func (o *EnforcementEventsRow) SetFirmGroupIdNil(b bool)`

 SetFirmGroupIdNil sets the value for FirmGroupId to be an explicit nil

### UnsetFirmGroupId
`func (o *EnforcementEventsRow) UnsetFirmGroupId()`

UnsetFirmGroupId ensures that no value is present for FirmGroupId, not even an explicit nil
### GetFirmGroupName

`func (o *EnforcementEventsRow) GetFirmGroupName() string`

GetFirmGroupName returns the FirmGroupName field if non-nil, zero value otherwise.

### GetFirmGroupNameOk

`func (o *EnforcementEventsRow) GetFirmGroupNameOk() (*string, bool)`

GetFirmGroupNameOk returns a tuple with the FirmGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmGroupName

`func (o *EnforcementEventsRow) SetFirmGroupName(v string)`

SetFirmGroupName sets FirmGroupName field to given value.

### HasFirmGroupName

`func (o *EnforcementEventsRow) HasFirmGroupName() bool`

HasFirmGroupName returns a boolean if a field has been set.

### SetFirmGroupNameNil

`func (o *EnforcementEventsRow) SetFirmGroupNameNil(b bool)`

 SetFirmGroupNameNil sets the value for FirmGroupName to be an explicit nil

### UnsetFirmGroupName
`func (o *EnforcementEventsRow) UnsetFirmGroupName()`

UnsetFirmGroupName ensures that no value is present for FirmGroupName, not even an explicit nil
### GetMatchTier

`func (o *EnforcementEventsRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *EnforcementEventsRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *EnforcementEventsRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *EnforcementEventsRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *EnforcementEventsRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *EnforcementEventsRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetOrgResolutionMethod

`func (o *EnforcementEventsRow) GetOrgResolutionMethod() string`

GetOrgResolutionMethod returns the OrgResolutionMethod field if non-nil, zero value otherwise.

### GetOrgResolutionMethodOk

`func (o *EnforcementEventsRow) GetOrgResolutionMethodOk() (*string, bool)`

GetOrgResolutionMethodOk returns a tuple with the OrgResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgResolutionMethod

`func (o *EnforcementEventsRow) SetOrgResolutionMethod(v string)`

SetOrgResolutionMethod sets OrgResolutionMethod field to given value.

### HasOrgResolutionMethod

`func (o *EnforcementEventsRow) HasOrgResolutionMethod() bool`

HasOrgResolutionMethod returns a boolean if a field has been set.

### SetOrgResolutionMethodNil

`func (o *EnforcementEventsRow) SetOrgResolutionMethodNil(b bool)`

 SetOrgResolutionMethodNil sets the value for OrgResolutionMethod to be an explicit nil

### UnsetOrgResolutionMethod
`func (o *EnforcementEventsRow) UnsetOrgResolutionMethod()`

UnsetOrgResolutionMethod ensures that no value is present for OrgResolutionMethod, not even an explicit nil
### GetIsOrgResolved

`func (o *EnforcementEventsRow) GetIsOrgResolved() bool`

GetIsOrgResolved returns the IsOrgResolved field if non-nil, zero value otherwise.

### GetIsOrgResolvedOk

`func (o *EnforcementEventsRow) GetIsOrgResolvedOk() (*bool, bool)`

GetIsOrgResolvedOk returns a tuple with the IsOrgResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgResolved

`func (o *EnforcementEventsRow) SetIsOrgResolved(v bool)`

SetIsOrgResolved sets IsOrgResolved field to given value.

### HasIsOrgResolved

`func (o *EnforcementEventsRow) HasIsOrgResolved() bool`

HasIsOrgResolved returns a boolean if a field has been set.

### SetIsOrgResolvedNil

`func (o *EnforcementEventsRow) SetIsOrgResolvedNil(b bool)`

 SetIsOrgResolvedNil sets the value for IsOrgResolved to be an explicit nil

### UnsetIsOrgResolved
`func (o *EnforcementEventsRow) UnsetIsOrgResolved()`

UnsetIsOrgResolved ensures that no value is present for IsOrgResolved, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


