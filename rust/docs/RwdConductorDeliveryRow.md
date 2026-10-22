# RwdConductorDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**conductor_name_key** | Option<**String**> | The normalised institution name key the conductor rollup is keyed on; every register institution whose name folds to it is one conductor. | [optional]
**conductor_id** | Option<**String**> | The conductor slug: the register's normalised institution name key with spaces as hyphens, the id in the web page address (/rwd-conductors/<id>). The entry key of the rwd_conductors object type; injective because no name key carries a hyphen. | [optional]
**conductor_name** | Option<**String**> | The conductor's display name: one register spelling of the key. | [optional]
**conductor_names** | Option<**Vec<String>**> | Every register spelling of the key, as the study lists name it. | [optional]
**mandate_stratum** | Option<**String**> | all (every study of the conductor), mandated (studies a regulator required) or voluntary (studies the registrant says it did not require); studies with an unknown mandate are in all only. | [optional]
**register_institution_ids** | Option<**Vec<String>**> | The register Institution IDs of every register row that folds to this conductor (more than one when the register lists the name twice). | [optional]
**register_row_count** | Option<**i32**> | How many register institution rows fold to this conductor; above one means the register lists the name more than once and the counts are per name, not per row. | [optional]
**declared_office_countries** | Option<**Vec<String>**> | The office country text the register institutions declare, one element per distinct declaration; the office is not where the conductor delivered studies (see rwd_conductor_footprint). | [optional]
**any_encepp_partner** | Option<**bool**> | True when any register row of the conductor declares itself an ENCePP (European Network of Centres for Pharmacoepidemiology and Pharmacovigilance) partner. | [optional]
**lifetime_studies** | Option<**i32**> | Distinct registered studies the conductor is named as conducting institution on. Not additive with additional_role_studies. | [optional]
**recent_studies** | Option<**i32**> | Conducting studies that started within five years of the snapshot. | [optional]
**ongoing_studies** | Option<**i32**> | Conducting studies whose status is Ongoing. | [optional]
**additional_role_studies** | Option<**i32**> | Distinct registered studies the conductor is named on as an additional institution; a different claim from conducting, never added to it. | [optional]
**mandated_studies** | Option<**i32**> | Conducting studies a regulator required (mandate_stratum mandated). | [optional]
**mandated_share** | Option<**f64**> | mandated_studies over lifetime_studies; null when the conductor has no conducting study. | [optional]
**final_report_pairs** | Option<**i32**> | Conducting studies with both a planned and an actual final report date. | [optional]
**backfilled_pairs** | Option<**i32**> | Of those pairs, the studies whose planned and actual date are equal to the day (a likely backfill). | [optional]
**comparable_pairs** | Option<**i32**> | Conducting studies with a final report slip (both dates present and different); the denominator of the delivery figures. | [optional]
**backfilled_share** | Option<**f64**> | backfilled_pairs over final_report_pairs; null when there is no pair. | [optional]
**n_floor** | Option<**i32**> | The floor on comparable_pairs below which the median, quartiles and late count are null (suppressed, not zero). | [optional]
**meets_n_floor** | Option<**bool**> | True when comparable_pairs reaches n_floor, so the slip figures are shown. | [optional]
**median_slip_days** | Option<**f64**> | Median final report slip in days (positive is late) over the conductor's comparable pairs; null below the floor. | [optional]
**p25_slip_days** | Option<**f64**> | First quartile of the final report slip in days; null below the floor. | [optional]
**p75_slip_days** | Option<**f64**> | Third quartile of the final report slip in days; null below the floor. | [optional]
**late_over_30d** | Option<**i32**> | Comparable studies whose final report slipped more than 30 days; null below the floor. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


