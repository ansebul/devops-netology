# Домашнее задание к занятию "6.5. Elasticsearch"

## Задача 1

В ответе приведите:
- текст Dockerfile манифеста
````
FROM elasticsearch:7.17.4
ADD elasticsearch.yml /usr/share/elasticsearch/config/
````
- ссылку на образ в репозитории dockerhub

> [https://hub.docker.com/r/ansebul/ansebul/tags](https://hub.docker.com/r/ansebul/ansebul/tags)

- ответ `elasticsearch` на запрос пути `/` в json виде
````bash
vagrant@server1:/opt/stack$ curl -X GET 'http://localhost:9200/'
{
  "name" : "netology_test",
  "cluster_name" : "docker-cluster",
  "cluster_uuid" : "uUyY8pifQAW1FJgE-JhCYA",
  "version" : {
    "number" : "7.17.4",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "79878662c54c886ae89206c685d9f1051a9d6411",
    "build_date" : "2022-05-18T18:04:20.964345128Z",
    "build_snapshot" : false,
    "lucene_version" : "8.11.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
````


## Задача 2

Получите список индексов и их статусов, используя API и **приведите в ответе** на задание.
````bash
vagrant@server1:~$ curl -X GET 'http://localhost:9200/_cat/indices?v'
health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   .geoip_databases lqP3PTaHS7q3AODv_GQLUg   1   0         41            8     46.2mb         46.2mb
green  open   ind-1            QMPRkc6sQ7SfRJnTbPsj_A   1   0          0            0       226b           226b
yellow open   ind-3            Vogt_TdlTp2MO5x-WCyrXA   4   2          0            0       904b           904b
yellow open   ind-2            H_E8H8MTQ92egVwj1V5R_g   2   1          0            0       452b           452b


vagrant@server1:~$ curl -X GET 'http://localhost:9200/_cluster/health/ind-1?pretty'
{
  "cluster_name" : "docker-cluster",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 1,
  "active_shards" : 1,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}

vagrant@server1:~$ curl -X GET 'http://localhost:9200/_cluster/health/ind-2?pretty'
{
  "cluster_name" : "docker-cluster",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 2,
  "active_shards" : 2,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 2,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 50.0
}

vagrant@server1:~$ curl -X GET 'http://localhost:9200/_cluster/health/ind-3?pretty'
{
  "cluster_name" : "docker-cluster",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 4,
  "active_shards" : 4,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 8,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 50.0
}

````

Получите состояние кластера `elasticsearch`, используя API.

````bash
$ curl -XGET localhost:9200/_cluster/health/?pretty=true
{
  "cluster_name" : "netology_test",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 7,
  "active_shards" : 7,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 41.17647058823529
}
````

Как вы думаете, почему часть индексов и кластер находится в состоянии yellow?
> Статус "yellow" у индексов, которым не хватает реплик. Если добавить - будет зелёным.

## Задача 3

**Приведите в ответе** запрос API и результат вызова API для создания репозитория.
```bash
vagrant@server1:/opt/stack$ curl -XPOST localhost:9200/_snapshot/netology_backup?pretty -H 'Content-Type: application/json' -d'{"type": "fs", "settings": { "location":"/var/lib/elasticsearch/data/snapshots" }}'
{
  "acknowledged" : true
}
vagrant@server1:~$ curl -X GET 'http://localhost:9200/_snapshot/netology_backup?pretty'
{
  "netology_backup" : {
    "type" : "fs",
    "settings" : {
      "location" : "/var/lib/elasticsearch/data/snapshots"
    }
  }
}
```


Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.
```bash
curl -X PUT localhost:9200/test -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 1,  "number_of_replicas": 0 }}'

vagrant@server1:~$ curl -X GET 'http://localhost:9200/_cat/indices?v'
health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   .geoip_databases lqP3PTaHS7q3AODv_GQLUg   1   0         41            7     44.9mb         44.9mb
green  open   test             oBI_F71BTOWMTsk6Fwk4wg   1   0          0            0       226b           226b


```


**Приведите в ответе** список файлов в директории со `snapshot`ами.
```bash
vagrant@server1:~$ ls -l es-data/snapshots/
total 48
-rw-rw-r-- 1 vagrant root  1425 May 30 08:29 index-0
-rw-rw-r-- 1 vagrant root     8 May 30 08:29 index.latest
drwxrwxr-x 6 vagrant root  4096 May 30 08:29 indices
-rw-rw-r-- 1 vagrant root 29238 May 30 08:29 meta-tly1wS65TW2egvqIbkY8oQ.dat
-rw-rw-r-- 1 vagrant root   712 May 30 08:29 snap-tly1wS65TW2egvqIbkY8oQ.dat

```
Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

```bash
vagrant@server1:~$ curl -X GET 'http://localhost:9200/_cat/indices?v'
health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test-2           tJfCkuMiRee1o5HvRqQuvA   1   0          0            0       226b           226b
green  open   .geoip_databases lqP3PTaHS7q3AODv_GQLUg   1   0         41            7     44.9mb         44.9mb

```

**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

> Мешает нам системный индекс, ни закрыть, ни удалить. Поэтому восстанавливать будем всё, кроме него.
```bash
vagrant@server1:~$ curl -X POST localhost:9200/_snapshot/netology_backup/elasticsearch/_restore?pretty
{
  "error" : {
    "root_cause" : [
      {
        "type" : "snapshot_restore_exception",
        "reason" : "[netology_backup:elasticsearch/tly1wS65TW2egvqIbkY8oQ] cannot restore index [.geoip_databases] because an open index with same name already exists in the cluster. Either close or delete the existing index or restore the index under a different name by providing a rename pattern and replacement name"
      }
    ],
    "type" : "snapshot_restore_exception",
    "reason" : "[netology_backup:elasticsearch/tly1wS65TW2egvqIbkY8oQ] cannot restore index [.geoip_databases] because an open index with same name already exists in the cluster. Either close or delete the existing index or restore the index under a different name by providing a rename pattern and replacement name"
  },
  "status" : 500
}

```

Делаю так:
```bash
vagrant@server1:~$ curl -X POST localhost:9200/_snapshot/netology_backup/elasticsearch/_restore?pretty -H 'Content-Type: application/json' -d'{"indices": "test"}'
{
  "accepted" : true
}

vagrant@server1:~$ curl -X GET 'http://localhost:9200/_cat/indices?v'
health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test-2           tJfCkuMiRee1o5HvRqQuvA   1   0          0            0       226b           226b
green  open   .geoip_databases lqP3PTaHS7q3AODv_GQLUg   1   0         41            7     44.9mb         44.9mb
green  open   test             JS1Te02tSz-teCFiQDvVNA   1   0          0            0       226b           226b


```