FROM golang:1.23-bookworm

# --- Wails 用システム依存（Linux ビルド + Windows クロスビルド） ---
RUN apt-get update && apt-get install -y \
    ca-certificates curl gnupg \
    libgtk-3-dev libwebkit2gtk-4.0-dev libwebkit2gtk-4.1-dev \
    build-essential pkg-config \
    gcc-mingw-w64-x86-64 nsis \
    && rm -rf /var/lib/apt/lists/*

# --- Node.js 22 LTS（Svelte 5 / Vite 7 は Node 20.19+ / 22.12+ を要求） ---
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - \
    && apt-get install -y nodejs \
    && rm -rf /var/lib/apt/lists/* \
    && node -v && npm -v

# --- Wails CLI ---
RUN go install github.com/wailsapp/wails/v2/cmd/wails@v2.12.0

ENV PATH=/root/go/bin:$PATH

# Go の実行は yaegi（純 Go 製インタプリタ）を本体に static リンクして行うため、
# 追加の Go ツールチェーンやトランスパイラは不要（go.mod の依存として取得される）。

WORKDIR /app
CMD ["bash"]
