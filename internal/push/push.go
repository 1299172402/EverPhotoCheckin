package push

import (
	"errors"
	"fmt"

	"github.com/winterssy/EverPhotoCheckin/internal/model"
	"github.com/winterssy/ghttp"
)

func Push(scKey, text, desp string) error {
	resp, err := ghttp.Get(fmt.Sprintf("https://sctapi.ftqq.com/%s.send", scKey),
		ghttp.WithQuery(ghttp.Params{
			"text": text,
			"desp": desp,
		}),
	)
	if err != nil {
		return err
	}

	pr := new(model.PushResponse)
	if err = resp.JSON(pr); err != nil {
		return err
	}

	if pr.Errno != 0 {
		return errors.New(pr.Errmsg)
	}

	return nil
}
