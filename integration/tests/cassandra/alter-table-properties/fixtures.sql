create table users (
  id int, 
  login varchar, 
  name varchar, 
  primary key (id)
) 
with 
  bloom_filter_fp_chance = 0.01 AND
  caching = {'keys': 'NONE', 'rows_per_partition': '10'} AND
  comment = 'test' AND 
  compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'} AND  
  compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'} AND 
  crc_check_chance = 1.0 AND 
  dclocal_read_repair_chance = 0.1 AND 
  default_time_to_live = 0 AND 
  max_index_interval = 2048 AND 
  memtable_flush_period_in_ms = 0 AND 
  min_index_interval = 128 AND 
  speculative_retry = '99PERCENTILE';
