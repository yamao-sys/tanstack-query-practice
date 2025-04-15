# tanstack-practice

Tanstack の検証

## コマンド類

### フロントエンド

OpenAPI から型ファイル生成

```
pnpx tsx generate_api_types.ts <openapi-file-path> <output-path>
```

## 参考
- https://zenn.dev/taisei_13046/books/133e9995b6aadf/viewer/2ce93a

## 使ってみた所管
- 設定は簡単
- featureパターンが良さそう
  - queryKeyを一元管理 & それをもとにしたuseQuery, useMutationのカスタムフックも作れるように
- App Routerではあまりハマらなさそう...
  - キャッシュはあくまでクライアントのものだから、SPAでなければ非同期の状態管理以外の旨みが薄れそう
  - ましてやApp Routerは非同期の管理はSuspenceも手段として取れるので
