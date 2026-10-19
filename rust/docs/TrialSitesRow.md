# TrialSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**site_key** | Option<**String**> | A site key built from the normalized facility name, city, state and country. | [optional]
**listing_ordinal** | Option<**i32**> | 1 for the first listing of the site on the trial; higher for a repeat listing of the same site on the same trial. | [optional]
**facility_name** | Option<**String**> | The facility name as filed. | [optional]
**city** | Option<**String**> | City. | [optional]
**state** | Option<**String**> | State or region. | [optional]
**country** | Option<**String**> | Country. | [optional]
**postal** | Option<**String**> | Postal code. | [optional]
**latitude** | Option<**f64**> | Latitude of the site, in degrees. | [optional]
**longitude** | Option<**f64**> | Longitude of the site, in degrees. | [optional]
**geo_precision** | Option<**String**> | How exact the coordinate is (postal, city and so on). | [optional]
**coord_source** | Option<**String**> | Where the coordinate came from: aact, geonames or als. | [optional]
**trial_status** | Option<**String**> | The trial's overall status code. | [optional]
**trial_start_year** | Option<**i32**> | Year the trial started. | [optional]
**trial_title** | Option<**String**> | The trial's brief title. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


