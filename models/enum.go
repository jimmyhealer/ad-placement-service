package models

type GenderType string

const (
	Male   GenderType = "M"
	Female GenderType = "F"
)

type PlatformType string

const (
	Android PlatformType = "android"
	IOS     PlatformType = "ios"
	Web     PlatformType = "web"
)

type CountryCode string

const (
	TW CountryCode = "TW"
	JP CountryCode = "JP"
	US CountryCode = "US"
)
