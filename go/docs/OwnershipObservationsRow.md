# OwnershipObservationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnershipObservationId** | Pointer to **NullableString** | The observation id. | [optional] 
**SourceObservationId** | Pointer to **NullableString** | The id the source gave the assertion. | [optional] 
**SupersedesObservationId** | Pointer to **NullableString** | The earlier observation this one amends or replaces. | [optional] 
**SupersededByObservationId** | Pointer to **NullableString** | The later observation that supersedes this one (the lowest when more than one claims it). | [optional] 
**SupersededByCount** | Pointer to **NullableInt32** | How many later observations claim to supersede this one; above 1 is a lineage conflict kept as filed. | [optional] 
**SupersessionChainId** | Pointer to **NullableString** | The id of the amendment thread this observation belongs to. | [optional] 
**SupersessionDepth** | Pointer to **NullableInt32** | Position in the amendment thread; 0 is the first filing. | [optional] 
**IsChainTerminal** | Pointer to **NullableBool** | True when nothing supersedes this observation. | [optional] 
**ChainTruncated** | Pointer to **NullableBool** | True when the amendment thread was cut at 25 steps. | [optional] 
**SourceSystem** | Pointer to **NullableString** | The source system that made the assertion. | [optional] 
**SourceClass** | Pointer to **NullableString** | The filing family: sec_forms_3_4_5, sec_schedules_13d_g, sec_form_adv, sec_def_14a or wikidata. | [optional] 
**SemanticKey** | Pointer to **NullableString** | The ownership semantic the assertion carries, from the project&#39;s controlled taxonomy. | [optional] 
**RelationshipType** | Pointer to **NullableString** | The owner&#39;s relationship to the company as the filing family states it. | [optional] 
**OwnershipDimension** | Pointer to **NullableString** | What the observation measures: economic_interest, voting_power, beneficial_ownership, investment_discretion, insider_reporting, consolidation, control or unspecified. Different dimensions are different measurements. | [optional] 
**Directness** | Pointer to **NullableString** | direct, indirect, mixed or unspecified. | [optional] 
**OwnerKind** | Pointer to **NullableString** | What the owner is: organization, natural_person, reporting_group, government or unknown. | [optional] 
**SubjectSourceRef** | Pointer to **NullableString** | The owned party as the source identifies it. | [optional] 
**SubjectName** | Pointer to **NullableString** | The owned company as filed. | [optional] 
**SubjectLegalEntityId** | Pointer to **NullableString** | The legal entity the source resolved the owned party to, when it did. | [optional] 
**SubjectCompanyId** | Pointer to **NullableString** | Company level: the subject company (null when unresolved). | [optional] 
**SubjectCompanyName** | Pointer to **NullableString** | Display name of the subject company. | [optional] 
**SubjectGroupId** | Pointer to **NullableString** | Group level: the subject company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**SubjectGroupName** | Pointer to **NullableString** | Display name of the subject company&#39;s group. | [optional] 
**SubjectOrgId** | Pointer to **NullableString** | Legal-entity level: the owned company&#39;s organization record. | [optional] 
**SubjectInvestorId** | Pointer to **NullableString** | The investor directory key of the owned party when it is an investor. | [optional] 
**SubjectResolutionMethod** | Pointer to **NullableString** | How the owned party was resolved. | [optional] 
**SubjectResolutionConfidence** | Pointer to **NullableFloat32** | Confidence of that resolution, 0 to 1. | [optional] 
**OwnerSourceRef** | Pointer to **NullableString** | The owner as the source identifies it. | [optional] 
**OwnerName** | Pointer to **NullableString** | The owner&#39;s name as filed. | [optional] 
**OwnerLegalEntityId** | Pointer to **NullableString** | The legal entity the source resolved the owner to, when it did. | [optional] 
**OwnerCompanyId** | Pointer to **NullableString** | Company level: the owner company (null when unresolved). | [optional] 
**OwnerCompanyName** | Pointer to **NullableString** | Display name of the owner company. | [optional] 
**OwnerGroupId** | Pointer to **NullableString** | Group level: the owner company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**OwnerGroupName** | Pointer to **NullableString** | Display name of the owner company&#39;s group. | [optional] 
**OwnerOrgId** | Pointer to **NullableString** | Legal-entity level: the owner&#39;s organization record. | [optional] 
**OwnerInvestorId** | Pointer to **NullableString** | The investor directory key of the owner when it is one. | [optional] 
**OwnerPersonId** | Pointer to **NullableString** | The key of an owner who is a natural person; null for other owners. | [optional] 
**PersonLinkMethod** | Pointer to **NullableString** | How the person key was linked: edgar_cik, adviser_crd, unresolved_person or not_person. | [optional] 
**PersonLinkConflict** | Pointer to **NullableBool** | True when the owner links to more than one person key. | [optional] 
**OwnerResolutionMethod** | Pointer to **NullableString** | How the owner was resolved. | [optional] 
**OwnerResolutionConfidence** | Pointer to **NullableFloat32** | Confidence of that resolution, 0 to 1. | [optional] 
**ReviewState** | Pointer to **NullableString** | The review state of the resolution. | [optional] 
**StakeValueType** | Pointer to **NullableString** | How the stake is stated: exact, range, band, explicitly_unknown, not_reported or not_applicable. No midpoint, zero or numeral is invented for a non-numeric state. | [optional] 
**PercentageBasis** | Pointer to **NullableString** | What the percentage is a share of: economic_interest, voting_power, beneficial_ownership, control or ownership_unspecified. | [optional] 
**StakePercent** | Pointer to **NullableFloat32** | The stake as an exact percentage, 0 to 100. | [optional] 
**StakeLowerPercent** | Pointer to **NullableFloat32** | The lower bound of a range or band stake, 0 to 100. | [optional] 
**StakeUpperPercent** | Pointer to **NullableFloat32** | The upper bound of a range or band stake, 0 to 100. | [optional] 
**StakeLowerInclusive** | Pointer to **NullableBool** | Whether the lower bound is included. | [optional] 
**StakeUpperInclusive** | Pointer to **NullableBool** | Whether the upper bound is included. | [optional] 
**StakeBandCode** | Pointer to **NullableString** | The source&#39;s band code, when the stake is a band. | [optional] 
**StakeDisplay** | Pointer to **NullableString** | A label for the stake that never invents a number, such as 5.5%, 5% to 10% or Not reported. | [optional] 
**SecurityTitle** | Pointer to **NullableString** | The security the stake is in, as filed. | [optional] 
**SharesAmount** | Pointer to **NullableFloat32** | The number of shares, when the source states one. | [optional] 
**ReportedValueAmount** | Pointer to **NullableFloat32** | The reported value, in reported_value_currency at reported_value_scale. | [optional] 
**ReportedValueCurrency** | Pointer to **NullableString** | The currency of the reported value. | [optional] 
**ReportedValueScale** | Pointer to **NullableString** | The scale of the reported value as filed. | [optional] 
**AnnouncedDate** | Pointer to **NullableString** | The date announced. | [optional] 
**EffectiveDate** | Pointer to **NullableString** | The effective date. | [optional] 
**FilingDate** | Pointer to **NullableString** | The filing date. | [optional] 
**AsOfDate** | Pointer to **NullableString** | The as-of date of the figure. | [optional] 
**CeasedDate** | Pointer to **NullableString** | The date the position ceased. | [optional] 
**ObservationDate** | Pointer to **NullableString** | The sort date: the first of as_of_date, effective_date, filing_date and announced_date that is present. | [optional] 
**ObservationDateBasis** | Pointer to **NullableString** | Which of those dates observation_date took, or unavailable. | [optional] 
**LifecycleStatus** | Pointer to **NullableString** | The lifecycle status the source states. | [optional] 
**SourceDocumentId** | Pointer to **NullableString** | The id of the source document. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The link to the filing. | [optional] 
**SourceContentHash** | Pointer to **NullableString** | A hash of the source content, for change detection. | [optional] 
**EvidenceQuote** | Pointer to **NullableString** | The quoted evidence from the source. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The tier vocabulary is none or licensable. Every row here is none: every source behind the ledger is open, and a data test reports any other tier. | [optional] 
**SourceLicense** | Pointer to **NullableString** | The licence of the source. | [optional] 
**CommercialUseAllowed** | Pointer to **NullableBool** | Whether the source licence allows commercial use. | [optional] 
**Attribution** | Pointer to **NullableString** | The source credit text for the row. | [optional] 
**ContainsPersonalData** | Pointer to **NullableBool** | A tag, never a gate: true when the source row carries a natural person&#39;s data. | [optional] 
**PrivacyStatus** | Pointer to **NullableString** | The source&#39;s own privacy classification of the row. | [optional] 
**IsCurrentEdge** | Pointer to **NullableBool** | True when this observation is the one the current ownership edge selected. | [optional] 
**CurrentEdgeState** | Pointer to **NullableString** | The state of the current edge when is_current_edge. | [optional] 

## Methods

### NewOwnershipObservationsRow

`func NewOwnershipObservationsRow() *OwnershipObservationsRow`

NewOwnershipObservationsRow instantiates a new OwnershipObservationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnershipObservationsRowWithDefaults

`func NewOwnershipObservationsRowWithDefaults() *OwnershipObservationsRow`

NewOwnershipObservationsRowWithDefaults instantiates a new OwnershipObservationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnershipObservationId

`func (o *OwnershipObservationsRow) GetOwnershipObservationId() string`

GetOwnershipObservationId returns the OwnershipObservationId field if non-nil, zero value otherwise.

### GetOwnershipObservationIdOk

`func (o *OwnershipObservationsRow) GetOwnershipObservationIdOk() (*string, bool)`

GetOwnershipObservationIdOk returns a tuple with the OwnershipObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipObservationId

`func (o *OwnershipObservationsRow) SetOwnershipObservationId(v string)`

SetOwnershipObservationId sets OwnershipObservationId field to given value.

### HasOwnershipObservationId

`func (o *OwnershipObservationsRow) HasOwnershipObservationId() bool`

HasOwnershipObservationId returns a boolean if a field has been set.

### SetOwnershipObservationIdNil

`func (o *OwnershipObservationsRow) SetOwnershipObservationIdNil(b bool)`

 SetOwnershipObservationIdNil sets the value for OwnershipObservationId to be an explicit nil

### UnsetOwnershipObservationId
`func (o *OwnershipObservationsRow) UnsetOwnershipObservationId()`

UnsetOwnershipObservationId ensures that no value is present for OwnershipObservationId, not even an explicit nil
### GetSourceObservationId

`func (o *OwnershipObservationsRow) GetSourceObservationId() string`

GetSourceObservationId returns the SourceObservationId field if non-nil, zero value otherwise.

### GetSourceObservationIdOk

`func (o *OwnershipObservationsRow) GetSourceObservationIdOk() (*string, bool)`

GetSourceObservationIdOk returns a tuple with the SourceObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObservationId

`func (o *OwnershipObservationsRow) SetSourceObservationId(v string)`

SetSourceObservationId sets SourceObservationId field to given value.

### HasSourceObservationId

`func (o *OwnershipObservationsRow) HasSourceObservationId() bool`

HasSourceObservationId returns a boolean if a field has been set.

### SetSourceObservationIdNil

