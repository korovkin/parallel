# parallel

where `xargs` meets `foreman` or `goreman` for controlled, concurrent execution of shell commands.

## simple example:

```
for f in `seq 1 10`; do echo "ls -la > /dev/null && sleep 1"; done | \
    $GOPATH/bin/parallel -j 5

[     109.494µs 11:25:29 0] concurrency limit: 5
[     154.492µs 11:25:29 0] reading from stdin...
[     209.559µs 11:25:29 4] run: 'ls -la > /dev/null && sleep 1'
[     235.505µs 11:25:29 1] run: 'ls -la > /dev/null && sleep 1'
[    1.651572ms 11:25:29 2] run: 'ls -la > /dev/null && sleep 1'
[    1.679727ms 11:25:29 3] run: 'ls -la > /dev/null && sleep 1'
[     5.23294ms 11:25:29 0] run: 'ls -la > /dev/null && sleep 1'
[  1.012679831s 11:25:30 4] done: dt: 1.012471996s
[  1.012741918s 11:25:30 4] run: 'ls -la > /dev/null && sleep 1'
[  1.015487034s 11:25:30 1] done: dt: 1.015267222s
[  1.015502072s 11:25:30 2] done: dt: 1.015274677s
[  1.015576779s 11:25:30 2] run: 'ls -la > /dev/null && sleep 1'
[  1.015596524s 11:25:30 1] run: 'ls -la > /dev/null && sleep 1'
[  1.019696765s 11:25:30 3] done: dt: 1.017444252s
[   1.01976208s 11:25:30 3] run: 'ls -la > /dev/null && sleep 1'
[  1.021896922s 11:25:30 0] done: dt: 1.020257745s
[  1.021989868s 11:25:30 0] run: 'ls -la > /dev/null && sleep 1'
[  2.030379779s 11:25:31 4] done: dt: 1.01763198s
[   2.03385578s 11:25:31 2] done: dt: 1.018273997s
[  2.033880667s 11:25:31 1] done: dt: 1.018269191s
[  2.036426394s 11:25:31 3] done: dt: 1.016662778s
[  2.036454777s 11:25:31 0] done: dt: 1.014434093s
[  2.036471295s 11:25:31 0] all done: dt: 2.036464688s

```

## decompress 10 large files serial:

```
 # ls -sh
total 2224400
222440 1.txt  222440 2.txt  222440 4.txt  222440 6.txt  222440 8.txt
222440 10.txt 222440 3.txt  222440 5.txt  222440 7.txt  222440 9.txt

 # time gzip *.txt
real	0m11.662s
user	0m11.001s
sys	0m0.564s

 # time gunzip *.gz
real	0m1.750s
user	0m0.755s
sys	0m0.466s
```
