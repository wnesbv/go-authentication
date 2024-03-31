package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "github.com/lib/pq"

	"go_authentication/chat"
	"go_authentication/sqlcsv"
	"go_authentication/search"
	"go_authentication/sitemap"
	"go_authentication/profile"
	"go_authentication/article"
	"go_authentication/owner_ssc"
	"go_authentication/subscription"
)


func main() {

	http.HandleFunc("/", profile.Home)
	http.HandleFunc("/alluser", profile.Alluser)
	http.HandleFunc("/signup", profile.Signup)
	http.HandleFunc("/login", profile.Login)
	http.HandleFunc("/auth", profile.AuthToken)
	http.HandleFunc("/update-name", profile.UpName)
	http.HandleFunc("/update-password", profile.UpPass)
	http.HandleFunc("/send-email", profile.EmailSend)
	http.HandleFunc("/verification", profile.VerifyEmail)
	http.HandleFunc("/delete-user", profile.DelUs)
	
	// art..
	http.HandleFunc("/article", article.HomeArticle)
	http.HandleFunc("/allarticle", article.Allarticle)
	http.HandleFunc("/detail-art", article.DetArt)
	http.HandleFunc("/author-id-article", article.UsAllArt)
	http.HandleFunc("/creativity", article.Creativity)
	http.HandleFunc("/update-art", article.UpArt)
	http.HandleFunc("/delete-art", article.DelArt)
	http.HandleFunc("/img-art", article.ImgArt)
	http.HandleFunc("/del-img-art", article.DelImgArt)
	http.HandleFunc("/csv-imp-art", sqlcsv.CsvImpArt)
	http.HandleFunc("/csv-exp-art", sqlcsv.ExpCsvArt)

	// chat..
	http.HandleFunc("/chat", chat.HomeChat)
	http.HandleFunc("/all-group", chat.GrAll)
	http.HandleFunc("/owner-group", chat.GrOwr)
	http.HandleFunc("/detail-group", chat.DetGr)
	
	http.HandleFunc("/creat-group", chat.Creativity)
	http.HandleFunc("/update-group", chat.UpGr)

	http.HandleFunc("/user", chat.UsChat)
	http.HandleFunc("/user/us", chat.UsMsg)
	http.HandleFunc("/group", chat.GrChat)
	http.HandleFunc("/group/rs", chat.GrMsg)

	// owner subscription..
	http.HandleFunc("/all-ssc", owner_ssc.OwrAllSsc)
	http.HandleFunc("/detail-ssc", owner_ssc.DetOwrSsc)
	http.HandleFunc("/del-ssc", owner_ssc.OwrDelSsc)
	
	http.HandleFunc("/adduser-ssc", owner_ssc.AddSscUs)
	http.HandleFunc("/addroom-ssc", owner_ssc.AddSscGr)
	http.HandleFunc("/up-owner-ssc", owner_ssc.OwrUpSsc)

	// subscription..
	http.HandleFunc("/subscription", subscription.AllSsc)
	http.HandleFunc("/user-ssc", subscription.ToUpUsSsc)
	http.HandleFunc("/group-ssc", subscription.ToUpGroupSsc)

	http.HandleFunc("/all-touser-ssc", subscription.ToUsAllSsc)
	http.HandleFunc("/all-toroom-ssc", subscription.ToGroupAllSsc)

	// search..
	http.HandleFunc("/search", search.SearchHandler)

	// sitemap.xml
	http.HandleFunc("/sitemap.xml", sitemap.SitemapHandler)

	http.HandleFunc("/onauth", article.OnAuth)

	// static..
	dir := http.Dir("./sfl/static")
	fls := http.FileServer(dir)
	http.Handle("/static/", http.StripPrefix("/static", fls))
	

	fmt.Println(" main goroutine..", runtime.NumGoroutine())

	fmt.Println("Server running port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


