-- name: GetMember :one
SELECT * FROM member
WHERE id = $1;

-- name: ListMembers :many
SELECT * FROM member
ORDER BY name;

-- name: CreateMember :one
INSERT INTO member (name, age)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateMember :one
UPDATE member
SET name = $2, age = $3
WHERE id = $1
RETURNING *;

-- name: DeleteMember :one
DELETE FROM member
WHERE id = $1
RETURNING *;

-- name: GetSkill :one
SELECT * FROM skill
WHERE id = $1;

-- name: ListSkills :many
SELECT * FROM skill
ORDER BY name;

-- name: CreateSkill :one
INSERT INTO skill (name, descs, member_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateSkill :one
UPDATE skill
SET name = $2, descs = $3, member_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteSkill :one
DELETE FROM skill
WHERE id = $1
RETURNING *;

-- name: ListSkillsByMemberID :many
SELECT skill.* FROM skill, member
WHERE member.id = skill.member_id AND skill.member_id = $1;