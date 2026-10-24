# InvestigatorInspectionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investigator_person_id** | Option<**String**> | The investigator's resolved person key, the id in the web page address (/people/<id>). The entry key of the investigators object type. | [optional]
**bimo_record_id** | Option<**String**> | The record id of the inspection and deficiency code. | [optional]
**bimo_investigator_id** | Option<**String**> | The CLIIL investigator code; two codes under one name may be different people. | [optional]
**inspected_investigator_name** | Option<**String**> | The inspected investigator's name as CLIIL files it ('last, first'). | [optional]
**inspection_date** | Option<[**String**](string.md)> | The inspection's date. | [optional]
**inspection_type** | Option<**String**> | The inspection type code as CLIIL files it (DA or FC). | [optional]
**inspection_classification** | Option<**String**> | The FDA inspection classification: NAI (no action indicated), VAI (voluntary action indicated) or OAI (official action indicated). | [optional]
**deficiency_code** | Option<**String**> | The deficiency code as CLIIL codes it (null when the record carries none). | [optional]
**has_deficiency** | Option<**bool**> | True when the inspection cited a deficiency. | [optional]
**site_name** | Option<**String**> | The inspected site's name. | [optional]
**site_city** | Option<**String**> | The inspected site's city. | [optional]
**site_state** | Option<**String**> | The inspected site's state. | [optional]
**site_country** | Option<**String**> | The inspected site's country. | [optional]
**match_tier** | Option<**String**> | How the CLIIL investigator was matched to the person: low is a bare-name match that reaches exactly one person (the only tier published; an ambiguous or unresolved name has no person and no row). | [optional]
**match_confidence** | Option<**f64**> | The match's confidence (0.45 for the low tier). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


