# OrgsCurrentOwnership200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerRollupId** | **string** |  | 
**Direction** | **string** |  | 
**SourceClass** | **string** |  | 
**SourceSystem** | **string** |  | 
**OwnershipDimension** | **string** |  | 
**RelationshipType** | **string** |  | 
**OwnerDisplayName** | **string** |  | 
**OwnerKind** | **string** |  | 
**OwnerOrgId** | **interface{}** |  | 
**OwnerInvestorId** | **interface{}** |  | 
**SubjectNameRaw** | **string** |  | 
**SubjectOrgId** | **string** |  | 
**ReportedPercent** | **NullableFloat32** |  | 
**PercentageBasis** | **NullableString** |  | 
**PercentBasis** | **string** |  | 
**EquitySharesStanding** | **float32** |  | 
**SharesBasis** | **string** |  | 
**EdgeState** | **string** |  | 
**PositionCount** | **int32** |  | 
**StandingPositionCount** | **int32** |  | 
**CeasedPositionCount** | **int32** |  | 
**NoLongerReportedPositionCount** | **int32** |  | 
**LastActivityDate** | **string** |  | 
**RestrictionTier** | **string** |  | 
**SourceLicense** | **string** |  | 
**Attribution** | **string** |  | 
**ContainsPersonalData** | Pointer to **bool** |  | [optional] 
**PrivacyStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgsCurrentOwnership200ResponseRowsInner

`func NewOrgsCurrentOwnership200ResponseRowsInner(ownerRollupId string, direction string, sourceClass string, sourceSystem string, ownershipDimension string, relationshipType string, ownerDisplayName string, ownerKind string, ownerOrgId interface{}, ownerInvestorId interface{}, subjectNameRaw string, subjectOrgId string, reportedPercent NullableFloat32, percentageBasis NullableString, percentBasis string, equitySharesStanding float32, sharesBasis string, edgeState string, positionCount int32, standingPositionCount int32, ceasedPositionCount int32, noLongerReportedPositionCount int32, lastActivityDate string, restrictionTier string, sourceLicense string, attribution string, ) *OrgsCurrentOwnership200ResponseRowsInner`

NewOrgsCurrentOwnership200ResponseRowsInner instantiates a new OrgsCurrentOwnership200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsCurrentOwnership200ResponseRowsInnerWithDefaults

`func NewOrgsCurrentOwnership200ResponseRowsInnerWithDefaults() *OrgsCurrentOwnership200ResponseRowsInner`

NewOrgsCurrentOwnership200ResponseRowsInnerWithDefaults instantiates a new OrgsCurrentOwnership200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerRollupId

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerRollupId() string`

GetOwnerRollupId returns the OwnerRollupId field if non-nil, zero value otherwise.

### GetOwnerRollupIdOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerRollupIdOk() (*string, bool)`

GetOwnerRollupIdOk returns a tuple with the OwnerRollupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRollupId

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetOwnerRollupId(v string)`

SetOwnerRollupId sets OwnerRollupId field to given value.


### GetDirection

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetSourceClass

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSourceClass() string`

GetSourceClass returns the SourceClass field if non-nil, zero value otherwise.

### GetSourceClassOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSourceClassOk() (*string, bool)`

GetSourceClassOk returns a tuple with the SourceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClass

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetSourceClass(v string)`

SetSourceClass sets SourceClass field to given value.


### GetSourceSystem

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSourceSystem() string`

GetSourceSystem returns the SourceSystem field if non-nil, zero value otherwise.

### GetSourceSystemOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSourceSystemOk() (*string, bool)`

GetSourceSystemOk returns a tuple with the SourceSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystem

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetSourceSystem(v string)`

SetSourceSystem sets SourceSystem field to given value.


### GetOwnershipDimension

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnershipDimension() string`

GetOwnershipDimension returns the OwnershipDimension field if non-nil, zero value otherwise.

### GetOwnershipDimensionOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnershipDimensionOk() (*string, bool)`

GetOwnershipDimensionOk returns a tuple with the OwnershipDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipDimension

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetOwnershipDimension(v string)`

SetOwnershipDimension sets OwnershipDimension field to given value.


### GetRelationshipType

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.


### GetOwnerDisplayName

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerDisplayName() string`

GetOwnerDisplayName returns the OwnerDisplayName field if non-nil, zero value otherwise.

### GetOwnerDisplayNameOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerDisplayNameOk() (*string, bool)`

GetOwnerDisplayNameOk returns a tuple with the OwnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerDisplayName

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetOwnerDisplayName(v string)`

SetOwnerDisplayName sets OwnerDisplayName field to given value.


### GetOwnerKind

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerKind() string`

GetOwnerKind returns the OwnerKind field if non-nil, zero value otherwise.

### GetOwnerKindOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerKindOk() (*string, bool)`

GetOwnerKindOk returns a tuple with the OwnerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerKind

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetOwnerKind(v string)`

SetOwnerKind sets OwnerKind field to given value.


### GetOwnerOrgId

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerOrgId() interface{}`

GetOwnerOrgId returns the OwnerOrgId field if non-nil, zero value otherwise.

### GetOwnerOrgIdOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerOrgIdOk() (*interface{}, bool)`

GetOwnerOrgIdOk returns a tuple with the OwnerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrgId

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetOwnerOrgId(v interface{})`

SetOwnerOrgId sets OwnerOrgId field to given value.


### SetOwnerOrgIdNil

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetOwnerOrgIdNil(b bool)`

 SetOwnerOrgIdNil sets the value for OwnerOrgId to be an explicit nil

### UnsetOwnerOrgId
`func (o *OrgsCurrentOwnership200ResponseRowsInner) UnsetOwnerOrgId()`

UnsetOwnerOrgId ensures that no value is present for OwnerOrgId, not even an explicit nil
### GetOwnerInvestorId

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerInvestorId() interface{}`

GetOwnerInvestorId returns the OwnerInvestorId field if non-nil, zero value otherwise.

### GetOwnerInvestorIdOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetOwnerInvestorIdOk() (*interface{}, bool)`

GetOwnerInvestorIdOk returns a tuple with the OwnerInvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInvestorId

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetOwnerInvestorId(v interface{})`

SetOwnerInvestorId sets OwnerInvestorId field to given value.


### SetOwnerInvestorIdNil

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetOwnerInvestorIdNil(b bool)`

 SetOwnerInvestorIdNil sets the value for OwnerInvestorId to be an explicit nil

### UnsetOwnerInvestorId
`func (o *OrgsCurrentOwnership200ResponseRowsInner) UnsetOwnerInvestorId()`

UnsetOwnerInvestorId ensures that no value is present for OwnerInvestorId, not even an explicit nil
### GetSubjectNameRaw

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSubjectNameRaw() string`

GetSubjectNameRaw returns the SubjectNameRaw field if non-nil, zero value otherwise.

### GetSubjectNameRawOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSubjectNameRawOk() (*string, bool)`

GetSubjectNameRawOk returns a tuple with the SubjectNameRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameRaw

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetSubjectNameRaw(v string)`

SetSubjectNameRaw sets SubjectNameRaw field to given value.


### GetSubjectOrgId

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSubjectOrgId() string`

GetSubjectOrgId returns the SubjectOrgId field if non-nil, zero value otherwise.

### GetSubjectOrgIdOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSubjectOrgIdOk() (*string, bool)`

GetSubjectOrgIdOk returns a tuple with the SubjectOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectOrgId

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetSubjectOrgId(v string)`

SetSubjectOrgId sets SubjectOrgId field to given value.


### GetReportedPercent

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetReportedPercent() float32`

GetReportedPercent returns the ReportedPercent field if non-nil, zero value otherwise.

### GetReportedPercentOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetReportedPercentOk() (*float32, bool)`

GetReportedPercentOk returns a tuple with the ReportedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedPercent

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetReportedPercent(v float32)`

SetReportedPercent sets ReportedPercent field to given value.


### SetReportedPercentNil

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetReportedPercentNil(b bool)`

 SetReportedPercentNil sets the value for ReportedPercent to be an explicit nil

### UnsetReportedPercent
`func (o *OrgsCurrentOwnership200ResponseRowsInner) UnsetReportedPercent()`

UnsetReportedPercent ensures that no value is present for ReportedPercent, not even an explicit nil
### GetPercentageBasis

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetPercentageBasis() string`

GetPercentageBasis returns the PercentageBasis field if non-nil, zero value otherwise.

### GetPercentageBasisOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetPercentageBasisOk() (*string, bool)`

GetPercentageBasisOk returns a tuple with the PercentageBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageBasis

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetPercentageBasis(v string)`

SetPercentageBasis sets PercentageBasis field to given value.


### SetPercentageBasisNil

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetPercentageBasisNil(b bool)`

 SetPercentageBasisNil sets the value for PercentageBasis to be an explicit nil

### UnsetPercentageBasis
`func (o *OrgsCurrentOwnership200ResponseRowsInner) UnsetPercentageBasis()`

UnsetPercentageBasis ensures that no value is present for PercentageBasis, not even an explicit nil
### GetPercentBasis

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetPercentBasis() string`

GetPercentBasis returns the PercentBasis field if non-nil, zero value otherwise.

### GetPercentBasisOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetPercentBasisOk() (*string, bool)`

GetPercentBasisOk returns a tuple with the PercentBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentBasis

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetPercentBasis(v string)`

SetPercentBasis sets PercentBasis field to given value.


### GetEquitySharesStanding

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetEquitySharesStanding() float32`

