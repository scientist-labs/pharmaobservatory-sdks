# TrialSiteHistoryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**as_of** | Option<[**String**](string.md)> | The registry snapshot date this row describes. | [optional]
**prev_as_of** | Option<[**String**](string.md)> | The trial's previous snapshot date (null for its first snapshot). | [optional]
**facilities_captured** | Option<**bool**> | True when the snapshot carried the site list; false means the site counts are unknown, not zero. | [optional]
**status_captured** | Option<**bool**> | True when the snapshot carried a recruitment status on the site rows. | [optional]
**comparable** | Option<**bool**> | True when the previous snapshot also carried the site list, so the add and drop counts are defined. | [optional]
**sites_listed** | Option<**i32**> | Distinct sites the trial listed at the snapshot (null where facilities_captured is false). | [optional]
**sites_added** | Option<**i32**> | Sites present at this snapshot and absent at the previous one (null where comparable is false). | [optional]
**sites_dropped** | Option<**i32**> | Sites present at the previous snapshot and absent at this one (null where comparable is false). | [optional]
**sites_status_populated** | Option<**i32**> | Sites whose row carried a recruitment status at the snapshot. | [optional]
**sites_recruiting** | Option<**i32**> | Sites with status RECRUITING at the snapshot. | [optional]
**sites_not_yet_recruiting** | Option<**i32**> | Sites with status NOT_YET_RECRUITING at the snapshot. | [optional]
**sites_active_not_recruiting** | Option<**i32**> | Sites with status ACTIVE_NOT_RECRUITING at the snapshot. | [optional]
**sites_completed** | Option<**i32**> | Sites with status COMPLETED at the snapshot. | [optional]
**sites_other_status** | Option<**i32**> | Sites with a status other than the four counted above. | [optional]
**countries_listed** | Option<**i32**> | Countries the listed sites span at the snapshot. | [optional]
**captured_columns** | Option<**Vec<String>**> | The registry study fields the snapshot's archive carried (null when no capture record exists for the snapshot; an empty list means it carried none of them). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


