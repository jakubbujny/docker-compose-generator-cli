package yml_operator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAppendAtEnd(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel.services.somewhere"
	source := "topLevel:\n   services:\n      somewhere:"
	toAppend :="toAppend:\n   someYml: file"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("topLevel:\n   services:\n      somewhere:\n         toAppend:\n            someYml: file\n", output)
}


func TestAppendWhenSomeBlockDeeper(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel.services.somewhere"
	source := "topLevel:\n   services:\n      somewhere:\n         someDeeper:"
	toAppend :="toAppend:\n   someYml: file"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("topLevel:\n   services:\n      somewhere:\n         toAppend:\n            someYml: file\n         someDeeper:\n", output)
}