`func (o *OwnershipObservationsRow) SetSourceObservationIdNil(b bool)`

 SetSourceObservationIdNil sets the value for SourceObservationId to be an explicit nil

### UnsetSourceObservationId
`func (o *OwnershipObservationsRow) UnsetSourceObservationId()`

UnsetSourceObservationId ensures that no value is present for SourceObservationId, not even an explicit nil
### GetSupersedesObservationId

`func (o *OwnershipObservationsRow) GetSupersedesObservationId() string`

GetSupersedesObservationId returns the SupersedesObservationId field if non-nil, zero value otherwise.

### GetSupersedesObservationIdOk

`func (o *OwnershipObservationsRow) GetSupersedesObservationIdOk() (*string, bool)`

GetSupersedesObservationIdOk returns a tuple with the SupersedesObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersedesObservationId

`func (o *OwnershipObservationsRow) SetSupersedesObservationId(v string)`

SetSupersedesObservationId sets SupersedesObservationId field to given value.

### HasSupersedesObservationId

`func (o *OwnershipObservationsRow) HasSupersedesObservationId() bool`

HasSupersedesObservationId returns a boolean if a field has been set.

### SetSupersedesObservationIdNil

`func (o *OwnershipObservationsRow) SetSupersedesObservationIdNil(b bool)`

 SetSupersedesObservationIdNil sets the value for SupersedesObservationId to be an explicit nil

### UnsetSupersedesObservationId
`func (o *OwnershipObservationsRow) UnsetSupersedesObservationId()`

UnsetSupersedesObservationId ensures that no value is present for SupersedesObservationId, not even an explicit nil
### GetSupersededByObservationId

`func (o *OwnershipObservationsRow) GetSupersededByObservationId() string`

GetSupersededByObservationId returns the SupersededByObservationId field if non-nil, zero value otherwise.

### GetSupersededByObservationIdOk

`func (o *OwnershipObservationsRow) GetSupersededByObservationIdOk() (*string, bool)`

GetSupersededByObservationIdOk returns a tuple with the SupersededByObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededByObservationId

`func (o *OwnershipObservationsRow) SetSupersededByObservationId(v string)`

SetSupersededByObservationId sets SupersededByObservationId field to given value.

### HasSupersededByObservationId

`func (o *OwnershipObservationsRow) HasSupersededByObservationId() bool`

HasSupersededByObservationId returns a boolean if a field has been set.

### SetSupersededByObservationIdNil

`func (o *OwnershipObservationsRow) SetSupersededByObservationIdNil(b bool)`

 SetSupersededByObservationIdNil sets the value for SupersededByObservationId to be an explicit nil

### UnsetSupersededByObservationId
`func (o *OwnershipObservationsRow) UnsetSupersededByObservationId()`

UnsetSupersededByObservationId ensures that no value is present for SupersededByObservationId, not even an explicit nil
### GetSupersededByCount

`func (o *OwnershipObservationsRow) GetSupersededByCount() int32`

GetSupersededByCount returns the SupersededByCount field if non-nil, zero value otherwise.

### GetSupersededByCountOk

`func (o *OwnershipObservationsRow) GetSupersededByCountOk() (*int32, bool)`

GetSupersededByCountOk returns a tuple with the SupersededByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededByCount

`func (o *OwnershipObservationsRow) SetSupersededByCount(v int32)`

SetSupersededByCount sets SupersededByCount field to given value.

### HasSupersededByCount

`func (o *OwnershipObservationsRow) HasSupersededByCount() bool`

HasSupersededByCount returns a boolean if a field has been set.

### SetSupersededByCountNil

`func (o *OwnershipObservationsRow) SetSupersededByCountNil(b bool)`

 SetSupersededByCountNil sets the value for SupersededByCount to be an explicit nil

### UnsetSupersededByCount
`func (o *OwnershipObservationsRow) UnsetSupersededByCount()`

UnsetSupersededByCount ensures that no value is present for SupersededByCount, not even an explicit nil
### GetSupersessionChainId

`func (o *OwnershipObservationsRow) GetSupersessionChainId() string`

GetSupersessionChainId returns the SupersessionChainId field if non-nil, zero value otherwise.

### GetSupersessionChainIdOk

`func (o *OwnershipObservationsRow) GetSupersessionChainIdOk() (*string, bool)`

GetSupersessionChainIdOk returns a tuple with the SupersessionChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersessionChainId

`func (o *OwnershipObservationsRow) SetSupersessionChainId(v string)`

SetSupersessionChainId sets SupersessionChainId field to given value.

### HasSupersessionChainId

`func (o *OwnershipObservationsRow) HasSupersessionChainId() bool`

HasSupersessionChainId returns a boolean if a field has been set.

### SetSupersessionChainIdNil

`func (o *OwnershipObservationsRow) SetSupersessionChainIdNil(b bool)`

 SetSupersessionChainIdNil sets the value for SupersessionChainId to be an explicit nil

### UnsetSupersessionChainId
`func (o *OwnershipObservationsRow) UnsetSupersessionChainId()`

UnsetSupersessionChainId ensures that no value is present for SupersessionChainId, not even an explicit nil
### GetSupersessionDepth

`func (o *OwnershipObservationsRow) GetSupersessionDepth() int32`

GetSupersessionDepth returns the SupersessionDepth field if non-nil, zero value otherwise.

### GetSupersessionDepthOk

`func (o *OwnershipObservationsRow) GetSupersessionDepthOk() (*int32, bool)`

GetSupersessionDepthOk returns a tuple with the SupersessionDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersessionDepth

`func (o *OwnershipObservationsRow) SetSupersessionDepth(v int32)`

SetSupersessionDepth sets SupersessionDepth field to given value.

### HasSupersessionDepth

`func (o *OwnershipObservationsRow) HasSupersessionDepth() bool`

HasSupersessionDepth returns a boolean if a field has been set.

### SetSupersessionDepthNil

`func (o *OwnershipObservationsRow) SetSupersessionDepthNil(b bool)`

 SetSupersessionDepthNil sets the value for SupersessionDepth to be an explicit nil

### UnsetSupersessionDepth
`func (o *OwnershipObservationsRow) UnsetSupersessionDepth()`

UnsetSupersessionDepth ensures that no value is present for SupersessionDepth, not even an explicit nil
### GetIsChainTerminal

`func (o *OwnershipObservationsRow) GetIsChainTerminal() bool`

GetIsChainTerminal returns the IsChainTerminal field if non-nil, zero value otherwise.

### GetIsChainTerminalOk

`func (o *OwnershipObservationsRow) GetIsChainTerminalOk() (*bool, bool)`

GetIsChainTerminalOk returns a tuple with the IsChainTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChainTerminal

`func (o *OwnershipObservationsRow) SetIsChainTerminal(v bool)`

SetIsChainTerminal sets IsChainTerminal field to given value.

### HasIsChainTerminal

`func (o *OwnershipObservationsRow) HasIsChainTerminal() bool`

HasIsChainTerminal returns a boolean if a field has been set.

### SetIsChainTerminalNil

`func (o *OwnershipObservationsRow) SetIsChainTerminalNil(b bool)`

 SetIsChainTerminalNil sets the value for IsChainTerminal to be an explicit nil

### UnsetIsChainTerminal
`func (o *OwnershipObservationsRow) UnsetIsChainTerminal()`

UnsetIsChainTerminal ensures that no value is present for IsChainTerminal, not even an explicit nil
### GetChainTruncated

`func (o *OwnershipObservationsRow) GetChainTruncated() bool`

GetChainTruncated returns the ChainTruncated field if non-nil, zero value otherwise.

### GetChainTruncatedOk

`func (o *OwnershipObservationsRow) GetChainTruncatedOk() (*bool, bool)`

GetChainTruncatedOk returns a tuple with the ChainTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainTruncated

`func (o *OwnershipObservationsRow) SetChainTruncated(v bool)`

SetChainTruncated sets ChainTruncated field to given value.

### HasChainTruncated

`func (o *OwnershipObservationsRow) HasChainTruncated() bool`

HasChainTruncated returns a boolean if a field has been set.

### SetChainTruncatedNil

`func (o *OwnershipObservationsRow) SetChainTruncatedNil(b bool)`

 SetChainTruncatedNil sets the value for ChainTruncated to be an explicit nil

### UnsetChainTruncated
`func (o *OwnershipObservationsRow) UnsetChainTruncated()`

UnsetChainTruncated ensures that no value is present for ChainTruncated, not even an explicit nil
### GetSourceSystem

`func (o *OwnershipObservationsRow) GetSourceSystem() string`

GetSourceSystem returns the SourceSystem field if non-nil, zero value otherwise.

### GetSourceSystemOk

`func (o *OwnershipObservationsRow) GetSourceSystemOk() (*string, bool)`

GetSourceSystemOk returns a tuple with the SourceSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystem

`func (o *OwnershipObservationsRow) SetSourceSystem(v string)`

SetSourceSystem sets SourceSystem field to given value.

### HasSourceSystem

`func (o *OwnershipObservationsRow) HasSourceSystem() bool`

HasSourceSystem returns a boolean if a field has been set.

### SetSourceSystemNil

`func (o *OwnershipObservationsRow) SetSourceSystemNil(b bool)`

 SetSourceSystemNil sets the value for SourceSystem to be an explicit nil

### UnsetSourceSystem
`func (o *OwnershipObservationsRow) UnsetSourceSystem()`

UnsetSourceSystem ensures that no value is present for SourceSystem, not even an explicit nil
### GetSourceClass

`func (o *OwnershipObservationsRow) GetSourceClass() string`

GetSourceClass returns the SourceClass field if non-nil, zero value otherwise.

### GetSourceClassOk

`func (o *OwnershipObservationsRow) GetSourceClassOk() (*string, bool)`

GetSourceClassOk returns a tuple with the SourceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClass

`func (o *OwnershipObservationsRow) SetSourceClass(v string)`

SetSourceClass sets SourceClass field to given value.

### HasSourceClass

`func (o *OwnershipObservationsRow) HasSourceClass() bool`

HasSourceClass returns a boolean if a field has been set.

### SetSourceClassNil

`func (o *OwnershipObservationsRow) SetSourceClassNil(b bool)`

 SetSourceClassNil sets the value for SourceClass to be an explicit nil

### UnsetSourceClass
`func (o *OwnershipObservationsRow) UnsetSourceClass()`

UnsetSourceClass ensures that no value is present for SourceClass, not even an explicit nil
### GetSemanticKey

`func (o *OwnershipObservationsRow) GetSemanticKey() string`

GetSemanticKey returns the SemanticKey field if non-nil, zero value otherwise.

### GetSemanticKeyOk

`func (o *OwnershipObservationsRow) GetSemanticKeyOk() (*string, bool)`

GetSemanticKeyOk returns a tuple with the SemanticKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticKey

`func (o *OwnershipObservationsRow) SetSemanticKey(v string)`

SetSemanticKey sets SemanticKey field to given value.

### HasSemanticKey

`func (o *OwnershipObservationsRow) HasSemanticKey() bool`

HasSemanticKey returns a boolean if a field has been set.

### SetSemanticKeyNil

`func (o *OwnershipObservationsRow) SetSemanticKeyNil(b bool)`

 SetSemanticKeyNil sets the value for SemanticKey to be an explicit nil

