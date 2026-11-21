# PharmaObservatory::RwdUsNetworksRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **network_key** | **String** | The network key, the entry key of the rwd_us_networks object type (n3c today). | [optional] |
| **network_source** | **String** | The source the network rows come from. | [optional] |
| **network_name** | **String** | The network&#39;s name. | [optional] |
| **network_operator** | **String** | The operating body. | [optional] |
| **member_total** | **Integer** | Institutions holding at least one published agreement, our count of the register. | [optional] |
| **members_with_dta** | **Integer** | Institutions holding a Data Transfer Agreement. | [optional] |
| **members_with_current_dta** | **Integer** | Institutions with a DTA row whose published term end is on or after the retrieval date. | [optional] |
| **members_with_dua** | **Integer** | Institutions holding a Data Use Agreement. | [optional] |
| **members_with_current_dua** | **Integer** | Institutions with a DUA row within its published term. | [optional] |
| **members_with_lhba** | **Integer** | Institutions holding an agreement published with the type code LHBA. | [optional] |
| **members_resolved_org** | **Integer** | Members resolved to an organization on the spine. | [optional] |
| **members_matched_rx** | **Integer** | Members matched to at least one RX provider, by an exact method visible in rwd_us_network_memberships. | [optional] |
| **published_dta_count** | **Integer** | N3C&#39;s own published count of DTAs, as N3C labels it. Differs from ours; not reconciled. | [optional] |
| **published_dua_count** | **Integer** | N3C&#39;s own published count of DUAs. Differs from ours; not reconciled. | [optional] |
| **published_institution_count** | **Integer** | N3C&#39;s own published count of institutions. Differs from ours; not reconciled. | [optional] |
| **published_project_count** | **Integer** | N3C&#39;s own published count of projects. Differs from ours; not reconciled. | [optional] |
| **as_of** | **Date** | The retrieval date of the agreement snapshot. | [optional] |
| **source_url** | **String** | The API endpoint the agreements came from. | [optional] |
| **agreements_url** | **String** | The public N3C dashboard. | [optional] |
| **signatories_page_url** | **String** | N3C&#39;s HTML signatories page, kept as a documented reconciliation and not ingested. | [optional] |
| **retrieved_at** | **Time** | When the snapshot was retrieved, in UTC. | [optional] |
| **attribution** | **String** | The credit line every surface carrying this data prints. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RwdUsNetworksRow.new(
  network_key: null,
  network_source: null,
  network_name: null,
  network_operator: null,
  member_total: null,
  members_with_dta: null,
  members_with_current_dta: null,
  members_with_dua: null,
  members_with_current_dua: null,
  members_with_lhba: null,
  members_resolved_org: null,
  members_matched_rx: null,
  published_dta_count: null,
  published_dua_count: null,
  published_institution_count: null,
  published_project_count: null,
  as_of: null,
  source_url: null,
  agreements_url: null,
  signatories_page_url: null,
  retrieved_at: null,
  attribution: null
)
```

