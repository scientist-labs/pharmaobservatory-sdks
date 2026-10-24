# PharmaObservatory::CompanyCompetitorActivityRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **company_id** | **String** | Company level: the company the row is about; the entry key for the companies object type. | [optional] |
| **company_name** | **String** | Display name of the company. | [optional] |
| **group_id** | **String** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the group. | [optional] |
| **condition_id** | **String** | The condition key, a MeSH (Medical Subject Headings) descriptor id (D followed by six to nine digits). | [optional] |
| **condition_name** | **String** | The MeSH descriptor name of the condition. | [optional] |
| **disease_rank** | **Integer** | The condition&#39;s rank among the company&#39;s disease areas by trial count (1 is the busiest). Only the top 10 are kept. | [optional] |
| **competitor_company_id** | **String** | Company level: the company active in the same disease area. | [optional] |
| **competitor_company_name** | **String** | Display name of the competitor&#39;s company. | [optional] |
| **competitor_group_id** | **String** | Group level: the competitor&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **competitor_group_name** | **String** | Display name of the competitor&#39;s group. | [optional] |
| **competitor_rank** | **Integer** | 0 for the company itself, 1 to 6 for the companies ranked by trial count across the company&#39;s top disease areas. | [optional] |
| **trials** | **Integer** | Distinct trials the company in this row has in the disease area. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::CompanyCompetitorActivityRow.new(
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  condition_id: null,
  condition_name: null,
  disease_rank: null,
  competitor_company_id: null,
  competitor_company_name: null,
  competitor_group_id: null,
  competitor_group_name: null,
  competitor_rank: null,
  trials: null
)
```

