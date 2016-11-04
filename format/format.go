package format
import (
	"../json"
	"regexp"
	"fmt"
)

func templateStringFormatByMap(template string,m map[string]interface{}) (string){
	
	
	for k, _ := range m {
		reg := regexp.MustCompile("{{" + k + "}}")
		s , ok := m[k].(string)
		if ok == false {
			fmt.Println("error", k, m[k], ok)
			continue
		}
		template = reg.ReplaceAllString(template, s)
	}
	return template;
}

func templateStringFormatByJSON(template string , j *json.JSON) (string){
	m , err := j.Map()
	if m == nil  || err != nil {
		return template
	}
	for k, _ := range m {
		reg := regexp.MustCompile("{{" + k + "}}")
		o := j.Get(k)
		if o == nil {
			continue
		}
		s, err := o.String()
		if err != nil {
			continue
		}
		template = reg.ReplaceAllString(template, s)
	}
	return template
}
