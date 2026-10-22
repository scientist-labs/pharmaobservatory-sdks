# PharmaObservatory::RareDiseaseOrganizationsRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **connection_kind** | **String** | disease_relation when connection_value is a relation to the disease, disease_function when it is a function the organization performs for it. | [optional] |
| **assertion_id** | **String** | The identifier of the assertion. | [optional] |
| **orpha_code** | **String** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] |
| **disease_name** | **String** | The disease&#39;s preferred English name. | [optional] |
| **organization_name** | **String** | The organization&#39;s name as filed. | [optional] |
| **organization_org_id** | **String** | Legal-entity level: the organization record the assertion is on. | [optional] |
| **organization_company_id** | **String** | Company level: the organization&#39;s company (null when unresolved). | [optional] |
| **organization_company_name** | **String** | Display name of the organization&#39;s company. | [optional] |
| **organization_group_id** | **String** | Group level: the organization&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] |
| **organization_group_name** | **String** | Display name of the organization&#39;s group. | [optional] |
| **connection_value** | **String** | The relation (a predicate) or the function, according to connection_kind. | [optional] |
| **evidence_type** | **String** | The kind of evidence the assertion rests on. | [optional] |
| **authority_class** | **String** | How authoritative the source is for the claim. | [optional] |
| **source_key** | **String** | The source the assertion was read from. | [optional] |
| **source_record_id** | **String** | The record in the source the assertion comes from. | [optional] |
| **source_url** | **String** | A link to the source record, where it has one. | [optional] |
| **evidence_trial_id** | **String** | The registry trial id the assertion names, where it names one. | [optional] |
| **observed_at** | **Time** | When the assertion was observed in the source (UTC). | [optional] |
| **last_verified_at** | **Time** | When the assertion was last verified against the source (UTC). | [optional] |
| **confidence** | **Float** | The derivation&#39;s confidence in the assertion. | [optional] |
| **license_name** | **String** | The licence of the source the assertion was read from. | [optional] |
| **attribution** | **String** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::RareDiseaseOrganizationsRow.new(
  connection_kind: null,
  assertion_id: null,
  orpha_code: null,
  disease_name: null,
  organization_name: null,
  organization_org_id: null,
  organization_company_id: null,
  organization_company_name: null,
  organization_group_id: null,
  organization_group_name: null,
  connection_value: null,
  evidence_type: null,
  authority_class: null,
  source_key: null,
  source_record_id: null,
  source_url: null,
  evidence_trial_id: null,
  observed_at: null,
  last_verified_at: null,
  confidence: null,
  license_name: null,
  attribution: null
)
```

