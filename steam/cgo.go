package steam

// #cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/../sdk/public/steam
// #cgo LDFLAGS: -Wl,-rpath,:. -L${SRCDIR}/.. -Wl,-Bdynamic -lsteam_api -static-libgcc -static-libstdc++
import "C"