### UnsetSemanticKey
`func (o *OwnershipObservationsRow) UnsetSemanticKey()`

UnsetSemanticKey ensures that no value is present for SemanticKey, not even an explicit nil
### GetRelationshipType

`func (o *OwnershipObservationsRow) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *OwnershipObservationsRow) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *OwnershipObservationsRow) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.

### HasRelationshipType

`func (o *OwnershipObservationsRow) HasRelationshipType() bool`

HasRelationshipType returns a boolean if a field has been set.

### SetRelationshipTypeNil

`func (o *OwnershipObservationsRow) SetRelationshipTypeNil(b bool)`

 SetRelationshipTypeNil sets the value for RelationshipType to be an explicit nil

### UnsetRelationshipType
`func (o *OwnershipObservationsRow) UnsetRelationshipType()`

UnsetRelationshipType ensures that no value is present for RelationshipType, not even an explicit nil
### GetOwnershipDimension

`func (o *OwnershipObservationsRow) GetOwnershipDimension() string`

GetOwnershipDimension returns the OwnershipDimension field if non-nil, zero value otherwise.

### GetOwnershipDimensionOk

`func (o *OwnershipObservationsRow) GetOwnershipDimensionOk() (*string, bool)`

GetOwnershipDimensionOk returns a tuple with the OwnershipDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipDimension

`func (o *OwnershipObservationsRow) SetOwnershipDimension(v string)`

SetOwnershipDimension sets OwnershipDimension field to given value.

### HasOwnershipDimension

`func (o *OwnershipObservationsRow) HasOwnershipDimension() bool`

HasOwnershipDimension returns a boolean if a field has been set.

### SetOwnershipDimensionNil

`func (o *OwnershipObservationsRow) SetOwnershipDimensionNil(b bool)`

 SetOwnershipDimensionNil sets the value for OwnershipDimension to be an explicit nil

### UnsetOwnershipDimension
`func (o *OwnershipObservationsRow) UnsetOwnershipDimension()`

UnsetOwnershipDimension ensures that no value is present for OwnershipDimension, not even an explicit nil
### GetDirectness

`func (o *OwnershipObservationsRow) GetDirectness() string`

GetDirectness returns the Directness field if non-nil, zero value otherwise.

### GetDirectnessOk

`func (o *OwnershipObservationsRow) GetDirectnessOk() (*string, bool)`

GetDirectnessOk returns a tuple with the Directness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectness

`func (o *OwnershipObservationsRow) SetDirectness(v string)`

SetDirectness sets Directness field to given value.

### HasDirectness

`func (o *OwnershipObservationsRow) HasDirectness() bool`

HasDirectness returns a boolean if a field has been set.

### SetDirectnessNil

`func (o *OwnershipObservationsRow) SetDirectnessNil(b bool)`

 SetDirectnessNil sets the value for Directness to be an explicit nil

### UnsetDirectness
`func (o *OwnershipObservationsRow) UnsetDirectness()`

UnsetDirectness ensures that no value is present for Directness, not even an explicit nil
### GetOwnerKind

`func (o *OwnershipObservationsRow) GetOwnerKind() string`

GetOwnerKind returns the OwnerKind field if non-nil, zero value otherwise.

### GetOwnerKindOk

`func (o *OwnershipObservationsRow) GetOwnerKindOk() (*string, bool)`

GetOwnerKindOk returns a tuple with the OwnerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerKind

`func (o *OwnershipObservationsRow) SetOwnerKind(v string)`

SetOwnerKind sets OwnerKind field to given value.

### HasOwnerKind

`func (o *OwnershipObservationsRow) HasOwnerKind() bool`

HasOwnerKind returns a boolean if a field has been set.

### SetOwnerKindNil

`func (o *OwnershipObservationsRow) SetOwnerKindNil(b bool)`

 SetOwnerKindNil sets the value for OwnerKind to be an explicit nil

### UnsetOwnerKind
`func (o *OwnershipObservationsRow) UnsetOwnerKind()`

UnsetOwnerKind ensures that no value is present for OwnerKind, not even an explicit nil
### GetSubjectSourceRef

`func (o *OwnershipObservationsRow) GetSubjectSourceRef() string`

GetSubjectSourceRef returns the SubjectSourceRef field if non-nil, zero value otherwise.

### GetSubjectSourceRefOk

`func (o *OwnershipObservationsRow) GetSubjectSourceRefOk() (*string, bool)`

GetSubjectSourceRefOk returns a tuple with the SubjectSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectSourceRef

`func (o *OwnershipObservationsRow) SetSubjectSourceRef(v string)`

SetSubjectSourceRef sets SubjectSourceRef field to given value.

### HasSubjectSourceRef

`func (o *OwnershipObservationsRow) HasSubjectSourceRef() bool`

HasSubjectSourceRef returns a boolean if a field has been set.

### SetSubjectSourceRefNil

`func (o *OwnershipObservationsRow) SetSubjectSourceRefNil(b bool)`

 SetSubjectSourceRefNil sets the value for SubjectSourceRef to be an explicit nil

### UnsetSubjectSourceRef
`func (o *OwnershipObservationsRow) UnsetSubjectSourceRef()`

UnsetSubjectSourceRef ensures that no value is present for SubjectSourceRef, not even an explicit nil
### GetSubjectName

`func (o *OwnershipObservationsRow) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *OwnershipObservationsRow) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *OwnershipObservationsRow) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *OwnershipObservationsRow) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### SetSubjectNameNil

`func (o *OwnershipObservationsRow) SetSubjectNameNil(b bool)`

 SetSubjectNameNil sets the value for SubjectName to be an explicit nil

### UnsetSubjectName
`func (o *OwnershipObservationsRow) UnsetSubjectName()`

UnsetSubjectName ensures that no value is present for SubjectName, not even an explicit nil
### GetSubjectLegalEntityId

`func (o *OwnershipObservationsRow) GetSubjectLegalEntityId() string`

GetSubjectLegalEntityId returns the SubjectLegalEntityId field if non-nil, zero value otherwise.

### GetSubjectLegalEntityIdOk

`func (o *OwnershipObservationsRow) GetSubjectLegalEntityIdOk() (*string, bool)`

GetSubjectLegalEntityIdOk returns a tuple with the SubjectLegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectLegalEntityId

`func (o *OwnershipObservationsRow) SetSubjectLegalEntityId(v string)`

SetSubjectLegalEntityId sets SubjectLegalEntityId field to given value.

### HasSubjectLegalEntityId

`func (o *OwnershipObservationsRow) HasSubjectLegalEntityId() bool`

HasSubjectLegalEntityId returns a boolean if a field has been set.

### SetSubjectLegalEntityIdNil

`func (o *OwnershipObservationsRow) SetSubjectLegalEntityIdNil(b bool)`

 SetSubjectLegalEntityIdNil sets the value for SubjectLegalEntityId to be an explicit nil

### UnsetSubjectLegalEntityId
`func (o *OwnershipObservationsRow) UnsetSubjectLegalEntityId()`

UnsetSubjectLegalEntityId ensures that no value is present for SubjectLegalEntityId, not even an explicit nil
### GetSubjectCompanyId

`func (o *OwnershipObservationsRow) GetSubjectCompanyId() string`

GetSubjectCompanyId returns the SubjectCompanyId field if non-nil, zero value otherwise.

### GetSubjectCompanyIdOk

`func (o *OwnershipObservationsRow) GetSubjectCompanyIdOk() (*string, bool)`

GetSubjectCompanyIdOk returns a tuple with the SubjectCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyId

`func (o *OwnershipObservationsRow) SetSubjectCompanyId(v string)`

SetSubjectCompanyId sets SubjectCompanyId field to given value.

### HasSubjectCompanyId

`func (o *OwnershipObservationsRow) HasSubjectCompanyId() bool`

HasSubjectCompanyId returns a boolean if a field has been set.

### SetSubjectCompanyIdNil

`func (o *OwnershipObservationsRow) SetSubjectCompanyIdNil(b bool)`

 SetSubjectCompanyIdNil sets the value for SubjectCompanyId to be an explicit nil

### UnsetSubjectCompanyId
`func (o *OwnershipObservationsRow) UnsetSubjectCompanyId()`

UnsetSubjectCompanyId ensures that no value is present for SubjectCompanyId, not even an explicit nil
### GetSubjectCompanyName

`func (o *OwnershipObservationsRow) GetSubjectCompanyName() string`

GetSubjectCompanyName returns the SubjectCompanyName field if non-nil, zero value otherwise.

### GetSubjectCompanyNameOk

`func (o *OwnershipObservationsRow) GetSubjectCompanyNameOk() (*string, bool)`

GetSubjectCompanyNameOk returns a tuple with the SubjectCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyName

`func (o *OwnershipObservationsRow) SetSubjectCompanyName(v string)`

SetSubjectCompanyName sets SubjectCompanyName field to given value.

### HasSubjectCompanyName

`func (o *OwnershipObservationsRow) HasSubjectCompanyName() bool`

HasSubjectCompanyName returns a boolean if a field has been set.

### SetSubjectCompanyNameNil

`func (o *OwnershipObservationsRow) SetSubjectCompanyNameNil(b bool)`

 SetSubjectCompanyNameNil sets the value for SubjectCompanyName to be an explicit nil

### UnsetSubjectCompanyName
`func (o *OwnershipObservationsRow) UnsetSubjectCompanyName()`

UnsetSubjectCompanyName ensures that no value is present for SubjectCompanyName, not even an explicit nil
### GetSubjectGroupId

`func (o *OwnershipObservationsRow) GetSubjectGroupId() string`

GetSubjectGroupId returns the SubjectGroupId field if non-nil, zero value otherwise.

### GetSubjectGroupIdOk

`func (o *OwnershipObservationsRow) GetSubjectGroupIdOk() (*string, bool)`

GetSubjectGroupIdOk returns a tuple with the SubjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupId

`func (o *OwnershipObservationsRow) SetSubjectGroupId(v string)`

SetSubjectGroupId sets SubjectGroupId field to given value.

### HasSubjectGroupId

`func (o *OwnershipObservationsRow) HasSubjectGroupId() bool`

HasSubjectGroupId returns a boolean if a field has been set.

### SetSubjectGroupIdNil

`func (o *OwnershipObservationsRow) SetSubjectGroupIdNil(b bool)`

 SetSubjectGroupIdNil sets the value for SubjectGroupId to be an explicit nil

### UnsetSubjectGroupId
`func (o *OwnershipObservationsRow) UnsetSubjectGroupId()`

UnsetSubjectGroupId ensures that no value is present for SubjectGroupId, not even an explicit nil
### GetSubjectGroupName

`func (o *OwnershipObservationsRow) GetSubjectGroupName() string`

GetSubjectGroupName returns the SubjectGroupName field if non-nil, zero value otherwise.

### GetSubjectGroupNameOk

`func (o *OwnershipObservationsRow) GetSubjectGroupNameOk() (*string, bool)`

GetSubjectGroupNameOk returns a tuple with the SubjectGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupName

`func (o *OwnershipObservationsRow) SetSubjectGroupName(v string)`

SetSubjectGroupName sets SubjectGroupName field to given value.

### HasSubjectGroupName

`func (o *OwnershipObservationsRow) HasSubjectGroupName() bool`

