package main

import (
	"startProject/Camera"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("Camera", func(t *testing.T) {
		Camera.RunAllTests(t)
	})

}
