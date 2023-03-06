package main

import (
    "context"
    "database/sql"
    "donc83/go-data-access/internal"
    "fmt"
    "github.com/go-sql-driver/mysql"
    "github.com/jackc/pgx/v5/pgxpool"
    "log"
    "os"
)

var db *sql.DB

func main() {
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }

    var err error

    dbPool, err := pgxpool.New(context.Background(),
        fmt.Sprintf("postgresql://%s:%s@127.0.0.1:5432/recordings?pool_max_conns=10",
            os.Getenv("DBUSER"), os.Getenv("DBPASS")))
    if err != nil {
        log.Fatal(err)
    }
    defer dbPool.Close()

    err = dbPool.Ping(context.Background())
    if err != nil {
        log.Fatal("Failed to ping postgres ", err)
    } else {
        log.Println("Connected to postgres!")
    }

    db, err = sql.Open("mysql", cfg.FormatDSN())
    defer db.Close()
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

    albums, err := albumsByArtist("John Coltrane")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Albums found: %v\n", albums)

    album, err := albumById(2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Album found %v\n", album)

    albID, err := addAlbum(internal.Album{
        Title:  "The Modern Sound of Betty Carter",
        Artist: "Betty Carter",
        Price:  49.99,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID of added album: %v\n", albID)

}

func albumsByArtist(name string) ([]internal.Album, error) {
    var albums []internal.Album

    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }

    for rows.Next() {
        var alb internal.Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %w", name, err)
        }
        albums = append(albums, alb)
    }
    return albums, nil
}

func albumById(id int64) (internal.Album, error) {
    var alb internal.Album

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumById %d: %w", id, err)
    }

    return alb, nil
}

func addAlbum(alb internal.Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %w", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %w", err)
    }
    return id, nil
}
