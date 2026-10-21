# SupplierSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**site_id** | Option<**String**> | The supplier site id. | [optional]
**provider_id** | Option<**String**> | The supplier id: the Scientist.com provider id as text. | [optional]
**supplier_name** | Option<**String**> | The supplier's name as listed on the marketplace. | [optional]
**category** | Option<**String**> | The kind of site as the supplier recorded it (Headquarters, Laboratory, Warehouse, Other). Examples, not a closed set. | [optional]
**city** | Option<**String**> | City as recorded. | [optional]
**state** | Option<**String**> | State or province as recorded. | [optional]
**country** | Option<**String**> | Country as recorded. | [optional]
**latitude** | Option<**f64**> | Latitude of the site. | [optional]
**longitude** | Option<**f64**> | Longitude of the site. | [optional]
**company_type** | Option<**String**> | The supplier's company type, repeated on each of its sites. | [optional]
**service_areas** | Option<**Vec<String>**> | The supplier's service areas, repeated on each of its sites. Filter with 'Biology' = ANY(service_areas). | [optional]
**hq_country** | Option<**String**> | ISO 3166-1 alpha-2 country code of the supplier's headquarters point, repeated on each of its sites. | [optional]
**hq_region** | Option<**String**> | The broad world region of the supplier's headquarters country, repeated on each of its sites; other when unknown. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


