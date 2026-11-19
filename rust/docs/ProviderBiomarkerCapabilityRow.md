# ProviderBiomarkerCapabilityRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**provider_biomarker_capability_id** | Option<**String**> | The evidence assertion id: a hash of the supplier, biomarker, lane, role, evidence reference kind, evidence reference and ladder rule. | [optional]
**provider_id** | Option<**String**> | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/<id>). A supplier is not a company. | [optional]
**provider_uuid** | Option<**String**> | The Scientist.com public provider uuid, the join key to the RX marketplace. Present on every row, including providers the marketplace does not publish. | [optional]
**provider_slug** | Option<**String**> | The supplier's marketplace slug. | [optional]
**provider_name** | Option<**String**> | The supplier's name. | [optional]
**route_key** | Option<**String**> | The biomarker's route key, the entry key for the biomarkers object type; the gene symbol for a gene biomarker. | [optional]
**biomarker_id** | Option<**String**> | The canonical biomarker entity id. | [optional]
**gene_symbol** | Option<**String**> | The gene symbol; null for a source-scoped biomarker that names no gene. | [optional]
**biomarker_display_name** | Option<**String**> | The biomarker's display name. | [optional]
**evidence_lane** | Option<**String**> | The source path of the row: ladder_trial and ladder_clinvar (the evidence ladder), fda_cdx (FDA companion diagnostic list), fda_bqp (FDA Biomarker Qualification Program), literature_assay (assay providers named in PubMed full text), supplier_survey (the supplier's own declaration) clinvar_submission (ClinVar clinical-testing submissions by a laboratory the supplier carries by name, one row per ClinVar record, tier null) and supplier_service_listing (a Scientist.com marketplace catalogue service the supplier listed, whose name names the gene; tier null, no evidence sense). Nothing is merged across lanes. | [optional]
**evidence_type** | Option<**String**> | The source's own evidence vocabulary: the ladder evidence_type on a ladder lane, literature_assay_provider on the literature lane, supplier_survey_declaration on the survey lane, rx_service_listing on the service listing lane. | [optional]
**org_role** | Option<**String**> | The supplier's role on the evidence: the ladder role on a ladder lane, the ledger provider role on the literature lane, respondent on the survey lane, or marketplace_listing on the service listing lane. | [optional]
**analytical_validation_evidence** | Option<**String**> | Sense 1, literature_assay lane only: the source's own label for how the assay is evidenced, never a number. literature_reported is a paper naming the assay, never validation. | [optional]
**clinical_use_evidence** | Option<**String**> | Sense 2: the ladder rung's name on a ladder lane, literature_reported on the literature lane (a literature mention, never validation), and ClinVar's own review status verbatim on the clinvar_submission lane. | [optional]
**regulatory_qualification** | Option<**String**> | Sense 3, fda_bqp lane only: fda_qualified, the FDA Biomarker Qualification Program qualified the biomarker. | [optional]
**companion_diagnostic** | Option<**String**> | Sense 4, fda_cdx lane only: fda_authorized_companion_diagnostic, the FDA companion diagnostic list names the supplier's device for the biomarker. | [optional]
**supplier_declared_status** | Option<**String**> | The supplier's own declaration, never mapped into the four evidence kinds: the survey's assay status verbatim on the supplier_survey lane, supplier_declared_service_listed on the supplier_service_listing lane (the supplier listed a catalogue service naming the gene). | [optional]
**tier** | Option<**i32**> | The ladder rung of this row, 1 to 4, a kind of evidence and never a score; set on the ladder_trial, ladder_clinvar, fda_cdx and fda_bqp lanes, null on literature_assay, supplier_survey, clinvar_submission and supplier_service_listing. Not rolled up: see the example query for the maximum within the ladder lanes. | [optional]
**tier_rule_id** | Option<**String**> | The ladder rule that set the tier of this row; null where tier is null. | [optional]
**evidence_ref** | Option<**String**> | The row key that resolves in the model evidence_ref_kind names. | [optional]
**evidence_ref_kind** | Option<**String**> | What evidence_ref is. | [optional]
**trial_id** | Option<**String**> | The trial's ClinicalTrials.gov registry id (NCT, National Clinical Trial, number) on trial-sourced ladder rows; null otherwise. | [optional]
**clinvar_variation_id** | Option<**String**> | The ClinVar variation id on ClinVar ladder rows; null otherwise. | [optional]
**device_application_numbers** | Option<**Vec<String>**> | The FDA device application numbers of a companion diagnostic row; null otherwise. | [optional]
**device_credential_labels** | Option<**Vec<String>**> | The FDA pathway labels (PMA, De Novo, 510(k)) of a companion diagnostic row; null otherwise. | [optional]
**source_url** | Option<**String**> | The source URL the evidence row carries, where it has one. | [optional]
**citation_url** | Option<**String**> | A resolvable citation: the trial page, the ClinVar variation page, the FDA device page or list, the PubMed page, or the supplier's marketplace page on a service listing. Null where the source offers none. | [optional]
**evidence_snippet** | Option<**String**> | A short evidence text (at most 500 characters), the paper title, or the catalogue service name on a service listing row, where the source has one. | [optional]
**most_recent_evidence_date** | Option<[**String**](string.md)> | literature_assay lane only: the publication date of the most recent cited paper. | [optional]
**match_method** | Option<**String**> | How the evidence reached the supplier: org_member_exact (the organization record of a member of the supplier), rx_name_key_exact (an exact name-key match of a source name to the supplier) or rx_provider_direct (the supplier's own marketplace listing, no name matched). A name match is not a verified identity. | [optional]
**match_confidence** | Option<**f64**> | The confidence of the upstream match rung where that rung defines one; null otherwise, never invented. | [optional]
**company_name_matched** | Option<**String**> | The source name that matched the supplier's name key, on name-matched rows. | [optional]
**source_license** | Option<**String**> | The license statement of the row's source lane. No lane carries a usage restriction. | [optional]
**attribution** | Option<**String**> | The credit the row's source lane is owed. | [optional]
**as_of** | Option<[**String**](string.md)> | The build date. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


