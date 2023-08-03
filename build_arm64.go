//go:build linux && arm64

package sherpa_ncnn

// #cgo LDFLAGS: -L ${SRCDIR}/lib/aarch64-unknown-linux-gnu -lsherpa-ncnn-c-api -lsherpa-ncnn-core -lkaldi-native-fbank-core -lncnn -Wl,-rpath,${SRCDIR}/lib/aarch64-unknown-linux-gnu
import "C"
