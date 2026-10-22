# PharmaObservatory::DrugPatentOwnersRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **owner_org_id** | **String** | Legal-entity level: the owning organization record as filed. | [optional] |
| **owner_name** | **String** | The owner&#39;s name. | [optional] |
| **owner_company_id** | **String** | Company level: the owner&#39;s company (null when the owner is not in the spine). | [optional] |
| **owner_company_name** | **String** | Display name of the owner&#39;s company. | [optional] |
| **owner_group_id** | **String** | Group level: the owner&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **owner_group_name** | **String** | Display name of the owner&#39;s group. | [optional] |
| **org_in_spine** | **Boolean** | True when the owner is an organization record in the spine. | [optional] |
| **n_patents** | **Integer** | Distinct declared patents the owner holds for the drug. | [optional] |
| **earliest_expiry** | **Date** | The earliest declared expiry. | [optional] |
| **latest_expiry** | **Date** | The latest declared expiry. | [optional] |
| **best_match_confidence** | **String** | The best assignee-to-organization match confidence (high, med, low). | [optional] |
| **link_bases** | **Array&lt;String&gt;** | The records that tie the patents to the drug (orange_book, uspto_pte). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugPatentOwnersRow.new(
  molecule_id: null,
  owner_org_id: null,
  owner_name: null,
  owner_company_id: null,
  owner_company_name: null,
  owner_group_id: null,
  owner_group_name: null,
  org_in_spine: null,
  n_patents: null,
  earliest_expiry: null,
  latest_expiry: null,
  best_match_confidence: null,
  link_bases: null
)
```

