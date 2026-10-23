# PharmaObservatory::LegalEntitySecFilingsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **org_id** | **String** | Legal-entity level: the SEC filer organization record. | [optional] |
| **org_name** | **String** | Display name of the filer organization record. | [optional] |
| **company_id** | **String** | Company level: the company the filer belongs to (null when unresolved). | [optional] |
| **company_name** | **String** | Display name of the filer&#39;s company. | [optional] |
| **group_id** | **String** | Group level: the filer&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the filer&#39;s group. | [optional] |
| **edgar_cik** | **String** | The filer&#39;s 10-digit SEC Central Index Key (CIK). | [optional] |
| **jurisdiction_of_inc** | **String** | The jurisdiction of incorporation as the latest filing states it (a US state or a country, free text); null when blank. | [optional] |
| **entity_type** | **String** | The entity type the latest filing states (Corporation, Limited Liability Company, Limited Partnership, General Partnership, Business Trust or Other); null when blank. | [optional] |
| **entity_type_other_desc** | **String** | The filer&#39;s free-text description when it ticked entity type Other. | [optional] |
| **year_of_inc_band** | **String** | The filing&#39;s year-of-incorporation choice (Over Five Years Ago, Within Last Five Years or Yet to Be Formed); null when blank. | [optional] |
| **year_of_inc** | **Integer** | The year of incorporation when the filing gives a four-digit year. | [optional] |
| **industry_group_types** | **Array&lt;String&gt;** | The industry groups the filer&#39;s in-scope filings name. | [optional] |
| **previous_names** | **Array&lt;String&gt;** | Earlier names the filer states, other than its current name. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::LegalEntitySecFilingsRow.new(
  org_id: null,
  org_name: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  edgar_cik: null,
  jurisdiction_of_inc: null,
  entity_type: null,
  entity_type_other_desc: null,
  year_of_inc_band: null,
  year_of_inc: null,
  industry_group_types: null,
  previous_names: null
)
```

