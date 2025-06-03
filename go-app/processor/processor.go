package processor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Processor struct {
	bot       *tgbotapi.BotAPI
	path      string
	channelID int64
}

func New() (*Processor, error) {
	tkntg := os.Getenv("TELEGRAM_TOKEN")
	c := os.Getenv("TELEGRAM_CHANNEL_ID")

	chanID, err := strconv.Atoi(c)
	if err != nil {
		return nil, fmt.Errorf("can't convert channel id to int:%w", err)
	}

	bot, err := tgbotapi.NewBotAPI(tkntg)
	if err != nil {
		return nil, fmt.Errorf("can't create bot:%w", err)
	}

	return &Processor{
		bot:       bot,
		channelID: int64(chanID),
	}, nil

}

func (p *Processor) LoadContent(url string) error {

	URL, err := normalizeURL(url)
	if err != nil {
		return err
	}

	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("can't get path:%w", err)
	}

	p.path = filepath.Join(path, "./downloads")

	cmd := exec.Command("docker", "exec", "py-app", "python", "download_instagram.py", URL)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()

	outBuf := out.String()
	stderrBuf := stderr.String()

	if err != nil {
		fmt.Printf("PY script OUT: %s", outBuf)
		fmt.Printf("PY script: %s", stderrBuf)
		return fmt.Errorf("can't run command:%w", err)

	}

	fmt.Printf("PY script OUT: %s", outBuf)

	path, err = os.Getwd()
	if err != nil {
		return fmt.Errorf("can't get path:%w", err)
	}

	return nil
}

func (p *Processor) content() ([]*os.File, error) {

	var files []*os.File

	fPaths, err := findFiles(p.path)
	if err != nil {
		return nil, fmt.Errorf("can't find files:%w", err)
	}

	for _, path := range fPaths {
		f, err := os.Open(filepath.Join(p.path, path))
		if err != nil {
			return nil, fmt.Errorf("can't open file:%w", err)
		}
		files = append(files, f)
	}

	return files, nil

}

func findFiles(path string) (u []string, err error) {

	dir, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening directory:%w", err)

	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("error reading directory:%w", err)

	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".mp4") || strings.HasSuffix(file.Name(), ".jpg") {
			u = append(u, file.Name())
		}
	}
	if len(u) == 0 {
		return nil, fmt.Errorf("file not found")
	}
	return
}

func normalizeURL(url string) (string, error) {
	if !strings.Contains(url, "instagram.com") {
		return "", fmt.Errorf("url is not instagram url:%s", url)
	}

	url = strings.ReplaceAll(url, "reels", "p")
	url = strings.ReplaceAll(url, "reel", "p")

	url = url[:strings.LastIndex(url, "/")+1]

	return url, nil
}
