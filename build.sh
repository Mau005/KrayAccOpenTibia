#!/usr/bin/env bash
set -euo pipefail

# Defaults
TARGET=""
VERSION="1.0.0"
FORCE="false"

# Parse args
while [[ $# -gt 0 ]]; do
  case "$1" in
    -t|--target) TARGET="${2:-}"; shift 2 ;;
    -v|--version) VERSION="${2:-}"; shift 2 ;;
    -y|--yes) FORCE="true"; shift ;;
    -h|--help)
      cat <<'EOF'
Usage:
  ./build.sh -t <target> [-v 1.2.3] [-y]

Targets:
  windows-amd64  -> Console EXE (optional icon via app.syso if windres is available)
  linux-amd64    -> Console binary + config.yml + www + .desktop (+ icon.png from icon.png or icon.iconset)
  linux-arm64    -> Same for ARM64
  darwin-amd64   -> Console binary + optional .app (uses icon.icns or auto-generates from icon.iconset on macOS)
  darwin-arm64   -> Same for Apple Silicon (M1/M2)
  all            -> Build everything (cleans dist/ root)

Examples:
  ./build.sh -t linux-amd64
  ./build.sh -t darwin-arm64 -v 1.0.2
  ./build.sh -t all -y
EOF
      exit 0 ;;
    *)
      echo "Unknown arg: $1" >&2; exit 1 ;;
  case esac
done

# Paths (relative to this script)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT="$SCRIPT_DIR"
APP="KrayACC"
PKG_DIR="$ROOT/cmd/server"
PKG="./cmd/server"
LDFLAGS="-s -w -X main.Version=${VERSION}"
DIST="$ROOT/dist"
ICON_PNG_ROOT="$ROOT/icon.png"
ICON_ICNS="$ROOT/icon.icns"
ICONSET_DIR="$ROOT/icon.iconset"
SYSO="$PKG_DIR/app.syso"

say() { printf '%s\n' "$*"; }

ensure_clean_dir() {
  local path="$1"
  if [[ -e "$path" ]]; then
    if [[ "$FORCE" != "true" ]]; then
      echo "[NOTICE] The folder '$path' and its contents will be deleted."
      read -r -p "Proceed? (y/N) " ans
      [[ "$ans" == "y" || "$ans" == "Y" || "$ans" == "s" || "$ans" == "S" ]] || { echo "Cancelled."; exit 1; }
    fi
    # Try remove; if fails (in use), move aside then remove
    if ! rm -rf -- "$path" 2>/dev/null; then
      echo "Warn: cannot delete '$path' (maybe in use). Renaming and retrying..."
      local backup="${path}_old"
      mv -f -- "$path" "$backup" || { echo "Failed to rename '$path'."; exit 1; }
      sleep 1
      rm -rf -- "$backup" || true
    fi
  fi
  mkdir -p -- "$path"
}

best_icon_png_path() {
  # Echo best PNG path for Linux .desktop
  if [[ -f "$ICON_PNG_ROOT" ]]; then
    echo "$ICON_PNG_ROOT"; return
  fi
  if [[ -d "$ICONSET_DIR" ]]; then
    local candidates=( \
      "icon_512x512@2x.png" "icon_512x512.png" \
      "icon_256x256@2x.png" "icon_256x256.png" \
      "icon_128x128@2x.png" "icon_128x128.png" \
      "icon_32x32@2x.png"   "icon_32x32.png" \
      "icon_16x16@2x.png"   "icon_16x16.png" )
    for f in "${candidates[@]}"; do
      [[ -f "$ICONSET_DIR/$f" ]] && { echo "$ICONSET_DIR/$f"; return; }
    done
  fi
  echo "" # none
}

ensure_icon_icns() {
  # Prints path to icon.icns if available/generated; else empty
  if [[ -f "$ICON_ICNS" ]]; then
    echo "$ICON_ICNS"; return
  fi
  if [[ -d "$ICONSET_DIR" ]] && command -v iconutil >/dev/null 2>&1; then
    say "Generating icon.icns from icon.iconset (iconutil)"
    ( cd "$ROOT" && iconutil -c icns "icon.iconset" -o "icon.icns" ) || true
    [[ -f "$ICON_ICNS" ]] && { echo "$ICON_ICNS"; return; }
  fi
  echo ""
}

copy_assets() {
  local dest="$1"
  cp -a "$ROOT/config.yml" "$dest/"
  cp -a "$ROOT/www"        "$dest/"

  # PNG for Linux / desktop
  local png; png="$(best_icon_png_path || true)"
  if [[ -n "$png" && -f "$png" ]]; then
    cp -a "$png" "$dest/icon.png"
  fi

  # For reference include icon.icns if present
  [[ -f "$ICON_ICNS" ]] && cp -a "$ICON_ICNS" "$dest/"

  # Optionally include icon.iconset for Mac users
  [[ -d "$ICONSET_DIR" ]] && cp -a "$ICONSET_DIR" "$dest/"
}

