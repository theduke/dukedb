package dukedb_test

import (
	"reflect"

	. "github.com/theduke/go-dukedb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/theduke/go-dukedb/backends/tests"
)

var _ = Describe("Utils", func() {

	Describe("CamelCaseToUnderscore", func() {
		It("Should be empty", func() {
			Expect(CamelCaseToUnderscore("")).To(Equal(""))
		})

		It("Should be single lower case letter from upper case", func() {
			Expect(CamelCaseToUnderscore("A")).To(Equal("a"))
		})

		It("Should be single lower case letter from lower case", func() {
			Expect(CamelCaseToUnderscore("a")).To(Equal("a"))
		})

		It("Should be two lower case letters", func() {
			Expect(CamelCaseToUnderscore("AA")).To(Equal("aa"))
		})

		It("Should be two parts with underscore", func() {
			Expect(CamelCaseToUnderscore("AaBb")).To(Equal("aa_bb"))
		})

		It("Should be three parts with underscore", func() {
			Expect(CamelCaseToUnderscore("AaBbCc")).To(Equal("aa_bb_cc"))
		})

		It("Should end with lower case", func() {
			Expect(CamelCaseToUnderscore("AaBbC")).To(Equal("aa_bb_c"))
		})
	})

	Describe("GetStructFieldValue", func() {
		type TestStruct struct {
			Val       string
			ValEmpty  string
			IntVal    int
			StructVal *TestStruct
		}

		var testStruct TestStruct

		BeforeEach(func() {
			testStruct = TestStruct{
				Val:       "test",
				IntVal:    33,
				StructVal: &TestStruct{Val: "test2"},
			}
		})

		Context("With nil struct argument", func() {
			It("Should fail on nil", func() {
				_, err := GetStructFieldValue(nil, "test")
				Expect(err).To(HaveOccurred())
				Expect(err.GetCode()).To(Equal("pointer_or_struct_expected"))
			})
		})

		Context("With pointer to non-struct", func() {
			It("Should fail on pointer to non-struct", func() {
				x := 22
				_, err := GetStructFieldValue(&x, "test")
				Expect(err).To(HaveOccurred())
				Expect(err.GetCode()).To(Equal("struct_expected"))
			})
		})

		Context("With non-struct arugment", func() {
			It("Should fail on non-struct argument", func() {
				_, err := GetStructFieldValue(22, "test")
				Expect(err).To(HaveOccurred())
				Expect(err.GetCode()).To(Equal("struct_expected"))
			})
		})

		Context("With valid fields", func() {
			It("Should be valid string", func() {
				Expect(GetStructFieldValue(&testStruct, "Val")).To(Equal("test"))
			})

			It("Should be valid int", func() {
				Expect(GetStructFieldValue(&testStruct, "IntVal")).To(Equal(33))
			})

			It("Should be pointer to struct", func() {
				Expect(GetStructFieldValue(&testStruct, "StructVal")).To(Equal(&TestStruct{Val: "test2"}))
			})
		})

		Context("With invalid fields", func() {
			It("Should fail on non-existant field", func() {
				_, err := GetStructFieldValue(&testStruct, "DoesNotExist")
				Expect(err).To(HaveOccurred())
				Expect(err.GetCode()).To(Equal("field_not_found"))
			})
		})
	})

	Describe("CompareValues", func() {
		It("Should eq with two strings", func() {
			a := interface{}("test")
			b := interface{}("test")
			Expect(CompareValues("eq", a, b)).To(BeTrue())
		})

		It("Should eq with two numbers", func() {
			a := interface{}(1)
			b := interface{}(uint32(1))
			Expect(CompareValues("eq", a, b)).To(BeTrue())
		})
	})

	Describe("CompareStringValues", func() {
		It("Should eq", func() {
			a := interface{}("test")
			b := interface{}("test")
			Expect(CompareStringValues("eq", a, b)).To(BeTrue())
		})
	})

	Describe("CompareNumericValues", func() {
		It("Should lt with ints", func() {
			a := interface{}(1)
			b := interface{}(2)
			Expect(CompareNumericValues("lt", a, b)).To(BeTrue())
		})

		It("Should gt with int64 and uint8", func() {
			a := interface{}(int64(1))
			b := interface{}(uint8(5))
			Expect(CompareNumericValues("lt", a, b)).To(BeTrue())
		})
	})

	Describe("SortStructSlice", func() {
		type Sortable struct {
			IntVal   int
			FloatVal float32
			StrVal   string
		}

		var sortables []interface{}

		BeforeEach(func() {
			sortables = []interface{}{
				Sortable{5, 5.1, "5"},
				Sortable{3, 3.1, "3"},
				Sortable{1, 1.1, "1"},
				Sortable{2, 2.1, "2"},
				Sortable{4, 4.1, "4"},
			}
		})

		It("Should sort asc by int field", func() {
			SortStructSlice(sortables, "IntVal", true)
			Expect((sortables[0]).(Sortable).IntVal).To(Equal(1))
			Expect(sortables[4].(Sortable).IntVal).To(Equal(5))
		})

		It("Should sort desc by int field", func() {
			SortStructSlice(sortables, "IntVal", false)
			Expect((sortables[0]).(Sortable).IntVal).To(Equal(5))
			Expect(sortables[4].(Sortable).IntVal).To(Equal(1))
		})

		It("Should sort asc by string field", func() {
			SortStructSlice(sortables, "StrVal", true)
			Expect((sortables[0]).(Sortable).StrVal).To(Equal("1"))
			Expect(sortables[4].(Sortable).StrVal).To(Equal("5"))
		})
	})

	Describe("ConvertStringToType", func() {
		It("Should convert int", func() {
			Expect(ConvertStringToType("-22", reflect.Int)).To(Equal(-22))
		})

		It("Should convert int64", func() {
			Expect(ConvertStringToType("-22", reflect.Int64)).To(Equal(int64(-22)))
		})

		It("Should convert uint", func() {
			Expect(ConvertStringToType("22", reflect.Uint)).To(Equal(uint(22)))
		})

		It("Should convert uint64", func() {
			Expect(ConvertStringToType("22", reflect.Uint64)).To(Equal(uint64(22)))
		})

		It("Should convert string", func() {
			Expect(ConvertStringToType("test", reflect.String)).To(Equal("test"))
		})

		It("Should error on invalid int", func() {
			_, err := ConvertStringToType("test", reflect.Int)
			Expect(err).To(HaveOccurred())
		})

		It("Should error on unsupported type", func() {
			_, err := ConvertStringToType("22", reflect.Ptr)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("cannot_convert_to_ptr"))
		})
	})

	Describe("SetStructFieldValueFromString", func() {
		type TestStruct struct {
			Val    string
			IntVal int
		}

		var testStruct TestStruct

		BeforeEach(func() {
			testStruct = TestStruct{}
		})

		It("Should error on non-pointer", func() {
			err := SetStructFieldValueFromString(22, "Val", "22")
			Expect(err).To(HaveOccurred())
			Expect(err.GetCode()).To(Equal("pointer_expected"))
		})

		It("Should error on non-struct pointer", func() {
			x := 22
			err := SetStructFieldValueFromString(&x, "Val", "22")
			Expect(err).To(HaveOccurred())
			Expect(err.GetCode()).To(Equal("pointer_to_struct_expected"))
		})

		It("Should error on inexistant field", func() {
			err := SetStructFieldValueFromString(&testStruct, "InvalidField", "22")
			Expect(err).To(HaveOccurred())
			Expect(err.GetCode()).To(Equal("unknown_field"))
		})

		It("Should error on unconvertable type", func() {
			err := SetStructFieldValueFromString(&testStruct, "IntVal", "xxx")
			Expect(err).To(HaveOccurred())
		})

		It("Should work with valid args for string field", func() {
			Expect(SetStructFieldValueFromString(&testStruct, "Val", "xxx")).ToNot(HaveOccurred())
			Expect(testStruct.Val).To(Equal("xxx"))
		})

		It("Should work with valid args for int field", func() {
			Expect(SetStructFieldValueFromString(&testStruct, "IntVal", "22")).ToNot(HaveOccurred())
			Expect(testStruct.IntVal).To(Equal(22))
		})

	})

	Describe("GetModelSliceFieldValues", func() {
		var modelSlice []interface{}

		BeforeEach(func() {
			modelSlice = []interface{}{&TestModel{
				ID:     1,
				StrVal: "str1",
				IntVal: 1,
			}, &TestModel{
				ID:     2,
				StrVal: "str2",
				IntVal: 2,
			}}
		})

		It("Should error on invalid field", func() {
			_, err := GetModelSliceFieldValues(modelSlice, "InvalidField")
			Expect(err).To(HaveOccurred())
			Expect(err.GetCode()).To(Equal("field_not_found"))
		})

		It("Should work for str field", func() {
			val := []interface{}{"str1", "str2"}
			Expect(GetModelSliceFieldValues(modelSlice, "StrVal")).To(Equal(val))
		})

		It("Should work for int field", func() {
			val := []interface{}{1, 2}
			Expect(GetModelSliceFieldValues(modelSlice, "IntVal")).To(Equal(val))
		})
	})

	Describe("FilterToSqlCondition", func() {

		It("Should convert eq", func() {
			Expect(FilterToSqlCondition("eq")).To(Equal("="))
		})

		It("Should convert neq", func() {
			Expect(FilterToSqlCondition("neq")).To(Equal("!="))
		})

		It("Should convert lt", func() {
			Expect(FilterToSqlCondition("lt")).To(Equal("<"))
		})

		It("Should convert lte", func() {
			Expect(FilterToSqlCondition("lte")).To(Equal("<="))
		})

		It("Should convert gt", func() {
			Expect(FilterToSqlCondition("gt")).To(Equal(">"))
		})

		It("Should convert gte", func() {
			Expect(FilterToSqlCondition("gte")).To(Equal(">="))
		})

		It("Should convert like", func() {
			Expect(FilterToSqlCondition("like")).To(Equal("LIKE"))
		})

		It("Should convert eq", func() {
			Expect(FilterToSqlCondition("eq")).To(Equal("="))
		})

		It("Should convert in", func() {
			Expect(FilterToSqlCondition("in")).To(Equal("IN"))
		})

		It("Should error on invalid filter", func() {
			_, err := FilterToSqlCondition("XXX")
			Expect(err).To(HaveOccurred())
			Expect(err.GetCode()).To(Equal("unknown_filter"))
		})
	})

	Describe("InterfaceToModelSlice", func() {
		var slice []interface{}

		var modelSlice []interface{}

		BeforeEach(func() {
			modelSlice = NewTestModelInterfaceSlice(1, 2)
			slice = []interface{}{
				modelSlice[0],
				modelSlice[1],
			}
		})

		It("Should error on non-slice/non pointer slice argument", func() {
			_, err := InterfaceToModelSlice(22)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("slice_expected"))
		})

		It("Should fail on non-model slice", func() {
			_, err := InterfaceToModelSlice([]int{1, 2})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("slice_values_do_not_implement_model_if"))
		})

		It("Should work with pointer to model slice", func() {
			Expect(InterfaceToModelSlice(&slice)).To(Equal(modelSlice))
		})

		It("Should work with model slice", func() {
			Expect(InterfaceToModelSlice(slice)).To(Equal(modelSlice))
		})
	})

	Describe("NewStruct", func() {

		It("Should error on non-struct", func() {
			_, err := NewStruct(22)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("struct_expected"))
		})

		It("Should error on ptr non-struct", func() {
			x := 22
			_, err := NewStruct(&x)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("struct_expected"))
		})

		It("Should build struct from pointer", func() {
			s, _ := NewStruct(TestModel{})
			Expect(s).To(Equal(&TestModel{}))
		})

	})

	Describe("NewSlice", func() {

		It("Should build int slice", func() {
			Expect(NewSlice(22)).To(Equal([]int{}))
		})

		It("Should build pointer to model slice", func() {
			Expect(NewSlice(&TestModel{})).To(Equal([]*TestModel{}))
		})
	})

	Describe("SetStructModelField", func() {

		var testParent *TestParent

		BeforeEach(func() {
			p := NewTestParent(1, false)
			testParent = &p
		})

		It("Should error on non-pointer", func() {
			err := SetStructModelField(22, "Child", []interface{}{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pointer_expected"))
		})

		It("Should error on pointer to non-struct", func() {
			x := 22
			err := SetStructModelField(&x, "Child", []interface{}{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pointer_to_struct_expected"))
		})

		It("Should error on unknown field", func() {
			err := SetStructModelField(testParent, "InvalidField", []interface{}{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("unknown_field"))
		})

		It("Should error on invalid target field type", func() {
			err := SetStructModelField(testParent, "StrVal", []interface{}{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("unsupported_field_type"))
		})

		It("Should set struct", func() {
			child := NewTestModel(1)
			SetStructModelField(testParent, "Child", []interface{}{&child})
			Expect(testParent.Child).To(Equal(child))
		})

		It("Should set struct pointer", func() {
			child := NewTestModel(1)
			SetStructModelField(testParent, "ChildPtr", []interface{}{&child})
			Expect(testParent.ChildPtr).To(Equal(&child))
		})

		It("Should set slice", func() {
			childSlice := NewTestModelSlice(1, 2)
			ifSlice := []interface{}{childSlice[0], childSlice[1]}

			err := SetStructModelField(testParent, "ChildSlice", ifSlice)
			Expect(err).ToNot(HaveOccurred())

			Expect(testParent.ChildSlice).To(Equal(childSlice))
		})

		It("Should set pointer slice", func() {
			childSlice := NewTestModelPtrSlice(1, 2)
			ifSlice := []interface{}{childSlice[0], childSlice[1]}
			SetStructModelField(testParent, "ChildSlicePtr", ifSlice)
			Expect(testParent.ChildSlicePtr).To(Equal(childSlice))
		})
	})
})
