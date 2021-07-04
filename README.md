# Airline Ticket Booking Application
   This project was generated with revel version 1.0.0.


## Requirements
* Go - https://golang.org/doc/install
* Revel Go Web Framework - https://github.com/revel
* MySQL - https://dev.mysql.com/downloads/installer/
* NodeJS (For npm Installer) - https://nodejs.org/en/download/
* Typescript - https://www.typescriptlang.org/download
* AngularJS - https://angularjs.org/
* Visual Studio Code (IDE) - https://code.visualstudio.com/download 
* Java JDK 1.8(Included in Ecplise) - https://www.oracle.com/sg/java/technologies/javase/javase-jdk8-downloads.html
* Ecplise (Running Test Cases) - https://www.eclipse.org/downloads/

## Included Jar files
* Selenium  - Located In \tests
* AngularJS - Located In \public\js
* Bootstrap - Located In \public\js

## How to Start the Web Server:

    revel run -a airlines

## Go to http://localhost:9000/ and you'll see:

    The Login Page for the application.

## Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory
	
    app_data/	  MySQL Database Dump
	
    messages/         Message files

    public/           Public static assets
            css/      CSS files
             js/      Javascript files
         images/      Image files
     typescript/      Typescript Javascript files

    tests/            Test suites
        airlineTest/  Selenium Test Suite
        jar files/    Jar files for setting up Selenium
	

## Help For revel installation.

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

## Typescript
### Install Typescript AngularJS typings
* Go node_modules folder and type in the command : npm install --save @types/angular
* [Help for AnuglarJS typings](https://www.npmjs.com/package/@types/angular)

### Command to compile typescript files 
    tsc -w

## Testing
* Download Ecplise and import airlinesTest into ecplise
* Testing file is done with java 1.8
* [Help with Ecplise project importing](https://dzone.com/articles/exporting-and-importing)
