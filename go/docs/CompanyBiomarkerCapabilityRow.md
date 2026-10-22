# CompanyBiomarkerCapabilityRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiomarkerValidationId** | Pointer to **NullableString** | The evidence assertion id. | [optional] 
**OrganizationName** | Pointer to **NullableString** | The organization&#39;s name as filed on the evidence. | [optional] 
**OrganizationOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record of the organization. | [optional] 
**OrganizationCompanyId** | Pointer to **NullableString** | Company level: the company the organization belongs to (null when unresolved). | [optional] 
**OrganizationCompanyName** | Pointer to **NullableString** | Display name of the organization&#39;s company. | [optional] 
**OrganizationGroupId** | Pointer to **NullableString** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**OrganizationGroupName** | Pointer to **NullableString** | Display name of the organization&#39;s group. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene&#39;s HGNC (Human Genome Organisation Gene Nomenclature Committee) symbol; null for evidence that names no gene. | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The biomarker&#39;s id (a hash, not the route key). | [optional] 
**BiomarkerDisplayName** | Pointer to **NullableString** | The biomarker&#39;s display name. | [optional] 
**BiomarkerKind** | Pointer to **NullableString** | The kind of biomarker. | [optional] 
**BiomarkerIdentityStatus** | Pointer to **NullableString** | How firmly the biomarker&#39;s identity is resolved. | [optional] 
**OrganizationRole** | Pointer to **NullableString** | The organization&#39;s role in the evidence (for example assay developer or testing laboratory). | [optional] 
**OrganizationMatchTier** | Pointer to **NullableString** | How the organization was matched to its record. | [optional] 
**Tier** | Pointer to **NullableInt32** | The evidence tier, higher for stronger kinds of evidence. A label of the kind of evidence, not a quality score. | [optional] 
**TierLabel** | Pointer to **NullableString** | The evidence tier in words. | [optional] 
**EvidenceType** | Pointer to **NullableString** | The kind of evidence (companion diagnostic authorization, trial endpoint, trial eligibility, trial specimen collection, trial text mention or ClinVar clinical testing). | [optional] 
**EvidenceCount** | Pointer to **NullableInt32** | Assertions that share this biomarker, organization and role. | [optional] 
**EvidenceRef** | Pointer to **NullableString** | The reference of the evidence (a trial id, a device application number or a project id). | [optional] 
**EvidenceRefKind** | Pointer to **NullableString** | What evidence_ref is. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial&#39;s registry id (NCT, National Clinical Trial, number) when the evidence is a trial. | [optional] 
**CitationUrl** | Pointer to **NullableString** | A link to the evidence. | [optional] 
**EvidenceSnippet** | Pointer to **NullableString** | A short excerpt of the evidence. | [optional] 
**BiomarkerRole** | Pointer to **NullableString** | The role the biomarker plays in the evidence. | [optional] 
**IsSurrogateEndpointClaim** | Pointer to **NullableBool** | Whether the evidence claims the biomarker as a surrogate endpoint. | [optional] 
**ClinvarScvCount** | Pointer to **NullableInt32** | ClinVar submissions behind a ClinVar row. | [optional] 

## Methods

### NewCompanyBiomarkerCapabilityRow

`func NewCompanyBiomarkerCapabilityRow() *CompanyBiomarkerCapabilityRow`

NewCompanyBiomarkerCapabilityRow instantiates a new CompanyBiomarkerCapabilityRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyBiomarkerCapabilityRowWithDefaults

`func NewCompanyBiomarkerCapabilityRowWithDefaults() *CompanyBiomarkerCapabilityRow`

NewCompanyBiomarkerCapabilityRowWithDefaults instantiates a new CompanyBiomarkerCapabilityRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiomarkerValidationId

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerValidationId() string`

GetBiomarkerValidationId returns the BiomarkerValidationId field if non-nil, zero value otherwise.

### GetBiomarkerValidationIdOk

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerValidationIdOk() (*string, bool)`

GetBiomarkerValidationIdOk returns a tuple with the BiomarkerValidationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerValidationId

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerValidationId(v string)`

SetBiomarkerValidationId sets BiomarkerValidationId field to given value.

### HasBiomarkerValidationId

`func (o *CompanyBiomarkerCapabilityRow) HasBiomarkerValidationId() bool`

HasBiomarkerValidationId returns a boolean if a field has been set.

### SetBiomarkerValidationIdNil

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerValidationIdNil(b bool)`

 SetBiomarkerValidationIdNil sets the value for BiomarkerValidationId to be an explicit nil

