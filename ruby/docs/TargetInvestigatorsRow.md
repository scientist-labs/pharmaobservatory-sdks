# PharmaObservatory::TargetInvestigatorsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **target_id** | **String** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. | [optional] |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **investigator_person_name** | **String** | The investigator&#39;s name as the registry filed it, resolved across the person&#39;s merged name fragments. | [optional] |
| **investigator_affiliation** | **String** | The investigator&#39;s affiliation as filed (the most-trialled affiliation across the merged fragments). | [optional] |
| **n_trials** | **Integer** | Distinct trials the person is named on that are in scope of the row (the target&#39;s trials, or the drug&#39;s). | [optional] |
| **is_principal_investigator** | **Boolean** | True when the person holds the principal investigator role on any registered trial (the person-level flag the investigators dataset and the people page publish), not only the trials of this row. | [optional] |
| **roles** | **Array&lt;String&gt;** | The distinct roles the person holds on the trials of this row, as the registry filed them (PRINCIPAL_INVESTIGATOR, STUDY_DIRECTOR, STUDY_CHAIR, SUB_INVESTIGATOR). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TargetInvestigatorsRow.new(
  target_id: null,
  investigator_person_id: null,
  investigator_person_name: null,
  investigator_affiliation: null,
  n_trials: null,
  is_principal_investigator: null,
  roles: null
)
```

