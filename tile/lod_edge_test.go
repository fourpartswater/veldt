package tile_test

import (
	"github.com/unchartedsoftware/veldt/tile"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EdgeLOD", func() {

	const lod = 3

	var input []float32
	var edges []float32
	var offsets []int
	var bytes []byte

	BeforeEach(func() {
		input = []float32{
			137.24, 7.07, 1.0, 224.49, 123.95, 1.0,
			124.51, 148.33, 2.0, 72.40, 22.15, 2.0,
			160.13, 77.59, 3.0, 128.77, 183.32, 3.0,
			65.36, 36.25, 4.0, 107.91, 250.01, 4.0,
			96.05, 198.40, 2.0, 66.70, 73.39, 2.0,
		}
		edges = []float32{
			65.36, 36.25, 4.0, 107.91, 250.01, 4.0,
			137.24, 7.07, 1.0, 224.49, 123.95, 1.0,
			160.13, 77.59, 3.0, 128.77, 183.32, 3.0,
			124.51, 148.33, 2.0, 72.40, 22.15, 2.0,
			96.05, 198.40, 2.0, 66.70, 73.39, 2.0,
		}
		offsets = []int{
			0, 0, 0, 0, 0, 0, 0, 24, 24, 24, 24, 24,
			24, 24, 24, 24, 24, 48, 48, 48, 48, 48,
			48, 48, 48, 48, 72, 72, 72, 72, 72, 72,
			72, 72, 72, 72, 72, 72, 96, 96, 96, 96,
			96, 96, 96, 96, 120, 120, 120, 120, 120,
			120, 120, 120, 120, 120, 120, 120, 120,
			120, 120, 120, 120, 120,
		}
		bytes = []byte{
			120, 0, 0, 0, 0, 1, 0, 0, 82, 184, 130, 66, 0, 0, 17, 66, 0, 0, 128,
			64, 236, 209, 215, 66, 143, 2, 122, 67, 0, 0, 128, 64, 113, 61, 9,
			67, 113, 61, 226, 64, 0, 0, 128, 63, 113, 125, 96, 67, 102, 230,
			247, 66, 0, 0, 128, 63, 72, 33, 32, 67, 20, 46, 155, 66, 0, 0, 64,
			64, 31, 197, 0, 67, 236, 81, 55, 67, 0, 0, 64, 64, 31, 5, 249, 66,
			123, 84, 20, 67, 0, 0, 0, 64, 205, 204, 144, 66, 51, 51, 177, 65,
			0, 0, 0, 64, 154, 25, 192, 66, 102, 102, 70, 67, 0, 0, 0, 64, 102,
			102, 133, 66, 174, 199, 146, 66, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 24,
			0, 0, 0, 24, 0, 0, 0, 24, 0, 0, 0, 24, 0, 0, 0, 24, 0, 0, 0, 24, 0,
			0, 0, 24, 0, 0, 0, 24, 0, 0, 0, 24, 0, 0, 0, 24, 0, 0, 0, 48, 0, 0,
			0, 48, 0, 0, 0, 48, 0, 0, 0, 48, 0, 0, 0, 48, 0, 0, 0, 48, 0, 0, 0,
			48, 0, 0, 0, 48, 0, 0, 0, 48, 0, 0, 0, 72, 0, 0, 0, 72, 0, 0, 0, 72,
			0, 0, 0, 72, 0, 0, 0, 72, 0, 0, 0, 72, 0, 0, 0, 72, 0, 0, 0, 72, 0,
			0, 0, 72, 0, 0, 0, 72, 0, 0, 0, 72, 0, 0, 0, 72, 0, 0, 0, 96, 0, 0,
			0, 96, 0, 0, 0, 96, 0, 0, 0, 96, 0, 0, 0, 96, 0, 0, 0, 96, 0, 0, 0,
			96, 0, 0, 0, 96, 0, 0, 0, 120, 0, 0, 0, 120, 0, 0, 0, 120, 0, 0, 0,
			120, 0, 0, 0, 120, 0, 0, 0, 120, 0, 0, 0, 120, 0, 0, 0, 120, 0, 0,
			0, 120, 0, 0, 0, 120, 0, 0, 0, 120, 0, 0, 0, 120, 0, 0, 0, 120, 0,
			0, 0, 120, 0, 0, 0, 120, 0, 0, 0, 120, 0, 0, 0, 120, 0, 0, 0, 120,
			0, 0, 0,
		}
	})

	Describe("EdgeLOD", func() {
		It("should sort the provided []float32 by morton code and return the sorted edges along with LOD offsets", func() {
			es, os := tile.EdgeLOD(input, lod)
			Expect(es).To(Equal(edges))
			Expect(os).To(Equal(offsets))
		})
	})

	Describe("EncodeEdgeLOD", func() {
		It("should sort the provided []float32 by morton code and encode the results along with appropriate LOD offsets", func() {
			bs := tile.EncodeEdgeLOD(input, lod)
			Expect(bs).To(Equal(bytes))
		})
	})

})
