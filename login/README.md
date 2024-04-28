สอบค่าวิชาชีพ ประจำปี 2024

Exam :	Dev034

ก่อนเริ่มให้ทำการติดตั้งของให้เรียบร้อย
ELK มีอยู่ใน other ในรูปเเบบ Docker-compose up ที่ root ได้เลยครับ ข้างใน FIX env สำหรับ User ไว้เเล้ว เเต่เป็น version 6.6
MSSQL อันนี้ อยากลองเเบบ command image run ดู จะได้เเตกต่างจาก ELK 

docker run --name sqlserver -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=Sa1angXD" -e TZ=Asia/Bangkok -e "MSSQL_AGENT_ENABLED=true" -p 1433:1433 -d mcr.microsoft.com/mssql/server:2019-latest

Frontend : Vue 3 CLI
lib include:
    axios
    vue-router
Run by local
npm i 
npm run

Backend : Go lang
lib include:
	github.com/denisenkom/go-mssqldb //connect mssql database
	github.com/gorilla/mux // Rounter for HTTP request multiplexer

Run by docker-compose -p exam2023 up --build
env on docker-compose.yaml
