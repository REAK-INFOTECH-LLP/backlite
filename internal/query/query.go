package query

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed schema.sql
var Schema string

const InsertTask = `
    INSERT INTO backlite_tasks 
        (id, created_at, queue, task, wait_until)
    VALUES ($1, $2, $3, $4, $5)
`

const SelectScheduledTasks = `
    SELECT 
        id, queue, task, attempts, wait_until, created_at, last_executed_at, NULL
    FROM 
        backlite_tasks
    WHERE
        claimed_at IS NULL
        OR claimed_at < $1
    ORDER BY
        wait_until ASC,
        id ASC
    LIMIT $2
    OFFSET $3
`

const DeleteTask = `
    DELETE FROM backlite_tasks
    WHERE id = $1
`

const InsertCompletedTask = `
    INSERT INTO backlite_tasks_completed
        (id, created_at, queue, last_executed_at, attempts, last_duration_micro, succeeded, task, expires_at, error)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

const TaskFailed = `
    UPDATE backlite_tasks
    SET 
        claimed_at = NULL, 
        wait_until = $1,
        last_executed_at = $2
    WHERE id = $3
`

const DeleteExpiredCompletedTasks = `
    DELETE FROM backlite_tasks_completed
    WHERE
        expires_at IS NOT NULL
        AND expires_at <= $1
`

const SelectTaskStatus = `
    SELECT
        claimed_at IS NOT NULL AS running,
        NULL AS success
    FROM backlite_tasks
    WHERE id = $1
    UNION ALL
    SELECT
        FALSE AS running,
        error IS NULL AS success
    FROM backlite_tasks_completed
    WHERE id = $2
`

func ClaimTasks(count int) string {
	const query = `
        UPDATE backlite_tasks
        SET
            claimed_at = $1,
            attempts = attempts + 1
        WHERE id IN (%s)
    `

	// PostgreSQL uses $2, $3, ... for parameter placeholders
	params := make([]string, count)
	for i := 0; i < count; i++ {
		params[i] = fmt.Sprintf("$%d", i+2)
	}
	paramStr := strings.Join(params, ",")
	return fmt.Sprintf(query, paramStr)
}
