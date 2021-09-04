CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1};
USE oauth;
CREATE TABLE access_tokens(access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);
