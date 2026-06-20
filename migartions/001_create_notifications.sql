CREATE TABLE notifications
(
    id VARCHAR(50) PRIMARY KEY,

    user_id VARCHAR(50) NOT NULL,

    channel VARCHAR(20) NOT NULL,

    message TEXT NOT NULL,

    status VARCHAR(20) NOT NULL,

    created_at TIMESTAMP NOT NULL,

    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_notifications_user
ON notifications(user_id);

CREATE INDEX idx_notifications_status
ON notifications(status);


ALTER TABLE notifications
ADD COLUMN retry_count INT NOT NULL DEFAULT 0;

ALTER TABLE notifications
ADD COLUMN next_retry_at TIMESTAMP NULL;

ALTER TABLE notifications
ALTER COLUMN next_retry_at
TYPE TIMESTAMPTZ;