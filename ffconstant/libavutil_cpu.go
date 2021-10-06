package ffconstant

const AV_CPU_FLAG_FORCE = 0x80000000 /* force usage of selected flags (OR) */

/* lower 16 bits - CPU features */
const AV_CPU_FLAG_MMX = 0x0001          ///< standard MMX
const AV_CPU_FLAG_MMXEXT = 0x0002       ///< SSE integer functions or AMD MMX ext
const AV_CPU_FLAG_MMX2 = 0x0002         ///< SSE integer functions or AMD MMX ext
const AV_CPU_FLAG_3DNOW = 0x0004        ///< AMD 3DNOW
const AV_CPU_FLAG_SSE = 0x0008          ///< SSE functions
const AV_CPU_FLAG_SSE2 = 0x0010         ///< PIV SSE2 functions
const AV_CPU_FLAG_SSE2SLOW = 0x40000000 ///< SSE2 supported, but usually not faster
///< than regular MMX/SSE (e.g. Core1)
const AV_CPU_FLAG_3DNOWEXT = 0x0020     ///< AMD 3DNowExt
const AV_CPU_FLAG_SSE3 = 0x0040         ///< Prescott SSE3 functions
const AV_CPU_FLAG_SSE3SLOW = 0x20000000 ///< SSE3 supported, but usually not faster
///< than regular MMX/SSE (e.g. Core1)
const AV_CPU_FLAG_SSSE3 = 0x0080        ///< Conroe SSSE3 functions
const AV_CPU_FLAG_SSSE3SLOW = 0x4000000 ///< SSSE3 supported, but usually not faster
const AV_CPU_FLAG_ATOM = 0x10000000     ///< Atom processor, some SSSE3 instructions are slower
const AV_CPU_FLAG_SSE4 = 0x0100         ///< Penryn SSE4.1 functions
const AV_CPU_FLAG_SSE42 = 0x0200        ///< Nehalem SSE4.2 functions
const AV_CPU_FLAG_AESNI = 0x80000       ///< Advanced Encryption Standard functions
const AV_CPU_FLAG_AVX = 0x4000          ///< AVX functions: requires OS support even if YMM registers aren't used
const AV_CPU_FLAG_AVXSLOW = 0x8000000   ///< AVX supported, but slow when using YMM registers (e.g. Bulldozer)
const AV_CPU_FLAG_XOP = 0x0400          ///< Bulldozer XOP functions
const AV_CPU_FLAG_FMA4 = 0x0800         ///< Bulldozer FMA4 functions
const AV_CPU_FLAG_CMOV = 0x1000         ///< supports cmov instruction
const AV_CPU_FLAG_AVX2 = 0x8000         ///< AVX2 functions: requires OS support even if YMM registers aren't used
const AV_CPU_FLAG_FMA3 = 0x10000        ///< Haswell FMA3 functions
const AV_CPU_FLAG_BMI1 = 0x20000        ///< Bit Manipulation Instruction Set 1
const AV_CPU_FLAG_BMI2 = 0x40000        ///< Bit Manipulation Instruction Set 2
const AV_CPU_FLAG_AVX512 = 0x100000     ///< AVX-512 functions: requires OS support even if YMM/ZMM registers aren't used

const AV_CPU_FLAG_ALTIVEC = 0x0001 ///< standard
const AV_CPU_FLAG_VSX = 0x0002     ///< ISA 2.06
const AV_CPU_FLAG_POWER8 = 0x0004  ///< ISA 2.07

const AV_CPU_FLAG_ARMV5TE = (1 << 0)
const AV_CPU_FLAG_ARMV6 = (1 << 1)
const AV_CPU_FLAG_ARMV6T2 = (1 << 2)
const AV_CPU_FLAG_VFP = (1 << 3)
const AV_CPU_FLAG_VFPV3 = (1 << 4)
const AV_CPU_FLAG_NEON = (1 << 5)
const AV_CPU_FLAG_ARMV8 = (1 << 6)
const AV_CPU_FLAG_VFP_VM = (1 << 7) ///< VFPv2 vector mode, deprecated in ARMv7-A and unavailable in various CPUs implementations
const AV_CPU_FLAG_SETEND = (1 << 16)

const AV_CPU_FLAG_MMI = (1 << 0)
const AV_CPU_FLAG_MSA = (1 << 1)
