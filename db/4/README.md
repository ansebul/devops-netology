# Домашнее задание к занятию "6.4. PostgreSQL"

## Задача 1

**Найдите и приведите** управляющие команды для:
- вывода списка БД
>  \l
 
- подключения к БД
> \c
 
- вывода списка таблиц
> \dt
 
- вывода описания содержимого таблиц
> \d+
 
- выхода из psql
> \q

## Задача 2

Используя таблицу pg_stats, найдите столбец таблицы orders с наибольшим средним значением размера элементов в байтах.
**Приведите в ответе** команду, которую вы использовали для вычисления и полученный результат.
> Как мы видим, это столбец `title`
```SQL
test_database=# select tablename, attname, avg_width from pg_stats where tablename='orders';
 tablename | attname | avg_width 
-----------+---------+-----------
 orders    | id      |         4
 orders    | title   |        16
 orders    | price   |         4
(3 rows)
```

## Задача 3

Вам предложили провести разбиение таблицы на 2 (шардировать на orders_1 - price>499 и orders_2 - price<=499).

Предложите SQL-транзакцию для проведения данной операции.

```SQL
BEGIN;
ALTER TABLE orders rename to orders_old;
CREATE TABLE orders (id integer NOT NULL, title character varying(80) NOT NULL, price integer DEFAULT 0) PARTITION BY RANGE (price);
CREATE TABLE orders_1 PARTITION OF orders FOR VALUES FROM ('499') TO (MAXVALUE);
CREATE TABLE orders_2 PARTITION OF orders FOR VALUES FROM ('0') TO ('499');
INSERT INTO orders (id, title, price) SELECT * FROM orders_old;
DROP TABLE orders_old;
COMMIT;
```

Можно ли было изначально исключить "ручное" разбиение при проектировании таблицы orders?
> ~~Теоретически - можно, но практически - очень сложно. При проектировании ещё нет такого объёма данных и как их делить? 99%, что одну таблицу полную получим и остальные пустые.~~

#### Можно попробовать 2 варианта:
> - `PARTITION BY RANGE` с заранее созданными вручную партициями и дефолтной партицией для значений, которые не попадают в созданные вручную.;
````SQL
CREATE TABLE orders (...) PARTITION BY RANGE (price);
CREATE TABLE orders_1 PARTITION OF orders FOR VALUES FROM ('0') TO (300);
CREATE TABLE orders_2 PARTITION OF orders FOR VALUES FROM ('300') TO ('600');
CREATE TABLE orders_2 PARTITION OF orders FOR VALUES FROM ('600') TO ('999');
CREATE TABLE orders_def PARTITION OF orders DEFAULT;
````
> - `PARTITION BY RANGE` с раширением `pg_partman` для автоматического создания новых таблиц  

````SQL
CREATE TABLE orders  (...) PARTITION BY RANGE (price);
SELECT partman.create_parent('public.orders', 'price', 'partman', '300');
...
````
 


## Задача 4

Используя утилиту `pg_dump` создайте бекап БД `test_database`.

> docker exec -i test-postgresql pg_dump -Cc  -Fp -U postgres -p5432 -h localhost test_database -f /var/lib/postgresql/data/test_database.sql

Как бы вы доработали бэкап-файл, чтобы добавить уникальность значения столбца `title` для таблиц `test_database`?

```diff

--- ./test_database.sql 2022-05-24 17:07:24.825528286 +0300
+++ ./test_database_modifed.sql 2022-05-24 17:08:38.699358040 +0300
@@ -46,7 +46,7 @@
 
 CREATE TABLE public.orders (
     id integer NOT NULL,
-    title character varying(80) NOT NULL,
+    title character varying(80) UNIQUE NOT NULL,
     price integer DEFAULT 0
 )
 PARTITION BY RANGE (price);
@@ -62,7 +62,7 @@
 
 CREATE TABLE public.orders_1 (
     id integer NOT NULL,
-    title character varying(80) NOT NULL,
+    title character varying(80) UNIQUE NOT NULL,
     price integer DEFAULT 0
 );
 ALTER TABLE ONLY public.orders ATTACH PARTITION public.orders_1 FOR VALUES FROM (499) TO (MAXVALUE);
@@ -76,7 +76,7 @@
 
 CREATE TABLE public.orders_2 (
     id integer NOT NULL,
-    title character varying(80) NOT NULL,
+    title character varying(80) UNIQUE NOT NULL,
     price integer DEFAULT 0
 );
 ALTER TABLE ONLY public.orders ATTACH PARTITION public.orders_2 FOR VALUES FROM (0) TO (499);

```

