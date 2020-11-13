package pg

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // required
)

// Repository is the application's data layer functionality.
type Repository interface {
	// Member queries
	CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error)
	DeleteMember(ctx context.Context, id int64) (Member, error)
	GetMember(ctx context.Context, id int64) (Member, error)
	ListMembers(ctx context.Context) ([]Member, error)
	UpdateMember(ctx context.Context, arg UpdateMemberParams) (Member, error)

	// Skill queries
	CreateSkill(ctx context.Context, arg CreateSkillParams) (Skill, error)
	DeleteSkill(ctx context.Context, id int64) (Skill, error)
	GetSkill(ctx context.Context, id int64) (Skill, error)
	ListSkills(ctx context.Context) ([]Skill, error)
	UpdateSkill(ctx context.Context, arg UpdateSkillParams) (Skill, error)
	ListSkillsByMemberID(ctx context.Context, memberID int64) ([]Skill, error)
}

type repoSvc struct {
	*Queries
	db *sql.DB
}

// NewRepository returns an implementation of the Repository interface.
func NewRepository(db *sql.DB) Repository {
	return &repoSvc{
		Queries: New(db),
		db:      db,
	}
}

// Open opens a database specified by the data source name.
// Format: host=foo port=5432 user=bar password=baz dbname=qux sslmode=disable
func Open(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}

// StringPtrToNullString converts *string to sql.NullString.
func StringPtrToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{}
}
