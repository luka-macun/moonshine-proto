DROP INDEX IF EXISTS idx_channels_created_by;
DROP INDEX IF EXISTS idx_messages_user;
DROP INDEX IF EXISTS idx_messages_channel_created;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS channels;
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "uuid-ossp";
