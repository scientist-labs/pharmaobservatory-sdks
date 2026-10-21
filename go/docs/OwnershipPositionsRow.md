# OwnershipPositionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionId** | Pointer to **NullableString** | The position id (the owner rollup id). | [optional] 
**SourceClass** | Pointer to **NullableString** | The filing family: sec_forms_3_4_5 (insiders), sec_schedules_13d_g (beneficial owners), sec_form_adv (investment adviser control persons), sec_def_14a (proxy-statement tables) or wikidata (ownership statements in Wikidata). | [optional] 
**OwnershipDimension** | Pointer to **NullableString** | What the position measures: insider_reporting, beneficial_ownership, voting_power, economic_interest or control. Positions of different dimensions are different measurements and are not added. | [optional] 
**RelationshipType** | Pointer to **NullableString** | The owner&#39;s relationship to the company as the filing family states it. | [optional] 
**SubjectName** | Pointer to **NullableString** | The owned company as filed. | [optional] 
**SubjectOrgId** | Pointer to **NullableString** | Legal-entity level: the owned company&#39;s organization record. | [optional] 
**SubjectCompanyId** | Pointer to **NullableString** | Company level: the owned company (null when unresolved). | [optional] 
**SubjectCompanyName** | Pointer to **NullableString** | Display name of the owned company. | [optional] 
**SubjectGroupId** | Pointer to **NullableString** | Group level: the owned company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**SubjectGroupName** | Pointer to **NullableString** | Display name of the owned company&#39;s group. | [optional] 
**SubjectInvestorId** | Pointer to **NullableString** | The key of the owned party in the investor directory when it is an investor (null otherwise). Filter on it for who is reported to own or control an investor. | [optional] 
**OwnerName** | Pointer to **NullableString** | The owner&#39;s name as filed. | [optional] 
**OwnerNameIsPublic** | Pointer to **NullableBool** | A label, not a gate: true when the filing marks the owner&#39;s privacy status public. It does not affect any other column; the owner&#39;s name, identity key and every other field below are populated whatever this value is. | [optional] 
**OwnerKind** | Pointer to **NullableString** | What the owner is: organization, natural_person, reporting_group or unknown. | [optional] 
**OwnerIdentityKey** | Pointer to **NullableString** | The owner&#39;s stable identity (an SEC registration number or an adviser registration number). | [optional] 
**OwnerPersonId** | Pointer to **NullableString** | The key of an owner who is a natural person; null for other owners. | [optional] 
**OwnerInvestorId** | Pointer to **NullableString** | The key of the owner in the investor directory when it is one. | [optional] 
**OwnerOrgId** | Pointer to **NullableString** | Legal-entity level: the owner&#39;s organization record (null when the owner is not an organization or is unresolved). | [optional] 
**OwnerCompanyId** | Pointer to **NullableString** | Company level: the owner&#39;s company (null when unresolved). | [optional] 
**OwnerCompanyName** | Pointer to **NullableString** | Display name of the owner&#39;s company. | [optional] 
**OwnerGroupId** | Pointer to **NullableString** | Group level: the owner&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**OwnerGroupName** | Pointer to **NullableString** | Display name of the owner&#39;s group. | [optional] 
**SharesBasis** | Pointer to **NullableString** | How the share figure was reached: sum_standing_balances (insider balances), group_reported_maximum (a joint filing&#39;s maximum, never a sum of its filers) or no_quantity (the source gives none). Read a share figure with its basis. | [optional] 
**Shares** | Pointer to **NullableFloat32** | Shares the owner holds, on the shares_basis. Null when the source gives no quantity. | [optional] 
**SharesDirect** | Pointer to **NullableFloat32** | Of those shares, the part held directly. | [optional] 
**SharesIndirect** | Pointer to **NullableFloat32** | Of those shares, the part held indirectly. | [optional] 
**DerivativeUnderlyingShares** | Pointer to **NullableFloat32** | Shares underlying derivative positions such as options, reported apart and never added to shares. | [optional] 
**PercentBasis** | Pointer to **NullableString** | How the percentage was reached, for example single_reporting_person, uniform_group_percent or single_band. | [optional] 
**ReportedPercent** | Pointer to **NullableFloat32** | The percentage the filing reports, on the percentage_basis. | [optional] 
**ReportedPercentLow** | Pointer to **NullableFloat32** | The low end of a reported band, for a source that reports a band. | [optional] 
**ReportedPercentHigh** | Pointer to **NullableFloat32** | The high end of a reported band, for a source that reports a band. | [optional] 
**PercentageBasis** | Pointer to **NullableString** | What the percentage is of: beneficial_ownership, voting_power or economic_interest. Percentages of different bases are not comparable. | [optional] 
**EdgeState** | Pointer to **NullableString** | standing, ceased or no_longer_reported, for the owner&#39;s position as a whole. | [optional] 
**PositionCount** | Pointer to **NullableInt32** | Positions the row consolidates. | [optional] 
**StandingPositionCount** | Pointer to **NullableInt32** | Of those, positions still standing. | [optional] 
**CeasedPositionCount** | Pointer to **NullableInt32** | Of those, positions reported ceased. | [optional] 
**NoLongerReportedPositionCount** | Pointer to **NullableInt32** | Of those, positions the filer no longer reports. | [optional] 
**FirstObservedDate** | Pointer to **NullableString** | Date the position was first reported. | [optional] 
**LastActivityDate** | Pointer to **NullableString** | Date of the latest reported activity. | [optional] 
**LastFilingDate** | Pointer to **NullableString** | Date of the latest filing. | [optional] 
**SourceUrl** | Pointer to **NullableString** | A link to the latest filing. | [optional] 
**ContainsPersonalData** | Pointer to **NullableBool** | True when the row names a natural person. | [optional] 
**PrivacyStatus** | Pointer to **NullableString** | The filing&#39;s privacy status for the owner: not_applicable, pending_review, approved_public_fields, restricted or protected. | [optional] 

## Methods

### NewOwnershipPositionsRow

`func NewOwnershipPositionsRow() *OwnershipPositionsRow`

NewOwnershipPositionsRow instantiates a new OwnershipPositionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnershipPositionsRowWithDefaults

`func NewOwnershipPositionsRowWithDefaults() *OwnershipPositionsRow`

NewOwnershipPositionsRowWithDefaults instantiates a new OwnershipPositionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *OwnershipPositionsRow) GetPositionId() string`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *OwnershipPositionsRow) GetPositionIdOk() (*string, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *OwnershipPositionsRow) SetPositionId(v string)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *OwnershipPositionsRow) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### SetPositionIdNil

`func (o *OwnershipPositionsRow) SetPositionIdNil(b bool)`

 SetPositionIdNil sets the value for PositionId to be an explicit nil

### UnsetPositionId
`func (o *OwnershipPositionsRow) UnsetPositionId()`

UnsetPositionId ensures that no value is present for PositionId, not even an explicit nil
### GetSourceClass

`func (o *OwnershipPositionsRow) GetSourceClass() string`

GetSourceClass returns the SourceClass field if non-nil, zero value otherwise.

### GetSourceClassOk

`func (o *OwnershipPositionsRow) GetSourceClassOk() (*string, bool)`

GetSourceClassOk returns a tuple with the SourceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClass

`func (o *OwnershipPositionsRow) SetSourceClass(v string)`

SetSourceClass sets SourceClass field to given value.

### HasSourceClass

`func (o *OwnershipPositionsRow) HasSourceClass() bool`

HasSourceClass returns a boolean if a field has been set.

### SetSourceClassNil

`func (o *OwnershipPositionsRow) SetSourceClassNil(b bool)`

 SetSourceClassNil sets the value for SourceClass to be an explicit nil

### UnsetSourceClass
`func (o *OwnershipPositionsRow) UnsetSourceClass()`

UnsetSourceClass ensures that no value is present for SourceClass, not even an explicit nil
### GetOwnershipDimension

`func (o *OwnershipPositionsRow) GetOwnershipDimension() string`

GetOwnershipDimension returns the OwnershipDimension field if non-nil, zero value otherwise.

### GetOwnershipDimensionOk

`func (o *OwnershipPositionsRow) GetOwnershipDimensionOk() (*string, bool)`

GetOwnershipDimensionOk returns a tuple with the OwnershipDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipDimension

`func (o *OwnershipPositionsRow) SetOwnershipDimension(v string)`

SetOwnershipDimension sets OwnershipDimension field to given value.

### HasOwnershipDimension

`func (o *OwnershipPositionsRow) HasOwnershipDimension() bool`

HasOwnershipDimension returns a boolean if a field has been set.

### SetOwnershipDimensionNil

`func (o *OwnershipPositionsRow) SetOwnershipDimensionNil(b bool)`

 SetOwnershipDimensionNil sets the value for OwnershipDimension to be an explicit nil

### UnsetOwnershipDimension
`func (o *OwnershipPositionsRow) UnsetOwnershipDimension()`

UnsetOwnershipDimension ensures that no value is present for OwnershipDimension, not even an explicit nil
### GetRelationshipType

`func (o *OwnershipPositionsRow) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *OwnershipPositionsRow) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *OwnershipPositionsRow) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.

