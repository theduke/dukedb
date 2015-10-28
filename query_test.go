package dukedb_test

import (
	db "github.com/theduke/go-dukedb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("db.Query", func() {

	It("Should set model", func() {
		q := db.Q("test_model")
		Expect(q.GetCollection()).To(Equal("test_model"))
	})

	It("Should set limit", func() {
		q := db.Q("test_model").Limit(22)
		Expect(q.GetLimit()).To(Equal(22))
	})

	/*
			It("Should set fields", func() {
				q := db.Q("test_model").Field("field1", "field2", "field3")
				Expect(q.GetFields()).To(Equal([]string{"field1", "field2", "field3"}))
			})

			It("Should set fields on LimitFields()", func() {
				q := db.Q("test_model").LimitFields("field1", "field2", "field3")
				Expect(q.GetFields()).To(Equal([]string{"field1", "field2", "field3"}))
			})

			It("Should limit fields", func() {
				q := db.Q("test_model").Fields("field1", "field2", "field3").LimitFields("field2", "field1")
				Expect(q.GetFields()).To(Equal([]string{"field1", "field2"}))
			})

		It("Should order", func() {
			q := db.Q("test_model").Order("field1", true).Order("field2", false)
			Expect(q.GetOrders()).To(Equal([]OrderSpec{
				Order("field1", true),
				Order("field2", false)}))
		})

		It("Should set order", func() {
			q := db.Q("test_model").Order("fieldxx", true).
				SetOrders(Order("field1", true), Order("field2", false))

			Expect(q.GetOrders()).To(Equal([]OrderSpec{
				Order("field1", true),
				Order("field2", false)}))
		})

		It("Should Filterdb.Q", func() {
			q := db.Q("test_model").Filterdb.Q(Eq("field1", "val")).Filterdb.Q(Neq("field2", 22), Eq("field3", 33))
			Expect(q.GetFilters()).To(Equal([]Filter{
				Eq("field1", "val"),
				Neq("field2", 22),
				Eq("field3", 33)}))
		})

		It("Should set filters", func() {
			q := db.Q("test_model").SetFilters(Eq("field1", "val"), Neq("field2", 22))
			Expect(q.GetFilters()).To(Equal([]Filter{
				Eq("field1", "val"),
				Neq("field2", 22)}))
		})

		It("Should Filter", func() {
			q := db.Q("test_model").Filter("field1", "val").Filter("field2", 22)
			Expect(q.GetFilters()).To(Equal([]Filter{
				Eq("field1", "val"),
				Eq("field2", 22)}))
		})

		It("Should FilterCond", func() {
			q := db.Q("test_model").FilterCond("field1", "=", "val").FilterCond("field2", "!=", 22)
			Expect(q.GetFilters()).To(Equal([]Filter{
				Eq("field1", "val"),
				Neq("field2", 22)}))
		})

		It("Should Anddb.Q", func() {
			q := db.Q("test_model").Anddb.Q(Eq("field1", "val")).Anddb.Q(Neq("field2", 22), Eq("field3", 33))
			Expect(q.GetFilters()).To(Equal([]Filter{
				Eq("field1", "val"),
				Neq("field2", 22),
				Eq("field3", 33)}))
		})

		It("Should And", func() {
			q := db.Q("test_model").And("field1", "val").And("field2", 22)
			Expect(q.GetFilters()).To(Equal([]Filter{
				Eq("field1", "val"),
				Eq("field2", 22)}))
		})

		It("Should AndCond", func() {
			q := db.Q("test_model").AndCond("field1", "=", "val").AndCond("field2", "!=", 22)
			Expect(q.GetFilters()).To(Equal([]Filter{
				Eq("field1", "val"),
				Neq("field2", 22)}))
		})

		It("Should Ordb.Q with no filters present", func() {
			q := db.Q("test_model").Ordb.Q(Eq("field1", "val"))
			Expect(q.GetFilters()).To(Equal([]Filter{
				Eq("field1", "val")}))
		})

		It("Should fail Ordb.Q with multiple filters present", func() {
			Expect(func() {
				db.Q("test_model").And("field1", "val").And("field2", 22).Ordb.Q(Eq("field1", "val"))
			}).To(Panic())
		})

		It("Should Ordb.Q with one filter present", func() {
			q := db.Q("test_model").And("field1", "val").Ordb.Q(Eq("field2", 22))
			Expect(q.GetFilters()).To(Equal([]Filter{
				Or(Eq("field1", "val"), Eq("field2", 22))}))
		})

		It("Should Ordb.Q with top level or present", func() {
			q := db.Q("test_model").Filterdb.Q(Or(Eq("field1", "val")))
			q.Ordb.Q(Eq("field2", 22))

			Expect(q.GetFilters()).To(Equal([]Filter{
				Or(Eq("field1", "val"), Eq("field2", 22))}))
		})

		It("Should Or", func() {
			q := db.Q("test_model").Or("field1", "val").Or("field2", 22)
			Expect(q.GetFilters()).To(Equal([]Filter{
				Or(Eq("field1", "val"), Eq("field2", 22))}))
		})

		It("Should OrCond", func() {
			q := db.Q("test_model").OrCond("field1", "=", "val").OrCond("field2", "!=", 22)
			Expect(q.GetFilters()).To(Equal([]Filter{
				Or(Eq("field1", "val"), Neq("field2", 22))}))
		})

		It("Should Not()", func() {
			q := db.Q("test_model").Not("field1", "val")
			Expect(q.GetFilters()).To(Equal([]Filter{Neq("field1", "val")}))
		})
	*/
})
