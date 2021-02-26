package prompt

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"net/url"
)

type SurveyURL struct{}

func NewSurveyURL() InputURL {
	return SurveyURL{}
}

func (SurveyURL) URL(name, defaultValue string) (string, error) {
	var value string

	validationQs := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: name,
				Default: defaultValue,
			},
			Validate: isValidURL,
		},
	}

	return value, survey.Ask(validationQs, &value)
}

func isValidURL(value interface{}) error {
	_, err := url.ParseRequestURI(value.(string))
	if err != nil {
		return errors.New("invalid URL")
	}
	return nil
}