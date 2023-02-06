package models

import "time"

type LeadMasterModal struct {
	LeadId            int       `json:"lead_id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	ContactNumber     string    `json:"contact_number"`
	Email             string    `json:"email"`
	Disposition       int       `json:"disposition"`
	AgentId           int       `json:"agent_id"`
	LeadSource        string    `json:"lead_source"`
	Remarks           string    `json:"remarks"`
	LockedBy          int       `json:"locked_by"`
	UpdatedOn         time.Time `json:"updated_on"`
	CreatedOn         time.Time `json:"created_on"`
	UpdatedBy         int       `json:"updated_by"`
	RnrCount          int       `json:"rnr_count"`
	LatestCreatedOn   time.Time `json:"latest_created_on"`
	CallDatetime      time.Time `json:"call_datetime"`
	TotalCallAttempt  int       `json:"total_call_attempt"`
	AgentTalkTime     int       `json:"agent_talk_time"`
	TotalAnsweredCall int       `json:"total_answered_call"`
	ChannelType       int       `json:"channel_type"`
	ChannelUser       int       `json:"channel_user"`
	NewLeadDatetime   string    `json:"new_lead_datetime"`
}

type MethodChan struct {
	Key   chan<- string
	Value chan<- interface{}
}

type RecommendedTest struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type ProductType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
