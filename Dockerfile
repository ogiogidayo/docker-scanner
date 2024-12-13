# Ubuntu 20.04をベースイメージとして使用
FROM ubuntu:20.04

# セキュリティ上の問題がある可能性のあるパッケージをインストール
RUN apt-get update && apt-get install -y \
    openssh-server=1:8.2p1-4ubuntu0.11 \
    vsftpd=3.0.3-12ubuntu2.1 \
    wget=1.20.3-1ubuntu2.1

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
CMD ["/bin/bash"]
