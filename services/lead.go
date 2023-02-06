package services

import (
	"lms/database"
	"lms/entity"
	"lms/models"
	"time"

	"github.com/rs/zerolog/log"
)

func LeadFollowupDetails(data chan<- []map[string]interface{}, agent_id int) {
	var result []map[string]interface{}
	now := time.Now()
	lastMonth := now.AddDate(0, 0, -30).Format("2006-01-02")
	threehoursMinus := now.Add(-3 * time.Hour).Format("2006-01-02")
	hours := now.Add(-1 * time.Hour).Format("2006-01-02 15:04:05")

	err := database.DB.Table("lead_follow_up fu").
		Select("lm.first_name,lm.last_name,fu.lead_id,fu.followup_time,al.disposition,dis.name as disposition_name,al.remarks,al.created_on").
		Joins("INNER JOIN lead_activity_log al ON (al.lead_id = fu.lead_id) Left join disposition dis on (al.disposition = dis.id) LEFT join lead_master lm on (fu.lead_id = lm.lead_id)").
		Where("fu.followup_assignee", agent_id).
		Where("fu.followed_by", "0").
		Where("al.latest", 1).
		Where("fu.followup_time between ? and ?", lastMonth, threehoursMinus).
		Where("fu.followed_by", "0").
		Where("CASE WHEN al.disposition = 5 THEN al.created_on < ? ELSE TRUE END", hours).
		Order("fu.followup_time DESC").
		Offset(100).Limit(5).
		Scan(&result).Error
	if err != nil {
		log.Error().Str("LeadFollowupDetails", "Lead Followup function query error").Msg(err.Error())
	}
	data <- result
}

func Leadlist(data chan<- []map[string]interface{}, request entity.Leadlist, userId int, return_type string) {
	var result []map[string]interface{}
	dispositions := LeadTerminationStage()

	query := database.DB.Table("lead_master lm").
		Joins("LEFT JOIN lead_activity_log lal ON lm.lead_id = lal.lead_id").
		Where("lm.agent_id", userId).
		Where("CASE WHEN lm.disposition = 5 THEN lm.updated_on < NOW() - INTERVAL 1 HOUR ELSE TRUE END").
		Where("lead_terminate != ?", 1).
		Where("lm.disposition NOT IN ?", dispositions)

	if return_type == "count" {
		query.Select("COUNT(DISTINCT lm.lead_id) as total_rows")
	} else {
		query.Select("lm.lead_id,lm.first_name,lm.last_name,lm.contact_number,lm.email,lm.disposition,dis.name as disposition_name,lm.agent_id,lm.lead_source,lm.remarks,lm.locked_by,lm.updated_on,lm.created_on,lm.updated_by,ls.source_name as lead_source_name,lm.rnr_count,lm.latest_created_on,lm.call_datetime,lm.total_call_attempt,lm.agent_talk_time,lm.total_answered_call,lm.channel_type,lm.channel_user,crm_users.name,crm_users.email").
			Joins("left join disposition dis on lm.disposition = dis.id").
			Joins("left join lead_source ls on ls.id = lm.lead_source").
			Joins("left join crm_users on lm.updated_by = crm_users.id").
			Group("lm.lead_id").
			Offset(request.Offset).Limit(request.Limit)
	}
	if request.Field != "" {
		if request.Field == "first_name" {
			query.Where("CONCAT(TRIM(lm.first_name), ' ', TRIM(lm.last_name)) LIKE ?", "%%"+request.Value+"%%")
		} else {
			query.Where("lm."+request.Field, request.Value)
		}
	}

	if request.LeadDate != "" {
		query.Where("new_lead_datetime REGEXP ?", request.LeadDate)
	}
	err := query.Order("lm.latest_created_on DESC").
		Scan(&result).Error

	if err != nil {
		log.Error().Str("LeadFollowupDetails", "Lead LeadList function query error").Msg(err.Error())
	}
	data <- result
}

func LeadTerminationStage() []int {
	var ids []int
	database.DB.Table("disposition").Select("id").
		Where("termination_stage", 1).
		Pluck("id", &ids)
	return ids
}

func LeadSource() interface{} {
	val := database.GetCache("leadSource")
	if val == nil {
		var leadSource []models.LeadSource
		err := database.DB.Table("lead_source").
			Select("id,source_name,inbound_visibility, is_premium, lead_assign_status,lead_assign_department,source_class").
			Where("active", 1).
			Order("source_name ASC").Find(&leadSource).Error
		if err == nil {
			database.SetCache("leadSource", leadSource, 4*60)
		}
		log.Info().Msg("=============>data comes from database")
		return leadSource
	} else {
		log.Info().Msg("=============>data comes from Redis cache")
		return val
	}
}

func GetDisposition() interface{} {
	val := database.GetCache("disposition_list")
	if val == nil {
		var disposition []models.Disposition
		err := database.DB.Table("disposition").
			Select("id,name,isactive,parent_id").
			Where("(disposition_type = ? || disposition_type = ?)", 0, 1).
			Order("name ASC").Find(&disposition).Error
		if err == nil {
			database.SetCache("disposition_list", disposition, 4*60)
		}
		return disposition
	} else {
		return val
	}
}

func GetLeadActivityLog(request entity.LeadActivityLog) []map[string]interface{} {
	var result []map[string]interface{}
	query := database.DB.Table("lead_activity_log lal").Select("lal.id,lal.lead_id,lal.disposition as disposition_id,dis.name as disposition_name,lal.ad_name as campaign_name, lal.created_by,lal.created_on,lal.remarks,lal.campaign_id,lal.lead_source,lal.recommended_test, lal.channel_user, lal.channel_type,ls.source_name as lead_source_name,lal.prescription_image,lal.agent_id,lal.utm_campaign, crm_users.name as created_by").
		Joins("left join lead_source ls on lal.lead_source = ls.id").
		Joins("left join disposition dis on dis.id = lal.disposition").
		Joins("left join crm_users on lal.agent_id = crm_users.id")
	if request.LeadId > 0 {
		query.Where("lal.lead_id", request.LeadId)
	}
	if request.Limit > 0 {
		query.Offset(request.Offset).Limit(request.Limit)
	}
	query.Order("lal.id DESC").
		Scan(&result)
	return result
}
