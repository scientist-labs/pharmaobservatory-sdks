# InstitutionalQuarterTotalsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**quarter_total_id** | Option<**String**> | The issuer and report period id. | [optional]
**issuer_org_id** | Option<**String**> | Legal-entity level: the issuer organization as the ownership identity seam resolves it, a resolution root; a listed subsidiary's securities may be keyed at its resolved parent. | [optional]
**issuer_company_id** | Option<**String**> | Company level: the issuer's company (null when the organization does not resolve to a company). | [optional]
**issuer_company_name** | Option<**String**> | Display name of the issuer's company. | [optional]
**issuer_group_id** | Option<**String**> | Group level: the issuer's current ultimate parent company; a root company's group is itself. | [optional]
**issuer_group_name** | Option<**String**> | Display name of the issuer's group. | [optional]
**report_period** | Option<[**String**](string.md)> | The Form 13F report period (quarter end) the total is measured at. | [optional]
**is_reference_period** | Option<**bool**> | True for the one quarter the ownership percentage reads: the most recent in which at least 80 percent of the prior quarter's filing managers had filed. | [optional]
**manager_count** | Option<**i64**> | Managers reporting a counted equity position in the issuer at the period. | [optional]
**position_count** | Option<**i64**> | Counted equity positions reported against the issuer at the period. | [optional]
**lane_position_count** | Option<**i64**> | Every position reported against the issuer at the period, counted or excluded. | [optional]
**shares_reported** | Option<**f64**> | Shares reported on Form 13F for the issuer at the period: counted equity positions only, summed within this issuer and period. | [optional]
**shares_otr** | Option<**f64**> | Of shares_reported, the shares under other (OTR) discretion. | [optional]
**shares_excluded_non_equity_issue** | Option<**f64**> | Shares reported against convertible notes, debt and preferred issues, excluded from shares_reported. | [optional]
**excluded_non_equity_issue_position_count** | Option<**i64**> | Positions excluded as non-equity issues. | [optional]
**excluded_derivative_position_count** | Option<**i64**> | Positions excluded because they are puts or calls. | [optional]
**excluded_principal_position_count** | Option<**i64**> | Positions excluded because they are reported as a principal amount rather than shares. | [optional]
**excluded_non_equity_cusip_count** | Option<**i64**> | Distinct CUSIPs excluded as non-equity issues. | [optional]
**equity_cusip_count** | Option<**i64**> | Distinct equity CUSIPs counted at the period. | [optional]
**latest_filing_date** | Option<[**String**](string.md)> | The latest Form 13F filing date among the period's positions. | [optional]
**aggregation_rule** | Option<**String**> | The rule the total was built under; always sum_within_issuer_period_only. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


