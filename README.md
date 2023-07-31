# Crypto Charts 
Periodically pulls OHLC data from a set of data sources. Generates simple line chart in SVG format and uploads it to a place of your choice.

## Purpose
Using canvas for visualizing a chart might be heavy for the browser especially if you have to render tens of charts on the same page. On top of that you have to fetch the data for each pair resulting in MBs of data and many requests. In environment where visualizing live data is not crucial server side SVG rendering sounds like a good idea.

## Info
Current implementation includes:
- bitstamp - data source. Requires to have .env file. Check the example in: /cmd/job/.env.example
- svg generating library
- google cloud storage - persistent store for SVGs
- crontab - runs the job every 5 minutes
- multistage Dockerfile for building and running the app as cron job

## Demo
7 day charts: [demo](https://ivivanov.github.io/crypto-charts/demo-7D.html)
