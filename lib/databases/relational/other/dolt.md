[dolt](https://github.com/dolthub/dolt)

*Description*: Dolt – Git for Data

*Labels*: #Go #Database

*Examples*:

```sql
mysql> create database getting_started;
Query OK, 1 row affected (0.04 sec)

mysql> use getting_started;
Database changed
mysql> create table employees (
    id int, 
    last_name varchar(255), 
    first_name varchar(255), 
    primary key(id));
Query OK, 0 rows affected (0.01 sec)

mysql> create table teams (
    id int, 
    team_name varchar(255), 
    primary key(id)); 
Query OK, 0 rows affected (0.00 sec)

mysql> create table employees_teams(
    team_id int, 
    employee_id int, 
    primary key(team_id, employee_id), 
    foreign key (team_id) references teams(id), 
    foreign key (employee_id) references employees(id));
Query OK, 0 rows affected (0.01 sec)

mysql> call dolt_add('teams', 'employees', 'employees_teams');
+--------+
| status |
+--------+
|      0 |
+--------+
1 row in set (0.03 sec)

mysql> call dolt_commit('-m', 'Created initial schema');
+----------------------------------+
| hash                             |
+----------------------------------+
| ne182jemgrlm8jnjmoubfqsstlfi1s98 |
+----------------------------------+
1 row in set (0.02 sec)

mysql> select * from dolt_log;
```
