# BiomarkerAssayProvidersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | Pointer to **NullableString** | The route key of the gene the biomarker resolved to, the entry key for the biomarkers object type. Null when the extraction did not resolve a gene. | [optional] 
**BiomarkerKey** | Pointer to **NullableString** | The grouping key of the biomarker. The biomarker&#39;s key when its gene resolved, otherwise raw: followed by the lower-cased name as written, so two spellings of one unresolved biomarker are two keys. | [optional] 
**BiomarkerId** | Pointer to **NullableString** | The biomarker&#39;s own key; null for a pair whose gene did not resolve. | [optional] 
**BiomarkerDisplayName** | Pointer to **NullableString** | The biomarker&#39;s name, as resolved or else as extracted. | [optional] 
**GeneResolutionStatus** | Pointer to **NullableString** | resolved: the gene resolved. unresolved: a gene symbol was extracted and matched nothing. not_applicable: no gene symbol was extracted. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The provider the rows are grouped on. org: followed by the organization key when the company resolves to an organization, otherwise name: followed by the normalised company name. | [optional] 
**ProviderRole** | Pointer to **NullableString** | What the papers name the provider as. assay_or_test_maker makes the assay or test; testing_service runs it as a service; both when the pair has papers of each kind. | [optional] 
**ProviderName** | Pointer to **NullableString** | The provider company name as it appears in the papers. | [optional] 
**ProviderOrgId** | Pointer to **NullableString** | The organization the provider&#39;s company name resolved to (the member organization, not its parent); null when the provider is ambiguous or unresolved. Reached by the rung named in provider_org_match_tier. | [optional] 
**ProviderOrgMatchTier** | Pointer to **NullableString** | How the provider was resolved to an organization, first rung that yields one organization wins. alias_exact: a reviewed brand alias, then exactly one organization carries the name. name_exact: exactly one organization carries the name. parent_unanimous: several organizations carry the name and all roll up to one parent, which is the organization. sponsor_key: the name is a trial-sponsor key. rx_provider_link: through its marketplace supplier. ambiguous: organizations under more than one parent carry the name, so none is chosen. unresolved: no organization, the provider stays a name. | [optional] 
**ProviderParentOrgId** | Pointer to **NullableString** | The resolved parent of provider_org_id in the organization hierarchy (the organization itself when it has none); null when the provider has no organization. The hierarchy is not guaranteed correct for every company, so read it with provider_org_match_tier. | [optional] 
**ProviderCompanyId** | Pointer to **NullableString** | The provider&#39;s company key. | [optional] 
**ProviderCompanyName** | Pointer to **NullableString** | The provider&#39;s company name. | [optional] 
**ProviderGroupId** | Pointer to **NullableString** | The provider&#39;s current ultimate parent group key. | [optional] 
**ProviderGroupName** | Pointer to **NullableString** | The provider&#39;s group name. | [optional] 
**RxMatched** | Pointer to **NullableBool** | Whether the provider name matches a supplier in the Scientist.com marketplace. | [optional] 
**RxProviderId** | Pointer to **NullableString** | The matched marketplace supplier&#39;s id; join to the suppliers dataset. | [optional] 
**RxProviderName** | Pointer to **NullableString** | The matched marketplace supplier&#39;s name. | [optional] 
**RxProviderSlug** | Pointer to **NullableString** | The matched supplier&#39;s marketplace profile slug. | [optional] 
**HasClinicalValidation** | Pointer to **NullableBool** | Whether any paper for the pair mentions clinical validation of the assay. | [optional] 
**HasAnalyticalValidation** | Pointer to **NullableBool** | Whether any paper for the pair mentions analytical validation of the assay. | [optional] 
**NPapers** | Pointer to **NullableInt64** | Papers naming the biomarker and the provider, dated or not. | [optional] 
**NDatedPapers** | Pointer to **NullableInt64** | Papers with a recoverable publication date. | [optional] 
**NUndatedPapers** | Pointer to **NullableInt64** | Papers with no recoverable publication date. n_dated_papers plus n_undated_papers is n_papers. | [optional] 
**MostRecentPublishedAt** | Pointer to **NullableTime** | The latest publication date among the dated papers, in UTC. Null when no paper is dated, which is not a statement that none was published recently. | [optional] 
**MostRecentPmid** | Pointer to **NullableString** | The PubMed id of the most recent dated paper. | [optional] 
**MostRecentTitle** | Pointer to **NullableString** | The title of that paper. | [optional] 
**MeetsCuratedThreshold** | Pointer to **NullableBool** | True when the provider matches a marketplace supplier or at least two papers name the pair. | [optional] 
**RegulatoryEvidenceKind** | Pointer to **NullableString** | The strongest US FDA or CMS record found for the pair. companion_diagnostic: the provider holds an FDA companion diagnostic for the biomarker. fda_cleared_specific_assay: a cleared or approved device of the provider that the paper names and that measures the biomarker. fda_cleared_for_biomarker: the provider holds a cleared or approved IVD under a product code for the biomarker. clia_testing_service: the provider is named as a testing service and its own organization holds a CLIA certificate. fda_ivd_company_only: the provider holds some cleared or approved IVD, none matched to the biomarker. none: no FDA record found. | [optional] 
**HasFdaRecordForBiomarker** | Pointer to **NullableBool** | True when regulatory_evidence_kind is one of the first four kinds, a record that concerns the biomarker rather than only the company. | [optional] 
**RegulatoryEvidenceReasons** | Pointer to **[]string** | Every level that held for the pair, strongest first; empty when none did. | [optional] 
**NFdaIvd510k** | Pointer to **NullableInt64** | The 510(k) clearances in the in-vitro diagnostic review panels held by the provider&#39;s organization (or, when it holds none, by organizations under its parent; fda_company_match_rung says which). Null for a provider that did not resolve to an organization. | [optional] 
**NFdaIvdPma** | Pointer to **NullableInt64** | The PMA approvals (distinct approval numbers) in the in-vitro diagnostic review panels held by the provider&#39;s organization or its parent family, on the same rung as n_fda_ivd_510k. Null for a provider that did not resolve to an organization. | [optional] 
**FdaCompanyMatchRung** | Pointer to **NullableString** | org: the counted records are held by the provider&#39;s own organization. parent: by organizations under the same resolved parent. Null when the provider holds none. | [optional] 
**NFdaAnalyte510k** | Pointer to **NullableInt64** | The provider&#39;s 510(k) clearances under a product code mapped to this biomarker. Null when the provider or the biomarker did not resolve. | [optional] 
**NFdaAnalytePma** | Pointer to **NullableInt64** | The provider&#39;s PMA approvals under a product code mapped to this biomarker. Null when the provider or the biomarker did not resolve. | [optional] 
**FdaAnalyteProductCodes** | Pointer to **[]string** | The FDA product codes behind n_fda_analyte_510k and n_fda_analyte_pma. | [optional] 
**FdaAnalyteFdaUrl** | Pointer to **NullableString** | The FDA database page of the provider&#39;s most recent record under a product code for this biomarker. | [optional] 
**SpecificAssayMatched** | Pointer to **NullableBool** | True when a paper names a cleared or approved device of this provider that measures this biomarker. The device name matches the paper&#39;s assay name, or the company&#39;s own quote, and the device&#39;s analyte agrees with the biomarker. | [optional] 
**SpecificAssayDeviceName** | Pointer to **NullableString** | The 510(k) device name or PMA trade name of the matched device; the best match when several matched. Null when no device matched. | [optional] 
**SpecificAssayKNumber** | Pointer to **NullableString** | The 510(k) number of the matched device; null when the match is a PMA or there is none. | [optional] 
**SpecificAssayPmaNumber** | Pointer to **NullableString** | The PMA number of the matched device; null when the match is a 510(k) or there is none. | [optional] 
**SpecificAssayFdaUrl** | Pointer to **NullableString** | The FDA database page of the matched device. | [optional] 
**SpecificAssaySimilarity** | Pointer to **NullableFloat32** | The trigram text similarity (0 to 1) of the larger of the two text clauses that matched the device, the paper&#39;s assay name or the company&#39;s quote against the device name. | [optional] 
**NCdxEntries** | Pointer to **NullableInt32** | The FDA companion diagnostic list entries for this biomarker from this provider&#39;s organization (or its parent family). Null when the provider or the biomarker did not resolve. | [optional] 
**CdxDiagnosticNames** | Pointer to **[]string** | The names of those companion diagnostics. | [optional] 
**NCliaCertificates** | Pointer to **NullableInt64** | The CLIA (US clinical laboratory) certificates held by the provider&#39;s organization, or by organizations under its parent when it holds none (clia_match_rung says which). A certificate licenses a lab; it is not evidence that any named assay is run. Null for a provider that did not resolve to an organization. | [optional] 
**CliaMatchRung** | Pointer to **NullableString** | org: the counted certificates are held by the provider&#39;s own organization. parent: by organizations under the same resolved parent. Null when the provider holds none. | [optional] 

