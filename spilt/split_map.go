package split

func GetType1() Type1 {
	type1 := map[string]interface{}{
		"db": map[string]interface{}{
			"tb": nil,
		},
		"db_ulu_3rddevice": map[string]interface{}{
			"tb_3rddevice": nil,
		},
		"db_ulu_3rddevice_set": map[string]interface{}{
			"tb_3rddevice_set": nil,
		},
		"db_ulu_b_widget": map[string]interface{}{
			"tb_b_widget":           nil,
			"tb_b_widget_order":     nil,
			"tb_b_widget_registere": nil,
		},
		"db_ulu_business": map[string]interface{}{
			"tb_business":          nil,
			"tb_business_detailed": nil,
		},
		"db_ulu_counter": map[string]interface{}{
			"tb_event_store_day_count":          nil,
			"tb_event_store_property_day_count": nil,
			"tb_plan_count":                     nil,
			"tb_store_plan_motion_count":        nil,
			"tb_store_plan_tuogang_count":       nil,
		},
		"db_ulu_device": map[string]interface{}{
			"tb_device": nil,
		},
		"db_ulu_device_bind_record": map[string]interface{}{
			"tb_device_bind_b": nil,
		},
		"db_ulu_device_channel": map[string]interface{}{
			"tb_device_channel": nil,
		},
		"db_ulu_device_local_account": map[string]interface{}{
			"tb_device_account": nil,
		},
		"db_ulu_device_set": map[string]interface{}{
			"tb_channel_set":  nil,
			"tb_config":       nil,
			"tb_device_cloud": nil,
			"tb_device_set":   nil,
		},
		"db_ulu_device_upgrade": map[string]interface{}{
			"tb_forced_version":  nil,
			"tb_special_version": nil,
		},
		"db_ulu_device_upgrade_status": map[string]interface{}{
			"tb_upgrade_status": nil,
		},
		"db_ulu_event": map[string]interface{}{
			"tb_event":          nil,
			"tb_event_detail":   nil,
			"tb_event_property": nil,
		},
		"db_ulu_event_comment": map[string]interface{}{
			"tb_event_comment": nil,
		},
		"db_ulu_event_handle": map[string]interface{}{
			"tb_event_handle":    nil,
			"tb_event_user_link": nil,
		},
		"db_ulu_event_pic": map[string]interface{}{
			"tb_event_pic": nil,
		},
		"db_ulu_examine": map[string]interface{}{
			"tb_examine":        nil,
			"tb_examine_detail": nil,
		},
		"db_ulu_examine_pic": map[string]interface{}{
			"tb_examine_pic": nil,
		},
		"db_ulu_examine_template": map[string]interface{}{
			"tb_examine_template": nil,
		},
		"db_ulu_keliu_counter": map[string]interface{}{
			"tb_keliu_business_dayentry_count":  nil,
			"tb_keliu_business_daypass_count":   nil,
			"tb_keliu_business_hourentry_count": nil,
			"tb_keliu_business_hourpass_count":  nil,
			"tb_keliu_store_dayentry_count":     nil,
			"tb_keliu_store_daypass_count":      nil,
			"tb_keliu_store_hourentry_count":    nil,
			"tb_keliu_store_hourpass_count":     nil,
		},
		"db_ulu_keliu_origin": map[string]interface{}{
			"tb_entry_device": nil,
			"tb_pass_device":  nil,
		},
		"db_ulu_mapped": map[string]interface{}{
			"tb_user_device": nil,
		},
		"db_ulu_message": map[string]interface{}{
			"tb_user_message": nil,
		},
		"db_ulu_message_client": map[string]interface{}{
			"tb_user_client": nil,
		},
		"db_ulu_offline": map[string]interface{}{
			"tb_offline_log": nil,
		},
		"db_ulu_pic": map[string]interface{}{
			"tb_pic": nil,
		},
		"db_ulu_pos_trade": map[string]interface{}{
			"tb_pos_trade":                   nil,
			"tb_pos_trade_commodity":         nil,
			"tb_pos_trade_promotion":         nil,
			"tb_pos_trade_settlement":        nil,
			"tb_user_store_channel_pos_bind": nil,
		},
		"db_ulu_posdevice": map[string]interface{}{
			"tb_posdevice": nil,
		},
		"db_ulu_posdevice_set": map[string]interface{}{
			"tb_posdevice_set": nil,
		},
		"db_ulu_roles_authority": map[string]interface{}{
			"tb_roles":             nil,
			"tb_roles_control":     nil,
			"tb_roles_relation":    nil,
			"tb_store_deleted_log": nil,
			"tb_user_group":        nil,
			"tb_user_group_log":    nil,
			"tb_user_module":       nil,
		},
		"db_ulu_stats": map[string]interface{}{
			"tb_user_act": nil,
		},
		"db_ulu_store": map[string]interface{}{
			"tb_common_property":        nil,
			"tb_device_digital":         nil,
			"tb_device_digital_mapping": nil,
			"tb_group":                  nil,
			"tb_store_detailed":         nil,
			"tb_store_property":         nil,
			"tb_store_property_temp":    nil,
		},
		"db_ulu_store_inspect_pic": map[string]interface{}{
			"tb_plan_intelligent_pic": nil,
			"tb_plan_pic":             nil,
		},
		"db_ulu_store_inspect_plan": map[string]interface{}{
			"tb_inspect_plan":            nil,
			"tb_inspect_plan_authority":  nil,
			"tb_inspect_plan_detail":     nil,
			"tb_plan_intelligent_device": nil,
		},
		"db_ulu_store_motion_pic": map[string]interface{}{
			"tb_motion_pic": nil,
		},
		"db_ulu_store_motion_plan": map[string]interface{}{
			"tb_motion_notice_link": nil,
			"tb_motion_plan":        nil,
		},
		"db_ulu_store_tuogang_pic": map[string]interface{}{
			"tb_tuogang_pic": nil,
		},
		"db_ulu_store_tuogang_plan": map[string]interface{}{
			"tb_tuogang_notice_link": nil,
			"tb_tuogang_plan":        nil,
			"tb_tuogang_plan_device": nil,
		},
		"db_ulu_user": map[string]interface{}{
			"tb_user":                 nil,
			"tb_user_printscreen_pic": nil,
			"tb_user_role":            nil,
		},
		"db_ulu_user_collect": map[string]interface{}{
			"tb_device_collect": nil,
			"tb_stroe_collect":  nil,
		},
		"db_ulu_widget": map[string]interface{}{
			"tb_widget_function": nil,
		},
	}
	return type1
}

func GetType2() Type2 {
	type2 := map[string]interface{}{
		"db_ulu_device_sn": map[string]interface{}{
			"tb_device_sn": nil,
		},
	}
	return type2
}

func GetType3() Type3 {
	type3 := map[string]interface{}{
		"db_ulu_mapped_myisam": map[string]interface{}{
			"tb_3rddid_autoid": nil,
			"tb_posdid_autoid": nil,
		},
	}
	return type3
}

func GetType4() Type4 {
	type4 := map[string]interface{}{
		"db_ulu_op_autoid": map[string]interface{}{
			"tb_appkey_autodid": nil,
		},
		"db_ulu_domain_autoid_myisam": map[string]interface{}{
			"tb_b_domain": nil,
		},
	}
	return type4
}

func GetType5() Type5 {
	type5 := map[string]interface{}{
		"db_ulu_mapped_myisam": map[string]interface{}{
			"tb_did_autoid": nil,
		},
	}
	return type5
}
