package errorhandler

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	ticket "prod/src/customlibrary/jira"
)

var ListName string
var ListId string

func ErrorCheck(err error) {
	if err != nil {
		orignalError := errors.New(fmt.Sprint(err))
		errorText := fmt.Sprintf(" %+v \n\n", orignalError)
		freshdeskError(errorText)
		fmt.Println(err.Error())
		fmt.Println(errorText)
		os.Exit(1)
	}
}

func freshdeskError(errorText string) {
	descriptionData := map[string]interface{}{}

	descriptionData["Reason"] = errorText

	name := "Admin"
	email := "admin@test.com"

	DataConfigs := map[string]interface{}{
		"name":  name,
		"email": email,

		"subject":         "Queue mechanism failed",
		"priority":        4,
		"descriptionData": descriptionData,
	}
	ticket.GenerateCommonTicketData(DataConfigs)
}
