# InvestigatorTrialsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investigator_person_id** | Option<**String**> | The investigator's resolved person key, the id in the web page address (/people/<id>). The entry key of the investigators object type. | [optional]
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional]
**roles** | Option<**Vec<String>**> | The roles the person holds on this trial as registered. | [optional]
**is_principal_investigator** | Option<**bool**> | True when the person is named principal investigator on this trial. | [optional]
**trial_title** | Option<**String**> | The trial's brief title. | [optional]
**phase_label** | Option<**String**> | The trial phase as a display label. | [optional]
**overall_status** | Option<**String**> | The registry's overall recruitment status. | [optional]
**study_type** | Option<**String**> | The registry's study type. | [optional]
**start_date** | Option<[**String**](string.md)> | Trial start date as registered. | [optional]
**start_year** | Option<**i32**> | Calendar year of the start date. | [optional]
**completion_date** | Option<[**String**](string.md)> | Trial completion date as registered. | [optional]
**enrollment** | Option<**i32**> | Enrollment as registered. | [optional]
**n_sites** | Option<**i64**> | The trial's distinct geocoded sites, the count the people page shows (0 when none is geocoded). | [optional]
**single_site_id** | Option<**String**> | The site key when the trial lists exactly one site (null otherwise); joins the sites dataset. | [optional]
**single_site_name** | Option<**String**> | The single site's facility name as the registry lists it. | [optional]
**single_site_city** | Option<**String**> | The single site's city. | [optional]
**single_site_state** | Option<**String**> | The single site's state or province. | [optional]
**single_site_country** | Option<**String**> | The single site's country. | [optional]
**sponsor_name** | Option<**String**> | The trial's lead sponsor as filed. | [optional]
**sponsor_org_id** | Option<**String**> | Legal-entity level: the organization record the lead sponsor resolved to (null when unresolved). | [optional]
**sponsor_company_id** | Option<**String**> | Company level: the company the lead sponsor belongs to (null when unresolved). | [optional]
**sponsor_company_name** | Option<**String**> | Display name of the sponsor's company. | [optional]
**sponsor_group_id** | Option<**String**> | Group level: the sponsor's current ultimate parent company; a root company's group is itself. | [optional]
**sponsor_group_name** | Option<**String**> | Display name of the sponsor's group. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


