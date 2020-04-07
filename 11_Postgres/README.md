This exercise was basic SQL commands in Postgres. An example problem an solution are below.

===========================

Ex1) Problem Statement

# hands-on exercise
1. delete all of your current tables.
1. READ ALL OF THIS: create a new table called employees with these fields ```id, name, score, salary``` AND give ```score``` a default value of 10 AND have the ```id``` field automatically increment.
1. add these records and then show all of the records
```
 id |  name  | score | salary 
----+--------+-------+--------
  1 | Daniel |    23 |  55000
  2 | Arin   |    25 |  65000
  3 | Juan   |    24 |  72000
  4 | Shen   |    26 |  64000
  5 | Myke   |    27 |  58000
  6 | McLeod |    26 |  72000
  7 | James  |    32 |  35000
```

===========================

Ex1) Solution

```
DROP TABLE employees, phonenumbers;
```

```
CREATE TABLE employees (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           TEXT    NOT NULL,
   SCORE            INT     DEFAULT 10 NOT NULL,
   SALARY         REAL
);
```

```
INSERT INTO employees (NAME,SCORE,SALARY) VALUES ('Daniel', 23, 55000.00), ('Arin', 25, 65000.00), ('Juan', 24, 72000.00), ('Shen', 26, 64000.00), ('Myke', 27, 58000.00), ('McLeod', 26, 72000.00), ('James', 32, 35000.00);
```

```
SELECT * FROM employees;
```

===========================