### HasRelationshipType

`func (o *OwnershipPositionsRow) HasRelationshipType() bool`

HasRelationshipType returns a boolean if a field has been set.

### SetRelationshipTypeNil

`func (o *OwnershipPositionsRow) SetRelationshipTypeNil(b bool)`

 SetRelationshipTypeNil sets the value for RelationshipType to be an explicit nil

### UnsetRelationshipType
`func (o *OwnershipPositionsRow) UnsetRelationshipType()`

UnsetRelationshipType ensures that no value is present for RelationshipType, not even an explicit nil
### GetSubjectName

`func (o *OwnershipPositionsRow) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *OwnershipPositionsRow) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *OwnershipPositionsRow) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *OwnershipPositionsRow) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### SetSubjectNameNil

`func (o *OwnershipPositionsRow) SetSubjectNameNil(b bool)`

 SetSubjectNameNil sets the value for SubjectName to be an explicit nil

### UnsetSubjectName
`func (o *OwnershipPositionsRow) UnsetSubjectName()`

UnsetSubjectName ensures that no value is present for SubjectName, not even an explicit nil
### GetSubjectOrgId

`func (o *OwnershipPositionsRow) GetSubjectOrgId() string`

GetSubjectOrgId returns the SubjectOrgId field if non-nil, zero value otherwise.

### GetSubjectOrgIdOk

`func (o *OwnershipPositionsRow) GetSubjectOrgIdOk() (*string, bool)`

GetSubjectOrgIdOk returns a tuple with the SubjectOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectOrgId

`func (o *OwnershipPositionsRow) SetSubjectOrgId(v string)`

SetSubjectOrgId sets SubjectOrgId field to given value.

### HasSubjectOrgId

`func (o *OwnershipPositionsRow) HasSubjectOrgId() bool`

HasSubjectOrgId returns a boolean if a field has been set.

### SetSubjectOrgIdNil

`func (o *OwnershipPositionsRow) SetSubjectOrgIdNil(b bool)`

 SetSubjectOrgIdNil sets the value for SubjectOrgId to be an explicit nil

### UnsetSubjectOrgId
`func (o *OwnershipPositionsRow) UnsetSubjectOrgId()`

UnsetSubjectOrgId ensures that no value is present for SubjectOrgId, not even an explicit nil
### GetSubjectCompanyId

`func (o *OwnershipPositionsRow) GetSubjectCompanyId() string`

GetSubjectCompanyId returns the SubjectCompanyId field if non-nil, zero value otherwise.

### GetSubjectCompanyIdOk

`func (o *OwnershipPositionsRow) GetSubjectCompanyIdOk() (*string, bool)`

GetSubjectCompanyIdOk returns a tuple with the SubjectCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyId

`func (o *OwnershipPositionsRow) SetSubjectCompanyId(v string)`

SetSubjectCompanyId sets SubjectCompanyId field to given value.

### HasSubjectCompanyId

`func (o *OwnershipPositionsRow) HasSubjectCompanyId() bool`

HasSubjectCompanyId returns a boolean if a field has been set.

### SetSubjectCompanyIdNil

`func (o *OwnershipPositionsRow) SetSubjectCompanyIdNil(b bool)`

 SetSubjectCompanyIdNil sets the value for SubjectCompanyId to be an explicit nil

### UnsetSubjectCompanyId
`func (o *OwnershipPositionsRow) UnsetSubjectCompanyId()`

UnsetSubjectCompanyId ensures that no value is present for SubjectCompanyId, not even an explicit nil
### GetSubjectCompanyName

`func (o *OwnershipPositionsRow) GetSubjectCompanyName() string`

GetSubjectCompanyName returns the SubjectCompanyName field if non-nil, zero value otherwise.

### GetSubjectCompanyNameOk

`func (o *OwnershipPositionsRow) GetSubjectCompanyNameOk() (*string, bool)`

GetSubjectCompanyNameOk returns a tuple with the SubjectCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyName

`func (o *OwnershipPositionsRow) SetSubjectCompanyName(v string)`

SetSubjectCompanyName sets SubjectCompanyName field to given value.

### HasSubjectCompanyName

`func (o *OwnershipPositionsRow) HasSubjectCompanyName() bool`

HasSubjectCompanyName returns a boolean if a field has been set.

### SetSubjectCompanyNameNil

`func (o *OwnershipPositionsRow) SetSubjectCompanyNameNil(b bool)`

 SetSubjectCompanyNameNil sets the value for SubjectCompanyName to be an explicit nil

### UnsetSubjectCompanyName
`func (o *OwnershipPositionsRow) UnsetSubjectCompanyName()`

UnsetSubjectCompanyName ensures that no value is present for SubjectCompanyName, not even an explicit nil
### GetSubjectGroupId

`func (o *OwnershipPositionsRow) GetSubjectGroupId() string`

GetSubjectGroupId returns the SubjectGroupId field if non-nil, zero value otherwise.

### GetSubjectGroupIdOk

`func (o *OwnershipPositionsRow) GetSubjectGroupIdOk() (*string, bool)`

GetSubjectGroupIdOk returns a tuple with the SubjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupId

`func (o *OwnershipPositionsRow) SetSubjectGroupId(v string)`

SetSubjectGroupId sets SubjectGroupId field to given value.

### HasSubjectGroupId

`func (o *OwnershipPositionsRow) HasSubjectGroupId() bool`

HasSubjectGroupId returns a boolean if a field has been set.

### SetSubjectGroupIdNil

`func (o *OwnershipPositionsRow) SetSubjectGroupIdNil(b bool)`

 SetSubjectGroupIdNil sets the value for SubjectGroupId to be an explicit nil

### UnsetSubjectGroupId
`func (o *OwnershipPositionsRow) UnsetSubjectGroupId()`

UnsetSubjectGroupId ensures that no value is present for SubjectGroupId, not even an explicit nil
### GetSubjectGroupName

`func (o *OwnershipPositionsRow) GetSubjectGroupName() string`

GetSubjectGroupName returns the SubjectGroupName field if non-nil, zero value otherwise.

### GetSubjectGroupNameOk

`func (o *OwnershipPositionsRow) GetSubjectGroupNameOk() (*string, bool)`

GetSubjectGroupNameOk returns a tuple with the SubjectGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupName

`func (o *OwnershipPositionsRow) SetSubjectGroupName(v string)`

SetSubjectGroupName sets SubjectGroupName field to given value.

### HasSubjectGroupName

`func (o *OwnershipPositionsRow) HasSubjectGroupName() bool`

HasSubjectGroupName returns a boolean if a field has been set.

### SetSubjectGroupNameNil

`func (o *OwnershipPositionsRow) SetSubjectGroupNameNil(b bool)`

 SetSubjectGroupNameNil sets the value for SubjectGroupName to be an explicit nil

### UnsetSubjectGroupName
`func (o *OwnershipPositionsRow) UnsetSubjectGroupName()`

UnsetSubjectGroupName ensures that no value is present for SubjectGroupName, not even an explicit nil
### GetSubjectInvestorId

`func (o *OwnershipPositionsRow) GetSubjectInvestorId() string`

GetSubjectInvestorId returns the SubjectInvestorId field if non-nil, zero value otherwise.

### GetSubjectInvestorIdOk

`func (o *OwnershipPositionsRow) GetSubjectInvestorIdOk() (*string, bool)`

GetSubjectInvestorIdOk returns a tuple with the SubjectInvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectInvestorId

`func (o *OwnershipPositionsRow) SetSubjectInvestorId(v string)`

SetSubjectInvestorId sets SubjectInvestorId field to given value.

### HasSubjectInvestorId

`func (o *OwnershipPositionsRow) HasSubjectInvestorId() bool`

HasSubjectInvestorId returns a boolean if a field has been set.

### SetSubjectInvestorIdNil

`func (o *OwnershipPositionsRow) SetSubjectInvestorIdNil(b bool)`

 SetSubjectInvestorIdNil sets the value for SubjectInvestorId to be an explicit nil

### UnsetSubjectInvestorId
`func (o *OwnershipPositionsRow) UnsetSubjectInvestorId()`

UnsetSubjectInvestorId ensures that no value is present for SubjectInvestorId, not even an explicit nil
### GetOwnerName

`func (o *OwnershipPositionsRow) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *OwnershipPositionsRow) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *OwnershipPositionsRow) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *OwnershipPositionsRow) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *OwnershipPositionsRow) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *OwnershipPositionsRow) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
### GetOwnerNameIsPublic

