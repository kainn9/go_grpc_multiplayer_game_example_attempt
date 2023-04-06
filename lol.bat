cd bots
set count=0
:launch
set /a count+=1
if %count%==40 goto end
start go run .
ping -n 2 127.0.0.1 > nul
goto launch
:end


@REM cd bots
@REM set count=0
@REM :launch
@REM set /a count+=1
@REM if %count%==10 goto end
@REM start "" /b cmd /c "go run . || goto error"
@REM ping -n 2 127.0.0.1 > nul
@REM goto launch
@REM :error
@REM echo Error launching bot %count%. Retrying in 5 seconds...
@REM ping -n 5 127.0.0.1 > nul
@REM goto launch
@REM :end