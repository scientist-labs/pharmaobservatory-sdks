# PharmaObservatory::OwnershipObservationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ownership_observation_id** | **String** | The observation id. | [optional] |
| **source_observation_id** | **String** | The id the source gave the assertion. | [optional] |
| **supersedes_observation_id** | **String** | The earlier observation this one amends or replaces. | [optional] |
| **superseded_by_observation_id** | **String** | The later observation that supersedes this one (the lowest when more than one claims it). | [optional] |
| **superseded_by_count** | **Integer** | How many later observations claim to supersede this one; above 1 is a lineage conflict kept as filed. | [optional] |
| **supersession_chain_id** | **String** | The id of the amendment thread this observation belongs to. | [optional] |
| **supersession_depth** | **Integer** | Position in the amendment thread; 0 is the first filing. | [optional] |
| **is_chain_terminal** | **Boolean** | True when nothing supersedes this observation. | [optional] |
| **chain_truncated** | **Boolean** | True when the amendment thread was cut at 25 steps. | [optional] |
| **source_system** | **String** | The source system that made the assertion. | [optional] |
| **source_class** | **String** | The filing family: sec_forms_3_4_5, sec_schedules_13d_g, sec_form_adv, sec_def_14a or wikidata. | [optional] |
| **semantic_key** | **String** | The ownership semantic the assertion carries, from the project&#39;s controlled taxonomy. | [optional] |
| **relationship_type** | **String** | The owner&#39;s relationship to the company as the filing family states it. | [optional] |
| **ownership_dimension** | **String** | What the observation measures: economic_interest, voting_power, beneficial_ownership, investment_discretion, insider_reporting, consolidation, control or unspecified. Different dimensions are different measurements. | [optional] |
| **directness** | **String** | direct, indirect, mixed or unspecified. | [optional] |
| **owner_kind** | **String** | What the owner is: organization, natural_person, reporting_group, government or unknown. | [optional] |
| **subject_source_ref** | **String** | The owned party as the source identifies it. | [optional] |
| **subject_name** | **String** | The owned company as filed. | [optional] |
| **subject_legal_entity_id** | **String** | The legal entity the source resolved the owned party to, when it did. | [optional] |
| **subject_company_id** | **String** | Company level: the subject company (null when unresolved). | [optional] |
| **subject_company_name** | **String** | Display name of the subject company. | [optional] |
| **subject_group_id** | **String** | Group level: the subject company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **subject_group_name** | **String** | Display name of the subject company&#39;s group. | [optional] |
| **subject_org_id** | **String** | Legal-entity level: the owned company&#39;s organization record. | [optional] |
| **subject_investor_id** | **String** | The investor directory key of the owned party when it is an investor. | [optional] |
| **subject_resolution_method** | **String** | How the owned party was resolved. | [optional] |
| **subject_resolution_confidence** | **Float** | Confidence of that resolution, 0 to 1. | [optional] |
| **owner_source_ref** | **String** | The owner as the source identifies it. | [optional] |
| **owner_name** | **String** | The owner&#39;s name as filed. | [optional] |
| **owner_legal_entity_id** | **String** | The legal entity the source resolved the owner to, when it did. | [optional] |
| **owner_company_id** | **String** | Company level: the owner company (null when unresolved). | [optional] |
| **owner_company_name** | **String** | Display name of the owner company. | [optional] |
| **owner_group_id** | **String** | Group level: the owner company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **owner_group_name** | **String** | Display name of the owner company&#39;s group. | [optional] |
| **owner_org_id** | **String** | Legal-entity level: the owner&#39;s organization record. | [optional] |
| **owner_investor_id** | **String** | The investor directory key of the owner when it is one. | [optional] |
| **owner_person_id** | **String** | The key of an owner who is a natural person; null for other owners. | [optional] |
| **person_link_method** | **String** | How the person key was linked: edgar_cik, adviser_crd, unresolved_person or not_person. | [optional] |
| **person_link_conflict** | **Boolean** | True when the owner links to more than one person key. | [optional] |
| **owner_resolution_method** | **String** | How the owner was resolved. | [optional] |
| **owner_resolution_confidence** | **Float** | Confidence of that resolution, 0 to 1. | [optional] |
| **review_state** | **String** | The review state of the resolution. | [optional] |
| **stake_value_type** | **String** | How the stake is stated: exact, range, band, explicitly_unknown, not_reported or not_applicable. No midpoint, zero or numeral is invented for a non-numeric state. | [optional] |
| **percentage_basis** | **String** | What the percentage is a share of: economic_interest, voting_power, beneficial_ownership, control or ownership_unspecified. | [optional] |
| **stake_percent** | **Float** | The stake as an exact percentage, 0 to 100. | [optional] |
| **stake_lower_percent** | **Float** | The lower bound of a range or band stake, 0 to 100. | [optional] |
| **stake_upper_percent** | **Float** | The upper bound of a range or band stake, 0 to 100. | [optional] |
| **stake_lower_inclusive** | **Boolean** | Whether the lower bound is included. | [optional] |
| **stake_upper_inclusive** | **Boolean** | Whether the upper bound is included. | [optional] |
| **stake_band_code** | **String** | The source&#39;s band code, when the stake is a band. | [optional] |
| **stake_display** | **String** | A label for the stake that never invents a number, such as 5.5%, 5% to 10% or Not reported. | [optional] |
| **security_title** | **String** | The security the stake is in, as filed. | [optional] |
| **shares_amount** | **Float** | The number of shares, when the source states one. | [optional] |
| **reported_value_amount** | **Float** | The reported value, in reported_value_currency at reported_value_scale. | [optional] |
| **reported_value_currency** | **String** | The currency of the reported value. | [optional] |
| **reported_value_scale** | **String** | The scale of the reported value as filed. | [optional] |
| **announced_date** | **Date** | The date announced. | [optional] |
| **effective_date** | **Date** | The effective date. | [optional] |
| **filing_date** | **Date** | The filing date. | [optional] |
| **as_of_date** | **Date** | The as-of date of the figure. | [optional] |
| **ceased_date** | **Date** | The date the position ceased. | [optional] |
| **observation_date** | **Date** | The sort date: the first of as_of_date, effective_date, filing_date and announced_date that is present. | [optional] |
| **observation_date_basis** | **String** | Which of those dates observation_date took, or unavailable. | [optional] |
| **lifecycle_status** | **String** | The lifecycle status the source states. | [optional] |
| **source_document_id** | **String** | The id of the source document. | [optional] |
| **source_url** | **String** | The link to the filing. | [optional] |
| **source_content_hash** | **String** | A hash of the source content, for change detection. | [optional] |
| **evidence_quote** | **String** | The quoted evidence from the source. | [optional] |
| **restriction_tier** | **String** | The tier vocabulary is none or licensable. Every row here is none: every source behind the ledger is open, and a data test reports any other tier. | [optional] |
| **source_license** | **String** | The licence of the source. | [optional] |
| **commercial_use_allowed** | **Boolean** | Whether the source licence allows commercial use. | [optional] |
| **attribution** | **String** | The source credit text for the row. | [optional] |
| **contains_personal_data** | **Boolean** | A tag, never a gate: true when the source row carries a natural person&#39;s data. | [optional] |
| **privacy_status** | **String** | The source&#39;s own privacy classification of the row. | [optional] |
| **is_current_edge** | **Boolean** | True when this observation is the one the current ownership edge selected. | [optional] |
| **current_edge_state** | **String** | The state of the current edge when is_current_edge. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OwnershipObservationsRow.new(
  ownership_observation_id: null,
  source_observation_id: null,
  supersedes_observation_id: null,
  superseded_by_observation_id: null,
  superseded_by_count: null,
  supersession_chain_id: null,
  supersession_depth: null,
  is_chain_terminal: null,
  chain_truncated: null,
  source_system: null,
  source_class: null,
  semantic_key: null,
  relationship_type: null,
  ownership_dimension: null,
  directness: null,
  owner_kind: null,
  subject_source_ref: null,
  subject_name: null,
  subject_legal_entity_id: null,
  subject_company_id: null,
  subject_company_name: null,
  subject_group_id: null,
  subject_group_name: null,
  subject_org_id: null,
  subject_investor_id: null,
  subject_resolution_method: null,
  subject_resolution_confidence: null,
  owner_source_ref: null,
  owner_name: null,
  owner_legal_entity_id: null,
  owner_company_id: null,
  owner_company_name: null,
  owner_group_id: null,
  owner_group_name: null,
  owner_org_id: null,
  owner_investor_id: null,
  owner_person_id: null,
  person_link_method: null,
  person_link_conflict: null,
  owner_resolution_method: null,
  owner_resolution_confidence: null,
  review_state: null,
  stake_value_type: null,
  percentage_basis: null,
  stake_percent: null,
  stake_lower_percent: null,
  stake_upper_percent: null,
  stake_lower_inclusive: null,
  stake_upper_inclusive: null,
  stake_band_code: null,
  stake_display: null,
  security_title: null,
  shares_amount: null,
  reported_value_amount: null,
  reported_value_currency: null,
  reported_value_scale: null,
  announced_date: null,
  effective_date: null,
  filing_date: null,
  as_of_date: null,
  ceased_date: null,
  observation_date: null,
  observation_date_basis: null,
  lifecycle_status: null,
  source_document_id: null,
  source_url: null,
  source_content_hash: null,
  evidence_quote: null,
  restriction_tier: null,
  source_license: null,
  commercial_use_allowed: null,
  attribution: null,
  contains_personal_data: null,
  privacy_status: null,
  is_current_edge: null,
  current_edge_state: null
)
```

