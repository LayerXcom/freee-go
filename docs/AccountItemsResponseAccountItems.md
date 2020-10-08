# AccountItemsResponseAccountItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 勘定科目ID | 
**Name** | **string** | 勘定科目名 (30文字以内) | 
**Shortcut** | Pointer to **string** | ショートカット1 (20文字以内) | [optional] 
**ShortcutNum** | Pointer to **string** | ショートカット2(勘定科目コード) (20文字以内) | [optional] 
**DefaultTaxId** | **int32** | デフォルト設定がされている税区分ID | [optional] 
**DefaultTaxCode** | **int32** | デフォルト設定がされている税区分コード | 
**AccountCategory** | **string** | 勘定科目カテゴリー | 
**AccountCategoryId** | **int32** | 勘定科目のカテゴリーID | 
**Categories** | **[]string** |  | 
**Available** | **bool** | 勘定科目の使用設定（true: 使用する、false: 使用しない） | 
**WalletableId** | Pointer to **int32** | 口座ID | 
**GroupName** | Pointer to **string** | 決算書表示名（小カテゴリー） | [optional] 
**CorrespondingIncomeName** | Pointer to **string** | 収入取引相手勘定科目名 | [optional] 
**CorrespondingIncomeId** | Pointer to **int32** | 収入取引相手勘定科目ID | [optional] 
**CorrespondingExpenseName** | Pointer to **string** | 支出取引相手勘定科目名 | [optional] 
**CorrespondingExpenseId** | Pointer to **int32** | 支出取引相手勘定科目ID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


