# ProviderRwdDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**provider_rwd_delivery_id** | Option<**String**> | The row id: a hash of provider_uuid, conductor_name_key and mandate_stratum. | [optional]
**provider_id** | Option<**String**> | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/<id>). The entry key of the suppliers object type. A supplier is not a company. | [optional]
**provider_uuid** | Option<**String**> | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional]
**provider_slug** | Option<**String**> | The supplier's marketplace slug. | [optional]
**provider_name** | Option<**String**> | The supplier's name. | [optional]
**institution_id** | Option<**String**> | The register Institution ID of the winning route to the conductor (strongest method, then lowest id). | [optional]
**match_method** | Option<**String**> | match_method says how the provider was tied to the register institution: org_member_exact (through the supplier's resolved organization, match_confidence carried from that match), rx_name_key_exact (an exact match of the register institution's name key to the supplier's, match_confidence null, a name match and not a verified identity) or org_unit_of_entry (an organization unit of the supplier's organization, one-directional, match_confidence the bridge's own figure). Where a provider reaches one conductor by several routes the strongest wins, in that order, and match_methods lists every route. | [optional]
**match_confidence** | Option<**f64**> | The bridge's own confidence for the winning route, 0 to 1; null on an rx_name_key_exact row, whose exact name match defines none. Passed through, never computed here. | [optional]
**conductor_name_key** | Option<**String**> | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. Not a vendor family. | [optional]
**conductor_id** | Option<**String**> | The conductor slug: the conductor name key with spaces as hyphens, the entry key of the rwd_conductors object type (the id in /app/rwd-conductors/<id>). | [optional]
**matched_institution_ids** | Option<**Vec<String>**> | Every register Institution ID that routes the provider to the conductor name key (more than one when the register lists the name twice or two methods reach it). | [optional]
**match_methods** | Option<**Vec<String>**> | Every match method that routes the provider to this conductor name key; match_method is the strongest of them. | [optional]
**conductor_name** | Option<**String**> | The conductor's display name: one register spelling of the key. | [optional]
**mandate_stratum** | Option<**String**> | all (every study of the conductor), mandated (studies a regulator required) or voluntary (studies the registrant says it did not require); studies with an unknown mandate are in all only. The mandated and voluntary books slip differently: choose a stratum. | [optional]
**lifetime_studies** | Option<**i32**> | Distinct registered studies the conductor is named as conducting institution on. | [optional]
**final_report_pairs** | Option<**i32**> | Conducting studies with both a planned and an actual final report date. | [optional]
**comparable_pairs** | Option<**i32**> | Conducting studies with a final report slip (both dates present and different); the N of the figures and the denominator of the floor. | [optional]
**backfilled_pairs** | Option<**i32**> | Of the final_report_pairs, the studies whose planned and actual date are equal to the day (a likely backfill, a form artefact). Excluded from every figure, counted here. | [optional]
**backfilled_share** | Option<**f64**> | backfilled_pairs over final_report_pairs; null when there is no pair. | [optional]
**mandated_share** | Option<**f64**> | The share of the conductor's conducting studies a regulator required; read it beside any figure to see how mandated the book is. Null when the conductor has no conducting study. | [optional]
**n_floor** | Option<**i32**> | The floor on comparable_pairs below which the four distribution figures are null (suppressed, not zero): 10. | [optional]
**meets_n_floor** | Option<**bool**> | True when comparable_pairs reaches n_floor, so the four figures are shown; false means they are null and the studies are in provider_rwd_delivery_study. | [optional]
**median_slip_days** | Option<**f64**> | Median final report slip in days against the planned date (positive is late) over the comparable pairs; null below the floor. | [optional]
**p25_slip_days** | Option<**f64**> | First quartile of the final report slip in days; null below the floor. | [optional]
**p75_slip_days** | Option<**f64**> | Third quartile of the final report slip in days; null below the floor. | [optional]
**late_over_30d** | Option<**i32**> | Comparable studies whose final report slipped more than 30 days; null below the floor. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


