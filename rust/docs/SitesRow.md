# SitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**site_id** | Option<**String**> | The site key: one facility spelling from the registry; the id the site page address carries. | [optional]
**site_name** | Option<**String**> | The facility name as filed in the registry. | [optional]
**city** | Option<**String**> | City as filed. | [optional]
**state** | Option<**String**> | State or province as filed. | [optional]
**country** | Option<**String**> | Country name as filed. | [optional]
**latitude** | Option<**f64**> | Latitude of the finest coordinate the record has. | [optional]
**longitude** | Option<**f64**> | Longitude of the finest coordinate the record has. | [optional]
**geo_precision** | Option<**String**> | How finely the coordinate places the site (city, postal and so on). | [optional]
**coord_source** | Option<**String**> | The source of the coordinate: aact (the registry's own), geonames (GeoNames) or als (Amazon Location Service). | [optional]
**is_location_source_asserted_only** | Option<**bool**> | True when the city the registry gave is a location that only the registry asserts for the institution. | [optional]
**site_org_id** | Option<**String**> | Legal-entity level: the organization record the site's institution resolved to (null when unresolved). | [optional]
**site_company_id** | Option<**String**> | Company level: the company the site's institution belongs to (null when unresolved). | [optional]
**site_company_name** | Option<**String**> | Display name of the site's institution's company. | [optional]
**site_group_id** | Option<**String**> | Group level: the site's institution's current ultimate parent company; a root company's group is itself. | [optional]
**site_group_name** | Option<**String**> | Display name of the site's institution's group. | [optional]
**n_trials** | Option<**i64**> | Registered trials that list the site. | [optional]
**n_active** | Option<**i64**> | Those trials whose status is ongoing. | [optional]
**first_year** | Option<**i32**> | Year of the earliest trial start at the site. | [optional]
**last_year** | Option<**i32**> | Year of the latest trial start at the site. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