HasSubjectGroupName returns a boolean if a field has been set.

### SetSubjectGroupNameNil

`func (o *OwnershipObservationsRow) SetSubjectGroupNameNil(b bool)`

 SetSubjectGroupNameNil sets the value for SubjectGroupName to be an explicit nil

### UnsetSubjectGroupName
`func (o *OwnershipObservationsRow) UnsetSubjectGroupName()`

UnsetSubjectGroupName ensures that no value is present for SubjectGroupName, not even an explicit nil
### GetSubjectOrgId

`func (o *OwnershipObservationsRow) GetSubjectOrgId() string`

GetSubjectOrgId returns the SubjectOrgId field if non-nil, zero value otherwise.

### GetSubjectOrgIdOk

`func (o *OwnershipObservationsRow) GetSubjectOrgIdOk() (*string, bool)`

GetSubjectOrgIdOk returns a tuple with the SubjectOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectOrgId

`func (o *OwnershipObservationsRow) SetSubjectOrgId(v string)`

SetSubjectOrgId sets SubjectOrgId field to given value.

### HasSubjectOrgId

`func (o *OwnershipObservationsRow) HasSubjectOrgId() bool`

HasSubjectOrgId returns a boolean if a field has been set.

### SetSubjectOrgIdNil

`func (o *OwnershipObservationsRow) SetSubjectOrgIdNil(b bool)`

 SetSubjectOrgIdNil sets the value for SubjectOrgId to be an explicit nil

### UnsetSubjectOrgId
`func (o *OwnershipObservationsRow) UnsetSubjectOrgId()`

UnsetSubjectOrgId ensures that no value is present for SubjectOrgId, not even an explicit nil
### GetSubjectInvestorId

`func (o *OwnershipObservationsRow) GetSubjectInvestorId() string`

GetSubjectInvestorId returns the SubjectInvestorId field if non-nil, zero value otherwise.

### GetSubjectInvestorIdOk

`func (o *OwnershipObservationsRow) GetSubjectInvestorIdOk() (*string, bool)`

GetSubjectInvestorIdOk returns a tuple with the SubjectInvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectInvestorId

`func (o *OwnershipObservationsRow) SetSubjectInvestorId(v string)`

SetSubjectInvestorId sets SubjectInvestorId field to given value.

### HasSubjectInvestorId

`func (o *OwnershipObservationsRow) HasSubjectInvestorId() bool`

HasSubjectInvestorId returns a boolean if a field has been set.

### SetSubjectInvestorIdNil

`func (o *OwnershipObservationsRow) SetSubjectInvestorIdNil(b bool)`

 SetSubjectInvestorIdNil sets the value for SubjectInvestorId to be an explicit nil

### UnsetSubjectInvestorId
`func (o *OwnershipObservationsRow) UnsetSubjectInvestorId()`

UnsetSubjectInvestorId ensures that no value is present for SubjectInvestorId, not even an explicit nil
### GetSubjectResolutionMethod

`func (o *OwnershipObservationsRow) GetSubjectResolutionMethod() string`

GetSubjectResolutionMethod returns the SubjectResolutionMethod field if non-nil, zero value otherwise.

### GetSubjectResolutionMethodOk

`func (o *OwnershipObservationsRow) GetSubjectResolutionMethodOk() (*string, bool)`

GetSubjectResolutionMethodOk returns a tuple with the SubjectResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectResolutionMethod

`func (o *OwnershipObservationsRow) SetSubjectResolutionMethod(v string)`

SetSubjectResolutionMethod sets SubjectResolutionMethod field to given value.

### HasSubjectResolutionMethod

`func (o *OwnershipObservationsRow) HasSubjectResolutionMethod() bool`

HasSubjectResolutionMethod returns a boolean if a field has been set.

### SetSubjectResolutionMethodNil

`func (o *OwnershipObservationsRow) SetSubjectResolutionMethodNil(b bool)`

 SetSubjectResolutionMethodNil sets the value for SubjectResolutionMethod to be an explicit nil

### UnsetSubjectResolutionMethod
`func (o *OwnershipObservationsRow) UnsetSubjectResolutionMethod()`

UnsetSubjectResolutionMethod ensures that no value is present for SubjectResolutionMethod, not even an explicit nil
### GetSubjectResolutionConfidence

`func (o *OwnershipObservationsRow) GetSubjectResolutionConfidence() float32`

GetSubjectResolutionConfidence returns the SubjectResolutionConfidence field if non-nil, zero value otherwise.

### GetSubjectResolutionConfidenceOk

`func (o *OwnershipObservationsRow) GetSubjectResolutionConfidenceOk() (*float32, bool)`

GetSubjectResolutionConfidenceOk returns a tuple with the SubjectResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectResolutionConfidence

`func (o *OwnershipObservationsRow) SetSubjectResolutionConfidence(v float32)`

SetSubjectResolutionConfidence sets SubjectResolutionConfidence field to given value.

### HasSubjectResolutionConfidence

`func (o *OwnershipObservationsRow) HasSubjectResolutionConfidence() bool`

HasSubjectResolutionConfidence returns a boolean if a field has been set.

### SetSubjectResolutionConfidenceNil

`func (o *OwnershipObservationsRow) SetSubjectResolutionConfidenceNil(b bool)`

 SetSubjectResolutionConfidenceNil sets the value for SubjectResolutionConfidence to be an explicit nil

### UnsetSubjectResolutionConfidence
`func (o *OwnershipObservationsRow) UnsetSubjectResolutionConfidence()`

UnsetSubjectResolutionConfidence ensures that no value is present for SubjectResolutionConfidence, not even an explicit nil
### GetOwnerSourceRef

`func (o *OwnershipObservationsRow) GetOwnerSourceRef() string`

GetOwnerSourceRef returns the OwnerSourceRef field if non-nil, zero value otherwise.

### GetOwnerSourceRefOk

`func (o *OwnershipObservationsRow) GetOwnerSourceRefOk() (*string, bool)`

GetOwnerSourceRefOk returns a tuple with the OwnerSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSourceRef

`func (o *OwnershipObservationsRow) SetOwnerSourceRef(v string)`

SetOwnerSourceRef sets OwnerSourceRef field to given value.

### HasOwnerSourceRef

`func (o *OwnershipObservationsRow) HasOwnerSourceRef() bool`

HasOwnerSourceRef returns a boolean if a field has been set.

### SetOwnerSourceRefNil

`func (o *OwnershipObservationsRow) SetOwnerSourceRefNil(b bool)`

 SetOwnerSourceRefNil sets the value for OwnerSourceRef to be an explicit nil

### UnsetOwnerSourceRef
`func (o *OwnershipObservationsRow) UnsetOwnerSourceRef()`

UnsetOwnerSourceRef ensures that no value is present for OwnerSourceRef, not even an explicit nil
### GetOwnerName

