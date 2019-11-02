# website-health-analyzer

This is an application which used for  analyzing  health of websites. It checks wether website is accessible or not. 

There is a scheduler added to check health of websites which run every 5 minutes.

## Start

You can start the whole stack using docker-compose  
`docker-compose up`  
This starts the web server, cron job to perform regular health checks and mongodb server

After this head on to [web-analyzer](http://localhost:9000/web/)

## Todo
Increase unit test coverage for the app

