# TrialPlThreeYearsResponseTrialPlThreeYearsBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**AccountGroupName** | **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**Partners** | [**[]TrialBsThreeYearsResponseTrialBsThreeYearsPartners**](trialBsThreeYearsResponse_trial_bs_three_years_partners.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Items** | [**[]TrialBsThreeYearsResponseTrialBsThreeYearsItems**](trialBsThreeYearsResponse_trial_bs_three_years_items.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Sections** | [**[]TrialPlThreeYearsResponseTrialPlThreeYearsSections**](trialPlThreeYearsResponse_trial_pl_three_years_sections.md) | breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**AccountCategoryName** | **string** | 勘定科目カテゴリー名 | [optional] 
**TotalLine** | **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 
**HierarchyLevel** | **int32** | 階層レベル | [optional] 
**ParentAccountCategoryName** | **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**TwoYearsBeforeClosingBalance** | **int32** | 前々年度期末残高 | [optional] 
**LastYearClosingBalance** | **int32** | 前年度期末残高 | [optional] 
**ClosingBalance** | **int32** | 期末残高 | [optional] 
**YearOnYear** | **float32** | 前年比 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


