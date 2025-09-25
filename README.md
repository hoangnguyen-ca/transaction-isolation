# transaction isolation

A set of experiments to test how PostgreSQL(PG) handles transaction isolation

From PostgreSQL docs (https://www.postgresql.org/docs/17/transaction-iso.html):

|Isolation Level|Dirty Read|Nonrepeatable Read|Phantom Read|Serialization Anomaly|
|---|---|---|---|---|
|Read uncommitted|Allowed, but not in PG|Possible|Possible|Possible|
|Read committed|Not possible|Possible|Possible|Possible|
|Repeatable read|Not possible|Not possible|Allowed, but not in PG|Possible|
|Serializable|Not possible|Not possible|Not possible|Not possible|