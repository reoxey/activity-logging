package mysql

import (
	"context"
	"strings"

	"activity_service/activity"
	proto "activity_service/proto/activity"
)

func (m mysqlRepo) Create(ctx context.Context, a *proto.Activity) (activityId int, err error) {

	var s []string
	s = append(s, "user_id = ?")
	s = append(s, "type = ?")
	s = append(s, "timestamp = ?")
	s = append(s, "duration = ?")

	rows, err := m.db.ExecContext(ctx, "INSERT "+m.table+" SET "+strings.Join(s, ","), a.User.Id, a.ActivityType, a.Timestamp, a.Duration)

	if err != nil {
		return 0, err
	}

	if n, _ := rows.RowsAffected(); n == 0 {
		return 0, NoRowsAffected
	}

	id, _ := rows.LastInsertId()

	return int(id), nil
}

func (m mysqlRepo) ById(ctx context.Context, id int) (a *proto.Activity, err error) {

	a = &proto.Activity{}

	a.User = &proto.User{}

	err = m.db.QueryRowContext(ctx, "SELECT activity_id, user_id, type, timestamp, duration " +
		"FROM "+m.table+" WHERE activity_id = ?", id).
		Scan(&a.Id, &a.User.Id, &a.ActivityType, &a.Timestamp, &a.Duration)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (m mysqlRepo) AllByUser(ctx context.Context, userId int) (activities activity.Activities, err error) {

	rows, err := m.db.QueryContext(ctx, "SELECT activity_id, user_id, type, timestamp, duration " +
		"FROM "+m.table+" WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		a := &proto.Activity{}

		a.User = &proto.User{}

		if err = rows.Scan(&a.Id, &a.User.Id, &a.ActivityType, &a.Timestamp, &a.Duration); err != nil {
			return activities, err
		}

		activities = append(activities, a)
	}

	return activities, nil
}
