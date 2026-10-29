# PharmaObservatory::BiomarkerAssayProvidersRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The route key of the gene the biomarker resolved to, the entry key for the biomarkers object type. Null when the extraction did not resolve a gene. | [optional] |
| **biomarker_key** | **String** | The grouping key of the biomarker. The biomarker&#39;s key when its gene resolved, otherwise raw: followed by the lower-cased name as written, so two spellings of one unresolved biomarker are two keys. | [optional] |
| **biomarker_id** | **String** | The biomarker&#39;s own key; null for a pair whose gene did not resolve. | [optional] |
| **biomarker_display_name** | **String** | The biomarker&#39;s name, as resolved or else as extracted. | [optional] |
| **gene_resolution_status** | **String** | resolved: the gene resolved. unresolved: a gene symbol was extracted and matched nothing. not_applicable: no gene symbol was extracted. | [optional] |
| **provider_key** | **String** | The provider the rows are grouped on. org: followed by the organization key when the company resolves to an organization, otherwise name: followed by the normalised company name. | [optional] |
| **provider_role** | **String** | What the papers name the provider as. assay_or_test_maker makes the assay or test; testing_service runs it as a service; both when the pair has papers of each kind. | [optional] |
| **provider_name** | **String** | The provider company name as it appears in the papers. | [optional] |
| **provider_org_id** | **String** | The organization the provider&#39;s company name resolved to (the member organization, not its parent); null when the provider is ambiguous or unresolved. Reached by the rung named in provider_org_match_tier. | [optional] |
| **provider_org_match_tier** | **String** | How the provider was resolved to an organization, first rung that yields one organization wins. alias_exact: a reviewed brand alias, then exactly one organization carries the name. name_exact: exactly one organization carries the name. parent_unanimous: several organizations carry the name and all roll up to one parent, which is the organization. sponsor_key: the name is a trial-sponsor key. rx_provider_link: through its marketplace supplier. ambiguous: organizations under more than one parent carry the name, so none is chosen. unresolved: no organization, the provider stays a name. | [optional] |
| **provider_parent_org_id** | **String** | The resolved parent of provider_org_id in the organization hierarchy (the organization itself when it has none); null when the provider has no organization. The hierarchy is not guaranteed correct for every company, so read it with provider_org_match_tier. | [optional] |
| **provider_company_id** | **String** | The provider&#39;s company key. | [optional] |
| **provider_company_name** | **String** | The provider&#39;s company name. | [optional] |
| **provider_group_id** | **String** | The provider&#39;s current ultimate parent group key. | [optional] |
| **provider_group_name** | **String** | The provider&#39;s group name. | [optional] |
| **rx_matched** | **Boolean** | Whether the provider name matches a supplier in the Scientist.com marketplace. | [optional] |
| **rx_provider_id** | **String** | The matched marketplace supplier&#39;s id; join to the suppliers dataset. | [optional] |
| **rx_provider_name** | **String** | The matched marketplace supplier&#39;s name. | [optional] |
| **rx_provider_slug** | **String** | The matched supplier&#39;s marketplace profile slug. | [optional] |
| **has_clinical_validation** | **Boolean** | Whether any paper for the pair mentions clinical validation of the assay. | [optional] |
| **has_analytical_validation** | **Boolean** | Whether any paper for the pair mentions analytical validation of the assay. | [optional] |
| **n_papers** | **Integer** | Papers naming the biomarker and the provider, dated or not. | [optional] |
| **n_dated_papers** | **Integer** | Papers with a recoverable publication date. | [optional] |
| **n_undated_papers** | **Integer** | Papers with no recoverable publication date. n_dated_papers plus n_undated_papers is n_papers. | [optional] |
| **most_recent_published_at** | **Time** | The latest publication date among the dated papers, in UTC. Null when no paper is dated, which is not a statement that none was published recently. | [optional] |
| **most_recent_pmid** | **String** | The PubMed id of the most recent dated paper. | [optional] |
| **most_recent_title** | **String** | The title of that paper. | [optional] |
| **meets_curated_threshold** | **Boolean** | True when the provider matches a marketplace supplier or at least two papers name the pair. | [optional] |
| **regulatory_evidence_kind** | **String** | The strongest US FDA or CMS record found for the pair. companion_diagnostic: the provider holds an FDA companion diagnostic for the biomarker. fda_cleared_specific_assay: a cleared or approved device of the provider that the paper names and that measures the biomarker. fda_cleared_for_biomarker: the provider holds a cleared or approved IVD under a product code for the biomarker. clia_testing_service: the provider is named as a testing service and its own organization holds a CLIA certificate. fda_ivd_company_only: the provider holds some cleared or approved IVD, none matched to the biomarker. none: no FDA record found. | [optional] |
| **has_fda_record_for_biomarker** | **Boolean** | True when regulatory_evidence_kind is one of the first four kinds, a record that concerns the biomarker rather than only the company. | [optional] |
| **regulatory_evidence_reasons** | **Array&lt;String&gt;** | Every level that held for the pair, strongest first; empty when none did. | [optional] |
| **n_fda_ivd_510k** | **Integer** | The 510(k) clearances in the in-vitro diagnostic review panels held by the provider&#39;s organization (or, when it holds none, by organizations under its parent; fda_company_match_rung says which). Null for a provider that did not resolve to an organization. | [optional] |
| **n_fda_ivd_pma** | **Integer** | The PMA approvals (distinct approval numbers) in the in-vitro diagnostic review panels held by the provider&#39;s organization or its parent family, on the same rung as n_fda_ivd_510k. Null for a provider that did not resolve to an organization. | [optional] |
| **fda_company_match_rung** | **String** | org: the counted records are held by the provider&#39;s own organization. parent: by organizations under the same resolved parent. Null when the provider holds none. | [optional] |
| **n_fda_analyte_510k** | **Integer** | The provider&#39;s 510(k) clearances under a product code mapped to this biomarker. Null when the provider or the biomarker did not resolve. | [optional] |
| **n_fda_analyte_pma** | **Integer** | The provider&#39;s PMA approvals under a product code mapped to this biomarker. Null when the provider or the biomarker did not resolve. | [optional] |
| **fda_analyte_product_codes** | **Array&lt;String&gt;** | The FDA product codes behind n_fda_analyte_510k and n_fda_analyte_pma. | [optional] |
| **fda_analyte_fda_url** | **String** | The FDA database page of the provider&#39;s most recent record under a product code for this biomarker. | [optional] |
| **specific_assay_matched** | **Boolean** | True when a paper names a cleared or approved device of this provider that measures this biomarker. The device name matches the paper&#39;s assay name, or the company&#39;s own quote, and the device&#39;s analyte agrees with the biomarker. | [optional] |
| **specific_assay_device_name** | **String** | The 510(k) device name or PMA trade name of the matched device; the best match when several matched. Null when no device matched. | [optional] |
| **specific_assay_k_number** | **String** | The 510(k) number of the matched device; null when the match is a PMA or there is none. | [optional] |
| **specific_assay_pma_number** | **String** | The PMA number of the matched device; null when the match is a 510(k) or there is none. | [optional] |
| **specific_assay_fda_url** | **String** | The FDA database page of the matched device. | [optional] |
| **specific_assay_similarity** | **Float** | The trigram text similarity (0 to 1) of the larger of the two text clauses that matched the device, the paper&#39;s assay name or the company&#39;s quote against the device name. | [optional] |
| **n_cdx_entries** | **Integer** | The FDA companion diagnostic list entries for this biomarker from this provider&#39;s organization (or its parent family). Null when the provider or the biomarker did not resolve. | [optional] |
| **cdx_diagnostic_names** | **Array&lt;String&gt;** | The names of those companion diagnostics. | [optional] |
| **n_clia_certificates** | **Integer** | The CLIA (US clinical laboratory) certificates held by the provider&#39;s organization, or by organizations under its parent when it holds none (clia_match_rung says which). A certificate licenses a lab; it is not evidence that any named assay is run. Null for a provider that did not resolve to an organization. | [optional] |
| **clia_match_rung** | **String** | org: the counted certificates are held by the provider&#39;s own organization. parent: by organizations under the same resolved parent. Null when the provider holds none. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerAssayProvidersRow.new(
  route_key: null,
  biomarker_key: null,
  biomarker_id: null,
  biomarker_display_name: null,
  gene_resolution_status: null,
  provider_key: null,
  provider_role: null,
  provider_name: null,
  provider_org_id: null,
  provider_org_match_tier: null,
  provider_parent_org_id: null,
  provider_company_id: null,
  provider_company_name: null,
  provider_group_id: null,
  provider_group_name: null,
  rx_matched: null,
  rx_provider_id: null,
  rx_provider_name: null,
  rx_provider_slug: null,
  has_clinical_validation: null,
  has_analytical_validation: null,
  n_papers: null,
  n_dated_papers: null,
  n_undated_papers: null,
  most_recent_published_at: null,
  most_recent_pmid: null,
  most_recent_title: null,
  meets_curated_threshold: null,
  regulatory_evidence_kind: null,
  has_fda_record_for_biomarker: null,
  regulatory_evidence_reasons: null,
  n_fda_ivd_510k: null,
  n_fda_ivd_pma: null,
  fda_company_match_rung: null,
  n_fda_analyte_510k: null,
  n_fda_analyte_pma: null,
  fda_analyte_product_codes: null,
  fda_analyte_fda_url: null,
  specific_assay_matched: null,
  specific_assay_device_name: null,
  specific_assay_k_number: null,
  specific_assay_pma_number: null,
  specific_assay_fda_url: null,
  specific_assay_similarity: null,
  n_cdx_entries: null,
  cdx_diagnostic_names: null,
  n_clia_certificates: null,
  clia_match_rung: null
)
```

