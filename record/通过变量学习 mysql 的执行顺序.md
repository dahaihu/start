# 通过变量学习 mysql 的执行顺序

《高性能mysql》中的6.7.9的自定义变量中有一句话让我有点懵逼，那就是:使用用户自定义变量的一个最常见的问题就是没有注意到在赋值和读取变量的时候可能是在不同的阶段。之后有一个对这句话的解释就是：这些出乎意料的结果可以在EXPLAIN语句中找到，注意看在Extra列中的'Using where', 'Using temporara'或者'Using filesort'。

说实话，反正我是没据explain中的extra列看懂赋值和读取变量到底是个什么顺序。但是还是可以大致猜测下主要的三个语句`select`,`where`和`order by`的执行顺序的。

**我觉得MySQL语句选取结果的过程可能如下：**

1.  **通过where条件筛选表中的数据行**
2.  **根据orderby的列进行排序**
3.  **通过select选取特定的列**

既然有了想法，那就可以开始验证了。测试用的是sakila数据库(MySQL官方提供的)，直接百度就可以安装和使用。

## 情况1：当在 select 和 where 的两个语句同时进行赋值的时候

```mysql
mysql> set @rownum := 0;
Query OK, 0 rows affected (0.00 sec)

mysql> select actor_id, first_name, @rownum:=@rownum+1 as cnt from actor where (@rownum:=@rownum+1) < 10;
+----------+------------+------+
| actor_id | first_name | cnt  |
+----------+------------+------+
|        1 | PENELOPE   |    2 |
|        2 | NICK       |    4 |
|        3 | ED         |    6 |
|        4 | JENNIFER   |    8 |
|        5 | JOHNNY     |   10 |
+----------+------------+------+
5 rows in set, 2 warnings (0.00 sec)
```

首先个人觉得肯定是先执行的where语句，因为肯定是先where语句对表进行筛选(where代表条件)，再进行select语句选取结果。如若不是这样，那么where语句(条件)有何意义？

假如不是首先执行 select ，那么满足条件的几条语句中的 cnt 应该均为奇数。所以肯定是where先执行，select后执行。

由于没有使用索引，所以数据是需要遍历每行找到满足条件的行的。在获取一行记录的过程中，需要使用`where`进行筛选，满足条件再进行`select`的操作。

## 情况2：当在 where 和 orderby 两个语句同时进行赋值的时候

```
mysql> set @rownum:=0;
Query OK, 0 rows affected (0.00 sec)

mysql> select actor_id, first_name, @rownum as rownum from actor where (@rownum:=@rownum+1) <= 9 order by first_name, least(0, @rownum:=@rownum+1);
+----------+------------+--------+
| actor_id | first_name | rownum |
+----------+------------+--------+
|        3 | ED         |      6 |
|        4 | JENNIFER   |      8 |
|        5 | JOHNNY     |     10 |
|        2 | NICK       |      4 |
|        1 | PENELOPE   |      2 |
+----------+------------+--------+
5 rows in set, 2 warnings (0.00 sec)
```

显然，这种情况下，`where` 中的赋值语句是先执行的。假如 `order by` 语句是先执行的，那么查询结果中应该返回的是 4 条记录而不是 5 条。变量 rownum = 10 的这一条记录就不会展示在结果中间了。

查询语句中使用了`least`的操作，在本次查询中`least`对于查询中的排序并不影响，仅仅是希望在排序的过程中完成变量的赋值操作。

通过 explain 语句可以知道，extra 列的结果如下：

```
mysql> explain select actor_id, first_name, @rownum as rownum from actor where (@rownum:=@rownum+1) <= 9 order by first_name, least(0, @rownum:=@rownum+1);
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+----------------------------------------------+
| id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra                                        |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+----------------------------------------------+
|  1 | SIMPLE      | actor | NULL       | ALL  | NULL          | NULL | NULL    | NULL |  200 |   100.00 | Using where; Using temporary; Using filesort |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+----------------------------------------------+
1 row in set, 3 warnings (0.01 sec)
```

由于在 `order by`的过程中添加了个赋值语句的过程，所以中间是多了个`Using temporary`的过程的。

## 情况3： 当只对select语句进行赋值的时候

```
mysql> select actor_id, first_name, (@rownum:=@rownum+1) as rownum from actor where @rownum <= 1 order by first_name;
+----------+-------------+--------+
| actor_id | first_name  | rownum |
+----------+-------------+--------+
|       71 | ADAM        |      1 |
|      132 | ADAM        |      2 |
|      165 | AL          |      3 |
|      173 | ALAN        |      4 |
|      125 | ALBERT      |      5 |
|      146 | ALBERT      |      6 |
|       29 | ALEC        |      7 |
|       65 | ANGELA      |      8 |
|      144 | ANGELA      |      9 |
|       76 | ANGELINA    |     10 |
```

