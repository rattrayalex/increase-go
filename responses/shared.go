package responses

import (
	"net/http"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/option"
)

type Page[T any] struct {
	Data []T `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       PageJSON
	cfg        *option.RequestConfig
	res        *http.Response
}

type PageJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Page[T] using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Page[T]) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *Page[T]) GetNextPage() (res *Page[T], err error) {
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("cursor", next))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Page[T]) SetPageConfig(cfg *option.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type PageAutoPager[T any] struct {
	page *Page[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageAutoPager[T any](page *Page[T], err error) *PageAutoPager[T] {
	return &PageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageAutoPager[T]) Next() bool {
	if len(r.page.Data) != 0 && r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageAutoPager[T]) Err() error {
	return r.err
}

func (r *PageAutoPager[T]) Index() int {
	return r.run
}

type PointOfServiceEntryMode string

const (
	PointOfServiceEntryModeManual                     PointOfServiceEntryMode = "manual"
	PointOfServiceEntryModeMagneticStripeNoCvv        PointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	PointOfServiceEntryModeOpticalCode                PointOfServiceEntryMode = "optical_code"
	PointOfServiceEntryModeIntegratedCircuitCard      PointOfServiceEntryMode = "integrated_circuit_card"
	PointOfServiceEntryModeContactless                PointOfServiceEntryMode = "contactless"
	PointOfServiceEntryModeCredentialOnFile           PointOfServiceEntryMode = "credential_on_file"
	PointOfServiceEntryModeMagneticStripe             PointOfServiceEntryMode = "magnetic_stripe"
	PointOfServiceEntryModeContactlessMagneticStripe  PointOfServiceEntryMode = "contactless_magnetic_stripe"
	PointOfServiceEntryModeIntegratedCircuitCardNoCvv PointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)