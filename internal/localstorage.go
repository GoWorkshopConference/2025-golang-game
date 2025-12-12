//go:build !js || !wasm

package internal

// GetParentLocalStorage は親ページのlocalStorageから値を取得します
// 通常ビルド時（非WASM）では常にfalseを返します
func GetParentLocalStorage(key string) (string, bool) {
	return "", false
}

// GetAllParentLocalStorage は親ページのlocalStorageの全データを取得します
// 通常ビルド時（非WASM）では空のmapを返します
func GetAllParentLocalStorage() map[string]string {
	return make(map[string]string)
}
