# PharmaObservatory::DrugVariantsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **member_id** | **String** | The id of the folded drug record: the drug_id other datasets carry. | [optional] |
| **member_name** | **String** | The folded record&#39;s display name. | [optional] |
| **role** | **String** | single for a one-to-one fold; component_of_compound for a combination arm listing the molecule. | [optional] |
| **method** | **String** | How the fold was made (for example rxnorm_synonym, unii_synonym, brand, override, registered_substance). | [optional] |
| **tier** | **Integer** | The resolver tier that made the fold (1 to 4); 4 is the model backstop. | [optional] |
| **confidence** | **Float** | The resolver&#39;s confidence for the method, from 0 to 1. | [optional] |
| **member_trials** | **Integer** | Registered trials of the member record. | [optional] |
| **member_max_phase** | **Integer** | The member&#39;s highest development phase. | [optional] |
| **member_approved** | **Boolean** | True when the member is approved somewhere. | [optional] |
| **member_biologic** | **Boolean** | True when the member is a biologic substance. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugVariantsRow.new(
  molecule_id: null,
  member_id: null,
  member_name: null,
  role: null,
  method: null,
  tier: null,
  confidence: null,
  member_trials: null,
  member_max_phase: null,
  member_approved: null,
  member_biologic: null
)
```

