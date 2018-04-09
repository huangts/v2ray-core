package ctlcmd

import (
	"context"
	"io"
	"os"
	"os/exec"

	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/platform"
	"v2ray.com/core/common/signal"
)

//go:generate go run $GOPATH/src/v2ray.com/core/common/errors/errorgen/main.go -pkg ctlcmd -path Command,Platform,CtlCmd

func Run(args []string, input io.Reader) (buf.MultiBuffer, error) {
	v2ctl := platform.GetToolLocation("v2ctl")
	if _, err := os.Stat(v2ctl); err != nil {
		return nil, err
	}

	errBuffer := &buf.MultiBuffer{}

	cmd := exec.Command(v2ctl, args...)
	cmd.Stderr = errBuffer
	cmd.SysProcAttr = getSysProcAttr()
	if input != nil {
		cmd.Stdin = input
	}

	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdoutReader.Close()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	var content buf.MultiBuffer
	loadTask := signal.ExecuteAsync(func() error {
		c, err := buf.ReadAllToMultiBuffer(stdoutReader)
		if err != nil {
			return err
		}
		content = c
		return nil
	})

	waitTask := signal.ExecuteAsync(func() error {
		if err := cmd.Wait(); err != nil {
			msg := "failed to execute v2ctl"
			if errBuffer.Len() > 0 {
				msg += ": " + errBuffer.String()
			}
			return newError(msg).Base(err)
		}
		return nil
	})

	if err := signal.ErrorOrFinish2(context.Background(), loadTask, waitTask); err != nil {
		return nil, err
	}

	return content, nil
}
