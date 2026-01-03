// Package users menyediakan service untuk autentikasi
package users

import (
	"backend-web-commision-kana/internal/repo"
	"backend-web-commision-kana/internal/utils/security"
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

func NewServices(repo *repo.Queries, db *pgxpool.Pool) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) Login(ctx context.Context, req ReqLogin) (*ResLogin, error) {
	user, err := s.repo.FindUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("username atau password salah")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	token, err := security.GenerateRandomToken(64)
	if err != nil {
		slog.Error("Gagal membuat token", "error", err)
		return nil, errors.New("terjadi kesalahan internal pada server")
	}

	savedToken, err := s.repo.SetToken(ctx, repo.SetTokenParams{
		UserID: user.ID,
		Token:  token,
	})
	if err != nil {
		slog.Error("Gagal menyimpan token ke database", "error", err)
		return nil, errors.New("terjadi kesalahan internal pada server")
	}

	return &ResLogin{
		Token:    savedToken,
		Username: user.Username,
	}, nil
}

func (s *svc) Logout(ctx context.Context, token string) (*ResLogout, error) {
	if err := s.repo.DeleteSessionByToken(ctx, token); err != nil {
		slog.Error("Gagal menghapus sesi", "error", err)
		return nil, errors.New("gagal logout, silahkan coba lagi")
	}

	return &ResLogout{
		Msg: "logout berhasil!",
	}, nil
}
