# PharmaObservatory::ApiV1SiteRecordGet200Response

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mode** | **String** |  | [optional] |
| **requested** | **String** |  | [optional] |
| **matched_count** | **Integer** |  | [optional] |
| **returned_count** | **Integer** |  | [optional] |
| **sites** | [**Array&lt;ApiV1SiteRecordGet200ResponseSitesInner&gt;**](ApiV1SiteRecordGet200ResponseSitesInner.md) |  | [optional] |
| **build_state** | [**ApiV1BiomarkerAssayScreenGet200ResponseBuildState**](ApiV1BiomarkerAssayScreenGet200ResponseBuildState.md) |  |  |
| **error** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |
| **detail_sections** | [**ApiV1SiteRecordGet200ResponseDetailSections**](ApiV1SiteRecordGet200ResponseDetailSections.md) |  | [optional] |
| **site_key_redirect** | [**ApiV1SiteRecordGet200ResponseSiteKeyRedirect**](ApiV1SiteRecordGet200ResponseSiteKeyRedirect.md) |  | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1SiteRecordGet200Response.new(
  mode: null,
  requested: null,
  matched_count: null,
  returned_count: null,
  sites: null,
  build_state: null,
  error: null,
  message: null,
  detail_sections: null,
  site_key_redirect: null
)
```

