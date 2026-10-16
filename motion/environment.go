package motion

import (
	"math"

	"github.com/akrck02/go-simulator/plugin"
	"github.com/akrck02/go-simulator/protocol"
	"github.com/akrck02/go-simulator/spatial"
	"github.com/akrck02/go-simulator/world"
)

// hoursPerTide is the semidiurnal tidal period; by default the tide keeps that proportion to the day.
const hoursPerTide = 12.42

// ---------------------------------------------------------------- day cycle

type dayCycle struct{ secondsPerDay float64 }

func dayCycleRule() plugin.Definition {
	return plugin.Definition{
		Description: protocol.RuleDescription{
			Name:       "day_cycle",
			Summary:    "Turns the sun around the celestial pole; the stars follow it.",
			Parameters: []protocol.ParameterDescription{number("seconds_per_day", 120, "how long a whole day lasts")},
			Reads:      readsEnvironment, Writes: []string{"environment"},
			Publishes: []protocol.OutputDescription{plugin.Environment("sun_direction", "the direction of the sun")},
		},
		New: func(parameters plugin.Parameters) (plugin.Rule, error) {
			return &dayCycle{secondsPerDay: positive(parameters.Number("seconds_per_day", 120), 120)}, nil
		},
	}
}

func (cycle *dayCycle) Step(context *plugin.Context) error {
	environment := context.World.Environment
	if environment.SunDirection == nil || environment.CelestialPole == nil {
		return nil
	}
	turn := spatial.RotationAround(*environment.CelestialPole, 2*math.Pi*context.Step/cycle.secondsPerDay)
	sun := turn.ApplyToDirection(*environment.SunDirection).Normalize()
	context.SetEnvironment(&world.Environment{SunDirection: &sun})
	return nil
}

// ---------------------------------------------------------------- gusts

type gusts struct {
	variation, veering float64
	elapsed            float64
	base               *world.Wind
}

func gustsRule() plugin.Definition {
	return plugin.Definition{
		Description: protocol.RuleDescription{
			Name:    "gusts",
			Summary: "Varies the wind strength and direction around the wind the world started with.",
			Parameters: []protocol.ParameterDescription{
				number("variation", 0.5, "how much the strength changes (0..1)"),
				number("veering", 0.35, "how far the direction swings, in radians"),
			},
			Reads: readsEnvironment, Writes: []string{"environment"},
			Publishes: []protocol.OutputDescription{plugin.Environment("wind", "the wind and its gusts")},
		},
		New: func(parameters plugin.Parameters) (plugin.Rule, error) {
			return &gusts{variation: parameters.Number("variation", 0.5), veering: parameters.Number("veering", 0.35)}, nil
		},
	}
}

func (rule *gusts) Start(context *plugin.Context) error {
	if wind := context.World.Environment.Wind; wind != nil {
		initial := *wind
		rule.base = &initial
	}
	return nil
}

func (rule *gusts) Step(context *plugin.Context) error {
	if rule.base == nil {
		return nil
	}
	rule.elapsed += context.Step
	strength := float64(rule.base.Strength) * (1 + rule.variation*smoothWaves(rule.elapsed*0.23))
	angle := rule.veering * smoothWaves(rule.elapsed*0.07+11)
	direction := spatial.RotationY(angle).ApplyToDirection(rule.base.Direction).Normalize()
	context.SetEnvironment(&world.Environment{Wind: &world.Wind{Direction: direction, Strength: float32(math.Max(0, strength))}})
	return nil
}

// smoothWaves is a sum of incommensurate sines, between −1 and 1, that never quite repeats.
func smoothWaves(time float64) float64 {
	return (math.Sin(time) + 0.6*math.Sin(time*2.31+1.7) + 0.3*math.Sin(time*5.13+4.2)) / 1.9
}

// ---------------------------------------------------------------- clouds

type cloudDrift struct{ speed float64 }

func cloudDriftRule() plugin.Definition {
	return plugin.Definition{
		Description: protocol.RuleDescription{
			Name:       "cloud_drift",
			Summary:    "Moves the clouds with the wind.",
			Parameters: []protocol.ParameterDescription{number("speed", 0.05, "cloud-pattern units per second at wind strength 1")},
			Reads:      readsEnvironment, Writes: []string{"environment"},
			Publishes: []protocol.OutputDescription{plugin.Environment("cloud_offset", "how far the clouds have drifted")},
		},
		New: func(parameters plugin.Parameters) (plugin.Rule, error) {
			return &cloudDrift{speed: positive(parameters.Number("speed", 0.05), 0.05)}, nil
		},
	}
}

func (drift *cloudDrift) Step(context *plugin.Context) error {
	environment := context.World.Environment
	if environment.CloudOffset == nil || environment.Wind == nil {
		return nil
	}
	wind := environment.Wind
	distance := float32(drift.speed * (0.2 + float64(wind.Strength)) * context.Step)
	offset := environment.CloudOffset.Add(spatial.Vector{wind.Direction[0] * distance, 0, wind.Direction[2] * distance})
	context.SetEnvironment(&world.Environment{CloudOffset: &offset})
	return nil
}

// ---------------------------------------------------------------- tides

type tides struct{ periodSeconds float64 }

func tidesRule() plugin.Definition {
	return plugin.Definition{
		Description: protocol.RuleDescription{
			Name:    "tides",
			Summary: "Raises and lowers the water: turns the tide angle a full turn per period. Heights and phases come with the scene.",
			Parameters: []protocol.ParameterDescription{
				number("period_seconds", 0, "seconds per tidal cycle (0 = 12.42 h in proportion to seconds_per_day)"),
				number("seconds_per_day", 120, "day length used when period_seconds is 0"),
			},
			Reads: readsEnvironment, Writes: []string{"environment"},
			Publishes: []protocol.OutputDescription{plugin.Environment("tide_angle", "where the tide is in its cycle")},
		},
		New: func(parameters plugin.Parameters) (plugin.Rule, error) {
			dayLength := positive(parameters.Number("seconds_per_day", 120), 120)
			return &tides{periodSeconds: positive(parameters.Number("period_seconds", 0), dayLength*hoursPerTide/24)}, nil
		},
	}
}

func (rule *tides) Step(context *plugin.Context) error {
	current := context.World.Environment.TideAngle
	if current == nil {
		return nil
	}
	angle := float32(math.Mod(float64(*current)+2*math.Pi*context.Step/rule.periodSeconds, 2*math.Pi))
	context.SetEnvironment(&world.Environment{TideAngle: &angle})
	return nil
}

func positive(value, fallback float64) float64 {
	if value <= 0 {
		return fallback
	}
	return value
}
