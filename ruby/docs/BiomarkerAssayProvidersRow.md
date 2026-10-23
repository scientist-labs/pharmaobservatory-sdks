# PharmaObservatory::BiomarkerAssayProvidersRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **route_key** | **String** | The route key of the gene the biomarker resolved to, the entry key for the biomarkers object type. Null when the extraction did not resolve a gene. | [optional] |
| **biomarker_key** | **String** | The grouping key of the biomarker. The biomarker&#39;s key when its gene resolved, otherwise raw: followed by the lower-cased name as written, so two spellings of one unresolved biomarker are two keys. | [optional] |
| **biomarker_id** | **String** | The biomarker&#39;s own key; null for a pair whose gene did not resolve. | [optional] |
| **biomarker_display_name** | **String** | The biomarker&#39;s name, as resolved or else as extracted. | [optional] |
| **gene_resolution_status** | **String** | resolved: the gene resolved. unresolved: a gene symbol was extracted and matched nothing. not_applicable: no gene symbol was extracted. | [optional] |
| **company_name_key** | **String** | The normalised company name the provider rows are grouped on. | [optional] |
| **provider_name** | **String** | The provider company name as it appears in the papers. | [optional] |
| **provider_org_id** | **String** | The provider&#39;s legal entity key, where the marketplace match resolves one. | [optional] |
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
| **assay_provider_class** | **String** | The company-level class. diagnostic_provider makes or supplies the diagnostic assay; reagent_supplier is the vendor whose kit appears in a methods section; context_dependent, pharma_or_other and not_a_company are the other decisions; unclassified means no decision yet. | [optional] |
| **assay_provider_class_source** | **String** | Where the class decision came from; null when unclassified. | [optional] |
| **assay_provider_class_decided_at** | **Time** | When the class was decided, in UTC; null when unclassified. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerAssayProvidersRow.new(
  route_key: null,
  biomarker_key: null,
  biomarker_id: null,
  biomarker_display_name: null,
  gene_resolution_status: null,
  company_name_key: null,
  provider_name: null,
  provider_org_id: null,
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
  assay_provider_class: null,
  assay_provider_class_source: null,
  assay_provider_class_decided_at: null
)
```

