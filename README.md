# Couchbase Time Series Management Service

## High level requirements

| Id            | Subject            | Comment                                              |
| ------------- |:------------------:| ----------------------------------------------------:|
| 0             | Data point storage | It needs to be possible to store data points by time |
| 0.1           | Tuple character    | A data point is basically a tuple of numeric values (metrics), whereby the storage key is a time stamp |
| 1             | Data retrieval     | It needs to be possible to retrieve data points by time|
| 1.1           | Range queries      | You should be able to provide a time range in order to retrieve data points|
| 1.2           | Filter by metric   | It needs to be possible to provide a set of metrics you are interested in|
| 1.3           | Metric retrieval   | It should be possible to list all available metrics|

