# OwnershipPositionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**position_id** | Option<**String**> | The position id. For a position whose owner name is withheld it is an opaque handle that derives nothing from the owner and is stable only while the other positions of the same owned company, filing family and dimension are unchanged. | [optional]
**source_class** | Option<**String**> | The filing family: sec_forms_3_4_5 (insiders), sec_schedules_13d_g (beneficial owners), sec_form_adv (investment adviser control persons) or sec_def_14a (proxy-statement tables). | [optional]
**ownership_dimension** | Option<**String**> | What the position measures: insider_reporting, beneficial_ownership, voting_power, economic_interest or control. Positions of different dimensions are different measurements and are not added. | [optional]
**relationship_type** | Option<**String**> | The owner's relationship to the company as the filing family states it. | [optional]
**subject_name** | Option<**String**> | The owned company as filed. | [optional]
**subject_org_id** | Option<**String**> | Legal-entity level: the owned company's organization record. | [optional]
**subject_company_id** | Option<**String**> | Company level: the owned company (null when unresolved). | [optional]
**subject_company_name** | Option<**String**> | Display name of the owned company. | [optional]
**subject_group_id** | Option<**String**> | Group level: the owned company's current ultimate parent; a root's group is itself. | [optional]
**subject_group_name** | Option<**String**> | Display name of the owned company's group. | [optional]
**owner_name** | Option<**String**> | The owner's name as filed. Null for an owner whose filing does not mark the name public (see owner_name_is_public). | [optional]
**owner_name_is_public** | Option<**bool**> | True when the filing marks the owner's name public. When false every column that could identify the owner is null (name, identity key, person id, investor id, organization and company columns, source_url and the dates); the position and its figures stay. | [optional]
**owner_kind** | Option<**String**> | What the owner is: organization, natural_person, reporting_group or unknown. | [optional]
**owner_identity_key** | Option<**String**> | The owner's stable identity (an SEC registration number or an adviser registration number); null when the name is withheld. | [optional]
**owner_person_id** | Option<**String**> | The key of an owner who is a natural person; null for other owners and when the name is withheld. | [optional]
**owner_investor_id** | Option<**String**> | The key of the owner in the investor directory when it is one; null when the name is withheld. | [optional]
**owner_org_id** | Option<**String**> | Legal-entity level: the owner's organization record (null when the owner is not an organization, is unresolved or is withheld). | [optional]
**owner_company_id** | Option<**String**> | Company level: the owner's company (null when unresolved or when the name is withheld). | [optional]
**owner_company_name** | Option<**String**> | Display name of the owner's company. | [optional]
**owner_group_id** | Option<**String**> | Group level: the owner's current ultimate parent; a root's group is itself (null when the owner is withheld). | [optional]
**owner_group_name** | Option<**String**> | Display name of the owner's group. | [optional]
**shares_basis** | Option<**String**> | How the share figure was reached: sum_standing_balances (insider balances), group_reported_maximum (a joint filing's maximum, never a sum of its filers) or no_quantity (the source gives none). Read a share figure with its basis. | [optional]
**shares** | Option<**f64**> | Shares the owner holds, on the shares_basis. Null when the source gives no quantity. | [optional]
**shares_direct** | Option<**f64**> | Of those shares, the part held directly. | [optional]
**shares_indirect** | Option<**f64**> | Of those shares, the part held indirectly. | [optional]
**derivative_underlying_shares** | Option<**f64**> | Shares underlying derivative positions such as options, reported apart and never added to shares. | [optional]
**percent_basis** | Option<**String**> | How the percentage was reached, for example single_reporting_person, uniform_group_percent or single_band. | [optional]
**reported_percent** | Option<**f64**> | The percentage the filing reports, on the percentage_basis. | [optional]
**reported_percent_low** | Option<**f64**> | The low end of a reported band, for a source that reports a band. | [optional]
**reported_percent_high** | Option<**f64**> | The high end of a reported band, for a source that reports a band. | [optional]
**percentage_basis** | Option<**String**> | What the percentage is of: beneficial_ownership, voting_power or economic_interest. Percentages of different bases are not comparable. | [optional]
**edge_state** | Option<**String**> | standing, ceased or no_longer_reported, for the owner's position as a whole. | [optional]
**position_count** | Option<**i32**> | Positions the row consolidates. | [optional]
**standing_position_count** | Option<**i32**> | Of those, positions still standing. | [optional]
**ceased_position_count** | Option<**i32**> | Of those, positions reported ceased. | [optional]
**no_longer_reported_position_count** | Option<**i32**> | Of those, positions the filer no longer reports. | [optional]
**first_observed_date** | Option<[**String**](string.md)> | Date the position was first reported; null when the owner's name is withheld. | [optional]
**last_activity_date** | Option<[**String**](string.md)> | Date of the latest reported activity; null when the owner's name is withheld. | [optional]
**last_filing_date** | Option<[**String**](string.md)> | Date of the latest filing; null when the owner's name is withheld. | [optional]
**source_url** | Option<**String**> | A link to the latest filing; null when the owner's name is withheld, because the filing names the owner. | [optional]
**contains_personal_data** | Option<**bool**> | True when the row names a natural person. | [optional]
**privacy_status** | Option<**String**> | The filing's privacy status for the owner: not_applicable, pending_review, approved_public_fields, restricted or protected. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


