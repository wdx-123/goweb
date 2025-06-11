package dao

func DeleteUserByID(userID int) error {
	query := `delete from user where id= ?`
	_, err := DB.Exec(query, userID)
	return err
}
