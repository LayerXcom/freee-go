# AccountItemParamsAccountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 勘定科目名 (30文字以内) | 
**Shortcut** | **string** | ショートカット1 (20文字以内) | [optional] 
**ShortcutNum** | **string** | ショートカット2(勘定科目コード)(20文字以内) | [optional] 
**TaxCode** | **int32** | 税区分コード | 
**GroupName** | **string** | 決算書表示名（小カテゴリー） Selectablesフォーム用選択項目情報エンドポイント(account_groups.name)で取得可能です | 
**AccountCategoryId** | **int32** | 勘定科目カテゴリーID Selectablesフォーム用選択項目情報エンドポイント(account_groups.account_category_id)で取得可能です | 
**CorrespondingIncomeId** | **int32** | 収入取引相手勘定科目ID | 
**CorrespondingExpenseId** | **int32** | 支出取引相手勘定科目ID | 
**AccumulatedDepAccountItemId** | **int32** | 減価償却累計額勘定科目ID（法人のみ利用可能） | [optional] 
**Searchable** | **int32** | 検索可能:2, 検索不可：3(登録時未指定の場合は2で登録されます。更新時未指定の場合はsearchableは変更されません。) | [optional] 
**Items** | [**[]AccountItemParamsAccountItemItems**](accountItemParams_account_item_items.md) | 品目 | [optional] 
**Partners** | [**[]AccountItemParamsAccountItemItems**](accountItemParams_account_item_items.md) | 取引先 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


