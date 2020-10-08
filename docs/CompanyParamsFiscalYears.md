# CompanyParamsFiscalYears

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseIndustryTemplate** | **bool** | 製造業向け機能（true: 使用する、false: 使用しない） | [optional] 
**IndirectWriteOffMethod** | **bool** | 固定資産の控除法（true\\: 間接控除法、false\\: 直接控除法）&lt;br&gt; 間接控除法は法人のみ対応しています。  | [optional] 
**IndirectWriteOffMethodType** | **bool** | 間接控除時の累計額(法人のみ)（true: 資産分類別、false: 共通）&#39;  | [optional] 
**StartDate** | **string** | 期首日 | [optional] 
**EndDate** | **string** | 期末日（決算日） | [optional] 
**AccountingPeriod** | **int32** | 期 | [optional] 
**DepreciationFraction** | **int32** | 減価償却端数処理法(法人のみ)(0: 切り捨て、1: 切り上げ) | [optional] 
**TaxFraction** | **int32** | 消費税端数処理方法（0: 切り上げ、1: 切り捨て, 2: 四捨五入） | [optional] 
**ReturnCode** | **int32** | 不動産所得使用区分（0: 一般、3: 一般/不動産） ※個人事業主のみ設定可能 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


