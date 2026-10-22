# PharmaObservatory::RareDiseasesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **disease_key** | **String** | The disease key (orpha: plus the ORPHA code). An attribute, not an id: the entry key is orpha_code. | [optional] |
| **preferred_name** | **String** | The disease&#39;s preferred English name. | [optional] |
| **aliases** | **Array&lt;String&gt;** | English synonyms Orphanet lists for the disease. | [optional] |
| **concept_type** | **String** | Whether the concept is a group of disorders, a disorder or a subtype. | [optional] |
| **status** | **String** | The concept&#39;s lifecycle state in Orphanet (active or deprecated). | [optional] |
| **ordo_entity_type** | **String** | The entity type the Orphanet Rare Disease Ontology gives the concept. | [optional] |
| **orphanet_disorder_type** | **String** | The disorder type in Orphanet&#39;s classification of rare diseases. | [optional] |
| **orphanet_disorder_group** | **String** | The disorder group in Orphanet&#39;s classification of rare diseases. | [optional] |
| **gene_evidence_state** | **String** | known_causal when at least one gene assertion is a disease-causing mutation assessed by Orphanet, weaker_only when only other gene assertions exist, none when there is no gene assertion. | [optional] |
| **n_genes** | **Integer** | Distinct genes asserted for the disease. | [optional] |
| **n_known_causal_genes** | **Integer** | Distinct genes whose assertion is a disease-causing mutation assessed by Orphanet. | [optional] |
| **prevalence_band** | **String** | The Orphanet point-prevalence band chosen for the disease, one of six literals; null when no documented band exists. | [optional] |
| **rarity_status_gb** | **String** | The disease&#39;s rarity status under the UK definition, from the latest assessment (rare, not_rare, unknown or conflicting); null when not assessed. | [optional] |
| **rarity_status_eu** | **String** | The disease&#39;s rarity status under the European Union definition, from the latest assessment (rare, not_rare, unknown or conflicting); null when not assessed. | [optional] |
| **trial_coverage_state** | **String** | no_bridge when the disease maps to no MeSH condition, bridged_no_trials when it maps to a condition with no trial, trials when at least one trial is linked. | [optional] |
| **n_trials_exact_direct** | **Integer** | Trials linked through an exact mapping whose own condition list names the disease. Count within this column only. | [optional] |
| **n_trials_exact_broader** | **Integer** | Trials linked through an exact mapping to a broader condition than the trial studies. Count within this column only. | [optional] |
| **n_trials_unspecified** | **Integer** | Trials linked through a MONDO cross-reference that states no relation. Count within this column only. | [optional] |
| **n_exact_mesh_conditions** | **Integer** | MeSH conditions the disease maps to with an exact relation. | [optional] |
| **merges_into_condition_id** | **String** | The single MeSH condition the web merges this disease into, set only when exactly one exact MeSH match exists. | [optional] |
| **icd10_coding_state** | **String** | How Orphanet aligns the disease to ICD-10: exact, directional, unstated or unaligned. | [optional] |
| **icd11_coding_state** | **String** | How Orphanet aligns the disease to ICD-11: exact, directional, unstated or unaligned. | [optional] |
| **source_release** | **String** | The Orphadata Science release the row was read from. | [optional] |
| **license_name** | **String** | The licence of the Orphadata Science release. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseasesRow.new(
  orpha_code: null,
  disease_key: null,
  preferred_name: null,
  aliases: null,
  concept_type: null,
  status: null,
  ordo_entity_type: null,
  orphanet_disorder_type: null,
  orphanet_disorder_group: null,
  gene_evidence_state: null,
  n_genes: null,
  n_known_causal_genes: null,
  prevalence_band: null,
  rarity_status_gb: null,
  rarity_status_eu: null,
  trial_coverage_state: null,
  n_trials_exact_direct: null,
  n_trials_exact_broader: null,
  n_trials_unspecified: null,
  n_exact_mesh_conditions: null,
  merges_into_condition_id: null,
  icd10_coding_state: null,
  icd11_coding_state: null,
  source_release: null,
  license_name: null,
  attribution: null
)
```

