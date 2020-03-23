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
}

func initializeArrayPage() ui.Control {
	//fields for the form
	arrayName := ui.NewEntry()
	eulaOrg := ui.NewEntry()
	eulaName := ui.NewEntry()
	eulaTitle := ui.NewEntry()
	eulaAccept := ui.NewCheckbox("yes")
	ntpServer := ui.NewEntry()
	timeZone := ui.NewEntry()
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

	//COMPANY INFO FIELDS//
	//define the group for the form
	/*
		group2 := ui.NewGroup("Company Info")
		group2.SetMargined(true)
		//add group to the vertical box
		vbox.Append(group2, true)
		//group1.SetChild(ui.NewNonWrappingMultilineEntry())
		//define the form for the group
		entryForm2 := ui.NewForm()
		entryForm2.SetPadded(true)
		//embed the company info form fields inside the second form group
		group2.SetChild(entryForm2)


	*/
	//NTP FIELDS//
	//group for NTP
	/*
		group7 := ui.NewGroup("NTP")
		group7.SetMargined(true)
		vbox.Append(group7, true)
		//defin the form for the group
		entryForm7 := ui.NewForm()
		entryForm7.SetPadded(true)
		//embed the ntp form fields inside the third form group
		group7.SetChild(entryForm7)
	*/
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

	//SMTP and DNS FORM//
	/*
		group8 := ui.NewGroup("Optional")
		group8.SetMargined(true)
		vbox.Append(group8, true)

		entryForm8 := ui.NewForm()
		entryForm8.SetPadded(true)
		group8.SetChild(entryForm8)

	*/
	//SUBMIT "GO" BUTTON//
	group9 := ui.NewGroup("Initialize Array")
	group9.SetMargined(true)
	vbox.Append(group9, true)

	entryForm9 := ui.NewForm()
	entryForm9.SetPadded(true)
	group9.SetChild(entryForm9)

	//submit and go button
	button2 := ui.NewButton("GO!")

	//used if initial prompt for ip is used.
	tempIP.SetText(ipAddress)

	entryForm9.Append("DHCP IP of Array ", tempIP, false)
	entryForm9.Append("READY, SET, ", button2, false)

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

	//initialize the array and do lots of other work
	button2.OnClicked(func(*ui.Button) {
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//responses from required text fields such as array name, EULA and ntpserver...
		responses := [6]string{arrayName.Text(), eulaOrg.Text(), eulaName.Text(), eulaTitle.Text(), ntpServer.Text(), timeZone.Text()}
		//responses from the IP config section for vi0, ct0, and ct1
		responsesIPs := [9]string{vir0IP.Text(), vir0SNM.Text(), vir0GW.Text(), ct0IP.Text(), ct0SNM.Text(), ct0GW.Text(), ct1IP.Text(), ct1SNM.Text(), ct1GW.Text()}
		//loop through these arrays and do form validation of being filled out and x number of characters... poor mans validator.
		for i, r := range responses {
			err := validate.Var(r, "required")
			if err != nil {
				fmt.Println(i)
				initResult.SetText("Please Fill out All Required Fields then presss GO again.")
				passed = false
			}
		}
		//add second array validation for ip addressing.
		for i, r := range responsesIPs {
			err := validate.Var(r, "required,gte=8")
			if err != nil {
				fmt.Println(i)
				initResult.SetText("Please Fill out All Required Fields then presss GO again.")
				passed = false
			}
		}
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

			/*for testing
			jsonStr := string(FAData)
			fmt.Println("The JSON data is:")
			fmt.Println(jsonStr)
			//ui.MsgBox(mainwin, "JSON To Send", jsonStr)
			initResult.SetText("JSON to send: \n" + jsonStr)
			*/

			//make the rest call with the json payload and stores response
			//resp := patchRestCall("http://"+tempIP.Text()+":8081/array-initial-config", FAData)
			resp := patchRestCall(tempIP.Text(), FAData)

			//update the initResult field with response.
			initResult.SetText("JSON Response: \n" + resp)

		}
	})

	return hbox
}

func setupUI() {
	mainwin = ui.NewWindow("Pure Storage Zero Touch Provisioner for Flash Array", 640, 480, true)
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

	tab.Append("Query Array", queryArrayPage())
	tab.SetMargined(0, true)

	tab.Append("Initialize Flash Array", initializeArrayPage())
	tab.SetMargined(2, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
