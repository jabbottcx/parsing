# python version is a non-working WIP.  Ignore at this time.

# Go Version
This downloads events via the API and sends them a file that is a single event per row.+
Install go.+
Copy all the files to the same directory.+
Update the file "retrieveConfig" to be the appropriate URL for your team, and your private key.  Included are details for Online Boutique, do not share with customer.+
The number of records is 50K, which you can adjust in retrieveWorking.go.  The query string can also be adjusted here along with time parameters for query.+
To run, the command "go run runAll.go"+

# Lots To Do
[] Move config for number of rows and query details to retrieveWorking
[] Adding argument for retrieveWorking to take a specific config file such as "<customer>Config"
[] Merge into single multi-step function
[] Extensive testing of patterns in Cx based on those designed from this output.

