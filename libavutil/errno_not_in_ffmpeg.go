package libavutil

/**
 * This file has no copyright assigned and is placed in the Public Domain.
 * This file is part of the mingw-w64 runtime package.
 * No warranty is given; refer to the file DISCLAIMER.PD within this package.
 */
// #ifndef _INC_ERRNO
// const _INC_ERRNO

// #include <crtdefs.h>

// #ifdef __cplusplus
// extern "C"
// {
// #endif

// #ifndef _CRT_ERRNO_DEFINED
// const _CRT_ERRNO_DEFINED
//     _CRTIMP extern int *__cdecl _errno(void);
// const errno (*_errno())

//     errno_t __cdecl _set_errno(int _Value);
//     errno_t __cdecl _get_errno(int *_Value);
// #endif /* _CRT_ERRNO_DEFINED */

const EPERM = 1
const ENOENT = 2
const ENOFILE = ENOENT
const ESRCH = 3
const EINTR = 4
const EIO = 5
const ENXIO = 6
const E2BIG = 7
const ENOEXEC = 8
const EBADF = 9
const ECHILD = 10
const EAGAIN = 11
const ENOMEM = 12
const EACCES = 13
const EFAULT = 14
const EBUSY = 16
const EEXIST = 17
const EXDEV = 18
const ENODEV = 19
const ENOTDIR = 20
const EISDIR = 21
const ENFILE = 23
const EMFILE = 24
const ENOTTY = 25
const EFBIG = 27
const ENOSPC = 28
const ESPIPE = 29
const EROFS = 30
const EMLINK = 31
const EPIPE = 32
const EDOM = 33
const EDEADLK = 36
const ENAMETOOLONG = 38
const ENOLCK = 39
const ENOSYS = 40
const ENOTEMPTY = 41

// #ifndef RC_INVOKED
// #if !defined(_SECURECRT_ERRCODE_VALUES_DEFINED)
// const _SECURECRT_ERRCODE_VALUES_DEFINED
const EINVAL = 22
const ERANGE = 34
const EILSEQ = 42
const STRUNCATE = 80

// #endif
// #endif

const EDEADLOCK = EDEADLK

/* Posix thread extensions.  */

// #ifndef ENOTSUP
const ENOTSUP = 129

// #endif

/* Extension defined as by report VC 10+ defines error-numbers.  */

// #ifndef EAFNOSUPPORT
const EAFNOSUPPORT = 102

// #endif

// #ifndef EADDRINUSE
const EADDRINUSE = 100

// #endif

// #ifndef EADDRNOTAVAIL
const EADDRNOTAVAIL = 101

// #endif

// #ifndef EISCONN
const EISCONN = 113

// #endif

// #ifndef ENOBUFS
const ENOBUFS = 119

// #endif

// #ifndef ECONNABORTED
const ECONNABORTED = 106

// #endif

// #ifndef EALREADY
const EALREADY = 103

// #endif

// #ifndef ECONNREFUSED
const ECONNREFUSED = 107

// #endif

// #ifndef ECONNRESET
const ECONNRESET = 108

// #endif

// #ifndef EDESTADDRREQ
const EDESTADDRREQ = 109

// #endif

// #ifndef EHOSTUNREACH
const EHOSTUNREACH = 110

// #endif

// #ifndef EMSGSIZE
const EMSGSIZE = 115

// #endif

// #ifndef ENETDOWN
const ENETDOWN = 116

// #endif

// #ifndef ENETRESET
const ENETRESET = 117

// #endif

// #ifndef ENETUNREACH
const ENETUNREACH = 118

// #endif

// #ifndef ENOPROTOOPT
const ENOPROTOOPT = 123

// #endif

// #ifndef ENOTSOCK
const ENOTSOCK = 128

// #endif

// #ifndef ENOTCONN
const ENOTCONN = 126

// #endif

// #ifndef ECANCELED
const ECANCELED = 105

// #endif

// #ifndef EINPROGRESS
const EINPROGRESS = 112

// #endif

// #ifndef EOPNOTSUPP
const EOPNOTSUPP = 130

// #endif

// #ifndef EWOULDBLOCK
const EWOULDBLOCK = 140

// #endif

// #ifndef EOWNERDEAD
const EOWNERDEAD = 133

// #endif

// #ifndef EPROTO
const EPROTO = 134

// #endif

// #ifndef EPROTONOSUPPORT
const EPROTONOSUPPORT = 135

// #endif

// #ifndef EBADMSG
const EBADMSG = 104

// #endif

// #ifndef EIDRM
const EIDRM = 111

// #endif

// #ifndef ENODATA
const ENODATA = 120

// #endif

// #ifndef ENOLINK
const ENOLINK = 121

// #endif

// #ifndef ENOMSG
const ENOMSG = 122

// #endif

// #ifndef ENOSR
const ENOSR = 124

// #endif

// #ifndef ENOSTR
const ENOSTR = 125

// #endif

// #ifndef ENOTRECOVERABLE
const ENOTRECOVERABLE = 127

// #endif

// #ifndef ETIME
const ETIME = 137

// #endif

// #ifndef ETXTBSY
const ETXTBSY = 139

// #endif

/* Defined as WSAETIMEDOUT.  */
// #ifndef ETIMEDOUT
const ETIMEDOUT = 138

// #endif

// #ifndef ELOOP
const ELOOP = 114

// #endif

// #ifndef EPROTOTYPE
const EPROTOTYPE = 136

// #endif

// #ifndef EOVERFLOW
const EOVERFLOW = 132

// #endif

// #ifdef __cplusplus
// }
// #endif
// #endif
