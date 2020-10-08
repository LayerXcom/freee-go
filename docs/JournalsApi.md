# \JournalsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadJournal**](JournalsApi.md#DownloadJournal) | **Get** /api/1/journals/reports/{id}/download | ダウンロード実行
[**GetJournalStatus**](JournalsApi.md#GetJournalStatus) | **Get** /api/1/journals/reports/{id}/status | ステータス確認
[**GetJournals**](JournalsApi.md#GetJournals) | **Get** /api/1/journals | ダウンロード要求



## DownloadJournal

> string DownloadJournal(ctx, id, companyId)

ダウンロード実行

 <h2 id=\"\">概要</h2>  <p>ダウンロードを実行する</p>  <p>＊このAPIは無料プランのアカウントではご利用になれません</p>  <h2 id=\"_2\">定義</h2>  <ul> <li>id : 受け付けID</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 受け付けID | 
**companyId** | **int32**| 事業所ID | 

### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournalStatus

> JournalStatusResponse GetJournalStatus(ctx, companyId, id, optional)

ステータス確認

 <h2 id=\"\">概要</h2>  <p>ダウンロードリクエストのステータスを確認する</p>  <p>＊このAPIは無料プランのアカウントではご利用になれません</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>status</p>  <ul> <li>enqueued : 実行待ち</li>  <li>working : 実行中</li>  <li>uploaded : 準備完了</li> </ul> </li>  <li> <p>id : 受け付けID</p> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**id** | **int32**| 受け付けID | 
 **optional** | ***GetJournalStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetJournalStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **visibleTags** | [**optional.Interface of []string**](string.md)| 補助科目やコメントとして出力する項目 | 
 **startDate** | **optional.String**| 取得開始日 (yyyy-mm-dd) | 
 **endDate** | **optional.String**| 取得終了日 (yyyy-mm-dd) | 

### Return type

[**JournalStatusResponse**](journalStatusResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournals

> JournalsResponse GetJournals(ctx, downloadType, companyId, optional)

ダウンロード要求

 <h2 id=\"\">概要</h2>  <p>ユーザーが所属する事業所の仕訳帳のダウンロードをリクエストします 生成されるファイルに関しては、<a href=\"https://support.freee.co.jp/hc/ja/articles/204599604#2\">ヘルプページ</a>をご参照ください</p>  <p>＊このAPIは無料プランのアカウントではご利用になれません</p>  <h2 id=\"_2\">定義</h2>  <ul>   <li>download_type     <ul>       <li>generic(freee Webからダウンロードできるものと同じ)</li>       <li>csv (yayoi形式)</li>       <li>pdf</li>     </ul>   </li>   <li>visible_tags : 指定しない場合は従来の仕様の仕訳帳が出力されます     <ul>       <li>partner : 取引先タグ</li>       <li>item : 品目タグ</li>       <li>tag : メモタグ</li>       <li>section : 部門タグ</li>       <li>description : 備考欄</li>       <li>wallet_txn_description : 明細の備考欄</li>       <li>         segment_1_tag : セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン)<br>         segment_2_tag : セグメント2(法人向け エンタープライズプラン)<br>         segment_3_tag : セグメント3(法人向け エンタープライズプラン)<br><br>         <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br>       </li>       <li>all : 指定された場合は上記の設定をすべて有効として扱いますが、セグメント1、セグメント2、セグメント3は含みません。セグメントが必要な場合はallではなく、segment_1_tag, segment_2_tag, segment_3_tagを指定してください。</li>     </ul>   </li>    <li>id : 受け付けID</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downloadType** | **string**| ダウンロード形式 | 
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetJournalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetJournalsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **visibleTags** | [**optional.Interface of []string**](string.md)| 補助科目やコメントとして出力する項目 | 
 **startDate** | **optional.String**| 取得開始日 (yyyy-mm-dd) | 
 **endDate** | **optional.String**| 取得終了日 (yyyy-mm-dd) | 

### Return type

[**JournalsResponse**](journalsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

