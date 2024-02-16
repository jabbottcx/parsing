
# Objective
This is intended to automate the process for downloading and preparing events to use RegEx tools to help write RegEx patterns.  It allows for the download of up to 50K records (instead of only 2K with a manual download).  It does not automate the creation of RegEx patterns.

# Dashboard
This [dashboard](https://onlineboutique.coralogix.com/#/dashboards/4SCEP7oosPIhxdgSiwyn3) is provided to highlight the common parsing conditions that need to be addressed.  This dashboard is not directly tied to the output of this module, but is intended to be used in concert.

![image](https://private-user-images.githubusercontent.com/125903661/305536674-cf1652a1-809d-431e-9765-b73964ad584e.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MDgxMTIwNzgsIm5iZiI6MTcwODExMTc3OCwicGF0aCI6Ii8xMjU5MDM2NjEvMzA1NTM2Njc0LWNmMTY1MmExLTgwOWQtNDMxZS05NzY1LWI3Mzk2NGFkNTg0ZS5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjQwMjE2JTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI0MDIxNlQxOTI5MzhaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT03N2Y4NTQyMWE1MjE2MGRhYTdkZDFjZmFiYWM4MzQ1OWJkYzJkMWJlMmJhNDA4MTQzMTVmMzdlZjY5NjE5NzZkJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCZhY3Rvcl9pZD0wJmtleV9pZD0wJnJlcG9faWQ9MCJ9._yX2S7GYANqzuQoCQnefTtTnraFLcn0BXJppTazgsuw)

# This is written in Go, which you must have [Installed](https://go.dev/doc/install)<br>
<br>
Clone this repo or copy all files herein to the same directory.<br>
<br>
Update the file "retrieveConfig" to be the appropriate URL for your team, and your private key. Included are details for Online Boutique. The default number of records is 50K. This is possible as this uses the DataPrime API and queries the archive specifically. Otherwise the limit is 12K. The query string can also be adjusted here along with time parameters for query.<br> The query can be changed to limit what is retrieved to only things that you want to target (e.g. only those events w/out a message field, etc.).<br>
<br>
To run, the command "go run runAll.go"<br>

# Output
The file FinalOutput.json contains the events to be used in the RegEx editor in your process.  There is an additional file, FinalAnalysis.json, which provides a summary of different patterns found and the occurance count in the data set.

Each time the process is run, these files will be overwritten, so save them if you want to retain them.

# To Do
[] Extensive testing of RegEx patterns based on those designed from this output.<br>
