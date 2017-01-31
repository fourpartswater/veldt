package promise_test

import (
	"fmt"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/unchartedsoftware/veldt/util/promise"
)

var _ = Describe("promise", func() {

	const (
		numConcurrent = 256
	)

	Describe("Resolve", func() {
		It("should accept an error value as an argument", func() {
			p := promise.NewPromise()
			p.Resolve(nil)
			Expect(p.Wait()).To(BeNil())
			p = promise.NewPromise()
			err := fmt.Errorf("error")
			p.Resolve(err)
			Expect(p.Wait()).To(Equal(err))
		})
		It("should ignore any resolutions after the first one issued", func() {
			p := promise.NewPromise()
			p.Resolve(nil)
			Expect(p.Wait()).To(BeNil())
			err := fmt.Errorf("error")
			p.Resolve(err)
			Expect(p.Wait()).To(BeNil())
		})
	})

	Describe("Wait", func() {
		It("should block until promise is resolved", func() {
			p := promise.NewPromise()
			err := fmt.Errorf("error")
			go func() {
				time.Sleep(time.Millisecond * 100)
				p.Resolve(err)
			}()
			Expect(p.Wait()).To(Equal(err))
		})
		It("should return previously resolved values", func() {
			p0 := promise.NewPromise()
			p0.Resolve(nil)
			Expect(p0.Wait()).To(BeNil())
			Expect(p0.Wait()).To(BeNil())
			Expect(p0.Wait()).To(BeNil())
			p1 := promise.NewPromise()
			err := fmt.Errorf("error")
			p1.Resolve(err)
			Expect(p1.Wait()).To(Equal(err))
			Expect(p1.Wait()).To(Equal(err))
			Expect(p1.Wait()).To(Equal(err))
		})
		It("should allow multiple routines to wait on a single resolve", func() {
			p := promise.NewPromise()
			wg := sync.WaitGroup{}
			err := fmt.Errorf("error")
			for i := 0; i < numConcurrent; i++ {
				wg.Add(1)
				go func() {
					Expect(p.Wait()).To(Equal(err))
					wg.Done()
				}()
			}
			time.Sleep(time.Millisecond * 100)
			p.Resolve(err)
			wg.Wait()
		})
	})

})
