# TrialPlResponseTrialPlBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**AccountGroupName** | **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**Partners** | [**[]TrialBsResponseTrialBsPartners**](trialBsResponse_trial_bs_partners.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Items** | [**[]TrialBsResponseTrialBsItems**](trialBsResponse_trial_bs_items.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Sections** | [**[]TrialPlResponseTrialPlSections**](trialPlResponse_trial_pl_sections.md) | breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**AccountCategoryName** | **string** | 勘定科目カテゴリー名 | [optional] 
**TotalLine** | **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 
**HierarchyLevel** | **int32** | 階層レベル | [optional] 
**ParentAccountCategoryName** | **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**OpeningBalance** | **int32** | 期首残高 | [optional] 
**DebitAmount** | **int32** | 借方金額 | [optional] 
**CreditAmount** | **int32** | 貸方金額 | [optional] 
**ClosingBalance** | **int32** | 期末残高 | [optional] 
**CompositionRatio** | **float32** | 構成比 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


