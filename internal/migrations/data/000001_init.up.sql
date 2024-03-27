begin;

create extension if not exists "uuid-ossp";

create table if not exists users(
                                    id UUID DEFAULT uuid_generate_v4() primary key ,
                                    name varchar(20) not null ,
                                    email varchar unique not null ,
                                    password varchar not null
);
CREATE INDEX if not exists id_users ON users (id);
CREATE TABLE IF NOT EXISTS refresh_tokens(
                                             id uuid DEFAULT uuid_generate_v4() primary key ,
                                             user_id uuid references users(id) on delete cascade not null ,
                                             token varchar(255) not null unique ,
                                             expires_at timestamp not null
);
CREATE INDEX if not exists id_token ON refresh_tokens (id);
CREATE table followers(
                          lead_id uuid references users (id) on delete cascade ,
                          follower uuid references users(id) on delete cascade,
                          primary key (lead_id,follower)
);
CREATE TABLE tweets (
                        id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                        user_id uuid REFERENCES users(id),
                        content TEXT,
                        media_url TEXT,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE likes (
                       user_id INTEGER REFERENCES users(id),
                       tweet_id INTEGER REFERENCES tweets(id),
                       PRIMARY KEY (user_id, tweet_id)
);

CREATE TABLE retweets (
                          user_id INTEGER REFERENCES users(id),
                          tweet_id uuid REFERENCES tweets(id),
                          PRIMARY KEY (user_id, tweet_id)
);
CREATE INDEX idx_tweet_content ON tweets USING gin(to_tsvector('english', content));

commit;