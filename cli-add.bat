@echo off
setlocal
set filePath=%1
for %%F in ("%filePath%") do (
    set "filename=%%~nF"  
    set "extension=%%~xF"
)
move "%filePath%" "C:\commands\%filename%%extension%"

@REM alias %filename% "%filePath%"