### UnsetBiomarkerValidationId
`func (o *CompanyBiomarkerCapabilityRow) UnsetBiomarkerValidationId()`

UnsetBiomarkerValidationId ensures that no value is present for BiomarkerValidationId, not even an explicit nil
### GetOrganizationName

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *CompanyBiomarkerCapabilityRow) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *CompanyBiomarkerCapabilityRow) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationOrgId

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationOrgId() string`

GetOrganizationOrgId returns the OrganizationOrgId field if non-nil, zero value otherwise.

### GetOrganizationOrgIdOk

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationOrgIdOk() (*string, bool)`

GetOrganizationOrgIdOk returns a tuple with the OrganizationOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationOrgId

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationOrgId(v string)`

SetOrganizationOrgId sets OrganizationOrgId field to given value.

### HasOrganizationOrgId

`func (o *CompanyBiomarkerCapabilityRow) HasOrganizationOrgId() bool`

HasOrganizationOrgId returns a boolean if a field has been set.

### SetOrganizationOrgIdNil

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationOrgIdNil(b bool)`

 SetOrganizationOrgIdNil sets the value for OrganizationOrgId to be an explicit nil

### UnsetOrganizationOrgId
`func (o *CompanyBiomarkerCapabilityRow) UnsetOrganizationOrgId()`

UnsetOrganizationOrgId ensures that no value is present for OrganizationOrgId, not even an explicit nil
### GetOrganizationCompanyId

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationCompanyId() string`

GetOrganizationCompanyId returns the OrganizationCompanyId field if non-nil, zero value otherwise.

### GetOrganizationCompanyIdOk

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationCompanyIdOk() (*string, bool)`

GetOrganizationCompanyIdOk returns a tuple with the OrganizationCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyId

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationCompanyId(v string)`

SetOrganizationCompanyId sets OrganizationCompanyId field to given value.

### HasOrganizationCompanyId

`func (o *CompanyBiomarkerCapabilityRow) HasOrganizationCompanyId() bool`

HasOrganizationCompanyId returns a boolean if a field has been set.

### SetOrganizationCompanyIdNil

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationCompanyIdNil(b bool)`

 SetOrganizationCompanyIdNil sets the value for OrganizationCompanyId to be an explicit nil

### UnsetOrganizationCompanyId
`func (o *CompanyBiomarkerCapabilityRow) UnsetOrganizationCompanyId()`

UnsetOrganizationCompanyId ensures that no value is present for OrganizationCompanyId, not even an explicit nil
### GetOrganizationCompanyName

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationCompanyName() string`

GetOrganizationCompanyName returns the OrganizationCompanyName field if non-nil, zero value otherwise.

### GetOrganizationCompanyNameOk

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationCompanyNameOk() (*string, bool)`

GetOrganizationCompanyNameOk returns a tuple with the OrganizationCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyName

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationCompanyName(v string)`

SetOrganizationCompanyName sets OrganizationCompanyName field to given value.

### HasOrganizationCompanyName

`func (o *CompanyBiomarkerCapabilityRow) HasOrganizationCompanyName() bool`

HasOrganizationCompanyName returns a boolean if a field has been set.

### SetOrganizationCompanyNameNil

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationCompanyNameNil(b bool)`

 SetOrganizationCompanyNameNil sets the value for OrganizationCompanyName to be an explicit nil

### UnsetOrganizationCompanyName
`func (o *CompanyBiomarkerCapabilityRow) UnsetOrganizationCompanyName()`

UnsetOrganizationCompanyName ensures that no value is present for OrganizationCompanyName, not even an explicit nil
### GetOrganizationGroupId

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationGroupId() string`

GetOrganizationGroupId returns the OrganizationGroupId field if non-nil, zero value otherwise.

### GetOrganizationGroupIdOk

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationGroupIdOk() (*string, bool)`

GetOrganizationGroupIdOk returns a tuple with the OrganizationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupId

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationGroupId(v string)`

SetOrganizationGroupId sets OrganizationGroupId field to given value.

### HasOrganizationGroupId

`func (o *CompanyBiomarkerCapabilityRow) HasOrganizationGroupId() bool`

HasOrganizationGroupId returns a boolean if a field has been set.

### SetOrganizationGroupIdNil

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationGroupIdNil(b bool)`

 SetOrganizationGroupIdNil sets the value for OrganizationGroupId to be an explicit nil

### UnsetOrganizationGroupId
`func (o *CompanyBiomarkerCapabilityRow) UnsetOrganizationGroupId()`

