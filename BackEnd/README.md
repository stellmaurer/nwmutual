# BackEnd
This part consists of a web server; it is a RESTful API written in Go utilizing the HTTP protocol.

The web server is public facing and its sole purpose is to return government job openings in Milwaukee, Wisconsin. The web server calls another API (https://data.usajobs.gov : a public facing API which returns government job opportunities), filters that data, and returns job titles in Milwaukee, Wisconsin to whomever asks for them.

## How to run
Open a terminal.
```
  cd <Path to nwmutual/BackEnd>
  ./BackEnd
```
