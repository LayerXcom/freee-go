# DealResponseDealReceipts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 証憑ID | 
**Status** | **string** | ステータス(unconfirmed:確認待ち、confirmed:確認済み、deleted:削除済み、ignored:無視) | 
**Description** | **string** | メモ | [optional] 
**MimeType** | **string** | MIMEタイプ | 
**IssueDate** | **string** | 発生日 | [optional] 
**Origin** | **string** | アップロード元種別 | 
**CreatedAt** | **string** | 作成日時（ISO8601形式） | 
**FileSrc** | **string** | ファイルのダウンロードURL（freeeにログインした状態でのみ閲覧可能です。） | 
**User** | [**DealResponseDealUser**](dealResponse_deal_user.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


