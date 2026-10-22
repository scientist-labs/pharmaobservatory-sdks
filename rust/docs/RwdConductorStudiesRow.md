# RwdConductorStudiesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**conductor_id** | Option<**String**> | The conductor slug: the register's normalised institution name key with spaces as hyphens, the id in the web page address (/rwd-conductors/<id>). The entry key of the rwd_conductors object type; injective because no name key carries a hyphen. | [optional]
**study_id** | Option<**String**> | The register's own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional]
**conductor_role** | Option<**String**> | conducting when the register names the conductor as an institution conducting the study, additional when as an additional institution. | [optional]
**study_title** | Option<**String**> | The study's title as registered. | [optional]
**study_status** | Option<**String**> | The study's status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional]
**mandate_stratum** | Option<**String**> | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional]
**study_start_actual_on** | Option<[**String**](string.md)> | Actual study start date. | [optional]
**final_report_slip_days** | Option<**i32**> | Actual minus planned final report date in days (positive is late); null when either date is missing or the two dates are equal (a backfill). | [optional]
**final_report_is_backfilled** | Option<**bool**> | True when the planned and actual final report dates are both present and equal to the day, which reads as a plan entered after the fact rather than perfect delivery. | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


