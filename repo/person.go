package person_repo

import (
	"context"
	"game-management/db"
	"game-management/model"
)

func GetPersons() (rs []*model.Person, err error) {
	rows, err := db.GetStorage().Conn().Query(context.TODO(), "select * from person")
	if err != nil {
		return rs, err
	}
	defer rows.Close()
	for rows.Next() {
		var person model.Person
		if err := rows.Scan(&person.Name); err != nil {
			return rs, err
		}
		rs = append(rs, &person)
	}
	if err := rows.Err(); err != nil {
		return rs, err
	}
	return rs, nil
}
