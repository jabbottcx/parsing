# python version is a non-working WIP.  Ignore at this time.

# Go Version
This downloads events via the API and sends them a file that is a single event per row.<br>
Install go.<br>
Copy all the files to the same directory.<br>
Update the file "retrieveConfig" to be the appropriate URL for your team, and your private key.  Included are details for Online Boutique, do not share with customer.<br>
The number of records is 50K, which you can adjust in retrieveWorking.go.  The query string can also be adjusted here along with time parameters for query.<br>
To run, the command "go run runAll.go"<br>

# Lots To Do
[] Move config for number of rows and query details to retrieveWorking<br>
[] Adding argument for retrieveWorking to take a specific config file such as "<customer>Config"<br>
[] Merge into single multi-step function<br>
[] Extensive testing of patterns in Cx based on those designed from this output.<br>

