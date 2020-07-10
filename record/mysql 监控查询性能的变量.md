# mysql 监控查询性能的变量

MySQL 的 SHOW STATUS 命令会展示两个容易混淆的数字，导致人们会问"这是什么"：

1. `Handler_read_rnd`
2. `Handler_read_rnd_next`

`Handler_read_rand` 计数的是函数`handler::rnd_pos()`的调用次数，这个函数会根据`fix position`来从表中获取一行，是一个随机读。`position`对于不同的存储引擎的含义是不同的。对于`MyISAM`，`position`意味着到文件头部开始的一个字节偏移量；对于`InnoDB`，这个意味着根据主键获取表中的一条记录。

`Handler_read_rnd_next` 表示的是函数`handler::rnd_next()`被调用的次数。这个基本上是一个游标的操作: 读取表(不是索引)中的`下一行`，这个操作会移动游标到下一次读取的位置，从而在下次调用的时候，可以获取到下一行。

为什么这两个函数会被调用？通常在排序操作进行的时候，会收集许多元组和其对应的`position`值，按照某种标准对这些元组进行排序(排序的操作应该不会是在获取表中所有记录的时候，再按照标准进行排序的，因为通常的排序字段比较少，如果把全部数据取出来进行排序的操作太耗费内存了)，然后会遍历排序好的列表，根据`position`(在此就是排序的标准的字段)来获取表记录的时候会进行`Handler_read_rnd`的操作。这个通常会发生在根据表中的随机指针获取行，虽然这个在数据全部都在内存的时候并不会造成随机的IO。通常发生在全表或者部分表的扫描的时候，会进行`Handler_read_rnd_next`的操作。

下面举一个`order by`的例子来详细说明下

```
 CREATE TABLE `test` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `data` varchar(32) DEFAULT NULL,
  `ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `file_sort` text,
  PRIMARY KEY (`id`),
  KEY `idx_data` (`data`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1
```

```
INSERT INTO `test` VALUES (1,'abc','2020-07-08 00:37:15','abcdefghijklmnopqrstuvwxyz'),(2,'abc','2020-07-08 00:37:15','bcdefghijklmnopqrstuvwxyza'),(3,'abd','2020-07-08 00:37:15','cdefghijklmnopqrstuvwxyzab'),(4,'acd','2020-07-08 00:37:15','defghijklmnopqrstuvwxyzabc'),(5,'def','2020-07-08 00:37:15','efghijklmnopqrstuvwxyzabcd'),(6,'pqr','2020-07-08 00:37:15','fghijklmnopqrstuvwxyzabcde'),(7,'stu','2020-07-08 00:37:15','ghijklmnopqrstuvwxyzabcdef'),(8,'vwx','2020-07-08 00:37:15','hijklmnopqrstuvwxyzabcdefg'),(9,'yza','2020-07-08 00:37:15','ijklmnopqrstuvwxyzabcdefgh'),(10,'def','2020-07-08 00:37:17','jklmnopqrstuvwxyzabcdefghi');
```

可以通过上面两个语句来进行创建表和插入数据的操作，再像下面一样执行操作：

```msyql
mysql> flush status;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from test order by file_sort desc;
+----+------+---------------------+----------------------------+
| id | data | ts                  | file_sort                  |
+----+------+---------------------+----------------------------+
| 10 | def  | 2020-07-07 17:37:17 | jklmnopqrstuvwxyzabcdefghi |
|  9 | yza  | 2020-07-07 17:37:15 | ijklmnopqrstuvwxyzabcdefgh |
|  8 | vwx  | 2020-07-07 17:37:15 | hijklmnopqrstuvwxyzabcdefg |
|  7 | stu  | 2020-07-07 17:37:15 | ghijklmnopqrstuvwxyzabcdef |
|  6 | pqr  | 2020-07-07 17:37:15 | fghijklmnopqrstuvwxyzabcde |
|  5 | def  | 2020-07-07 17:37:15 | efghijklmnopqrstuvwxyzabcd |
|  4 | acd  | 2020-07-07 17:37:15 | defghijklmnopqrstuvwxyzabc |
|  3 | abd  | 2020-07-07 17:37:15 | cdefghijklmnopqrstuvwxyzab |
|  2 | abc  | 2020-07-07 17:37:15 | bcdefghijklmnopqrstuvwxyza |
|  1 | abc  | 2020-07-07 17:37:15 | abcdefghijklmnopqrstuvwxyz |
+----+------+---------------------+----------------------------+
10 rows in set (0.00 sec)

mysql> show session status like "Handler_read%";
+-----------------------+-------+
| Variable_name         | Value |
+-----------------------+-------+
| Handler_read_first    | 1     |
| Handler_read_key      | 11    |
| Handler_read_last     | 0     |
| Handler_read_next     | 0     |
| Handler_read_prev     | 0     |
| Handler_read_rnd      | 10    |
| Handler_read_rnd_next | 11    |
+-----------------------+-------+
7 rows in set (0.00 sec)


```



