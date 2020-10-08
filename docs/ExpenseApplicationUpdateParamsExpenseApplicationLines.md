# ExpenseApplicationUpdateParamsExpenseApplicationLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 経費申請の項目行ID: 既存項目行を更新する場合に指定します。IDを指定しない項目行は、新規行として扱われ追加されます。また、expense_application_linesに含まれない既存の項目行は削除されます。更新後も残したい行は、必ず経費申請の項目行IDを指定してexpense_application_linesに含めてください。 | [optional] 
**TransactionDate** | **string** | 日付 (yyyy-mm-dd) | [optional] 
**Description** | **string** | 内容 (250文字以内) | [optional] 
**Amount** | **int32** | 金額 | [optional] 
**ExpenseApplicationLineTemplateId** | **int32** | 経費科目ID | [optional] 
**ReceiptId** | **int32** | 証憑ID | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


