// WARNING: This code is auto-generated from the Heroku Platform API JSON Schema
// by a Ruby script (gen/gen.rb). Changes should be made to the generation
// script rather than the generated files.

package heroku

import (
	"strings"
	"time"
)

// Add-ons represent add-ons that have been provisioned for an app.
type Addon struct {
	// config vars associated with this application
	ConfigVars []string `json:"config_vars"`

	// when add-on was updated
	CreatedAt time.Time `json:"created_at"`

	// unique identifier of add-on
	Id string `json:"id"`

	// name of the add-on unique within its app
	Name string `json:"name"`

	// message returned from provider
	Message string `json:"message"`

	// identity of add-on plan
	Plan struct {
		Id    string          `json:"id"`
		Name  string          `json:"name"`
		Price *AddonPlanPrice `json:"price,omitempty`
	} `json:"plan"`

	App App `json:"app"`

	// id of this add-on with its provider
	ProviderId string `json:"provider_id"`

	// when add-on was updated
	UpdatedAt time.Time `json:"updated_at"`
}

type AddonPlanPrice struct {
	Cents int `json:"cents"`
}

// Create a new add-on.
//
// appIdentity is the unique identifier of the Addon's App. plan is the unique
// identifier of this plan or unique name of this plan. options is the struct of
// optional parameters for this action.
func (c *Client) AddonCreate(appIdentity string, plan string, options *AddonCreateOpts) (*Addon, error) {
	params := struct {
		Plan       string             `json:"plan"`
		Config     *map[string]string `json:"config,omitempty"`
		Attachment *map[string]string `json:"attachment,omitempty"`
	}{
		Plan: plan,
	}
	if options != nil && options.Config != nil {
		params.Config = options.Config
	}
	if options != nil && options.Attachment != nil {
		params.Attachment = options.Attachment
	}
	var addonRes Addon
	return &addonRes, c.Post(&addonRes, "/apps/"+appIdentity+"/addons", params)
}

// AddonCreateOpts holds the optional parameters for AddonCreate
type AddonCreateOpts struct {
	// custom add-on provisioning options
	Config     *map[string]string `json:"config,omitempty"`
	Attachment *map[string]string `json:"attachment,omitempty"`
}

// Delete an existing add-on.
//
// appIdentity is the unique identifier of the Addon's App. addonIdentity is the
// unique identifier of the Addon.
func (c *Client) AddonDelete(appIdentity string, addonIdentity string) error {
	return c.Delete("/apps/" + appIdentity + "/addons/" + addonIdentity)
}

// Info for an existing add-on.
//
// appIdentity is the unique identifier of the Addon's App. addonIdentity is the
// unique identifier of the Addon.
func (c *Client) AddonInfo(appIdentity string, addonIdentity string) (*Addon, error) {
	var addon Addon
	if strings.Contains(addonIdentity, "::") {
		// identity specifies a different app, use global endpoint
		return c.GlobalAddonInfo(addonIdentity)
	} else {
		return &addon, c.Get(&addon, "/apps/"+appIdentity+"/addons/"+addonIdentity)
	}
}

func (c *Client) GlobalAddonInfo(addonIdentity string) (*Addon, error) {
	var addon Addon
	return &addon, c.Get(&addon, "/addons/"+addonIdentity)
}

// List existing add-ons.
//
// appIdentity is the unique identifier of the Addon's App. lr is an optional
// ListRange that sets the Range options for the paginated list of results.
func (c *Client) AddonList(appIdentity string, lr *ListRange) ([]Addon, error) {
	req, err := c.NewRequest("GET", "/apps/"+appIdentity+"/addons", nil)
	if err != nil {
		return nil, err
	}

	if lr != nil {
		lr.SetHeader(req)
	}

	var addonsRes []Addon
	return addonsRes, c.DoReq(req, &addonsRes)
}

// Change add-on plan. Some add-ons may not support changing plans. In that
// case, an error will be returned.
//
// appIdentity is the unique identifier of the Addon's App. addonIdentity is the
// unique identifier of the Addon. plan is the unique identifier of this plan or
// unique name of this plan.
func (c *Client) AddonUpdate(appIdentity string, addonIdentity string, plan string) (*Addon, error) {
	params := struct {
		Plan string `json:"plan"`
	}{
		Plan: plan,
	}
	var addonRes Addon
	return &addonRes, c.Patch(&addonRes, "/apps/"+appIdentity+"/addons/"+addonIdentity, params)
}

// List all addons belonging to the current user
func (c *Client) AddonListAll() ([]Addon, error) {
	var addons []Addon
	return addons, c.Get(&addons, "/addons")
}
