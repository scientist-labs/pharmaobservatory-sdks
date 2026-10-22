# RareDiseaseOrganizationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionKind** | Pointer to **NullableString** | disease_relation when connection_value is a relation to the disease, disease_function when it is a function the organization performs for it. | [optional] 
**AssertionId** | Pointer to **NullableString** | The identifier of the assertion. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**DiseaseName** | Pointer to **NullableString** | The disease&#39;s preferred English name. | [optional] 
**OrganizationName** | Pointer to **NullableString** | The organization&#39;s name as filed. | [optional] 
**OrganizationOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record the assertion is on. | [optional] 
**OrganizationCompanyId** | Pointer to **NullableString** | Company level: the organization&#39;s company (null when unresolved). | [optional] 
**OrganizationCompanyName** | Pointer to **NullableString** | Display name of the organization&#39;s company. | [optional] 
**OrganizationGroupId** | Pointer to **NullableString** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**OrganizationGroupName** | Pointer to **NullableString** | Display name of the organization&#39;s group. | [optional] 
**ConnectionValue** | Pointer to **NullableString** | The relation (a predicate) or the function, according to connection_kind. | [optional] 
**EvidenceType** | Pointer to **NullableString** | The kind of evidence the assertion rests on. | [optional] 
**AuthorityClass** | Pointer to **NullableString** | How authoritative the source is for the claim. | [optional] 
**SourceKey** | Pointer to **NullableString** | The source the assertion was read from. | [optional] 
**SourceRecordId** | Pointer to **NullableString** | The record in the source the assertion comes from. | [optional] 
**SourceUrl** | Pointer to **NullableString** | A link to the source record, where it has one. | [optional] 
**EvidenceTrialId** | Pointer to **NullableString** | The registry trial id the assertion names, where it names one. | [optional] 
**ObservedAt** | Pointer to **NullableTime** | When the assertion was observed in the source (UTC). | [optional] 
**LastVerifiedAt** | Pointer to **NullableTime** | When the assertion was last verified against the source (UTC). | [optional] 
**Confidence** | Pointer to **NullableFloat32** | The derivation&#39;s confidence in the assertion. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the source the assertion was read from. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseaseOrganizationsRow

`func NewRareDiseaseOrganizationsRow() *RareDiseaseOrganizationsRow`

NewRareDiseaseOrganizationsRow instantiates a new RareDiseaseOrganizationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseOrganizationsRowWithDefaults

`func NewRareDiseaseOrganizationsRowWithDefaults() *RareDiseaseOrganizationsRow`

NewRareDiseaseOrganizationsRowWithDefaults instantiates a new RareDiseaseOrganizationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionKind

`func (o *RareDiseaseOrganizationsRow) GetConnectionKind() string`

GetConnectionKind returns the ConnectionKind field if non-nil, zero value otherwise.

### GetConnectionKindOk

`func (o *RareDiseaseOrganizationsRow) GetConnectionKindOk() (*string, bool)`

GetConnectionKindOk returns a tuple with the ConnectionKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionKind

`func (o *RareDiseaseOrganizationsRow) SetConnectionKind(v string)`

SetConnectionKind sets ConnectionKind field to given value.

### HasConnectionKind

`func (o *RareDiseaseOrganizationsRow) HasConnectionKind() bool`

HasConnectionKind returns a boolean if a field has been set.

### SetConnectionKindNil

`func (o *RareDiseaseOrganizationsRow) SetConnectionKindNil(b bool)`

 SetConnectionKindNil sets the value for ConnectionKind to be an explicit nil

### UnsetConnectionKind
`func (o *RareDiseaseOrganizationsRow) UnsetConnectionKind()`

UnsetConnectionKind ensures that no value is present for ConnectionKind, not even an explicit nil
### GetAssertionId

`func (o *RareDiseaseOrganizationsRow) GetAssertionId() string`

GetAssertionId returns the AssertionId field if non-nil, zero value otherwise.

### GetAssertionIdOk

`func (o *RareDiseaseOrganizationsRow) GetAssertionIdOk() (*string, bool)`

GetAssertionIdOk returns a tuple with the AssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionId

`func (o *RareDiseaseOrganizationsRow) SetAssertionId(v string)`

SetAssertionId sets AssertionId field to given value.

### HasAssertionId

