## Quick start

### Docker-compose

just run `docker-compose up --build -d`

### Go build

Builds like any go program. Takes three env variables:
```
PORT
DEVICES_PATH
HOST
```

`DEVICES_PATH` is the path to the csv file. By default it will use port 6733 and the path `./devices.csv` relative to the run directory. For host, it will use 127.0.0.1, however, when running docker compose, it will use 0.0.0.0 so it binds correctly.


## Tests
To run tests with coverage, run `go test ./... -cover`

## Results
The results of running the program with the simulator can be found in `results.txt`. The results match the expected output 100%

## Analysis

### Challenge
This exercise took just a little longer than 1 working day. 

The main challenge was deciding how to represent the data and whether to use any external databases. After looking at the requirements and calculations logic, a simple in memory noSQL db seemed to be the best fit.

### Extending the data model
Currently, each heartbeat timestamp is not stored neither is the sent_at timestamp. This was done to keep the stats lookup more efficient. In a real world example, all the data would be stored either in an RDS db, a noSQL db or file storage for future analysis. However, this would not change the nature of the `summary db` (which is what is implemented in this project). As in, one solution would be to store all heart beats in an heart beat RDS table, sharded and partitioned by device and timestamp. The table would be used to get deep insights or generate reports. There will still be a summary db for fast lookup, either in a cache db (like Redis) or a noSQL db (like Dyammo). The program would ensure whenever raw data is ingested, the summary db is also updated.

The project also writes to the db one entry at a time, which would not be efficient at scale. Once at high capacity, the backend can be divided into a server and a collector. The server would handle the reads of the device, and for every write operation, it would write to an event queue. The collector would then read in batches from the event queue and batch insert the data, both in the big db and the `summary db`

### Time complexity
Since the data is represented via noSQL db and every read and write does not expand the data, its bottle neck is just the lookup and save. Therefore the time complexity is `O(1)` for best case and `O(n)` for worse case.

## Client Code
This repo uses openapi for client file generation. To generate the client, run:

```
./generate.sh
```

Unfortunately, there was a bug in the openapi.json spec. It seems that upload_time for UploadStatsRequest is of type int32 when it should be int64. To mitigate this, a custom struct was made instead of using the client. In a real world example, the openapi.json would be fixed.

## Decisions

### Structure
Ideally the business logic should be separate from the datastore package which should only care about querying and saving data. However, since the business logic was simple, the logic was not separated. In a real world example, the service will call the model package which will call the datastore package.

### Database choice
Given that the operations are just a read and write for specific devices, without any relational logic, a NoSQL data model was chosen. For simplicity, an in-memory map was chosen.

To avoid read/write concurrency conflicts, read/write locks are used. There is also a test to ensure these operations can be done in parallel.

Also, given the nature of the calculation, there was no need to store all timestamps and average them everytime, instead just the min, max, sum and counts are stored. This makes it very efficient and performant.

### Sum data type
Instead of storing all the up times, a sum and a count is stored to easily increment and calculate the average. Since the up_time being provided in the test cases were bigger than int32 max value, the sum is stored as a big.Int. This ensures that the sum remains accurate and does not overflow.

## Future Enhancements
The software can be improved by:
- Improving logging (having different log levels)
- Storing the raw data
- Health check endpoint
- Endpoint to indicate whether the device is online or not (the problem stated what offline means, but there was no use of it)
- Summary endpoint (how many devices offline vs online)
