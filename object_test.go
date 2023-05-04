package goattribute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type inputTestStruct struct {
	Std  int8   `json:"std"`
	Name string `json:"name"`
}

type outputTestStruct struct {
	Filename string `json:"filename"`
}

type configTestStruct struct {
	Name   string             `json:"name"`
	Input  *inputTestStruct   `json:"input"`
	Output []outputTestStruct `json:"output"`
}

func TestSetAttr(t *testing.T) {
	var config = &configTestStruct{
		Name: "example1",
		Input: &inputTestStruct{
			Std:  1,
			Name: "example2",
		},
		Output: []outputTestStruct{
			{
				Filename: "example3",
			},
			{
				Filename: "example4",
			},
		},
	}

	assert.Nil(t, New(config).SetAttr("Name", "test1"))
	assert.Equal(t, "test1", config.Name)

	assert.Nil(t, New(config).SetAttr("Input.Name", "test2"))
	assert.Equal(t, "test2", config.Input.Name)

	assert.Nil(t, New(config).SetAttr("Input.Std", 2))
	assert.Equal(t, int8(2), config.Input.Std)

	assert.Nil(t, New(config).SetAttr("Output[0].Filename", "test3"))
	assert.Equal(t, "test3", config.Output[0].Filename)

	assert.Nil(t, New(config).SetAttr("Output[1].Filename", "test4"))
	assert.Equal(t, "test4", config.Output[1].Filename)

	assert.NotNil(t, New(config).SetAttr("Input", "test"))
	assert.NotNil(t, New(config).SetAttr("Output[2].Filename", "test5"))
	assert.NotNil(t, New(config).SetAttr("NotExists", "123123"))
	assert.NotNil(t, New(config).SetAttr("Input.NotExists", "123123"))
	assert.NotNil(t, New(config).SetAttr("NotExists.Test", "123123"))
}

func TestGetAttr(t *testing.T) {
	var config = &configTestStruct{
		Name: "example1",
		Input: &inputTestStruct{
			Std:  1,
			Name: "example2",
		},
		Output: []outputTestStruct{
			{
				Filename: "example3",
			},
			{
				Filename: "example4",
			},
		},
	}

	var attr = New(config)
	name, err := attr.GetAttr("Name")
	assert.NoError(t, err)
	assert.Equal(t, "example1", name)

	std, err := attr.GetAttr("Input.Std")
	assert.NoError(t, err)
	assert.Equal(t, int8(1), std)

	name, err = attr.GetAttr("Input.Name")
	assert.NoError(t, err)
	assert.Equal(t, "example2", name)

	filename, err := attr.GetAttr("Output[0].Filename")
	assert.NoError(t, err)
	assert.Equal(t, "example3", filename)

	filename, err = attr.GetAttr("Output[1].Filename")
	assert.NoError(t, err)
	assert.Equal(t, "example4", filename)

	_, err = attr.GetAttr("NotExists")
	assert.NotNil(t, err)

	_, err = attr.GetAttr("NotExists.Test")
	assert.NotNil(t, err)

	_, err = attr.GetAttr("Input.NotExists")
	assert.NotNil(t, err)

	_, err = attr.GetAttr("Test.NotExists")
	assert.NotNil(t, err)

	_, err = attr.GetAttr("Output[3].Filename")
	assert.NotNil(t, err)
}

func TestNilSet(t *testing.T) {
	var config configTestStruct
	var attr = New(&config)
	assert.NoError(t, attr.SetAttr("Name", "hello"))
	assert.NotNil(t, attr.SetAttr("Input.Name", "world"))
	assert.NotNil(t, attr.SetAttr("Input.Std", 2))
	assert.NotNil(t, attr.SetAttr("Output[0].Filename", "test.txt"))

	config.Input = &inputTestStruct{}
	assert.NoError(t, attr.SetAttr("Input.Name", "world"))
	assert.Equal(t, "world", config.Input.Name)
	name, err := attr.GetAttr("Input.Name")
	assert.NoError(t, err)
	assert.Equal(t, "world", name)
}
