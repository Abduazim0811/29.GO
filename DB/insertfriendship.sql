INSERT INTO friendships (user_id, friend_id, status)
SELECT $1, $2, 'pending'
WHERE NOT EXISTS (
    SELECT 1 FROM friendships
    WHERE user_id = $1 AND friend_id = $2
);