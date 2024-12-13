## OSパッケージ脆弱性レポート

### my-docker-image (Ubuntu 16.04)
- **検出された脆弱性**: 0
    - **UNKNOWN**: 0
    - **LOW**: 0
    - **MEDIUM**: 0
    - **HIGH**: 0
    - **CRITICAL**: 0

### 機密情報の検出

| ファイルパス | 警告レベル |
|--------------|------------|
| /etc/ssh/ssh_host_dsa_key | 🔴 HIGH |
| /etc/ssh/ssh_host_ecdsa_key | 🔴 HIGH |
| /etc/ssh/ssh_host_ed25519_key | 🔴 HIGH |
| /etc/ssh/ssh_host_rsa_key | 🔴 HIGH |
| /etc/ssl/private/ssl-cert-snakeoil.key | 🔴 HIGH |

#### 詳細

- **/etc/ssh/ssh_host_dsa_key**
    - 内容: DSAプライベートキー
    - `RUN /bin/sh -c apt-get update && apt-get`によって追加
- **/etc/ssh/ssh_host_ecdsa_key**
    - 内容: ECプライベートキー
    - `RUN /bin/sh -c apt-get update && apt-get`によって追加
- **/etc/ssh/ssh_host_ed25519_key**
    - 内容: OpenSSHプライベートキー
    - `RUN /bin/sh -c apt-get update && apt-get`によって追加
- **/etc/ssh/ssh_host_rsa_key**
    - 内容: RSAプライベートキー
    - `RUN /bin/sh -c apt-get update && apt-get`によって追加
- **/etc/ssl/private/ssl-cert-snakeoil.key**
    - 内容: プライベートキー
    - `RUN /bin/sh -c apt-get update && apt-get`によって追加


