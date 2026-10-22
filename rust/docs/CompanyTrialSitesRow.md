# CompanyTrialSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | The company key (company:<key>) of the lead sponsor. | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the current ultimate parent company; a root company's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**roster_root_id** | Option<**String**> | The Clinical Research Map's organization key for these rows: the group when it is a public roster root, else the company when it is one, else null. | [optional]
**site_id** | Option<**String**> | The site key: one facility spelling from the registry; the id the site page address carries. | [optional]
**facility_name** | Option<**String**> | The facility name as filed in the registry. | [optional]
**city** | Option<**String**> | City as filed. | [optional]
**state** | Option<**String**> | State or province as filed. | [optional]
**country** | Option<**String**> | Country name as filed. | [optional]
**site_country_code** | Option<**String**> | ISO 3166-1 alpha-2 code of the site's country, from the registry's country name; null when the name resolves to no country. | [optional]
**latitude** | Option<**f64**> | Latitude of the finest coordinate the site record has. | [optional]
**longitude** | Option<**f64**> | Longitude of the finest coordinate the site record has. | [optional]
**geo_precision** | Option<**String**> | How finely the coordinate places the site (city, postal and so on). | [optional]
**coord_source** | Option<**String**> | The source of the coordinate: aact (the registry's own), geonames (GeoNames) or als (Amazon Location Service). | [optional]
**n_trials** | Option<**i32**> | Distinct trials the company leads that list the site. | [optional]
**site_rank** | Option<**i32**> | 1 is the company's busiest located site (most led trials, then site key); dense per company. | [optional]
**n_sites** | Option<**i32**> | The company's located site count, repeated on each of its rows. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


