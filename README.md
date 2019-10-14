This is a project which takes a public facing API output and puts it into a filterable list.

# How to run the project

NOTE: This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 7.3.8 and [npm](https://www.npmjs.com/get-npm) version 6.9.0. Using other versions of those packages in this project may result in issues.

## Running the BackEnd
Open up a terminal.
```
  cd <Path to nwmutual/BackEnd>
  ./BackEnd
```

## Running the FrontEnd
Make sure you have [npm](https://www.npmjs.com/get-npm) installed first. Open up a terminal.
```
  npm install -g @angular/cli@7.3.8
  cd <Path to nwmutual/FrontEnd>
  npm install
  ng serve
```
Now open up a compatible web browser, type `http://localhost:4200/` into the address bar, and wait for the website to load.

**Project Parts**

**(1) BackEnd**\
This part consists of a web server; it is a RESTful API written in Go utilizing the HTTP protocol.\
\
The web server is public facing and its sole purpose is to return government job openings in Milwaukee, Wisconsin.  The web server calls another API (https://data.usajobs.gov : a public facing API which returns government job opportunities), filters that data, and returns job titles in Milwaukee, Wisconsin to whomever asks for them.\
\
**(2) FrontEnd**\
This part consists of an Angular web application which is **compatible with Google Chrome, Firefox, Safari, Microsoft Edge, and Internet Explorer**.\
\
The web application queries the BackEnd part of this project and receives any open government jobs in Milwaukee, Wisconsin.  It then displays those open government jobs by title.
