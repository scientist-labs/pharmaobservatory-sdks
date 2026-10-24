# AnimalWelfareAssurancesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**institution_org_id** | Option<**String**> | Legal-entity level: the organization record the OLAW institution resolved to; the entry key for the legal_entities object type. | [optional]
**institution_company_id** | Option<**String**> | Company level: the organization's company (null when the organization is not in a company); the entry key for the companies object type. | [optional]
**institution_company_name** | Option<**String**> | Display name of the organization's company. | [optional]
**institution_group_id** | Option<**String**> | Group level: the organization's current ultimate parent; a root's group is itself. | [optional]
**institution_group_name** | Option<**String**> | Display name of the organization's group. | [optional]
**source** | Option<**String**> | The animal-welfare source: nih_olaw today. A second source (the Canadian Council on Animal Care list) can be added as another value without changing the grain. | [optional]
**assurance_id** | Option<**String**> | The OLAW assurance id of the organization's best evidence (a primary assurance beats a covered component). | [optional]
**assurance_status_code** | Option<**String**> | OLAW's raw two-digit assurance status code as published (for example 02 or 03). Never decoded, because no OLAW page documents it. | [optional]
**assurance_kind** | Option<**String**> | OLAW's PRIMARY_OR_COVERED_COMP field as published: PRIMARY or COVERED_COMP (a covered component of a primary assurance). | [optional]
**is_domestic** | Option<**bool**> | True for a domestic (US) assurance, false for a foreign one, from OLAW's assurance type (D or F). | [optional]
**match_tier** | Option<**String**> | How well the OLAW institution name reached the organization: high (name, country and state agree), medium (name and country agree) or low (name only). The web badge shows high and medium. | [optional]
**assurance_count** | Option<**i64**> | How many OLAW assurance rows resolved to this organization (a university and its covered components). | [optional]
**assurance_ids** | Option<**Vec<String>**> | Every OLAW assurance id that resolved to this organization, sorted. | [optional]
**institution_names** | Option<**Vec<String>**> | The institution names as OLAW lists them, for every assurance that resolved to this organization. | [optional]
**source_asof** | Option<**String**> | When OLAW's lookup tool says its list was last updated; the source does not state a timezone. | [optional]
**source_credit** | Option<**String**> | The plain source credit line to show beside any figure from this dataset. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