build_windows() {
  local outdir="$DIST/${APP}_windows_amd64"
  ensure_clean_dir "$outdir"

  # Optional: generate app.syso if cross windres is available
  if command -v x86_64-w64-mingw32-windres >/dev/null 2>&1; then
    ( cd "$PKG_DIR" && x86_64-w64-mingw32-windres app.rc -O coff -o app.syso ) || true
  else
    say "[!] windres for Windows not found; building without .rc resources."
  fi

  GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
    go build -trimpath -ldflags="$LDFLAGS" -o "$outdir/${APP}.exe" "$PKG"

  copy_assets "$outdir"
  # Avoid leaking .syso to other targets
  [[ -f "$SYSO" ]] && rm -f -- "$SYSO" || true
}

build_linux() {
  local arch="$1"
  local outdir="$DIST/${APP}_linux_${arch}"
  ensure_clean_dir "$outdir"
  [[ -f "$SYSO" ]] && rm -f -- "$SYSO" || true

  GOOS=linux GOARCH="$arch" CGO_ENABLED=0 \
    go build -trimpath -ldflags="$LDFLAGS" -o "$outdir/$APP" "$PKG"

  copy_assets "$outdir"

  # .desktop
  local binPath iconPath
  binPath="$(python3 - <<PY
import os,sys
p=os.path.join("$outdir","$APP"); print(p.replace('\\\\','/'))
PY
)"
  if [[ -f "$outdir/icon.png" ]]; then
    iconPath="$outdir/icon.png"
  elif [[ -f "$outdir/icon.iconset/icon_512x512.png" ]]; then
    iconPath="$outdir/icon.iconset/icon_512x512.png"
  else
    iconPath="$outdir/icon.png"
  fi
  iconPath="$(python3 - <<PY
import os
print(r"$iconPath".replace('\\\\','/'))
PY
)"
  cat > "$outdir/${APP}.desktop" <<EOF
[Desktop Entry]
Name=$APP
Comment=$APP console client
Exec="$binPath"
Icon="$iconPath"
Terminal=true
Type=Application
Categories=Game;
EOF
}

build_mac() {
  local arch="$1"
  local outdir="$DIST/${APP}_darwin_${arch}"
  ensure_clean_dir "$outdir"
  [[ -f "$SYSO" ]] && rm -f -- "$SYSO" || true

  GOOS=darwin GOARCH="$arch" CGO_ENABLED=0 \
    go build -trimpath -ldflags="$LDFLAGS" -o "$outdir/$APP" "$PKG"

  copy_assets "$outdir"

  # Try to ensure icon.icns (auto-generate if on macOS with iconutil)
  local icnsPath
  icnsPath="$(ensure_icon_icns)"

  if [[ -n "$icnsPath" && -f "$icnsPath" ]]; then
    local appRoot="$DIST/${APP}_${arch}.app"
    local appDir="$appRoot/Contents"
    ensure_clean_dir "$appRoot"
    mkdir -p "$appDir/MacOS" "$appDir/Resources"
    cp -a "$outdir/$APP"          "$appDir/MacOS/$APP"
    cp -a "$icnsPath"             "$appDir/Resources/icon.icns"
    cat > "$appDir/Info.plist" <<'PLIST'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
 "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>CFBundleExecutable</key><string>KrayACC</string>
  <key>CFBundleIconFile</key><string>icon.icns</string>
  <key>CFBundleIdentifier</key><string>com.ainhosoft.krayacc</string>
  <key>CFBundleName</key><string>KrayACC</string>
  <key>CFBundlePackageType</key><string>APPL</string>
</dict>
</plist>
PLIST
    say "macOS .app bundle created with icon."
  else
    say "[!] icon.icns not found and could not be generated here."
    if [[ -d "$ICONSET_DIR" ]]; then
      say "    -> On macOS, run:  iconutil -c icns icon.iconset -o icon.icns"
      say "    -> Re-run:        ./build.sh -t darwin-$arch"
    else
      say "    -> Provide icon.icns or icon.iconset in repo root to bundle an app icon."
    fi
  fi
}

build_all() {
  ensure_clean_dir "$DIST"
  build_windows
  build_linux amd64
  build_linux arm64
  build_mac amd64
  build_mac arm64
  say "Done. Check dist/."
}

# Router
case "${TARGET,,}" in
  "")        echo "Use -t|--target (run with -h for help)"; exit 1 ;;
  windows-amd64) build_windows; say "Done. Check dist/." ;;
  linux-amd64)   build_linux amd64; say "Done. Check dist/." ;;
  linux-arm64)   build_linux arm64; say "Done. Check dist/." ;;
  darwin-amd64)  build_mac amd64;   say "Done. Check dist/." ;;
  darwin-arm64)  build_mac arm64;   say "Done. Check dist/." ;;
  all)           build_all ;;
  *)             echo "Unknown target: ${TARGET}"; exit 1 ;;
esac