## Methods

### NewBiomarkerAssayProvidersRow

`func NewBiomarkerAssayProvidersRow() *BiomarkerAssayProvidersRow`

NewBiomarkerAssayProvidersRow instantiates a new BiomarkerAssayProvidersRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkerAssayProvidersRowWithDefaults

`func NewBiomarkerAssayProvidersRowWithDefaults() *BiomarkerAssayProvidersRow`

NewBiomarkerAssayProvidersRowWithDefaults instantiates a new BiomarkerAssayProvidersRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkerAssayProvidersRow) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkerAssayProvidersRow) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkerAssayProvidersRow) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.

### HasRouteKey

`func (o *BiomarkerAssayProvidersRow) HasRouteKey() bool`

HasRouteKey returns a boolean if a field has been set.

### SetRouteKeyNil

`func (o *BiomarkerAssayProvidersRow) SetRouteKeyNil(b bool)`

 SetRouteKeyNil sets the value for RouteKey to be an explicit nil

### UnsetRouteKey
`func (o *BiomarkerAssayProvidersRow) UnsetRouteKey()`

UnsetRouteKey ensures that no value is present for RouteKey, not even an explicit nil
### GetBiomarkerKey

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerKey() string`

GetBiomarkerKey returns the BiomarkerKey field if non-nil, zero value otherwise.

### GetBiomarkerKeyOk

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerKeyOk() (*string, bool)`

GetBiomarkerKeyOk returns a tuple with the BiomarkerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerKey

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerKey(v string)`

SetBiomarkerKey sets BiomarkerKey field to given value.

### HasBiomarkerKey

`func (o *BiomarkerAssayProvidersRow) HasBiomarkerKey() bool`

HasBiomarkerKey returns a boolean if a field has been set.

### SetBiomarkerKeyNil

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerKeyNil(b bool)`

 SetBiomarkerKeyNil sets the value for BiomarkerKey to be an explicit nil

### UnsetBiomarkerKey
`func (o *BiomarkerAssayProvidersRow) UnsetBiomarkerKey()`

UnsetBiomarkerKey ensures that no value is present for BiomarkerKey, not even an explicit nil
### GetBiomarkerId

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerId() string`

GetBiomarkerId returns the BiomarkerId field if non-nil, zero value otherwise.

### GetBiomarkerIdOk

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerIdOk() (*string, bool)`

GetBiomarkerIdOk returns a tuple with the BiomarkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerId

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerId(v string)`

SetBiomarkerId sets BiomarkerId field to given value.

### HasBiomarkerId

`func (o *BiomarkerAssayProvidersRow) HasBiomarkerId() bool`

HasBiomarkerId returns a boolean if a field has been set.

### SetBiomarkerIdNil

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerIdNil(b bool)`

 SetBiomarkerIdNil sets the value for BiomarkerId to be an explicit nil

### UnsetBiomarkerId
`func (o *BiomarkerAssayProvidersRow) UnsetBiomarkerId()`

UnsetBiomarkerId ensures that no value is present for BiomarkerId, not even an explicit nil
### GetBiomarkerDisplayName

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerDisplayName() string`

GetBiomarkerDisplayName returns the BiomarkerDisplayName field if non-nil, zero value otherwise.

### GetBiomarkerDisplayNameOk

`func (o *BiomarkerAssayProvidersRow) GetBiomarkerDisplayNameOk() (*string, bool)`

GetBiomarkerDisplayNameOk returns a tuple with the BiomarkerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerDisplayName

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerDisplayName(v string)`

SetBiomarkerDisplayName sets BiomarkerDisplayName field to given value.

### HasBiomarkerDisplayName

`func (o *BiomarkerAssayProvidersRow) HasBiomarkerDisplayName() bool`

HasBiomarkerDisplayName returns a boolean if a field has been set.

### SetBiomarkerDisplayNameNil

`func (o *BiomarkerAssayProvidersRow) SetBiomarkerDisplayNameNil(b bool)`

 SetBiomarkerDisplayNameNil sets the value for BiomarkerDisplayName to be an explicit nil

### UnsetBiomarkerDisplayName
`func (o *BiomarkerAssayProvidersRow) UnsetBiomarkerDisplayName()`

UnsetBiomarkerDisplayName ensures that no value is present for BiomarkerDisplayName, not even an explicit nil
### GetGeneResolutionStatus

`func (o *BiomarkerAssayProvidersRow) GetGeneResolutionStatus() string`

GetGeneResolutionStatus returns the GeneResolutionStatus field if non-nil, zero value otherwise.

### GetGeneResolutionStatusOk

`func (o *BiomarkerAssayProvidersRow) GetGeneResolutionStatusOk() (*string, bool)`

GetGeneResolutionStatusOk returns a tuple with the GeneResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneResolutionStatus

`func (o *BiomarkerAssayProvidersRow) SetGeneResolutionStatus(v string)`

SetGeneResolutionStatus sets GeneResolutionStatus field to given value.

### HasGeneResolutionStatus

`func (o *BiomarkerAssayProvidersRow) HasGeneResolutionStatus() bool`

HasGeneResolutionStatus returns a boolean if a field has been set.

### SetGeneResolutionStatusNil

`func (o *BiomarkerAssayProvidersRow) SetGeneResolutionStatusNil(b bool)`

 SetGeneResolutionStatusNil sets the value for GeneResolutionStatus to be an explicit nil

### UnsetGeneResolutionStatus
`func (o *BiomarkerAssayProvidersRow) UnsetGeneResolutionStatus()`

UnsetGeneResolutionStatus ensures that no value is present for GeneResolutionStatus, not even an explicit nil
### GetProviderKey

`func (o *BiomarkerAssayProvidersRow) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *BiomarkerAssayProvidersRow) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *BiomarkerAssayProvidersRow) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *BiomarkerAssayProvidersRow) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *BiomarkerAssayProvidersRow) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *BiomarkerAssayProvidersRow) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetProviderRole

`func (o *BiomarkerAssayProvidersRow) GetProviderRole() string`

GetProviderRole returns the ProviderRole field if non-nil, zero value otherwise.

### GetProviderRoleOk

`func (o *BiomarkerAssayProvidersRow) GetProviderRoleOk() (*string, bool)`

GetProviderRoleOk returns a tuple with the ProviderRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRole

`func (o *BiomarkerAssayProvidersRow) SetProviderRole(v string)`

SetProviderRole sets ProviderRole field to given value.

### HasProviderRole

`func (o *BiomarkerAssayProvidersRow) HasProviderRole() bool`

HasProviderRole returns a boolean if a field has been set.

### SetProviderRoleNil

`func (o *BiomarkerAssayProvidersRow) SetProviderRoleNil(b bool)`

 SetProviderRoleNil sets the value for ProviderRole to be an explicit nil

### UnsetProviderRole
`func (o *BiomarkerAssayProvidersRow) UnsetProviderRole()`

UnsetProviderRole ensures that no value is present for ProviderRole, not even an explicit nil
### GetProviderName

`func (o *BiomarkerAssayProvidersRow) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *BiomarkerAssayProvidersRow) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *BiomarkerAssayProvidersRow) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *BiomarkerAssayProvidersRow) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *BiomarkerAssayProvidersRow) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *BiomarkerAssayProvidersRow) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetProviderOrgId

`func (o *BiomarkerAssayProvidersRow) GetProviderOrgId() string`

GetProviderOrgId returns the ProviderOrgId field if non-nil, zero value otherwise.

### GetProviderOrgIdOk

`func (o *BiomarkerAssayProvidersRow) GetProviderOrgIdOk() (*string, bool)`

GetProviderOrgIdOk returns a tuple with the ProviderOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOrgId

`func (o *BiomarkerAssayProvidersRow) SetProviderOrgId(v string)`

SetProviderOrgId sets ProviderOrgId field to given value.

### HasProviderOrgId

`func (o *BiomarkerAssayProvidersRow) HasProviderOrgId() bool`

