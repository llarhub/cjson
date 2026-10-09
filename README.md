# cJSON

LLGo bindings for cJSON 1.7.19.

## Installation

Install [LLGo](https://github.com/xgo-dev/llgo#readme) first, then install the native cJSON library and `pkg-config` for your platform. The binding links to `libcjson` through `pkg-config`.

These bindings are generated from cJSON **1.7.19**, and the tests expect that exact library version. Package manager versions vary; if yours differs, use the source installation below to match the bindings.

### macOS (Homebrew)

```bash
brew install cjson pkg-config
```

### Linux

Ubuntu / Debian ([libcjson-dev](https://packages.debian.org/libcjson-dev)):

```bash
sudo apt update
sudo apt install libcjson-dev pkg-config
```

Fedora ([cjson-devel](https://packages.fedoraproject.org/pkgs/cjson/cjson-devel/)):

```bash
sudo dnf install cjson-devel pkgconf-pkg-config
```

Arch Linux ([cjson](https://archlinux.org/packages/extra/x86_64/cjson/)):

```bash
sudo pacman -S cjson pkgconf
```

### Windows

Set up the matching LLGo toolchain and `pkg-config` using the [LLGo Windows guide](https://github.com/xgo-dev/llgo/blob/main/WINDOWS.md), then install cJSON for the same architecture and ABI.

For MinGW, run the following in an MSYS2 **CLANG64** shell (amd64):

```bash
pacman -S mingw-w64-clang-x86_64-cjson mingw-w64-clang-x86_64-pkgconf
```

For arm64, use a **CLANGARM64** shell and replace `mingw-w64-clang-x86_64-` with `mingw-w64-clang-aarch64-`. See the [MSYS2 cJSON packages](https://packages.msys2.org/base/mingw-w64-cjson). Keep the matching `clang64/bin` or `clangarm64/bin` directory on `PATH` when building and running programs so the cJSON DLL can be found.

For MSVC, install cJSON with your configured vcpkg (run from its directory in PowerShell):

```powershell
.\vcpkg.exe install cjson:x64-windows pkgconf:x64-windows
```

Use `arm64-windows` instead of `x64-windows` for arm64. Configure `pkg-config` to use that triplet's metadata and add its `bin` directory to `PATH`, as described in the LLGo Windows guide. Keep MSVC and MinGW libraries in their respective toolchain environments.

If using WSL, follow the Linux instructions inside your WSL distribution and run Go and LLGo there.

### From source (Linux / macOS, cJSON 1.7.19)

With Git, CMake, a C compiler, a build tool such as Make, and `pkg-config` installed, build the [upstream v1.7.19 release](https://github.com/DaveGamble/cJSON/tree/v1.7.19#building):

```bash
git clone --depth 1 --branch v1.7.19 https://github.com/DaveGamble/cJSON.git cJSON-1.7.19
cmake -S cJSON-1.7.19 -B cJSON-1.7.19/build \
  -DCMAKE_POLICY_VERSION_MINIMUM=3.5 \
  -DCMAKE_BUILD_TYPE=Release \
  -DCMAKE_INSTALL_PREFIX="$HOME/.local" \
  -DCMAKE_INSTALL_LIBDIR=lib \
  -DBUILD_SHARED_LIBS=ON \
  -DENABLE_CJSON_TEST=OFF
cmake --build cJSON-1.7.19/build
cmake --install cJSON-1.7.19/build
export PKG_CONFIG_PATH="$HOME/.local/lib/pkgconfig${PKG_CONFIG_PATH:+:$PKG_CONFIG_PATH}"
```

On Linux, also make the shared library available at runtime:

```bash
export LD_LIBRARY_PATH="$HOME/.local/lib${LD_LIBRARY_PATH:+:$LD_LIBRARY_PATH}"
```

Keep these environment variables set in the terminal where you build and run your program, or add them to your shell profile.

### Install the Go binding

Verify that `pkg-config` finds the intended cJSON installation:

```bash
pkg-config --modversion libcjson
pkg-config --libs libcjson
```

Then, from your Go module directory:

```bash
go get github.com/llarhub/cjson
```

## Usage

Parse JSON and read a string field:

```go
package main

import (
	"github.com/goplus/lib/c"
	"github.com/llarhub/cjson"
)

func main() {
	root := cjson.Parse(c.AllocaCStr(`{"name":"LLGo"}`))
	if root == nil {
		panic("invalid JSON")
	}
	defer root.Delete()

	name := root.ObjectItemCaseSensitive(c.AllocaCStr("name"))
	if name != nil && name.IsString() != 0 {
		println(c.GoString(name.StringValue()))
	}
}
```

Save the example as `main.go` and run it with `llgo run .`.
