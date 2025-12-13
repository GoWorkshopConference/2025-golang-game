# golang-game
Go の Ebitengine で作るミニゲーム

## 開発手順

1. このリポジトリをクローンしてきてください
2. 素材となる画像をダウンロードしてきて、`internal/assets` 以下に配置してください
   - 展示したゲームではいらすとやさんから以下の画像を使用しました。
   - [エビフライのイラスト`ebi_fry_3_rich.png`](https://www.irasutoya.com/2012/03/blog-post_5621.html)
   - [エビフライのイラスト（3本）`ebi_fry_3.png`](https://www.irasutoya.com/2019/09/3.html)
   - [大きな口を開けている人のイラスト（男性）`big_mouse.png`](https://www.irasutoya.com/2015/07/blog-post_314.html)
   - [ソース・ウスターソースのイラスト`sauce.png`](https://www.irasutoya.com/2014/06/blog-post_2392.html)
   - [ウイルスのイラスト`virus.png`](https://www.irasutoya.com/2013/01/blog-post_4644.html)
   - [ウイルスに感染したEメールのイラスト`virus_computer.png`](https://www.irasutoya.com/2016/03/e_14.html)
3. `go mod tidy` で依存パッケージをインストールしてください
4. ローカルでプログラムを動かす場合は `make run` もしくは `go run cmd/game/main.go` でゲームを始められます
   - クリック（タップ）は動作しません
5. Web サイトにゲームをアップロードするときは、 `make build` もしくは `env GOOS=js GOARCH=wasm go build -o webpage/game.wasm github.com/GoWorkshopConference/golang-game/cmd/game` で wasm に変換し、Web サイト上で表示してください
