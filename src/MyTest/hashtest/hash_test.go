package hashtest

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
	"io/ioutil"
	"testing"
	"time"
)

func Test_sm3(t *testing.T) {
	var hash = sm3.New()
	b, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\837.mp4")
	fmt.Println(len(b))
	start := time.Now()
	hash.Write(b)
	hash.Sum(nil)
	end := time.Now()
	total := end.Sub(start)
	ti := total.Nanoseconds()
	fmt.Println(ti)
}
func BenchmarkSm3Big(b *testing.B) {
	var hash = sm3.New()

	v, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\418.mp4")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.Write(v)
		hash.Sum(nil)
		hash.Reset()
	}
}
func BenchmarkSm3Small(b *testing.B) {
	var hash = sm3.New()

	v, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\837.mp4")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.Write(v)
		hash.Sum(nil)
		hash.Reset()
	}

}
func BenchmarkMd5Small(b *testing.B) {
	var hash = md5.New()
	v, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\837.mp4")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.Write(v)
		hash.Sum(nil)
		hash.Reset()
	}
}
func BenchmarkMd5Big(b *testing.B) {
	var hash = md5.New()
	v, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\837.mp4")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.Write(v)
		hash.Sum(nil)
		hash.Reset()
	}

}
func BenchmarkSHA1Big(b *testing.B) {
	var hash = sha1.New()
	v, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\837.mp44")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.Write(v)
		hash.Sum(nil)
		hash.Reset()
	}

}
func BenchmarkSHA1Small(b *testing.B) {
	var hash = sha1.New()
	v, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\837.mp4")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.Write(v)
		hash.Sum(nil)
		hash.Reset()
	}
}
func BenchmarkSHA256Big(b *testing.B) {
	var hash = sha256.New()
	v, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\837.mp4")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.Write(v)
		hash.Sum(nil)
		hash.Reset()
	}

}
func BenchmarkSHA256Small(b *testing.B) {
	var hash = sha256.New()
	v, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\837.mp4")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.Write(v)
		hash.Sum(nil)
		hash.Reset()
	}
}
func ReadVideo(filename string) []byte {
	b, _ := ioutil.ReadFile("E:\\200617workproject\\java\\src\\main\\java\\chain\\tj\\file\\video\\418.mp4")
	return b
}
