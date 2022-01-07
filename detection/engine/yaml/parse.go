package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Rules map[string][]Rule

type Rule struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Condition   string `yaml:"condition"`
	Priority    int    `yaml:"priority"`
}

func ReadK8sRulesFromYml(filePath string) (rules Rules) {
	rules = Rules{}
	rulesContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	err = yaml.Unmarshal(rulesContent, rules)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return rules

}
