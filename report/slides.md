---
marp: true
theme: default
---

# Pala, An Election Management System
By Adithya Nair, Kausik Muthukumar, Ananthapadmanabhan Nair

---

## Research
1. Golang
3. Gin
2. wrk

---
## Why Golang
Flask
```sh
❯ ./wrk -t12 -c400 -d30s http://127.0.0.1:5000/voters/
Running 30s test @ http://127.0.0.1:5000/voters/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   434.04ms  120.53ms 745.01ms   82.33%
    Req/Sec    76.11     48.29   353.00     69.45%
  25894 requests in 30.04s, 9.58MB read
  Non-2xx or 3xx responses: 25894
Requests/sec:    861.88
Transfer/sec:    326.59KB
```
---
## Why Golang
Node.js
```sh
❯ ./wrk -t12 -c400 -d30s http://127.0.0.1:3000/voters/
Running 30s test @ http://127.0.0.1:3000/voters/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    91.25ms   84.58ms   1.99s    98.95%
    Req/Sec   376.13     86.99     1.15k    88.59%
  130990 requests in 30.04s, 39.60MB read
  Socket errors: connect 0, read 0, write 0, timeout 187
Requests/sec:   4360.39
Transfer/sec:      1.32MB
```

---
## Why Golang

Golang
```sh
❯ wrk -t12 -c400 -d30s http://127.0.0.1:8080/voters/
Running 30s test @ http://127.0.0.1:8080/voters/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    16.24ms   13.15ms 147.72ms   66.07%
    Req/Sec     2.24k   256.41     8.39k    87.90%
  804250 requests in 30.10s, 221.66MB read
Requests/sec:  26720.77
Transfer/sec:      7.36MB
```

---

## What Has Been Done?
- The backend has been fully implemented
- Mapping of Data model is complete

---

## What Needs To Be Done?
- Frontend 
- The DB architecture needs to be swapped.
- Benchmarking with wrk

---
## What Needs To Be Done
- Concurrency with goroutines
- Investigating and implementing load balancer
- Implement scheduling algorithms for voting booth allocation.
- Monitoring system

---

## Workload Distribution

1. Backend implementation - Adithya
2. Data Implementation - Kausik
3. Frontend Development - Ananthapadmanabhan
