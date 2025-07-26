# 运行环境
# go1.23.0 windows/amd64
# 依赖安装步骤
# go get -u gorm.io/gorm
# go get -u gorm.io/driver/sqlite
# go get -u gorm.io/driver/mysql
# go get -u github.com/gin-gonic/gin
# 启动方式
# go run main.go
# 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码
# 测试用例及结果
![alt text](610a89db06d3dd738d182d4c9b02505.png)
![alt text](f1994903e93e747b40539f6b339527e.png)
![alt text](94a19a80778deb037ffc1a454969fdb.png)
# 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。
![alt text](4407ae7fc727be51584c747711a676d.png)
![alt text](67e3b96769cd4ee7a12ed4d28b7533e.png)
![alt text](dbb1be1b61e4cdb9e066f7154a383e5.png)
# 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
![alt text](5cca8a73b6d6c93808a2ecc86d4e9ac-1.png)
![alt text](9bc54854ad8b13643742a8d2b8d2bf2.png)
![alt text](58dc8d0aa27ebf6f421b088d87b2d7b.png)
![alt text](8e68b86b95e0b4271742b6fddc0ea3b.png)
![alt text](bde6c383d31e13a5df47cb28230c310.png)
# 实现文章的读取功能，支持获取所有文章列表。
![alt text](999ef9b96a91a15419608a598f23b31.png)
![alt text](7b652f8093ec24f020541fab7fe1c09.png)
![alt text](ccc23de0f953557edd23aaeb816cc26.png)
# 实现文章的读取功能，支持获取单个文章的详细信息。
![alt text](e49ca7639214f05ec694d3efd375652.png)
![alt text](af01a05de3eae73ae10ab6aad7dab54.png)
![alt text](c1bb505a73277d07946b62c8167738d.png)
![alt text](ccc23de0f953557edd23aaeb816cc26.png)
# 实现文章的更新功能，只有文章的作者才能更新自己的文章。
![alt text](e6961e8cba9fda65e2d723169aa892a.png)
![alt text](7c1e36c4f74e1e49b4604a1343a8790.png)
![alt text](6ec87e29474c97aeea7aa369d0bc6a5.png)
![alt text](55e26bebb99811bae754ce729c8e879.png)
# 实现文章的删除功能，只有文章的作者才能删除自己的文章。
![alt text](c8145c8bba41735bafd447d57dac66d.png)
![alt text](5aabd788d230c3b7adff22561cef4a6.png)
![alt text](e99eefac66a9b1762c12ee805aa2e74.png)
![alt text](3291ea87642071e9997f399e4b71d99.png)
# 实现评论的创建功能，已认证的用户可以对文章发表评论。
![alt text](df018f203483eae6853d79ed08a56ff.png)
![alt text](6f04c5fb8b8fce7caba7a54a48ec9be.png)
![alt text](d3933d54a16c4996c088ae6bdadefb8.png)
![alt text](b42dc67f6a32ed28960003f76fe9129.png)
# 实现评论的读取功能，支持获取某篇文章的所有评论列表。
![alt text](2d9046daf4d58862ecb0db9a178ce1d.png)
![alt text](40081a222ebf5fec2f845c9d3d1ab2b.png)
![alt text](8d347518af29659e21145701fdad65d.png)
![alt text](9387c60c0df247bb2376cf9c5a6a59b.png)
