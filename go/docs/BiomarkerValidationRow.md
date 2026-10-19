# BiomarkerValidationRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiomarkerValidationId** | Pointer to **NullableString** | The evidence assertion id. | [optional] 
**RouteKey** | Pointer to **NullableString** | The biomarker&#39;s route key, the entry key for the biomarkers object type; for a gene it is the gene symbol. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol (HGNC, the HUGO Gene Nomenclature Committee). | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The biomarker key. | [optional] 
**BiomarkerDisplayName** | Pointer to **NullableString** | The biomarker&#39;s display name. | [optional] 
**BiomarkerKind** | Pointer to **NullableString** | The kind of biomarker the evidence is filed against. | [optional] 
**OrganizationName** | Pointer to **NullableString** | The organization that can supply the evidence, as the source names it (null when the source names none that resolved). | [optional] 
**OrganizationOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record (null when unresolved). | [optional] 
**OrganizationCompanyId** | Pointer to **NullableString** | Company level: the organization&#39;s company (null when unresolved). | [optional] 
**OrganizationCompanyName** | Pointer to **NullableString** | Display name of the organization&#39;s company. | [optional] 
**OrganizationGroupId** | Pointer to **NullableString** | Group level: the organization&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**OrganizationGroupName** | Pointer to **NullableString** | Display name of the organization&#39;s group. | [optional] 
**OrganizationRole** | Pointer to **NullableString** | What the organization is to the evidence: lead_sponsor, collaborator, contract_vendor, assay_manufacturer, qualification_submitter, testing_lab or publishing_org. | [optional] 
**OrganizationMatchTier** | Pointer to **NullableString** | How the organization name matched an organization record. | [optional] 
**Tier** | Pointer to **NullableInt32** | The rung on the validation ladder, 1 to 4. A rung is a kind of evidence, never a score: rows are not ranked across organizations or evidence types. | [optional] 
**TierLabel** | Pointer to **NullableString** | The rung&#39;s name, for example fda_authorized_companion_diagnostic or named_in_inclusion_criteria. | [optional] 
**EvidenceType** | Pointer to **NullableString** | The evidence lane: fda_bqp_qualification, fda_cdx_authorization, aact_endpoint_role, aact_eligibility_section, aact_biospecimen, trial_mention or clinvar_clinical_testing. | [optional] 
**EvidenceCount** | Pointer to **NullableInt32** | Sibling rows sharing this biomarker, organization and role. Not the number of source records behind a ClinVar row (clinvar_scv_count). | [optional] 
**EvidenceRef** | Pointer to **NullableString** | The source reference the row cites. | [optional] 
**EvidenceRefKind** | Pointer to **NullableString** | What kind of reference evidence_ref is. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial key (an NCT id) when the evidence is a trial; null otherwise. | [optional] 
**CitationUrl** | Pointer to **NullableString** | A link to the cited record. | [optional] 
**CitationVerified** | Pointer to **NullableBool** | False for the FDA Biomarker Qualification Program arm, whose captured links are known to go stale; true otherwise. A link marked false may not resolve. | [optional] 
**EvidenceSnippet** | Pointer to **NullableString** | The text span the evidence rests on, where the source gives one. | [optional] 
**BiomarkerRole** | Pointer to **NullableString** | The biomarker&#39;s role in the evidence, where the source states it. | [optional] 
**IsSurrogateEndpointClaim** | Pointer to **NullableBool** | True when the evidence is a claim that the biomarker is a surrogate endpoint. | [optional] 
**ClinvarScvCount** | Pointer to **NullableInt32** | For a ClinVar row, the number of ClinVar submission records the row collapses; null otherwise. | [optional] 

## Methods

### NewBiomarkerValidationRow

`func NewBiomarkerValidationRow() *BiomarkerValidationRow`

NewBiomarkerValidationRow instantiates a new BiomarkerValidationRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerValidationRowWithDefaults

`func NewBiomarkerValidationRowWithDefaults() *BiomarkerValidationRow`

NewBiomarkerValidationRowWithDefaults instantiates a new BiomarkerValidationRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiomarkerValidationId

`func (o *BiomarkerValidationRow) GetBiomarkerValidationId() string`

