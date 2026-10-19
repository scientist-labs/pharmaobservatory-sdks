# BiomarkerValidationRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**biomarker_validation_id** | Option<**String**> | The evidence assertion id. | [optional]
**route_key** | Option<**String**> | The biomarker's route key, the entry key for the biomarkers object type; for a gene it is the gene symbol. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol (HGNC, the HUGO Gene Nomenclature Committee). | [optional]
**biomarker_id** | Option<**String**> | The biomarker key. | [optional]
**biomarker_display_name** | Option<**String**> | The biomarker's display name. | [optional]
**biomarker_kind** | Option<**String**> | The kind of biomarker the evidence is filed against. | [optional]
**organization_name** | Option<**String**> | The organization that can supply the evidence, as the source names it (null when the source names none that resolved). | [optional]
**organization_org_id** | Option<**String**> | Legal-entity level: the organization record (null when unresolved). | [optional]
**organization_company_id** | Option<**String**> | Company level: the organization's company (null when unresolved). | [optional]
**organization_company_name** | Option<**String**> | Display name of the organization's company. | [optional]
**organization_group_id** | Option<**String**> | Group level: the organization's current ultimate parent; a root's group is itself. | [optional]
**organization_group_name** | Option<**String**> | Display name of the organization's group. | [optional]
**organization_role** | Option<**String**> | What the organization is to the evidence: lead_sponsor, collaborator, contract_vendor, assay_manufacturer, qualification_submitter, testing_lab or publishing_org. | [optional]
**organization_match_tier** | Option<**String**> | How the organization name matched an organization record. | [optional]
**tier** | Option<**i32**> | The rung on the validation ladder, 1 to 4. A rung is a kind of evidence, never a score: rows are not ranked across organizations or evidence types. | [optional]
**tier_label** | Option<**String**> | The rung's name, for example fda_authorized_companion_diagnostic or named_in_inclusion_criteria. | [optional]
**evidence_type** | Option<**String**> | The evidence lane: fda_bqp_qualification, fda_cdx_authorization, aact_endpoint_role, aact_eligibility_section, aact_biospecimen, trial_mention or clinvar_clinical_testing. | [optional]
**evidence_count** | Option<**i32**> | Sibling rows sharing this biomarker, organization and role. Not the number of source records behind a ClinVar row (clinvar_scv_count). | [optional]
**evidence_ref** | Option<**String**> | The source reference the row cites. | [optional]
**evidence_ref_kind** | Option<**String**> | What kind of reference evidence_ref is. | [optional]
**trial_id** | Option<**String**> | The trial key (an NCT id) when the evidence is a trial; null otherwise. | [optional]
**citation_url** | Option<**String**> | A link to the cited record. | [optional]
**citation_verified** | Option<**bool**> | False for the FDA Biomarker Qualification Program arm, whose captured links are known to go stale; true otherwise. A link marked false may not resolve. | [optional]
**evidence_snippet** | Option<**String**> | The text span the evidence rests on, where the source gives one. | [optional]
**biomarker_role** | Option<**String**> | The biomarker's role in the evidence, where the source states it. | [optional]
**is_surrogate_endpoint_claim** | Option<**bool**> | True when the evidence is a claim that the biomarker is a surrogate endpoint. | [optional]
**clinvar_scv_count** | Option<**i32**> | For a ClinVar row, the number of ClinVar submission records the row collapses; null otherwise. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


