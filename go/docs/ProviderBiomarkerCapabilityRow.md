# ProviderBiomarkerCapabilityRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderBiomarkerCapabilityId** | Pointer to **NullableString** | The evidence assertion id: a hash of the supplier, biomarker, lane, role, evidence reference kind, evidence reference and ladder rule. | [optional] 
**ProviderId** | Pointer to **NullableString** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). A supplier is not a company. | [optional] 
**ProviderUuid** | Pointer to **NullableString** | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] 
**ProviderSlug** | Pointer to **NullableString** | The supplier&#39;s marketplace slug. | [optional] 
**ProviderName** | Pointer to **NullableString** | The supplier&#39;s name. | [optional] 
**RouteKey** | Pointer to **NullableString** | The biomarker&#39;s route key, the entry key for the biomarkers object type; the gene symbol for a gene biomarker. | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The canonical biomarker entity id. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol; null for a source-scoped biomarker that names no gene. | [optional] 
**BiomarkerDisplayName** | Pointer to **NullableString** | The biomarker&#39;s display name. | [optional] 
**EvidenceLane** | Pointer to **NullableString** | The source path of the row: ladder_trial and ladder_clinvar (the evidence ladder), fda_cdx (FDA companion diagnostic list), fda_bqp (FDA Biomarker Qualification Program), literature_assay (assay providers named in PubMed full text), supplier_survey (the supplier&#39;s own declaration) clinvar_submission (ClinVar clinical-testing submissions by a laboratory the supplier carries by name, one row per ClinVar record, tier null) and supplier_service_listing (a Scientist.com marketplace catalogue service the supplier listed, whose name names the gene; tier null, no evidence sense). Nothing is merged across lanes. | [optional] 
**EvidenceType** | Pointer to **NullableString** | The source&#39;s own evidence vocabulary: the ladder evidence_type on a ladder lane, literature_assay_provider on the literature lane, supplier_survey_declaration on the survey lane, rx_service_listing on the service listing lane. | [optional] 
**OrgRole** | Pointer to **NullableString** | The supplier&#39;s role on the evidence: the ladder role on a ladder lane, the ledger provider role on the literature lane, respondent on the survey lane, or marketplace_listing on the service listing lane. | [optional] 
**AnalyticalValidationEvidence** | Pointer to **NullableString** | Sense 1, literature_assay lane only: the source&#39;s own label for how the assay is evidenced, never a number. literature_reported is a paper naming the assay, never validation. | [optional] 
**ClinicalUseEvidence** | Pointer to **NullableString** | Sense 2: the ladder rung&#39;s name on a ladder lane, literature_reported on the literature lane (a literature mention, never validation), and ClinVar&#39;s own review status verbatim on the clinvar_submission lane. | [optional] 
**RegulatoryQualification** | Pointer to **NullableString** | Sense 3, fda_bqp lane only: fda_qualified, the FDA Biomarker Qualification Program qualified the biomarker. | [optional] 
**CompanionDiagnostic** | Pointer to **NullableString** | Sense 4, fda_cdx lane only: fda_authorized_companion_diagnostic, the FDA companion diagnostic list names the supplier&#39;s device for the biomarker. | [optional] 
**SupplierDeclaredStatus** | Pointer to **NullableString** | The supplier&#39;s own declaration, never mapped into the four evidence kinds: the survey&#39;s assay status verbatim on the supplier_survey lane, supplier_declared_service_listed on the supplier_service_listing lane (the supplier listed a catalogue service naming the gene). | [optional] 
**Tier** | Pointer to **NullableInt32** | The ladder rung of this row, 1 to 4, a kind of evidence and never a score; set on the ladder_trial, ladder_clinvar, fda_cdx and fda_bqp lanes, null on literature_assay, supplier_survey, clinvar_submission and supplier_service_listing. Not rolled up: see the example query for the maximum within the ladder lanes. | [optional] 
**TierRuleId** | Pointer to **NullableString** | The ladder rule that set the tier of this row; null where tier is null. | [optional] 
**EvidenceRef** | Pointer to **NullableString** | The row key that resolves in the model evidence_ref_kind names. | [optional] 
**EvidenceRefKind** | Pointer to **NullableString** | What evidence_ref is. | [optional] 
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number) on trial-sourced ladder rows; null otherwise. | [optional] 
**ClinvarVariationId** | Pointer to **NullableString** | The ClinVar variation id on ClinVar ladder rows; null otherwise. | [optional] 
**DeviceApplicationNumbers** | Pointer to **[]string** | The FDA device application numbers of a companion diagnostic row; null otherwise. | [optional] 
**DeviceCredentialLabels** | Pointer to **[]string** | The FDA pathway labels (PMA, De Novo, 510(k)) of a companion diagnostic row; null otherwise. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The source URL the evidence row carries, where it has one. | [optional] 
**CitationUrl** | Pointer to **NullableString** | A resolvable citation: the trial page, the ClinVar variation page, the FDA device page or list, the PubMed page, or the supplier&#39;s marketplace page on a service listing. Null where the source offers none. | [optional] 
**EvidenceSnippet** | Pointer to **NullableString** | A short evidence text (at most 500 characters), the paper title, or the catalogue service name on a service listing row, where the source has one. | [optional] 
**MostRecentEvidenceDate** | Pointer to **NullableString** | literature_assay lane only: the publication date of the most recent cited paper. | [optional] 
**MatchMethod** | Pointer to **NullableString** | How the evidence reached the supplier: org_member_exact (the organization record of a member of the supplier), rx_name_key_exact (an exact name-key match of a source name to the supplier) or rx_provider_direct (the supplier&#39;s own marketplace listing, no name matched). A name match is not a verified identity. | [optional] 
**MatchConfidence** | Pointer to **NullableFloat32** | The confidence of the upstream match rung where that rung defines one; null otherwise, never invented. | [optional] 
**CompanyNameMatched** | Pointer to **NullableString** | The source name that matched the supplier&#39;s name key, on name-matched rows. | [optional] 
**SourceLicense** | Pointer to **NullableString** | The license statement of the row&#39;s source lane. No lane carries a usage restriction. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit the row&#39;s source lane is owed. | [optional] 
**AsOf** | Pointer to **NullableString** | The build date. | [optional] 

## Methods

### NewProviderBiomarkerCapabilityRow

`func NewProviderBiomarkerCapabilityRow() *ProviderBiomarkerCapabilityRow`

NewProviderBiomarkerCapabilityRow instantiates a new ProviderBiomarkerCapabilityRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderBiomarkerCapabilityRowWithDefaults

`func NewProviderBiomarkerCapabilityRowWithDefaults() *ProviderBiomarkerCapabilityRow`

NewProviderBiomarkerCapabilityRowWithDefaults instantiates a new ProviderBiomarkerCapabilityRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderBiomarkerCapabilityId

`func (o *ProviderBiomarkerCapabilityRow) GetProviderBiomarkerCapabilityId() string`

GetProviderBiomarkerCapabilityId returns the ProviderBiomarkerCapabilityId field if non-nil, zero value otherwise.

### GetProviderBiomarkerCapabilityIdOk

`func (o *ProviderBiomarkerCapabilityRow) GetProviderBiomarkerCapabilityIdOk() (*string, bool)`

GetProviderBiomarkerCapabilityIdOk returns a tuple with the ProviderBiomarkerCapabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderBiomarkerCapabilityId

`func (o *ProviderBiomarkerCapabilityRow) SetProviderBiomarkerCapabilityId(v string)`

SetProviderBiomarkerCapabilityId sets ProviderBiomarkerCapabilityId field to given value.

### HasProviderBiomarkerCapabilityId

`func (o *ProviderBiomarkerCapabilityRow) HasProviderBiomarkerCapabilityId() bool`

HasProviderBiomarkerCapabilityId returns a boolean if a field has been set.

### SetProviderBiomarkerCapabilityIdNil

`func (o *ProviderBiomarkerCapabilityRow) SetProviderBiomarkerCapabilityIdNil(b bool)`

 SetProviderBiomarkerCapabilityIdNil sets the value for ProviderBiomarkerCapabilityId to be an explicit nil

### UnsetProviderBiomarkerCapabilityId
`func (o *ProviderBiomarkerCapabilityRow) UnsetProviderBiomarkerCapabilityId()`

UnsetProviderBiomarkerCapabilityId ensures that no value is present for ProviderBiomarkerCapabilityId, not even an explicit nil
### GetProviderId

`func (o *ProviderBiomarkerCapabilityRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ProviderBiomarkerCapabilityRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ProviderBiomarkerCapabilityRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ProviderBiomarkerCapabilityRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ProviderBiomarkerCapabilityRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ProviderBiomarkerCapabilityRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderUuid

`func (o *ProviderBiomarkerCapabilityRow) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *ProviderBiomarkerCapabilityRow) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *ProviderBiomarkerCapabilityRow) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.

