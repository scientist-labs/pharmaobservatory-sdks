# PharmaObservatory::CompanyTrialSitesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | The company key (company:&lt;key&gt;) of the lead sponsor. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **roster_root_id** | **String** | The Clinical Research Map&#39;s organization key for these rows: the group when it is a public roster root, else the company when it is one, else null. | [optional] |
| **site_id** | **String** | The site key: one facility spelling from the registry; the id the site page address carries. | [optional] |
| **facility_name** | **String** | The facility name as filed in the registry. | [optional] |
| **city** | **String** | City as filed. | [optional] |
| **state** | **String** | State or province as filed. | [optional] |
| **country** | **String** | Country name as filed. | [optional] |
| **site_country_code** | **String** | ISO 3166-1 alpha-2 code of the site&#39;s country, from the registry&#39;s country name; null when the name resolves to no country. | [optional] |
| **latitude** | **Float** | Latitude of the finest coordinate the site record has. | [optional] |
| **longitude** | **Float** | Longitude of the finest coordinate the site record has. | [optional] |
| **geo_precision** | **String** | How finely the coordinate places the site (city, postal and so on). | [optional] |
| **coord_source** | **String** | The source of the coordinate: aact (the registry&#39;s own), geonames (GeoNames) or als (Amazon Location Service). | [optional] |
| **n_trials** | **Integer** | Distinct trials the company leads that list the site. | [optional] |
| **site_rank** | **Integer** | 1 is the company&#39;s busiest located site (most led trials, then site key); dense per company. | [optional] |
| **n_sites** | **Integer** | The company&#39;s located site count, repeated on each of its rows. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyTrialSitesRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  roster_root_id: null,
  site_id: null,
  facility_name: null,
  city: null,
  state: null,
  country: null,
  site_country_code: null,
  latitude: null,
  longitude: null,
  geo_precision: null,
  coord_source: null,
  n_trials: null,
  site_rank: null,
  n_sites: null
)
```

