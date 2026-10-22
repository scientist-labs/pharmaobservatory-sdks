# CompanyRareDiseaseDeclarationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssertionId** | Pointer to **NullableString** | The declaration assertion id. | [optional] 
**OrganizationName** | Pointer to **NullableString** | The organization&#39;s name as the source filed it. | [optional] 
**OrganizationOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record of the organization. | [optional] 
**OrganizationCompanyId** | Pointer to **NullableString** | Company level: the company the organization belongs to (null when unresolved). | [optional] 
**OrganizationCompanyName** | Pointer to **NullableString** | Display name of the organization&#39;s company. | [optional] 
**OrganizationGroupId** | Pointer to **NullableString** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**OrganizationGroupName** | Pointer to **NullableString** | Display name of the organization&#39;s group. | [optional] 
**DeclaredFunction** | Pointer to **NullableString** | The function the source states for the organization. | [optional] 
**EvidenceType** | Pointer to **NullableString** | The kind of evidence the source gives. | [optional] 
**AuthorityClass** | Pointer to **NullableString** | How authoritative the source is for the claim. | [optional] 
**SourceKey** | Pointer to **NullableString** | The source the assertion comes from. | [optional] 
**SourceRecordId** | Pointer to **NullableString** | The source&#39;s own record id. | [optional] 
**SourceUrl** | Pointer to **NullableString** | A link to the source record. | [optional] 
**ObservedAt** | Pointer to **NullableTime** | When the assertion was observed (UTC). | [optional] 
**LastVerifiedAt** | Pointer to **NullableTime** | When the assertion was last verified (UTC). | [optional] 
**Confidence** | Pointer to **NullableFloat32** | The source&#39;s or the extraction&#39;s own confidence; null when none is stated. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the source. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit line for the source. | [optional] 

## Methods

### NewCompanyRareDiseaseDeclarationsRow

`func NewCompanyRareDiseaseDeclarationsRow() *CompanyRareDiseaseDeclarationsRow`

NewCompanyRareDiseaseDeclarationsRow instantiates a new CompanyRareDiseaseDeclarationsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyRareDiseaseDeclarationsRowWithDefaults

`func NewCompanyRareDiseaseDeclarationsRowWithDefaults() *CompanyRareDiseaseDeclarationsRow`

NewCompanyRareDiseaseDeclarationsRowWithDefaults instantiates a new CompanyRareDiseaseDeclarationsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertionId

`func (o *CompanyRareDiseaseDeclarationsRow) GetAssertionId() string`

GetAssertionId returns the AssertionId field if non-nil, zero value otherwise.

### GetAssertionIdOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetAssertionIdOk() (*string, bool)`

GetAssertionIdOk returns a tuple with the AssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionId

`func (o *CompanyRareDiseaseDeclarationsRow) SetAssertionId(v string)`

SetAssertionId sets AssertionId field to given value.

### HasAssertionId

`func (o *CompanyRareDiseaseDeclarationsRow) HasAssertionId() bool`

HasAssertionId returns a boolean if a field has been set.

### SetAssertionIdNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetAssertionIdNil(b bool)`

 SetAssertionIdNil sets the value for AssertionId to be an explicit nil

### UnsetAssertionId
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetAssertionId()`

UnsetAssertionId ensures that no value is present for AssertionId, not even an explicit nil
### GetOrganizationName

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *CompanyRareDiseaseDeclarationsRow) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationOrgId

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationOrgId() string`

GetOrganizationOrgId returns the OrganizationOrgId field if non-nil, zero value otherwise.

### GetOrganizationOrgIdOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationOrgIdOk() (*string, bool)`

GetOrganizationOrgIdOk returns a tuple with the OrganizationOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationOrgId

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationOrgId(v string)`

SetOrganizationOrgId sets OrganizationOrgId field to given value.

### HasOrganizationOrgId

`func (o *CompanyRareDiseaseDeclarationsRow) HasOrganizationOrgId() bool`

HasOrganizationOrgId returns a boolean if a field has been set.

### SetOrganizationOrgIdNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationOrgIdNil(b bool)`

 SetOrganizationOrgIdNil sets the value for OrganizationOrgId to be an explicit nil

### UnsetOrganizationOrgId
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetOrganizationOrgId()`

UnsetOrganizationOrgId ensures that no value is present for OrganizationOrgId, not even an explicit nil
### GetOrganizationCompanyId

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationCompanyId() string`

GetOrganizationCompanyId returns the OrganizationCompanyId field if non-nil, zero value otherwise.

### GetOrganizationCompanyIdOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationCompanyIdOk() (*string, bool)`

GetOrganizationCompanyIdOk returns a tuple with the OrganizationCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyId

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationCompanyId(v string)`

SetOrganizationCompanyId sets OrganizationCompanyId field to given value.

### HasOrganizationCompanyId

`func (o *CompanyRareDiseaseDeclarationsRow) HasOrganizationCompanyId() bool`

HasOrganizationCompanyId returns a boolean if a field has been set.

### SetOrganizationCompanyIdNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationCompanyIdNil(b bool)`

 SetOrganizationCompanyIdNil sets the value for OrganizationCompanyId to be an explicit nil

