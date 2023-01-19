# practical benchmark

Benchmark run on macosX

- Processor: 3.41 GHz Intel Core i5
- RAM: 16 GB 2400 MHz DDR4

## Result

```
goos: darwin
goarch: amd64
pkg: github.com/h4ckm03d/golang-playground/6-practical-benchmark
BenchmarkBase64Encode-4                           687727              1461 ns/op
BenchmarkBase64Decode-4                           678280              1664 ns/op
BenchmarkCompressWrite-4                          142774             10300 ns/op
BenchmarkCompressRead-4                            68023             16344 ns/op
BenchmarkWriteFile-4                                   3         341506472 ns/op
BenchmarkWriteFileBuffered-4                         183           6605207 ns/op
BenchmarkReadFile-4                                    4         252704252 ns/op
BenchmarkReadFileBuffered-4                          192           5698784 ns/op
BenchmarkHashMD5-4                                734155              1685 ns/op
BenchmarkHashSHA1-4                               880683              1427 ns/op
BenchmarkHashSHA256-4                             390978              2939 ns/op
BenchmarkHashSHA512-4                             559508              2214 ns/op
BenchmarkHashSHA3256-4                            305725              4133 ns/op
BenchmarkHashSHA3512-4                            174043              7201 ns/op
BenchmarkHashBLAKE2b256-4                         916990              1494 ns/op
BenchmarkHashBLAKE2b512-4                         935786              1425 ns/op
BenchmarkHTTP-4                                    10593            121536 ns/op
BenchmarkHTTPNoKeepAlive-4                          3434            294166 ns/op
BenchmarkHTTPSNoKeepAlive-4                          410           3059177 ns/op
BenchmarkMapStringKeys-4                        12447582                96.2 ns/op
BenchmarkMapIntKeys-4                           20010943                76.0 ns/op
BenchmarkObjectCreationNoPool-4                 24793286                44.7 ns/op
BenchmarkObjectCreationPool-4                   60609782                19.4 ns/op
BenchmarkParseBool-4                            316795712                3.83 ns/op
BenchmarkParseInt-4                             29305872                52.8 ns/op
BenchmarkParseFloat-4                           17233749                60.5 ns/op
BenchmarkMathRand-4                             71137954                17.0 ns/op
BenchmarkCryptoRand-4                            3863286               345 ns/op
BenchmarkMathRandString-4                        9977487               104 ns/op
BenchmarkCryptoRandString-4                       881052              1244 ns/op
BenchmarkRegexMatchString-4                       187429              6583 ns/op
BenchmarkRegexMatchStringCompiled-4              2184663               562 ns/op
BenchmarkSerializationJSONMarshal-4              2164740               579 ns/op
BenchmarkSerializationJSONUnmarshal-4             523017              2191 ns/op
BenchmarkSerializationProtoBufMarshal-4          6781549               176 ns/op
BenchmarkSerializationProtoBufUnmarshal-4        3365281               353 ns/op
BenchmarkSerializationMessagePackMarshal-4       7418779               172 ns/op
BenchmarkSerializationMessagePackUnmarshal-4     4820314               220 ns/op
BenchmarkSerializationGobMarshal-4               2044286               570 ns/op
BenchmarkSerializationGobUnmarshal-4             1318915               874 ns/op
BenchmarkSliceAppend-4                          625260250                1.68 ns/op
BenchmarkSliceAppendPrealloc-4                  1000000000               0.860 ns/op
BenchmarkSort1000-4                                10000            114650 ns/op
BenchmarkSort10000-4                                 904           1335605 ns/op
BenchmarkSort100000-4                                 69          17160820 ns/op
BenchmarkSort1000000-4                                 5         205216102 ns/op
BenchmarkConcatString-4                          8596020               144 ns/op
BenchmarkConcatBuffer-4                         140304704                7.71 ns/op
BenchmarkConcatBytesGuess-4                     858581476                1.55 ns/op
BenchmarkConcatBytesFix-4                       1000000000               1.08 ns/op
BenchmarkConcatBuilder-4                        460746288                3.07 ns/op
BenchmarkExecute-4                                461127              2734 ns/op
BenchmarkParse-4                                 2729244               397 ns/op
PASS
ok      github.com/h4ckm03d/golang-playground/6-practical-benchmark     87.882s
```
