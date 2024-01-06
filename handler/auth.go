package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func onSuccessfulAuth(c echo.Context) error {

}

func SocialAuth(c echo.Context) error {
	gothic.GetProviderName = func(r *http.Request) (string, error) {
		param := c.Param("provider")

		if param == "" {
			return "", providerNotAvailableErr()
		}

		return param, nil
	}

	if gothUser, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request()); err == nil {
		fmt.Println(gothUser)
	} else {
		gothic.BeginAuthHandler(c.Response().Writer, c.Request())
	}

	return nil
}

func SocialAuthCallback(c echo.Context) error {
	gothic.GetProviderName = func(r *http.Request) (string, error) {
		param := c.Param("provider")

		if param == "" {
			return "", providerNotAvailableErr()
		}

		return param, nil
	}

	gothUser, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())

	if err != nil {
		return err
	}

	fmt.Println(gothUser)
	return nil
}

func SocialAuthLogout(c echo.Context) error {
	gothic.GetProviderName = func(r *http.Request) (string, error) {
		param := c.Param("provider")

		if param == "" {
			return "", providerNotAvailableErr()
		}

		return param, nil
	}

	err := gothic.Logout(c.Response().Writer, c.Request())

	if err != nil {
		return err
	}

	return nil
}
