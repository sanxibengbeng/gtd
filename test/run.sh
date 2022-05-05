## https://github.com/timotta/wrk-scripts
wrk -c 101 -t 4 -d 10s -s multi.lua --latency http://localhost:8890