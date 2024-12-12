## DockerScanner 要件定義

### Goal
- OSSを活用してDockerイメージのセキュリティ診断とDevSecOpsの実践
- 分析内容
  - ベースイメージのセキュリティ診断
  - パッケージの脆弱性診断
  - ファイルの検査 (credential, secret, ...)


### Non-Goal
- 低レイヤの脆弱性診断
- 0からのScannerの実装

### Requirement
- Dockerイメージのセキュリティ診断
  - ベースイメージのセキュリティ診断
  - パッケージの脆弱性診断
  - ファイルの検査
- reportの出力
  - PRにコメント（トリガーはPullRequest Openとする）

### Non-Functional Requirement
- OSSを活用する
- GitHub ActionsでのCI/CDの想定・実装


## DockerScanner 設計
- OS脆弱性スキャン
- npmパッケージスキャン

```mermaid
sequenceDiagram
    participant User
    participant GitHub
    participant サードパーティ

    User->>GitHub: プルリクエスト作成/更新
    GitHub->>GitHub: ワークフロー開始 (Dockerfile変更時)

    GitHub->>GitHub: コードのチェックアウト (actions/checkout)
    
    GitHub->>サードパーティ: Docker Buildxのセットアップ (docker/setup-buildx-action)
    
    GitHub->>GitHub: Dockerレイヤーキャッシュの設定 (actions/cache)
    
    GitHub->>サードパーティ: Dockerイメージのビルド (docker/build-push-action)
    
    GitHub->>サードパーティ: Trivyのセットアップ (aquasecurity/setup-trivy)
    
    GitHub->>GitHub: OS脆弱性スキャン実行 (trivy image)
    
    GitHub->>GitHub: npm依存関係のキャッシュ設定 (actions/cache)
    
    GitHub->>GitHub: npmパッケージスキャン実行
        Note over GitHub: package-lock.jsonが存在する場合
    
    GitHub->>GitHub: スキャン結果のフォーマット
    
    GitHub->>サードパーティ: PRコメント投稿 (marocchino/sticky-pull-request-comment)
    
    GitHub->>GitHub: 一時ファイルとキャッシュの整理

```