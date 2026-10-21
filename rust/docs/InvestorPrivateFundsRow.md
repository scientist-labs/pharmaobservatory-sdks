# InvestorPrivateFundsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fund_id** | Option<**String**> | The fund key, advfund:<PF-ID>. | [optional]
**pf_id** | Option<**String**> | The SEC private fund identification number (PF-ID). | [optional]
**fund_name** | Option<**String**> | The fund's name as the adviser filed it. | [optional]
**fund_type** | Option<**String**> | The fund type the adviser reports on Form ADV Schedule D. | [optional]
**gross_asset_value** | Option<**f64**> | The fund's gross asset value in whole US dollars, as filed. | [optional]
**beneficial_owner_count** | Option<**i32**> | The number of beneficial owners the fund reports. It counts the fund's own investors, not investors in this directory. | [optional]
**is_master** | Option<**bool**> | True when the fund is a master fund. | [optional]
**is_feeder** | Option<**bool**> | True when the fund is a feeder fund. | [optional]
**master_fund_name** | Option<**String**> | The master fund's name when this fund is a feeder. | [optional]
**master_pf_id** | Option<**String**> | The master fund's PF-ID when this fund is a feeder. | [optional]
**adviser_count** | Option<**i64**> | Advisers that report this PF-ID. | [optional]
**investor_id** | Option<**String**> | The investor key of the adviser the fund is listed under. | [optional]
**as_of_date** | Option<[**String**](string.md)> | The date of the filing the fund row comes from. | [optional]
**source_url** | Option<**String**> | A link to the adviser's Form ADV filing. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


