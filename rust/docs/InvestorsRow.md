# InvestorsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**investor_id** | Option<**String**> | The investor key, the id the web page uses: adv:<CRD> for a Form ADV adviser, lei:<LEI> where one is known, agency:<slug> for an NIH RePORTER funding agency, investor:<hash> for a name-resolved investor. | [optional]
**legal_name** | Option<**String**> | The investor's legal or display name. | [optional]
**investor_type** | Option<**String**> | venture_capital, private_equity, hedge_fund (from the plurality of the adviser's Form ADV Item 7.B fund types), government_agency or other. | [optional]
**source** | Option<**String**> | The source lane: sec_adv (Form ADV advisers) or reporter (NIH RePORTER funding agencies). | [optional]
**country** | Option<**String**> | The adviser's main-office country, or United States for an NIH funding agency. | [optional]
**raum_usd** | Option<**f64**> | Regulatory assets under management (RAUM) in whole US dollars, as the adviser reports it on Form ADV. Null for an agency or an adviser that reports none. | [optional]
**raum_band** | Option<**String**> | The RAUM band the amount falls in. | [optional]
**is_era** | Option<**bool**> | True for an exempt reporting adviser (ERA), false for a registered adviser, null for an agency. | [optional]
**is_life_sciences_focused** | Option<**bool**> | True when the investor is life-sciences focused, false when it is not, null when no signal exists. Read it with is_life_sciences_focused_basis. | [optional]
**is_life_sciences_focused_basis** | Option<**String**> | Where is_life_sciences_focused came from: human or agent (a recorded classification), override (a curator's call), edge_share (the share of its portfolio edges in life-sciences issuers) or classifier (the Form 13F biomedical value share). Null exactly when the flag is null. | [optional]
**private_fund_count** | Option<**i32**> | Private funds the adviser reports on Form ADV Item 7.B. | [optional]
**total_private_fund_gav** | Option<**f64**> | The adviser's total private-fund gross asset value (GAV) in whole US dollars, a firm-level figure. | [optional]
**portfolio_co_count** | Option<**i64**> | Distinct portfolio companies named against the investor in S-1 and 424B offering holder tables. 0 usually means no named-holder link was observed. | [optional]
**holdings_position_count** | Option<**i64**> | Current Form 13F positions reported by the managers that resolve to this investor, summed over its manager keys. 0 when none resolves. | [optional]
**org_id** | Option<**String**> | Legal-entity level: the investor's organization record where the link resolves (null otherwise). An advfirm:<CRD> shell resolves to no legal-entity record and no company. | [optional]
**company_id** | Option<**String**> | Company level: the investor's company (null when unresolved). | [optional]
**company_name** | Option<**String**> | Display name of the investor's company. | [optional]
**group_id** | Option<**String**> | Group level: the investor's company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the investor's group. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


