# parallel

where `xargs` meets `foreman` or `goreman` for controlled, concurrent execution of shell commands.

inspired by https://www.gnu.org/software/parallel/

## builds

[![Build Status](https://travis-ci.org/korovkin/parallel.svg)](https://travis-ci.org/korovkin/limiter)


## install:

executing the following will install the tool into: `$GOPATH/bin/parallel`
```
go get github.com:korovkin/parallel.git/...
```

## compress/decompress 10 large files serial:

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

## compress/decompress 10 large files concurrently:

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
