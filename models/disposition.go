package models

type Disposition struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	ParentId         int    `json:"parent_id"`
	Isactive         int    `json:"isactive"`
	DispositionType  int    `json:"disposition_type"`
	TerminationStage int    `json:"termination_stage"`
	DisplayText      string `json:"display_text"`
}

type ConfigMaster struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueType string `json:"value_type"`
	Module    string `json:"module"`
}

type LeadSource struct {
	Id                   int    `json:"id"`
	SourceName           string `json:"source_name"`
	IsPremium            int    `json:"is_premium"`
	InboundVisibility    int    `json:"inbound_visibility"`
	LeadAssignStatus     int    `json:"lead_assign_status"`
	LeadAssignDepartment int    `json:"lead_assign_department"`
	SourceClass          string `json:"source_class"`
}
