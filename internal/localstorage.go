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

// GetServerUrl はサーバーのベースURLを取得します
// 通常ビルド時（非WASM）ではデフォルト値を返します
func GetServerUrl() string {
	return "http://localhost:8080"
}

// SendScoreToServer はスコアをサーバーに送信します
// 通常ビルド時（非WASM）では何もしません
func SendScoreToServer(username string, score int, lifeTime int) error {
	return nil
}
