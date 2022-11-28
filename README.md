usermessage表：
ID 主键，name 姓名,password 密码,secretProtection 密保
messageboard表:
senderID 发送者ID,sendername 发送者姓名,
receiveID 接收者ID,receivename 接收者姓名,message 留言,
commenterID评论者ID,commenterName 评论者姓名,comment 评论
数据库说明：
用户表没什么好说，留言板表可根据留言者和发送者确实唯一留言内容，根据留言者和发送者以及评论者可确定唯一评论内容
亮点：
1.根据用户名和密保找回密码
