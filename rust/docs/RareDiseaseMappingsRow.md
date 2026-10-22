# RareDiseaseMappingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**xref_assertion_id** | Option<**String**> | The identifier of the cross-reference assertion. | [optional]
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**target_namespace** | Option<**String**> | The vocabulary the target code belongs to (ICD-10, ICD-11, MESH, MONDO, UMLS, GARD, ORPHA and others). | [optional]
**target_code** | Option<**String**> | The code in the target vocabulary. | [optional]
**target_curie** | Option<**String**> | The target as a prefixed identifier (namespace and code). | [optional]
**mapping_relation** | Option<**String**> | The relation the source states from the disease to the target: exact, narrower, broader, undetermined or unspecified. | [optional]
**assertion_source** | Option<**String**> | The file the assertion was read from: ordo, orphadata_product1 or mondo. | [optional]
**confidence** | Option<**f64**> | The source's confidence in the assertion, where it states one. | [optional]
**condition_id** | Option<**String**> | The MeSH descriptor id of the matching condition, on MESH rows; the entry key of the conditions dataset. | [optional]
**condition_name** | Option<**String**> | The matching condition's name, on MESH rows. | [optional]
**source_release** | Option<**String**> | The Orphadata Science release the row was read from. | [optional]
**source_url** | Option<**String**> | The URL of the source file the row was read from. | [optional]
**license_name** | Option<**String**> | The licence the assertion's file is published under. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