HasProviderOrgId returns a boolean if a field has been set.

### SetProviderOrgIdNil

`func (o *BiomarkerAssayProvidersRow) SetProviderOrgIdNil(b bool)`

 SetProviderOrgIdNil sets the value for ProviderOrgId to be an explicit nil

### UnsetProviderOrgId
`func (o *BiomarkerAssayProvidersRow) UnsetProviderOrgId()`

UnsetProviderOrgId ensures that no value is present for ProviderOrgId, not even an explicit nil
### GetProviderOrgMatchTier

`func (o *BiomarkerAssayProvidersRow) GetProviderOrgMatchTier() string`

GetProviderOrgMatchTier returns the ProviderOrgMatchTier field if non-nil, zero value otherwise.

### GetProviderOrgMatchTierOk

`func (o *BiomarkerAssayProvidersRow) GetProviderOrgMatchTierOk() (*string, bool)`

GetProviderOrgMatchTierOk returns a tuple with the ProviderOrgMatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOrgMatchTier

`func (o *BiomarkerAssayProvidersRow) SetProviderOrgMatchTier(v string)`

SetProviderOrgMatchTier sets ProviderOrgMatchTier field to given value.

### HasProviderOrgMatchTier

`func (o *BiomarkerAssayProvidersRow) HasProviderOrgMatchTier() bool`

HasProviderOrgMatchTier returns a boolean if a field has been set.

### SetProviderOrgMatchTierNil

`func (o *BiomarkerAssayProvidersRow) SetProviderOrgMatchTierNil(b bool)`

 SetProviderOrgMatchTierNil sets the value for ProviderOrgMatchTier to be an explicit nil

### UnsetProviderOrgMatchTier
`func (o *BiomarkerAssayProvidersRow) UnsetProviderOrgMatchTier()`

UnsetProviderOrgMatchTier ensures that no value is present for ProviderOrgMatchTier, not even an explicit nil
### GetProviderParentOrgId

`func (o *BiomarkerAssayProvidersRow) GetProviderParentOrgId() string`

GetProviderParentOrgId returns the ProviderParentOrgId field if non-nil, zero value otherwise.

### GetProviderParentOrgIdOk

`func (o *BiomarkerAssayProvidersRow) GetProviderParentOrgIdOk() (*string, bool)`

GetProviderParentOrgIdOk returns a tuple with the ProviderParentOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderParentOrgId

`func (o *BiomarkerAssayProvidersRow) SetProviderParentOrgId(v string)`

SetProviderParentOrgId sets ProviderParentOrgId field to given value.

### HasProviderParentOrgId

`func (o *BiomarkerAssayProvidersRow) HasProviderParentOrgId() bool`

HasProviderParentOrgId returns a boolean if a field has been set.

### SetProviderParentOrgIdNil

`func (o *BiomarkerAssayProvidersRow) SetProviderParentOrgIdNil(b bool)`

 SetProviderParentOrgIdNil sets the value for ProviderParentOrgId to be an explicit nil

### UnsetProviderParentOrgId
`func (o *BiomarkerAssayProvidersRow) UnsetProviderParentOrgId()`

UnsetProviderParentOrgId ensures that no value is present for ProviderParentOrgId, not even an explicit nil
### GetProviderCompanyId

`func (o *BiomarkerAssayProvidersRow) GetProviderCompanyId() string`

GetProviderCompanyId returns the ProviderCompanyId field if non-nil, zero value otherwise.

### GetProviderCompanyIdOk

`func (o *BiomarkerAssayProvidersRow) GetProviderCompanyIdOk() (*string, bool)`

GetProviderCompanyIdOk returns a tuple with the ProviderCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCompanyId

`func (o *BiomarkerAssayProvidersRow) SetProviderCompanyId(v string)`

SetProviderCompanyId sets ProviderCompanyId field to given value.

### HasProviderCompanyId

`func (o *BiomarkerAssayProvidersRow) HasProviderCompanyId() bool`

HasProviderCompanyId returns a boolean if a field has been set.

### SetProviderCompanyIdNil

`func (o *BiomarkerAssayProvidersRow) SetProviderCompanyIdNil(b bool)`

 SetProviderCompanyIdNil sets the value for ProviderCompanyId to be an explicit nil

### UnsetProviderCompanyId
`func (o *BiomarkerAssayProvidersRow) UnsetProviderCompanyId()`

UnsetProviderCompanyId ensures that no value is present for ProviderCompanyId, not even an explicit nil
### GetProviderCompanyName

`func (o *BiomarkerAssayProvidersRow) GetProviderCompanyName() string`

GetProviderCompanyName returns the ProviderCompanyName field if non-nil, zero value otherwise.

### GetProviderCompanyNameOk

`func (o *BiomarkerAssayProvidersRow) GetProviderCompanyNameOk() (*string, bool)`

GetProviderCompanyNameOk returns a tuple with the ProviderCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCompanyName

`func (o *BiomarkerAssayProvidersRow) SetProviderCompanyName(v string)`

SetProviderCompanyName sets ProviderCompanyName field to given value.

### HasProviderCompanyName

`func (o *BiomarkerAssayProvidersRow) HasProviderCompanyName() bool`

HasProviderCompanyName returns a boolean if a field has been set.

### SetProviderCompanyNameNil

`func (o *BiomarkerAssayProvidersRow) SetProviderCompanyNameNil(b bool)`

 SetProviderCompanyNameNil sets the value for ProviderCompanyName to be an explicit nil

### UnsetProviderCompanyName
`func (o *BiomarkerAssayProvidersRow) UnsetProviderCompanyName()`

UnsetProviderCompanyName ensures that no value is present for ProviderCompanyName, not even an explicit nil
### GetProviderGroupId

`func (o *BiomarkerAssayProvidersRow) GetProviderGroupId() string`

GetProviderGroupId returns the ProviderGroupId field if non-nil, zero value otherwise.

### GetProviderGroupIdOk

`func (o *BiomarkerAssayProvidersRow) GetProviderGroupIdOk() (*string, bool)`

GetProviderGroupIdOk returns a tuple with the ProviderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderGroupId

`func (o *BiomarkerAssayProvidersRow) SetProviderGroupId(v string)`

SetProviderGroupId sets ProviderGroupId field to given value.

### HasProviderGroupId

`func (o *BiomarkerAssayProvidersRow) HasProviderGroupId() bool`

HasProviderGroupId returns a boolean if a field has been set.

### SetProviderGroupIdNil

`func (o *BiomarkerAssayProvidersRow) SetProviderGroupIdNil(b bool)`

 SetProviderGroupIdNil sets the value for ProviderGroupId to be an explicit nil

### UnsetProviderGroupId
`func (o *BiomarkerAssayProvidersRow) UnsetProviderGroupId()`

UnsetProviderGroupId ensures that no value is present for ProviderGroupId, not even an explicit nil
### GetProviderGroupName

`func (o *BiomarkerAssayProvidersRow) GetProviderGroupName() string`

GetProviderGroupName returns the ProviderGroupName field if non-nil, zero value otherwise.

### GetProviderGroupNameOk

`func (o *BiomarkerAssayProvidersRow) GetProviderGroupNameOk() (*string, bool)`

GetProviderGroupNameOk returns a tuple with the ProviderGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderGroupName

`func (o *BiomarkerAssayProvidersRow) SetProviderGroupName(v string)`

SetProviderGroupName sets ProviderGroupName field to given value.

### HasProviderGroupName

`func (o *BiomarkerAssayProvidersRow) HasProviderGroupName() bool`

HasProviderGroupName returns a boolean if a field has been set.

### SetProviderGroupNameNil

`func (o *BiomarkerAssayProvidersRow) SetProviderGroupNameNil(b bool)`

 SetProviderGroupNameNil sets the value for ProviderGroupName to be an explicit nil

### UnsetProviderGroupName
`func (o *BiomarkerAssayProvidersRow) UnsetProviderGroupName()`

UnsetProviderGroupName ensures that no value is present for ProviderGroupName, not even an explicit nil
### GetRxMatched

`func (o *BiomarkerAssayProvidersRow) GetRxMatched() bool`

GetRxMatched returns the RxMatched field if non-nil, zero value otherwise.

### GetRxMatchedOk

`func (o *BiomarkerAssayProvidersRow) GetRxMatchedOk() (*bool, bool)`

GetRxMatchedOk returns a tuple with the RxMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMatched

