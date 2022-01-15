package sets

import "testing"

func TestIntersect(t *testing.T) {

	setA := NewByteSet()
	testString := "Application"
	testStringArr := []byte(testString)
	for _, v := range testStringArr {
		setA.Add(v)
	}

	expected := NewByteSet()
	expArr := []byte{'A', 'p', 'l', 'i', 'c','a', 't', 'o', 'n'}
	for _, v := range expArr {
		expected.Add(v)
	}

	if !setA.Equal(expected) {
		t.Error()
	}

}
