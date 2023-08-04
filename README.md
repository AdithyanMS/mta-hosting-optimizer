<h1 align="center" id="title">mta-hosting-optimizer</h1>

<p id="description">This project implements a server which has a HTTP/REST endpoint used to retrieve hostnames having less or equals X active IP addresses. The Hosts and their respective data is fetched from sample_data/sample_data.go.</p>

<h2>Usage instructions:</h2>

<p>1. Download the repository</p>

```
git clone https://github.com/AdithyanMS/mta-hosting-optimizer.git
```

<p>2. cd into mta-hosting-optimizer</p>

```
cd mta-hosting-optimizer
```

<p>3. run the code</p>

```
go run .
```
<p></p>

That's it :) . now you have a server running at the port specified in main.go. <br>
<h3>Default value of X is 1. For changing this value, add a .env file in the project directory and set the desired value of X in variable MIN_IP_COUNT</h3>
.env

```
MIN_IP_COUNT=0
```
<h2>Endpoints:</h2>
<h3>/inefficient_hosts</h3> 
Request Method:"GET"<br>
Request Body: nil<br>
Request Parameters:nil <br>
Response: []string