`func (o *OwnershipObservationsRow) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *OwnershipObservationsRow) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *OwnershipObservationsRow) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *OwnershipObservationsRow) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *OwnershipObservationsRow) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *OwnershipObservationsRow) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
### GetOwnerLegalEntityId

`func (o *OwnershipObservationsRow) GetOwnerLegalEntityId() string`

GetOwnerLegalEntityId returns the OwnerLegalEntityId field if non-nil, zero value otherwise.

### GetOwnerLegalEntityIdOk

`func (o *OwnershipObservationsRow) GetOwnerLegalEntityIdOk() (*string, bool)`

GetOwnerLegalEntityIdOk returns a tuple with the OwnerLegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerLegalEntityId

`func (o *OwnershipObservationsRow) SetOwnerLegalEntityId(v string)`

SetOwnerLegalEntityId sets OwnerLegalEntityId field to given value.

### HasOwnerLegalEntityId

`func (o *OwnershipObservationsRow) HasOwnerLegalEntityId() bool`

HasOwnerLegalEntityId returns a boolean if a field has been set.

### SetOwnerLegalEntityIdNil

`func (o *OwnershipObservationsRow) SetOwnerLegalEntityIdNil(b bool)`

 SetOwnerLegalEntityIdNil sets the value for OwnerLegalEntityId to be an explicit nil

### UnsetOwnerLegalEntityId
`func (o *OwnershipObservationsRow) UnsetOwnerLegalEntityId()`

UnsetOwnerLegalEntityId ensures that no value is present for OwnerLegalEntityId, not even an explicit nil
### GetOwnerCompanyId

`func (o *OwnershipObservationsRow) GetOwnerCompanyId() string`

GetOwnerCompanyId returns the OwnerCompanyId field if non-nil, zero value otherwise.

### GetOwnerCompanyIdOk

`func (o *OwnershipObservationsRow) GetOwnerCompanyIdOk() (*string, bool)`

GetOwnerCompanyIdOk returns a tuple with the OwnerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerCompanyId

`func (o *OwnershipObservationsRow) SetOwnerCompanyId(v string)`

SetOwnerCompanyId sets OwnerCompanyId field to given value.

### HasOwnerCompanyId

`func (o *OwnershipObservationsRow) HasOwnerCompanyId() bool`

HasOwnerCompanyId returns a boolean if a field has been set.

### SetOwnerCompanyIdNil

`func (o *OwnershipObservationsRow) SetOwnerCompanyIdNil(b bool)`

 SetOwnerCompanyIdNil sets the value for OwnerCompanyId to be an explicit nil

### UnsetOwnerCompanyId
`func (o *OwnershipObservationsRow) UnsetOwnerCompanyId()`

UnsetOwnerCompanyId ensures that no value is present for OwnerCompanyId, not even an explicit nil
### GetOwnerCompanyName

`func (o *OwnershipObservationsRow) GetOwnerCompanyName() string`

GetOwnerCompanyName returns the OwnerCompanyName field if non-nil, zero value otherwise.

### GetOwnerCompanyNameOk

`func (o *OwnershipObservationsRow) GetOwnerCompanyNameOk() (*string, bool)`

GetOwnerCompanyNameOk returns a tuple with the OwnerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerCompanyName

`func (o *OwnershipObservationsRow) SetOwnerCompanyName(v string)`

SetOwnerCompanyName sets OwnerCompanyName field to given value.

### HasOwnerCompanyName

`func (o *OwnershipObservationsRow) HasOwnerCompanyName() bool`

HasOwnerCompanyName returns a boolean if a field has been set.

### SetOwnerCompanyNameNil

`func (o *OwnershipObservationsRow) SetOwnerCompanyNameNil(b bool)`

 SetOwnerCompanyNameNil sets the value for OwnerCompanyName to be an explicit nil

### UnsetOwnerCompanyName
`func (o *OwnershipObservationsRow) UnsetOwnerCompanyName()`

UnsetOwnerCompanyName ensures that no value is present for OwnerCompanyName, not even an explicit nil
### GetOwnerGroupId

`func (o *OwnershipObservationsRow) GetOwnerGroupId() string`

GetOwnerGroupId returns the OwnerGroupId field if non-nil, zero value otherwise.

### GetOwnerGroupIdOk

`func (o *OwnershipObservationsRow) GetOwnerGroupIdOk() (*string, bool)`

GetOwnerGroupIdOk returns a tuple with the OwnerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupId

`func (o *OwnershipObservationsRow) SetOwnerGroupId(v string)`

SetOwnerGroupId sets OwnerGroupId field to given value.

### HasOwnerGroupId

`func (o *OwnershipObservationsRow) HasOwnerGroupId() bool`

HasOwnerGroupId returns a boolean if a field has been set.

### SetOwnerGroupIdNil

`func (o *OwnershipObservationsRow) SetOwnerGroupIdNil(b bool)`

 SetOwnerGroupIdNil sets the value for OwnerGroupId to be an explicit nil

### UnsetOwnerGroupId
`func (o *OwnershipObservationsRow) UnsetOwnerGroupId()`

UnsetOwnerGroupId ensures that no value is present for OwnerGroupId, not even an explicit nil
### GetOwnerGroupName

`func (o *OwnershipObservationsRow) GetOwnerGroupName() string`

GetOwnerGroupName returns the OwnerGroupName field if non-nil, zero value otherwise.

### GetOwnerGroupNameOk

`func (o *OwnershipObservationsRow) GetOwnerGroupNameOk() (*string, bool)`

GetOwnerGroupNameOk returns a tuple with the OwnerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupName

`func (o *OwnershipObservationsRow) SetOwnerGroupName(v string)`

SetOwnerGroupName sets OwnerGroupName field to given value.

### HasOwnerGroupName

`func (o *OwnershipObservationsRow) HasOwnerGroupName() bool`

HasOwnerGroupName returns a boolean if a field has been set.

### SetOwnerGroupNameNil

`func (o *OwnershipObservationsRow) SetOwnerGroupNameNil(b bool)`

 SetOwnerGroupNameNil sets the value for OwnerGroupName to be an explicit nil

### UnsetOwnerGroupName
`func (o *OwnershipObservationsRow) UnsetOwnerGroupName()`

UnsetOwnerGroupName ensures that no value is present for OwnerGroupName, not even an explicit nil
### GetOwnerOrgId

`func (o *OwnershipObservationsRow) GetOwnerOrgId() string`

GetOwnerOrgId returns the OwnerOrgId field if non-nil, zero value otherwise.

### GetOwnerOrgIdOk

`func (o *OwnershipObservationsRow) GetOwnerOrgIdOk() (*string, bool)`

GetOwnerOrgIdOk returns a tuple with the OwnerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrgId

`func (o *OwnershipObservationsRow) SetOwnerOrgId(v string)`

SetOwnerOrgId sets OwnerOrgId field to given value.

### HasOwnerOrgId

`func (o *OwnershipObservationsRow) HasOwnerOrgId() bool`

HasOwnerOrgId returns a boolean if a field has been set.

### SetOwnerOrgIdNil

`func (o *OwnershipObservationsRow) SetOwnerOrgIdNil(b bool)`

 SetOwnerOrgIdNil sets the value for OwnerOrgId to be an explicit nil

### UnsetOwnerOrgId
`func (o *OwnershipObservationsRow) UnsetOwnerOrgId()`

UnsetOwnerOrgId ensures that no value is present for OwnerOrgId, not even an explicit nil
### GetOwnerInvestorId

`func (o *OwnershipObservationsRow) GetOwnerInvestorId() string`

GetOwnerInvestorId returns the OwnerInvestorId field if non-nil, zero value otherwise.

### GetOwnerInvestorIdOk

`func (o *OwnershipObservationsRow) GetOwnerInvestorIdOk() (*string, bool)`

GetOwnerInvestorIdOk returns a tuple with the OwnerInvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInvestorId

`func (o *OwnershipObservationsRow) SetOwnerInvestorId(v string)`

SetOwnerInvestorId sets OwnerInvestorId field to given value.

### HasOwnerInvestorId

`func (o *OwnershipObservationsRow) HasOwnerInvestorId() bool`

HasOwnerInvestorId returns a boolean if a field has been set.

### SetOwnerInvestorIdNil

`func (o *OwnershipObservationsRow) SetOwnerInvestorIdNil(b bool)`

 SetOwnerInvestorIdNil sets the value for OwnerInvestorId to be an explicit nil

### UnsetOwnerInvestorId
`func (o *OwnershipObservationsRow) UnsetOwnerInvestorId()`

UnsetOwnerInvestorId ensures that no value is present for OwnerInvestorId, not even an explicit nil
### GetOwnerPersonId

`func (o *OwnershipObservationsRow) GetOwnerPersonId() string`

GetOwnerPersonId returns the OwnerPersonId field if non-nil, zero value otherwise.

### GetOwnerPersonIdOk

`func (o *OwnershipObservationsRow) GetOwnerPersonIdOk() (*string, bool)`

GetOwnerPersonIdOk returns a tuple with the OwnerPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPersonId

`func (o *OwnershipObservationsRow) SetOwnerPersonId(v string)`

SetOwnerPersonId sets OwnerPersonId field to given value.

### HasOwnerPersonId

`func (o *OwnershipObservationsRow) HasOwnerPersonId() bool`

HasOwnerPersonId returns a boolean if a field has been set.

### SetOwnerPersonIdNil

`func (o *OwnershipObservationsRow) SetOwnerPersonIdNil(b bool)`

 SetOwnerPersonIdNil sets the value for OwnerPersonId to be an explicit nil

### UnsetOwnerPersonId
`func (o *OwnershipObservationsRow) UnsetOwnerPersonId()`

UnsetOwnerPersonId ensures that no value is present for OwnerPersonId, not even an explicit nil
### GetPersonLinkMethod

`func (o *OwnershipObservationsRow) GetPersonLinkMethod() string`

GetPersonLinkMethod returns the PersonLinkMethod field if non-nil, zero value otherwise.

### GetPersonLinkMethodOk

`func (o *OwnershipObservationsRow) GetPersonLinkMethodOk() (*string, bool)`

GetPersonLinkMethodOk returns a tuple with the PersonLinkMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonLinkMethod

`func (o *OwnershipObservationsRow) SetPersonLinkMethod(v string)`

SetPersonLinkMethod sets PersonLinkMethod field to given value.

### HasPersonLinkMethod

`func (o *OwnershipObservationsRow) HasPersonLinkMethod() bool`

HasPersonLinkMethod returns a boolean if a field has been set.

### SetPersonLinkMethodNil

`func (o *OwnershipObservationsRow) SetPersonLinkMethodNil(b bool)`

 SetPersonLinkMethodNil sets the value for PersonLinkMethod to be an explicit nil

### UnsetPersonLinkMethod
`func (o *OwnershipObservationsRow) UnsetPersonLinkMethod()`

UnsetPersonLinkMethod ensures that no value is present for PersonLinkMethod, not even an explicit nil
### GetPersonLinkConflict

`func (o *OwnershipObservationsRow) GetPersonLinkConflict() bool`

GetPersonLinkConflict returns the PersonLinkConflict field if non-nil, zero value otherwise.

### GetPersonLinkConflictOk

`func (o *OwnershipObservationsRow) GetPersonLinkConflictOk() (*bool, bool)`

GetPersonLinkConflictOk returns a tuple with the PersonLinkConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonLinkConflict

`func (o *OwnershipObservationsRow) SetPersonLinkConflict(v bool)`

SetPersonLinkConflict sets PersonLinkConflict field to given value.

### HasPersonLinkConflict

`func (o *OwnershipObservationsRow) HasPersonLinkConflict() bool`

HasPersonLinkConflict returns a boolean if a field has been set.

### SetPersonLinkConflictNil

`func (o *OwnershipObservationsRow) SetPersonLinkConflictNil(b bool)`

 SetPersonLinkConflictNil sets the value for PersonLinkConflict to be an explicit nil

### UnsetPersonLinkConflict
`func (o *OwnershipObservationsRow) UnsetPersonLinkConflict()`

UnsetPersonLinkConflict ensures that no value is present for PersonLinkConflict, not even an explicit nil
### GetOwnerResolutionMethod

`func (o *OwnershipObservationsRow) GetOwnerResolutionMethod() string`

GetOwnerResolutionMethod returns the OwnerResolutionMethod field if non-nil, zero value otherwise.

### GetOwnerResolutionMethodOk

`func (o *OwnershipObservationsRow) GetOwnerResolutionMethodOk() (*string, bool)`

GetOwnerResolutionMethodOk returns a tuple with the OwnerResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerResolutionMethod

`func (o *OwnershipObservationsRow) SetOwnerResolutionMethod(v string)`

SetOwnerResolutionMethod sets OwnerResolutionMethod field to given value.

### HasOwnerResolutionMethod

`func (o *OwnershipObservationsRow) HasOwnerResolutionMethod() bool`

HasOwnerResolutionMethod returns a boolean if a field has been set.

### SetOwnerResolutionMethodNil

`func (o *OwnershipObservationsRow) SetOwnerResolutionMethodNil(b bool)`

 SetOwnerResolutionMethodNil sets the value for OwnerResolutionMethod to be an explicit nil

### UnsetOwnerResolutionMethod
`func (o *OwnershipObservationsRow) UnsetOwnerResolutionMethod()`

UnsetOwnerResolutionMethod ensures that no value is present for OwnerResolutionMethod, not even an explicit nil
### GetOwnerResolutionConfidence

`func (o *OwnershipObservationsRow) GetOwnerResolutionConfidence() float32`

GetOwnerResolutionConfidence returns the OwnerResolutionConfidence field if non-nil, zero value otherwise.

### GetOwnerResolutionConfidenceOk

`func (o *OwnershipObservationsRow) GetOwnerResolutionConfidenceOk() (*float32, bool)`

GetOwnerResolutionConfidenceOk returns a tuple with the OwnerResolutionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerResolutionConfidence

`func (o *OwnershipObservationsRow) SetOwnerResolutionConfidence(v float32)`

SetOwnerResolutionConfidence sets OwnerResolutionConfidence field to given value.

### HasOwnerResolutionConfidence

`func (o *OwnershipObservationsRow) HasOwnerResolutionConfidence() bool`

HasOwnerResolutionConfidence returns a boolean if a field has been set.

### SetOwnerResolutionConfidenceNil

`func (o *OwnershipObservationsRow) SetOwnerResolutionConfidenceNil(b bool)`

 SetOwnerResolutionConfidenceNil sets the value for OwnerResolutionConfidence to be an explicit nil

### UnsetOwnerResolutionConfidence
`func (o *OwnershipObservationsRow) UnsetOwnerResolutionConfidence()`

UnsetOwnerResolutionConfidence ensures that no value is present for OwnerResolutionConfidence, not even an explicit nil
### GetReviewState

`func (o *OwnershipObservationsRow) GetReviewState() string`

GetReviewState returns the ReviewState field if non-nil, zero value otherwise.

### GetReviewStateOk

`func (o *OwnershipObservationsRow) GetReviewStateOk() (*string, bool)`

GetReviewStateOk returns a tuple with the ReviewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewState

`func (o *OwnershipObservationsRow) SetReviewState(v string)`

SetReviewState sets ReviewState field to given value.

### HasReviewState

`func (o *OwnershipObservationsRow) HasReviewState() bool`

HasReviewState returns a boolean if a field has been set.

### SetReviewStateNil

`func (o *OwnershipObservationsRow) SetReviewStateNil(b bool)`

 SetReviewStateNil sets the value for ReviewState to be an explicit nil

### UnsetReviewState
`func (o *OwnershipObservationsRow) UnsetReviewState()`

UnsetReviewState ensures that no value is present for ReviewState, not even an explicit nil
### GetStakeValueType

`func (o *OwnershipObservationsRow) GetStakeValueType() string`

GetStakeValueType returns the StakeValueType field if non-nil, zero value otherwise.

### GetStakeValueTypeOk

`func (o *OwnershipObservationsRow) GetStakeValueTypeOk() (*string, bool)`

GetStakeValueTypeOk returns a tuple with the StakeValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeValueType

`func (o *OwnershipObservationsRow) SetStakeValueType(v string)`

SetStakeValueType sets StakeValueType field to given value.

### HasStakeValueType

`func (o *OwnershipObservationsRow) HasStakeValueType() bool`

HasStakeValueType returns a boolean if a field has been set.

### SetStakeValueTypeNil

`func (o *OwnershipObservationsRow) SetStakeValueTypeNil(b bool)`

 SetStakeValueTypeNil sets the value for StakeValueType to be an explicit nil

### UnsetStakeValueType
`func (o *OwnershipObservationsRow) UnsetStakeValueType()`

UnsetStakeValueType ensures that no value is present for StakeValueType, not even an explicit nil
### GetPercentageBasis

`func (o *OwnershipObservationsRow) GetPercentageBasis() string`

GetPercentageBasis returns the PercentageBasis field if non-nil, zero value otherwise.

### GetPercentageBasisOk

`func (o *OwnershipObservationsRow) GetPercentageBasisOk() (*string, bool)`

GetPercentageBasisOk returns a tuple with the PercentageBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageBasis

`func (o *OwnershipObservationsRow) SetPercentageBasis(v string)`

SetPercentageBasis sets PercentageBasis field to given value.

### HasPercentageBasis

`func (o *OwnershipObservationsRow) HasPercentageBasis() bool`

HasPercentageBasis returns a boolean if a field has been set.

### SetPercentageBasisNil

`func (o *OwnershipObservationsRow) SetPercentageBasisNil(b bool)`

 SetPercentageBasisNil sets the value for PercentageBasis to be an explicit nil

### UnsetPercentageBasis
`func (o *OwnershipObservationsRow) UnsetPercentageBasis()`

UnsetPercentageBasis ensures that no value is present for PercentageBasis, not even an explicit nil
### GetStakePercent

`func (o *OwnershipObservationsRow) GetStakePercent() float32`

GetStakePercent returns the StakePercent field if non-nil, zero value otherwise.

### GetStakePercentOk

`func (o *OwnershipObservationsRow) GetStakePercentOk() (*float32, bool)`

GetStakePercentOk returns a tuple with the StakePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakePercent

`func (o *OwnershipObservationsRow) SetStakePercent(v float32)`

SetStakePercent sets StakePercent field to given value.

### HasStakePercent

`func (o *OwnershipObservationsRow) HasStakePercent() bool`

HasStakePercent returns a boolean if a field has been set.

### SetStakePercentNil

`func (o *OwnershipObservationsRow) SetStakePercentNil(b bool)`

 SetStakePercentNil sets the value for StakePercent to be an explicit nil

### UnsetStakePercent
`func (o *OwnershipObservationsRow) UnsetStakePercent()`

UnsetStakePercent ensures that no value is present for StakePercent, not even an explicit nil
### GetStakeLowerPercent

`func (o *OwnershipObservationsRow) GetStakeLowerPercent() float32`

GetStakeLowerPercent returns the StakeLowerPercent field if non-nil, zero value otherwise.

### GetStakeLowerPercentOk

`func (o *OwnershipObservationsRow) GetStakeLowerPercentOk() (*float32, bool)`

GetStakeLowerPercentOk returns a tuple with the StakeLowerPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeLowerPercent

`func (o *OwnershipObservationsRow) SetStakeLowerPercent(v float32)`

SetStakeLowerPercent sets StakeLowerPercent field to given value.

### HasStakeLowerPercent

`func (o *OwnershipObservationsRow) HasStakeLowerPercent() bool`

HasStakeLowerPercent returns a boolean if a field has been set.

### SetStakeLowerPercentNil

`func (o *OwnershipObservationsRow) SetStakeLowerPercentNil(b bool)`

 SetStakeLowerPercentNil sets the value for StakeLowerPercent to be an explicit nil

### UnsetStakeLowerPercent
`func (o *OwnershipObservationsRow) UnsetStakeLowerPercent()`

UnsetStakeLowerPercent ensures that no value is present for StakeLowerPercent, not even an explicit nil
### GetStakeUpperPercent

`func (o *OwnershipObservationsRow) GetStakeUpperPercent() float32`

GetStakeUpperPercent returns the StakeUpperPercent field if non-nil, zero value otherwise.

### GetStakeUpperPercentOk

`func (o *OwnershipObservationsRow) GetStakeUpperPercentOk() (*float32, bool)`

GetStakeUpperPercentOk returns a tuple with the StakeUpperPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeUpperPercent

`func (o *OwnershipObservationsRow) SetStakeUpperPercent(v float32)`

SetStakeUpperPercent sets StakeUpperPercent field to given value.

### HasStakeUpperPercent

`func (o *OwnershipObservationsRow) HasStakeUpperPercent() bool`

HasStakeUpperPercent returns a boolean if a field has been set.

### SetStakeUpperPercentNil

`func (o *OwnershipObservationsRow) SetStakeUpperPercentNil(b bool)`

 SetStakeUpperPercentNil sets the value for StakeUpperPercent to be an explicit nil

### UnsetStakeUpperPercent
`func (o *OwnershipObservationsRow) UnsetStakeUpperPercent()`

UnsetStakeUpperPercent ensures that no value is present for StakeUpperPercent, not even an explicit nil
### GetStakeLowerInclusive

`func (o *OwnershipObservationsRow) GetStakeLowerInclusive() bool`

GetStakeLowerInclusive returns the StakeLowerInclusive field if non-nil, zero value otherwise.

### GetStakeLowerInclusiveOk

`func (o *OwnershipObservationsRow) GetStakeLowerInclusiveOk() (*bool, bool)`

GetStakeLowerInclusiveOk returns a tuple with the StakeLowerInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeLowerInclusive

`func (o *OwnershipObservationsRow) SetStakeLowerInclusive(v bool)`

SetStakeLowerInclusive sets StakeLowerInclusive field to given value.

### HasStakeLowerInclusive

`func (o *OwnershipObservationsRow) HasStakeLowerInclusive() bool`

HasStakeLowerInclusive returns a boolean if a field has been set.

### SetStakeLowerInclusiveNil

`func (o *OwnershipObservationsRow) SetStakeLowerInclusiveNil(b bool)`

 SetStakeLowerInclusiveNil sets the value for StakeLowerInclusive to be an explicit nil

### UnsetStakeLowerInclusive
`func (o *OwnershipObservationsRow) UnsetStakeLowerInclusive()`

UnsetStakeLowerInclusive ensures that no value is present for StakeLowerInclusive, not even an explicit nil
### GetStakeUpperInclusive

`func (o *OwnershipObservationsRow) GetStakeUpperInclusive() bool`

GetStakeUpperInclusive returns the StakeUpperInclusive field if non-nil, zero value otherwise.

### GetStakeUpperInclusiveOk

`func (o *OwnershipObservationsRow) GetStakeUpperInclusiveOk() (*bool, bool)`

GetStakeUpperInclusiveOk returns a tuple with the StakeUpperInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeUpperInclusive

`func (o *OwnershipObservationsRow) SetStakeUpperInclusive(v bool)`

SetStakeUpperInclusive sets StakeUpperInclusive field to given value.

### HasStakeUpperInclusive

`func (o *OwnershipObservationsRow) HasStakeUpperInclusive() bool`

HasStakeUpperInclusive returns a boolean if a field has been set.

### SetStakeUpperInclusiveNil

`func (o *OwnershipObservationsRow) SetStakeUpperInclusiveNil(b bool)`

 SetStakeUpperInclusiveNil sets the value for StakeUpperInclusive to be an explicit nil

### UnsetStakeUpperInclusive
`func (o *OwnershipObservationsRow) UnsetStakeUpperInclusive()`

UnsetStakeUpperInclusive ensures that no value is present for StakeUpperInclusive, not even an explicit nil
### GetStakeBandCode

`func (o *OwnershipObservationsRow) GetStakeBandCode() string`

GetStakeBandCode returns the StakeBandCode field if non-nil, zero value otherwise.

### GetStakeBandCodeOk

`func (o *OwnershipObservationsRow) GetStakeBandCodeOk() (*string, bool)`

GetStakeBandCodeOk returns a tuple with the StakeBandCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeBandCode

`func (o *OwnershipObservationsRow) SetStakeBandCode(v string)`

SetStakeBandCode sets StakeBandCode field to given value.

### HasStakeBandCode

`func (o *OwnershipObservationsRow) HasStakeBandCode() bool`

HasStakeBandCode returns a boolean if a field has been set.

### SetStakeBandCodeNil

`func (o *OwnershipObservationsRow) SetStakeBandCodeNil(b bool)`

 SetStakeBandCodeNil sets the value for StakeBandCode to be an explicit nil

### UnsetStakeBandCode
`func (o *OwnershipObservationsRow) UnsetStakeBandCode()`

UnsetStakeBandCode ensures that no value is present for StakeBandCode, not even an explicit nil
### GetStakeDisplay

`func (o *OwnershipObservationsRow) GetStakeDisplay() string`

GetStakeDisplay returns the StakeDisplay field if non-nil, zero value otherwise.

### GetStakeDisplayOk

`func (o *OwnershipObservationsRow) GetStakeDisplayOk() (*string, bool)`

GetStakeDisplayOk returns a tuple with the StakeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeDisplay

`func (o *OwnershipObservationsRow) SetStakeDisplay(v string)`

SetStakeDisplay sets StakeDisplay field to given value.

### HasStakeDisplay

`func (o *OwnershipObservationsRow) HasStakeDisplay() bool`

HasStakeDisplay returns a boolean if a field has been set.

### SetStakeDisplayNil

`func (o *OwnershipObservationsRow) SetStakeDisplayNil(b bool)`

 SetStakeDisplayNil sets the value for StakeDisplay to be an explicit nil

### UnsetStakeDisplay
`func (o *OwnershipObservationsRow) UnsetStakeDisplay()`

UnsetStakeDisplay ensures that no value is present for StakeDisplay, not even an explicit nil
### GetSecurityTitle

`func (o *OwnershipObservationsRow) GetSecurityTitle() string`

GetSecurityTitle returns the SecurityTitle field if non-nil, zero value otherwise.

### GetSecurityTitleOk

`func (o *OwnershipObservationsRow) GetSecurityTitleOk() (*string, bool)`

GetSecurityTitleOk returns a tuple with the SecurityTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityTitle

`func (o *OwnershipObservationsRow) SetSecurityTitle(v string)`

SetSecurityTitle sets SecurityTitle field to given value.

### HasSecurityTitle

`func (o *OwnershipObservationsRow) HasSecurityTitle() bool`

HasSecurityTitle returns a boolean if a field has been set.

### SetSecurityTitleNil

`func (o *OwnershipObservationsRow) SetSecurityTitleNil(b bool)`

 SetSecurityTitleNil sets the value for SecurityTitle to be an explicit nil

### UnsetSecurityTitle
`func (o *OwnershipObservationsRow) UnsetSecurityTitle()`

UnsetSecurityTitle ensures that no value is present for SecurityTitle, not even an explicit nil
### GetSharesAmount

`func (o *OwnershipObservationsRow) GetSharesAmount() float32`

GetSharesAmount returns the SharesAmount field if non-nil, zero value otherwise.

### GetSharesAmountOk

`func (o *OwnershipObservationsRow) GetSharesAmountOk() (*float32, bool)`

GetSharesAmountOk returns a tuple with the SharesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesAmount

`func (o *OwnershipObservationsRow) SetSharesAmount(v float32)`

SetSharesAmount sets SharesAmount field to given value.

### HasSharesAmount

`func (o *OwnershipObservationsRow) HasSharesAmount() bool`

HasSharesAmount returns a boolean if a field has been set.

### SetSharesAmountNil

`func (o *OwnershipObservationsRow) SetSharesAmountNil(b bool)`

 SetSharesAmountNil sets the value for SharesAmount to be an explicit nil

### UnsetSharesAmount
`func (o *OwnershipObservationsRow) UnsetSharesAmount()`

UnsetSharesAmount ensures that no value is present for SharesAmount, not even an explicit nil
### GetReportedValueAmount

`func (o *OwnershipObservationsRow) GetReportedValueAmount() float32`

GetReportedValueAmount returns the ReportedValueAmount field if non-nil, zero value otherwise.

### GetReportedValueAmountOk

`func (o *OwnershipObservationsRow) GetReportedValueAmountOk() (*float32, bool)`

GetReportedValueAmountOk returns a tuple with the ReportedValueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedValueAmount

`func (o *OwnershipObservationsRow) SetReportedValueAmount(v float32)`

SetReportedValueAmount sets ReportedValueAmount field to given value.

### HasReportedValueAmount

`func (o *OwnershipObservationsRow) HasReportedValueAmount() bool`

HasReportedValueAmount returns a boolean if a field has been set.

### SetReportedValueAmountNil

`func (o *OwnershipObservationsRow) SetReportedValueAmountNil(b bool)`

 SetReportedValueAmountNil sets the value for ReportedValueAmount to be an explicit nil

### UnsetReportedValueAmount
`func (o *OwnershipObservationsRow) UnsetReportedValueAmount()`

UnsetReportedValueAmount ensures that no value is present for ReportedValueAmount, not even an explicit nil
### GetReportedValueCurrency

`func (o *OwnershipObservationsRow) GetReportedValueCurrency() string`

GetReportedValueCurrency returns the ReportedValueCurrency field if non-nil, zero value otherwise.

### GetReportedValueCurrencyOk

`func (o *OwnershipObservationsRow) GetReportedValueCurrencyOk() (*string, bool)`

GetReportedValueCurrencyOk returns a tuple with the ReportedValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedValueCurrency

`func (o *OwnershipObservationsRow) SetReportedValueCurrency(v string)`

SetReportedValueCurrency sets ReportedValueCurrency field to given value.

### HasReportedValueCurrency

`func (o *OwnershipObservationsRow) HasReportedValueCurrency() bool`

HasReportedValueCurrency returns a boolean if a field has been set.

### SetReportedValueCurrencyNil

`func (o *OwnershipObservationsRow) SetReportedValueCurrencyNil(b bool)`

 SetReportedValueCurrencyNil sets the value for ReportedValueCurrency to be an explicit nil

### UnsetReportedValueCurrency
`func (o *OwnershipObservationsRow) UnsetReportedValueCurrency()`

UnsetReportedValueCurrency ensures that no value is present for ReportedValueCurrency, not even an explicit nil
### GetReportedValueScale

`func (o *OwnershipObservationsRow) GetReportedValueScale() string`

GetReportedValueScale returns the ReportedValueScale field if non-nil, zero value otherwise.

### GetReportedValueScaleOk

`func (o *OwnershipObservationsRow) GetReportedValueScaleOk() (*string, bool)`

GetReportedValueScaleOk returns a tuple with the ReportedValueScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedValueScale

`func (o *OwnershipObservationsRow) SetReportedValueScale(v string)`

SetReportedValueScale sets ReportedValueScale field to given value.

### HasReportedValueScale

`func (o *OwnershipObservationsRow) HasReportedValueScale() bool`

HasReportedValueScale returns a boolean if a field has been set.

### SetReportedValueScaleNil

`func (o *OwnershipObservationsRow) SetReportedValueScaleNil(b bool)`

 SetReportedValueScaleNil sets the value for ReportedValueScale to be an explicit nil

### UnsetReportedValueScale
`func (o *OwnershipObservationsRow) UnsetReportedValueScale()`

UnsetReportedValueScale ensures that no value is present for ReportedValueScale, not even an explicit nil
### GetAnnouncedDate

`func (o *OwnershipObservationsRow) GetAnnouncedDate() string`

GetAnnouncedDate returns the AnnouncedDate field if non-nil, zero value otherwise.

### GetAnnouncedDateOk

`func (o *OwnershipObservationsRow) GetAnnouncedDateOk() (*string, bool)`

GetAnnouncedDateOk returns a tuple with the AnnouncedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncedDate

`func (o *OwnershipObservationsRow) SetAnnouncedDate(v string)`

SetAnnouncedDate sets AnnouncedDate field to given value.

### HasAnnouncedDate

`func (o *OwnershipObservationsRow) HasAnnouncedDate() bool`

HasAnnouncedDate returns a boolean if a field has been set.

### SetAnnouncedDateNil

`func (o *OwnershipObservationsRow) SetAnnouncedDateNil(b bool)`

 SetAnnouncedDateNil sets the value for AnnouncedDate to be an explicit nil

### UnsetAnnouncedDate
`func (o *OwnershipObservationsRow) UnsetAnnouncedDate()`

UnsetAnnouncedDate ensures that no value is present for AnnouncedDate, not even an explicit nil
### GetEffectiveDate

`func (o *OwnershipObservationsRow) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *OwnershipObservationsRow) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *OwnershipObservationsRow) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *OwnershipObservationsRow) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### SetEffectiveDateNil

