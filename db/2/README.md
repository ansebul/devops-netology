# Домашнее задание к занятию "6.2. SQL"

## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 12) c 2 volume, 
в который будут складываться данные БД и бэкапы.

Приведите получившуюся команду или docker-compose манифест.

```yml
version: '3.1'

volumes:
  pg_data:
  pg_backup:

services:
  postgres:
    image: postgres:12.10
    restart: always
    environment:
      - POSTGRES_PASSWORD=rjyec
      - POSTGRES_USER=postgres
      - POSTGRES_DB=postgres
    volumes:
      - pg_data:/var/lib/postgresql/data
      - pg_backup:/var/lib/postgresql/backups
    ports:
      - ${POSTGRES_PORT:-5432}:5432

```


## Задача 2

- итоговый список БД после выполнения пунктов выше,

```
test_db=# \l
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 test_db   | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
(4 rows)
```
- описание таблиц (describe)

```
test_db=# \d+ orders
                                                        Table "public.orders"
 Column |          Type          | Collation | Nullable |              Default               | Storage  | Stats target | Description 
--------+------------------------+-----------+----------+------------------------------------+----------+--------------+-------------
 id     | integer                |           | not null | nextval('orders_id_seq'::regclass) | plain    |              | 
 title  | character varying(255) |           |          |                                    | extended |              | 
 price  | integer                |           |          |                                    | plain    |              | 
Indexes:
    "id_pkey" PRIMARY KEY, btree (id)
Access method: heap

test_db=# \d+ clients
                                                        Table "public.clients"
 Column  |          Type          | Collation | Nullable |               Default               | Storage  | Stats target | Description 
---------+------------------------+-----------+----------+-------------------------------------+----------+--------------+-------------
 id      | integer                |           | not null | nextval('clients_id_seq'::regclass) | plain    |              | 
 surname | character varying(255) |           |          |                                     | extended |              | 
 country | character varying(255) |           |          |                                     | extended |              | 
 order   | integer                |           |          |                                     | plain    |              | 
Indexes:
    "id_clients_pkey" PRIMARY KEY, btree (id)
    "country_idx" btree (country)
Access method: heap


```
- SQL-запрос для выдачи списка пользователей с правами над таблицами test_db
```sql
SELECT * FROM information_schema.table_privileges WHERE table_name IN ('orders', 'clients');
```

- список пользователей с правами над таблицами test_db

```
 grantor  |     grantee      | table_catalog | table_schema | table_name | privilege_type | is_grantable | with_hierarchy                                                                            
----------+------------------+---------------+--------------+------------+----------------+--------------+----------------
 postgres | postgres         | test_db       | public       | orders     | INSERT         | YES          | NO
 postgres | postgres         | test_db       | public       | orders     | SELECT         | YES          | YES
 postgres | postgres         | test_db       | public       | orders     | UPDATE         | YES          | NO
 postgres | postgres         | test_db       | public       | orders     | DELETE         | YES          | NO
 postgres | postgres         | test_db       | public       | orders     | TRUNCATE       | YES          | NO
 postgres | postgres         | test_db       | public       | orders     | REFERENCES     | YES          | NO
 postgres | postgres         | test_db       | public       | orders     | TRIGGER        | YES          | NO
 postgres | test-admin-user  | test_db       | public       | orders     | INSERT         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | SELECT         | NO           | YES
 postgres | test-admin-user  | test_db       | public       | orders     | UPDATE         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | DELETE         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | TRUNCATE       | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | REFERENCES     | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | TRIGGER        | NO           | NO
 postgres | test-simple-user | test_db       | public       | orders     | INSERT         | NO           | NO
 postgres | test-simple-user | test_db       | public       | orders     | SELECT         | NO           | YES
 postgres | test-simple-user | test_db       | public       | orders     | UPDATE         | NO           | NO
 postgres | test-simple-user | test_db       | public       | orders     | DELETE         | NO           | NO
 postgres | postgres         | test_db       | public       | clients    | INSERT         | YES          | NO
 postgres | postgres         | test_db       | public       | clients    | SELECT         | YES          | YES
 postgres | postgres         | test_db       | public       | clients    | UPDATE         | YES          | NO
 postgres | postgres         | test_db       | public       | clients    | DELETE         | YES          | NO
 postgres | postgres         | test_db       | public       | clients    | TRUNCATE       | YES          | NO
 postgres | postgres         | test_db       | public       | clients    | REFERENCES     | YES          | NO
 postgres | postgres         | test_db       | public       | clients    | TRIGGER        | YES          | NO
 postgres | test-admin-user  | test_db       | public       | clients    | INSERT         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | SELECT         | NO           | YES
 postgres | test-admin-user  | test_db       | public       | clients    | UPDATE         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | DELETE         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | TRUNCATE       | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | REFERENCES     | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | TRIGGER        | NO           | NO
 postgres | test-simple-user | test_db       | public       | clients    | INSERT         | NO           | NO
 postgres | test-simple-user | test_db       | public       | clients    | SELECT         | NO           | YES
 postgres | test-simple-user | test_db       | public       | clients    | UPDATE         | NO           | NO
 postgres | test-simple-user | test_db       | public       | clients    | DELETE         | NO           | NO
(36 rows)
```


