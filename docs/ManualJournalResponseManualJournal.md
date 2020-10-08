# ManualJournalResponseManualJournal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 振替伝票ID | 
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**Adjustment** | **bool** | 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳） | 
**TxnNumber** | Pointer to **string** | 仕訳番号 | 
**Details** | [**[]ManualJournalResponseManualJournalDetails**](manualJournalResponse_manual_journal_details.md) | 貸借行一覧（配列）: 貸借合わせて100行まで登録できます。 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


