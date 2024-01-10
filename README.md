<h1>LAIN.(SOCIAL MEDIA WEBSITE)</h1>
<h2>FRONTEND</h2>
<p>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp  This website frontend end design was made up of HTML ,CSS and JAVASCRIPT. Frameworks were used at beginining for backend purpose , then styling the website was done from scratch without any framework. </p>
<h3>NOTE</h3>
<P>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp 1)The website only works in laptop since we havent used media queries.  <br>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp 2)There might be some implmentation css error like logout button would come out of navgiation bar or the text area where you <br> &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp  message would exceed the screen.  <br>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp 3)Inorder to fix the error, keep the zoom value to given following browsers <br>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp -CHROME (ZOOM-100%) <BR> &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp -MICROSOFT EDGE (ZOOM-100%) <BR> &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp -FIREFOX (ZOOM 90%-100%)<BR> &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp 4)These errors were identified using only 3 browsers, if there is any css implementation error,try to zoom out and it depends upon <br>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp laptop also. </P> <h3>SHORT CUT KEYS</h3><P>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp 1)To create post u can use TAB key <br>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp 2)To send post OR To send comment u can use ENTER key <BR> &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp 3)To remove text area or post area u can use ESCAPE key </P><BR>
<h2>BACKEND</h2>
<P>&nbsp &nbsp &nbsp &nbsp &nbsp &nbsp This website backend part was done using GO and COCKROACH DB</P>

# LAIN.
Loquacious Amigos Incessantly Nattering.<br>
A place for SJEC students to be open with their thoughts.

## Table of Contents
- [About](#about)
  - [Prerequisites](#prerequisites)
  - [Building](#building)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

## About
The website uses GO for backend paired with CockroachDB, CSS and very few javascript for Frontend. Styling the website was done from scratch without any frameworks<br>
We tried everything we could to deploy the website but there weren't any free services available for GO hosting, so for now, The website runs locally.


### Prerequisites
Clone the repository with
git clone
You need to have CockroachDB installed, along with latest version of GO

### Building
Run cockroach.
```bash
cockroach start-single-node --insecure --http-addr=localhost:4000
```
In a new terminal, Run the go build command
```bash
go build ./cmd/lain && ./lain
```


```bash
# Example installation steps
git clone https://github.com/your-username/your-project.git
cd your-project
npm install
