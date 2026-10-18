# PharmaObservatory::OrgsMemberIdentity200ResponseContext

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** |  |  |
| **name** | **String** |  |  |
| **countries** | **Array&lt;String&gt;** |  |  |
| **legal_forms** | **Array&lt;String&gt;** |  |  |
| **identifiers** | [**Array&lt;OrgsMemberIdentity200ResponseContextIdentifiersInner&gt;**](OrgsMemberIdentity200ResponseContextIdentifiersInner.md) |  |  |
| **parent** | [**OrgsMemberIdentity200ResponseContextParent**](OrgsMemberIdentity200ResponseContextParent.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::OrgsMemberIdentity200ResponseContext.new(
  company_id: null,
  name: null,
  countries: null,
  legal_forms: null,
  identifiers: null,
  parent: null
)
```

