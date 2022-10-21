package helpers

import (
	"MyGramHacktiv8/pkg/errs"
	"github.com/asaskevich/govalidator"
	url2 "net/url"
)

func ValidateStruct(payload interface{}) errs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func ValidateUrl(rawUrl string) errs.MessageErr {
	_, err := url2.ParseRequestURI(rawUrl)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}
	//fmt.Println(u)
	return nil
}
