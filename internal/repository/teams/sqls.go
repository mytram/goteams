package teams

const sqlSelectTeams = `
SELECT t.id, t.name, t.created_at, t.updated_at,
       COALESCE(
         JSON_AGG(JSON_BUILD_OBJECT(
           'id',        p.id,
           'name',      p.name,
           'createdAt', p.created_at,
           'updatedAt', p.updated_at
         )) FILTER (WHERE p.id IS NOT NULL),
         NULL
       )
FROM   teams t
LEFT
JOIN   team_players tp ON (t.id = tp.team_id)
LEFT
JOIN   players p ON (p.id = tp.player_id)
WHERE  t.deleted_at IS NULL
GROUP BY t.id
`

const sqlSelectTeamById = `
SELECT t.id, t.name, t.created_at, t.updated_at,
       COALESCE(
         JSON_AGG(JSON_BUILD_OBJECT(
           'id',        p.id,
           'name',      p.name,
           'createdAt', p.created_at,
           'updatedAt', p.updated_at
         )) FILTER (WHERE p.id IS NOT NULL),
         NULL
       )
FROM   teams t
LEFT
JOIN   team_players tp ON (t.id = tp.team_id)
LEFT
JOIN   players p ON (p.id = tp.player_id)
WHERE  t.deleted_at IS NULL
AND    t.id = $1
GROUP BY t.id
`

const sqlInsertTeam = `
INSERT INTO teams(name, created_at, updated_at)
VALUES($1, NOW(), NOW())
RETURNING id, name, created_at, updated_at, NULL
`

const sqlInsertTeamPlayer = `
INSERT INTO team_players(team_id, player_id)
VALUES($1, $2) ON CONFLICT DO NOTHING
`

const sqlDeleteTeamPlayers = `
DELETE FROM team_players WHERE team_id = $1
`

const sqlUpdateTeamPlayer = `
UPDATE teams
SET name = $2, updated_at = NOW()
WHERE id = $1
`