GetBiomarkerValidationId returns the BiomarkerValidationId field if non-nil, zero value otherwise.

### GetBiomarkerValidationIdOk

`func (o *BiomarkerValidationRow) GetBiomarkerValidationIdOk() (*string, bool)`

GetBiomarkerValidationIdOk returns a tuple with the BiomarkerValidationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerValidationId

`func (o *BiomarkerValidationRow) SetBiomarkerValidationId(v string)`

SetBiomarkerValidationId sets BiomarkerValidationId field to given value.

### HasBiomarkerValidationId

`func (o *BiomarkerValidationRow) HasBiomarkerValidationId() bool`

HasBiomarkerValidationId returns a boolean if a field has been set.

### SetBiomarkerValidationIdNil

`func (o *BiomarkerValidationRow) SetBiomarkerValidationIdNil(b bool)`

 SetBiomarkerValidationIdNil sets the value for BiomarkerValidationId to be an explicit nil

### UnsetBiomarkerValidationId
`func (o *BiomarkerValidationRow) UnsetBiomarkerValidationId()`

UnsetBiomarkerValidationId ensures that no value is present for BiomarkerValidationId, not even an explicit nil
### GetRouteKey

`func (o *BiomarkerValidationRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerValidationRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerValidationRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerValidationRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerValidationRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerValidationRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetGeneSymbol

`func (o *BiomarkerValidationRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkerValidationRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkerValidationRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *BiomarkerValidationRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *BiomarkerValidationRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkerValidationRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetBiomarkerId

`func (o *BiomarkerValidationRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkerValidationRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkerValidationRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkerValidationRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkerValidationRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkerValidationRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetBiomarkerDisplayName

`func (o *BiomarkerValidationRow) GetBiomarkerDisplayName() string`

GetBiomarkerDisplayName returns the BiomarkerDisplayName field if non-nil, zero value otherwise.

### GetBiomarkerDisplayNameOk

`func (o *BiomarkerValidationRow) GetBiomarkerDisplayNameOk() (*string, bool)`

GetBiomarkerDisplayNameOk returns a tuple with the BiomarkerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerDisplayName

`func (o *BiomarkerValidationRow) SetBiomarkerDisplayName(v string)`

SetBiomarkerDisplayName sets BiomarkerDisplayName field to given value.

### HasBiomarkerDisplayName

`func (o *BiomarkerValidationRow) HasBiomarkerDisplayName() bool`

HasBiomarkerDisplayName returns a boolean if a field has been set.

### SetBiomarkerDisplayNameNil

`func (o *BiomarkerValidationRow) SetBiomarkerDisplayNameNil(b bool)`

 SetBiomarkerDisplayNameNil sets the value for BiomarkerDisplayName to be an explicit nil

### UnsetBiomarkerDisplayName
`func (o *BiomarkerValidationRow) UnsetBiomarkerDisplayName()`

UnsetBiomarkerDisplayName ensures that no value is present for BiomarkerDisplayName, not even an explicit nil
### GetBiomarkerKind

`func (o *BiomarkerValidationRow) GetBiomarkerKind() string`

GetBiomarkerKind returns the BiomarkerKind field if non-nil, zero value otherwise.

### GetBiomarkerKindOk

`func (o *BiomarkerValidationRow) GetBiomarkerKindOk() (*string, bool)`

GetBiomarkerKindOk returns a tuple with the BiomarkerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerKind

`func (o *BiomarkerValidationRow) SetBiomarkerKind(v string)`

SetBiomarkerKind sets BiomarkerKind field to given value.

### HasBiomarkerKind

`func (o *BiomarkerValidationRow) HasBiomarkerKind() bool`

HasBiomarkerKind returns a boolean if a field has been set.

### SetBiomarkerKindNil

`func (o *BiomarkerValidationRow) SetBiomarkerKindNil(b bool)`

 SetBiomarkerKindNil sets the value for BiomarkerKind to be an explicit nil

### UnsetBiomarkerKind
`func (o *BiomarkerValidationRow) UnsetBiomarkerKind()`

UnsetBiomarkerKind ensures that no value is present for BiomarkerKind, not even an explicit nil
### GetOrganizationName

`func (o *BiomarkerValidationRow) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *BiomarkerValidationRow) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *BiomarkerValidationRow) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *BiomarkerValidationRow) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *BiomarkerValidationRow) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *BiomarkerValidationRow) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationOrgId

