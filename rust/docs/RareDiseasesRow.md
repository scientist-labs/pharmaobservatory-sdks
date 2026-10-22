# RareDiseasesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**orpha_code** | Option<**String**> | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional]
**disease_key** | Option<**String**> | The disease key (orpha: plus the ORPHA code). An attribute, not an id: the entry key is orpha_code. | [optional]
**preferred_name** | Option<**String**> | The disease's preferred English name. | [optional]
**aliases** | Option<**Vec<String>**> | English synonyms Orphanet lists for the disease. | [optional]
**concept_type** | Option<**String**> | Whether the concept is a group of disorders, a disorder or a subtype. | [optional]
**status** | Option<**String**> | The concept's lifecycle state in Orphanet (active or deprecated). | [optional]
**ordo_entity_type** | Option<**String**> | The entity type the Orphanet Rare Disease Ontology gives the concept. | [optional]
**orphanet_disorder_type** | Option<**String**> | The disorder type in Orphanet's classification of rare diseases. | [optional]
**orphanet_disorder_group** | Option<**String**> | The disorder group in Orphanet's classification of rare diseases. | [optional]
**gene_evidence_state** | Option<**String**> | known_causal when at least one gene assertion is a disease-causing mutation assessed by Orphanet, weaker_only when only other gene assertions exist, none when there is no gene assertion. | [optional]
**n_genes** | Option<**i64**> | Distinct genes asserted for the disease. | [optional]
**n_known_causal_genes** | Option<**i64**> | Distinct genes whose assertion is a disease-causing mutation assessed by Orphanet. | [optional]
**prevalence_band** | Option<**String**> | The Orphanet point-prevalence band chosen for the disease, one of six literals; null when no documented band exists. | [optional]
**rarity_status_gb** | Option<**String**> | The disease's rarity status under the UK definition, from the latest assessment (rare, not_rare, unknown or conflicting); null when not assessed. | [optional]
**rarity_status_eu** | Option<**String**> | The disease's rarity status under the European Union definition, from the latest assessment (rare, not_rare, unknown or conflicting); null when not assessed. | [optional]
**trial_coverage_state** | Option<**String**> | no_bridge when the disease maps to no MeSH condition, bridged_no_trials when it maps to a condition with no trial, trials when at least one trial is linked. | [optional]
**n_trials_exact_direct** | Option<**i64**> | Trials linked through an exact mapping whose own condition list names the disease. Count within this column only. | [optional]
**n_trials_exact_broader** | Option<**i64**> | Trials linked through an exact mapping to a broader condition than the trial studies. Count within this column only. | [optional]
**n_trials_unspecified** | Option<**i64**> | Trials linked through a MONDO cross-reference that states no relation. Count within this column only. | [optional]
**n_exact_mesh_conditions** | Option<**i64**> | MeSH conditions the disease maps to with an exact relation. | [optional]
**merges_into_condition_id** | Option<**String**> | The single MeSH condition the web merges this disease into, set only when exactly one exact MeSH match exists. | [optional]
**icd10_coding_state** | Option<**String**> | How Orphanet aligns the disease to ICD-10: exact, directional, unstated or unaligned. | [optional]
**icd11_coding_state** | Option<**String**> | How Orphanet aligns the disease to ICD-11: exact, directional, unstated or unaligned. | [optional]
**source_release** | Option<**String**> | The Orphadata Science release the row was read from. | [optional]
**license_name** | Option<**String**> | The licence of the Orphadata Science release. | [optional]
**attribution** | Option<**String**> | The attribution line the source requires; quote it with any figure taken from the row. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


