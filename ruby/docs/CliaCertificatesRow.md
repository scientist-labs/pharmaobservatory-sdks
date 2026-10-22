# PharmaObservatory::CliaCertificatesRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **clia_number** | **String** | The ten-character CLIA certificate number. | [optional] |
| **lab_name** | **String** | The laboratory&#39;s name as listed on the certificate. | [optional] |
| **firm_name** | **String** | The organization the certificate resolved to (its canonical name). | [optional] |
| **firm_org_id** | **String** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] |
| **firm_company_id** | **String** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] |
| **firm_company_name** | **String** | Display name of the firm&#39;s company. | [optional] |
| **firm_group_id** | **String** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **firm_group_name** | **String** | Display name of the firm&#39;s group. | [optional] |
| **match_tier** | **String** | How the certificate was attached to the organization: high, medium or low. | [optional] |
| **match_confidence** | **Float** | The matcher&#39;s confidence, 0 to 1. | [optional] |
| **city** | **String** | The laboratory&#39;s city. | [optional] |
| **state_region** | **String** | The laboratory&#39;s state. | [optional] |
| **postal5** | **String** | The five-digit postal code. | [optional] |
| **country** | **String** | The country code. | [optional] |
| **latitude** | **Float** | Latitude in degrees (WGS84), for the map tool. | [optional] |
| **longitude** | **Float** | Longitude in degrees (WGS84), for the map tool. | [optional] |
| **geocode_precision** | **String** | What a pin locates (an address, a postal code or a city). | [optional] |
| **coord_source** | **String** | Where the coordinates came from. | [optional] |
| **is_active** | **Boolean** | True when the certificate is active on the file. | [optional] |
| **certificate_type** | **String** | The CLIA certificate type (for example Compliance, Accreditation or Waiver). | [optional] |
| **complexity_ceiling** | **String** | The highest test complexity the certificate permits. | [optional] |
| **accreditor_claimed** | **Array&lt;String&gt;** | The accreditation body the file names. | [optional] |
| **accreditor_confirmed** | **Array&lt;String&gt;** | The accreditation body confirmed against the accreditor&#39;s own list. | [optional] |
| **has_cap_accreditation** | **Boolean** | True when the laboratory holds College of American Pathologists (CAP) accreditation. | [optional] |
| **has_ashi_accreditation** | **Boolean** | True when the laboratory holds American Society for Histocompatibility and Immunogenetics (ASHI) accreditation. | [optional] |
| **has_aabb_accreditation** | **Boolean** | True when the laboratory holds AABB (formerly American Association of Blood Banks) accreditation. | [optional] |
| **facility_type** | **String** | The laboratory&#39;s facility type on the file. | [optional] |
| **annual_test_volume** | **Integer** | Reported annual test volume. | [optional] |
| **test_volume_source** | **String** | Where the volume figure came from. | [optional] |
| **lab_site_count** | **Integer** | How many sites the laboratory reports. | [optional] |
| **pos_vintage** | **String** | The vintage of the CMS Provider of Services file. | [optional] |
| **as_of** | **Date** | The build date of the figures. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CliaCertificatesRow.new(
  clia_number: null,
  lab_name: null,
  firm_name: null,
  firm_org_id: null,
  firm_company_id: null,
  firm_company_name: null,
  firm_group_id: null,
  firm_group_name: null,
  match_tier: null,
  match_confidence: null,
  city: null,
  state_region: null,
  postal5: null,
  country: null,
  latitude: null,
  longitude: null,
  geocode_precision: null,
  coord_source: null,
  is_active: null,
  certificate_type: null,
  complexity_ceiling: null,
  accreditor_claimed: null,
  accreditor_confirmed: null,
  has_cap_accreditation: null,
  has_ashi_accreditation: null,
  has_aabb_accreditation: null,
  facility_type: null,
  annual_test_volume: null,
  test_volume_source: null,
  lab_site_count: null,
  pos_vintage: null,
  as_of: null
)
```

