// Package motion holds the rules of the motion plugin.
package motion

import (
	"github.com/akrck02/go-simulator/plugin"
	"github.com/akrck02/go-simulator/protocol"
	"github.com/akrck02/go-simulator/world"
)

// Version of the plugin.
const Version = "1.0.0"

// Plugin returns the motion plugin with all its rules.
func Plugin() *plugin.Plugin {
	return &plugin.Plugin{
		Name:    "motion",
		Version: Version,
		Summary: "Day cycle, wind gusts, drifting clouds, tides and the motions entities ask for (bob, spin, path).",
		Rules:   []plugin.Definition{dayCycleRule(), gustsRule(), cloudDriftRule(), tidesRule(), entityMotionRule()},
	}
}

var readsEnvironment = world.Interest{Environment: true}

func number(name string, fallback float64, summary string) protocol.ParameterDescription {
	return protocol.ParameterDescription{Name: name, Type: "number", Default: fallback, Summary: summary}
}
