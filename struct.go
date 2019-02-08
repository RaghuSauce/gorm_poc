package main

type APIReturnError struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

// PRList stores a list of pull requests associated with a single team.
type PRList struct {
	Team string        `json:"team"`
	PRS  []GithubEvent `json:"prs"`
}

// GithubEvent defines the document stored in couchbase for tracking data
// about a particular pull request submitted to Github.
type GithubEvent struct {
	Action           string          `json:"action"`
	Number           int             `json:"number"`
	Team             string          `json:"team"`
	State            string          `json:"state"`
	ReviewStart      string          `json:"review_start"`
	ReviewEnd        string          `json:"review_end"`
	MergedTime       string          `json:"merged_time"`
	ChangeRequestors []string        `json:"change_requestors"`
	ChangeRequested  bool            `json:"change_requested"`
	ChangeApprovers  []string        `json:"change_approver"`
	MergedBy         string          `json:"merged_by"`
	PullRequest      PullRequest     `json:"pull_request"`
	Repository       Repository      `json:"repository"`
	Issue            Issue           `json:"issue"`
	Comment          Comment         `json:"comment"`
	Notifications    Notifications   `json:"notifications"`
	Review           Review          `json:"review"`
	Sender           Login           `json:"sender"`
	Reactions        []SlackReaction `json:"reactions"`

	// JiraTicketComment stores the ID of the Github comment posted
	// to the PR indicating which tickets were referenced.
	JiraTicketComment int `json:"jira_ticket_comment"`

	// ReminderRequester records which slack user posted the link
	//   in the first place.
	ReminderRequester string `json:"reminder_requester"`
	// ReminderChannel records the room to which the PR Manager Bot
	//    should post reminders. If this is blank, PR Bot will ignore the PR.
	ReminderChannel string `json:"reminder_channel"`

	// ReminderMessages stores the timestamps of the messages in the
	// 	ReminderChannel that are associated with the PR. It's used for
	//	knowing which messages should receive reactions and threaded responses.
	ReminderMessages []string `json:"reminder_messages"`
}

// SlackReaction stores reactions already posted to the message by the bot.
type SlackReaction struct {
	Timestamp string `json:"ts"`
	Reaction  string `json:"reaction"`
}

// PullRequest stores information within a Github event object.
type PullRequest struct {
	Title    string `json:"title"`
	URL      string `json:"html_url" validate:"nonzero"`
	User     User   `json:"user"`
	Assignee Login  `json:"assignee"`
	Head     struct {
		Ref string `json:"ref"`
		SHA string `json:"sha"`
	} `json:"head"`
	CreatedDateTime string `json:"created_at"`
	Merged          bool   `json:"merged"`
	Commits         int    `json:"commits"`
	Additions       int    `json:"additions"`
	Deletions       int    `json:"deletions"`
	ChangedFiles    int    `json:"changed_files"`
	MergedBy        struct {
		Login  string `json:"login"`
		Avatar string `json:"avatar_url"`
	} `json:"merged_by"`
}

// Login stores information within a Github event object.
type Login struct {
	Login string `json:"login"`
}

// Issue stores information about a single pull request  within a Github event object.
type Issue struct {
	PullRequest struct {
		URL string `json:"html_url"`
	} `json:"pull_request"`
}

// Comment stores information about a PR comment within a Github event object.
type Comment struct {
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

// User stores information about the creator of an item within a Github event object.
type User struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

// Repository stores information about the parent repo within a Github event object.
type Repository struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
}

// Review stores information about a single PR's approval within a Github event object.
type Review struct {
	Body      string   `json:"body"`
	State     string   `json:"state"`
	User      User     `json:"user"`
	Reviewers []string `json:"reviewers"`
}

// Notifications tracks which reminders have been posted to slack for an event.
type Notifications struct {
	Hour2    Hour2Notification    `json:"hour2"`
	Hour4    Hour4Notification    `json:"hour4"`
	Hour12   Hour12Notification   `json:"hour12"`
	Beyond12 Beyond12Notification `json:"beyond12"`
}

type Hour2Notification bool
type Hour4Notification bool
type Hour12Notification bool
type Beyond12Notification struct {
	LastNotification string `json:"last_notification"`
}

// OptionalFlags stores information about optional parameters that can be sent with webhooks.
type OptionalFlags struct {
	DisableBranchDeletion bool               `json:"disableBranchDeletion"`
	RequiredReviewers     int                `json:"requiredReviewers"`
	SlackNotifications    SlackNotifications `json:"slackNotify"`
}

// SlackNotifications - Optional Parameter for notifying slack when a PR does various things
type SlackNotifications []struct {
	WebhookURL string   `json:"webhook_url"`
	Channels   []string `json:"channels"`
	Actions    []string `json:"actions"`
}