UnsetOrganizationGroupId ensures that no value is present for OrganizationGroupId, not even an explicit nil
### GetOrganizationGroupName

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationGroupName() string`

GetOrganizationGroupName returns the OrganizationGroupName field if non-nil, zero value otherwise.

### GetOrganizationGroupNameOk

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationGroupNameOk() (*string, bool)`

GetOrganizationGroupNameOk returns a tuple with the OrganizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupName

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationGroupName(v string)`

SetOrganizationGroupName sets OrganizationGroupName field to given value.

### HasOrganizationGroupName

`func (o *CompanyBiomarkerCapabilityRow) HasOrganizationGroupName() bool`

HasOrganizationGroupName returns a boolean if a field has been set.

### SetOrganizationGroupNameNil

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationGroupNameNil(b bool)`

 SetOrganizationGroupNameNil sets the value for OrganizationGroupName to be an explicit nil

### UnsetOrganizationGroupName
`func (o *CompanyBiomarkerCapabilityRow) UnsetOrganizationGroupName()`

UnsetOrganizationGroupName ensures that no value is present for OrganizationGroupName, not even an explicit nil
### GetGeneSymbol

`func (o *CompanyBiomarkerCapabilityRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *CompanyBiomarkerCapabilityRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *CompanyBiomarkerCapabilityRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *CompanyBiomarkerCapabilityRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *CompanyBiomarkerCapabilityRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *CompanyBiomarkerCapabilityRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetBiomarkerId

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *CompanyBiomarkerCapabilityRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *CompanyBiomarkerCapabilityRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetBiomarkerDisplayName

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerDisplayName() string`

GetBiomarkerDisplayName returns the BiomarkerDisplayName field if non-nil, zero value otherwise.

### GetBiomarkerDisplayNameOk

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerDisplayNameOk() (*string, bool)`

GetBiomarkerDisplayNameOk returns a tuple with the BiomarkerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerDisplayName

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerDisplayName(v string)`

SetBiomarkerDisplayName sets BiomarkerDisplayName field to given value.

### HasBiomarkerDisplayName

`func (o *CompanyBiomarkerCapabilityRow) HasBiomarkerDisplayName() bool`

HasBiomarkerDisplayName returns a boolean if a field has been set.

### SetBiomarkerDisplayNameNil

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerDisplayNameNil(b bool)`

 SetBiomarkerDisplayNameNil sets the value for BiomarkerDisplayName to be an explicit nil

### UnsetBiomarkerDisplayName
`func (o *CompanyBiomarkerCapabilityRow) UnsetBiomarkerDisplayName()`

UnsetBiomarkerDisplayName ensures that no value is present for BiomarkerDisplayName, not even an explicit nil
### GetBiomarkerKind

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerKind() string`

GetBiomarkerKind returns the BiomarkerKind field if non-nil, zero value otherwise.

### GetBiomarkerKindOk

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerKindOk() (*string, bool)`

GetBiomarkerKindOk returns a tuple with the BiomarkerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerKind

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerKind(v string)`

SetBiomarkerKind sets BiomarkerKind field to given value.

### HasBiomarkerKind

`func (o *CompanyBiomarkerCapabilityRow) HasBiomarkerKind() bool`

HasBiomarkerKind returns a boolean if a field has been set.

### SetBiomarkerKindNil

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerKindNil(b bool)`

 SetBiomarkerKindNil sets the value for BiomarkerKind to be an explicit nil

### UnsetBiomarkerKind
`func (o *CompanyBiomarkerCapabilityRow) UnsetBiomarkerKind()`

UnsetBiomarkerKind ensures that no value is present for BiomarkerKind, not even an explicit nil
### GetBiomarkerIdentityStatus

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerIdentityStatus() string`

GetBiomarkerIdentityStatus returns the BiomarkerIdentityStatus field if non-nil, zero value otherwise.

### GetBiomarkerIdentityStatusOk

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerIdentityStatusOk() (*string, bool)`

GetBiomarkerIdentityStatusOk returns a tuple with the BiomarkerIdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerIdentityStatus

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerIdentityStatus(v string)`

SetBiomarkerIdentityStatus sets BiomarkerIdentityStatus field to given value.

### HasBiomarkerIdentityStatus

`func (o *CompanyBiomarkerCapabilityRow) HasBiomarkerIdentityStatus() bool`

HasBiomarkerIdentityStatus returns a boolean if a field has been set.

