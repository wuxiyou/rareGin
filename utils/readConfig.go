package utils

import "goini"

func ReadConfigInfo(sect, name string) string {
	configInfo := goini.SetConfig("./config/config.ini")

	info := configInfo.GetValue(sect, name)

	return info
}