`func (o *BiomarkerAssayProvidersRow) SetRxMatched(v bool)`

SetRxMatched sets RxMatched field to given value.

### HasRxMatched

`func (o *BiomarkerAssayProvidersRow) HasRxMatched() bool`

HasRxMatched returns a boolean if a field has been set.

### SetRxMatchedNil

`func (o *BiomarkerAssayProvidersRow) SetRxMatchedNil(b bool)`

 SetRxMatchedNil sets the value for RxMatched to be an explicit nil

### UnsetRxMatched
`func (o *BiomarkerAssayProvidersRow) UnsetRxMatched()`

UnsetRxMatched ensures that no value is present for RxMatched, not even an explicit nil
### GetRxProviderId

`func (o *BiomarkerAssayProvidersRow) GetRxProviderId() string`

GetRxProviderId returns the RxProviderId field if non-nil, zero value otherwise.

### GetRxProviderIdOk

`func (o *BiomarkerAssayProvidersRow) GetRxProviderIdOk() (*string, bool)`

GetRxProviderIdOk returns a tuple with the RxProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxProviderId

`func (o *BiomarkerAssayProvidersRow) SetRxProviderId(v string)`

SetRxProviderId sets RxProviderId field to given value.

### HasRxProviderId

`func (o *BiomarkerAssayProvidersRow) HasRxProviderId() bool`

HasRxProviderId returns a boolean if a field has been set.

### SetRxProviderIdNil

`func (o *BiomarkerAssayProvidersRow) SetRxProviderIdNil(b bool)`

 SetRxProviderIdNil sets the value for RxProviderId to be an explicit nil

### UnsetRxProviderId
`func (o *BiomarkerAssayProvidersRow) UnsetRxProviderId()`

UnsetRxProviderId ensures that no value is present for RxProviderId, not even an explicit nil
### GetRxProviderName

`func (o *BiomarkerAssayProvidersRow) GetRxProviderName() string`

GetRxProviderName returns the RxProviderName field if non-nil, zero value otherwise.

### GetRxProviderNameOk

`func (o *BiomarkerAssayProvidersRow) GetRxProviderNameOk() (*string, bool)`

GetRxProviderNameOk returns a tuple with the RxProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxProviderName

`func (o *BiomarkerAssayProvidersRow) SetRxProviderName(v string)`

SetRxProviderName sets RxProviderName field to given value.

### HasRxProviderName

`func (o *BiomarkerAssayProvidersRow) HasRxProviderName() bool`

HasRxProviderName returns a boolean if a field has been set.

### SetRxProviderNameNil

`func (o *BiomarkerAssayProvidersRow) SetRxProviderNameNil(b bool)`

 SetRxProviderNameNil sets the value for RxProviderName to be an explicit nil

### UnsetRxProviderName
`func (o *BiomarkerAssayProvidersRow) UnsetRxProviderName()`

UnsetRxProviderName ensures that no value is present for RxProviderName, not even an explicit nil
### GetRxProviderSlug

`func (o *BiomarkerAssayProvidersRow) GetRxProviderSlug() string`

GetRxProviderSlug returns the RxProviderSlug field if non-nil, zero value otherwise.

### GetRxProviderSlugOk

`func (o *BiomarkerAssayProvidersRow) GetRxProviderSlugOk() (*string, bool)`

GetRxProviderSlugOk returns a tuple with the RxProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxProviderSlug

`func (o *BiomarkerAssayProvidersRow) SetRxProviderSlug(v string)`

SetRxProviderSlug sets RxProviderSlug field to given value.

### HasRxProviderSlug

`func (o *BiomarkerAssayProvidersRow) HasRxProviderSlug() bool`

HasRxProviderSlug returns a boolean if a field has been set.

### SetRxProviderSlugNil

`func (o *BiomarkerAssayProvidersRow) SetRxProviderSlugNil(b bool)`

 SetRxProviderSlugNil sets the value for RxProviderSlug to be an explicit nil

### UnsetRxProviderSlug
`func (o *BiomarkerAssayProvidersRow) UnsetRxProviderSlug()`

UnsetRxProviderSlug ensures that no value is present for RxProviderSlug, not even an explicit nil
### GetHasClinicalValidation

`func (o *BiomarkerAssayProvidersRow) GetHasClinicalValidation() bool`

GetHasClinicalValidation returns the HasClinicalValidation field if non-nil, zero value otherwise.

### GetHasClinicalValidationOk

`func (o *BiomarkerAssayProvidersRow) GetHasClinicalValidationOk() (*bool, bool)`

GetHasClinicalValidationOk returns a tuple with the HasClinicalValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClinicalValidation

`func (o *BiomarkerAssayProvidersRow) SetHasClinicalValidation(v bool)`

SetHasClinicalValidation sets HasClinicalValidation field to given value.

### HasHasClinicalValidation

`func (o *BiomarkerAssayProvidersRow) HasHasClinicalValidation() bool`

HasHasClinicalValidation returns a boolean if a field has been set.

### SetHasClinicalValidationNil

`func (o *BiomarkerAssayProvidersRow) SetHasClinicalValidationNil(b bool)`

 SetHasClinicalValidationNil sets the value for HasClinicalValidation to be an explicit nil

### UnsetHasClinicalValidation
`func (o *BiomarkerAssayProvidersRow) UnsetHasClinicalValidation()`

UnsetHasClinicalValidation ensures that no value is present for HasClinicalValidation, not even an explicit nil
### GetHasAnalyticalValidation

`func (o *BiomarkerAssayProvidersRow) GetHasAnalyticalValidation() bool`

GetHasAnalyticalValidation returns the HasAnalyticalValidation field if non-nil, zero value otherwise.

### GetHasAnalyticalValidationOk

`func (o *BiomarkerAssayProvidersRow) GetHasAnalyticalValidationOk() (*bool, bool)`

GetHasAnalyticalValidationOk returns a tuple with the HasAnalyticalValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnalyticalValidation

`func (o *BiomarkerAssayProvidersRow) SetHasAnalyticalValidation(v bool)`

SetHasAnalyticalValidation sets HasAnalyticalValidation field to given value.

### HasHasAnalyticalValidation

`func (o *BiomarkerAssayProvidersRow) HasHasAnalyticalValidation() bool`

HasHasAnalyticalValidation returns a boolean if a field has been set.

### SetHasAnalyticalValidationNil

`func (o *BiomarkerAssayProvidersRow) SetHasAnalyticalValidationNil(b bool)`

 SetHasAnalyticalValidationNil sets the value for HasAnalyticalValidation to be an explicit nil

### UnsetHasAnalyticalValidation
`func (o *BiomarkerAssayProvidersRow) UnsetHasAnalyticalValidation()`

UnsetHasAnalyticalValidation ensures that no value is present for HasAnalyticalValidation, not even an explicit nil
### GetNPapers

`func (o *BiomarkerAssayProvidersRow) GetNPapers() int64`

GetNPapers returns the NPapers field if non-nil, zero value otherwise.

### GetNPapersOk

`func (o *BiomarkerAssayProvidersRow) GetNPapersOk() (*int64, bool)`

GetNPapersOk returns a tuple with the NPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPapers

`func (o *BiomarkerAssayProvidersRow) SetNPapers(v int64)`

SetNPapers sets NPapers field to given value.

### HasNPapers

`func (o *BiomarkerAssayProvidersRow) HasNPapers() bool`

HasNPapers returns a boolean if a field has been set.

### SetNPapersNil

`func (o *BiomarkerAssayProvidersRow) SetNPapersNil(b bool)`

 SetNPapersNil sets the value for NPapers to be an explicit nil

### UnsetNPapers
`func (o *BiomarkerAssayProvidersRow) UnsetNPapers()`

UnsetNPapers ensures that no value is present for NPapers, not even an explicit nil
### GetNDatedPapers

`func (o *BiomarkerAssayProvidersRow) GetNDatedPapers() int64`

GetNDatedPapers returns the NDatedPapers field if non-nil, zero value otherwise.

### GetNDatedPapersOk

`func (o *BiomarkerAssayProvidersRow) GetNDatedPapersOk() (*int64, bool)`

GetNDatedPapersOk returns a tuple with the NDatedPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDatedPapers

`func (o *BiomarkerAssayProvidersRow) SetNDatedPapers(v int64)`

SetNDatedPapers sets NDatedPapers field to given value.

### HasNDatedPapers

