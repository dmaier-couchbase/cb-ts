# Couchbase Time Series Management Service

I lastly implemented a demo service for vizalizing time series data from a Couchbase Bucket in Grafana:

* http://nosqlgeek.blogspot.de/2016/08/time-series-data-management-with.html
* http://nosqlgeek.blogspot.de/2016/09/visualizing-time-series-data-from.html

So one obvious consequence was to develop a more generic service for such purposes.


## High level requirements

The following table shows some high level requirements for such a service.

| Id            | Subject            | Comment                                              |
| ------------- |:------------------:| ----------------------------------------------------:|
| 0             | Data point storage | It needs to be possible to store data points by time |
| 0.1           | Tuple character    | A data point is basically a tuple of numeric values (measurement), whereby the storage key is a time stamp |
| 0.2            | Multi-point     | It needs to be possible to store multiple data points in one step|
| 0.3            | Tag support   | It should be possible to add tags to data points when storing them|
| 0.4            | Updates  | The same measurement at exactly the same time overrides an existing measurement|
| 1             | Data retrieval     | It needs to be possible to retrieve data points by time|
| 1.1           | Range queries      | You should be able to provide a time range in order to retrieve data points|
| 1.2           | Measurment support  | Multiple measurements should be supported|
| 1.2.1           | Filter by measurement   | It needs to be possible to provide a set of measurments you are interested in|
| 1.2.2           | Measurement retrieval   | It should be possible to list all available metrics|
| 1.3           | Tag (annotation) support  | It should be possible to tag measurements at a specific time|
| 1.3.1          | Filter by tags  | We need to be able to filter by tags|
| 1.3.2           | Tag retrieval  | It should be possible to list all tags those are belonging to a measurement|


## Is there a Standard?

It does not look like it. Looking on another time series database, the API looks for instance as the following one:

```
POST /write
db = <db_name>
data = '<measurement>,[<tag_1>=<t_val_1> ... <tag_n>=<t_val_n>], value=<value> <epoch ts> ...'

GET /query
db = <db_name>
pretty = [true|false]
q = SELECT <asterisk_value_or_tags> FROM <measurement> WHERE <condition_on_tags>
```

## API proposal

The following API seems to be suitable

### Connect

```
POST /init/<dbname>

{
   "admin" : "<admin>",
   "admin_pwd" : <"admin_pwd">,
   "db_pwd" : "<db_pwd>"
}

GET /<dbname>

{ 
  "db" : "<dbname>",
  "connected" : true|false,
}

GET /

{  "databases" : ["<db_name>", ...] }
```


### Create

```
POST /write/<dbname>

{
  "measurement" : "<measurement>",
  "tags" : [ { name = "<tag_1>", value = "<t_val_1>"} , ... ],
  "value" : <value>,
  "ts" : <ts>
}

or

[
 {
    "measurement" : "<measurement>",
    "tags" : [ { name = "<tag_1>", value = "<t_val_1>"} , ... ],
    "value" : <value>,
    "ts" : <ts>
 },
...
]
```

### Update

```
PUT /write/<dbname>/<measurement>/<ts>

{
  "tags" : [ { name = "<tag_1>", value = "<t_val_1>"} , ... ],
  "value" : <value>
}
```

### List

```
GET /list/<dbname>/measurements
GET /list/<dbname>/tags
```

### Query

```
GET /query/<dbname>?include_value=<true|false>&include_tags=<tag_list|true|false>&filters=<tag_name=tag_value,...>&ts_from=<from>&ts_to=<to>&measurements=<measurement_1,...>
GET /query/<dbname>/<measurement>?include_value=<true|false>&include_tags=<tag_list|true|false>&filters=<tag_name=tag_value,...>&ts_from=<from>&ts_to=<to>

{
    "results": [
        {
            "series": [
                {
                    "name": "<measurement>",
                    "columns": [
                        "time",
                        "value",
                        "<tag_1>",
                        ...
                    ],
                    "values": [
                        [
                            "<timestamp, eg. 2015-01-29T21:55:43.702900257Z>",
                            "<value, e.g. 0.55>",
                            "<tag_val_1>",
                            ...
                            
                        ],
                        ...
                    ]
                }
            ],
            ...
        }
    ]
}
```

