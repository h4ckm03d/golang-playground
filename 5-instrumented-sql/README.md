
# Instrumented SQL

Experiment purpose:
- How to log every event to database
- List all query

Main logic to get query:
```go
logger := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
		switch msg {
		case instrumentedsql.OpSQLConnExec,
			instrumentedsql.OpSQLConnQuery,
			instrumentedsql.OpSQLStmtExec:
			if len(keyvals) > 1 && keyvals[0] != "query" {
				return
			}
			query := keyvals[1].(string)
			log.Println(query)
			if ok := queries[query]; !ok && strings.Index(query, "where") > -1 {
				queries[query] = true
			}
		}
	})
```

```
6:44 5-instrumented-sql (master) •/• go run main.go
2020/01/03 16:45:18 sql-connector-connect [err <nil> duration 1.089484ms]
2020/01/03 16:45:18 sql-conn-exec [query 
        create table foo (id integer not null primary key, name text);
        delete from foo;
         err <nil> duration 2.504259ms args {}]
2020/01/03 16:45:18 
        create table foo (id integer not null primary key, name text);
        delete from foo;

2020/01/03 16:45:18 sql-tx-begin [err <nil> duration 14.749µs]
2020/01/03 16:45:18 sql-prepare [query insert into foo(id, name) values(?, ?) err <nil> duration 40.164µs]
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 191.031µs args {[int64 0], [string "こんにちわ世界000"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.868µs args {[int64 1], [string "こんにちわ世界001"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.412µs args {[int64 2], [string "こんにちわ世界002"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 11.435µs args {[int64 3], [string "こんにちわ世界003"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.22µs args {[int64 4], [string "こんにちわ世界004"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.934µs args {[int64 5], [string "こんにちわ世界005"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.86µs args {[int64 6], [string "こんにちわ世界006"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.854µs args {[int64 7], [string "こんにちわ世界007"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.902µs args {[int64 8], [string "こんにちわ世界008"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.971µs args {[int64 9], [string "こんにちわ世界009"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.855µs args {[int64 10], [string "こんにちわ世界010"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.896µs args {[int64 11], [string "こんにちわ世界011"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 6.276µs args {[int64 12], [string "こんにちわ世界012"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 8.159µs args {[int64 13], [string "こんにちわ世界013"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.107µs args {[int64 14], [string "こんにちわ世界014"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.796µs args {[int64 15], [string "こんにちわ世界015"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.305µs args {[int64 16], [string "こんにちわ世界016"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.309µs args {[int64 17], [string "こんにちわ世界017"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.091µs args {[int64 18], [string "こんにちわ世界018"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.239µs args {[int64 19], [string "こんにちわ世界019"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.807µs args {[int64 20], [string "こんにちわ世界020"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.101µs args {[int64 21], [string "こんにちわ世界021"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.947µs args {[int64 22], [string "こんにちわ世界022"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.064µs args {[int64 23], [string "こんにちわ世界023"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.456µs args {[int64 24], [string "こんにちわ世界024"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.958µs args {[int64 25], [string "こんにちわ世界025"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.997µs args {[int64 26], [string "こんにちわ世界026"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.242µs args {[int64 27], [string "こんにちわ世界027"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.746µs args {[int64 28], [string "こんにちわ世界028"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.001µs args {[int64 29], [string "こんにちわ世界029"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.335µs args {[int64 30], [string "こんにちわ世界030"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.911µs args {[int64 31], [string "こんにちわ世界031"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 6.505µs args {[int64 32], [string "こんにちわ世界032"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.595µs args {[int64 33], [string "こんにちわ世界033"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.174µs args {[int64 34], [string "こんにちわ世界034"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.9µs args {[int64 35], [string "こんにちわ世界035"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.631µs args {[int64 36], [string "こんにちわ世界036"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 14.192µs args {[int64 37], [string "こんにちわ世界037"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.061µs args {[int64 38], [string "こんにちわ世界038"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.208µs args {[int64 39], [string "こんにちわ世界039"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.943µs args {[int64 40], [string "こんにちわ世界040"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.53µs args {[int64 41], [string "こんにちわ世界041"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 15.856µs args {[int64 42], [string "こんにちわ世界042"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.538µs args {[int64 43], [string "こんにちわ世界043"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.102µs args {[int64 44], [string "こんにちわ世界044"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.988µs args {[int64 45], [string "こんにちわ世界045"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 52.97µs args {[int64 46], [string "こんにちわ世界046"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.953µs args {[int64 47], [string "こんにちわ世界047"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.195µs args {[int64 48], [string "こんにちわ世界048"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.988µs args {[int64 49], [string "こんにちわ世界049"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.008µs args {[int64 50], [string "こんにちわ世界050"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.088µs args {[int64 51], [string "こんにちわ世界051"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.912µs args {[int64 52], [string "こんにちわ世界052"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.952µs args {[int64 53], [string "こんにちわ世界053"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 6.333µs args {[int64 54], [string "こんにちわ世界054"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.561µs args {[int64 55], [string "こんにちわ世界055"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.138µs args {[int64 56], [string "こんにちわ世界056"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.897µs args {[int64 57], [string "こんにちわ世界057"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.327µs args {[int64 58], [string "こんにちわ世界058"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.023µs args {[int64 59], [string "こんにちわ世界059"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.167µs args {[int64 60], [string "こんにちわ世界060"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.057µs args {[int64 61], [string "こんにちわ世界061"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.907µs args {[int64 62], [string "こんにちわ世界062"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.816µs args {[int64 63], [string "こんにちわ世界063"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.41µs args {[int64 64], [string "こんにちわ世界064"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.068µs args {[int64 65], [string "こんにちわ世界065"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.915µs args {[int64 66], [string "こんにちわ世界066"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.376µs args {[int64 67], [string "こんにちわ世界067"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.471µs args {[int64 68], [string "こんにちわ世界068"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.998µs args {[int64 69], [string "こんにちわ世界069"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.952µs args {[int64 70], [string "こんにちわ世界070"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.168µs args {[int64 71], [string "こんにちわ世界071"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 5.76µs args {[int64 72], [string "こんにちわ世界072"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.509µs args {[int64 73], [string "こんにちわ世界073"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.631µs args {[int64 74], [string "こんにちわ世界074"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.075µs args {[int64 75], [string "こんにちわ世界075"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.602µs args {[int64 76], [string "こんにちわ世界076"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.976µs args {[int64 77], [string "こんにちわ世界077"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.948µs args {[int64 78], [string "こんにちわ世界078"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.919µs args {[int64 79], [string "こんにちわ世界079"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.934µs args {[int64 80], [string "こんにちわ世界080"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 3.877µs args {[int64 81], [string "こんにちわ世界081"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.031µs args {[int64 82], [string "こんにちわ世界082"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.898µs args {[int64 83], [string "こんにちわ世界083"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.927µs args {[int64 84], [string "こんにちわ世界084"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.214µs args {[int64 85], [string "こんにちわ世界085"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 13.054µs args {[int64 86], [string "こんにちわ世界086"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.727µs args {[int64 87], [string "こんにちわ世界087"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.146µs args {[int64 88], [string "こんにちわ世界088"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.09µs args {[int64 89], [string "こんにちわ世界089"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 4.627µs args {[int64 90], [string "こんにちわ世界090"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.279µs args {[int64 91], [string "こんにちわ世界091"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.061µs args {[int64 92], [string "こんにちわ世界092"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.473µs args {[int64 93], [string "こんにちわ世界093"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.957µs args {[int64 94], [string "こんにちわ世界094"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 14.493µs args {[int64 95], [string "こんにちわ世界095"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.768µs args {[int64 96], [string "こんにちわ世界096"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 2.096µs args {[int64 97], [string "こんにちわ世界097"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 1.925µs args {[int64 98], [string "こんにちわ世界098"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-stmt-exec [query insert into foo(id, name) values(?, ?) err <nil> duration 9.944µs args {[int64 99], [string "こんにちわ世界099"]}]
2020/01/03 16:45:18 insert into foo(id, name) values(?, ?)
2020/01/03 16:45:18 sql-tx-commit [err <nil> duration 570.78µs]
2020/01/03 16:45:18 sql-stmt-close [err <nil> duration 2.14µs]
2020/01/03 16:45:18 sql-conn-query [query select id, name from foo err <nil> duration 27.331µs args {}]
2020/01/03 16:45:18 select id, name from foo
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 312.694µs]
0 こんにちわ世界000
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 3.849µs]
1 こんにちわ世界001
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.647µs]
2 こんにちわ世界002
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.363µs]
3 こんにちわ世界003
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.337µs]
4 こんにちわ世界004
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.279µs]
5 こんにちわ世界005
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.317µs]
6 こんにちわ世界006
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.303µs]
7 こんにちわ世界007
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.334µs]
8 こんにちわ世界008
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.335µs]
9 こんにちわ世界009
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.407µs]
10 こんにちわ世界010
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.466µs]
11 こんにちわ世界011
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.365µs]
12 こんにちわ世界012
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 20.538µs]
13 こんにちわ世界013
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.944µs]
14 こんにちわ世界014
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.413µs]
15 こんにちわ世界015
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.485µs]
16 こんにちわ世界016
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 6.65µs]
17 こんにちわ世界017
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.498µs]
18 こんにちわ世界018
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.439µs]
19 こんにちわ世界019
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.387µs]
20 こんにちわ世界020
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.421µs]
21 こんにちわ世界021
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.345µs]
22 こんにちわ世界022
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.337µs]
23 こんにちわ世界023
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 4.407µs]
24 こんにちわ世界024
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.7µs]
25 こんにちわ世界025
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.471µs]
26 こんにちわ世界026
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.751µs]
27 こんにちわ世界027
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.974µs]
28 こんにちわ世界028
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.401µs]
29 こんにちわ世界029
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.383µs]
30 こんにちわ世界030
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.421µs]
31 こんにちわ世界031
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.403µs]
32 こんにちわ世界032
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.368µs]
33 こんにちわ世界033
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.324µs]
34 こんにちわ世界034
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.448µs]
35 こんにちわ世界035
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.325µs]
36 こんにちわ世界036
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.368µs]
37 こんにちわ世界037
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.428µs]
38 こんにちわ世界038
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.342µs]
39 こんにちわ世界039
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.292µs]
40 こんにちわ世界040
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 4.384µs]
41 こんにちわ世界041
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.832µs]
42 こんにちわ世界042
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.545µs]
43 こんにちわ世界043
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.452µs]
44 こんにちわ世界044
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.411µs]
45 こんにちわ世界045
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.351µs]
46 こんにちわ世界046
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.302µs]
47 こんにちわ世界047
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.31µs]
48 こんにちわ世界048
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.331µs]
49 こんにちわ世界049
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.323µs]
50 こんにちわ世界050
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.29µs]
51 こんにちわ世界051
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 43.465µs]
52 こんにちわ世界052
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.758µs]
53 こんにちわ世界053
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.394µs]
54 こんにちわ世界054
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.31µs]
55 こんにちわ世界055
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.388µs]
56 こんにちわ世界056
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.344µs]
57 こんにちわ世界057
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.312µs]
58 こんにちわ世界058
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.415µs]
59 こんにちわ世界059
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.348µs]
60 こんにちわ世界060
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.606µs]
61 こんにちわ世界061
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.423µs]
62 こんにちわ世界062
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.315µs]
63 こんにちわ世界063
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.498µs]
64 こんにちわ世界064
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.322µs]
65 こんにちわ世界065
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.319µs]
66 こんにちわ世界066
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 37.226µs]
67 こんにちわ世界067
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 2.212µs]
68 こんにちわ世界068
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.417µs]
69 こんにちわ世界069
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.349µs]
70 こんにちわ世界070
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.568µs]
71 こんにちわ世界071
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.382µs]
72 こんにちわ世界072
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.618µs]
73 こんにちわ世界073
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.37µs]
74 こんにちわ世界074
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.375µs]
75 こんにちわ世界075
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.354µs]
76 こんにちわ世界076
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.333µs]
77 こんにちわ世界077
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.287µs]
78 こんにちわ世界078
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.332µs]
79 こんにちわ世界079
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.405µs]
80 こんにちわ世界080
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.398µs]
81 こんにちわ世界081
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.355µs]
82 こんにちわ世界082
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.345µs]
83 こんにちわ世界083
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.377µs]
84 こんにちわ世界084
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.308µs]
85 こんにちわ世界085
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.336µs]
86 こんにちわ世界086
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.392µs]
87 こんにちわ世界087
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 19.452µs]
88 こんにちわ世界088
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 4.36µs]
89 こんにちわ世界089
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.469µs]
90 こんにちわ世界090
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.43µs]
91 こんにちわ世界091
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.611µs]
92 こんにちわ世界092
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.32µs]
93 こんにちわ世界093
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.351µs]
94 こんにちわ世界094
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.356µs]
95 こんにちわ世界095
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.342µs]
96 こんにちわ世界096
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 11.307µs]
97 こんにちわ世界097
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 2.062µs]
98 こんにちわ世界098
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.65µs]
99 こんにちわ世界099
2020/01/03 16:45:18 sql-rows-next [err EOF duration 7.824µs]
2020/01/03 16:45:18 sql-conn-query [query select id, name from foo where name=? err <nil> duration 42.839µs args {[string "bajindul"]}]
2020/01/03 16:45:18 select id, name from foo where name=?
2020/01/03 16:45:18 sql-rows-next [err EOF duration 56.874µs]
2020/01/03 16:45:18 sql-prepare [query select name from foo where id = ? err <nil> duration 12.007µs]
2020/01/03 16:45:18 sql-stmt-query [query select name from foo where id = ? err <nil> duration 2.576µs args {[string "3"]}]
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 19.188µs]
こんにちわ世界003
2020/01/03 16:45:18 sql-conn-exec [query delete from foo err <nil> duration 1.360468ms args {}]
2020/01/03 16:45:18 delete from foo
2020/01/03 16:45:18 sql-conn-exec [query insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz') err <nil> duration 2.579839ms args {}]
2020/01/03 16:45:18 insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')
2020/01/03 16:45:18 sql-conn-query [query select id, name from foo err <nil> duration 49.574µs args {}]
2020/01/03 16:45:18 select id, name from foo
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 37.139µs]
1 foo
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.801µs]
2 bar
2020/01/03 16:45:18 sql-rows-next [err <nil> duration 1.385µs]
3 baz
2020/01/03 16:45:18 sql-rows-next [err EOF duration 3.985µs]
2020/01/03 16:45:18 sql-stmt-close [err <nil> duration 1.237µs]

===list known Query===
"select id, name from foo where name=?" 
```