### SetBiomarkerIdentityStatusNil

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerIdentityStatusNil(b bool)`

 SetBiomarkerIdentityStatusNil sets the value for BiomarkerIdentityStatus to be an explicit nil

### UnsetBiomarkerIdentityStatus
`func (o *CompanyBiomarkerCapabilityRow) UnsetBiomarkerIdentityStatus()`

UnsetBiomarkerIdentityStatus ensures that no value is present for BiomarkerIdentityStatus, not even an explicit nil
### GetOrganizationRole

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationRole() string`

GetOrganizationRole returns the OrganizationRole field if non-nil, zero value otherwise.

### GetOrganizationRoleOk

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationRoleOk() (*string, bool)`

GetOrganizationRoleOk returns a tuple with the OrganizationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRole

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationRole(v string)`

SetOrganizationRole sets OrganizationRole field to given value.

### HasOrganizationRole

`func (o *CompanyBiomarkerCapabilityRow) HasOrganizationRole() bool`

HasOrganizationRole returns a boolean if a field has been set.

### SetOrganizationRoleNil

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationRoleNil(b bool)`

 SetOrganizationRoleNil sets the value for OrganizationRole to be an explicit nil

### UnsetOrganizationRole
`func (o *CompanyBiomarkerCapabilityRow) UnsetOrganizationRole()`

UnsetOrganizationRole ensures that no value is present for OrganizationRole, not even an explicit nil
### GetOrganizationMatchTier

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationMatchTier() string`

GetOrganizationMatchTier returns the OrganizationMatchTier field if non-nil, zero value otherwise.

### GetOrganizationMatchTierOk

`func (o *CompanyBiomarkerCapabilityRow) GetOrganizationMatchTierOk() (*string, bool)`

GetOrganizationMatchTierOk returns a tuple with the OrganizationMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMatchTier

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationMatchTier(v string)`

SetOrganizationMatchTier sets OrganizationMatchTier field to given value.

### HasOrganizationMatchTier

`func (o *CompanyBiomarkerCapabilityRow) HasOrganizationMatchTier() bool`

HasOrganizationMatchTier returns a boolean if a field has been set.

### SetOrganizationMatchTierNil

`func (o *CompanyBiomarkerCapabilityRow) SetOrganizationMatchTierNil(b bool)`

 SetOrganizationMatchTierNil sets the value for OrganizationMatchTier to be an explicit nil

### UnsetOrganizationMatchTier
`func (o *CompanyBiomarkerCapabilityRow) UnsetOrganizationMatchTier()`

UnsetOrganizationMatchTier ensures that no value is present for OrganizationMatchTier, not even an explicit nil
### GetTier

`func (o *CompanyBiomarkerCapabilityRow) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CompanyBiomarkerCapabilityRow) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CompanyBiomarkerCapabilityRow) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CompanyBiomarkerCapabilityRow) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *CompanyBiomarkerCapabilityRow) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *CompanyBiomarkerCapabilityRow) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetTierLabel

`func (o *CompanyBiomarkerCapabilityRow) GetTierLabel() string`

GetTierLabel returns the TierLabel field if non-nil, zero value otherwise.

### GetTierLabelOk

`func (o *CompanyBiomarkerCapabilityRow) GetTierLabelOk() (*string, bool)`

GetTierLabelOk returns a tuple with the TierLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierLabel

`func (o *CompanyBiomarkerCapabilityRow) SetTierLabel(v string)`

SetTierLabel sets TierLabel field to given value.

### HasTierLabel

`func (o *CompanyBiomarkerCapabilityRow) HasTierLabel() bool`

HasTierLabel returns a boolean if a field has been set.

### SetTierLabelNil

`func (o *CompanyBiomarkerCapabilityRow) SetTierLabelNil(b bool)`

 SetTierLabelNil sets the value for TierLabel to be an explicit nil

### UnsetTierLabel
`func (o *CompanyBiomarkerCapabilityRow) UnsetTierLabel()`

UnsetTierLabel ensures that no value is present for TierLabel, not even an explicit nil
### GetEvidenceType

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *CompanyBiomarkerCapabilityRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *CompanyBiomarkerCapabilityRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetEvidenceCount

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceCount() int32`

GetEvidenceCount returns the EvidenceCount field if non-nil, zero value otherwise.

### GetEvidenceCountOk

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceCountOk() (*int32, bool)`

