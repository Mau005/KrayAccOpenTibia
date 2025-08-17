#!/usr/bin/env bash
set -euo pipefail

# =========================
#  Vars y parsing de args
# =========================
TARGET=""
VERSION="1.0.0"
FORCE="false"

while [[ $# -gt 0 ]]; do
  case "$1" in
    -t|--target) TARGET="${2:-}"; shift 2 ;;
    -v|--version) VERSION="${2:-}"; shift 2 ;;
    -f|--force) FORCE="true"; shift ;;
    -h|--help) SHOW_HELP="1"; shift ;;
    *) echo "Arg desconocido: $1"; SHOW_HELP="1"; shift ;;
  esac
done

TARGET_NORM="$(echo -n "${TARGET}" | tr '[:upper:]' '[:lower:]')"
echo "Received target: [${TARGET_NORM}]"

# Raíz del repo = carpeta del script
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

APP="KrayACC"

# Paquetes Go (rutas repo-relative para go build)
PKG_DIR="${ROOT}/cmd/server"
PKG="./cmd/server"

PKG_MANIFEST_DIR="${ROOT}/cmd/manifest"
PKG_MANIFEST="./cmd/manifest"
TOOL_NAME="manifest"

LD="-s -w -X main.Version=${VERSION}"
SYSO="${PKG_DIR}/app.syso"  # Solo relevante p/Windows
DIST="${ROOT}/dist"
ICON_PNG_ROOT="${ROOT}/icon.png"
ICON_ICNS="${ROOT}/icon.icns"
ICON_SET_DIR="${ROOT}/icon.iconset"

say(){ echo -e "$*"; }

ensure_clean_dir(){
  local path="$1"
  if [[ -e "$path" ]]; then
    if [[ "${FORCE}" != "true" ]]; then
      say "[NOTICE] La carpeta '$path' y su contenido serán eliminados."
      read -r -p "¿Continuar? (Y/N) " ans
      case "$ans" in
        y|Y|s|S) ;;
        *) echo "Cancelado por el usuario."; exit 1 ;;
      esac
    fi
    # Intentar borrar, si falla, renombrar y reintentar
    if ! rm -rf -- "$path" 2>/dev/null; then
      say "No se pudo borrar '$path'. ¿En uso? Renombrando y reintentando..."
      local backup="${path}_old"
      mv -- "$path" "$backup"
      sleep 1
      rm -rf -- "$backup" || true
    fi
  fi
  mkdir -p -- "$path"
}

# Retorna la mejor PNG disponible (para Linux/Desktop)
get_icon_png_path(){
  if [[ -f "$ICON_PNG_ROOT" ]]; then
    echo "$ICON_PNG_ROOT"; return 0
  fi
  if [[ -d "$ICON_SET_DIR" ]]; then
    local candidates=(
      "icon_512x512@2x.png" "icon_512x512.png"
      "icon_256x256@2x.png" "icon_256x256.png"
      "icon_128x128@2x.png" "icon_128x128.png"
      "icon_32x32@2x.png"   "icon_32x32.png"
      "icon_16x16@2x.png"   "icon_16x16.png"
    )
    for f in "${candidates[@]}"; do
      local p="${ICON_SET_DIR}/${f}"
      if [[ -f "$p" ]]; then echo "$p"; return 0; fi
    done
  fi
  echo ""  # no icon
}

# Garantiza icon.icns si es posible (macOS con iconutil)
# Devuelve ruta a icon.icns o vacío si no se pudo
ensure_icon_icns(){
  if [[ -f "$ICON_ICNS" ]]; then
    echo "$ICON_ICNS"; return 0
  fi
  if [[ -d "$ICON_SET_DIR" ]] && command -v iconutil >/dev/null 2>&1; then
    say "Generando icon.icns desde icon.iconset (iconutil)"
    ( cd "$ROOT" && iconutil -c icns "icon.iconset" -o "icon.icns" ) || true
    if [[ -f "$ICON_ICNS" ]]; then
      echo "$ICON_ICNS"; return 0
    fi
  fi
  echo ""
}

