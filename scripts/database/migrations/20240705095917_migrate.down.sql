
-- init_schema.down.sql
UPDATE schema_migrations SET version = <last_successful_version> WHERE version = -1;

-- Drop watch_history table
DROP TABLE IF EXISTS watch_history;

-- Drop notification table
DROP TABLE IF EXISTS notification;

-- Drop comment table
DROP TABLE IF EXISTS comment;

-- Drop video table
DROP TABLE IF EXISTS video;

-- Drop subscriptions table
DROP TABLE IF EXISTS subscriptions;

-- Drop channel table
DROP TABLE IF EXISTS channel;

-- Drop users table
DROP TABLE IF EXISTS users;

-- Drop types
DROP TYPE IF EXISTS NotificationStatus;
DROP TYPE IF EXISTS NotificationType;
DROP TYPE IF EXISTS VideoStatus;

-- Disable extension for auto-incrementing primary keys
DROP EXTENSION IF EXISTS "uuid-ossp";
