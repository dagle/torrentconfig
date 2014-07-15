package torrentConfig

import (
	"github.com/steeve/libtorrent-go"
	"os"
	"encoding/json"
	"net/url"
	"io"
)

//This file is just structs so we can use it for different tools.
//These should be pure.

type Params struct {
	Uri string
	Path string
	Sequential bool
	Max_uploads int
	Max_connections int
	Upload_limit int
	Download_limit int
	Ratio int
	Priority int
	Queue bool
}

// I find this kinda broken
func p2p(param *Params) libtorrent.Add_torrent_params {
    tp := libtorrent.NewAdd_torrent_params()
	url, err := url.Parse(param.Uri)
	if err != nil {
		// do something
	}
    switch url.Scheme {
    case "file":
        ti := libtorrent.NewTorrent_info(url.Path)
        tp.SetTi(ti)
    case "link":
        tp.SetUrl(param.Uri)
    }
    tp.SetSave_path(param.Path)
    return tp
}

type Stats struct {
    Upload_rate int
    Download_rate int
    Total_download int // should be larger?
    Total_upload int
    Num_peers int
    Num_torrents int
    Paused bool
    Dht bool
}

type Config struct {
    Root string
    Max_upload_speed int
    Max_download_speed int
    Max_connections int
    Download_path string
    Encryption int
    Encryption_type int //
    No_sparsefile bool
    Portlower int
    Portupper int
    Minpeer int
    Maxpeer int
    Ratio int
    //portsRandom bool
    Checkhash bool
    Dht bool
    Peer_exchange bool
    // Will be both in the client and server for now
    defParams Params
	//resume bool
}

func readConfFile(path string) (*Config, error) {
    c := new(Config)
    f, err := os.Open(path + "/.t9fs/config.json")
    if err != nil {
        return nil, err
    }
    defer f.Close()
    c.defaultConfig(path)
    c.parseConfig(f)
    return c, nil
}

func (c *Config) parseConfig(f io.Reader) error {
    d := json.NewDecoder(f)
    e := d.Decode(c);
	return e
}

func (c *Config) defaultConfig(path string) {
    c.Root = "t9fs"
    c.Max_upload_speed = 0
    c.Max_download_speed = 0
    c.Max_connections = 0
    c.Download_path = path
    c.Encryption = 0
	c.Encryption_type = 0
	c.No_sparsefile = false
    c.Portlower = 6900
    c.Portupper = 6999
	c.Minpeer = 0
	c.Maxpeer = 0
	c.Ratio = 0
	c.Checkhash = false
	c.Dht = false
	c.Peer_exchange = false
	// defaultparams
}
