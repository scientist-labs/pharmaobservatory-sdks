# OrgsOwnershipHistory200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnershipObservationId** | **string** |  | 
**Direction** | **string** |  | 
**SourceClass** | **string** |  | 
**SourceSystem** | **string** |  | 
**OwnershipDimension** | **string** |  | 
**RelationshipType** | **string** |  | 
**OwnerNameRaw** | **string** |  | 
**OwnerSourceRef** | **string** |  | 
**SubjectNameRaw** | **string** |  | 
**SubjectSourceRef** | **string** |  | 
**ReviewState** | **string** |  | 
**StakeValueType** | **string** |  | 
**StakePercent** | **NullableFloat32** |  | 
**PercentageBasis** | **NullableString** |  | 
**ObservationDate** | **string** |  | 
**IsCurrentEdge** | **bool** |  | 
**CurrentEdgeState** | **NullableString** |  | 
**RestrictionTier** | **string** |  | 
**SourceLicense** | **string** |  | 
**Attribution** | **string** |  | 
**ContainsPersonalData** | Pointer to **bool** |  | [optional] 
**PrivacyStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgsOwnershipHistory200ResponseRowsInner

`func NewOrgsOwnershipHistory200ResponseRowsInner(ownershipObservationId string, direction string, sourceClass string, sourceSystem string, ownershipDimension string, relationshipType string, ownerNameRaw string, ownerSourceRef string, subjectNameRaw string, subjectSourceRef string, reviewState string, stakeValueType string, stakePercent NullableFloat32, percentageBasis NullableString, observationDate string, isCurrentEdge bool, currentEdgeState NullableString, restrictionTier string, sourceLicense string, attribution string, ) *OrgsOwnershipHistory200ResponseRowsInner`

NewOrgsOwnershipHistory200ResponseRowsInner instantiates a new OrgsOwnershipHistory200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsOwnershipHistory200ResponseRowsInnerWithDefaults

`func NewOrgsOwnershipHistory200ResponseRowsInnerWithDefaults() *OrgsOwnershipHistory200ResponseRowsInner`

NewOrgsOwnershipHistory200ResponseRowsInnerWithDefaults instantiates a new OrgsOwnershipHistory200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnershipObservationId

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetOwnershipObservationId() string`

GetOwnershipObservationId returns the OwnershipObservationId field if non-nil, zero value otherwise.

### GetOwnershipObservationIdOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetOwnershipObservationIdOk() (*string, bool)`

GetOwnershipObservationIdOk returns a tuple with the OwnershipObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipObservationId

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetOwnershipObservationId(v string)`

SetOwnershipObservationId sets OwnershipObservationId field to given value.


### GetDirection

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetSourceClass

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSourceClass() string`

GetSourceClass returns the SourceClass field if non-nil, zero value otherwise.

### GetSourceClassOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSourceClassOk() (*string, bool)`

GetSourceClassOk returns a tuple with the SourceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClass

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetSourceClass(v string)`

SetSourceClass sets SourceClass field to given value.


### GetSourceSystem

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSourceSystem() string`

GetSourceSystem returns the SourceSystem field if non-nil, zero value otherwise.

### GetSourceSystemOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSourceSystemOk() (*string, bool)`

GetSourceSystemOk returns a tuple with the SourceSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystem

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetSourceSystem(v string)`

SetSourceSystem sets SourceSystem field to given value.


### GetOwnershipDimension

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetOwnershipDimension() string`

GetOwnershipDimension returns the OwnershipDimension field if non-nil, zero value otherwise.

### GetOwnershipDimensionOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetOwnershipDimensionOk() (*string, bool)`

GetOwnershipDimensionOk returns a tuple with the OwnershipDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipDimension

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetOwnershipDimension(v string)`

SetOwnershipDimension sets OwnershipDimension field to given value.


### GetRelationshipType

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.


### GetOwnerNameRaw

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetOwnerNameRaw() string`

GetOwnerNameRaw returns the OwnerNameRaw field if non-nil, zero value otherwise.

### GetOwnerNameRawOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetOwnerNameRawOk() (*string, bool)`

GetOwnerNameRawOk returns a tuple with the OwnerNameRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNameRaw

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetOwnerNameRaw(v string)`

SetOwnerNameRaw sets OwnerNameRaw field to given value.


### GetOwnerSourceRef

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetOwnerSourceRef() string`

GetOwnerSourceRef returns the OwnerSourceRef field if non-nil, zero value otherwise.

### GetOwnerSourceRefOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetOwnerSourceRefOk() (*string, bool)`

GetOwnerSourceRefOk returns a tuple with the OwnerSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSourceRef

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetOwnerSourceRef(v string)`

SetOwnerSourceRef sets OwnerSourceRef field to given value.


### GetSubjectNameRaw

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSubjectNameRaw() string`

GetSubjectNameRaw returns the SubjectNameRaw field if non-nil, zero value otherwise.

### GetSubjectNameRawOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSubjectNameRawOk() (*string, bool)`

GetSubjectNameRawOk returns a tuple with the SubjectNameRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameRaw

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetSubjectNameRaw(v string)`

SetSubjectNameRaw sets SubjectNameRaw field to given value.


### GetSubjectSourceRef

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSubjectSourceRef() string`

GetSubjectSourceRef returns the SubjectSourceRef field if non-nil, zero value otherwise.

### GetSubjectSourceRefOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSubjectSourceRefOk() (*string, bool)`

