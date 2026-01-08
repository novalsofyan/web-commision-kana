// Package scheduler untuk scheduler
package scheduler

import (
	"backend-web-commision-kana/internal/repo"
	"context"
	"log/slog"
	"time"
)

// SessionCleanupScheduler menangani cleanup periodik untuk session yang expired
type SessionCleanupScheduler struct {
	queries  *repo.Queries
	interval time.Duration
	stopChan chan struct{}
}

// NewSessionCleanupScheduler membuat session cleanup scheduler baru
// interval: seberapa sering cleanup dijalankan (contoh: 24 jam)
func NewSessionCleanupScheduler(queries *repo.Queries, interval time.Duration) *SessionCleanupScheduler {
	return &SessionCleanupScheduler{
		queries:  queries,
		interval: interval,
		stopChan: make(chan struct{}),
	}
}

// Start memulai cleanup scheduler di goroutine
func (s *SessionCleanupScheduler) Start() {
	go s.run()
	slog.Info("Session cleanup scheduler dimulai", "interval", s.interval)
}

// Stop menghentikan cleanup scheduler secara graceful
func (s *SessionCleanupScheduler) Stop() {
	close(s.stopChan)
	slog.Info("Session cleanup scheduler dihentikan")
}

// run adalah loop utama yang menjalankan cleanup secara periodik
func (s *SessionCleanupScheduler) run() {
	// Jalankan cleanup langsung saat startup
	s.cleanup()

	// Buat ticker untuk eksekusi periodik
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.cleanup()
		case <-s.stopChan:
			return
		}
	}
}

// cleanup menjalankan operasi cleanup yang sebenarnya
func (s *SessionCleanupScheduler) cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	slog.Info("Menjalankan cleanup expired sessions...")

	err := s.queries.CleanupExpiredSessions(ctx)
	if err != nil {
		slog.Error("Gagal cleanup expired sessions", "error", err)
		return
	}

	slog.Info("Cleanup expired sessions berhasil")
}
