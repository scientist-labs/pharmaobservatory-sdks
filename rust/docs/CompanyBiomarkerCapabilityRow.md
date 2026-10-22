# CompanyBiomarkerCapabilityRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**biomarker_validation_id** | Option<**String**> | The evidence assertion id. | [optional]
**organization_name** | Option<**String**> | The organization's name as filed on the evidence. | [optional]
**organization_org_id** | Option<**String**> | Legal-entity level: the organization record of the organization. | [optional]
**organization_company_id** | Option<**String**> | Company level: the company the organization belongs to (null when unresolved). | [optional]
**organization_company_name** | Option<**String**> | Display name of the organization's company. | [optional]
**organization_group_id** | Option<**String**> | Group level: the organization's current ultimate parent company; a root company's group is itself. | [optional]
**organization_group_name** | Option<**String**> | Display name of the organization's group. | [optional]
**gene_symbol** | Option<**String**> | The gene's HGNC (Human Genome Organisation Gene Nomenclature Committee) symbol; null for evidence that names no gene. | [optional]
**biomarker_id** | Option<**String**> | The biomarker's id (a hash, not the route key). | [optional]
**biomarker_display_name** | Option<**String**> | The biomarker's display name. | [optional]
**biomarker_kind** | Option<**String**> | The kind of biomarker. | [optional]
**biomarker_identity_status** | Option<**String**> | How firmly the biomarker's identity is resolved. | [optional]
**organization_role** | Option<**String**> | The organization's role in the evidence (for example assay developer or testing laboratory). | [optional]
**organization_match_tier** | Option<**String**> | How the organization was matched to its record. | [optional]
**tier** | Option<**i32**> | The evidence tier, higher for stronger kinds of evidence. A label of the kind of evidence, not a quality score. | [optional]
**tier_label** | Option<**String**> | The evidence tier in words. | [optional]
**evidence_type** | Option<**String**> | The kind of evidence (companion diagnostic authorization, trial endpoint, trial eligibility, trial specimen collection, trial text mention or ClinVar clinical testing). | [optional]
**evidence_count** | Option<**i32**> | Assertions that share this biomarker, organization and role. | [optional]
**evidence_ref** | Option<**String**> | The reference of the evidence (a trial id, a device application number or a project id). | [optional]
**evidence_ref_kind** | Option<**String**> | What evidence_ref is. | [optional]
**trial_id** | Option<**String**> | The trial's registry id (NCT, National Clinical Trial, number) when the evidence is a trial. | [optional]
**citation_url** | Option<**String**> | A link to the evidence. | [optional]
**evidence_snippet** | Option<**String**> | A short excerpt of the evidence. | [optional]
**biomarker_role** | Option<**String**> | The role the biomarker plays in the evidence. | [optional]
**is_surrogate_endpoint_claim** | Option<**bool**> | Whether the evidence claims the biomarker as a surrogate endpoint. | [optional]
**clinvar_scv_count** | Option<**i32**> | ClinVar submissions behind a ClinVar row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


