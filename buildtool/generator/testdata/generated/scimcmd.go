// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type ScimCmd struct {
	GetProvisioningDetailsForUser     ScimGetProvisioningDetailsForUserCmd     `cmd:""`
	ListProvisionedIdentities         ScimListProvisionedIdentitiesCmd         `cmd:""`
	ProvisionAndInviteUsers           ScimProvisionAndInviteUsersCmd           `cmd:""`
	RemoveUserFromOrg                 ScimRemoveUserFromOrgCmd                 `cmd:""`
	ReplaceProvisionedUserInformation ScimReplaceProvisionedUserInformationCmd `cmd:""`
	UpdateUserAttribute               ScimUpdateUserAttributeCmd               `cmd:""`
}

type ScimGetProvisioningDetailsForUserCmd struct {
	Org        string `required:"" name:"org"`
	ScimUserId int64  `required:"" name:"scim_user_id"`
	internal.BaseCmd
}

func (c *ScimGetProvisioningDetailsForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users/:scim_user_id")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("scim_user_id", c.ScimUserId)
	return c.DoRequest("GET")
}

type ScimListProvisionedIdentitiesCmd struct {
	Count      int64  `name:"count"`
	Filter     string `name:"filter"`
	Org        string `required:"" name:"org"`
	StartIndex int64  `name:"startIndex"`
	internal.BaseCmd
}

func (c *ScimListProvisionedIdentitiesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users")
	c.UpdateURLQuery("count", c.Count)
	c.UpdateURLQuery("filter", c.Filter)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("startIndex", c.StartIndex)
	return c.DoRequest("GET")
}

type ScimProvisionAndInviteUsersCmd struct {
	Org string `required:"" name:"org"`
	internal.BaseCmd
}

func (c *ScimProvisionAndInviteUsersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users")
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("POST")
}

type ScimRemoveUserFromOrgCmd struct {
	Org        string `required:"" name:"org"`
	ScimUserId int64  `required:"" name:"scim_user_id"`
	internal.BaseCmd
}

func (c *ScimRemoveUserFromOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users/:scim_user_id")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("scim_user_id", c.ScimUserId)
	return c.DoRequest("DELETE")
}

type ScimReplaceProvisionedUserInformationCmd struct {
	Org        string `required:"" name:"org"`
	ScimUserId int64  `required:"" name:"scim_user_id"`
	internal.BaseCmd
}

func (c *ScimReplaceProvisionedUserInformationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users/:scim_user_id")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("scim_user_id", c.ScimUserId)
	return c.DoRequest("PUT")
}

type ScimUpdateUserAttributeCmd struct {
	Org        string `required:"" name:"org"`
	ScimUserId int64  `required:"" name:"scim_user_id"`
	internal.BaseCmd
}

func (c *ScimUpdateUserAttributeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/scim/v2/organizations/:org/Users/:scim_user_id")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("scim_user_id", c.ScimUserId)
	return c.DoRequest("PATCH")
}