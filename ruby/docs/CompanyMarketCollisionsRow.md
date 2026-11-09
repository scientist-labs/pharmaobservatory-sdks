# PharmaObservatory::CompanyMarketCollisionsRow

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
| **condition_ui** | **String** | The shared MeSH indication (a MeSH descriptor id): both assets are owned, live and Phase 3 or 4 in it. | [optional] |
| **condition_name** | **String** | The MeSH name of the shared indication. | [optional] |
| **target_symbols** | **Array&lt;String&gt;** | Gene symbols of the company drug&#39;s recorded targets. Never empty, and none is a target of the competitor drug. | [optional] |
| **competitor_target_symbols** | **Array&lt;String&gt;** | Gene symbols of the competitor drug&#39;s recorded targets. | [optional] |
| **stage** | **Integer** | The less advanced of the two programs&#39; furthest phase: 6 Phase 4, 5 Phase 3, 4 Phase 2/3, 3 Phase 2, 2 Phase 1/2, 1 Phase 1, 0 Early Phase 1, -1 phase not applicable. | [optional] |
| **last_year** | **Integer** | The later of the two programs&#39; last trial start year. | [optional] |
| **n_head_to_head_trials** | **Integer** | Trials led by either company in which one drug is experimental and the other an active comparator. | [optional] |
| **ownership_basis** | **String** | Why the company owns the drug. experimental_arm and regulatory_holder rest on evidence; inferred_lead is an inference (no arm data says otherwise). | [optional] |
| **competitor_ownership_basis** | **String** | Why the competitor owns its drug, same vocabulary as ownership_basis. | [optional] |
| **indication_idf** | **Float** | The rarity weight of the indication: the natural log of (companies holding a live, owned Phase 3+ asset with a recorded target) over (companies holding one in this indication). | [optional] |
| **market_weight** | **Float** | The weight of the pair, 0.5 times indication_idf times factors for recency and a head-to-head trial. There is no stage factor, both programs are Phase 3 or 4. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyMarketCollisionsRow.new(
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
  condition_ui: null,
  condition_name: null,
  target_symbols: null,
  competitor_target_symbols: null,
  stage: null,
  last_year: null,
  n_head_to_head_trials: null,
  ownership_basis: null,
  competitor_ownership_basis: null,
  indication_idf: null,
  market_weight: null
)
```

