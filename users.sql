-- name: CreateUser :one
INSERT INTO public.users (name, dob)
VALUES ($1, $2)
RETURNING id, name, dob;

-- name: GetUserByID :one
SELECT id, name, dob FROM public.users WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, dob FROM public.users ORDER BY id;

-- name: UpdateUser :one
UPDATE public.users
SET name = $2, dob = $3
WHERE id = $1
RETURNING id, name, dob;

-- name: DeleteUser :exec
DELETE FROM public.users WHERE id = $1;
