package db

type User struct {
	UserID                    int
	StravaUserID              float64
	AccessToken, RefreshToken string
	ExpiresTo                 int64
}

// SaveUser inserts a user to database
func SaveUser(user *User) error {
	dbase := NewDatabase()
	args := QueryArgs{user.AccessToken, user.RefreshToken, user.ExpiresTo}
	row, err := dbase.ExecCRUD(insertUserQuery(), args)
	if err != nil {
		return err
	}

	row.Scan(&user.UserID)
	return nil
}

func insertUserQuery() string {
	return "insert into \"strava-gate\".\"users\"(access_token, refresh_token, expires_to) " +
		"values($1, $2, $3) " +
		"returning user_id"
}
