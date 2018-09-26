%module steam

// include modules
%include cmalloc.i

// macro to both include in wrapper and swig preprocessor
%define INCL(path)
%{#include path%}
%include path
%enddef

// define some required macros
#define POSIX

// delete some classes that don't translate properly
%rename("$ignore", regextarget=1, fullname=1) "CSteamID::.*";
%rename("$ignore", regextarget=1, fullname=1) "CGameID::.*";

// include the headers
INCL("../sdk/public/steam/steam_api.h")

// need to get the size of this types
%sizeof(CreateItemResult_t)
%sizeof(SubmitItemUpdateResult_t)
