package motion

import (
	"math"

	"github.com/akrck02/go-simulator/plugin"
	"github.com/akrck02/go-simulator/protocol"
	"github.com/akrck02/go-simulator/spatial"
	"github.com/akrck02/go-simulator/world"
)

// The properties an entity sets to ask for a motion (glTF node extras arrive as properties):
//
//	"bob":  {"height": h, "period": seconds, "roll": radians}
//	"spin": {"axis": [x, y, z], "speed": radians per second}
//	"path": {"points": [[x, y, z], ...], "speed": units per second, "loop": true}
//
// An entity takes one motion; when it asks for several, bob wins over spin and spin over path.
var motionProperties = []string{"bob", "spin", "path"}

// motion places an entity, given its rest transform and the seconds since it started moving.
type motion interface {
	transformAt(rest spatial.Matrix, elapsed float64) spatial.Matrix
}

type movingEntity struct {
	rest    spatial.Matrix
	motion  motion
	elapsed float64
}

// entityMotion moves every entity that asks for a motion, including the ones added later.
type entityMotion struct {
	moving map[world.Identifier]*movingEntity
}

func entityMotionRule() plugin.Definition {
	return plugin.Definition{
		Description: protocol.RuleDescription{
			Name:    "entity_motion",
			Summary: "Moves the entities that ask for it in their properties: bob (float and roll), spin (turn on an axis) and path (follow points).",
			// it moves entities other plugins or the scene made, so it publishes nothing of its own
			Reads: world.Interest{Properties: motionProperties},
		},
		New: func(plugin.Parameters) (plugin.Rule, error) {
			return &entityMotion{moving: map[world.Identifier]*movingEntity{}}, nil
		},
	}
}

func (rule *entityMotion) Step(context *plugin.Context) error {
	rule.followTheWorld(context.World)
	for _, entity := range context.World.Entities() {
		moving, found := rule.moving[entity.Identifier]
		if !found {
			continue
		}
		moving.elapsed += context.Step
		transform := moving.motion.transformAt(moving.rest, moving.elapsed)
		context.Modify(&world.Entity{Identifier: entity.Identifier, Transform: &transform})
	}
	return nil
}

// followTheWorld starts moving new entities that ask for a motion and forgets removed ones.
func (rule *entityMotion) followTheWorld(current *world.World) {
	for identifier := range rule.moving {
		if current.Entity(identifier) == nil {
			delete(rule.moving, identifier)
		}
	}
	for _, entity := range current.Entities() {
		if _, known := rule.moving[entity.Identifier]; known {
			continue
		}
		if chosen := motionOf(entity); chosen != nil {
			rule.moving[entity.Identifier] = &movingEntity{rest: restOf(entity), motion: chosen}
		}
	}
}

func restOf(entity *world.Entity) spatial.Matrix {
	if entity.Transform == nil || entity.Transform.IsZero() {
		return spatial.Identity()
	}
	return *entity.Transform
}

func motionOf(entity *world.Entity) motion {
	if values, asks := entity.Properties["bob"].(map[string]any); asks {
		return newBobbing(plugin.Parameters(values), restOf(entity))
	}
	if values, asks := entity.Properties["spin"].(map[string]any); asks {
		return newSpin(plugin.Parameters(values))
	}
	if values, asks := entity.Properties["path"].(map[string]any); asks {
		return newPath(plugin.Parameters(values))
	}
	return nil
}

// ---------------------------------------------------------------- bob

type bobbing struct {
	height, period, roll, phase float64
}

func newBobbing(values plugin.Parameters, rest spatial.Matrix) *bobbing {
	return &bobbing{height: values.Number("height", 0.001), period: positive(values.Number("period", 4), 4),
		roll: values.Number("roll", 0.03), phase: positionPhase(rest)}
}

func (motion *bobbing) transformAt(rest spatial.Matrix, elapsed float64) spatial.Matrix {
	wave := 2*math.Pi*elapsed/motion.period + motion.phase
	lift := spatial.Translation(spatial.Vector{0, float32(motion.height * math.Sin(wave)), 0})
	roll := spatial.RotationX(motion.roll * math.Sin(wave*0.7+1))
	return lift.Multiply(rest).Multiply(roll)
}

// positionPhase derives a phase from where an entity is, so neighbours do not move together.
func positionPhase(transform spatial.Matrix) float64 {
	position := transform.Position()
	return math.Mod(float64(position[0])*12.9898+float64(position[2])*78.233, 2*math.Pi)
}

// ---------------------------------------------------------------- spin

type spin struct {
	axis  spatial.Vector
	speed float64
}

func newSpin(values plugin.Parameters) *spin {
	return &spin{axis: values.Vector("axis", spatial.Vector{0, 1, 0}).Normalize(), speed: values.Number("speed", 1)}
}

func (motion *spin) transformAt(rest spatial.Matrix, elapsed float64) spatial.Matrix {
	return rest.Multiply(spatial.RotationAround(motion.axis, motion.speed*elapsed))
}

// ---------------------------------------------------------------- path

// path follows waypoints at a constant speed, facing where it goes (the model's +Z forward).
// Without a loop it turns around at the last point.
type path struct {
	route   []spatial.Vector
	lengths []float64
	total   float64
	speed   float64
}

func newPath(values plugin.Parameters) *path {
	points := pointsOf(values["points"])
	route := append([]spatial.Vector(nil), points...)
	if values.Boolean("loop", true) && len(points) > 1 {
		route = append(route, points[0])
	} else {
		for index := len(points) - 2; index >= 0; index-- {
			route = append(route, points[index])
		}
	}
	motion := &path{route: route, speed: values.Number("speed", 1)}
	for index := 1; index < len(route); index++ {
		length := float64(route[index].Subtract(route[index-1]).Length())
		motion.lengths = append(motion.lengths, length)
		motion.total += length
	}
	return motion
}

func pointsOf(value any) []spatial.Vector {
	list, _ := value.([]any)
	points := make([]spatial.Vector, 0, len(list))
	for _, entry := range list {
		points = append(points, plugin.VectorFrom(entry, spatial.Vector{}))
	}
	return points
}

func (motion *path) transformAt(rest spatial.Matrix, elapsed float64) spatial.Matrix {
	if motion.total == 0 {
		return rest
	}
	position, heading := motion.pointAt(math.Mod(motion.speed*elapsed, motion.total))
	return spatial.Translation(position).Multiply(spatial.RotationY(heading)).Multiply(rest.WithoutTranslation())
}

// pointAt returns the position at a distance along the route and the heading around Y.
func (motion *path) pointAt(distance float64) (spatial.Vector, float64) {
	for index, length := range motion.lengths {
		if distance <= length || index == len(motion.lengths)-1 {
			from, to := motion.route[index], motion.route[index+1]
			amount := 0.0
			if length > 0 {
				amount = distance / length
			}
			position := from.Add(to.Subtract(from).Scale(float32(amount)))
			return position, math.Atan2(float64(to[0]-from[0]), float64(to[2]-from[2]))
		}
		distance -= length
	}
	return motion.route[0], 0
}
