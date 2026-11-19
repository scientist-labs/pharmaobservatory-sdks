# PharmaObservatory::InvestmentGraphEdgesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **edge_key** | **String** | The stable handle of one edge row: md5 of from_node_id, to_node_id, edge_type, source and source_ref. | [optional] |
| **from_node_id** | **String** | The node the edge starts at: an investor id or a company id (the resolved parent organization, or investor_site: and the normalized name for a website-only company the organization spine has not admitted). | [optional] |
| **to_node_id** | **String** | The node the edge ends at. Same id forms as from_node_id. invested_in goes from an investor to a company; co_invested is stored in both directions between investors. | [optional] |
| **from_investor_id** | **String** | The investor key when the from node is an investor (every edge today), null otherwise. | [optional] |
| **to_investor_id** | **String** | The investor key when the to node is an investor (a co_invested edge), null otherwise. | [optional] |
| **edge_type** | **String** | invested_in, co_invested, acquired, merged_with or licensed. acquired, merged_with and licensed have no rows until the M&amp;A events dataset lands. There is no exited type: a listing that disappears from a website is never an end. | [optional] |
| **relationship_type** | **String** | The lane&#39;s own word for the relationship, so a public-market holding is never read as a venture round: portfolio_listed (investor website), holder_5pct and selling_stockholder (offering holder tables), institutional_position (13F), beneficial_owner (13D/G), co_appearance (co_invested). | [optional] |
| **source** | **String** | The lane that states the edge: investor_site, sec_edgar_holder_table (also the lane of the derived co_invested edges), sec_13f, sec_13dg, sec_form_d, sec_8k, company_press_release or reconciled. sec_form_d, sec_8k and company_press_release have no rows yet; reconciled carries the press-release and SEC deal metadata for the pairs it enriches and is the preferred lane of such a pair. | [optional] |
| **source_ref** | **String** | The evidence key within the lane: the portfolio company key for the website and holder-table lanes, the ownership observation id for 13F and 13D/G, and the literal co_invested_pairs for a derived co-investment. | [optional] |
| **source_url** | **String** | The page the edge was read from. Null for a derived co_invested edge, which has no page of its own. | [optional] |
| **event_date** | **Date** | The date the lane states, of the kind in event_date_kind. Null exactly when event_date_kind is unknown. | [optional] |
| **event_date_kind** | **String** | What event_date is: transaction (a date the lane states for the deal itself), announcement, filing (the filing date), as_of (a report period end, as 13F states), form_d_filing (the filing date of a Form D offering by the company, never an investment date), first_observed (the first crawl that saw a website listing, never an investment date) or unknown (no date, event_date null). | [optional] |
| **observed_through** | **Date** | The latest date the relationship was still observed in the lane (the last crawl, the last filing). Null when the lane gives none. Not an end date. | [optional] |
| **event_stage** | **String** | Where a deal stands (announced, signed, closed, completed or terminated), so an announced acquisition is never drawn as a completed one. Null for every lane today. | [optional] |
| **round_label** | **String** | The funding round named by the company&#39;s own press release (reconciled lane); null for every lane that does not state one. | [optional] |
| **amount_value** | **Float** | The ROUND amount in amount_currency, never the investor&#39;s own cheque (reconciled lane): the press-release round total, else the amount sold in a Form D offering. Null for every lane that does not state one. | [optional] |
| **amount_currency** | **String** | The currency of amount_value as the source states it, never defaulted to USD. Null while amount_value is null. | [optional] |
| **is_lead** | **Boolean** | Whether the investor led or co-led the round per the company&#39;s press release (reconciled lane): true for lead or co-lead, false for a named participant, null where the role is unstated or the lane does not state it. | [optional] |
| **shared_company_count** | **Integer** | For a derived co_invested edge, the companies the two investors are both named against in offering filings. Null for every other edge. | [optional] |
| **is_derived** | **Boolean** | True for an edge computed from other rows (co_invested), false for an edge a source states. | [optional] |
| **pair_key** | **String** | from_node_id, to_node_id and edge_type joined with a pipe: every lane that states the same relationship shares one pair_key. | [optional] |
| **lane_rank** | **Integer** | The lane&#39;s precedence when several lanes state one pair (lower wins): 13D/G 1, 13F 2, offering filings and derived co-investment 3, investor website 7. | [optional] |
| **is_preferred** | **Boolean** | True on exactly one row per pair_key, the lowest lane_rank (then the latest event_date): the row a picture draws. The other rows stay as evidence. | [optional] |
| **source_license** | **String** | The license class of the lane: public_domain_us_government or first_party_website_facts. | [optional] |
| **attribution** | **String** | The plain credit line for the lane. Show it with the edge for every reader. | [optional] |
| **from_name** | **String** | Display name of the from node. | [optional] |
| **from_org_id** | **String** | Legal-entity level: the from node&#39;s organization record (null for an investor with none and for an unadmitted website-only company). | [optional] |
| **from_company_id** | **String** | Company level: the from node&#39;s company record (null where unresolved). | [optional] |
| **from_company_name** | **String** | Display name of the from node&#39;s company. | [optional] |
| **from_group_id** | **String** | Group level: the from node&#39;s company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **from_group_name** | **String** | Display name of the from node&#39;s group. | [optional] |
| **to_name** | **String** | Display name of the to node. | [optional] |
| **to_org_id** | **String** | Legal-entity level: the to node&#39;s organization record (null for an investor with none and for an unadmitted website-only company). | [optional] |
| **to_company_id** | **String** | Company level: the to node&#39;s company record (null where unresolved). | [optional] |
| **to_company_name** | **String** | Display name of the to node&#39;s company. | [optional] |
| **to_group_id** | **String** | Group level: the to node&#39;s company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **to_group_name** | **String** | Display name of the to node&#39;s group. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestmentGraphEdgesRow.new(
  edge_key: null,
  from_node_id: null,
  to_node_id: null,
  from_investor_id: null,
  to_investor_id: null,
  edge_type: null,
  relationship_type: null,
  source: null,
  source_ref: null,
  source_url: null,
  event_date: null,
  event_date_kind: null,
  observed_through: null,
  event_stage: null,
  round_label: null,
  amount_value: null,
  amount_currency: null,
  is_lead: null,
  shared_company_count: null,
  is_derived: null,
  pair_key: null,
  lane_rank: null,
  is_preferred: null,
  source_license: null,
  attribution: null,
  from_name: null,
  from_org_id: null,
  from_company_id: null,
  from_company_name: null,
  from_group_id: null,
  from_group_name: null,
  to_name: null,
  to_org_id: null,
  to_company_id: null,
  to_company_name: null,
  to_group_id: null,
  to_group_name: null
)
```

