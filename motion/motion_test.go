package motion

import (
	"math"
	"testing"

	"github.com/akrck02/go-simulator/configuration"
	"github.com/akrck02/go-simulator/host"
	"github.com/akrck02/go-simulator/spatial"
	"github.com/akrck02/go-simulator/world"
)

// running starts the motion plugin inside the test with a rules file and an initial world.
func running(t *testing.T, rules string, initial *world.World) *host.Engine {
	t.Helper()
	settings, err := configuration.Parse([]byte(rules))
	if err != nil {
		t.Fatal(err)
	}
	engine, err := host.Build(settings, initial, map[string]host.Builtin{"motion": Plugin().ServeConnection})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = engine.Close() })
	if _, err := engine.Start(); err != nil {
		t.Fatal(err)
	}
	return engine
}

func rulesWith(rules string) string {
	return "plugins: [{name: motion, builtin: true}]\nstages:\n  - name: motion\n    clock: realtime\n    rules:\n" + rules
}

func skyWorld() *world.World {
	initial := world.New()
	sun, pole, offset, angle := spatial.Vector{0.6, 0.8, 0}, spatial.Vector{0, 0.7, -0.71}, spatial.Vector{}, float32(0)
	initial.Environment = world.Environment{SunDirection: &sun, CelestialPole: &pole, CloudOffset: &offset, TideAngle: &angle,
		Wind: &world.Wind{Direction: spatial.Vector{1, 0, 0}, Strength: 1}}
	return initial
}

func TestDayCycleTurnsTheSun(t *testing.T) {
	engine := running(t, rulesWith("      - {rule: day_cycle, with: {seconds_per_day: 100}}"), skyWorld())
	before := *engine.World.Environment.SunDirection
	if _, err := engine.Advance(50); err != nil {
		t.Fatal(err)
	}
	if after := *engine.World.Environment.SunDirection; before.Dot(after) > 0.99 {
		t.Fatalf("half a day should move the sun")
	}
}

func TestCloudsDriftWithTheWind(t *testing.T) {
	engine := running(t, rulesWith("      - {rule: cloud_drift, with: {speed: 1}}"), skyWorld())
	_, _ = engine.Advance(2)
	if offset := *engine.World.Environment.CloudOffset; offset[0] <= 0 || offset[2] != 0 {
		t.Fatalf("clouds should drift east, offset %v", offset)
	}
}

func TestTidesTurnAndCanBeSwitchedOff(t *testing.T) {
	engine := running(t, rulesWith("      - {rule: tides, with: {period_seconds: 40}, toggle_key: M}"), skyWorld())
	_, _ = engine.Advance(10)
	if angle := *engine.World.Environment.TideAngle; math.Abs(float64(angle)-math.Pi/2) > 1e-5 {
		t.Fatalf("a quarter of the period should turn a quarter: %v", angle)
	}
	engine.Toggle("M")
	_, _ = engine.Advance(10)
	if angle := *engine.World.Environment.TideAngle; math.Abs(float64(angle)-math.Pi/2) > 1e-5 {
		t.Fatalf("switched-off tides should not move")
	}
}

func TestGustsChangeTheWindAroundItsStart(t *testing.T) {
	engine := running(t, rulesWith("      - {rule: gusts, with: {variation: 0.5, veering: 0.35}}"), skyWorld())
	_, _ = engine.Advance(3)
	wind := engine.World.Environment.Wind
	if wind.Strength == 1 || wind.Strength < 0.4 || wind.Strength > 1.6 {
		t.Fatalf("the strength should vary within ±50%%: %v", wind.Strength)
	}
}

func TestEntitiesFollowTheirPathAndLaterOnesMoveToo(t *testing.T) {
	initial := world.New()
	path := map[string]any{"points": []any{[]any{0.0, 0.0, 0.0}, []any{10.0, 0.0, 0.0}}, "speed": 1.0, "loop": false}
	_ = initial.ApplyChange(world.AddEntity(&world.Entity{Identifier: "ferry", Properties: map[string]any{"path": path}}))
	engine := running(t, rulesWith("      - {rule: entity_motion}"), initial)
	_, _ = engine.Advance(5)
	if position := engine.World.Entity("ferry").Transform.Position(); math.Abs(float64(position[0]-5)) > 1e-4 {
		t.Fatalf("after 5 s the ferry should be halfway, got %v", position)
	}
	_, _ = engine.Advance(10)
	if position := engine.World.Entity("ferry").Transform.Position(); math.Abs(float64(position[0]-5)) > 1e-4 {
		t.Fatalf("going back, after 15 s the ferry should be halfway again, got %v", position)
	}
}

func TestBobbingWinsOverOtherMotions(t *testing.T) {
	both := &world.Entity{Identifier: "boat", Properties: map[string]any{
		"bob":  map[string]any{"height": 0.5},
		"spin": map[string]any{"speed": 1.0},
	}}
	if _, isBobbing := motionOf(both).(*bobbing); !isBobbing {
		t.Fatalf("bob should win over spin")
	}
	if motionOf(&world.Entity{Identifier: "rock"}) != nil {
		t.Fatalf("an entity that asks for nothing does not move")
	}
}
