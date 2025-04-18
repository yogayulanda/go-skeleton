package grpcgateway

// func RunServerGrpcGW(ctx context.Context, container *di.Container) error {
// 	mux := runtime.NewServeMux()

// 	// Middleware wrapper
// 	handler := middleware.ChainMiddleware(
// 		middleware.HTTPRequestLogger(container.Log),
// 		middleware.HTTPPanicRecovery(container.Log),
// 		// Add other middleware here
// 	)(mux)

// 	grpcAddr := fmt.Sprintf("localhost:%s", container.Config.GRPC_PORT)

// 	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

// 	if err := v1pb.RegisterTransactionHistoryServiceHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts); err != nil {
// 		return fmt.Errorf("failed to register TrxHistory handler: %w", err)
// 	}
// 	if err := v1pb.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts); err != nil {
// 		return fmt.Errorf("failed to register Health handler: %w", err)
// 	}

// 	// @auto:inject:handler

// 	srv := &http.Server{
// 		Addr:         fmt.Sprintf(":%s", container.Config.HTTP_PORT),
// 		Handler:      handler,
// 		ReadTimeout:  15 * time.Second,
// 		WriteTimeout: 15 * time.Second,
// 		IdleTimeout:  60 * time.Second,
// 	}

// 	// Optional TLS setup if enabled
// 	if container.Config.ENABLE_TLS {
// 		srv.TLSConfig = &tls.Config{
// 			MinVersion: tls.VersionTLS12,
// 		}
// 	}
// 	utils.LogAvailableEndpoints()
// 	// Log service started AFTER all init
// 	container.Log.Info("✅ go-skeleton service started successfully",
// 		zap.String("version", "v1.0.0"),
// 		zap.String("time", time.Now().Format(time.RFC3339)),
// 	)

// 	// Graceful shutdown
// 	idleConnsClosed := make(chan struct{})
// 	go func() {
// 		stop := make(chan os.Signal, 1)
// 		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
// 		<-stop

// 		container.Log.Info("🛑 Shutting down HTTP server...")

// 		ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

// 		if err := srv.Shutdown(ctxTimeout); err != nil {
// 			container.Log.Error("❌ Failed to shutdown HTTP server gracefully", zap.Error(err))
// 		}
// 		close(idleConnsClosed)
// 	}()

// 	// Run server
// 	var err error
// 	if container.Config.ENABLE_TLS {
// 		err = srv.ListenAndServeTLS(container.Config.TLS_CERT_PATH, container.Config.TLS_KEY_PATH)
// 	} else {
// 		err = srv.ListenAndServe()
// 	}

// 	if err != http.ErrServerClosed {
// 		return fmt.Errorf("HTTP server failed: %w", err)
// 	}

// 	<-idleConnsClosed
// 	return nil
// }