`func (o *BiomarkerValidationRow) GetOrganizationOrgId() string`

GetOrganizationOrgId returns the OrganizationOrgId field if non-nil, zero value otherwise.

### GetOrganizationOrgIdOk

`func (o *BiomarkerValidationRow) GetOrganizationOrgIdOk() (*string, bool)`

GetOrganizationOrgIdOk returns a tuple with the OrganizationOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationOrgId

`func (o *BiomarkerValidationRow) SetOrganizationOrgId(v string)`

SetOrganizationOrgId sets OrganizationOrgId field to given value.

### HasOrganizationOrgId

`func (o *BiomarkerValidationRow) HasOrganizationOrgId() bool`

HasOrganizationOrgId returns a boolean if a field has been set.

### SetOrganizationOrgIdNil

`func (o *BiomarkerValidationRow) SetOrganizationOrgIdNil(b bool)`

 SetOrganizationOrgIdNil sets the value for OrganizationOrgId to be an explicit nil

### UnsetOrganizationOrgId
`func (o *BiomarkerValidationRow) UnsetOrganizationOrgId()`

UnsetOrganizationOrgId ensures that no value is present for OrganizationOrgId, not even an explicit nil
### GetOrganizationCompanyId

`func (o *BiomarkerValidationRow) GetOrganizationCompanyId() string`

GetOrganizationCompanyId returns the OrganizationCompanyId field if non-nil, zero value otherwise.

### GetOrganizationCompanyIdOk

`func (o *BiomarkerValidationRow) GetOrganizationCompanyIdOk() (*string, bool)`

GetOrganizationCompanyIdOk returns a tuple with the OrganizationCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyId

`func (o *BiomarkerValidationRow) SetOrganizationCompanyId(v string)`

SetOrganizationCompanyId sets OrganizationCompanyId field to given value.

### HasOrganizationCompanyId

`func (o *BiomarkerValidationRow) HasOrganizationCompanyId() bool`

HasOrganizationCompanyId returns a boolean if a field has been set.

### SetOrganizationCompanyIdNil

`func (o *BiomarkerValidationRow) SetOrganizationCompanyIdNil(b bool)`

 SetOrganizationCompanyIdNil sets the value for OrganizationCompanyId to be an explicit nil

### UnsetOrganizationCompanyId
`func (o *BiomarkerValidationRow) UnsetOrganizationCompanyId()`

UnsetOrganizationCompanyId ensures that no value is present for OrganizationCompanyId, not even an explicit nil
### GetOrganizationCompanyName

`func (o *BiomarkerValidationRow) GetOrganizationCompanyName() string`

GetOrganizationCompanyName returns the OrganizationCompanyName field if non-nil, zero value otherwise.

### GetOrganizationCompanyNameOk

`func (o *BiomarkerValidationRow) GetOrganizationCompanyNameOk() (*string, bool)`

GetOrganizationCompanyNameOk returns a tuple with the OrganizationCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCompanyName

`func (o *BiomarkerValidationRow) SetOrganizationCompanyName(v string)`

SetOrganizationCompanyName sets OrganizationCompanyName field to given value.

### HasOrganizationCompanyName

`func (o *BiomarkerValidationRow) HasOrganizationCompanyName() bool`

HasOrganizationCompanyName returns a boolean if a field has been set.

### SetOrganizationCompanyNameNil

`func (o *BiomarkerValidationRow) SetOrganizationCompanyNameNil(b bool)`

 SetOrganizationCompanyNameNil sets the value for OrganizationCompanyName to be an explicit nil

### UnsetOrganizationCompanyName
`func (o *BiomarkerValidationRow) UnsetOrganizationCompanyName()`

UnsetOrganizationCompanyName ensures that no value is present for OrganizationCompanyName, not even an explicit nil
### GetOrganizationGroupId

`func (o *BiomarkerValidationRow) GetOrganizationGroupId() string`

GetOrganizationGroupId returns the OrganizationGroupId field if non-nil, zero value otherwise.

### GetOrganizationGroupIdOk

