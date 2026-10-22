# PharmaObservatory::RecallsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **recall_number** | **String** | The FDA recall number. | [optional] |
| **domain** | **String** | The regulatory domain of the recalled product (drug or device here). | [optional] |
| **status** | **String** | The recall status (Ongoing, Completed or Terminated). | [optional] |
| **recall_class** | **String** | The severity class: Class I, II or III. | [optional] |
| **product_type** | **String** | The product type the FDA reports (Drugs, Devices and so on). | [optional] |
| **firm_name** | **String** | The firm&#39;s name as filed with the regulator (the recalling firm). | [optional] |
| **firm_org_id** | **String** | Legal-entity level: the organization record the firm resolved to (null when the firm did not resolve). | [optional] |
| **firm_company_id** | **String** | Company level: the firm&#39;s company (null when unresolved); the entry key for the companies object type. | [optional] |
| **firm_company_name** | **String** | Display name of the firm&#39;s company. | [optional] |
| **firm_group_id** | **String** | Group level: the firm&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **firm_group_name** | **String** | Display name of the firm&#39;s group. | [optional] |
| **is_org_resolved** | **Boolean** | True when the firm was attached to an organization. A false row is kept with null firm ids and the name as filed, so an empty read for a company is not proof of no enforcement history. | [optional] |
| **match_tier** | **String** | How well the firm name was attached to an organization: high, medium or low name-and-geography matches, registry (the operating firm of a registered facility), ambiguous (several candidates, none chosen) or unresolved. Warning letters, untitled letters and debarments attach by firm name alone. | [optional] |
| **org_resolution_method** | **String** | The route that attached the firm: fei_bridge (the FDA establishment identifier, FEI, bridge), fei_firm_mint or fei_firm_fold (the operating firm&#39;s registry record), firm_name_geo (name and geography), firm_name_only (name alone) or unresolved. | [optional] |
| **product_description** | **String** | The recalled product as described in the report. | [optional] |
| **reason_for_recall** | **String** | The reason the FDA gives for the recall. | [optional] |
| **recall_initiation_date** | **Date** | The date the firm began the recall. | [optional] |
| **report_date** | **Date** | The date the FDA reported the recall. | [optional] |
| **termination_date** | **Date** | The date the recall was terminated, where it was. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RecallsRow.new(
  recall_number: null,
  domain: null,
  status: null,
  recall_class: null,
  product_type: null,
  firm_name: null,
  firm_org_id: null,
  firm_company_id: null,
  firm_company_name: null,
  firm_group_id: null,
  firm_group_name: null,
  is_org_resolved: null,
  match_tier: null,
  org_resolution_method: null,
  product_description: null,
  reason_for_recall: null,
  recall_initiation_date: null,
  report_date: null,
  termination_date: null
)
```

