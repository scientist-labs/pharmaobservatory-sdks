# PharmaObservatory::CompanyUsPatentsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **patent_id** | **String** | The PatentsView patent id. | [optional] |
| **patent_number** | **String** | The US patent number. | [optional] |
| **patent_title** | **String** | The patent title. | [optional] |
| **patent_date** | **Date** | The grant date. | [optional] |
| **is_biomedical** | **Boolean** | True when the patent carries a biomedical CPC class; the split is never a filter on what belongs to the portfolio. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyUsPatentsRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  patent_id: null,
  patent_number: null,
  patent_title: null,
  patent_date: null,
  is_biomedical: null
)
```

