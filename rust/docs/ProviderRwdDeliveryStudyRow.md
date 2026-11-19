# ProviderRwdDeliveryStudyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**provider_rwd_delivery_study_id** | Option<**String**> | The row id: a hash of provider_uuid, conductor_name_key and study_id. | [optional]
**provider_id** | Option<**String**> | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/<id>). The entry key of the suppliers object type. A supplier is not a company. | [optional]
**provider_uuid** | Option<**String**> | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional]
**provider_slug** | Option<**String**> | The supplier's marketplace slug. | [optional]
**provider_name** | Option<**String**> | The supplier's name. | [optional]
**institution_id** | Option<**String**> | The register Institution ID of the winning route to the conductor (strongest method, then lowest id). | [optional]
**match_method** | Option<**String**> | match_method says how the provider was tied to the register institution: org_member_exact (through the supplier's resolved organization, match_confidence carried from that match), rx_name_key_exact (an exact match of the register institution's name key to the supplier's, match_confidence null, a name match and not a verified identity) or org_unit_of_entry (an organization unit of the supplier's organization, one-directional, match_confidence the bridge's own figure). Where a provider reaches one conductor by several routes the strongest wins, in that order, and match_methods lists every route. | [optional]
**match_confidence** | Option<**f64**> | The bridge's own confidence for the winning route, 0 to 1; null on an rx_name_key_exact row, whose exact name match defines none. Passed through, never computed here. | [optional]
**conductor_name_key** | Option<**String**> | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. Not a vendor family. | [optional]
**study_id** | Option<**String**> | The register's own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional]
**eu_pas_number** | Option<**String**> | The EU PAS (post-authorisation study) register number of the study. | [optional]
**study_mandate_stratum** | Option<**String**> | The study's own mandate: mandated (a regulator required it), voluntary, or unknown (the register does not say). | [optional]
**final_report_planned_on** | Option<[**String**](string.md)> | Planned date of the final study report. | [optional]
**final_report_actual_on** | Option<[**String**](string.md)> | Actual date of the final study report. | [optional]
**final_report_slip_days** | Option<**i32**> | Actual minus planned final report date in days (positive is late, negative is early). A plain fact about one study, never a statistic. | [optional]
**all_meets_n_floor** | Option<**bool**> | True when the provider and conductor's all row meets the floor of 10 comparable pairs. | [optional]
**stratum_meets_n_floor** | Option<**bool**> | True when the row of the study's own mandate stratum meets the floor; null for a study whose mandate is unknown (it belongs to the all row only). | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


