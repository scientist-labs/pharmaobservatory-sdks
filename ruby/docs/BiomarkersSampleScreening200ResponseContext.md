# PharmaObservatory::BiomarkersSampleScreening200ResponseContext

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **coverage_counts** | [**BiomarkersSampleScreening200ResponseContextCoverageCounts**](BiomarkersSampleScreening200ResponseContextCoverageCounts.md) |  |  |
| **requested_positive_samples** | **Integer** |  |  |
| **zero_variants** | **Array&lt;Object&gt;** |  |  |
| **multiple_variants** | **Array&lt;String&gt;** |  |  |
| **source_error_variants** | **Array&lt;Object&gt;** |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::BiomarkersSampleScreening200ResponseContext.new(
  coverage_counts: null,
  requested_positive_samples: null,
  zero_variants: null,
  multiple_variants: null,
  source_error_variants: null
)
```

