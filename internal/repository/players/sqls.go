package players

const sqlSelectPlayers = `
SELECT id, name, number, created_at, updated_at
FROM   players
WHERE  deleted_at IS NULL
`

const sqlSelectPlayerById = `
SELECT id, name, number, created_at, updated_at
FROM   players
WHERE  deleted_at IS NULL
AND    id = $1
`

const sqlInsertPlayer = `
INSERT INTO players(name, number, created_at, updated_at)
VALUES($1, $2, NOW(), NOW())
RETURNING id, name, number, created_at, updated_at
`
