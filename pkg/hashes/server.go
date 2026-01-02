package hashes

import (
	"fmt"
	"log"
	"strconv"
	"sync/atomic"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config  *Config
	store   *Store
	counter struct {
		hashes  uint64
		msisdns uint64
	}
}

type response struct {
	Value    string `json:"value,omitempty"`
	ErrorID  byte   `json:"error_id,omitempty"`
	ErrorMsg string `json:"error_msg,omitempty"`
}

func NewServer(cfg *Config, store *Store) *Server {
	return &Server{
		config: cfg,
		store:  store,
	}
}

func (s *Server) Start() error {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Get("/", s.handleRoot())
	app.Get("/metrics", s.handleMetrics())
	app.Get("/hashes/:msisdn", s.handleGetHash())
	app.Get("/msisdns/:hash", s.handleGetMsisdn())

	log.Printf("🌐 HTTP server listening on %s", s.config.Addr)
	if err := app.Listen(s.config.Addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
		return err
	}
	return nil
}

// --- Handlers ---

func (s *Server) handleRoot() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	}
}

func (s *Server) handleMetrics() fiber.Handler {
	return func(c *fiber.Ctx) error {
		hashes := atomic.LoadUint64(&s.counter.hashes)
		msisdns := atomic.LoadUint64(&s.counter.msisdns)

		return c.Status(fiber.StatusOK).SendString(fmt.Sprintf(
			"hashes_total %d\nmsisdns_total %d\n", hashes, msisdns,
		))
	}
}

func (s *Server) handleGetMsisdn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		atomic.AddUint64(&s.counter.msisdns, 1)
		hash := c.Params("hash")

		msisdn, ok := s.store.Msisdn(hash)
		if !ok {
			return c.Status(fiber.StatusNotFound).JSON(response{
				ErrorID:  1,
				ErrorMsg: fmt.Sprintf("not found: %s", hash),
			})
		}

		return c.JSON(response{Value: s.config.CC + msisdn})
	}
}

func (s *Server) handleGetHash() fiber.Handler {
	return func(c *fiber.Ctx) error {
		atomic.AddUint64(&s.counter.hashes, 1)
		msisdn := c.Params("msisdn")

		// --- Валідації ---
		if !validateMsisdnLen(msisdn, s.config.MsisdnLenMin, s.config.MsisdnLenMax) {
			return badRequest(c, 2, "not supported MSISDN length", msisdn)
		}
		if cc, ok := validateCC(msisdn, s.config.CC); !ok {
			return badRequest(c, 3, "not supported country code", cc)
		}
		if ndc, ok := validateNDC(msisdn, s.config.NDCS); !ok {
			return badRequest(c, 4, "not supported NDC", ndc)
		}

		// --- Формування хешу ---
		hash := s.store.Hash(msisdn[3:])
		return c.JSON(response{Value: hash})
	}
}

// --- Helpers ---

func badRequest(c *fiber.Ctx, id byte, msg, value string) error {
	return c.Status(fiber.StatusBadRequest).JSON(response{
		ErrorID:  id,
		ErrorMsg: fmt.Sprintf("%s: %s", msg, value),
	})
}

// --- Validation ---

func validateMsisdnLen(msisdn string, min, max int) bool {
	l := len(msisdn)
	return l >= min && l <= max
}

func validateCC(msisdn, expectedCC string) (string, bool) {
	if len(msisdn) < 3 {
		return "", false
	}
	cc := msisdn[:3]
	return cc, cc == expectedCC
}

func validateNDC(msisdn string, ndcs []int) (string, bool) {
	if len(msisdn) < 5 {
		return "", false
	}
	ndcStr := msisdn[3:5]
	ndc, err := strconv.Atoi(ndcStr)
	if err != nil {
		log.Printf("invalid NDC: %s (%v)", ndcStr, err)
		return ndcStr, false
	}

	for _, n := range ndcs {
		if n == ndc {
			return ndcStr, true
		}
	}
	return ndcStr, false
}
