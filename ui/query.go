package ui

const selectTask = `
	SELECT 
	    id,
	    queue,
	    task,
	    attempts,
	    wait_until,
	    created_at,
	    last_executed_at,
	    claimed_at
	FROM 
	    backlite_tasks
	WHERE
	    id = $1
`

const selectCompletedTask = `
	SELECT
	    id,
		created_at,
		queue,
		last_executed_at,
		attempts,
		last_duration_micro,
		succeeded,
		task,
		expires_at,
		error
	FROM
	    backlite_tasks_completed 
	WHERE
	    id = $1
`

const selectRunningTasks = `
	SELECT 
	    id,
	    queue,
	    null,
	    attempts,
	    wait_until,
	    created_at,
	    last_executed_at,
	    claimed_at
	FROM 
	    backlite_tasks
	WHERE
	    claimed_at IS NOT NULL
	LIMIT $1
	OFFSET $2 
`
const selectCompletedTasks = `
	SELECT
	    id,
		created_at,
		queue text,
		last_executed_at,
		attempts,
		last_duration_micro,
		succeeded,
		null,
		expires_at,
		error
	FROM
	    backlite_tasks_completed 
	WHERE
	    succeeded = $1
	ORDER BY
	    last_executed_at DESC
	LIMIT $2
	OFFSET $3
`
