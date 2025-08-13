param(
  [string]$Target = "",
  [string]$Version = "1.0.0",
  [switch]$Force
)

$ErrorActionPreference = "Stop"

# Normalize Target (supports -Target value and -Target=value)
$Target = ($Target | Out-String).Trim()
$TargetNorm = $Target.ToLowerInvariant()
Write-Host "Received target: [$TargetNorm]"

# --- Base paths (always relative to this script) ---
$Root        = $PSScriptRoot
$App         = "KrayACC"
$PkgDir      = Join-Path $Root "cmd/server"
$Pkg         = "./cmd/server"   # go build uses repo-relative paths
$Ld          = "-s -w -X main.Version=$Version"
$Syso        = Join-Path $PkgDir "app.syso"
$Dist        = Join-Path $Root "dist"
$IconPngRoot = Join-Path $Root "icon.png"
$IconIcns    = Join-Path $Root "icon.icns"
$IconSetDir  = Join-Path $Root "icon.iconset"

function Say($msg){ Write-Host $msg }

function Ensure-CleanDir($path){
  if (Test-Path $path) {
    if (-not $Force) {
      Say "[NOTICE] The folder '$path' and its contents will be deleted."
      $ans = Read-Host "Proceed? (Y/N)"
      if ($ans -notin @("y","Y","s","S")) { throw "Cancelled by user." }
    }
    try {
      Remove-Item -Recurse -Force $path
    } catch {
      Write-Warning "Cannot delete '$path'. It might be in use. Renaming and retrying..."
      $backup = "${path}_old"
      try {
        Rename-Item $path $backup -ErrorAction Stop
        Start-Sleep -Seconds 1
        Remove-Item -Recurse -Force $backup -ErrorAction SilentlyContinue
      } catch {
        throw "Failed to remove or rename '$path'. Close any running binaries or open windows and try again."
      }
    }
  }
  New-Item -ItemType Directory $path | Out-Null
}

# Picks the best available PNG path for Linux/Desktop use.
function Get-IconPngPath {
  if (Test-Path $IconPngRoot) { return $IconPngRoot }
  if (Test-Path $IconSetDir) {
    # Prefer the largest common sizes
    $candidates = @(
      "icon_512x512@2x.png","icon_512x512.png",
      "icon_256x256@2x.png","icon_256x256.png",
      "icon_128x128@2x.png","icon_128x128.png",
      "icon_32x32@2x.png","icon_32x32.png",
      "icon_16x16@2x.png","icon_16x16.png"
    )
    foreach($f in $candidates){
      $p = Join-Path $IconSetDir $f
      if (Test-Path $p) { return $p }
    }
  }
  return $null
}

# Tries to ensure icon.icns exists:
# - If already present, OK.
# - Else if icon.iconset exists and iconutil is available (macOS), generate icon.icns in root.
# Returns: full path to icon.icns if available, otherwise $null.
function Ensure-IconIcns {
  if (Test-Path $IconIcns) { return $IconIcns }
  if ((Test-Path $IconSetDir) -and (Get-Command iconutil -ErrorAction SilentlyContinue)) {
    Say "Generating icon.icns from icon.iconset (iconutil)"
    Push-Location $Root
    try {
      & iconutil -c icns "icon.iconset" -o "icon.icns"
    } finally {
      Pop-Location
    }
    if (Test-Path $IconIcns) { return $IconIcns }
  }
  return $null
}

function Copy-Assets($dest){
  Copy-Item (Join-Path $Root "config.yml") $dest
  Copy-Item (Join-Path $Root "www") -Recurse $dest

  # Copy PNG for Linux/desktop use
  $png = Get-IconPngPath
  if ($png) {
    Copy-Item $png (Join-Path $dest "icon.png")
  }

  # If you keep icon.icns in root, also copy it for reference
  if (Test-Path $IconIcns) {
    Copy-Item $IconIcns $dest
  }

  # Optionally include the icon.iconset folder for Mac users to regen icns if needed
  if (Test-Path $IconSetDir) {
    Copy-Item $IconSetDir (Join-Path $dest "icon.iconset") -Recurse
  }
}

function Build-Windows {
  $outdir = Join-Path $Dist "${App}_windows_amd64"
  Ensure-CleanDir $outdir

  Push-Location $PkgDir
  if (Get-Command windres -ErrorAction SilentlyContinue) {
    & windres app.rc -O coff -o app.syso
  } elseif (Get-Command rc -ErrorAction SilentlyContinue) {
    & rc /nologo /fo app.res app.rc
    & cvtres /MACHINE:X64 /OUT:app.syso app.res
    if (Test-Path "app.res") { Remove-Item "app.res" }
  } else {
    Say "[!] windres/rc.exe not found. If app.syso already exists, it will be used."
  }
  Pop-Location

  $env:GOOS="windows"; $env:GOARCH="amd64"; $env:CGO_ENABLED="0"
  Say "==> Building Windows/amd64 (console)"
  go build -trimpath -ldflags="$Ld" -o (Join-Path $outdir "${App}.exe") $Pkg

  Copy-Assets $outdir

  if (Test-Path $Syso) { Remove-Item $Syso } # avoid polluting other targets
}

