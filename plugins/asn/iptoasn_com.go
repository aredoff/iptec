package asn

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/netip"
	"strconv"

	clog "github.com/aredoff/iptec/log"
)

const (
	url = "https://iptoasn.com/data/ip2asn-combined.tsv.gz"
)

func download() error {
	clog.Info("Start update asn database...")
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	gzipReader, err := gzip.NewReader(bytes.NewReader(body))
	if err != nil {
		return err
	}

	// output, e2 := io.ReadAll(gzipReader)
	// if e2 != nil {
	// 	fmt.Println(e2)
	// }

	// result := string(output)
	db := asnDb{}
	r := csv.NewReader(gzipReader)
	r.Comma = '\t'
	for {
		row, err := r.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		start, err := netip.ParseAddr(row[0])
		if err != nil {
			return fmt.Errorf("invalid start address #%s: %s", row[0], err)
		}
		end, err := netip.ParseAddr(row[1])
		if err != nil {
			return fmt.Errorf("invalid end address #%s: %s", row[1], err)
		}
		number, err := strconv.Atoi(row[2])
		if err != nil {
			return fmt.Errorf("invalid number address #%s: %s", row[2], err)
		}
		db = append(db, AS{
			Start:       start,
			End:         end,
			Number:      number,
			CountryCode: row[3],
			Description: row[4],
		})
		// db.Sort()

	}
	fmt.Println(db.Len())

	test_ip, _ := netip.ParseAddr("8.8.8.8")
	fmt.Println(db.Find(test_ip))

	// for _, v := range strings.Split(result, "\n") {
	// 	fmt.Println(v)
	// 	rows := strings.Split(v, "      ")
	// 	fmt.Println(rows)
	// }
	// fmt.Println(result)
	return nil
}