`func (o *OwnershipPositionsRow) GetOwnerNameIsPublic() bool`

GetOwnerNameIsPublic returns the OwnerNameIsPublic field if non-nil, zero value otherwise.

### GetOwnerNameIsPublicOk

`func (o *OwnershipPositionsRow) GetOwnerNameIsPublicOk() (*bool, bool)`

GetOwnerNameIsPublicOk returns a tuple with the OwnerNameIsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNameIsPublic

`func (o *OwnershipPositionsRow) SetOwnerNameIsPublic(v bool)`

SetOwnerNameIsPublic sets OwnerNameIsPublic field to given value.

### HasOwnerNameIsPublic

`func (o *OwnershipPositionsRow) HasOwnerNameIsPublic() bool`

HasOwnerNameIsPublic returns a boolean if a field has been set.

### SetOwnerNameIsPublicNil

`func (o *OwnershipPositionsRow) SetOwnerNameIsPublicNil(b bool)`

 SetOwnerNameIsPublicNil sets the value for OwnerNameIsPublic to be an explicit nil

### UnsetOwnerNameIsPublic
`func (o *OwnershipPositionsRow) UnsetOwnerNameIsPublic()`

UnsetOwnerNameIsPublic ensures that no value is present for OwnerNameIsPublic, not even an explicit nil
### GetOwnerKind

`func (o *OwnershipPositionsRow) GetOwnerKind() string`

GetOwnerKind returns the OwnerKind field if non-nil, zero value otherwise.

### GetOwnerKindOk

`func (o *OwnershipPositionsRow) GetOwnerKindOk() (*string, bool)`

GetOwnerKindOk returns a tuple with the OwnerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerKind

