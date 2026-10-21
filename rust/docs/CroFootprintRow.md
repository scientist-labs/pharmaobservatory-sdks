# CroFootprintRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**vendor_company_id** | Option<**String**> | Company level: the contract research organization's company. | [optional]
**vendor_company_name** | Option<**String**> | Display name of the contract research organization's company. | [optional]
**vendor_group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**vendor_group_name** | Option<**String**> | Display name of the company's group. | [optional]
**country** | Option<**String**> | The country name as the registry's facility roster spells it. | [optional]
**n_trials** | Option<**i32**> | Distinct trials the company delivered with a facility in the country. | [optional]
**n_sites** | Option<**i32**> | Distinct sites (facility keys) in the country across those trials. | [optional]
**latitude** | Option<**f64**> | Latitude of the country's approximate centroid. Null for a country with no centroid. | [optional]
**longitude** | Option<**f64**> | Longitude of the country's approximate centroid. Null for a country with no centroid. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


