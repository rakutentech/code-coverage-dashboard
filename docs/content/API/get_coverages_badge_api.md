---
title: "GET /coverages-api/badge"
---


Get the Badge

```
query:"org_name" json:"org_name"  validate:"required"
query:"repo_name" json:"repo_name"  validate:"required"
query:"branch_name" json:"branch_name"  validate:"required"
query:"language" json:"language"  validate:"required,oneof=go php js"
query:"subtitle" json:"subtitle"`
```

```curl
localhost:3006/coverages-api/badge\
?org_name=<org-name>\
&repo_name=<repo-name>\
&branch_name=<branch-name>\
&language=<language>
```