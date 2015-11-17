// WARNING: This code is auto-generated from the Heroku Platform API JSON Schema
// by a Ruby script (gen/gen.rb). Changes should be made to the generation
// script rather than the generated files.

package heroku

import "time"

type AddonAttachment struct {
	ID        string          `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	App       App             `json:"app"`
	Addon     AttachmentAddon `json:"addon"`
	Name      string          `json:"name"`
}

type ShortApp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AttachmentAddon struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	App  App    `json:"app"`
}

// Create a new add-on.
//
// appIdentity is the unique identifier of the Addon's App. plan is the unique
// identifier of this plan or unique name of this plan. options is the struct of
// optional parameters for this action.
// func (c *Client) AddonCreate(appIdentity string, plan string, options *AddonCreateOpts) (*Addon, error) {
// 	params := struct {
// 		Plan   string             `json:"plan"`
// 		Config *map[string]string `json:"config,omitempty"`
// 	}{
// 		Plan: plan,
// 	}
// 	if options != nil {
// 		params.Config = options.Config
// 	}
// 	var addonRes Addon
// 	return &addonRes, c.Post(&addonRes, "/apps/"+appIdentity+"/addons", params)
// }

// // AddonCreateOpts holds the optional parameters for AddonCreate
// type AddonCreateOpts struct {
// 	// custom add-on provisioning options
// 	Config *map[string]string `json:"config,omitempty"`
// }

// // Delete an existing add-on.
// //
// // appIdentity is the unique identifier of the Addon's App. addonIdentity is the
// // unique identifier of the Addon.
// func (c *Client) AddonDelete(appIdentity string, addonIdentity string) error {
// 	return c.Delete("/apps/" + appIdentity + "/addons/" + addonIdentity)
// }

// List all attachments attached to the given user's apps
func (c *Client) AddonAttachmentList() ([]AddonAttachment, error) {
	var attachments []AddonAttachment
	return attachments, c.Get(&attachments, "/addon-attachments")
}

// List existing add-ons.
//
// appIdentity is the unique identifier of the Addon's App. lr is an optional
// ListRange that sets the Range options for the paginated list of results.
// func (c *Client) AddonList(appIdentity string, lr *ListRange) ([]Addon, error) {
// 	req, err := c.NewRequest("GET", "/apps/"+appIdentity+"/addons", nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if lr != nil {
// 		lr.SetHeader(req)
// 	}

// 	var addonsRes []Addon
// 	return addonsRes, c.DoReq(req, &addonsRes)
// }

// // Change add-on plan. Some add-ons may not support changing plans. In that
// // case, an error will be returned.
// //
// // appIdentity is the unique identifier of the Addon's App. addonIdentity is the
// // unique identifier of the Addon. plan is the unique identifier of this plan or
// // unique name of this plan.
// func (c *Client) AddonUpdate(appIdentity string, addonIdentity string, plan string) (*Addon, error) {
// 	params := struct {
// 		Plan string `json:"plan"`
// 	}{
// 		Plan: plan,
// 	}
// 	var addonRes Addon
// 	return &addonRes, c.Patch(&addonRes, "/apps/"+appIdentity+"/addons/"+addonIdentity, params)
// }
