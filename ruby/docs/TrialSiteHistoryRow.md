# PharmaObservatory::TrialSiteHistoryRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **trial_id** | **String** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] |
| **as_of** | **Date** | The registry snapshot date this row describes. | [optional] |
| **prev_as_of** | **Date** | The trial&#39;s previous snapshot date (null for its first snapshot). | [optional] |
| **facilities_captured** | **Boolean** | True when the snapshot carried the site list; false means the site counts are unknown, not zero. | [optional] |
| **status_captured** | **Boolean** | True when the snapshot carried a recruitment status on the site rows. | [optional] |
| **comparable** | **Boolean** | True when the previous snapshot also carried the site list, so the add and drop counts are defined. | [optional] |
| **sites_listed** | **Integer** | Distinct sites the trial listed at the snapshot (null where facilities_captured is false). | [optional] |
| **sites_added** | **Integer** | Sites present at this snapshot and absent at the previous one (null where comparable is false). | [optional] |
| **sites_dropped** | **Integer** | Sites present at the previous snapshot and absent at this one (null where comparable is false). | [optional] |
| **sites_status_populated** | **Integer** | Sites whose row carried a recruitment status at the snapshot. | [optional] |
| **sites_recruiting** | **Integer** | Sites with status RECRUITING at the snapshot. | [optional] |
| **sites_not_yet_recruiting** | **Integer** | Sites with status NOT_YET_RECRUITING at the snapshot. | [optional] |
| **sites_active_not_recruiting** | **Integer** | Sites with status ACTIVE_NOT_RECRUITING at the snapshot. | [optional] |
| **sites_completed** | **Integer** | Sites with status COMPLETED at the snapshot. | [optional] |
| **sites_other_status** | **Integer** | Sites with a status other than the four counted above. | [optional] |
| **countries_listed** | **Integer** | Countries the listed sites span at the snapshot. | [optional] |
| **captured_columns** | **Array&lt;String&gt;** | The registry study fields the snapshot&#39;s archive carried (null when no capture record exists for the snapshot; an empty list means it carried none of them). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::TrialSiteHistoryRow.new(
  trial_id: null,
  as_of: null,
  prev_as_of: null,
  facilities_captured: null,
  status_captured: null,
  comparable: null,
  sites_listed: null,
  sites_added: null,
  sites_dropped: null,
  sites_status_populated: null,
  sites_recruiting: null,
  sites_not_yet_recruiting: null,
  sites_active_not_recruiting: null,
  sites_completed: null,
  sites_other_status: null,
  countries_listed: null,
  captured_columns: null
)
```

