# \ApprovalRequestsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApprovalRequest**](ApprovalRequestsApi.md#CreateApprovalRequest) | **Post** /api/1/approval_requests | 各種申請の作成
[**DestroyApprovalRequest**](ApprovalRequestsApi.md#DestroyApprovalRequest) | **Delete** /api/1/approval_requests/{id} | 各種申請の削除
[**GetApprovalRequest**](ApprovalRequestsApi.md#GetApprovalRequest) | **Get** /api/1/approval_requests/{id} | 各種申請の取得
[**GetApprovalRequestForm**](ApprovalRequestsApi.md#GetApprovalRequestForm) | **Get** /api/1/approval_requests/forms/{id} | 各種申請の申請フォームの取得
[**GetApprovalRequestForms**](ApprovalRequestsApi.md#GetApprovalRequestForms) | **Get** /api/1/approval_requests/forms | 各種申請の申請フォーム一覧の取得
[**GetApprovalRequests**](ApprovalRequestsApi.md#GetApprovalRequests) | **Get** /api/1/approval_requests | 各種申請の一覧
[**UpdateApprovalRequest**](ApprovalRequestsApi.md#UpdateApprovalRequest) | **Put** /api/1/approval_requests/{id} | 各種申請の更新
[**UpdateApprovalRequestAction**](ApprovalRequestsApi.md#UpdateApprovalRequestAction) | **Post** /api/1/approval_requests/{id}/actions | 各種申請の承認操作



## CreateApprovalRequest

> ApprovalRequestResponse CreateApprovalRequest(ctx, optional)

各種申請の作成

 <h2 id=\"_1\">概要</h2>  <p>指定した事業所の各種申請を作成する</p>  <p>各種申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-approval-requests\" target=\"_blank\">会計freeeの各種申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIでは、各種申請を作成することができます。</li>   <li>     申請項目(request_items)については、申請フォームで設定された項目のIDとそれに対応する値を入力してください。     <ul>       <li>タイトル(title)：文字列(改行なし) 例)予算申請</li>       <li>1行コメント(single_line)：文字列(改行なし) 例)予算に関する申請</li>       <li>複数行コメント(multi_line)：文字列(改行あり)       <br>       &nbsp;&nbsp;例)<br>       &nbsp;&nbsp;予算に関する申請<br>       &nbsp;&nbsp;申請日 2020-01-01<br>       </li>       <li>プルダウン(select)： プルダウンの選択肢の名前(改行なし) 例)開発部</li>       <li>日付(date)： 日付形式 例)2020-01-01</li>       <li>金額(amount)： 数値(申請フォームで設定した上限・下限金額内の値, 改行なし) 例)10000</li>       <li>添付ファイル(receipt)： ファイルボックスAPIのID 例)1</li>     </ul>   </li>   <li>     申請ステータス(下書き、申請中)の指定と変更、及び承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）は以下を参考にして行ってください。     <ul>       <li>         承認操作は申請ステータスが申請中、承認済み、却下のものだけが対象です。         <ul>           <li>             初回申請の場合             <ul><li>申請の作成（POST）</li></ul>           </li>           <li>             作成済みの申請の申請ステータス変更・更新する場合             <ul><li>申請の更新（PUT）</li></ul>           </li>           <li>             申請中、承認済み、却下の申請の承認操作を行う場合             <ul><li>承認操作の実行（POST）</li></ul>           </li>         </ul>       </li>       <li>申請の削除（DELETE）が可能なのは申請ステータスが下書き、差戻しの場合のみです</li>     </ul>   </li>   <li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している各種申請は本API経由で作成ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>Web画面やAPIを通じて申請フォームが変更されると、本APIで使用する項目ID（request_itemsで指定するid）も変更されます。本API利用前に各種申請の取得APIを利用し、最新の申請フォームを取得することを推奨します。</li>   <li>本APIはエンタープライズプランをご利用の事業所のみ利用可能です。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateApprovalRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateApprovalRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **approvalRequestCreateParams** | [**optional.Interface of ApprovalRequestCreateParams**](ApprovalRequestCreateParams.md)| 各種申請の作成 | 

### Return type

[**ApprovalRequestResponse**](approvalRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyApprovalRequest

> DestroyApprovalRequest(ctx, id, companyId)

各種申請の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の各種申請を削除する</p>  <p>各種申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-approval-requests\" target=\"_blank\">会計freeeの各種申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>     申請ステータス(下書き、申請中)の指定と変更、及び承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）は以下を参考にして行ってください。     <ul>       <li>         承認操作は申請ステータスが申請中、承認済み、却下のものだけが対象です。         <ul>           <li>             初回申請の場合             <ul><li>申請の作成（POST）</li></ul>           </li>           <li>             作成済みの申請の申請ステータス変更・更新する場合             <ul><li>申請の更新（PUT）</li></ul>           </li>           <li>             申請中、承認済み、却下の申請の承認操作を行う場合             <ul><li>承認操作の実行（POST）</li></ul>           </li>         </ul>       </li>       <li>申請の削除（DELETE）が可能なのは申請ステータスが下書き、差戻しの場合のみです</li>     </ul>   </li>   <li>本APIはエンタープライズプランをご利用の事業所のみ利用可能です。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 各種申請ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovalRequest

> ApprovalRequestResponse GetApprovalRequest(ctx, id, companyId)

各種申請の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の各種申請を取得する</p>  <p>各種申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-approval-requests\" target=\"_blank\">会計freeeの各種申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している各種申請と申請経路はAPI経由で参照は可能ですが、作成と更新、承認ステータスの変更ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>本APIはエンタープライズプランをご利用の事業所のみ利用可能です。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 各種申請ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**ApprovalRequestResponse**](approvalRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovalRequestForm

> ApprovalRequestFormResponse GetApprovalRequestForm(ctx, id, companyId)

各種申請の申請フォームの取得

 <h2 id=\"_1\">概要</h2>  <p>指定した事業所の各種申請の申請フォームを取得する</p>  <p>各種申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-approval-requests\" target=\"_blank\">会計freeeの各種申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIはエンタープライズプランをご利用の事業所のみ利用可能です。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 申請フォームID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**ApprovalRequestFormResponse**](approvalRequestFormResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovalRequestForms

> InlineResponse20018 GetApprovalRequestForms(ctx, companyId)

各種申請の申請フォーム一覧の取得

 <h2 id=\"_1\">概要</h2>  <p>指定した事業所の各種申請の申請フォーム一覧を取得する</p>  <p>各種申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-approval-requests\" target=\"_blank\">会計freeeの各種申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIはエンタープライズプランをご利用の事業所のみ利用可能です。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovalRequests

> ApprovalRequestsIndexResponse GetApprovalRequests(ctx, companyId, optional)

各種申請の一覧

 <h2 id=\"_1\">概要</h2>  <p>指定した事業所の各種申請一覧を取得する</p>  <p>各種申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-approval-requests\" target=\"_blank\">会計freeeの各種申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIでは、各種申請の一覧を取得することができます。</li>   <li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している各種申請と申請経路はAPI経由で参照は可能ですが、作成と更新、承認ステータスの変更ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>本APIはエンタープライズプランをご利用の事業所のみ利用可能です。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetApprovalRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApprovalRequestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
 **applicationNumber** | **optional.Int32**| 申請No. | 
 **title** | **optional.String**| 申請タイトル | 
 **formId** | **optional.Int32**| 申請フォームID | 
 **startApplicationDate** | **optional.String**| 申請日で絞込：開始日(yyyy-mm-dd) | 
 **endApplicationDate** | **optional.String**| 申請日で絞込：終了日(yyyy-mm-dd) | 
 **applicantId** | **optional.Int32**| 申請者のユーザーID | 
 **approverId** | **optional.Int32**| 承認者のユーザーID | 
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 500) | 

### Return type

[**ApprovalRequestsIndexResponse**](approvalRequestsIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApprovalRequest

> ApprovalRequestResponse UpdateApprovalRequest(ctx, id, approvalRequestUpdateParams)

各種申請の更新

 <h2 id=\"_1\">概要</h2>  <p>指定した事業所の各種申請を更新する</p>  <p>各種申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-approval-requests\" target=\"_blank\">会計freeeの各種申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIでは、各種申請を更新することができます。</li>   <li>     申請項目(request_items)については、各種申請の取得APIで取得したrequest_items.idとそれに対応する値を入力してください。     <ul>       <li>タイトル(title)：文字列(改行なし) 例)予算申請</li>       <li>1行コメント(single_line)：文字列(改行なし) 例)予算に関する申請</li>       <li>複数行コメント(multi_line)：文字列(改行あり)       <br>       &nbsp;&nbsp;例)<br>       &nbsp;&nbsp;予算に関する申請<br>       &nbsp;&nbsp;申請日 2020-01-01<br>       </li>       <li>プルダウン(select)： プルダウンの選択肢の名前(改行なし) 例)開発部</li>       <li>日付(date)： 日付形式 例)2020-01-01</li>       <li>金額(amount)： 数値(申請フォームで設定した上限・下限金額内の値, 改行なし) 例)10000</li>       <li>添付ファイル(receipt)： ファイルボックスAPIのID 例)1</li>     </ul>   </li>   <li>本APIでは、status(申請ステータス): draft:下書き, feedback:差戻しのみ更新可能です。</li>   <li>     申請ステータス(下書き、申請中)の指定と変更、及び承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）は以下を参考にして行ってください。     <ul>       <li>         承認操作は申請ステータスが申請中、承認済み、却下のものだけが対象です。         <ul>           <li>             初回申請の場合             <ul><li>申請の作成（POST）</li></ul>           </li>           <li>             作成済みの申請の申請ステータス変更・更新する場合             <ul><li>申請の更新（PUT）</li></ul>           </li>           <li>             申請中、承認済み、却下の申請の承認操作を行う場合             <ul><li>承認操作の実行（POST）</li></ul>           </li>         </ul>       </li>       <li>申請の削除（DELETE）が可能なのは申請ステータスが下書き、差戻しの場合のみです</li>     </ul>   </li>   <li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している各種申請は本API経由で更新ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>Web画面やAPIを通じて申請フォームが変更されると、本APIで使用する項目ID（request_itemsで指定するid）も変更されます。本APIで使用する項目IDは申請作成時点の項目IDです。本API利用前に各種申請の取得APIを利用し、申請作成時点の項目IDを取得することを推奨します。</li>   <li>本APIはエンタープライズプランをご利用の事業所のみ利用可能です。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 各種申請ID | 
**approvalRequestUpdateParams** | [**ApprovalRequestUpdateParams**](ApprovalRequestUpdateParams.md)| 各種申請の更新 | 

### Return type

[**ApprovalRequestResponse**](approvalRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApprovalRequestAction

> ApprovalRequestResponse UpdateApprovalRequestAction(ctx, id, approvalRequestActionCreateParams)

各種申請の承認操作

 <h2 id=\"_1\">概要</h2>  <p>指定した事業所の各種申請の承認操作を行う</p>  <p>各種申請APIの使い方については、<a href=\"https://developer.freee.co.jp/tips/accounting-approval-requests\" target=\"_blank\">会計freeeの各種申請APIの使い方</a>をご参照ください</p>  <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIでは、各種申請の承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）を行うことができます。</li>   <li>     申請ステータス(下書き、申請中)の指定と変更、及び承認操作（承認する、却下する、申請者へ差し戻す、代理承認する、承認済み・却下済みを取り消す）は以下を参考にして行ってください。     <ul>       <li>         承認操作は申請ステータスが申請中、承認済み、却下のものだけが対象です。         <ul>           <li>             初回申請の場合             <ul><li>申請の作成（POST）</li></ul>           </li>           <li>             作成済みの申請の申請ステータス変更・更新する場合             <ul><li>申請の更新（PUT）</li></ul>           </li>           <li>             申請中、承認済み、却下の申請の承認操作を行う場合             <ul><li>承認操作の実行（POST）</li></ul>           </li>         </ul>       </li>       <li>申請の削除（DELETE）が可能なのは申請ステータスが下書き、差戻しの場合のみです</li>     </ul>   </li> 　<li>     申請経路、承認者の指定として部門役職データ連携を活用し、以下のいずれかを利用している各種申請はAPI経由で承認ステータスの変更ができません。     <ul>       <li>役職指定（申請者の所属部門）</li>       <li>役職指定（申請時に部門指定）</li>       <li>部門および役職指定</li>     </ul>   </li>   <li>本APIはエンタープライズプランをご利用の事業所のみ利用可能です。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 各種申請ID | 
**approvalRequestActionCreateParams** | [**ApprovalRequestActionCreateParams**](ApprovalRequestActionCreateParams.md)| 各種申請の承認操作 | 

### Return type

[**ApprovalRequestResponse**](approvalRequestResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

