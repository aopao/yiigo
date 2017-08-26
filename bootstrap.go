package yiigo

// Bootstrap start components
func Bootstrap(mysql bool, mongo bool, redis bool) error {
	loadEnv("env.ini")
	initLogger("log.xml")

	if mysql {
		if err := initMySQL(); err != nil {
			return err
		}
	}

	if mongo {
		if err := initMongo(); err != nil {
			return err
		}
	}

	if redis {
		initRedis()
	}

	return nil
}