`func (o *RareDiseaseOrganizationsRow) HasAssertionId() bool`

HasAssertionId returns a boolean if a field has been set.

### SetAssertionIdNil

`func (o *RareDiseaseOrganizationsRow) SetAssertionIdNil(b bool)`

 SetAssertionIdNil sets the value for AssertionId to be an explicit nil

### UnsetAssertionId
`func (o *RareDiseaseOrganizationsRow) UnsetAssertionId()`

UnsetAssertionId ensures that no value is present for AssertionId, not even an explicit nil
### GetOrphaCode

`func (o *RareDiseaseOrganizationsRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseOrganizationsRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseOrganizationsRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseOrganizationsRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseOrganizationsRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseOrganizationsRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetDiseaseName

`func (o *RareDiseaseOrganizationsRow) GetDiseaseName() string`

GetDiseaseName returns the DiseaseName field if non-nil, zero value otherwise.

### GetDiseaseNameOk

`func (o *RareDiseaseOrganizationsRow) GetDiseaseNameOk() (*string, bool)`

GetDiseaseNameOk returns a tuple with the DiseaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiseaseName

`func (o *RareDiseaseOrganizationsRow) SetDiseaseName(v string)`

SetDiseaseName sets DiseaseName field to given value.

### HasDiseaseName

`func (o *RareDiseaseOrganizationsRow) HasDiseaseName() bool`

HasDiseaseName returns a boolean if a field has been set.

### SetDiseaseNameNil

`func (o *RareDiseaseOrganizationsRow) SetDiseaseNameNil(b bool)`

 SetDiseaseNameNil sets the value for DiseaseName to be an explicit nil

### UnsetDiseaseName
`func (o *RareDiseaseOrganizationsRow) UnsetDiseaseName()`

UnsetDiseaseName ensures that no value is present for DiseaseName, not even an explicit nil
### GetOrganizationName

`func (o *RareDiseaseOrganizationsRow) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *RareDiseaseOrganizationsRow) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *RareDiseaseOrganizationsRow) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *RareDiseaseOrganizationsRow) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *RareDiseaseOrganizationsRow) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *RareDiseaseOrganizationsRow) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationOrgId

`func (o *RareDiseaseOrganizationsRow) GetOrganizationOrgId() string`

GetOrganizationOrgId returns the OrganizationOrgId field if non-nil, zero value otherwise.

### GetOrganizationOrgIdOk

`func (o *RareDiseaseOrganizationsRow) GetOrganizationOrgIdOk() (*string, bool)`

GetOrganizationOrgIdOk returns a tuple with the OrganizationOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationOrgId

`func (o *RareDiseaseOrganizationsRow) SetOrganizationOrgId(v string)`

SetOrganizationOrgId sets OrganizationOrgId field to given value.

### HasOrganizationOrgId

`func (o *RareDiseaseOrganizationsRow) HasOrganizationOrgId() bool`

HasOrganizationOrgId returns a boolean if a field has been set.

### SetOrganizationOrgIdNil

`func (o *RareDiseaseOrganizationsRow) SetOrganizationOrgIdNil(b bool)`

 SetOrganizationOrgIdNil sets the value for OrganizationOrgId to be an explicit nil

### UnsetOrganizationOrgId
`func (o *RareDiseaseOrganizationsRow) UnsetOrganizationOrgId()`

UnsetOrganizationOrgId ensures that no value is present for OrganizationOrgId, not even an explicit nil
### GetOrganizationCompanyId

`func (o *RareDiseaseOrganizationsRow) GetOrganizationCompanyId() string`

GetOrganizationCompanyId returns the OrganizationCompanyId field if non-nil, zero value otherwise.

### GetOrganizationCompanyIdOk

`func (o *RareDiseaseOrganizationsRow) GetOrganizationCompanyIdOk() (*string, bool)`

GetOrganizationCompanyIdOk returns a tuple with the OrganizationCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyId

`func (o *RareDiseaseOrganizationsRow) SetOrganizationCompanyId(v string)`

SetOrganizationCompanyId sets OrganizationCompanyId field to given value.

### HasOrganizationCompanyId

`func (o *RareDiseaseOrganizationsRow) HasOrganizationCompanyId() bool`

HasOrganizationCompanyId returns a boolean if a field has been set.

