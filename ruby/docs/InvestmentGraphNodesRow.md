# PharmaObservatory::InvestmentGraphNodesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **node_id** | **String** | The node key: an investor id (adv:, agency:, lei:, investor:), a company&#39;s organization id (the resolved parent), or investor_site: followed by the normalized printed name for an unadmitted website-only company. | [optional] |
| **node_kind** | **String** | investor or company. | [optional] |
| **display_name** | **String** | The investor&#39;s legal name, the company&#39;s resolved parent name, or for an unadmitted website-only company the lowest name as printed on the investors&#39; pages. | [optional] |
| **investor_type** | **String** | venture_capital, private_equity, hedge_fund, government_agency or other. Null for a company node. | [optional] |
| **investor_id** | **String** | The investor key for an investor node (the same as node_id), null for a company node. | [optional] |
| **org_id** | **String** | Legal-entity level: a company node&#39;s organization (the same as node_id), or an investor node&#39;s own organization resolved to its parent. Null for an unadmitted website-only company and for an investor with no organization. | [optional] |
| **company_id** | **String** | Company level: the company record of org_id (null where unresolved). | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **same_entity_node_id** | **String** | The other node of one real-world entity: an investor that is also a company with edges points at its company node and the reverse. An annotation, not an edge. Null when there is none. | [optional] |
| **node_resolution** | **String** | spine: the id is an organization or investor id. unadmitted_listing: a website-only company the organization spine has not admitted, keyed on its printed name (same-printed-name companies merge into one node; the id will change on admission). | [optional] |
| **edge_count** | **Integer** | Stored edge rows incident to the node, as the sum of out_edge_count and in_edge_count. A co_invested edge is stored in both directions, so a co-investor counts it once in each. | [optional] |
| **out_edge_count** | **Integer** | Stored edge rows that start at the node. | [optional] |
| **in_edge_count** | **Integer** | Stored edge rows that end at the node. | [optional] |
| **node_source_count** | **Integer** | Distinct sources among the node&#39;s edges. | [optional] |
| **sources** | **Array&lt;String&gt;** | The sources that state the node&#39;s edges. | [optional] |
| **unlinked_company_count** | **Integer** | For an investor node, its portfolio rows that cannot become a node (a company known only by an SEC filer number, or an organization with no resolved parent); shown as a count and never guessed at. 0 for a company node. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestmentGraphNodesRow.new(
  node_id: null,
  node_kind: null,
  display_name: null,
  investor_type: null,
  investor_id: null,
  org_id: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  same_entity_node_id: null,
  node_resolution: null,
  edge_count: null,
  out_edge_count: null,
  in_edge_count: null,
  node_source_count: null,
  sources: null,
  unlinked_company_count: null
)
```

