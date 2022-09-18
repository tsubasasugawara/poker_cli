CREATE TABLE IF NOT EXISTS participants (
    room_id VARCHAR(36) NOT NULL,
    player_id VARCHAR(36) NOT NULL UNIQUE,
    FOREIGN KEY(room_id) REFERENCES rooms(id)
);
