package gojira

import (
	"encoding/json"
	"log"

	"github.com/franela/goreq"
)

const (
	ISSUE_TYPE_TASK      = "Task"
	ISSUE_TYPE_BUG       = "Bug"
	ISSUE_TYPE_INVENTION = "Invention"
)

type Project struct {
	Key string `json:"key"`
}

type Type struct {
	Name string `json:"name"`
}

type Issue struct {
	Project Project `json:"project"`
	Summary string  `json:"summary"`
	Desc    string  `json:"description"`
	Type    Type    `json:"issuetype"`
}

type JIRAIssue struct {
	Id   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

type JIRAPostWrap struct {
	Fields Issue `json:"fields"`
}

type JIRAConfig struct {
	URL          string
	AuthUsername string
	AuthPassword string
}

func (i *Issue) String() string {
	return i.Project.Key + " " + i.Summary + " " + i.Desc + " " + i.Type.Name
}

func NewIssue(projectKey string, summary string, desc string, issueType string) Issue {
	return Issue{Project{projectKey}, summary, desc, Type{issueType}}
}

func PostNewIssue(jC JIRAConfig, issue Issue) (JIRAIssue, error) {
	b, err := json.Marshal(&JIRAPostWrap{issue})
	// HTTP request
	res, err := goreq.Request{
		Method:            "POST",
		Uri:               jC.URL,
		BasicAuthUsername: jC.AuthUsername,
		BasicAuthPassword: jC.AuthPassword,
		ContentType:       "application/json",
		Body:              b,
	}.Do()
	jiraI := JIRAIssue{}
	res.Body.FromJsonTo(&jiraI)
	if err != nil {
		log.Fatal(err)
	}
	return jiraI, err
}
