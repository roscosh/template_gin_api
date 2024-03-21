package sql

type SQL struct {
	Users *UsersSQL
}

func (s *SQL) Close() {
	s.Users.baseSQL.pool.Close()
}

func NewSQL(dsn string) (*SQL, error) {
	pool, err := newPostgresPool(dsn)
	if err != nil {
		return nil, err
	}
	return &SQL{
		Users: NewUsersSQL(pool),
	}, nil
}
