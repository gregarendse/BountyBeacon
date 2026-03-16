package operations

type LoggedInUserForBreadcrumbsVariables struct{}

type Account struct {
	Number string `json:"number"`
}

type LoggedInUserForBreadcrumbsData struct {
	Viewer struct {
		Accounts []Account `json:"accounts"`
	} `json:"viewer"`
}

func NewLoggedInUserForBreadcrumbsRequest() GraphQLRequest[LoggedInUserForBreadcrumbsVariables] {
	return GraphQLRequest[LoggedInUserForBreadcrumbsVariables]{
		OperationName: "getLoggedInUserForBreadcrumbs",
		Variables:     LoggedInUserForBreadcrumbsVariables{},
		Query: `query getLoggedInUserForBreadcrumbs {
			viewer {
				accounts {
					number
				}
			}
		}`,
	}
}
