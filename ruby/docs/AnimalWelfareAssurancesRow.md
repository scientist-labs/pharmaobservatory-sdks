# PharmaObservatory::AnimalWelfareAssurancesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **institution_org_id** | **String** | Legal-entity level: the organization record the OLAW institution resolved to; the entry key for the legal_entities object type. | [optional] |
| **institution_company_id** | **String** | Company level: the organization&#39;s company (null when the organization is not in a company); the entry key for the companies object type. | [optional] |
| **institution_company_name** | **String** | Display name of the organization&#39;s company. | [optional] |
| **institution_group_id** | **String** | Group level: the organization&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **institution_group_name** | **String** | Display name of the organization&#39;s group. | [optional] |
| **source** | **String** | The animal-welfare source: nih_olaw today. A second source (the Canadian Council on Animal Care list) can be added as another value without changing the grain. | [optional] |
| **assurance_id** | **String** | The OLAW assurance id of the organization&#39;s best evidence (a primary assurance beats a covered component). | [optional] |
| **assurance_status_code** | **String** | OLAW&#39;s raw two-digit assurance status code as published (for example 02 or 03). Never decoded, because no OLAW page documents it. | [optional] |
| **assurance_kind** | **String** | OLAW&#39;s PRIMARY_OR_COVERED_COMP field as published: PRIMARY or COVERED_COMP (a covered component of a primary assurance). | [optional] |
| **is_domestic** | **Boolean** | True for a domestic (US) assurance, false for a foreign one, from OLAW&#39;s assurance type (D or F). | [optional] |
| **match_tier** | **String** | How well the OLAW institution name reached the organization: high (name, country and state agree), medium (name and country agree) or low (name only). The web badge shows high and medium. | [optional] |
| **assurance_count** | **Integer** | How many OLAW assurance rows resolved to this organization (a university and its covered components). | [optional] |
| **assurance_ids** | **Array&lt;String&gt;** | Every OLAW assurance id that resolved to this organization, sorted. | [optional] |
| **institution_names** | **Array&lt;String&gt;** | The institution names as OLAW lists them, for every assurance that resolved to this organization. | [optional] |
| **source_asof** | **Time** | When OLAW&#39;s lookup tool says its list was last updated; the source does not state a timezone. | [optional] |
| **source_credit** | **String** | The plain source credit line to show beside any figure from this dataset. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::AnimalWelfareAssurancesRow.new(
  institution_org_id: null,
  institution_company_id: null,
  institution_company_name: null,
  institution_group_id: null,
  institution_group_name: null,
  source: null,
  assurance_id: null,
  assurance_status_code: null,
  assurance_kind: null,
  is_domestic: null,
  match_tier: null,
  assurance_count: null,
  assurance_ids: null,
  institution_names: null,
  source_asof: null,
  source_credit: null
)
```

