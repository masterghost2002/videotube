-- Migration file for PostgreSQL database

-- Enable extension for auto-incrementing primary keys
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create users table
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR NOT NULL UNIQUE,
  full_name VARCHAR NOT NULL,
  email VARCHAR NOT NULL UNIQUE,
  password VARCHAR NOT NULL,
  profile_url VARCHAR,
  channel_id BIGINT UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create an index on the email column
CREATE INDEX idx_users_email ON users(email);

-- Create channel table
CREATE TABLE channel (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT UNIQUE REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  name VARCHAR(64) NOT NULL,
  logo VARCHAR,
  subscriber_count BIGINT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add foreign key constraint to users for channel_id
ALTER TABLE users ADD CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES channel(id) ON DELETE CASCADE;

-- Create subscriptions table
CREATE TABLE subscriptions (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  channel_id BIGINT NOT NULL REFERENCES channel(id) ON DELETE CASCADE,
  subscribed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (user_id, channel_id)
);

-- Create video table
CREATE TYPE VideoStatus AS ENUM ('processing', 'available');

CREATE TABLE video (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR NOT NULL,
  description VARCHAR,
  status VideoStatus DEFAULT 'processing',
  comments_available BOOLEAN DEFAULT TRUE,
  duration_seconds BIGINT,
  thumbnail VARCHAR,
  _1080p_url VARCHAR,
  _720p_url VARCHAR,
  _480p_url VARCHAR,
  _360p_url VARCHAR,
  channel_id BIGINT NOT NULL REFERENCES channel(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create comment table
CREATE TABLE comment (
  id BIGSERIAL PRIMARY KEY,
  text VARCHAR NOT NULL,
  video_id BIGINT NOT NULL REFERENCES video(id) ON DELETE CASCADE,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  parent_id BIGINT REFERENCES comment(id) ON DELETE CASCADE DEFAULT NULL,
  replies_count BIGINT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (user_id, video_id)
);

-- Create notification table
CREATE TYPE NotificationStatus AS ENUM ('READ', 'UNREAD');
CREATE TYPE NotificationType AS ENUM ('COMMENT', 'NEW_VIDEO', 'VIDEO_STATUS', 'SUBSCRIPTION');

CREATE TABLE notification (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  type NotificationType DEFAULT 'NEW_VIDEO',
  content VARCHAR NOT NULL,
  status NotificationStatus DEFAULT 'UNREAD',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create watch_history table
CREATE TABLE watch_history (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  video_id BIGINT NOT NULL REFERENCES video(id) ON DELETE CASCADE,
  watched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
