# \ItemsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItem**](ItemsApi.md#CreateItem) | **Post** /api/1/items | 品目の作成
[**DestroyItem**](ItemsApi.md#DestroyItem) | **Delete** /api/1/items/{id} | 品目の削除
[**GetItem**](ItemsApi.md#GetItem) | **Get** /api/1/items/{id} | 品目の取得
[**GetItems**](ItemsApi.md#GetItems) | **Get** /api/1/items | 品目一覧の取得
[**UpdateItem**](ItemsApi.md#UpdateItem) | **Put** /api/1/items/{id} | 品目の更新



## CreateItem

> ItemResponse CreateItem(ctx, optional)

品目の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所の品目を作成する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemParams** | [**optional.Interface of ItemParams**](ItemParams.md)| 品目の作成 | 

### Return type

[**ItemResponse**](itemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyItem

> DestroyItem(ctx, id, companyId)

品目の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の品目を削除する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 品目ID | 
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


## GetItem

> ItemResponse GetItem(ctx, companyId, id)

品目の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の品目を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**id** | **int32**| 品目ID | 

### Return type

[**ItemResponse**](itemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> InlineResponse2008 GetItems(ctx, companyId, optional)

品目一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の品目一覧を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItem

> ItemResponse UpdateItem(ctx, id, optional)

品目の更新

 <h2 id=\"\">概要</h2>  <p>指定した事業所の品目を更新する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 品目ID | 
 **optional** | ***UpdateItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemParams** | [**optional.Interface of ItemParams**](ItemParams.md)| 品目の更新 | 

### Return type

[**ItemResponse**](itemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