function Build-Linux($arch){
  $outdir = Join-Path $Dist "${App}_linux_${arch}"
  Ensure-CleanDir $outdir

  if (Test-Path $Syso) { Remove-Item $Syso } # syso is Windows-only

  $env:GOOS="linux"; $env:GOARCH=$arch; $env:CGO_ENABLED="0"
  Say "==> Building Linux/$arch (console)"
  go build -trimpath -ldflags="$Ld" -o (Join-Path $outdir $App) $Pkg

  Copy-Assets $outdir

  # Create .desktop (absolute paths and forward slashes)
  $binPath  = (Join-Path $outdir $App).Replace('\','/')
  $iconGuess = Join-Path $outdir "icon.png"
  if (-not (Test-Path $iconGuess)) {
    # fallback to an .iconset png if we copied the folder
    $iconGuess = Join-Path $outdir "icon.iconset/icon_512x512.png"
  }
  $iconPath = $iconGuess.Replace('\','/')

@"
[Desktop Entry]
Name=$App
Comment=$App console client
Exec="$binPath"
Icon="$iconPath"
Terminal=true
Type=Application
Categories=Game;
"@ | Set-Content (Join-Path $outdir "${App}.desktop") -Encoding UTF8
}

function Build-Mac($arch){
  $outdir = Join-Path $Dist "${App}_darwin_${arch}"
  Ensure-CleanDir $outdir

  if (Test-Path $Syso) { Remove-Item $Syso } # syso is Windows-only

  $env:GOOS="darwin"; $env:GOARCH=$arch; $env:CGO_ENABLED="0"
  Say "==> Building macOS/$arch (console)"
  go build -trimpath -ldflags="$Ld" -o (Join-Path $outdir $App) $Pkg

  Copy-Assets $outdir

  # Try to ensure icon.icns exists (generate from icon.iconset if possible)
  $icnsPath = Ensure-IconIcns

  if ($icnsPath) {
    $appRoot = Join-Path $Dist "${App}_${arch}.app"
    $appDir  = Join-Path $appRoot "Contents"
    Ensure-CleanDir $appRoot
    New-Item -ItemType Directory -Force (Join-Path $appDir "MacOS")     | Out-Null
    New-Item -ItemType Directory -Force (Join-Path $appDir "Resources") | Out-Null
    Copy-Item (Join-Path $outdir $App) (Join-Path $appDir "MacOS/${App}")
    Copy-Item $icnsPath (Join-Path $appDir "Resources/icon.icns")
@"
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
 "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>CFBundleExecutable</key><string>${App}</string>
  <key>CFBundleIconFile</key><string>icon.icns</string>
  <key>CFBundleIdentifier</key><string>com.ainhosoft.krayacc</string>
  <key>CFBundleName</key><string>${App}</string>
  <key>CFBundlePackageType</key><string>APPL</string>
</dict>
</plist>
"@ | Set-Content (Join-Path $appDir "Info.plist") -Encoding UTF8
    Say "macOS .app bundle created with icon."
  } else {
    Say "[!] icon.icns not found and could not be generated here."
    if (Test-Path $IconSetDir) {
      Say "    -> On macOS, run:  iconutil -c icns icon.iconset -o icon.icns"
      Say "    -> Re-run:        .\build.ps1 -Target darwin-$arch"
    } else {
      Say "    -> Provide icon.icns or icon.iconset in repo root to bundle an app icon."
    }
  }
}

function Build-All {
  Ensure-CleanDir $Dist
  Build-Windows
  Build-Linux "amd64"
  Build-Linux "arm64"
  Build-Mac   "amd64"
  Build-Mac   "arm64"
  Say "Done. Check dist/."
}

function Show-Help {
@"
Usage:
  .\build.ps1 -Target <target> [-Version 1.2.3] [-Force]

Targets:
  windows-amd64  -> Console EXE with icon/version (app.syso)
  linux-amd64    -> Console binary + config.yml + www + .desktop (+ icon.png from icon.png or icon.iconset)
  linux-arm64    -> Same for ARM64
  darwin-amd64   -> Console binary + optional .app (uses icon.icns or auto-generates from icon.iconset on macOS)
  darwin-arm64   -> Same for Apple Silicon (M1/M2)
  all            -> Build everything (cleans dist/ root)

Examples:
  .\build.ps1 -Target windows-amd64 -Version 1.0.1
  .\build.ps1 -Target linux-amd64
  .\build.ps1 -Target all -Force
"@ | Write-Host
}

switch ($TargetNorm) {
  ""              { Show-Help }
  "windows-amd64" { Build-Windows; Say "Done. Check dist/."; }
  "linux-amd64"   { Build-Linux "amd64"; Say "Done. Check dist/."; }
  "linux-arm64"   { Build-Linux "arm64"; Say "Done. Check dist/."; }
  "darwin-amd64"  { Build-Mac "amd64"; Say "Done. Check dist/."; }
  "darwin-arm64"  { Build-Mac "arm64"; Say "Done. Check dist/."; }
  "all"           { Build-All }
  default         { Show-Help; throw "Unknown target: $Target" }
}
