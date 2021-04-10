package mysql

import (
	"context"
	"strings"

	proto "user_service/proto/user"
	"user_service/user"
)

func (m *mysqlRepo) Create(ctx context.Context, user *proto.User) (userId int, err error) {

	var s []string
	s = append(s, "name = ?")
	s = append(s, "email = ?")
	s = append(s, "phone = ?")

	rows, err := m.db.ExecContext(ctx, "INSERT "+m.table+" SET "+strings.Join(s, ","), user.Name, user.Email, user.Phone)

	if err != nil {
		return 0, err
	}

	if n, _ := rows.RowsAffected(); n == 0 {
		return 0, NoRowsAffected
	}

	id, _ := rows.LastInsertId()

	return int(id), nil
}

func (m *mysqlRepo) ById(ctx context.Context, id int) (user *proto.User, err error) {

	user = &proto.User{}

	err = m.db.QueryRowContext(ctx, "SELECT user_id, name, email, phone " +
		"FROM "+m.table+" WHERE user_id = ?", id).
		Scan(&user.Id, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *mysqlRepo) All(ctx context.Context) (users user.Users, err error) {

	rows, err := m.db.QueryContext(ctx, "SELECT user_id, name, email, phone " +
		"FROM "+m.table+" WHERE 1")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &proto.User{}

		if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}
