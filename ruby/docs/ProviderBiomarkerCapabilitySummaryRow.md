# PharmaObservatory::ProviderBiomarkerCapabilitySummaryRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **summary_id** | **String** | The summary row id: a hash of the grain columns. | [optional] |
| **provider_id** | **String** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). A supplier is not a company. | [optional] |
| **provider_uuid** | **String** | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional] |
| **provider_slug** | **String** | The supplier&#39;s marketplace slug. | [optional] |
| **provider_name** | **String** | The supplier&#39;s name. | [optional] |
| **route_key** | **String** | The biomarker&#39;s route key, the entry key for the biomarkers object type; the gene symbol for a gene biomarker. | [optional] |
| **biomarker_id** | **String** | The canonical biomarker entity id. | [optional] |
| **gene_symbol** | **String** | The gene symbol; null for a source-scoped biomarker that names no gene. | [optional] |
| **biomarker_display_name** | **String** | The biomarker&#39;s display name. | [optional] |
| **evidence_lane** | **String** | The source path of the group: ladder_trial and ladder_clinvar (the evidence ladder), fda_cdx (FDA companion diagnostic list), fda_bqp (FDA Biomarker Qualification Program), literature_assay (assay providers named in PubMed full text), supplier_survey (the supplier&#39;s own declaration) clinvar_submission (ClinVar clinical-testing submissions by a laboratory the supplier carries by name, tier null) and supplier_service_listing (a Scientist.com marketplace catalogue service the supplier listed, whose name names the gene; tier null, no evidence sense). Nothing is merged across lanes. | [optional] |
| **org_role** | **String** | The role the supplier played in the evidence: a ladder role, the literature provider role verbatim, respondent for the survey, or marketplace_listing for a service listing. | [optional] |
| **match_method** | **String** | How the evidence was attached to the supplier: org_member_exact (through the supplier&#39;s resolved organization), rx_name_key_exact (an exact name match, not a verified identity) or rx_provider_direct (the supplier&#39;s own marketplace listing, no name matched). | [optional] |
| **analytical_validation_evidence** | **String** | Analytical validation of the group (literature_assay lane only): fda_cleared_specific_assay, fda_cleared_for_biomarker or literature_reported. | [optional] |
| **clinical_use_evidence** | **String** | Clinical-use evidence of the group: the ladder tier label, literature_reported, or on the clinvar_submission lane the ClinVar review status verbatim. | [optional] |
| **regulatory_qualification** | **String** | FDA qualification (fda_bqp lane only): fda_qualified. | [optional] |
| **companion_diagnostic** | **String** | FDA companion diagnostic authorization (fda_cdx lane only): fda_authorized_companion_diagnostic. | [optional] |
| **supplier_declared_status** | **String** | The supplier&#39;s own declaration, never mapped into the four evidence kinds: the survey&#39;s assay status on supplier_survey, supplier_declared_service_listed on supplier_service_listing. | [optional] |
| **tier** | **Integer** | The ladder rung of the group, 1 to 4, a kind of evidence and never a score; null on literature_assay, supplier_survey, clinvar_submission and supplier_service_listing. | [optional] |
| **assertion_count** | **Integer** | How many evidence assertions the row stands for; sums to the row count of provider_biomarker_capability under the same filter. | [optional] |
| **distinct_evidence_ref_count** | **Integer** | How many distinct evidence references (one ClinVar record, one trial, one paper) the assertions cite. | [optional] |
| **example_evidence_ref** | **String** | The smallest evidence reference of the group: a deterministic example to link, not a ranking. | [optional] |
| **example_evidence_ref_kind** | **String** | Which identifier example_evidence_ref is: nct_id, clinvar_scv, fda_cdx_entry_biomarker, fda_bqp_project, pmid, assay_ledger_row, survey_response_biomarker or rx_ware. | [optional] |
| **example_citation_url** | **String** | The citation page of the example reference; null where the source offers no resolvable page. | [optional] |
| **most_recent_evidence_date** | **Date** | The latest publication date in the group (literature_assay lane only). | [optional] |
| **source_license** | **String** | The license statement of the group&#39;s source lane. No lane carries a usage restriction. | [optional] |
| **attribution** | **String** | The credit the group&#39;s source lane is owed. | [optional] |
| **as_of** | **Date** | The build date. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ProviderBiomarkerCapabilitySummaryRow.new(
  summary_id: null,
  provider_id: null,
  provider_uuid: null,
  provider_slug: null,
  provider_name: null,
  route_key: null,
  biomarker_id: null,
  gene_symbol: null,
  biomarker_display_name: null,
  evidence_lane: null,
  org_role: null,
  match_method: null,
  analytical_validation_evidence: null,
  clinical_use_evidence: null,
  regulatory_qualification: null,
  companion_diagnostic: null,
  supplier_declared_status: null,
  tier: null,
  assertion_count: null,
  distinct_evidence_ref_count: null,
  example_evidence_ref: null,
  example_evidence_ref_kind: null,
  example_citation_url: null,
  most_recent_evidence_date: null,
  source_license: null,
  attribution: null,
  as_of: null
)
```

