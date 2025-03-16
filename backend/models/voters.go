package models

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB

type Voter struct {
	ID     string `json:"id"`
	VName  string `json:"VNAME"`
	Choice int    `json:"choice"`
}

func GetVoterByID(vid string) (Voter, error) {
	var voter Voter
	err := DB.QueryRow("SELECT voter from voters WHERE vid = ?", vid).Scan(&voter.ID, &voter.VName, &voter.Choice)
	if err != sql.ErrNoRows {
		fmt.Println("No such voter")
		return Voter{}, err
	}

	return voter, nil
}

func QueryVoters() ([]Voter, error) {
	rows, err := DB.Query("SELECT * FROM voters")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	voters := make([]Voter, 0)

	for rows.Next() {
		singleVoter := Voter{}
		err = rows.Scan(&singleVoter.ID, &singleVoter.VName, &singleVoter.Choice)

		if err != nil {
			return nil, err
		}

		voters = append(voters, singleVoter)

	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return voters, nil
}

func AddVoter(newVoter *Voter) (int64, error) {
	sql := `INSERT INTO voters (vid, vname, vote)
            VALUES (?, ?, ?)`
	res, err := DB.Exec(sql, newVoter.ID, newVoter.VName, newVoter.Choice)
	fmt.Println(res)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func DeleteVoter(VID *string) (int64, error) {
	sql := `DELETE FROM voters WHERE vid == ?`

	res, err := DB.Exec(sql, VID)
	fmt.Println(res)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func main() {
	// Connect to the SQLite database
	err := ConnectDB()
	fmt.Println("Connected to the SQLite database successfully.")

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = CreateTableOfVoters()

	if err != nil {
		fmt.Println(err)
		return
	}

}

func CreateTableOfVoters() (sql.Result, error) {

	sql := `CREATE TABLE IF NOT EXISTS voters (
			vid TEXT PRIMARY KEY,
			vname TEXT NOT NULL,
			vote INTEGER NOT NULL) STRICT;`

	return DB.Exec(sql)
}

func ConnectDB() error {
	db, err := sql.Open("sqlite", "./my.db")

	if err != nil {
		return err
	}

	DB = db
	return nil
}