`func (o *BiomarkerValidationRow) GetOrganizationGroupIdOk() (*string, bool)`

GetOrganizationGroupIdOk returns a tuple with the OrganizationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupId

`func (o *BiomarkerValidationRow) SetOrganizationGroupId(v string)`

SetOrganizationGroupId sets OrganizationGroupId field to given value.

### HasOrganizationGroupId

`func (o *BiomarkerValidationRow) HasOrganizationGroupId() bool`

HasOrganizationGroupId returns a boolean if a field has been set.

### SetOrganizationGroupIdNil

`func (o *BiomarkerValidationRow) SetOrganizationGroupIdNil(b bool)`

 SetOrganizationGroupIdNil sets the value for OrganizationGroupId to be an explicit nil

### UnsetOrganizationGroupId
`func (o *BiomarkerValidationRow) UnsetOrganizationGroupId()`

UnsetOrganizationGroupId ensures that no value is present for OrganizationGroupId, not even an explicit nil
### GetOrganizationGroupName

`func (o *BiomarkerValidationRow) GetOrganizationGroupName() string`

GetOrganizationGroupName returns the OrganizationGroupName field if non-nil, zero value otherwise.

### GetOrganizationGroupNameOk

`func (o *BiomarkerValidationRow) GetOrganizationGroupNameOk() (*string, bool)`

GetOrganizationGroupNameOk returns a tuple with the OrganizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGroupName

`func (o *BiomarkerValidationRow) SetOrganizationGroupName(v string)`

SetOrganizationGroupName sets OrganizationGroupName field to given value.

### HasOrganizationGroupName

`func (o *BiomarkerValidationRow) HasOrganizationGroupName() bool`

HasOrganizationGroupName returns a boolean if a field has been set.

### SetOrganizationGroupNameNil

`func (o *BiomarkerValidationRow) SetOrganizationGroupNameNil(b bool)`

 SetOrganizationGroupNameNil sets the value for OrganizationGroupName to be an explicit nil

### UnsetOrganizationGroupName
`func (o *BiomarkerValidationRow) UnsetOrganizationGroupName()`

UnsetOrganizationGroupName ensures that no value is present for OrganizationGroupName, not even an explicit nil
### GetOrganizationRole

`func (o *BiomarkerValidationRow) GetOrganizationRole() string`

GetOrganizationRole returns the OrganizationRole field if non-nil, zero value otherwise.

### GetOrganizationRoleOk

`func (o *BiomarkerValidationRow) GetOrganizationRoleOk() (*string, bool)`

GetOrganizationRoleOk returns a tuple with the OrganizationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRole

`func (o *BiomarkerValidationRow) SetOrganizationRole(v string)`

SetOrganizationRole sets OrganizationRole field to given value.

### HasOrganizationRole

`func (o *BiomarkerValidationRow) HasOrganizationRole() bool`

HasOrganizationRole returns a boolean if a field has been set.

### SetOrganizationRoleNil

`func (o *BiomarkerValidationRow) SetOrganizationRoleNil(b bool)`

 SetOrganizationRoleNil sets the value for OrganizationRole to be an explicit nil

### UnsetOrganizationRole
`func (o *BiomarkerValidationRow) UnsetOrganizationRole()`

UnsetOrganizationRole ensures that no value is present for OrganizationRole, not even an explicit nil
### GetOrganizationMatchTier

`func (o *BiomarkerValidationRow) GetOrganizationMatchTier() string`

GetOrganizationMatchTier returns the OrganizationMatchTier field if non-nil, zero value otherwise.

### GetOrganizationMatchTierOk

`func (o *BiomarkerValidationRow) GetOrganizationMatchTierOk() (*string, bool)`

GetOrganizationMatchTierOk returns a tuple with the OrganizationMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMatchTier

`func (o *BiomarkerValidationRow) SetOrganizationMatchTier(v string)`

SetOrganizationMatchTier sets OrganizationMatchTier field to given value.

### HasOrganizationMatchTier

`func (o *BiomarkerValidationRow) HasOrganizationMatchTier() bool`

HasOrganizationMatchTier returns a boolean if a field has been set.

### SetOrganizationMatchTierNil

`func (o *BiomarkerValidationRow) SetOrganizationMatchTierNil(b bool)`

 SetOrganizationMatchTierNil sets the value for OrganizationMatchTier to be an explicit nil

