# ベースイメージとして既知の脆弱性を含む古いバージョンを使用
FROM ubuntu:20.04

# 必要なパッケージをインストール（古いバージョンを意図的に指定）
RUN apt-get update && apt-get install -y \
    openssh-server=1:7.2p2-4ubuntu2.10 \
    vsftpd=3.0.3-3ubuntu2 \
    wget=1.17.1-1ubuntu1.5

# rootユーザーで動作（非推奨）
USER root

# 不適切な権限設定（すべてのユーザーがアクセス可能）
RUN chmod 777 /etc/passwd

# 不要なポートを公開
EXPOSE 22 21 80

# 環境変数に機密情報を直接記載（非常に危険）
ENV DB_PASSWORD="password123"

# デフォルトコマンド
CMD ["/bin/bash"]
