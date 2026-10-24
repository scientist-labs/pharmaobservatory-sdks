# PharmaObservatory::CompanyFirmographicsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **founded_year** | **Integer** | The earliest of the ROR and Wikidata founding years, else our provider profile, else the GLEIF entity creation year last (GLEIF dates re-registrations and reused shells wrongly). | [optional] |
| **founded_year_source** | **String** | The source that supplied founded year. | [optional] |
| **status** | **String** | public when an organization record of the company carries an SEC filer id (not a Form D-only filer), else private. | [optional] |
| **hq_city** | **String** | Headquarters city: GLEIF, then SEC filing business address, then ROR, then Wikidata. | [optional] |
| **hq_city_source** | **String** | The source that supplied hq city. | [optional] |
| **hq_country** | **String** | Headquarters country: GLEIF, then ROR, then Wikidata. | [optional] |
| **hq_country_source** | **String** | The source that supplied hq country. | [optional] |
| **primary_industry** | **String** | Primary industry: SEC standard industrial classification description, then ROR organization type, then Wikidata industry. | [optional] |
| **primary_industry_source** | **String** | The source that supplied primary industry. | [optional] |
| **website** | **String** | Website: ROR, then Wikidata, then our provider profile. | [optional] |
| **website_source** | **String** | The source that supplied website. | [optional] |
| **employees** | **Integer** | Latest employee headcount from Wikidata; a zero or negative value is treated as absent. | [optional] |
| **employees_source** | **String** | The source that supplied employees. | [optional] |
| **employees_as_of** | **Date** | The date Wikidata states for the headcount. | [optional] |
| **verticals** | **Array&lt;String&gt;** | Service areas from our provider profile (the pipe-separated list, split). | [optional] |
| **verticals_source** | **String** | The source that supplied verticals. | [optional] |
| **legal_name** | **String** | Legal name: GLEIF, then Wikidata, then our provider profile&#39;s official name. | [optional] |
| **legal_name_source** | **String** | The source that supplied legal name. | [optional] |
| **corporate_parent** | **String** | Corporate parent name: GLEIF group name, then Wikidata parent, then our provider profile&#39;s parent. This is a stated source value, not the company&#39;s ladder group. | [optional] |
| **corporate_parent_source** | **String** | The source that supplied corporate parent. | [optional] |
| **legal_form** | **String** | Legal form: GLEIF, then Wikidata. | [optional] |
| **legal_form_source** | **String** | The source that supplied legal form. | [optional] |
| **wikidata_qid** | **String** | The Wikidata item id of the organization record whose profile supplies the five fields below. | [optional] |
| **wikipedia_url** | **String** | Wikipedia page: ROR, then Wikidata. | [optional] |
| **linkedin** | **String** | LinkedIn organization id from the Wikidata profile. | [optional] |
| **x_handle** | **String** | X (Twitter) username from the Wikidata profile. | [optional] |
| **facebook** | **String** | Facebook id from the Wikidata profile. | [optional] |
| **short_description** | **String** | The Wikidata short description. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyFirmographicsRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  founded_year: null,
  founded_year_source: null,
  status: null,
  hq_city: null,
  hq_city_source: null,
  hq_country: null,
  hq_country_source: null,
  primary_industry: null,
  primary_industry_source: null,
  website: null,
  website_source: null,
  employees: null,
  employees_source: null,
  employees_as_of: null,
  verticals: null,
  verticals_source: null,
  legal_name: null,
  legal_name_source: null,
  corporate_parent: null,
  corporate_parent_source: null,
  legal_form: null,
  legal_form_source: null,
  wikidata_qid: null,
  wikipedia_url: null,
  linkedin: null,
  x_handle: null,
  facebook: null,
  short_description: null
)
```

