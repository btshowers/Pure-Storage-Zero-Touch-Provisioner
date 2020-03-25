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
	//might use for auth query...
	//req.Header.Set("Authorization", "Bearer b7d03a6947b217efb6f3ec3bd3504582")

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

//used for v1 build. incorportated into the init page.
/*
func queryArrayPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	button := ui.NewButton("Run Query")
	hbox.Append(button, false)
	//hbox.Append(ui.NewCheckbox("Checkbox"), false)

	vbox.Append(ui.NewLabel("Enter the DHCP Ip of the Array below to test connectivity"), false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	group := ui.NewGroup("Entries")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	input := ui.NewEntry()
	input.SetText(ipAddress)
	queryResult := ui.NewMultilineEntry()
	queryResult.SetReadOnly(true)

	entryForm.Append("Array IP", input, false)
	entryForm.Append("Query Results", queryResult, true)

	button.OnClicked(func(*ui.Button) {
		queryResult.SetText("Processing please wait...")
		ipAddress = input.Text()
		//result := getAPICall("https://" + ipAddress + ":8081/array-initial-config")
		result := getAPICall(ipAddress)
		defer queryResult.SetText(result)

	})

	return vbox
}*/

func initializeArrayPage() ui.Control {
	//fields for the form
	arrayName := ui.NewEntry()
	eulaOrg := ui.NewEntry()
	eulaName := ui.NewEntry()
	eulaTitle := ui.NewEntry()
	eulaAccept := ui.NewCheckbox("yes")
	ntpServer := ui.NewEntry()
	timeZone := ui.NewEntry()
	//set default timezone
	timeZone.SetText("America/Los_Angeles")
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

	//for demo only
	//tempIP.SetText("https://pureapisim.azurewebsites.net/api/array-initial-config")

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

		//for demo purposes only
		//result := getAPICall(ipAddress)

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
		err6 := validate.Var(timeZone.Text(), "required")
		if err6 != nil {
			initResult.SetText("Please provide a valid Timezone")
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
		/*err1 := validate.Var(arrayName.Text(), "required")
		if err1 != nil {
			initResult.SetText("Please provide the Array Name")
			passed = false
		}*/
		var rxPat = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,54}[a-zA-Z0-9])?$`)
		if !rxPat.MatchString(arrayName.Text()) {
			initResult.SetText("ArrayName has invalid characters.")
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
			FA.Timezone = timeZone.Text()
			FA.SMTP.RelayHost = smtpRelay.Text()
			FA.SMTP.SenderDomain = smtpDomain.Text()
			FA.AlertEmails = ae
			FA.EulaAcceptance.Accepted = eulaAccept.Checked()
			FA.EulaAcceptance.AcceptedBy.FullName = eulaName.Text()
			FA.EulaAcceptance.AcceptedBy.Organization = eulaOrg.Text()

			//marshal (json encode) the map into a json string
			FAData, err := json.Marshal(FA)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			//make the rest call with the json payload and stores response
			resp := patchRestCall("http://"+tempIP.Text()+":8081/array-initial-config", FAData)

			//rest call for demo purposes.
			//resp := patchRestCall(tempIP.Text(), FAData)

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

	//tab.Append("Query Array", queryArrayPage())
	//tab.SetMargined(0, true)

	tab.Append("ZTP Flash Array", initializeArrayPage())
	tab.SetMargined(0, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
