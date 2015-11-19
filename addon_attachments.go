// WARNING: This code is auto-generated from the Heroku Platform API JSON Schema
// by a Ruby script (gen/gen.rb). Changes should be made to the generation
// script rather than the generated files.

package heroku

import (
	"strings"
	"time"
)

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
	Plan struct {
		Name string `json:"name"`
	} `json:"plan"`
}

// Create a new attachment.
//
// appIdentity is the unique identifier of the attachment's App. addonName is the name of the addon to attach to.
// attachment name is the name you want. empty string is default
func (c *Client) AddonAttachmentCreate(appIdentity, addonName, attachmentName string, attachForce bool) (*AddonAttachment, error) {
	type attachmentPostBody struct {
		App   string `json:"app"`
		Addon string `json:"addon"`
		Name  string `json:"name,omitempty"`
		Force bool   `json:"force,omitempty"`
	}
	var params attachmentPostBody
	params.App = appIdentity
	params.Addon = addonName
	params.Name = attachmentName
	params.Force = attachForce
	var attachmentRes AddonAttachment
	return &attachmentRes, c.Post(&attachmentRes, "/addon-attachments", params)
}

// List all attachments attached to the given user's apps
func (c *Client) AddonAttachmentListAll() ([]AddonAttachment, error) {
	var attachments []AddonAttachment
	return attachments, c.Get(&attachments, "/addon-attachments")
}

// List all attachments for the provided addon identity
func (c *Client) AddonAttachmentList(addonIdentity string) ([]AddonAttachment, error) {
	var attachments []AddonAttachment
	return attachments, c.Get(&attachments, "/addons/"+addonIdentity+"/addon-attachments")
}

// Get info on an attachment. Will use the unscoped endpoint if there is a :: in the identity string.
func (c *Client) AddonAttachmentInfo(appIdentity string, attachmentIdentity string) (*AddonAttachment, error) {
	var attachment AddonAttachment
	if strings.Contains(attachmentIdentity, "::") {
		// identity specifies a different app, use endpoint for all accessible addons
		return &attachment, c.Get(&attachment, "/addon-attachments/"+attachmentIdentity)
	} else {
		return &attachment, c.Get(&attachment, "/apps/"+appIdentity+"/addon-attachments/"+attachmentIdentity)
	}
}

// Delete an existing attachment.
//
// appIdentity is the unique identifier of the Attachment's App. attachmentIdentity is the
// unique identifier of the attachment.
func (c *Client) AddonAttachmentDelete(attachmentIdentity string) error {
	return c.Delete("/addon-attachments/" + attachmentIdentity)
}