`func (o *OwnershipPositionsRow) SetOwnerKind(v string)`

SetOwnerKind sets OwnerKind field to given value.

### HasOwnerKind

`func (o *OwnershipPositionsRow) HasOwnerKind() bool`

HasOwnerKind returns a boolean if a field has been set.

### SetOwnerKindNil

`func (o *OwnershipPositionsRow) SetOwnerKindNil(b bool)`

 SetOwnerKindNil sets the value for OwnerKind to be an explicit nil

### UnsetOwnerKind
`func (o *OwnershipPositionsRow) UnsetOwnerKind()`

UnsetOwnerKind ensures that no value is present for OwnerKind, not even an explicit nil
### GetOwnerIdentityKey

`func (o *OwnershipPositionsRow) GetOwnerIdentityKey() string`

GetOwnerIdentityKey returns the OwnerIdentityKey field if non-nil, zero value otherwise.

### GetOwnerIdentityKeyOk

`func (o *OwnershipPositionsRow) GetOwnerIdentityKeyOk() (*string, bool)`

GetOwnerIdentityKeyOk returns a tuple with the OwnerIdentityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdentityKey

`func (o *OwnershipPositionsRow) SetOwnerIdentityKey(v string)`

SetOwnerIdentityKey sets OwnerIdentityKey field to given value.

### HasOwnerIdentityKey

`func (o *OwnershipPositionsRow) HasOwnerIdentityKey() bool`

HasOwnerIdentityKey returns a boolean if a field has been set.

### SetOwnerIdentityKeyNil

`func (o *OwnershipPositionsRow) SetOwnerIdentityKeyNil(b bool)`

 SetOwnerIdentityKeyNil sets the value for OwnerIdentityKey to be an explicit nil

### UnsetOwnerIdentityKey
`func (o *OwnershipPositionsRow) UnsetOwnerIdentityKey()`

UnsetOwnerIdentityKey ensures that no value is present for OwnerIdentityKey, not even an explicit nil
### GetOwnerPersonId

`func (o *OwnershipPositionsRow) GetOwnerPersonId() string`

GetOwnerPersonId returns the OwnerPersonId field if non-nil, zero value otherwise.

### GetOwnerPersonIdOk

`func (o *OwnershipPositionsRow) GetOwnerPersonIdOk() (*string, bool)`

GetOwnerPersonIdOk returns a tuple with the OwnerPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPersonId

`func (o *OwnershipPositionsRow) SetOwnerPersonId(v string)`

SetOwnerPersonId sets OwnerPersonId field to given value.

### HasOwnerPersonId

`func (o *OwnershipPositionsRow) HasOwnerPersonId() bool`

HasOwnerPersonId returns a boolean if a field has been set.

### SetOwnerPersonIdNil

`func (o *OwnershipPositionsRow) SetOwnerPersonIdNil(b bool)`

 SetOwnerPersonIdNil sets the value for OwnerPersonId to be an explicit nil

### UnsetOwnerPersonId
`func (o *OwnershipPositionsRow) UnsetOwnerPersonId()`

UnsetOwnerPersonId ensures that no value is present for OwnerPersonId, not even an explicit nil
### GetOwnerInvestorId

`func (o *OwnershipPositionsRow) GetOwnerInvestorId() string`

GetOwnerInvestorId returns the OwnerInvestorId field if non-nil, zero value otherwise.

### GetOwnerInvestorIdOk

`func (o *OwnershipPositionsRow) GetOwnerInvestorIdOk() (*string, bool)`

GetOwnerInvestorIdOk returns a tuple with the OwnerInvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInvestorId

`func (o *OwnershipPositionsRow) SetOwnerInvestorId(v string)`

SetOwnerInvestorId sets OwnerInvestorId field to given value.

### HasOwnerInvestorId

`func (o *OwnershipPositionsRow) HasOwnerInvestorId() bool`

HasOwnerInvestorId returns a boolean if a field has been set.

### SetOwnerInvestorIdNil

`func (o *OwnershipPositionsRow) SetOwnerInvestorIdNil(b bool)`

 SetOwnerInvestorIdNil sets the value for OwnerInvestorId to be an explicit nil

### UnsetOwnerInvestorId
`func (o *OwnershipPositionsRow) UnsetOwnerInvestorId()`

UnsetOwnerInvestorId ensures that no value is present for OwnerInvestorId, not even an explicit nil
### GetOwnerOrgId

`func (o *OwnershipPositionsRow) GetOwnerOrgId() string`

GetOwnerOrgId returns the OwnerOrgId field if non-nil, zero value otherwise.

### GetOwnerOrgIdOk

`func (o *OwnershipPositionsRow) GetOwnerOrgIdOk() (*string, bool)`

GetOwnerOrgIdOk returns a tuple with the OwnerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrgId

`func (o *OwnershipPositionsRow) SetOwnerOrgId(v string)`

SetOwnerOrgId sets OwnerOrgId field to given value.

### HasOwnerOrgId

`func (o *OwnershipPositionsRow) HasOwnerOrgId() bool`

HasOwnerOrgId returns a boolean if a field has been set.

### SetOwnerOrgIdNil

`func (o *OwnershipPositionsRow) SetOwnerOrgIdNil(b bool)`

 SetOwnerOrgIdNil sets the value for OwnerOrgId to be an explicit nil

### UnsetOwnerOrgId
`func (o *OwnershipPositionsRow) UnsetOwnerOrgId()`

UnsetOwnerOrgId ensures that no value is present for OwnerOrgId, not even an explicit nil
### GetOwnerCompanyId

`func (o *OwnershipPositionsRow) GetOwnerCompanyId() string`

GetOwnerCompanyId returns the OwnerCompanyId field if non-nil, zero value otherwise.

### GetOwnerCompanyIdOk

`func (o *OwnershipPositionsRow) GetOwnerCompanyIdOk() (*string, bool)`

GetOwnerCompanyIdOk returns a tuple with the OwnerCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerCompanyId

`func (o *OwnershipPositionsRow) SetOwnerCompanyId(v string)`

SetOwnerCompanyId sets OwnerCompanyId field to given value.

### HasOwnerCompanyId

`func (o *OwnershipPositionsRow) HasOwnerCompanyId() bool`

HasOwnerCompanyId returns a boolean if a field has been set.

### SetOwnerCompanyIdNil

`func (o *OwnershipPositionsRow) SetOwnerCompanyIdNil(b bool)`

 SetOwnerCompanyIdNil sets the value for OwnerCompanyId to be an explicit nil

