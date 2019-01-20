package parser

import (
	"learnGo/crawler/engine"
	"learnGo/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(\d+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(\d+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(\d+)kg</div>`)
var incomeRe = regexp.MustCompile(`(\d+-\d+元)`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\p{Han}]+)\(\d+.\d+-\d+.\d+\)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(离异|未婚|丧偶)</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(大[^<]+|高[^<]+)</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+/[^<]+)</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>工作地:([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^<]+房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^<]+车)</div>`)
//var guessRe = regexp.MustCompile(``)
//var idUrlRe = regexp.MustCompile(``)

//传入个人信息页面的URL,解析个人信息
func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	//处理各个字段
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Marriage = extractString(contents, marriageRe)

	profile.Income = extractString(contents, incomeRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []interface{}{profile}, //因为里面的Items只接受interface,所以这里转一下
	}

	return result
}

//处理拿到的信息
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}