GetEvidenceCountOk returns a tuple with the EvidenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCount

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceCount(v int32)`

SetEvidenceCount sets EvidenceCount field to given value.

### HasEvidenceCount

`func (o *CompanyBiomarkerCapabilityRow) HasEvidenceCount() bool`

HasEvidenceCount returns a boolean if a field has been set.

### SetEvidenceCountNil

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceCountNil(b bool)`

 SetEvidenceCountNil sets the value for EvidenceCount to be an explicit nil

### UnsetEvidenceCount
`func (o *CompanyBiomarkerCapabilityRow) UnsetEvidenceCount()`

UnsetEvidenceCount ensures that no value is present for EvidenceCount, not even an explicit nil
### GetEvidenceRef

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceRef() string`

GetEvidenceRef returns the EvidenceRef field if non-nil, zero value otherwise.

### GetEvidenceRefOk

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceRefOk() (*string, bool)`

GetEvidenceRefOk returns a tuple with the EvidenceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRef

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceRef(v string)`

SetEvidenceRef sets EvidenceRef field to given value.

### HasEvidenceRef

`func (o *CompanyBiomarkerCapabilityRow) HasEvidenceRef() bool`

HasEvidenceRef returns a boolean if a field has been set.

### SetEvidenceRefNil

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceRefNil(b bool)`

 SetEvidenceRefNil sets the value for EvidenceRef to be an explicit nil

### UnsetEvidenceRef
`func (o *CompanyBiomarkerCapabilityRow) UnsetEvidenceRef()`

UnsetEvidenceRef ensures that no value is present for EvidenceRef, not even an explicit nil
### GetEvidenceRefKind

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceRefKind() string`

GetEvidenceRefKind returns the EvidenceRefKind field if non-nil, zero value otherwise.

### GetEvidenceRefKindOk

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceRefKindOk() (*string, bool)`

GetEvidenceRefKindOk returns a tuple with the EvidenceRefKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRefKind

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceRefKind(v string)`

SetEvidenceRefKind sets EvidenceRefKind field to given value.

### HasEvidenceRefKind

`func (o *CompanyBiomarkerCapabilityRow) HasEvidenceRefKind() bool`

HasEvidenceRefKind returns a boolean if a field has been set.

### SetEvidenceRefKindNil

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceRefKindNil(b bool)`

 SetEvidenceRefKindNil sets the value for EvidenceRefKind to be an explicit nil

### UnsetEvidenceRefKind
`func (o *CompanyBiomarkerCapabilityRow) UnsetEvidenceRefKind()`

UnsetEvidenceRefKind ensures that no value is present for EvidenceRefKind, not even an explicit nil
### GetTrialId

`func (o *CompanyBiomarkerCapabilityRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *CompanyBiomarkerCapabilityRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *CompanyBiomarkerCapabilityRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *CompanyBiomarkerCapabilityRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *CompanyBiomarkerCapabilityRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *CompanyBiomarkerCapabilityRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetCitationUrl

`func (o *CompanyBiomarkerCapabilityRow) GetCitationUrl() string`

GetCitationUrl returns the CitationUrl field if non-nil, zero value otherwise.

### GetCitationUrlOk

`func (o *CompanyBiomarkerCapabilityRow) GetCitationUrlOk() (*string, bool)`

GetCitationUrlOk returns a tuple with the CitationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationUrl

`func (o *CompanyBiomarkerCapabilityRow) SetCitationUrl(v string)`

SetCitationUrl sets CitationUrl field to given value.

### HasCitationUrl

`func (o *CompanyBiomarkerCapabilityRow) HasCitationUrl() bool`

HasCitationUrl returns a boolean if a field has been set.

### SetCitationUrlNil

`func (o *CompanyBiomarkerCapabilityRow) SetCitationUrlNil(b bool)`

 SetCitationUrlNil sets the value for CitationUrl to be an explicit nil

### UnsetCitationUrl
`func (o *CompanyBiomarkerCapabilityRow) UnsetCitationUrl()`

UnsetCitationUrl ensures that no value is present for CitationUrl, not even an explicit nil
### GetEvidenceSnippet

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceSnippet() string`

GetEvidenceSnippet returns the EvidenceSnippet field if non-nil, zero value otherwise.

### GetEvidenceSnippetOk

`func (o *CompanyBiomarkerCapabilityRow) GetEvidenceSnippetOk() (*string, bool)`

