// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import internal "github.com/octo-cli/octo-cli/internal"

type RateLimitCmd struct {
	Get RateLimitGetCmd `cmd:""`
}

type RateLimitGetCmd struct {
	internal.BaseCmd
}

func (c *RateLimitGetCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/rate_limit")
	return c.DoRequest("GET")
}
