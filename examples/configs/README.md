# Postgres

Running test Postgres in Docker:

```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=backlokto -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=backlokto -d postgres
```

Connecting to test Postgres:

```bash
docker exec -it postgres psql -U backlokto -d backlokto
```

Create table for test data:

```sql
CREATE TABLE backlokto_data (
  column1 VARCHAR(255),
  column2 VARCHAR(255),
  column3 INT
); 
```

Insert test data:

```sql
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('FIRST_ROW_1', 'FIRST_ROW_2', 10);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('SECOND_ROW_1', 'SECOND_ROW_2', 15);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('THIRD_ROW_1', 'THIRD_ROW_2', 20);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('FOURTH_ROW_1', 'FOURTH_ROW_2', 25);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('FIFTH_ROW_1', 'FIFTH_ROW_2', 30);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('SIXTH_ROW_1', 'SIXTH_ROW_2', 35); 
```

# MySQL

Running test MySQL in Docker and logging in (for password promt just press enter):

```bash
docker run --name mysql -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -d mysql
docker exec -it mysql mysql -u root -p
```

Creating database, user and password:

```sql
CREATE DATABASE IF NOT EXISTS backlokto;
CREATE USER IF NOT EXISTS 'backlokto'@'%' IDENTIFIED BY 'qwerty';
GRANT ALL PRIVILEGES ON backlokto.* TO 'backlokto'@'%' WITH GRANT OPTION;
```

Creating test table and inserting data:

```sql
USE backlokto;
CREATE TABLE backlokto_data (
    id INT AUTO_INCREMENT PRIMARY KEY,
    column1 VARCHAR(255),
    column2 INT,
    column3 VARCHAR(255)
);

INSERT INTO backlokto_data (column1, column2, column3) VALUES
('Sample 1', 10, '2024-03-02 12:00:00'),
('Sample 2', 20, '2024-03-02 12:15:00'),
('Sample 3', 30, '2024-03-02 12:30:00'),
('Sample 4', 40, '2024-03-02 12:45:00'),
('Sample 5', 50, '2024-03-02 13:00:00');
```