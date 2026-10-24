# PharmaObservatory::InvestigatorInspectionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **bimo_record_id** | **String** | The record id of the inspection and deficiency code. | [optional] |
| **bimo_investigator_id** | **String** | The CLIIL investigator code; two codes under one name may be different people. | [optional] |
| **inspected_investigator_name** | **String** | The inspected investigator&#39;s name as CLIIL files it (&#39;last, first&#39;). | [optional] |
| **inspection_date** | **Date** | The inspection&#39;s date. | [optional] |
| **inspection_type** | **String** | The inspection type code as CLIIL files it (DA or FC). | [optional] |
| **inspection_classification** | **String** | The FDA inspection classification: NAI (no action indicated), VAI (voluntary action indicated) or OAI (official action indicated). | [optional] |
| **deficiency_code** | **String** | The deficiency code as CLIIL codes it (null when the record carries none). | [optional] |
| **has_deficiency** | **Boolean** | True when the inspection cited a deficiency. | [optional] |
| **site_name** | **String** | The inspected site&#39;s name. | [optional] |
| **site_city** | **String** | The inspected site&#39;s city. | [optional] |
| **site_state** | **String** | The inspected site&#39;s state. | [optional] |
| **site_country** | **String** | The inspected site&#39;s country. | [optional] |
| **match_tier** | **String** | How the CLIIL investigator was matched to the person: low is a bare-name match that reaches exactly one person (the only tier published; an ambiguous or unresolved name has no person and no row). | [optional] |
| **match_confidence** | **Float** | The match&#39;s confidence (0.45 for the low tier). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestigatorInspectionsRow.new(
  investigator_person_id: null,
  bimo_record_id: null,
  bimo_investigator_id: null,
  inspected_investigator_name: null,
  inspection_date: null,
  inspection_type: null,
  inspection_classification: null,
  deficiency_code: null,
  has_deficiency: null,
  site_name: null,
  site_city: null,
  site_state: null,
  site_country: null,
  match_tier: null,
  match_confidence: null
)
```

