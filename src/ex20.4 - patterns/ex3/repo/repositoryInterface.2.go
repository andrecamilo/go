package main

// Repository pattern example - usage in main function
func main() {
	projectsRepo := inmem.ProjectRepository{}
	newProject, _ := projectsRepo.Store(
		project.Project{
			Name: "First project",
		},
	)
	newProject.Name = "First project - updated!"
	_, _ = projectsRepo.Store(newProject)
}
