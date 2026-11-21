# ConditionSiteYearsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional]
**site_id** | Option<**String**> | The site key (a normalized facility, city and country). | [optional]
**start_year** | Option<**i32**> | Year the trials started. Null for a trial with no start date. | [optional]
**site_name** | Option<**String**> | The facility name as the registry lists it. | [optional]
**city** | Option<**String**> | City as listed. | [optional]
**state** | Option<**String**> | State or province as listed. | [optional]
**country** | Option<**String**> | Country as listed. | [optional]
**latitude** | Option<**f64**> | Latitude of the finest coordinate the site record has. | [optional]
**longitude** | Option<**f64**> | Longitude of the finest coordinate the site record has. | [optional]
**n_trials** | Option<**i32**> | Trials of the condition that started that year and list the site. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


