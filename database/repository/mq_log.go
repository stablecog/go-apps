package repository

import (
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/ent/mqlog"
)

func (r *Repository) GetQueuePosition(messageId uuid.UUID) (position int, total int, err error) {
	query := `
	SELECT row_num, total
	FROM (
			SELECT message_id,
						 ROW_NUMBER() OVER (ORDER BY priority DESC, created_at ASC) as row_num,
						 COUNT(*) OVER () as total
			FROM mq_log
	) AS subquery
	WHERE message_id = $1
	`

	rows, err := r.DB.QueryContext(r.Ctx, query, messageId)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&position, &total)
		if err != nil {
			return 0, 0, err
		}
	} else {
		// No rows for the given message ID means not in the queue
		// For the total, we need another query
		count, err := r.DB.MqLog.Query().Count(r.Ctx)
		if err != nil {
			return 0, 0, err
		}
		return 0, count, nil
	}

	// Check for any errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return 0, 0, err
	}

	return position, total, nil
}

// Add to queue log
func (r *Repository) AddToQueueLog(messageId uuid.UUID, priority int, DB *ent.Client) (*ent.MqLog, error) {
	if DB == nil {
		DB = r.DB
	}
	return DB.MqLog.Create().SetMessageID(messageId).SetPriority(priority).Save(r.Ctx)
}

// Delete from queue log
func (r *Repository) DeleteFromQueueLog(messageId uuid.UUID, DB *ent.Client) (int, error) {
	if DB == nil {
		DB = r.DB
	}
	return DB.MqLog.Delete().Where(mqlog.MessageIDEQ(messageId)).Exec(r.Ctx)
}

// Set is_processing
func (r *Repository) SetIsProcessingInQueueLog(messageId uuid.UUID, isProcessing bool, DB *ent.Client) (int, error) {
	if DB == nil {
		DB = r.DB
	}
	return DB.MqLog.Update().Where(mqlog.MessageIDEQ(messageId)).SetIsProcessing(isProcessing).Save(r.Ctx)
}
