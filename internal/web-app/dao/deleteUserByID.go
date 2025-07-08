package dao

func DeleteUserByID(userID int) error {
	DBmutex.Lock()
	defer DBmutex.Unlock()
	query := `delete from user where id= ?`
	_, err := DB.Exec(query, userID)
	return err
}
