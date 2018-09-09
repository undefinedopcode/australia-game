package database

import "database/sql"

func Random4DistinctPopulations() (*sql.Rows, error) {

	sql := `select * 
	from suburb_pops_geo
	where locality not in (
		select locality
		from (
			select locality, count(*)
			from suburb_population
			group by population   
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
