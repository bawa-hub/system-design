# Exhaustive Redis masterclass curriculum (course outline)

Course intro & how to use this syllabus
Fundamentals & data types (strings, lists, sets, zsets, hashes, bitmaps, HLL, geo, streams) ← we start here
Installation & running Redis (binary, Docker, cloud-managed Redis)

Redis client usage & connection best practices (connection pooling, timeouts, pipelining)

Persistence & durability (RDB snapshots, AOF, rewrite, trade-offs)

Memory management & tuning (maxmemory, policies, memory fragmentation, jemalloc)

Eviction strategies & caching patterns (cache-aside, write-through, write-back, TTL strategies)

Replication & high-availability (primary-replica, partial sync, failover basics)

Redis Sentinel & HA operations (setup, failover behavior, caveats)

Redis Cluster & sharding internals (hash slots, rebalancing, cross-slot ops, topology)

Transactions, optimistic locking, WATCH, MULTI/EXEC

Scripting with Lua (EVAL, atomic scripts, common pitfalls)

Pub/Sub basics + advanced patterns (fanout, channel design, scaling pub/sub)

Streams & consumer groups (XADD, XREADGROUP, trimming, acking, design patterns)

Advanced data structures & specialized modules: RedisJSON, RediSearch, RedisTimeSeries, RedisBloom, RedisGraph, RedisGears

Distributed locks & coordination patterns (Redlock, pitfalls, leader election)

Queues, job processing and durable work queues with Streams (designing worker groups, retries, DLQ)

Rate limiting & sliding window algorithms (token bucket, leaky bucket, fixed window, sliding log)

Real-time analytics patterns (rolling counters, top-k, sketches, HyperLogLog)

Geospatial indexing & queries with GEO* commands

Observability & monitoring (INFO, slowlog, latency, Prometheus exporters, alerts)

Benchmarking & performance tooling (redis-benchmark, memtier_benchmark, profiling)

Security & hardening (AUTH, ACLs, TLS, network segregation, secrets)

Deployment & ops at scale (Docker Compose, k8s StatefulSets, Helm, rolling upgrades, backup/restore)

Scaling patterns & architecture decisions (multi-tier cache, cache warming, hot keys, client-side sharding)

Internals deep-dive (data encodings, event loop, I/O threads, copy-on-write on fork, RDB/AOF internals, replication protocol PSYNC)

Redis modules: building custom modules (C API), when to write a module vs use existing ones

Redis enterprise / Active-Active (CRDT) & Redis on Flash (concepts)

Testing & chaos engineering (fault injection, failover tests, consistency testing)

Cost & capacity planning (memory sizing, instance sizing, cost tradeoffs)

Common anti-patterns & gotchas (hot keys, long blocking commands, large key scans)

Interview prep — top Redis questions + whiteboard problems + system design patterns using Redis

Capstone projects (build full products using Redis):

Real-time leaderboard + analytics

Distributed job queue with guaranteed delivery

Chat / presence system using Streams + Pub/Sub

Time-series analytics pipeline with RedisTimeSeries

Next steps & continuous learning (papers, RedisConf talks, books, Redis University courses)