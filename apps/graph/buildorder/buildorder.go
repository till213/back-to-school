package main

import (
	"fmt"
	"regexp"

	"github.com/golang-collections/collections/stack"
)

type color int

const (
	white color = iota
	gray
	black
)

type project struct {
	id           string
	dependencies []*project
	mark         color
}

func dfs(p *project, sorted *stack.Stack) (acyclic bool) {
	if p.mark == black {
		// Already completely dealt with
		return true
	} else if p.mark == gray {
		// Graph is not a DAG -> cycle detected
		// Push the vertex which closes the circle
		sorted.Push(p)
		return false
	}

	p.mark = gray
	acyclic = true
	for _, dep := range p.dependencies {
		// Recursively visit all dependent projects first ("depth first")
		acyclic = dfs(dep, sorted)
		if !acyclic {
			break
		}
	}
	p.mark = black
	sorted.Push(p)
	return acyclic
}

func buildOrder(projects []*project) (orderedProjects []*project, acyclic bool) {

	sorted := stack.New()

	for _, p := range projects {
		p.mark = white
	}
	for _, p := range projects {
		if p.mark != black {
			acyclic = dfs(p, sorted)
			if !acyclic {
				break
			}
		}
	}
	orderedProjects = make([]*project, 0, sorted.Len())
	l := sorted.Len()
	for i := 0; i < l; i++ {
		orderedProjects = append(orderedProjects, sorted.Pop().(*project))
	}
	return orderedProjects, acyclic
}

func parseDependency(dependency string) (project, dependentOnProject string) {
	// (A, B)
	// this is not the most efficient way to parse a simple string like "(A, B)",
	// but regular expresssions are cool ;)
	re := regexp.MustCompile(`\((.*?), (.*?)\)`)
	match := re.FindStringSubmatch(dependency)
	project = match[1]
	dependentOnProject = match[2]
	return
}

func findProject(projects []*project, id string) *project {
	for _, p := range projects {
		if p.id == id {
			return p
		}
	}
	return nil
}

func createProjects(projectIDs, dependencies []string) []*project {
	projects := make([]*project, len(projectIDs))

	// projects
	for i, id := range projectIDs {
		projects[i] = new(project)
		projects[i].id = id
		projects[i].dependencies = make([]*project, 0, 4)
	}
	// dependencies
	for _, d := range dependencies {
		pID, depID := parseDependency(d)
		project := findProject(projects, pID)
		depProject := findProject(projects, depID)
		project.dependencies = append(project.dependencies, depProject)
	}
	return projects
}

func main() {
	projectIDs := []string{"A", "B", "C", "D", "E", "F"}
	// (A dependent on B)
	dependencies := []string{"(A, D)", "(F, B)", "(B, D)", "(F, A)", "(D, C)"}

	projects := createProjects(projectIDs, dependencies)
	orderedProjects, acyclic := buildOrder(projects)

	if acyclic {
		fmt.Println("Build order")
		for _, v := range orderedProjects {
			fmt.Printf("%s, ", v.id)
		}
		fmt.Println()
	} else {
		fmt.Println("Mutual dependency exists.")
	}
}
