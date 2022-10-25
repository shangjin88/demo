package main

import "fmt"

// RunnerOptions defines the options supported during the execution of a
// kubeadm composable workflows
type RunnerOptions struct {
	// SkipPhases defines the list of phases to be excluded by execution (if empty, none).
	SkipPhases []string
}

type Phase struct {
	// name of the phase.
	// Phase name should be unique among peer phases (phases belonging to
	// the same workflow or phases belonging to the same parent phase).
	Name string
}

type Runner struct {
	// Options that regulate the runner behavior.
	Options RunnerOptions

	// Phases composing the workflow to be managed by the runner.
	Phases []Phase
}

func NewRunner() *Runner {
	return &Runner{
		Phases: []Phase{},
	}
}

func (r *Runner) initAllPhase() {
	r.Phases = append(r.Phases, Phase{Name: "check"}, Phase{Name: "install"}, Phase{Name: "output"})
}

func (r *Runner) computePhaseRunFlags() map[string]bool {
	phaseRunFlags := map[string]bool{}

	for _, p := range r.Phases {
		phaseRunFlags[p.Name] = true
	}

	for _, skipPhase := range r.Options.SkipPhases {
		phaseRunFlags[skipPhase] = false
	}

	return phaseRunFlags
}

func (r *Runner) run() {
	r.initAllPhase()
	phaseRunFlags := r.computePhaseRunFlags()

	for _, phase := range r.Phases {
		if run, ok := phaseRunFlags[phase.Name]; !run || !ok {
			continue
		}
		fmt.Printf("[%v] Starting...\n", phase)
	}

}

func main() {
	r := NewRunner()
	r.Options.SkipPhases = append(r.Options.SkipPhases, "install")
	r.run()

}
