package goattribute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type inputTestStruct struct {
	Std  int8   `json:"std" yaml:"std"`
	Name string `json:"name" yaml:"name"`
}

type outputTestStruct struct {
	Filename string `json:"filename" yaml:"filename"`
}

type userTestStruct struct {
	Username string `json:"username" yaml:"username"`
	Age      int    `json:"age" yaml:"age"`
}

type configTestStruct struct {
	Name       string             `json:"name" yaml:"name"`
	Input      *inputTestStruct   `json:"input"`
	Output     []outputTestStruct `json:"output" yaml:"output"`
	LinkedUser userTestStruct     `json:"linked_user" yaml:"linked_user"`

	private string
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
		LinkedUser: userTestStruct{
			Username: "example",
			Age:      29,
		},
		private: "example5",
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

	assert.Nil(t, New(config).SetAttr("LinkedUser.Username", "test5"))
	assert.Equal(t, "test5", config.LinkedUser.Username)

	assert.Nil(t, New(config).SetAttr("LinkedUser.Age", 33))
	assert.Equal(t, 33, config.LinkedUser.Age)

	// private field cannot be written
	assert.NotNil(t, New(config).SetAttr("private", "test6"))

	assert.NotNil(t, New(config).SetAttr("Input", "test"))
	assert.NotNil(t, New(config).SetAttr("Output[2].Filename", "test5"))
	assert.NotNil(t, New(config).SetAttr("NotExists", "123123"))
	assert.NotNil(t, New(config).SetAttr("Input.NotExists", "123123"))
	assert.NotNil(t, New(config).SetAttr("NotExists.Test", "123123"))
}

func TestSetAttrWithTag(t *testing.T) {
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
		LinkedUser: userTestStruct{
			Username: "example",
			Age:      29,
		},
	}

	assert.Nil(t, NewWithTag(config, "json").SetAttr("name", "test1"))
	assert.Equal(t, "test1", config.Name)

	assert.Nil(t, NewWithTag(config, "json").SetAttr("input.name", "test2"))
	assert.Equal(t, "test2", config.Input.Name)

	assert.Nil(t, NewWithTag(config, "json").SetAttr("input.std", 2))
	assert.Equal(t, int8(2), config.Input.Std)

	assert.Nil(t, NewWithTag(config, "json").SetAttr("output[0].filename", "test3"))
	assert.Equal(t, "test3", config.Output[0].Filename)

	assert.Nil(t, NewWithTag(config, "json").SetAttr("output[1].filename", "test4"))
	assert.Equal(t, "test4", config.Output[1].Filename)

	assert.Nil(t, NewWithTag(config, "json").SetAttr("linked_user.username", "test5"))
	assert.Equal(t, "test5", config.LinkedUser.Username)

	assert.Nil(t, NewWithTag(config, "json").SetAttr("linked_user.age", 33))
	assert.Equal(t, 33, config.LinkedUser.Age)

	assert.NotNil(t, NewWithTag(config, "json").SetAttr("input", "test"))
	assert.NotNil(t, NewWithTag(config, "json").SetAttr("output[2].filename", "test5"))
	assert.NotNil(t, NewWithTag(config, "json").SetAttr("not_exists", "123123"))
	assert.NotNil(t, NewWithTag(config, "json").SetAttr("input.not_exists", "123123"))
	assert.NotNil(t, NewWithTag(config, "json").SetAttr("not_exists.test", "123123"))

	assert.Nil(t, NewWithTag(config, "yaml").SetAttr("output[0].filename", "test6"))
	assert.Equal(t, "test6", config.Output[0].Filename)

	assert.NotNil(t, NewWithTag(config, "none").SetAttr("name", "nononono"))
	assert.NotNil(t, NewWithTag(config, "none").SetAttr("output[0].filename", "nononono2"))

	assert.NotNil(t, NewWithTag(config, "yaml").SetAttr("input.std", 3))
	assert.Nil(t, NewWithTag(config, "yaml").SetAttr("Input.std", 4))
	assert.Equal(t, int8(4), config.Input.Std)
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
		LinkedUser: userTestStruct{
			Username: "example5",
			Age:      38,
		},
		private: "example6",
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

	username, err := attr.GetAttr("LinkedUser.Username")
	assert.NoError(t, err)
	assert.Equal(t, "example5", username)

	age, err := attr.GetAttr("LinkedUser.Age")
	assert.NoError(t, err)
	assert.Equal(t, 38, age)

	_, err = attr.GetAttr("private")
	assert.NotNil(t, err)

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

func TestGetAttrWithTag(t *testing.T) {
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
		LinkedUser: userTestStruct{
			Username: "example5",
			Age:      38,
		},
	}

	var attr = NewWithTag(config, "json")
	name, err := attr.GetAttr("name")
	assert.NoError(t, err)
	assert.Equal(t, "example1", name)

	std, err := attr.GetAttr("input.std")
	assert.NoError(t, err)
	assert.Equal(t, int8(1), std)

	name, err = attr.GetAttr("input.name")
	assert.NoError(t, err)
	assert.Equal(t, "example2", name)

	filename, err := attr.GetAttr("output[0].filename")
	assert.NoError(t, err)
	assert.Equal(t, "example3", filename)

	filename, err = attr.GetAttr("output[1].filename")
	assert.NoError(t, err)
	assert.Equal(t, "example4", filename)

	username, err := attr.GetAttr("linked_user.username")
	assert.NoError(t, err)
	assert.Equal(t, "example5", username)

	age, err := attr.GetAttr("linked_user.age")
	assert.NoError(t, err)
	assert.Equal(t, 38, age)

	_, err = attr.GetAttr("not_exists")
	assert.NotNil(t, err)

	_, err = attr.GetAttr("not_exists.test")
	assert.NotNil(t, err)

	_, err = attr.GetAttr("input.not_exists")
	assert.NotNil(t, err)

	_, err = attr.GetAttr("test.not_exists")
	assert.NotNil(t, err)

	_, err = attr.GetAttr("output[3].filename")
	assert.NotNil(t, err)
}

func TestNilSet(t *testing.T) {
	var config configTestStruct
	var attr = New(&config)
	assert.NoError(t, attr.SetAttr("Name", "hello"))
	assert.NoError(t, attr.SetAttr("LinkedUser.Username", "test2"))
	assert.NoError(t, attr.SetAttr("LinkedUser.Age", 45))
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

func TestNilSetWithTag(t *testing.T) {
	var config configTestStruct
	var attr = NewWithTag(&config, "json")
	assert.NoError(t, attr.SetAttr("name", "hello"))
	assert.NoError(t, attr.SetAttr("linked_user.username", "test2"))
	assert.NoError(t, attr.SetAttr("linked_user.age", 45))
	assert.NotNil(t, attr.SetAttr("input.name", "world"))
	assert.NotNil(t, attr.SetAttr("input.std", 2))
	assert.NotNil(t, attr.SetAttr("output[0].filename", "test.txt"))

	config.Input = &inputTestStruct{}
	assert.NoError(t, attr.SetAttr("input.name", "world"))
	assert.Equal(t, "world", config.Input.Name)
	name, err := attr.GetAttr("input.name")
	assert.NoError(t, err)
	assert.Equal(t, "world", name)
}
