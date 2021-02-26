package prompt

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"regexp"
)

type SurveyEmail struct{}

func NewSurveyEmail() InputEmail {
	return SurveyEmail{}
}

func (SurveyEmail) Email(name string) (string, error) {

	var value string

	validationQs := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: name,
			},
			Validate: isValidSurveyEmail,
		},
	}

	return value, survey.Ask(validationQs, &value)
}

func isValidSurveyEmail(email interface{}) error {
	rgx := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !rgx.MatchString(email.(string)) {
		return fmt.Errorf("%s is not a valid email", email)
	}
	return nil
}
