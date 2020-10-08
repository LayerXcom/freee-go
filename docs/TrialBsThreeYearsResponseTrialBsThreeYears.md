# TrialBsThreeYearsResponseTrialBsThreeYears

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**FiscalYear** | **int32** | 会計年度(条件に指定した時、または条件に月、日条件がない時のみ含まれる） | [optional] 
**StartMonth** | **int32** | 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 
**EndMonth** | **int32** | 発生月で絞込：終了会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 
**StartDate** | **string** | 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**EndDate** | **string** | 発生日で絞込：終了日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**AccountItemDisplayType** | **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる） | [optional] 
**BreakdownDisplayType** | **string** | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item）(条件に指定した時のみ含まれる） | [optional] 
**PartnerId** | **int32** | 取引先ID(条件に指定した時のみ含まれる） | [optional] 
**PartnerCode** | **string** | 取引先コード(条件に指定した時のみ含まれる） | [optional] 
**ItemId** | **int32** | 品目ID(条件に指定した時のみ含まれる） | [optional] 
**Adjustment** | **string** | 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる） | [optional] 
**CreatedAt** | **string** | 作成日時 | [optional] 
**Balances** | [**[]TrialBsThreeYearsResponseTrialBsThreeYearsBalances**](trialBsThreeYearsResponse_trial_bs_three_years_balances.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


