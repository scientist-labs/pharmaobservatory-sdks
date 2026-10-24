# InvestigatorDebarmentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investigator_person_id** | Option<**String**> | The investigator's resolved person key, the id in the web page address (/people/<id>). The entry key of the investigators object type. | [optional]
**debarment_id** | Option<**String**> | The debarment record id. | [optional]
**debarred_name** | Option<**String**> | The debarred individual's name as the FDA list spells it. | [optional]
**effective_date** | Option<[**String**](string.md)> | The debarment's effective date. | [optional]
**term** | Option<**String**> | The debarment term as the FDA states it. | [optional]
**fr_citation** | Option<**String**> | The Federal Register citation. | [optional]
**source_url** | Option<**String**> | A link to the source record. | [optional]
**match_tier** | Option<**String**> | How the debarred individual was matched to the investigator: low is a bare-name match that reaches exactly one person (the only tier published; an ambiguous or unresolved name has no person and no row). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


