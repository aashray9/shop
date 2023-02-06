package models

import "time"

type LeadFollowUp struct {
	Id                  int       `json:"id"`
	LeadId              int       `json:"lead_id"`
	FollowupTime        time.Time `json:"followup_time"`
	IsFollowed          int       `json:"is_followed"`
	FollowupAssignee    int       `json:"followup_assignee"`
	OldFollowupAssignee int       `json:"old_followup_assignee"`
	FollowedBy          int       `json:"followed_by"`
	FollowedOn          time.Time `json:"followed_on"`
	CreatedOn           time.Time `json:"created_on"`
	UnattemptedFollowup int       `json:"unattempted_followup"`
	Priority            int       `json:"priority"`
}
