package yml_operator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestAppendSimple(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel"
	source := "topLevel:\n  services:\n    somewhere: test"
	toAppend :="toAppend: something\n"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("topLevel:\n  services:\n    somewhere: test\n  toAppend: something\n", output)
}

func TestAppendAtEnd(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel.services"
	source := "topLevel:\n  services:\n    somewhere: test"
	toAppend :="toAppend:\n  someYml: file"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("topLevel:\n  services:\n    somewhere: test\n    toAppend:\n      someYml: file\n", output)
}


func TestAppendWhenSomeBlockDeeper(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel.services.somewhere"
	source := "topLevel:\n  services:\n    somewhere:\n      someDeeper: test"
	toAppend :="toAppend:\n  someYml: file"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("topLevel:\n  services:\n    somewhere:\n      someDeeper: test\n      toAppend:\n        someYml: file\n", output)
}

func TestAppendWhenPathTopLevel(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel"
	source := "topLevel:\n  services:\n    somewhere: test"
	toAppend :="toAppend:\n  someYml: file"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("topLevel:\n  services:\n    somewhere: test\n  toAppend:\n    someYml: file\n", output)
}

func TestAppendWhenIsVersion(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "topLevel"
	source := "version: '3'\ntopLevel:\n  is: string"
	toAppend :="toAppend:\n   someYml: file"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("topLevel:\n  is: string\n  toAppend:\n    someYml: file\nversion: \"3\"\n", output)
}

func TestAppendWhenDoubleSameName(t *testing.T) {
	assert := assert.New(t)

	//given
	path := "replace"
	source := "something:\n  deeper:\n    replace:\n      here: test\nreplace:\n  something: test"
	toAppend :="toAppend: test"
	//when
	output,err := AppendToYmlInSection(toAppend, source, path)
	//then

	assert.Nil(err)
	assert.Equal("something:\n  deeper:\n    replace:\n      here: test\nreplace:\n  something: test\n  toAppend: test\n", output)
}