GetSubjectSourceRefOk returns a tuple with the SubjectSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectSourceRef

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetSubjectSourceRef(v string)`

SetSubjectSourceRef sets SubjectSourceRef field to given value.


### GetReviewState

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetReviewState() string`

GetReviewState returns the ReviewState field if non-nil, zero value otherwise.

### GetReviewStateOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetReviewStateOk() (*string, bool)`

GetReviewStateOk returns a tuple with the ReviewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewState

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetReviewState(v string)`

SetReviewState sets ReviewState field to given value.


### GetStakeValueType

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetStakeValueType() string`

GetStakeValueType returns the StakeValueType field if non-nil, zero value otherwise.

### GetStakeValueTypeOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetStakeValueTypeOk() (*string, bool)`

GetStakeValueTypeOk returns a tuple with the StakeValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeValueType

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetStakeValueType(v string)`

SetStakeValueType sets StakeValueType field to given value.


### GetStakePercent

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetStakePercent() float32`

GetStakePercent returns the StakePercent field if non-nil, zero value otherwise.

### GetStakePercentOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetStakePercentOk() (*float32, bool)`

GetStakePercentOk returns a tuple with the StakePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakePercent

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetStakePercent(v float32)`

SetStakePercent sets StakePercent field to given value.


### SetStakePercentNil

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetStakePercentNil(b bool)`

 SetStakePercentNil sets the value for StakePercent to be an explicit nil

### UnsetStakePercent
`func (o *OrgsOwnershipHistory200ResponseRowsInner) UnsetStakePercent()`

UnsetStakePercent ensures that no value is present for StakePercent, not even an explicit nil
### GetPercentageBasis

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetPercentageBasis() string`

GetPercentageBasis returns the PercentageBasis field if non-nil, zero value otherwise.

### GetPercentageBasisOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetPercentageBasisOk() (*string, bool)`

GetPercentageBasisOk returns a tuple with the PercentageBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageBasis

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetPercentageBasis(v string)`

SetPercentageBasis sets PercentageBasis field to given value.


### SetPercentageBasisNil

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetPercentageBasisNil(b bool)`

 SetPercentageBasisNil sets the value for PercentageBasis to be an explicit nil

### UnsetPercentageBasis
`func (o *OrgsOwnershipHistory200ResponseRowsInner) UnsetPercentageBasis()`

UnsetPercentageBasis ensures that no value is present for PercentageBasis, not even an explicit nil
### GetObservationDate

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetObservationDate() string`

GetObservationDate returns the ObservationDate field if non-nil, zero value otherwise.

### GetObservationDateOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetObservationDateOk() (*string, bool)`

GetObservationDateOk returns a tuple with the ObservationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDate

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetObservationDate(v string)`

SetObservationDate sets ObservationDate field to given value.


### GetIsCurrentEdge

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetIsCurrentEdge() bool`

GetIsCurrentEdge returns the IsCurrentEdge field if non-nil, zero value otherwise.

### GetIsCurrentEdgeOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetIsCurrentEdgeOk() (*bool, bool)`

GetIsCurrentEdgeOk returns a tuple with the IsCurrentEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrentEdge

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetIsCurrentEdge(v bool)`

SetIsCurrentEdge sets IsCurrentEdge field to given value.


### GetCurrentEdgeState

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetCurrentEdgeState() string`

GetCurrentEdgeState returns the CurrentEdgeState field if non-nil, zero value otherwise.

### GetCurrentEdgeStateOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetCurrentEdgeStateOk() (*string, bool)`

GetCurrentEdgeStateOk returns a tuple with the CurrentEdgeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEdgeState

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetCurrentEdgeState(v string)`

SetCurrentEdgeState sets CurrentEdgeState field to given value.


### SetCurrentEdgeStateNil

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetCurrentEdgeStateNil(b bool)`

 SetCurrentEdgeStateNil sets the value for CurrentEdgeState to be an explicit nil

### UnsetCurrentEdgeState
`func (o *OrgsOwnershipHistory200ResponseRowsInner) UnsetCurrentEdgeState()`

UnsetCurrentEdgeState ensures that no value is present for CurrentEdgeState, not even an explicit nil
### GetRestrictionTier

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.


### GetSourceLicense

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.


### GetAttribution

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.


### GetContainsPersonalData

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetContainsPersonalData() bool`

GetContainsPersonalData returns the ContainsPersonalData field if non-nil, zero value otherwise.

### GetContainsPersonalDataOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetContainsPersonalDataOk() (*bool, bool)`

GetContainsPersonalDataOk returns a tuple with the ContainsPersonalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsPersonalData

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetContainsPersonalData(v bool)`

SetContainsPersonalData sets ContainsPersonalData field to given value.

### HasContainsPersonalData

`func (o *OrgsOwnershipHistory200ResponseRowsInner) HasContainsPersonalData() bool`

HasContainsPersonalData returns a boolean if a field has been set.

### GetPrivacyStatus

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetPrivacyStatus() string`

GetPrivacyStatus returns the PrivacyStatus field if non-nil, zero value otherwise.

### GetPrivacyStatusOk

`func (o *OrgsOwnershipHistory200ResponseRowsInner) GetPrivacyStatusOk() (*string, bool)`

GetPrivacyStatusOk returns a tuple with the PrivacyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyStatus

`func (o *OrgsOwnershipHistory200ResponseRowsInner) SetPrivacyStatus(v string)`

SetPrivacyStatus sets PrivacyStatus field to given value.

### HasPrivacyStatus

`func (o *OrgsOwnershipHistory200ResponseRowsInner) HasPrivacyStatus() bool`

HasPrivacyStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


