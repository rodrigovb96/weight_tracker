# Weight Tracker
---

Input weights daily and then you could plot the weakly measures.

For plotting: 

localhost:8090/<Month>/<Week>

<Month> -> 01, 02, ... 12;
<Week> -> 0, 1, 2, ... 4;


For input a weight;

localhost:PORT/input/<Weight>

<Weight> -> float value i.e 75.6


# Build & Usage:
--- 

Export the input port, for example: 

```bash
	export PORT=8080
```

Install the go module:  
```bash
	go install
```

Build: 
```bash
	go build 
```

Then run it: 
```bash
	./weight_tracker &
```
