package database

import "database/sql"

func Random4DistinctMedianRents() (*sql.Rows, error) {

	sql := `select * 
	from median_rents_geo
	where locality not in (
		select locality
		from (
			select locality, count(*)
			from median_rents
			group by amount   
			having count(*) = 1  
		)
	)
	order by random() limit 4;`

	rows, err := DB.Query(sql)

	if err != nil {
		return nil, err
	}

	return rows, nil

}
