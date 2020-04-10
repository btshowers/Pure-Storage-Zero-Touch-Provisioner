//Brandon Showers
//March 22 2020
//v1

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"gopkg.in/go-playground/validator.v9"
)

//FLASH ARRAY VARS//
var mainwin *ui.Window
var ipAddress = ""

//END FLASH ARRAY VARS//

//FLASH BLADE VARS//
//Global Vars//
var ipAddressFB = ""
var xAuthToken = ""
var loginUrl = ""
var apiUrl = ""

//END FLASH BLADE VARS//

//FLASH ARRAY FUNCTIONS//
func timeZones() []string {
	tz := []string{"Africa/Abidjan", "Africa/Accra", "Africa/Addis_Ababa", "Africa/Algiers", "Africa/Asmara", "Africa/Bamako", "Africa/Bangui", "Africa/Banjul", "Africa/Bissau", "Africa/Blantyre", "Africa/Brazzaville", "Africa/Bujumbura", "Africa/Cairo", "Africa/Casablanca", "Africa/Ceuta", "Africa/Conakry", "Africa/Dakar", "Africa/Dar_es_Salaam", "Africa/Djibouti", "Africa/Douala", "Africa/El_Aaiun", "Africa/Freetown", "Africa/Gaborone", "Africa/Harare", "Africa/Johannesburg", "Africa/Juba", "Africa/Kampala", "Africa/Khartoum", "Africa/Kigali", "Africa/Kinshasa", "Africa/Lagos", "Africa/Libreville", "Africa/Lome", "Africa/Luanda", "Africa/Lubumbashi", "Africa/Lusaka", "Africa/Malabo", "Africa/Maputo", "Africa/Maseru", "Africa/Mbabane", "Africa/Mogadishu", "Africa/Monrovia", "Africa/Nairobi", "Africa/Ndjamena", "Africa/Niamey", "Africa/Nouakchott", "Africa/Ouagadougou", "Africa/Porto-Novo", "Africa/Sao_Tome", "Africa/Tripoli", "Africa/Tunis", "Africa/Windhoek", "America/Adak", "America/Anchorage", "America/Anguilla", "America/Antigua", "America/Araguaina", "America/Argentina/Buenos_Aires", "America/Argentina/Catamarca", "America/Argentina/Cordoba", "America/Argentina/Jujuy", "America/Argentina/La_Rioja", "America/Argentina/Mendoza", "America/Argentina/Rio_Gallegos", "America/Argentina/Salta", "America/Argentina/San_Juan", "America/Argentina/San_Luis", "America/Argentina/Tucuman", "America/Argentina/Ushuaia", "America/Aruba", "America/Asuncion", "America/Atikokan", "America/Bahia", "America/Bahia_Banderas", "America/Barbados", "America/Belem", "America/Belize", "America/Blanc-Sablon", "America/Boa_Vista", "America/Bogota", "America/Boise", "America/Cambridge_Bay", "America/Campo_Grande", "America/Cancun", "America/Caracas", "America/Cayenne", "America/Cayman", "America/Chicago", "America/Chihuahua", "America/Costa_Rica", "America/Creston", "America/Cuiaba", "America/Curacao", "America/Danmarkshavn", "America/Dawson", "America/Dawson_Creek", "America/Denver", "America/Detroit", "America/Dominica", "America/Edmonton", "America/Eirunepe", "America/El_Salvador", "America/Fort_Nelson", "America/Fortaleza", "America/Glace_Bay", "America/Godthab", "America/Goose_Bay", "America/Grand_Turk", "America/Grenada", "America/Guadeloupe", "America/Guatemala", "America/Guayaquil", "America/Guyana", "America/Halifax", "America/Havana", "America/Hermosillo", "America/Indiana/Indianapolis", "America/Indiana/Knox", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Tell_City", "America/Indiana/Vevay", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Inuvik", "America/Iqaluit", "America/Jamaica", "America/Juneau", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Kralendijk", "America/La_Paz", "America/Lima", "America/Los_Angeles", "America/Lower_Princes", "America/Maceio", "America/Managua", "America/Manaus", "America/Marigot", "America/Martinique", "America/Matamoros", "America/Mazatlan", "America/Menominee", "America/Merida", "America/Metlakatla", "America/Mexico_City", "America/Miquelon", "America/Moncton", "America/Monterrey", "America/Montevideo", "America/Montserrat", "America/Nassau", "America/New_York", "America/Nipigon", "America/Nome", "America/Noronha", "America/North_Dakota/Beulah", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/Ojinaga", "America/Panama", "America/Pangnirtung", "America/Paramaribo", "America/Phoenix", "America/Port_of_Spain", "America/Port-au-Prince", "America/Porto_Velho", "America/Puerto_Rico", "America/Punta_Arenas", "America/Rainy_River", "America/Rankin_Inlet", "America/Recife", "America/Regina", "America/Resolute", "America/Rio_Branco", "America/Santarem", "America/Santiago", "America/Santo_Domingo", "America/Sao_Paulo", "America/Scoresbysund", "America/Sitka", "America/St_Barthelemy", "America/St_Johns", "America/St_Kitts", "America/St_Lucia", "America/St_Thomas", "America/St_Vincent", "America/Swift_Current", "America/Tegucigalpa", "America/Thule", "America/Thunder_Bay", "America/Tijuana", "America/Toronto", "America/Tortola", "America/Vancouver", "America/Whitehorse", "America/Winnipeg", "America/Yakutat", "America/Yellowknife", "Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Macquarie", "Antarctica/Mawson", "Antarctica/McMurdo", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/Syowa", "Antarctica/Troll", "Antarctica/Vostok", "Arctic/Longyearbyen", "Asia/Aden", "Asia/Almaty", "Asia/Amman", "Asia/Anadyr", "Asia/Aqtau", "Asia/Aqtobe", "Asia/Ashgabat", "Asia/Atyrau", "Asia/Baghdad", "Asia/Bahrain", "Asia/Baku", "Asia/Bangkok", "Asia/Barnaul", "Asia/Beirut", "Asia/Bishkek", "Asia/Brunei", "Asia/Chita", "Asia/Choibalsan", "Asia/Colombo", "Asia/Damascus", "Asia/Dhaka", "Asia/Dili", "Asia/Dubai", "Asia/Dushanbe", "Asia/Famagusta", "Asia/Gaza", "Asia/Hebron", "Asia/Ho_Chi_Minh", "Asia/Hong_Kong", "Asia/Hovd", "Asia/Irkutsk", "Asia/Jakarta", "Asia/Jayapura", "Asia/Jerusalem", "Asia/Kabul", "Asia/Kamchatka", "Asia/Karachi", "Asia/Kathmandu", "Asia/Khandyga", "Asia/Kolkata", "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Kuching", "Asia/Kuwait", "Asia/Macau", "Asia/Magadan", "Asia/Makassar", "Asia/Manila", "Asia/Muscat", "Asia/Nicosia", "Asia/Novokuznetsk", "Asia/Novosibirsk", "Asia/Omsk", "Asia/Oral", "Asia/Phnom_Penh", "Asia/Pontianak", "Asia/Pyongyang", "Asia/Qatar", "Asia/Qostanay", "Asia/Qyzylorda", "Asia/Riyadh", "Asia/Sakhalin", "Asia/Samarkand", "Asia/Seoul", "Asia/Shanghai", "Asia/Singapore", "Asia/Srednekolymsk", "Asia/Taipei", "Asia/Tashkent", "Asia/Tbilisi", "Asia/Tehran", "Asia/Thimphu", "Asia/Tokyo", "Asia/Tomsk", "Asia/Ulaanbaatar", "Asia/Urumqi", "Asia/Ust-Nera", "Asia/Vientiane", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yangon", "Asia/Yekaterinburg", "Asia/Yerevan", "Atlantic/Azores", "Atlantic/Bermuda", "Atlantic/Canary", "Atlantic/Cape_Verde", "Atlantic/Faroe", "Atlantic/Madeira", "Atlantic/Reykjavik", "Atlantic/South_Georgia", "Atlantic/St_Helena", "Atlantic/Stanley", "Australia/Adelaide", "Australia/Brisbane", "Australia/Broken_Hill", "Australia/Currie", "Australia/Darwin", "Australia/Eucla", "Australia/Hobart", "Australia/Lindeman", "Australia/Lord_Howe", "Australia/Melbourne", "Australia/Perth", "Australia/Sydney", "Europe/Amsterdam", "Europe/Andorra", "Europe/Astrakhan", "Europe/Athens", "Europe/Belgrade", "Europe/Berlin", "Europe/Bratislava", "Europe/Brussels", "Europe/Bucharest", "Europe/Budapest", "Europe/Busingen", "Europe/Chisinau", "Europe/Copenhagen", "Europe/Dublin", "Europe/Gibraltar", "Europe/Guernsey", "Europe/Helsinki", "Europe/Isle_of_Man", "Europe/Istanbul", "Europe/Jersey", "Europe/Kaliningrad", "Europe/Kiev", "Europe/Kirov", "Europe/Lisbon", "Europe/Ljubljana", "Europe/London", "Europe/Luxembourg", "Europe/Madrid", "Europe/Malta", "Europe/Mariehamn", "Europe/Minsk", "Europe/Monaco", "Europe/Moscow", "Europe/Oslo", "Europe/Paris", "Europe/Podgorica", "Europe/Prague", "Europe/Riga", "Europe/Rome", "Europe/Samara", "Europe/San_Marino", "Europe/Sarajevo", "Europe/Saratov", "Europe/Simferopol", "Europe/Skopje", "Europe/Sofia", "Europe/Stockholm", "Europe/Tallinn", "Europe/Tirane", "Europe/Ulyanovsk", "Europe/Uzhgorod", "Europe/Vaduz", "Europe/Vatican", "Europe/Vienna", "Europe/Vilnius", "Europe/Volgograd", "Europe/Warsaw", "Europe/Zagreb", "Europe/Zaporozhye", "Europe/Zurich", "Indian/Antananarivo", "Indian/Chagos", "Indian/Christmas", "Indian/Cocos", "Indian/Comoro", "Indian/Kerguelen", "Indian/Mahe", "Indian/Maldives", "Indian/Mauritius", "Indian/Mayotte", "Indian/Reunion", "Pacific/Apia", "Pacific/Auckland", "Pacific/Bougainville", "Pacific/Chatham", "Pacific/Chuuk", "Pacific/Easter", "Pacific/Efate", "Pacific/Enderbury", "Pacific/Fakaofo", "Pacific/Fiji", "Pacific/Funafuti", "Pacific/Galapagos", "Pacific/Gambier", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Honolulu", "Pacific/Kiritimati", "Pacific/Kosrae", "Pacific/Kwajalein", "Pacific/Majuro", "Pacific/Marquesas", "Pacific/Midway", "Pacific/Nauru", "Pacific/Niue", "Pacific/Norfolk", "Pacific/Noumea", "Pacific/Pago_Pago", "Pacific/Palau", "Pacific/Pitcairn", "Pacific/Pohnpei", "Pacific/Port_Moresby", "Pacific/Rarotonga", "Pacific/Saipan", "Pacific/Tahiti", "Pacific/Tarawa", "Pacific/Tongatapu", "Pacific/Wake", "Pacific/Wallis"}
	return tz
}

func getAPICallFA(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	//log.Println(string(body))
	return string(body)
}

func patchRestCallFA(url string, data []byte) string {

	body := bytes.NewReader(data)

	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	req.Header.Set("Content-Type", "application/json")

	//make the rest call
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	//wait then close the connection to free space.
	defer resp.Body.Close()

	//converting http.response body to byte array so it can be cast to string to return. *sheesh
	respData, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2.Error())
		return err2.Error()
	}
	//finally return the response as a string back to the console app.
	return string(respData)
}

