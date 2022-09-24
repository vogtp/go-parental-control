

rem go build -ldflags -H=windowsgui -o .\go-win-idle.exe .\cmd\idle\ 
go build -ldflags -H=windowsgui -o go-win-systray.exe .\cmd\systray\

if %errorlevel% neq 0 exit /b %errorlevel%

rem copy .\go-win-idle.exe "C:\ProgramData\go-parental-controls\"
copy .\go-win-systray.exe "C:\ProgramData\go-parental-controls\"
copy .\go-win-systray.exe "C:\"