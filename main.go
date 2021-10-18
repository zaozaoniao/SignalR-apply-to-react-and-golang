	shub := services.AppHub{Cache: syncSer.Cache}
	sHubSrv, err := signalr.NewServer(context.TODO(),
		//	signalr.SimpleHubFactory(&shub),    // 简单的Hub，是作为原型，这是瞬态Hub， 每次都初始化一个新的hub struct
		//  signalr.HubFactory(services.NewHub),  // 具备灵活自定义生成Hub的能力， DI
		signalr.InsecureSkipVerify(true),
		signalr.UseHub(&shub), // 这是单例hub
		signalr.KeepAliveInterval(2*time.Second),
		signalr.Logger(kitlog.NewLogfmtLogger(os.Stderr), true))
	sHubSrv.MapHTTP(mux, "/realtime")



            //  push msg to web clients
            if cs := sHubServer.HubClients(); cs != nil {
	c := cs.All()
	if c != nil {
	     c.Send("receive", ts.Format("2006/01/02 15:04:05"))
	}
            }