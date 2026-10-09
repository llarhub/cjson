// Minimal size_t declaration used while generating cJSON bindings.
// __SIZE_TYPE__ is provided by Clang for the current target architecture.
#ifndef SETUP_PKGS_CJSON_STDDEF_H
#define SETUP_PKGS_CJSON_STDDEF_H
typedef __SIZE_TYPE__ size_t;
#endif
