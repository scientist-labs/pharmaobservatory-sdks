# CliaCertificatesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**clia_number** | Option<**String**> | The ten-character CLIA certificate number. | [optional]
**lab_name** | Option<**String**> | The laboratory's name as listed on the certificate. | [optional]
**firm_name** | Option<**String**> | The organization the certificate resolved to (its canonical name). | [optional]
**firm_org_id** | Option<**String**> | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional]
**firm_company_id** | Option<**String**> | Company level: the firm's company (null when unresolved); the entry key for the companies object type. | [optional]
**firm_company_name** | Option<**String**> | Display name of the firm's company. | [optional]
**firm_group_id** | Option<**String**> | Group level: the firm's current ultimate parent; a root's group is itself. | [optional]
**firm_group_name** | Option<**String**> | Display name of the firm's group. | [optional]
**match_tier** | Option<**String**> | How the certificate was attached to the organization: high, medium or low. | [optional]
**match_confidence** | Option<**f64**> | The matcher's confidence, 0 to 1. | [optional]
**city** | Option<**String**> | The laboratory's city. | [optional]
**state_region** | Option<**String**> | The laboratory's state. | [optional]
**postal5** | Option<**String**> | The five-digit postal code. | [optional]
**country** | Option<**String**> | The country code. | [optional]
**latitude** | Option<**f64**> | Latitude in degrees (WGS84), for the map tool. | [optional]
**longitude** | Option<**f64**> | Longitude in degrees (WGS84), for the map tool. | [optional]
**geocode_precision** | Option<**String**> | What a pin locates (an address, a postal code or a city). | [optional]
**coord_source** | Option<**String**> | Where the coordinates came from. | [optional]
**is_active** | Option<**bool**> | True when the certificate is active on the file. | [optional]
**certificate_type** | Option<**String**> | The CLIA certificate type (for example Compliance, Accreditation or Waiver). | [optional]
**complexity_ceiling** | Option<**String**> | The highest test complexity the certificate permits. | [optional]
**accreditor_claimed** | Option<**Vec<String>**> | The accreditation body the file names. | [optional]
**accreditor_confirmed** | Option<**Vec<String>**> | The accreditation body confirmed against the accreditor's own list. | [optional]
**has_cap_accreditation** | Option<**bool**> | True when the laboratory holds College of American Pathologists (CAP) accreditation. | [optional]
**has_ashi_accreditation** | Option<**bool**> | True when the laboratory holds American Society for Histocompatibility and Immunogenetics (ASHI) accreditation. | [optional]
**has_aabb_accreditation** | Option<**bool**> | True when the laboratory holds AABB (formerly American Association of Blood Banks) accreditation. | [optional]
**facility_type** | Option<**String**> | The laboratory's facility type on the file. | [optional]
**annual_test_volume** | Option<**i64**> | Reported annual test volume. | [optional]
**test_volume_source** | Option<**String**> | Where the volume figure came from. | [optional]
**lab_site_count** | Option<**i32**> | How many sites the laboratory reports. | [optional]
**pos_vintage** | Option<**String**> | The vintage of the CMS Provider of Services file. | [optional]
**as_of** | Option<[**String**](string.md)> | The build date of the figures. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