### UnsetOwnerCompanyId
`func (o *OwnershipPositionsRow) UnsetOwnerCompanyId()`

UnsetOwnerCompanyId ensures that no value is present for OwnerCompanyId, not even an explicit nil
### GetOwnerCompanyName

`func (o *OwnershipPositionsRow) GetOwnerCompanyName() string`

GetOwnerCompanyName returns the OwnerCompanyName field if non-nil, zero value otherwise.

### GetOwnerCompanyNameOk

`func (o *OwnershipPositionsRow) GetOwnerCompanyNameOk() (*string, bool)`

GetOwnerCompanyNameOk returns a tuple with the OwnerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerCompanyName

`func (o *OwnershipPositionsRow) SetOwnerCompanyName(v string)`

SetOwnerCompanyName sets OwnerCompanyName field to given value.

### HasOwnerCompanyName

`func (o *OwnershipPositionsRow) HasOwnerCompanyName() bool`

HasOwnerCompanyName returns a boolean if a field has been set.

### SetOwnerCompanyNameNil

`func (o *OwnershipPositionsRow) SetOwnerCompanyNameNil(b bool)`

 SetOwnerCompanyNameNil sets the value for OwnerCompanyName to be an explicit nil

### UnsetOwnerCompanyName
`func (o *OwnershipPositionsRow) UnsetOwnerCompanyName()`

UnsetOwnerCompanyName ensures that no value is present for OwnerCompanyName, not even an explicit nil
### GetOwnerGroupId

`func (o *OwnershipPositionsRow) GetOwnerGroupId() string`

GetOwnerGroupId returns the OwnerGroupId field if non-nil, zero value otherwise.

### GetOwnerGroupIdOk

`func (o *OwnershipPositionsRow) GetOwnerGroupIdOk() (*string, bool)`

GetOwnerGroupIdOk returns a tuple with the OwnerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupId

`func (o *OwnershipPositionsRow) SetOwnerGroupId(v string)`

SetOwnerGroupId sets OwnerGroupId field to given value.

### HasOwnerGroupId

`func (o *OwnershipPositionsRow) HasOwnerGroupId() bool`

HasOwnerGroupId returns a boolean if a field has been set.

### SetOwnerGroupIdNil

`func (o *OwnershipPositionsRow) SetOwnerGroupIdNil(b bool)`

 SetOwnerGroupIdNil sets the value for OwnerGroupId to be an explicit nil

### UnsetOwnerGroupId
`func (o *OwnershipPositionsRow) UnsetOwnerGroupId()`

UnsetOwnerGroupId ensures that no value is present for OwnerGroupId, not even an explicit nil
### GetOwnerGroupName

`func (o *OwnershipPositionsRow) GetOwnerGroupName() string`

GetOwnerGroupName returns the OwnerGroupName field if non-nil, zero value otherwise.

### GetOwnerGroupNameOk

`func (o *OwnershipPositionsRow) GetOwnerGroupNameOk() (*string, bool)`

GetOwnerGroupNameOk returns a tuple with the OwnerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupName

`func (o *OwnershipPositionsRow) SetOwnerGroupName(v string)`

SetOwnerGroupName sets OwnerGroupName field to given value.

### HasOwnerGroupName

`func (o *OwnershipPositionsRow) HasOwnerGroupName() bool`

HasOwnerGroupName returns a boolean if a field has been set.

### SetOwnerGroupNameNil

`func (o *OwnershipPositionsRow) SetOwnerGroupNameNil(b bool)`

 SetOwnerGroupNameNil sets the value for OwnerGroupName to be an explicit nil

### UnsetOwnerGroupName
`func (o *OwnershipPositionsRow) UnsetOwnerGroupName()`

UnsetOwnerGroupName ensures that no value is present for OwnerGroupName, not even an explicit nil
### GetSharesBasis

`func (o *OwnershipPositionsRow) GetSharesBasis() string`

GetSharesBasis returns the SharesBasis field if non-nil, zero value otherwise.

### GetSharesBasisOk

`func (o *OwnershipPositionsRow) GetSharesBasisOk() (*string, bool)`

GetSharesBasisOk returns a tuple with the SharesBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesBasis

`func (o *OwnershipPositionsRow) SetSharesBasis(v string)`

SetSharesBasis sets SharesBasis field to given value.

### HasSharesBasis

`func (o *OwnershipPositionsRow) HasSharesBasis() bool`

HasSharesBasis returns a boolean if a field has been set.

### SetSharesBasisNil

`func (o *OwnershipPositionsRow) SetSharesBasisNil(b bool)`

 SetSharesBasisNil sets the value for SharesBasis to be an explicit nil

### UnsetSharesBasis
`func (o *OwnershipPositionsRow) UnsetSharesBasis()`

UnsetSharesBasis ensures that no value is present for SharesBasis, not even an explicit nil
### GetShares

`func (o *OwnershipPositionsRow) GetShares() float32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *OwnershipPositionsRow) GetSharesOk() (*float32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *OwnershipPositionsRow) SetShares(v float32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *OwnershipPositionsRow) HasShares() bool`

HasShares returns a boolean if a field has been set.

### SetSharesNil

`func (o *OwnershipPositionsRow) SetSharesNil(b bool)`

 SetSharesNil sets the value for Shares to be an explicit nil

### UnsetShares
`func (o *OwnershipPositionsRow) UnsetShares()`

UnsetShares ensures that no value is present for Shares, not even an explicit nil
### GetSharesDirect

`func (o *OwnershipPositionsRow) GetSharesDirect() float32`

GetSharesDirect returns the SharesDirect field if non-nil, zero value otherwise.

### GetSharesDirectOk

`func (o *OwnershipPositionsRow) GetSharesDirectOk() (*float32, bool)`

GetSharesDirectOk returns a tuple with the SharesDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesDirect

`func (o *OwnershipPositionsRow) SetSharesDirect(v float32)`

SetSharesDirect sets SharesDirect field to given value.

### HasSharesDirect

`func (o *OwnershipPositionsRow) HasSharesDirect() bool`

HasSharesDirect returns a boolean if a field has been set.

### SetSharesDirectNil

`func (o *OwnershipPositionsRow) SetSharesDirectNil(b bool)`

 SetSharesDirectNil sets the value for SharesDirect to be an explicit nil

### UnsetSharesDirect
`func (o *OwnershipPositionsRow) UnsetSharesDirect()`

UnsetSharesDirect ensures that no value is present for SharesDirect, not even an explicit nil
### GetSharesIndirect

