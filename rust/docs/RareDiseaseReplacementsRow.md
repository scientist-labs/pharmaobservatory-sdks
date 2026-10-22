# RareDiseaseReplacementsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**replacement_assertion_id** | Option<**String**> | The identifier of the replacement assertion. | [optional]
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**preferred_name** | Option<**String**> | The deprecated disease's preferred English name. | [optional]
**replacement_orpha_code** | Option<**String**> | The ORPHA code of the concept that replaces the deprecated one. | [optional]
**replacement_preferred_name** | Option<**String**> | The replacement concept's preferred English name. | [optional]
**replacement_status** | Option<**String**> | The replacement concept's lifecycle state. | [optional]
**predicate** | Option<**String**> | moved_to when Orphanet moved the concept to the replacement, referred_to when it referred the concept there. | [optional]
**assertion_source** | Option<**String**> | The Orphanet file the assertion was read from (ordo or orphadata_product1). | [optional]
**source_release** | Option<**String**> | The Orphadata Science release the row was read from. | [optional]
**source_url** | Option<**String**> | The URL of the source file the row was read from. | [optional]
**license_name** | Option<**String**> | The licence the assertion's file is published under. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