### HasProviderUuid

`func (o *ProviderBiomarkerCapabilityRow) HasProviderUuid() bool`

HasProviderUuid returns a boolean if a field has been set.

### SetProviderUuidNil

`func (o *ProviderBiomarkerCapabilityRow) SetProviderUuidNil(b bool)`

 SetProviderUuidNil sets the value for ProviderUuid to be an explicit nil

### UnsetProviderUuid
`func (o *ProviderBiomarkerCapabilityRow) UnsetProviderUuid()`

UnsetProviderUuid ensures that no value is present for ProviderUuid, not even an explicit nil
### GetProviderSlug

`func (o *ProviderBiomarkerCapabilityRow) GetProviderSlug() string`

GetProviderSlug returns the ProviderSlug field if non-nil, zero value otherwise.

### GetProviderSlugOk

`func (o *ProviderBiomarkerCapabilityRow) GetProviderSlugOk() (*string, bool)`

GetProviderSlugOk returns a tuple with the ProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSlug

`func (o *ProviderBiomarkerCapabilityRow) SetProviderSlug(v string)`

SetProviderSlug sets ProviderSlug field to given value.

### HasProviderSlug

`func (o *ProviderBiomarkerCapabilityRow) HasProviderSlug() bool`

HasProviderSlug returns a boolean if a field has been set.

