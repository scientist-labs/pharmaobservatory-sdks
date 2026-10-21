# PharmaObservatory::SuppliersRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **provider_id** | **String** | The supplier id: the Scientist.com provider id as text, the id in the supplier page address (/suppliers/&lt;id&gt;). A supplier is not a company. | [optional] |
| **name** | **String** | The supplier&#39;s name as listed on the marketplace. | [optional] |
| **company_type** | **String** | The supplier&#39;s company type as the marketplace spells it (small-business, academic-core and so on). An open set; null when the supplier lists none. | [optional] |
| **marketplace_tier** | **String** | The supplier&#39;s marketplace tier word; null when the marketplace assigns none. | [optional] |
| **service_areas** | **Array&lt;String&gt;** | The service areas the supplier lists, one array element each (empty when none). Filter with &#39;Biology&#39; &#x3D; ANY(service_areas). | [optional] |
| **hq_country** | **String** | ISO 3166-1 alpha-2 country code of the supplier&#39;s headquarters point (null when unknown). | [optional] |
| **hq_region** | **String** | The broad world region of hq_country, from the country_regions reference table; other when the country is unknown or in none of the six regions. | [optional] |
| **hq_latitude** | **Float** | Latitude of the supplier&#39;s headquarters point: the Headquarters site, else the first site by category then site id. | [optional] |
| **hq_longitude** | **Float** | Longitude of the supplier&#39;s headquarters point. | [optional] |
| **n_sites** | **Integer** | How many located sites the supplier has. | [optional] |
| **n_countries** | **Integer** | How many distinct countries the supplier&#39;s sites are in. | [optional] |
| **is_cro** | **Boolean** | True when the supplier&#39;s organization is a contract research organization (CRO) under the one organization-level definition. | [optional] |
| **is_trial_sponsor** | **Boolean** | True when the supplier name matches a registered trial sponsor. A name match, not a verified identity. | [optional] |
| **n_trials** | **Integer** | Registered trials where an organization matched to the supplier by name is lead sponsor or collaborator. | [optional] |
| **n_lead_trials** | **Integer** | The subset of n_trials where the matched organization is the lead sponsor. | [optional] |
| **website** | **String** | The website the supplier listed; a third-party address, not vouched for. | [optional] |
| **profile_url** | **String** | The supplier&#39;s page on the Scientist.com marketplace; the marketplace asks a visitor to sign in. | [optional] |
| **logo_path** | **String** | The path of the supplier&#39;s logo on our own origin (/supplier-logos/&lt;sha&gt;.png); null when it has none. | [optional] |
| **has_description** | **Boolean** | True when the supplier wrote a profile description. The description text is not published. | [optional] |
| **is_veteran_owned** | **Boolean** | SAM.gov self-certified veteran-owned designation. | [optional] |
| **is_woman_owned** | **Boolean** | SAM.gov self-certified woman-owned designation. | [optional] |
| **is_minority_owned** | **Boolean** | SAM.gov self-certified minority-owned designation (8(a) or small disadvantaged business). | [optional] |
| **is_hubzone** | **Boolean** | SAM.gov self-certified HUBZone (historically underutilized business zone) designation. | [optional] |
| **supplier_org_id** | **String** | Legal-entity level: the organization record the supplier resolved to. Null when the supplier has no organization record of its own. | [optional] |
| **supplier_company_id** | **String** | Company level: the company the supplier&#39;s organization belongs to (null when none). | [optional] |
| **supplier_company_name** | **String** | Display name of the supplier&#39;s company. | [optional] |
| **supplier_group_id** | **String** | Group level: the supplier&#39;s company&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **supplier_group_name** | **String** | Display name of the supplier&#39;s group. | [optional] |
| **org_link_method** | **String** | How the supplier was linked to its organization record, so a caller can apply a confidence floor. The supplier pages on organization pages show only provider_ror_exact, provider_edgar_exact and supplier_curated. | [optional] |
| **has_public_page** | **Boolean** | True when the supplier has a public page on the web (every row of this dataset does). | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SuppliersRow.new(
  provider_id: null,
  name: null,
  company_type: null,
  marketplace_tier: null,
  service_areas: null,
  hq_country: null,
  hq_region: null,
  hq_latitude: null,
  hq_longitude: null,
  n_sites: null,
  n_countries: null,
  is_cro: null,
  is_trial_sponsor: null,
  n_trials: null,
  n_lead_trials: null,
  website: null,
  profile_url: null,
  logo_path: null,
  has_description: null,
  is_veteran_owned: null,
  is_woman_owned: null,
  is_minority_owned: null,
  is_hubzone: null,
  supplier_org_id: null,
  supplier_company_id: null,
  supplier_company_name: null,
  supplier_group_id: null,
  supplier_group_name: null,
  org_link_method: null,
  has_public_page: null
)
```

