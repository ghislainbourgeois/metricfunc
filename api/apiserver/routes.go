// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package apiserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/omec-project/logger_util"
	"github.com/omec-project/metricfunc/logger"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := logger_util.NewGinWithLogrus(logger.ApiSrvLog)
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nmetric-func/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"GetSubscriberSummary",
		strings.ToUpper("Get"),
		"/subscriber/:imsi",
		GetSubscriberSummary,
	},

	{
		"GetSubscriberAll",
		strings.ToUpper("Get"),
		"/subscriber/all",
		GetSubscriberAll,
	},

	{
		"GetNfStatus",
		strings.ToUpper("Get"),
		"/nfstatus/:type",
		GetNfStatus,
	},

	{
		"GetNfStatusAll",
		strings.ToUpper("Get"),
		"/nfstatus/all",
		GetNfStatusAll,
	},
	{
		"GetNfServiceStatsSummary",
		strings.ToUpper("Get"),
		"/nfServiceStatsSummary/:type",
		GetNfServiceStatsSummary,
	},
	{
		"GetNfServiceStatsDetail",
		strings.ToUpper("Get"),
		"/nfServiceStatsDetail/:type",
		GetNfServiceStatsDetail,
	},

	{
		"GetNfServiceStatsAll",
		strings.ToUpper("Get"),
		"/nfServiceStats/all",
		GetNfServiceStatsAll,
	},

	{
		"TestIPs",
		strings.ToUpper("Post"),
		"/testIPs",
		PushTestIPs,
	},
}

/* APIs
State Info of subscriber(not as SMF level but higher level like connected/idle/active)
Stats covered
In memory Database details
Prometheus
*/
