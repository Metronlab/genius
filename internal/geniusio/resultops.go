package geniusio

import (
	"fmt"
	"github.com/Metronlab/genius/internal/geniuserr"
	"github.com/Metronlab/genius/internal/geniustypes"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
)

type GenerationWriteFunc func(spec geniustypes.TmplSpecPaths, generated []byte) error

func GetGenerationWriteFunc(dryRun bool) GenerationWriteFunc {
	if dryRun {
		log.Println("genius dry run enabled")
		return cmpResult
	}
	return writeResult
}

func cmpResult(spec geniustypes.TmplSpecPaths, generated []byte) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("comparing file %s, %w", spec.Out, err)
		}
	}()

	expectedRes, err := ioutil.ReadFile(spec.Out)
	if err != nil {
		return err
	}

	if assert.ObjectsAreEqual(expectedRes, generated) {
		return nil
	}

	diff, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(expectedRes)),
		B:        difflib.SplitLines(string(generated)),
		FromFile: "expect",
		FromDate: "",
		ToFile:   "have",
		ToDate:   "",
		Context:  1,
	})
	if err != nil {
		return err
	}

	return fmt.Errorf("%w: %s", geniuserr.ErrDryMismatch, diff)
}

func writeResult(spec geniustypes.TmplSpecPaths, generated []byte) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("writing file %s, %w", spec.Out, err)
		}
	}()

	stat, err := os.Stat(spec.In)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(spec.Out, generated, stat.Mode())
}
