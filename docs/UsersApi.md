# \UsersApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /api/1/users | 事業所に所属するユーザー一覧の取得
[**GetUsersCapabilities**](UsersApi.md#GetUsersCapabilities) | **Get** /api/1/users/capabilities | ログインユーザーの権限の取得
[**GetUsersMe**](UsersApi.md#GetUsersMe) | **Get** /api/1/users/me | ログインユーザー情報の取得
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /api/1/users/me | ユーザー情報の更新



## GetUsers

> InlineResponse2006 GetUsers(ctx, companyId, optional)

事業所に所属するユーザー一覧の取得

 <h2 id=\"\">概要</h2>  <p>事業所に所属するユーザーの一覧を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersCapabilities

> InlineResponse2007 GetUsersCapabilities(ctx, companyId)

ログインユーザーの権限の取得

 <h2 id=\"\">概要</h2>  <p>ユーザーの権限情報を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersMe

> MeResponse GetUsersMe(ctx, optional)

ログインユーザー情報の取得

 <h2 id=\"\">概要</h2>  <p>ユーザーの情報を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUsersMeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersMeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companies** | **optional.Bool**| 取得情報にユーザーが所属する事業所一覧を含める | 

### Return type

[**MeResponse**](meResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserResponse UpdateUser(ctx, optional)

ユーザー情報の更新

 <h2 id=\"\">概要</h2>  <p>ユーザー情報を更新する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userParams** | [**optional.Interface of UserParams**](UserParams.md)| ユーザー情報の更新 | 

### Return type

[**UserResponse**](userResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

