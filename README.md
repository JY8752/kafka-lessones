# [Apache Kafka をはじめる](https://www.amazon.co.jp/Apache-Kafka%E3%82%92%E3%81%AF%E3%81%98%E3%82%81%E3%82%8B-%E6%8A%80%E8%A1%93%E3%81%AE%E6%B3%89%E3%82%B7%E3%83%AA%E3%83%BC%E3%82%BA%EF%BC%88NextPublishing%EF%BC%89-%E4%BD%90%E3%80%85%E6%9C%A8-%E5%81%A5%E5%A4%AA%E6%9C%97-ebook/dp/B0BDKQDJR7)の学習用レポジトリ


## トピックの作成

```
docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic ticket-order --partitions 3 --replication-factor 1
```