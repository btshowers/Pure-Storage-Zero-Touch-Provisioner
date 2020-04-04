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

//Global Vars//
var mainwin *ui.Window
var ipAddress = ""
var xAuthToken = ""
var loginUrl = ""
var apiUrl = ""

//TODO not sure if this method is usefull.. may just move to the login button function
func urlParser(url string) []string {
	//returns a slice broken out by forward slash in the url
	parts := strings.Split(url, "/")
	return parts
}

//TODO consider returning a byte slice from all functions rather than string to enable
//parsing of the key/values to present a structured response vs the entire json.

//Go lang doesn't support option parameters so I had to create multiple function
//with the same purpose to take certain parameters.
//Get Rest Function takes 2 parameters: url and xauth token and returns a response string//
/*func getAPICallOrig(url string, xAuthToken string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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
	//log.Println(string(body))
	return string(body)
}*/

//Get rest Function takes 2 parameters but returns an http response object//
func getAPICall(url string, xAuthToken string) []byte {
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
func postAPICallLogin(url string, apiToken string) string {
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
	if len(resp.Header["x-auth-token"]) > 0 {
		s := resp.Header["x-auth-token"]
		t := strings.Replace(s[0], "[", "", -1)
		t = strings.Replace(t, "]", "", -1)
		xAuthToken = t
	}
	//fmt.Printf("%v", resp.Header["X-Auth-Token"])
	//log.Println(string(body))
	return string(body)
}

//Post rest function takes 2 parameters and returns a string//
/*func postAPICallORIG(url string, apiToken string) string {
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

	fmt.Printf("%v", resp.Header["X-Auth-Token"])
	//log.Println(string(body))
	return string(body)
}*/

//Post rest function takes 3 parameters and returns a string//
func postAPICall2(url string, xAuthToken string, data []byte) string {

	body := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-auth-token", xAuthToken)
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

//Post rest function takes 2 parameters returns a string//
/*func postAPICall(url string, apiToken string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		//return err.Error()
	}
	req.Header.Set("api-token", apiToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		//return err.Error()
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	//log.Println(string(body))
	return body
}*/

//patch rest function takes 3 parameters and returns a string response//
//patch function takes 3 parameters and returns string//
func patchAPICall(url string, xAuthToken string, data []byte) string {

	body := bytes.NewReader(data)

	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-auth-token", xAuthToken)
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

//delete rest function takes 2 parameters and returns a string//
//delete function takes 2 parameters and returns a string//
func deleteAPICall(url string, xAuthToken string) string {
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
	//log.Println(string(body))
	return string(body)
}

//This is the main worker function that builds out the form and interacts with the user responses//
func initializeArrayPage() ui.Control {
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
	//loginGroup.Hide()
	loginForm := ui.NewForm()
	loginForm.SetPadded(true)
	loginGroup.SetChild(loginForm)

	apiToken := ui.NewEntry()
	//variables
	//for demo only
	//apiToken.SetText("2PDoD5iaokKDwGh9uNqt1jpDTNpgshfiOzO643z5ch92Mwycl7veBA==")
	xAuthTokenField := ui.NewEntry()
	//apiUrl := ui.NewEntry() //URL for api endpoint
	loginSubmitButton := ui.NewButton("Create Session")
	getAPIVersionsButton := ui.NewButton("Query Versions")
	//loginApplyTokenButton := ui.NewButton("Apply X-Auth-Token")
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
	//arrayForm.Append("api token", apiToken, false)
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
	//arrayForm.Append("api token", apiToken, false)
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
	//lagFormNew.Append("", lagPostButton, false)
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
	//lagFormNew.Append("", lagPostButton, false)
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

	xAuthTokenLabel := ui.NewLabel("")
	apiUrlLabel := ui.NewLabel("")
	xAuthTokenField.SetReadOnly(true)
	//multiline field for showing results of patch api call and form validation messages.
	entryForm9.Append("API URL: ", apiUrlLabel, false)
	entryForm9.Append("X-Auth-Token", xAuthTokenLabel, false)

	//button1 := ui.NewButton("Query")
	//entryForm9.Append("", ui.NewLabel(""), false)

	//submit and go button
	//button2 := ui.NewButton("Initialize")

	//sets the initResults console to readonly
	initResult.SetReadOnly(true)
	//multiline field for showing results of patch api call and form validation messages.
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

	//Buttons from Forms//

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
			//apiUrlLabel.SetText(apiUrl)
			//make the rest call
			resp := getAPICall("https://"+managementIP.Text()+"/api/api_version", apiToken.Text())

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
			resp := postAPICallLogin(loginUrl+"/login", apiToken.Text())
			//set the response in the display of the app
			initResult.SetText(string(resp))
			//slice up the respose to parse the a-auth-token from the string
			/*parts := strings.Split(string(resp), `"`)
			//fmt.Println(parts[3])
			//set the fields to the token that was parsed
			xAuthToken = parts[3]*/
			xAuthTokenField.SetText(xAuthToken)
			xAuthTokenLabel.SetText(xAuthToken)
		}

	})

	arrayGetButton.OnClicked(func(*ui.Button) {
		result := getAPICall(apiUrl+"/arrays", xAuthToken)
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

			result := patchAPICall(apiUrl+"/arrays", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	dnsGetButton.OnClicked(func(*ui.Button) {
		result := getAPICall(apiUrl+"/dns", xAuthToken)
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

			result := patchAPICall(apiUrl+"/dns", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	//HWC Buttons
	hwcGetButton.OnClicked(func(*ui.Button) {
		//name := hwcName.Text()
		//form validation object instantiation
		var passed bool = true
		validate := validator.New()

		//validate ntp servers
		err2 := validate.Var(hwcName.Text(), "required")
		if err2 != nil {
			initResult.SetText("Please provide the Array name")
			passed = false
		}
		if passed {
			result := getAPICall(apiUrl+"/hardware-connectors?names="+hwcName.Text(), xAuthToken)
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
		/*err1 := validate.Var(hwcLaneSpeed.Text(), "required")
		if err1 != nil {
			initResult.SetText("Please provide the Lane Speed")
			passed = false
		}
		err3 := validate.Var(hwcPortCount.Text(), "required")
		if err1 != nil {
			initResult.SetText("Please provide the Port Count")
			passed = false
		}*/

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

			result := patchAPICall(apiUrl+"/hardware-connectors?names="+hwcName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})
	//Lag Buttons
	lagGetButton.OnClicked(func(*ui.Button) {

		result := getAPICall(apiUrl+"/link-aggregation-groups", xAuthToken)
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
		/*err3 := validate.Var(hwcPortCount.Text(), "required")
		if err1 != nil {
			initResult.SetText("Please provide the Port Count")
			passed = false
		}*/

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

			result := postAPICall2(apiUrl+"/link-aggregation-groups?names="+lagNameNew.Text(), xAuthToken, pNameSlice)
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

			result := patchAPICall(apiUrl+"/link-aggregation-groups?names="+lagNameExisting.Text(), xAuthToken, pNameSlice)
			initResult.SetText(result)
		}
	})

	lagDeleteButton.OnClicked(func(*ui.Button) {

		result := deleteAPICall(apiUrl+"/link-aggregation-groups?names="+lagNameDelete.Text(), xAuthToken)
		initResult.SetText(result)

	})

	subnetGetButton.OnClicked(func(*ui.Button) {
		result := getAPICall(apiUrl+"/subnets", xAuthToken)
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

			result := postAPICall2(apiUrl+"/subnets?names="+subnetName.Text(), xAuthToken, FBData)
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

			result := patchAPICall(apiUrl+"/subnets?names="+subnetName.Text(), xAuthToken, FBData)
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
			result := deleteAPICall(apiUrl+"/subnets?names="+subnetName.Text(), xAuthToken)
			initResult.SetText(result)
		}
	})

	nicGetButton.OnClicked(func(*ui.Button) {
		result := getAPICall(apiUrl+"/network-interfaces", xAuthToken)
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

			//slices for multiple entry fields

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

			result := patchAPICall(apiUrl+"/network-interfaces?names="+nicName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	smtpGetButton.OnClicked(func(*ui.Button) {
		result := getAPICall(apiUrl+"/smtp", xAuthToken)
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

			//slices for multiple entry fields

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

			result := patchAPICall(apiUrl+"/smtp", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	supportGetButton.OnClicked(func(*ui.Button) {
		result := getAPICall(apiUrl+"/support", xAuthToken)
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

			//slices for multiple entry fields

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

			result := patchAPICall(apiUrl+"/support", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	awGetButton.OnClicked(func(*ui.Button) {
		result := getAPICall(apiUrl+"/alert-watchers", xAuthToken)
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

			//slices for multiple entry fields

			//initialize FAS struct object
			FB := &FAB{}
			FB.Enabled = awIsEnabled

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICall(apiUrl+"/alert-watchers?names="+awName.Text(), xAuthToken, FBData)
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
			result := deleteAPICall(apiUrl+"/alert-watchers?names="+awName.Text(), xAuthToken)
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

			//slices for multiple entry fields

			//initialize FAS struct object
			FB := &FAB{}
			FB.Enabled = awIsEnabled

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := postAPICall2(apiUrl+"/alert-watchers?names="+awName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	adminsGetButton.OnClicked(func(*ui.Button) {
		result := getAPICall(apiUrl+"/admins", xAuthToken)
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

			//slices for multiple entry fields

			//initialize FAS struct object
			FB := &FAB{}
			FB.CToken = adminsCreateTokenIsEnabled

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICall(apiUrl+"/admins?names="+adminName.Text(), xAuthToken, FBData)
			initResult.SetText(result)
		}
	})

	finalGetButton.OnClicked(func(*ui.Button) {
		result := getAPICall(apiUrl+"/setup/validation", xAuthToken)
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

			//slices for multiple entry fields

			//initialize FAS struct object
			FB := &FAB{}
			FB.Complete = finalSetupCompleteIsComplete

			//marshal (json encode) the map into a json string
			FBData, err := json.Marshal(FB)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := patchAPICall(apiUrl+"/setup/finalization", xAuthToken, FBData)
			initResult.SetText(result)
		}
	})
	//END Buttons//

	return hbox
}

//define and launch the gui //
func setupUI() {
	mainwin = ui.NewWindow("Pure Storage Zero Touch Provisioner for Flash Blade", 900, 480, true)
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

	tab.Append("ZTP Flash Blade", initializeArrayPage())
	tab.SetMargined(0, true)

	mainwin.Show()
}

//Main//
func main() {
	ui.Main(setupUI)
}