### SetOrganizationCompanyIdNil

`func (o *RareDiseaseOrganizationsRow) SetOrganizationCompanyIdNil(b bool)`

 SetOrganizationCompanyIdNil sets the value for OrganizationCompanyId to be an explicit nil

### UnsetOrganizationCompanyId
`func (o *RareDiseaseOrganizationsRow) UnsetOrganizationCompanyId()`

UnsetOrganizationCompanyId ensures that no value is present for OrganizationCompanyId, not even an explicit nil
### GetOrganizationCompanyName

`func (o *RareDiseaseOrganizationsRow) GetOrganizationCompanyName() string`

GetOrganizationCompanyName returns the OrganizationCompanyName field if non-nil, zero value otherwise.

### GetOrganizationCompanyNameOk

`func (o *RareDiseaseOrganizationsRow) GetOrganizationCompanyNameOk() (*string, bool)`

GetOrganizationCompanyNameOk returns a tuple with the OrganizationCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyName

`func (o *RareDiseaseOrganizationsRow) SetOrganizationCompanyName(v string)`

SetOrganizationCompanyName sets OrganizationCompanyName field to given value.

### HasOrganizationCompanyName

`func (o *RareDiseaseOrganizationsRow) HasOrganizationCompanyName() bool`

HasOrganizationCompanyName returns a boolean if a field has been set.

### SetOrganizationCompanyNameNil

`func (o *RareDiseaseOrganizationsRow) SetOrganizationCompanyNameNil(b bool)`

 SetOrganizationCompanyNameNil sets the value for OrganizationCompanyName to be an explicit nil

### UnsetOrganizationCompanyName
`func (o *RareDiseaseOrganizationsRow) UnsetOrganizationCompanyName()`

UnsetOrganizationCompanyName ensures that no value is present for OrganizationCompanyName, not even an explicit nil
### GetOrganizationGroupId

`func (o *RareDiseaseOrganizationsRow) GetOrganizationGroupId() string`

GetOrganizationGroupId returns the OrganizationGroupId field if non-nil, zero value otherwise.

### GetOrganizationGroupIdOk

`func (o *RareDiseaseOrganizationsRow) GetOrganizationGroupIdOk() (*string, bool)`

GetOrganizationGroupIdOk returns a tuple with the OrganizationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupId

`func (o *RareDiseaseOrganizationsRow) SetOrganizationGroupId(v string)`

SetOrganizationGroupId sets OrganizationGroupId field to given value.

### HasOrganizationGroupId

`func (o *RareDiseaseOrganizationsRow) HasOrganizationGroupId() bool`

HasOrganizationGroupId returns a boolean if a field has been set.

### SetOrganizationGroupIdNil

`func (o *RareDiseaseOrganizationsRow) SetOrganizationGroupIdNil(b bool)`

 SetOrganizationGroupIdNil sets the value for OrganizationGroupId to be an explicit nil

### UnsetOrganizationGroupId
`func (o *RareDiseaseOrganizationsRow) UnsetOrganizationGroupId()`

UnsetOrganizationGroupId ensures that no value is present for OrganizationGroupId, not even an explicit nil
### GetOrganizationGroupName

`func (o *RareDiseaseOrganizationsRow) GetOrganizationGroupName() string`

GetOrganizationGroupName returns the OrganizationGroupName field if non-nil, zero value otherwise.

### GetOrganizationGroupNameOk

`func (o *RareDiseaseOrganizationsRow) GetOrganizationGroupNameOk() (*string, bool)`

GetOrganizationGroupNameOk returns a tuple with the OrganizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupName

`func (o *RareDiseaseOrganizationsRow) SetOrganizationGroupName(v string)`

SetOrganizationGroupName sets OrganizationGroupName field to given value.

### HasOrganizationGroupName

`func (o *RareDiseaseOrganizationsRow) HasOrganizationGroupName() bool`

HasOrganizationGroupName returns a boolean if a field has been set.

### SetOrganizationGroupNameNil

`func (o *RareDiseaseOrganizationsRow) SetOrganizationGroupNameNil(b bool)`

 SetOrganizationGroupNameNil sets the value for OrganizationGroupName to be an explicit nil

### UnsetOrganizationGroupName
`func (o *RareDiseaseOrganizationsRow) UnsetOrganizationGroupName()`

