package dokeshi

// configuration of the generator
type Config struct {
	Generator struct {
		Repository string
		Temp       string
		Branch     string
		Dest       string
		RSS        bool
	}

	Blog struct {
		URL       string
		Lang      string
		Descp     string
		Datefmt   string
		Title     string
		Author    string
		Frontpage int
		Statics   struct {
			Files []struct {
				Src  string
				Dest string
			}
			Templates []struct {
				Src  string
				Dest string
			}
		}
	}
}
