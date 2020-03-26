REM Usage: golangspellupdatespell.bat "[module]" "[name]"
REM Example: golangspellupdatespell.bat "github.com/danilovalente/golangspell-core" "golangspell-core"
del /f /q /s %HOMEDRIVE%%HOMEPATH%\go\src\%1\*.* > NUL
rmdir /q /s %HOMEDRIVE%%HOMEPATH%\go\src\%1
del /f /q %HOMEDRIVE%%HOMEPATH%\go\bin\%2.exe 
SET GO111MODULE=off
golangspell updatespell %1 %2
SET GO111MODULE=on