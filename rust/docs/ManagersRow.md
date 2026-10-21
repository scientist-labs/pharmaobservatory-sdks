# ManagersRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**manager_id** | Option<**String**> | The manager key, the id the web page uses: edgar_cik:<10-digit filer CIK>, or party:sec-cik:<ref> for a filer whose CIK is malformed. The same value as holder_identity_key on institutional_holdings. | [optional]
**name** | Option<**String**> | The manager's name as filed (the longest variant), falling back to the manager key. | [optional]
**name_variant_count** | Option<**i32**> | How many spellings of the name the filings carry. | [optional]
**investor_id** | Option<**String**> | The manager's investor key where the link resolves; null when the manager resolves to no investor or its rows disagree (see investor_link_conflict). | [optional]
**investor_link_conflict** | Option<**bool**> | True when the manager's own holding rows resolve to different investors, so no investor is linked. | [optional]
**org_id** | Option<**String**> | Legal-entity level: the manager's organization record when all its rows agree on one (null otherwise). | [optional]
**company_id** | Option<**String**> | Company level: the manager's company (null when unresolved). | [optional]
**company_name** | Option<**String**> | Display name of the manager's company. | [optional]
**group_id** | Option<**String**> | Group level: the manager's company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the manager's group. | [optional]
**position_count** | Option<**i32**> | Current positions the manager reports, one per issuer security. | [optional]
**issuer_count** | Option<**i32**> | Distinct issuer securities (CUSIPs) across those positions. | [optional]
**valued_position_count** | Option<**i32**> | Of those positions, the ones with a reported value (a position whose filing scale was ambiguous has none). | [optional]
**first_observed_date** | Option<[**String**](string.md)> | The earliest filing date a current position was first observed. | [optional]
**last_filing_date** | Option<[**String**](string.md)> | The latest filing date among the manager's positions. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


