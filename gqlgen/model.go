// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

type MemberInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type SkillInput struct {
	Name     string `json:"name"`
	Descs    string `json:"descs"`
	MemberID int64  `json:"member_id"`
}
