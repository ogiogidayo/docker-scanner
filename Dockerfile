# Ubuntu 20.04をベースイメージを使用（脆弱性のあるイメージ）
FROM ubuntu:20.04

# セキュリティ上の問題がある可能性のあるパッケージをインストール
RUN apt-get update && apt-get install -y \
    openssh-server=1:8.2p1-4ubuntu0.11 \
    wget=1.20.3-1ubuntu2.1 \
    npm=6.14.4+ds-1ubuntu2

# タイムゾーンを設定
RUN ln -fs /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    dpkg-reconfigure --frontend noninteractive tzdata

# rootユーザーで動作（非推奨）
USER root

# 不適切な権限設定（すべてのユーザーがアクセス可能）
RUN chmod 777 /etc/passwd

# 不要なポートを公開
EXPOSE 22 21 80

# 環境変数に機密情報を直接記載
ENV DB_PASSWORD="password123"

# 脆弱なディレクトリに作業ディレクトリを設定
WORKDIR /proc/self/fd/8

# npmで脆弱なパッケージをインストール
RUN npm install -g lodash@4.17.20

# 任意のスクリプトを実行
CMD ["/bin/bash"]
