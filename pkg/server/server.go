package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yykhomenko/hashes/pkg/config"
	"github.com/yykhomenko/hashes/pkg/store"
	"log"
	"strconv"
	"sync/atomic"
)

type Server struct {
	config  *config.Config
	store   *store.Store
	counter *counter
}

type counter struct {
	hashes  uint64
	msisdns uint64
}

type response struct {
	Value    string `json:"value,omitempty"`
	ErrorID  byte   `json:"errorID,omitempty"`
	ErrorMsg string `json:"errorMsg,omitempty"`
}

func New(c *config.Config, s *store.Store) *Server {
	return &Server{
		config:  c,
		store:   s,
		counter: &counter{},
	}
}

func (s *Server) Start() {
	log.Println("http-server listening...")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/metrics", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString(fmt.Sprintf(
			"hashes_total %d\nmsisdns_total %d\n",
			s.counter.hashes,
			s.counter.msisdns,
		))
	})

	app.Get("/hashes/:msisdn", func(c *fiber.Ctx) error {
		atomic.AddUint64(&s.counter.hashes, 1)
		msisdn := c.Params("msisdn")

		if !validateMsisdnLen(msisdn, s.config.MsisdnLenMin, s.config.MsisdnLenMax) {
			return c.Status(fiber.StatusBadRequest).JSON(response{ErrorID: 2, ErrorMsg: "Not supported MSISDN format: " + msisdn})
		}

		if cc, ok := validateCC(msisdn, s.config.CC); !ok {
			return c.Status(fiber.StatusBadRequest).JSON(response{ErrorID: 3, ErrorMsg: "Not supported CC: " + cc})
		}

		if ndc, ok := validateNDC(msisdn, s.config.NDCS); !ok {
			return c.Status(fiber.StatusBadRequest).JSON(response{ErrorID: 4, ErrorMsg: "Not supported NDC: " + ndc})
		}

		return c.Status(fiber.StatusOK).JSON(response{Value: s.store.Hash(msisdn[3:])})
	})

	app.Get("/msisdns/:hash", func(c *fiber.Ctx) error {
		atomic.AddUint64(&s.counter.msisdns, 1)
		hash := c.Params("hash")
		msisdn, exists := s.store.Msisdn(hash)

		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(response{ErrorID: 1, ErrorMsg: "Not found"})
		}

		return c.Status(fiber.StatusOK).JSON(response{Value: s.config.CC + msisdn})
	})

	log.Fatal(app.Listen(s.config.Addr))
}

func validateMsisdnLen(msisdn string, min, max int) bool {
	l := len(msisdn)
	return min <= l && l <= max
}

func validateCC(msisdn, confCC string) (string, bool) {
	cc := msisdn[:3]
	if cc == confCC {
		return cc, true
	} else {
		return cc, false
	}
}

func validateNDC(msisdn string, ndcs []int) (string, bool) {
	ndcStr := msisdn[3:5]

	ndc, err := strconv.Atoi(ndcStr)
	if err != nil {
		log.Println(err)
	}

	for _, n := range ndcs {
		if ndc == n {
			return ndcStr, true
		}
	}

	return ndcStr, false
}
