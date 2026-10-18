# PharmaObservatory::OrgsList200ResponseRowsInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **canonical_org_id** | **String** |  |  |
| **name** | **String** |  |  |
| **kind** | **String** |  |  |
| **org_type** | **String** |  |  |
| **org_kinds** | **Array&lt;String&gt;** |  |  |
| **country** | **String** |  |  |
| **ticker** | **String** |  |  |
| **lead_trials** | **Integer** |  |  |
| **recent_lead_trials** | **Integer** |  |  |
| **lead_enrollment** | **Integer** |  |  |
| **collaborator_trials** | **Object** |  |  |
| **group_size** | **Integer** |  |  |
| **web_url** | **String** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsList200ResponseRowsInner.new(
  id: null,
  canonical_org_id: null,
  name: null,
  kind: null,
  org_type: null,
  org_kinds: null,
  country: null,
  ticker: null,
  lead_trials: null,
  recent_lead_trials: null,
  lead_enrollment: null,
  collaborator_trials: null,
  group_size: null,
  web_url: null
)
```

