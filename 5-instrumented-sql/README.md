
# Instrumented SQL

Experiment purpose:
- How to log every event to database
- List all query

```
14:52 5-instrumented-sql (master) •|• go run main.go
2019/11/01 14:52:24 sql-connector-connect [err <nil> duration 882.897µs]
2019/11/01 14:52:24 sql-conn-exec [query 
        create table foo (id integer not null primary key, name text);
        delete from foo;
         err <nil> duration 1.73524ms args {}]
2019/11/01 14:52:24 
        create table foo (id integer not null primary key, name text);
        delete from foo;

2019/11/01 14:52:24 sql-tx-begin [err <nil> duration 33.307µs]
2019/11/01 14:52:24 sql-prepare [query insert into foo(id, name) values(?, ?) err <nil> duration 38.97µs]
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 97.795µs args {[int64 0], [string "こんにちわ世界000"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.33µs args {[int64 1], [string "こんにちわ世界001"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.646µs args {[int64 2], [string "こんにちわ世界002"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 13.519µs args {[int64 3], [string "こんにちわ世界003"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.848µs args {[int64 4], [string "こんにちわ世界004"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.282µs args {[int64 5], [string "こんにちわ世界005"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.062µs args {[int64 6], [string "こんにちわ世界006"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.112µs args {[int64 7], [string "こんにちわ世界007"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 7.707µs args {[int64 8], [string "こんにちわ世界008"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.47µs args {[int64 9], [string "こんにちわ世界009"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.994µs args {[int64 10], [string "こんにちわ世界010"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.173µs args {[int64 11], [string "こんにちわ世界011"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 6.325µs args {[int64 12], [string "こんにちわ世界012"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.18µs args {[int64 13], [string "こんにちわ世界013"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.376µs args {[int64 14], [string "こんにちわ世界014"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.069µs args {[int64 15], [string "こんにちわ世界015"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.073µs args {[int64 16], [string "こんにちわ世界016"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 17.368µs args {[int64 17], [string "こんにちわ世界017"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.086µs args {[int64 18], [string "こんにちわ世界018"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.433µs args {[int64 19], [string "こんにちわ世界019"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.145µs args {[int64 20], [string "こんにちわ世界020"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.973µs args {[int64 21], [string "こんにちわ世界021"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.997µs args {[int64 22], [string "こんにちわ世界022"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.383µs args {[int64 23], [string "こんにちわ世界023"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.962µs args {[int64 24], [string "こんにちわ世界024"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.172µs args {[int64 25], [string "こんにちわ世界025"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.855µs args {[int64 26], [string "こんにちわ世界026"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.255µs args {[int64 27], [string "こんにちわ世界027"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.852µs args {[int64 28], [string "こんにちわ世界028"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.396µs args {[int64 29], [string "こんにちわ世界029"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.212µs args {[int64 30], [string "こんにちわ世界030"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.19µs args {[int64 31], [string "こんにちわ世界031"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.375µs args {[int64 32], [string "こんにちわ世界032"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.163µs args {[int64 33], [string "こんにちわ世界033"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.125µs args {[int64 34], [string "こんにちわ世界034"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.127µs args {[int64 35], [string "こんにちわ世界035"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.549µs args {[int64 36], [string "こんにちわ世界036"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.073µs args {[int64 37], [string "こんにちわ世界037"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 31.109µs args {[int64 38], [string "こんにちわ世界038"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.323µs args {[int64 39], [string "こんにちわ世界039"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.139µs args {[int64 40], [string "こんにちわ世界040"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.278µs args {[int64 41], [string "こんにちわ世界041"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.149µs args {[int64 42], [string "こんにちわ世界042"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.664µs args {[int64 43], [string "こんにちわ世界043"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.755µs args {[int64 44], [string "こんにちわ世界044"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.476µs args {[int64 45], [string "こんにちわ世界045"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.719µs args {[int64 46], [string "こんにちわ世界046"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 14.855µs args {[int64 47], [string "こんにちわ世界047"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.358µs args {[int64 48], [string "こんにちわ世界048"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 14.673µs args {[int64 49], [string "こんにちわ世界049"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.262µs args {[int64 50], [string "こんにちわ世界050"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.202µs args {[int64 51], [string "こんにちわ世界051"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.183µs args {[int64 52], [string "こんにちわ世界052"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 14.557µs args {[int64 53], [string "こんにちわ世界053"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.349µs args {[int64 54], [string "こんにちわ世界054"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.126µs args {[int64 55], [string "こんにちわ世界055"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.179µs args {[int64 56], [string "こんにちわ世界056"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.229µs args {[int64 57], [string "こんにちわ世界057"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.319µs args {[int64 58], [string "こんにちわ世界058"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.465µs args {[int64 59], [string "こんにちわ世界059"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.299µs args {[int64 60], [string "こんにちわ世界060"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.148µs args {[int64 61], [string "こんにちわ世界061"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.281µs args {[int64 62], [string "こんにちわ世界062"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.124µs args {[int64 63], [string "こんにちわ世界063"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.527µs args {[int64 64], [string "こんにちわ世界064"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.381µs args {[int64 65], [string "こんにちわ世界065"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.208µs args {[int64 66], [string "こんにちわ世界066"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.626µs args {[int64 67], [string "こんにちわ世界067"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.246µs args {[int64 68], [string "こんにちわ世界068"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.226µs args {[int64 69], [string "こんにちわ世界069"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 37.478µs args {[int64 70], [string "こんにちわ世界070"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 15.667µs args {[int64 71], [string "こんにちわ世界071"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.402µs args {[int64 72], [string "こんにちわ世界072"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.644µs args {[int64 73], [string "こんにちわ世界073"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.314µs args {[int64 74], [string "こんにちわ世界074"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.522µs args {[int64 75], [string "こんにちわ世界075"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.158µs args {[int64 76], [string "こんにちわ世界076"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.179µs args {[int64 77], [string "こんにちわ世界077"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.246µs args {[int64 78], [string "こんにちわ世界078"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.261µs args {[int64 79], [string "こんにちわ世界079"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.143µs args {[int64 80], [string "こんにちわ世界080"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.401µs args {[int64 81], [string "こんにちわ世界081"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.489µs args {[int64 82], [string "こんにちわ世界082"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.363µs args {[int64 83], [string "こんにちわ世界083"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 14.417µs args {[int64 84], [string "こんにちわ世界084"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.102µs args {[int64 85], [string "こんにちわ世界085"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.122µs args {[int64 86], [string "こんにちわ世界086"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.207µs args {[int64 87], [string "こんにちわ世界087"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.096µs args {[int64 88], [string "こんにちわ世界088"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.338µs args {[int64 89], [string "こんにちわ世界089"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.17µs args {[int64 90], [string "こんにちわ世界090"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.288µs args {[int64 91], [string "こんにちわ世界091"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.353µs args {[int64 92], [string "こんにちわ世界092"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.346µs args {[int64 93], [string "こんにちわ世界093"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.445µs args {[int64 94], [string "こんにちわ世界094"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 14.597µs args {[int64 95], [string "こんにちわ世界095"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.386µs args {[int64 96], [string "こんにちわ世界096"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.057µs args {[int64 97], [string "こんにちわ世界097"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.092µs args {[int64 98], [string "こんにちわ世界098"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.165µs args {[int64 99], [string "こんにちわ世界099"]}]
2019/11/01 14:52:24 insert into foo(id, name) values(?, ?)
2019/11/01 14:52:24 sql-tx-commit [err <nil> duration 393.15µs]
2019/11/01 14:52:24 sql-stmt-close [err <nil> duration 1.651µs]
2019/11/01 14:52:24 sql-conn-query [query select id, name from foo err <nil> duration 47.204µs args {}]
2019/11/01 14:52:24 select id, name from foo
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 44.949µs]
0 こんにちわ世界000
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 2.241µs]
1 こんにちわ世界001
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.52µs]
2 こんにちわ世界002
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.47µs]
3 こんにちわ世界003
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.933µs]
4 こんにちわ世界004
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.582µs]
5 こんにちわ世界005
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.474µs]
6 こんにちわ世界006
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.462µs]
7 こんにちわ世界007
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.453µs]
8 こんにちわ世界008
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.456µs]
9 こんにちわ世界009
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.427µs]
10 こんにちわ世界010
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.399µs]
11 こんにちわ世界011
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.491µs]
12 こんにちわ世界012
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.475µs]
13 こんにちわ世界013
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.499µs]
14 こんにちわ世界014
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.399µs]
15 こんにちわ世界015
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.635µs]
16 こんにちわ世界016
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.443µs]
17 こんにちわ世界017
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.499µs]
18 こんにちわ世界018
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.473µs]
19 こんにちわ世界019
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.442µs]
20 こんにちわ世界020
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.487µs]
21 こんにちわ世界021
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.509µs]
22 こんにちわ世界022
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.469µs]
23 こんにちわ世界023
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.442µs]
24 こんにちわ世界024
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.543µs]
25 こんにちわ世界025
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 15.623µs]
26 こんにちわ世界026
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.462µs]
27 こんにちわ世界027
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.444µs]
28 こんにちわ世界028
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.387µs]
29 こんにちわ世界029
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.498µs]
30 こんにちわ世界030
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.48µs]
31 こんにちわ世界031
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.522µs]
32 こんにちわ世界032
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.437µs]
33 こんにちわ世界033
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.482µs]
34 こんにちわ世界034
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.479µs]
35 こんにちわ世界035
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.497µs]
36 こんにちわ世界036
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.406µs]
37 こんにちわ世界037
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.428µs]
38 こんにちわ世界038
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.462µs]
39 こんにちわ世界039
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.477µs]
40 こんにちわ世界040
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.432µs]
41 こんにちわ世界041
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 15.929µs]
42 こんにちわ世界042
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 4.361µs]
43 こんにちわ世界043
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.407µs]
44 こんにちわ世界044
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.444µs]
45 こんにちわ世界045
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.418µs]
46 こんにちわ世界046
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 4.092µs]
47 こんにちわ世界047
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.697µs]
48 こんにちわ世界048
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.538µs]
49 こんにちわ世界049
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.502µs]
50 こんにちわ世界050
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.473µs]
51 こんにちわ世界051
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.501µs]
52 こんにちわ世界052
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.402µs]
53 こんにちわ世界053
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.442µs]
54 こんにちわ世界054
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.4µs]
55 こんにちわ世界055
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.466µs]
56 こんにちわ世界056
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.819µs]
57 こんにちわ世界057
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 3.932µs]
58 こんにちわ世界058
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.755µs]
59 こんにちわ世界059
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.532µs]
60 こんにちわ世界060
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.528µs]
61 こんにちわ世界061
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.466µs]
62 こんにちわ世界062
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.369µs]
63 こんにちわ世界063
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.378µs]
64 こんにちわ世界064
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.385µs]
65 こんにちわ世界065
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.419µs]
66 こんにちわ世界066
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.515µs]
67 こんにちわ世界067
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.477µs]
68 こんにちわ世界068
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.339µs]
69 こんにちわ世界069
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.311µs]
70 こんにちわ世界070
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.358µs]
71 こんにちわ世界071
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.31µs]
72 こんにちわ世界072
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.389µs]
73 こんにちわ世界073
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.319µs]
74 こんにちわ世界074
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 4.289µs]
75 こんにちわ世界075
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.875µs]
76 こんにちわ世界076
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.459µs]
77 こんにちわ世界077
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.424µs]
78 こんにちわ世界078
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.503µs]
79 こんにちわ世界079
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.515µs]
80 こんにちわ世界080
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.437µs]
81 こんにちわ世界081
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.412µs]
82 こんにちわ世界082
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.411µs]
83 こんにちわ世界083
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.455µs]
84 こんにちわ世界084
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.429µs]
85 こんにちわ世界085
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.398µs]
86 こんにちわ世界086
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.476µs]
87 こんにちわ世界087
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.478µs]
88 こんにちわ世界088
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.405µs]
89 こんにちわ世界089
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.504µs]
90 こんにちわ世界090
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.48µs]
91 こんにちわ世界091
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.402µs]
92 こんにちわ世界092
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.414µs]
93 こんにちわ世界093
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.479µs]
94 こんにちわ世界094
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.495µs]
95 こんにちわ世界095
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.462µs]
96 こんにちわ世界096
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.542µs]
97 こんにちわ世界097
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.606µs]
98 こんにちわ世界098
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.477µs]
99 こんにちわ世界099
2019/11/01 14:52:24 sql-rows-next [err EOF duration 10.801µs]
2019/11/01 14:52:24 sql-prepare [query select name from foo where id = ? err <nil> duration 34.575µs]
2019/11/01 14:52:24 sql-stmt-query [query select name from foo where id = ? err <nil> duration 3.07µs args {[string "3"]}]
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 71.021µs]
こんにちわ世界003
2019/11/01 14:52:24 sql-conn-exec [query delete from foo err <nil> duration 563.872µs args {}]
2019/11/01 14:52:24 delete from foo
2019/11/01 14:52:24 sql-conn-exec [query insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz') err <nil> duration 407.986µs args {}]
2019/11/01 14:52:24 insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')
2019/11/01 14:52:24 sql-conn-query [query select id, name from foo err <nil> duration 19.188µs args {}]
2019/11/01 14:52:24 select id, name from foo
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 27.87µs]
1 foo
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 2.337µs]
2 bar
2019/11/01 14:52:24 sql-rows-next [err <nil> duration 1.529µs]
3 baz
2019/11/01 14:52:24 sql-rows-next [err EOF duration 4.202µs]
2019/11/01 14:52:24 sql-stmt-close [err <nil> duration 1.167µs]

===list Query===
"\n\tcreate table foo (id integer not null primary key, name text);\n\tdelete from foo;\n\t" 
"insert into foo(id, name) values(?, ?)" 
"select id, name from foo" 
"delete from foo" 
"insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')" 
```