copy_assets(){
  local dest="$1"
  [[ -f "${ROOT}/config.yml" ]] && cp -f -- "${ROOT}/config.yml" "$dest"
  [[ -d "${ROOT}/www" ]] && cp -R -- "${ROOT}/www" "$dest"

  # PNG para Linux/desktop
  local png; png="$(get_icon_png_path)"
  if [[ -n "$png" ]]; then
    cp -f -- "$png" "${dest}/icon.png"
  fi

  # Copiar icns si está en root (referencia)
  [[ -f "$ICON_ICNS" ]] && cp -f -- "$ICON_ICNS" "$dest"

  # Opcional: incluir icon.iconset
  if [[ -d "$ICON_SET_DIR" ]]; then
    cp -R -- "$ICON_SET_DIR" "${dest}/icon.iconset"
  fi
}

build_windows(){
  local outdir="${DIST}/${APP}_windows_amd64"
  ensure_clean_dir "$outdir"

  # Nota: generar app.syso en Linux/Mac requiere windres o herramientas equivalentes.
  # Si no existen, seguir sin icon/version resource (Windows igual compila).
  if command -v windres >/dev/null 2>&1; then
    ( cd "$PKG_DIR" && windres app.rc -O coff -o app.syso ) || true
  else
    say "[!] 'windres' no encontrado. Se compilará sin app.syso (icon/version)."
  fi

  # Server
  say "==> Building Windows/amd64 (console) - server"
  GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
    go build -trimpath -ldflags="${LD}" -o "${outdir}/${APP}.exe" "${PKG}"

  # Tool (manifest)
  say "==> Building Windows/amd64 (console) - tool '${TOOL_NAME}'"
  GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
    go build -trimpath -ldflags="${LD}" -o "${outdir}/${TOOL_NAME}.exe" "${PKG_MANIFEST}"

  copy_assets "$outdir"

  # Limpiar syso para no contaminar otros targets
  [[ -f "$SYSO" ]] && rm -f -- "$SYSO" || true
}

build_linux(){
  local arch="$1"
  local outdir="${DIST}/${APP}_linux_${arch}"
  ensure_clean_dir "$outdir"

  # syso solo para Windows; evitar contaminar
  [[ -f "$SYSO" ]] && rm -f -- "$SYSO" || true

  # Server
  say "==> Building Linux/${arch} (console) - server"
  GOOS=linux GOARCH="${arch}" CGO_ENABLED=0 \
    go build -trimpath -ldflags="${LD}" -o "${outdir}/${APP}" "${PKG}"

  # Tool (manifest)
  say "==> Building Linux/${arch} (console) - tool '${TOOL_NAME}'"
  GOOS=linux GOARCH="${arch}" CGO_ENABLED=0 \
    go build -trimpath -ldflags="${LD}" -o "${outdir}/${TOOL_NAME}" "${PKG_MANIFEST}"

  copy_assets "$outdir"

  # .desktop
  local bin_path icon_guess icon_path
  bin_path="${outdir}/${APP}"
  icon_guess="${outdir}/icon.png"
  if [[ ! -f "$icon_guess" ]]; then
    icon_guess="${outdir}/icon.iconset/icon_512x512.png"
  fi
  icon_path="$icon_guess"

  cat > "${outdir}/${APP}.desktop" <<EOF
[Desktop Entry]
Name=${APP}
Comment=${APP} console client
Exec="${bin_path}"
Icon="${icon_path}"
Terminal=true
Type=Application
Categories=Game;
EOF
}