`func (o *OwnershipPositionsRow) GetSharesIndirect() float32`

GetSharesIndirect returns the SharesIndirect field if non-nil, zero value otherwise.

### GetSharesIndirectOk

`func (o *OwnershipPositionsRow) GetSharesIndirectOk() (*float32, bool)`

GetSharesIndirectOk returns a tuple with the SharesIndirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesIndirect

`func (o *OwnershipPositionsRow) SetSharesIndirect(v float32)`

SetSharesIndirect sets SharesIndirect field to given value.

### HasSharesIndirect

`func (o *OwnershipPositionsRow) HasSharesIndirect() bool`

HasSharesIndirect returns a boolean if a field has been set.

### SetSharesIndirectNil

`func (o *OwnershipPositionsRow) SetSharesIndirectNil(b bool)`

 SetSharesIndirectNil sets the value for SharesIndirect to be an explicit nil

### UnsetSharesIndirect
`func (o *OwnershipPositionsRow) UnsetSharesIndirect()`

UnsetSharesIndirect ensures that no value is present for SharesIndirect, not even an explicit nil
### GetDerivativeUnderlyingShares

`func (o *OwnershipPositionsRow) GetDerivativeUnderlyingShares() float32`

GetDerivativeUnderlyingShares returns the DerivativeUnderlyingShares field if non-nil, zero value otherwise.

### GetDerivativeUnderlyingSharesOk

`func (o *OwnershipPositionsRow) GetDerivativeUnderlyingSharesOk() (*float32, bool)`

GetDerivativeUnderlyingSharesOk returns a tuple with the DerivativeUnderlyingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivativeUnderlyingShares

`func (o *OwnershipPositionsRow) SetDerivativeUnderlyingShares(v float32)`

SetDerivativeUnderlyingShares sets DerivativeUnderlyingShares field to given value.

### HasDerivativeUnderlyingShares

`func (o *OwnershipPositionsRow) HasDerivativeUnderlyingShares() bool`

HasDerivativeUnderlyingShares returns a boolean if a field has been set.

### SetDerivativeUnderlyingSharesNil

`func (o *OwnershipPositionsRow) SetDerivativeUnderlyingSharesNil(b bool)`

 SetDerivativeUnderlyingSharesNil sets the value for DerivativeUnderlyingShares to be an explicit nil

### UnsetDerivativeUnderlyingShares
`func (o *OwnershipPositionsRow) UnsetDerivativeUnderlyingShares()`

UnsetDerivativeUnderlyingShares ensures that no value is present for DerivativeUnderlyingShares, not even an explicit nil
### GetPercentBasis

`func (o *OwnershipPositionsRow) GetPercentBasis() string`

GetPercentBasis returns the PercentBasis field if non-nil, zero value otherwise.

### GetPercentBasisOk

`func (o *OwnershipPositionsRow) GetPercentBasisOk() (*string, bool)`

GetPercentBasisOk returns a tuple with the PercentBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentBasis

`func (o *OwnershipPositionsRow) SetPercentBasis(v string)`

SetPercentBasis sets PercentBasis field to given value.

### HasPercentBasis

`func (o *OwnershipPositionsRow) HasPercentBasis() bool`

HasPercentBasis returns a boolean if a field has been set.

### SetPercentBasisNil

`func (o *OwnershipPositionsRow) SetPercentBasisNil(b bool)`

 SetPercentBasisNil sets the value for PercentBasis to be an explicit nil

### UnsetPercentBasis
`func (o *OwnershipPositionsRow) UnsetPercentBasis()`

UnsetPercentBasis ensures that no value is present for PercentBasis, not even an explicit nil
### GetReportedPercent

`func (o *OwnershipPositionsRow) GetReportedPercent() float32`

GetReportedPercent returns the ReportedPercent field if non-nil, zero value otherwise.

### GetReportedPercentOk

`func (o *OwnershipPositionsRow) GetReportedPercentOk() (*float32, bool)`

GetReportedPercentOk returns a tuple with the ReportedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedPercent

`func (o *OwnershipPositionsRow) SetReportedPercent(v float32)`

SetReportedPercent sets ReportedPercent field to given value.

### HasReportedPercent

`func (o *OwnershipPositionsRow) HasReportedPercent() bool`

HasReportedPercent returns a boolean if a field has been set.

### SetReportedPercentNil

`func (o *OwnershipPositionsRow) SetReportedPercentNil(b bool)`

 SetReportedPercentNil sets the value for ReportedPercent to be an explicit nil

### UnsetReportedPercent
`func (o *OwnershipPositionsRow) UnsetReportedPercent()`

UnsetReportedPercent ensures that no value is present for ReportedPercent, not even an explicit nil
### GetReportedPercentLow

`func (o *OwnershipPositionsRow) GetReportedPercentLow() float32`

GetReportedPercentLow returns the ReportedPercentLow field if non-nil, zero value otherwise.

### GetReportedPercentLowOk

`func (o *OwnershipPositionsRow) GetReportedPercentLowOk() (*float32, bool)`

GetReportedPercentLowOk returns a tuple with the ReportedPercentLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedPercentLow

`func (o *OwnershipPositionsRow) SetReportedPercentLow(v float32)`

SetReportedPercentLow sets ReportedPercentLow field to given value.

### HasReportedPercentLow

`func (o *OwnershipPositionsRow) HasReportedPercentLow() bool`

HasReportedPercentLow returns a boolean if a field has been set.

### SetReportedPercentLowNil

`func (o *OwnershipPositionsRow) SetReportedPercentLowNil(b bool)`

 SetReportedPercentLowNil sets the value for ReportedPercentLow to be an explicit nil

### UnsetReportedPercentLow
`func (o *OwnershipPositionsRow) UnsetReportedPercentLow()`

UnsetReportedPercentLow ensures that no value is present for ReportedPercentLow, not even an explicit nil
### GetReportedPercentHigh

`func (o *OwnershipPositionsRow) GetReportedPercentHigh() float32`

GetReportedPercentHigh returns the ReportedPercentHigh field if non-nil, zero value otherwise.

### GetReportedPercentHighOk

`func (o *OwnershipPositionsRow) GetReportedPercentHighOk() (*float32, bool)`

GetReportedPercentHighOk returns a tuple with the ReportedPercentHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedPercentHigh

`func (o *OwnershipPositionsRow) SetReportedPercentHigh(v float32)`

SetReportedPercentHigh sets ReportedPercentHigh field to given value.

### HasReportedPercentHigh

