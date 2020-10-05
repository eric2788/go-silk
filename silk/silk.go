// Package silk provides ...
package silk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
)

type SilkEncoder struct {
	codecDir    string
	encoderPath string
	cachePath	string
}

func downloadCodec(url string, path string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, body, os.ModePerm)
	return
}

func (s *SilkEncoder)Init(cachePath, codecPath string) error {

	appPath, err := os.Executable()
	appPath = path.Dir(appPath)
	if err != nil {
		return  err
	}

	s.cachePath = path.Join(appPath, cachePath)
	s.codecDir = path.Join(appPath, codecPath)

	if !FileExist(s.codecDir) {
		_ = os.MkdirAll(s.codecDir, os.ModePerm)
	}

	if !FileExist(s.cachePath) {
		_ = os.MkdirAll(s.cachePath, os.ModePerm)
	}

	goos := runtime.GOOS
	arch := runtime.GOARCH

	var encoderFile string

	if goos == "windows" && arch == "amd64" {
		encoderFile = "windows-encoder.exe"
	} else if goos == "linux" && arch == "amd64" {
		encoderFile = "linux-amd64-encoder"
	} else if goos == "linux" && arch == "arm64" {
		encoderFile = "linux-arm64-encoder"
	} else {
		return errors.New(fmt.Sprintf("%s %s is not supported.", goos, arch))
	}

	s.encoderPath = path.Join(s.codecDir, encoderFile)

	if !FileExist(s.encoderPath) {
		if err = downloadCodec("https://cdn.jsdelivr.net/gh/wdvxdr1123/tosilk/codec/" + encoderFile, s.encoderPath); err != nil {
			return errors.New("下载依赖失败")
		}
	}
	fmt.Println(s.encoderPath)
	return nil
}

func (s *SilkEncoder)EncodeToSilkWithCache(rec []byte, tempName string) ([]byte, error) {
	// 1. 写入缓存文件
	rawPath := path.Join(s.cachePath, tempName + ".wav")

	err := ioutil.WriteFile(rawPath, rec, os.ModePerm)
	if err != nil {
		return nil, err
	}
	// os.Remove(rawPath)
	// 2.转换pcm
	pcmPath := path.Join(s.cachePath, tempName + ".pcm")
	cmd := exec.Command("ffmpeg", "-i", rawPath , "-f", "s16le", "-ar", "24000", "-ac", "1", "-acodec", "pcm_s16le", pcmPath)
	err = cmd.Run()
	fmt.Println(err)
	// efer os.Remove(pcmPath)

	// 3. 转silk
	silkPath := path.Join(s.cachePath, tempName + ".silk")
	cmd = exec.Command(s.encoderPath, pcmPath, silkPath, "-quiet", "-tencent")
	err = cmd.Run()
	fmt.Println(err)

	return ioutil.ReadFile(silkPath)
}
