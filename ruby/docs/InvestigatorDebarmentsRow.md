# PharmaObservatory::InvestigatorDebarmentsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **debarment_id** | **String** | The debarment record id. | [optional] |
| **debarred_name** | **String** | The debarred individual&#39;s name as the FDA list spells it. | [optional] |
| **effective_date** | **Date** | The debarment&#39;s effective date. | [optional] |
| **term** | **String** | The debarment term as the FDA states it. | [optional] |
| **fr_citation** | **String** | The Federal Register citation. | [optional] |
| **source_url** | **String** | A link to the source record. | [optional] |
| **match_tier** | **String** | How the debarred individual was matched to the investigator: low is a bare-name match that reaches exactly one person (the only tier published; an ambiguous or unresolved name has no person and no row). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestigatorDebarmentsRow.new(
  investigator_person_id: null,
  debarment_id: null,
  debarred_name: null,
  effective_date: null,
  term: null,
  fr_citation: null,
  source_url: null,
  match_tier: null
)
```