这个结果是比较让人吃惊的，因为返回了全部数据记录(节选了部分展示结果)。`因为 order by 引入了文件排序，而 where 条件是在文件排序操作之前取值的，所以这个查询会返回表中的所有记录。`

通过 expalin 语句的 extra 列可以知道，对比情况 2 的时候，是少了`Using temporary`的过程的。

```
mysql> explain select actor_id, first_name, (@rownum:=@rownum+1) as rownum from actor where @rownum<=1 order by first_name;
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-----------------------------+
| id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra                       |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-----------------------------+
|  1 | SIMPLE      | actor | NULL       | ALL  | NULL          | NULL | NULL    | NULL |  200 |   100.00 | Using where; Using filesort |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-----------------------------+
1 row in set, 2 warnings (0.00 sec)
```



## 结论

三个特定语句的执行书序是：先进行where，再进行orderby，最后再进行select。

**MySQL关联执行的策略:MySQL对任何关联都执行嵌套循环关联操作，即MySQL先在一个表中循环取出单条数据，然后再嵌套循环到下一个表中寻找匹配的行，依次下去，直到找到所有表中匹配的行为止。然后根据各个表匹配的行，返回查询中需要的各个列。---摘自高性能MySQL中6.4.3的MySQL如何执行关联查询**



## 附录

《高性能MySQL》中的一个查询比较意思，本地测试的结果和书上的结果并不一致，差别甚至有点大

书：

```
mysql> set @rownum:=0;
mysql> select actor_id, @rownum as rownum from actor where (@rownum:=@rownum+1) <= 1;
+----------+------------+
| actor_id | rownum |
+----------+------------+
|        1 | 1      |
+----------+------------+

```

本地：

```
mysql> set @rownum:=0;
Query OK, 0 rows affected (0.00 sec)

mysql> select actor_id, @rownum as rownum from actor where (@rownum:=@rownum+1) <= 1;
+----------+--------+
| actor_id | rownum |
+----------+--------+
|       58 |      1 |
+----------+--------+
1 row in set, 1 warning (0.00 sec)
```

查看使用的索引的时候，发现了问题，查询的时候并没有使用索引`PRIMARY`，而是使用了索引`idx_actor_last_name`

```
mysql> explain select actor_id, @rownum as rownum from actor where (@rownum:=@rownum+1) <= 1;
+----+-------------+-------+------------+-------+---------------+---------------------+---------+------+------+----------+--------------------------+
| id | select_type | table | partitions | type  | possible_keys | key                 | key_len | ref  | rows | filtered | Extra                    |
+----+-------------+-------+------------+-------+---------------+---------------------+---------+------+------+----------+--------------------------+
|  1 | SIMPLE      | actor | NULL       | index | NULL          | idx_actor_last_name | 137     | NULL |  200 |   100.00 | Using where; Using index |
+----+-------------+-------+------------+-------+---------------+---------------------+---------+------+------+----------+--------------------------+
1 row in set, 2 warnings (0.00 sec)
```

要想和书上一致，则需要强制使用索引`PRIMARY`了。

```
mysql> set @rownum:=0;
Query OK, 0 rows affected (0.00 sec)

mysql> select actor_id, @rownum as rownum from actor force index (`PRIMARY`) where (@rownum:=@rownum+1) <= 1;
+----------+--------+
| actor_id | rownum |
+----------+--------+
|        1 |      1 |
+----------+--------+
1 row in set, 1 warning (0.00 sec)

mysql> explain select actor_id, @rownum as rownum from actor force index (`PRIMARY`) where (@rownum:=@rownum+1) <= 1;
+----+-------------+-------+------------+-------+---------------+---------+---------+------+------+----------+--------------------------+
| id | select_type | table | partitions | type  | possible_keys | key     | key_len | ref  | rows | filtered | Extra                    |
+----+-------------+-------+------------+-------+---------------+---------+---------+------+------+----------+--------------------------+
|  1 | SIMPLE      | actor | NULL       | index | NULL          | PRIMARY | 2       | NULL |  200 |   100.00 | Using where; Using index |
+----+-------------+-------+------------+-------+---------------+---------+---------+------+------+----------+--------------------------+
1 row in set, 2 warnings (0.00 sec)
```

