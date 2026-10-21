# PharmaObservatory::ManagersRow

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **manager_id** | **String** | The manager key, the id the web page uses: edgar_cik:&lt;10-digit filer CIK&gt;, or party:sec-cik:&lt;ref&gt; for a filer whose CIK is malformed. The same value as holder_identity_key on institutional_holdings. | [optional] |
| **name** | **String** | The manager&#39;s name as filed (the longest variant), falling back to the manager key. | [optional] |
| **name_variant_count** | **Integer** | How many spellings of the name the filings carry. | [optional] |
| **investor_id** | **String** | The manager&#39;s investor key where the link resolves; null when the manager resolves to no investor or its rows disagree (see investor_link_conflict). | [optional] |
| **investor_link_conflict** | **Boolean** | True when the manager&#39;s own holding rows resolve to different investors, so no investor is linked. | [optional] |
| **org_id** | **String** | Legal-entity level: the manager&#39;s organization record when all its rows agree on one (null otherwise). | [optional] |
| **company_id** | **String** | Company level: the manager&#39;s company (null when unresolved). | [optional] |
| **company_name** | **String** | Display name of the manager&#39;s company. | [optional] |
| **group_id** | **String** | Group level: the manager&#39;s company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] |
| **group_name** | **String** | Display name of the manager&#39;s group. | [optional] |
| **position_count** | **Integer** | Current positions the manager reports, one per issuer security. | [optional] |
| **issuer_count** | **Integer** | Distinct issuer securities (CUSIPs) across those positions. | [optional] |
| **valued_position_count** | **Integer** | Of those positions, the ones with a reported value (a position whose filing scale was ambiguous has none). | [optional] |
| **first_observed_date** | **Date** | The earliest filing date a current position was first observed. | [optional] |
| **last_filing_date** | **Date** | The latest filing date among the manager&#39;s positions. | [optional] |

## Example

```ruby
require 'pharmaobservatory'

instance = PharmaObservatory::ManagersRow.new(
  manager_id: null,
  name: null,
  name_variant_count: null,
  investor_id: null,
  investor_link_conflict: null,
  org_id: null,
  company_id: null,
  company_name: null,
  group_id: null,
  group_name: null,
  position_count: null,
  issuer_count: null,
  valued_position_count: null,
  first_observed_date: null,
  last_filing_date: null
)
```

