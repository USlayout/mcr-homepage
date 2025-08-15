# Docker使い方ガイド - MCR部ホームページ開発

## 目次
1. [Dockerとは](#dockerとは)
2. [インストール](#インストール)
3. [基本コマンド](#基本コマンド)
4. [MCR部ホームページ開発環境構築](#mcr部ホームページ開発環境構築)
5. [よくあるトラブル](#よくあるトラブル)

## Dockerとは

Dockerは、アプリケーションを**コンテナ**という軽量な仮想環境で実行するツールです。

### イメージとコンテナの関係

**重要な概念**: Dockerは「設計図（イメージ）」から「実際の家（コンテナ）」を作る仕組みです。

```
📋 Dockerイメージ (設計図)
   ↓ docker run
🏠 Dockerコンテナ (実際に動く環境)
```

#### Dockerイメージとは
- **設計図のようなもの**: アプリケーションの実行に必要な全てが含まれたテンプレート
- **内容**: OS、ライブラリ、アプリケーション、設定ファイルなど
- **特徴**: 読み取り専用、変更されない
- **例**: `node:18` = Node.js 18がインストールされたLinux環境の設計図

#### Dockerコンテナとは
- **実際に動く環境**: イメージから作られた実行可能なインスタンス
- **特徴**: 実際にプロセスが動く、ファイルの読み書きができる
- **例**: `node:18`イメージから作ったコンテナでNode.jsアプリケーションが動く

#### 具体例で理解
```bash
# 1. イメージをダウンロード（設計図を入手）
docker pull node:18

# 2. イメージからコンテナを作成＆実行（設計図から家を建てて住む）
docker run -it node:18 bash

# 3. 同じイメージから複数のコンテナを作ることも可能
docker run -d node:18 node app.js  # アプリケーション用
docker run -it node:18 bash        # 開発用シェル
```

#### なぜイメージをダウンロードするのか？
- **再利用性**: 一度ダウンロードすれば、何度でもコンテナを作れる
- **効率性**: 必要な環境が全て含まれているので、個別にNode.jsをインストールする必要がない
- **一貫性**: どのマシンでも同じ環境のコンテナが作れる

### メリット
- **環境の統一**: チーム全員が同じ環境で開発できる
- **依存関係の管理**: Node.js、Go、データベースなどを簡単にセットアップ
- **本番環境との一致**: 開発環境と本番環境を同じにできる

## インストール

### Windows
1. [Docker Desktop for Windows](https://www.docker.com/products/docker-desktop/)をダウンロード
2. インストーラーを実行
3. WSL 2バックエンドを有効にする
4. 再起動後、Docker Desktopを起動

### 動作確認
```bash
docker --version
docker run hello-world
```

## 基本コマンド

### イメージ操作（設計図の管理）
```bash
# イメージ一覧表示（ダウンロード済みの設計図を確認）
docker images

# イメージをダウンロード（Docker Hubから設計図を入手）
docker pull node:18

# 具体例: Node.js 18の環境設計図をダウンロード
# これだけではまだコンテナは作られない！
docker pull node:18

# イメージを削除（不要な設計図を削除）
docker rmi <image_id>
```

### コンテナ操作（実際の環境の管理）
```bash
# コンテナ実行（設計図からコンテナを作成＆起動）
docker run -it node:18 bash
# ↑ node:18イメージから新しいコンテナを作り、bashシェルで操作

# バックグラウンドでコンテナ実行
docker run -d node:18 node app.js
# ↑ node:18イメージからコンテナを作り、app.jsを実行

# 実行中のコンテナ一覧（今動いている環境を確認）
docker ps

# 全てのコンテナ一覧（停止中も含む）
docker ps -a

# コンテナ停止（環境を一時停止）
docker stop <container_id>

# コンテナ削除（環境を完全削除）
docker rm <container_id>

# 停止したコンテナを再開
docker start <container_id>
```

### 実践例: Node.jsアプリケーションを動かす

#### Linux/Mac (bash)
```bash
# 1. Node.js環境の設計図をダウンロード
docker pull node:18

# 2. 現在のディレクトリをコンテナにマウントして開発環境作成
docker run -it -v $(pwd):/app -w /app node:18 bash

# 3. コンテナ内でNode.jsアプリケーション実行
# （コンテナの中で）npm install
# （コンテナの中で）npm start
```

#### Windows (PowerShell) 
```powershell
# 1. Node.js環境の設計図をダウンロード
docker pull node:18

# 2. 現在のディレクトリをコンテナにマウントして開発環境作成（PowerShell用）
docker run -it -v "${PWD}:/app" -w /app node:18 bash
# または
docker run -it -v "$(Get-Location):/app" -w /app node:18 bash

# 3. コンテナ内でNode.jsアプリケーション実行
# （コンテナの中で）npm install
# （コンテナの中で）npm start
```

#### 重要な違い
- **Linux/Mac**: `$(pwd)` = 現在のディレクトリ
- **PowerShell**: `${PWD}` または `$(Get-Location)` = 現在のディレクトリ
- **注意**: PowerShellでは`$(pwd)`は動作しません！

### ログとデバッグ
```bash
# コンテナのログ確認
docker logs <container_id>

# コンテナに入る
docker exec -it <container_id> bash
```

## MCR部ホームページ開発環境構築

### 1. Docker Compose設定

**重要**: `docker-compose.yml`は**コンテナの中ではなく、Windowsのローカル環境**（プロジェクトルート）に作成します！

```
📁 C:\Users\masan\Documents\mcr-homepage\  ← ここに作成！
├── docker-compose.yml                      ← このファイルを作る
├── frontend/
├── backend/
└── database/
```

プロジェクトルートに`docker-compose.yml`を作成：

```yaml
version: '3.8'

services:
  # フロントエンド（Node.js + TypeScript）
  frontend:
    image: node:18
    working_dir: /app
    volumes:
      - ./frontend:/app
      - /app/node_modules
    ports:
      - "3000:3000"
    command: npm run dev
    environment:
      - NODE_ENV=development

  # バックエンド（Go）
  backend:
    image: golang:1.20
    working_dir: /app
    volumes:
      - ./backend:/app
    ports:
      - "8080:8080"
    command: go run main.go
    environment:
      - GO_ENV=development
    depends_on:
      - database

  # データベース（MySQL）
  database:
    image: mysql:8
    environment:
      MYSQL_DATABASE: mcr_homepage
      MYSQL_USER: mcr_user
      MYSQL_PASSWORD: mcr_password
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  # WordPress（CMS検討中）
  wordpress:
    image: wordpress:latest
    ports:
      - "8000:80"
    environment:
      WORDPRESS_DB_HOST: database
      WORDPRESS_DB_USER: mcr_user
      WORDPRESS_DB_PASSWORD: mcr_password
      WORDPRESS_DB_NAME: mcr_wordpress
    depends_on:
      - database

volumes:
  postgres_data:
```

### 2. 開発環境起動手順

**操作場所**: 全てWindowsのPowerShell（コンテナの外）で実行

```powershell
# 1. プロジェクトディレクトリに移動（Windows環境）
cd C:\Users\masan\Documents\mcr-homepage

# 2. docker-compose.ymlがあることを確認
ls docker-compose.yml

# 3. 全サービス起動（複数のコンテナを一度に起動）
docker-compose up -d

# 4. ログ確認
docker-compose logs -f

# 5. 特定サービスのログ確認
docker-compose logs frontend
```

#### Docker Composeの仕組み
```
Windows環境（あなたのPC）
├── docker-compose.yml          ← 設定ファイル（ここに作る）
├── frontend/                   ← ソースコード
├── backend/                    ← ソースコード
└── Docker Engine
    ├── 🐳 frontendコンテナ    ← Node.js環境
    ├── 🐳 backendコンテナ     ← Go環境  
    └── 🐳 databaseコンテナ    ← PostgreSQL環境
```

### 3. 各サービスへのアクセス

- **フロントエンド**: http://localhost:3000
- **バックエンドAPI**: http://localhost:8080
- **WordPress**: http://localhost:8000
- **データベース**: localhost:5432

### 4. 開発時のよく使うコマンド

```bash
# サービス再起動
docker-compose restart frontend

# 特定サービスでコマンド実行
docker-compose exec frontend npm install
docker-compose exec backend go mod tidy

# データベースに接続
docker-compose exec database mysql -U mcr_user -d mcr_homepage

# 全サービス停止・削除
docker-compose down

# ボリュームも含めて完全削除
docker-compose down -v
```

## プロジェクト固有の設定

### フロントエンド用Dockerfile
```dockerfile
# frontend/Dockerfile
FROM node:18-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]
```

### バックエンド用Dockerfile
```dockerfile
# backend/Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
```

## よくあるトラブル

### 0. PowerShellでのパス指定エラー
```
エラー: docker: invalid reference format
```

**原因**: PowerShellでLinux用の`$(pwd)`コマンドを使用

**解決方法**:
```powershell
# ❌ 間違い（Linux/Mac用）
docker run -it -v $(pwd):/app -w /app node:18 bash

# ✅ 正解（PowerShell用）
docker run -it -v "${PWD}:/app" -w /app node:18 bash
# または
docker run -it -v "$(Get-Location):/app" -w /app node:18 bash
```

### 0.1. version警告とボリューム未定義エラー
```
エラー1: the attribute `version` is obsolete, it will be ignored
エラー2: service "database" refers to undefined volume mysql_data: invalid compose project
```

**解決方法**:
```yaml
# ❌ 間違い
version: '3.8'  # この行を削除
volumes:
  postgres_data:  # 間違ったボリューム名

# ✅ 正解
# version行は削除
volumes:
  mysql_data:  # 正しいボリューム名
```

### 0.2. frontend/backendディレクトリが存在しないエラー
```
エラー: npm error enoent Could not read package.json
```

**原因**: `./frontend/`や`./backend/`ディレクトリがまだ作成されていない

**解決方法**:
```powershell
# 必要なディレクトリを作成
mkdir frontend, backend

# または個別に
mkdir frontend
mkdir backend

# 一時的にコンテナを停止
docker-compose down

# ディレクトリ作成後に再起動
docker-compose up -d
```

### 1. ポートが使用中エラー
```bash
# ポートを使用しているプロセスを確認
netstat -ano | findstr :3000

# プロセスを終了
taskkill /PID <PID> /F
```

### 2. ボリュームマウントの問題
```bash
# Windowsの場合、パスの区切り文字に注意
# × C:\Users\...
# ○ /c/Users/... または C:/Users/...
```

### 3. 権限エラー
```bash
# コンテナ内でのファイル権限を確認
docker-compose exec frontend ls -la

# ユーザー権限でコンテナ実行
docker-compose exec --user $(id -u):$(id -g) frontend bash
```

### 4. メモリ不足
```bash
# 不要なコンテナ・イメージを削除
docker system prune -a

# ボリュームも削除
docker system prune -a --volumes
```

## 部員向け開発フロー

### 1. 初回セットアップ
```bash
# リポジトリクローン
git clone https://github.com/USlayout/mcr-homepage.git
cd mcr-homepage

# 環境起動
docker-compose up -d

# 依存関係インストール（必要に応じて）
docker-compose exec frontend npm install
docker-compose exec backend go mod download
```

### 2. 日常の開発
```bash
# 開発開始
docker-compose up -d

# コード変更後の確認
# → ホットリロードで自動反映

# 開発終了
docker-compose down
```

### 3. 新機能追加時
```bash
# 新しい依存関係追加
docker-compose exec frontend npm install <package>
docker-compose exec backend go get <package>

# イメージ再ビルド（必要に応じて）
docker-compose build
```

## 本番環境デプロイ用設定

### docker-compose.prod.yml
```yaml
version: '3.8'

services:
  frontend:
    build: ./frontend
    environment:
      - NODE_ENV=production
    restart: unless-stopped

  backend:
    build: ./backend
    environment:
      - GO_ENV=production
    restart: unless-stopped

  database:
    image: postgres:15
    environment:
      POSTGRES_DB: ${DB_NAME}
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

volumes:
  postgres_data:
```

### 本番デプロイコマンド
```bash
# 本番用設定で起動
docker-compose -f docker-compose.prod.yml up -d

# SSL証明書設定（Let's Encrypt使用例）
docker run -it --rm \
  -v /etc/letsencrypt:/etc/letsencrypt \
  certbot/certbot certonly --webroot \
  -w /var/www/html -d your-domain.com
```

## 参考リンク

- [Docker公式ドキュメント](https://docs.docker.com/)
- [Docker Compose公式ガイド](https://docs.docker.com/compose/)
- [Node.js Docker化ベストプラクティス](https://nodejs.org/en/docs/guides/nodejs-docker-webapp/)
- [Go Docker化ガイド](https://docs.docker.com/language/golang/)

---

**注意**: このガイドはMCR部ホームページ開発用に特化しています。不明な点があれば部内で相談してください。