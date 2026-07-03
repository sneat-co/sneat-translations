package trans

// Translation keys for the invitus extension (space-join invites).
//
// These back the Sneat space-join INVITE EMAIL delivered when someone invites a
// person to join a space (team/family/etc). They are DISTINCT from the legacy
// DebtsTracker EMAIL_INVITE_* keys, which reference debtstracker.io + an invite
// CODE and would render the wrong content for a Sneat space-join invite (which
// uses a PIN + a /join/{spaceType} link).
//
// Per this file's convention the const name IS the translation key.
//
//   - InvitusSpaceInviteEmailSubject uses fmt-style %s placeholders (rendered via
//     the map translator's fmt.Sprintf path). Args, in order:
//     %s = inviter display name, %s = space title.
//   - InvitusSpaceInviteEmailHtml uses Go html/template {{.Field}} placeholders
//     (rendered via the map translator's single-struct template path). Fields:
//     {{.FromHTML}}   = inviter display name (optionally a mailto anchor)
//     {{.SpaceTitle}} = title of the space being joined
//     {{.HostPath}}   = app host + path used to build the join link
//     {{.SpaceType}}  = space type, drives the /join/{type} link
//     {{.ID}}         = invite ID (the join link "id" query value)
//     {{.PinCode}}    = personal PIN code (appears twice: in the link and in the body)
const (
	// InvitusSpaceInviteEmailSubject is the space-join invite email subject.
	InvitusSpaceInviteEmailSubject = "InvitusSpaceInviteEmailSubject"
	// InvitusSpaceInviteEmailHtml is the space-join invite email HTML body.
	InvitusSpaceInviteEmailHtml = "InvitusSpaceInviteEmailHtml"
)
