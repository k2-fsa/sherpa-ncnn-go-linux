//go:build !android && linux && amd64 && !musl

package sherpa_ncnn

// #cgo LDFLAGS: -L ${SRCDIR}/lib/x86_64-unknown-linux-gnu -lsherpa-ncnn-c-api -lsherpa-ncnn-core -lkaldi-native-fbank-core -lncnn -Wl,-rpath,${SRCDIR}/lib/x86_64-unknown-linux-gnu
import "C"
