# InvestigatorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investigator_person_id** | Option<**String**> | The investigator's resolved person key, the id in the web page address (/people/<id>). The entry key of the investigators object type. | [optional]
**investigator_person_name** | Option<**String**> | The investigator's name as the registry filed it, resolved across the person's merged name fragments. | [optional]
**investigator_affiliation** | Option<**String**> | The investigator's affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional]
**roles** | Option<**Vec<String>**> | Distinct registry roles the person holds across their trials (PRINCIPAL_INVESTIGATOR, STUDY_DIRECTOR, STUDY_CHAIR, SUB_INVESTIGATOR); empty when the registry names none. | [optional]
**is_principal_investigator** | Option<**bool**> | True when the person is named principal investigator on at least one trial. | [optional]
**n_trials** | Option<**i64**> | Distinct registered trials across the person's merged fragments. | [optional]
**n_trial_publications** | Option<**i64**> | Distinct publications the person's trials register (trial-attributed, not authorship). | [optional]
**n_authored_publications** | Option<**i64**> | Publications linked to the person as an author. | [optional]
**merged_fragment_count** | Option<**i64**> | Name fragments the resolution merged into this person. | [optional]
**npi_anchored** | Option<**bool**> | True when the person's fragments were merged on a US National Provider Identifier. | [optional]
**orcid_anchored** | Option<**bool**> | True when the person's fragments were merged on an ORCID. | [optional]
**identity_anchor** | Option<**String**> | npi when the person is NPI-anchored, orcid when ORCID-anchored only, none otherwise. | [optional]
**primary_specialty** | Option<**String**> | The clinician's primary specialty in the CMS Doctors and Clinicians file, upper-case as filed (null without an NPI or clinician record). | [optional]
**has_nih_funding** | Option<**bool**> | True when the person holds at least one corroborated NIH RePORTER project. | [optional]
**nih_award_dollars** | Option<**f64**> | Corroborated NIH RePORTER award dollars (zero when only uncorroborated links exist). | [optional]
**nih_project_count** | Option<**i64**> | Corroborated NIH RePORTER projects. | [optional]
**nih_award_dollars_incl_uncorroborated** | Option<**f64**> | NIH RePORTER award dollars including projects linked on organization affinity alone. | [optional]
**nih_project_count_incl_uncorroborated** | Option<**i64**> | NIH RePORTER projects including those linked on organization affinity alone. | [optional]
**nih_corroborated** | Option<**bool**> | True when at least one contributing RePORTER link is confirmed by a trial or a publication; null when the person has no funding record. | [optional]
**nih_award_count** | Option<**i64**> | Distinct NIH awards (null when the person has no funding record). | [optional]
**nih_first_fiscal_year** | Option<**i32**> | First fiscal year of the person's NIH projects. | [optional]
**nih_last_fiscal_year** | Option<**i32**> | Last fiscal year of the person's NIH projects. | [optional]
**nih_top_ic** | Option<**String**> | The NIH institute or center that administers the most of the person's projects. | [optional]
**has_hospital_affiliation** | Option<**bool**> | True when the person has a CMS facility affiliation (NPI-keyed, so false for a person with no NPI by lookup, not by fact). | [optional]
**n_affiliated_hospitals** | Option<**i64**> | CMS hospitals the person is affiliated with. | [optional]
**primary_ccn** | Option<**String**> | The lowest affiliated CMS Certification Number, the key into the hospitals dataset. | [optional]
**has_expertise_topics** | Option<**bool**> | True when the person has authored papers with MeSH topics (investigator_topics has rows for them). | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


