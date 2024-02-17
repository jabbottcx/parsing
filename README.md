
# Objective
This script automates the process of downloading sample events and preparing them for use in developing RegEx patterns. An additional benefit is the increase in the number of sample records-- from 2,000 available in the export UI to 50,000 using the [direct archive query API](https://coralogix.com/docs/direct-query-http-api/). Using a larger number of sample records allows for the development of more inclusive RegEx patterns.

# Dashboard
This [dashboard](https://onlineboutique.coralogix.com/#/dashboards/4SCEP7oosPIhxdgSiwyn3) is provided to highlight the common parsing conditions that need to be addressed. This dashboard is not directly tied to the output of this module, but is intended to be used in concert.

![image](https://private-user-images.githubusercontent.com/125903661/305536674-cf1652a1-809d-431e-9765-b73964ad584e.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MDgxMTIwNzgsIm5iZiI6MTcwODExMTc3OCwicGF0aCI6Ii8xMjU5MDM2NjEvMzA1NTM2Njc0LWNmMTY1MmExLTgwOWQtNDMxZS05NzY1LWI3Mzk2NGFkNTg0ZS5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjQwMjE2JTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI0MDIxNlQxOTI5MzhaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT03N2Y4NTQyMWE1MjE2MGRhYTdkZDFjZmFiYWM4MzQ1OWJkYzJkMWJlMmJhNDA4MTQzMTVmMzdlZjY5NjE5NzZkJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCZhY3Rvcl9pZD0wJmtleV9pZD0wJnJlcG9faWQ9MCJ9._yX2S7GYANqzuQoCQnefTtTnraFLcn0BXJppTazgsuw)

# Usage
<br>
1. This utility is written in Go, which you must install, available at: https://go.dev/doc/install. <br>
<br>
2. Clone this repo, or download all files herein to the same directory.<br>
<br>
3. Update the file "retrieveConfig" to include:<br>
&nbsp&nbsp&nbspa. the appropriate _url_ for the target team (required)<br>
&nbsp&nbsp&nbspb. the __privateKey__ for that team (required)<br>
&nbsp&nbsp&nbspc. the __startDate__ and __endDate__<br><br>
  Optionally, you can revise the query parameters to target specific focus areas.<br><br>
&nbsp&nbsp&nbspd. the __query__, in the appropriate format as determined by the #syntax element. The default is DataPrime. The query can be used to target priority events, such as those without a message field.<br>
&nbsp&nbsp&nbspe. which __tier__ should be queried.  Note that tiers other than archive allow fewer results (no more than 12,000)<br>
4. To run, in a terminal within the directory containing the repo contents, run the command "go run runAll.go".  This runs four scripts in sequence to extract the data, create one line per event, and de-escape the content.<br>
<br>
# Output
The output consistents of two files, one FinalOutout, which contains the individual events. This file can be used in RegEx101 or other tool.<br> <br>
The second file, FinalAnalysis, provides a summary of different JSON patterns of the individual events.<br> 
Each time the process is run, these files will be overwritten, so save them if you want to retain them. <br>

# To Do
[] Extensive testing of RegEx patterns based on those designed from this output.<br>
