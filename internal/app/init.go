package app

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"html/template"

	"github.com/TwiN/go-color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"

	"github.com/fadhilaf/s-tech/common/env"

	"github.com/fadhilaf/s-tech/internal/middleware"

	"github.com/fadhilaf/s-tech/common/validations"
)

type App struct {
	Config   env.Config
	delivery deliveries
	usecase  usecases
	db       *sql.DB
}

//go:embed templates/*.gohtml
var templateFiles embed.FS


//go:embed assets/*
var assetFiles embed.FS

func (app *App) createHttpHandlers() *gin.Engine {
	router := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowHeaders = append(corsCfg.AllowHeaders, "Accept")

	corsCfg.AllowAllOrigins = true

	router.Use(cors.New(corsCfg))

	router.Use(middleware.SessionMiddleware(app.db))
	router.Use(middleware.SaveAndLoadSessionMiddleware())

	apiRouterGroup := router.Group("/api")
	app.apiHandler(apiRouterGroup)

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB

  // Membaca file dari variabel embed, bukan dari harddisk
	templ := template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))
	router.SetHTMLTemplate(templ)
	// router.LoadHTMLGlob(app.Config.AppStatic + "/template/*.gohtml")

	// untuk asset: logo, css, js
	router.StaticFS("/assets", http.FS(assetFiles))

	// untuk lokasi simpan file yang diupload (janlup kasi opsi if, kalau semisal kita simpan static file di cloud lain bukan di server yang sama)
	if !app.Config.IsStaticCloud {
		router.Static("/public", app.Config.AppStaticPath)
	}

	viewRouterGroup := router.Group("/")
	app.viewHandler(viewRouterGroup)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.gohtml", gin.H{})
	})

	// >> mesin untuk menampilkan semua routes yang ada
	routes := router.Routes()
	if gin.Mode() == gin.DebugMode {
		fmt.Println()
		for _, v := range routes {
			path := color.InBold(v.Path)
			method := color.InYellow(fmt.Sprintf("%-6s", v.Method))
			fmt.Println(method, path)
		}
		fmt.Println()
	}
	// << mesin untuk menampilkan semua routes yang ada selesai

	return router
}

func (app *App) StartServer() {
	if app.Config.Env == env.EnvProd {
		fmt.Println(
			color.Ize(color.Yellow, color.InBold("\nAPP RUN IN PRODUCTION MODE\n")),
		)
	} else {
		fmt.Println(
			color.Ize(color.Red, color.InBold("\nAPP RUN IN DEVELOPMENT MODE\n")),
		)
	}

	osSignalChan := make(chan os.Signal, 1)
	// signal.Notify akan memasukkan nilai ke channel osSignalChan ketika ada signal interrupt yang masuk (SIGINT (Ctrl+C) atau SIGTERM (kill))
	// dibuat ini, biar program tidak langsung berhenti ketika function main selesai dijalankan. (nunggu di line 98 sampai ada signal interrupt)
	signal.Notify(osSignalChan, syscall.SIGINT, syscall.SIGTERM)

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//validations itu settingan custom, utk validasi email, dll buatan kito dewek di folder validation
		validations.InitValidations(validator)
	}
	router := app.createHttpHandlers()
	address := fmt.Sprintf("%s:%s", app.Config.AppHost, app.Config.AppPort)
	log.Printf("Server listening on %v\n", address)

	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalf("Cannot start server %v\n", err)
		}
	}()

	//ngeluarin nilai channel, kalo channel masih kosong, jadi nunggu disini sampe ada signal interrupt
	<-osSignalChan
	err := srv.Close()
	if err != nil {
		log.Fatalf("cannot shutdown server %v", err)
	}
	fmt.Println()
	log.Println("Server exiting")
}

func New(config env.Config, db *sql.DB) App {

	app := App{}
	app.Config = config

	app.db = db
	app.initUsecase()
	app.initDelivery()

	return app
}