`func (o *OwnershipObservationsRow) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *OwnershipObservationsRow) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetFilingDate

`func (o *OwnershipObservationsRow) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *OwnershipObservationsRow) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *OwnershipObservationsRow) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *OwnershipObservationsRow) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### SetFilingDateNil

`func (o *OwnershipObservationsRow) SetFilingDateNil(b bool)`

 SetFilingDateNil sets the value for FilingDate to be an explicit nil

### UnsetFilingDate
`func (o *OwnershipObservationsRow) UnsetFilingDate()`

UnsetFilingDate ensures that no value is present for FilingDate, not even an explicit nil
### GetAsOfDate

`func (o *OwnershipObservationsRow) GetAsOfDate() string`

GetAsOfDate returns the AsOfDate field if non-nil, zero value otherwise.

### GetAsOfDateOk

`func (o *OwnershipObservationsRow) GetAsOfDateOk() (*string, bool)`

GetAsOfDateOk returns a tuple with the AsOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfDate

`func (o *OwnershipObservationsRow) SetAsOfDate(v string)`

SetAsOfDate sets AsOfDate field to given value.

### HasAsOfDate

`func (o *OwnershipObservationsRow) HasAsOfDate() bool`

HasAsOfDate returns a boolean if a field has been set.

### SetAsOfDateNil

