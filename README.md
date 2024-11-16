# go-kafka-demo

## Why?

Implements simple Kafka foo with Go for demonstration purposes.

## Run Locally

* Docker is required to run locally

```bash
docker compose up -d
[+] Running 4/4
 ✔ Network redpanda-quickstart-one-broker_redpanda_network  Created
 ✔ Volume "redpanda-quickstart-one-broker_redpanda-0"       Created
 ✔ Container redpanda-0                                     Started
 ✔ Container redpanda-console                               Started

# Run Quick Demo
go run main.go
2024/11/16 14:20:39 Producing event:  GenericEvent[ID=2, Message=08a97619-7a19-40d7-b94b-07dfaeec2a38, timestamp=2024-11-16 14:20:39.625848 +0000 GMT m=+0.004930803]
2024/11/16 14:20:39 Producing event:  GenericEvent[ID=5, Message=77834248-f41c-44df-9058-c8b4b2311368, timestamp=2024-11-16 14:20:39.625892 +0000 GMT m=+0.004975258]
2024/11/16 14:20:39 Producing event:  GenericEvent[ID=4, Message=ed4e4a20-e438-44a5-85e6-115e523e0564, timestamp=2024-11-16 14:20:39.625959 +0000 GMT m=+0.005042058]
2024/11/16 14:20:39 Producing event:  GenericEvent[ID=3, Message=f1ac8b5b-05f9-4766-a97a-dad754e88e99, timestamp=2024-11-16 14:20:39.62587 +0000 GMT m=+0.004953109]
2024/11/16 14:20:39 Producing event:  GenericEvent[ID=1, Message=48dc29a1-415b-45ad-b6d7-06aa1403cfac, timestamp=2024-11-16 14:20:39.625898 +0000 GMT m=+0.004981163]
2024/11/16 14:20:48 Event received: GenericEvent[ID=4, Message=ed4e4a20-e438-44a5-85e6-115e523e0564, timestamp=0001-01-01 00:00:00 +0000 UTC]
2024/11/16 14:20:48 Event received: GenericEvent[ID=2, Message=08a97619-7a19-40d7-b94b-07dfaeec2a38, timestamp=0001-01-01 00:00:00 +0000 UTC]
2024/11/16 14:20:48 Event received: GenericEvent[ID=1, Message=48dc29a1-415b-45ad-b6d7-06aa1403cfac, timestamp=0001-01-01 00:00:00 +0000 UTC]
2024/11/16 14:20:48 Event received: GenericEvent[ID=5, Message=77834248-f41c-44df-9058-c8b4b2311368, timestamp=0001-01-01 00:00:00 +0000 UTC]
2024/11/16 14:20:48 Event received: GenericEvent[ID=3, Message=f1ac8b5b-05f9-4766-a97a-dad754e88e99, timestamp=0001-01-01 00:00:00 +0000 UTC]
```

* When finished testing run the following to clean the local environment:

```bash
docker compose down -v
```
