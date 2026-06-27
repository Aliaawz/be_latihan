package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"http://localhost:3000",
	"http://127.0.0.1:3000",
	"belatihan-production-b816.up.railway.app",
	"https://belatihan-production-66bb.up.railway.app",
	"https://my-fe-orpin.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
