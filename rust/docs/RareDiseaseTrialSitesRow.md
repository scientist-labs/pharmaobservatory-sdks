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
**latitude** | Option<**f64**> | The site's distance-ready latitude in degrees: one real filed location, the one most of the site's registry filings agree on, the same point the site has on the disease page map and in the nearest-site list. Null when the site is not mappable (its filings disagree, or are coarser than a city). Not the mean of the site's locations. | [optional]
**longitude** | Option<**f64**> | The site's distance-ready longitude in degrees, paired with latitude; null exactly when latitude is. | [optional]
**coord_source** | Option<**String**> | The source of the latitude and longitude: aact (the registry's own), geonames (GeoNames) or als (Amazon Location Service); null when latitude is. | [optional]
**condition_ids** | Option<**Vec<String>**> | The MeSH condition ids through which the disease reaches the site's trials. | [optional]
**n_trials** | Option<**i32**> | Registered trials of this disease and arm that list the site. | [optional]
**n_trials_direct** | Option<**i32**> | Those trials whose own condition list names the mapped condition. | [optional]
**n_active_trials** | Option<**i32**> | Those trials that are recruiting, not yet recruiting, active or enrolling by invitation. | [optional]
**n_open_trials_direct** | Option<**i32**> | Those directly tagged trials that are recruiting or not yet recruiting, the statuses a person can volunteer for. A subset of n_trials_direct; active but closed and invitation-only trials are not counted. | [optional]
**first_trial_year** | Option<**i32**> | The earliest start year among those trials. | [optional]
**last_trial_year** | Option<**i32**> | The latest start year among those trials. | [optional]
**last_update_posted_date** | Option<[**String**](string.md)> | The latest registry update among those trials. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


