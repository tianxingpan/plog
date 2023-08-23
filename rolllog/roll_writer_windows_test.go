//go:build windows
// +build windows

package rolllog

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRollWriterResume(t *testing.T) {
	logDir, logName := t.TempDir(), "test_app.log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		require.Nil(t, os.Mkdir(logDir, os.ModeDir))
	}

	createAndLog := func() {
		w, err := NewRollWriter(filepath.Join(logDir, logName),
			WithMaxSize(1),
			WithMaxAge(1),
			WithMaxBackups(20),
		)
		require.Nil(t, err)
		log.SetOutput(w)
		const testTimes = 40000
		for i := 0; i < testTimes; i++ {
			log.Printf("this is a test log: %d\n", i)
		}
		time.Sleep(time.Millisecond)
		require.Nil(t, w.Close())
	}

	createAndLog()
	createAndLog()

	var tmpCnt int
	require.Nil(t, filepath.WalkDir(logDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if strings.Contains(filepath.Base(path), "tmp-") {
			tmpCnt++
		}
		return nil
	}))
	require.Less(t, tmpCnt, 2)
}

func TestRollWriterBackupExistingNonlinkFile(t *testing.T) {
	logDir, logName := t.TempDir(), "test_app.log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		require.Nil(t, os.Mkdir(logDir, os.ModeDir))
	}

	f, err := os.Create(filepath.Join(logDir, logName))
	require.Nil(t, err)
	require.Nil(t, f.Close())

	w, err := NewRollWriter(filepath.Join(logDir, logName),
		WithMaxSize(1),
		WithMaxAge(1),
		WithMaxBackups(20),
	)
	require.Nil(t, err)
	log.SetOutput(w)
	const testTimes = 40000
	for i := 0; i < testTimes; i++ {
		log.Printf("this is a test log: %d\n", i)
	}
	time.Sleep(time.Millisecond)
	require.Nil(t, w.Close())

	var bkCnt int
	require.Nil(t, filepath.WalkDir(logDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if strings.HasPrefix(filepath.Base(path), "bk-") {
			bkCnt++
		}
		return nil
	}))
	require.Equal(t, 1, bkCnt)
}
