# PharmaObservatory::InvestmentGraphNodeEdgesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **node_id** | **String** | The endpoint this row is read from: the from node when direction is out, the to node when it is in. | [optional] |
| **direction** | **String** | out when node_id is the edge&#39;s from node, in when it is the to node. | [optional] |
| **edge_key** | **String** | The stable handle of the edge, the same on its two rows. | [optional] |
| **node_kind** | **String** | investor or company: the kind of node_id. | [optional] |
| **investor_id** | **String** | The investor key when node_id is an investor, null otherwise. | [optional] |
| **node_name** | **String** | Display name of node_id. | [optional] |
| **org_id** | **String** | Legal-entity level: node_id&#39;s organization (null for an investor with none and for an unadmitted website-only company). | [optional] |
| **company_id** | **String** | Company level: node_id&#39;s company record (null where unresolved). | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **counterpart_node_id** | **String** | The node at the other end of the edge. | [optional] |
| **counterpart_kind** | **String** | investor or company: the kind of counterpart_node_id. | [optional] |
| **counterpart_name** | **String** | Display name of the counterpart node. | [optional] |
| **from_node_id** | **String** | The node the edge starts at: an investor id or a company id (the resolved parent organization, or investor_site: and the normalized name for a website-only company the organization spine has not admitted). | [optional] |
| **to_node_id** | **String** | The node the edge ends at. Same id forms as from_node_id. invested_in goes from an investor to a company; co_invested is stored in both directions between investors. | [optional] |
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

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestmentGraphNodeEdgesRow.new(
  node_id: null,
  direction: null,
  edge_key: null,
  node_kind: null,
  investor_id: null,
  node_name: null,
  org_id: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  counterpart_node_id: null,
  counterpart_kind: null,
  counterpart_name: null,
  from_node_id: null,
  to_node_id: null,
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
  attribution: null
)
```

