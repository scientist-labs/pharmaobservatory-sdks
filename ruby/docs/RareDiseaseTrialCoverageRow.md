# PharmaObservatory::RareDiseaseTrialCoverageRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **coverage_state** | **String** | no_bridge, bridged_no_trials or trials. | [optional] |
| **mapping_arms** | **Array&lt;String&gt;** | The mapping arms the disease has a condition under (exact, unspecified). | [optional] |
| **mapping_target_curie** | **String** | The MeSH condition or conditions the disease maps to (prefixed identifiers, comma separated), the exact ones where any exist. | [optional] |
| **mapping_target_name** | **String** | The names of those MeSH conditions. | [optional] |
| **mapping_sources** | **Array&lt;String&gt;** | The files that assert the mapping. | [optional] |
| **mapping_source_releases** | **Array&lt;String&gt;** | The releases of those files. | [optional] |
| **n_trials_exact_direct** | **Integer** | Trials linked through an exact mapping whose own condition list names the disease. Count within this column only. | [optional] |
| **n_trials_exact_broader** | **Integer** | Trials linked through an exact mapping to a broader condition than the trial studies. Count within this column only. | [optional] |
| **n_trials_unspecified** | **Integer** | Trials linked through a MONDO cross-reference that states no relation. Count within this column only. | [optional] |
| **n_open_direct** | **Integer** | Exact direct trials that are recruiting or not yet recruiting. | [optional] |
| **n_ongoing_direct** | **Integer** | Exact direct trials that are active but not recruiting, or enrolling by invitation. | [optional] |
| **n_unverified_2y_direct** | **Integer** | Exact direct trials whose record was last updated two or more years before the snapshot. | [optional] |
| **direct_share** | **Float** | The exact direct trials as a share of the exact direct and exact broader trials; null when both are zero. | [optional] |
| **n_sponsors** | **Integer** | Sponsoring organizations of the exact direct trials. | [optional] |
| **snapshot_date** | **Date** | The date of the registry snapshot. | [optional] |
| **freshness_as_of** | **Date** | The registry&#39;s latest update date the snapshot reflects. | [optional] |
| **license_name** | **String** | The licence of the Orphadata Science release the mapping rests on. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |
| **trials_license_name** | **String** | The licence of the trial registry data. | [optional] |
| **trials_attribution** | **String** | The attribution line the trial registry requires. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseTrialCoverageRow.new(
  orpha_code: null,
  coverage_state: null,
  mapping_arms: null,
  mapping_target_curie: null,
  mapping_target_name: null,
  mapping_sources: null,
  mapping_source_releases: null,
  n_trials_exact_direct: null,
  n_trials_exact_broader: null,
  n_trials_unspecified: null,
  n_open_direct: null,
  n_ongoing_direct: null,
  n_unverified_2y_direct: null,
  direct_share: null,
  n_sponsors: null,
  snapshot_date: null,
  freshness_as_of: null,
  license_name: null,
  attribution: null,
  trials_license_name: null,
  trials_attribution: null
)
```

