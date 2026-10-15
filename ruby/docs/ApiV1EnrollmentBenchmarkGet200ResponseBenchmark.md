# PharmaObservatory::ApiV1EnrollmentBenchmarkGet200ResponseBenchmark

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cohort_level** | **String** |  |  |
| **cohort_label** | **String** |  |  |
| **condition_name** | **String** |  |  |
| **phase** | **String** |  |  |
| **country** | **String** |  |  |
| **country_backed** | **Boolean** |  |  |
| **accrual_rate_per_site_month** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkAccrualRatePerSiteMonth**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkAccrualRatePerSiteMonth.md) |  |  |
| **typical_duration_months** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkTypicalDurationMonths**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkTypicalDurationMonths.md) |  |  |
| **enrollment** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment.md) |  |  |
| **active_sites** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkEnrollment.md) |  |  |
| **cohort_size** | **Integer** |  |  |
| **cohort_size_total_indication** | **Integer** |  |  |
| **confidence** | [**ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkConfidence**](ApiV1EnrollmentBenchmarkGet200ResponseBenchmarkConfidence.md) |  |  |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ApiV1EnrollmentBenchmarkGet200ResponseBenchmark.new(
  cohort_level: null,
  cohort_label: null,
  condition_name: null,
  phase: null,
  country: null,
  country_backed: null,
  accrual_rate_per_site_month: null,
  typical_duration_months: null,
  enrollment: null,
  active_sites: null,
  cohort_size: null,
  cohort_size_total_indication: null,
  confidence: null
)
```