GetEquitySharesStanding returns the EquitySharesStanding field if non-nil, zero value otherwise.

### GetEquitySharesStandingOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetEquitySharesStandingOk() (*float32, bool)`

GetEquitySharesStandingOk returns a tuple with the EquitySharesStanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquitySharesStanding

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetEquitySharesStanding(v float32)`

SetEquitySharesStanding sets EquitySharesStanding field to given value.


### GetSharesBasis

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSharesBasis() string`

GetSharesBasis returns the SharesBasis field if non-nil, zero value otherwise.

### GetSharesBasisOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSharesBasisOk() (*string, bool)`

GetSharesBasisOk returns a tuple with the SharesBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesBasis

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetSharesBasis(v string)`

SetSharesBasis sets SharesBasis field to given value.


### GetEdgeState

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetEdgeState() string`

GetEdgeState returns the EdgeState field if non-nil, zero value otherwise.

### GetEdgeStateOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetEdgeStateOk() (*string, bool)`

GetEdgeStateOk returns a tuple with the EdgeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeState

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetEdgeState(v string)`

SetEdgeState sets EdgeState field to given value.


### GetPositionCount

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetPositionCount() int32`

GetPositionCount returns the PositionCount field if non-nil, zero value otherwise.

### GetPositionCountOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetPositionCountOk() (*int32, bool)`

GetPositionCountOk returns a tuple with the PositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionCount

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetPositionCount(v int32)`

SetPositionCount sets PositionCount field to given value.


### GetStandingPositionCount

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetStandingPositionCount() int32`

GetStandingPositionCount returns the StandingPositionCount field if non-nil, zero value otherwise.

### GetStandingPositionCountOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetStandingPositionCountOk() (*int32, bool)`

GetStandingPositionCountOk returns a tuple with the StandingPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandingPositionCount

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetStandingPositionCount(v int32)`

SetStandingPositionCount sets StandingPositionCount field to given value.


### GetCeasedPositionCount

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetCeasedPositionCount() int32`

GetCeasedPositionCount returns the CeasedPositionCount field if non-nil, zero value otherwise.

### GetCeasedPositionCountOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetCeasedPositionCountOk() (*int32, bool)`

GetCeasedPositionCountOk returns a tuple with the CeasedPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeasedPositionCount

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetCeasedPositionCount(v int32)`

SetCeasedPositionCount sets CeasedPositionCount field to given value.


### GetNoLongerReportedPositionCount

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetNoLongerReportedPositionCount() int32`

GetNoLongerReportedPositionCount returns the NoLongerReportedPositionCount field if non-nil, zero value otherwise.

### GetNoLongerReportedPositionCountOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetNoLongerReportedPositionCountOk() (*int32, bool)`

GetNoLongerReportedPositionCountOk returns a tuple with the NoLongerReportedPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLongerReportedPositionCount

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetNoLongerReportedPositionCount(v int32)`

SetNoLongerReportedPositionCount sets NoLongerReportedPositionCount field to given value.


### GetLastActivityDate

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetLastActivityDate() string`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetLastActivityDateOk() (*string, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetLastActivityDate(v string)`

SetLastActivityDate sets LastActivityDate field to given value.


### GetRestrictionTier

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.


### GetSourceLicense

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.


### GetAttribution

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.


### GetContainsPersonalData

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetContainsPersonalData() bool`

GetContainsPersonalData returns the ContainsPersonalData field if non-nil, zero value otherwise.

### GetContainsPersonalDataOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetContainsPersonalDataOk() (*bool, bool)`

GetContainsPersonalDataOk returns a tuple with the ContainsPersonalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsPersonalData

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetContainsPersonalData(v bool)`

SetContainsPersonalData sets ContainsPersonalData field to given value.

### HasContainsPersonalData

`func (o *OrgsCurrentOwnership200ResponseRowsInner) HasContainsPersonalData() bool`

HasContainsPersonalData returns a boolean if a field has been set.

### GetPrivacyStatus

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetPrivacyStatus() string`

GetPrivacyStatus returns the PrivacyStatus field if non-nil, zero value otherwise.

### GetPrivacyStatusOk

`func (o *OrgsCurrentOwnership200ResponseRowsInner) GetPrivacyStatusOk() (*string, bool)`

GetPrivacyStatusOk returns a tuple with the PrivacyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyStatus

`func (o *OrgsCurrentOwnership200ResponseRowsInner) SetPrivacyStatus(v string)`

SetPrivacyStatus sets PrivacyStatus field to given value.

### HasPrivacyStatus

`func (o *OrgsCurrentOwnership200ResponseRowsInner) HasPrivacyStatus() bool`

HasPrivacyStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


