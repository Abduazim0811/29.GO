CREATE TABLE IF NOT EXISTS friendships (
    user_id INT,
    friend_id INT,
    status VARCHAR(50) NOT NULL,
    PRIMARY KEY (user_id, friend_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (friend_id) REFERENCES users(user_id)
);

