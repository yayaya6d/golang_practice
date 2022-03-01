package channel

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type chanExampleSuite struct {
	suite.Suite
}

func TestChanExampleSuite(t *testing.T) {
	suite.Run(t, new(chanExampleSuite))
}

func (s *chanExampleSuite) TestFuturePromise_CallFuturePromiseFunc_ReturnExpectedValue() {
	// act
	actual := future_promise()

	// assert
	s.Equal(500, actual)
}

func (s *chanExampleSuite) TestFuturePromise2_CallFuturePromiseChanFunc_ReturnExpectedValue() {
	// act
	actual := future_promise_2()

	// assert
	s.Equal(30, actual)
}

func (s *chanExampleSuite) TestGetFirstInput_CallGetFirstInputFunc_ReturnExpectedValue() {
	// act
	actual := getFirstInput()

	// assert
	s.Equal(1, actual)
}

func (s *chanExampleSuite) TestGetResponseWithError_CallGetResponseWithErrorFunc_ReturnExpectedValue() {
	// act
	actual_i, actual_e := getResponseWithError()

	// assert
	s.Equal(-1, actual_i)
	s.Equal(actual_e.Error(), "err")
}

func (s *chanExampleSuite) TestOneToOneChanInform_CallOneToOneChanInformFunc_ReturnExpectedValue() {
	// act
	i1, i2 := oneToOneChanInform()

	// assert
	s.Equal(9, i1)
	s.Equal(0, i2)
}

func (s *chanExampleSuite) TestOneToOneChanInform2_CallOneToOneChanInform2Func_ReturnExpectedValue() {
	// act
	actual := oneToOneChanInform_2()

	// assert
	s.Equal(1, actual)
}

func (s *chanExampleSuite) TestInformGroupByCloseChan_CallInformGroupByCloseChan_ReturnExpectedValue() {
	// act
	actual := informGroupByCloseChan()

	// assert
	s.Equal(6, actual)
}
