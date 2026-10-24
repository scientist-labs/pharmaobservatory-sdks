# PharmaObservatory::InvestigatorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **investigator_person_name** | **String** | The investigator&#39;s name as the registry filed it, resolved across the person&#39;s merged name fragments. | [optional] |
| **investigator_affiliation** | **String** | The investigator&#39;s affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional] |
| **roles** | **Array&lt;String&gt;** | Distinct registry roles the person holds across their trials (PRINCIPAL_INVESTIGATOR, STUDY_DIRECTOR, STUDY_CHAIR, SUB_INVESTIGATOR); empty when the registry names none. | [optional] |
| **is_principal_investigator** | **Boolean** | True when the person is named principal investigator on at least one trial. | [optional] |
| **n_trials** | **Integer** | Distinct registered trials across the person&#39;s merged fragments. | [optional] |
| **n_trial_publications** | **Integer** | Distinct publications the person&#39;s trials register (trial-attributed, not authorship). | [optional] |
| **n_authored_publications** | **Integer** | Publications linked to the person as an author. | [optional] |
| **merged_fragment_count** | **Integer** | Name fragments the resolution merged into this person. | [optional] |
| **npi_anchored** | **Boolean** | True when the person&#39;s fragments were merged on a US National Provider Identifier. | [optional] |
| **orcid_anchored** | **Boolean** | True when the person&#39;s fragments were merged on an ORCID. | [optional] |
| **identity_anchor** | **String** | npi when the person is NPI-anchored, orcid when ORCID-anchored only, none otherwise. | [optional] |
| **primary_specialty** | **String** | The clinician&#39;s primary specialty in the CMS Doctors and Clinicians file, upper-case as filed (null without an NPI or clinician record). | [optional] |
| **has_nih_funding** | **Boolean** | True when the person holds at least one corroborated NIH RePORTER project. | [optional] |
| **nih_award_dollars** | **Float** | Corroborated NIH RePORTER award dollars (zero when only uncorroborated links exist). | [optional] |
| **nih_project_count** | **Integer** | Corroborated NIH RePORTER projects. | [optional] |
| **nih_award_dollars_incl_uncorroborated** | **Float** | NIH RePORTER award dollars including projects linked on organization affinity alone. | [optional] |
| **nih_project_count_incl_uncorroborated** | **Integer** | NIH RePORTER projects including those linked on organization affinity alone. | [optional] |
| **nih_corroborated** | **Boolean** | True when at least one contributing RePORTER link is confirmed by a trial or a publication; null when the person has no funding record. | [optional] |
| **nih_award_count** | **Integer** | Distinct NIH awards (null when the person has no funding record). | [optional] |
| **nih_first_fiscal_year** | **Integer** | First fiscal year of the person&#39;s NIH projects. | [optional] |
| **nih_last_fiscal_year** | **Integer** | Last fiscal year of the person&#39;s NIH projects. | [optional] |
| **nih_top_ic** | **String** | The NIH institute or center that administers the most of the person&#39;s projects. | [optional] |
| **has_hospital_affiliation** | **Boolean** | True when the person has a CMS facility affiliation (NPI-keyed, so false for a person with no NPI by lookup, not by fact). | [optional] |
| **n_affiliated_hospitals** | **Integer** | CMS hospitals the person is affiliated with. | [optional] |
| **primary_ccn** | **String** | The lowest affiliated CMS Certification Number, the key into the hospitals dataset. | [optional] |
| **has_expertise_topics** | **Boolean** | True when the person has authored papers with MeSH topics (investigator_topics has rows for them). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestigatorsRow.new(
  investigator_person_id: null,
  investigator_person_name: null,
  investigator_affiliation: null,
  roles: null,
  is_principal_investigator: null,
  n_trials: null,
  n_trial_publications: null,
  n_authored_publications: null,
  merged_fragment_count: null,
  npi_anchored: null,
  orcid_anchored: null,
  identity_anchor: null,
  primary_specialty: null,
  has_nih_funding: null,
  nih_award_dollars: null,
  nih_project_count: null,
  nih_award_dollars_incl_uncorroborated: null,
  nih_project_count_incl_uncorroborated: null,
  nih_corroborated: null,
  nih_award_count: null,
  nih_first_fiscal_year: null,
  nih_last_fiscal_year: null,
  nih_top_ic: null,
  has_hospital_affiliation: null,
  n_affiliated_hospitals: null,
  primary_ccn: null,
  has_expertise_topics: null
)
```