`func (o *BiomarkerAssayProvidersRow) HasNDatedPapers() bool`

HasNDatedPapers returns a boolean if a field has been set.

### SetNDatedPapersNil

`func (o *BiomarkerAssayProvidersRow) SetNDatedPapersNil(b bool)`

 SetNDatedPapersNil sets the value for NDatedPapers to be an explicit nil

### UnsetNDatedPapers
`func (o *BiomarkerAssayProvidersRow) UnsetNDatedPapers()`

UnsetNDatedPapers ensures that no value is present for NDatedPapers, not even an explicit nil
### GetNUndatedPapers

`func (o *BiomarkerAssayProvidersRow) GetNUndatedPapers() int64`

GetNUndatedPapers returns the NUndatedPapers field if non-nil, zero value otherwise.

### GetNUndatedPapersOk

`func (o *BiomarkerAssayProvidersRow) GetNUndatedPapersOk() (*int64, bool)`

GetNUndatedPapersOk returns a tuple with the NUndatedPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNUndatedPapers

`func (o *BiomarkerAssayProvidersRow) SetNUndatedPapers(v int64)`

SetNUndatedPapers sets NUndatedPapers field to given value.

### HasNUndatedPapers

`func (o *BiomarkerAssayProvidersRow) HasNUndatedPapers() bool`

HasNUndatedPapers returns a boolean if a field has been set.

### SetNUndatedPapersNil

`func (o *BiomarkerAssayProvidersRow) SetNUndatedPapersNil(b bool)`

 SetNUndatedPapersNil sets the value for NUndatedPapers to be an explicit nil

### UnsetNUndatedPapers
`func (o *BiomarkerAssayProvidersRow) UnsetNUndatedPapers()`

UnsetNUndatedPapers ensures that no value is present for NUndatedPapers, not even an explicit nil
### GetMostRecentPublishedAt

`func (o *BiomarkerAssayProvidersRow) GetMostRecentPublishedAt() time.Time`

GetMostRecentPublishedAt returns the MostRecentPublishedAt field if non-nil, zero value otherwise.

### GetMostRecentPublishedAtOk

`func (o *BiomarkerAssayProvidersRow) GetMostRecentPublishedAtOk() (*time.Time, bool)`

GetMostRecentPublishedAtOk returns a tuple with the MostRecentPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentPublishedAt

`func (o *BiomarkerAssayProvidersRow) SetMostRecentPublishedAt(v time.Time)`

SetMostRecentPublishedAt sets MostRecentPublishedAt field to given value.

### HasMostRecentPublishedAt

`func (o *BiomarkerAssayProvidersRow) HasMostRecentPublishedAt() bool`

HasMostRecentPublishedAt returns a boolean if a field has been set.

### SetMostRecentPublishedAtNil

`func (o *BiomarkerAssayProvidersRow) SetMostRecentPublishedAtNil(b bool)`

 SetMostRecentPublishedAtNil sets the value for MostRecentPublishedAt to be an explicit nil

### UnsetMostRecentPublishedAt
`func (o *BiomarkerAssayProvidersRow) UnsetMostRecentPublishedAt()`

UnsetMostRecentPublishedAt ensures that no value is present for MostRecentPublishedAt, not even an explicit nil
### GetMostRecentPmid

`func (o *BiomarkerAssayProvidersRow) GetMostRecentPmid() string`

GetMostRecentPmid returns the MostRecentPmid field if non-nil, zero value otherwise.

### GetMostRecentPmidOk

`func (o *BiomarkerAssayProvidersRow) GetMostRecentPmidOk() (*string, bool)`

GetMostRecentPmidOk returns a tuple with the MostRecentPmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentPmid

`func (o *BiomarkerAssayProvidersRow) SetMostRecentPmid(v string)`

SetMostRecentPmid sets MostRecentPmid field to given value.

### HasMostRecentPmid

`func (o *BiomarkerAssayProvidersRow) HasMostRecentPmid() bool`

HasMostRecentPmid returns a boolean if a field has been set.

### SetMostRecentPmidNil

`func (o *BiomarkerAssayProvidersRow) SetMostRecentPmidNil(b bool)`

 SetMostRecentPmidNil sets the value for MostRecentPmid to be an explicit nil

### UnsetMostRecentPmid
`func (o *BiomarkerAssayProvidersRow) UnsetMostRecentPmid()`

UnsetMostRecentPmid ensures that no value is present for MostRecentPmid, not even an explicit nil
### GetMostRecentTitle

`func (o *BiomarkerAssayProvidersRow) GetMostRecentTitle() string`

GetMostRecentTitle returns the MostRecentTitle field if non-nil, zero value otherwise.

### GetMostRecentTitleOk

`func (o *BiomarkerAssayProvidersRow) GetMostRecentTitleOk() (*string, bool)`

GetMostRecentTitleOk returns a tuple with the MostRecentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentTitle

`func (o *BiomarkerAssayProvidersRow) SetMostRecentTitle(v string)`

SetMostRecentTitle sets MostRecentTitle field to given value.

### HasMostRecentTitle

`func (o *BiomarkerAssayProvidersRow) HasMostRecentTitle() bool`

HasMostRecentTitle returns a boolean if a field has been set.

### SetMostRecentTitleNil

`func (o *BiomarkerAssayProvidersRow) SetMostRecentTitleNil(b bool)`

 SetMostRecentTitleNil sets the value for MostRecentTitle to be an explicit nil

### UnsetMostRecentTitle
`func (o *BiomarkerAssayProvidersRow) UnsetMostRecentTitle()`

UnsetMostRecentTitle ensures that no value is present for MostRecentTitle, not even an explicit nil
### GetMeetsCuratedThreshold

`func (o *BiomarkerAssayProvidersRow) GetMeetsCuratedThreshold() bool`

GetMeetsCuratedThreshold returns the MeetsCuratedThreshold field if non-nil, zero value otherwise.

### GetMeetsCuratedThresholdOk

`func (o *BiomarkerAssayProvidersRow) GetMeetsCuratedThresholdOk() (*bool, bool)`

GetMeetsCuratedThresholdOk returns a tuple with the MeetsCuratedThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetsCuratedThreshold

`func (o *BiomarkerAssayProvidersRow) SetMeetsCuratedThreshold(v bool)`

SetMeetsCuratedThreshold sets MeetsCuratedThreshold field to given value.

### HasMeetsCuratedThreshold

`func (o *BiomarkerAssayProvidersRow) HasMeetsCuratedThreshold() bool`

HasMeetsCuratedThreshold returns a boolean if a field has been set.

### SetMeetsCuratedThresholdNil

`func (o *BiomarkerAssayProvidersRow) SetMeetsCuratedThresholdNil(b bool)`

 SetMeetsCuratedThresholdNil sets the value for MeetsCuratedThreshold to be an explicit nil

### UnsetMeetsCuratedThreshold
`func (o *BiomarkerAssayProvidersRow) UnsetMeetsCuratedThreshold()`

UnsetMeetsCuratedThreshold ensures that no value is present for MeetsCuratedThreshold, not even an explicit nil
### GetRegulatoryEvidenceKind

`func (o *BiomarkerAssayProvidersRow) GetRegulatoryEvidenceKind() string`

GetRegulatoryEvidenceKind returns the RegulatoryEvidenceKind field if non-nil, zero value otherwise.

### GetRegulatoryEvidenceKindOk

`func (o *BiomarkerAssayProvidersRow) GetRegulatoryEvidenceKindOk() (*string, bool)`

GetRegulatoryEvidenceKindOk returns a tuple with the RegulatoryEvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryEvidenceKind

`func (o *BiomarkerAssayProvidersRow) SetRegulatoryEvidenceKind(v string)`

SetRegulatoryEvidenceKind sets RegulatoryEvidenceKind field to given value.

### HasRegulatoryEvidenceKind

`func (o *BiomarkerAssayProvidersRow) HasRegulatoryEvidenceKind() bool`

HasRegulatoryEvidenceKind returns a boolean if a field has been set.

### SetRegulatoryEvidenceKindNil

`func (o *BiomarkerAssayProvidersRow) SetRegulatoryEvidenceKindNil(b bool)`

 SetRegulatoryEvidenceKindNil sets the value for RegulatoryEvidenceKind to be an explicit nil

### UnsetRegulatoryEvidenceKind
`func (o *BiomarkerAssayProvidersRow) UnsetRegulatoryEvidenceKind()`

