# PharmaObservatory::WarningLettersRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **letter_id** | **String** | The letter id. | [optional] |
| **letter_url** | **String** | The FDA page of the letter. | [optional] |
| **firm_name** | **String** | The firm&#39;s name as filed with the regulator. | [optional] |
| **firm_org_id** | **String** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] |
| **firm_company_id** | **String** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] |
| **firm_company_name** | **String** | Display name of the firm&#39;s company. | [optional] |
| **firm_group_id** | **String** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **firm_group_name** | **String** | Display name of the firm&#39;s group. | [optional] |
| **is_org_resolved** | **Boolean** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] |
| **match_tier** | **String** | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous (several candidates, none chosen) or unresolved. Warning letters, untitled letters and debarments attach by firm name alone. | [optional] |
| **issuing_office** | **String** | The FDA office that issued the letter (a center or a district office). | [optional] |
| **subject** | **String** | The letter&#39;s subject line. | [optional] |
| **excerpt** | **String** | An excerpt of the letter text. | [optional] |
| **issue_date** | **Date** | The date the letter was issued. | [optional] |
| **posted_date** | **Date** | The date the letter was posted on the FDA site. | [optional] |
| **response_letter_url** | **String** | The firm&#39;s response letter on the FDA site, where one is posted. | [optional] |
| **closeout_letter_url** | **String** | The FDA close-out letter, where one is posted. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::WarningLettersRow.new(
  letter_id: null,
  letter_url: null,
  firm_name: null,
  firm_org_id: null,
  firm_company_id: null,
  firm_company_name: null,
  firm_group_id: null,
  firm_group_name: null,
  is_org_resolved: null,
  match_tier: null,
  issuing_office: null,
  subject: null,
  excerpt: null,
  issue_date: null,
  posted_date: null,
  response_letter_url: null,
  closeout_letter_url: null
)
```

