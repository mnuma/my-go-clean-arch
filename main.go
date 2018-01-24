package main

import (
  "github.com/labstack/echo"
  "github.com/mnuma/my-go-clean-arch/author/repository"

  articleRepository "github.com/mnuma/my-go-clean-arch/article/repository"
  httpDeliver "github.com/mnuma/my-go-clean-arch/article/delivery/http"

  "database/sql"
  "fmt"
  "net/url"
  "github.com/k0kubun/pp"
  "github.com/mnuma/my-go-clean-arch/article/usecase"
)

func main() {

  connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "user", "pass", "localhost", "3306", "sample")
  val := url.Values{}
  val.Add("parseTime", "1")
  val.Add("loc", "Asia/Tokyo")
  dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

  dbConn, err := sql.Open(`mysql`, dsn)
  if err != nil {
    pp.Println(err)
  }
  defer dbConn.Close()

  e := echo.New()
  authorRepository := repository.AuthorRepository(dbConn)
  ar := articleRepository.NewMysqlArticleRepository(dbConn)
  au := usecase.NewArticleUsecase(ar, authorRepository)

  httpDeliver.NewArticleHttpHandler(e, au)

  e.Start(":1323")
}
