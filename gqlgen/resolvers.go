package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/hunght/gqlgen-sqlc-example/pg"
)

// Resolver connects individual resolvers with the datalayer.
type Resolver struct {
	Repository pg.Repository
}

func (r *memberResolver) Skills(ctx context.Context, obj *pg.Member) ([]pg.Skill, error) {
	return r.Repository.ListSkillsByMemberID(ctx, obj.ID)
}

func (r *mutationResolver) CreateMember(ctx context.Context, data MemberInput) (*pg.Member, error) {
	member, err := r.Repository.CreateMember(ctx, pg.CreateMemberParams{
		Name: data.Name,
		Age:  int32(data.Age),
	})
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *mutationResolver) UpdateMember(ctx context.Context, id int64, data MemberInput) (*pg.Member, error) {
	member, err := r.Repository.UpdateMember(ctx, pg.UpdateMemberParams{
		ID:   id,
		Name: data.Name,
		Age:  int32(data.Age),
	})
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *mutationResolver) DeleteMember(ctx context.Context, id int64) (*pg.Member, error) {
	member, err := r.Repository.DeleteMember(ctx, id)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *mutationResolver) CreateSkill(ctx context.Context, data SkillInput) (*pg.Skill, error) {
	skill, err := r.Repository.CreateSkill(ctx, pg.CreateSkillParams{
		Name:     data.Name,
		Descs:    data.Descs,
		MemberID: data.MemberID,
	})
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *mutationResolver) UpdateSkill(ctx context.Context, id int64, data SkillInput) (*pg.Skill, error) {
	skill, err := r.Repository.UpdateSkill(ctx, pg.UpdateSkillParams{
		ID:       id,
		Name:     data.Name,
		Descs:    data.Descs,
		MemberID: data.MemberID,
	})
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *mutationResolver) DeleteSkill(ctx context.Context, id int64) (*pg.Skill, error) {
	skill, err := r.Repository.DeleteSkill(ctx, id)
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *queryResolver) Member(ctx context.Context, id int64) (*pg.Member, error) {
	member, err := r.Repository.GetMember(ctx, id)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *queryResolver) Members(ctx context.Context) ([]pg.Member, error) {
	return r.Repository.ListMembers(ctx)
}

func (r *queryResolver) Skill(ctx context.Context, id int64) (*pg.Skill, error) {
	skill, err := r.Repository.GetSkill(ctx, id)
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *queryResolver) Skills(ctx context.Context) ([]pg.Skill, error) {
	return r.Repository.ListSkills(ctx)
}

// Member returns MemberResolver implementation.
func (r *Resolver) Member() MemberResolver { return &memberResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type memberResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
