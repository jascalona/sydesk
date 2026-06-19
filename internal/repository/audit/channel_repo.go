package audit

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/audit"
)

type channelRepo struct {
	DB *sql.DB
}

func NewChannelRepo(db *sql.DB) audit.InterfaceChannel {
	return &channelRepo{DB: db}
}

func (r *channelRepo) GetAll(ctx context.Context) ([]*audit.ChannelRep, error) {
	query := `SELECT id, name, created_at FROM channel_reception ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Println("error al correr el query context", err.Error())
		return nil, err
	}

	defer rows.Close()
	channel := make([]*audit.ChannelRep, 0)

	for rows.Next() {
		cp := &audit.ChannelRep{}
		err := rows.Scan(
			&cp.ID,
			&cp.NAME,
			&cp.CREATEDAT,
		)
		if err != nil {
			return nil, err
		}

		channel = append(channel, cp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return channel, nil
}
