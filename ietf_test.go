package ietflt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetISO6391(t *testing.T) {
	assert := assert.New(t)

	code := GetISO6391("af-NA")
	assert.Equal(code, "af")
}

func TestGetLangTagExtensions(t *testing.T) {
	assert := assert.New(t)

	tags := GetLangTags("af")
	assert.Equal(tags, []string{"af", "af-NA", "af-ZA"}, "not equal")
}

func TestGetISO3166Alpha2(t *testing.T) {
	assert := assert.New(t)

	territory := GetGetISO31661Alpha2("az-Cyrl-AZ")
	assert.Equal(territory, "AZ", "not equal")

	territory = GetGetISO31661Alpha2("ca-ES-VALENCIA")
	assert.Equal(territory, "ES", "not equal")
}
