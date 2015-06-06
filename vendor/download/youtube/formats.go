package youtube

type Format struct{
	Code string	
	Info FormatInfo
}

type FormatInfo struct{
	Ext string
	Width uint16
	Height uint16
	Note string
	Preference int16
	Fps int16
	Vcodec string
	Container string
	Acodec string
	Abr int16
	Protocol string
}