package trans

// Translation keys for the eventius extension (event RSVPs).
//
// These back the host-facing "a new RSVP was received" notification: a short
// message naming the guest + their answer, plus the localized answer labels
// (going / maybe / declined). Per this file's convention the const name IS the
// translation key. The message templates take fmt-style %s placeholders (the map
// translator renders via fmt.Sprintf).
const (
	// EventiusHostRsvpReceivedSubject is the email subject for the host RSVP alert.
	EventiusHostRsvpReceivedSubject = "EventiusHostRsvpReceivedSubject"
	// EventiusHostRsvpReceived is the host alert when the guest's name is known.
	// Args: %s = guest display name, %s = localized answer label.
	EventiusHostRsvpReceived = "EventiusHostRsvpReceived"
	// EventiusHostRsvpReceivedNoName is the host alert when no guest name is
	// available (e.g. an invitee-link RSVP whose name lives host-side).
	// Args: %s = localized answer label.
	EventiusHostRsvpReceivedNoName = "EventiusHostRsvpReceivedNoName"

	// EventiusRsvpAnswerGoing is the localized label for RSVP status "yes".
	EventiusRsvpAnswerGoing = "EventiusRsvpAnswerGoing"
	// EventiusRsvpAnswerMaybe is the localized label for RSVP status "maybe".
	EventiusRsvpAnswerMaybe = "EventiusRsvpAnswerMaybe"
	// EventiusRsvpAnswerDeclined is the localized label for RSVP status "no".
	EventiusRsvpAnswerDeclined = "EventiusRsvpAnswerDeclined"
)
