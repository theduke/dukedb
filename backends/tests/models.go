package tests

import(
	"fmt"
	"strconv"

	. "github.com/theduke/go-dukedb"
)

type BaseModel struct {
}

func (i BaseModel) GetID() string {
	return ""
}
func (i BaseModel) SetID(x string) error {
	return nil
}
func (i BaseModel) Collection() string {
	return "base_model"
}

type TestModel struct {
	ID uint64

	// For inferred belongs-to
	TestParentID uint64 `db:"ignore-zero"`

	// For explicit has-one/belongs-to
	MyParent   *TestParent `db:"has-one:MyParentID:ID"`
	MyParentID uint64      `db:"ignore-zero"`

	StrVal string
	IntVal int64
}

func (t TestModel) Collection() string {
	return "test_models"
}

func (t TestModel) GetID() string {
	if t.ID == 0 {
		return ""
	}
	return strconv.FormatUint(t.ID, 10)
}

func (t *TestModel) SetID(x string) error {
	id, err := strconv.ParseUint(x, 10, 64)
	if err != nil {
		return err
	}
	t.ID = id
	return nil
}

type TestParent struct {
	TestModel

	Child   TestModel
	ChildID uint64

	ChildPtr *TestModel

	ChildSlice    []TestModel
	ChildSlice2   []TestModel  `db:"belongs-to:ID:MyParentID"`
	ChildSlicePtr []*TestModel `db:"m2m"`
}

func (t TestParent) Collection() string {
	return "test_parents"
}

func NewTestModel(index int) TestModel {
	return TestModel{
		StrVal: fmt.Sprintf("str%v", index),
		IntVal: int64(index),
	}
}

func NewTestParent(index int, withChildren bool) TestParent {
	base := NewTestModel(index)

	var child TestModel
	var childPtr *TestModel
	var childSlice []TestModel
	var childPtrSlice []*TestModel
	if withChildren {
		child = NewTestModel(index*10 + 1)

		child2 := NewTestModel(index*10 + 2)
		childPtr = &child2

		childSlice = NewTestModelSlice(index*10+3, 2)
		childPtrSlice = NewTestModelPtrSlice(index*10+5, 2)
	}

	return TestParent{
		TestModel: base,

		Child:    child,
		ChildPtr: childPtr,

		ChildSlice:    childSlice,
		ChildSlicePtr: childPtrSlice,
	}
}

func NewTestModelSlice(startIndex int, count int) []TestModel {
	slice := make([]TestModel, 0)
	for i := startIndex; i < startIndex+count; i++ {
		slice = append(slice, NewTestModel(i))
	}

	return slice
}

func NewTestModelPtrSlice(startIndex int, count int) []*TestModel {
	slice := make([]*TestModel, 0)
	for i := startIndex; i < startIndex+count; i++ {
		model := NewTestModel(i)
		slice = append(slice, &model)
	}

	return slice
}

func NewTestModelInterfaceSlice(startIndex int, count int) []Model {
	slice := make([]Model, 0)
	for i := startIndex; i < startIndex+count; i++ {
		model := NewTestModel(i)
		slice = append(slice, &model)
	}

	return slice
}

func NewTestParentSlice(startIndex int, count int, withChildren bool) []TestParent {
	slice := make([]TestParent, 0)
	for i := startIndex; i < startIndex+count; i++ {
		slice = append(slice, NewTestParent(i, withChildren))
	}

	return slice
}

func NewTestParentPtrSlice(startIndex int, count int, withChildren bool) []*TestParent {
	slice := make([]*TestParent, 0)
	for i := startIndex; i < startIndex+count; i++ {
		parent := NewTestParent(i, withChildren)
		slice = append(slice, &parent)
	}

	return slice
}

func NewTestParentInterfaceSlice(startIndex int, count int, withChildren bool) []Model {
	slice := make([]Model, 0)
	for i := startIndex; i < startIndex+count; i++ {
		model := NewTestParent(i, withChildren)
		slice = append(slice, &model)
	}

	return slice
}
