# docker-scanner

```mermaid
sequenceDiagram
    participant user as User
    participant github as GitHub
    participant actions as Actions(CI/CD)
    participant trivy as Trivy
        
    user ->>+github: Push Dockerfile
    github ->>actions: Trigger
    par Dockerfile解析
    actions ->>trivy: パッケージファイルの診断
    trivy ->>actions: response
    actions ->>actions: Dockerfile解析
    end
   actions ->>github: report if PR
```