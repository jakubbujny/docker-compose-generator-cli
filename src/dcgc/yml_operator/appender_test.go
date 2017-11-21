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

func TestAppendWhenPathTopLevel(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel"
	source := "topLevel:\n   services:\n      somewhere:"
	toAppend :="toAppend:\n   someYml: file"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("topLevel:\n   toAppend:\n      someYml: file\n   services:\n      somewhere:\n", output)
}

func TestAppendWhenOnlyOneElement(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel"
	source := "topLevel:"
	toAppend :="toAppend:\n   someYml: file"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("topLevel:\n   toAppend:\n      someYml: file\n", output)
}

func TestAppendWhenIsVersion(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel"
	source := "version:'3'\ntopLevel:"
	toAppend :="toAppend:\n   someYml: file"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("version:'3'\ntopLevel:\n   toAppend:\n      someYml: file\n", output)
}

func TestAppendWhenDoubleSameName(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "replace"
	source := "something:\n   deeper:\n      replace:\nreplace:"
	toAppend :="toAppend:"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("something:\n   deeper:\n      replace:\nreplace:\n   toAppend:\n", output)
}