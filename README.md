# australia-game
A game about data.

Data:
* ABS population data (2011)
* Victorian government (median rents)
* AIHW life expectancy.

To build and run (requires Go 1.7+ for context):
```
cd src/quiz
go build .
./quiz
```

Quiz server runs on port 3350 (Ballarats post code).

Navigate to http://localhost:3350/ 

Todo: 
* More splash screens / transition animations
* Fix bugs
* More data sets
* Improve geocoding thru google maps api
