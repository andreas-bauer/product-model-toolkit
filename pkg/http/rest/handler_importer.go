// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/pkg/importing/convert"
	"github.com/osrgroup/product-model-toolkit/pkg/importing/convert/composer"
	"github.com/spdx/tools-golang/tvloader"
)

func importSPDX(c echo.Context) error {
	r := c.Request().Body
	doc, err := tvloader.Load2_1(r)
	if err != nil {
		msg := fmt.Sprintf("Error while parsing SPDX body: %v", err)
		c.Error(errors.New(msg))
	}
	msg := fmt.Sprintf("Successfully parsed SDPX document.\nFound %v packages", len(doc.Packages))

	return c.String(http.StatusOK, msg)
}

func importComposer(c echo.Context) error {
	r := c.Request().Body

	var conv convert.Converter = new(composer.Composer)
	prod, err := conv.Convert(r)
	if err != nil {
		msg := fmt.Sprintf("Error while parsing Composer JSON body: %v", err)
		c.Error(errors.New(msg))
	}

	msg := fmt.Sprintf("Successfully parsed Composer JSON.\nFound %v packages", len(prod.Components))

	return c.String(http.StatusOK, msg)
}