//END FLASH ARRAY FUNCTIONS//

//FLASH BLADE FUNCTIONS//
//TODO not sure if this method is usefull.. may just move to the login button function
func urlParser(url string) []string {
	//returns a slice broken out by forward slash in the url
	parts := strings.Split(url, "/")
	return parts
}

func getAPICallFB(url string, xAuthToken string) []byte {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte(err.Error())
	}
	req.Header.Set("x-auth-token", xAuthToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return []byte(err.Error())
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err2.Error())
		return []byte(err2.Error())
	}

	//log.Println(string(body))
	return body
}

//Post rest function specifically for logon only takes 2 parameters and returns a string//
func postAPICallLoginFB(url string, apiToken string) string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	req.Header.Set("api-token", apiToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	//Sets the x-auth-token from the header response
	if len(resp.Header["X-Auth-Token"]) > 0 {
		s := resp.Header["X-Auth-Token"]
		t := strings.Replace(s[0], "[", "", -1)
		t = strings.Replace(t, "]", "", -1)
		xAuthToken = t
	}

	return string(body)
}

//Post rest function takes 3 parameters and returns a string//
func postAPICall2FB(url string, xAuthToken string, data []byte) string {

	body := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-auth-token", xAuthToken)

	//make the rest call
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	//wait then close the connection to free space.
	defer resp.Body.Close()

	//converting http.response body to byte array so it can be cast to string to return. *sheesh
	respData, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2.Error())
		return err2.Error()
	}
	//finally return the response as a string back to the console app.
	return string(respData)
}

//patch rest function takes 3 parameters and returns a string response//
func patchAPICallFB(url string, xAuthToken string, data []byte) string {

	body := bytes.NewReader(data)

	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-auth-token", xAuthToken)

	//make the rest call
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	//wait then close the connection to free space.
	defer resp.Body.Close()

	//converting http.response body to byte array so it can be cast to string to return. *sheesh
	respData, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2.Error())
		return err2.Error()
	}
	//finally return the response as a string back to the console app.
	return string(respData)
}

