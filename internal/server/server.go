package server

import (
	"context"
	"errors"
	"net"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/charmbracelet/wish/recover"
	"golang.org/x/sync/errgroup"

	"github.com/davidsbond/kingdom/internal/game"
	"github.com/davidsbond/kingdom/internal/game/scene"
	"github.com/davidsbond/kingdom/internal/game/window"
)

type (
	Config struct {
		Host    string
		Port    int
		KeyPath string
		Logger  *log.Logger
	}
)

func Run(ctx context.Context, config Config) error {
	logger := config.Logger
	log.SetDefault(logger)

	state := game.NewState()

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(config.Host, strconv.Itoa(config.Port))),
		wish.WithHostKeyPath(config.KeyPath),
		wish.WithMiddleware(
			recover.MiddlewareWithLogger(logger,
				bubbletea.Middleware(func(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
					pty, _, _ := sess.Pty()

					w := window.New(pty.Window.Width, pty.Window.Height)
					player := state.Player(sess.User())

					sctx := scene.Context{
						Window: w,
						Player: player,
						State:  state,
					}

					return scene.Splash(sctx), []tea.ProgramOption{
						tea.WithAltScreen(),
						tea.WithContext(ctx),
						tea.WithFPS(60),
					}
				}),
			),
			activeterm.Middleware(),
			logging.StructuredMiddlewareWithLogger(logger, log.DebugLevel),
		),
	)
	if err != nil {
		return err
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		logger.
			With("host", config.Host, "port", config.Port).
			Info("starting server")

		return s.ListenAndServe()
	})

	group.Go(func() error {
		<-ctx.Done()

		sCtx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()

		logger.Warn("server shutting down")
		return s.Shutdown(sCtx)
	})

	err = group.Wait()
	switch {
	case errors.Is(err, context.Canceled), errors.Is(err, ssh.ErrServerClosed):
		return nil
	case err != nil:
		return err
	default:
		return nil
	}
}