## Задача 3

Используя SQL синтаксис - наполните таблицы следующими тестовыми данными

```sql
INSERT INTO orders (title, price) VALUES ('Шоколад', 10);
INSERT INTO orders (title, price) VALUES ('Принтер', 3000);
INSERT INTO orders (title, price) VALUES ('Книга', 500);
INSERT INTO orders (title, price) VALUES ('Монитор', 7000);
INSERT INTO orders (title, price) VALUES ('Гитара', 4000);

INSERT INTO clients (surname, country) VALUES ('Иванов Иван Иванович', 'USA');
INSERT INTO clients (surname, country) VALUES ('Петров Петр Петрович', 'Canada');
INSERT INTO clients (surname, country) VALUES ('Иоганн Себастьян Бах', 'Japan');
INSERT INTO clients (surname, country) VALUES ('Ронни Джеймс Дио', 'Russia');
INSERT INTO clients (surname, country) VALUES ('Ritchie Blackmore', 'Russia');
```

Используя SQL синтаксис - вычислите количество записей для каждой таблицы 

```
test_db=# SELECT COUNT(*) FROM orders;                                                                                                                                                               
 count                                                                                                                                                                                               
-------                                                                                                                                                                                              
     5                                                                                                                                                                                               
(1 row)                                                                                                                                                                                              
                                                                                                                                                                                                     
test_db=# SELECT COUNT(*) FROM clients;
 count 
-------
     5
(1 row)
```

## Задача 4

Часть пользователей из таблицы clients решили оформить заказы из таблицы orders.
Используя foreign keys свяжите записи из таблиц.

Приведите SQL-запросы для выполнения данных операций:
```sql
test_db=# UPDATE clients SET "order"=(SELECT id FROM orders WHERE title='Книга') WHERE surname='Иванов Иван Иванович';
UPDATE 1
test_db=# UPDATE clients SET "order"=(SELECT id FROM orders WHERE title='Монитор') WHERE surname='Петров Петр Петрович';
UPDATE 1
test_db=# UPDATE clients SET "order"=(SELECT id FROM orders WHERE title='Гитара') WHERE surname='Иоганн Себастьян Бах';
UPDATE 1
```

Приведите SQL-запрос для выдачи всех пользователей, которые совершили заказ, а также вывод данного запроса:

```sql
test_db=# SELECT c.surname, o.title FROM clients c INNER JOIN orders o ON c.order = o.id WHERE c.order IS NOT NULL;
       surname        |  title  
----------------------+---------
 Иванов Иван Иванович | Книга
 Петров Петр Петрович | Монитор
 Иоганн Себастьян Бах | Гитара
(3 rows)
```
 
## Задача 5

Получите полную информацию по выполнению запроса выдачи всех пользователей из задачи 4 
(используя директиву EXPLAIN).

Приведите получившийся результат и объясните что значат полученные значения.

```
test_db=# EXPLAIN SELECT c.surname, o.title FROM clients c INNER JOIN orders o ON c.order = o.id WHERE c.order IS NOT NULL;
                               QUERY PLAN                                
-------------------------------------------------------------------------
 Hash Join  (cost=11.57..24.20 rows=70 width=1032)
   Hash Cond: (o.id = c."order")
   ->  Seq Scan on orders o  (cost=0.00..11.40 rows=140 width=520)
   ->  Hash  (cost=10.70..10.70 rows=70 width=520)
         ->  Seq Scan on clients c  (cost=0.00..10.70 rows=70 width=520)
               Filter: ("order" IS NOT NULL)
(6 rows)
```

> План запроса говорит нам, что СУБД сначала сделает объединение Join с условием o.id = c."order", затем произведёт последовательное сканирование таблицы orders, потом clients, далее отфильтрует по условию ("order" IS NOT NULL). Мы видим, что никакие индексы не используются.



## Задача 6

Создайте бэкап БД test_db и поместите его в volume, предназначенный для бэкапов (см. Задачу 1).

Остановите контейнер с PostgreSQL (но не удаляйте volumes).

Поднимите новый пустой контейнер с PostgreSQL.

Восстановите БД test_db в новом контейнере.

Приведите список операций, который вы применяли для бэкапа данных и восстановления. 

```bash
docker-compose exec postgres pg_dump -U postgres -Cc -Fp -b -U postgres -v -p 5432 -h localhost test_db -f /var/lib/postgresql/backups/test_db.sql
docker-compose exec postgres /bin/bash -c 'pg_dumpall -r -U postgres -p5432 -h localhost | grep "test.*user" | tee /var/lib/postgresql/backups/roles.sql'

docker-compose stop
cd /var/lib/docker/volumes/stack_pg_data/_data && rm -fvr ./*
docker-compose -f /opt/stack/docker-compose.yaml up -d

docker-compose exec postgres /bin/bash -c 'psql -p 5432 -h localhost -U postgres -d postgres < /var/lib/postgresql/backups/roles.sql'
docker-compose exec postgres /bin/bash -c 'psql -p 5432 -h localhost -U postgres -d postgres < /var/lib/postgresql/backups/test_db.sql'

```
