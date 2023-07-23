# 開発

```sh
docker compose build --no-cache
```

```sh
docker compose up -d
```

```sh
docker compose down
```

# フロントエンド

SvelteによるTODOアプリ（なぜかSvelteKitではじめてしまったがとくに意味はない）

## 機能

**MUST**

- TODOの表示
- TODOの追加
- TODOの削除

**WANT**

- TODOの編集
- ドラッグ&ドロップで並べ替え
- ログイン機能
    - ユーザー毎のTODO管理