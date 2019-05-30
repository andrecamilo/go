// Repository pattern example - repository interface
package repo
type Project struct {
    ID   int64
    Name string
}
type IRepository interface {
    FindAll() ([]Project, error)
    Store(Project) (Project, error)
    Delete(Project) error
}

// Repository pattern example - in-memory implementation
package inmem
type ProjectRepository struct {
    projects  map[int64]string
    highestID int64
}
