package concat

import (
	"bytes"
	"strings"
	"testing"
)

func ConcatOperator(b *testing.B, numStr int) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		str := "a"
		for j := 0; j < numStr; j++ {
			str += "a"
		}
	}
}

func ConcatJoin(b *testing.B, numStr int) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		strArr := []string{"a"}
		for j := 0; j < numStr; j++ {
			strArr = append(strArr, "a")
		}

		strings.Join(strArr, "")
	}
}

func ConcatBuffer(b *testing.B, numStr int) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf := bytes.NewBufferString("a")
		for j := 0; j < numStr; j++ {
			buf.WriteString("a")
		}
	}
}

func BenchmarkOperator10(b *testing.B) {
	ConcatOperator(b, 10)
}

func BenchmarkJoin10(b *testing.B) {
	ConcatJoin(b, 10)
}

func BenchmarkBuffer10(b *testing.B) {
	ConcatBuffer(b, 10)
}

func BenchmarkOperator1000(b *testing.B) {
	ConcatOperator(b, 1000)
}

func BenchmarkJoin1000(b *testing.B) {
	ConcatJoin(b, 1000)
}

func BenchmarkBuffer1000(b *testing.B) {
	ConcatBuffer(b, 1000)
}

func BenchmarkOperator10000(b *testing.B) {
	ConcatOperator(b, 10000)
}

func BenchmarkJoin10000(b *testing.B) {
	ConcatJoin(b, 10000)
}

func BenchmarkBuffer10000(b *testing.B) {
	ConcatBuffer(b, 10000)
}
