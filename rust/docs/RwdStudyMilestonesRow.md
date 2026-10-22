# RwdStudyMilestonesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**study_id** | Option<**String**> | The register's own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional]
**milestone_ordinal** | Option<**i32**> | The milestone's position in the study timeline. | [optional]
**milestone_key** | Option<**String**> | Which milestone: funding_contract, study_start, data_analysis_start, interim_report or final_report. | [optional]
**milestone_label** | Option<**String**> | The milestone's label. | [optional]
**planned_on** | Option<[**String**](string.md)> | The planned date. | [optional]
**actual_on** | Option<[**String**](string.md)> | The actual date. | [optional]
**has_pair** | Option<**bool**> | True when both dates are present. | [optional]
**is_backfilled** | Option<**bool**> | True when planned and actual are both present and equal to the day (a plan entered after the fact). | [optional]
**is_comparable_pair** | Option<**bool**> | True when both dates are present and differ, so the slip is a measurement; every delivery distribution is computed over these rows. | [optional]
**slip_days** | Option<**i32**> | Actual minus planned in days (positive is late); null when either date is missing or the two are equal. | [optional]
**is_late_over_30d** | Option<**bool**> | True when the slip is more than 30 days; null when the pair is not comparable. | [optional]
**required_by_regulator** | Option<**String**> | The registrant's answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional]
**mandate_stratum** | Option<**String**> | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional]
**rmp_requirement** | Option<**String**> | The risk management plan (RMP) category the registrant says required the study, as typed. | [optional]
**study_status** | Option<**String**> | The study's status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional]
**snapshot_label** | Option<**String**> | The label of the catalogue snapshot the row was read from (a date). | [optional]
**accessed_month** | Option<**String**> | The month and year the catalogue was accessed (YYYY-MM), which EMA's site asks to be cited with its URL. | [optional]
**restriction_tier** | Option<**String**> | The license restriction tier of the row's source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional]
**attribution** | Option<**String**> | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