GetEvidenceSnippetOk returns a tuple with the EvidenceSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSnippet

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceSnippet(v string)`

SetEvidenceSnippet sets EvidenceSnippet field to given value.

### HasEvidenceSnippet

`func (o *CompanyBiomarkerCapabilityRow) HasEvidenceSnippet() bool`

HasEvidenceSnippet returns a boolean if a field has been set.

### SetEvidenceSnippetNil

`func (o *CompanyBiomarkerCapabilityRow) SetEvidenceSnippetNil(b bool)`

 SetEvidenceSnippetNil sets the value for EvidenceSnippet to be an explicit nil

### UnsetEvidenceSnippet
`func (o *CompanyBiomarkerCapabilityRow) UnsetEvidenceSnippet()`

UnsetEvidenceSnippet ensures that no value is present for EvidenceSnippet, not even an explicit nil
### GetBiomarkerRole

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerRole() string`

GetBiomarkerRole returns the BiomarkerRole field if non-nil, zero value otherwise.

### GetBiomarkerRoleOk

`func (o *CompanyBiomarkerCapabilityRow) GetBiomarkerRoleOk() (*string, bool)`

GetBiomarkerRoleOk returns a tuple with the BiomarkerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerRole

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerRole(v string)`

SetBiomarkerRole sets BiomarkerRole field to given value.

### HasBiomarkerRole

`func (o *CompanyBiomarkerCapabilityRow) HasBiomarkerRole() bool`

HasBiomarkerRole returns a boolean if a field has been set.

### SetBiomarkerRoleNil

`func (o *CompanyBiomarkerCapabilityRow) SetBiomarkerRoleNil(b bool)`

 SetBiomarkerRoleNil sets the value for BiomarkerRole to be an explicit nil

### UnsetBiomarkerRole
`func (o *CompanyBiomarkerCapabilityRow) UnsetBiomarkerRole()`

UnsetBiomarkerRole ensures that no value is present for BiomarkerRole, not even an explicit nil
### GetIsSurrogateEndpointClaim

`func (o *CompanyBiomarkerCapabilityRow) GetIsSurrogateEndpointClaim() bool`

GetIsSurrogateEndpointClaim returns the IsSurrogateEndpointClaim field if non-nil, zero value otherwise.

### GetIsSurrogateEndpointClaimOk

`func (o *CompanyBiomarkerCapabilityRow) GetIsSurrogateEndpointClaimOk() (*bool, bool)`

GetIsSurrogateEndpointClaimOk returns a tuple with the IsSurrogateEndpointClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSurrogateEndpointClaim

`func (o *CompanyBiomarkerCapabilityRow) SetIsSurrogateEndpointClaim(v bool)`

SetIsSurrogateEndpointClaim sets IsSurrogateEndpointClaim field to given value.

### HasIsSurrogateEndpointClaim

`func (o *CompanyBiomarkerCapabilityRow) HasIsSurrogateEndpointClaim() bool`

HasIsSurrogateEndpointClaim returns a boolean if a field has been set.

### SetIsSurrogateEndpointClaimNil

`func (o *CompanyBiomarkerCapabilityRow) SetIsSurrogateEndpointClaimNil(b bool)`

 SetIsSurrogateEndpointClaimNil sets the value for IsSurrogateEndpointClaim to be an explicit nil

### UnsetIsSurrogateEndpointClaim
`func (o *CompanyBiomarkerCapabilityRow) UnsetIsSurrogateEndpointClaim()`

UnsetIsSurrogateEndpointClaim ensures that no value is present for IsSurrogateEndpointClaim, not even an explicit nil
### GetClinvarScvCount

`func (o *CompanyBiomarkerCapabilityRow) GetClinvarScvCount() int32`

GetClinvarScvCount returns the ClinvarScvCount field if non-nil, zero value otherwise.

### GetClinvarScvCountOk

`func (o *CompanyBiomarkerCapabilityRow) GetClinvarScvCountOk() (*int32, bool)`

GetClinvarScvCountOk returns a tuple with the ClinvarScvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinvarScvCount

`func (o *CompanyBiomarkerCapabilityRow) SetClinvarScvCount(v int32)`

SetClinvarScvCount sets ClinvarScvCount field to given value.

### HasClinvarScvCount

`func (o *CompanyBiomarkerCapabilityRow) HasClinvarScvCount() bool`

HasClinvarScvCount returns a boolean if a field has been set.

### SetClinvarScvCountNil

`func (o *CompanyBiomarkerCapabilityRow) SetClinvarScvCountNil(b bool)`

 SetClinvarScvCountNil sets the value for ClinvarScvCount to be an explicit nil

### UnsetClinvarScvCount
`func (o *CompanyBiomarkerCapabilityRow) UnsetClinvarScvCount()`

UnsetClinvarScvCount ensures that no value is present for ClinvarScvCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


