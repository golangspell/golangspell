REM Usage: golangspellupdate.bat
del /f /q /s %HOMEDRIVE%%HOMEPATH%\go\src\github.com\golangspell\golangspell\*.* > NUL
rmdir /q /s %HOMEDRIVE%%HOMEPATH%\go\src\github.com\golangspell\golangspell
del /f /q %HOMEDRIVE%%HOMEPATH%\go\bin\golangspell.exe 
go install "github.com/golangspell/golangspell@latest"