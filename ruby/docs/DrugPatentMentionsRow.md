# PharmaObservatory::DrugPatentMentionsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **molecule_id** | **String** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] |
| **patent_number** | **String** | The patent publication number. | [optional] |
| **country** | **String** | The patent office country code. | [optional] |
| **publication_date** | **Date** | The publication date. | [optional] |
| **title** | **String** | The patent title as filed. | [optional] |
| **assignee** | **String** | The assignee as filed. Not an ownership statement. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::DrugPatentMentionsRow.new(
  molecule_id: null,
  patent_number: null,
  country: null,
  publication_date: null,
  title: null,
  assignee: null
)
```