### UnsetOrganizationCompanyId
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetOrganizationCompanyId()`

UnsetOrganizationCompanyId ensures that no value is present for OrganizationCompanyId, not even an explicit nil
### GetOrganizationCompanyName

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationCompanyName() string`

GetOrganizationCompanyName returns the OrganizationCompanyName field if non-nil, zero value otherwise.

### GetOrganizationCompanyNameOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationCompanyNameOk() (*string, bool)`

GetOrganizationCompanyNameOk returns a tuple with the OrganizationCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyName

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationCompanyName(v string)`

SetOrganizationCompanyName sets OrganizationCompanyName field to given value.

### HasOrganizationCompanyName

`func (o *CompanyRareDiseaseDeclarationsRow) HasOrganizationCompanyName() bool`

HasOrganizationCompanyName returns a boolean if a field has been set.

### SetOrganizationCompanyNameNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationCompanyNameNil(b bool)`

 SetOrganizationCompanyNameNil sets the value for OrganizationCompanyName to be an explicit nil

### UnsetOrganizationCompanyName
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetOrganizationCompanyName()`

UnsetOrganizationCompanyName ensures that no value is present for OrganizationCompanyName, not even an explicit nil
### GetOrganizationGroupId

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationGroupId() string`

GetOrganizationGroupId returns the OrganizationGroupId field if non-nil, zero value otherwise.

### GetOrganizationGroupIdOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationGroupIdOk() (*string, bool)`

GetOrganizationGroupIdOk returns a tuple with the OrganizationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupId

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationGroupId(v string)`

SetOrganizationGroupId sets OrganizationGroupId field to given value.

### HasOrganizationGroupId

`func (o *CompanyRareDiseaseDeclarationsRow) HasOrganizationGroupId() bool`

HasOrganizationGroupId returns a boolean if a field has been set.

### SetOrganizationGroupIdNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationGroupIdNil(b bool)`

 SetOrganizationGroupIdNil sets the value for OrganizationGroupId to be an explicit nil

### UnsetOrganizationGroupId
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetOrganizationGroupId()`

UnsetOrganizationGroupId ensures that no value is present for OrganizationGroupId, not even an explicit nil
### GetOrganizationGroupName

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationGroupName() string`

GetOrganizationGroupName returns the OrganizationGroupName field if non-nil, zero value otherwise.

### GetOrganizationGroupNameOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetOrganizationGroupNameOk() (*string, bool)`

GetOrganizationGroupNameOk returns a tuple with the OrganizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupName

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationGroupName(v string)`

SetOrganizationGroupName sets OrganizationGroupName field to given value.

### HasOrganizationGroupName

`func (o *CompanyRareDiseaseDeclarationsRow) HasOrganizationGroupName() bool`

HasOrganizationGroupName returns a boolean if a field has been set.

### SetOrganizationGroupNameNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetOrganizationGroupNameNil(b bool)`

 SetOrganizationGroupNameNil sets the value for OrganizationGroupName to be an explicit nil

### UnsetOrganizationGroupName
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetOrganizationGroupName()`

UnsetOrganizationGroupName ensures that no value is present for OrganizationGroupName, not even an explicit nil
### GetDeclaredFunction

`func (o *CompanyRareDiseaseDeclarationsRow) GetDeclaredFunction() string`

GetDeclaredFunction returns the DeclaredFunction field if non-nil, zero value otherwise.

### GetDeclaredFunctionOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetDeclaredFunctionOk() (*string, bool)`

GetDeclaredFunctionOk returns a tuple with the DeclaredFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredFunction

`func (o *CompanyRareDiseaseDeclarationsRow) SetDeclaredFunction(v string)`

SetDeclaredFunction sets DeclaredFunction field to given value.

### HasDeclaredFunction

`func (o *CompanyRareDiseaseDeclarationsRow) HasDeclaredFunction() bool`

HasDeclaredFunction returns a boolean if a field has been set.

### SetDeclaredFunctionNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetDeclaredFunctionNil(b bool)`

 SetDeclaredFunctionNil sets the value for DeclaredFunction to be an explicit nil

