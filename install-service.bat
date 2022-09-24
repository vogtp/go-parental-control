@echo off
rem run this script as admin

go build -o .\go-win-session.exe .\cmd\session\ 

if %errorlevel% neq 0 exit /b %errorlevel%

sc create go-win-session binpath= "%CD%\go-win-session.exe" start= auto DisplayName= "go-win-session"
sc description go-win-session "go-partental-control"
net start go-win-session
sc query go-win-session


:exit