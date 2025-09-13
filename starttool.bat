@echo off
rem 启动第一个应用
cd C:\work\GoWord\DikTok_2
start go run main.go

rem 启动第二个应用
cd C:\work\GoWord\DikTok_2\rpc\contact
start go run contact.go

rem 启动第三个应用
cd C:\work\GoWord\DikTok_2\rpc\user
start go run user.go

rem 启动第四个应用
cd C:\work\GoWord\DikTok_2\rpc\video
start go run video.go

exit