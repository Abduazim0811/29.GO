UPDATE friendships SET status = 'accepted'
WHERE user_id = $1 AND friend_id = $2 AND status = 'pending';