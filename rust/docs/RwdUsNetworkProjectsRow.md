# RwdUsNetworkProjectsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**project_id** | Option<**String**> | The N3C project id (a Data Use Request uid), the grain. | [optional]
**network_key** | Option<**String**> | The network key, the entry key of the rwd_us_networks object type. | [optional]
**title** | Option<**String**> | The project title as N3C publishes it. | [optional]
**institution_name** | Option<**String**> | The accessing institution as published. | [optional]
**institution_name_key** | Option<**String**> | The normalized name key; joins to rwd_us_network_memberships.institution_name_key. | [optional]
**pi_name** | Option<**String**> | Personal data: the principal investigator's name as N3C's public dashboard publishes it. Landed whole, tagged, never a gate (#5143). | [optional]
**workspace_status** | Option<**String**> | The project's workspace status as N3C publishes it. | [optional]
**project_type** | Option<**String**> | The project type as N3C publishes it. | [optional]
**is_task_team** | Option<**bool**> | Whether N3C marks the project as a task team. | [optional]
**contains_personal_data** | Option<**bool**> | True wherever pi_name is populated. A tag, never a gate (#5143). | [optional]
**personal_data_fields** | Option<**String**> | The personal-data columns of the row, as a JSON list: [\"pi_name\"]. | [optional]
**as_of** | Option<[**String**](string.md)> | The retrieval date of the roster snapshot. | [optional]
**source_url** | Option<**String**> | The API endpoint the roster came from. | [optional]
**retrieved_at** | Option<**String**> | When the snapshot was retrieved, in UTC. | [optional]
**org_id** | Option<**String**> | The organization spine id of the accessing institution, or null. Carried, never minted. | [optional]
**org_name** | Option<**String**> | The spine organization's name. | [optional]
**institution_ror_id** | Option<**String**> | The bare ROR id N3C published for the institution where it is a network member; null otherwise. | [optional]
**org_match_rung** | Option<**String**> | How the institution reached the organization: ror_exact, xref, match_keys, ambiguous or unresolved. | [optional]
**org_match_tier** | Option<**String**> | The tier of the organization match. | [optional]
**org_match_confidence** | Option<**f64**> | The organization match confidence. | [optional]
**institution_is_network_member** | Option<**bool**> | Whether the accessing institution itself holds a published agreement. | [optional]
**pi_person_id** | Option<**String**> | A resolved person on the spine, by name plus organization; see pi_match_method. The entry key of the investigators object type; null when unresolved. | [optional]
**pi_match_method** | Option<**String**> | How the PI was matched: name_org_spine (name and organization agreed with one spine person), name_org_reporter (a unique NIH RePORTER PI at the same organization), ambiguous (more than one candidate, never picked) or unresolved. | [optional]
**pi_corroborated** | Option<**bool**> | False on every row: no shared trial or paper corroborates the name-plus-organization match. | [optional]
**pi_reporter_pi_id** | Option<**String**> | NIH RePORTER's persistent PI_ID, an external key published as a labelled attribute, never as a resolved person; set where a unique RePORTER PI shares the name and organization. | [optional]
**candidate_person_count** | Option<**i32**> | On an ambiguous row, the number of distinct spine persons the evidence could not choose between; otherwise 0. | [optional]
**also_matched_by_reporter** | Option<**bool**> | True on a name_org_spine row whose person the RePORTER arm also reached. | [optional]
**project_total** | Option<**i32**> | Projects in the roster, the denominator of the resolution rate. | [optional]
**projects_with_institution_on_spine** | Option<**i32**> | Projects whose institution resolved to a spine organization. | [optional]
**projects_with_resolved_pi** | Option<**i32**> | Projects whose PI resolved to a spine person; over project_total it is the measured resolution rate. | [optional]
**projects_with_reporter_pi_id** | Option<**i32**> | Projects carrying a RePORTER PI_ID. | [optional]
**attribution** | Option<**String**> | The credit line every surface carrying this data prints. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


