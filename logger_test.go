
package gologger_test

import (
	testing "testing"
	golog "github.com/kmcsr/go-logger"
)

func TestLogger(t *testing.T){
	logger := golog.NewLogger("Test1")
	if logger.Name() != "Test1" {
		t.Fatalf(`logger.Name() == "%s" expect "Test1"`, logger.Name())
	}
	t.Logf("logger.Level() == %s", logger.Level())
	t.Logf("logger.ErrLevel() == %s", logger.ErrLevel())
	logger.SetLevel(golog.LEVEL_DEBUG)
	t.Logf("logger.Level() == %s", logger.Level())
	logger.Debugf("This is a number of %d.", 123)
	logger.Log("This message has", nil, "prefix.")
	logger.Infof("This is an '%s'", "info.")
	logger.Warn("I have some warn for you")
	logger.Errorf("Run away %s, now", "gnajom")
	logger.Stackf("Oh no, come on see %s stacks:", "these")
}


