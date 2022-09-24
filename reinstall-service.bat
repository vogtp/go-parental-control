@echo off

go build -o .\go-win-session.exe .\cmd\session\ 

if %errorlevel% neq 0 exit /b %errorlevel%

net stop go-win-session
net start go-win-session