UnsetOrganizationGroupName ensures that no value is present for OrganizationGroupName, not even an explicit nil
### GetConnectionValue

`func (o *RareDiseaseOrganizationsRow) GetConnectionValue() string`

GetConnectionValue returns the ConnectionValue field if non-nil, zero value otherwise.

### GetConnectionValueOk

`func (o *RareDiseaseOrganizationsRow) GetConnectionValueOk() (*string, bool)`

GetConnectionValueOk returns a tuple with the ConnectionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionValue

`func (o *RareDiseaseOrganizationsRow) SetConnectionValue(v string)`

SetConnectionValue sets ConnectionValue field to given value.

### HasConnectionValue

`func (o *RareDiseaseOrganizationsRow) HasConnectionValue() bool`

HasConnectionValue returns a boolean if a field has been set.

### SetConnectionValueNil

`func (o *RareDiseaseOrganizationsRow) SetConnectionValueNil(b bool)`

 SetConnectionValueNil sets the value for ConnectionValue to be an explicit nil

### UnsetConnectionValue
`func (o *RareDiseaseOrganizationsRow) UnsetConnectionValue()`

UnsetConnectionValue ensures that no value is present for ConnectionValue, not even an explicit nil
### GetEvidenceType

`func (o *RareDiseaseOrganizationsRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *RareDiseaseOrganizationsRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *RareDiseaseOrganizationsRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *RareDiseaseOrganizationsRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *RareDiseaseOrganizationsRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *RareDiseaseOrganizationsRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetAuthorityClass

`func (o *RareDiseaseOrganizationsRow) GetAuthorityClass() string`

GetAuthorityClass returns the AuthorityClass field if non-nil, zero value otherwise.

### GetAuthorityClassOk

`func (o *RareDiseaseOrganizationsRow) GetAuthorityClassOk() (*string, bool)`

GetAuthorityClassOk returns a tuple with the AuthorityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityClass

`func (o *RareDiseaseOrganizationsRow) SetAuthorityClass(v string)`

SetAuthorityClass sets AuthorityClass field to given value.

### HasAuthorityClass

`func (o *RareDiseaseOrganizationsRow) HasAuthorityClass() bool`

HasAuthorityClass returns a boolean if a field has been set.

### SetAuthorityClassNil

`func (o *RareDiseaseOrganizationsRow) SetAuthorityClassNil(b bool)`

 SetAuthorityClassNil sets the value for AuthorityClass to be an explicit nil

### UnsetAuthorityClass
`func (o *RareDiseaseOrganizationsRow) UnsetAuthorityClass()`

UnsetAuthorityClass ensures that no value is present for AuthorityClass, not even an explicit nil
### GetSourceKey

`func (o *RareDiseaseOrganizationsRow) GetSourceKey() string`

GetSourceKey returns the SourceKey field if non-nil, zero value otherwise.

### GetSourceKeyOk

`func (o *RareDiseaseOrganizationsRow) GetSourceKeyOk() (*string, bool)`

GetSourceKeyOk returns a tuple with the SourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKey

`func (o *RareDiseaseOrganizationsRow) SetSourceKey(v string)`

SetSourceKey sets SourceKey field to given value.

### HasSourceKey

`func (o *RareDiseaseOrganizationsRow) HasSourceKey() bool`

HasSourceKey returns a boolean if a field has been set.

### SetSourceKeyNil

`func (o *RareDiseaseOrganizationsRow) SetSourceKeyNil(b bool)`

 SetSourceKeyNil sets the value for SourceKey to be an explicit nil

### UnsetSourceKey
`func (o *RareDiseaseOrganizationsRow) UnsetSourceKey()`

UnsetSourceKey ensures that no value is present for SourceKey, not even an explicit nil
### GetSourceRecordId

`func (o *RareDiseaseOrganizationsRow) GetSourceRecordId() string`

GetSourceRecordId returns the SourceRecordId field if non-nil, zero value otherwise.

### GetSourceRecordIdOk

`func (o *RareDiseaseOrganizationsRow) GetSourceRecordIdOk() (*string, bool)`

GetSourceRecordIdOk returns a tuple with the SourceRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRecordId

`func (o *RareDiseaseOrganizationsRow) SetSourceRecordId(v string)`

SetSourceRecordId sets SourceRecordId field to given value.

### HasSourceRecordId

`func (o *RareDiseaseOrganizationsRow) HasSourceRecordId() bool`

HasSourceRecordId returns a boolean if a field has been set.

### SetSourceRecordIdNil

`func (o *RareDiseaseOrganizationsRow) SetSourceRecordIdNil(b bool)`

 SetSourceRecordIdNil sets the value for SourceRecordId to be an explicit nil

### UnsetSourceRecordId
`func (o *RareDiseaseOrganizationsRow) UnsetSourceRecordId()`

UnsetSourceRecordId ensures that no value is present for SourceRecordId, not even an explicit nil
### GetSourceUrl

`func (o *RareDiseaseOrganizationsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RareDiseaseOrganizationsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RareDiseaseOrganizationsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RareDiseaseOrganizationsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *RareDiseaseOrganizationsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *RareDiseaseOrganizationsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetEvidenceTrialId

`func (o *RareDiseaseOrganizationsRow) GetEvidenceTrialId() string`

GetEvidenceTrialId returns the EvidenceTrialId field if non-nil, zero value otherwise.

### GetEvidenceTrialIdOk

`func (o *RareDiseaseOrganizationsRow) GetEvidenceTrialIdOk() (*string, bool)`

GetEvidenceTrialIdOk returns a tuple with the EvidenceTrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceTrialId

`func (o *RareDiseaseOrganizationsRow) SetEvidenceTrialId(v string)`

SetEvidenceTrialId sets EvidenceTrialId field to given value.

### HasEvidenceTrialId

`func (o *RareDiseaseOrganizationsRow) HasEvidenceTrialId() bool`

HasEvidenceTrialId returns a boolean if a field has been set.

### SetEvidenceTrialIdNil

`func (o *RareDiseaseOrganizationsRow) SetEvidenceTrialIdNil(b bool)`

 SetEvidenceTrialIdNil sets the value for EvidenceTrialId to be an explicit nil

### UnsetEvidenceTrialId
`func (o *RareDiseaseOrganizationsRow) UnsetEvidenceTrialId()`

UnsetEvidenceTrialId ensures that no value is present for EvidenceTrialId, not even an explicit nil
### GetObservedAt

`func (o *RareDiseaseOrganizationsRow) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *RareDiseaseOrganizationsRow) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *RareDiseaseOrganizationsRow) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *RareDiseaseOrganizationsRow) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### SetObservedAtNil

`func (o *RareDiseaseOrganizationsRow) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *RareDiseaseOrganizationsRow) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil
### GetLastVerifiedAt

