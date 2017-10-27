package ietflt

// ISO639-1
func GetISO6391(lang string) string {
	return iso6391s[lang]
}

// Language Tags
func GetLangTags(langType string) []string {
	return langTags[langType]
}

// ISO3166-1 alpha-2
func GetGetISO31661Alpha2(lang string) string {
	return regions[lang]
}
