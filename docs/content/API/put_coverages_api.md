---
title: "PUT /coverages-api"
---

Update the current branches from Github Actions.
**Use this API** in the end of action or on a, action cronjob to delete old branches from the dashboard.

```sh
curl \
-X PUT \
-H 'Content-Type: application/json' \
-H 'Authorization: GITHUB_TOKEN' \
-d '{
  "org_name": "rakutentech",
  "repo_name": "laravel-request-docs",
  "github_api_url": "api.github.com",
  "commit_hash": "abc",
  "active_branches": ["develop", "master"]
}' \
"http://localhost:3006/coverages-api"
```
