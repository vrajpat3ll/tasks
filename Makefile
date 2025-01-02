tasks:
	@go build -o "tasks.exe" "./cmd"
	@cli-add.bat tasks.exe