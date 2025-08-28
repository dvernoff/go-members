@echo off
setlocal enabledelayedexpansion

echo ------------------------------------------------------
echo        Discord Member Checker - Build
echo ------------------------------------------------------

set APP_NAME=go-members
set VERSION=1.0.0

REM Create builds directory
if not exist builds mkdir builds

echo.
echo [+] Building for Windows (amd64)...
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w" -o builds\%APP_NAME%.exe .
if %errorlevel% == 0 (
    echo [✓] Windows build completed successfully!
    echo     Output: builds\%APP_NAME%.exe
) else (
    echo [!] Windows build failed!
)

echo.
echo [+] Building for Linux (amd64)...
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w" -o builds\%APP_NAME%-linux .
if %errorlevel% == 0 (
    echo [✓] Linux build completed successfully!
    echo     Output: builds\%APP_NAME%-linux
) else (
    echo [!] Linux build failed!
)

REM Clear temporary env vars
set GOOS=
set GOARCH=

echo.
echo [+] Moving .env into builds folder (if exists)...
if exist ".env" (
    move /Y ".env" "builds\.env" >nul
    if %errorlevel% == 0 (
        echo [✓] .env moved to builds\.env
    ) else (
        echo [!] Failed to move .env
    )
) else (
    echo [i] No .env file found in project root.
)

echo.
echo ------------------------------------------------------
echo        Build finished
echo ------------------------------------------------------
echo.
echo Build artifacts:
dir /B builds\

echo.
echo Remember to set your bot token in builds\.env:
echo DISCORD_TOKEN=your_bot_token_here
echo.

pause
