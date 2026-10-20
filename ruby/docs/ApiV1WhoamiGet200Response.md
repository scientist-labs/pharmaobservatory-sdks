# PharmaObservatory::ApiV1WhoamiGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ok** | **Boolean** |  |  |
| **sub** | **String** |  |  |
| **roles** | **Array&lt;String&gt;** |  |  |
| **name** | **Object** |  | [optional] |
| **email** | **Object** |  | [optional] |
| **organization** | **Object** |  | [optional] |
| **plan** | **Object** |  | [optional] |
| **limits** | [**ApiV1WhoamiGet200ResponseLimits**](ApiV1WhoamiGet200ResponseLimits.md) |  | [optional] |
| **quota** | [**ApiV1WhoamiGet200ResponseQuota**](ApiV1WhoamiGet200ResponseQuota.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1WhoamiGet200Response.new(
  ok: null,
  sub: null,
  roles: null,
  name: null,
  email: null,
  organization: null,
  plan: null,
  limits: null,
  quota: null
)
```