`func (o *OwnershipPositionsRow) HasReportedPercentHigh() bool`

HasReportedPercentHigh returns a boolean if a field has been set.

### SetReportedPercentHighNil

`func (o *OwnershipPositionsRow) SetReportedPercentHighNil(b bool)`

 SetReportedPercentHighNil sets the value for ReportedPercentHigh to be an explicit nil

### UnsetReportedPercentHigh
`func (o *OwnershipPositionsRow) UnsetReportedPercentHigh()`

UnsetReportedPercentHigh ensures that no value is present for ReportedPercentHigh, not even an explicit nil
### GetPercentageBasis

`func (o *OwnershipPositionsRow) GetPercentageBasis() string`

GetPercentageBasis returns the PercentageBasis field if non-nil, zero value otherwise.

### GetPercentageBasisOk

`func (o *OwnershipPositionsRow) GetPercentageBasisOk() (*string, bool)`

GetPercentageBasisOk returns a tuple with the PercentageBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageBasis

`func (o *OwnershipPositionsRow) SetPercentageBasis(v string)`

SetPercentageBasis sets PercentageBasis field to given value.

### HasPercentageBasis

`func (o *OwnershipPositionsRow) HasPercentageBasis() bool`

HasPercentageBasis returns a boolean if a field has been set.

### SetPercentageBasisNil

`func (o *OwnershipPositionsRow) SetPercentageBasisNil(b bool)`

 SetPercentageBasisNil sets the value for PercentageBasis to be an explicit nil

### UnsetPercentageBasis
`func (o *OwnershipPositionsRow) UnsetPercentageBasis()`

UnsetPercentageBasis ensures that no value is present for PercentageBasis, not even an explicit nil
### GetEdgeState

`func (o *OwnershipPositionsRow) GetEdgeState() string`

GetEdgeState returns the EdgeState field if non-nil, zero value otherwise.

### GetEdgeStateOk

`func (o *OwnershipPositionsRow) GetEdgeStateOk() (*string, bool)`

GetEdgeStateOk returns a tuple with the EdgeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeState

`func (o *OwnershipPositionsRow) SetEdgeState(v string)`

SetEdgeState sets EdgeState field to given value.

### HasEdgeState

`func (o *OwnershipPositionsRow) HasEdgeState() bool`

HasEdgeState returns a boolean if a field has been set.

### SetEdgeStateNil

`func (o *OwnershipPositionsRow) SetEdgeStateNil(b bool)`

 SetEdgeStateNil sets the value for EdgeState to be an explicit nil

### UnsetEdgeState
`func (o *OwnershipPositionsRow) UnsetEdgeState()`

UnsetEdgeState ensures that no value is present for EdgeState, not even an explicit nil
### GetPositionCount

`func (o *OwnershipPositionsRow) GetPositionCount() int32`

GetPositionCount returns the PositionCount field if non-nil, zero value otherwise.

### GetPositionCountOk

`func (o *OwnershipPositionsRow) GetPositionCountOk() (*int32, bool)`

GetPositionCountOk returns a tuple with the PositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionCount

`func (o *OwnershipPositionsRow) SetPositionCount(v int32)`

SetPositionCount sets PositionCount field to given value.

### HasPositionCount

`func (o *OwnershipPositionsRow) HasPositionCount() bool`

HasPositionCount returns a boolean if a field has been set.

### SetPositionCountNil

`func (o *OwnershipPositionsRow) SetPositionCountNil(b bool)`

 SetPositionCountNil sets the value for PositionCount to be an explicit nil

### UnsetPositionCount
`func (o *OwnershipPositionsRow) UnsetPositionCount()`

UnsetPositionCount ensures that no value is present for PositionCount, not even an explicit nil
### GetStandingPositionCount

`func (o *OwnershipPositionsRow) GetStandingPositionCount() int32`

GetStandingPositionCount returns the StandingPositionCount field if non-nil, zero value otherwise.

### GetStandingPositionCountOk

`func (o *OwnershipPositionsRow) GetStandingPositionCountOk() (*int32, bool)`

GetStandingPositionCountOk returns a tuple with the StandingPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandingPositionCount

`func (o *OwnershipPositionsRow) SetStandingPositionCount(v int32)`

SetStandingPositionCount sets StandingPositionCount field to given value.

### HasStandingPositionCount

`func (o *OwnershipPositionsRow) HasStandingPositionCount() bool`

HasStandingPositionCount returns a boolean if a field has been set.

### SetStandingPositionCountNil

`func (o *OwnershipPositionsRow) SetStandingPositionCountNil(b bool)`

 SetStandingPositionCountNil sets the value for StandingPositionCount to be an explicit nil

### UnsetStandingPositionCount
`func (o *OwnershipPositionsRow) UnsetStandingPositionCount()`

UnsetStandingPositionCount ensures that no value is present for StandingPositionCount, not even an explicit nil
### GetCeasedPositionCount

`func (o *OwnershipPositionsRow) GetCeasedPositionCount() int32`

GetCeasedPositionCount returns the CeasedPositionCount field if non-nil, zero value otherwise.

### GetCeasedPositionCountOk

`func (o *OwnershipPositionsRow) GetCeasedPositionCountOk() (*int32, bool)`

GetCeasedPositionCountOk returns a tuple with the CeasedPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeasedPositionCount

`func (o *OwnershipPositionsRow) SetCeasedPositionCount(v int32)`

SetCeasedPositionCount sets CeasedPositionCount field to given value.

### HasCeasedPositionCount

`func (o *OwnershipPositionsRow) HasCeasedPositionCount() bool`

HasCeasedPositionCount returns a boolean if a field has been set.

### SetCeasedPositionCountNil

`func (o *OwnershipPositionsRow) SetCeasedPositionCountNil(b bool)`

 SetCeasedPositionCountNil sets the value for CeasedPositionCount to be an explicit nil

### UnsetCeasedPositionCount
`func (o *OwnershipPositionsRow) UnsetCeasedPositionCount()`

UnsetCeasedPositionCount ensures that no value is present for CeasedPositionCount, not even an explicit nil
### GetNoLongerReportedPositionCount

`func (o *OwnershipPositionsRow) GetNoLongerReportedPositionCount() int32`

GetNoLongerReportedPositionCount returns the NoLongerReportedPositionCount field if non-nil, zero value otherwise.

### GetNoLongerReportedPositionCountOk

`func (o *OwnershipPositionsRow) GetNoLongerReportedPositionCountOk() (*int32, bool)`