UnsetRegulatoryEvidenceKind ensures that no value is present for RegulatoryEvidenceKind, not even an explicit nil
### GetHasFdaRecordForBiomarker

`func (o *BiomarkerAssayProvidersRow) GetHasFdaRecordForBiomarker() bool`

GetHasFdaRecordForBiomarker returns the HasFdaRecordForBiomarker field if non-nil, zero value otherwise.

### GetHasFdaRecordForBiomarkerOk

`func (o *BiomarkerAssayProvidersRow) GetHasFdaRecordForBiomarkerOk() (*bool, bool)`

GetHasFdaRecordForBiomarkerOk returns a tuple with the HasFdaRecordForBiomarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFdaRecordForBiomarker

`func (o *BiomarkerAssayProvidersRow) SetHasFdaRecordForBiomarker(v bool)`

SetHasFdaRecordForBiomarker sets HasFdaRecordForBiomarker field to given value.

### HasHasFdaRecordForBiomarker

`func (o *BiomarkerAssayProvidersRow) HasHasFdaRecordForBiomarker() bool`

HasHasFdaRecordForBiomarker returns a boolean if a field has been set.

### SetHasFdaRecordForBiomarkerNil

`func (o *BiomarkerAssayProvidersRow) SetHasFdaRecordForBiomarkerNil(b bool)`

 SetHasFdaRecordForBiomarkerNil sets the value for HasFdaRecordForBiomarker to be an explicit nil

### UnsetHasFdaRecordForBiomarker
`func (o *BiomarkerAssayProvidersRow) UnsetHasFdaRecordForBiomarker()`

UnsetHasFdaRecordForBiomarker ensures that no value is present for HasFdaRecordForBiomarker, not even an explicit nil
### GetRegulatoryEvidenceReasons

`func (o *BiomarkerAssayProvidersRow) GetRegulatoryEvidenceReasons() []string`

GetRegulatoryEvidenceReasons returns the RegulatoryEvidenceReasons field if non-nil, zero value otherwise.

### GetRegulatoryEvidenceReasonsOk

`func (o *BiomarkerAssayProvidersRow) GetRegulatoryEvidenceReasonsOk() (*[]string, bool)`

GetRegulatoryEvidenceReasonsOk returns a tuple with the RegulatoryEvidenceReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryEvidenceReasons

`func (o *BiomarkerAssayProvidersRow) SetRegulatoryEvidenceReasons(v []string)`

SetRegulatoryEvidenceReasons sets RegulatoryEvidenceReasons field to given value.

### HasRegulatoryEvidenceReasons

`func (o *BiomarkerAssayProvidersRow) HasRegulatoryEvidenceReasons() bool`

HasRegulatoryEvidenceReasons returns a boolean if a field has been set.

### SetRegulatoryEvidenceReasonsNil

`func (o *BiomarkerAssayProvidersRow) SetRegulatoryEvidenceReasonsNil(b bool)`

 SetRegulatoryEvidenceReasonsNil sets the value for RegulatoryEvidenceReasons to be an explicit nil

### UnsetRegulatoryEvidenceReasons
`func (o *BiomarkerAssayProvidersRow) UnsetRegulatoryEvidenceReasons()`

UnsetRegulatoryEvidenceReasons ensures that no value is present for RegulatoryEvidenceReasons, not even an explicit nil
### GetNFdaIvd510k

`func (o *BiomarkerAssayProvidersRow) GetNFdaIvd510k() int64`

GetNFdaIvd510k returns the NFdaIvd510k field if non-nil, zero value otherwise.

### GetNFdaIvd510kOk

`func (o *BiomarkerAssayProvidersRow) GetNFdaIvd510kOk() (*int64, bool)`

GetNFdaIvd510kOk returns a tuple with the NFdaIvd510k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaIvd510k

`func (o *BiomarkerAssayProvidersRow) SetNFdaIvd510k(v int64)`

SetNFdaIvd510k sets NFdaIvd510k field to given value.

### HasNFdaIvd510k

`func (o *BiomarkerAssayProvidersRow) HasNFdaIvd510k() bool`

HasNFdaIvd510k returns a boolean if a field has been set.

### SetNFdaIvd510kNil

`func (o *BiomarkerAssayProvidersRow) SetNFdaIvd510kNil(b bool)`

 SetNFdaIvd510kNil sets the value for NFdaIvd510k to be an explicit nil

### UnsetNFdaIvd510k
`func (o *BiomarkerAssayProvidersRow) UnsetNFdaIvd510k()`

UnsetNFdaIvd510k ensures that no value is present for NFdaIvd510k, not even an explicit nil
### GetNFdaIvdPma

`func (o *BiomarkerAssayProvidersRow) GetNFdaIvdPma() int64`

GetNFdaIvdPma returns the NFdaIvdPma field if non-nil, zero value otherwise.

### GetNFdaIvdPmaOk

`func (o *BiomarkerAssayProvidersRow) GetNFdaIvdPmaOk() (*int64, bool)`

GetNFdaIvdPmaOk returns a tuple with the NFdaIvdPma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaIvdPma

`func (o *BiomarkerAssayProvidersRow) SetNFdaIvdPma(v int64)`

SetNFdaIvdPma sets NFdaIvdPma field to given value.

### HasNFdaIvdPma

`func (o *BiomarkerAssayProvidersRow) HasNFdaIvdPma() bool`

HasNFdaIvdPma returns a boolean if a field has been set.

### SetNFdaIvdPmaNil

`func (o *BiomarkerAssayProvidersRow) SetNFdaIvdPmaNil(b bool)`

 SetNFdaIvdPmaNil sets the value for NFdaIvdPma to be an explicit nil

### UnsetNFdaIvdPma
`func (o *BiomarkerAssayProvidersRow) UnsetNFdaIvdPma()`

UnsetNFdaIvdPma ensures that no value is present for NFdaIvdPma, not even an explicit nil
### GetFdaCompanyMatchRung

`func (o *BiomarkerAssayProvidersRow) GetFdaCompanyMatchRung() string`

GetFdaCompanyMatchRung returns the FdaCompanyMatchRung field if non-nil, zero value otherwise.

### GetFdaCompanyMatchRungOk

`func (o *BiomarkerAssayProvidersRow) GetFdaCompanyMatchRungOk() (*string, bool)`

GetFdaCompanyMatchRungOk returns a tuple with the FdaCompanyMatchRung field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdaCompanyMatchRung

`func (o *BiomarkerAssayProvidersRow) SetFdaCompanyMatchRung(v string)`

SetFdaCompanyMatchRung sets FdaCompanyMatchRung field to given value.

### HasFdaCompanyMatchRung

`func (o *BiomarkerAssayProvidersRow) HasFdaCompanyMatchRung() bool`

HasFdaCompanyMatchRung returns a boolean if a field has been set.

### SetFdaCompanyMatchRungNil

`func (o *BiomarkerAssayProvidersRow) SetFdaCompanyMatchRungNil(b bool)`

 SetFdaCompanyMatchRungNil sets the value for FdaCompanyMatchRung to be an explicit nil

### UnsetFdaCompanyMatchRung
`func (o *BiomarkerAssayProvidersRow) UnsetFdaCompanyMatchRung()`

UnsetFdaCompanyMatchRung ensures that no value is present for FdaCompanyMatchRung, not even an explicit nil
### GetNFdaAnalyte510k

`func (o *BiomarkerAssayProvidersRow) GetNFdaAnalyte510k() int64`

GetNFdaAnalyte510k returns the NFdaAnalyte510k field if non-nil, zero value otherwise.

### GetNFdaAnalyte510kOk

`func (o *BiomarkerAssayProvidersRow) GetNFdaAnalyte510kOk() (*int64, bool)`

GetNFdaAnalyte510kOk returns a tuple with the NFdaAnalyte510k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaAnalyte510k

`func (o *BiomarkerAssayProvidersRow) SetNFdaAnalyte510k(v int64)`

SetNFdaAnalyte510k sets NFdaAnalyte510k field to given value.

### HasNFdaAnalyte510k

`func (o *BiomarkerAssayProvidersRow) HasNFdaAnalyte510k() bool`

HasNFdaAnalyte510k returns a boolean if a field has been set.

### SetNFdaAnalyte510kNil

`func (o *BiomarkerAssayProvidersRow) SetNFdaAnalyte510kNil(b bool)`

 SetNFdaAnalyte510kNil sets the value for NFdaAnalyte510k to be an explicit nil