`func (o *RareDiseaseOrganizationsRow) GetLastVerifiedAt() time.Time`

GetLastVerifiedAt returns the LastVerifiedAt field if non-nil, zero value otherwise.

### GetLastVerifiedAtOk

`func (o *RareDiseaseOrganizationsRow) GetLastVerifiedAtOk() (*time.Time, bool)`

GetLastVerifiedAtOk returns a tuple with the LastVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedAt

`func (o *RareDiseaseOrganizationsRow) SetLastVerifiedAt(v time.Time)`

SetLastVerifiedAt sets LastVerifiedAt field to given value.

### HasLastVerifiedAt

`func (o *RareDiseaseOrganizationsRow) HasLastVerifiedAt() bool`

HasLastVerifiedAt returns a boolean if a field has been set.

### SetLastVerifiedAtNil

`func (o *RareDiseaseOrganizationsRow) SetLastVerifiedAtNil(b bool)`

 SetLastVerifiedAtNil sets the value for LastVerifiedAt to be an explicit nil

### UnsetLastVerifiedAt
`func (o *RareDiseaseOrganizationsRow) UnsetLastVerifiedAt()`

UnsetLastVerifiedAt ensures that no value is present for LastVerifiedAt, not even an explicit nil
### GetConfidence

`func (o *RareDiseaseOrganizationsRow) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RareDiseaseOrganizationsRow) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RareDiseaseOrganizationsRow) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RareDiseaseOrganizationsRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *RareDiseaseOrganizationsRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *RareDiseaseOrganizationsRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseOrganizationsRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseOrganizationsRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseOrganizationsRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseOrganizationsRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseOrganizationsRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseOrganizationsRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseOrganizationsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseOrganizationsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseOrganizationsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseOrganizationsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseOrganizationsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseOrganizationsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


