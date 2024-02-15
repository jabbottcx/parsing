
# Objective
This is intended to automate the process for downloading and preparing events to use RegEx tools to help write RegEx patterns.  It allows for the download of up to 50K records (instaed of only 2K with a manual download).  It does not automate the creation of RegEx patterns.

# This is written in Go, which you must have [Installed](https://go.dev/doc/install)<br>
<br>
Clone this repo or copy all files herein to the same directory.<br>
<br>
Update the file "retrieveConfig" to be the appropriate URL for your team, and your private key.  Included are details for Online Boutique, do not share with customer.<br>
<br>
The number of records is 50K, which you can adjust in retrieveWorking.go.  This is possible as this uses the DataPrime API and queries the archive specifically.  Otherwise the limit is 12K. The query string can also be adjusted here along with time parameters for query.<br>  The query, which will be externalized, provides an opportunity to limit your output to only things that you want to target (e.g. only those events w/out a message field, or those events that are multi-year, or whatever).<br>
<br>
To run, the command "go run runAll.go"<br>

# Output
The file FinalOutput.json contains the events to be used in the RegEx editor in your process.  There is an additional file, FinalAnalysis.json, which provides a summary of different patterns found and the occurance count in the data set.

Each time the process is run, these files will be overwritten, so save them if you want to retain them.

# To Do
[] Extensive testing of RegEx patterns based on those designed from this output.<br>