### UnsetDeclaredFunction
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetDeclaredFunction()`

UnsetDeclaredFunction ensures that no value is present for DeclaredFunction, not even an explicit nil
### GetEvidenceType

`func (o *CompanyRareDiseaseDeclarationsRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *CompanyRareDiseaseDeclarationsRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *CompanyRareDiseaseDeclarationsRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetAuthorityClass

`func (o *CompanyRareDiseaseDeclarationsRow) GetAuthorityClass() string`

GetAuthorityClass returns the AuthorityClass field if non-nil, zero value otherwise.

### GetAuthorityClassOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetAuthorityClassOk() (*string, bool)`

GetAuthorityClassOk returns a tuple with the AuthorityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityClass

`func (o *CompanyRareDiseaseDeclarationsRow) SetAuthorityClass(v string)`

SetAuthorityClass sets AuthorityClass field to given value.

### HasAuthorityClass

`func (o *CompanyRareDiseaseDeclarationsRow) HasAuthorityClass() bool`

HasAuthorityClass returns a boolean if a field has been set.

### SetAuthorityClassNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetAuthorityClassNil(b bool)`

 SetAuthorityClassNil sets the value for AuthorityClass to be an explicit nil

### UnsetAuthorityClass
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetAuthorityClass()`

UnsetAuthorityClass ensures that no value is present for AuthorityClass, not even an explicit nil
### GetSourceKey

`func (o *CompanyRareDiseaseDeclarationsRow) GetSourceKey() string`

GetSourceKey returns the SourceKey field if non-nil, zero value otherwise.

### GetSourceKeyOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetSourceKeyOk() (*string, bool)`

GetSourceKeyOk returns a tuple with the SourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKey

`func (o *CompanyRareDiseaseDeclarationsRow) SetSourceKey(v string)`

SetSourceKey sets SourceKey field to given value.

### HasSourceKey

`func (o *CompanyRareDiseaseDeclarationsRow) HasSourceKey() bool`

HasSourceKey returns a boolean if a field has been set.

### SetSourceKeyNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetSourceKeyNil(b bool)`

 SetSourceKeyNil sets the value for SourceKey to be an explicit nil

### UnsetSourceKey
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetSourceKey()`

UnsetSourceKey ensures that no value is present for SourceKey, not even an explicit nil
### GetSourceRecordId

`func (o *CompanyRareDiseaseDeclarationsRow) GetSourceRecordId() string`

GetSourceRecordId returns the SourceRecordId field if non-nil, zero value otherwise.

### GetSourceRecordIdOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetSourceRecordIdOk() (*string, bool)`

GetSourceRecordIdOk returns a tuple with the SourceRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRecordId

`func (o *CompanyRareDiseaseDeclarationsRow) SetSourceRecordId(v string)`

SetSourceRecordId sets SourceRecordId field to given value.

### HasSourceRecordId

`func (o *CompanyRareDiseaseDeclarationsRow) HasSourceRecordId() bool`

HasSourceRecordId returns a boolean if a field has been set.

### SetSourceRecordIdNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetSourceRecordIdNil(b bool)`

 SetSourceRecordIdNil sets the value for SourceRecordId to be an explicit nil

### UnsetSourceRecordId
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetSourceRecordId()`

UnsetSourceRecordId ensures that no value is present for SourceRecordId, not even an explicit nil
### GetSourceUrl

`func (o *CompanyRareDiseaseDeclarationsRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CompanyRareDiseaseDeclarationsRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *CompanyRareDiseaseDeclarationsRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetObservedAt

`func (o *CompanyRareDiseaseDeclarationsRow) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *CompanyRareDiseaseDeclarationsRow) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *CompanyRareDiseaseDeclarationsRow) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### SetObservedAtNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil
### GetLastVerifiedAt

`func (o *CompanyRareDiseaseDeclarationsRow) GetLastVerifiedAt() time.Time`

GetLastVerifiedAt returns the LastVerifiedAt field if non-nil, zero value otherwise.

### GetLastVerifiedAtOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetLastVerifiedAtOk() (*time.Time, bool)`

GetLastVerifiedAtOk returns a tuple with the LastVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedAt

`func (o *CompanyRareDiseaseDeclarationsRow) SetLastVerifiedAt(v time.Time)`

SetLastVerifiedAt sets LastVerifiedAt field to given value.

### HasLastVerifiedAt

`func (o *CompanyRareDiseaseDeclarationsRow) HasLastVerifiedAt() bool`

HasLastVerifiedAt returns a boolean if a field has been set.

### SetLastVerifiedAtNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetLastVerifiedAtNil(b bool)`

 SetLastVerifiedAtNil sets the value for LastVerifiedAt to be an explicit nil

### UnsetLastVerifiedAt
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetLastVerifiedAt()`

UnsetLastVerifiedAt ensures that no value is present for LastVerifiedAt, not even an explicit nil
### GetConfidence

`func (o *CompanyRareDiseaseDeclarationsRow) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CompanyRareDiseaseDeclarationsRow) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *CompanyRareDiseaseDeclarationsRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetLicenseName

`func (o *CompanyRareDiseaseDeclarationsRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *CompanyRareDiseaseDeclarationsRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *CompanyRareDiseaseDeclarationsRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *CompanyRareDiseaseDeclarationsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *CompanyRareDiseaseDeclarationsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *CompanyRareDiseaseDeclarationsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *CompanyRareDiseaseDeclarationsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *CompanyRareDiseaseDeclarationsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *CompanyRareDiseaseDeclarationsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


