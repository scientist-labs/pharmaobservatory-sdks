# PharmaObservatory::CompanyTrialGeographyRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | The company key (company:&lt;key&gt;). | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **is_public_roster_root** | **Boolean** | True when the company is a public roster root: an organization the public Clinical Research Map lists. | [optional] |
| **roster_root_id** | **String** | The Clinical Research Map&#39;s organization key for this company: the group when it is a public roster root, else the company when it is one, else null. | [optional] |
| **roster_lead_trials** | **Integer** | For a public roster root, the lead-trial count the public organizations list sorts on; null for a company that is not one. | [optional] |
| **domicile_country** | **String** | ISO 3166-1 alpha-2 code of the company&#39;s legal-entity domicile; null when unknown. | [optional] |
| **site_country** | **String** | ISO 3166-1 alpha-2 code of the country holding the most of the company&#39;s own located led-trial sites; null when none resolves. | [optional] |
| **map_country** | **String** | The country the map places the company under: domicile_country, else site_country; null when neither is known. | [optional] |
| **map_country_basis** | **String** | Which signal map_country came from: domicile or trial_sites; null exactly when map_country is null. | [optional] |
| **n_located_sites** | **Integer** | The company&#39;s own located led-trial sites. | [optional] |
| **n_unlocated_sites** | **Integer** | The company&#39;s led-trial sites with no coordinate, which the map does not draw. | [optional] |
| **n_site_countries** | **Integer** | Distinct resolved countries among the company&#39;s located led-trial sites. | [optional] |
| **n_unresolved_sites** | **Integer** | Located led-trial sites whose country name resolves to no country code; they do not vote for site_country. | [optional] |
| **roster_map_country** | **String** | For a public roster root, the country the public Clinical Research Map shows the organization under, as published; null for a non-root. | [optional] |
| **roster_map_country_basis** | **String** | For a public roster root, which signal the public map&#39;s country came from: domicile or trial_sites; null for a non-root or an unknown country. | [optional] |
| **roster_n_located_sites** | **Integer** | For a public roster root, the located sites the public map draws the organization at (its group fold); null for a non-root. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyTrialGeographyRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  is_public_roster_root: null,
  roster_root_id: null,
  roster_lead_trials: null,
  domicile_country: null,
  site_country: null,
  map_country: null,
  map_country_basis: null,
  n_located_sites: null,
  n_unlocated_sites: null,
  n_site_countries: null,
  n_unresolved_sites: null,
  roster_map_country: null,
  roster_map_country_basis: null,
  roster_n_located_sites: null
)
```

