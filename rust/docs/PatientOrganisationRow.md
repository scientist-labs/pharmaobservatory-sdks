# PatientOrganisationRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_kind** | Option<**String**> | Which kind of page the row is listed on: rare_disease (a rare disease page, entered on orpha_code) or condition (a condition page, entered on condition_id). | [optional]
**target_key** | Option<**String**> | The orpha_code of a rare_disease row or the condition_id of a condition row: the key the row is entered on. Always present, so it can stand in a sort or a cursor. | [optional]
**orpha_code** | Option<**String**> | The ORPHA code of the disease the row is listed under, the digits in the web page address (/rare-diseases/558); null on a condition row. The entry key of the rare_diseases object type. | [optional]
**condition_id** | Option<**String**> | The MeSH descriptor id of the condition the row is listed under (D followed by six to nine digits); null on a rare disease row. The entry key of the conditions object type. | [optional]
**patient_organisation_id** | Option<**String**> | The organisation's id: a spine organisation id (edgar:, ror:, sponsor: and the like) when the organisation folded onto an organisation we already hold, else pag:<registrable domain of its website>. | [optional]
**org_id** | Option<**String**> | The legal-entity id of the organisation when it folded onto one we already hold, the join to legal_entities; null for a pag: organisation, which has no spine record. | [optional]
**name** | Option<**String**> | The organisation's name as its own website states it. | [optional]
**website_url** | Option<**String**> | The organisation's website as it states it. | [optional]
**country** | Option<**String**> | The country the organisation declares (ISO 3166-1 alpha-2), else its primary location's country. | [optional]
**region** | Option<**String**> | The region of its primary location. | [optional]
**city** | Option<**String**> | The city of its primary location. | [optional]
**location_count** | Option<**i32**> | How many locations the organisation lists. | [optional]
**languages** | Option<**Vec<String>**> | The languages the organisation says it works in, as it states them. | [optional]
**phones** | Option<**Vec<String>**> | The organisation's own phone numbers, as stated. Organisation-level channels only: no person's line. | [optional]
**emails** | Option<**Vec<String>**> | The organisation's own email addresses, as stated, minus any address whose local part looks like a person's name. Organisation-level channels only. | [optional]
**contact_urls** | Option<**Vec<String>**> | The organisation's own contact page addresses. | [optional]
**social_urls** | Option<**Vec<String>**> | The organisation's own social media addresses. | [optional]
**service_codes** | Option<**Vec<String>**> | The services the organisation's own pages say it offers, from a closed list: support_groups, helpline, information_and_education, patient_registry, research_funding, advocacy_and_policy, financial_assistance, care_navigation, clinical_trial_matching, other. | [optional]
**other_service_labels** | Option<**Vec<String>**> | For a service coded other, the organisation's own words for it. Empty when it states none. | [optional]
**link_kind** | Option<**String**> | How the organisation reaches the disease or condition the row is listed under: declared (its own page names it), mapped (a condition row: it declares a rare disease this condition is mapped to; the mapped_* columns say which), broader (a rare disease row: it declared a broader parent disease) or narrower (a rare disease row: it declared a narrower child disease). A broader or narrower row never says the organisation declares THIS disease. | [optional]
**link_rank** | Option<**i32**> | The page's group order: 1 declared, 2 mapped, 3 broader, 4 narrower. | [optional]
**declared_scope** | Option<**String**> | How much of its work the organisation says the disease is: sole_focus, one_of_several or umbrella; null when it states none. | [optional]
**edge_verdict** | Option<**String**> | Whether the organisation's pages were judged to SERVE people with the disease it declared: serves, unclear (asked and not callable) or unjudged (not yet asked). Only unclear and unjudged rows carry a 'not yet confirmed' label on the page; an edge judged to be a mention only is not a row. A judgement of ours about the organisation's pages, not the organisation's own statement. On a broader, narrower or mapped row it describes the declared disease. | [optional]
**declared_name** | Option<**String**> | The disease name the organisation actually declared. On a broader, narrower or mapped row it is NOT this page's disease. | [optional]
**mapped_orpha_code** | Option<**String**> | On a mapped row, the ORPHA code of the declared rare disease the condition is mapped to; null otherwise. | [optional]
**mapped_disease_name** | Option<**String**> | On a mapped row, the preferred name of that rare disease; null otherwise. | [optional]
**mapping_relation** | Option<**String**> | On a mapped row, how the rare disease relates to the condition (exact, broader or narrower); null otherwise. | [optional]
**source_url** | Option<**String**> | The page of the organisation's own website the declaration came from. | [optional]
**retrieved_on** | Option<[**String**](string.md)> | The date that page was retrieved. | [optional]
**credit** | Option<**String**> | The source credit line for the row: Source: <organisation name>, <page address>, retrieved <date>. Mandatory wherever the row is shown. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


