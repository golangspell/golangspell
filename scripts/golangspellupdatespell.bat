REM Usage: golangspellupdatespell.bat "[module]" "[name]"
REM Example: golangspellupdatespell.bat "github.com/golangspell/golangspell-core" "golangspell-core"
del /f /q /s %HOMEDRIVE%%HOMEPATH%\go\src\%1\*.* > NUL
rmdir /q /s %HOMEDRIVE%%HOMEPATH%\go\src\%1
del /f /q %HOMEDRIVE%%HOMEPATH%\go\bin\%2.exe 
golangspell updatespell %1 %2