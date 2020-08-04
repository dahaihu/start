# mysql 监控查询性能的变量 Handler_read*

## 变量详解

MySQL 的 SHOW STATUS 命令会展示两个容易混淆的数字，导致人们会问"这是什么"：

1. `Handler_read_rnd`
2. `Handler_read_rnd_next`

`Handler_read_rnd` 计数的是函数`handler::rnd_pos()`的调用次数，这个函数会根据`fix position`来从表中获取一行，是一个随机读。`position`对于不同的存储引擎的含义是不同的。对于`MyISAM`，`position`意味着到文件头部开始的一个字节偏移量；对于`InnoDB`，这个意味着根据主键获取表中的一条记录。

`Handler_read_rnd_next` 表示的是函数`handler::rnd_next()`被调用的次数。这个基本上是一个游标的操作: 读取表(不是索引)中的`下一行`，这个操作会移动游标到下一次读取的位置，从而在下次调用的时候，可以获取到下一行。

为什么这两个函数会被调用？通常在排序操作进行的时候，会收集许多元组和其对应的`position`值，按照某种标准(group by)对这些元组进行排序，然后会遍历排序好的列表，根据`position`来获取表记录的时候会进行`Handler_read_rnd`的操作(后文会讲到的 file\_sort排序算法之一)。这个通常会发生在根据表中的随机指针获取行，除了数据全部都在内存的时候会造成随机的IO。通常发生在全表或者部分表的扫描的时候，会进行`Handler_read_rnd_next`的操作。

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

上面的解释结果中`Handler_read_first`和`Handler_read_key`中的`1`以及`Handler_read_rnd_next`的`11`还是比较好理解的。

`Handler_read_first`和`Handler_read_last`：比较简单，一个是读取索引的第一个元素，另一个是读取索引的最后一个元素。

`Handler_read_key`: 也比较简单，就是读取索引的某一个元素。

`Handler_read_next`和`Handler_read_prev`：还是比较简单，因为索引是有序的，所以会在读取索引的时候，按照向前后者向后的顺序读取索引。

在这个查询展示的变量结果中就是，因为`select`的是表中的所有元素，所以`Handler_read_rnd_next`的结果是`11`(至于为什么不是10，因为游标得走11次才知道第11次是没有值的)。再由于有了索引，所以会利用索引读取到表中的第一行的记录(便于读取表中的后续记录)，于是就有了`Handler_read_first`的`1`和`Handler_read_key`中的`1`。还剩下两个值就是`Handler_read_key`中剩下的`10`，和`Handler_read_rnd`中的`10`。这两个值确实不好理解，因为鬼知道到底怎么查询的。那咱就需要知道这个`怎么查询`的具体过程了，因为使用到了排序，所以就需要知道 MySQL 的排序的具体实现了。

## 排序

MySQL 中的排序算法有两个，一个是 **Original Filesort Algorithm**， 另外一个是 **Modified Filesort Algorithm**。MySQL 的优化器会决定到底使用哪种算法。`通常对于涉及 BLOB 和 TEXT 类型的时候，使用 the original filesort algorithm，其他时候使用 modified filesort algorithm。`

### **The Original Filesort Algorithm**

此算法会取值主键和排序的键，根据排序的键对主键进行排序，然后根据主键的顺序从表中取行。只会读取行两次，一次是取主键和排序的键，另外一次是排序好后根据主键的顺序取值输出。

### The Modified Filesort Algorithm

此算法会取值所有需要的列，根据排序的键对取到的行进行排序，然后按照排序的顺序进行输出。只会读取行一次。

## 继续

有了排序的结果，就可以知道剩余的两个变量的值的原因了。由于`file_sort`列是 TEXT 类型，所以使用的是 **Original Filesort Algorithm**。主键和`file_sort`排序好之后，就会根据主键来进行取行，所以就有了`Handler_read_key`中的 10，和`Handler_read_rnd`中的10。到此全部变量就可以完美的解释了。

