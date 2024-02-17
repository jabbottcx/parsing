
# Objective
This script automates the process of downloading sample events and preparing them for use in developing RegEx patterns. An additional benefit is the increase in the number of sample records-- from 2,000 available in the export UI to 50,000 using the [direct archive query API](https://coralogix.com/docs/direct-query-http-api/). Using a larger number of sample records allows for the development of more inclusive RegEx patterns.

# Dashboard
This [dashboard](https://onlineboutique.coralogix.com/#/dashboards/4SCEP7oosPIhxdgSiwyn3) is provided to highlight the common parsing conditions that need to be addressed. This dashboard is not directly tied to the output of this module, but is intended to be used in concert.

<img width="1908" alt="Screenshot 2024-02-16 at 2 27 37 PM" src="https://github.com/jabbottcx/parsing/assets/125903661/11fc0360-5fb5-406f-97e6-105047108345">


# Usage
<br>
1. This utility is written in Go, which must be installed: https://go.dev/doc/install.<br>
<br>
2. Clone this repo, or download all files herein to the same directory.<br>
<br>
3. Update the file "retrieveConfig" to include:<br>
&nbsp&nbsp&nbspa. the appropriate <strong>url</strong> for the target team (required)<br>
&nbsp&nbsp&nbspb. the <strong>privateKey</strong> for that team (required)<br>
&nbsp&nbsp&nbspc. the <strong>startDate</strong> and  <strong>endDate</strong><br><br>
  Optionally, you can revise the query parameters to target specific focus areas.<br><br>
&nbsp&nbsp&nbspd. the <strong>query</strong>, in the appropriate format as determined by the #syntax element. The default is DataPrime. The query can be used to target priority events, such as those without a message field.<br>
&nbsp&nbsp&nbspe. which <strong>tier</strong> should be queried.  Note that tiers other than archive allow fewer results (no more than 12,000)<br><br>
4. To run, in a terminal within the directory containing the repo contents, run the command "go run runAll.go".  This runs four scripts in sequence to extract the data, create one line per event, and de-escape the content.<br>
<br>

# Output
The output consistents of two files:<br><br>
<strong>FinalOutout</strong>, which contains the individual events. This file can be used in RegEx101 or other tool.<br> <br>
<strong>FinalAnalysis</strong>, which provides a summary of different JSON patterns of the individual events.<br><br> 
Each time the process is run, these files will be overwritten, so save them if you want to retain them. <br>