`func (o *OwnershipObservationsRow) SetAsOfDateNil(b bool)`

 SetAsOfDateNil sets the value for AsOfDate to be an explicit nil

### UnsetAsOfDate
`func (o *OwnershipObservationsRow) UnsetAsOfDate()`

UnsetAsOfDate ensures that no value is present for AsOfDate, not even an explicit nil
### GetCeasedDate

`func (o *OwnershipObservationsRow) GetCeasedDate() string`

GetCeasedDate returns the CeasedDate field if non-nil, zero value otherwise.

### GetCeasedDateOk

`func (o *OwnershipObservationsRow) GetCeasedDateOk() (*string, bool)`

GetCeasedDateOk returns a tuple with the CeasedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeasedDate

`func (o *OwnershipObservationsRow) SetCeasedDate(v string)`

SetCeasedDate sets CeasedDate field to given value.

### HasCeasedDate

`func (o *OwnershipObservationsRow) HasCeasedDate() bool`

HasCeasedDate returns a boolean if a field has been set.

### SetCeasedDateNil

`func (o *OwnershipObservationsRow) SetCeasedDateNil(b bool)`

 SetCeasedDateNil sets the value for CeasedDate to be an explicit nil

### UnsetCeasedDate
`func (o *OwnershipObservationsRow) UnsetCeasedDate()`

UnsetCeasedDate ensures that no value is present for CeasedDate, not even an explicit nil
### GetObservationDate

`func (o *OwnershipObservationsRow) GetObservationDate() string`

GetObservationDate returns the ObservationDate field if non-nil, zero value otherwise.

### GetObservationDateOk

`func (o *OwnershipObservationsRow) GetObservationDateOk() (*string, bool)`

GetObservationDateOk returns a tuple with the ObservationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDate

`func (o *OwnershipObservationsRow) SetObservationDate(v string)`

SetObservationDate sets ObservationDate field to given value.

### HasObservationDate

`func (o *OwnershipObservationsRow) HasObservationDate() bool`

HasObservationDate returns a boolean if a field has been set.

### SetObservationDateNil

`func (o *OwnershipObservationsRow) SetObservationDateNil(b bool)`

 SetObservationDateNil sets the value for ObservationDate to be an explicit nil

### UnsetObservationDate
`func (o *OwnershipObservationsRow) UnsetObservationDate()`

UnsetObservationDate ensures that no value is present for ObservationDate, not even an explicit nil
### GetObservationDateBasis

`func (o *OwnershipObservationsRow) GetObservationDateBasis() string`

GetObservationDateBasis returns the ObservationDateBasis field if non-nil, zero value otherwise.

### GetObservationDateBasisOk

`func (o *OwnershipObservationsRow) GetObservationDateBasisOk() (*string, bool)`

GetObservationDateBasisOk returns a tuple with the ObservationDateBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDateBasis

`func (o *OwnershipObservationsRow) SetObservationDateBasis(v string)`

SetObservationDateBasis sets ObservationDateBasis field to given value.

### HasObservationDateBasis

`func (o *OwnershipObservationsRow) HasObservationDateBasis() bool`

HasObservationDateBasis returns a boolean if a field has been set.

### SetObservationDateBasisNil

`func (o *OwnershipObservationsRow) SetObservationDateBasisNil(b bool)`

 SetObservationDateBasisNil sets the value for ObservationDateBasis to be an explicit nil

### UnsetObservationDateBasis
`func (o *OwnershipObservationsRow) UnsetObservationDateBasis()`

UnsetObservationDateBasis ensures that no value is present for ObservationDateBasis, not even an explicit nil
### GetLifecycleStatus

`func (o *OwnershipObservationsRow) GetLifecycleStatus() string`

GetLifecycleStatus returns the LifecycleStatus field if non-nil, zero value otherwise.

### GetLifecycleStatusOk

`func (o *OwnershipObservationsRow) GetLifecycleStatusOk() (*string, bool)`

GetLifecycleStatusOk returns a tuple with the LifecycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStatus

