SELECT user.ID, UserName, ParentUserName FROM user
LEFT JOIN (SELECT ID, UserName AS ParentUserName FROM user) AS parent
ON parent.ID = user.parent
ORDER BY user.ID;