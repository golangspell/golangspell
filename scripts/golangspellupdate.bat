REM Usage: golangspellupdate.bat
del /f /q /s %HOMEDRIVE%%HOMEPATH%\go\src\github.com\danilovalente\golangspell\*.* > NUL
rmdir /q /s %HOMEDRIVE%%HOMEPATH%\go\src\github.com\danilovalente\golangspell
del /f /q %HOMEDRIVE%%HOMEPATH%\go\bin\golangspell.exe 
go get "github.com/danilovalente/golangspell"