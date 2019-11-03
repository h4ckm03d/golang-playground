
# Instrumented SQL

Experiment purpose:
- How to log every event to database
- List all query

```
5-instrumented-sql git:(master) ✗ go run main.go          
2019/11/03 17:11:31 sql-connector-connect [err <nil> duration 711.712µs]
2019/11/03 17:11:31 sql-conn-exec [query 
        create table foo (id integer not null primary key, name text);
        delete from foo;
         err <nil> duration 1.320493ms args {}]
2019/11/03 17:11:31 
        create table foo (id integer not null primary key, name text);
        delete from foo;

2019/11/03 17:11:31 sql-tx-begin [err <nil> duration 8.646µs]
2019/11/03 17:11:31 sql-prepare [query insert into foo(id, name) values(?, ?) err <nil> duration 21.473µs]
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 148.111µs args {[int64 0], [string "こんにちわ世界000"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.302µs args {[int64 1], [string "こんにちわ世界001"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.207µs args {[int64 2], [string "こんにちわ世界002"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.963µs args {[int64 3], [string "こんにちわ世界003"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.789µs args {[int64 4], [string "こんにちわ世界004"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 11.208µs args {[int64 5], [string "こんにちわ世界005"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.605µs args {[int64 6], [string "こんにちわ世界006"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.169µs args {[int64 7], [string "こんにちわ世界007"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.485µs args {[int64 8], [string "こんにちわ世界008"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 7.43µs args {[int64 9], [string "こんにちわ世界009"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.194µs args {[int64 10], [string "こんにちわ世界010"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.918µs args {[int64 11], [string "こんにちわ世界011"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.112µs args {[int64 12], [string "こんにちわ世界012"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.651µs args {[int64 13], [string "こんにちわ世界013"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.12µs args {[int64 14], [string "こんにちわ世界014"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.099µs args {[int64 15], [string "こんにちわ世界015"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.721µs args {[int64 16], [string "こんにちわ世界016"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.093µs args {[int64 17], [string "こんにちわ世界017"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.547µs args {[int64 18], [string "こんにちわ世界018"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.965µs args {[int64 19], [string "こんにちわ世界019"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.887µs args {[int64 20], [string "こんにちわ世界020"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.161µs args {[int64 21], [string "こんにちわ世界021"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.502µs args {[int64 22], [string "こんにちわ世界022"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.358µs args {[int64 23], [string "こんにちわ世界023"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.188µs args {[int64 24], [string "こんにちわ世界024"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.388µs args {[int64 25], [string "こんにちわ世界025"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.186µs args {[int64 26], [string "こんにちわ世界026"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.253µs args {[int64 27], [string "こんにちわ世界027"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.095µs args {[int64 28], [string "こんにちわ世界028"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.426µs args {[int64 29], [string "こんにちわ世界029"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.81µs args {[int64 30], [string "こんにちわ世界030"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.788µs args {[int64 31], [string "こんにちわ世界031"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.362µs args {[int64 32], [string "こんにちわ世界032"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.936µs args {[int64 33], [string "こんにちわ世界033"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2µs args {[int64 34], [string "こんにちわ世界034"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.563µs args {[int64 35], [string "こんにちわ世界035"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.045µs args {[int64 36], [string "こんにちわ世界036"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.866µs args {[int64 37], [string "こんにちわ世界037"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.005µs args {[int64 38], [string "こんにちわ世界038"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.929µs args {[int64 39], [string "こんにちわ世界039"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.326µs args {[int64 40], [string "こんにちわ世界040"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.475µs args {[int64 41], [string "こんにちわ世界041"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.335µs args {[int64 42], [string "こんにちわ世界042"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.294µs args {[int64 43], [string "こんにちわ世界043"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 9.368µs args {[int64 44], [string "こんにちわ世界044"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.007µs args {[int64 45], [string "こんにちわ世界045"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.034µs args {[int64 46], [string "こんにちわ世界046"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.831µs args {[int64 47], [string "こんにちわ世界047"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.823µs args {[int64 48], [string "こんにちわ世界048"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.804µs args {[int64 49], [string "こんにちわ世界049"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.147µs args {[int64 50], [string "こんにちわ世界050"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.397µs args {[int64 51], [string "こんにちわ世界051"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.901µs args {[int64 52], [string "こんにちわ世界052"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.176µs args {[int64 53], [string "こんにちわ世界053"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.819µs args {[int64 54], [string "こんにちわ世界054"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.072µs args {[int64 55], [string "こんにちわ世界055"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.888µs args {[int64 56], [string "こんにちわ世界056"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.198µs args {[int64 57], [string "こんにちわ世界057"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.913µs args {[int64 58], [string "こんにちわ世界058"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.437µs args {[int64 59], [string "こんにちわ世界059"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.246µs args {[int64 60], [string "こんにちわ世界060"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.862µs args {[int64 61], [string "こんにちわ世界061"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.549µs args {[int64 62], [string "こんにちわ世界062"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.671µs args {[int64 63], [string "こんにちわ世界063"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 10.282µs args {[int64 64], [string "こんにちわ世界064"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.493µs args {[int64 65], [string "こんにちわ世界065"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 15.119µs args {[int64 66], [string "こんにちわ世界066"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.762µs args {[int64 67], [string "こんにちわ世界067"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.602µs args {[int64 68], [string "こんにちわ世界068"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.947µs args {[int64 69], [string "こんにちわ世界069"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.132µs args {[int64 70], [string "こんにちわ世界070"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:31 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.232µs args {[int64 71], [string "こんにちわ世界071"]}]
2019/11/03 17:11:31 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.628µs args {[int64 72], [string "こんにちわ世界072"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.072µs args {[int64 73], [string "こんにちわ世界073"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.102µs args {[int64 74], [string "こんにちわ世界074"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.006µs args {[int64 75], [string "こんにちわ世界075"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.308µs args {[int64 76], [string "こんにちわ世界076"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.682µs args {[int64 77], [string "こんにちわ世界077"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.585µs args {[int64 78], [string "こんにちわ世界078"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.252µs args {[int64 79], [string "こんにちわ世界079"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.198µs args {[int64 80], [string "こんにちわ世界080"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.171µs args {[int64 81], [string "こんにちわ世界081"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.006µs args {[int64 82], [string "こんにちわ世界082"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.721µs args {[int64 83], [string "こんにちわ世界083"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.544µs args {[int64 84], [string "こんにちわ世界084"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.315µs args {[int64 85], [string "こんにちわ世界085"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.392µs args {[int64 86], [string "こんにちわ世界086"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.793µs args {[int64 87], [string "こんにちわ世界087"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.841µs args {[int64 88], [string "こんにちわ世界088"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.082µs args {[int64 89], [string "こんにちわ世界089"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.05µs args {[int64 90], [string "こんにちわ世界090"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.126µs args {[int64 91], [string "こんにちわ世界091"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.093µs args {[int64 92], [string "こんにちわ世界092"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.965µs args {[int64 93], [string "こんにちわ世界093"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.983µs args {[int64 94], [string "こんにちわ世界094"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.834µs args {[int64 95], [string "こんにちわ世界095"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.146µs args {[int64 96], [string "こんにちわ世界096"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.101µs args {[int64 97], [string "こんにちわ世界097"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.037µs args {[int64 98], [string "こんにちわ世界098"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.973µs args {[int64 99], [string "こんにちわ世界099"]}]
2019/11/03 17:11:32 insert into foo(id, name) values(?, ?)
2019/11/03 17:11:32 sql-tx-commit [err <nil> duration 510.326µs]
2019/11/03 17:11:32 sql-stmt-close [err <nil> duration 1.375µs]
2019/11/03 17:11:32 sql-conn-query [query select id, name from foo err <nil> duration 30.756µs args {}]
2019/11/03 17:11:32 select id, name from foo
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 41.654µs]
0 こんにちわ世界000
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 2.455µs]
1 こんにちわ世界001
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.53µs]
2 こんにちわ世界002
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.494µs]
3 こんにちわ世界003
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.289µs]
4 こんにちわ世界004
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.295µs]
5 こんにちわ世界005
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.283µs]
6 こんにちわ世界006
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.316µs]
7 こんにちわ世界007
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 3.42µs]
8 こんにちわ世界008
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.472µs]
9 こんにちわ世界009
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.271µs]
10 こんにちわ世界010
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.201µs]
11 こんにちわ世界011
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.22µs]
12 こんにちわ世界012
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.2µs]
13 こんにちわ世界013
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.224µs]
14 こんにちわ世界014
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.288µs]
15 こんにちわ世界015
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.379µs]
16 こんにちわ世界016
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.262µs]
17 こんにちわ世界017
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.341µs]
18 こんにちわ世界018
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.31µs]
19 こんにちわ世界019
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.312µs]
20 こんにちわ世界020
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.228µs]
21 こんにちわ世界021
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.422µs]
22 こんにちわ世界022
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.247µs]
23 こんにちわ世界023
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.255µs]
24 こんにちわ世界024
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 2.177µs]
25 こんにちわ世界025
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.423µs]
26 こんにちわ世界026
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.251µs]
27 こんにちわ世界027
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.281µs]
28 こんにちわ世界028
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.243µs]
29 こんにちわ世界029
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.282µs]
30 こんにちわ世界030
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.57µs]
31 こんにちわ世界031
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 2.11µs]
32 こんにちわ世界032
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 2.374µs]
33 こんにちわ世界033
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.536µs]
34 こんにちわ世界034
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.232µs]
35 こんにちわ世界035
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.365µs]
36 こんにちわ世界036
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 2.151µs]
37 こんにちわ世界037
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.402µs]
38 こんにちわ世界038
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.998µs]
39 こんにちわ世界039
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.365µs]
40 こんにちわ世界040
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.338µs]
41 こんにちわ世界041
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 11.358µs]
42 こんにちわ世界042
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 4.058µs]
43 こんにちわ世界043
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 5.94µs]
44 こんにちわ世界044
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.439µs]
45 こんにちわ世界045
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.248µs]
46 こんにちわ世界046
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.297µs]
47 こんにちわ世界047
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.321µs]
48 こんにちわ世界048
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.241µs]
49 こんにちわ世界049
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.238µs]
50 こんにちわ世界050
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.307µs]
51 こんにちわ世界051
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.301µs]
52 こんにちわ世界052
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.214µs]
53 こんにちわ世界053
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.311µs]
54 こんにちわ世界054
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.308µs]
55 こんにちわ世界055
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.42µs]
56 こんにちわ世界056
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.267µs]
57 こんにちわ世界057
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.295µs]
58 こんにちわ世界058
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.296µs]
59 こんにちわ世界059
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.285µs]
60 こんにちわ世界060
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.269µs]
61 こんにちわ世界061
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.226µs]
62 こんにちわ世界062
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.321µs]
63 こんにちわ世界063
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.323µs]
64 こんにちわ世界064
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.189µs]
65 こんにちわ世界065
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.33µs]
66 こんにちわ世界066
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.279µs]
67 こんにちわ世界067
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.246µs]
68 こんにちわ世界068
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.303µs]
69 こんにちわ世界069
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 5.099µs]
70 こんにちわ世界070
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 2.003µs]
71 こんにちわ世界071
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.35µs]
72 こんにちわ世界072
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.227µs]
73 こんにちわ世界073
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.245µs]
74 こんにちわ世界074
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.3µs]
75 こんにちわ世界075
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.236µs]
76 こんにちわ世界076
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.287µs]
77 こんにちわ世界077
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.265µs]
78 こんにちわ世界078
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.337µs]
79 こんにちわ世界079
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.238µs]
80 こんにちわ世界080
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.22µs]
81 こんにちわ世界081
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.251µs]
82 こんにちわ世界082
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.212µs]
83 こんにちわ世界083
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.242µs]
84 こんにちわ世界084
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 2.112µs]
85 こんにちわ世界085
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.418µs]
86 こんにちわ世界086
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.175µs]
87 こんにちわ世界087
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.193µs]
88 こんにちわ世界088
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.171µs]
89 こんにちわ世界089
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.154µs]
90 こんにちわ世界090
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.237µs]
91 こんにちわ世界091
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.23µs]
92 こんにちわ世界092
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.48µs]
93 こんにちわ世界093
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.196µs]
94 こんにちわ世界094
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.307µs]
95 こんにちわ世界095
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.182µs]
96 こんにちわ世界096
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.284µs]
97 こんにちわ世界097
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.367µs]
98 こんにちわ世界098
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.257µs]
99 こんにちわ世界099
2019/11/03 17:11:32 sql-rows-next [err EOF duration 5.626µs]
2019/11/03 17:11:32 sql-conn-query [query select id, name from foo where name=? err <nil> duration 36.456µs args {[string "bajindul"]}]
2019/11/03 17:11:32 select id, name from foo where name=?
2019/11/03 17:11:32 sql-rows-next [err EOF duration 39.559µs]
2019/11/03 17:11:32 sql-prepare [query select name from foo where id = ? err <nil> duration 10.833µs]
2019/11/03 17:11:32 sql-stmt-query [query select name from foo where id = ? err <nil> duration 2.519µs args {[string "3"]}]
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 23.48µs]
こんにちわ世界003
2019/11/03 17:11:32 sql-conn-exec [query delete from foo err <nil> duration 582.841µs args {}]
2019/11/03 17:11:32 delete from foo
2019/11/03 17:11:32 sql-conn-exec [query insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz') err <nil> duration 496.685µs args {}]
2019/11/03 17:11:32 insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')
2019/11/03 17:11:32 sql-conn-query [query select id, name from foo err <nil> duration 14.957µs args {}]
2019/11/03 17:11:32 select id, name from foo
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 31.471µs]
1 foo
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 2.128µs]
2 bar
2019/11/03 17:11:32 sql-rows-next [err <nil> duration 1.335µs]
3 baz
2019/11/03 17:11:32 sql-rows-next [err EOF duration 3.197µs]
2019/11/03 17:11:32 sql-stmt-close [err <nil> duration 857ns]

===list known Query===
"select id, name from foo" 
"select id, name from foo where name=?"
```