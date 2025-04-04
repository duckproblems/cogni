package ai

type State string

const (
	IDLE        State = "Idle"
	CHASING     State = "Chasing"
	PATROLLING  State = "Patrolling"
	UNCONSCIOUS State = "Unconscious"
)
