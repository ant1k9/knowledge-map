[Autovacuum Tuning Basics](https://www.2ndquadrant.com/en/blog/autovacuum-tuning-basics/)

*Description*: The main focus of this blog post is tuning – what configuration options are there, rules of thumb, etc.

*Labels*: #PostgreSQL #Autovacuum #Performance

*Examples*:

```sql
autovacuum_vacuum_scale_factor = 0   # the relative size of dead tuples
autovacuum_vacuum_threshold = 10000  # the absolute size of dead tuples
autovacuum_vacuum_cost_delay = 20ms  # the delay between autovacuum passes
autovacuum_vacuum_cost_limit = 200   # cost summed from params below:

vacuum_cost_page_hit = 1
vacuum_cost_page_miss = 10
vacuum_cost_page_dirty = 20
```
