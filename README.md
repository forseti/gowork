# gowork
Go Workout is a project I created to re-learn and sharpen 
my knowledge in `go` programming language.

[The course](https://www.udemy.com/course/pemrograman-go-lang-pemula-sampai-mahir/) 
I'm learning from is made by 
[an awesome programmer, Eko Kurniawan](https://programmerzamannow.com/). 
I highly recommend it to anyone who can understand Bahasa Indonesia.

### Docker Compose
These are the commands you can use with `docker-compose`.
#### MySQL
Exercise 05 uses MySQL. To run the `db` service:
```shell
docker-compose run --service-ports entity
```
To stop the database
```shell
docker-compose down
```
To remove the `db`:
```shell
docker volume rm gowork_my-entity
```
### Development
#### MySQL
Once you successfully run a MySQL container, 
you can use `mysql` console to access the database (the password is `root`):
```shell
mysql -h127.0.0.1 -P3306 -u root -p
```
Then in `mysql` console, you can access the db schema, 
the default name is `gowork_db`:
```sql
use gowork_db;
```
You can instantiate any SQL scripts used by this project:
```sql
source @db/customer.sql
```
