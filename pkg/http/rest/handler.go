// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/osrgroup/product-model-toolkit/pkg/importing"
	"github.com/osrgroup/product-model-toolkit/pkg/querying"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

// Handler handle all request for the given route group.
func Handler(g *echo.Group, qSrv querying.Service, iSrv importing.Service) {
	g.GET("/", handleEntryPoint)
	g.GET("/version", handleVersion)
	g.GET("/health", handleHealth)

	g.GET("/products", findAllProducts(qSrv))
	g.GET("/products/:id", findProductByID(qSrv))
	g.POST("/products/spdx", importSPDX(iSrv))
	g.POST("/products/composer", importComposer(iSrv))
}

func handleEntryPoint(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Echo().Routers())
}

func handleVersion(c echo.Context) error {
	return c.String(http.StatusOK, version.Name())
}

func handleHealth(c echo.Context) error {
	type status struct {
		Status string `json:"status"`
	}

	up := status{Status: "UP"}
	return c.JSON(http.StatusOK, up)
}

func findAllProducts(q querying.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		prods, err := q.FindAllProducts()
		if err != nil {
			c.Error(err)
		}

		return c.JSON(http.StatusOK, prods)
	}
}

func findProductByID(q querying.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idstr, err := strconv.Atoi(id)
		if err != nil {
			c.Error(err)
		}

		prod, err := q.FindProductByID(idstr)
		if err != nil {
			c.Error(err)
		}

		return c.JSON(http.StatusOK, prod)
	}
}
