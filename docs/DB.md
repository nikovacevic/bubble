[Home](https://github.com/nikovacevic/bubble) | [Docs](https://github.com/nikovacevic/bubble/blob/master/docs/)

## Database

### Local installation
https://www.postgresql.org/download/linux/ubuntu/

```
$ sudo systemctl start postgresql@13-main
$ sudo su postgres
$ psql
```
```sql
CREATE USER niko;
CREATE DATABASE bubble WITH OWNER=niko;
```
```
$ psql -d bubble
bubble=>
```


