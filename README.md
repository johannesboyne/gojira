#GO Jira

*gojira* is a tiny go JIRA wraper package, at the moment only Issue Creation is implemented. If you want more functionality you are welcome to participate.

##Usage

```go
jiraConfig  := gojira.JIRAConfig{"http://localhost:8090/rest/api/2/issue/", "username", "password"}
issue       := gojira.NewIssue("Project-Key", "A nice summary text", "And a propper description, including a link: [http://jira.atlassian.com]", gojira.ISSUE_TYPE_TASK)
jiraI, _    := gojira.PostNewIssue(jiraConfig, issue)
```

##License

MIT
