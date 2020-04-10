# Flash Array and Flash Blade Zero Touch Provisioners

Pure Storage Flash Array and Flash Blade Zero Touch Provisioner. This lighweight application allows engineers to remotely initialize a new FA or FB using this cross-platform console application.

## Getting Started

Download the application binary relevant to your desktop OS and execute it.  Fill out the fields and your off.

The single console application has two tabs: Flash Array and Flash Blade.

#### Flash Array

![alt text](assets/FA.PNG)

The ZTP Flash Array tab is comprised of 3 columns.  Columnm 1 is the general configurations sections, Column 2 is the IP Addressing section and Column 3 is the Action and Output section.

To begin you will enter the Temp DHCP IP of Flash Array in the "DHCP IP of Array" field in column 3.  You can then Query the Array to make sure you have connectivity.

Next you will fill out the remainder of the fileds in Column 1 and Column 2 according to the specifics of the Array and environment.

When ready to initialize, press the Initialize Array button in Column 3 and view the results in the output window below. 

#### Flash Blade

![alt text1](assets/FB.PNG)

Select the ZTP Flash Blade Tab.

The ZTP Flash Blade tab is comprised of 3 columns as well.  Column 1 provides the 12 Steps necessary to configure a Flash Blade in order of operation.  Each step contains a button that will display the relevant form and controls for the associated step.  Column 2 is the Form fields and controls to query and apply required configurations relevant to each of the 12 steps.  Column 3 provides the output of each action.

To begin you will can either query the Flash Blade to obtain the API versions and auto generate the API url for the "Array API URL" field or you can manually enter the API url in the "Array API URL" field.  If you manually enter the API URL you must include the version and have the format as [http://\<fqdn or IP to array\>/api/\<api version\>]

Once you have the API url in place you will need to provide the "API Token" that will be used to login to the array and generate the "x-auth-token" that will be used for the remainder of the configuration.

When you have both the API token and the API url in place, clicking the "Create Session" button should result in the API URL and x-auth-token labels in column 3 displaying the resulting values.

When the above is successful you are ready to move to Step 2 - Step 12 in succession.

### Prerequisites

Windows 10 or higher, Linux with UI, or MAC OS.

```
For Windows, download and launch the appropriate .exe, fill out the form and go.
```

```
For Linux, (desktop GUI is required) download and launch appropriate ZTP binary, set it to executable "chmod +x", execute it and fill out the form.
```

```
For MacOS, download the appropriate ZTP binary, open terminal, set the binary to executable "chmod +x", execute it and fill out the form.
```

## Screen Shots

FlashArray

![alt text](assets/FA.PNG)

FlashBlade

![alt text1](assets/FB.PNG)

## Built With

* [UI](https://github.com/andlabs/ui/) - The GUI framework used
* [Validator v9](https://gopkg.in/go-playground/validator.v9) - form validation
* [GOLang](https://golang.org/) - Core
* [Crossbuild](https://github.com/magJ/go-ui-crossbuild) - Compiled with Go UI Crossbuild


## Authors

* **Brandon Showers* - - [GIBLAB](https://gitlab.com/btshowers)  OR  [About Me](https://bshowers.com)


## Acknowledgments

* Thanks Simon Dodsley for the hackathon to get me to learn GOLang!

