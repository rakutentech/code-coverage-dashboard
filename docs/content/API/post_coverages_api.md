---
title: "POST /coverages-api"
---

POST the archive tar that has coverage HTML and Cover XML to the API endpoint.<br>
API stores the HTML for static view. Cover XML for parsing total percentage.
Responds with the coverages for this org/repo for all branches.

**Use this API** to create new coverage from Github Actions.



```
curl -X POST -H 'Content-Type: multipart/form-data' -H 'Authorization: GITHUB_TOKEN' \
--form file=@/Users/pulkit.kathuria/git/code-coverage-dashboard/server/test_data/go_coverage.tar.gz \
"localhost:3000/coverages-api\
?org_name=rakuten\
&branch_name=develop\
&github_api_url=https://api.github.com\
&repo_name=laravel-request-docs\
&commit_hash=a123...5716217f903ebffe42c3f113066103eca9\
&commit_author=pulkit.kathuria\
&language=go\
&coverage_xml_file_name=coverage.xml"
```

Example Response

```json
{
  "coverage": {
    "id": 15,
    "org_name": "rakuten",
    "repo_name": "laravel-request-docs",
    "branch_name": "develop",
    "commit_hash": "a123...5716217f903ebffe42c3f113066103eca9",
    "commit_author": "pulkit.kathuria",
    "language": "php",
    "percentage": 33.03,
    "created_at": "2022-04-07T09:21:56.879+09:00",
    "updated_at": "2022-04-07T09:21:56.879+09:00"
  },
  "data": [
    {
      "id": 3,
      "org_name": "rakuten",
      "repo_name": "laravel-request-docs",
      "branch_name": "develop",
      "commit_hash": "a123...5716217f903ebffe42c3f113066103eca9",
      "commit_author": "pulkit.kathuria",
      "language": "php",
      "percentage": 43.03,
      "created_at": "2022-04-05T21:10:27.716515+09:00",
      "updated_at": "2022-04-05T21:10:48.367029+09:00"
    },
    {
      "id": 2,
      "org_name": "rakuten",
      "repo_name": "laravel-request-docs",
      "branch_name": "master",
      "commit_hash": "a123...5716217f903ebffe42c3f113066103eca9",
      "commit_author": "pulkit.kathuria",
      "language": "php",
      "percentage": 33.03,
      "created_at": "2022-04-05T21:10:27.716515+09:00",
      "updated_at": "2022-04-05T21:10:48.367029+09:00"
    },
    {
        ...repo data for other branches
    }
  ]
}
```