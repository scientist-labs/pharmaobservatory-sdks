# InvestmentGraphEdgesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**edge_key** | Option<**String**> | The stable handle of one edge row: md5 of from_node_id, to_node_id, edge_type, source and source_ref. | [optional]
**from_node_id** | Option<**String**> | The node the edge starts at: an investor id or a company id (the resolved parent organization, or investor_site: and the normalized name for a website-only company the organization spine has not admitted). | [optional]
**to_node_id** | Option<**String**> | The node the edge ends at. Same id forms as from_node_id. invested_in goes from an investor to a company; co_invested is stored in both directions between investors. | [optional]
**from_investor_id** | Option<**String**> | The investor key when the from node is an investor (every edge today), null otherwise. | [optional]
**to_investor_id** | Option<**String**> | The investor key when the to node is an investor (a co_invested edge), null otherwise. | [optional]
**edge_type** | Option<**String**> | invested_in, co_invested, acquired, merged_with or licensed. acquired, merged_with and licensed have no rows until the M&A events dataset lands. There is no exited type: a listing that disappears from a website is never an end. | [optional]
**relationship_type** | Option<**String**> | The lane's own word for the relationship, so a public-market holding is never read as a venture round: portfolio_listed (investor website), holder_5pct and selling_stockholder (offering holder tables), institutional_position (13F), beneficial_owner (13D/G), co_appearance (co_invested). | [optional]
**source** | Option<**String**> | The lane that states the edge: investor_site, sec_edgar_holder_table (also the lane of the derived co_invested edges), sec_13f, sec_13dg, sec_form_d, sec_8k, company_press_release or reconciled. sec_form_d, sec_8k and company_press_release have no rows yet; reconciled carries the press-release and SEC deal metadata for the pairs it enriches and is the preferred lane of such a pair. | [optional]
**source_ref** | Option<**String**> | The evidence key within the lane: the portfolio company key for the website and holder-table lanes, the ownership observation id for 13F and 13D/G, and the literal co_invested_pairs for a derived co-investment. | [optional]
**source_url** | Option<**String**> | The page the edge was read from. Null for a derived co_invested edge, which has no page of its own. | [optional]
**event_date** | Option<[**String**](string.md)> | The date the lane states, of the kind in event_date_kind. Null exactly when event_date_kind is unknown. | [optional]
**event_date_kind** | Option<**String**> | What event_date is: transaction (a date the lane states for the deal itself), announcement, filing (the filing date), as_of (a report period end, as 13F states), form_d_filing (the filing date of a Form D offering by the company, never an investment date), first_observed (the first crawl that saw a website listing, never an investment date) or unknown (no date, event_date null). | [optional]
**observed_through** | Option<[**String**](string.md)> | The latest date the relationship was still observed in the lane (the last crawl, the last filing). Null when the lane gives none. Not an end date. | [optional]
**event_stage** | Option<**String**> | Where a deal stands (announced, signed, closed, completed or terminated), so an announced acquisition is never drawn as a completed one. Null for every lane today. | [optional]
**round_label** | Option<**String**> | The funding round named by the company's own press release (reconciled lane); null for every lane that does not state one. | [optional]
**amount_value** | Option<**f64**> | The ROUND amount in amount_currency, never the investor's own cheque (reconciled lane): the press-release round total, else the amount sold in a Form D offering. Null for every lane that does not state one. | [optional]
**amount_currency** | Option<**String**> | The currency of amount_value as the source states it, never defaulted to USD. Null while amount_value is null. | [optional]
**is_lead** | Option<**bool**> | Whether the investor led or co-led the round per the company's press release (reconciled lane): true for lead or co-lead, false for a named participant, null where the role is unstated or the lane does not state it. | [optional]
**shared_company_count** | Option<**i32**> | For a derived co_invested edge, the companies the two investors are both named against in offering filings. Null for every other edge. | [optional]
**is_derived** | Option<**bool**> | True for an edge computed from other rows (co_invested), false for an edge a source states. | [optional]
**pair_key** | Option<**String**> | from_node_id, to_node_id and edge_type joined with a pipe: every lane that states the same relationship shares one pair_key. | [optional]
**lane_rank** | Option<**i32**> | The lane's precedence when several lanes state one pair (lower wins): 13D/G 1, 13F 2, offering filings and derived co-investment 3, investor website 7. | [optional]
**is_preferred** | Option<**bool**> | True on exactly one row per pair_key, the lowest lane_rank (then the latest event_date): the row a picture draws. The other rows stay as evidence. | [optional]
**source_license** | Option<**String**> | The license class of the lane: public_domain_us_government or first_party_website_facts. | [optional]
**attribution** | Option<**String**> | The plain credit line for the lane. Show it with the edge for every reader. | [optional]
**from_name** | Option<**String**> | Display name of the from node. | [optional]
**from_org_id** | Option<**String**> | Legal-entity level: the from node's organization record (null for an investor with none and for an unadmitted website-only company). | [optional]
**from_company_id** | Option<**String**> | Company level: the from node's company record (null where unresolved). | [optional]
**from_company_name** | Option<**String**> | Display name of the from node's company. | [optional]
**from_group_id** | Option<**String**> | Group level: the from node's company's current ultimate parent; a root's group is itself. | [optional]
**from_group_name** | Option<**String**> | Display name of the from node's group. | [optional]
**to_name** | Option<**String**> | Display name of the to node. | [optional]
**to_org_id** | Option<**String**> | Legal-entity level: the to node's organization record (null for an investor with none and for an unadmitted website-only company). | [optional]
**to_company_id** | Option<**String**> | Company level: the to node's company record (null where unresolved). | [optional]
**to_company_name** | Option<**String**> | Display name of the to node's company. | [optional]
**to_group_id** | Option<**String**> | Group level: the to node's company's current ultimate parent; a root's group is itself. | [optional]
**to_group_name** | Option<**String**> | Display name of the to node's group. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


