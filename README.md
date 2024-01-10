# LAIN.
Loquacious Amigos Incessantly Nattering.<br>
A place for SJEC students to be open with their thoughts.

## Table of Contents
- [About](#about)
  - [Prerequisites](#prerequisites)
  - [Building](#building)
- [Usage](#usage)

## About
The website uses GO for backend paired with CockroachDB, CSS and very few javascript for Frontend. Styling the website was done from scratch without any frameworks<br>
We tried everything we could to deploy the website but there weren't any free services available for GO hosting, so for now, The website runs locally.


## Prerequisites
Clone the repository with
```bash
git clone https://github.com/nearlynithin/lain.git
```
You need to have CockroachDB installed, along with latest version of GO

## Building
Run cockroach.
```bash
cockroach start-single-node --insecure --http-addr=localhost:4000
```
In a new terminal, Run the go build command
```bash
go build ./cmd/lain && ./lain
```

## Usage
<P>There might be some css implmentation errors like logout button would come out of navgiation bar or the text area where you message would exceed the screen.  <br>In order to fix the error, keep the zoom value to given following browsers <br>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp * CHROME (ZOOM-100%) <BR> &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp * MICROSOFT EDGE (ZOOM-100%) <BR> &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp * MOZILLA FIREFOX (ZOOM 90%-100%)<BR> &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp <br> These errors were identified using only 3 browsers, if there is any css implementation error,try to zoom out and it depends upon the screen ratio too. </P> 
<h3>SHORT CUT KEYS</h3>
* To create post u can use TAB key <br> * To send post OR To send comment u can use ENTER key <BR> * 3)To remove text area or post area u can use ESCAPE key
