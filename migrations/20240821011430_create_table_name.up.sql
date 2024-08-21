create table user(id UInt64, name String, date Date) ENGINE = MergeTree() PARTITION BY toYYYYMMDD(date)
ORDER BY
  (date, id);
