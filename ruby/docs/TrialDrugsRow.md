# PharmaObservatory::TrialDrugsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **drug_key** | **String** | The molecule key, or the drug name as registered prefixed unresolved: for a drug that reached no molecule. | [optional] |
| **molecule_id** | **String** | The molecule key the drugs resolved to (null for a drug with no molecule). | [optional] |
| **molecule_name** | **String** | The molecule&#39;s display name (null for a drug with no molecule). | [optional] |
| **registered_as** | **Array&lt;String&gt;** | The drug names as registered on the trial that led to this row. | [optional] |
| **matched_from_name_part** | **Boolean** | Whether every match to the molecule came from one part of a multi-part drug name. | [optional] |
| **brand_count** | **Integer** | Number of brand names recorded for the molecule. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialDrugsRow.new(
  trial_id: null,
  drug_key: null,
  molecule_id: null,
  molecule_name: null,
  registered_as: null,
  matched_from_name_part: null,
  brand_count: null
)
```

