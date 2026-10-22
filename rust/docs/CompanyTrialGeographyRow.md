# CompanyTrialGeographyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | The company key (company:<key>). | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the current ultimate parent company; a root company's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**is_public_roster_root** | Option<**bool**> | True when the company is a public roster root: an organization the public Clinical Research Map lists. | [optional]
**roster_root_id** | Option<**String**> | The Clinical Research Map's organization key for this company: the group when it is a public roster root, else the company when it is one, else null. | [optional]
**roster_lead_trials** | Option<**i32**> | For a public roster root, the lead-trial count the public organizations list sorts on; null for a company that is not one. | [optional]
**domicile_country** | Option<**String**> | ISO 3166-1 alpha-2 code of the company's legal-entity domicile; null when unknown. | [optional]
**site_country** | Option<**String**> | ISO 3166-1 alpha-2 code of the country holding the most of the company's own located led-trial sites; null when none resolves. | [optional]
**map_country** | Option<**String**> | The country the map places the company under: domicile_country, else site_country; null when neither is known. | [optional]
**map_country_basis** | Option<**String**> | Which signal map_country came from: domicile or trial_sites; null exactly when map_country is null. | [optional]
**n_located_sites** | Option<**i32**> | The company's own located led-trial sites. | [optional]
**n_unlocated_sites** | Option<**i32**> | The company's led-trial sites with no coordinate, which the map does not draw. | [optional]
**n_site_countries** | Option<**i32**> | Distinct resolved countries among the company's located led-trial sites. | [optional]
**n_unresolved_sites** | Option<**i32**> | Located led-trial sites whose country name resolves to no country code; they do not vote for site_country. | [optional]
**roster_map_country** | Option<**String**> | For a public roster root, the country the public Clinical Research Map shows the organization under, as published; null for a non-root. | [optional]
**roster_map_country_basis** | Option<**String**> | For a public roster root, which signal the public map's country came from: domicile or trial_sites; null for a non-root or an unknown country. | [optional]
**roster_n_located_sites** | Option<**i32**> | For a public roster root, the located sites the public map draws the organization at (its group fold); null for a non-root. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


