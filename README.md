This is a project which takes a public facing API output and puts it into a filterable list.

**Project Parts**

**(1) BackEnd**
This part consists of a web server; it is a RESTful API written in Go utilizing the HTTP protocol.

The web server is public facing and its sole purpose is to return government job openings in Milwaukee, Wisconsin.  The web server calls another API (https://data.usajobs.gov : a public facing API which returns government job opportunities), filters that data, and returns job titles in Milwaukee, Wisconsin to whomever asks for them.  

**(2) FrontEnd**
This part consists of an Angular web application which is compatible with **Google Chrome, Firefox, Safari, Microsoft Edge, and Internet Explorer**.

The web application queries the BackEnd part of this project and receives any open government jobs in Milwaukee, Wisconsin.  It then displays those open government jobs by title.
