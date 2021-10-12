[Basics Of Tuning Checkpoints](https://www.2ndquadrant.com/en/blog/basics-of-tuning-checkpoints/)

*Description*: On systems doing non-trivial number of writes, tuning checkpoints is crucial for getting good performance. So let me walk you through the checkpoints – what they do and how to tune them in PostgreSQL.

*Labels*: #PostgreSQL #Checkpoints #Performance

*Examples*:

```sql
checkpoint_timeout = 30min          -- the period for starting checkpoints
max_wal_size = 1GB                  -- the max size of WAL triggering new checkpoints ahead of timeout
checkpoint_completion_target = 0.5  -- how long the checkpoint will take time until it starts fsync
```
