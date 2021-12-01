package freshdesk

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	cg "prod/src/customlibrary/configuration"
)

func GenerateCommonTicketData(data map[string]interface{}) {

	postdata := map[string]interface{}{}
	if len(data) > 0 {
		reporterName := data["name"].(string)
		reporterEmail := data["email"].(string)
		htmlBody := ""
		htmlBody += "From : " + reporterName + " <" + reporterEmail + ">"
		if _, ok := data["descriptionData"]; ok {
			descriptionData := data["descriptionData"].(map[string]interface{})
			for key, smtpVal := range descriptionData {
				htmlBody += key + " : " + smtpVal.(string) + "\n"
			}
		}

		postdata["description"] = htmlBody
		postdata["subject"] = data["subject"].(string)
	}
	//fmt.Println(postdata)
	sent(postdata)

}

func sent(data map[string]interface{}) {

	base := "https://" + cg.Config.Jira.Host
	tp := jira.BasicAuthTransport{
		Username: cg.Config.Jira.Username,
		Password: cg.Config.Jira.ApiKey,
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		// panic(err)
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				AccountID: cg.Config.Jira.AssigneeId,
			},

			Description: data["description"].(string),
			Type: jira.IssueType{
				Name: "Bug",
			},
			Project: jira.Project{
				Key: cg.Config.Jira.ProjectKey,
			},
			Summary: data["subject"].(string),
		},
	}
	issue, _, err := jiraClient.Issue.Create(&i)
	if err != nil {
		// panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)

}
