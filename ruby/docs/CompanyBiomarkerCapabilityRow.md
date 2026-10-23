# PharmaObservatory::CompanyBiomarkerCapabilityRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **biomarker_validation_id** | **String** | The evidence assertion id. | [optional] |
| **organization_name** | **String** | The organization&#39;s name as filed on the evidence. | [optional] |
| **organization_org_id** | **String** | Legal-entity level: the organization record of the organization. | [optional] |
| **organization_company_id** | **String** | Company level: the company the organization belongs to (null when unresolved). | [optional] |
| **organization_company_name** | **String** | Display name of the organization&#39;s company. | [optional] |
| **organization_group_id** | **String** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **organization_group_name** | **String** | Display name of the organization&#39;s group. | [optional] |
| **gene_symbol** | **String** | The gene&#39;s HGNC (Human Genome Organisation Gene Nomenclature Committee) symbol; null for evidence that names no gene. | [optional] |
| **biomarker_id** | **String** | The biomarker&#39;s id (a hash, not the route key). | [optional] |
| **biomarker_display_name** | **String** | The biomarker&#39;s display name. | [optional] |
| **biomarker_kind** | **String** | The kind of biomarker. | [optional] |
| **biomarker_identity_status** | **String** | How firmly the biomarker&#39;s identity is resolved. | [optional] |
| **organization_role** | **String** | The organization&#39;s role in the evidence (for example assay developer or testing laboratory). | [optional] |
| **organization_match_tier** | **String** | How the organization was matched to its record. | [optional] |
| **tier** | **Integer** | The evidence tier, higher for stronger kinds of evidence. A label of the kind of evidence, not a quality score. | [optional] |
| **tier_label** | **String** | The evidence tier in words. | [optional] |
| **evidence_type** | **String** | The kind of evidence (companion diagnostic authorization, trial endpoint, trial eligibility, trial specimen collection, trial text mention or ClinVar clinical testing). | [optional] |
| **evidence_count** | **Integer** | Assertions that share this biomarker, organization and role. | [optional] |
| **evidence_ref** | **String** | The reference of the evidence (a trial id, a device application number or a project id). | [optional] |
| **evidence_ref_kind** | **String** | What evidence_ref is. | [optional] |
| **trial_id** | **String** | The trial&#39;s registry id (NCT, National Clinical Trial, number) when the evidence is a trial. | [optional] |
| **citation_url** | **String** | A link to the evidence. | [optional] |
| **evidence_snippet** | **String** | A short excerpt of the evidence. | [optional] |
| **biomarker_role** | **String** | The role the biomarker plays in the evidence. | [optional] |
| **is_surrogate_endpoint_claim** | **Boolean** | Whether the evidence claims the biomarker as a surrogate endpoint. | [optional] |
| **clinvar_scv_count** | **Integer** | ClinVar submissions behind a ClinVar row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyBiomarkerCapabilityRow.new(
  biomarker_validation_id: null,
  organization_name: null,
  organization_org_id: null,
  organization_company_id: null,
  organization_company_name: null,
  organization_group_id: null,
  organization_group_name: null,
  gene_symbol: null,
  biomarker_id: null,
  biomarker_display_name: null,
  biomarker_kind: null,
  biomarker_identity_status: null,
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
  evidence_snippet: null,
  biomarker_role: null,
  is_surrogate_endpoint_claim: null,
  clinvar_scv_count: null
)
```

