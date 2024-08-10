
# This script matches the linux/macOS version of bin/lint
# Please keep them in sync

Write-Output "Running linter..."

if (-not (Get-Command golangci-lint -ErrorAction SilentlyContinue)) {
    Write-Host -ForegroundColor Red "[FAIL] golangci-lint is not installed."
    Write-Host -ForegroundColor Yellow "       Please install golangci-lint by downloading it from : https://github.com/golangci/golangci-lint/releases"
    exit 1
}

$startTime = Get-Date

golangci-lint run
$exit_code = $LASTEXITCODE

$endTime = Get-Date
$elapsedTime = New-TimeSpan -Start $startTime -End $endTime
Write-Output " Linter finished in $elapsedTime."

if ($exit_code -eq 0) {
    Write-Host -BackgroundColor Green -ForegroundColor White "[PASS]  Linter has not found any issues."
}
else
{
    Write-Host -BackgroundColor Red -ForegroundColor White "[FAIL]  Linter has found issues."
    exit $exit_code
}
