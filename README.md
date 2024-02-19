# GitHub Events Tracker
# Ohad Bar 18.2.24

This Go application fetches events from GitHub's API and stores them in a DB.

First Step:
Add your github private token to:
1. config constant.go file
2. docker-compose.yml
See below exmaple:
environment:
  - GITHUB_TOKEN=your_actual_github_token_here
and in the constants place. (constants.go - 	Token      = "" // Replace with your GitHub token)

Local build test:
1. Build:Run from root folder: go build -o myapp 
2. Run: Run from root folder:  ./myapp

Docker build: (shut down the server locally before)
1. open docker desktop
2. docker-compose build
3. docker-compose up
4. Go to http://localhost:8080/events for test.

FrontEnd Build and run:
1. Run from root folder: npm install
2. Run from root folder: npm start
3. Go to http://localhost:3000/




