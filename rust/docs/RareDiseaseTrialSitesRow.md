# RareDiseaseTrialSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**mapping_arm** | Option<**String**> | exact or unspecified: the mapping arm the site's trials belong to. | [optional]
**site_id** | Option<**String**> | The site key: one facility spelling from the registry; joins the sites dataset. | [optional]
**site_rank** | Option<**i32**> | The site's rank by trial count within its disease and arm (1 is the most trials); a position in a list, not a quality score. | [optional]
**site_name** | Option<**String**> | The facility name as filed in the registry. | [optional]
**city** | Option<**String**> | The facility's city as filed. | [optional]
**state** | Option<**String**> | The facility's state or region as filed. | [optional]
**country** | Option<**String**> | The facility's country as filed. | [optional]
**latitude** | Option<**f64**> | The site's latitude in degrees; null when no coordinate is known. | [optional]
**longitude** | Option<**f64**> | The site's longitude in degrees; null when no coordinate is known. | [optional]
**coord_source** | Option<**String**> | The source of the coordinate: aact (the registry's own), geonames (GeoNames) or als (Amazon Location Service). | [optional]
**condition_ids** | Option<**Vec<String>**> | The MeSH condition ids through which the disease reaches the site's trials. | [optional]
**n_trials** | Option<**i32**> | Registered trials of this disease and arm that list the site. | [optional]
**n_trials_direct** | Option<**i32**> | Those trials whose own condition list names the mapped condition. | [optional]
**n_active_trials** | Option<**i32**> | Those trials that are recruiting, not yet recruiting, active or enrolling by invitation. | [optional]
**first_trial_year** | Option<**i32**> | The earliest start year among those trials. | [optional]
**last_trial_year** | Option<**i32**> | The latest start year among those trials. | [optional]
**last_update_posted_date** | Option<[**String**](string.md)> | The latest registry update among those trials. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


