---
title: "Reverse Proxy"
---

# Using it in a reverse proxy

```conf
# serve api for GHE actions
ProxyPass         /coverages-api  http://localhost:3006
ProxyPassReverse  /coverages-api  http://localhost:3006
```