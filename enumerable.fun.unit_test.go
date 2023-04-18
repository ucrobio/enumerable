package enumerable

import (
	"testing"

	. "expect.ucrob.io"
	. "spec.ucrob.io"
)

func TestEnumerableFunUnitSuite(t *testing.T) {
	slice := Var(Slice(1, 2, 3))

	Run(
		Define(
			"fun unit suite",

			Let(slice, func() []int { return Slice(1, 2, 3) }),

			Describe(
				".Slice",

				Inline(func() error { return Value(slice.Value()[0]).To(BeEq(1)) }),
				Inline(func() error { return Value(slice.Value()[1]).To(BeEq(2)) }),
				Inline(func() error { return Value(slice.Value()[2]).To(BeEq(3)) }),
			),

			Describe(
				".Each",

				It("interacts with all elements", func() error {
					sum := 0

					Each(func(value int) { sum += value }, slice.Value())

					return Value(sum).To(BeEq(6))
				}),
			),

			Describe(
				".Map",

				It("maps all elements", func() error {
					sum := 0

					Each(func(value int) { sum += value },
						Map(func(value int) int { return value * 3 }, slice.Value()),
					)

					return Value(sum).To(BeEq(18))
				}),
			),
		),

		func(err error) { t.Fatal(err) },
	)
}
