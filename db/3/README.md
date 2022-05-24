# Домашнее задание к занятию "6.3. MySQL"


## Задача 1

Найдите команду для выдачи статуса БД и **приведите в ответе** из ее вывода версию сервера БД.
> mysql  Ver 8.0.29 for Linux on x86_64 (MySQL Community Server - GPL)

**Приведите в ответе** количество записей с `price` > 300.

```SQL
mysql> SELECT COUNT(*) FROM orders WHERE price > 300;
+----------+
| COUNT(*) |
+----------+
|        1 |
+----------+
1 row in set (0.00 sec)
```

## Задача 2
    
Используя таблицу INFORMATION_SCHEMA.USER_ATTRIBUTES получите данные по пользователю `test` и 
**приведите в ответе к задаче**.

```SQL
mysql> SELECT * FROM INFORMATION_SCHEMA.USER_ATTRIBUTES WHERE USER = 'test' AND HOST = 'localhost'\G
*************************** 1. row ***************************
     USER: test
     HOST: localhost
ATTRIBUTE: {"fname": "James", "lname": "Pretty"}
1 row in set (0.00 sec)
```

## Задача 3

Исследуйте, какой `engine` используется в таблице БД `test_db` и **приведите в ответе**.
```SQL
mysql> SELECT ENGINE FROM information_schema.TABLES WHERE TABLE_NAME = 'orders';
+--------+
| ENGINE |
+--------+
| InnoDB |
+--------+
1 row in set (0.00 sec)
```

Измените `engine` и **приведите время выполнения и запрос на изменения из профайлера в ответе**:
- на `MyISAM`
- на `InnoDB`

```SQL
mysql> ALTER TABLE orders ENGINE = MyISAM;
Query OK, 5 rows affected (0.05 sec)
Records: 5  Duplicates: 0  Warnings: 0

mysql> ALTER TABLE orders ENGINE = InnoDB;
Query OK, 5 rows affected (0.06 sec)
Records: 5  Duplicates: 0  Warnings: 0

mysql> SHOW PROFILES;
+----------+------------+--------------------------------------------------------------------------+
| Query_ID | Duration   | Query                                                                    |
+----------+------------+--------------------------------------------------------------------------+
|       17 | 0.00262450 | SELECT ENGINE FROM information_schema.TABLES WHERE TABLE_NAME = 'orders' |
|       18 | 0.05336850 | ALTER TABLE orders ENGINE = MyISAM                                       |
|       19 | 0.06524325 | ALTER TABLE orders ENGINE = InnoDB                                       |
+----------+------------+--------------------------------------------------------------------------+
3 rows in set, 1 warning (0.00 sec)

```


## Задача 4 

Приведите в ответе измененный файл `my.cnf`.

````
root@71bcc6cd110a:/# cat /etc/mysql/my.cnf   

[mysqld]
pid-file        = /var/run/mysqld/mysqld.pid
socket          = /var/run/mysqld/mysqld.sock
datadir         = /var/lib/mysql
secure-file-priv= NULL
#
#  Скорость IO важнее сохранности данных
innodb_flush_method = O_DSYNC
#
#  Нужна компрессия таблиц для экономии места на диске
innodb_file_per_table = 1
#
#  Размер буффера с незакомиченными транзакциями 1 Мб
innodb_log_buffer_size = 1M
#
#  Буффер кеширования 30% от ОЗУ
innodb_buffer_pool_size = 200M
#
#  Размер файла логов операций 100 Мб
innodb_log_file_size = 100M

# Custom config should go here
!includedir /etc/mysql/conf.d/


````

