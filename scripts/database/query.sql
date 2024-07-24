-- users.sql
-- name: CreateUser :one
INSERT INTO users (username, full_name, email, password, profile_url, channel_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, username, full_name, email, password, profile_url, channel_id, created_at, updated_at;

-- name: GetUserById :one
SELECT id, username, full_name, email, password, profile_url, channel_id, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, username, full_name, email, password, profile_url, channel_id, created_at, updated_at
FROM users
WHERE email = $1;

-- name: UpdateUser :exec
UPDATE users
SET username = $2, full_name = $3, email = $4, password = $5, profile_url = $6, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- channel.sql
-- name: CreateChannel :one
INSERT INTO channel (user_id, name, logo, subscriber_count)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, name, logo, subscriber_count, created_at, updated_at;

-- name: GetChannelById :one
SELECT id, user_id, name, logo, subscriber_count, created_at, updated_at
FROM channel
WHERE id = $1;

-- name: GetChannelByUserId :one
SELECT id, user_id, name, logo, subscriber_count, created_at, updated_at
FROM channel
WHERE user_id = $1;
-- name: UpdateChannel :exec
UPDATE channel
SET name = $2, logo = $3, subscriber_count = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
-- name: GetChannels :many
SELECT id, user_id, name, logo, subscriber_count, created_at, updated_at
FROM channel;
-- name: DeleteChannel :exec
DELETE FROM channel
WHERE id = $1;
-- subscriptions.sql
-- name: CreateSubscription :one
INSERT INTO subscriptions (user_id, channel_id)
VALUES ($1, $2)
RETURNING id, user_id, channel_id, subscribed_at, created_at;

-- name: GetSubscription :one
SELECT id, user_id, channel_id, subscribed_at, created_at
FROM subscriptions
WHERE id = $1;

-- name: GetSubscriptionsByUserId :many
SELECT id, user_id, channel_id, subscribed_at, created_at
FROM subscriptions
WHERE user_id = $1;

-- name: GetSubscriptionsByChannelId :many
SELECT id, user_id, channel_id, subscribed_at, created_at
FROM subscriptions
WHERE channel_id = $1;

-- name: DeleteSubscription :exec
DELETE FROM subscriptions
WHERE user_id = $1 AND channel_id = $2;

-- video.sql
-- name: CreateVideo :one
INSERT INTO video (title, description, status, comments_available, duration_seconds, thumbnail, _1080p_url, _720p_url, _480p_url, _360p_url, channel_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, title, description, status, comments_available, duration_seconds, thumbnail, _1080p_url, _720p_url, _480p_url, _360p_url, channel_id, created_at, updated_at;

-- name: GetVideoById :one
SELECT id, title, description, status, comments_available, duration_seconds, thumbnail, _1080p_url, _720p_url, _480p_url, _360p_url, channel_id, created_at, updated_at
FROM video
WHERE id = $1;
-- name: GetVideoByChannelId :many 
SELECT id, title, description, status, comments_available, duration_seconds, thumbnail, _1080p_url, _720p_url, _480p_url, _360p_url, channel_id, created_at, updated_at
FROM video
WHERE channel_id = $1;

-- name: UpdateVideo :exec
UPDATE video
SET title = $2, description = $3, status = $4, comments_available = $5, duration_seconds = $6, thumbnail = $7, _1080p_url = $8, _720p_url = $9, _480p_url = $10, _360p_url = $11, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteVideo :exec
DELETE FROM video
WHERE id = $1;

-- comment.sql
-- name: CreateComment :one
INSERT INTO comment (text, video_id, user_id, parent_id, replies_count)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, text, video_id, user_id, parent_id, replies_count, created_at;

-- name: GetCommentById :one
SELECT id, text, video_id, user_id, parent_id, replies_count, created_at
FROM comment
WHERE id = $1;

-- name: GetCommentsByVideoId :many
SELECT id, text, video_id, user_id, parent_id, replies_count, created_at
FROM comment
WHERE video_id = $1;

-- name: UpdateComment :exec
UPDATE comment
SET text = $2, replies_count = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteComment :exec
DELETE FROM comment
WHERE id = $1;


-- notification.sql
-- name: CreateNotification :one
INSERT INTO notification (user_id, type, content, status)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, type, content, status, created_at, updated_at;

-- name: GetNotificationById :one
SELECT id, user_id, type, content, status, created_at, updated_at
FROM notification
WHERE id = $1;

-- name: GetNotificationsByUserId :many
SELECT id, user_id, type, content, status, created_at, updated_at
FROM notification
WHERE user_id = $1;


-- name: UpdateNotification :exec
UPDATE notification
SET type = $2, content = $3, status = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteNotification :exec
DELETE FROM notification
WHERE id = $1;

-- watch_history.sql
-- name: CreateWatchHistory :one
INSERT INTO watch_history (user_id, video_id, watched_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, video_id, watched_at, created_at, updated_at;

-- name: GetWatchHistoryById :one
SELECT id, user_id, video_id, watched_at, created_at, updated_at
FROM watch_history
WHERE id = $1;

-- name: GetWatchHistoryByUserId :many
SELECT id, user_id, video_id, watched_at, created_at, updated_at
FROM watch_history
WHERE user_id = $1;

-- name: UpdateWatchHistory :exec
UPDATE watch_history
SET watched_at = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteWatchHistory :exec
DELETE FROM watch_history
WHERE id = $1;

