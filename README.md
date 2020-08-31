我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

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
  "time" : <ts>
}

or

[
 {
    "measurement" : "<measurement>",
    "tags" : [ { name = "<tag_1>", value = "<t_val_1>"} , ... ],
    "value" : <value>,
    "time" : <ts>
 },
...
]
```

### Update

```
PUT /write/<dbname>/<measurement>/<time>

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
GET /query/<dbname>?include_value=<true|false>&include_tags=<tag_list|true|false>&filters=<tag_name=tag_value,...>&time_from=<from>&time_to=<to>&measurements=<measurement_1,...>
GET /query/<dbname>/<measurement>?include_value=<true|false>&include_tags=<tag_list|true|false>&filters=<tag_name=tag_value,...>&time_from=<from>&time_to=<to>

{
    "results": [
        {
            "series": [
                {
                    "measurement": "<measurement>",
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

