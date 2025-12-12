//go:build js && wasm

package internal

import (
	"syscall/js"
)

// GetParentLocalStorage は親ページのlocalStorageから値を取得します
// iframe内のWASMから親ページのlocalStorageにアクセスするために使用します
func GetParentLocalStorage(key string) (string, bool) {
	// JavaScriptのグローバル変数 window.parentLocalStorage にアクセス
	parentLocalStorage := js.Global().Get("parentLocalStorage")

	// parentLocalStorageが存在しない、またはundefinedの場合はfalseを返す
	if parentLocalStorage.IsUndefined() || parentLocalStorage.IsNull() {
		return "", false
	}

	// キーに対応する値を取得
	value := parentLocalStorage.Get(key)

	// 値が存在しない場合はfalseを返す
	if value.IsUndefined() || value.IsNull() {
		return "", false
	}

	// 文字列として返す
	return value.String(), true
}

// GetAllParentLocalStorage は親ページのlocalStorageの全データを取得します
func GetAllParentLocalStorage() map[string]string {
	result := make(map[string]string)

	// JavaScriptのグローバル変数 window.parentLocalStorage にアクセス
	parentLocalStorage := js.Global().Get("parentLocalStorage")

	// parentLocalStorageが存在しない、またはundefinedの場合は空のmapを返す
	if parentLocalStorage.IsUndefined() || parentLocalStorage.IsNull() {
		return result
	}

	// Object.keys()を使って全てのキーを取得
	keys := js.Global().Get("Object").Call("keys", parentLocalStorage)
	length := keys.Length()

	// 各キーと値を取得
	for i := 0; i < length; i++ {
		key := keys.Index(i).String()
		value := parentLocalStorage.Get(key)
		if !value.IsUndefined() && !value.IsNull() {
			result[key] = value.String()
		}
	}

	return result
}

// GetServerUrl はサーバーのベースURLを取得します
func GetServerUrl() string {
	serverUrl := js.Global().Get("serverUrl")
	if serverUrl.IsUndefined() || serverUrl.IsNull() {
		return "http://localhost:8080"
	}
	return serverUrl.String()
}

// SendScoreToServer はスコアをサーバーに送信します（非同期）
func SendScoreToServer(username string, score int, lifeTime int) error {
	serverUrl := GetServerUrl()
	url := serverUrl + "/score"

	// リクエストボディを作成
	requestBodyObj := js.Global().Get("Object").New()
	requestBodyObj.Set("username", username)
	requestBodyObj.Set("score", score)
	requestBodyObj.Set("life_time", lifeTime)

	// JSONに変換
	json := js.Global().Get("JSON")
	jsonBody := json.Call("stringify", requestBodyObj)

	// fetch APIを使用してPOSTリクエストを送信
	fetch := js.Global().Get("fetch")
	if fetch.IsUndefined() {
		return nil // エラーを無視（非同期処理のため）
	}

	// リクエストオプションを作成
	options := js.Global().Get("Object").New()
	options.Set("method", "POST")
	headers := js.Global().Get("Object").New()
	headers.Set("Content-Type", "application/json")
	options.Set("headers", headers)
	options.Set("body", jsonBody)

	// Promiseを返すfetchを呼び出し（非同期で実行）
	promise := fetch.Invoke(url, options)

	// エラーハンドリング（非同期のため、エラーはログに出力するだけ）
	promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := args[0]
		ok := response.Get("ok")
		if !ok.IsUndefined() && !ok.Bool() {
			js.Global().Get("console").Call("error", "Failed to send score to server")
		}
		return nil
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("console").Call("error", "Error sending score:", args[0])
		return nil
	}))

	return nil // 非同期処理のため、常にnilを返す
}
