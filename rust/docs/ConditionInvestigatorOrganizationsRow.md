# ConditionInvestigatorOrganizationsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**condition_id** | Option<**String**> | The condition's MeSH descriptor id, the id in the conditions page address (D001943). The entry key of the conditions object type. | [optional]
**condition_name** | Option<**String**> | The condition's MeSH name. | [optional]
**organization_name** | Option<**String**> | The organization's name in the registry that resolved it (ROR, else SEC EDGAR), never the investigator's affiliation text as filed. | [optional]
**organization_org_id** | Option<**String**> | Legal-entity level: the organization record the investigators' affiliation resolved to. | [optional]
**organization_company_id** | Option<**String**> | Company level: the company the organization belongs to (null when unresolved). | [optional]
**organization_company_name** | Option<**String**> | Display name of the organization's company. | [optional]
**organization_group_id** | Option<**String**> | Group level: the organization's current ultimate parent company; a root company's group is itself. | [optional]
**organization_group_name** | Option<**String**> | Display name of the organization's group. | [optional]
**n_investigators** | Option<**i64**> | Distinct investigators at the organization named on a trial of the condition, in any role. Always at least 2. | [optional]
**n_principal_investigators** | Option<**i64**> | Of n_investigators, those who hold the principal investigator role on one of the condition's trials. | [optional]
**n_investigators_total** | Option<**i64**> | Distinct investigators across the condition's trials, at any organization or none. A constant of the condition repeated on every row: never sum it. | [optional]
**n_at_resolved_organizations** | Option<**i64**> | Of n_investigators_total, those whose affiliation resolved to an organization at a high-confidence tier, listed or not. A constant of the condition. | [optional]
**n_at_unresolved_affiliations** | Option<**i64**> | Of n_investigators_total, those with no resolved organization, including those with no affiliation on file. A constant of the condition. | [optional]
**n_principal_investigators_total** | Option<**i64**> | Of n_investigators_total, those who hold the principal investigator role on one of the condition's trials. A constant of the condition. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


