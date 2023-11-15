package implement

import (
	"context"
	"os"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	v1 "github.com/superjcd/calendarservice/genproto/v1"
	"github.com/superjcd/calendarservice/service"
	"gorm.io/gorm"
)

var dbFile = "fake.db"

type TestSuite struct {
	suite.Suite
	Dbfile string
	cs     service.CalendarService
}

func (suite *TestSuite) SetupSuite() {
	file, err := os.Create(dbFile)
	assert.Nil(suite.T(), err)
	defer file.Close()

	suite.Dbfile = dbFile
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	assert.Nil(suite.T(), err)

	service, err := NewCalendarService(db)
	assert.Nil(suite.T(), err)
	suite.cs = service
}

func (suite *TestSuite) TearDownSuite() {
	var err error
	err = suite.cs.Close()
	assert.Nil(suite.T(), err)
	err = os.Remove(dbFile)
	assert.Nil(suite.T(), err)
}

func (suite *TestSuite) TestCreateCalendarItem() {
	rq := v1.CreateCalendarItemRequest{
		Creator: "jack",
		Date:    "2023-10-30",
		Content: "do something about that",
	}

	err := suite.cs.CreateCalendarItem(context.Background(), &rq)
	assert.Nil(suite.T(), err)

}

func (suite *TestSuite) TestListCalendarItem() {
	rq := v1.ListCalendarItemsRequest{
		Creator: "jack",
	}
	items, err := suite.cs.ListCalendarItems(context.Background(), &rq)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(items))
}

func (suite *TestSuite) TestUpdtCalendarItem() {
	var err error
	rq := v1.UpdateCalendarItemRequest{
		Creator: "jack",
		Date:    "2023-10-30",
		Content: "do something about this",
	}

	err = suite.cs.UpdateCalendarItem(context.Background(), &rq)
	assert.Nil(suite.T(), err)

	rq2 := v1.ListCalendarItemsRequest{
		Creator: "jack",
	}
	items, err2 := suite.cs.ListCalendarItems(context.Background(), &rq2)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), 1, len(items))
	assert.Equal(suite.T(), "do something about this", items[0].Content)
}

func TestFakeStoreSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
