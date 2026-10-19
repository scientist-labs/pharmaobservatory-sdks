# CompaniesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**company_id** | Option<**String**> | The company key (company:<key>). | [optional]
**company_name** | Option<**String**> | The company's display name: its own legal entity's name where one is known. | [optional]
**name_basis** | Option<**String**> | Which source supplied the name: gleif_legal_name, edgar_registrant, ror or sponsor_display_name. | [optional]
**country** | Option<**String**> | Country of the company's own legal entity, as the source states it. | [optional]
**org_type** | Option<**String**> | The organization type of the company's representative record. | [optional]
**org_kinds** | Option<**Vec<String>**> | The roles the company's organizations play (Manufacturer, Sponsor and so on). | [optional]
**edgar_cik** | Option<**String**> | SEC EDGAR central index key of the representative record, when it files. | [optional]
**edgar_ticker** | Option<**String**> | Exchange ticker of the representative record, when listed. | [optional]
**representative_org_id** | Option<**String**> | Legal-entity level: the organization record the company page shows. | [optional]
**parent_company_id** | Option<**String**> | The direct owner company (part_of), null for a root. | [optional]
**parent_company_name** | Option<**String**> | Display name of the direct owner company. | [optional]
**part_of_since** | Option<[**String**](string.md)> | Date the ownership link began, when every hop is dated. | [optional]
**part_of_since_precision** | Option<**String**> | Precision of part_of_since (day, month or year). | [optional]
**group_id** | Option<**String**> | Group level: the current ultimate parent company; a root company's group is itself. Current ownership only; dated links are in company_members. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**rollup_depth** | Option<**i32**> | Hops from this company to its ultimate parent (0 for a root). | [optional]
**n_legal_entities** | Option<**i64**> | Registered legal entities that are members of the company. | [optional]
**n_member_companies** | Option<**i64**> | Companies folded into this one. | [optional]
**num_lead_trials** | Option<**i64**> | Trials the company's organizations lead. | [optional]
**num_collaborator_trials** | Option<**i64**> | Trials the company's organizations collaborate on. | [optional]
**num_trials** | Option<**i64**> | Trials the company's organizations lead or collaborate on. | [optional]
**total_lead_enrollment** | Option<**i64**> | Planned enrollment across led trials. | [optional]
**num_recent_lead_trials** | Option<**i64**> | Led trials started in the recent window. | [optional]
**activity_score** | Option<**f64**> | Trial activity score (interventional lead trials and enrollment, weighted toward recent). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


