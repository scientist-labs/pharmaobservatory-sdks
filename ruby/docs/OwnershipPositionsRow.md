# PharmaObservatory::OwnershipPositionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **position_id** | **String** | The position id. For a position whose owner name is withheld it is an opaque handle that derives nothing from the owner and is stable only while the other positions of the same owned company, filing family and dimension are unchanged. | [optional] |
| **source_class** | **String** | The filing family: sec_forms_3_4_5 (insiders), sec_schedules_13d_g (beneficial owners), sec_form_adv (investment adviser control persons) or sec_def_14a (proxy-statement tables). | [optional] |
| **ownership_dimension** | **String** | What the position measures: insider_reporting, beneficial_ownership, voting_power, economic_interest or control. Positions of different dimensions are different measurements and are not added. | [optional] |
| **relationship_type** | **String** | The owner&#39;s relationship to the company as the filing family states it. | [optional] |
| **subject_name** | **String** | The owned company as filed. | [optional] |
| **subject_org_id** | **String** | Legal-entity level: the owned company&#39;s organization record. | [optional] |
| **subject_company_id** | **String** | Company level: the owned company (null when unresolved). | [optional] |
| **subject_company_name** | **String** | Display name of the owned company. | [optional] |
| **subject_group_id** | **String** | Group level: the owned company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **subject_group_name** | **String** | Display name of the owned company&#39;s group. | [optional] |
| **owner_name** | **String** | The owner&#39;s name as filed. Null for an owner whose filing does not mark the name public (see owner_name_is_public). | [optional] |
| **owner_name_is_public** | **Boolean** | True when the filing marks the owner&#39;s name public. When false every column that could identify the owner is null (name, identity key, person id, investor id, organization and company columns, source_url and the dates); the position and its figures stay. | [optional] |
| **owner_kind** | **String** | What the owner is: organization, natural_person, reporting_group or unknown. | [optional] |
| **owner_identity_key** | **String** | The owner&#39;s stable identity (an SEC registration number or an adviser registration number); null when the name is withheld. | [optional] |
| **owner_person_id** | **String** | The key of an owner who is a natural person; null for other owners and when the name is withheld. | [optional] |
| **owner_investor_id** | **String** | The key of the owner in the investor directory when it is one; null when the name is withheld. | [optional] |
| **owner_org_id** | **String** | Legal-entity level: the owner&#39;s organization record (null when the owner is not an organization, is unresolved or is withheld). | [optional] |
| **owner_company_id** | **String** | Company level: the owner&#39;s company (null when unresolved or when the name is withheld). | [optional] |
| **owner_company_name** | **String** | Display name of the owner&#39;s company. | [optional] |
| **owner_group_id** | **String** | Group level: the owner&#39;s current ultimate parent; a root&#39;s group is itself (null when the owner is withheld). | [optional] |
| **owner_group_name** | **String** | Display name of the owner&#39;s group. | [optional] |
| **shares_basis** | **String** | How the share figure was reached: sum_standing_balances (insider balances), group_reported_maximum (a joint filing&#39;s maximum, never a sum of its filers) or no_quantity (the source gives none). Read a share figure with its basis. | [optional] |
| **shares** | **Float** | Shares the owner holds, on the shares_basis. Null when the source gives no quantity. | [optional] |
| **shares_direct** | **Float** | Of those shares, the part held directly. | [optional] |
| **shares_indirect** | **Float** | Of those shares, the part held indirectly. | [optional] |
| **derivative_underlying_shares** | **Float** | Shares underlying derivative positions such as options, reported apart and never added to shares. | [optional] |
| **percent_basis** | **String** | How the percentage was reached, for example single_reporting_person, uniform_group_percent or single_band. | [optional] |
| **reported_percent** | **Float** | The percentage the filing reports, on the percentage_basis. | [optional] |
| **reported_percent_low** | **Float** | The low end of a reported band, for a source that reports a band. | [optional] |
| **reported_percent_high** | **Float** | The high end of a reported band, for a source that reports a band. | [optional] |
| **percentage_basis** | **String** | What the percentage is of: beneficial_ownership, voting_power or economic_interest. Percentages of different bases are not comparable. | [optional] |
| **edge_state** | **String** | standing, ceased or no_longer_reported, for the owner&#39;s position as a whole. | [optional] |
| **position_count** | **Integer** | Positions the row consolidates. | [optional] |
| **standing_position_count** | **Integer** | Of those, positions still standing. | [optional] |
| **ceased_position_count** | **Integer** | Of those, positions reported ceased. | [optional] |
| **no_longer_reported_position_count** | **Integer** | Of those, positions the filer no longer reports. | [optional] |
| **first_observed_date** | **Date** | Date the position was first reported; null when the owner&#39;s name is withheld. | [optional] |
| **last_activity_date** | **Date** | Date of the latest reported activity; null when the owner&#39;s name is withheld. | [optional] |
| **last_filing_date** | **Date** | Date of the latest filing; null when the owner&#39;s name is withheld. | [optional] |
| **source_url** | **String** | A link to the latest filing; null when the owner&#39;s name is withheld, because the filing names the owner. | [optional] |
| **contains_personal_data** | **Boolean** | True when the row names a natural person. | [optional] |
| **privacy_status** | **String** | The filing&#39;s privacy status for the owner: not_applicable, pending_review, approved_public_fields, restricted or protected. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OwnershipPositionsRow.new(
  position_id: null,
  source_class: null,
  ownership_dimension: null,
  relationship_type: null,
  subject_name: null,
  subject_org_id: null,
  subject_company_id: null,
  subject_company_name: null,
  subject_group_id: null,
  subject_group_name: null,
  owner_name: null,
  owner_name_is_public: null,
  owner_kind: null,
  owner_identity_key: null,
  owner_person_id: null,
  owner_investor_id: null,
  owner_org_id: null,
  owner_company_id: null,
  owner_company_name: null,
  owner_group_id: null,
  owner_group_name: null,
  shares_basis: null,
  shares: null,
  shares_direct: null,
  shares_indirect: null,
  derivative_underlying_shares: null,
  percent_basis: null,
  reported_percent: null,
  reported_percent_low: null,
  reported_percent_high: null,
  percentage_basis: null,
  edge_state: null,
  position_count: null,
  standing_position_count: null,
  ceased_position_count: null,
  no_longer_reported_position_count: null,
  first_observed_date: null,
  last_activity_date: null,
  last_filing_date: null,
  source_url: null,
  contains_personal_data: null,
  privacy_status: null
)
```