### UnsetNFdaAnalyte510k
`func (o *BiomarkerAssayProvidersRow) UnsetNFdaAnalyte510k()`

UnsetNFdaAnalyte510k ensures that no value is present for NFdaAnalyte510k, not even an explicit nil
### GetNFdaAnalytePma

`func (o *BiomarkerAssayProvidersRow) GetNFdaAnalytePma() int64`

GetNFdaAnalytePma returns the NFdaAnalytePma field if non-nil, zero value otherwise.

### GetNFdaAnalytePmaOk

`func (o *BiomarkerAssayProvidersRow) GetNFdaAnalytePmaOk() (*int64, bool)`

GetNFdaAnalytePmaOk returns a tuple with the NFdaAnalytePma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaAnalytePma

`func (o *BiomarkerAssayProvidersRow) SetNFdaAnalytePma(v int64)`

SetNFdaAnalytePma sets NFdaAnalytePma field to given value.

### HasNFdaAnalytePma

`func (o *BiomarkerAssayProvidersRow) HasNFdaAnalytePma() bool`

HasNFdaAnalytePma returns a boolean if a field has been set.

### SetNFdaAnalytePmaNil

`func (o *BiomarkerAssayProvidersRow) SetNFdaAnalytePmaNil(b bool)`

 SetNFdaAnalytePmaNil sets the value for NFdaAnalytePma to be an explicit nil

### UnsetNFdaAnalytePma
`func (o *BiomarkerAssayProvidersRow) UnsetNFdaAnalytePma()`

UnsetNFdaAnalytePma ensures that no value is present for NFdaAnalytePma, not even an explicit nil
### GetFdaAnalyteProductCodes

`func (o *BiomarkerAssayProvidersRow) GetFdaAnalyteProductCodes() []string`

GetFdaAnalyteProductCodes returns the FdaAnalyteProductCodes field if non-nil, zero value otherwise.

### GetFdaAnalyteProductCodesOk

`func (o *BiomarkerAssayProvidersRow) GetFdaAnalyteProductCodesOk() (*[]string, bool)`

GetFdaAnalyteProductCodesOk returns a tuple with the FdaAnalyteProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdaAnalyteProductCodes

`func (o *BiomarkerAssayProvidersRow) SetFdaAnalyteProductCodes(v []string)`

SetFdaAnalyteProductCodes sets FdaAnalyteProductCodes field to given value.

### HasFdaAnalyteProductCodes

`func (o *BiomarkerAssayProvidersRow) HasFdaAnalyteProductCodes() bool`

HasFdaAnalyteProductCodes returns a boolean if a field has been set.

### SetFdaAnalyteProductCodesNil

`func (o *BiomarkerAssayProvidersRow) SetFdaAnalyteProductCodesNil(b bool)`

 SetFdaAnalyteProductCodesNil sets the value for FdaAnalyteProductCodes to be an explicit nil

### UnsetFdaAnalyteProductCodes
`func (o *BiomarkerAssayProvidersRow) UnsetFdaAnalyteProductCodes()`

UnsetFdaAnalyteProductCodes ensures that no value is present for FdaAnalyteProductCodes, not even an explicit nil
### GetFdaAnalyteFdaUrl

`func (o *BiomarkerAssayProvidersRow) GetFdaAnalyteFdaUrl() string`

GetFdaAnalyteFdaUrl returns the FdaAnalyteFdaUrl field if non-nil, zero value otherwise.

### GetFdaAnalyteFdaUrlOk

`func (o *BiomarkerAssayProvidersRow) GetFdaAnalyteFdaUrlOk() (*string, bool)`

GetFdaAnalyteFdaUrlOk returns a tuple with the FdaAnalyteFdaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdaAnalyteFdaUrl

`func (o *BiomarkerAssayProvidersRow) SetFdaAnalyteFdaUrl(v string)`

SetFdaAnalyteFdaUrl sets FdaAnalyteFdaUrl field to given value.

### HasFdaAnalyteFdaUrl

`func (o *BiomarkerAssayProvidersRow) HasFdaAnalyteFdaUrl() bool`

HasFdaAnalyteFdaUrl returns a boolean if a field has been set.

### SetFdaAnalyteFdaUrlNil

`func (o *BiomarkerAssayProvidersRow) SetFdaAnalyteFdaUrlNil(b bool)`

 SetFdaAnalyteFdaUrlNil sets the value for FdaAnalyteFdaUrl to be an explicit nil

### UnsetFdaAnalyteFdaUrl
`func (o *BiomarkerAssayProvidersRow) UnsetFdaAnalyteFdaUrl()`

UnsetFdaAnalyteFdaUrl ensures that no value is present for FdaAnalyteFdaUrl, not even an explicit nil
### GetSpecificAssayMatched

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayMatched() bool`

GetSpecificAssayMatched returns the SpecificAssayMatched field if non-nil, zero value otherwise.

### GetSpecificAssayMatchedOk

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayMatchedOk() (*bool, bool)`

GetSpecificAssayMatchedOk returns a tuple with the SpecificAssayMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificAssayMatched

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayMatched(v bool)`

SetSpecificAssayMatched sets SpecificAssayMatched field to given value.

### HasSpecificAssayMatched

`func (o *BiomarkerAssayProvidersRow) HasSpecificAssayMatched() bool`

HasSpecificAssayMatched returns a boolean if a field has been set.

### SetSpecificAssayMatchedNil

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayMatchedNil(b bool)`

 SetSpecificAssayMatchedNil sets the value for SpecificAssayMatched to be an explicit nil

### UnsetSpecificAssayMatched
`func (o *BiomarkerAssayProvidersRow) UnsetSpecificAssayMatched()`

UnsetSpecificAssayMatched ensures that no value is present for SpecificAssayMatched, not even an explicit nil
### GetSpecificAssayDeviceName

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayDeviceName() string`

GetSpecificAssayDeviceName returns the SpecificAssayDeviceName field if non-nil, zero value otherwise.

### GetSpecificAssayDeviceNameOk

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayDeviceNameOk() (*string, bool)`

GetSpecificAssayDeviceNameOk returns a tuple with the SpecificAssayDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificAssayDeviceName

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayDeviceName(v string)`

SetSpecificAssayDeviceName sets SpecificAssayDeviceName field to given value.

### HasSpecificAssayDeviceName

`func (o *BiomarkerAssayProvidersRow) HasSpecificAssayDeviceName() bool`

HasSpecificAssayDeviceName returns a boolean if a field has been set.

### SetSpecificAssayDeviceNameNil

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayDeviceNameNil(b bool)`

 SetSpecificAssayDeviceNameNil sets the value for SpecificAssayDeviceName to be an explicit nil

### UnsetSpecificAssayDeviceName
`func (o *BiomarkerAssayProvidersRow) UnsetSpecificAssayDeviceName()`

UnsetSpecificAssayDeviceName ensures that no value is present for SpecificAssayDeviceName, not even an explicit nil
### GetSpecificAssayKNumber

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayKNumber() string`

GetSpecificAssayKNumber returns the SpecificAssayKNumber field if non-nil, zero value otherwise.

### GetSpecificAssayKNumberOk

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayKNumberOk() (*string, bool)`

GetSpecificAssayKNumberOk returns a tuple with the SpecificAssayKNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificAssayKNumber

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayKNumber(v string)`

SetSpecificAssayKNumber sets SpecificAssayKNumber field to given value.

### HasSpecificAssayKNumber

`func (o *BiomarkerAssayProvidersRow) HasSpecificAssayKNumber() bool`

HasSpecificAssayKNumber returns a boolean if a field has been set.

### SetSpecificAssayKNumberNil

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayKNumberNil(b bool)`

 SetSpecificAssayKNumberNil sets the value for SpecificAssayKNumber to be an explicit nil

### UnsetSpecificAssayKNumber
`func (o *BiomarkerAssayProvidersRow) UnsetSpecificAssayKNumber()`

UnsetSpecificAssayKNumber ensures that no value is present for SpecificAssayKNumber, not even an explicit nil
### GetSpecificAssayPmaNumber

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayPmaNumber() string`

GetSpecificAssayPmaNumber returns the SpecificAssayPmaNumber field if non-nil, zero value otherwise.

### GetSpecificAssayPmaNumberOk

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayPmaNumberOk() (*string, bool)`

