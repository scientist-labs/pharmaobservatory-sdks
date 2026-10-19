# PharmaObservatory::SitesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **site_id** | **String** | The site key: one facility spelling from the registry; the id the site page address carries. | [optional] |
| **site_name** | **String** | The facility name as filed in the registry. | [optional] |
| **city** | **String** | City as filed. | [optional] |
| **state** | **String** | State or province as filed. | [optional] |
| **country** | **String** | Country name as filed. | [optional] |
| **latitude** | **Float** | Latitude of the finest coordinate the record has. | [optional] |
| **longitude** | **Float** | Longitude of the finest coordinate the record has. | [optional] |
| **geo_precision** | **String** | How finely the coordinate places the site (city, postal and so on). | [optional] |
| **coord_source** | **String** | The source of the coordinate: aact (the registry&#39;s own), geonames (GeoNames) or als (Amazon Location Service). | [optional] |
| **is_location_source_asserted_only** | **Boolean** | True when the city the registry gave is a location that only the registry asserts for the institution. | [optional] |
| **site_org_id** | **String** | Legal-entity level: the organization record the site&#39;s institution resolved to (null when unresolved). | [optional] |
| **site_company_id** | **String** | Company level: the company the site&#39;s institution belongs to (null when unresolved). | [optional] |
| **site_company_name** | **String** | Display name of the site&#39;s institution&#39;s company. | [optional] |
| **site_group_id** | **String** | Group level: the site&#39;s institution&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **site_group_name** | **String** | Display name of the site&#39;s institution&#39;s group. | [optional] |
| **n_trials** | **Integer** | Registered trials that list the site. | [optional] |
| **n_active** | **Integer** | Those trials whose status is ongoing. | [optional] |
| **first_year** | **Integer** | Year of the earliest trial start at the site. | [optional] |
| **last_year** | **Integer** | Year of the latest trial start at the site. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SitesRow.new(
  site_id: null,
  site_name: null,
  city: null,
  state: null,
  country: null,
  latitude: null,
  longitude: null,
  geo_precision: null,
  coord_source: null,
  is_location_source_asserted_only: null,
  site_org_id: null,
  site_company_id: null,
  site_company_name: null,
  site_group_id: null,
  site_group_name: null,
  n_trials: null,
  n_active: null,
  first_year: null,
  last_year: null
)
```

