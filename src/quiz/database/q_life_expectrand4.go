package database

import (
	"database/sql"
	"log"
)

func Random4DistinctLifeExpectancy() (*sql.Rows, error) {

	sql := `select * 
	from life_expectancy_geo
	where locality not in (
		select locality
		from (
			select locality, count(*)
			from life_expectancy
			group by age   
			having count(*) = 1  
		)
	)
	order by random() limit 4;`

	rows, err := DB.Query(sql)

	if err != nil {
		log.Printf("Error in query: %v", err)
		return nil, err
	}

	return rows, nil

}
