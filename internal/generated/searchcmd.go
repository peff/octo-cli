// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import (
	"github.com/octo-cli/octo-cli/internal"
)

type SearchCmd struct {
	Code                  SearchCodeCmd                  `cmd:""`
	Commits               SearchCommitsCmd               `cmd:""`
	IssuesAndPullRequests SearchIssuesAndPullRequestsCmd `cmd:""`
	Labels                SearchLabelsCmd                `cmd:""`
	Repos                 SearchReposCmd                 `cmd:""`
	Topics                SearchTopicsCmd                `cmd:""`
	Users                 SearchUsersCmd                 `cmd:""`
}

type SearchCodeCmd struct {
	Q       string `required:"" name:"q"`
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchCodeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/code")
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type SearchCommitsCmd struct {
	Cloak   bool   `required:"" name:"cloak-preview"`
	Q       string `required:"" name:"q"`
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchCommitsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/commits")
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("cloak", c.Cloak)
	return c.DoRequest("GET")
}

type SearchIssuesAndPullRequestsCmd struct {
	Q       string `required:"" name:"q"`
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchIssuesAndPullRequestsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/issues")
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type SearchLabelsCmd struct {
	Q            string `required:"" name:"q"`
	RepositoryId int64  `required:"" name:"repository_id"`
	Order        string `name:"order"`
	Sort         string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchLabelsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/labels")
	c.UpdateURLQuery("repository_id", c.RepositoryId)
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("order", c.Order)
	return c.DoRequest("GET")
}

type SearchReposCmd struct {
	Mercy   bool   `name:"mercy-preview"`
	Q       string `required:"" name:"q"`
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchReposCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/repositories")
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("mercy", c.Mercy)
	return c.DoRequest("GET")
}

type SearchTopicsCmd struct {
	Mercy bool   `name:"mercy-preview"`
	Q     string `required:"" name:"q"`
	internal.BaseCmd
}

func (c *SearchTopicsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/topics")
	c.UpdateURLQuery("q", c.Q)
	c.UpdatePreview("mercy", c.Mercy)
	return c.DoRequest("GET")
}

type SearchUsersCmd struct {
	Q       string `required:"" name:"q"`
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchUsersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/users")
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}
