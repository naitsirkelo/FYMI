# FYMI: Find Your Movie Info


### IMT2681 Project (Marius Håkonsen & Ole K Larsen)

![Build Status](https://img.shields.io/badge/build-deployed-green.svg)


## Overview
The plan for the FYMI bot was to create a service that would allow the user to access useful information about an IMDb title. The service should be available straight from a smartphone without using a browser, and should only require one line of text to get the information you need.
This feature is implemented using the Slack application with custom slash, "/", commands, usable from any phone, tablet or browser.


## Future updates
* The application is currently only available on a private Slack channel, but will in the future be open for any Slack user to install to their channel.
* Implementation of MongoDB to allow users to see a log of previously searched titles, and create their own list of watched movies and series.
* Allow user to specify what kind of information that will be returned to a given Slack channel.


## Assumptions
Every call to the created API will utilize the same user key provided by the OMDb API structure, which allows 1000 requests every 24 hours.


## Usage

##### /fymihelp
Writes all available commands to the Slack channel.
##### /fymiid    <IMDb movie id>    
Returns useful info from a movie with corresponding ID.
(Example: tt0068646)
##### /fymititle <IMDb movie title>     
Gets the first resulting movie containing the title.
(Example: "The Godfather")
##### /fymisearch <IMDb movie title>
Shows a drop-down list of movies containing the title, giving the user the option to find the correct choice between the 10 first movies found.
(Example: "Star Wars")


## Reflection

#### What went well

We began the project by setting up a framework that deployed to Heroku and allowed for custom http handlers. The modularity of the project and integrating the Slack application was then set up without any major difficulties, since we at this stage already had a working deployment.
The transfer to the mobile version of Slack was a success when the information that is returned to the user was made more compact and practical.

#### Challenging aspects

Constructing the correct JSON structure to make sure the Slack would respond with a drop-down list was the most challenging part, and this was of great value as a learning experience in handling JSON.

#### Work hours: 41


## Tools
* Golang
* OMDb API
* Heroku
* Slack
