# Tracking service
Go REST API service test task.

## Overview
Project structure is heavily influenced by Domain Driven Design 
and Go DDD approach described in [Go-With-Domain E-book](https://threedots.tech/go-with-the-domain).
On application level, basic yet robust CQRS implemented which 
allows for clear separation between commands and queries.
This enables future refactoring and scaling, for example usage of 
remote sources for data reads and writes, such as database RO replica.
Domain layer in this project happens to be quite thin, and, 
arguably, dumb. This is mostly due to overall small scope of a test 
task and inability to dig deeper into domain knowledge.

## User activity metrics calculation
Recurrent user activity metrics calculation implemented as simple 
command which calculates metric values for each user and stores them.
Each time calculation command is called, it calculates those metrics for 
previous 4-hour time window. Command operation intentionally made 
idempotent, so multiple calls won't be harmful.
Decision to calculate 4-hour time window boundaries in-app is made 
to not rely on external trigger accuracy. This guaranties that each metric
record is calculated exactly for 4-hour interval and no events are missed
or counted twice.

In this test task, metrics job is triggered by gocron library, but generally 
it's designed to be triggered externally by http request. This approach 
allows for horizontal scaling of the app without multiple simultaneous runs of 
a job. Also there's separate calculate_metrics cmd which can be packaged to
it's own Docker container and executed kubernetes cronJob, for example. 

## Run instructions
To start application fully in docker, run:
```bash
make start
```
or
```bash
docker-compose up -d
```

API will be available at http://localhost:8080.
React app with for querying user events will be at http://localhost:3000 
Simple adminer ui for postgres will be at http://localhost:8000

Postgres credentials for local development are
user `tracking_service`
password `password`
DB `tracking_service`

To see application's logs, run:
```bash
docker-compose logs -f backend
```

To stop all docker containers, run:
```bash
make stop
```
or 
```bash
docker-compose down 
```


## API 
Api docs are available as HTTPie collection in `httpie-collection-tracking-service.yaml` file.

### Create event
```bash
curl --request POST \
  --url http://localhost:8080/api/event \
  --header 'Content-Type: application/json' \
  --data '{
  "user_id": 321,
  "action": "test_action",
  "metadata": {
    "page": "/hello/world"
  }
}'
```

### List events
```bash
curl --request GET \
  --url 'http://localhost:8080/api/events?user_id=321&from=2025-10-23T13%3A07%3A00.000Z&till=2025-10-257T22%3A07%3A00.000Z'
```


### Trigger user activity metrics calculation
```bash
curl --request POST \
  --url http://localhost:8080/api/metrics
```


### List user activity metrics
```bash 
curl --request GET \
  --url 'http://localhost:8080/api/metrics?user_id=321&from=2025-10-23T13%3A07%3A18.392Z&till=2025-10-25T22%3A07%3A18.392Z'
```


## Day-by-day progress
### Day 1
- Created repo with initial project structure
- Added Event domain layer
- Added application layer with basic CQRS
- Adeed application handlers for event creation and retrieval
- Added in-memory implementation for event repository
- Added http_api port implementation with endpoints for application handlers

### Day 2
- Added postgres migrations for events table
- Added sqlc code generation for database access 
- Created postgres implementation for event repository and updated tests
- Added domain layer and database structure for user activity metrics
- Created application handler for metrics calculation

### Day 3
- Added metrics list api
- Changed http router to chi
- Improved docker-compose configuartion
- Automated migrations in postgres container
- Created minimal react frontend
