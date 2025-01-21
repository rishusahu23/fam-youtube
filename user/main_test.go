package user

import (
	"context"
	"flag"
	"github.com/golang/mock/gomock"
	"github.com/rishusahu23/fam-youtube/config"
	mocks "github.com/rishusahu23/fam-youtube/user/mocks/dao"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	//ctx := context.Background()
	//conf := Initialise(ctx)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func Initialise(ctx context.Context) *config.Config {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}
	return conf
}

type MockDependencies struct {
	dao *mocks.MockUserDao
	txn *mocks2.MockTransactionManager
}

func GetServiceWithDependencies(t *testing.T) (*Service, *MockDependencies) {
	ctr := gomock.NewController(t)
	dao := mocks.NewMockUserDao(ctr)
	txn := mocks2.NewMockTransactionManager(ctr)
	svc := &Service{
		dao:        dao,
		txnManager: txn,
	}
	md := &MockDependencies{
		dao: dao,
		txn: txn,
	}
	return svc, md
}
