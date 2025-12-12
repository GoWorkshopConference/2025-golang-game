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
