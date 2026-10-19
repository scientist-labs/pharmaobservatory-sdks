# SiteHospitalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**site_id** | Option<**String**> | The site key: one facility spelling from the registry. | [optional]
**ccn** | Option<**String**> | CMS Certification Number of the matched hospital. | [optional]
**hospital_name** | Option<**String**> | The hospital's name in the CMS registry. | [optional]
**hospital_type** | Option<**String**> | The CMS hospital type. | [optional]
**ownership** | Option<**String**> | The CMS ownership category. | [optional]
**bed_count** | Option<**i32**> | Certified bed count in the latest registry file. | [optional]
**is_teaching** | Option<**bool**> | Whether the hospital is a teaching hospital. | [optional]
**hospital_latitude** | Option<**f64**> | Latitude of the hospital. | [optional]
**hospital_longitude** | Option<**f64**> | Longitude of the hospital. | [optional]
**geocode_method** | Option<**String**> | How the hospital's coordinate was placed. | [optional]
**ccn_n_trials** | Option<**i64**> | Registered trials across the hospital's matched site records. | [optional]
**ccn_n_site_variants** | Option<**i64**> | Site records matched to the same hospital. | [optional]
**is_primary_site** | Option<**bool**> | True for the matched site record with the most trials for this hospital. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


