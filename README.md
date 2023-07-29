# TODOアプリ（ Svelte + Go + Docker ）

<img src="./resource/banner.png">

学習を目的としたSvelteとGoによるTODOアプリです。Svelteはフロントエンド処理を担当させ、GoはTODOを管理するAPIを提供しています。双方のコンテナを作成しDockerのコンテナ間通信を行わせる。

## Svelte

フロントエンドフレームワークとして、Svelteを触ってみたかったので選択。かなり直感的に書けるのとReactよりVueに近いみたい（Vueを触ったことがない）なので、Reactよりとっつきやすいと思った。

ただ、大規模なアプリケーションになると、1つのファイルが縦に長くなってしまって辛くなりそうな感じ。普段Reactを触っているせいでそう感じるのかも。Next触ったことない上、ReactはReact Nativeばかり触っているのでWeb API系に対する型の付け方が全然わからなかった。（FormDataとかFetch APIとか）

あと、やっぱり新しいフレームワークを触るのはおもしろい。しかし、SvelteKit使う必要はなかった。

### 事前学習

本格的に触り始める前に結構事前にさわったこと

- [Svelteチュートリアル](https://svelte.dev/tutorial/basics)
- [実践Svelte入門](https://gihyo.jp/book/2023/978-4-297-13495-2)

実践Svelte入門はとくにオススメ

## Go

Goは今回の製作で一番苦労した。なぜなら、これまでフロントエンド系の言語・環境しか触ってこなかったので知識ゼロからのスタートだった。

**難しかったポイント**

- Go言語自体の理解
- Go言語のパッケージ管理
- httpリクエスト関連
- サーバーサイドでのデータの取り回し
- Go言語の設計

業務でJavaでも触っていれば苦労しないんだろうけど、どれもこれもはじめてで久しぶりにプログラミングの難しさを感じ、いい意味で焦りを感じることができた。

本当はタスク管理もDBサーバーもたてて、mySQLとかしたかったけど今回はそこまでいけなかった。個人的にはDBサーバーも大きな壁のひとつだ。今回のTODO管理はインメモリで再起動したら消える仕様で妥協した。クライアントサイドに持たせるのは何度もやってきたのでもういいかなと思った。

クリーンアーキテクチャの本とか読んでも素晴らしさが理解できなかったけど、今回の実装をおこなった上で再読すれば少しは理解できるのかもしれないと思った。

### 事前学習

- [A Tour Of Go](https://go-tour-jp.appspot.com/welcome/1)
- [初めてのGo言語](https://www.oreilly.co.jp/books/9784814400041/)
- [詳解Ｇｏ言語Ｗｅｂアプリケーション開発](https://www.kinokuniya.co.jp/f/dsg-01-9784863543720)

正直Goは完全初心者向けのドキュメントがかなり少ないから優良なものが少ない。今の自分のレベル（サーバーサイド何も知らない）にマッチしている参考書を見つけることができなかった。

## Docker

Dockerは今回かなり理解できたように思う。完全には理解できていないけど、初歩的なところを体感しながら理解できたのは良かった。

Docker imageに必要なもの詰め込んで、Docker Composeで動かす程度の理解だが・・。開発するには環境構築が必要なのは変わりないけど、プロダクトの環境をそのままDockerに乗っけることができたので、今回はまんぞく。

Dockerとデプロイ周りのことがさっぱりなので、今後触れる機会があれば触ってみたい。サーバーサイドコードを本番環境にデプロイする日は来るのだろうか。

### 事前学習

とくになし。公式ドキュメントが結構充実している。

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

# バックエンド

Goで実装されたインメモリTODO管理APIサーバー

|エンドポイント|機能|IF|
|:---:|:---:|:---:|
|/task/list|タスクのリストを取得|-|
|/task/add|タスクのリストを取得|`{"title":"タスク名"}`|
|/task/remove|タスクのリストを取得|`{"id": 0}`|

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