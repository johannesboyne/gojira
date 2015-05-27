package gojira

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestJira(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println()
		match, _ := regexp.MatchString("Basic", string(r.Header.Get("Authorization")))
		if match == true {
			fmt.Fprintln(w, `{"id":"11603","key":"CPIN-11","self":"https://localserver/2/issue/11603"}`)
		} else {
			fmt.Fprintln(w, `{"error": "missing basic auth"}`)
		}
	}))
	defer ts.Close()

	jiraConfig := JIRAConfig{ts.URL, "username", "password"}
	issue := NewIssue("Project-Key", "A nice summary text", "And a propper description, including a link: [http://jira.atlassian.com]", ISSUE_TYPE_TASK)
	jiraI, _ := PostNewIssue(jiraConfig, issue)

	if jiraI.Self != "https://localserver/2/issue/11603" {
		t.Errorf("Wrongly returned URL %v", jiraI.Self)
	}
}