### UnsetOrganizationMatchTier
`func (o *BiomarkerValidationRow) UnsetOrganizationMatchTier()`

UnsetOrganizationMatchTier ensures that no value is present for OrganizationMatchTier, not even an explicit nil
### GetTier

`func (o *BiomarkerValidationRow) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *BiomarkerValidationRow) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *BiomarkerValidationRow) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *BiomarkerValidationRow) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *BiomarkerValidationRow) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *BiomarkerValidationRow) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetTierLabel

`func (o *BiomarkerValidationRow) GetTierLabel() string`

GetTierLabel returns the TierLabel field if non-nil, zero value otherwise.

### GetTierLabelOk

`func (o *BiomarkerValidationRow) GetTierLabelOk() (*string, bool)`

GetTierLabelOk returns a tuple with the TierLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierLabel

`func (o *BiomarkerValidationRow) SetTierLabel(v string)`

SetTierLabel sets TierLabel field to given value.

### HasTierLabel

`func (o *BiomarkerValidationRow) HasTierLabel() bool`

HasTierLabel returns a boolean if a field has been set.

### SetTierLabelNil

`func (o *BiomarkerValidationRow) SetTierLabelNil(b bool)`

 SetTierLabelNil sets the value for TierLabel to be an explicit nil

### UnsetTierLabel
`func (o *BiomarkerValidationRow) UnsetTierLabel()`

UnsetTierLabel ensures that no value is present for TierLabel, not even an explicit nil
### GetEvidenceType

`func (o *BiomarkerValidationRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *BiomarkerValidationRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *BiomarkerValidationRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *BiomarkerValidationRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *BiomarkerValidationRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *BiomarkerValidationRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetEvidenceCount

`func (o *BiomarkerValidationRow) GetEvidenceCount() int32`

GetEvidenceCount returns the EvidenceCount field if non-nil, zero value otherwise.

### GetEvidenceCountOk

`func (o *BiomarkerValidationRow) GetEvidenceCountOk() (*int32, bool)`

GetEvidenceCountOk returns a tuple with the EvidenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCount

`func (o *BiomarkerValidationRow) SetEvidenceCount(v int32)`

SetEvidenceCount sets EvidenceCount field to given value.

### HasEvidenceCount

`func (o *BiomarkerValidationRow) HasEvidenceCount() bool`

HasEvidenceCount returns a boolean if a field has been set.

### SetEvidenceCountNil

`func (o *BiomarkerValidationRow) SetEvidenceCountNil(b bool)`

 SetEvidenceCountNil sets the value for EvidenceCount to be an explicit nil

### UnsetEvidenceCount
`func (o *BiomarkerValidationRow) UnsetEvidenceCount()`

UnsetEvidenceCount ensures that no value is present for EvidenceCount, not even an explicit nil
### GetEvidenceRef

`func (o *BiomarkerValidationRow) GetEvidenceRef() string`

GetEvidenceRef returns the EvidenceRef field if non-nil, zero value otherwise.

### GetEvidenceRefOk

`func (o *BiomarkerValidationRow) GetEvidenceRefOk() (*string, bool)`

GetEvidenceRefOk returns a tuple with the EvidenceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRef

`func (o *BiomarkerValidationRow) SetEvidenceRef(v string)`

SetEvidenceRef sets EvidenceRef field to given value.

### HasEvidenceRef

`func (o *BiomarkerValidationRow) HasEvidenceRef() bool`

HasEvidenceRef returns a boolean if a field has been set.

### SetEvidenceRefNil

`func (o *BiomarkerValidationRow) SetEvidenceRefNil(b bool)`

 SetEvidenceRefNil sets the value for EvidenceRef to be an explicit nil

### UnsetEvidenceRef
`func (o *BiomarkerValidationRow) UnsetEvidenceRef()`

UnsetEvidenceRef ensures that no value is present for EvidenceRef, not even an explicit nil
### GetEvidenceRefKind

`func (o *BiomarkerValidationRow) GetEvidenceRefKind() string`

GetEvidenceRefKind returns the EvidenceRefKind field if non-nil, zero value otherwise.

### GetEvidenceRefKindOk

`func (o *BiomarkerValidationRow) GetEvidenceRefKindOk() (*string, bool)`

GetEvidenceRefKindOk returns a tuple with the EvidenceRefKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRefKind

`func (o *BiomarkerValidationRow) SetEvidenceRefKind(v string)`

SetEvidenceRefKind sets EvidenceRefKind field to given value.

### HasEvidenceRefKind

`func (o *BiomarkerValidationRow) HasEvidenceRefKind() bool`

HasEvidenceRefKind returns a boolean if a field has been set.

### SetEvidenceRefKindNil

`func (o *BiomarkerValidationRow) SetEvidenceRefKindNil(b bool)`

 SetEvidenceRefKindNil sets the value for EvidenceRefKind to be an explicit nil

### UnsetEvidenceRefKind
`func (o *BiomarkerValidationRow) UnsetEvidenceRefKind()`

UnsetEvidenceRefKind ensures that no value is present for EvidenceRefKind, not even an explicit nil
### GetTrialId

`func (o *BiomarkerValidationRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *BiomarkerValidationRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *BiomarkerValidationRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *BiomarkerValidationRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *BiomarkerValidationRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *BiomarkerValidationRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetCitationUrl

`func (o *BiomarkerValidationRow) GetCitationUrl() string`

GetCitationUrl returns the CitationUrl field if non-nil, zero value otherwise.

### GetCitationUrlOk

`func (o *BiomarkerValidationRow) GetCitationUrlOk() (*string, bool)`

GetCitationUrlOk returns a tuple with the CitationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationUrl

`func (o *BiomarkerValidationRow) SetCitationUrl(v string)`

SetCitationUrl sets CitationUrl field to given value.

### HasCitationUrl

`func (o *BiomarkerValidationRow) HasCitationUrl() bool`

HasCitationUrl returns a boolean if a field has been set.

### SetCitationUrlNil

`func (o *BiomarkerValidationRow) SetCitationUrlNil(b bool)`

 SetCitationUrlNil sets the value for CitationUrl to be an explicit nil

### UnsetCitationUrl
`func (o *BiomarkerValidationRow) UnsetCitationUrl()`

UnsetCitationUrl ensures that no value is present for CitationUrl, not even an explicit nil
### GetCitationVerified

`func (o *BiomarkerValidationRow) GetCitationVerified() bool`

GetCitationVerified returns the CitationVerified field if non-nil, zero value otherwise.

### GetCitationVerifiedOk

`func (o *BiomarkerValidationRow) GetCitationVerifiedOk() (*bool, bool)`

GetCitationVerifiedOk returns a tuple with the CitationVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationVerified

`func (o *BiomarkerValidationRow) SetCitationVerified(v bool)`

SetCitationVerified sets CitationVerified field to given value.

### HasCitationVerified

`func (o *BiomarkerValidationRow) HasCitationVerified() bool`

HasCitationVerified returns a boolean if a field has been set.

### SetCitationVerifiedNil

`func (o *BiomarkerValidationRow) SetCitationVerifiedNil(b bool)`

 SetCitationVerifiedNil sets the value for CitationVerified to be an explicit nil

### UnsetCitationVerified
`func (o *BiomarkerValidationRow) UnsetCitationVerified()`

UnsetCitationVerified ensures that no value is present for CitationVerified, not even an explicit nil
### GetEvidenceSnippet

`func (o *BiomarkerValidationRow) GetEvidenceSnippet() string`

GetEvidenceSnippet returns the EvidenceSnippet field if non-nil, zero value otherwise.

### GetEvidenceSnippetOk

`func (o *BiomarkerValidationRow) GetEvidenceSnippetOk() (*string, bool)`

GetEvidenceSnippetOk returns a tuple with the EvidenceSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSnippet

`func (o *BiomarkerValidationRow) SetEvidenceSnippet(v string)`

SetEvidenceSnippet sets EvidenceSnippet field to given value.

### HasEvidenceSnippet

`func (o *BiomarkerValidationRow) HasEvidenceSnippet() bool`

HasEvidenceSnippet returns a boolean if a field has been set.

### SetEvidenceSnippetNil

`func (o *BiomarkerValidationRow) SetEvidenceSnippetNil(b bool)`

 SetEvidenceSnippetNil sets the value for EvidenceSnippet to be an explicit nil

### UnsetEvidenceSnippet
`func (o *BiomarkerValidationRow) UnsetEvidenceSnippet()`

UnsetEvidenceSnippet ensures that no value is present for EvidenceSnippet, not even an explicit nil
### GetBiomarkerRole

`func (o *BiomarkerValidationRow) GetBiomarkerRole() string`

GetBiomarkerRole returns the BiomarkerRole field if non-nil, zero value otherwise.

### GetBiomarkerRoleOk

`func (o *BiomarkerValidationRow) GetBiomarkerRoleOk() (*string, bool)`

GetBiomarkerRoleOk returns a tuple with the BiomarkerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerRole

`func (o *BiomarkerValidationRow) SetBiomarkerRole(v string)`

SetBiomarkerRole sets BiomarkerRole field to given value.

### HasBiomarkerRole

`func (o *BiomarkerValidationRow) HasBiomarkerRole() bool`

HasBiomarkerRole returns a boolean if a field has been set.

### SetBiomarkerRoleNil

`func (o *BiomarkerValidationRow) SetBiomarkerRoleNil(b bool)`

 SetBiomarkerRoleNil sets the value for BiomarkerRole to be an explicit nil

### UnsetBiomarkerRole
`func (o *BiomarkerValidationRow) UnsetBiomarkerRole()`

UnsetBiomarkerRole ensures that no value is present for BiomarkerRole, not even an explicit nil
### GetIsSurrogateEndpointClaim

`func (o *BiomarkerValidationRow) GetIsSurrogateEndpointClaim() bool`

GetIsSurrogateEndpointClaim returns the IsSurrogateEndpointClaim field if non-nil, zero value otherwise.

### GetIsSurrogateEndpointClaimOk

`func (o *BiomarkerValidationRow) GetIsSurrogateEndpointClaimOk() (*bool, bool)`

GetIsSurrogateEndpointClaimOk returns a tuple with the IsSurrogateEndpointClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSurrogateEndpointClaim

`func (o *BiomarkerValidationRow) SetIsSurrogateEndpointClaim(v bool)`

SetIsSurrogateEndpointClaim sets IsSurrogateEndpointClaim field to given value.

### HasIsSurrogateEndpointClaim

`func (o *BiomarkerValidationRow) HasIsSurrogateEndpointClaim() bool`

HasIsSurrogateEndpointClaim returns a boolean if a field has been set.

### SetIsSurrogateEndpointClaimNil

`func (o *BiomarkerValidationRow) SetIsSurrogateEndpointClaimNil(b bool)`

 SetIsSurrogateEndpointClaimNil sets the value for IsSurrogateEndpointClaim to be an explicit nil

### UnsetIsSurrogateEndpointClaim
`func (o *BiomarkerValidationRow) UnsetIsSurrogateEndpointClaim()`

UnsetIsSurrogateEndpointClaim ensures that no value is present for IsSurrogateEndpointClaim, not even an explicit nil
### GetClinvarScvCount

`func (o *BiomarkerValidationRow) GetClinvarScvCount() int32`

GetClinvarScvCount returns the ClinvarScvCount field if non-nil, zero value otherwise.

### GetClinvarScvCountOk

`func (o *BiomarkerValidationRow) GetClinvarScvCountOk() (*int32, bool)`

GetClinvarScvCountOk returns a tuple with the ClinvarScvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinvarScvCount

`func (o *BiomarkerValidationRow) SetClinvarScvCount(v int32)`

SetClinvarScvCount sets ClinvarScvCount field to given value.

### HasClinvarScvCount

`func (o *BiomarkerValidationRow) HasClinvarScvCount() bool`

HasClinvarScvCount returns a boolean if a field has been set.

### SetClinvarScvCountNil

`func (o *BiomarkerValidationRow) SetClinvarScvCountNil(b bool)`

 SetClinvarScvCountNil sets the value for ClinvarScvCount to be an explicit nil

### UnsetClinvarScvCount
`func (o *BiomarkerValidationRow) UnsetClinvarScvCount()`

UnsetClinvarScvCount ensures that no value is present for ClinvarScvCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


