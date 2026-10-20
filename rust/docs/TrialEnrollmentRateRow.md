# TrialEnrollmentRateRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**phase** | Option<**String**> | The registry's phase code (PHASE1, PHASE2, PHASE2/PHASE3, NA and so on). Null when the registry gives no phase (observational and expanded-access trials); phase_label carries the display label for those. | [optional]
**phase_label** | Option<**String**> | The trial phase as a display label (Phase 2, Phase 3, Observational, Not applicable, Expanded access, Unknown). The same trial's phase column holds the registry code. | [optional]
**lead_sponsor_name** | Option<**String**> | The lead sponsor name as filed. | [optional]
**overall_status_summary** | Option<**String**> | A coarse status label (Completed and so on). | [optional]
**actual_enrollment** | Option<**i32**> | Enrollment typed ACTUAL. | [optional]
**enrollment_type** | Option<**String**> | ACTUAL or ESTIMATED. | [optional]
**enrollment_is_plausible** | Option<**bool**> | False when the registered enrollment is a placeholder value rather than a count. | [optional]
**active_sites** | Option<**i32**> | Number of sites the rate is computed over. | [optional]
**facility_country_count** | Option<**i32**> | Number of countries with a listed site. | [optional]
**has_us_facility** | Option<**bool**> | Whether any listed site is in the United States. | [optional]
**start_date** | Option<[**String**](string.md)> | Trial start date. | [optional]
**primary_completion_date** | Option<[**String**](string.md)> | Primary completion date. | [optional]
**completion_date** | Option<[**String**](string.md)> | Completion date. | [optional]
**primary_completion_date_type** | Option<**String**> | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry's older ANTICIPATED is published as ESTIMATED. | [optional]
**completion_date_type** | Option<**String**> | ACTUAL (a date that has occurred) or ESTIMATED (a planned date). The registry's older ANTICIPATED is published as ESTIMATED. | [optional]
**is_rate_computable** | Option<**bool**> | Whether the inputs support a rate. | [optional]
**rate_flag_reason** | Option<**String**> | Why a rate could not be computed (null when it could). | [optional]
**duration_months** | Option<**f64**> | Months from start to primary completion. | [optional]
**duration_months_overall** | Option<**f64**> | Months from start to completion. | [optional]
**pts_per_month** | Option<**f64**> | Participants enrolled per month. | [optional]
**pts_per_site_month** | Option<**f64**> | Participants enrolled per site per month. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


