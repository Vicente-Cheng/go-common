package gocommon

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SliceFuncs struct {
	suite.Suite
}

func TestSliceFuncs(t *testing.T) {
	suite.Run(t, new(SliceFuncs))
}

func (s *SliceFuncs) SetupSuite() {
	// you could do something here before all tests
}

func (s *SliceFuncs) TestSliceContentCmpInt() {
	a := []int{1, 1, 2, 3}
	b := []int{3, 2, 1, 1}
	c := []int{1, 2, 3, 4}
	d := []int{}
	e := []int{}
	f := []int{1, 1, 2}
	g := []int{1, 2, 2}
	fCpy := []int{1, 1, 2}
	gCpy := []int{1, 2, 2}

	s.Equal(true, SliceContentCmp(a, b), "SliceContentCmp should return true")
	s.Equal(false, SliceContentCmp(a, c), "SliceContentCmp should return false")
	s.Equal(true, SliceContentCmp(d, e), "SliceContentCmp should return true")
	s.Equal(false, SliceContentCmp(f, g), "SliceContentCmp should return false")
	s.Equal(f, fCpy, "original slice should not change.")
	s.Equal(g, gCpy, "original slice should not change.")

}

func (s *SliceFuncs) TestSliceContentCmpString() {
	a := []string{"a", "b", "c"}
	b := []string{"b", "c", "a"}
	c := []string{"a", "b", "c", "d"}
	d := []string{}
	e := []string{}
	f := []string{"a", "a", "b"}
	g := []string{"a", "b", "b"}
	fCpy := []string{"a", "a", "b"}
	gCpy := []string{"a", "b", "b"}

	s.Equal(true, SliceContentCmp(a, b), "SliceContentCmp should return true")
	s.Equal(false, SliceContentCmp(a, c), "SliceContentCmp should return false")
	s.Equal(true, SliceContentCmp(d, e), "SliceContentCmp should return true")
	s.Equal(false, SliceContentCmp(f, g), "SliceContentCmp should return false")
	s.Equal(f, fCpy, "original slice should not change.")
	s.Equal(g, gCpy, "original slice should not change.")
}

func (s *SliceFuncs) TestSliceDedupeInt() {
	a := []int{1, 1, 2, 2, 3, 3}
	b := []int{1, 2, 3, 4, 5, 6}
	c := []int{1, 2, 3, 3, 4}

	require.Equal(s.T(), []int{1, 2, 3}, SliceDedupe(a), "SliceDedupe should return the same slice with {1, 2, 3}")
	require.Equal(s.T(), []int{1, 2, 3, 4, 5, 6}, SliceDedupe(b), "SliceDedupe should return the same slice with {1, 2, 3, 4, 5, 6}")
	require.Equal(s.T(), []int{1, 2, 3, 4}, SliceDedupe(c), "SliceDedupe should return the same slice with {1, 2, 3, 4}")
}

func (s *SliceFuncs) TestSliceDedupeString() {
	a := []string{"foo", "bar", "bar", "foo", "roll", "roll"}
	b := []string{"apple", "book", "clock", "duck", "escape", "field"}
	c := []string{"foo", "bar", "roll", "roll", "desk"}

	require.Equal(s.T(), []string{"foo", "bar", "roll"}, SliceDedupe(a), "SliceDedupe should return the same slice with {foo, bar, roll}")
	require.Equal(s.T(), []string{"apple", "book", "clock", "duck", "escape", "field"}, SliceDedupe(b), "SliceDedupe should return the same slice with {apple, book, clock, duck, escape, field}")
	require.Equal(s.T(), []string{"foo", "bar", "roll", "desk"}, SliceDedupe(c), "SliceDedupe should return the same slice with {foo, bar, roll, desk}")
}
