# 6章 Kafka Connect

## Kafka Connect コンテナビルド

```
docker build -t ird-cp-kafka-connect:1.0.0 .
```

## プラグインの確認

```
% curl http://localhost:8083/connector-plugins/ | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   805  100   805    0     0   6207      0 --:--:-- --:--:-- --:--:--  7000
[
  {
    "class": "com.github.jcustenborder.kafka.connect.redis.RedisSinkConnector",
    "type": "sink",
    "version": "0.0.0.0"
  },
  {
    "class": "io.confluent.connect.jdbc.JdbcSinkConnector",
    "type": "sink",
    "version": "10.5.1"
  },
  {
    "class": "io.confluent.connect.s3.S3SinkConnector",
    "type": "sink",
    "version": "10.0.11"
  },
```

## Source Connectorの作成

```
curl -XPOST --url localhost:8083/connectors \
-H 'content-type: application/json' \
-d @create_source_connector.json

{"name":"postgresql-irdappdb-ticket-orders-source-connector","config":{"name":"postgresql-irdappdb-ticket-orders-source-connector","connector.class":"JdbcSourceConnector","connection.url":"jdbc:postgresql://source-database:5432/irdappdb","connection.user":"connectuser","connection.password":"secretdayo","topic.prefix":"source-postgresql-","table.whitelist":"ticket_orders","mode":"incrementing","incrementing.column.name":"order_id","validate.non.null":"true","transforms":"createKey,extractUserId","transforms.createKey.type":"org.apache.kafka.connect.transforms.ValueToKey","transforms.createKey.fields":"user_id","transforms.extractUserId.type":"org.apache.kafka.connect.transforms.ExtractField$Key","transforms.extractUserId.field":"user_id"},"tasks":[],"type":"source"}%     
```

## Connector一覧の確認

```
curl localhost:8083/connectors | jq
[
  "postgresql-irdappdb-ticket-orders-source-connector"
]
```

## Source Connectorの実行状況の確認

```
curl localhost:8083/connectors/postgresql-irdappdb-ticket-orders-source-connector/status | jq
{
  "name": "postgresql-irdappdb-ticket-orders-source-connector",
  "connector": {
    "state": "RUNNING",
    "worker_id": "localhost:8083"
  },
  "tasks": [
    {
      "id": 0,
      "state": "RUNNING",
      "worker_id": "localhost:8083"
    }
  ],
  "type": "source"
}
```

## DBにレコードをインサート

```
cat ./create-data.sql | docker exec -i source-database psql -U connectuser -d irdappdb
```

## Topick一覧を確認

```
docker exec cli kafka-topics --bootstrap-server broker-1:9092 --list 

__consumer_offsets
connect-config
connect-offsets
connect-status
source-postgresql-ticket_orders
```

## Topic内のイベントを確認

```
docker exec cli kafka-console-consumer --bootstrap-server broker-1:9092 --topic source-postgresql-ticket_orders --from-beginning --property print.key=true --property key.separator=" : "

1 : {"schema":{"type":"struct","fields":[{"type":"int32","optional":false,"field":"order_id"},{"type":"int32","optional":false,"field":"user_id"},{"type":"int32","optional":false,"field":"content_id"},{"type":"int64","optional":false,"name":"org.apache.kafka.connect.data.Timestamp","version":1,"field":"created_timestamp"}],"optional":false,"name":"ticket_orders"},"payload":{"order_id":1,"user_id":1,"content_id":1,"created_timestamp":1690362213547}}
2 : {"schema":{"type":"struct","fields":[{"type":"int32","optional":false,"field":"order_id"},{"type":"int32","optional":false,"field":"user_id"},{"type":"int32","optional":false,"field":"content_id"},{"type":"int64","optional":false,"name":"org.apache.kafka.connect.data.Timestamp","version":1,"field":"created_timestamp"}],"optional":false,"name":"ticket_orders"},"payload":{"order_id":2,"user_id":2,"content_id":3,"created_timestamp":1690362213547}}
1 : {"schema":{"type":"struct","fields":[{"type":"int32","optional":false,"field":"order_id"},{"type":"int32","optional":false,"field":"user_id"},{"type":"int32","optional":false,"field":"content_id"},{"type":"int64","optional":false,"name":"org.apache.kafka.connect.data.Timestamp","version":1,"field":"created_timestamp"}],"optional":false,"name":"ticket_orders"},"payload":{"order_id":3,"user_id":1,"content_id":2,"created_timestamp":1690362213547}}
3 : {"schema":{"type":"struct","fields":[{"type":"int32","optional":false,"field":"order_id"},{"type":"int32","optional":false,"field":"user_id"},{"type":"int32","optional":false,"field":"content_id"},{"type":"int64","optional":false,"name":"org.apache.kafka.connect.data.Timestamp","version":1,"field":"created_timestamp"}],"optional":false,"name":"ticket_orders"},"payload":{"order_id":4,"user_id":3,"content_id":3,"created_timestamp":1690362213547}}
4 : {"schema":{"type":"struct","fields":[{"type":"int32","optional":false,"field":"order_id"},{"type":"int32","optional":false,"field":"user_id"},{"type":"int32","optional":false,"field":"content_id"},{"type":"int64","optional":false,"name":"org.apache.kafka.connect.data.Timestamp","version":1,"field":"created_timestamp"}],"optional":false,"name":"ticket_orders"},"payload":{"order_id":5,"user_id":4,"content_id":6,"created_timestamp":1690362213547}}
```

## Sink Connector(Redis)の作成

```
curl -XPOST --url localhost:8083/connectors \
  -H 'content-type: application/json' \
  -d @create_sink_connector_redis.json

{"name":"redis-ticket-order-events-sink-connector","config":{"name":"redis-ticket-order-events-sink-connector","connector.class":"com.github.jcustenborder.kafka.connect.redis.RedisSinkConnector","topics":"source-postgresql-ticket_orders","tasks.max":"1","key.converter":"org.apache.kafka.connect.storage.StringConverter","value.converter":"org.apache.kafka.connect.storage.StringConverter","redis.hosts":"sink-database"},"tasks":[],"type":"sink"}
```

## Sink Connector(MinIO)の作成

```
curl -XPOST --url localhost:8083/connectors \
  -H 'content-type: application/json' \
  -d @create_sink_connector_s3.json
```