//delete function takes 2 parameters and returns a string//
func deleteAPICallFB(url string, xAuthToken string) string {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	req.Header.Set("x-auth-token", xAuthToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	return string(body)
}

//END FLASH BLADE FUNCTIONS//

func initializeFATab() ui.Control {
	//fields for the form
	arrayName := ui.NewEntry()
	eulaOrg := ui.NewEntry()
	eulaName := ui.NewEntry()
	eulaTitle := ui.NewEntry()
	eulaAccept := ui.NewCheckbox("yes")
	ntpServer := ui.NewEntry()
	vir0IP := ui.NewEntry()
	vir0SNM := ui.NewEntry()
	vir0GW := ui.NewEntry()
	ct0IP := ui.NewEntry()
	ct0SNM := ui.NewEntry()
	ct0GW := ui.NewEntry()
	ct1IP := ui.NewEntry()
	ct1SNM := ui.NewEntry()
	ct1GW := ui.NewEntry()
	dnsDomain := ui.NewEntry()
	dnsServer := ui.NewEntry()
	smtpRelay := ui.NewEntry()
	smtpDomain := ui.NewEntry()
	smtpAlertEmail := ui.NewEntry()
	tempIP := ui.NewEntry() //dhcp ip address
	initResult := ui.NewMultilineEntry()
	timeZone := ui.NewCombobox()
	tz := timeZones()
	for i, v := range tz {
		timeZone.Append(v)
		i++
	}

	//first column definition
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	//define vertical box inside column similar to a div
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	//ARRAY NAME FIELD//
	//define the group for the form
	group1 := ui.NewGroup("General Configs")
	group1.SetMargined(true)

	//add group to the vertical box
	vbox.Append(group1, true)

	//define the form for the group
	entryForm1 := ui.NewForm()
	entryForm1.SetPadded(true)

	//embed the array name form field inside the first form group
	group1.SetChild(entryForm1)
	entryForm1.Append("FlashArray Name", arrayName, false)
	entryForm1.Append("", ui.NewLabel(""), false)
	entryForm1.Append("Organization Name", eulaOrg, false)
	entryForm1.Append("Your Name", eulaName, false)
	entryForm1.Append("Your Title", eulaTitle, false)
	entryForm1.Append("You accept EULA", eulaAccept, false)
	entryForm1.Append("", ui.NewLabel(""), false)
	entryForm1.Append("NTP Time Server(s)", ntpServer, false)
	entryForm1.Append("", ui.NewLabel("*Comma seperated for multiple entries"), false)
	entryForm1.Append("TimeZone", timeZone, false)
	entryForm1.Append("OPTIONAL FIELDS", ui.NewLabel("________________________________________"), false)
	entryForm1.Append("", ui.NewLabel(""), false)
	entryForm1.Append("DNS Domain", dnsDomain, false)
	entryForm1.Append("DNS Name Server(s)", dnsServer, false)
	entryForm1.Append("", ui.NewLabel("*Comma seperated for multiple entries"), false)
	entryForm1.Append("", ui.NewLabel(""), false)
	entryForm1.Append("SMTP Relay Host", smtpRelay, false)
	entryForm1.Append("SMTP sender domain", smtpDomain, false)
	entryForm1.Append("Alert Email Address(s)", smtpAlertEmail, false)
	entryForm1.Append("", ui.NewLabel("*Comma seperated for multiple entries"), false)

	//seperator line
	hbox.Append(ui.NewVerticalSeparator(), false)

	//Middle column
	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	//VIR0IP FORM//
	group3 := ui.NewGroup("Virtual Nic 0")
	group3.SetMargined(true)
	vbox.Append(group3, true)

	entryForm3 := ui.NewForm()
	entryForm3.SetPadded(true)
	group3.SetChild(entryForm3)

	//autofill button to copy contents to ct0 and ct1 ip configs
	button := ui.NewButton("Autofill")
	entryForm3.Append("IP Address", vir0IP, false)
	entryForm3.Append("Subnet Mask", vir0SNM, false)
	entryForm3.Append("Default Gateway", vir0GW, false)
	entryForm3.Append("Replicate below", button, false)

	//CT0 FORM//
	group5 := ui.NewGroup("Controller 0")
	group5.SetMargined(true)
	vbox.Append(group5, true)
	entryForm5 := ui.NewForm()
	entryForm5.SetPadded(true)
	group5.SetChild(entryForm5)

	entryForm5.Append("IP Address", ct0IP, false)
	entryForm5.Append("Subnet Mask", ct0SNM, false)
	entryForm5.Append("Default Gateway", ct0GW, false)

	//CT1 FORM//
	group6 := ui.NewGroup("Controller 1")
	group6.SetMargined(true)
	vbox.Append(group6, true)
	entryForm6 := ui.NewForm()
	entryForm6.SetPadded(true)
	group6.SetChild(entryForm6)

	entryForm6.Append("IP Address", ct1IP, false)
	entryForm6.Append("Subnet Mask", ct1SNM, false)
	entryForm6.Append("Default Gateway", ct1GW, false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	//third column
	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	//SUBMIT "GO" BUTTON//
	group9 := ui.NewGroup("Initialize Array")
	group9.SetMargined(true)
	vbox.Append(group9, true)

	entryForm9 := ui.NewForm()
	entryForm9.SetPadded(true)
	group9.SetChild(entryForm9)

	button1 := ui.NewButton("Query")
	entryForm9.Append("", ui.NewLabel(""), false)

	//submit and go button
	button2 := ui.NewButton("Initialize")

	entryForm9.Append("DHCP IP of Array ", tempIP, false)
	entryForm9.Append("Query First, ", button1, false)
	entryForm9.Append("Configure Array ", button2, false)

	//multiline field for showing results of patch api call and form validation messages.
	//sets the initResults console to readonly
	initResult.SetReadOnly(true)
	entryForm9.Append("Init Results", initResult, true)

	//autofill IP config button actions
	//used to replicate the ip info from vi0 to ct0 and ct1
	button.OnClicked(func(*ui.Button) {

		ct0IP.SetText(vir0IP.Text())
		ct0SNM.SetText(vir0SNM.Text())
		ct0GW.SetText(vir0GW.Text())
		ct1IP.SetText(vir0IP.Text())
		ct1SNM.SetText(vir0SNM.Text())
		ct1GW.SetText(vir0GW.Text())

	})

	button1.OnClicked(func(*ui.Button) {
		initResult.SetText("Processing please wait...")
		ipAddress = tempIP.Text()
		//query the FA
		result := getAPICallFA("http://" + ipAddress + ":8081/array-initial-config")

		initResult.SetText(result)

	})

	//initialize the array and do lots of other work
	button2.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate Controller 1 Gateway
		err7 := validate.Var(ct1GW.Text(), "required,ipv4")
		if err7 != nil {
			initResult.SetText("Please provide a valid Gateway for Controller 1")
			passed = false
		}
		//validate Controller 1 SN
		err8 := validate.Var(ct1SNM.Text(), "required,ipv4")
		if err8 != nil {
			initResult.SetText("Please provide a valid Subnet Mask for Controller 1")
			passed = false
		}
		//validate Controller 1 IP
		err9 := validate.Var(ct1IP.Text(), "required,ipv4")
		if err9 != nil {
			initResult.SetText("Please provide a valid IP Address for Controller 1")
			passed = false
		}
		//validate Controller 0 Gateway
		err10 := validate.Var(ct0GW.Text(), "required,ipv4")
		if err10 != nil {
			initResult.SetText("Please provide a valid Gateway for Controller 0")
			passed = false
		}
		//validate Controller 0 SN
		err11 := validate.Var(ct0SNM.Text(), "required,ipv4")
		if err11 != nil {
			initResult.SetText("Please provide a valid Subnet Mask for Controller 0")
			passed = false
		}
		//validate Controller 0 IP
		err12 := validate.Var(ct0IP.Text(), "required,ipv4")
		if err12 != nil {
			initResult.SetText("Please provide a valid IP Address for Controller 0")
			passed = false
		}
		//validate Virtual 0 Gateway
		err13 := validate.Var(vir0GW.Text(), "required,ipv4")
		if err13 != nil {
			initResult.SetText("Please provide a valid Gateway for Virtual 0")
			passed = false
		}
		//validate Virtual 0 SN
		err14 := validate.Var(vir0SNM.Text(), "required,ipv4")
		if err14 != nil {
			initResult.SetText("Please provide a valid Subnet Mask for Virtual 0")
			passed = false
		}
		//validate Virtual 0 IP
		err15 := validate.Var(vir0IP.Text(), "required,ipv4")
		if err15 != nil {
			initResult.SetText("Please provide a valid IP Address for Virtual 0")
			passed = false
		}
		//validate Ntp server
		err5 := validate.Var(ntpServer.Text(), "required")
		if err5 != nil {
			initResult.SetText("Please provide an NTP Server")
			passed = false
		}
		//validate eula
		if eulaAccept.Checked() != true {
			initResult.SetText("You must accept the terms of our EULA")
			passed = false
		}
		//validate Eula Title
		err4 := validate.Var(eulaTitle.Text(), "required")
		if err4 != nil {
			initResult.SetText("Please provide your Job Title")
			passed = false
		}
		//validate Eula Name
		err3 := validate.Var(eulaName.Text(), "required")
		if err3 != nil {
			initResult.SetText("Please provide your Full Name")
			passed = false
		}
		//validate Eula Org Name
		err2 := validate.Var(eulaOrg.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide your Organization Name")
			passed = false
		}
		//validate Array Name
		var rxPat = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,54}[a-zA-Z0-9])?$`)
		if !rxPat.MatchString(arrayName.Text()) {
			initResult.SetText("ArrayName has blank or contains invalid characters.  It must begin with a number or letter, can contain a dash in the body of the name, but must also end with a number or letter.   No more than 55 characters in length.")
		}
		//validate DHCP Boot IP
		err0 := validate.Var(tempIP.Text(), "required")
		if err0 != nil {
			initResult.SetText("Please provide a valid IP Address for the DHCP boot IP")
			passed = false
		}

		//if all validation above passes then proceed...
		if passed == true {
			//cool site to generate struct from json https://mholt.github.io/json-to-go/
			//define the flash array json structure
			type FAS struct {
				ArrayName string `json:"array_name"`
				Ct0Eth0   struct {
					Address string `json:"address"`
					Netmask string `json:"netmask"`
					Gateway string `json:"gateway"`
				} `json:"ct0.eth0"`
				Ct1Eth0 struct {
					Address string `json:"address"`
					Netmask string `json:"netmask"`
					Gateway string `json:"gateway"`
				} `json:"ct1.eth0"`
				Vir0 struct {
					Address string `json:"address"`
					Netmask string `json:"netmask"`
					Gateway string `json:"gateway"`
				} `json:"vir0"`
				DNS struct {
					Domain      string   `json:"domain"`
					Nameservers []string `json:"nameservers"`
				} `json:"dns"`
				NtpServers []string `json:"ntp_servers"`
				Timezone   string   `json:"timezone"`
				SMTP       struct {
					RelayHost    string `json:"relay_host"`
					SenderDomain string `json:"sender_domain"`
				} `json:"smtp"`
				AlertEmails    []string `json:"alert_emails"`
				EulaAcceptance struct {
					Accepted   bool `json:"accepted"`
					AcceptedBy struct {
						Organization string `json:"organization"`
						FullName     string `json:"full_name"`
						JobTitle     string `json:"job_title"`
					} `json:"accepted_by"`
				} `json:"eula_acceptance"`
			}

			//slices for multiple entry fields
			//split string into slice(array) *need to add conditional here
			ns := strings.Split(dnsServer.Text(), ",")
			ntp := strings.Split(ntpServer.Text(), ",")
			ae := strings.Split(smtpAlertEmail.Text(), ",")

			//initialize FAS struct object
			FA := &FAS{}
			FA.ArrayName = arrayName.Text()
			FA.Ct0Eth0.Address = ct0IP.Text()
			FA.Ct0Eth0.Netmask = ct0SNM.Text()
			FA.Ct0Eth0.Gateway = ct0GW.Text()
			FA.Ct1Eth0.Address = ct0IP.Text()
			FA.Ct1Eth0.Netmask = ct0SNM.Text()
			FA.Ct1Eth0.Gateway = ct0GW.Text()
			FA.Vir0.Address = ct0IP.Text()
			FA.Vir0.Netmask = ct0SNM.Text()
			FA.Vir0.Gateway = ct0GW.Text()
			FA.DNS.Domain = dnsDomain.Text()
			FA.DNS.Nameservers = ns
			FA.NtpServers = ntp
			FA.Timezone = tz[timeZone.Selected()]
			FA.SMTP.RelayHost = smtpRelay.Text()
			FA.SMTP.SenderDomain = smtpDomain.Text()
			FA.AlertEmails = ae
			FA.EulaAcceptance.Accepted = eulaAccept.Checked()
			FA.EulaAcceptance.AcceptedBy.FullName = eulaName.Text()
			FA.EulaAcceptance.AcceptedBy.Organization = eulaOrg.Text()
			FA.EulaAcceptance.AcceptedBy.JobTitle = eulaTitle.Text()

			//marshal (json encode) the map into a json string
			FAData, err := json.Marshal(FA)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			//Prod make the rest call with the json payload and stores response
			resp := patchRestCallFA("http://"+tempIP.Text()+":8081/array-initial-config", FAData)

			//update the initResult field with response.
			initResult.SetText("JSON Response: \n" + resp)

		}
	})

	return hbox
}

func initializeFBTab() ui.Control {
	//fields for the form
	arrayName := ui.NewEntry()
	ntpServer := ui.NewEntry()

	dnsDomain := ui.NewEntry()
	dnsServer := ui.NewEntry()

	initResult := ui.NewMultilineEntry()

	//first column definition
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	//define vertical box inside column similar to a div
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	//Buttons Group for Left Column//

	//define the group for the form
	buttonGroup := ui.NewGroup("Form Controls")
	buttonGroup.SetMargined(true)

	//add group to the vertical box
	vbox.Append(buttonGroup, true)

	///Form Instantiation///
	//define the form for the button group
	buttonForm := ui.NewForm()
	buttonForm.SetPadded(true)

	///Button Definition Login///
	//embed the login form field inside the first form group
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 1 Login", ui.NewLabel(""), false)
	login := ui.NewButton("Login Page")
	buttonForm.Append("Apply API Key", login, false)
	//seperator line
	hbox.Append(ui.NewVerticalSeparator(), false)
	///End Button Definition///

	///Button Definition Array///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 2 Array Config", ui.NewLabel(""), false)
	array := ui.NewButton("Array Form")
	buttonForm.Append("Array Form", array, false)
	///End Button Definition///

	///Button Definition DNS///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 3 DNS Config", ui.NewLabel(""), false)
	dns := ui.NewButton("DNS Form")
	buttonForm.Append("DNS Form", dns, false)
	///End Button Definition///

	///Button Definition Hardware Connectors///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 4 Hardware-Ctl Config", ui.NewLabel(""), false)
	hwc := ui.NewButton("HWC Form")
	buttonForm.Append("HWC Form", hwc, false)
	///End Button Definition///

	///Button Definition Link Aggregation///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 5 LAG Config", ui.NewLabel(""), false)
	lag := ui.NewButton("LAG Form")
	buttonForm.Append("LAG Form", lag, false)
	///End Button Definition///

	///Button Definition Subnets Aggregation///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 6 Subnet Config", ui.NewLabel(""), false)
	subnet := ui.NewButton("Subnet Form")
	buttonForm.Append("Subnet Form", subnet, false)
	///End Button Definition///

	///Button Definition Network Interfaces///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 7 Network Config", ui.NewLabel(""), false)
	network := ui.NewButton("NIC Form")
	buttonForm.Append("NIC Form", network, false)
	///End Button Definition///

	///Button Definition smtp///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 8 SMTP Config", ui.NewLabel(""), false)
	smtp := ui.NewButton("SMTP Form")
	buttonForm.Append("SMTP Form", smtp, false)
	///End Button Definition///

	///Button Definition support///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 9 Support Config", ui.NewLabel(""), false)
	support := ui.NewButton("Support Form")
	buttonForm.Append("Phonehome Form", support, false)
	///End Button Definition///

	///Button Definition alert watchers///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 10 Alerts Config", ui.NewLabel(""), false)
	aw := ui.NewButton("Alerts Form")
	buttonForm.Append("Alerts Form", aw, false)
	///End Button Definition///

	///Button Definition admin///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 11 Admins", ui.NewLabel(""), false)
	admins := ui.NewButton("Admins Form")
	buttonForm.Append("Admins Form", admins, false)
	///End Button Definition///

	///Button Definition validation and finalization///
	buttonGroup.SetChild(buttonForm)
	buttonForm.Append("STEP 12 Final Step", ui.NewLabel(""), false)
	final := ui.NewButton("Finalize Form")
	buttonForm.Append("Finalize Form", final, false)
	///End Button Definition///

	//Middle column
	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	//Login FORM//
	loginGroup := ui.NewGroup("Login")
	loginGroup.SetMargined(false)
	vbox.Append(loginGroup, false)
	loginForm := ui.NewForm()
	loginForm.SetPadded(true)
	loginGroup.SetChild(loginForm)

	//variables
	apiToken := ui.NewEntry()
	xAuthTokenField := ui.NewEntry()
	loginSubmitButton := ui.NewButton("Create Session")
	getAPIVersionsButton := ui.NewButton("Query Versions")
	apiUrlForm := ui.NewEntry()
	managementIP := ui.NewEntry()
	//append variables to form
	loginForm.Append("Array API URL", apiUrlForm, false)
	loginForm.Append("", ui.NewLabel("format:  https://10.1.1.100/api/1.8"), false)
	loginForm.Append("api token", apiToken, false)
	//arrayForm.Append("", ui.NewLabel(""), false)
	loginForm.Append("", loginSubmitButton, false)
	loginForm.Append("", ui.NewLabel(""), false)
	loginForm.Append("", ui.NewLabel(""), false)
	loginForm.Append("IP of FB to Query", managementIP, false)
	loginForm.Append("Get API Versions and build api url", getAPIVersionsButton, false)

	//Array Form//
	arrayGroup := ui.NewGroup("Array Config")
	arrayGroup.SetMargined(false)
	vbox.Append(arrayGroup, false)
	arrayGroup.Hide()
	arrayForm := ui.NewForm()
	arrayForm.SetPadded(true)
	arrayGroup.SetChild(arrayForm)
	arrayForm.Append("Array Name", arrayName, false)
	arrayForm.Append("NTP Servers", ntpServer, false)
	arrayForm.Append("", ui.NewLabel("*Comma seperated for multiple entries"), false)
	arrayGetButton := ui.NewButton("Query Array")
	arrayPatchButton := ui.NewButton("Apply To Array")
	arrayForm.Append("", arrayPatchButton, false)
	arrayForm.Append("", ui.NewLabel(""), false)
	arrayForm.Append("", arrayGetButton, false)
	//end Array Form//

	//DNS Form//
	dnsGroup := ui.NewGroup("DNS Config")
	dnsGroup.SetMargined(false)
	vbox.Append(dnsGroup, false)
	dnsGroup.Hide()
	dnsForm := ui.NewForm()
	dnsForm.SetPadded(true)
	dnsGroup.SetChild(dnsForm)
	dnsForm.Append("DNS Domain Name", dnsDomain, false)
	dnsForm.Append("DNS Servers", dnsServer, false)
	dnsForm.Append("", ui.NewLabel("*Comma seperated for multiple entries"), false)
	dnsGetButton := ui.NewButton("Query Array")
	dnsPatchButton := ui.NewButton("Apply To Array")
	dnsForm.Append("", dnsPatchButton, false)
	dnsForm.Append("", ui.NewLabel(""), false)
	dnsForm.Append("", dnsGetButton, false)
	//end DNS Form//

	//Hardware connector form//
	hwcName := ui.NewEntry()
	hwcLaneSpeed := ui.NewEntry()
	hwcPortCount := ui.NewEntry()
	hwcGroup := ui.NewGroup("Hardware-Connector Config")
	hwcGroup.SetMargined(false)
	vbox.Append(hwcGroup, false)
	hwcGroup.Hide()
	hwcForm := ui.NewForm()
	hwcForm.SetPadded(true)
	hwcGroup.SetChild(hwcForm)
	hwcForm.Append("Hardware Connector Name", hwcName, false)
	hwcForm.Append("Lane Speed", hwcLaneSpeed, false)
	hwcForm.Append("Port Count", hwcPortCount, false)
	hwcGetButton := ui.NewButton("Query Array")
	hwcPatchButton := ui.NewButton("Apply To Array")
	hwcForm.Append("", hwcPatchButton, false)
	hwcForm.Append("", ui.NewLabel(""), false)
	hwcForm.Append("", hwcGetButton, false)
	//END hardware-Connectors Form//

	//Link Aggregation Groups connector form//
	//LAG Init group and form
	lagNew := ui.NewButton("Create New LAG")
	lagExisting := ui.NewButton("Update Existing")
	lagGetButton := ui.NewButton("Query LAG")
	lagDelete := ui.NewButton("Delete LAG")

	lagGroupInit := ui.NewGroup("LAG Options")
	lagGroupInit.SetMargined(false)
	vbox.Append(lagGroupInit, false)
	lagGroupInit.Hide()
	lagFormInit := ui.NewForm()
	lagFormInit.SetPadded(true)
	lagGroupInit.SetChild(lagFormInit)
	lagFormInit.Append("", lagNew, false)
	lagFormInit.Append("", lagExisting, false)
	lagFormInit.Append("", lagDelete, false)
	lagFormInit.Append("", lagGetButton, false)

	lagNameNew := ui.NewEntry()
	lagNameExisting := ui.NewEntry()
	lagPortsNew := ui.NewEntry()
	lagPortsExisting := ui.NewEntry()
	lagAddRemove := ui.NewCombobox()
	lagAddRemove.Append("Add Ports")
	lagAddRemove.Append("Remove Ports")

	//lag create new group and form
	lagGroupNew := ui.NewGroup("New LAG Config")
	lagGroupNew.SetMargined(false)
	vbox.Append(lagGroupNew, false)
	lagGroupNew.Hide()
	lagFormNew := ui.NewForm()
	lagFormNew.SetPadded(true)
	lagGroupNew.SetChild(lagFormNew)
	lagFormNew.Append("LAG Name", lagNameNew, false)
	lagFormNew.Append("Lag Port Name(s)", lagPortsNew, false)
	lagFormNew.Append("", ui.NewLabel("*Comma seperated for multiple entries"), false)
	lagPostButton := ui.NewButton("Create New LAG")
	lagFormNew.Append("", lagPostButton, false)

	//lag modify existing group and form
	lagGroupExisting := ui.NewGroup("Existing LAG Config")
	lagGroupExisting.SetMargined(false)
	vbox.Append(lagGroupExisting, false)
	lagGroupExisting.Hide()
	lagFormExisting := ui.NewForm()
	lagFormExisting.SetPadded(true)
	lagGroupExisting.SetChild(lagFormExisting)

	lagFormNew.Append("", ui.NewLabel(""), false)
	lagFormExisting.Append("", ui.NewLabel(""), false)
	lagFormExisting.Append("LAG Name", lagNameExisting, false)
	lagFormExisting.Append("Lag Port Name(s)", lagPortsExisting, false)
	lagFormExisting.Append("", ui.NewLabel("*Comma seperated for multiple entries"), false)

	lagFormExisting.Append("", ui.NewLabel(""), false)
	lagFormExisting.Append("Modify Ports", lagAddRemove, false)
	lagPatchButton := ui.NewButton("Update LAG Ports")
	lagFormExisting.Append("", lagPatchButton, false)

	//lag create delete group and form
	lagNameDelete := ui.NewEntry()
	lagDeleteConfirm := ui.NewCheckbox("Yes")
	lagGroupDelete := ui.NewGroup("New LAG Config")
	lagGroupDelete.SetMargined(false)
	vbox.Append(lagGroupDelete, false)
	lagGroupDelete.Hide()
	lagFormDelete := ui.NewForm()
	lagFormDelete.SetPadded(true)
	lagGroupDelete.SetChild(lagFormDelete)
	lagFormDelete.Append("LAG Name", lagNameDelete, false)
	lagFormDelete.Append("Confirm Delete", lagDeleteConfirm, false)
	lagDeleteButton := ui.NewButton("Delete LAG")
	lagFormDelete.Append("", lagDeleteButton, false)
	//END link aggrigation Form//

	//subnets Form//
	subnetGateway := ui.NewEntry()
	subnetLag := ui.NewEntry()
	subnetMtu := ui.NewEntry()
	subnetPrefix := ui.NewEntry()
	subnetVlan := ui.NewEntry()
	subnetName := ui.NewEntry()
	subnetInterfaceName := ui.NewEntry()
	subnetEnabled := ui.NewCombobox()
	subnetEnabled.Append("true")
	subnetEnabled.Append("false")
	subnetServices := ui.NewEntry()

	subnetGroup := ui.NewGroup("Subnet Config")
	subnetGroup.SetMargined(false)
	vbox.Append(subnetGroup, false)
	subnetGroup.Hide()
	subnetForm := ui.NewForm()
	subnetForm.SetPadded(true)
	subnetGroup.SetChild(subnetForm)
	subnetForm.Append("Subnet Name", subnetName, false)
	subnetForm.Append("Enabled", subnetEnabled, false)
	subnetForm.Append("Gateway", subnetGateway, false)
	subnetForm.Append("Interface Name", subnetInterfaceName, false)
	subnetForm.Append("LAG Name", subnetLag, false)
	subnetForm.Append("MTU", subnetMtu, false)
	subnetForm.Append("Prefix", subnetPrefix, false)
	subnetForm.Append("VLAN", subnetVlan, false)
	subnetForm.Append("Services", subnetServices, false)

	subnetGetButton := ui.NewButton("Query")
	subnetPatchButton := ui.NewButton("Update Existing")
	subnetPostButton := ui.NewButton("Create New")
	subnetDeleteButton := ui.NewButton("Delete")

	subnetForm.Append("", subnetPostButton, false)
	subnetForm.Append("", subnetPatchButton, false)
	subnetForm.Append("", subnetGetButton, false)
	subnetForm.Append("", subnetDeleteButton, false)
	//end subnets Form//

	//network interfaces Form//
	nicAddress := ui.NewEntry()
	nicName := ui.NewEntry()
	nicStatus := ui.NewEntry()
	nicStatus.SetText("enabled")
	nicType := ui.NewEntry()
	nicType.SetText("vip")

	nicGroup := ui.NewGroup("Net Interface Config")
	nicGroup.SetMargined(false)
	vbox.Append(nicGroup, false)
	nicGroup.Hide()
	nicForm := ui.NewForm()
	nicForm.SetPadded(true)
	nicGroup.SetChild(nicForm)
	nicForm.Append("Interface Name", nicName, false)
	nicForm.Append("IP Address", nicAddress, false)
	nicForm.Append("Status", nicStatus, false)
	nicForm.Append("Type", nicType, false)

	nicGetButton := ui.NewButton("Query Array")
	nicPatchButton := ui.NewButton("Update NIC")

	nicForm.Append("", nicPatchButton, false)
	nicForm.Append("", nicGetButton, false)
	//end network interfaces Form//

	//smtp Form//
	smtpRelayHost := ui.NewEntry()
	smtpSenderDomain := ui.NewEntry()

	smtpGroup := ui.NewGroup("SMTP Config")
	smtpGroup.SetMargined(false)
	vbox.Append(smtpGroup, false)
	smtpGroup.Hide()
	smtpForm := ui.NewForm()
	smtpForm.SetPadded(true)
	smtpGroup.SetChild(smtpForm)
	smtpForm.Append("Relay Host", smtpRelayHost, false)
	smtpForm.Append("Sender Domain", smtpSenderDomain, false)

	smtpGetButton := ui.NewButton("Query")
	smtpPatchButton := ui.NewButton("Create New")

	smtpForm.Append("", smtpPatchButton, false)
	smtpForm.Append("", smtpGetButton, false)
	//end smtp Form//

	//support Form//
	supportPhoneHome := ui.NewCombobox()
	supportPhoneHome.Append("true")
	supportPhoneHome.Append("false")
	supportProxy := ui.NewEntry()

	supportGroup := ui.NewGroup("Support Config")
	supportGroup.SetMargined(false)
	vbox.Append(supportGroup, false)
	supportGroup.Hide()
	supportForm := ui.NewForm()
	supportForm.SetPadded(true)
	supportGroup.SetChild(supportForm)
	supportForm.Append("Enable Phone Home", supportPhoneHome, false)
	supportForm.Append("Proxy Server", supportProxy, false)
	supportForm.Append("", ui.NewLabel("*Comma seperated for multiple entries"), false)
	supportGetButton := ui.NewButton("Query Array")
	supportPatchButton := ui.NewButton("Apply To Array")

	supportForm.Append("", supportPatchButton, false)
	supportForm.Append("", supportGetButton, false)
	//end support Form//

	//alert watchers Form//
	awName := ui.NewEntry()
	awEnabled := ui.NewCombobox()
	awEnabled.Append("true")
	awEnabled.Append("false")

	awGroup := ui.NewGroup("Alert Watchers Config")
	awGroup.SetMargined(false)
	vbox.Append(awGroup, false)
	awGroup.Hide()
	awForm := ui.NewForm()
	awForm.SetPadded(true)
	awGroup.SetChild(awForm)
	awForm.Append("Email Address", awName, false)
	awForm.Append("Enabled", awEnabled, false)

	awGetButton := ui.NewButton("Query")
	awPatchButton := ui.NewButton("Update Existing")
	awDeleteButton := ui.NewButton("Delete Alert Watcher")
	awPostButton := ui.NewButton("New Alert Watcher")

	awForm.Append("", awPostButton, false)
	awForm.Append("", awPatchButton, false)
	awForm.Append("", awDeleteButton, false)
	awForm.Append("", awGetButton, false)
	//end alert watchers Form//

	//Admins Form//
	adminName := ui.NewEntry()
	adminsCreateToken := ui.NewCombobox()
	adminsCreateToken.Append("true")
	adminsCreateToken.Append("false")

	adminsGroup := ui.NewGroup("Admins Config")
	adminsGroup.SetMargined(false)
	vbox.Append(adminsGroup, false)
	adminsGroup.Hide()
	adminsForm := ui.NewForm()
	adminsForm.SetPadded(true)
	adminsGroup.SetChild(adminsForm)

	adminsForm.Append("Admin UserName", adminName, false)
	adminsForm.Append("Create API Token", adminsCreateToken, false)

	adminsGetButton := ui.NewButton("Query Admins")
	adminsPatchButton := ui.NewButton("Update Admin")

	adminsForm.Append("", adminsPatchButton, false)
	adminsForm.Append("", adminsGetButton, false)
	//end admins Form//

	//finalization and validation Form//
	finalSetupComplete := ui.NewCombobox()
	finalSetupComplete.Append("true")
	finalSetupComplete.Append("false")

	finalGroup := ui.NewGroup("Validate and Finalize")
	finalGroup.SetMargined(false)
	vbox.Append(finalGroup, false)
	finalGroup.Hide()
	finalForm := ui.NewForm()
	finalForm.SetPadded(true)
	finalGroup.SetChild(finalForm)
	finalForm.Append("Setup Complete", finalSetupComplete, false)

	finalGetButton := ui.NewButton("Validate")
	finalPatchButton := ui.NewButton("Finalize Setup")

	finalForm.Append("", finalPatchButton, false)
	finalForm.Append("", finalGetButton, false)
	//end finalization and validation Form//

	hbox.Append(ui.NewVerticalSeparator(), false)

	//third column
	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	//SUBMIT "GO" BUTTON//
	group9 := ui.NewGroup("Initialize Array")
	group9.SetMargined(true)
	vbox.Append(group9, true)

	entryForm9 := ui.NewForm()
	entryForm9.SetPadded(true)
	group9.SetChild(entryForm9)

	//labels used to display the api url and x-auth token in third column
	xAuthTokenLabel := ui.NewLabel("")
	apiUrlLabel := ui.NewLabel("")
	xAuthTokenField.SetReadOnly(true)
	entryForm9.Append("API URL: ", apiUrlLabel, false)
	entryForm9.Append("X-Auth-Token", xAuthTokenLabel, false)

	//multiline field for showing results of patch api call and form validation messages.
	//sets the initResults console to readonly
	initResult.SetReadOnly(true)
	entryForm9.Append("Init Results", initResult, true)

	//Login Form Button
	login.OnClicked(func(*ui.Button) {
		loginGroup.Show()
		arrayGroup.Hide()
		dnsGroup.Hide()
		hwcGroup.Hide()
		lagGroupNew.Hide()
		lagGroupExisting.Hide()
		lagGroupInit.Hide()
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupDelete.Hide()
		initResult.SetText("Please fill out the form to obtain an x-auth-token to proceed.")

	})

	//arrays Form Button
	array.OnClicked(func(*ui.Button) {
		arrayGroup.Show()
		loginGroup.Hide()
		dnsGroup.Hide()
		hwcGroup.Hide()
		lagGroupNew.Hide()
		lagGroupExisting.Hide()
		lagGroupInit.Hide()
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupDelete.Hide()
		initResult.SetText("Query or update array info")

	})

	//DNS Form Button
	dns.OnClicked(func(*ui.Button) {
		dnsGroup.Show()
		arrayGroup.Hide()
		loginGroup.Hide()
		hwcGroup.Hide()
		lagGroupNew.Hide()
		lagGroupExisting.Hide()
		lagGroupInit.Hide()
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupDelete.Hide()
		initResult.SetText("Query or update array info")

	})

	//HWC Form Button
	hwc.OnClicked(func(*ui.Button) {
		hwcGroup.Show()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		lagGroupNew.Hide()
		lagGroupExisting.Hide()
		lagGroupInit.Hide()
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupDelete.Hide()
		initResult.SetText("Query or update array info")

	})

	//LAG Form Button
	lag.OnClicked(func(*ui.Button) {
		lagGroupInit.Show()
		lagGroupNew.Hide()
		lagGroupExisting.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupDelete.Hide()
		initResult.SetText("Query or update array info")

	})

	//LAG New Form Button
	lagNew.OnClicked(func(*ui.Button) {
		lagGroupInit.Show()
		lagGroupNew.Show()
		lagGroupExisting.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupDelete.Hide()
		initResult.SetText("Query or update array info")

	})

	//LAG existing Form Button
	lagExisting.OnClicked(func(*ui.Button) {
		lagGroupInit.Show()
		lagGroupExisting.Show()
		lagGroupNew.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupDelete.Hide()
		initResult.SetText("Query or update array info")

	})

	//LAG delete Form Button
	lagDelete.OnClicked(func(*ui.Button) {
		lagGroupDelete.Show()
		lagGroupInit.Show()
		lagGroupExisting.Hide()
		lagGroupNew.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		initResult.SetText("Query or update array info")

	})

	//Subnet Form Button
	subnet.OnClicked(func(*ui.Button) {
		finalGroup.Hide()
		subnetGroup.Show()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupInit.Hide()
		lagGroupExisting.Hide()
		lagGroupNew.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		initResult.SetText("Query or update array info")

	})

	//Network Init Form Button
	network.OnClicked(func(*ui.Button) {
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Show()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupInit.Hide()
		lagGroupExisting.Hide()
		lagGroupNew.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		initResult.SetText("Query or update array info")

	})

	//SMTP Form Button
	smtp.OnClicked(func(*ui.Button) {
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Show()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupInit.Hide()
		lagGroupExisting.Hide()
		lagGroupNew.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		initResult.SetText("Query or update array info")

	})

	//Support Form Button
	support.OnClicked(func(*ui.Button) {
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Show()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupInit.Hide()
		lagGroupExisting.Hide()
		lagGroupNew.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		initResult.SetText("Query or update array info")

	})

	//Alert Watchers Form Button
	aw.OnClicked(func(*ui.Button) {
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Show()
		adminsGroup.Hide()
		lagGroupInit.Hide()
		lagGroupExisting.Hide()
		lagGroupNew.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		initResult.SetText("Query or update array info")

	})

	//Admins Form Button
	admins.OnClicked(func(*ui.Button) {
		lagGroupInit.Hide()
		lagGroupExisting.Hide()
		lagGroupNew.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		finalGroup.Hide()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Show()
		initResult.SetText("Query or update array info")

	})

	//Final Validation Form Button
	final.OnClicked(func(*ui.Button) {
		finalGroup.Show()
		subnetGroup.Hide()
		nicGroup.Hide()
		smtpGroup.Hide()
		supportGroup.Hide()
		awGroup.Hide()
		adminsGroup.Hide()
		lagGroupInit.Hide()
		lagGroupExisting.Hide()
		lagGroupNew.Hide()
		hwcGroup.Hide()
		dnsGroup.Hide()
		arrayGroup.Hide()
		loginGroup.Hide()
		initResult.SetText("Query or update array info")

	})

	//Buttons Actions from Forms//

	//QUERY FOR API VERSIONS//
	getAPIVersionsButton.OnClicked(func(*ui.Button) {
		//make sure the api endpoints are in the right format
		passed := true
		validate := validator.New()
		err := validate.Var(managementIP.Text(), "required,ipv4")
		if err != nil {
			initResult.SetText("Please provide a valid IP Address for the FB management endpoint")
			passed = false
		}
		if passed == true {
			//make the rest call
			resp := getAPICallFB("https://"+managementIP.Text()+"/api/api_version", apiToken.Text())

			type Version struct {
				Versions []string `json:"versions"`
			}

			var version Version
			err := json.Unmarshal(resp, &version)
			if err == nil {
				fmt.Println(err)
			}
			fmt.Printf("%v", (len(version.Versions)))
			if len(version.Versions) > 0 {
				apiUrlForm.SetText("https://" + managementIP.Text() + "/api/" + version.Versions[(len(version.Versions)-1)])
				apiUrlLabel.SetText("https://" + managementIP.Text() + "/api/" + version.Versions[(len(version.Versions)-1)])
			}
			//set the response in the display of the app
			initResult.SetText(string(resp))

		}

	})

	//LOGIN SUBMIT//
	loginSubmitButton.OnClicked(func(*ui.Button) {
		//make sure the api endpoints are in the right format
		passed := true
		url := urlParser(apiUrlForm.Text())
		if len(url) > 4 {
			loginUrl = (url[0] + "//" + url[2] + "/" + url[3])
			apiUrl = (url[0] + "//" + url[2] + "/" + url[3] + "/" + url[4])
		} else {
			initResult.SetText("please enter a valid API url that includes the version.  e.g. https://purefb01.example.com/api/1.8")
			passed = false
		}
		if passed == true {
			apiUrlLabel.SetText(apiUrl)
			fmt.Println(apiUrl)

			//make the rest call
			resp := postAPICallLoginFB(loginUrl+"/login", apiToken.Text())
			//set the response in the display of the app
			initResult.SetText(string(resp))
			xAuthTokenField.SetText(xAuthToken)
			xAuthTokenLabel.SetText(xAuthToken)
		}

	})

	arrayGetButton.OnClicked(func(*ui.Button) {
		result := getAPICallFB(apiUrl+"/arrays", xAuthToken)
		initResult.SetText(string(result))
	})

	arrayPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err1 := validate.Var(ntpServer.Text(), "required")
		if err1 != nil {
			initResult.SetText("Please provide the Ntp server(s)")
			passed = false
		}
		//validate Array Name
		var rxPat = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,54}[a-zA-Z0-9])?$`)
		if !rxPat.MatchString(arrayName.Text()) {
			initResult.SetText("Array Name has blank or contains invalid characters.  It must begin with a number or letter, can contain a dash in the body of the name, but must also end with a number or letter.   No more than 55 characters in length.")
			passed = false
		}
		if passed == true {
			//struct here
			type FAB struct {
				Name      string   `json:"name"`
				NtpServer []string `json:"ntp_servers"`
			}

			//slices for multiple entry fields
			//split string into slice(array) *need to add conditional here
			ntp := strings.Split(ntpServer.Text(), ",")

			//initialize FAS struct object
			FB := &FAB{}
			FB.Name = arrayName.Text()
			FB.NtpServer = ntp

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/arrays", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	dnsGetButton.OnClicked(func(*ui.Button) {
		result := getAPICallFB(apiUrl+"/dns", xAuthToken)
		initResult.SetText(string(result))
	})

	dnsPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(dnsServer.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide the dns server(s)")
			passed = false
		}
		err1 := validate.Var(dnsDomain.Text(), "required")
		if err1 != nil {
			initResult.SetText("Please provide the dns domain")
			passed = false
		}
		if passed == true {

			//struct here
			type FAB struct {
				Domain      string   `json:"domain"`
				Nameservers []string `json:"nameservers"`
			}

			//slices for multiple entry fields
			//split string into slice(array) *need to add conditional here
			dns := strings.Split(dnsServer.Text(), ",")

			//initialize FAS struct object
			FB := &FAB{}
			FB.Domain = dnsDomain.Text()
			FB.Nameservers = dns

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/dns", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	//HWC Buttons
	hwcGetButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(hwcName.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide a Hardware Connector name to query for")
			passed = false
		}
		if passed {
			result := getAPICallFB(apiUrl+"/hardware-connectors?names="+hwcName.Text(), xAuthToken)
			initResult.SetText(string(result))
		}
	})

	hwcPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(hwcName.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide the Array name")
			passed = false
		}

		if passed == true {

			//struct here
			type FAB struct {
				Lanespeed string `json:"lane_speed"`
				Portcount string `json:"port_count"`
			}

			//initialize FAS struct object
			FB := &FAB{}
			FB.Lanespeed = hwcLaneSpeed.Text()
			FB.Portcount = hwcPortCount.Text()

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/hardware-connectors?names="+hwcName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})
	//Lag Buttons
	lagGetButton.OnClicked(func(*ui.Button) {

		result := getAPICallFB(apiUrl+"/link-aggregation-groups", xAuthToken)
		initResult.SetText(string(result))

	})

	lagPostButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(lagNameNew.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide the Array name")
			passed = false
		}
		err1 := validate.Var(lagPortsNew.Text(), "required")
		if err1 != nil {
			initResult.SetText("Please provide the Port Name(s)")
			passed = false
		}

		if passed == true {
			//this was tricky and probably not the best way to accompish this but it works.
			portNames := strings.Split(lagPortsNew.Text(), ",")
			var pName = `{"ports": [`
			for i, v := range portNames {
				i++
				pName += `{"name": "`
				pName += v
				pName += `"}`
				if i < len(portNames) {
					pName += `,`
				}
			}
			pName += `]}`
			pNameSlice := []byte(pName)

			result := postAPICall2FB(apiUrl+"/link-aggregation-groups?names="+lagNameNew.Text(), xAuthToken, pNameSlice)
			initResult.SetText(result)
		}
	})

	lagPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(lagNameExisting.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide the Array name")
			passed = false
		}

		if passed == true {

			//this was tricky and probably not the best way to accompish this but it works.
			portNames := strings.Split(lagPortsExisting.Text(), ",")
			var pName = ""
			if lagAddRemove.Selected() == 0 {
				pName += `{"add_ports":[`
			}
			if lagAddRemove.Selected() == 1 {
				pName += `{"remove_ports":[`
			}
			for i, v := range portNames {
				i++
				pName += `{"name":"`
				pName += v
				pName += `"}`
				if i < len(portNames) {
					pName += `,`
				}
			}
			pName += `]}`
			pNameSlice := []byte(pName)

			result := patchAPICallFB(apiUrl+"/link-aggregation-groups?names="+lagNameExisting.Text(), xAuthToken, pNameSlice)
			initResult.SetText(result)
		}
	})

	lagDeleteButton.OnClicked(func(*ui.Button) {

		result := deleteAPICallFB(apiUrl+"/link-aggregation-groups?names="+lagNameDelete.Text(), xAuthToken)
		initResult.SetText(result)

	})

	subnetGetButton.OnClicked(func(*ui.Button) {
		result := getAPICallFB(apiUrl+"/subnets", xAuthToken)
		initResult.SetText(string(result))
	})

	subnetPostButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(subnetGateway.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide the gateway")
			passed = false
		}
		err4 := validate.Var(subnetPrefix.Text(), "required")
		if err4 != nil {
			initResult.SetText("Please provide the Prefix")
			passed = false
		}
		err5 := validate.Var(subnetName.Text(), "required")
		if err5 != nil {
			initResult.SetText("Please provide the Subnet Name")
			passed = false
		}

		if passed == true {

			type FAB struct {
				Enabled    bool   `json:"enabled"`
				Gateway    string `json:"gateway"`
				Interfaces struct {
					Name string `json:"name"`
				} `json:"interfaces"`
				LinkAggregationGroup struct {
					Name string `json:"link_aggregation_group"`
				} `json:"link_aggregation_group"`
				Mtu      int      `json:"mtu"`
				Prefix   string   `json:"prefix"`
				Services []string `json:"services"`
				Vlan     int      `json:"vlan"`
			}

			//slices for multiple entry fields
			//split string into slice(array) *need to add conditional here
			svc := strings.Split(subnetServices.Text(), ",")
			var isEnabled bool
			if subnetEnabled.Selected() == 0 {
				isEnabled = true
			}
			if subnetEnabled.Selected() == 1 {
				isEnabled = false
			}
			mtuInt, err := strconv.Atoi(subnetMtu.Text())
			vlanInt, err := strconv.Atoi(subnetVlan.Text())
			//initialize FAS struct object
			FB := &FAB{}
			FB.Enabled = isEnabled
			FB.Gateway = subnetGateway.Text()
			FB.Interfaces.Name = subnetInterfaceName.Text()
			FB.LinkAggregationGroup.Name = subnetLag.Text()
			FB.Mtu = mtuInt
			FB.Prefix = subnetPrefix.Text()
			FB.Services = svc
			FB.Vlan = vlanInt

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := postAPICall2FB(apiUrl+"/subnets?names="+subnetName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	subnetPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(subnetGateway.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide the gateway")
			passed = false
		}
		err4 := validate.Var(subnetPrefix.Text(), "required")
		if err4 != nil {
			initResult.SetText("Please provide the Prefix")
			passed = false
		}
		err5 := validate.Var(subnetName.Text(), "required")
		if err5 != nil {
			initResult.SetText("Please provide the Subnet Name")
			passed = false
		}

		if passed == true {

			type FAB struct {
				Enabled    bool   `json:"enabled"`
				Gateway    string `json:"gateway"`
				Interfaces struct {
					Name string `json:"name"`
				} `json:"interfaces"`
				LinkAggregationGroup struct {
					Name string `json:"name"`
				} `json:"link_aggregation_group"`
				Mtu      int      `json:"mtu"`
				Prefix   string   `json:"prefix"`
				Services []string `json:"services"`
				Vlan     int      `json:"vlan"`
			}

			//slices for multiple entry fields
			//split string into slice(array) *need to add conditional here
			svc := strings.Split(subnetServices.Text(), ",")
			var isEnabled bool
			if subnetEnabled.Selected() == 0 {
				isEnabled = true
			}
			if subnetEnabled.Selected() == 1 {
				isEnabled = false
			}
			mtuInt, err := strconv.Atoi(subnetMtu.Text())
			vlanInt, err := strconv.Atoi(subnetVlan.Text())
			//initialize FAS struct object
			FB := &FAB{}
			FB.Enabled = isEnabled
			FB.Gateway = subnetGateway.Text()
			FB.Interfaces.Name = subnetInterfaceName.Text()
			FB.LinkAggregationGroup.Name = subnetLag.Text()
			FB.Mtu = mtuInt
			FB.Prefix = subnetPrefix.Text()
			FB.Services = svc
			FB.Vlan = vlanInt

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/subnets?names="+subnetName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	subnetDeleteButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		err := validate.Var(subnetName.Text(), "required")
		if err != nil {
			initResult.SetText("Please provide the Subnet Name To Delete")
			passed = false
		}
		if passed == true {
			result := deleteAPICallFB(apiUrl+"/subnets?names="+subnetName.Text(), xAuthToken)
			initResult.SetText(result)
		}
	})

	nicGetButton.OnClicked(func(*ui.Button) {
		result := getAPICallFB(apiUrl+"/network-interfaces", xAuthToken)
		initResult.SetText(string(result))
	})

	nicPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(nicName.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide the Name")
			passed = false
		}
		err4 := validate.Var(nicAddress.Text(), "required")
		if err4 != nil {
			initResult.SetText("Please provide the IP")
			passed = false
		}
		err5 := validate.Var(nicType.Text(), "required")
		if err5 != nil {
			initResult.SetText("Please provide the Type")
			passed = false
		}
		err6 := validate.Var(nicStatus.Text(), "required")
		if err6 != nil {
			initResult.SetText("Please provide the Status")
			passed = false
		}

		if passed == true {

			type FAB struct {
				Address string `json:"address"`
				Type    string `json:"type"`
				Status  string `json:"status"`
			}

			//initialize FAS struct object
			FB := &FAB{}
			FB.Address = nicAddress.Text()
			FB.Type = nicType.Text()
			FB.Status = nicStatus.Text()

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/network-interfaces?names="+nicName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	smtpGetButton.OnClicked(func(*ui.Button) {
		result := getAPICallFB(apiUrl+"/smtp", xAuthToken)
		initResult.SetText(string(result))
	})

	smtpPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(smtpRelayHost.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide the SMTP Relay Host Name")
			passed = false
		}
		err4 := validate.Var(smtpSenderDomain.Text(), "required")
		if err4 != nil {
			initResult.SetText("Please provide the sender domain")
			passed = false
		}

		if passed == true {

			type FAB struct {
				Relay  string `json:"relay_host"`
				Domain string `json:"sender_domain"`
			}

			//initialize FAS struct object
			FB := &FAB{}
			FB.Relay = smtpRelayHost.Text()
			FB.Domain = smtpSenderDomain.Text()

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/smtp", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	supportGetButton.OnClicked(func(*ui.Button) {
		result := getAPICallFB(apiUrl+"/support", xAuthToken)
		initResult.SetText(string(result))
	})

	supportPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err4 := validate.Var(supportProxy.Text(), "required")
		if err4 != nil {
			initResult.SetText("Please provide the proxy server")
			passed = false
		}
		if supportPhoneHome.Selected() == -1 {
			initResult.SetText("Please provide select true or false for phone home")
			passed = false
		}

		var phoneHome = ""
		if supportPhoneHome.Selected() == 0 {
			phoneHome = "true"
		}
		if supportPhoneHome.Selected() == 1 {
			phoneHome = "false"
		}

		if passed == true {

			type FAB struct {
				Phonehome string `json:"phonehome_enabled"`
				Proxy     string `json:"proxy"`
			}

			//initialize FAS struct object
			FB := &FAB{}
			FB.Phonehome = phoneHome
			FB.Proxy = supportProxy.Text()

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/support", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	awGetButton.OnClicked(func(*ui.Button) {
		result := getAPICallFB(apiUrl+"/alert-watchers", xAuthToken)
		initResult.SetText(string(result))
	})

	awPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(awName.Text(), "email")
		if err2 != nil {
			initResult.SetText("Please provide an email address")
			passed = false
		}
		if awEnabled.Selected() == -1 {
			initResult.SetText("Please provide select true or false for enabled")
			passed = false
		}

		var awIsEnabled = ""
		if awEnabled.Selected() == 0 {
			awIsEnabled = "true"
		}
		if awEnabled.Selected() == 1 {
			awIsEnabled = "false"
		}

		if passed == true {

			type FAB struct {
				Enabled string `json:"enabled"`
			}

			//initialize FAS struct object
			FB := &FAB{}
			FB.Enabled = awIsEnabled

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/alert-watchers?names="+awName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	awDeleteButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err := validate.Var(awName.Text(), "email")
		if err != nil {
			initResult.SetText("Please provide an email address")
			passed = false
		}
		if passed == true {
			result := deleteAPICallFB(apiUrl+"/alert-watchers?names="+awName.Text(), xAuthToken)
			initResult.SetText(result)
		}
	})

	awPostButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(awName.Text(), "email")
		if err2 != nil {
			initResult.SetText("Please provide an email address")
			passed = false
		}

		if awEnabled.Selected() == -1 {
			initResult.SetText("Please provide select true or false for enabled")
			passed = false
		}

		var awIsEnabled = ""
		if awEnabled.Selected() == 0 {
			awIsEnabled = "true"
		}
		if awEnabled.Selected() == 1 {
			awIsEnabled = "false"
		}
		if passed == true {

			type FAB struct {
				Enabled string `json:"enabled"`
			}

			//initialize FAS struct object
			FB := &FAB{}
			FB.Enabled = awIsEnabled

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := postAPICall2FB(apiUrl+"/alert-watchers?names="+awName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	adminsGetButton.OnClicked(func(*ui.Button) {
		result := getAPICallFB(apiUrl+"/admins", xAuthToken)
		initResult.SetText(string(result))
	})

	adminsPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(adminName.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide an admin username")
			passed = false
		}
		if adminsCreateToken.Selected() == -1 {
			initResult.SetText("Please provide select true or false for enabled")
			passed = false
		}

		var adminsCreateTokenIsEnabled = ""
		if adminsCreateToken.Selected() == 0 {
			adminsCreateTokenIsEnabled = "true"
		}
		if adminsCreateToken.Selected() == 1 {
			adminsCreateTokenIsEnabled = "false"
		}

		if passed == true {

			type FAB struct {
				CToken string `json:"create_api_token"`
			}

			//initialize FAS struct object
			FB := &FAB{}
			FB.CToken = adminsCreateTokenIsEnabled

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/admins?names="+adminName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	finalGetButton.OnClicked(func(*ui.Button) {
		result := getAPICallFB(apiUrl+"/setup/validation", xAuthToken)
		initResult.SetText(string(result))
	})

	finalPatchButton.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true

		//validate ntp servers
		if finalSetupComplete.Selected() == -1 {
			initResult.SetText("Please provide select true or false")
			passed = false
		}

		var finalSetupCompleteIsComplete = ""
		if finalSetupComplete.Selected() == 0 {
			finalSetupCompleteIsComplete = "true"
		}
		if finalSetupComplete.Selected() == 1 {
			finalSetupCompleteIsComplete = "false"
		}

		if passed == true {

			type FAB struct {
				Complete string `json:"setup_completed"`
			}

			//initialize FAS struct object
			FB := &FAB{}
			FB.Complete = finalSetupCompleteIsComplete

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICallFB(apiUrl+"/setup/finalization", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})
	//END Button Actions from forms//

	return hbox
}

//define the UI interface//
func setupUI() {
	mainwin = ui.NewWindow("Pure Storage Zero Touch Provisioner for Flash Array", 800, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("ZTP Flash Array", initializeFATab())
	tab.SetMargined(0, true)

	tab.Append("ZTP Flash Blade", initializeFBTab())
	tab.SetMargined(0, true)

	mainwin.Show()
}

//main//
func main() {
	ui.Main(setupUI)
}
