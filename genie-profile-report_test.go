package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestExtractDataFromProfile(t *testing.T) {
	exampleProfile := `<Genie>
  <Profile Account="WIZARDMAN" Password="kjsldfjldjfldf" Character="Wizzy" Game="DRF">
    <Layout FileName="default.layout" />
  </Profile>
</Genie>`
	expectedCD := charData{
		acct: "WIZARDMAN",
		char: "Wizzy",
		game: "DRF",
	}
	cd, err := extractProfileData(exampleProfile)
	assert.Equal(t, nil, err)
	assert.True(t, reflect.DeepEqual(expectedCD, cd))
}

func TestExtractDataFromProfileErr(t *testing.T) {
	exampleProfile := `some malformed file`
	expectedCD := charData{}
	cd, err := extractProfileData(exampleProfile)
	assert.True(t, err != nil)
	assert.True(t, reflect.DeepEqual(expectedCD, cd))
}