GetNoLongerReportedPositionCountOk returns a tuple with the NoLongerReportedPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLongerReportedPositionCount

`func (o *OwnershipPositionsRow) SetNoLongerReportedPositionCount(v int32)`

SetNoLongerReportedPositionCount sets NoLongerReportedPositionCount field to given value.

### HasNoLongerReportedPositionCount

`func (o *OwnershipPositionsRow) HasNoLongerReportedPositionCount() bool`

HasNoLongerReportedPositionCount returns a boolean if a field has been set.

### SetNoLongerReportedPositionCountNil

`func (o *OwnershipPositionsRow) SetNoLongerReportedPositionCountNil(b bool)`

 SetNoLongerReportedPositionCountNil sets the value for NoLongerReportedPositionCount to be an explicit nil

### UnsetNoLongerReportedPositionCount
`func (o *OwnershipPositionsRow) UnsetNoLongerReportedPositionCount()`

UnsetNoLongerReportedPositionCount ensures that no value is present for NoLongerReportedPositionCount, not even an explicit nil
### GetFirstObservedDate

`func (o *OwnershipPositionsRow) GetFirstObservedDate() string`

GetFirstObservedDate returns the FirstObservedDate field if non-nil, zero value otherwise.

### GetFirstObservedDateOk

`func (o *OwnershipPositionsRow) GetFirstObservedDateOk() (*string, bool)`

GetFirstObservedDateOk returns a tuple with the FirstObservedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstObservedDate

`func (o *OwnershipPositionsRow) SetFirstObservedDate(v string)`

SetFirstObservedDate sets FirstObservedDate field to given value.

### HasFirstObservedDate

`func (o *OwnershipPositionsRow) HasFirstObservedDate() bool`

HasFirstObservedDate returns a boolean if a field has been set.

### SetFirstObservedDateNil

`func (o *OwnershipPositionsRow) SetFirstObservedDateNil(b bool)`

 SetFirstObservedDateNil sets the value for FirstObservedDate to be an explicit nil

### UnsetFirstObservedDate
`func (o *OwnershipPositionsRow) UnsetFirstObservedDate()`

UnsetFirstObservedDate ensures that no value is present for FirstObservedDate, not even an explicit nil
### GetLastActivityDate

`func (o *OwnershipPositionsRow) GetLastActivityDate() string`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *OwnershipPositionsRow) GetLastActivityDateOk() (*string, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *OwnershipPositionsRow) SetLastActivityDate(v string)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *OwnershipPositionsRow) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### SetLastActivityDateNil

`func (o *OwnershipPositionsRow) SetLastActivityDateNil(b bool)`

 SetLastActivityDateNil sets the value for LastActivityDate to be an explicit nil

### UnsetLastActivityDate
`func (o *OwnershipPositionsRow) UnsetLastActivityDate()`

UnsetLastActivityDate ensures that no value is present for LastActivityDate, not even an explicit nil
### GetLastFilingDate

`func (o *OwnershipPositionsRow) GetLastFilingDate() string`

GetLastFilingDate returns the LastFilingDate field if non-nil, zero value otherwise.

### GetLastFilingDateOk

`func (o *OwnershipPositionsRow) GetLastFilingDateOk() (*string, bool)`

GetLastFilingDateOk returns a tuple with the LastFilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFilingDate

`func (o *OwnershipPositionsRow) SetLastFilingDate(v string)`

SetLastFilingDate sets LastFilingDate field to given value.

### HasLastFilingDate

`func (o *OwnershipPositionsRow) HasLastFilingDate() bool`

HasLastFilingDate returns a boolean if a field has been set.

### SetLastFilingDateNil

`func (o *OwnershipPositionsRow) SetLastFilingDateNil(b bool)`

 SetLastFilingDateNil sets the value for LastFilingDate to be an explicit nil

### UnsetLastFilingDate
`func (o *OwnershipPositionsRow) UnsetLastFilingDate()`

UnsetLastFilingDate ensures that no value is present for LastFilingDate, not even an explicit nil
### GetSourceUrl

`func (o *OwnershipPositionsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *OwnershipPositionsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *OwnershipPositionsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *OwnershipPositionsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *OwnershipPositionsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *OwnershipPositionsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetContainsPersonalData

`func (o *OwnershipPositionsRow) GetContainsPersonalData() bool`

GetContainsPersonalData returns the ContainsPersonalData field if non-nil, zero value otherwise.

### GetContainsPersonalDataOk

`func (o *OwnershipPositionsRow) GetContainsPersonalDataOk() (*bool, bool)`

GetContainsPersonalDataOk returns a tuple with the ContainsPersonalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsPersonalData

`func (o *OwnershipPositionsRow) SetContainsPersonalData(v bool)`

SetContainsPersonalData sets ContainsPersonalData field to given value.

### HasContainsPersonalData

`func (o *OwnershipPositionsRow) HasContainsPersonalData() bool`

HasContainsPersonalData returns a boolean if a field has been set.

### SetContainsPersonalDataNil

`func (o *OwnershipPositionsRow) SetContainsPersonalDataNil(b bool)`

 SetContainsPersonalDataNil sets the value for ContainsPersonalData to be an explicit nil

### UnsetContainsPersonalData
`func (o *OwnershipPositionsRow) UnsetContainsPersonalData()`

UnsetContainsPersonalData ensures that no value is present for ContainsPersonalData, not even an explicit nil
### GetPrivacyStatus

`func (o *OwnershipPositionsRow) GetPrivacyStatus() string`

GetPrivacyStatus returns the PrivacyStatus field if non-nil, zero value otherwise.

### GetPrivacyStatusOk

`func (o *OwnershipPositionsRow) GetPrivacyStatusOk() (*string, bool)`

GetPrivacyStatusOk returns a tuple with the PrivacyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyStatus

`func (o *OwnershipPositionsRow) SetPrivacyStatus(v string)`

SetPrivacyStatus sets PrivacyStatus field to given value.

### HasPrivacyStatus

`func (o *OwnershipPositionsRow) HasPrivacyStatus() bool`

HasPrivacyStatus returns a boolean if a field has been set.

### SetPrivacyStatusNil

`func (o *OwnershipPositionsRow) SetPrivacyStatusNil(b bool)`

 SetPrivacyStatusNil sets the value for PrivacyStatus to be an explicit nil

### UnsetPrivacyStatus
`func (o *OwnershipPositionsRow) UnsetPrivacyStatus()`

UnsetPrivacyStatus ensures that no value is present for PrivacyStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


