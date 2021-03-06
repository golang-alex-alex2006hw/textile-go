package migrations

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	_ "github.com/mutecomm/go-sqlcipher"
	"github.com/textileio/textile-go/crypto"
	libp2pc "gx/ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5/go-libp2p-crypto"
	"os"
	"path"
)

type thread struct {
	id   string
	name string
	sk   []byte
}

type threadRow struct {
	Name string `json:"name"`
	Sk   string `json:"sk"`
}

type photoRow struct {
	Id  string `json:"id"`
	Key string `json:"key"`
}

type Major005 struct{}

func (Major005) Up(repoPath string, pinCode string, testnet bool) error {
	var dbPath string
	if testnet {
		dbPath = path.Join(repoPath, "datastore", "testnet.db")
	} else {
		dbPath = path.Join(repoPath, "datastore", "mainnet.db")
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	if pinCode != "" {
		if _, err := db.Exec("pragma key='" + pinCode + "';"); err != nil {
			return err
		}
	}

	// collect thread secrets
	var threads []*thread
	var defaults []*thread
	rows, err := db.Query("select id, name, sk from threads;")
	if err != nil {
		return err
	}
	for rows.Next() {
		var id, name string
		var sk []byte
		if err := rows.Scan(&id, &name, &sk); err != nil {
			return err
		}
		threads = append(threads, &thread{id: id, name: name, sk: sk})
	}

	// write to file
	tfile, err := os.Create(path.Join(repoPath, "migration005_threads.ndjson"))
	if err != nil {
		return err
	}
	defer tfile.Close()
	for _, thrd := range threads {
		if thrd.name == "default" {
			defaults = append(defaults, thrd)
		}
		sk64 := base64.StdEncoding.EncodeToString(thrd.sk)
		row, err := json.Marshal(&threadRow{Name: thrd.name, Sk: sk64})
		if err != nil {
			return err
		}
		if _, err = tfile.Write(append(row[:], []byte("\n")[:]...)); err != nil {
			return err
		}
	}

	// collect default thread photo blocks
	var photos []*photoRow
	for _, thread := range defaults {
		sk, err := libp2pc.UnmarshalPrivateKey(thread.sk)
		if err != nil {
			return err
		}
		rows, err := db.Query("select dataId, dataKeyCipher from blocks where threadId='" + thread.id + "' and type=4;")
		if err != nil {
			return err
		}
		for rows.Next() {
			var id string
			var keyCipher []byte
			if err := rows.Scan(&id, &keyCipher); err != nil {
				return err
			}
			key, err := crypto.Decrypt(sk, keyCipher)
			if err != nil {
				return err
			}
			photos = append(photos, &photoRow{Id: id, Key: string(key)})
		}
	}

	// write to file
	bfile, err := os.Create(path.Join(repoPath, "migration005_default_photos.ndjson"))
	if err != nil {
		return err
	}
	defer bfile.Close()
	for _, photo := range photos {
		row, err := json.Marshal(photo)
		if err != nil {
			return err
		}
		if _, err = bfile.Write(append(row[:], []byte("\n")[:]...)); err != nil {
			return err
		}
	}

	// blast repo sans logs
	return blastRepo(repoPath)
}

func (Major005) Down(repoPath string, pinCode string, testnet bool) error {
	return ErrorCannotMigrateDown
}

func (Major005) Major() bool {
	return true
}
