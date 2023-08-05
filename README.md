# Crypto Charts 
Periodically pulls OHLC data from a set of data sources. Generates chart in SVG format and uploads it to a place of your choice.

## Purpose
Using canvas for visualizing a chart might be heavy for the browser especially if you have to render tens of charts on the same page. On top of that you have to fetch the data for each pair resulting in MBs of data and many requests. In environment where visualizing live data is not crucial server side SVG rendering sounds like a good idea.

## Use
```
go build .

```
```
crypto-charts --help 
```

### Config priority:
- All flags have default values
- Values provided from yaml config override default values
- Values provided via flags have hightest priority

## Info
Current implementation includes:
- fetchers:
    - bitstamp: pairs should be listed in .crypto-charts.yaml config
    - osmosis: wip
    - feel free to implement any other OHLC sources and pass them to the `job.NewJob` func
- svg generators
    - simple line charts
    - more advanced charts with grids, prices, BB, SMA, ...
- google cloud storage - persistent store for SVGs
- .crypto-charts.yaml - contains config for fetchers, generators and uploader
- ./deploy/Dockerfile - multistage Dockerfile for building and running the app as cron job
- ./deploy/crontab - runs the job every 5 minutes

## Demo
- 7 day simple charts: [demo](https://ivivanov.github.io/crypto-charts/demo-simple-7D-period.html)
- max period advanced: [demo](https://ivivanov.github.io/crypto-charts/demo-advanced-max-period.html)


## TODOs:
- extract all chart generators parameters in config
- add compose to build and deploy with multiple configurations
- add tests
- add build workflows
