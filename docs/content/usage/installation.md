---
title: "Github Actions"
---

<br>

To use the Code Coverage Dashboard,
you need to host the Code Coverage Dashboard server in your environment first and then add the following 2 actions to the repository.

<br>

# __on push / on pull request__ : Publish code coverage report

The Action triggered by on push / on pull request will post a report to the Code Coverage Dashboard that hosted on your environment.
This will comment the coverage result to PR if on pull request.

```yml
steps:
  - name: Code Coverage Dashboard
    uses: rakutentech/code-coverage-dashboard@v1
    with:
      language: "go"
      api_host: "https://<your-host>/coverages-api"
      report_dir: "build"
      coverage_xml_file_name: "coverage.xml"
```

## Parameters (on push / on pull request)

| Name                   | Required | Example                                   | Description                                        |
| :--------------------- | :------: | :---------------------------------------- | :------------------------------------------------- |
| api_host               |    ◯     | https://< your-host >/coverages-api       | Code Coverage API URL that is hosted on your environment.|
| language               |    ◯     | go                                        | `go` / `php` / `js` is supported.                        |
| report_dir             |    ◯     | build                                     | Directory path where the report is located. Should define by a relative path from the root directory.|
| coverage_xml_file_name |    ◯     | coverage.xml                              | Coverage coverage report name that should be under report_dir.|
| ui_host                |    ✕     | https://< your-host >/coverages-ui        | Optional. Code Coverage UI host. [ Default: < api_host >/../coverages-ui ] |
| working_dir            |    ✕     | ./sub_directory                           | Optional. Used for specifying the sub directory that working on. [ Default: ./ (current directory) ].|
| skip_pr_comment        |    ✕     | true                                      | Optional. Enable to skip Pull Request comment. [ Default: false ]|

<br>

# __on delete__ : Sync code coverage report

The Action triggered by on delete will sync branches information with the Code Coverage Dashboard that hosted on your environment.

```yml
steps:
    - name: Sync active branches
      uses: rakutentech/code-coverage-dashboard@v1
      with:
        api_host: "https://<your-host>/coverages-api"
```

## Parameters (on delete)

| Name                   | Required | Example                                   | Description                                        |
| :--------------------- | :------: | :---------------------------------------- | :------------------------------------------------- |
| api_host               |    ◯     | https://< your-host >/coverages-api       | Code Coverage API URL that is hosted on your environment.|
