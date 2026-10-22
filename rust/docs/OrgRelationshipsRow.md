# OrgRelationshipsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**subject_org_id** | Option<**String**> | Legal-entity level: the subject's organization record. | [optional]
**subject_name** | Option<**String**> | The organization's name as filed. | [optional]
**subject_company_id** | Option<**String**> | Company level: the subject's company (null when unresolved). | [optional]
**subject_company_name** | Option<**String**> | Display name of the subject's company. | [optional]
**subject_group_id** | Option<**String**> | Group level: the subject's current ultimate parent; a root's group is itself. | [optional]
**subject_group_name** | Option<**String**> | Display name of the subject's group. | [optional]
**object_org_id** | Option<**String**> | Legal-entity level: the object's organization record. | [optional]
**object_name** | Option<**String**> | The organization's name as filed. | [optional]
**object_company_id** | Option<**String**> | Company level: the object's company (null when unresolved). | [optional]
**object_company_name** | Option<**String**> | Display name of the object's company. | [optional]
**object_group_id** | Option<**String**> | Group level: the object's current ultimate parent; a root's group is itself. | [optional]
**object_group_name** | Option<**String**> | Display name of the object's group. | [optional]
**predicate** | Option<**String**> | The directed relationship from subject to object: parent_of and ultimate_parent_of (corporate ownership, the subject is the parent), linked_charity_of (a charity linked to the main charity), uses_vendor (a sponsor and the contract research organization inferred for its trials), co_sponsors_trial_with (two sponsors of one trial), project_partner_of and awards_to (UK research grants). | [optional]
**edge_family** | Option<**String**> | The family of the predicate: corporate_structure (parent_of, ultimate_parent_of, linked_charity_of), contract_vendor (uses_vendor), co_sponsorship (co_sponsors_trial_with) or research_funding (project_partner_of, awards_to). | [optional]
**edge_scope** | Option<**String**> | general for an edge from a source that covers organizations at large (GLEIF, Charity Commission, the contract research organization (CRO) inference); rare_disease_scoped for an edge from a source read only for organizations in the rare-disease evidence set (ROR, co-sponsorship, UK research grants). A rare_disease_scoped edge is a true edge and an incomplete picture: an organization with none of them has not been shown to have none. | [optional]
**source_key** | Option<**String**> | The source that asserted the edge: gleif (Global Legal Entity Identifier Foundation Golden Copy), charity_commission, aact_cro_inference (inferred from ClinicalTrials.gov filings against a curated CRO list), ror (Research Organization Registry), aact (ClinicalTrials.gov) or ukri_gtr (UK Research and Innovation Gateway to Research). | [optional]
**evidence_type** | Option<**String**> | The kind of evidence behind the edge: registry (a register entry), trial, project or grant. | [optional]
**authority_class** | Option<**String**> | authoritative when the source itself filed the relationship; derived when the edge is an inference (the CRO edge is a curated-list name match, not a filing). | [optional]
**evidence_count** | Option<**i64**> | The source records behind the edge. Counted within one source, so it does not add across sources. | [optional]
**context_count** | Option<**i64**> | Distinct source record contexts among them (for example distinct trials). | [optional]
**confidence** | Option<**f64**> | The highest confidence among the edge's records, 0 to 1, where the source states one. Null for sources that state none. A GLEIF ownership-resolution score and a CRO name-match score are different measures; do not compare them. | [optional]
**first_observed_at** | Option<**String**> | When the earliest of the edge's records was observed in the source. | [optional]
**last_observed_at** | Option<**String**> | When the latest of the edge's records was observed in the source. | [optional]
**license_name** | Option<**String**> | The licence of the source behind the edge. | [optional]
**attribution** | Option<**String**> | The credit line the source requires for the edge. Show it wherever the edge is shown: the Charity Commission and UK Research and Innovation edges require it under the Open Government Licence v3.0. | [optional]
**restriction_tier** | Option<**String**> | The tier vocabulary is none or licensable. Every row here is none: every source behind these edges is open, and a test on the source assertions reports any restricted source that is ever added. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


