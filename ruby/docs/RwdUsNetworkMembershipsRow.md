# PharmaObservatory::RwdUsNetworkMembershipsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **membership_assertion_id** | **String** | The assertion id: a hash of the network, institution, provider uuid and match method. | [optional] |
| **network_source** | **String** | The source the member rows come from. | [optional] |
| **network_key** | **String** | The network key, the entry key of the rwd_us_networks object type. | [optional] |
| **network_name** | **String** | The network&#39;s name. | [optional] |
| **network_operator** | **String** | The operating body. | [optional] |
| **membership_kind** | **String** | The strongest agreement the institution holds. membership_kind is the strongest agreement the institution holds (dta_signatory, else dua_signatory, else lhba_signatory), never data contributor. | [optional] |
| **has_dta** | **Boolean** | The institution holds at least one published DTA (Data Transfer Agreement) row. | [optional] |
| **has_dua** | **Boolean** | The institution holds at least one published DUA (Data Use Agreement) row. | [optional] |
| **has_lhba** | **Boolean** | The institution holds at least one agreement published with the type code LHBA, which N3C does not expand. | [optional] |
| **has_current_dta** | **Boolean** | A DTA row&#39;s published term end is on or after the retrieval date. | [optional] |
| **has_current_dua** | **Boolean** | A DUA row is within its published term. | [optional] |
| **has_current_lhba** | **Boolean** | An LHBA row is within its published term. | [optional] |
| **first_executed_on** | **Date** | The earliest published execution date over the institution&#39;s agreement rows. | [optional] |
| **latest_term_ends_on** | **Date** | The latest published term end over the institution&#39;s rows. | [optional] |
| **agreement_count** | **Integer** | Published agreement rows folded into this member. | [optional] |
| **institution_name** | **String** | The institution as N3C publishes it (the ROR name). | [optional] |
| **institution_name_key** | **String** | The normalized name key, the join key to other name-keyed datasets. | [optional] |
| **ror_id** | **String** | The bare ROR id (Research Organization Registry) N3C published; null when it published a domain instead. | [optional] |
| **institution_domain** | **String** | The organization&#39;s email domain as published, an organization domain and not a personal address. | [optional] |
| **is_lead** | **Boolean** | Null for N3C, which has no lead concept; false would be a false No. | [optional] |
| **as_of** | **Date** | The retrieval date of the agreement snapshot. | [optional] |
| **source_url** | **String** | The API endpoint the agreements came from. | [optional] |
| **retrieved_at** | **Time** | When the snapshot was retrieved, in UTC. | [optional] |
| **org_id** | **String** | The organization spine id the member resolved to, or null. Carried, never minted. | [optional] |
| **org_name** | **String** | The spine organization&#39;s name. | [optional] |
| **org_match_rung** | **String** | How the member reached the organization: ror_exact (the source&#39;s own ROR id), xref, match_keys, ambiguous or unresolved. | [optional] |
| **org_match_tier** | **String** | The tier of the organization match (ror_exact is high, stated as the source&#39;s own identifier). | [optional] |
| **org_match_confidence** | **Float** | The organization match confidence (0.95 on ror_exact). | [optional] |
| **provider_id** | **String** | The RX supplier id as text, the entry key of the suppliers object type. Null on an unmatched row. | [optional] |
| **provider_uuid** | **String** | The Scientist.com public provider uuid. Null on an unmatched row. | [optional] |
| **provider_slug** | **String** | The supplier&#39;s marketplace slug. | [optional] |
| **provider_name** | **String** | The supplier&#39;s name. | [optional] |
| **match_method** | **String** | How the supplier was matched: org_member_exact (through the resolved organization), rx_name_key_exact (an exact name match, not a verified identity) or unmatched (no provider). | [optional] |
| **also_matched_by_name** | **Boolean** | True on an org_member_exact row whose provider also matches by name; null on an unmatched row. | [optional] |
| **match_confidence** | **Float** | The organization match confidence passed through on org_member_exact; null on rx_name_key_exact and unmatched. | [optional] |
| **member_total** | **Integer** | Members of the network, the register-wide denominator printed beside any count. | [optional] |
| **members_with_current_dta** | **Integer** | Network members with a DTA within its published term. | [optional] |
| **members_with_current_dua** | **Integer** | Network members with a DUA within its published term. | [optional] |
| **members_resolved_org** | **Integer** | Network members resolved to an organization. | [optional] |
| **members_matched_rx** | **Integer** | Network members matched to at least one RX provider. | [optional] |
| **attribution** | **String** | The credit line every surface carrying this data prints. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdUsNetworkMembershipsRow.new(
  membership_assertion_id: null,
  network_source: null,
  network_key: null,
  network_name: null,
  network_operator: null,
  membership_kind: null,
  has_dta: null,
  has_dua: null,
  has_lhba: null,
  has_current_dta: null,
  has_current_dua: null,
  has_current_lhba: null,
  first_executed_on: null,
  latest_term_ends_on: null,
  agreement_count: null,
  institution_name: null,
  institution_name_key: null,
  ror_id: null,
  institution_domain: null,
  is_lead: null,
  as_of: null,
  source_url: null,
  retrieved_at: null,
  org_id: null,
  org_name: null,
  org_match_rung: null,
  org_match_tier: null,
  org_match_confidence: null,
  provider_id: null,
  provider_uuid: null,
  provider_slug: null,
  provider_name: null,
  match_method: null,
  also_matched_by_name: null,
  match_confidence: null,
  member_total: null,
  members_with_current_dta: null,
  members_with_current_dua: null,
  members_resolved_org: null,
  members_matched_rx: null,
  attribution: null
)
```

