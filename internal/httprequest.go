package internal

import (
	"bytes"
	"encoding/json"
	"log"
	"net/url"

	"github.com/antonholmquist/jason"
)

var errorStr = map[string]string{
	"INVALID_DATA":      "ошибка валидации переданных данных",
	"LIMIT_ERROR":       "Превышен лимит запросов",
	"METHOD_FAILED":     "внутренняя ошибка при выполнении метода",
	"AUTH_ERROR":        "ошибка авторизации",
	"INCORRECT_REQUEST": "Некорректный запрос к API",
	"NO_SUCH_METHOD":    "указанного метода не существует",
}

func init() {
	_ = errorStr
}

type status struct {
	Status     string
	Error_code string
	Error_text string
}
type subDomain struct {
	Result []struct {
		Domain_id int64
		Fqdn      string
		Id        int64
	}
}
type accountInfo struct {
	Result struct {
		User_domains int64 // Фактическое кол-во доменов
		Plan_domain  int64 // Максимальное кол-во доменов
	}
}
type resSubDomain struct {
	status
	Answer struct {
		subDomain
		status
	}
}
type resAccountInfo struct {
	status
	Answer struct {
		accountInfo
		status
	}
}
type resDnsGetData struct {
	status
	Answer struct {
		dnsData
		status
	}
}
type inputData map[string]interface{}

type requestParam struct {
	url           string
	login         string
	passwd        string
	input_format  string
	output_format string
}

func getUrlTemplate(urlStart string) requestParam {
	return requestParam{
		url:           urlStart,
		login:         Pass.login,
		passwd:        Pass.password,
		input_format:  "json",
		output_format: "json",
	}
}

func getUrlRequest(urlStart string, iData inputData) string {
	r := getUrlTemplate(urlStart)
	var inputData string
	if iData != nil {
		if iDat, err := json.Marshal(iData); err != nil {
			log.Println("Error getUrlRequest: ", err.Error())
			return ""
		} else {
			inputData = url.QueryEscape(string(iDat))
		}
	}
	if iData != nil {
		str := r.url + "?login=" + r.login + "&passwd=" + r.passwd + "&input_format=json&output_format=json&input_data=" + inputData
		return str
	} else {
		str := r.url + "?login=" + r.login + "&passwd=" + r.passwd + "&output_format=json"
		return str
	}
}

func GetDnsGetData(domainName string) {
	//resreq := resDnsGetData{}
	// idata := inputData{
	// 	"fqdn": "_acme-challenge." + domainName,
	// }
	idata := inputData{
		"fqdn": domainName,
	}
	res := getUrlRequest("https://api.beget.com/api/dns/getData", idata)
	if res1, err := sendRequest(res); err != nil {
		log.Println(err.Error())
	} else {

		log.Println("----------- getDnsGetData ------------------")
		log.Println(res1)
		printPrettier(res1)
	}
}

// Информация о пользователе
func GetAccountInfoReq() {
	//req := getUrlRequest("https://api.beget.com/api/user/getAccountInfo", nil)
	//resreq := resAccountInfo{}
	req := getUrlRequest("https://api.beget.com/api/user/getAccountInfo", nil)
	if res, err := sendRequest(req); err != nil {
		log.Println(err.Error())
		return
	} else {
		log.Println("----------- getAccountInfoReq ------------------")
		printPrettier2(res)
	}
}

// информация о DNS  _acme-challenge
func RequestData() {
	//getDnsGetData("arkadii.ru")

	// if errorCode, ok := result["error_code"]; ok {
	// 	log.Println(errorStr[errorCode.(string)])
	// }
}

// вывод на консоль
func printPrettier2(result interface{}) {
	// красивый вывод на экран

	if res, err := json.MarshalIndent(result, "", "   "); err != nil {
		log.Println("Error-5", err.Error())
	} else {
		log.Println(string(res))
	}
}
func printPrettier(result *jason.Object) {
	// красивый вывод на экран
	strRes := result.String()
	prettier := &bytes.Buffer{}
	if err := json.Indent(prettier, []byte(strRes), "", "   "); err != nil { // MarshalIndent(www, "", "   "); err != nil {
		log.Println("Error-50", err.Error())
	} else {
		log.Println("[*]>", prettier)
	}
}

func GetSubDomain() {
	// https://api.beget.com/api/domain/getSubdomainList?login=userlogin&passwd=password&output_format=json
	resreq := resSubDomain{}
	req := getUrlRequest("https://api.beget.com/api/domain/getSubdomainList", nil)
	if _, err := sendRequest(req); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("----------- GetSubDomain ------------------")
	printPrettier2(resreq)

}

//https://api.beget.com/api/dns/changeRecords?login=userlogin&passwd=password&input_format=json&output_format=json&input_data={"fqdn":"beget.ru","records":{"TXT":[{"priority":10,"value":"TXT record"}]}}

func SetDnsTxtData(domainName string) {
	resreq := resDnsGetData{}
	thisMap := make(map[string][]map[string]interface{})
	aMap := map[string]interface{}{
		"priority": 10,
		"value":    "TXT re 12345 T record 777 unicode",
	}
	thisMap["TXT"] = append(thisMap["Txt"], aMap)

	idata := inputData{
		"fqdn":    "anisoftware.ru",
		"records": thisMap,
	}

	res := getUrlRequest("https://api.beget.com/api/dns/changeRecords", idata)
	if _, err := sendRequest(res); err != nil {
		log.Println("!>", err.Error())
	}

	log.Println("----------- SetDnsTxtData ------------------")
	printPrettier2(resreq)
}
