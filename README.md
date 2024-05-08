```
wrkb hashes http://lo:8080/hashes/380501234567

Process "hashes" starts with:
cpu(s)  171.207405
threads 24
mem(b)  837599232
disk(b) 5027954

num|    rps|  latency|  cpu| thr|         rss
---------------------------------------------
  1|  36880|  26.39µs| 0.93|  24|   837599232
  2|  66500|  28.12µs| 1.50|  24|   837599232
  3|  96450|  28.02µs| 2.10|  24|   837599232
  4| 113670|  30.23µs| 2.30|  24|   837599232
  5| 118730|  34.56µs| 2.36|  24|   837599232
  6| 129090|  37.66µs| 2.18|  24|   837599232
  7| 141990|  42.03µs| 2.24|  24|   837599232
  8| 155840| 115.37µs| 2.17|  24|   837648384
  9| 152530|  96.92µs| 2.20|  24|   842055680
 10| 154740|  95.59µs| 2.25|  24|   842104832
 11| 143250| 126.35µs| 2.27|  24|   842104832
 12| 153880| 171.51µs| 2.27|  24|   842104832
 13| 138710| 112.83µs| 2.28|  24|   843218944
 14| 136310|  75.21µs| 2.32|  24|   843218944
 15| 126810| 117.99µs| 2.35|  24|   843218944
 16| 139150|   98.4µs| 2.35|  24|   843218944
 32| 123790| 179.32µs| 2.42|  24|   843218944
 64| 129169| 320.09µs| 2.68|  24|   843218944

Best:
 10| 154740|  95.59µs| 2.25|  24|   842104832
```

```
wrkb hashes http://lo:8080/msisdns/a10ad80c494e707e665d264962b51630

Process "hashes" starts with:
cpu(s)  210.366011
threads 24
mem(b)  843218944
disk(b) 5027954

num|    rps|  latency|  cpu| thr|         rss
---------------------------------------------
  1|  34690|  42.17µs| 0.91|  24|   843251712
  2|  65150|  28.67µs| 1.50|  24|   843251712
  3|  95500|  28.48µs| 2.13|  24|   843251712
  4| 111610|  30.97µs| 2.34|  24|   843251712
  5| 116070|   35.5µs| 2.16|  24|   843251712
  6| 131380|  38.68µs| 2.24|  24|   843251712
  7| 135610|  42.72µs| 2.44|  24|   843251712
  8| 149230| 134.31µs| 2.29|  24|   843268096
  9| 152840|  99.14µs| 2.25|  24|   843284480
 10| 148790| 294.94µs| 2.47|  24|   843284480
 11| 142430| 106.71µs| 2.24|  24|   843825152
 12| 149530| 237.22µs| 2.55|  24|   843825152
 13| 142460| 671.92µs| 2.29|  24|   843825152
 14| 133680| 120.26µs| 2.30|  24|   844365824
 15| 133490| 232.26µs| 2.37|  24|   844365824
 16| 128750|  90.11µs| 2.35|  24|   844382208
 32| 121280| 185.49µs| 2.53|  24|   844382208
 64| 121890| 353.53µs| 2.72|  24|   844382208

Best:
  9| 152840|  99.14µs| 2.25|  24|   843284480

```
