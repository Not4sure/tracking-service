The Task: Service for Tracking User Activity

Build a Go REST API service that records user activity events and produces daily aggregated statistics.

The service should be able to:
Accept creation of an activity event.
Example event:
{
  "user_id": 42,
  "action": "page_view",
  "metadata": {"page": "/home"}
  ... (The candidate may modify or extend the event structure as needed.)
}
Retrieve events filtered by user and date range.
Run a background job \ cron job every 4 hours that calculates and saves the number of events created by each user during that period.

Requirements:
Go service runnable via Docker.
Relational database (PostgreSQL or MySQL).
Minimal React client to check user events.

Optional (nice to have):
Metrics and logs monitoring using Grafana.

Deliverables:
Public repository link.
README with: brief description, run instructions (local + Docker), sample requests, daily job description, notes on any optional parts.
Example env file.

Deadline:
2–3 days.
