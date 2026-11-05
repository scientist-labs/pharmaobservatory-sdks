# PharmaObservatory::CompanyAssetCollisionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **molecule_id** | **String** | Molecule level: the company&#39;s asset, the resolved drug (a drug with two canonical ids is one asset). | [optional] |
| **molecule_name** | **String** | The asset&#39;s display name, from the company&#39;s most-trialled record of it. | [optional] |
| **competitor_company_id** | **String** | Company level: the company whose asset collides. | [optional] |
| **competitor_company_name** | **String** | Display name of the competitor&#39;s company. | [optional] |
| **competitor_group_id** | **String** | Group level: the competitor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **competitor_group_name** | **String** | Display name of the competitor&#39;s group. | [optional] |
| **competitor_molecule_id** | **String** | Molecule level: the competitor&#39;s asset, the resolved drug. | [optional] |
| **competitor_molecule_name** | **String** | The competitor asset&#39;s display name. | [optional] |
| **target_family** | **String** | The company drug&#39;s target family of the shared gene (the Open Targets target name, so a taxane&#39;s tubulin isotypes are one family; the ChEMBL mechanism text for a drug Open Targets lacks; the gene symbol for a curated investigational asset). | [optional] |
| **shared_symbols** | **Array&lt;String&gt;** | Gene symbols of the shared genes within the family. | [optional] |
| **action_type** | **String** | The company drug&#39;s action on the family (Open Targets or ChEMBL vocabulary); null when no source states one. | [optional] |
| **competitor_action_type** | **String** | The competitor drug&#39;s action on the shared genes; null when no source states one. | [optional] |
| **same_action** | **Boolean** | True when the two drugs have the same action type on a shared gene of the family. | [optional] |
| **n_shared_indications** | **Integer** | Disease areas both programs are in, among live programs. | [optional] |
| **top_shared_indications** | **Array&lt;String&gt;** | Up to 3 shared disease area names, by the pair&#39;s combined trial count. An empty list when none. | [optional] |
| **stage** | **Integer** | The less advanced of the two programs&#39; furthest phase: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1, -1 phase not applicable. | [optional] |
| **last_year** | **Integer** | The later of the two programs&#39; last trial start year. | [optional] |
| **n_head_to_head_trials** | **Integer** | Trials led by either company in which one drug is experimental and the other an active comparator. | [optional] |
| **ownership_basis** | **String** | Why the company owns the drug. experimental_arm and regulatory_holder rest on evidence; inferred_lead is an inference (no arm data says otherwise). | [optional] |
| **competitor_ownership_basis** | **String** | Why the competitor owns its drug, same vocabulary as ownership_basis. | [optional] |
| **family_idf** | **Float** | The rarity weight of the family: the natural log of (companies holding any live, owned target family) over (companies holding this one). | [optional] |
| **collision_weight** | **Float** | The weight of the pair, family_idf times a factor for same action, a shared indication, the stage, recency and a head-to-head trial. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyAssetCollisionsRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  molecule_id: null,
  molecule_name: null,
  competitor_company_id: null,
  competitor_company_name: null,
  competitor_group_id: null,
  competitor_group_name: null,
  competitor_molecule_id: null,
  competitor_molecule_name: null,
  target_family: null,
  shared_symbols: null,
  action_type: null,
  competitor_action_type: null,
  same_action: null,
  n_shared_indications: null,
  top_shared_indications: null,
  stage: null,
  last_year: null,
  n_head_to_head_trials: null,
  ownership_basis: null,
  competitor_ownership_basis: null,
  family_idf: null,
  collision_weight: null
)
```