build_mac(){
  local arch="$1"
  local outdir="${DIST}/${APP}_darwin_${arch}"
  ensure_clean_dir "$outdir"

  # syso solo para Windows
  [[ -f "$SYSO" ]] && rm -f -- "$SYSO" || true

  # Server
  say "==> Building macOS/${arch} (console) - server"
  GOOS=darwin GOARCH="${arch}" CGO_ENABLED=0 \
    go build -trimpath -ldflags="${LD}" -o "${outdir}/${APP}" "${PKG}"

  # Tool (manifest)
  say "==> Building macOS/${arch} (console) - tool '${TOOL_NAME}'"
  GOOS=darwin GOARCH="${arch}" CGO_ENABLED=0 \
    go build -trimpath -ldflags="${LD}" -o "${outdir}/${TOOL_NAME}" "${PKG_MANIFEST}"

  copy_assets "$outdir"

  # Intentar preparar .app si existe o se puede generar icon.icns
  local icns_path; icns_path="$(ensure_icon_icns)"
  if [[ -n "$icns_path" ]]; then
    local app_root="${DIST}/${APP}_${arch}.app"
    local app_dir="${app_root}/Contents"
    ensure_clean_dir "$app_root"
    mkdir -p "${app_dir}/MacOS" "${app_dir}/Resources"
    cp -f -- "${outdir}/${APP}" "${app_dir}/MacOS/${APP}"
    cp -f -- "${icns_path}" "${app_dir}/Resources/icon.icns"
    cat > "${app_dir}/Info.plist" <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
 "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>CFBundleExecutable</key><string>${APP}</string>
  <key>CFBundleIconFile</key><string>icon.icns</string>
  <key>CFBundleIdentifier</key><string>com.ainhosoft.krayacc</string>
  <key>CFBundleName</key><string>${APP}</string>
  <key>CFBundlePackageType</key><string>APPL</string>
</dict>
</plist>
EOF
    say "macOS .app bundle creado con icono."
  else
    say "[!] icon.icns no encontrado / no generado."
    if [[ -d "$ICON_SET_DIR" ]]; then
      say "    -> En macOS, ejecuta:  iconutil -c icns icon.iconset -o icon.icns"
      say "    -> Re-ejecuta:         ./build.sh -t darwin-${arch}"
    else
      say "    -> Provee icon.icns o icon.iconset en la raíz del repo para incluir un icono."
    fi
  fi
}

build_all(){
  ensure_clean_dir "$DIST"
  build_windows
  build_linux "amd64"
  build_linux "arm64"
  build_mac "amd64"
  build_mac "arm64"
  say "Done. Revisa dist/."
}

show_help(){
cat <<'EOF'
Usage:
  ./build.sh -t <target> [-v 1.2.3] [-f]

Targets:
  windows-amd64  -> EXE consola (sin app.syso si no hay 'windres' disponible)
  linux-amd64    -> Binario + config.yml + www + .desktop (+ icon.png de icon.png o icon.iconset)
  linux-arm64    -> Igual para ARM64
  darwin-amd64   -> Binario + opcional .app (usa icon.icns o genera desde icon.iconset en macOS)
  darwin-arm64   -> Igual para Apple Silicon (M1/M2)
  all            -> Compila todo (limpia dist/ root)

Salida por target (dist/<target>/):
  - 'KrayACC' / 'KrayACC.exe'  (server)
  - 'manifest' / 'manifest.exe' (tool de cmd/manifest)
  - assets: config.yml, www/, iconos y .desktop (Linux)

Ejemplos:
  ./build.sh -t windows-amd64 -v 1.0.1
  ./build.sh -t linux-amd64
  ./build.sh -t all -f
EOF
}

if [[ "${SHOW_HELP:-}" == "1" || -z "${TARGET_NORM}" ]]; then
  show_help
  exit 0
fi

case "$TARGET_NORM" in
  windows-amd64) build_windows; say "Done. Revisa dist/." ;;
  linux-amd64)   build_linux "amd64"; say "Done. Revisa dist/." ;;
  linux-arm64)   build_linux "arm64"; say "Done. Revisa dist/." ;;
  darwin-amd64)  build_mac "amd64";   say "Done. Revisa dist/." ;;
  darwin-arm64)  build_mac "arm64";   say "Done. Revisa dist/." ;;
  all)           build_all ;;
  *) show_help; echo "Target desconocido: ${TARGET}"; exit 1 ;;
esac
