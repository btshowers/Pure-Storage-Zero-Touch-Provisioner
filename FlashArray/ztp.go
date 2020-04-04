//Brandon Showers
//March 22 2020
//v1

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"gopkg.in/go-playground/validator.v9"
)

var mainwin *ui.Window
var ipAddress = ""

func timeZones() []string {
	tz := []string{"Africa/Abidjan", "Africa/Accra", "Africa/Addis_Ababa", "Africa/Algiers", "Africa/Asmara", "Africa/Bamako", "Africa/Bangui", "Africa/Banjul", "Africa/Bissau", "Africa/Blantyre", "Africa/Brazzaville", "Africa/Bujumbura", "Africa/Cairo", "Africa/Casablanca", "Africa/Ceuta", "Africa/Conakry", "Africa/Dakar", "Africa/Dar_es_Salaam", "Africa/Djibouti", "Africa/Douala", "Africa/El_Aaiun", "Africa/Freetown", "Africa/Gaborone", "Africa/Harare", "Africa/Johannesburg", "Africa/Juba", "Africa/Kampala", "Africa/Khartoum", "Africa/Kigali", "Africa/Kinshasa", "Africa/Lagos", "Africa/Libreville", "Africa/Lome", "Africa/Luanda", "Africa/Lubumbashi", "Africa/Lusaka", "Africa/Malabo", "Africa/Maputo", "Africa/Maseru", "Africa/Mbabane", "Africa/Mogadishu", "Africa/Monrovia", "Africa/Nairobi", "Africa/Ndjamena", "Africa/Niamey", "Africa/Nouakchott", "Africa/Ouagadougou", "Africa/Porto-Novo", "Africa/Sao_Tome", "Africa/Tripoli", "Africa/Tunis", "Africa/Windhoek", "America/Adak", "America/Anchorage", "America/Anguilla", "America/Antigua", "America/Araguaina", "America/Argentina/Buenos_Aires", "America/Argentina/Catamarca", "America/Argentina/Cordoba", "America/Argentina/Jujuy", "America/Argentina/La_Rioja", "America/Argentina/Mendoza", "America/Argentina/Rio_Gallegos", "America/Argentina/Salta", "America/Argentina/San_Juan", "America/Argentina/San_Luis", "America/Argentina/Tucuman", "America/Argentina/Ushuaia", "America/Aruba", "America/Asuncion", "America/Atikokan", "America/Bahia", "America/Bahia_Banderas", "America/Barbados", "America/Belem", "America/Belize", "America/Blanc-Sablon", "America/Boa_Vista", "America/Bogota", "America/Boise", "America/Cambridge_Bay", "America/Campo_Grande", "America/Cancun", "America/Caracas", "America/Cayenne", "America/Cayman", "America/Chicago", "America/Chihuahua", "America/Costa_Rica", "America/Creston", "America/Cuiaba", "America/Curacao", "America/Danmarkshavn", "America/Dawson", "America/Dawson_Creek", "America/Denver", "America/Detroit", "America/Dominica", "America/Edmonton", "America/Eirunepe", "America/El_Salvador", "America/Fort_Nelson", "America/Fortaleza", "America/Glace_Bay", "America/Godthab", "America/Goose_Bay", "America/Grand_Turk", "America/Grenada", "America/Guadeloupe", "America/Guatemala", "America/Guayaquil", "America/Guyana", "America/Halifax", "America/Havana", "America/Hermosillo", "America/Indiana/Indianapolis", "America/Indiana/Knox", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Tell_City", "America/Indiana/Vevay", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Inuvik", "America/Iqaluit", "America/Jamaica", "America/Juneau", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Kralendijk", "America/La_Paz", "America/Lima", "America/Los_Angeles", "America/Lower_Princes", "America/Maceio", "America/Managua", "America/Manaus", "America/Marigot", "America/Martinique", "America/Matamoros", "America/Mazatlan", "America/Menominee", "America/Merida", "America/Metlakatla", "America/Mexico_City", "America/Miquelon", "America/Moncton", "America/Monterrey", "America/Montevideo", "America/Montserrat", "America/Nassau", "America/New_York", "America/Nipigon", "America/Nome", "America/Noronha", "America/North_Dakota/Beulah", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/Ojinaga", "America/Panama", "America/Pangnirtung", "America/Paramaribo", "America/Phoenix", "America/Port_of_Spain", "America/Port-au-Prince", "America/Porto_Velho", "America/Puerto_Rico", "America/Punta_Arenas", "America/Rainy_River", "America/Rankin_Inlet", "America/Recife", "America/Regina", "America/Resolute", "America/Rio_Branco", "America/Santarem", "America/Santiago", "America/Santo_Domingo", "America/Sao_Paulo", "America/Scoresbysund", "America/Sitka", "America/St_Barthelemy", "America/St_Johns", "America/St_Kitts", "America/St_Lucia", "America/St_Thomas", "America/St_Vincent", "America/Swift_Current", "America/Tegucigalpa", "America/Thule", "America/Thunder_Bay", "America/Tijuana", "America/Toronto", "America/Tortola", "America/Vancouver", "America/Whitehorse", "America/Winnipeg", "America/Yakutat", "America/Yellowknife", "Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Macquarie", "Antarctica/Mawson", "Antarctica/McMurdo", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/Syowa", "Antarctica/Troll", "Antarctica/Vostok", "Arctic/Longyearbyen", "Asia/Aden", "Asia/Almaty", "Asia/Amman", "Asia/Anadyr", "Asia/Aqtau", "Asia/Aqtobe", "Asia/Ashgabat", "Asia/Atyrau", "Asia/Baghdad", "Asia/Bahrain", "Asia/Baku", "Asia/Bangkok", "Asia/Barnaul", "Asia/Beirut", "Asia/Bishkek", "Asia/Brunei", "Asia/Chita", "Asia/Choibalsan", "Asia/Colombo", "Asia/Damascus", "Asia/Dhaka", "Asia/Dili", "Asia/Dubai", "Asia/Dushanbe", "Asia/Famagusta", "Asia/Gaza", "Asia/Hebron", "Asia/Ho_Chi_Minh", "Asia/Hong_Kong", "Asia/Hovd", "Asia/Irkutsk", "Asia/Jakarta", "Asia/Jayapura", "Asia/Jerusalem", "Asia/Kabul", "Asia/Kamchatka", "Asia/Karachi", "Asia/Kathmandu", "Asia/Khandyga", "Asia/Kolkata", "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Kuching", "Asia/Kuwait", "Asia/Macau", "Asia/Magadan", "Asia/Makassar", "Asia/Manila", "Asia/Muscat", "Asia/Nicosia", "Asia/Novokuznetsk", "Asia/Novosibirsk", "Asia/Omsk", "Asia/Oral", "Asia/Phnom_Penh", "Asia/Pontianak", "Asia/Pyongyang", "Asia/Qatar", "Asia/Qostanay", "Asia/Qyzylorda", "Asia/Riyadh", "Asia/Sakhalin", "Asia/Samarkand", "Asia/Seoul", "Asia/Shanghai", "Asia/Singapore", "Asia/Srednekolymsk", "Asia/Taipei", "Asia/Tashkent", "Asia/Tbilisi", "Asia/Tehran", "Asia/Thimphu", "Asia/Tokyo", "Asia/Tomsk", "Asia/Ulaanbaatar", "Asia/Urumqi", "Asia/Ust-Nera", "Asia/Vientiane", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yangon", "Asia/Yekaterinburg", "Asia/Yerevan", "Atlantic/Azores", "Atlantic/Bermuda", "Atlantic/Canary", "Atlantic/Cape_Verde", "Atlantic/Faroe", "Atlantic/Madeira", "Atlantic/Reykjavik", "Atlantic/South_Georgia", "Atlantic/St_Helena", "Atlantic/Stanley", "Australia/Adelaide", "Australia/Brisbane", "Australia/Broken_Hill", "Australia/Currie", "Australia/Darwin", "Australia/Eucla", "Australia/Hobart", "Australia/Lindeman", "Australia/Lord_Howe", "Australia/Melbourne", "Australia/Perth", "Australia/Sydney", "Europe/Amsterdam", "Europe/Andorra", "Europe/Astrakhan", "Europe/Athens", "Europe/Belgrade", "Europe/Berlin", "Europe/Bratislava", "Europe/Brussels", "Europe/Bucharest", "Europe/Budapest", "Europe/Busingen", "Europe/Chisinau", "Europe/Copenhagen", "Europe/Dublin", "Europe/Gibraltar", "Europe/Guernsey", "Europe/Helsinki", "Europe/Isle_of_Man", "Europe/Istanbul", "Europe/Jersey", "Europe/Kaliningrad", "Europe/Kiev", "Europe/Kirov", "Europe/Lisbon", "Europe/Ljubljana", "Europe/London", "Europe/Luxembourg", "Europe/Madrid", "Europe/Malta", "Europe/Mariehamn", "Europe/Minsk", "Europe/Monaco", "Europe/Moscow", "Europe/Oslo", "Europe/Paris", "Europe/Podgorica", "Europe/Prague", "Europe/Riga", "Europe/Rome", "Europe/Samara", "Europe/San_Marino", "Europe/Sarajevo", "Europe/Saratov", "Europe/Simferopol", "Europe/Skopje", "Europe/Sofia", "Europe/Stockholm", "Europe/Tallinn", "Europe/Tirane", "Europe/Ulyanovsk", "Europe/Uzhgorod", "Europe/Vaduz", "Europe/Vatican", "Europe/Vienna", "Europe/Vilnius", "Europe/Volgograd", "Europe/Warsaw", "Europe/Zagreb", "Europe/Zaporozhye", "Europe/Zurich", "Indian/Antananarivo", "Indian/Chagos", "Indian/Christmas", "Indian/Cocos", "Indian/Comoro", "Indian/Kerguelen", "Indian/Mahe", "Indian/Maldives", "Indian/Mauritius", "Indian/Mayotte", "Indian/Reunion", "Pacific/Apia", "Pacific/Auckland", "Pacific/Bougainville", "Pacific/Chatham", "Pacific/Chuuk", "Pacific/Easter", "Pacific/Efate", "Pacific/Enderbury", "Pacific/Fakaofo", "Pacific/Fiji", "Pacific/Funafuti", "Pacific/Galapagos", "Pacific/Gambier", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Honolulu", "Pacific/Kiritimati", "Pacific/Kosrae", "Pacific/Kwajalein", "Pacific/Majuro", "Pacific/Marquesas", "Pacific/Midway", "Pacific/Nauru", "Pacific/Niue", "Pacific/Norfolk", "Pacific/Noumea", "Pacific/Pago_Pago", "Pacific/Palau", "Pacific/Pitcairn", "Pacific/Pohnpei", "Pacific/Port_Moresby", "Pacific/Rarotonga", "Pacific/Saipan", "Pacific/Tahiti", "Pacific/Tarawa", "Pacific/Tongatapu", "Pacific/Wake", "Pacific/Wallis"}
	return tz
}

func getAPICall(url string) string {
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

func patchRestCall(url string, data []byte) string {

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

func initializeArrayPage() ui.Control {
	//fields for the form
	arrayName := ui.NewEntry()
	eulaOrg := ui.NewEntry()
	eulaName := ui.NewEntry()
	eulaTitle := ui.NewEntry()
	eulaAccept := ui.NewCheckbox("yes")
	ntpServer := ui.NewEntry()
	//timeZone := ui.NewEntry()
	//set default timezone
	//timeZone.SetText("America/Los_Angeles")
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

	//sets the initResults console to readonly
	initResult.SetReadOnly(true)
	//multiline field for showing results of patch api call and form validation messages.
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
		result := getAPICall("http://" + ipAddress + ":8081/array-initial-config")

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
		//validate TimeZone
		/*err6 := validate.Var(timeZone.Text, "required")
		if err6 != nil {
			initResult.SetText("Please provide a valid Timezone")
			passed = false
		}*/
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
		/*err1 := validate.Var(arrayName.Text(), "required")
		if err1 != nil {
			initResult.SetText("Please provide the Array Name")
			passed = false
		}*/
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
			resp := patchRestCall("http://"+tempIP.Text()+":8081/array-initial-config", FAData)

			//update the initResult field with response.
			initResult.SetText("JSON Response: \n" + resp)

		}
	})

	return hbox
}

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

	tab.Append("ZTP Flash Array", initializeArrayPage())
	tab.SetMargined(0, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