### SetProviderSlugNil

`func (o *ProviderBiomarkerCapabilityRow) SetProviderSlugNil(b bool)`

 SetProviderSlugNil sets the value for ProviderSlug to be an explicit nil

### UnsetProviderSlug
`func (o *ProviderBiomarkerCapabilityRow) UnsetProviderSlug()`

UnsetProviderSlug ensures that no value is present for ProviderSlug, not even an explicit nil
### GetProviderName

`func (o *ProviderBiomarkerCapabilityRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ProviderBiomarkerCapabilityRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ProviderBiomarkerCapabilityRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ProviderBiomarkerCapabilityRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ProviderBiomarkerCapabilityRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ProviderBiomarkerCapabilityRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetRouteKey

`func (o *ProviderBiomarkerCapabilityRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *ProviderBiomarkerCapabilityRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *ProviderBiomarkerCapabilityRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *ProviderBiomarkerCapabilityRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *ProviderBiomarkerCapabilityRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *ProviderBiomarkerCapabilityRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerId

`func (o *ProviderBiomarkerCapabilityRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *ProviderBiomarkerCapabilityRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *ProviderBiomarkerCapabilityRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *ProviderBiomarkerCapabilityRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *ProviderBiomarkerCapabilityRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *ProviderBiomarkerCapabilityRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetGeneSymbol

`func (o *ProviderBiomarkerCapabilityRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *ProviderBiomarkerCapabilityRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *ProviderBiomarkerCapabilityRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *ProviderBiomarkerCapabilityRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *ProviderBiomarkerCapabilityRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *ProviderBiomarkerCapabilityRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetBiomarkerDisplayName

`func (o *ProviderBiomarkerCapabilityRow) GetBiomarkerDisplayName() string`

GetBiomarkerDisplayName returns the BiomarkerDisplayName field if non-nil, zero value otherwise.

### GetBiomarkerDisplayNameOk

`func (o *ProviderBiomarkerCapabilityRow) GetBiomarkerDisplayNameOk() (*string, bool)`

GetBiomarkerDisplayNameOk returns a tuple with the BiomarkerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerDisplayName

`func (o *ProviderBiomarkerCapabilityRow) SetBiomarkerDisplayName(v string)`

SetBiomarkerDisplayName sets BiomarkerDisplayName field to given value.

### HasBiomarkerDisplayName

`func (o *ProviderBiomarkerCapabilityRow) HasBiomarkerDisplayName() bool`

HasBiomarkerDisplayName returns a boolean if a field has been set.

### SetBiomarkerDisplayNameNil

`func (o *ProviderBiomarkerCapabilityRow) SetBiomarkerDisplayNameNil(b bool)`

 SetBiomarkerDisplayNameNil sets the value for BiomarkerDisplayName to be an explicit nil

### UnsetBiomarkerDisplayName
`func (o *ProviderBiomarkerCapabilityRow) UnsetBiomarkerDisplayName()`

UnsetBiomarkerDisplayName ensures that no value is present for BiomarkerDisplayName, not even an explicit nil
### GetEvidenceLane

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceLane() string`

GetEvidenceLane returns the EvidenceLane field if non-nil, zero value otherwise.

### GetEvidenceLaneOk

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceLaneOk() (*string, bool)`

GetEvidenceLaneOk returns a tuple with the EvidenceLane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceLane

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceLane(v string)`

SetEvidenceLane sets EvidenceLane field to given value.

### HasEvidenceLane

`func (o *ProviderBiomarkerCapabilityRow) HasEvidenceLane() bool`

HasEvidenceLane returns a boolean if a field has been set.

### SetEvidenceLaneNil

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceLaneNil(b bool)`

 SetEvidenceLaneNil sets the value for EvidenceLane to be an explicit nil

### UnsetEvidenceLane
`func (o *ProviderBiomarkerCapabilityRow) UnsetEvidenceLane()`

UnsetEvidenceLane ensures that no value is present for EvidenceLane, not even an explicit nil
### GetEvidenceType

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *ProviderBiomarkerCapabilityRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *ProviderBiomarkerCapabilityRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetOrgRole

`func (o *ProviderBiomarkerCapabilityRow) GetOrgRole() string`

GetOrgRole returns the OrgRole field if non-nil, zero value otherwise.

### GetOrgRoleOk

`func (o *ProviderBiomarkerCapabilityRow) GetOrgRoleOk() (*string, bool)`

GetOrgRoleOk returns a tuple with the OrgRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRole

`func (o *ProviderBiomarkerCapabilityRow) SetOrgRole(v string)`

SetOrgRole sets OrgRole field to given value.

### HasOrgRole

`func (o *ProviderBiomarkerCapabilityRow) HasOrgRole() bool`

HasOrgRole returns a boolean if a field has been set.

### SetOrgRoleNil

`func (o *ProviderBiomarkerCapabilityRow) SetOrgRoleNil(b bool)`

 SetOrgRoleNil sets the value for OrgRole to be an explicit nil

### UnsetOrgRole
`func (o *ProviderBiomarkerCapabilityRow) UnsetOrgRole()`

UnsetOrgRole ensures that no value is present for OrgRole, not even an explicit nil
### GetAnalyticalValidationEvidence

`func (o *ProviderBiomarkerCapabilityRow) GetAnalyticalValidationEvidence() string`

GetAnalyticalValidationEvidence returns the AnalyticalValidationEvidence field if non-nil, zero value otherwise.

### GetAnalyticalValidationEvidenceOk

`func (o *ProviderBiomarkerCapabilityRow) GetAnalyticalValidationEvidenceOk() (*string, bool)`

GetAnalyticalValidationEvidenceOk returns a tuple with the AnalyticalValidationEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticalValidationEvidence

`func (o *ProviderBiomarkerCapabilityRow) SetAnalyticalValidationEvidence(v string)`

SetAnalyticalValidationEvidence sets AnalyticalValidationEvidence field to given value.

### HasAnalyticalValidationEvidence

`func (o *ProviderBiomarkerCapabilityRow) HasAnalyticalValidationEvidence() bool`

HasAnalyticalValidationEvidence returns a boolean if a field has been set.

### SetAnalyticalValidationEvidenceNil

`func (o *ProviderBiomarkerCapabilityRow) SetAnalyticalValidationEvidenceNil(b bool)`

 SetAnalyticalValidationEvidenceNil sets the value for AnalyticalValidationEvidence to be an explicit nil

### UnsetAnalyticalValidationEvidence
`func (o *ProviderBiomarkerCapabilityRow) UnsetAnalyticalValidationEvidence()`

UnsetAnalyticalValidationEvidence ensures that no value is present for AnalyticalValidationEvidence, not even an explicit nil
### GetClinicalUseEvidence

`func (o *ProviderBiomarkerCapabilityRow) GetClinicalUseEvidence() string`

GetClinicalUseEvidence returns the ClinicalUseEvidence field if non-nil, zero value otherwise.

### GetClinicalUseEvidenceOk

`func (o *ProviderBiomarkerCapabilityRow) GetClinicalUseEvidenceOk() (*string, bool)`

GetClinicalUseEvidenceOk returns a tuple with the ClinicalUseEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinicalUseEvidence

`func (o *ProviderBiomarkerCapabilityRow) SetClinicalUseEvidence(v string)`

SetClinicalUseEvidence sets ClinicalUseEvidence field to given value.

### HasClinicalUseEvidence

`func (o *ProviderBiomarkerCapabilityRow) HasClinicalUseEvidence() bool`

HasClinicalUseEvidence returns a boolean if a field has been set.

### SetClinicalUseEvidenceNil

`func (o *ProviderBiomarkerCapabilityRow) SetClinicalUseEvidenceNil(b bool)`

 SetClinicalUseEvidenceNil sets the value for ClinicalUseEvidence to be an explicit nil

### UnsetClinicalUseEvidence
`func (o *ProviderBiomarkerCapabilityRow) UnsetClinicalUseEvidence()`

UnsetClinicalUseEvidence ensures that no value is present for ClinicalUseEvidence, not even an explicit nil
### GetRegulatoryQualification

`func (o *ProviderBiomarkerCapabilityRow) GetRegulatoryQualification() string`

GetRegulatoryQualification returns the RegulatoryQualification field if non-nil, zero value otherwise.

### GetRegulatoryQualificationOk

`func (o *ProviderBiomarkerCapabilityRow) GetRegulatoryQualificationOk() (*string, bool)`

GetRegulatoryQualificationOk returns a tuple with the RegulatoryQualification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryQualification

`func (o *ProviderBiomarkerCapabilityRow) SetRegulatoryQualification(v string)`

SetRegulatoryQualification sets RegulatoryQualification field to given value.

### HasRegulatoryQualification

`func (o *ProviderBiomarkerCapabilityRow) HasRegulatoryQualification() bool`

HasRegulatoryQualification returns a boolean if a field has been set.

### SetRegulatoryQualificationNil

`func (o *ProviderBiomarkerCapabilityRow) SetRegulatoryQualificationNil(b bool)`

 SetRegulatoryQualificationNil sets the value for RegulatoryQualification to be an explicit nil

### UnsetRegulatoryQualification
`func (o *ProviderBiomarkerCapabilityRow) UnsetRegulatoryQualification()`

UnsetRegulatoryQualification ensures that no value is present for RegulatoryQualification, not even an explicit nil
### GetCompanionDiagnostic

`func (o *ProviderBiomarkerCapabilityRow) GetCompanionDiagnostic() string`

GetCompanionDiagnostic returns the CompanionDiagnostic field if non-nil, zero value otherwise.

### GetCompanionDiagnosticOk

`func (o *ProviderBiomarkerCapabilityRow) GetCompanionDiagnosticOk() (*string, bool)`

GetCompanionDiagnosticOk returns a tuple with the CompanionDiagnostic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanionDiagnostic

`func (o *ProviderBiomarkerCapabilityRow) SetCompanionDiagnostic(v string)`

SetCompanionDiagnostic sets CompanionDiagnostic field to given value.

### HasCompanionDiagnostic

`func (o *ProviderBiomarkerCapabilityRow) HasCompanionDiagnostic() bool`

HasCompanionDiagnostic returns a boolean if a field has been set.

### SetCompanionDiagnosticNil

`func (o *ProviderBiomarkerCapabilityRow) SetCompanionDiagnosticNil(b bool)`

 SetCompanionDiagnosticNil sets the value for CompanionDiagnostic to be an explicit nil

### UnsetCompanionDiagnostic
`func (o *ProviderBiomarkerCapabilityRow) UnsetCompanionDiagnostic()`

UnsetCompanionDiagnostic ensures that no value is present for CompanionDiagnostic, not even an explicit nil
### GetSupplierDeclaredStatus

`func (o *ProviderBiomarkerCapabilityRow) GetSupplierDeclaredStatus() string`

GetSupplierDeclaredStatus returns the SupplierDeclaredStatus field if non-nil, zero value otherwise.

### GetSupplierDeclaredStatusOk

`func (o *ProviderBiomarkerCapabilityRow) GetSupplierDeclaredStatusOk() (*string, bool)`

GetSupplierDeclaredStatusOk returns a tuple with the SupplierDeclaredStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierDeclaredStatus

`func (o *ProviderBiomarkerCapabilityRow) SetSupplierDeclaredStatus(v string)`

SetSupplierDeclaredStatus sets SupplierDeclaredStatus field to given value.

### HasSupplierDeclaredStatus

`func (o *ProviderBiomarkerCapabilityRow) HasSupplierDeclaredStatus() bool`

HasSupplierDeclaredStatus returns a boolean if a field has been set.

### SetSupplierDeclaredStatusNil

`func (o *ProviderBiomarkerCapabilityRow) SetSupplierDeclaredStatusNil(b bool)`

 SetSupplierDeclaredStatusNil sets the value for SupplierDeclaredStatus to be an explicit nil

### UnsetSupplierDeclaredStatus
`func (o *ProviderBiomarkerCapabilityRow) UnsetSupplierDeclaredStatus()`

UnsetSupplierDeclaredStatus ensures that no value is present for SupplierDeclaredStatus, not even an explicit nil
### GetTier

`func (o *ProviderBiomarkerCapabilityRow) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ProviderBiomarkerCapabilityRow) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ProviderBiomarkerCapabilityRow) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ProviderBiomarkerCapabilityRow) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *ProviderBiomarkerCapabilityRow) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *ProviderBiomarkerCapabilityRow) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetTierRuleId

`func (o *ProviderBiomarkerCapabilityRow) GetTierRuleId() string`

GetTierRuleId returns the TierRuleId field if non-nil, zero value otherwise.

### GetTierRuleIdOk

`func (o *ProviderBiomarkerCapabilityRow) GetTierRuleIdOk() (*string, bool)`

GetTierRuleIdOk returns a tuple with the TierRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierRuleId

`func (o *ProviderBiomarkerCapabilityRow) SetTierRuleId(v string)`

SetTierRuleId sets TierRuleId field to given value.

### HasTierRuleId

`func (o *ProviderBiomarkerCapabilityRow) HasTierRuleId() bool`

HasTierRuleId returns a boolean if a field has been set.

### SetTierRuleIdNil

`func (o *ProviderBiomarkerCapabilityRow) SetTierRuleIdNil(b bool)`

 SetTierRuleIdNil sets the value for TierRuleId to be an explicit nil

### UnsetTierRuleId
`func (o *ProviderBiomarkerCapabilityRow) UnsetTierRuleId()`

UnsetTierRuleId ensures that no value is present for TierRuleId, not even an explicit nil
### GetEvidenceRef

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceRef() string`

GetEvidenceRef returns the EvidenceRef field if non-nil, zero value otherwise.

### GetEvidenceRefOk

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceRefOk() (*string, bool)`

GetEvidenceRefOk returns a tuple with the EvidenceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRef

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceRef(v string)`

SetEvidenceRef sets EvidenceRef field to given value.

### HasEvidenceRef

`func (o *ProviderBiomarkerCapabilityRow) HasEvidenceRef() bool`

HasEvidenceRef returns a boolean if a field has been set.

### SetEvidenceRefNil

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceRefNil(b bool)`

 SetEvidenceRefNil sets the value for EvidenceRef to be an explicit nil

### UnsetEvidenceRef
`func (o *ProviderBiomarkerCapabilityRow) UnsetEvidenceRef()`

UnsetEvidenceRef ensures that no value is present for EvidenceRef, not even an explicit nil
### GetEvidenceRefKind

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceRefKind() string`

GetEvidenceRefKind returns the EvidenceRefKind field if non-nil, zero value otherwise.

### GetEvidenceRefKindOk

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceRefKindOk() (*string, bool)`

GetEvidenceRefKindOk returns a tuple with the EvidenceRefKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRefKind

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceRefKind(v string)`

SetEvidenceRefKind sets EvidenceRefKind field to given value.

### HasEvidenceRefKind

`func (o *ProviderBiomarkerCapabilityRow) HasEvidenceRefKind() bool`

HasEvidenceRefKind returns a boolean if a field has been set.

### SetEvidenceRefKindNil

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceRefKindNil(b bool)`

 SetEvidenceRefKindNil sets the value for EvidenceRefKind to be an explicit nil

### UnsetEvidenceRefKind
`func (o *ProviderBiomarkerCapabilityRow) UnsetEvidenceRefKind()`

UnsetEvidenceRefKind ensures that no value is present for EvidenceRefKind, not even an explicit nil
### GetTrialId

`func (o *ProviderBiomarkerCapabilityRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *ProviderBiomarkerCapabilityRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *ProviderBiomarkerCapabilityRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *ProviderBiomarkerCapabilityRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *ProviderBiomarkerCapabilityRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *ProviderBiomarkerCapabilityRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetClinvarVariationId

`func (o *ProviderBiomarkerCapabilityRow) GetClinvarVariationId() string`

GetClinvarVariationId returns the ClinvarVariationId field if non-nil, zero value otherwise.

### GetClinvarVariationIdOk

`func (o *ProviderBiomarkerCapabilityRow) GetClinvarVariationIdOk() (*string, bool)`

GetClinvarVariationIdOk returns a tuple with the ClinvarVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinvarVariationId

`func (o *ProviderBiomarkerCapabilityRow) SetClinvarVariationId(v string)`

SetClinvarVariationId sets ClinvarVariationId field to given value.

### HasClinvarVariationId

`func (o *ProviderBiomarkerCapabilityRow) HasClinvarVariationId() bool`

HasClinvarVariationId returns a boolean if a field has been set.

### SetClinvarVariationIdNil

`func (o *ProviderBiomarkerCapabilityRow) SetClinvarVariationIdNil(b bool)`

 SetClinvarVariationIdNil sets the value for ClinvarVariationId to be an explicit nil

### UnsetClinvarVariationId
`func (o *ProviderBiomarkerCapabilityRow) UnsetClinvarVariationId()`

UnsetClinvarVariationId ensures that no value is present for ClinvarVariationId, not even an explicit nil
### GetDeviceApplicationNumbers

`func (o *ProviderBiomarkerCapabilityRow) GetDeviceApplicationNumbers() []string`

GetDeviceApplicationNumbers returns the DeviceApplicationNumbers field if non-nil, zero value otherwise.

### GetDeviceApplicationNumbersOk

`func (o *ProviderBiomarkerCapabilityRow) GetDeviceApplicationNumbersOk() (*[]string, bool)`

GetDeviceApplicationNumbersOk returns a tuple with the DeviceApplicationNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceApplicationNumbers

`func (o *ProviderBiomarkerCapabilityRow) SetDeviceApplicationNumbers(v []string)`

SetDeviceApplicationNumbers sets DeviceApplicationNumbers field to given value.

### HasDeviceApplicationNumbers

`func (o *ProviderBiomarkerCapabilityRow) HasDeviceApplicationNumbers() bool`

HasDeviceApplicationNumbers returns a boolean if a field has been set.

### SetDeviceApplicationNumbersNil

`func (o *ProviderBiomarkerCapabilityRow) SetDeviceApplicationNumbersNil(b bool)`

 SetDeviceApplicationNumbersNil sets the value for DeviceApplicationNumbers to be an explicit nil

### UnsetDeviceApplicationNumbers
`func (o *ProviderBiomarkerCapabilityRow) UnsetDeviceApplicationNumbers()`

UnsetDeviceApplicationNumbers ensures that no value is present for DeviceApplicationNumbers, not even an explicit nil
### GetDeviceCredentialLabels

`func (o *ProviderBiomarkerCapabilityRow) GetDeviceCredentialLabels() []string`

GetDeviceCredentialLabels returns the DeviceCredentialLabels field if non-nil, zero value otherwise.

### GetDeviceCredentialLabelsOk

`func (o *ProviderBiomarkerCapabilityRow) GetDeviceCredentialLabelsOk() (*[]string, bool)`

GetDeviceCredentialLabelsOk returns a tuple with the DeviceCredentialLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCredentialLabels

`func (o *ProviderBiomarkerCapabilityRow) SetDeviceCredentialLabels(v []string)`

SetDeviceCredentialLabels sets DeviceCredentialLabels field to given value.

### HasDeviceCredentialLabels

`func (o *ProviderBiomarkerCapabilityRow) HasDeviceCredentialLabels() bool`

HasDeviceCredentialLabels returns a boolean if a field has been set.

### SetDeviceCredentialLabelsNil

`func (o *ProviderBiomarkerCapabilityRow) SetDeviceCredentialLabelsNil(b bool)`

 SetDeviceCredentialLabelsNil sets the value for DeviceCredentialLabels to be an explicit nil

### UnsetDeviceCredentialLabels
`func (o *ProviderBiomarkerCapabilityRow) UnsetDeviceCredentialLabels()`

UnsetDeviceCredentialLabels ensures that no value is present for DeviceCredentialLabels, not even an explicit nil
### GetSourceUrl

`func (o *ProviderBiomarkerCapabilityRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ProviderBiomarkerCapabilityRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ProviderBiomarkerCapabilityRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *ProviderBiomarkerCapabilityRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *ProviderBiomarkerCapabilityRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *ProviderBiomarkerCapabilityRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetCitationUrl

`func (o *ProviderBiomarkerCapabilityRow) GetCitationUrl() string`

GetCitationUrl returns the CitationUrl field if non-nil, zero value otherwise.

### GetCitationUrlOk

`func (o *ProviderBiomarkerCapabilityRow) GetCitationUrlOk() (*string, bool)`

GetCitationUrlOk returns a tuple with the CitationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationUrl

`func (o *ProviderBiomarkerCapabilityRow) SetCitationUrl(v string)`

SetCitationUrl sets CitationUrl field to given value.

### HasCitationUrl

`func (o *ProviderBiomarkerCapabilityRow) HasCitationUrl() bool`

HasCitationUrl returns a boolean if a field has been set.

### SetCitationUrlNil

`func (o *ProviderBiomarkerCapabilityRow) SetCitationUrlNil(b bool)`

 SetCitationUrlNil sets the value for CitationUrl to be an explicit nil

### UnsetCitationUrl
`func (o *ProviderBiomarkerCapabilityRow) UnsetCitationUrl()`

UnsetCitationUrl ensures that no value is present for CitationUrl, not even an explicit nil
### GetEvidenceSnippet

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceSnippet() string`

GetEvidenceSnippet returns the EvidenceSnippet field if non-nil, zero value otherwise.

### GetEvidenceSnippetOk

`func (o *ProviderBiomarkerCapabilityRow) GetEvidenceSnippetOk() (*string, bool)`

GetEvidenceSnippetOk returns a tuple with the EvidenceSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSnippet

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceSnippet(v string)`

SetEvidenceSnippet sets EvidenceSnippet field to given value.

### HasEvidenceSnippet

`func (o *ProviderBiomarkerCapabilityRow) HasEvidenceSnippet() bool`

HasEvidenceSnippet returns a boolean if a field has been set.

### SetEvidenceSnippetNil

`func (o *ProviderBiomarkerCapabilityRow) SetEvidenceSnippetNil(b bool)`

 SetEvidenceSnippetNil sets the value for EvidenceSnippet to be an explicit nil

### UnsetEvidenceSnippet
`func (o *ProviderBiomarkerCapabilityRow) UnsetEvidenceSnippet()`

UnsetEvidenceSnippet ensures that no value is present for EvidenceSnippet, not even an explicit nil
### GetMostRecentEvidenceDate

`func (o *ProviderBiomarkerCapabilityRow) GetMostRecentEvidenceDate() string`

GetMostRecentEvidenceDate returns the MostRecentEvidenceDate field if non-nil, zero value otherwise.

### GetMostRecentEvidenceDateOk

`func (o *ProviderBiomarkerCapabilityRow) GetMostRecentEvidenceDateOk() (*string, bool)`

GetMostRecentEvidenceDateOk returns a tuple with the MostRecentEvidenceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentEvidenceDate

`func (o *ProviderBiomarkerCapabilityRow) SetMostRecentEvidenceDate(v string)`

SetMostRecentEvidenceDate sets MostRecentEvidenceDate field to given value.

### HasMostRecentEvidenceDate

`func (o *ProviderBiomarkerCapabilityRow) HasMostRecentEvidenceDate() bool`

HasMostRecentEvidenceDate returns a boolean if a field has been set.

### SetMostRecentEvidenceDateNil

`func (o *ProviderBiomarkerCapabilityRow) SetMostRecentEvidenceDateNil(b bool)`

 SetMostRecentEvidenceDateNil sets the value for MostRecentEvidenceDate to be an explicit nil

### UnsetMostRecentEvidenceDate
`func (o *ProviderBiomarkerCapabilityRow) UnsetMostRecentEvidenceDate()`

UnsetMostRecentEvidenceDate ensures that no value is present for MostRecentEvidenceDate, not even an explicit nil
### GetMatchMethod

`func (o *ProviderBiomarkerCapabilityRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *ProviderBiomarkerCapabilityRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *ProviderBiomarkerCapabilityRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *ProviderBiomarkerCapabilityRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *ProviderBiomarkerCapabilityRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *ProviderBiomarkerCapabilityRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil
### GetMatchConfidence

`func (o *ProviderBiomarkerCapabilityRow) GetMatchConfidence() float32`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *ProviderBiomarkerCapabilityRow) GetMatchConfidenceOk() (*float32, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *ProviderBiomarkerCapabilityRow) SetMatchConfidence(v float32)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *ProviderBiomarkerCapabilityRow) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### SetMatchConfidenceNil

`func (o *ProviderBiomarkerCapabilityRow) SetMatchConfidenceNil(b bool)`

 SetMatchConfidenceNil sets the value for MatchConfidence to be an explicit nil

### UnsetMatchConfidence
`func (o *ProviderBiomarkerCapabilityRow) UnsetMatchConfidence()`

UnsetMatchConfidence ensures that no value is present for MatchConfidence, not even an explicit nil
### GetCompanyNameMatched

`func (o *ProviderBiomarkerCapabilityRow) GetCompanyNameMatched() string`

GetCompanyNameMatched returns the CompanyNameMatched field if non-nil, zero value otherwise.

### GetCompanyNameMatchedOk

`func (o *ProviderBiomarkerCapabilityRow) GetCompanyNameMatchedOk() (*string, bool)`

GetCompanyNameMatchedOk returns a tuple with the CompanyNameMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNameMatched

`func (o *ProviderBiomarkerCapabilityRow) SetCompanyNameMatched(v string)`

SetCompanyNameMatched sets CompanyNameMatched field to given value.

### HasCompanyNameMatched

`func (o *ProviderBiomarkerCapabilityRow) HasCompanyNameMatched() bool`

HasCompanyNameMatched returns a boolean if a field has been set.

### SetCompanyNameMatchedNil

`func (o *ProviderBiomarkerCapabilityRow) SetCompanyNameMatchedNil(b bool)`

 SetCompanyNameMatchedNil sets the value for CompanyNameMatched to be an explicit nil

### UnsetCompanyNameMatched
`func (o *ProviderBiomarkerCapabilityRow) UnsetCompanyNameMatched()`

UnsetCompanyNameMatched ensures that no value is present for CompanyNameMatched, not even an explicit nil
### GetSourceLicense

`func (o *ProviderBiomarkerCapabilityRow) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *ProviderBiomarkerCapabilityRow) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *ProviderBiomarkerCapabilityRow) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.

### HasSourceLicense

`func (o *ProviderBiomarkerCapabilityRow) HasSourceLicense() bool`

HasSourceLicense returns a boolean if a field has been set.

### SetSourceLicenseNil

`func (o *ProviderBiomarkerCapabilityRow) SetSourceLicenseNil(b bool)`

 SetSourceLicenseNil sets the value for SourceLicense to be an explicit nil

### UnsetSourceLicense
`func (o *ProviderBiomarkerCapabilityRow) UnsetSourceLicense()`

UnsetSourceLicense ensures that no value is present for SourceLicense, not even an explicit nil
### GetAttribution

`func (o *ProviderBiomarkerCapabilityRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ProviderBiomarkerCapabilityRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ProviderBiomarkerCapabilityRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ProviderBiomarkerCapabilityRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *ProviderBiomarkerCapabilityRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *ProviderBiomarkerCapabilityRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetAsOf

`func (o *ProviderBiomarkerCapabilityRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ProviderBiomarkerCapabilityRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ProviderBiomarkerCapabilityRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *ProviderBiomarkerCapabilityRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *ProviderBiomarkerCapabilityRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *ProviderBiomarkerCapabilityRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


