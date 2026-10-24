# PharmaObservatory::InvestigatorHospitalsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **investigator_person_id** | **String** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] |
| **ccn** | **String** | CMS Certification Number, the six-character Medicare provider id the hospital page address carries. | [optional] |
| **hospital_name** | **String** | The hospital&#39;s facility name in the CMS directory. | [optional] |
| **city** | **String** | The hospital&#39;s city in the CMS directory. | [optional] |
| **state** | **String** | The hospital&#39;s state in the CMS directory. | [optional] |
| **facility_types** | **Array&lt;String&gt;** | The CMS facility types of the affiliation. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::InvestigatorHospitalsRow.new(
  investigator_person_id: null,
  ccn: null,
  hospital_name: null,
  city: null,
  state: null,
  facility_types: null
)
```

