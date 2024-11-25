# One Billion Row Challenge

## Introduction

Messing around with the one billion row challenge in Go

TLDR 
Take a data set of one billion rows that contains a weather station name and a temperature
Find the average temperature for each weather station
Find the min and max temperature for each weather station
output the results for each weather station in alphabetical order

## First pass

First pass at this problem I went with a simple approach.
Read the file line by line, split each line to get the weather station name and temperature
save the count of each weather station, the sum of temperatures at each weather station 
and then the min and max temperature at each weather station.

With this approach I don't need to read in all the data I can just store the station once and
then update the values as I move through each line if required.

Finally I can format the final string output by sorting the keys of the map alphabetically and then
appending the values to a string following the requested format station/min/max/avg

Initial Time: processing took: time.Duration=3m30.078314333  (210.069s)
