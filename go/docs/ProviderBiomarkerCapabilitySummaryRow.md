# ProviderBiomarkerCapabilitySummaryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryId** | Pointer to **NullableString** | The summary row id: a hash of the grain columns. | [optional] 
**ProviderId** | Pointer to **NullableString** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). A supplier is not a company. | [optional] 
**ProviderUuid** | Pointer to **NullableString** | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] 
**ProviderSlug** | Pointer to **NullableString** | The supplier&#39;s marketplace slug. | [optional] 
**ProviderName** | Pointer to **NullableString** | The supplier&#39;s name. | [optional] 
**RouteKey** | Pointer to **NullableString** | The biomarker&#39;s route key, the entry key for the biomarkers object type; the gene symbol for a gene biomarker. | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The canonical biomarker entity id. | [optional] 
**GeneSymbol** | Pointer to **NullableString** | The gene symbol; null for a source-scoped biomarker that names no gene. | [optional] 
**BiomarkerDisplayName** | Pointer to **NullableString** | The biomarker&#39;s display name. | [optional] 
**EvidenceLane** | Pointer to **NullableString** | The source path of the group: ladder_trial and ladder_clinvar (the evidence ladder), fda_cdx (FDA companion diagnostic list), fda_bqp (FDA Biomarker Qualification Program), literature_assay (assay providers named in PubMed full text), supplier_survey (the supplier&#39;s own declaration) clinvar_submission (ClinVar clinical-testing submissions by a laboratory the supplier carries by name, tier null) and supplier_service_listing (a Scientist.com marketplace catalogue service the supplier listed, whose name names the gene; tier null, no evidence sense). Nothing is merged across lanes. | [optional] 
**OrgRole** | Pointer to **NullableString** | The role the supplier played in the evidence: a ladder role, the literature provider role verbatim, respondent for the survey, or marketplace_listing for a service listing. | [optional] 
**MatchMethod** | Pointer to **NullableString** | How the evidence was attached to the supplier: org_member_exact (through the supplier&#39;s resolved organization), rx_name_key_exact (an exact name match, not a verified identity) or rx_provider_direct (the supplier&#39;s own marketplace listing, no name matched). | [optional] 
**AnalyticalValidationEvidence** | Pointer to **NullableString** | Analytical validation of the group (literature_assay lane only): fda_cleared_specific_assay, fda_cleared_for_biomarker or literature_reported. | [optional] 
**ClinicalUseEvidence** | Pointer to **NullableString** | Clinical-use evidence of the group: the ladder tier label, literature_reported, or on the clinvar_submission lane the ClinVar review status verbatim. | [optional] 
**RegulatoryQualification** | Pointer to **NullableString** | FDA qualification (fda_bqp lane only): fda_qualified. | [optional] 
**CompanionDiagnostic** | Pointer to **NullableString** | FDA companion diagnostic authorization (fda_cdx lane only): fda_authorized_companion_diagnostic. | [optional] 
**SupplierDeclaredStatus** | Pointer to **NullableString** | The supplier&#39;s own declaration, never mapped into the four evidence kinds: the survey&#39;s assay status on supplier_survey, supplier_declared_service_listed on supplier_service_listing. | [optional] 
**Tier** | Pointer to **NullableInt32** | The ladder rung of the group, 1 to 4, a kind of evidence and never a score; null on literature_assay, supplier_survey, clinvar_submission and supplier_service_listing. | [optional] 
**AssertionCount** | Pointer to **NullableInt64** | How many evidence assertions the row stands for; sums to the row count of provider_biomarker_capability under the same filter. | [optional] 
**DistinctEvidenceRefCount** | Pointer to **NullableInt64** | How many distinct evidence references (one ClinVar record, one trial, one paper) the assertions cite. | [optional] 
**ExampleEvidenceRef** | Pointer to **NullableString** | The smallest evidence reference of the group: a deterministic example to link, not a ranking. | [optional] 
**ExampleEvidenceRefKind** | Pointer to **NullableString** | Which identifier example_evidence_ref is: nct_id, clinvar_scv, fda_cdx_entry_biomarker, fda_bqp_project, pmid, assay_ledger_row, survey_response_biomarker or rx_ware. | [optional] 
**ExampleCitationUrl** | Pointer to **NullableString** | The citation page of the example reference; null where the source offers no resolvable page. | [optional] 
**MostRecentEvidenceDate** | Pointer to **NullableString** | The latest publication date in the group (literature_assay lane only). | [optional] 
**SourceLicense** | Pointer to **NullableString** | The license statement of the group&#39;s source lane. No lane carries a usage restriction. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit the group&#39;s source lane is owed. | [optional] 
**AsOf** | Pointer to **NullableString** | The build date. | [optional] 

## Methods

### NewProviderBiomarkerCapabilitySummaryRow

`func NewProviderBiomarkerCapabilitySummaryRow() *ProviderBiomarkerCapabilitySummaryRow`

NewProviderBiomarkerCapabilitySummaryRow instantiates a new ProviderBiomarkerCapabilitySummaryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderBiomarkerCapabilitySummaryRowWithDefaults

`func NewProviderBiomarkerCapabilitySummaryRowWithDefaults() *ProviderBiomarkerCapabilitySummaryRow`

NewProviderBiomarkerCapabilitySummaryRowWithDefaults instantiates a new ProviderBiomarkerCapabilitySummaryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryId

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetSummaryId() string`

GetSummaryId returns the SummaryId field if non-nil, zero value otherwise.

### GetSummaryIdOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetSummaryIdOk() (*string, bool)`

GetSummaryIdOk returns a tuple with the SummaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryId

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetSummaryId(v string)`

SetSummaryId sets SummaryId field to given value.

### HasSummaryId

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasSummaryId() bool`

HasSummaryId returns a boolean if a field has been set.

### SetSummaryIdNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetSummaryIdNil(b bool)`

 SetSummaryIdNil sets the value for SummaryId to be an explicit nil

### UnsetSummaryId
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetSummaryId()`

UnsetSummaryId ensures that no value is present for SummaryId, not even an explicit nil
### GetProviderId

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderUuid

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.

### HasProviderUuid

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasProviderUuid() bool`

HasProviderUuid returns a boolean if a field has been set.

### SetProviderUuidNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetProviderUuidNil(b bool)`

 SetProviderUuidNil sets the value for ProviderUuid to be an explicit nil

### UnsetProviderUuid
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetProviderUuid()`

UnsetProviderUuid ensures that no value is present for ProviderUuid, not even an explicit nil
### GetProviderSlug

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetProviderSlug() string`

GetProviderSlug returns the ProviderSlug field if non-nil, zero value otherwise.

### GetProviderSlugOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetProviderSlugOk() (*string, bool)`

GetProviderSlugOk returns a tuple with the ProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSlug

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetProviderSlug(v string)`

SetProviderSlug sets ProviderSlug field to given value.

### HasProviderSlug

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasProviderSlug() bool`

HasProviderSlug returns a boolean if a field has been set.

### SetProviderSlugNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetProviderSlugNil(b bool)`

 SetProviderSlugNil sets the value for ProviderSlug to be an explicit nil

### UnsetProviderSlug
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetProviderSlug()`

UnsetProviderSlug ensures that no value is present for ProviderSlug, not even an explicit nil
### GetProviderName

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetRouteKey

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerId

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetGeneSymbol

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### SetGeneSymbolNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetBiomarkerDisplayName

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetBiomarkerDisplayName() string`

GetBiomarkerDisplayName returns the BiomarkerDisplayName field if non-nil, zero value otherwise.

### GetBiomarkerDisplayNameOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetBiomarkerDisplayNameOk() (*string, bool)`

GetBiomarkerDisplayNameOk returns a tuple with the BiomarkerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerDisplayName

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetBiomarkerDisplayName(v string)`

SetBiomarkerDisplayName sets BiomarkerDisplayName field to given value.

### HasBiomarkerDisplayName

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasBiomarkerDisplayName() bool`

HasBiomarkerDisplayName returns a boolean if a field has been set.

### SetBiomarkerDisplayNameNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetBiomarkerDisplayNameNil(b bool)`

 SetBiomarkerDisplayNameNil sets the value for BiomarkerDisplayName to be an explicit nil

### UnsetBiomarkerDisplayName
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetBiomarkerDisplayName()`

UnsetBiomarkerDisplayName ensures that no value is present for BiomarkerDisplayName, not even an explicit nil
### GetEvidenceLane

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetEvidenceLane() string`

GetEvidenceLane returns the EvidenceLane field if non-nil, zero value otherwise.

### GetEvidenceLaneOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetEvidenceLaneOk() (*string, bool)`

GetEvidenceLaneOk returns a tuple with the EvidenceLane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceLane

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetEvidenceLane(v string)`

SetEvidenceLane sets EvidenceLane field to given value.

### HasEvidenceLane

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasEvidenceLane() bool`

HasEvidenceLane returns a boolean if a field has been set.

### SetEvidenceLaneNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetEvidenceLaneNil(b bool)`

 SetEvidenceLaneNil sets the value for EvidenceLane to be an explicit nil

### UnsetEvidenceLane
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetEvidenceLane()`

UnsetEvidenceLane ensures that no value is present for EvidenceLane, not even an explicit nil
### GetOrgRole

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetOrgRole() string`

GetOrgRole returns the OrgRole field if non-nil, zero value otherwise.

### GetOrgRoleOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetOrgRoleOk() (*string, bool)`

GetOrgRoleOk returns a tuple with the OrgRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRole

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetOrgRole(v string)`

SetOrgRole sets OrgRole field to given value.

### HasOrgRole

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasOrgRole() bool`

HasOrgRole returns a boolean if a field has been set.

### SetOrgRoleNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetOrgRoleNil(b bool)`

 SetOrgRoleNil sets the value for OrgRole to be an explicit nil

### UnsetOrgRole
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetOrgRole()`

UnsetOrgRole ensures that no value is present for OrgRole, not even an explicit nil
### GetMatchMethod

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### SetMatchMethodNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetMatchMethodNil(b bool)`

 SetMatchMethodNil sets the value for MatchMethod to be an explicit nil

### UnsetMatchMethod
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetMatchMethod()`

UnsetMatchMethod ensures that no value is present for MatchMethod, not even an explicit nil
### GetAnalyticalValidationEvidence

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetAnalyticalValidationEvidence() string`

GetAnalyticalValidationEvidence returns the AnalyticalValidationEvidence field if non-nil, zero value otherwise.

### GetAnalyticalValidationEvidenceOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetAnalyticalValidationEvidenceOk() (*string, bool)`

GetAnalyticalValidationEvidenceOk returns a tuple with the AnalyticalValidationEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticalValidationEvidence

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetAnalyticalValidationEvidence(v string)`

SetAnalyticalValidationEvidence sets AnalyticalValidationEvidence field to given value.

### HasAnalyticalValidationEvidence

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasAnalyticalValidationEvidence() bool`

HasAnalyticalValidationEvidence returns a boolean if a field has been set.

### SetAnalyticalValidationEvidenceNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetAnalyticalValidationEvidenceNil(b bool)`

 SetAnalyticalValidationEvidenceNil sets the value for AnalyticalValidationEvidence to be an explicit nil

### UnsetAnalyticalValidationEvidence
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetAnalyticalValidationEvidence()`

UnsetAnalyticalValidationEvidence ensures that no value is present for AnalyticalValidationEvidence, not even an explicit nil
### GetClinicalUseEvidence

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetClinicalUseEvidence() string`

GetClinicalUseEvidence returns the ClinicalUseEvidence field if non-nil, zero value otherwise.

### GetClinicalUseEvidenceOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetClinicalUseEvidenceOk() (*string, bool)`

GetClinicalUseEvidenceOk returns a tuple with the ClinicalUseEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinicalUseEvidence

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetClinicalUseEvidence(v string)`

SetClinicalUseEvidence sets ClinicalUseEvidence field to given value.

### HasClinicalUseEvidence

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasClinicalUseEvidence() bool`

HasClinicalUseEvidence returns a boolean if a field has been set.

### SetClinicalUseEvidenceNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetClinicalUseEvidenceNil(b bool)`

 SetClinicalUseEvidenceNil sets the value for ClinicalUseEvidence to be an explicit nil

### UnsetClinicalUseEvidence
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetClinicalUseEvidence()`

UnsetClinicalUseEvidence ensures that no value is present for ClinicalUseEvidence, not even an explicit nil
### GetRegulatoryQualification

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetRegulatoryQualification() string`

GetRegulatoryQualification returns the RegulatoryQualification field if non-nil, zero value otherwise.

### GetRegulatoryQualificationOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetRegulatoryQualificationOk() (*string, bool)`

GetRegulatoryQualificationOk returns a tuple with the RegulatoryQualification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryQualification

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetRegulatoryQualification(v string)`

SetRegulatoryQualification sets RegulatoryQualification field to given value.

### HasRegulatoryQualification

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasRegulatoryQualification() bool`

HasRegulatoryQualification returns a boolean if a field has been set.

### SetRegulatoryQualificationNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetRegulatoryQualificationNil(b bool)`

 SetRegulatoryQualificationNil sets the value for RegulatoryQualification to be an explicit nil

### UnsetRegulatoryQualification
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetRegulatoryQualification()`

UnsetRegulatoryQualification ensures that no value is present for RegulatoryQualification, not even an explicit nil
### GetCompanionDiagnostic

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetCompanionDiagnostic() string`

GetCompanionDiagnostic returns the CompanionDiagnostic field if non-nil, zero value otherwise.

### GetCompanionDiagnosticOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetCompanionDiagnosticOk() (*string, bool)`

GetCompanionDiagnosticOk returns a tuple with the CompanionDiagnostic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanionDiagnostic

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetCompanionDiagnostic(v string)`

SetCompanionDiagnostic sets CompanionDiagnostic field to given value.

### HasCompanionDiagnostic

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasCompanionDiagnostic() bool`

HasCompanionDiagnostic returns a boolean if a field has been set.

### SetCompanionDiagnosticNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetCompanionDiagnosticNil(b bool)`

 SetCompanionDiagnosticNil sets the value for CompanionDiagnostic to be an explicit nil

### UnsetCompanionDiagnostic
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetCompanionDiagnostic()`

UnsetCompanionDiagnostic ensures that no value is present for CompanionDiagnostic, not even an explicit nil
### GetSupplierDeclaredStatus

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetSupplierDeclaredStatus() string`

GetSupplierDeclaredStatus returns the SupplierDeclaredStatus field if non-nil, zero value otherwise.

### GetSupplierDeclaredStatusOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetSupplierDeclaredStatusOk() (*string, bool)`

GetSupplierDeclaredStatusOk returns a tuple with the SupplierDeclaredStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierDeclaredStatus

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetSupplierDeclaredStatus(v string)`

SetSupplierDeclaredStatus sets SupplierDeclaredStatus field to given value.

### HasSupplierDeclaredStatus

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasSupplierDeclaredStatus() bool`

HasSupplierDeclaredStatus returns a boolean if a field has been set.

### SetSupplierDeclaredStatusNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetSupplierDeclaredStatusNil(b bool)`

 SetSupplierDeclaredStatusNil sets the value for SupplierDeclaredStatus to be an explicit nil

### UnsetSupplierDeclaredStatus
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetSupplierDeclaredStatus()`

UnsetSupplierDeclaredStatus ensures that no value is present for SupplierDeclaredStatus, not even an explicit nil
### GetTier

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetAssertionCount

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetAssertionCount() int64`

GetAssertionCount returns the AssertionCount field if non-nil, zero value otherwise.

### GetAssertionCountOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetAssertionCountOk() (*int64, bool)`

GetAssertionCountOk returns a tuple with the AssertionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionCount

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetAssertionCount(v int64)`

SetAssertionCount sets AssertionCount field to given value.

### HasAssertionCount

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasAssertionCount() bool`

HasAssertionCount returns a boolean if a field has been set.

### SetAssertionCountNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetAssertionCountNil(b bool)`

 SetAssertionCountNil sets the value for AssertionCount to be an explicit nil

### UnsetAssertionCount
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetAssertionCount()`

UnsetAssertionCount ensures that no value is present for AssertionCount, not even an explicit nil
### GetDistinctEvidenceRefCount

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetDistinctEvidenceRefCount() int64`

GetDistinctEvidenceRefCount returns the DistinctEvidenceRefCount field if non-nil, zero value otherwise.

### GetDistinctEvidenceRefCountOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetDistinctEvidenceRefCountOk() (*int64, bool)`

GetDistinctEvidenceRefCountOk returns a tuple with the DistinctEvidenceRefCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctEvidenceRefCount

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetDistinctEvidenceRefCount(v int64)`

SetDistinctEvidenceRefCount sets DistinctEvidenceRefCount field to given value.

### HasDistinctEvidenceRefCount

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasDistinctEvidenceRefCount() bool`

HasDistinctEvidenceRefCount returns a boolean if a field has been set.

### SetDistinctEvidenceRefCountNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetDistinctEvidenceRefCountNil(b bool)`

 SetDistinctEvidenceRefCountNil sets the value for DistinctEvidenceRefCount to be an explicit nil

### UnsetDistinctEvidenceRefCount
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetDistinctEvidenceRefCount()`

UnsetDistinctEvidenceRefCount ensures that no value is present for DistinctEvidenceRefCount, not even an explicit nil
### GetExampleEvidenceRef

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetExampleEvidenceRef() string`

GetExampleEvidenceRef returns the ExampleEvidenceRef field if non-nil, zero value otherwise.

### GetExampleEvidenceRefOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetExampleEvidenceRefOk() (*string, bool)`

GetExampleEvidenceRefOk returns a tuple with the ExampleEvidenceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleEvidenceRef

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetExampleEvidenceRef(v string)`

SetExampleEvidenceRef sets ExampleEvidenceRef field to given value.

### HasExampleEvidenceRef

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasExampleEvidenceRef() bool`

HasExampleEvidenceRef returns a boolean if a field has been set.

### SetExampleEvidenceRefNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetExampleEvidenceRefNil(b bool)`

 SetExampleEvidenceRefNil sets the value for ExampleEvidenceRef to be an explicit nil

### UnsetExampleEvidenceRef
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetExampleEvidenceRef()`

UnsetExampleEvidenceRef ensures that no value is present for ExampleEvidenceRef, not even an explicit nil
### GetExampleEvidenceRefKind

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetExampleEvidenceRefKind() string`

GetExampleEvidenceRefKind returns the ExampleEvidenceRefKind field if non-nil, zero value otherwise.

### GetExampleEvidenceRefKindOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetExampleEvidenceRefKindOk() (*string, bool)`

GetExampleEvidenceRefKindOk returns a tuple with the ExampleEvidenceRefKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleEvidenceRefKind

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetExampleEvidenceRefKind(v string)`

SetExampleEvidenceRefKind sets ExampleEvidenceRefKind field to given value.

### HasExampleEvidenceRefKind

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasExampleEvidenceRefKind() bool`

HasExampleEvidenceRefKind returns a boolean if a field has been set.

### SetExampleEvidenceRefKindNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetExampleEvidenceRefKindNil(b bool)`

 SetExampleEvidenceRefKindNil sets the value for ExampleEvidenceRefKind to be an explicit nil

### UnsetExampleEvidenceRefKind
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetExampleEvidenceRefKind()`

UnsetExampleEvidenceRefKind ensures that no value is present for ExampleEvidenceRefKind, not even an explicit nil
### GetExampleCitationUrl

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetExampleCitationUrl() string`

GetExampleCitationUrl returns the ExampleCitationUrl field if non-nil, zero value otherwise.

### GetExampleCitationUrlOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetExampleCitationUrlOk() (*string, bool)`

GetExampleCitationUrlOk returns a tuple with the ExampleCitationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleCitationUrl

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetExampleCitationUrl(v string)`

SetExampleCitationUrl sets ExampleCitationUrl field to given value.

### HasExampleCitationUrl

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasExampleCitationUrl() bool`

HasExampleCitationUrl returns a boolean if a field has been set.

### SetExampleCitationUrlNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetExampleCitationUrlNil(b bool)`

 SetExampleCitationUrlNil sets the value for ExampleCitationUrl to be an explicit nil

### UnsetExampleCitationUrl
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetExampleCitationUrl()`

UnsetExampleCitationUrl ensures that no value is present for ExampleCitationUrl, not even an explicit nil
### GetMostRecentEvidenceDate

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetMostRecentEvidenceDate() string`

GetMostRecentEvidenceDate returns the MostRecentEvidenceDate field if non-nil, zero value otherwise.

### GetMostRecentEvidenceDateOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetMostRecentEvidenceDateOk() (*string, bool)`

GetMostRecentEvidenceDateOk returns a tuple with the MostRecentEvidenceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentEvidenceDate

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetMostRecentEvidenceDate(v string)`

SetMostRecentEvidenceDate sets MostRecentEvidenceDate field to given value.

### HasMostRecentEvidenceDate

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasMostRecentEvidenceDate() bool`

HasMostRecentEvidenceDate returns a boolean if a field has been set.

### SetMostRecentEvidenceDateNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetMostRecentEvidenceDateNil(b bool)`

 SetMostRecentEvidenceDateNil sets the value for MostRecentEvidenceDate to be an explicit nil

### UnsetMostRecentEvidenceDate
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetMostRecentEvidenceDate()`

UnsetMostRecentEvidenceDate ensures that no value is present for MostRecentEvidenceDate, not even an explicit nil
### GetSourceLicense

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.

### HasSourceLicense

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasSourceLicense() bool`

HasSourceLicense returns a boolean if a field has been set.

### SetSourceLicenseNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetSourceLicenseNil(b bool)`

 SetSourceLicenseNil sets the value for SourceLicense to be an explicit nil

### UnsetSourceLicense
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetSourceLicense()`

UnsetSourceLicense ensures that no value is present for SourceLicense, not even an explicit nil
### GetAttribution

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetAsOf

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ProviderBiomarkerCapabilitySummaryRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *ProviderBiomarkerCapabilitySummaryRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *ProviderBiomarkerCapabilitySummaryRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *ProviderBiomarkerCapabilitySummaryRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


