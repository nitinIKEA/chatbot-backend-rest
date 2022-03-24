package service

import "strings"

func (s *Service) PrepareRoutes() []Route {
	return Routes{
		Route{
			"Index",
			"GET",
			"/v2/",
			s.Index,
		},

		Route{
			"FaqDetailedfaqGet",
			strings.ToUpper("Get"),
			"/v2/faq/{detailedfaq}",
			s.FaqDetailedfaqGet,
		},

		Route{
			"QueryOrderstatusManordrefGet",
			strings.ToUpper("Get"),
			"/v2/query/orderstatus/manordref",
			s.QueryOrderstatusManordrefGet,
		},

		Route{
			"QueryOrderstatusOrdidrefGet",
			strings.ToUpper("Get"),
			"/v2/query/orderstatus/ordidref",
			s.QueryOrderstatusOrdidrefGet,
		},

		Route{
			"QueryOrderstatusOrdrefGet",
			strings.ToUpper("Get"),
			"/v2/query/orderstatus/ordref",
			s.QueryOrderstatusOrdrefGet,
		},

		Route{
			"QueryOrderstatusWrkordrefGet",
			strings.ToUpper("Get"),
			"/v2/query/orderstatus/wrkordref",
			s.QueryOrderstatusWrkordrefGet,
		},

		Route{
			"QueryPickingroupGet",
			strings.ToUpper("Get"),
			"/v2/query/pickingroup",
			s.QueryPickingroupGet,
		},

		Route{
			"QueryStockavailabilityAvailableGet",
			strings.ToUpper("Get"),
			"/v2/query/stockavailability/available",
			s.QueryStockavailabilityAvailableGet,
		},

		Route{
			"QueryStockavailabilityBlockedGet",
			strings.ToUpper("Get"),
			"/v2/query/stockavailability/blocked",
			s.QueryStockavailabilityBlockedGet,
		},
	}
}
