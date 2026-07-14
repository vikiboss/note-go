package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"example.com/go-learning/week12/internal/filesync"
)

func main() {
	source := flag.String("source", "", "source directory")
	target := flag.String("target", "", "target directory")
	dryRun := flag.Bool("dry-run", false, "print changes without writing")
	flag.Parse()
	if *source == "" || *target == "" {
		fmt.Fprintln(os.Stderr, "usage: filesync -source DIR -target DIR [-dry-run]")
		os.Exit(2)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	actions, err := filesync.Sync(ctx, *source, *target, *dryRun)
	if err != nil {
		logger.Error("sync failed", "error", err)
		os.Exit(1)
	}
	for _, action := range actions {
		fmt.Printf("%s\t%s\n", action.Kind, action.Path)
	}
	logger.Info("sync complete", "dry_run", *dryRun, "actions", len(actions))
}
