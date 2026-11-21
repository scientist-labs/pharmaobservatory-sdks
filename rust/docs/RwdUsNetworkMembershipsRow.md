# RwdUsNetworkMembershipsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**membership_assertion_id** | Option<**String**> | The assertion id: a hash of the network, institution, provider uuid and match method. | [optional]
**network_source** | Option<**String**> | The source the member rows come from. | [optional]
**network_key** | Option<**String**> | The network key, the entry key of the rwd_us_networks object type. | [optional]
**network_name** | Option<**String**> | The network's name. | [optional]
**network_operator** | Option<**String**> | The operating body. | [optional]
**membership_kind** | Option<**String**> | The strongest agreement the institution holds. membership_kind is the strongest agreement the institution holds (dta_signatory, else dua_signatory, else lhba_signatory), never data contributor. | [optional]
**has_dta** | Option<**bool**> | The institution holds at least one published DTA (Data Transfer Agreement) row. | [optional]
**has_dua** | Option<**bool**> | The institution holds at least one published DUA (Data Use Agreement) row. | [optional]
**has_lhba** | Option<**bool**> | The institution holds at least one agreement published with the type code LHBA, which N3C does not expand. | [optional]
**has_current_dta** | Option<**bool**> | A DTA row's published term end is on or after the retrieval date. | [optional]
**has_current_dua** | Option<**bool**> | A DUA row is within its published term. | [optional]
**has_current_lhba** | Option<**bool**> | An LHBA row is within its published term. | [optional]
**first_executed_on** | Option<[**String**](string.md)> | The earliest published execution date over the institution's agreement rows. | [optional]
**latest_term_ends_on** | Option<[**String**](string.md)> | The latest published term end over the institution's rows. | [optional]
**agreement_count** | Option<**i32**> | Published agreement rows folded into this member. | [optional]
**institution_name** | Option<**String**> | The institution as N3C publishes it (the ROR name). | [optional]
**institution_name_key** | Option<**String**> | The normalized name key, the join key to other name-keyed datasets. | [optional]
**ror_id** | Option<**String**> | The bare ROR id (Research Organization Registry) N3C published; null when it published a domain instead. | [optional]
**institution_domain** | Option<**String**> | The organization's email domain as published, an organization domain and not a personal address. | [optional]
**is_lead** | Option<**bool**> | Null for N3C, which has no lead concept; false would be a false No. | [optional]
**as_of** | Option<[**String**](string.md)> | The retrieval date of the agreement snapshot. | [optional]
**source_url** | Option<**String**> | The API endpoint the agreements came from. | [optional]
**retrieved_at** | Option<**String**> | When the snapshot was retrieved, in UTC. | [optional]
**org_id** | Option<**String**> | The organization spine id the member resolved to, or null. Carried, never minted. | [optional]
**org_name** | Option<**String**> | The spine organization's name. | [optional]
**org_match_rung** | Option<**String**> | How the member reached the organization: ror_exact (the source's own ROR id), xref, match_keys, ambiguous or unresolved. | [optional]
**org_match_tier** | Option<**String**> | The tier of the organization match (ror_exact is high, stated as the source's own identifier). | [optional]
**org_match_confidence** | Option<**f64**> | The organization match confidence (0.95 on ror_exact). | [optional]
**provider_id** | Option<**String**> | The RX supplier id as text, the entry key of the suppliers object type. Null on an unmatched row. | [optional]
**provider_uuid** | Option<**String**> | The Scientist.com public provider uuid. Null on an unmatched row. | [optional]
**provider_slug** | Option<**String**> | The supplier's marketplace slug. | [optional]
**provider_name** | Option<**String**> | The supplier's name. | [optional]
**match_method** | Option<**String**> | How the supplier was matched: org_member_exact (through the resolved organization), rx_name_key_exact (an exact name match, not a verified identity) or unmatched (no provider). | [optional]
**also_matched_by_name** | Option<**bool**> | True on an org_member_exact row whose provider also matches by name; null on an unmatched row. | [optional]
**match_confidence** | Option<**f64**> | The organization match confidence passed through on org_member_exact; null on rx_name_key_exact and unmatched. | [optional]
**member_total** | Option<**i32**> | Members of the network, the register-wide denominator printed beside any count. | [optional]
**members_with_current_dta** | Option<**i32**> | Network members with a DTA within its published term. | [optional]
**members_with_current_dua** | Option<**i32**> | Network members with a DUA within its published term. | [optional]
**members_resolved_org** | Option<**i32**> | Network members resolved to an organization. | [optional]
**members_matched_rx** | Option<**i32**> | Network members matched to at least one RX provider. | [optional]
**attribution** | Option<**String**> | The credit line every surface carrying this data prints. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