`func (o *OwnershipObservationsRow) SetLifecycleStatus(v string)`

SetLifecycleStatus sets LifecycleStatus field to given value.

### HasLifecycleStatus

`func (o *OwnershipObservationsRow) HasLifecycleStatus() bool`

HasLifecycleStatus returns a boolean if a field has been set.

### SetLifecycleStatusNil

`func (o *OwnershipObservationsRow) SetLifecycleStatusNil(b bool)`

 SetLifecycleStatusNil sets the value for LifecycleStatus to be an explicit nil

### UnsetLifecycleStatus
`func (o *OwnershipObservationsRow) UnsetLifecycleStatus()`

UnsetLifecycleStatus ensures that no value is present for LifecycleStatus, not even an explicit nil
### GetSourceDocumentId

`func (o *OwnershipObservationsRow) GetSourceDocumentId() string`

GetSourceDocumentId returns the SourceDocumentId field if non-nil, zero value otherwise.

### GetSourceDocumentIdOk

`func (o *OwnershipObservationsRow) GetSourceDocumentIdOk() (*string, bool)`

GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentId

`func (o *OwnershipObservationsRow) SetSourceDocumentId(v string)`

SetSourceDocumentId sets SourceDocumentId field to given value.

### HasSourceDocumentId

`func (o *OwnershipObservationsRow) HasSourceDocumentId() bool`

HasSourceDocumentId returns a boolean if a field has been set.

### SetSourceDocumentIdNil

`func (o *OwnershipObservationsRow) SetSourceDocumentIdNil(b bool)`

 SetSourceDocumentIdNil sets the value for SourceDocumentId to be an explicit nil

### UnsetSourceDocumentId
`func (o *OwnershipObservationsRow) UnsetSourceDocumentId()`

UnsetSourceDocumentId ensures that no value is present for SourceDocumentId, not even an explicit nil
### GetSourceUrl

`func (o *OwnershipObservationsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *OwnershipObservationsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *OwnershipObservationsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *OwnershipObservationsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *OwnershipObservationsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *OwnershipObservationsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetSourceContentHash

`func (o *OwnershipObservationsRow) GetSourceContentHash() string`

GetSourceContentHash returns the SourceContentHash field if non-nil, zero value otherwise.

### GetSourceContentHashOk

`func (o *OwnershipObservationsRow) GetSourceContentHashOk() (*string, bool)`

GetSourceContentHashOk returns a tuple with the SourceContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContentHash

`func (o *OwnershipObservationsRow) SetSourceContentHash(v string)`

SetSourceContentHash sets SourceContentHash field to given value.

### HasSourceContentHash

`func (o *OwnershipObservationsRow) HasSourceContentHash() bool`

HasSourceContentHash returns a boolean if a field has been set.

### SetSourceContentHashNil

`func (o *OwnershipObservationsRow) SetSourceContentHashNil(b bool)`

 SetSourceContentHashNil sets the value for SourceContentHash to be an explicit nil

### UnsetSourceContentHash
`func (o *OwnershipObservationsRow) UnsetSourceContentHash()`

UnsetSourceContentHash ensures that no value is present for SourceContentHash, not even an explicit nil
### GetEvidenceQuote

`func (o *OwnershipObservationsRow) GetEvidenceQuote() string`

GetEvidenceQuote returns the EvidenceQuote field if non-nil, zero value otherwise.

### GetEvidenceQuoteOk

`func (o *OwnershipObservationsRow) GetEvidenceQuoteOk() (*string, bool)`

GetEvidenceQuoteOk returns a tuple with the EvidenceQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceQuote

`func (o *OwnershipObservationsRow) SetEvidenceQuote(v string)`

SetEvidenceQuote sets EvidenceQuote field to given value.

### HasEvidenceQuote

`func (o *OwnershipObservationsRow) HasEvidenceQuote() bool`

HasEvidenceQuote returns a boolean if a field has been set.

### SetEvidenceQuoteNil

`func (o *OwnershipObservationsRow) SetEvidenceQuoteNil(b bool)`

 SetEvidenceQuoteNil sets the value for EvidenceQuote to be an explicit nil

### UnsetEvidenceQuote
`func (o *OwnershipObservationsRow) UnsetEvidenceQuote()`

UnsetEvidenceQuote ensures that no value is present for EvidenceQuote, not even an explicit nil
### GetRestrictionTier

`func (o *OwnershipObservationsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *OwnershipObservationsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *OwnershipObservationsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *OwnershipObservationsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *OwnershipObservationsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *OwnershipObservationsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetSourceLicense

`func (o *OwnershipObservationsRow) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *OwnershipObservationsRow) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *OwnershipObservationsRow) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.

### HasSourceLicense

`func (o *OwnershipObservationsRow) HasSourceLicense() bool`

HasSourceLicense returns a boolean if a field has been set.

### SetSourceLicenseNil

`func (o *OwnershipObservationsRow) SetSourceLicenseNil(b bool)`

 SetSourceLicenseNil sets the value for SourceLicense to be an explicit nil

### UnsetSourceLicense
`func (o *OwnershipObservationsRow) UnsetSourceLicense()`

UnsetSourceLicense ensures that no value is present for SourceLicense, not even an explicit nil
### GetCommercialUseAllowed

`func (o *OwnershipObservationsRow) GetCommercialUseAllowed() bool`

GetCommercialUseAllowed returns the CommercialUseAllowed field if non-nil, zero value otherwise.

### GetCommercialUseAllowedOk

`func (o *OwnershipObservationsRow) GetCommercialUseAllowedOk() (*bool, bool)`

GetCommercialUseAllowedOk returns a tuple with the CommercialUseAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialUseAllowed

`func (o *OwnershipObservationsRow) SetCommercialUseAllowed(v bool)`

SetCommercialUseAllowed sets CommercialUseAllowed field to given value.

### HasCommercialUseAllowed

`func (o *OwnershipObservationsRow) HasCommercialUseAllowed() bool`

HasCommercialUseAllowed returns a boolean if a field has been set.

### SetCommercialUseAllowedNil

`func (o *OwnershipObservationsRow) SetCommercialUseAllowedNil(b bool)`

 SetCommercialUseAllowedNil sets the value for CommercialUseAllowed to be an explicit nil

### UnsetCommercialUseAllowed
`func (o *OwnershipObservationsRow) UnsetCommercialUseAllowed()`

UnsetCommercialUseAllowed ensures that no value is present for CommercialUseAllowed, not even an explicit nil
### GetAttribution

`func (o *OwnershipObservationsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OwnershipObservationsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OwnershipObservationsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *OwnershipObservationsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *OwnershipObservationsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *OwnershipObservationsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetContainsPersonalData

`func (o *OwnershipObservationsRow) GetContainsPersonalData() bool`

GetContainsPersonalData returns the ContainsPersonalData field if non-nil, zero value otherwise.

### GetContainsPersonalDataOk

`func (o *OwnershipObservationsRow) GetContainsPersonalDataOk() (*bool, bool)`

GetContainsPersonalDataOk returns a tuple with the ContainsPersonalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsPersonalData

`func (o *OwnershipObservationsRow) SetContainsPersonalData(v bool)`

SetContainsPersonalData sets ContainsPersonalData field to given value.

### HasContainsPersonalData

`func (o *OwnershipObservationsRow) HasContainsPersonalData() bool`

HasContainsPersonalData returns a boolean if a field has been set.

### SetContainsPersonalDataNil

`func (o *OwnershipObservationsRow) SetContainsPersonalDataNil(b bool)`

 SetContainsPersonalDataNil sets the value for ContainsPersonalData to be an explicit nil

### UnsetContainsPersonalData
`func (o *OwnershipObservationsRow) UnsetContainsPersonalData()`

UnsetContainsPersonalData ensures that no value is present for ContainsPersonalData, not even an explicit nil
### GetPrivacyStatus

`func (o *OwnershipObservationsRow) GetPrivacyStatus() string`

GetPrivacyStatus returns the PrivacyStatus field if non-nil, zero value otherwise.

### GetPrivacyStatusOk

`func (o *OwnershipObservationsRow) GetPrivacyStatusOk() (*string, bool)`

GetPrivacyStatusOk returns a tuple with the PrivacyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyStatus

`func (o *OwnershipObservationsRow) SetPrivacyStatus(v string)`

SetPrivacyStatus sets PrivacyStatus field to given value.

### HasPrivacyStatus

`func (o *OwnershipObservationsRow) HasPrivacyStatus() bool`

HasPrivacyStatus returns a boolean if a field has been set.

### SetPrivacyStatusNil

`func (o *OwnershipObservationsRow) SetPrivacyStatusNil(b bool)`

 SetPrivacyStatusNil sets the value for PrivacyStatus to be an explicit nil

### UnsetPrivacyStatus
`func (o *OwnershipObservationsRow) UnsetPrivacyStatus()`

UnsetPrivacyStatus ensures that no value is present for PrivacyStatus, not even an explicit nil
### GetIsCurrentEdge

`func (o *OwnershipObservationsRow) GetIsCurrentEdge() bool`

GetIsCurrentEdge returns the IsCurrentEdge field if non-nil, zero value otherwise.

### GetIsCurrentEdgeOk

`func (o *OwnershipObservationsRow) GetIsCurrentEdgeOk() (*bool, bool)`

GetIsCurrentEdgeOk returns a tuple with the IsCurrentEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrentEdge

`func (o *OwnershipObservationsRow) SetIsCurrentEdge(v bool)`

SetIsCurrentEdge sets IsCurrentEdge field to given value.

### HasIsCurrentEdge

`func (o *OwnershipObservationsRow) HasIsCurrentEdge() bool`

HasIsCurrentEdge returns a boolean if a field has been set.

### SetIsCurrentEdgeNil

`func (o *OwnershipObservationsRow) SetIsCurrentEdgeNil(b bool)`

 SetIsCurrentEdgeNil sets the value for IsCurrentEdge to be an explicit nil

### UnsetIsCurrentEdge
`func (o *OwnershipObservationsRow) UnsetIsCurrentEdge()`

UnsetIsCurrentEdge ensures that no value is present for IsCurrentEdge, not even an explicit nil
### GetCurrentEdgeState

`func (o *OwnershipObservationsRow) GetCurrentEdgeState() string`

GetCurrentEdgeState returns the CurrentEdgeState field if non-nil, zero value otherwise.

### GetCurrentEdgeStateOk

`func (o *OwnershipObservationsRow) GetCurrentEdgeStateOk() (*string, bool)`

GetCurrentEdgeStateOk returns a tuple with the CurrentEdgeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEdgeState

`func (o *OwnershipObservationsRow) SetCurrentEdgeState(v string)`

SetCurrentEdgeState sets CurrentEdgeState field to given value.

### HasCurrentEdgeState

`func (o *OwnershipObservationsRow) HasCurrentEdgeState() bool`

HasCurrentEdgeState returns a boolean if a field has been set.

### SetCurrentEdgeStateNil

`func (o *OwnershipObservationsRow) SetCurrentEdgeStateNil(b bool)`

 SetCurrentEdgeStateNil sets the value for CurrentEdgeState to be an explicit nil

### UnsetCurrentEdgeState
`func (o *OwnershipObservationsRow) UnsetCurrentEdgeState()`

UnsetCurrentEdgeState ensures that no value is present for CurrentEdgeState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


