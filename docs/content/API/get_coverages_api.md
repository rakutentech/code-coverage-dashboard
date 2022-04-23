---
title: "GET /coverages-api"
---

Paginated list of coverages for all orgs and repos. <br>
**Use this API** for the dashboard to show all coverages in your organization.


```sh
curl localhost:3000/coverages-api?p=1
```

Request

```
query:"org_name"  hint:"To filter by org name"`
query:"repo_name"  hint:"To filter by repository name"`
query:"full"  hint:"To include all history for trends"`
query:"p"  validate:"gte=0"  message:"p greater than 0" hint:"Page number for pagination"`
```

Response will be a paginated result of coverages stored inside

```json
{
  "has_next": false,
  "data": {
    "rakutentech/laravel-request-docs": [
      {
        "id": 1,
        "org_name": "rakutentech",
        "repo_name": "laravel-request-docs",
        "branch_name": "develop",
        "commit_hash": "a124...ebffe42c3f113066103eca9",
        "commit_author": "james.bond",
        "language": "php",
        "percentage": 33.03,
        "created_at": "2022-04-04T23:17:04.400701+09:00",
        "updated_at": "2022-04-05T17:15:59.62489+09:00"
      },
      {
        "id": 2,
        "org_name": "rakutentech",
        "repo_name": "laravel-request-docs",
        "branch_name": "feature/test",
        "commit_hash": "a125...ebffe42c3f113066103eca9",
        "commit_author": "james.bond",
        "language": "php",
        "percentage": 43.03,
        "created_at": "2022-04-05T14:10:58.254005+09:00",
        "updated_at": "2022-04-05T14:10:58.254005+09:00"
      }
    ],
    "rakutentech/awesome_app": [
      {
        "id": 3,
        "org_name": "rakutentech",
        "repo_name": "awesome_app",
        "branch_name": "develop",
        "commit_hash": "a123...ebffe42c3f113066103eca9",
        "commit_author": "james.bond",
        "language": "go",
        "percentage": 30.02,
        "created_at": "2022-04-05T21:10:27.716515+09:00",
        "updated_at": "2022-04-05T21:10:48.367029+09:00"
      }
    ]
  }
}
```