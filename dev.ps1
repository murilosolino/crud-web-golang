# Script para iniciar o servidor com hot reload usando Air
# Uso: .\dev.ps1

$airPath = "$env:USERPROFILE\go\bin\air.exe"

if (-not (Test-Path $airPath)) {
    Write-Host "Air não encontrado. Instalando..." -ForegroundColor Yellow
    go install github.com/air-verse/air@latest
    Write-Host "Air instalado com sucesso!" -ForegroundColor Green
}

Write-Host "Iniciando servidor com hot reload na porta 8080..." -ForegroundColor Cyan
Write-Host "Pressione Ctrl+C para parar o servidor" -ForegroundColor Gray
Write-Host ""

& $airPath
