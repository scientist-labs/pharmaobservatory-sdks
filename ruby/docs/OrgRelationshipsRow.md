# PharmaObservatory::OrgRelationshipsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **subject_org_id** | **String** | Legal-entity level: the subject&#39;s organization record. | [optional] |
| **subject_name** | **String** | The organization&#39;s name as filed. | [optional] |
| **subject_company_id** | **String** | Company level: the subject&#39;s company (null when unresolved). | [optional] |
| **subject_company_name** | **String** | Display name of the subject&#39;s company. | [optional] |
| **subject_group_id** | **String** | Group level: the subject&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **subject_group_name** | **String** | Display name of the subject&#39;s group. | [optional] |
| **object_org_id** | **String** | Legal-entity level: the object&#39;s organization record. | [optional] |
| **object_name** | **String** | The organization&#39;s name as filed. | [optional] |
| **object_company_id** | **String** | Company level: the object&#39;s company (null when unresolved). | [optional] |
| **object_company_name** | **String** | Display name of the object&#39;s company. | [optional] |
| **object_group_id** | **String** | Group level: the object&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **object_group_name** | **String** | Display name of the object&#39;s group. | [optional] |
| **predicate** | **String** | The directed relationship from subject to object: parent_of and ultimate_parent_of (corporate ownership, the subject is the parent), linked_charity_of (a charity linked to the main charity), uses_vendor (a sponsor and the contract research organization inferred for its trials), co_sponsors_trial_with (two sponsors of one trial), project_partner_of and awards_to (UK research grants). | [optional] |
| **edge_family** | **String** | The family of the predicate: corporate_structure (parent_of, ultimate_parent_of, linked_charity_of), contract_vendor (uses_vendor), co_sponsorship (co_sponsors_trial_with) or research_funding (project_partner_of, awards_to). | [optional] |
| **edge_scope** | **String** | general for an edge from a source that covers organizations at large (GLEIF, Charity Commission, the contract research organization (CRO) inference); rare_disease_scoped for an edge from a source read only for organizations in the rare-disease evidence set (ROR, co-sponsorship, UK research grants). A rare_disease_scoped edge is a true edge and an incomplete picture: an organization with none of them has not been shown to have none. | [optional] |
| **source_key** | **String** | The source that asserted the edge: gleif (Global Legal Entity Identifier Foundation Golden Copy), charity_commission, aact_cro_inference (inferred from ClinicalTrials.gov filings against a curated CRO list), ror (Research Organization Registry), aact (ClinicalTrials.gov) or ukri_gtr (UK Research and Innovation Gateway to Research). | [optional] |
| **evidence_type** | **String** | The kind of evidence behind the edge: registry (a register entry), trial, project or grant. | [optional] |
| **authority_class** | **String** | authoritative when the source itself filed the relationship; derived when the edge is an inference (the CRO edge is a curated-list name match, not a filing). | [optional] |
| **evidence_count** | **Integer** | The source records behind the edge. Counted within one source, so it does not add across sources. | [optional] |
| **context_count** | **Integer** | Distinct source record contexts among them (for example distinct trials). | [optional] |
| **confidence** | **Float** | The highest confidence among the edge&#39;s records, 0 to 1, where the source states one. Null for sources that state none. A GLEIF ownership-resolution score and a CRO name-match score are different measures; do not compare them. | [optional] |
| **first_observed_at** | **Time** | When the earliest of the edge&#39;s records was observed in the source. | [optional] |
| **last_observed_at** | **Time** | When the latest of the edge&#39;s records was observed in the source. | [optional] |
| **license_name** | **String** | The licence of the source behind the edge. | [optional] |
| **attribution** | **String** | The credit line the source requires for the edge. Show it wherever the edge is shown: the Charity Commission and UK Research and Innovation edges require it under the Open Government Licence v3.0. | [optional] |
| **restriction_tier** | **String** | The tier vocabulary is none or licensable. Every row here is none: every source behind these edges is open, and a test on the source assertions reports any restricted source that is ever added. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgRelationshipsRow.new(
  subject_org_id: null,
  subject_name: null,
  subject_company_id: null,
  subject_company_name: null,
  subject_group_id: null,
  subject_group_name: null,
  object_org_id: null,
  object_name: null,
  object_company_id: null,
  object_company_name: null,
  object_group_id: null,
  object_group_name: null,
  predicate: null,
  edge_family: null,
  edge_scope: null,
  source_key: null,
  evidence_type: null,
  authority_class: null,
  evidence_count: null,
  context_count: null,
  confidence: null,
  first_observed_at: null,
  last_observed_at: null,
  license_name: null,
  attribution: null,
  restriction_tier: null
)
```

