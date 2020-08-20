// Code generated by "esc -o generator/data/tpl/tpl.go -modtime 0 -pkg=tpl generator/template"; DO NOT EDIT.

package tpl

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
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
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
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

	"/generator/template/echo_enum.gogo": {
		name:    "echo_enum.gogo",
		local:   "generator/template/echo_enum.gogo",
		size:    491,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/3SQMWvDMBSEZ79fcYQMEjTOnuKpaSFLUmjoEjIo9qsxtSUjy0MQ+u9FtuNQSLZD3H36
eOs13kzBKFmzVY4LXK5orXFGtdUrtgfsD0e8b3fHlKhV+a8qGd6nn2MMgchd2+FprxoOAZV2RLnRnYOg
xPsVrNIlI/2ouC46hBBfb23vl7eYRci3qnseKyuwLkIgSfTT6xwij6LzVOLL2UqXQqIbAjwlWjXcYZOh
Ue1prp7Hgqfkic9daIPFnBcv02D0SAJRYtn1VmP45xSFzhSe+cXDChkPEtWmpai0k0NVUqCHOv9xy5m3
62a4kLgYU8MTAEzoYZBld4XpA9ZFxP4FAAD//9bZDJPrAQAA
`,
	},

	"/generator/template/echo_service.gogo": {
		name:    "echo_service.gogo",
		local:   "generator/template/echo_service.gogo",
		size:    2781,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7xWb2/bthN+LX2Kq3/9taRjy4q3IksCI8tiu8uA/MFi7MXCNJOls01EJg2KSp3K7Gcf
TrI9O033L8HeCNTd8bnnHh4PbLXgRCcIY1RoIosJDB9gZrTV0UweQvcCzi8G0OueDgLfn0XxXTRGKIrg
slo65/tyOtPGAvO9msExzmc13/dqY2kn+TCI9bSVRsPMRvFdC+OJrm37HrT+pHVrlXG9GOuaz30/1ior
oXvTSKaQWSPVGDpQ+8AYY9dR89Nx89ebhRDJ4vrV/4R4/f83b4WoC7EjRFOIVkeIIyE+3P5WCLFwn28W
10LMizA8Dl2TVt29ft/Ror8fLk397snS1O2vTP1e393wHSZE8J8n5XW+YEyIebvNqWhahQsh5uE+r9Nf
mNAn4vxoy7fDjxijzOEu4YXf0WdIn5g+SMbdkRDzvRHVMm/vljzb35Dj3bBi/S6hvz18Xg0LJsQmlf1t
KlWO0fNycM7r/1CdSlPOv/+ilZ5V68thbSM1F0IEi9vF52dh1l+OHudCBHxno+CXEu4lRXtBwZ4t1tHr
cqrdR4ZmmplXU60D1dwMzvLMnujpTKbIShen6FaLBu55NEXnQGZgJwhSWTSjKEaItbKRVBlEaVq6yGB0
mqLJfPsww83N612F7xVFE0ykxgjBGdqJTjJwjszBQNoUnWM0roMTrSzObQPqRRGcqlluBw8zdI4DI8tF
btemopAjUAhBzxhtyAa1mnPV1rWN4lAlzvGKA6qEEjvff5rRKFcx3K6LuP0xUkmKhmXmHori9dLMoWS7
dPZpT+F7Bm1uFBAEi2GzHg4MjQEkVpxCPangoAMKP7JHhfrkHFEodCAOfpAqYVLxw9LyqgNKpiWAZzCb
EcaJnk61KgsuKLpcHcCb9bo4wyyLxnhAEJUyjJP43opxHPx0dXHOvm2HDSBY7nueWxK5j9KeMZRIquCX
KJVJZJHxw5XjryittixpVbv+VnKd2yfPGOiQSymXR0spM3MfbDRT3ACpCIlO+WmMtc76cRFbtMJwmazi
5a26yJUct4PbFKxzy32PWmyj3+ha/YxjmVk0W9crzzABq2EoVQJG55YuUtmEX4QzhHrZVL14ohtQdeS6
IQvfa7Ug+yhtPCFAerzEFojX8eUpNQOaBkmRZ/SoIaC3GXRxFOWprdw+CXLbAH1HimJQWQNWZd0K5YcU
RXqtwqDq5j9eVMFW6lKSPxkD61O6QnMvY6zO6fLialCdFQbvewNWK9+CduJcrfGVW8q3L/rXsd/31tCU
5l9gPxoovwcAAP//F5mGs90KAAA=
`,
	},

	"/generator/template/echo_struct.gogo": {
		name:    "echo_struct.gogo",
		local:   "generator/template/echo_struct.gogo",
		size:    924,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7SSUY+TQBDHn7ufYiSmgeaO+oypyeWghqTX01r7YszdWqZ1lS7csDXWzX53M1s4qNFH
eSDDf2bn/5tZplO4rQqEPWokabCALyeoqTKVrNVrSO9heb+GLM3XsRC13H6XewRr43fn0DkhplMWbkvZ
NEt5YMmcavxDg8bQcWvAipG110BS7xHiucKyaMA5VuO1MiWXcniqOXr81lQ6CayNz12Cx/Nx1AUfcuKv
vcTuqLcQEkysfTmEiOAtmt4ojAZWVgAAqB0QzGagVdkq/PyQBL+Qqo0s+xPPWUJzJN0VePmcbBMU95Yt
csv/DHq5qwg2slSFNBhGMOnijKgi3h8SNZDM4NPniZ/ZJ6z792Kveay467PCp6Mi9AAjnneAx6MHAZuM
kC3mebZIH1bZ+4/5KktZZe8ZyLpGXYT8dQXjAYUPeYrkou8V+DQvLoExukiM3MVNdpT41IPOKzpIAwEe
pCqDDvcF/cxYiO+k2X79YEjpfTj0inr8fLm5WeTpQ3Z3ky/+N30Xqx2UqH33CN7AK4/T/grji8u0/t0k
wLXON23rtCqFE78DAAD//3AqAjycAwAA
`,
	},

	"/generator/template/echo_v4/service.gogo": {
		name:    "service.gogo",
		local:   "generator/template/echo_v4/service.gogo",
		size:    2781,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7RWTW/jNhA9i79iauQgbbVSsEgvDnxos242BTYJNkZ7XDDS2CIikwpJ2UkF/vdiKPlD
sh1k0fYmc+YN38y8GTpN4UrlCAuUqLnFHB5fodLKKl6J8UJdwuc7uL2bwfTzzSxhrOLZE18gNE1y3346
x5omuVlWSlvjf5wZGE8gcY4JfwohCzArFIwWwhb1Y5KpZVryR2N59pSSJV1djFjEWJpS4Fu+ROdAGLAF
gpAW9ZxnCJmSlgtpgJelN9GBVmWJ2jD7WuE+eItqWNA0H0HMIfm1tsU3fK6Fxtw5Ou/cyRBmQFSSKyUt
vtgIQtQaUGulozYEyg71ETSXC4TkK9pC5QacYz7aTNgSnRuEikHjM3G7kVVtr9XstULnIgg1morO72q7
Z2gaMQeJkEzpbjqD0ci5GB7F3/6IIP5jh/DUYjjCmLg5xk6UYF7LDL736vD9C5d5iTo0erUraNRm9FXk
eYlrrvF3QjYs0GhrLYEChRJfbOvXxSCn6OCEYD3cG5X3vgH9noDRq2TQsoiRWcx95j9NQIqyRWx6fmaS
L9xcqeVSSV8z6mAQBGkKOJ4QLAlJsUnPJYK1KEuouBQZReHGoLZCSZhzUcawLkRWkEClsrAuuIU1wppL
62MTnRjUE7xxwSXZW6abUmTJHw93t+HF+XkMGHmT26bSaW/f+8FqIRfhL95f61YSYeSRztelc6W2UKl8
PFLDLmDTHCp5oImBHs76ghi09Ue6SkMxnoDEddifje4CSqHr7ASy5Dch81Djc3R50Ox+ATcFaRNOPzCq
4EAKf/JS5NziThJiDiteTrUmThqfk41LGF1uLP07TUWuw97eK796nGt6d4y7GO6Q78Wnc9oQpuo6FHTj
zD6krL3n6EaA/ZXQ2wDj7ahs95FfQRE7OipvT8r/OCg/OCdHqoad2IdD8p4Z8SPS5X6ktm2ttiv3tNz8
vG78di3csGFD/0/nu3bvDaN/Q9IUvuFCGIu69w7WBnOwCh6FzEGr2tKL5+f0wD1E+ODnbpoVKobBEm9Y
cID4S9jiXuNcvIToATGMRhE7QWfnfYoYrIUtIKuNVUuovOsJrvs3n2Ydd0HA+GZSEqef9IUXU3KtVV2F
He5nGDVN8oB6JTK857ZwbhS/8epFvff+1IO/r50udque+7uH2Wjf48wc0kyup7OQaLURj7A6yqg09H8r
QA/vpffeCK3S9r9PZ3I9fUcilO6/ycTj/6NUBn96/gkAAP//4IxIqN0KAAA=
`,
	},

	"/generator/template/go/enum.gogo": {
		name:    "enum.gogo",
		local:   "generator/template/go/enum.gogo",
		size:    494,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/3SQsWrzMBSFZ9+nOIQMEvxxdv94alrIkhQauoQMin1rTG3JyPIQhN69yHYcCs12EOd8
+rjbLV5MyahYs1WOS1xv6KxxRnV1Vpn/2B1xOJ7wutufUqJOFd+qYnifvk8xBCJ368ang2o5BNTaERVG
9w6CEu83sEpXjPSt5qbsEUJ8vbe9X99jHiGfqhl4qmzAugyBJNHXoAuIIrouU4kPZ2tdCYl+DPCUaNVy
jyxHq7rzUr1MBU/JE5+HUIbVklf/5sHkkQSixLIbrMb4zzkKXSg884u3FTIeJKrNS1FrJ8eqpEB/6vzG
rRfevl/gQuJqTANPADCjx0GePxTmD1iXEfsTAAD///3pkaPuAQAA
`,
	},

	"/generator/template/go/service.gogo": {
		name:    "service.gogo",
		local:   "generator/template/go/service.gogo",
		size:    3138,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7RWTW/jNhM+m79iXiN4K6VeKVi0Fwc+bBM3mwKbGBujPQaMPLaIyKRCUnGyAv97MZRk
S7KdzaLtTRrOxzPzzAwZx3ChFggrlKi5xQU8vEKulVU8F+OVOofLW7i5ncP08noeMZbz5JGvEMoymlWf
zrGyjK7XudLW+J8TA+MJRM4x4aUQsMFwJWxaPESJWscZfzCWJ48xJqkads9elfqmVNwg2H6s1JCFjMUx
Rb7ha3QOhAGbIghpUS95gpAoabmQBniW+SMSaJVlqA2zrzm2jbdWJRuU5QcQS4g+FTb9ik+F0LhwjuS1
Oh0ECRDg6EJJiy82hAC1BtRa6bBygbK2+gCayxVC9AVtqhYGnGPe21zYDJ3ruRqBxifCdi3zwl6p+WuO
zoUQaDQ5yW8L2zooS7EEiRBNKTbJYDh0bgQP4psXkYn/2Fl4aCM4gJiwOcaOlGBZyATuO3W4/8zlIkMd
GP28K2hYZfRFLBYZbrjG38myZAONttASyFEg8cVWerUPUgr3JGTWsXuj8l53QP8TMPo56lEWMjoWS5/5
/yYgRVZZNJyfmOgzNxdqvVbS14wYHAwGcQw4npBZFFBLRx2VEDYiyyDnUiTkhRuD2golYclFNoJNKpKU
GlQqC5uUW9ggbLi03jfBGYF6hDcCnNN5hbQpRRL9cXd7E/zy8WwEGPojt02l7r229p3VQq6CX8/OPPFV
SwSht3S+LrUq0UKl8v6oG3YOy3K/k3s90euHk25D9Gj9EVZpKMYTkLgJurNRB6AUamYnkES/CbkIND6F
53tkH+CatFt8+0kbT6DPxEz5ReFcuTUYw/+332Wrro2jPa7Id9jAwMzgIc2GpZ1iTahj8SkjUS+BP3km
FtziLgmxhGeeTbWmRDQ+RY1KEJ43J+2qfD/nToxx7eN7aTqPtoJ/GrMqzsGdBe2l1dlR4+0wbzemX5Ih
OzjMb8/yfzjKPzjJRwfZ9Vl/1xT7Ia5zP1DbqlbbS6FLfQcI+W70dhQ2aFhf/+PZju7WuvC3XBzDV1wJ
Y1F3burC4AKsggchF6BVYelO9ptkTz1AOPWbYZqkagS9a6Zkgz2Lv4RNZxqX4iVAbzCC4TBkR+DstI8B
g42wKSSFsWoNuVc9grUd+TjqUe0EjCfTJxHHYDbCJikFJ3ligar7aXZN6wX1iGgtjJArvyd/MnCJS15k
tjpmRO79tvmiShoFFYiO6q4FGzWo1urudRV1Qntejz+LVlXEK62KPKgz+xmGZRndoX4WCc64TZ0bjt54
OYSdN9OxR1O7u2vfVX/Pbu/mw7bGidmHGV1N5wHBqjweQHUQUb2h0Zt30nuvh2oW2t/HM7maviMRSvef
ZOLt/6VUeg/HvwMAAP//TNNGskIMAAA=
`,
	},

	"/generator/template/go/struct.gogo": {
		name:    "struct.gogo",
		local:   "generator/template/go/struct.gogo",
		size:    1068,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7RTQW/bPAw917+Cn/GhsIPW2dlDBhSJMxhI0y3NchmGVosZT5stubQyLBP03wcqduMU
3XE5BDQfyff0RI3HMNUFQokKSRgs4OsBGtJGi0ampX4LsztY3q0hm+XrJAgasf0hSgRrkw/H0LnA2iSv
G02mdS4YjxmcVqJtl6Jm2BwafJGD1tB+a8AGF9ZeAwlVIiRziVXRgnOcTdbSVFzK4aHh6PF7q1UaWpsc
p4SPx3ZUBTexkFdmBbu92kJEMLL2/6GIGN6jORFF8YDKBgAAcgcEkwkoWXUZ/v0UBL+R9EZUp45nlNDs
SfUFPn0EO4CSE2UnudMfWCt3kGxEJQthcIVPe0lYOPd8gnMTY+hLoxhGfZwRaWJjkaiFdAKfv4y8GR6w
7u+OX8Nr9B5kIwa62ZMwZJILZIp5ni1mD6vs46d8lc04y9wTEE2Dqoj46wouByp8yKdIz+ZegYfZ0RQu
0cXBhTu74l4lPp2EzjXVwkCItZBV2Mv9j35lnEhuhdl+uzckVRkNueKT/Hy5uVnks4fs9iZf/Gv1fSx3
UKHy02N4B2+8nG5HLs8u0/r/NgWudX5oV6dk5VcIVeGfId9f3k51XWvlm+79Mxts0OjlCvmyKOYHKVXJ
IrrZoUfCwfw/AQAA//8uGBAWLAQAAA==
`,
	},

	"/generator/template/go_client.gogo": {
		name:    "go_client.gogo",
		local:   "generator/template/go_client.gogo",
		size:    2379,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6yVX2/bNhfGr8VPcV7BeCGurlwU240BXzSJNwRYkyJxdlMUKC0d2VwlUiOpZo7A7z7w
jyTbWbeg2B1N8hz+znMeHy0WsNlzDVwDg4rXCDsUqJjBErYHaJU0krUcsq+oNJci77qnLi9ksxiOKFks
4JcxiJkl9D3kG94gWOsOr27h5nYD66vrTU5Iy4ovbIfQ9/mHsLSWEN60UhnISJJuDwZ1SpIURSFLLnaL
37UUfkMpqfwRlwsuO8Nr90OgWeyNaVNCCTGH1uU23NQI+Q1r0FrQRnWFgZ4AALCWP9z96va42BFLSNWJ
ArIWfjgLo3CP5p2/nXWqjhEUepK0ecyygk7VLkl8N7+UzVqpZw8nff8aFBM7hFkFyxUMF3/mWJfaKeXY
zsH73uedVdbCZyfDMu37eJh+jiGvAUVp7VQKulJOSSisnXgZjVU4JIWmUwLSQjaNFODVTV2WiTV/j1qz
HXrAf5f2WZH/aXVuzSuQCjKuL/jTWqmQirqNUG7csPZYirOufluJLX+aZJiePl1FZdaia3TcmqHL7SsO
jwxahQNrgQtDSCGF9h4/SjMplIy1HweunKl+Y3WH4cpAQodmF7L0xhs966s6KU+wBrWja1j7cbz6KVzo
SfINngloCUeNmceAwJFYMurn3/nogD5NZjznu5QlZtQJcqR8xoWh/iolJypPOCfZBnkoXOsxd0ZhK2Ud
vRgz+4DVaiKYOgv2bDl63uxlqcc3/WSYjQWc+SlT+McVM8z/565F25nNofUXM4V6PLntzHg0dyYLRvPT
xFn/3qiwvVyB+52/Z0rvWT2kpyThlb/wvxUI7qpMRt157WNDM9yoWq5gHFGvfPMe7q6tTZ3genzHDc38
g9TGjbc5pKxta14ww6UIM3cOfhjnN/h40VUVqiyi0pfiJCVWqEChzi9kecgva6kxoyQUfXEwONKEgZ7f
ISvf1a7uEPLil/QjN8XeP3VvmOn0pfcKSQqmEd6+ebP0gaElyxX8/6wrvbO8eyc24EE0sQUTagynJPkb
qOdUDmvYjaFzd+wN6rF+DFjbMM0CFRcl/gn5besaoSF89tKXAYZE38Pnd0P4Ed5bj1eE2RrwZqdflpdQ
hfjvpgrhE9VPQy9PAqXyRs1SLgwqwWrQqL5i/KPBElJ4FYfiiOZ8XGLFutr8U8ZOfBHyUYD2vvIzJaXO
dMffhr8CAAD//1ut0G9LCQAA
`,
	},

	"/generator/template/markdown.gomd": {
		name:    "markdown.gomd",
		local:   "generator/template/markdown.gomd",
		size:    1858,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7RVQWsTQRS+7694Nj20hd3ciwpiK1baLqSp50yzL8lqdmY7OxuokwHtxYNgLiqi4KmI
J6soiFLsnzFpbv4FmZ2dZDcthZQ6p8y877037/u+ndy84bruUr0TJhAmQKAVdhHaSJETgQHsHUDMmWAk
DmGphzwJGfXS9EnqNVlUtaHlZde97ZhSaz5s+3VYX9uoe9mxIyUntI2wGKGA1VvgbaHosCBRCqSERZaK
OBX3QuwGiQ630W48PwvVD2IEpaSsruggiA5CyyBanEVgCkCESULamIBgsIfQZLSHXM8gGDxKGIWVqlJO
BaTUF/G2SYRKOU6lUoGzz0/HH5+NBoPx6Ze/J+8c14LusihCKixufPxj9PVwt7ZpQI0ctVvbUKpRhIze
/ByeDGwpwXbjGHlGgHdfiHhLBLbkcHA4ep01LdBkedjKR/JbGQdZ/ga1lJS3erZ8uCQfDqQMW4D7MDk7
l+PWfL/uSok0sHhqrmBnh4UFpZbyuhM+lvMUpw8x4SRCgRyo7gB94LifhhwD/Vvom0IfAkyaPIxFyKjT
h1W3tKC/WvhtolBCOVK6YPjxjDt0byntpPkq9s62UnprRBAzrJkvTGoYIxHgbZI97Cp1h3NyYCnIUqay
67aWm+oKtBhH0uxY+yENjKvy7AyjW2iCp1FjjNNXw/cfzo5+/fn9QsvtNBoNbUtHyog8xgc7/nb5W1BK
Q4rJxivj4++jty/ncEzhK5rZz+GZYtJ/MY0xynRdahkdn/HQVQ1zbQ6ZxyD5k3Vh2JaIzCs5Y6N1mkYX
OABpGmVvq44nE131cUlYmmPPCVU8LCmVTTJRqUe6KV6qTkGZqyvyULe5Dq4txvQtUzk8ej769kn/Pxli
vB3kvbA54+LsBT8XmrZx/gUAAP//nIj5n0IHAAA=
`,
	},

	"/generator/template/php_client.gophp": {
		name:    "php_client.gophp",
		local:   "generator/template/php_client.gophp",
		size:    3870,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5RWWW/jNhB+16+YGgZsB46SbR+6iOsU2cRtA+RYbJMCxWZh0PI4ZiuRikht4gj87wUp
idZB2a4egpic45uPc/3ya7yOvZMTeFhTAVQAgRUNEZ6RYUIkLmGxgTjhkpOYnsXrOAgpMgnD75gIypmf
pu+pH/DopBQaaWtX93B3/wCzq+sH3/MYiVDEJEDIMv+ORPin/qHUxPNSgfA35++cP33WBi5iOjGHt5vL
G7IQTzOWRubPxPNOjo7gFoUgzyjg6OjEy7JjSAh7RvDLc6W8ICRCQJZJKkME41ApI0tXQMUn+j5LkvIc
8E0iWwoo3T/l9zyZvQUYS8qZUcVQYK5/yaOd+pc8ijhzmWBLpYBGcYgRMllRKcB7mQcAUInqN4rhUoBS
5kJTjIF+lH5BpOawVDHmvVwwXYQ0gFXKAu0dKKNySJKEbKCfoIg5EzjKFc3fnV71R1cwpEKgHFr9rz2L
ofdtNKpYKq3RFeAL+DdkgSH0bi4+zW7mX2afZxcPs6texbj++nJNxfG5NQlTMHiHo0lNbsUTJMEaOnAA
ERVumqAqwKi4X/yDgQT/ikjysIkRGohyVFEMU2D4uk2nUl6pJrZS4/jc8F3B0SX3nYR0SSS6LdUZ+foN
pkZt4oxJ56czApeVRva0bJlMqp4q7wB/B1PreOwDKW5qllS7k3Kfbhf9HfE5YLsdO6y1GG2fKs9911HT
W/D/t5B/KCq5EU6rhuU64a/mZWyj+t2MhdD2tWFvYPUHenwwLgHfqJC9CqfdkeVn7mbXiFegnFtfwyxz
ZppSjbYPWWb8VYqxQZfrVRv10fECzzVETbsJyjRhLfOTWuRbLprGJZ8XDdBtN7887NF3FmdDrpLHMD13
FM0W2PiQotljb7ynKNono5JBPdLzK70X6BWhtRSYw66NwE5uLbVv9AacCQnVJMky/y8Spo4RvAVW+rU7
SHOMr6WML81GNXHn2HxuPCdpIIf9BRH4mFCYwuDDjz/7p/6p/+Hs4+nH00FHUm/NFy3WFvIf9mZYe4FG
WpXfQLuepwkdmEcsgIzbcpJGyFNpxH46rQuM2m9YEtcP9NIUoV6a4GwKfr5kFY+gGo3iFuWa2+MmZdUm
4V+zOJV5mutm/dIiak3YMsQEplsD23Eyhv6CvhtU4wIja3ZJ01ExiuVmqzdybR761k46/z6VFplz/Cco
mtOtU2zXHlH2oQRF/VLZndbiL4N1418Ui7ENgrIlvoF/b0aBgJ7R7XXEU2qXQVlfbdl88FiNvbCLh3Gv
e3nO1BLs2LEoVTbc3Nzu/daG1V5i9myIda2SDZdPt25Bzo4VrmCoaELZDit7Zvsj+5fxVwY5OMjJk5sY
z6Dn78rKXcvjQZ7vOCyJJEXy4tKvrRNFr2xP2W2/Oz4PSBhexFRXz8vYjCHd827lUqle/vvxy7X5v+wB
tisVzCnvvwAAAP//2W6eJx4PAAA=
`,
	},

	"/generator/template/spring_service.gojava": {
		name:    "spring_service.gojava",
		local:   "generator/template/spring_service.gojava",
		size:    856,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6yRTY7iMBCF9zlFKauwGHOAbBADQiwgaMgFKkkRLKDssSuNUOS7t8JfsuhuqRFZRX5+
z++rGo/hr6kIamJyKFRBcQHrjBi0OoVZBussh/lsmasoslgesCZoW7W5/YaQRpE+WeMEjKuVt05zvXN4
orNxB3WmQhWaK4XMRlC0YbUgWaG1muv0t9aN8S97/5G3hj1NTXV5wfy/IS83b2Sb4qhLwMKLw1KgPKL3
3VTWeKIQpugJ2ggAoG3/gEOuCdSKZG8qDyE8Fb0DJlBbch+6pPxiCeLFPI8fdyYD3iS+Tl32IcSjmzok
up7ce7WtyhqxjXSJIfTFurhkMmDppCX3NzWP7sW7z5E0jnt7onmUXtUegbjq2v6ItMm2A6Z+/e9AWpAk
b4V4LvXbJ794L71HPZJC9BkAAP//3o5aN1gDAAA=
`,
	},

	"/generator/template/spring_struct.gojava": {
		name:    "spring_struct.gojava",
		local:   "generator/template/spring_struct.gojava",
		size:    585,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5SPwWoCMRCG73mKOdqD8QGWQkFbqBT1sC8w7o7baDYJk1mphLx7ibqtUmhpTjOB+f7v
n81g7luCjhwxCrWwPUFgLx6DqWCxhtW6hufFa62VCtgcsCNISW8uY86VUqYPngUa3+sdRiH+6K3eY3OI
3ml0zguK8U4vo3dzJhTP1b+ONuwDsZy+s/Z4RD2IsfrNRKmUCsPWmgYaizEWv3kZVthTzpAUAEBKU2B0
HYF+MWTbmPP5P7A5ohDsjENbTpd4xPoUzpcJdIFA6TlCyLWQszrvTzedLrSLx73BpKzeCQ+NbJCxz/nh
D6vy5N1EfaMAj78KlfWqldIdE6Zj1y+7m5IdScHWRmzhTka18phkYPczNl9jxuysPgMAAP//ioLHmkkC
AAA=
`,
	},

	"/generator/template/ts/helper.gots": {
		name:    "helper.gots",
		local:   "generator/template/ts/helper.gots",
		size:    790,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xSX2vCMBB/76e4vbUi8V3XMSYyN9AOt6e9jNBebYdNyiWdk5jvPtLaEnGCQaS9+/25
/K7BZDQKwP3gsebEKzCEqpZCoYX+CXSBgESShlIAo0mAv7UkDXkjUl1K0UGWXGS7UmxDJIqm8EayKhXe
C/xBegATAACUeXiHRHA8Og4bfOI4hkZkmJcCs+gEdodQNyR6MUb4jaluHWYtxrb/O9SQcc27mqaDp+Dq
EMPre7JmNSeFoe/MXLvXgpTrtIAQo0v+Bcn3V/uyI/ogpblulC+VcoVQaF3PZYbs6eXza7HZJJvpALh+
Y8MYQ6IxVKgU3+K0nYud3mw0C4wpc2DPKNiSq7msKikWbi3WXvGfJ6tVsv5vBBdnN8aCCGKoeO0Jfhxq
DL3Y+uN2O7Aic9Zr11KQ3Lsczzr29ssP0Y6BMTZYWW+ODHPe7PQtiXbfkDEoslNENrB/AQAA//8kueYK
FgMAAA==
`,
	},

	"/generator/template/ts/helper_common_partials.gots": {
		name:    "helper_common_partials.gots",
		local:   "generator/template/ts/helper_common_partials.gots",
		size:    3664,
		modtime: 0,
		compressed: `
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
UA4AAA==
`,
	},

	"/generator/template/ts/helper_fetch.gots": {
		name:    "helper_fetch.gots",
		local:   "generator/template/ts/helper_fetch.gots",
		size:    873,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4SST0+DQBDF73yK10sLjdneW2mMTRM91JrGkxezwtBiZJfsDlHS7Hc3FIr0DzrhsLx5
85udyXqT8dhD9eEul0Zm2BuyuVaWHI4n8I5AxmjTSh7GE4++c20Y0pYqQlKoiFOtauODVPFnqrb+4S+Y
4tnoLLV0K1U5x94DgDTxB3W6EaowxIVpIAfRNV74g9otYskSwyG4zEkn6IhhGGKk3z8o4lEXWlX/2kRG
1sotdR1V1I4miRCXFa3dtSf7lXK0O+ItSy7sOTiSlrBjzhc6JnH/+Pq23GzWm+mJ6Xz42R+IxXq1Wj/1
USKtLDespTEI4fsBwvnZrY7BpuzJdK6UyXyhs0xXQG1eypw6Cw1mV8sdIlnvJvi/QaFiSlJFcQ/rQnWB
H1yI1ZNqBw+uN+1f8vVGjd/fCyEORTcQQrRdXGd613mxp23cTwAAAP//S/mhHGkDAAA=
`,
	},

	"/generator/template/ts/objs.gots": {
		name:    "objs.gots",
		local:   "generator/template/ts/objs.gots",
		size:    727,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xQQWsTQRi9z694FKF1Mcm90EO0KwhBigQvIjLJfkkHd2fXmVmxDAOCFQ1YKZh66EFP
QqHQ9iRo/TnZ3f4Mma3ZbaUgOKePN++9772vFwQswHBbaExETBAaU5KkuKEIox2sZio1Kc/E6jXaOJWG
C6nB4xhmmxBxw6GNyscmV4QRCTlFrimCkDWhdTUamtRLMSbNAnT+57EAF6ffys/vFuffq/nX8v3+4ufH
ZVIW4PKn+PC22D/ubz0oZnvV8enF2ZtqflTOXhe/Dqr5UXW4Wx6clXsn1fmn8stucXK4+DFjQY+xXg8k
80QzaztQXE4J3dADcI7RqyxVpibA2u5DnpBzsAwArvDvC4qjWnD5sSRu+Pkxj3Ny7k4jIhl5qmPW+rHj
nA9R39TsZHQtySY3fOhBr/D4LVIqVd4f6xtYm5K5lyZJKsMGrpffbsMLaUhN+JhuaAAxwZrQVyz+yJdd
ngsZrcPadm1TkmTUzP+6hLeA0b5J28k5a8UE9ALdAR9RjJVB/244ePYo3Ar7w3BzxbknT//ec8Px2O8A
AAD//zVVtfvXAgAA
`,
	},

	"/generator/template/ts/service_axios.gots": {
		name:    "service_axios.gots",
		local:   "generator/template/ts/service_axios.gots",
		size:    3038,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xVbW8TVxb+7PkVBwthO3ImZKVFrIODsiy7ZLWQaJPVIqGoupk5ti+M7wz3XidYZiQo
DRQpQCQi1DZpoWqpUHlJK7WUBFL+jF/CJ/5CdV/GHkOoyIf4zr3nPOc572MjI84IzNeogAoNEKiAKjLk
RKIPi03IRTyUIYloTouhkfJCJgllAio8ZBKZD1Oz0+CFPoKsEQnLIb8Ay1TWQNYQArrICW9CjlyiocgV
1SXHSsixCFTmBHC82KAcfaNsxRQVyoQkQYA+UKahIh6eR09aLgOm2jQVsMyplMiU+HwzwjmP0yiR1jIR
D5eojwIILBJBPWgIUkWohNy4QIIACPOhTprAEH0g/vmGkHVkEojnhdynrAoyBBGhRyvUSyglThhJ5oOg
skEkDZkzAqMf/+eMwN7Ww+69G+2Xz3vr97ufr7V3bidJcEbAvHRWVzprjzs3b/Ueb+39dK23/mhqdrr3
1Wftl9/1Hlx9+2q1s/28vfu6t/6o9+RJ+8XN7r3tzs5dHdi3r1ahs/Gg+/T7N5tX9n642n799d7WVf3U
efpFZ/NRe+f2m2+3exvP2i+eDgxeXzHYBnUouHtbDw1TK3rnx+6dtfbvG4be1Oy0Ydi5v9PdfDJg+Oyb
N1+udK+tdK7/2rmztXdt1/DpPtju3nrWWfmtvXvX8DBn9fTzp+0Xtzprq3tXVtu7m50bO92NX7rr287I
mEPrUcglaD+K0IIpdfgviihkAovJ58UGCnkiZBVahVhVbz0pt4kEoeUAALRanLAqwkHZjLAIBxfDMIBS
GfJVlNNa8B9EEhUGAe4/G8xTmRaFOLbao0YT4rhob5D5cewkRt2xVss9ERAhzpA6xvHM4vk0h35p/48H
RUDOQ36KMD9QxTeAqGEQIc9NOI62dRplLfQFlCFbRZmFy5CtIfH1IYw0QX32MUCJ+hiFwghGDftLpFfL
Ok6AUjWIsq/walJGpbGx8b/9xR0/ctQdH/+re+Rw6ejho4ezE47jhUxIULaQK/MmhNmzozbg6I/+n8pa
tgTZs6f/c0rKyD5knXjCcfCSdtqgzKH8u7FbgnxD/ReSU1YtQHnSHqGsX/SNMTVg2uDBhJPhKBucmY+4
b6Bi0wQC5SlNNm84l6B17gI2E1MLySEuWHyLN7OoOt0lQtAqs7qiaB0vaFPGCZIuPp04ZeTYPJSBsOZk
ntun0nCZHpuf1E7NO2Xoy6TcNOAcxSc+kQTKkMi46tvJ0ArkVSWElZRQuQw5405OuZPJSN7Uv0mU/j03
c8aNCBeYT7QKE04mE4OnigHyaPQyssbD5T6yFnEypuDffRqORD8C6YQWoa7rtZTUbREiwkldlFSMtNez
PKxTgcdUzByT9EQrEX4vOp5u7tJ+DZ8UpibMCROVkNeT2EMJ8tp1BagOxb5skmh9EU+kjBkuYsrzMNID
AcpwzjRV0fRUMWmpBZcyL2j4KPJGq5DC0VGa1Q5BeR/U49ZbKEGD+VihDH3DQ6X8wPsKhw7BgQM2QgOn
TWzcKDFkDsatFJvIhH0/Jn2o44bzOSMBRHxwliyYtKVcLFoihT5a6X20j5lgC24VpYG3iCYo6Z61zugr
Vxd0HjkfVE1KdGjMKilDMDY/rqwhUx0yrGu6CVJ//S4NgyX0k1bdbyQoNEv5HSq28l2Lkh9CS6kM9+if
AqnRpXDcdyGsk7pp9eLykrWkFt5gScFoHGsBsxkHW89eH7TZK5WTWkgUaAXwol6dc8iXqIenpQ/ujElo
AbL/OjlvhPu7MwWlC8EgmRWqBXSu1LuGPam+1DJOocb7DP5WyzXrNp+MmlbLnWZRQyrlOC6U+jOn1XJn
GrL/ApeB4RLy4WGTGmdQTu/svN1IRci2WoOAxnFW31gWWZsGm6v0rDRVrZRNKLRmMvKI+BDNSZVFHSgd
sj8CAAD//w3QI9XeCwAA
`,
	},

	"/generator/template/ts/service_fetch.gots": {
		name:    "service_fetch.gots",
		local:   "generator/template/ts/service_fetch.gots",
		size:    2436,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xVXW8bRRe+319xXquSP+Q41XtpN5VKASUI6qgxohLiYrx7bE+yntnOzCa1lpFaSlqQ
UhqJCIm2okWoqKKoAQkKaQn8GX+kV/wFtDPjrNcuEr5azzznnOd8PbNcqXgVaPWohA4NEaiELjIURGEA
7QEUI8EVJxEtGhhalM+ZIpRJ6AjOFLIALqyvgc8DBNUjCna42IIdqnrQQeX3zG2HC1httdYhlqSL0rnL
ghlrKmFHUKWQAWXQGkS44QsaKYc2mEjwbRqgBAJtIqlvHRr/hgUJQyAsgD4ZAEMMgASbsVR9ZAqI73MR
UNYFxUFG6NMO9VOPm+grEHg1pgItkgUgqYqJopx5FVj67z+vAieHj8df3R6+fD45eDj+bH/44otpHb0K
2JvR3u5o/+no8zuTp4cnP92cHDy5sL42uffp8OV3k0c3/v5jb3T0fHj81+Tgianh6Ma9tHoWmTm5tWvx
Fpkr2MnhYxvdQe/+ML67P/zzvg15YX3N+ho9fDF+8GMW9dk3r77eHd/cHd36dXT38OTm8asH10++vzF+
dDS+82y0+9vw+MtX3x5N7rvv9OrnT4a/3xnt751c3xsePxjdfjG+/8v44MirLHu0H3GhIPEAAJJEENZF
OKMGEVbhTJvzEOorUOqiWjPAN4kiaRISam/HzE9rL8taO+slawlaV90JskBrT6dz2IdibTlJahdDIuUl
0ketm+1NWWyccjgdtvdFWAUUgotVwoIwHYfMRQ/DCEWx4XkhqnTCUjisQKFN/C1kQaHh+ZxJBT0kAQoJ
Ky67wpWly3g1RqkwWPqAql6hDoUr7727qlTkLgqWd+Gi2Rq1lKZarEORRFFIfTNqy5uSs4KnG56H1wzv
jisEbKB6w9IpxSKsg1SCsm7Zhc+YxiJseHrRXqJaNZxLlnodkg+3cDB19NH0Q09dClSxYNBsp9tRI1LS
LnO2suryL5tQpjf+tPJpT7M+wJLWnre8DLFEqwcekQPmZ8TSlT23xlpmKpqxSj/OlwwFiWKb+jilZuvX
R9XjQf4sIoL0ZR2sG69ch3XB+1TiOecQPgaG2yjOu+RsE2NTsZnBKLk6Vqehqy5cuWHMlBg4B5kTk9Rl
lDzcRlgBskOoO0wbVZ3Bz7Ivrjc3WsVq7q7Ng0Ed3tloXqrZ5GhnULKplfNI14bTM11ueHO0hGUUpEuV
p+W41tJZK80a0g6UcgipiIol/G9lBf5/9mx5LhUbJyJCYvBWuk+ny/AaqvV87OkcLcDTFZ/DGr1YAFpu
Lbym5uDZxb8ZvdZgEcy35oB8axEkMKACfYXBHDi7WDQyC5xDp5OyAAuIIvVcI/MYnfuneoLvQGmmH25o
81C31rnxIHK6eNbCon2Svt4lFKK8MPS9VDqzttvpyomqscsIWHazZtNQTkDs45AJv9ZWV4zTVFPMQ2Es
zULXmpF9IEDrBbFLkpp9BEpTZUiS2hqLbJJazwhEktSasTq9mRcKVy4jUnkfVZgzPV8qJEmmg1oXqlA4
ZVKoOpUymmkeL6ON/wQAAP//H/GYo4QJAAA=
`,
	},

	"/generator/template/ts/service_wechat.gots": {
		name:    "service_wechat.gots",
		local:   "generator/template/ts/service_wechat.gots",
		size:    1703,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xUTY8bRRC9968oWZFmxjjtXSSiZcwgBQiwB9gVu5wQh95x2Z4w7p509+wHQx8QhA8R
SA4rDnAAhJBWikRy4pBA8mfsODf+Auru8cx4JRAXy1P16nV11Xs97PdJH54/+O3Z918uHv+xOv/p2Vf3
Fo++K6TQghUZ6YPPLO/cXt67v/z629X9B88ffrY6v7i+v7v64fPF419XP3/69593Fk9+XD75ffH0l+XD
u6uLb5aP7i7+ero6v2gZvrjtwT5+eFbgQSqzQpP+kGTzQkgNFQEAqCrJ+BThij4rcABXjoTIIU4gnKLe
dcA3mGaWQAF9s+SpzgRXkTF19VVfCcYM6gjysTHEwESKOQR0WFX09Zwp9S6bozF7RzdVMGp6gClylEzj
+zIfAEop5NuMj/OMT6GlmGFeoAxGhKSCKw2FFPNMZZMzSCA8OY0gedVdR6IuJYdwjnomxk24TYjCtt9J
NCmOJ7DvaBHCUKIS+TEOJN7EVG/gAU5OP/AHfBi2QQBKqacfdIKqTFNUKraMngZqbh8xXfCEZXkMIUrp
of50920amInIxh+bMcSsR3Nyut8ZTjMoO6UG8R7eKlFpSLroMJA+HESEHDMJR0zZrUACvZnWRTwcbr/8
It2+tkO3t1+i17bina2drd6IEDx1u5zU4oAD1K/52rCUeQxKy4xPo3qCLW0p85Ft3B7GSj07FB8ht8dp
++dfmK+vgaFDbbJ3WVza8TuRpmsJWnG3goSrpkZ4G7QSN8YX+lXbqt7+3sFhz1XYTDYBvOV8coDyOEvx
HT0Guld4g0DvrRsdcJemk/BecQCnfZt3lDfsl3Vdh9F2enkiVUW9r8KCSTZXsY3s8qLUttiYKF6r+pWq
onulbjLwCXA8RumVnaOGzrIg6RozrHc2gF5VtZM0pucidQe9aERaDzYqCytPXNr6MdMsBt/qAPxM4qCq
6vEYEwxghmyMMq4Cu2ohs4+ZvWkQQ+BWCgG80C7amIjqGXLrptal2QRC+yqJiXUbtadCkiQQiCPrqCBq
3KzlWcfZde/1xGjHqZ6DKbg0xcaMkDKdziDE6L/4nJ/XdNFoXUza3/9RYiLqDrMPQ3vn4RBm9uVE/4ra
ixeCK+yybryv7l1xdNYlTotOlf8EAAD//zZ+2l2nBgAA
`,
	},

	"/generator/template/yii2/Module.gophp": {
		name:    "Module.gophp",
		local:   "generator/template/yii2/Module.gophp",
		size:    1725,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5RUXW/bOgx916/gBQLECZLm3bnJvV2xrwJpi3YdUMxBodh0ok2WNEleOxj+74NkJ7YT
B8X0Yonk4aHJI/37n9opQgTN0CgaIxTFxQ3N8MEdynJOSG4Qnhib+81vxqIX3ET3aJQUBhvrhhqM3klp
jdVUfRYWdUpjnBMyG48JjIEqBplMco6QYMoEs0wKiDk1hsB4RvwOVlUEvloUiYHokLp2sExxzFBYA6dk
pCAAAJ7QrTEU/zOxQ81sIuOyNs78V+UbzmIYxFJYLTlHfXPowQKG3TZETZQZzkkngdKYstd5x1b957Pr
6vyvS0pzEfvebDCVGi+rQwADWu1GPrj6VbdYCsE/imoUNgzbmKBGjGpIF+aWRptrASnlpi7UrfKwe2Is
DAdUqelS1xOfLlOpM2phAXsRhOGH2/vV5Zfn64fbm7o7R2iD9kpmSgo3ueBbp4ihxp85GjuExRK6Lu/2
wvDOqJGfR4Sh901OMYpq42bVn9KHUKU4i6lr0Oy7kcLHDvcM10aKO59keJp+3cOIgm44Xkn5g+FXylni
M/ukvr3nIUanbwOOKIeotdSfqEg46rcbdyznXQU00ft2mrOM61FrqoMaDAsY2B0z0+UWbdAtaDQ/I4Kj
uMkhWwuxNznNbZmxqIOWt5as1Xmt2JL0X5796xA48v2daUh8RVu0j5qvqKBbxzJd0iS5zzkeS1TgSyO+
j1rm6lFzFwhBT9+rF8E3vu5QZYEeDWiZW7w7AbQekD6QK/HM1ItiCgOD+heL0Y08XFw8VCevgLLsRWgq
tgiHyBXanUxMT7DnL4o2Q1le1PIqy0ptVLFZY5r0MqJIoCf/kc7Xzcu1nlT3YrQfe0n+BAAA//86oGPc
vQYAAA==
`,
	},

	"/generator/template/yii2/RequestHandler.gophp": {
		name:    "RequestHandler.gophp",
		local:   "generator/template/yii2/RequestHandler.gophp",
		size:    373,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xPwUrEMBS85yvewYMWbD9gBT3qQQXXYy+hGW2hec3mJaKE9++y6VpQPOzpMW9mmJmb
2zAGY9h6SLADqJT2yXrsj0B1Z0yWv8/eLw6z7IwZZitCLzhkSLq37GZEwmcCO6FxxdL/5oshIirlmqLl
d1C7R/yYBrSPSOPiRLUKuqaplxq6CzZaT2tqX0pNPfah9oFDTq9fAap0EXHYLBEpR/7X85zTZjrpu3rf
Mg9pWvhnrurlOZlXtE6qrTuafJjhwYlGRFRGT5PBTtWo+Q4AAP//qCrWrHUBAAA=
`,
	},

	"/generator/template/yii2/controllers/ApiController.gophp": {
		name:    "ApiController.gophp",
		local:   "generator/template/yii2/controllers/ApiController.gophp",
		size:    1401,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7xUy27bOhDd6ysGRgBJwbWzl67VpkHRdtE2aLspoiBgqHFEQCYZcuRHBf57IUqWJduL
ootqY3LmzOMMz/j/N7rUQSDZGq1mHKFpFl/YGr+3F+dyriQZVVVobBoEtT0DrFWBlU2966cQ3WEvRL7F
5/xuiO4BSv1SKr83itStFmkQ8IpZC7daHKGAO0JZWDiagiYAANBGbBghXD2VTBY+a2evnyvBYVVLTkJJ
EFJQFHtXF9h+V1QKO88OobAEiVvIT+h8w9caLX3sQFGc+njX1bm5vu7SXUPzVsgSjaBCcdcbby4284wl
2whl7HlHgwuWoJlBSUkygqcDUqwg8pN6wp2wZKNZnjcNWs40wqT//Lamsu9+FsfHYtOCDyGrqURJgjNS
JnyEJTxMsO0X+pohLLOzOY3KJImHtd4o/m+S5PFIwf0NmTtl7D8gMyrzp2SGo0GqjRx1c1CMf+vGMPmC
sPiMVKrCOndRIsz/NA0JqvoROHcuF4OvsGyXLEmumNbzzHRiTSeQ1tKLu9vNvGkGTrD4JHVNP/baVziP
nGd+edrrPHtBeqeK/T0zbG2j+CJ8wypRMMKTZK2kTzZunvVjdy46xJ9o3AcKaYlJjmp1kcDXmgYGMZxo
wqC93NL4qTyI1BMzhu3HoKNGqTRq62d4+K/KP6BEw6r3O466fa5o1uejvUZQKwgHeiEIyZUxyGkxi6eC
QFk4F7jgdwAAAP//RNUysnkFAAA=
`,
	},

	"/generator/template/yii2/handlers/ErrorHandler.gophp": {
		name:    "ErrorHandler.gophp",
		local:   "generator/template/yii2/handlers/ErrorHandler.gophp",
		size:    819,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6RQXWvqQBB9318xiA8Grhe53L6YmmJF6EtLwb4IgTJuRrs02Vl2NuAH+e8lBmusVFqc
p5k95+w5M7d37s0pZbEgcagJdjv4+4QFzfZTVcVKlUIwNyZuGuYtc/rsOfDYmVgpnaMITL1n/4A2y8kD
rQPZTCDdGJMuUChtw2qnAABcuciNhmVpdTBswZPNyE/Xmlw997p0aKM9v1HVZZbQQsFYCWg18RIOsdJ7
s91bfn4XtfR1zY0ZDrvoXD/xJI6tUD+RgKGUCWcEI/g/GMQnkm5NhBEcrftJ4Ff0Hje96MitgHKhn6Sc
cFGwvTrov+uC/tbv5rvDNAYnUF2dgkRwRZ1R0k60ovDYAL3oz7lIAur3c8mLR01jmQVv7OqrsL2burhQ
hgHrC9Uv8WWqkM0OV6tUpT4CAAD//zSBgW8zAwAA
`,
	},

	"/generator/template/yii2/handlers/RequestHandler.gophp": {
		name:    "RequestHandler.gophp",
		local:   "generator/template/yii2/handlers/RequestHandler.gophp",
		size:    241,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/0yOMU4EMQxFe5/CBQUUzAUGCUEFBQgBDdI03sSwI2XibJwUYPnuiKw0ovR/tt+/uS3H
Apk21kKB0QynZ9r4bUzuM0DXc3xPyv/RsknkpPNY+BD5EVleqjS5K+sMQAdtlULDkEgVX/nUWdsD5Zi4
ggEiotk1VspfjNMTt6NEdR9gPy79kNaAnz2HtkpGs1HP/fJsX8zG/78Qp8dcenv/LuyOF5VPV/Ou4Rzd
weE3AAD//+j458PxAAAA
`,
	},

	"/generator/template/yii2/models/enum.gophp": {
		name:    "enum.gophp",
		local:   "generator/template/yii2/models/enum.gophp",
		size:    146,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/zyLsQrCQBBE+/2K+QHzA1EsRCu1tLJZc4sKlzW4OVCW/Xe5A9O8GZh56+30mKiY4PTd
Hflm172WsaEnGjKbwb3lmUdBVxkB+cyiyVCP5AQA7iu8We+C7vCUnAwRbRheajPc/+6m9gvnIhH9ooqm
CAr6BQAA//8rGSPnkgAAAA==
`,
	},

	"/generator/template/yii2/models/error.gophp": {
		name:    "error.gophp",
		local:   "generator/template/yii2/models/error.gophp",
		size:    2240,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7RUXW/TMBR9z6+4RJGWSG3/QFjRpgVeBkywF8RQ5aW3q1FiB9vNuln+76jOkiaOO1oQ
edhU+36ce3zOffuuWldBwEiJsiI5gtYw+0RK/Gp/GZMGwUYifOP8mfO7G8EVv6hoGgR5QaQEre3/XUaT
B8YAbhWypYQ2+u6SPmdCcJFtc6wU5QxoWRVYIlO9qI8oJXnAQAcAAFpPQRD2gDB7T7FYSjDGXlSCK8wV
LiHS2vbcoWxTkC2NCZrAzX1Bc1htWN70ZFTFRAjyBJFAWXEmMWkS7d9Xu+4+uoKYSokq7vK/hx2G8EeS
9Cq11egKqPx8/xNzBbMrosjtU4X9qoNQ/AWza3KPBYTXF5fZ9eJLdpNd3GZXoZMR1aSAc/ADSYehak3l
dC5RLbqYGCwRi5JUccdQHNUJaIhUWcE5MHwcPm8L3pg4SW3UdG45jeokBYFqI5g9TsFMLMAEktQzJxYS
3XEajB2+P/Z/NbvF5X8lLyQrm+OBDsiE+MA7uOOPO5nAf3dAwjUp6JIojE/W7ZsX4TpMjSSr1oI/Wu47
X35AhoIUnXnj8KzLPwMqgXEFuKVShb15zd864TQvNDJe0UKhGE03GUm7ns73HHaiXZFCYgpmJA2vDMZq
65X0F/Cra3Dqvn9z5t+AjiqGzj6CPUua1s1cBzymtUXTW7Ku5jyedTfyARk/DAC7dbtNMiyfDojZU+UW
V3xh5ztUt7k8zjn/eYn3dgWcz30rGV6E25FST+f7AZtF61p6ctzKdZp7ZL1vNPmXlek0Gpn0JGPsvqSV
ggl+BwAA//+3ZdI+wAgAAA==
`,
	},

	"/generator/template/yii2/models/message.gophp": {
		name:    "message.gophp",
		local:   "generator/template/yii2/models/message.gophp",
		size:    2204,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/7RUXW/TMBR9z6+4RJGWSG3/QFhR0QovBSbYC2KoctPb1iixg+1mG5b/O6qzpInjjhZE
HjbVvh/nHp97Xr8pd2XASIGyJBmC1jD5SAr8Yn8ZkwbBXiJ85fwX5/e3gis+K2kaBFlOpASt7f9DRp0H
xgAtyhwLZEpCk3D/AaUkWwx0AACg9RgEYVuEyTuK+VqCMfaiFFxhpnANkda24AFCk4JsbUxQB+5XOc1g
s2eZopwBZVTFRAjyBJFAWXImMakT7d8Xux4+uoGYSokqbvO/hS2G8HuSdCo11egGqPy0+oGZgskNUeTu
qcRu1V4o/oTJgqwwh3AxeztfLD/Pb+ezu/lN6GREFcnhGvxA0n6o2lE5nkpUyzYmBkvEsiBl3DIUR1UC
GiJVlHANDB/6b9eANyZOUhs1nlpOoypJQaDaC2aPUzAjCzCBJPXMiblEd5waY4vvj/1fzG5w+V/JC8nK
5nygPTIhPvEO7vjDTibw352QcEVyuiYK44t1++pZuA5TA8mqneAPlvt2L98jQ0Hy+WOGpZVJeNXmXwGV
wLgCfKRShZ15zd9uwmW7UMt4Q3OFYjDdaCDtajw9ctiKdkNyiSmYgTS8MhiqrVPSX8Cvrt6p+/71md8B
HVX0N/sM9ixpWtdzndgxrS2ajsm6mvPsrOvIJ2S87QF267ZO0i+f9og5UuUWV3xp5ztVt748b3P+s4l3
vAKupz5LhmfhtqRU4+lxwNpo3ZUenWe5TnOPrI+NRv9imU6jwZJetBiHL2mkYILfAQAA//+cn2fxnAgA
AA==
`,
	},

	"/generator/template": {
		name:  "template",
		local: `generator/template`,
		isDir: true,
	},

	"/generator/template/echo_v4": {
		name:  "echo_v4",
		local: `generator/template/echo_v4`,
		isDir: true,
	},

	"/generator/template/go": {
		name:  "go",
		local: `generator/template/go`,
		isDir: true,
	},

	"/generator/template/ts": {
		name:  "ts",
		local: `generator/template/ts`,
		isDir: true,
	},

	"/generator/template/yii2": {
		name:  "yii2",
		local: `generator/template/yii2`,
		isDir: true,
	},

	"/generator/template/yii2/controllers": {
		name:  "controllers",
		local: `generator/template/yii2/controllers`,
		isDir: true,
	},

	"/generator/template/yii2/handlers": {
		name:  "handlers",
		local: `generator/template/yii2/handlers`,
		isDir: true,
	},

	"/generator/template/yii2/models": {
		name:  "models",
		local: `generator/template/yii2/models`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"generator/template": {
		_escData["/generator/template/echo_enum.gogo"],
		_escData["/generator/template/echo_service.gogo"],
		_escData["/generator/template/echo_struct.gogo"],
		_escData["/generator/template/echo_v4"],
		_escData["/generator/template/go"],
		_escData["/generator/template/go_client.gogo"],
		_escData["/generator/template/markdown.gomd"],
		_escData["/generator/template/php_client.gophp"],
		_escData["/generator/template/spring_service.gojava"],
		_escData["/generator/template/spring_struct.gojava"],
		_escData["/generator/template/ts"],
		_escData["/generator/template/yii2"],
	},

	"generator/template/echo_v4": {
		_escData["/generator/template/echo_v4/service.gogo"],
	},

	"generator/template/go": {
		_escData["/generator/template/go/enum.gogo"],
		_escData["/generator/template/go/service.gogo"],
		_escData["/generator/template/go/struct.gogo"],
	},

	"generator/template/ts": {
		_escData["/generator/template/ts/helper.gots"],
		_escData["/generator/template/ts/helper_common_partials.gots"],
		_escData["/generator/template/ts/helper_fetch.gots"],
		_escData["/generator/template/ts/objs.gots"],
		_escData["/generator/template/ts/service_axios.gots"],
		_escData["/generator/template/ts/service_fetch.gots"],
		_escData["/generator/template/ts/service_wechat.gots"],
	},

	"generator/template/yii2": {
		_escData["/generator/template/yii2/Module.gophp"],
		_escData["/generator/template/yii2/RequestHandler.gophp"],
		_escData["/generator/template/yii2/controllers"],
		_escData["/generator/template/yii2/handlers"],
		_escData["/generator/template/yii2/models"],
	},

	"generator/template/yii2/controllers": {
		_escData["/generator/template/yii2/controllers/ApiController.gophp"],
	},

	"generator/template/yii2/handlers": {
		_escData["/generator/template/yii2/handlers/ErrorHandler.gophp"],
		_escData["/generator/template/yii2/handlers/RequestHandler.gophp"],
	},

	"generator/template/yii2/models": {
		_escData["/generator/template/yii2/models/enum.gophp"],
		_escData["/generator/template/yii2/models/error.gophp"],
		_escData["/generator/template/yii2/models/message.gophp"],
	},
}
