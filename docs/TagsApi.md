# \TagsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagsApi.md#CreateTag) | **Post** /api/1/tags | メモタグの作成
[**DestroyTag**](TagsApi.md#DestroyTag) | **Delete** /api/1/tags/{id} | メモタグの削除
[**GetTag**](TagsApi.md#GetTag) | **Get** /api/1/tags/{id} | メモタグの詳細情報の取得
[**GetTags**](TagsApi.md#GetTags) | **Get** /api/1/tags | メモタグ一覧の取得
[**UpdateTag**](TagsApi.md#UpdateTag) | **Put** /api/1/tags/{id} | メモタグの更新



## CreateTag

> TagResponse CreateTag(ctx, tagParams)

メモタグの作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所のメモタグを作成する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagParams** | [**TagParams**](TagParams.md)| メモタグの作成 | 

### Return type

[**TagResponse**](tagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTag

> DestroyTag(ctx, id, companyId)

メモタグの削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所のメモタグを削除する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| タグID | 
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


## GetTag

> TagResponse GetTag(ctx, id, companyId)

メモタグの詳細情報の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所のメモタグを取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| タグID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**TagResponse**](tagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> InlineResponse200 GetTags(ctx, companyId, optional)

メモタグ一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所のメモタグ一覧を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> TagResponse UpdateTag(ctx, id, optional)

メモタグの更新

 <h2 id=\"\">概要</h2>  <p>指定した事業所のメモタグを更新する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| メモタグID | 
 **optional** | ***UpdateTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagParams** | [**optional.Interface of TagParams**](TagParams.md)| メモタグの更新 | 

### Return type

[**TagResponse**](tagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

