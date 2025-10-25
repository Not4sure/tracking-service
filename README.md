# Tracking service
Go REST API service test task.

Project structure is heavily influenced by Domain Driven Design 
and Go DDD approach described in [Go-With-Domain E-book](https://threedots.tech/go-with-the-domain).
On application level, basic yet robust CQRS implemented which 
allows for clear separation between commands and queries.
This enables future refactoring and scaling, for example usage of 
remote sources for data reads and writes, such as database RO replica.
Domain layer in this project happens to be quite thin, and, 
arguably, dumb. This is mostly due to overall small scope of a test 
task and inability to dig deeper into domain knowledge.

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
it's own Docker container to be run as kubernetes cronJob. 


# Day-by-day progress
## Day 1
- Created repo with initial project structure
- Added Event domain layer
- Added application layer with basic CQRS
- Adeed application handlers for event creation and retrieval
- Added in-memory implementation for event repository
- Added http_api port implementation with endpoints for application handlers

## Day 2
- Added postgres migrations for events table
- Added sqlc code generation for database access 
- Created postgres implementation for event repository and updated tests
- Added domain layer and database structure for user activity metrics
- Created application handler for metrics calculation

## Day 3
- Added metrics list api
- Changed http router to chi
- Improved docker-compose configuartion
- Automated migrations in postgres container
- Created minimal react frontend
