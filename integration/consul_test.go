package main

import (
	"net/http"
	"os/exec"
	"time"

	checker "github.com/vdemeester/shakers"
	check "gopkg.in/check.v1"
)

func (s *ConsulSuite) TestSimpleConfiguration(c *check.C) {
	cmd := exec.Command(traefikBinary, "fixtures/consul/simple.toml")
	err := cmd.Start()
	c.Assert(err, checker.IsNil)

	time.Sleep(500 * time.Millisecond)
	// TODO validate : run on 80
	resp, err := http.Get("http://127.0.0.1/")

	// Expected a 404 as we did not comfigure anything
	c.Assert(err, checker.IsNil)
	c.Assert(resp.StatusCode, checker.Equals, 404)

	killErr := cmd.Process.Kill()
	c.Assert(killErr, checker.IsNil)
}
