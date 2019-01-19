package parser

import (
	"io/ioutil"
	"learnGo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents,"心痛的感觉")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element;but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	profile.Name = "心痛的感觉"
	expected := model.Profile{
		Age:        47,
		Height:     157,
		Weight:     55,
		Income:     "3001-5000元",
		Name:       "心痛的感觉",
		Xinzuo:     "射手座",
		Occupation: "教育/科研",
		Marriage:   "离异",
		House:      "已购房",
		Hokou:      "阿坝汶川",
		Education:  "大学本科",
		Car:        "未买车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
