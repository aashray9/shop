package services

import (
	"lms/common"
	"lms/config"
	"lms/database"

	"github.com/golang-module/carbon/v2"
)

func TodayOnHoldBookinks(userId int) interface{} {
	var result []map[string]interface{}
	err := database.DB.Table("deal_order_management dom").
		Select("order_id,billing_cust_name,remarks,sample_collection_time").
		Joins("left join booking_remarks br on br.booking_id = dom.order_id").
		Where("delivery_status =? ", config.ONHOLD).
		Where("dom.created_by = ? ", userId).
		Where("sample_collection_time BETWEEN ? AND ?", carbon.Now().StartOfDay().ToDateTimeString(), carbon.Now().EndOfDay().ToDateTimeString()).
		Order("dom.order_id").
		Group("br.created_at DESC").
		Scan(&result).
		Error
	if err != nil {
		common.Log.Error().Str("TodayOnHoldBookinks", " query error").Msg(err.Error())
	}
	return result
}

func AgentReminderBookings(userId int) []map[string]interface{} {
	var result []map[string]interface{}
	err := database.DB.Table("deal_order_management").
		Select("order_id,billing_cust_name,sample_collection_time").
		Where("delivery_status =? ", config.ORDERBOOKED).
		Where("created_by = ? ", userId).
		Where("sample_collection_time BETWEEN ? AND ?",
			carbon.Now().AddDays(2).StartOfDay().ToDateTimeString(),
			carbon.Now().AddDays(2).EndOfDay().ToDateTimeString()).
		Scan(&result).
		Error
	if err != nil {
		common.Log.Error().Str("AgentReminderBookings", " query error").Msg(err.Error())
	}
	return result
}
