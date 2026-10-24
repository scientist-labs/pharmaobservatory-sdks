# SuppliersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**provider_id** | Option<**String**> | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/<id>). A supplier is not a company. | [optional]
**name** | Option<**String**> | The supplier's name as listed on the marketplace. | [optional]
**company_type** | Option<**String**> | The supplier's company type as the marketplace spells it (small-business, academic-core and so on). An open set; null when the supplier lists none. | [optional]
**marketplace_tier** | Option<**String**> | The supplier's marketplace tier word; null when the marketplace assigns none. | [optional]
**service_areas** | Option<**Vec<String>**> | The service areas the supplier lists, one array element each (empty when none). Filter with 'Biology' = ANY(service_areas). | [optional]
**hq_country** | Option<**String**> | ISO 3166-1 alpha-2 country code of the supplier's headquarters point (null when unknown). | [optional]
**hq_region** | Option<**String**> | The broad world region of hq_country, from the country_regions reference table; other when the country is unknown or in none of the six regions. | [optional]
**hq_latitude** | Option<**f64**> | Latitude of the supplier's headquarters point: the Headquarters site, else the first site by category then site id. | [optional]
**hq_longitude** | Option<**f64**> | Longitude of the supplier's headquarters point. | [optional]
**n_sites** | Option<**i64**> | How many located sites the supplier has. | [optional]
**n_countries** | Option<**i64**> | How many distinct countries the supplier's sites are in. | [optional]
**is_cro** | Option<**bool**> | True when the supplier's organization is a contract research organization (CRO) under the one organization-level definition. | [optional]
**is_trial_sponsor** | Option<**bool**> | True when the supplier name matches a registered trial sponsor. A name match, not a verified identity. | [optional]
**n_trials** | Option<**i64**> | Registered trials where an organization matched to the supplier by name is lead sponsor or collaborator. | [optional]
**n_lead_trials** | Option<**i64**> | The subset of n_trials where the matched organization is the lead sponsor. | [optional]
**website** | Option<**String**> | The website the supplier listed; a third-party address, not vouched for. | [optional]
**profile_url** | Option<**String**> | The supplier's page on the Scientist.com marketplace; the marketplace asks a visitor to sign in. | [optional]
**logo_path** | Option<**String**> | The path of the supplier's logo on our own origin (/supplier-logos/<sha>.png); null when it has none. | [optional]
**has_description** | Option<**bool**> | True when the supplier wrote a profile description. The description text is not published. | [optional]
**is_veteran_owned** | Option<**bool**> | SAM.gov self-certified veteran-owned designation. | [optional]
**is_woman_owned** | Option<**bool**> | SAM.gov self-certified woman-owned designation. | [optional]
**is_minority_owned** | Option<**bool**> | SAM.gov self-certified minority-owned designation (8(a) or small disadvantaged business). | [optional]
**is_hubzone** | Option<**bool**> | SAM.gov self-certified HUBZone (historically underutilized business zone) designation. | [optional]
**supplier_org_id** | Option<**String**> | Legal-entity level: the organization record the supplier resolved to. Null when the supplier has no organization record of its own. | [optional]
**supplier_company_id** | Option<**String**> | Company level: the company the supplier's organization belongs to (null when none). | [optional]
**supplier_company_name** | Option<**String**> | Display name of the supplier's company. | [optional]
**supplier_group_id** | Option<**String**> | Group level: the supplier's company's current ultimate parent company; a root company's group is itself. | [optional]
**supplier_group_name** | Option<**String**> | Display name of the supplier's group. | [optional]
**org_link_method** | Option<**String**> | How the supplier was linked to its organization record, so a caller can apply a confidence floor. The supplier pages on organization pages show only provider_ror_exact, provider_edgar_exact and supplier_curated. | [optional]
**phs_animal_welfare_assurance_id** | Option<**String**> | The NIH Office of Laboratory Animal Welfare (OLAW) Public Health Service Animal Welfare Assurance id the supplier page badges (#5286). Set only when the supplier is linked to its organization with confidence (provider_ror_exact, provider_edgar_exact or supplier_curated) and the OLAW institution matched that organization at a high or medium tier, the same rule as the page. Null means no badge, never 'not assured': OLAW's list is matched to the organizations we hold, and most institutions do not resolve. Full detail, with the raw status code, is the animal_welfare_assurances dataset. | [optional]
**has_public_page** | Option<**bool**> | True when the supplier has a public page on the web (every row of this dataset does). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


