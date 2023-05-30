package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("yajed", func(t *testing.T) {
		result := HelloWorld("yajed")
		assert.Equal(t,"hello yajed",result,"harusnya sih sama")
	})
	t.Run("tono", func(t *testing.T) {
		result := HelloWorld("tono")
		assert.Equal(t,"hello tono",result,"harusnya sih sama")
	})
	t.Run("Toni", func(t *testing.T) {
		result := HelloWorld("toni")
		assert.Equal(t,"hello toni",result,"harusnya sih sama")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("before")
	m.Run()
	fmt.Println("after")
}


func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("waduh pake windows nih")
	}
	fmt.Println("apakah dicetak?")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("yajed")
	require.Equal(t,"hello yajed",result,"harusnya sih sama")
	fmt.Println("assert sudah selesai dipanggil")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("yajed")
	assert.Equal(t,"hello yajed",result,"harusnya sih sama")
	fmt.Println("assert sudah selesai dipanggil")
}

func TestHelloWorld(t *testing.T) {

	result := HelloWorld("yajed")

	if result != "hello yajed" {
		// t.Fail()
		t.Error("harusnya hello yajed")
	}
	

	fmt.Println("sudah dieksekusi")
	
}

func TestHelloWorldToni(t *testing.T) {

	result := HelloWorld("toni")

	if result != "hello toni" {
	// t.FailNow()
	t.Fatal("harusnya hello toni")
	}

	fmt.Println("sudah dieksekusi")



}


