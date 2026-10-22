# HospitalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ccn** | Option<**String**> | CMS Certification Number, the six-character Medicare provider id the hospital page address carries. | [optional]
**hospital_name** | Option<**String**> | The hospital's facility name in the CMS registry. | [optional]
**city** | Option<**String**> | The hospital's city. | [optional]
**state** | Option<**String**> | The hospital's two-letter state or territory code. | [optional]
**bed_count** | Option<**i32**> | Certified bed count in the latest registry file. Null when the file has no bed record. | [optional]
**hospital_type** | Option<**String**> | The CMS hospital type. | [optional]
**ownership** | Option<**String**> | The CMS ownership category. | [optional]
**is_teaching** | Option<**bool**> | Whether the hospital is a teaching hospital. | [optional]
**emergency_services** | Option<**String**> | Whether the hospital provides emergency services, as CMS files it. Null where CMS left it blank. | [optional]
**latitude** | Option<**f64**> | Latitude of the hospital's own coordinate. Null exactly when geocode_method is unmatched. | [optional]
**longitude** | Option<**f64**> | Longitude of the hospital's own coordinate. Null exactly when geocode_method is unmatched. | [optional]
**geocode_method** | Option<**String**> | How the coordinate was placed: als (Amazon Location Service address match), zip_centroid (coarse ZIP-code centroid) or unmatched (no coordinate). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


