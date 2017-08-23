package store

import (
	"testing"
)

var schSTore = SqlSchoolStore{Database}

func TestSqlSchoolStore_RetrieveBestPerfomingSchool(t *testing.T) {
	res := <-schSTore.RetrieveBestPerfomingSchool(map[string]interface{}{"From": 2016, "To": 2016})
	if res.Err != nil {
		t.Errorf("Best perfoming school test failed with", res.Err)
	}
	t.Log("Best Perfroming school", res.Data)
}

func TestSqlSchoolStore_RankAllSchools(t *testing.T) {
	res := <-schSTore.RankAllSchools()
	if res.Err != nil {
		t.Errorf("Rank schools test failed with", res.Err)
	}
	t.Log("Ranking", res.Data)
}
