# parallel

where `xargs` meets `foreman` or `goreman` for controlled, concurrent execution of shell commands.

inspired by https://www.gnu.org/software/parallel/

## builds

[![Build Status](https://travis-ci.org/korovkin/parallel.svg)](https://travis-ci.org/korovkin/parallel)


## build:
```
go build cmd/*.go
```

or just:

```
make travis
```

## Demos:

### compress/decompress 10 large files (serially):

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

### compress/decompress 10 large files (concurrently with parallel):

```
time for f in *.txt ; do echo gzip $f; done | $GOPATH/bin/parallel -j 10
[       98.86µs 11:36:51 0] concurrency limit: 10
[     141.504µs 11:36:51 0] reading from stdin...
[     215.997µs 11:36:51 9] run: 'gzip 9.txt'
[     258.119µs 11:36:51 1] run: 'gzip 10.txt'
[     279.051µs 11:36:51 3] run: 'gzip 3.txt'
[    5.380117ms 11:36:51 7] run: 'gzip 7.txt'
[    5.446718ms 11:36:51 5] run: 'gzip 5.txt'
[    5.522043ms 11:36:51 2] run: 'gzip 2.txt'
[    7.693391ms 11:36:51 8] run: 'gzip 8.txt'
[    9.895031ms 11:36:51 6] run: 'gzip 6.txt'
[   10.003485ms 11:36:51 0] run: 'gzip 1.txt'
[   10.040754ms 11:36:51 4] run: 'gzip 4.txt'
[  2.670755505s 11:36:54 9] done: dt: 2.670561141s
[  2.723745718s 11:36:54 1] done: dt: 2.723551945s
[  2.732728457s 11:36:54 3] done: dt: 2.732508999s
[  2.748645989s 11:36:54 8] done: dt: 2.748406607s
[  2.762459047s 11:36:54 5] done: dt: 2.762266508s
[  2.777302364s 11:36:54 7] done: dt: 2.777068367s
[  2.785686994s 11:36:54 4] done: dt: 2.785451804s
[  2.796846523s 11:36:54 2] done: dt: 2.796563159s
[  2.810195841s 11:36:54 0] done: dt: 2.809967849s
[  2.817260242s 11:36:54 6] done: dt: 2.81703108s
[  2.817275381s 11:36:54 0] all done: dt: 2.817269836s

real	0m2.831s
user	0m16.502s
sys	0m0.831s


time for f in *.txt.gz ; do echo gunzip $f; done | $GOPATH/bin/parallel -j 10
[     105.518µs 11:41:18 0] concurrency limit: 10
[     150.091µs 11:41:18 0] reading from stdin...
[     232.374µs 11:41:18 9] run: 'gunzip 9.txt.gz'
[      279.31µs 11:41:18 1] run: 'gunzip 10.txt.gz'
[    1.434552ms 11:41:18 4] run: 'gunzip 4.txt.gz'
[    1.506747ms 11:41:18 6] run: 'gunzip 6.txt.gz'
[    4.912507ms 11:41:18 2] run: 'gunzip 2.txt.gz'
[    8.693375ms 11:41:18 0] run: 'gunzip 1.txt.gz'
[    8.733623ms 11:41:18 5] run: 'gunzip 5.txt.gz'
[   10.506841ms 11:41:18 3] run: 'gunzip 3.txt.gz'
[   12.224203ms 11:41:18 7] run: 'gunzip 7.txt.gz'
[   14.152521ms 11:41:18 8] run: 'gunzip 8.txt.gz'
[  1.191842637s 11:41:19 1] done: dt: 1.191617208s
[  1.223837314s 11:41:19 9] done: dt: 1.223637668s
[  1.240858566s 11:41:19 4] done: dt: 1.240603349s
[  1.255985175s 11:41:19 8] done: dt: 1.255724068s
[  1.273008068s 11:41:19 5] done: dt: 1.272772052s
[  1.282475414s 11:41:19 2] done: dt: 1.282202009s
[  1.288824013s 11:41:19 0] done: dt: 1.288580036s
[  1.294614678s 11:41:19 6] done: dt: 1.294366962s
[  1.295771009s 11:41:19 3] done: dt: 1.295523197s
[  1.299966607s 11:41:19 7] done: dt: 1.29971294s
[  1.299979953s 11:41:19 0] all done: dt: 1.299975556s

real	0m1.312s
user	0m0.954s
sys	0m0.674s

```


### run concurrent jobs (pings)


```
echo -e " ping www.google.com\n ping apple.com\n \n\n" | ./parallel -j 10 -v

[56.609µs        15:06:07 000 I] concurrency limit: 10
[204.437µs       15:06:07 004 I] start: ''
[231.207µs       15:06:07 001 I] start: 'ping apple.com'
[269.186µs       15:06:07 002 I] start: ''
[307.086µs       15:06:07 000 I] start: 'ping www.google.com'
[400.384µs       15:06:07 003 I] start: ''
[676.3µs         15:06:07 004 I] execute: done: dt: 478.343µs
[749.079µs       15:06:07 002 I] execute: done: dt: 517.862µs
[851.008µs       15:06:07 003 I] execute: done: dt: 508.883µs
[29.482275ms     15:06:07 001 I] PING apple.com (17.253.144.10) 56(84) bytes of data.
[29.496675ms     15:06:07 001 I] 64 bytes from icloud.com (17.253.144.10): icmp_seq=1 ttl=58 time=26.2 ms
[99.010999ms     15:06:07 000 I] PING www.google.com(yx-in-x69.1e100.net (2607:f8b0:4002:c08::69)) 56 data bytes
[99.028179ms     15:06:07 000 I] 64 bytes from yx-in-x69.1e100.net (2607:f8b0:4002:c08::69): icmp_seq=1 ttl=103 time=92.8 ms
[1.027270463s    15:06:07 001 I] 64 bytes from icloud.com (17.253.144.10): icmp_seq=2 ttl=58 time=22.4 ms
[1.100067039s    15:06:07 000 I] 64 bytes from yx-in-x69.1e100.net (2607:f8b0:4002:c08::69): icmp_seq=2 ttl=102 time=92.8 ms

```
