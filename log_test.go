package plog

import "testing"

func TestLog(t *testing.T) {
	l := WithFields("uid", "10012")

	l.Trace("helloworld")
	l.Debug("helloworld")
	l.Info("helloworld")
	l.Warn("helloworld")
	l.Error("helloworld")
	l.Tracef("helloworld")
	l.Debugf("helloworld")
	l.Infof("helloworld")
	l.Warnf("helloworld")
	l.Errorf("helloworld")
}
