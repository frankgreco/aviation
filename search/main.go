//package main
package search

// import (
// 	"bufio"
// 	"database/sql"
// 	"fmt"
// 	"os"

// 	"github.com/spf13/cobra"
// 	//"github.com/frankgreco/aviation/utils/db"
// )

// var (
// 	config = struct {
// 		cfg Config
// 		db  string
// 	}{
// 		cfg: Config{
// 			Filters: Filters{
// 				TailNumber: new(string),
// 			},
// 		},
// 	}

// 	rootCmd = &cobra.Command{
// 		Use:  "search",
// 		RunE: search,
// 	}
// )

// func search(cmd *cobra.Command, args []string) error {
// 	dbase, err := sql.Open("pgx", config.db)
// 	if err != nil {
// 		return err
// 	}
// 	defer dbase.Close()
// 	dbase.SetMaxOpenConns(1)

// 	dbOption := Database(dbase)
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Print("> ")
// 		text, err := reader.ReadString('\n')
// 		if err != nil {
// 			return err
// 		}

// 		if text == "\\q\n" {
// 			return nil
// 		}

// 		results, err := New(append(Parse(text), dbOption)...)
// 		if err != nil {
// 			return err
// 		}

// 		for _, result := range results.Results {
// 			fmt.Printf("make=%s		model=%s	tail_number=%s\n", result.Aircraft.Manufacturer, result.Aircraft.Model, result.Registration.Id)
// 		}
// 	}

// 	return nil
// }

// func main() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
// 		os.Exit(1)
// 	}
// }

// func init() {
// 	rootCmd.Flags().StringVar(&config.db, "db_url", "", "database connection string")
// 	rootCmd.MarkFlagRequired("db_url")
// }
