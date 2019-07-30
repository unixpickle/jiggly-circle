package main

type Particle struct {
	Position Vector2
	Velocity Vector2
}

type Scene struct {
	Particles []*Particle
	Radius    float64
}

func NewScene(numParticles int, radius float64) *Scene {
	res := &Scene{
		Particles: make([]*Particle, numParticles),
		Radius:    radius,
	}
	for i := range res.Particles {
		res.Particles[i] = &Particle{
			Position: RandomCircleVector2().Scale(radius),
			Velocity: RandomCircleVector2().Scale(radius),
		}
	}
	return res
}

func (p *Scene) Step(t float64) {
	for _, particle := range p.Particles {
		particle.Position = particle.Position.Add(particle.Velocity.Scale(t))
		norm := particle.Position.Norm()
		if norm > p.Radius {
			// Prevent the particle from escaping, since this is
			// not an exact simulation and it might be past the
			// outskirts of the scene.
			particle.Position = particle.Position.Scale(p.Radius / norm)

			// Bounce the particle against the wall.
			reflectDir := particle.Position.Scale(1 / p.Radius)
			tangent := reflectDir.OrthogonalVector()
			particle.Velocity = reflectDir.Scale(-reflectDir.Dot(particle.Velocity)).Add(
				tangent.Scale(tangent.Dot(particle.Velocity)),
			)
		}
	}
}
