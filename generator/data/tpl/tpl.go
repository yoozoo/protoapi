// Code generated by "esc -o generator/data/tpl/tpl.go -pkg=tpl generator/template"; DO NOT EDIT.

package tpl

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/generator/template/helper.gots": {
		local:   "generator/template/helper.gots",
		size:    2363,
		modtime: 1530687262,
		compressed: `
H4sIAAAAAAAC/4xWbW/bNhD+7l9xE7ZKil0p2zqgUKamL+u6DkMdxMknwx8Y62SzpkmNopxoTf77cKQk
y29bggC27p57Ht4LT47PzgZA//C2YJqtYcMEMCiN5nJhzRpNpWUJTALKucowa7xglszAnEm4Q2BFgTID
o4AVHCotBnAWD/ChUNpAXsm54aolCDZMJA1J2H6BbwMAAKfWAG+vP39Q60JJlIaCwshiHK4QbI5B/MOr
83jBR+C/9Y+6f37n3Mlx90+v4sUI/O9PeD+44NEJ97kNHh73/vLeBU9PuH9z7pkfXgyeBgPXCXhfcZEB
g9vrv+CubupK9bHtKanCZomAMus1rdICEmu/YyXax1xpayhRb/gce+CGyOE1/l1haUDdfcW5iQCjRWRj
/0Ah1HXjvVeVyKjNFOKgoHIwdYE7wP7AJNBNDuSVEHCrRdvrl1bi08cbSnGFdbxhokIoGNel5cAHti4E
JvRAKVFsCt7SmCKJY6HmTCxVaZLX56/PPQIxvYAUvkm2xgS8e5SLe+TeCCSfr5zBcCa9J8K250rfwFHC
S2JJO5IXxJF2FEcHe4ESNTP4Cc2tFr/evAkq3c34qCl5Ajf7485zCL5z3rAx9W5BpcWFtT0N7IdAQ1Sm
hBSms4uBs1IpA3KtsAYu4ZCOnBvWcLWyY9fxQiujqJHRkpXje3mlVYHa1MEK67BPQn+0HNJGYLrCeral
bI7YsltkmoKkzj8+2lFRObRmv5IZ5lxi5u9ruNyPEtskR0RSbv1xTIK8BKY1q/cPERk1sQUPQic8bcb3
HaFnB+orSG0dh+BPZ37vFICixOPg/RLZ9myYOF4dAkS50h/ZfBlsaAh3Sbt8MmZwx2NTOpnQ71zgYT5W
EVKguM+TcRt6cQDqZB3djr9JnvTbPjpthz0p+udk/CVy887zOtjsyT7tPNnBjoqqXAbNa2IVUhtSH4bd
iyMMtyUN27vRjkaJmjPB/8Hsyq24tCH9qrgM/Be0aLtbtw/u50D7c5gC3eGIywwfxnngX/qu4C9/hEvw
L31IgChheKDbP9f2Jm93/N7ifubabixf2Br7CKBl1YO1K6nDdTuqAzbL9dm7ta+cgmcX/sTZrH9HMgVv
wmqL8Z6zbOM+XdSG/ueSvVLl0S3bO+fW2D/c//zosG0HL/ZcT7uch+BF3nCHid7Y/wYAAP//Ubm8iDsJ
AAA=
`,
	},

	"/generator/template/interface.gots": {
		local:   "generator/template/interface.gots",
		size:    134,
		modtime: 1530687262,
		compressed: `
H4sIAAAAAAAC/6quVihKzEtPVdBzSSxJDKksSC1WqK3lSq0oyC8qUcjMK0ktSktMTlWortbzS8xNra1V
qOZSUFBQQOhzy0zNSQFrUoBIQBVagdTATYXJV1crpOalgHi1XAg2IAAA//9CAdkdhgAAAA==
`,
	},

	"/generator/template/spring_service.gojava": {
		local:   "generator/template/spring_service.gojava",
		size:    708,
		modtime: 1530692503,
		compressed: `
H4sIAAAAAAAC/6yRPW7DMAyFd52CyOQM0QW8BFmCDmmDohegbdYRklCqRCMIBN698A9qD+3QtJoEPT29
T48B6zO2BDnb47hVLY1x1+CjgI+tTSE6bt8jXunm49neqLKV48YisxcU59nuSQ4YguO2/K316NPD3ldK
wXOinW/uD5g/Okoyek3oqourAaskEWuB+oIp9a0845VUd5gIsgEAyDkitwT2QHLyTYKN6iBsF18pVkOh
clJdrUd1CTucTJE525dOQidv90Cqc2b/XLFdYPbSE883Ha8npn5Fki7ybC8cr8tBVTMSzFP6D7w9SfE3
oK+yf8z4JqCchrAB4gZUjZrPAAAA//8s8GFExAIAAA==
`,
	},

	"/generator/template/spring_struct.gojava": {
		local:   "generator/template/spring_struct.gojava",
		size:    392,
		modtime: 1530689747,
		compressed: `
H4sIAAAAAAAC/3yPsU7EMAyG9zzFPx5D8wKICYmBAd1wL2BypkTkQpSkJ6HI747clKoFcZ7sRP6/z4nc
B42M1uyxtyL3xqTpNXgHF6gU/XvU5oUuLIJmAKC1THFk2CfP4VwwiMzvKfsrVcabjxR09ZmudPpK82aD
1RAoo4cM4HiGiOnLHbsHHnT8jDVPrh4p00XkbpUYsNNYJLTquy92Q8TDTb6Oi8W/p612m5tGrhp78jVo
7uFHTStznXL8i5XfcDHfAQAA///dOjtLiAEAAA==
`,
	},

	"/generator/template/vue.gots": {
		local:   "generator/template/vue.gots",
		size:    1199,
		modtime: 1530167803,
		compressed: `
H4sIAAAAAAAC/6xTT2+cTgy98yks9JMC0gZyjMgvUdX866FqojTJpephCl4WFcZTjyd/ROe7Vwy7FFaN
eumext5nv+dn03SGWKCHc+oMadSygkeH4GHN1MHBk8NDw2SQ5fWwwpJYCfHBSbQtfHR4h5Yclzir4G3q
N66Hvs/u8IdDK97DagytIW3R+4kuy/s+u1Ci7l8NXjUtej/vUaNGVoK3ZOWB29WUuMYhnrX5gK1BPoii
R4eZs5jMhKYnUfRumjfCl9C9wrVyrUDZKmsHeefD45PqBn34IqgrG6zpIwCAPIeStBV2pRCH1CxO0i1s
+FlnkJM0xD7aVb9XFuGB2xAbbp6UIHxTFh+4LcAKN7qGU4g3IqbI85ZK1W7ISnF8dHwUT22uL+8hvPs+
u3K6lIb0KDoxilVXLIxPC7hl6hqLH5vv+P9iCV++wk+4ZCY+gz6a1Oc53N9c3BRQoSB3jUZwFqFGyQ1Z
AVWWxNUgVgh4JJoXN+uAH7BTOjgVUsth9/abyKax2daTFcTLpcRjZjl0nJ78gbzGfe4a36AeL2mPOTg5
nM0/MYVRHGsIFP8N681qlGRUlGayQZ1M2BE/rghOz2ZXNZ+HWsxaqpMdcmbCHukOkVVKFCi7/A6XZX61
CHE4jb9rCLC3BWh8Ho8siT9Th7IZTHpm0nW8V+SnaPuHj/yvAAAA//8kA0wHrwQAAA==
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/generator": {
		isDir: true,
		local: "generator",
	},

	"/generator/template": {
		isDir: true,
		local: "generator/template",
	},
}
