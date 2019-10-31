
# Instrumented SQL

Experiment purpose:
- How to log every event to database
- List all query

```

2019/10/31 16:37:18 
        create table foo (id integer not null primary key, name text);
        delete from foo;

2019/10/31 16:37:18 sql-tx-begin [err <nil> duration 27.121µs]
2019/10/31 16:37:18 sql-prepare [query insert into foo(id, name) values(?, ?) err <nil> duration 38.606µs]
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 96.785µs args {[int64 0], [string "こんにちわ世界000"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.492µs args {[int64 1], [string "こんにちわ世界001"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.443µs args {[int64 2], [string "こんにちわ世界002"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 12.873µs args {[int64 3], [string "こんにちわ世界003"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.439µs args {[int64 4], [string "こんにちわ世界004"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.942µs args {[int64 5], [string "こんにちわ世界005"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.112µs args {[int64 6], [string "こんにちわ世界006"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.838µs args {[int64 7], [string "こんにちわ世界007"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 6.09µs args {[int64 8], [string "こんにちわ世界008"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.501µs args {[int64 9], [string "こんにちわ世界009"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.981µs args {[int64 10], [string "こんにちわ世界010"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.969µs args {[int64 11], [string "こんにちわ世界011"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.055µs args {[int64 12], [string "こんにちわ世界012"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.204µs args {[int64 13], [string "こんにちわ世界013"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.859µs args {[int64 14], [string "こんにちわ世界014"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.63µs args {[int64 15], [string "こんにちわ世界015"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.986µs args {[int64 16], [string "こんにちわ世界016"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.934µs args {[int64 17], [string "こんにちわ世界017"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.867µs args {[int64 18], [string "こんにちわ世界018"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.913µs args {[int64 19], [string "こんにちわ世界019"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.945µs args {[int64 20], [string "こんにちわ世界020"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.976µs args {[int64 21], [string "こんにちわ世界021"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.239µs args {[int64 22], [string "こんにちわ世界022"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.328µs args {[int64 23], [string "こんにちわ世界023"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.934µs args {[int64 24], [string "こんにちわ世界024"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.307µs args {[int64 25], [string "こんにちわ世界025"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.846µs args {[int64 26], [string "こんにちわ世界026"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.772µs args {[int64 27], [string "こんにちわ世界027"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 9.776µs args {[int64 28], [string "こんにちわ世界028"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.269µs args {[int64 29], [string "こんにちわ世界029"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.996µs args {[int64 30], [string "こんにちわ世界030"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.066µs args {[int64 31], [string "こんにちわ世界031"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.965µs args {[int64 32], [string "こんにちわ世界032"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.974µs args {[int64 33], [string "こんにちわ世界033"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.93µs args {[int64 34], [string "こんにちわ世界034"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.034µs args {[int64 35], [string "こんにちわ世界035"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.94µs args {[int64 36], [string "こんにちわ世界036"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.797µs args {[int64 37], [string "こんにちわ世界037"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.033µs args {[int64 38], [string "こんにちわ世界038"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.939µs args {[int64 39], [string "こんにちわ世界039"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 10.014µs args {[int64 40], [string "こんにちわ世界040"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.155µs args {[int64 41], [string "こんにちわ世界041"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.077µs args {[int64 42], [string "こんにちわ世界042"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.405µs args {[int64 43], [string "こんにちわ世界043"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.895µs args {[int64 44], [string "こんにちわ世界044"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.993µs args {[int64 45], [string "こんにちわ世界045"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.203µs args {[int64 46], [string "こんにちわ世界046"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.98µs args {[int64 47], [string "こんにちわ世界047"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.869µs args {[int64 48], [string "こんにちわ世界048"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.413µs args {[int64 49], [string "こんにちわ世界049"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.466µs args {[int64 50], [string "こんにちわ世界050"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 14.865µs args {[int64 51], [string "こんにちわ世界051"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 6.211µs args {[int64 52], [string "こんにちわ世界052"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.138µs args {[int64 53], [string "こんにちわ世界053"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.599µs args {[int64 54], [string "こんにちわ世界054"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.14µs args {[int64 55], [string "こんにちわ世界055"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.153µs args {[int64 56], [string "こんにちわ世界056"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.041µs args {[int64 57], [string "こんにちわ世界057"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.959µs args {[int64 58], [string "こんにちわ世界058"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.27µs args {[int64 59], [string "こんにちわ世界059"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 16.889µs args {[int64 60], [string "こんにちわ世界060"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.156µs args {[int64 61], [string "こんにちわ世界061"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.238µs args {[int64 62], [string "こんにちわ世界062"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.959µs args {[int64 63], [string "こんにちわ世界063"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.475µs args {[int64 64], [string "こんにちわ世界064"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.113µs args {[int64 65], [string "こんにちわ世界065"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.739µs args {[int64 66], [string "こんにちわ世界066"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.621µs args {[int64 67], [string "こんにちわ世界067"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.051µs args {[int64 68], [string "こんにちわ世界068"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.992µs args {[int64 69], [string "こんにちわ世界069"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 8.701µs args {[int64 70], [string "こんにちわ世界070"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.579µs args {[int64 71], [string "こんにちわ世界071"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.047µs args {[int64 72], [string "こんにちわ世界072"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.318µs args {[int64 73], [string "こんにちわ世界073"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.991µs args {[int64 74], [string "こんにちわ世界074"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.001µs args {[int64 75], [string "こんにちわ世界075"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.899µs args {[int64 76], [string "こんにちわ世界076"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.947µs args {[int64 77], [string "こんにちわ世界077"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.399µs args {[int64 78], [string "こんにちわ世界078"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.016µs args {[int64 79], [string "こんにちわ世界079"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.067µs args {[int64 80], [string "こんにちわ世界080"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.201µs args {[int64 81], [string "こんにちわ世界081"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.981µs args {[int64 82], [string "こんにちわ世界082"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.993µs args {[int64 83], [string "こんにちわ世界083"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.244µs args {[int64 84], [string "こんにちわ世界084"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.008µs args {[int64 85], [string "こんにちわ世界085"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.028µs args {[int64 86], [string "こんにちわ世界086"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.039µs args {[int64 87], [string "こんにちわ世界087"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.19µs args {[int64 88], [string "こんにちわ世界088"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 6.294µs args {[int64 89], [string "こんにちわ世界089"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.171µs args {[int64 90], [string "こんにちわ世界090"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.065µs args {[int64 91], [string "こんにちわ世界091"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.909µs args {[int64 92], [string "こんにちわ世界092"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 23.844µs args {[int64 93], [string "こんにちわ世界093"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 14.541µs args {[int64 94], [string "こんにちわ世界094"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.982µs args {[int64 95], [string "こんにちわ世界095"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.908µs args {[int64 96], [string "こんにちわ世界096"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.999µs args {[int64 97], [string "こんにちわ世界097"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.281µs args {[int64 98], [string "こんにちわ世界098"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.414µs args {[int64 99], [string "こんにちわ世界099"]}]
2019/10/31 16:37:18 insert into foo(id, name) values(?, ?)
2019/10/31 16:37:18 sql-tx-commit [err <nil> duration 432.478µs]
2019/10/31 16:37:18 sql-stmt-close [err <nil> duration 1.688µs]
2019/10/31 16:37:18 sql-conn-query [query select id, name from foo err <nil> duration 35.67µs args {}]
2019/10/31 16:37:18 select id, name from foo
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 45.065µs]
0 こんにちわ世界000
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 2.143µs]
1 こんにちわ世界001
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.701µs]
2 こんにちわ世界002
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 6.539µs]
3 こんにちわ世界003
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.798µs]
4 こんにちわ世界004
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.532µs]
5 こんにちわ世界005
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.546µs]
6 こんにちわ世界006
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.729µs]
7 こんにちわ世界007
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 18.83µs]
8 こんにちわ世界008
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.898µs]
9 こんにちわ世界009
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.73µs]
10 こんにちわ世界010
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 11.415µs]
11 こんにちわ世界011
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 2.691µs]
12 こんにちわ世界012
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.452µs]
13 こんにちわ世界013
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.341µs]
14 こんにちわ世界014
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.356µs]
15 こんにちわ世界015
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 2.119µs]
16 こんにちわ世界016
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.448µs]
17 こんにちわ世界017
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.427µs]
18 こんにちわ世界018
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.521µs]
19 こんにちわ世界019
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.422µs]
20 こんにちわ世界020
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.468µs]
21 こんにちわ世界021
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.53µs]
22 こんにちわ世界022
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.483µs]
23 こんにちわ世界023
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.681µs]
24 こんにちわ世界024
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.382µs]
25 こんにちわ世界025
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.439µs]
26 こんにちわ世界026
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.447µs]
27 こんにちわ世界027
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.476µs]
28 こんにちわ世界028
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.386µs]
29 こんにちわ世界029
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.431µs]
30 こんにちわ世界030
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.468µs]
31 こんにちわ世界031
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.439µs]
32 こんにちわ世界032
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.443µs]
33 こんにちわ世界033
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.478µs]
34 こんにちわ世界034
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.457µs]
35 こんにちわ世界035
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.568µs]
36 こんにちわ世界036
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.484µs]
37 こんにちわ世界037
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.517µs]
38 こんにちわ世界038
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.401µs]
39 こんにちわ世界039
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.481µs]
40 こんにちわ世界040
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.516µs]
41 こんにちわ世界041
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.476µs]
42 こんにちわ世界042
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.438µs]
43 こんにちわ世界043
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.521µs]
44 こんにちわ世界044
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.388µs]
45 こんにちわ世界045
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.652µs]
46 こんにちわ世界046
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.499µs]
47 こんにちわ世界047
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.427µs]
48 こんにちわ世界048
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.419µs]
49 こんにちわ世界049
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.464µs]
50 こんにちわ世界050
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.471µs]
51 こんにちわ世界051
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.528µs]
52 こんにちわ世界052
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.426µs]
53 こんにちわ世界053
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.487µs]
54 こんにちわ世界054
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.551µs]
55 こんにちわ世界055
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.467µs]
56 こんにちわ世界056
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.457µs]
57 こんにちわ世界057
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.505µs]
58 こんにちわ世界058
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.429µs]
59 こんにちわ世界059
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.447µs]
60 こんにちわ世界060
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.524µs]
61 こんにちわ世界061
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.46µs]
62 こんにちわ世界062
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.45µs]
63 こんにちわ世界063
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.509µs]
64 こんにちわ世界064
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.442µs]
65 こんにちわ世界065
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.766µs]
66 こんにちわ世界066
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.52µs]
67 こんにちわ世界067
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.436µs]
68 こんにちわ世界068
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.467µs]
69 こんにちわ世界069
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.496µs]
70 こんにちわ世界070
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.556µs]
71 こんにちわ世界071
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.542µs]
72 こんにちわ世界072
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.506µs]
73 こんにちわ世界073
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.466µs]
74 こんにちわ世界074
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.532µs]
75 こんにちわ世界075
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.486µs]
76 こんにちわ世界076
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.482µs]
77 こんにちわ世界077
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.582µs]
78 こんにちわ世界078
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.626µs]
79 こんにちわ世界079
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.456µs]
80 こんにちわ世界080
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.336µs]
81 こんにちわ世界081
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.395µs]
82 こんにちわ世界082
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.313µs]
83 こんにちわ世界083
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.611µs]
84 こんにちわ世界084
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.377µs]
85 こんにちわ世界085
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.356µs]
86 こんにちわ世界086
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.343µs]
87 こんにちわ世界087
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.403µs]
88 こんにちわ世界088
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.329µs]
89 こんにちわ世界089
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.359µs]
90 こんにちわ世界090
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.355µs]
91 こんにちわ世界091
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.33µs]
92 こんにちわ世界092
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.388µs]
93 こんにちわ世界093
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.431µs]
94 こんにちわ世界094
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.812µs]
95 こんにちわ世界095
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.452µs]
96 こんにちわ世界096
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.419µs]
97 こんにちわ世界097
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.392µs]
98 こんにちわ世界098
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.484µs]
99 こんにちわ世界099
2019/10/31 16:37:18 sql-rows-next [err EOF duration 7.452µs]
2019/10/31 16:37:18 sql-prepare [query select name from foo where id = ? err <nil> duration 24.584µs]
2019/10/31 16:37:18 sql-stmt-query [query select name from foo where id = ? err <nil> duration 3.105µs args {[string "3"]}]
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 34.019µs]
こんにちわ世界003
2019/10/31 16:37:18 sql-conn-exec [query delete from foo err <nil> duration 481.165µs args {}]
2019/10/31 16:37:18 delete from foo
2019/10/31 16:37:18 sql-conn-exec [query insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz') err <nil> duration 595.152µs args {}]
2019/10/31 16:37:18 insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')
2019/10/31 16:37:18 sql-conn-query [query select id, name from foo err <nil> duration 60.72µs args {}]
2019/10/31 16:37:18 select id, name from foo
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 29.557µs]
1 foo
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 2.07µs]
2 bar
2019/10/31 16:37:18 sql-rows-next [err <nil> duration 1.568µs]
3 baz
2019/10/31 16:37:18 sql-rows-next [err EOF duration 18.227µs]
list Query
"insert into foo(id, name) values(?, ?)" 
"select id, name from foo" 
"delete from foo" 
"insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')" 
"\n\tcreate table foo (id integer not null primary key, name text);\n\tdelete from foo;\n\t" 
2019/10/31 16:37:18 sql-stmt-close [err <nil> duration 1.595µs]
```