//go:build linux && arm && !arm7

package sherpa_ncnn

// #cgo LDFLAGS: -L ${SRCDIR}/lib/arm-unknown-linux-gnueabihf -lsherpa-ncnn-c-api -lsherpa-ncnn-core -lkaldi-native-fbank-core -lncnn -Wl,-rpath,${SRCDIR}/lib/arm-unknown-linux-gnueabihf
import "C"
