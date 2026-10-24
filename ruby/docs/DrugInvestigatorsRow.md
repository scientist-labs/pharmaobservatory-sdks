# PharmaObservatory::DrugInvestigatorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The drug&#39;s resolved id (the same value as molecule_id on the drugs dataset). The entry key for the drugs object type. | [optional] |
| **drug_name** | **String** | The drug&#39;s name as the drugs dataset names it. | [optional] |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **investigator_person_name** | **String** | The investigator&#39;s name as the registry filed it, resolved across the person&#39;s merged name fragments. | [optional] |
| **investigator_affiliation** | **String** | The investigator&#39;s affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional] |
| **n_trials** | **Integer** | Distinct trials the person is named on that are in scope of the row (the target&#39;s trials, or the drug&#39;s). | [optional] |
| **is_principal_investigator** | **Boolean** | True when the person holds the principal investigator role on any registered trial (the person-level flag the investigators dataset and the people page publish), not only the trials of this row. | [optional] |
| **roles** | **Array&lt;String&gt;** | The distinct roles the person holds on the trials of this row, as the registry filed them (PRINCIPAL_INVESTIGATOR, STUDY_DIRECTOR, STUDY_CHAIR, SUB_INVESTIGATOR). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugInvestigatorsRow.new(
  molecule_id: null,
  drug_name: null,
  investigator_person_id: null,
  investigator_person_name: null,
  investigator_affiliation: null,
  n_trials: null,
  is_principal_investigator: null,
  roles: null
)
```

