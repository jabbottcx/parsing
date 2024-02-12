# python version is a non-working WIP.  Ignore at this time.

# Go Version
This downloads events via the API and sends them a file that is a single event per row.<br>
Install go.<br>
<br>
Clone this repo or copy all files herein to the same directory.<br>
<br>
Update the file "retrieveConfig" to be the appropriate URL for your team, and your private key.  Included are details for Online Boutique, do not share with customer.<br>
<br>
The number of records is 50K, which you can adjust in retrieveWorking.go.  This is possible as this uses the DataPrime API and queries the archive specifically.  Otherwise the limit is 12K. The query string can also be adjusted here along with time parameters for query.<br>  The query, which will be externalized, provides an opportunity to limit your output to only things that you want to target (e.g. only those events w/out a message field, or those events that are multi-year, or whatever).
<br>
To run, the command "go run runAll.go"<br>

# Lots To Do
[] Move config for number of rows and query details to retrieveWorking<br>
[] Adding argument for retrieveWorking to take a specific config file such as "<customer>Config"<br>
[] Merge into single multi-step function<br>
[] Extensive testing of RegEx patterns based on those designed from this output.<br>
[] Add summary analytics to include number of events w/message, num w/body, etc.