那么可以再增加一列为 VARCHAR，并删除原本的 TEXT 的列，验证下在没有 TEXT 的时候排序算法是否使用的是 **Modified Filesort Algorithm**。

```
mysql> alter table test add column test_column varchar(100) not null default '';

mysql> UPDATE test SET test_column = 'abcdefghijklmnopqrstuvwxyz' WHERE id = 1;
mysql> UPDATE test SET test_column = 'bcdefghijklmnopqrstuvwxyza' WHERE id = 2;
mysql> UPDATE test SET test_column = 'cdefghijklmnopqrstuvwxyzab' WHERE id = 3;
mysql> UPDATE test SET test_column = 'defghijklmnopqrstuvwxyzabc' WHERE id = 4;
mysql> UPDATE test SET test_column = 'efghijklmnopqrstuvwxyzabcd' WHERE id = 5;
mysql> UPDATE test SET test_column = 'fghijklmnopqrstuvwxyzabcde' WHERE id = 6;
mysql> UPDATE test SET test_column = 'ghijklmnopqrstuvwxyzabcdef' WHERE id = 7;
mysql> UPDATE test SET test_column = 'hijklmnopqrstuvwxyzabcdefg' WHERE id = 8;
mysql> UPDATE test SET test_column = 'ijklmnopqrstuvwxyzabcdefgh' WHERE id = 9;
mysql> UPDATE test SET test_column = 'jklmnopqrstuvwxyzabcdefghi' WHERE id = 10;

mysql> select * from test order by test_column;
+----+------+---------------------+----------------------------+
| id | data | ts                  | test_column                |
+----+------+---------------------+----------------------------+
|  1 | abc  | 2020-07-11 11:34:27 | abcdefghijklmnopqrstuvwxyz |
|  2 | abc  | 2020-07-11 11:34:27 | bcdefghijklmnopqrstuvwxyza |
|  3 | abd  | 2020-07-11 11:34:27 | cdefghijklmnopqrstuvwxyzab |
|  4 | acd  | 2020-07-11 11:34:27 | defghijklmnopqrstuvwxyzabc |
|  5 | def  | 2020-07-11 11:34:27 | efghijklmnopqrstuvwxyzabcd |
|  6 | pqr  | 2020-07-11 11:34:27 | fghijklmnopqrstuvwxyzabcde |
|  7 | stu  | 2020-07-11 11:34:27 | ghijklmnopqrstuvwxyzabcdef |
|  8 | vwx  | 2020-07-11 11:34:27 | hijklmnopqrstuvwxyzabcdefg |
|  9 | yza  | 2020-07-11 11:34:27 | ijklmnopqrstuvwxyzabcdefgh |
| 10 | def  | 2020-07-11 11:34:28 | jklmnopqrstuvwxyzabcdefghi |
+----+------+---------------------+----------------------------+
10 rows in set (0.00 sec)

mysql> show session status like 'Handler_read%';
+-----------------------+-------+
| Variable_name         | Value |
+-----------------------+-------+
| Handler_read_first    | 1     |
| Handler_read_key      | 1     |
| Handler_read_last     | 0     |
| Handler_read_next     | 0     |
| Handler_read_prev     | 0     |
| Handler_read_rnd      | 0     |
| Handler_read_rnd_next | 11    |
+-----------------------+-------+
7 rows in set (0.01 sec)
```

可以看到在删除列类型为 TEXT 的列`file_sort`之后，再按照相同的值的列`test_column`排序，输出的结果中就没有了第二次读取数据表的`Handler_read_key`和`Handler_read_rnd`的过程了。

## 参考文章

1. https://www.valinv.com/dev/mysql-mysql-filesort-algorithms
2. https://jin-yang.github.io/post/mysql-handler.html
3. http://www.fromdual.com/mysql-handler-read-status-variables
4. https://www.percona.com/blog/2010/06/15/what-does-handler_read_rnd-mean/