GetSpecificAssayPmaNumberOk returns a tuple with the SpecificAssayPmaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificAssayPmaNumber

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayPmaNumber(v string)`

SetSpecificAssayPmaNumber sets SpecificAssayPmaNumber field to given value.

### HasSpecificAssayPmaNumber

`func (o *BiomarkerAssayProvidersRow) HasSpecificAssayPmaNumber() bool`

HasSpecificAssayPmaNumber returns a boolean if a field has been set.

### SetSpecificAssayPmaNumberNil

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayPmaNumberNil(b bool)`

 SetSpecificAssayPmaNumberNil sets the value for SpecificAssayPmaNumber to be an explicit nil

### UnsetSpecificAssayPmaNumber
`func (o *BiomarkerAssayProvidersRow) UnsetSpecificAssayPmaNumber()`

UnsetSpecificAssayPmaNumber ensures that no value is present for SpecificAssayPmaNumber, not even an explicit nil
### GetSpecificAssayFdaUrl

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayFdaUrl() string`

GetSpecificAssayFdaUrl returns the SpecificAssayFdaUrl field if non-nil, zero value otherwise.

### GetSpecificAssayFdaUrlOk

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssayFdaUrlOk() (*string, bool)`

GetSpecificAssayFdaUrlOk returns a tuple with the SpecificAssayFdaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificAssayFdaUrl

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayFdaUrl(v string)`

SetSpecificAssayFdaUrl sets SpecificAssayFdaUrl field to given value.

### HasSpecificAssayFdaUrl

`func (o *BiomarkerAssayProvidersRow) HasSpecificAssayFdaUrl() bool`

HasSpecificAssayFdaUrl returns a boolean if a field has been set.

### SetSpecificAssayFdaUrlNil

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssayFdaUrlNil(b bool)`

 SetSpecificAssayFdaUrlNil sets the value for SpecificAssayFdaUrl to be an explicit nil

### UnsetSpecificAssayFdaUrl
`func (o *BiomarkerAssayProvidersRow) UnsetSpecificAssayFdaUrl()`

UnsetSpecificAssayFdaUrl ensures that no value is present for SpecificAssayFdaUrl, not even an explicit nil
### GetSpecificAssaySimilarity

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssaySimilarity() float32`

GetSpecificAssaySimilarity returns the SpecificAssaySimilarity field if non-nil, zero value otherwise.

### GetSpecificAssaySimilarityOk

`func (o *BiomarkerAssayProvidersRow) GetSpecificAssaySimilarityOk() (*float32, bool)`

GetSpecificAssaySimilarityOk returns a tuple with the SpecificAssaySimilarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificAssaySimilarity

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssaySimilarity(v float32)`

SetSpecificAssaySimilarity sets SpecificAssaySimilarity field to given value.

### HasSpecificAssaySimilarity

`func (o *BiomarkerAssayProvidersRow) HasSpecificAssaySimilarity() bool`

HasSpecificAssaySimilarity returns a boolean if a field has been set.

### SetSpecificAssaySimilarityNil

`func (o *BiomarkerAssayProvidersRow) SetSpecificAssaySimilarityNil(b bool)`

 SetSpecificAssaySimilarityNil sets the value for SpecificAssaySimilarity to be an explicit nil

### UnsetSpecificAssaySimilarity
`func (o *BiomarkerAssayProvidersRow) UnsetSpecificAssaySimilarity()`

UnsetSpecificAssaySimilarity ensures that no value is present for SpecificAssaySimilarity, not even an explicit nil
### GetNCdxEntries

`func (o *BiomarkerAssayProvidersRow) GetNCdxEntries() int32`

GetNCdxEntries returns the NCdxEntries field if non-nil, zero value otherwise.

### GetNCdxEntriesOk

`func (o *BiomarkerAssayProvidersRow) GetNCdxEntriesOk() (*int32, bool)`

GetNCdxEntriesOk returns a tuple with the NCdxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCdxEntries

`func (o *BiomarkerAssayProvidersRow) SetNCdxEntries(v int32)`

SetNCdxEntries sets NCdxEntries field to given value.

### HasNCdxEntries

`func (o *BiomarkerAssayProvidersRow) HasNCdxEntries() bool`

HasNCdxEntries returns a boolean if a field has been set.

### SetNCdxEntriesNil

`func (o *BiomarkerAssayProvidersRow) SetNCdxEntriesNil(b bool)`

 SetNCdxEntriesNil sets the value for NCdxEntries to be an explicit nil

### UnsetNCdxEntries
`func (o *BiomarkerAssayProvidersRow) UnsetNCdxEntries()`

UnsetNCdxEntries ensures that no value is present for NCdxEntries, not even an explicit nil
### GetCdxDiagnosticNames

`func (o *BiomarkerAssayProvidersRow) GetCdxDiagnosticNames() []string`

GetCdxDiagnosticNames returns the CdxDiagnosticNames field if non-nil, zero value otherwise.

### GetCdxDiagnosticNamesOk

`func (o *BiomarkerAssayProvidersRow) GetCdxDiagnosticNamesOk() (*[]string, bool)`

GetCdxDiagnosticNamesOk returns a tuple with the CdxDiagnosticNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdxDiagnosticNames

`func (o *BiomarkerAssayProvidersRow) SetCdxDiagnosticNames(v []string)`

SetCdxDiagnosticNames sets CdxDiagnosticNames field to given value.

### HasCdxDiagnosticNames

`func (o *BiomarkerAssayProvidersRow) HasCdxDiagnosticNames() bool`

HasCdxDiagnosticNames returns a boolean if a field has been set.

### SetCdxDiagnosticNamesNil

`func (o *BiomarkerAssayProvidersRow) SetCdxDiagnosticNamesNil(b bool)`

 SetCdxDiagnosticNamesNil sets the value for CdxDiagnosticNames to be an explicit nil

### UnsetCdxDiagnosticNames
`func (o *BiomarkerAssayProvidersRow) UnsetCdxDiagnosticNames()`

UnsetCdxDiagnosticNames ensures that no value is present for CdxDiagnosticNames, not even an explicit nil
### GetNCliaCertificates

`func (o *BiomarkerAssayProvidersRow) GetNCliaCertificates() int64`

GetNCliaCertificates returns the NCliaCertificates field if non-nil, zero value otherwise.

### GetNCliaCertificatesOk

`func (o *BiomarkerAssayProvidersRow) GetNCliaCertificatesOk() (*int64, bool)`

GetNCliaCertificatesOk returns a tuple with the NCliaCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCliaCertificates

`func (o *BiomarkerAssayProvidersRow) SetNCliaCertificates(v int64)`

SetNCliaCertificates sets NCliaCertificates field to given value.

### HasNCliaCertificates

`func (o *BiomarkerAssayProvidersRow) HasNCliaCertificates() bool`

HasNCliaCertificates returns a boolean if a field has been set.

### SetNCliaCertificatesNil

`func (o *BiomarkerAssayProvidersRow) SetNCliaCertificatesNil(b bool)`

 SetNCliaCertificatesNil sets the value for NCliaCertificates to be an explicit nil

### UnsetNCliaCertificates
`func (o *BiomarkerAssayProvidersRow) UnsetNCliaCertificates()`

UnsetNCliaCertificates ensures that no value is present for NCliaCertificates, not even an explicit nil
### GetCliaMatchRung

`func (o *BiomarkerAssayProvidersRow) GetCliaMatchRung() string`

GetCliaMatchRung returns the CliaMatchRung field if non-nil, zero value otherwise.

### GetCliaMatchRungOk

`func (o *BiomarkerAssayProvidersRow) GetCliaMatchRungOk() (*string, bool)`

GetCliaMatchRungOk returns a tuple with the CliaMatchRung field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliaMatchRung

`func (o *BiomarkerAssayProvidersRow) SetCliaMatchRung(v string)`

SetCliaMatchRung sets CliaMatchRung field to given value.

### HasCliaMatchRung

`func (o *BiomarkerAssayProvidersRow) HasCliaMatchRung() bool`

HasCliaMatchRung returns a boolean if a field has been set.

### SetCliaMatchRungNil

`func (o *BiomarkerAssayProvidersRow) SetCliaMatchRungNil(b bool)`

 SetCliaMatchRungNil sets the value for CliaMatchRung to be an explicit nil

### UnsetCliaMatchRung
`func (o *BiomarkerAssayProvidersRow) UnsetCliaMatchRung()`

UnsetCliaMatchRung ensures that no value is present for CliaMatchRung, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


