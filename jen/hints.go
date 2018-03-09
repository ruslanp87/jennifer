package jen

/*
	This file is generated by genjen - do not edit!
*/

// Hints is a map containing hints for the names of all standard library packages
var Hints = map[string]string{
	"archive/tar":                       "tar",
	"archive/zip":                       "zip",
	"bufio":                             "bufio",
	"builtin":                           "builtin",
	"bytes":                             "bytes",
	"cmd/asm/internal/arch":             "arch",
	"cmd/asm/internal/asm":              "asm",
	"cmd/asm/internal/flags":            "flags",
	"cmd/asm/internal/lex":              "lex",
	"cmd/compile/internal/amd64":        "amd64",
	"cmd/compile/internal/arm":          "arm",
	"cmd/compile/internal/arm64":        "arm64",
	"cmd/compile/internal/gc":           "gc",
	"cmd/compile/internal/mips":         "mips",
	"cmd/compile/internal/mips64":       "mips64",
	"cmd/compile/internal/ppc64":        "ppc64",
	"cmd/compile/internal/s390x":        "s390x",
	"cmd/compile/internal/ssa":          "ssa",
	"cmd/compile/internal/syntax":       "syntax",
	"cmd/compile/internal/test":         "test",
	"cmd/compile/internal/types":        "types",
	"cmd/compile/internal/x86":          "x86",
	"cmd/go/internal/base":              "base",
	"cmd/go/internal/bug":               "bug",
	"cmd/go/internal/cache":             "cache",
	"cmd/go/internal/cfg":               "cfg",
	"cmd/go/internal/clean":             "clean",
	"cmd/go/internal/cmdflag":           "cmdflag",
	"cmd/go/internal/doc":               "doc",
	"cmd/go/internal/envcmd":            "envcmd",
	"cmd/go/internal/fix":               "fix",
	"cmd/go/internal/fmtcmd":            "fmtcmd",
	"cmd/go/internal/generate":          "generate",
	"cmd/go/internal/get":               "get",
	"cmd/go/internal/help":              "help",
	"cmd/go/internal/list":              "list",
	"cmd/go/internal/load":              "load",
	"cmd/go/internal/run":               "run",
	"cmd/go/internal/str":               "str",
	"cmd/go/internal/test":              "test",
	"cmd/go/internal/tool":              "tool",
	"cmd/go/internal/version":           "version",
	"cmd/go/internal/vet":               "vet",
	"cmd/go/internal/web":               "web",
	"cmd/go/internal/work":              "work",
	"cmd/internal/bio":                  "bio",
	"cmd/internal/browser":              "browser",
	"cmd/internal/buildid":              "buildid",
	"cmd/internal/dwarf":                "dwarf",
	"cmd/internal/edit":                 "edit",
	"cmd/internal/gcprog":               "gcprog",
	"cmd/internal/goobj":                "goobj",
	"cmd/internal/obj":                  "obj",
	"cmd/internal/obj/arm":              "arm",
	"cmd/internal/obj/arm64":            "arm64",
	"cmd/internal/obj/mips":             "mips",
	"cmd/internal/obj/ppc64":            "ppc64",
	"cmd/internal/obj/s390x":            "s390x",
	"cmd/internal/obj/x86":              "x86",
	"cmd/internal/objabi":               "objabi",
	"cmd/internal/objfile":              "objfile",
	"cmd/internal/src":                  "src",
	"cmd/internal/sys":                  "sys",
	"cmd/internal/test2json":            "test2json",
	"cmd/link/internal/amd64":           "amd64",
	"cmd/link/internal/arm":             "arm",
	"cmd/link/internal/arm64":           "arm64",
	"cmd/link/internal/ld":              "ld",
	"cmd/link/internal/loadelf":         "loadelf",
	"cmd/link/internal/loadmacho":       "loadmacho",
	"cmd/link/internal/loadpe":          "loadpe",
	"cmd/link/internal/mips":            "mips",
	"cmd/link/internal/mips64":          "mips64",
	"cmd/link/internal/objfile":         "objfile",
	"cmd/link/internal/ppc64":           "ppc64",
	"cmd/link/internal/s390x":           "s390x",
	"cmd/link/internal/sym":             "sym",
	"cmd/link/internal/x86":             "x86",
	"cmd/vet/internal/cfg":              "cfg",
	"cmd/vet/internal/whitelist":        "whitelist",
	"compress/bzip2":                    "bzip2",
	"compress/flate":                    "flate",
	"compress/gzip":                     "gzip",
	"compress/lzw":                      "lzw",
	"compress/zlib":                     "zlib",
	"container/heap":                    "heap",
	"container/list":                    "list",
	"container/ring":                    "ring",
	"context":                           "context",
	"crypto":                            "crypto",
	"crypto/aes":                        "aes",
	"crypto/cipher":                     "cipher",
	"crypto/des":                        "des",
	"crypto/dsa":                        "dsa",
	"crypto/ecdsa":                      "ecdsa",
	"crypto/elliptic":                   "elliptic",
	"crypto/hmac":                       "hmac",
	"crypto/internal/cipherhw":          "cipherhw",
	"crypto/md5":                        "md5",
	"crypto/rand":                       "rand",
	"crypto/rc4":                        "rc4",
	"crypto/rsa":                        "rsa",
	"crypto/sha1":                       "sha1",
	"crypto/sha256":                     "sha256",
	"crypto/sha512":                     "sha512",
	"crypto/subtle":                     "subtle",
	"crypto/tls":                        "tls",
	"crypto/x509":                       "x509",
	"crypto/x509/pkix":                  "pkix",
	"database/sql":                      "sql",
	"database/sql/driver":               "driver",
	"debug/dwarf":                       "dwarf",
	"debug/elf":                         "elf",
	"debug/gosym":                       "gosym",
	"debug/macho":                       "macho",
	"debug/pe":                          "pe",
	"debug/plan9obj":                    "plan9obj",
	"encoding":                          "encoding",
	"encoding/ascii85":                  "ascii85",
	"encoding/asn1":                     "asn1",
	"encoding/base32":                   "base32",
	"encoding/base64":                   "base64",
	"encoding/binary":                   "binary",
	"encoding/csv":                      "csv",
	"encoding/gob":                      "gob",
	"encoding/hex":                      "hex",
	"encoding/json":                     "json",
	"encoding/pem":                      "pem",
	"encoding/xml":                      "xml",
	"errors":                            "errors",
	"expvar":                            "expvar",
	"flag":                              "flag",
	"fmt":                               "fmt",
	"go/ast":                            "ast",
	"go/build":                          "build",
	"go/constant":                       "constant",
	"go/doc":                            "doc",
	"go/format":                         "format",
	"go/importer":                       "importer",
	"go/internal/gccgoimporter":         "gccgoimporter",
	"go/internal/gcimporter":            "gcimporter",
	"go/internal/srcimporter":           "srcimporter",
	"go/parser":                         "parser",
	"go/printer":                        "printer",
	"go/scanner":                        "scanner",
	"go/token":                          "token",
	"go/types":                          "types",
	"hash":                              "hash",
	"hash/adler32":                      "adler32",
	"hash/crc32":                        "crc32",
	"hash/crc64":                        "crc64",
	"hash/fnv":                          "fnv",
	"html":                              "html",
	"html/template":                     "template",
	"image":                             "image",
	"image/color":                       "color",
	"image/color/palette":               "palette",
	"image/draw":                        "draw",
	"image/gif":                         "gif",
	"image/internal/imageutil":          "imageutil",
	"image/jpeg":                        "jpeg",
	"image/png":                         "png",
	"index/suffixarray":                 "suffixarray",
	"internal/cpu":                      "cpu",
	"internal/nettrace":                 "nettrace",
	"internal/poll":                     "poll",
	"internal/race":                     "race",
	"internal/singleflight":             "singleflight",
	"internal/syscall/windows":          "windows",
	"internal/syscall/windows/registry": "registry",
	"internal/syscall/windows/sysdll":   "sysdll",
	"internal/testenv":                  "testenv",
	"internal/testlog":                  "testlog",
	"internal/trace":                    "trace",
	"io":                                "io",
	"io/ioutil":                         "ioutil",
	"log":                               "log",
	"log/syslog":                        "syslog",
	"math":                              "math",
	"math/big":                          "big",
	"math/bits":                         "bits",
	"math/cmplx":                        "cmplx",
	"math/rand":                         "rand",
	"mime":                              "mime",
	"mime/multipart":                    "multipart",
	"mime/quotedprintable":              "quotedprintable",
	"net":                            "net",
	"net/http":                       "http",
	"net/http/cgi":                   "cgi",
	"net/http/cookiejar":             "cookiejar",
	"net/http/fcgi":                  "fcgi",
	"net/http/httptest":              "httptest",
	"net/http/httptrace":             "httptrace",
	"net/http/httputil":              "httputil",
	"net/http/internal":              "internal",
	"net/http/pprof":                 "pprof",
	"net/internal/socktest":          "socktest",
	"net/mail":                       "mail",
	"net/rpc":                        "rpc",
	"net/rpc/jsonrpc":                "jsonrpc",
	"net/smtp":                       "smtp",
	"net/textproto":                  "textproto",
	"net/url":                        "url",
	"os":                             "os",
	"os/exec":                        "exec",
	"os/signal":                      "signal",
	"os/signal/internal/pty":         "pty",
	"os/user":                        "user",
	"path":                           "path",
	"path/filepath":                  "filepath",
	"plugin":                         "plugin",
	"reflect":                        "reflect",
	"regexp":                         "regexp",
	"regexp/syntax":                  "syntax",
	"runtime":                        "runtime",
	"runtime/cgo":                    "cgo",
	"runtime/debug":                  "debug",
	"runtime/internal/atomic":        "atomic",
	"runtime/internal/sys":           "sys",
	"runtime/pprof":                  "pprof",
	"runtime/pprof/internal/profile": "profile",
	"runtime/race":                   "race",
	"runtime/trace":                  "trace",
	"sort":                           "sort",
	"strconv":                        "strconv",
	"strings":                        "strings",
	"sync":                           "sync",
	"sync/atomic":                    "atomic",
	"syscall":                        "syscall",
	"testing":                        "testing",
	"testing/internal/testdeps":      "testdeps",
	"testing/iotest":                 "iotest",
	"testing/quick":                  "quick",
	"text/scanner":                   "scanner",
	"text/tabwriter":                 "tabwriter",
	"text/template":                  "template",
	"text/template/parse":            "parse",
	"time":                           "time",
	"unicode":                        "unicode",
	"unicode/utf16":                  "utf16",
	"unicode/utf8":                   "utf8",
	"unsafe":                         "unsafe",
}
