# PharmaObservatory::BiomarkerValidationRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **biomarker_validation_id** | **String** | The evidence assertion id. | [optional] |
| **route_key** | **String** | The biomarker&#39;s route key, the entry key for the biomarkers object type; for a gene it is the gene symbol. | [optional] |
| **gene_symbol** | **String** | The gene symbol (HGNC, the HUGO Gene Nomenclature Committee). | [optional] |
| **biomarker_id** | **String** | The biomarker key. | [optional] |
| **biomarker_display_name** | **String** | The biomarker&#39;s display name. | [optional] |
| **biomarker_kind** | **String** | The kind of biomarker the evidence is filed against. | [optional] |
| **organization_name** | **String** | The organization that can supply the evidence, as the source names it (null when the source names none that resolved). | [optional] |
| **organization_org_id** | **String** | Legal-entity level: the organization record (null when unresolved). | [optional] |
| **organization_company_id** | **String** | Company level: the organization&#39;s company (null when unresolved). | [optional] |
| **organization_company_name** | **String** | Display name of the organization&#39;s company. | [optional] |
| **organization_group_id** | **String** | Group level: the organization&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **organization_group_name** | **String** | Display name of the organization&#39;s group. | [optional] |
| **organization_role** | **String** | What the organization is to the evidence: lead_sponsor, collaborator, contract_vendor, assay_manufacturer, qualification_submitter, testing_lab or publishing_org. | [optional] |
| **organization_match_tier** | **String** | How the organization name matched an organization record. | [optional] |
| **tier** | **Integer** | The rung on the validation ladder, 1 to 4. A rung is a kind of evidence, never a score: rows are not ranked across organizations or evidence types. | [optional] |
| **tier_label** | **String** | The rung&#39;s name, for example fda_authorized_companion_diagnostic or named_in_inclusion_criteria. | [optional] |
| **evidence_type** | **String** | The evidence lane: fda_bqp_qualification, fda_cdx_authorization, aact_endpoint_role, aact_eligibility_section, aact_biospecimen, trial_mention or clinvar_clinical_testing. | [optional] |
| **evidence_count** | **Integer** | Sibling rows sharing this biomarker, organization and role. Not the number of source records behind a ClinVar row (clinvar_scv_count). | [optional] |
| **evidence_ref** | **String** | The source reference the row cites. | [optional] |
| **evidence_ref_kind** | **String** | What kind of reference evidence_ref is. | [optional] |
| **trial_id** | **String** | The trial key (an NCT id) when the evidence is a trial; null otherwise. | [optional] |
| **citation_url** | **String** | A link to the cited record. | [optional] |
| **citation_verified** | **Boolean** | False for the FDA Biomarker Qualification Program arm, whose captured links are known to go stale; true otherwise. A link marked false may not resolve. | [optional] |
| **evidence_snippet** | **String** | The text span the evidence rests on, where the source gives one. | [optional] |
| **biomarker_role** | **String** | The biomarker&#39;s role in the evidence, where the source states it. | [optional] |
| **is_surrogate_endpoint_claim** | **Boolean** | True when the evidence is a claim that the biomarker is a surrogate endpoint. | [optional] |
| **clinvar_scv_count** | **Integer** | For a ClinVar row, the number of ClinVar submission records the row collapses; null otherwise. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkerValidationRow.new(
  biomarker_validation_id: null,
  route_key: null,
  gene_symbol: null,
  biomarker_id: null,
  biomarker_display_name: null,
  biomarker_kind: null,
  organization_name: null,
  organization_org_id: null,
  organization_company_id: null,
  organization_company_name: null,
  organization_group_id: null,
  organization_group_name: null,
  organization_role: null,
  organization_match_tier: null,
  tier: null,
  tier_label: null,
  evidence_type: null,
  evidence_count: null,
  evidence_ref: null,
  evidence_ref_kind: null,
  trial_id: null,
  citation_url: null,
  citation_verified: null,
  evidence_snippet: null,
  biomarker_role: null,
  is_surrogate_endpoint_claim: null,
  clinvar_scv_count: null
)
```

