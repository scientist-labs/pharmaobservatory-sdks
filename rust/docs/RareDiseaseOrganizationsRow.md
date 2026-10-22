# RareDiseaseOrganizationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**connection_kind** | Option<**String**> | disease_relation when connection_value is a relation to the disease, disease_function when it is a function the organization performs for it. | [optional]
**assertion_id** | Option<**String**> | The identifier of the assertion. | [optional]
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**disease_name** | Option<**String**> | The disease's preferred English name. | [optional]
**organization_name** | Option<**String**> | The organization's name as filed. | [optional]
**organization_org_id** | Option<**String**> | Legal-entity level: the organization record the assertion is on. | [optional]
**organization_company_id** | Option<**String**> | Company level: the organization's company (null when unresolved). | [optional]
**organization_company_name** | Option<**String**> | Display name of the organization's company. | [optional]
**organization_group_id** | Option<**String**> | Group level: the organization's current ultimate parent company; a root company's group is itself. | [optional]
**organization_group_name** | Option<**String**> | Display name of the organization's group. | [optional]
**connection_value** | Option<**String**> | The relation (a predicate) or the function, according to connection_kind. | [optional]
**evidence_type** | Option<**String**> | The kind of evidence the assertion rests on. | [optional]
**authority_class** | Option<**String**> | How authoritative the source is for the claim. | [optional]
**source_key** | Option<**String**> | The source the assertion was read from. | [optional]
**source_record_id** | Option<**String**> | The record in the source the assertion comes from. | [optional]
**source_url** | Option<**String**> | A link to the source record, where it has one. | [optional]
**evidence_trial_id** | Option<**String**> | The registry trial id the assertion names, where it names one. | [optional]
**observed_at** | Option<**String**> | When the assertion was observed in the source (UTC). | [optional]
**last_verified_at** | Option<**String**> | When the assertion was last verified against the source (UTC). | [optional]
**confidence** | Option<**f64**> | The derivation's confidence in the assertion. | [optional]
**license_name** | Option<**String**> | The licence of the source the assertion was read from. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


