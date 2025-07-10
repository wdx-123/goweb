package dao

import "log"

func DeleteUserByID(userID int) error {
	DBmutex.Lock()
	defer DBmutex.Unlock()
	query := `delete from user where id= ?`
	_, err := DB.Exec(query, userID)
	if err != nil {
		log.Println(err)
	}
	return err
}
