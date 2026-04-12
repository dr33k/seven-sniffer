PORT SNIFFER

Scans ports for a given IP address host logging their status on the console

Writes TCP packets to files generated for each port

RUN

```
$ go install ./cmd/s-sniff
$ s-sniff -host google.com -ports=80,90,100
```

Compiled with: go1.25.7
