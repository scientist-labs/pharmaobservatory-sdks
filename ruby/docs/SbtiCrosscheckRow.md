# PharmaObservatory::SbtiCrosscheckRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **org_id** | **String** | Legal-entity level: the organization the cross-check is for. | [optional] |
| **org_name** | **String** | Display name of the organization record. | [optional] |
| **company_id** | **String** | Company level: the company the organization belongs to (null when unresolved). | [optional] |
| **company_name** | **String** | Display name of the organization&#39;s company. | [optional] |
| **group_id** | **String** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the organization&#39;s group. | [optional] |
| **reported_company_name** | **String** | The company name as the SBTi list or the disclosure gives it. | [optional] |
| **recon_status** | **String** | The cross-check result: agree, extracted_only, sbti_only, target_year_conflict or self_claim_conflict. | [optional] |
| **has_net_zero_target** | **Boolean** | Whether a net-zero target is stated (SBTi&#39;s value where it lists the organization, else the disclosure&#39;s). | [optional] |
| **net_zero_year** | **Integer** | The net-zero target year. | [optional] |
| **has_near_term_target** | **Boolean** | Whether a near-term target is stated. | [optional] |
| **near_term_target_year** | **Integer** | The near-term target year. | [optional] |
| **sbti_validated** | **Boolean** | Whether SBTi lists the targets as validated; null when SBTi does not list the organization. | [optional] |
| **claimed_sbti_status** | **String** | The SBTi status the organization&#39;s own filing claims. | [optional] |
| **in_sbti** | **Boolean** | Whether the SBTi list carries the organization. | [optional] |
| **in_extracted** | **Boolean** | Whether an extraction of the organization&#39;s own filing exists. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::SbtiCrosscheckRow.new(
  org_id: null,
  org_name: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  reported_company_name: null,
  recon_status: null,
  has_net_zero_target: null,
  net_zero_year: null,
  has_near_term_target: null,
  near_term_target_year: null,
  sbti_validated: null,
  claimed_sbti_status: null,
  in_sbti: null,
  in_extracted: null
)
```

